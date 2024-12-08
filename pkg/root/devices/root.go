package devices

import (
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

func init() {

}

var EdgedeviceCmd = &cobra.Command{
	Use:     "devices",
	Aliases: []string{"d", "device"},
	Short:   "edgedevice info in current kubernetes cluster",
	Long:    `show edgedevice info in current kubernetes cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		listCmd.Run(cmd, args)
	},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(shifu.CheckShifuInstalled())
	},
}
