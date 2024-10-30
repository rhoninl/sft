package edgedevices

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/edgenesis/shifu/pkg/k8s/api/v1alpha1"
	"github.com/olekukonko/tablewriter"
	"github.com/rhoninl/shifucli/cmd/k8s"
	"github.com/rhoninl/shifucli/cmd/utils/logger"
	"github.com/spf13/cobra"
)

var (
	edgedeviceLogTemplate = "%s\t%s\t%s\t%s\t%s\n"
	protocol              string
	status                string
	headers               = []string{"Name", "Protocol", "Address", "Status", "AGE"}
)

func init() {
	listCmd.Flags().StringVarP(&protocol, "protocol", "p", "", "filter by protocol")
	listCmd.Flags().StringVarP(&status, "status", "s", "", "filter by status")
	EdgedeviceCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "edgedevice info in current kubernetes cluster",
	Long:    `show edgedevice info in current kubernetes cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		edgedevices, err := k8s.GetEdgedevices()
		if err != nil {
			cobra.CheckErr(err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Append(headers)
		table.SetBorder(false)
		table.SetColumnSeparator("")

		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 10, ' ', 0)

		for _, edgedevice := range edgedevices {
			if protocol != "" && string(*edgedevice.Spec.Protocol) != protocol {
				continue
			}

			var address string
			if edgedevice.Spec.Address != nil {
				address = *edgedevice.Spec.Address
			} else {
				address = "N/A"
			}

			var phase v1alpha1.EdgeDevicePhase
			if edgedevice.Status.EdgeDevicePhase != nil {
				phase = *edgedevice.Status.EdgeDevicePhase
			} else {
				phase = "N/A"
			}

			table.Append([]string{edgedevice.Name, string(*edgedevice.Spec.Protocol), address, logger.StatusWithColor(string(phase)), TimeToAge(edgedevice.CreationTimestamp.Time)})
			writer.Flush()
		}

		table.Render()
	},
}

func TimeToAge(createTime time.Time) string {
	return DurationToMaxUnitString(time.Since(createTime).Round(time.Second))
}

func DurationToMaxUnitString(d time.Duration) string {
	// Calculate the individual units
	days := d / (24 * time.Hour)
	hours := d / time.Hour
	minutes := d / time.Minute
	seconds := d / time.Second

	// Return the largest unit representation
	if days > 0 {
		return fmt.Sprintf("%dd", days)
	} else if hours > 0 {
		return fmt.Sprintf("%dh", hours)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm", minutes)
	} else {
		return fmt.Sprintf("%ds", seconds)
	}
}
