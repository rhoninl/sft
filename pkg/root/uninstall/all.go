package uninstall

import (
	"github.com/spf13/cobra"
)

func init() {
	UninstallCmd.AddCommand(uninstallAllCmd)
}

var uninstallAllCmd = &cobra.Command{
	Use:     "all",
	Aliases: []string{"a"},
	Short:   "uninstall all component in kubernetes cluster",
	Long:    "uninstall all component in kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		UninstallCmd.Run(cmd, args)
	},
}
