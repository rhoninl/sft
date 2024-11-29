package describe

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/address"
	"github.com/rhoninl/sft/pkg/utils/logger"
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
	fmt.Println("Address:", address.GetRealDeviceAddress(*device))
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

	if len(tmpSettings) == 0 {
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
		if blocks[i] == "" {
			continue
		}
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
