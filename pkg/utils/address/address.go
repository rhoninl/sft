package address

import (
	"fmt"
	"net"
	"slices"
	"strings"

	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/logger"
)

var internalHost = []string{
	"localhost",
	"0.0.0.0",
}

var addressHallmark = []string{
	"IP",
	"HOST",
	"Address",
	"ADDRESS",
}

func containsAny(arr []string, str string) bool {
	for _, v := range arr {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}

// return str contains n number of arr
func containsN(arr []string, str string) int {
	count := 0
	for _, v := range arr {
		if strings.Contains(str, v) {
			count++
		}
	}
	return count
}

var passiveProtocol = []string{
	"LwM2M",
}

func GetRealDeviceAddress(device k8s.Device) string {
	address := device.EdgeDevice.Spec.Address
	var protocol = *device.EdgeDevice.Spec.Protocol
	if slices.Contains(passiveProtocol, string(protocol)) {
		logger.Debugf(logger.Verbose, "device %s is a passive device, which is hit by %s protocol", device.EdgeDevice.Name, protocol)
		return "[P]" + passiveDeviceServerAddress(&device)
	}

	if address == nil || *address == "" {
		logger.Debugf(logger.Verbose, "device %s has no address", device.EdgeDevice.Name)
		return "N/A"
	}

	host, port, err := net.SplitHostPort(*address)
	if err != nil {
		logger.Debugf(logger.Verbose, "device %s has invalid address: %s, using it as fallback", device.EdgeDevice.Name, *address)
		return *address
	}

	if !containsAny(internalHost, host) {
		logger.Debugf(logger.Verbose, "device %s has external address: %s", device.EdgeDevice.Name, *address)
		return *address
	}

	logger.Debugf(logger.Verbose, "device %s is using a driver, find the driver container", device.EdgeDevice.Name)

	envs := make(map[string]string)

	// Find the container that exposes the port
	for _, container := range device.Deployment.Spec.Template.Spec.Containers {
		var mark bool
		if container.Ports != nil {
			for _, containerPort := range container.Ports {
				if string(containerPort.ContainerPort) == port {
					clear(envs)
					mark = true
				}
			}
		}
		for _, env := range container.Env {
			envs[env.Name] = env.Value
		}

		if mark {
			break
		}
	}

	maxKey := ""
	maxKeyN := 0
	for key := range envs {
		n := containsN(addressHallmark, key)
		if n == 0 {
			delete(envs, key)
		}

		if n > maxKeyN {
			maxKey = key
			maxKeyN = n
		}
	}
	if maxKeyN == 0 {
		logger.Debugf(logger.Verbose, "device %s has no valid address, using passive device server address", device.EdgeDevice.Name)
		return "[P]" + passiveDeviceServerAddress(&device)
	}

	logger.Debugf(logger.Verbose, "device %s has valid address: %s", device.EdgeDevice.Name, envs[maxKey])
	return envs[maxKey]
}

func passiveDeviceServerAddress(device *k8s.Device) string {
	for _, service := range device.Services.Items {
		if service.Spec.Type != "NodePort" {
			continue
		}

		for _, port := range service.Spec.Ports {
			if port.NodePort == 0 {
				continue
			} else if port.TargetPort.IntVal == 8080 {
				continue
			}

			return fmt.Sprintf("localhost:%d", port.NodePort)
		}
	}

	service := device.Services.Items[0]

	return fmt.Sprintf("%s.%s.svc.cluster.local:%d", service.Name, service.Namespace, service.Spec.Ports[0].TargetPort.IntVal)
}
