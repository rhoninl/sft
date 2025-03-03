package install

import (
	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

func init() {
	InstallCmd.AddCommand(installTelemetryServiceCmd)
	installTelemetryServiceCmd.Flags().BoolVarP(&ignoreIfExists, "ignore-if-exists", "i", false, "ignore if the resource already exists")
}

var installTelemetryServiceCmd = &cobra.Command{
	Use:     "telemetryservice",
	Aliases: []string{"ts"},
	Short:   "install telemetryservice component in kubernetes cluster",
	Long:    "install telemetryservice component in kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		yamlContent, err := shifu.Resource(shifu.TypeTelemetryService).SetVersion(version).GetDeployYaml()
		if err != nil {
			logger.Debug(logger.Verbose, err)
			logger.Println("Failed to retrieve telemetryservice YAML")
			return
		}

		_, err = k8s.ApplyYaml(yamlContent, ignoreIfExists)
		if err != nil {
			logger.Debug(logger.Verbose, err)
			logger.Println("Failed to install telemetryservice component")
			return
		}

		logger.Println("TelemetryService component install successfully")
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := shifu.CheckTelemetryServiceInstalled(); err != nil {
			cobra.CheckErr(err)
		}
	},
}
