package k8s

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/edgenesis/shifu/pkg/k8s/api/v1alpha1"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type Device struct {
	EdgeDevice v1alpha1.EdgeDevice
	Deployment appv1.Deployment
	ConfigMap  corev1.ConfigMap
}

func GetAllByDeviceName(deviceName string) (*Device, error) {
	client, _, err := NewClient()
	if err != nil {
		return nil, err
	}

	gvr := schema.GroupVersionResource{
		Group:    "shifu.edgenesis.io",
		Version:  "v1alpha1",
		Resource: "edgedevices",
	}

	namespace := "devices"

	edgedeviceRow, err := client.Resource(gvr).Namespace(namespace).Get(context.TODO(), deviceName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return nil, fmt.Errorf("edgedevice %s not found", deviceName)
	}
	var edgedevice v1alpha1.EdgeDevice
	edgedeviceByte, err := json.Marshal(edgedeviceRow.Object)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(edgedeviceByte, &edgedevice); err != nil {
		return nil, err
	}

	deployment, err := GetDeployByEnv("EDGEDEVICE_NAME", deviceName)
	if errors.IsNotFound(err) {
		return nil, fmt.Errorf("deployment %s not found", deviceName)
	}

	if deployment[0].Spec.Template.Spec.Volumes[0].VolumeSource.ConfigMap == nil {
		return nil, fmt.Errorf("configmap not found")
	}

	configmapName := deployment[0].Spec.Template.Spec.Volumes[0].VolumeSource.ConfigMap.Name
	configmap, err := GetConfigmapByName(configmapName)
	if err != nil {
		return nil, err
	}

	return &Device{
		EdgeDevice: edgedevice,
		Deployment: deployment[0],
		ConfigMap:  *configmap,
	}, nil
}

func GetConfigmapByName(name string) (*corev1.ConfigMap, error) {
	clientset, _, err := NewClientSet()
	if err != nil {
		return nil, err
	}

	obj, err := clientset.CoreV1().ConfigMaps("deviceshifu").Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return obj, nil
}
