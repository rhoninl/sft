package uninstall

import (
	"github.com/spf13/cobra"
)

func init() {

}

var UninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Aliases: []string{"ui"},
	Short:   "way to uninstall shifu component in kubernetes cluster",
	Long:    "way to uninstall shifu component in kubernetes cluster",
}
