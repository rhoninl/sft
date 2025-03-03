package uninstall

import (
	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

func init() {
	UninstallCmd.AddCommand(UninstallShifuCmd)
	UninstallShifuCmd.Flags().BoolVarP(&ignoreIfNotExists, "ignore-if-not-exists", "i", false, "ignore if the resource not exists")
}

var UninstallShifuCmd = &cobra.Command{
	Use:     "shifu",
	Aliases: []string{"sf"},
	Short:   "uninstall shifu component in kubernetes cluster",
	Long:    "uninstall shifu component in kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if shifu.CheckTelemetryServiceInstalled() == nil {
			uninstallTelemetryServiceCmd.Run(cmd, args)
		}

		if err := UninstallShifu(); err != nil {
			logger.Debug(logger.Verbose, err)
			logger.Println("Failed to uninstall shifu component")
			return
		}

		logger.Println("Shifu component uninstalled successfully")
	},
}

func UninstallShifu() error {
	yamlContent, err := shifu.Resource(shifu.TypeShifu).GetDeployYaml()
	if err != nil {
		return err
	}

	_, err = k8s.DeleteYaml(string(yamlContent), ignoreIfNotExists)
	if err != nil {
		return err
	}

	return nil
}
