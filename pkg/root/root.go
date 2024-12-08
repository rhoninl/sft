package root

import (
	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/root/describe"
	"github.com/rhoninl/sft/pkg/root/devices"
	"github.com/rhoninl/sft/pkg/root/forward"
	"github.com/rhoninl/sft/pkg/root/install"
	"github.com/rhoninl/sft/pkg/root/logs"
	"github.com/rhoninl/sft/pkg/root/restart"
	"github.com/rhoninl/sft/pkg/root/uninstall"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "sft",
		Short: "shifuctl controls the shifu manager",
		Long:  `shifuctl controls the shifu manager`,
	}
)

// Execute executes the root command.
func Execute() error {
	return RootCmd.Execute()
}
func init() {
	cobra.EnableCommandSorting = false
	RootCmd.CompletionOptions.DisableDefaultCmd = false
	RootCmd.PersistentFlags().StringVar(&k8s.KubeConfigPath, "config", "", "kubeconfig file (default use KUBECONFIG env or ~/.kube/config)")
	RootCmd.PersistentFlags().IntVarP(&logger.DebugLogLevel, "verbose", "v", 0, "debug level, availables: 1, 2")

	RootCmd.AddCommand(devices.EdgedeviceCmd)
	RootCmd.AddCommand(describe.DescribeCmd)
	RootCmd.AddCommand(restart.RestartCmd)
	RootCmd.AddCommand(statusCmd)
	RootCmd.AddCommand(forward.ForwardCmd)
	RootCmd.AddCommand(logs.LogsCmd)
	RootCmd.AddCommand(install.InstallCmd)
	RootCmd.AddCommand(uninstall.UninstallCmd)
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(webCmd)
}
