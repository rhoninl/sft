package logs

import (
	"fmt"
	"strconv"

	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/spf13/cobra"
	appv1 "k8s.io/api/apps/v1"
)

var follow bool
var container string

func init() {
	LogsCmd.Flags().BoolVarP(&follow, "follow", "f", false, "Specify if the logs should be streamed.")
	LogsCmd.Flags().StringVarP(&container, "container", "c", "1", "Specify which container will log")
}

var LogsCmd = &cobra.Command{
	Use:     "logs",
	Aliases: []string{"log"},
	Short:   "Display logs of the Shifu component in the Kubernetes cluster",
	Long:    "Display logs of the Shifu component in the Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Device name is required")
			return
		}

		deviceName := args[0]
		deployments, err := k8s.GetDeployByEnv("EDGEDEVICE_NAME", deviceName)
		if err != nil {
			fmt.Printf("Error retrieving deployments: %v\n", err)
			return
		}

		if len(deployments) == 0 {
			fmt.Println("Invalid device name or no deployments found")
			return
		}

		containerName := GetContainerName(deployments[0], container)
		if err := k8s.GetDeploymentLogs("deviceshifu", deployments[0].Name, containerName, follow); err != nil {
			fmt.Printf("Error retrieving logs: %v\n", err)
		}
	},
	ValidArgs: k8s.GetValidDeviceNames(),
}

func GetContainerName(deployments appv1.Deployment, flag string) string {
	containers := deployments.Spec.Template.Spec.Containers
	var containerName = flag

	if containerNumber, err := strconv.Atoi(flag); err != nil {
		// TODO: Do nothing now
	} else if containerNumber <= len(containers) {
		containerName = containers[containerNumber-1].Name
	}

	return containerName
}
