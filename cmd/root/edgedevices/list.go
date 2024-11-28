package edgedevices

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/rhoninl/shifucli/cmd/k8s"
	"github.com/rhoninl/shifucli/cmd/utils/logger"
	"github.com/spf13/cobra"
)

var (
	edgedeviceLogTemplate = "%s\t%s\t%s\t%s\t%s\n"
	protocol              string
	status                string
	headers               = []string{"\rName", "Protocol", "Address", "Status", "AGE"}
)

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
		edgedevices, err := k8s.GetEdgedevices()
		if err != nil {
			cobra.CheckErr(err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Append(headers)
		table.SetBorder(false)
		table.SetColumnSeparator("")

		for _, edgedevice := range edgedevices {
			if protocol != "" && edgedevice.Spec.Protocol != nil && string(*edgedevice.Spec.Protocol) != protocol {
				continue
			}

			address := "N/A"
			if edgedevice.Spec.Address != nil {
				address = *edgedevice.Spec.Address
			}

			phase := "N/A"
			if edgedevice.Status.EdgeDevicePhase != nil {
				phase = string(*edgedevice.Status.EdgeDevicePhase)
			}

			table.Append([]string{
				"\r" + edgedevice.Name,
				string(*edgedevice.Spec.Protocol),
				address,
				logger.StatusWithColor(phase),
				TimeToAge(edgedevice.CreationTimestamp.Time),
			})
		}

		table.Render()
	},
}

func TimeToAge(createTime time.Time) string {
	return DurationToMaxUnitString(time.Since(createTime).Round(time.Second))
}

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
