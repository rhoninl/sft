package uninstall

import (
	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

func init() {
	UninstallCmd.AddCommand(uninstallTelemetryServiceCmd)
	uninstallTelemetryServiceCmd.Flags().BoolVarP(&ignoreIfNotExists, "ignore-if-not-exists", "i", false, "ignore if the resource not exists")
}

var uninstallTelemetryServiceCmd = &cobra.Command{
	Use:     "telemetryservice",
	Aliases: []string{"ts"},
	Short:   "uninstall telemetryservice component in kubernetes cluster",
	Long:    "uninstall telemetryservice component in kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		yamlContent, err := shifu.Resource(shifu.TypeTelemetryService).GetDeployYaml()
		if err != nil {
			logger.Debug(logger.Verbose, err)
			logger.Println("Failed to retrieve telemetryservice YAML")
			return
		}

		_, err = k8s.DeleteYaml(yamlContent, ignoreIfNotExists)
		if err != nil {
			logger.Debug(logger.Verbose, err)
			logger.Println("Failed to uninstall telemetryservice component")
			return
		}

		logger.Println("TelemetryService component uninstalled successfully")
	},
}
