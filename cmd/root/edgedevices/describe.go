package edgedevices

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	EdgedeviceCmd.AddCommand(describeCmd)
}

var describeCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"desc"},
	Short:   "show edgedevice detail info",
	Long:    `show edgedevice detail info`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(args)
	},
}
