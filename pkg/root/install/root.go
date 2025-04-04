package install

import (
	"github.com/spf13/cobra"
)

var (
	version        string
	ignoreIfExists bool
)

const (
	EmptyVersion = ""
)

func init() {
	installAllCmd.Flags().StringVar(&version, "version", EmptyVersion, "install all components")
}

var InstallCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i"},
	Short:   "way to install shifu component in kubernetes cluster",
	Long:    "way to install shifu component in kubernetes cluster",
}
