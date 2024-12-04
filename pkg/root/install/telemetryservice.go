package install

import (
	"os/exec"

	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

func init() {
	InstallCmd.AddCommand(installTelemetryServiceCmd)
}

var installTelemetryServiceCmd = &cobra.Command{
	Use:     "telemetryservice",
	Aliases: []string{"ts"},
	Short:   "install telemetryservice component in kubernetes cluster",
	Long:    "install telemetryservice component in kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		yamlFile := "https://raw.githubusercontent.com/Edgenesis/shifu/" + shifu.GetLatestShifuVersion() + "/pkg/telemetryservice/install/telemetryservice_install.yaml"
		kcmd := exec.Command("kubectl", "apply", "-f", yamlFile)
		if err := kcmd.Run(); err != nil {
			logger.Println("Failed to install telemetryservice component, please install shifu component first")
			return
		}
		logger.Println("TelemetryService component installed successfully")
	},
}
