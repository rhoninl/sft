package forward

import (
	"fmt"
	"strings"

	"github.com/rhoninl/shifucli/cmd/k8s"
	"github.com/spf13/cobra"
)

func init() {
}

var ForwardCmd = &cobra.Command{
	Use:     "forward",
	Aliases: []string{"frd", "f"},
	Short:   "forward deviceshifu api to local",
	Long:    `forward deviceshifu api to local`,
	Run: func(cmd *cobra.Command, args []string) {
		edgeDeviceName := args[0]
		forwardPort := args[1]
		ports := strings.Split(forwardPort, ":")
		if len(ports) != 2 {
			fmt.Println("invalid port")
			return
		}

		deployments, err := k8s.GetDeployByEnv("EDGEDEVICE_NAME", edgeDeviceName)
		if err != nil {
			fmt.Println("failed to get deployment", err)
			return
		}

		if len(deployments) == 0 {
			fmt.Println("no deployment found")
			return
		}

		pods, err := k8s.GetPodsByDeployment("deviceshifu", deployments[0].Name)
		if err != nil {
			fmt.Println("failed to get pod", err)
			return
		}

		k8s.PortForwardPod("deviceshifu", pods[0].Name, ports[0], ports[1])
	},
}
