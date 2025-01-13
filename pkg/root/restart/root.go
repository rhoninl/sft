package restart

import (
	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

var RestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart shifu",
	Long:  `restart shifu`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logger.Println("Please provide a device name")
			return
		}

		deviceName := args[0]
		err := RestartDeviceShifu(deviceName)
		cobra.CheckErr(err)
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

func RestartDeviceShifu(deviceName string) error {
	deployment, err := k8s.GetDeployByEnv("EDGEDEVICE_NAME", deviceName)
	if err != nil {
		logger.Println("Failed to get deployment", err)
		return err
	}

	if err := k8s.RestartDeployment(deployment[0].Name, deployment[0].Namespace); err != nil {
		logger.Println("Failed to restart deployment", err)
		return err
	}

	return nil
}
