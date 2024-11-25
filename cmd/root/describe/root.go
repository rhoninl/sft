package describe

import (
	"fmt"
	"net"
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
	Short:   "show edgedevice detail info in current kubernetes cluster",
	Long:    `show edgedevice detail info in current kubernetes cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		deviceName := args[0]
		device, err := k8s.GetAllByDeviceName(deviceName)
		if err != nil {
			fmt.Println(err)
			return
		}

		var connectionSettings []byte
		if device.EdgeDevice.Spec.ProtocolSettings != nil {
			connectionSettings, err = yaml.Marshal(device.EdgeDevice.Spec.ProtocolSettings)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		var gatewaySettings []byte
		if device.EdgeDevice.Spec.GatewaySettings != nil {
			gatewaySettings, err = yaml.Marshal(device.EdgeDevice.Spec.GatewaySettings)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		fmt.Println("Name: ", device.EdgeDevice.Name)
		var protocol string = string(*device.EdgeDevice.Spec.Protocol)
		if gatewaySettings != nil {
			protocol += " -> " + *device.EdgeDevice.Spec.GatewaySettings.Protocol
		}

		fmt.Println("Status: ", logger.StatusWithColor(string(*device.EdgeDevice.Status.EdgeDevicePhase)))
		fmt.Println("Address: ", getRealDeviceAddress(*device))
		fmt.Println("Protocol: ", protocol)

		var gatewayInfo, deviceInfo string
		if len(connectionSettings) != 0 {
			deviceInfo += string(connectionSettings)
		}
		if len(gatewaySettings) != 0 {
			gatewayInfo += string(gatewaySettings)
		}

		alignMultiLineStrings(deviceInfo, gatewayInfo)
		fmt.Println("===========Container===========")
		var contianerSettings []string
		for _, container := range device.Deployment.Spec.Template.Spec.Containers {
			var containerSetting string
			containerSetting += fmt.Sprintf("Name: %s\n", container.Name)
			containerSetting += fmt.Sprintf("Image: %s\n", container.Image)
			for _, env := range container.Env {
				containerSetting += fmt.Sprintf("\tEnv: %s: %s\n", env.Name, env.Value)
			}

			contianerSettings = append(contianerSettings, containerSetting)
		}

		alignMultiLineStrings(contianerSettings...)
		fmt.Println("==============API==============")
		fmt.Print(device.ConfigMap.Data["instructions"])
		fmt.Print(device.ConfigMap.Data["telemetries"])
	},
	ValidArgs: func() []string {
		edgedevices, err := k8s.GetEdgedevices()
		if err != nil {
			return nil
		}

		var deviceNames []string
		for _, edgedevice := range edgedevices {
			deviceNames = append(deviceNames, edgedevice.Name)
		}

		return deviceNames
	}(),
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

func getRealDeviceAddress(device k8s.Device) string {
	address := device.EdgeDevice.Spec.Address

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
