package forward

import (
	"fmt"
	"net"
	"strings"

	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/rhoninl/sft/pkg/utils/shifu"
	"github.com/spf13/cobra"
)

func init() {
}

var ForwardCmd = &cobra.Command{
	Use:     "forward",
	Aliases: []string{"frd", "f"},
	Short:   "Forward the deviceshifu API to a local port",
	Long:    `Forward the deviceshifu API to a local port for easier access and testing`,
	Run: func(cmd *cobra.Command, args []string) {
		edgeDeviceName := args[0]
		var forwardPort string
		if len(args) > 1 {
			forwardPort = args[1]
		}

		deployments, err := k8s.GetDeployByEnv("EDGEDEVICE_NAME", edgeDeviceName)
		if err != nil {
			logger.Printf("Error: Failed to retrieve deployment for edge device '%s': %v\n", edgeDeviceName, err)
			return
		}

		if len(deployments) == 0 {
			logger.Printf("Error: No deployment found for edge device '%s'\n", edgeDeviceName)
			return
		}

		if forwardPort == "" {
			if len(deployments[0].Spec.Template.Spec.Containers) == 0 || len(deployments[0].Spec.Template.Spec.Containers[0].Ports) == 0 {
				logger.Println("Error: No container ports found in the deployment")
				return
			}
			containerPort := deployments[0].Spec.Template.Spec.Containers[0].Ports[0].ContainerPort
			localPort := 3000
			for {
				if !isPortInUse(localPort) {
					forwardPort = fmt.Sprintf("%d:%d", localPort, containerPort)
					break
				}
				localPort++
			}
		}

		ports := strings.Split(forwardPort, ":")
		if len(ports) != 2 {
			logger.Println("Error: Invalid port format. Expected format is 'localPort:remotePort'")
			return
		}

		pods, err := k8s.GetPodsByDeployment("deviceshifu", deployments[0].Name)
		if err != nil {
			logger.Printf("Error: Failed to retrieve pods for deployment '%s': %v\n", deployments[0].Name, err)
			return
		}

		logger.Printf("Initiating port-forwarding for pod '%s' on ports %s -> %s\n", pods[0].Name, ports[0], ports[1])
		if err := k8s.PortForwardPod("deviceshifu", pods[0].Name, ports[0], ports[1]); err != nil {
			logger.Printf("Error: Failed to port-forward pod '%s': %v\n", pods[0].Name, err)
			return
		}
	},

	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return k8s.GetValidDeviceNames(), cobra.ShellCompDirectiveNoFileComp
	},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(shifu.CheckShifuInstalled())
	},
}

func isPortInUse(port int) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
