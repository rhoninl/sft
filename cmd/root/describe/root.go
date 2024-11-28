package describe

import (
	"encoding/json"
	"fmt"
	"net"
	"slices"
	"strings"

	"github.com/rhoninl/shifucli/cmd/k8s"
	"github.com/rhoninl/shifucli/cmd/utils/logger"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	DescribeCmd.CompletionOptions.DisableDefaultCmd = false
}

var DescribeCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"desc"},
	Short:   "Show detailed information of an edgedevice in the current Kubernetes cluster",
	Long:    `Show detailed information of an edgedevice in the current Kubernetes cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: Device name is required")
			return
		}

		deviceName := args[0]
		device, err := k8s.GetAllByDeviceName(deviceName)
		if err != nil {
			fmt.Printf("Error retrieving device: %v\n", err)
			return
		}

		printDeviceDetails(device)
	},
	ValidArgs: k8s.GetValidDeviceNames(),
}

func printDeviceDetails(device *k8s.Device) {
	fmt.Println("Name:", device.EdgeDevice.Name)
	fmt.Println("Status:", logger.StatusWithColor(string(*device.EdgeDevice.Status.EdgeDevicePhase)))
	fmt.Println("Address:", getRealDeviceAddress(*device))
	fmt.Println("Protocol:", getProtocol(device))

	connectionSettings, err := marshalSettings(device.EdgeDevice.Spec.ProtocolSettings)
	if err != nil {
		fmt.Printf("Error marshalling protocol settings: %v\n", err)
		return
	}

	gatewaySettings, err := marshalSettings(device.EdgeDevice.Spec.GatewaySettings)
	if err != nil {
		fmt.Printf("Error marshalling gateway settings: %v\n", err)
		return
	}

	alignMultiLineStrings(string(connectionSettings), string(gatewaySettings))
	printContainerSettings(device)
	printAPIInfo(device)
}

func getProtocol(device *k8s.Device) string {
	protocol := string(*device.EdgeDevice.Spec.Protocol)
	if device.EdgeDevice.Spec.GatewaySettings != nil {
		protocol += " -> " + *device.EdgeDevice.Spec.GatewaySettings.Protocol
	}
	return protocol
}

func marshalSettings(settings interface{}) ([]byte, error) {
	var tmpSettings = map[string]interface{}{}

	data, err := json.Marshal(settings)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &tmpSettings)
	if err != nil {
		return nil, err
	}

	if settings == nil {
		return nil, nil
	}
	return yaml.Marshal(tmpSettings)
}

func printContainerSettings(device *k8s.Device) {
	fmt.Println("===========Container===========")
	var containerSettings []string
	for _, container := range device.Deployment.Spec.Template.Spec.Containers {
		containerSetting := fmt.Sprintf("Name: %s\nImage: %s\n", container.Name, container.Image)
		for _, env := range container.Env {
			containerSetting += fmt.Sprintf("\tEnv: %s: %s\n", env.Name, env.Value)
		}
		containerSettings = append(containerSettings, containerSetting)
	}
	alignMultiLineStrings(containerSettings...)
}

func printAPIInfo(device *k8s.Device) {
	fmt.Println("==============API==============")
	instructions := strings.Split(device.ConfigMap.Data["instructions"], "\n")
	for i, line := range instructions {
		if i >= 10 {
			fmt.Println("......")
			break
		}
		fmt.Println(line)
	}
	fmt.Print(device.ConfigMap.Data["telemetries"])
}

// Align multiple strings and print them as blocks
func alignMultiLineStrings(blocks ...string) {
	// Trim spaces and split each block into lines
	lines := make([][]string, len(blocks))
	maxWidths := make([]int, len(blocks)) // To store the maximum width of each block

	for i := range blocks {
		blocks[i] = strings.ReplaceAll(blocks[i], "\t", "  ")
		block := blocks[i]
		// Split into lines and trim spaces
		blockLines := strings.Split(strings.TrimSpace(block), "\n")
		lines[i] = blockLines

		// Find the maximum width of lines in this block
		for _, line := range blockLines {
			if len(line) > maxWidths[i] {
				maxWidths[i] = len(line)
			}
		}
	}

	// Determine the maximum number of lines across all blocks
	maxLines := 0
	for _, blockLines := range lines {
		if len(blockLines) > maxLines {
			maxLines = len(blockLines)
		}
	}

	// Align and print each line
	for i := 0; i < maxLines; i++ {
		var row []string

		for j, blockLines := range lines {
			line := ""
			if i < len(blockLines) {
				line = blockLines[i]
			}

			// Align the line to the max width of its block
			row = append(row, fmt.Sprintf("%-*s", maxWidths[j], line))
		}

		// Join the aligned lines with separators and print
		fmt.Println(strings.Join(row, " | "))
	}
}

var internalHost = []string{
	"localhost",
	"0.0.0.0",
	"",
}

var addressHallmark = []string{
	"IP",
	"HOST",
	"Address",
}

func containsAny(arr []string, str string) bool {
	for _, v := range arr {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}

var passiveProtocol = []string{
	"LwM2M",
}

func getRealDeviceAddress(device k8s.Device) string {
	address := device.EdgeDevice.Spec.Address

	if slices.Contains(passiveProtocol, string(*device.EdgeDevice.Spec.Protocol)) {
		return "(passive device)" + passiveDeviceServerAddress(&device)
	}

	if address == nil {
		return "N/A"
	} else if *address == "" {
		return "N/A"
	}

	host, port, err := net.SplitHostPort(*address)
	if err != nil {
		return *address
	}

	if !containsAny(internalHost, host) {
		return *address
	}

	envs := make(map[string]string)

	// Find the container that exposes the port
	for _, container := range device.Deployment.Spec.Template.Spec.Containers {
		var mark bool
		if container.Ports != nil {
			for _, containerPort := range container.Ports {
				if string(containerPort.ContainerPort) == port {
					clear(envs)
					mark = true
				}
			}
		}
		for _, env := range container.Env {
			envs[env.Name] = env.Value
		}

		if mark {
			break
		}
	}

	for key := range envs {
		if !containsAny(addressHallmark, key) {
			delete(envs, key)
		}
	}

	if len(envs) == 0 {
		return "(passive device)" + passiveDeviceServerAddress(&device)
	}

	for _, value := range envs {
		return value
	}

	return "N/A"
}

func passiveDeviceServerAddress(device *k8s.Device) string {
	for _, service := range device.Services.Items {
		if service.Spec.Type != "NodePort" {
			continue
		}

		for _, port := range service.Spec.Ports {
			if port.NodePort == 0 {
				continue
			} else if port.TargetPort.IntVal == 8080 {
				continue
			}

			return fmt.Sprintf("localhost:%d", port.NodePort)
		}
	}

	service := device.Services.Items[0]

	return fmt.Sprintf("%s.%s.svc.cluster.local:%d", service.Name, service.Namespace, service.Spec.Ports[0].TargetPort.IntVal)
}
