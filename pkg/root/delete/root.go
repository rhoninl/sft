package delete

import (
	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

func init() {
	DeleteCommand.CompletionOptions.DisableDefaultCmd = false
}

var DeleteCommand = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "delete the device in the current Kubernetes cluster",
	Long:    `delete the device in the current Kubernetes cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			logger.Println("Error: Device name is required")
			return
		}

		deviceName := args[0]
		device, err := k8s.GetAllByDeviceName(deviceName)
		if err != nil {
			logger.Printf("Error retrieving device: %v\n", err)
			return
		}

		_ = device
		// printDeviceDetails(device)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		if err := shifu.CheckShifuInstalled(); err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		deviceNames := k8s.GetValidDeviceNames()
		return deviceNames, cobra.ShellCompDirectiveNoFileComp
	},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(shifu.CheckShifuInstalled())
	},
}
