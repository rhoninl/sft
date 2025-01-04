// Package edgedevices provides functionality to manage edge devices in a Kubernetes cluster.
package devices

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/address"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/spf13/cobra"
)

var (
	// protocol is the flag value for filtering edge devices by protocol
	protocol string
	// status is the flag value for filtering edge devices by status
	status string
	// headers defines the column headers for the edge device list table
	headers = []string{"\rName", "Protocol", "Address", "Status", "Age"}
)

type DeviceInfo struct {
	Name     string
	Protocol string
	Address  string
	Status   string
	Age      string
}

func init() {
	listCmd.Flags().StringVarP(&protocol, "protocol", "p", "", "Filter by protocol")
	listCmd.Flags().StringVarP(&status, "status", "s", "", "Filter by status")
	EdgedeviceCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Display edgedevice info in the current Kubernetes cluster",
	Long:    "Show detailed edgedevice information in the current Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.Append(headers)
		table.SetBorder(false)
		table.SetColumnSeparator("")

		devices, err := ListDevices()
		if err != nil {
			logger.Printf("Error listing devices: %v\n", err)
			return
		}

		for _, device := range devices {
			table.Append([]string{
				"\r" + device.Name,
				device.Protocol,
				device.Address,
				device.Status,
				device.Age,
			})
		}

		table.Render()
	},
}

// TimeToAge converts a creation timestamp to a human-readable age string
func TimeToAge(createTime time.Time) string {
	return DurationToMaxUnitString(time.Since(createTime).Round(time.Second))
}

// DurationToMaxUnitString converts a duration to a human-readable string using the largest suitable unit
// Returns a string in the format of "Xd", "Xh", "Xm", or "Xs" where X is the number of days, hours, minutes, or seconds
func DurationToMaxUnitString(d time.Duration) string {
	switch {
	case d >= 24*time.Hour:
		return fmt.Sprintf("%dd", d/(24*time.Hour))
	case d >= time.Hour:
		return fmt.Sprintf("%dh", d/time.Hour)
	case d >= time.Minute:
		return fmt.Sprintf("%dm", d/time.Minute)
	default:
		return fmt.Sprintf("%ds", d/time.Second)
	}
}

func ListDevices() ([]DeviceInfo, error) {
	edgedevices, err := k8s.GetEdgedevices()
	if err != nil {
		return nil, err
	}

	devices := make([]DeviceInfo, 0)

	for _, edgedevice := range edgedevices {
		if protocol != "" && edgedevice.Spec.Protocol != nil && string(*edgedevice.Spec.Protocol) != protocol {
			continue
		}

		device, err := k8s.GetAllByDeviceName(edgedevice.Name)
		if err != nil {
			logger.Printf("Error retrieving device: %v\n", err)
			return nil, err
		}

		address := address.GetRealDeviceAddress(*device)

		phase := "N/A"
		if edgedevice.Status.EdgeDevicePhase != nil {
			phase = string(*edgedevice.Status.EdgeDevicePhase)
		}

		devices = append(devices, DeviceInfo{
			Name:     edgedevice.Name,
			Protocol: string(*edgedevice.Spec.Protocol),
			Address:  address,
			Status:   phase,
			Age:      TimeToAge(edgedevice.CreationTimestamp.Time),
		})
	}

	return devices, nil
}
