package root

import (
	"github.com/rhoninl/shifucli/cmd/k8s"
	"github.com/rhoninl/shifucli/cmd/root/describe"
	"github.com/rhoninl/shifucli/cmd/root/edgedevices"
	"github.com/rhoninl/shifucli/cmd/root/forward"
	"github.com/rhoninl/shifucli/cmd/root/install"
	"github.com/rhoninl/shifucli/cmd/root/logs"
	"github.com/rhoninl/shifucli/cmd/root/uninstall"
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
	RootCmd.AddCommand(edgedevices.EdgedeviceCmd)
	RootCmd.AddCommand(serviceCmd)
	RootCmd.AddCommand(describe.DescribeCmd)
	RootCmd.AddCommand(statusCmd)
	RootCmd.AddCommand(forward.ForwardCmd)
	RootCmd.AddCommand(logs.LogsCmd)
	RootCmd.AddCommand(install.InstallCmd)
	RootCmd.AddCommand(uninstall.UninstallCmd)
	RootCmd.AddCommand(versionCmd)
}
