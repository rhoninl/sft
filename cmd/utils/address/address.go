package address

import (
	"fmt"
	"net"
	"slices"
	"strings"

	"github.com/rhoninl/shifucli/cmd/k8s"
)

var internalHost = []string{
	"localhost",
	"0.0.0.0",
	"",
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

	if slices.Contains(passiveProtocol, string(*device.EdgeDevice.Spec.Protocol)) {
		return "[P]" + passiveDeviceServerAddress(&device)
	}

	if address == nil {
		return "N/A"
	} else if *address == "" {
		return "N/A"
	}

	host, port, err := net.SplitHostPort(*address)
	if err != nil {
		return *address
	}

	if !containsAny(internalHost, host) {
		return *address
	}

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
		return "[P]" + passiveDeviceServerAddress(&device)
	}

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
