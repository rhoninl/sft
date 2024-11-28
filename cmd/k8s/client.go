package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/edgenesis/shifu/pkg/k8s/api/v1alpha1"
	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var KubeConfigPath string

func NewDynamicClient() (*dynamic.DynamicClient, error) {
	if len(KubeConfigPath) == 0 {
		KubeConfigPath = os.Getenv("KUBECONFIG")
		if KubeConfigPath == "" {
			if home := homedir.HomeDir(); home != "" {
				KubeConfigPath = filepath.Join(home, ".kube", "config")
			}
		}
	}

	config, err := clientcmd.BuildConfigFromFlags("", KubeConfigPath)
	if err != nil {
		return nil, err
	}

	dyClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return dyClient, nil
}

func NewDiscoveryClient() (*discovery.DiscoveryClient, error) {
	if len(KubeConfigPath) == 0 {
		KubeConfigPath = os.Getenv("KUBECONFIG")
		if KubeConfigPath == "" {
			if home := homedir.HomeDir(); home != "" {
				KubeConfigPath = filepath.Join(home, ".kube", "config")
			}
		}
	}

	config, err := clientcmd.BuildConfigFromFlags("", KubeConfigPath)
	if err != nil {
		return nil, err
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create discovery client: %v", err)
	}

	return discoveryClient, nil
}

func NewClientSet() (*kubernetes.Clientset, *rest.Config, error) {
	if len(KubeConfigPath) == 0 {
		KubeConfigPath = os.Getenv("KUBECONFIG")
		if KubeConfigPath == "" {
			if home := homedir.HomeDir(); home != "" {
				KubeConfigPath = filepath.Join(home, ".kube", "config")
			}
		}
	}

	config, err := clientcmd.BuildConfigFromFlags("", KubeConfigPath)
	if err != nil {
		return nil, nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}

	return clientset, config, nil
}

func NewClient() (*dynamic.DynamicClient, *discovery.DiscoveryClient, error) {
	if len(KubeConfigPath) == 0 {
		KubeConfigPath = os.Getenv("KUBECONFIG")
		if KubeConfigPath == "" {
			if home := homedir.HomeDir(); home != "" {
				KubeConfigPath = filepath.Join(home, ".kube", "config")
			}
		}
	}

	config, err := clientcmd.BuildConfigFromFlags("", KubeConfigPath)
	if err != nil {
		return nil, nil, err
	}

	dyClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}
	// Create a discovery client
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create discovery client: %v", err)
	}

	// Return both clients
	return dyClient, discoveryClient, nil
}

func GetEdgedevices() ([]v1alpha1.EdgeDevice, error) {
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
	edgedevices, err := client.Resource(gvr).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var edgedeviceList []v1alpha1.EdgeDevice
	for _, edgedeviceItem := range edgedevices.Items {
		edgedeviceByte, err := json.Marshal(edgedeviceItem.Object)
		if err != nil {
			return nil, err
		}
		var edgedevice v1alpha1.EdgeDevice
		if err := json.Unmarshal(edgedeviceByte, &edgedevice); err != nil {
			return nil, err
		}

		edgedeviceList = append(edgedeviceList, edgedevice)
	}

	return edgedeviceList, nil
}

func GetResource(name string, namespace string, group string, version string, resource string) (*unstructured.Unstructured, error) {
	client, _, err := NewClient()
	if err != nil {
		return nil, err
	}

	gvr := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}

	obj, err := client.Resource(gvr).Namespace(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// isNamespacedResource determines if a resource is namespaced or cluster-scoped
func isNamespacedResource(resource string) bool {
	// List of cluster-scoped resources
	clusterScopedResources := map[string]bool{
		"namespaces":                true,
		"customresourcedefinitions": true,
		"clusterroles":              true,
		"clusterrolebindings":       true,
		"nodes":                     true,
		"persistentvolumes":         true,
		// Add more cluster-scoped resources as needed
	}

	if _, ok := clusterScopedResources[resource]; ok {
		return false
	}
	return true
}

func GetDeploymentPods(namespace string, deploymentName string) ([]v1.Pod, error) {
	clientset, _, err := NewClientSet()
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	// 1. Get the Deployment resource
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get deployment: %v", err)
	}

	// 2. Get the list of Pods associated with the Deployment
	labelSelector := deployment.Spec.Selector.MatchLabels
	labelSelectorStr := metav1.FormatLabelSelector(&metav1.LabelSelector{MatchLabels: labelSelector})
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelectorStr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list pods for deployment: %v", err)
	}

	return pods.Items, nil
}

func GetDeploymentFirstReplicaStatus(namespace, deploymentName string) (string, error) {
	clientset, _, err := NewClientSet()
	if err != nil {
		return "", fmt.Errorf("failed to create clientset: %w", err)
	}

	// 1. 获取 Deployment 资源
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to get deployment: %v", err)
	}

	// 2. 获取与 Deployment 关联的 Pod 列表
	labelSelector := deployment.Spec.Selector.MatchLabels
	labelSelectorStr := metav1.FormatLabelSelector(&metav1.LabelSelector{MatchLabels: labelSelector})
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelectorStr,
	})
	if err != nil {
		return "", fmt.Errorf("failed to list pods for deployment: %v", err)
	}

	if len(pods.Items) == 0 {
		return "", fmt.Errorf("no pods found for deployment %s", deploymentName)
	}

	firstPod := pods.Items[0]

	if firstPod.Status.Phase != v1.PodRunning {
		containerErrors := getContainerErrorReasons(firstPod)
		if len(containerErrors) > 0 {
			return containerErrors[0], nil
		}
	}

	var containerErrors []string
	for _, containerStatus := range firstPod.Status.ContainerStatuses {
		if !containerStatus.Ready {
			containerErrors = append(containerErrors, fmt.Sprintf("Container %s is not ready", containerStatus.Name))
		}
		if containerStatus.State.Waiting != nil {
			containerErrors = append(containerErrors, fmt.Sprintf("Container %s is in waiting state: %s", containerStatus.Name, containerStatus.State.Waiting.Reason))
		}
		if containerStatus.State.Terminated != nil {
			containerErrors = append(containerErrors, fmt.Sprintf("Container %s is terminated: %s", containerStatus.Name, containerStatus.State.Terminated.Reason))
		}
	}

	if len(containerErrors) > 0 {
		return containerErrors[0], nil
	}

	return "Running", nil
}

func getContainerErrorReasons(pod v1.Pod) []string {
	var reasons []string

	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.State.Waiting != nil && containerStatus.State.Waiting.Reason != "" {
			reasons = append(reasons, containerStatus.State.Waiting.Reason)
		}
		if containerStatus.State.Terminated != nil && containerStatus.State.Terminated.Reason != "" {
			reasons = append(reasons, containerStatus.State.Terminated.Reason)
		}
	}

	return reasons
}

func GetDeployByEnv(key string, value string) ([]appv1.Deployment, error) {
	clientset, _, err := NewClientSet()
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var result []appv1.Deployment

	for _, deployment := range deployments.Items {
		for _, container := range deployment.Spec.Template.Spec.Containers {
			for _, env := range container.Env {
				if env.Name == key && env.Value == value {
					result = append(result, deployment)
					break
				}
			}
		}
	}

	return result, nil
}

func GetPodsByDeployment(namespace, deploymentName string) ([]v1.Pod, error) {
	clientset, _, err := NewClientSet()
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %w", err)
	}

	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get deployment: %v", err)
	}

	labelSelector := deployment.Spec.Selector.MatchLabels
	labelSelectorStr := metav1.FormatLabelSelector(&metav1.LabelSelector{MatchLabels: labelSelector})
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelectorStr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list pods for deployment: %v", err)
	}

	return pods.Items, nil
}
