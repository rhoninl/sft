package describe

import (
	"fmt"
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

		fmt.Println(logger.WithColor(logger.Blue, "************Device*************"))
		fmt.Println("Name: ", device.EdgeDevice.Name)
		var protocol string = string(*device.EdgeDevice.Spec.Protocol)
		if gatewaySettings != nil {
			protocol += " -> " + *device.EdgeDevice.Spec.GatewaySettings.Protocol
		}

		fmt.Println("Protocol: ", protocol)
		fmt.Println("Address: ", *device.EdgeDevice.Spec.Address)

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
		fmt.Println(logger.WithColor(logger.Purple, "**************END**************"))
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
