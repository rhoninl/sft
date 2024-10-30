package root

import (
	"fmt"
	"strings"

	"github.com/rhoninl/shifucli/cmd/k8s"
	"github.com/rhoninl/shifucli/cmd/utils/logger"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func init() {
}

var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"sts"},
	Short:   "show shifu status",
	Long:    `show shifu status`,
	Run: func(cmd *cobra.Command, args []string) {
		var message string
		version, status, err := GetShifuInfo()
		if err != nil {
			message = logger.WithColor(logger.Red, "Not Installed")
		} else {
			message = version + " " + logger.StatusWithColor(status)
		}

		fmt.Println("Shifu: ", message)
	},
}

func GetShifuInfo() (string, string, error) {
	deployment, err := k8s.GetResource("shifu-crd-controller-manager", "shifu-crd-system", "apps", "v1", "deployments")
	if err != nil {
		fmt.Println("Failed to get shifu-crd-controller-manager deployment", err)
		return "", "", err
	}

	shifuVersion, err := getShifuVersionFromDeployment(deployment)
	if err != nil {
		fmt.Println("Failed to get shifu version", err)
		return "", "", err
	}

	status, err := k8s.GetDeploymentFirstReplicaStatus("shifu-crd-system", "shifu-crd-controller-manager")
	if err != nil {
		fmt.Println("Failed to get shifu status", err)
		return "", "", err
	}

	return shifuVersion, status, nil
}

func getShifuVersionFromDeployment(deployment *unstructured.Unstructured) (string, error) {
	containers, found, err := unstructured.NestedSlice(deployment.Object, "spec", "template", "spec", "containers")
	if !found || err != nil {
		fmt.Printf("Error retrieving containers: %v\n", err)
		return "", err
	}
	// Ensure that there is at least one container
	if len(containers) == 0 {
		fmt.Println("No containers found in the deployment")
		return "", nil
	}

	// Get the second container (it's of type map[string]interface{})
	firstContainer, ok := containers[1].(map[string]interface{})
	if !ok {
		fmt.Println("Error casting first container")
		return "", nil
	}

	image, found, err := unstructured.NestedString(firstContainer, "image")
	if !found || err != nil {
		fmt.Printf("Error retrieving image: %v\n", err)
		return "", err
	}

	return strings.Split(image, ":")[1], nil
}
