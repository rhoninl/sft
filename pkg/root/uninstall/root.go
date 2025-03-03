package uninstall

import (
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

var (
	ignoreIfNotExists bool
)

func init() {
	UninstallCmd.Flags().BoolVarP(&ignoreIfNotExists, "ignore-if-not-exists", "i", false, "ignore if the resource not exists")
}

var UninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Aliases: []string{"ui"},
	Short:   "way to uninstall shifu component in kubernetes cluster",
	Long:    "way to uninstall shifu component in kubernetes cluster",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(shifu.CheckShifuInstalled())
	},
}
