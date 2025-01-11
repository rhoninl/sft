package forward

import (
	"context"
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
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		edgeDeviceName := args[0]
		var devicePort, localPort string
		if len(args) > 1 {
			forwardPort := args[1]
			ports := strings.Split(forwardPort, ":")
			if len(ports) != 2 {
				logger.Println("Error: Invalid port format. Expected format is 'localPort:devicePort'")
				return
			}
			devicePort = ports[1]
			localPort = ports[0]
		}

		readyChan := make(chan struct{})
		defer close(readyChan)

		if err := ForwardPort(ctx, edgeDeviceName, devicePort, localPort, readyChan); err != nil {
			logger.Println("Error: Failed to forward port:", err)
		}
	},

	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		if err := shifu.CheckShifuInstalled(); err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		deviceNames := k8s.GetValidDeviceNames()
		return deviceNames, cobra.ShellCompDirectiveNoFileComp
	},

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(shifu.CheckShifuInstalled())
	},
}

func ForwardPort(ctx context.Context, deviceName string, devicePort string, localPort string, readyChan chan struct{}) error {
	deployments, err := k8s.GetDeployByEnv("EDGEDEVICE_NAME", deviceName)
	if err != nil {
		logger.Printf("Error: Failed to retrieve deployment for edge device '%s': %v\n", deviceName, err)
		return err
	}

	if len(deployments) == 0 {
		logger.Printf("Error: No deployment found for edge device '%s'\n", deviceName)
		return err
	}

	// if devicePort == 0 {
	// 	if len(deployments[0].Spec.Template.Spec.Containers) == 0 || len(deployments[0].Spec.Template.Spec.Containers[0].Ports) == 0 {
	// 		logger.Println("Error: No container ports found in the deployment")
	// 		return err
	// 	}
	// 	containerPort := deployments[0].Spec.Template.Spec.Containers[0].Ports[0].ContainerPort
	// 	localPort := 3000
	// 	for {
	// 		if !isPortInUse(localPort) {
	// 			break
	// 		}
	// 		localPort++
	// 	}
	// }

	// ports := strings.Split(forwardPort, ":")
	// if len(ports) != 2 {
	// 	logger.Println("Error: Invalid port format. Expected format is 'localPort:remotePort'")
	// 	return
	// }

	pods, err := k8s.GetPodsByDeployment("deviceshifu", deployments[0].Name)
	if err != nil {
		logger.Printf("Error: Failed to retrieve pods for deployment '%s': %v\n", deployments[0].Name, err)
		return err
	}

	logger.Printf("Initiating port-forwarding for pod '%s' on ports %s -> %s \n", deviceName, devicePort, localPort)
	if err := k8s.PortForwardPod(ctx, "deviceshifu", pods[0].Name, devicePort, localPort, readyChan); err != nil {
		logger.Printf("Error: Failed to port-forward pod '%s': %v\n", pods[0].Name, err)
		return err
	}

	return nil
}

func isPortInUse(port int) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
