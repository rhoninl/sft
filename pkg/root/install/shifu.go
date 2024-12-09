package install

import (
	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

func init() {
	InstallCmd.AddCommand(installShifuCmd)
}

var installShifuCmd = &cobra.Command{
	Use:     "shifu",
	Aliases: []string{"sf"},
	Short:   "install shifu component in kubernetes cluster",
	Long:    "install shifu component in kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		yamlContent, err := shifu.Resource(shifu.TypeShifu).SetVersion(version).GetDeployYaml()
		if err != nil {
			logger.Debugf(logger.Verbose, "Failed to install shifu component: %v", err)
			logger.Println("Failed to install shifu component")
			return
		}

		_, err = k8s.ApplyYaml(string(yamlContent))
		if err != nil {
			logger.Debugf(logger.Verbose, "Failed to install shifu component: %v", err)
			logger.Println("Failed to install shifu component")
			return
		}

		logger.Println("Shifu component installed successfully")
	},
}
