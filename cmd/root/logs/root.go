package logs

import (
	"fmt"
	"strconv"

	"github.com/rhoninl/shifucli/cmd/k8s"
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
	Short:   "way to show logs of shifu component in kubernetes cluster",
	Long:    "way to show logs of shifu component in kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		deviceName := args[0]
		deployments, err := k8s.GetDeployByEnv("EDGEDEVICE_NAME", deviceName)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(deployments) == 0 {
			fmt.Println("invalid device name")
			return
		}

		containerName := GetContainerName(deployments[0], container)
		k8s.GetDeploymentLogs("deviceshifu", deployments[0].Name, containerName, follow)
	},
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
