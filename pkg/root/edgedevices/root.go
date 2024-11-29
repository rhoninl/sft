package edgedevices

import (
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
}
