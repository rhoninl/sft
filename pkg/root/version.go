package root

import (
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/spf13/cobra"
)

func init() {
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show sft version",
	Long:  `All software has versions. This is ShifuTool's`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Println("v0.0.1")
	},
}
