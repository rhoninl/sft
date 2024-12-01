package uninstall

import (
	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

func init() {
	UninstallCmd.AddCommand(UninstallShifuCmd)
}

var UninstallShifuCmd = &cobra.Command{
	Use:     "shifu",
	Aliases: []string{"sf"},
	Short:   "uninstall shifu component in kubernetes cluster",
	Long:    "uninstall shifu component in kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		yamlContent, err := shifu.GetInstallYaml(shifu.GetLatestShifuVersion())
		if err != nil {
			logger.Debug(err)
			logger.Println("Failed to install shifu component")
			return
		}

		_, err = k8s.DeleteYaml(string(yamlContent))
		if err != nil {
			logger.Debug(err)
			logger.Println("Failed to install shifu component")
			return
		}

		logger.Println("Shifu component installed successfully")
	},
}
