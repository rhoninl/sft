package install

import (
	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

func init() {
	InstallCmd.AddCommand(installShifuCmd)
	installShifuCmd.Flags().BoolVarP(&ignoreIfExists, "ignore-if-exists", "i", false, "ignore if the resource already exists")
}

var installShifuCmd = &cobra.Command{
	Use:     "shifu",
	Aliases: []string{"sf"},
	Short:   "install shifu component in kubernetes cluster",
	Long:    "install shifu component in kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		InstallShifu(version)
	},
}

func InstallShifu(version string) error {
	yamlContent, err := shifu.Resource(shifu.TypeShifu).SetVersion(version).GetDeployYaml()
	if err != nil {
		logger.Debugf(logger.Verbose, "Failed to install shifu component: %v", err)
		logger.Println("Failed to install shifu component")
		return err
	}

	_, err = k8s.ApplyYaml(string(yamlContent), ignoreIfExists)
	if err != nil {
		logger.Debugf(logger.Verbose, "Failed to install shifu component: %v", err)
		logger.Println("Failed to install shifu component")
		return err
	}

	logger.Println("Shifu component installed successfully")
	return nil
}
