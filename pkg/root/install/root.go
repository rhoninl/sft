package install

import (
	"github.com/spf13/cobra"
)

func init() {

}

var InstallCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i"},
	Short:   "way to install shifu component in kubernetes cluster",
	Long:    "way to install shifu component in kubernetes cluster",
}
