package uninstall

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	UninstallCmd.AddCommand(installAllCmd)
}

var installAllCmd = &cobra.Command{
	Use:     "all",
	Aliases: []string{"a"},
	Short:   "uninstall all component in kubernetes cluster",
	Long:    "uninstall all component in kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("uninstall all")
	},
}
