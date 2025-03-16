package shifu

import (
	"fmt"

	"github.com/edgenesis/shifu/pkg/logger"
	"github.com/rhoninl/sft/pkg/k8s"
)

func DeleteDevice(deviceName string) error {
	device, err := k8s.GetAllByDeviceName(deviceName)
	if err != nil {
		return fmt.Errorf("failed to get all by device name: %w", err)
	}

	// Delete EdgeDevice
	edgeDeviceYaml, err := k8s.ToYaml(device.EdgeDevice)
	if err != nil {
		return fmt.Errorf("failed to convert EdgeDevice to yaml: %w", err)
	}
	if _, err := k8s.DeleteYaml(edgeDeviceYaml, true); err != nil {
		return fmt.Errorf("failed to delete EdgeDevice: %w", err)
	}

	// Delete Deployment
	deploymentYaml, err := k8s.ToYaml(device.Deployment)
	if err != nil {
		return fmt.Errorf("failed to convert Deployment to yaml: %w", err)
	}
	if _, err := k8s.DeleteYaml(deploymentYaml, true); err != nil {
		logger.Infof("failed to delete Deployment: %w", err)
		return fmt.Errorf("failed to delete Deployment: %w", err)
	}

	// Delete ConfigMap
	configMapYaml, err := k8s.ToYaml(device.ConfigMap)
	if err != nil {
		return fmt.Errorf("failed to convert ConfigMap to yaml: %w", err)
	}
	if _, err := k8s.DeleteYaml(configMapYaml, true); err != nil {
		return fmt.Errorf("failed to delete ConfigMap: %w", err)
	}

	// Delete Services
	for _, service := range device.Services.Items {
		serviceYaml, err := k8s.ToYaml(service)
		if err != nil {
			return fmt.Errorf("failed to convert Service to yaml: %w", err)
		}
		if _, err := k8s.DeleteYaml(serviceYaml, true); err != nil {
			return fmt.Errorf("failed to delete Service: %w", err)
		}
	}

	return nil
}
