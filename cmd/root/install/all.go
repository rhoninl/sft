package install

import (
	"github.com/spf13/cobra"
)

func init() {
	InstallCmd.AddCommand(installAllCmd)
}

var installAllCmd = &cobra.Command{
	Use:     "all",
	Aliases: []string{"a"},
	Short:   "install all component in kubernetes cluster",
	Long:    "install all component in kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		installShifuCmd.Run(cmd, args)
		installTelemetryServiceCmd.Run(cmd, args)
	},
}
