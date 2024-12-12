package k8s

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

var LogLine int64

func GetDeploymentLogs(namespace, deploymentName, containerName string, follow bool) error {
	dynamicClient, err := NewDynamicClient()
	if err != nil {
		return fmt.Errorf("failed to create clientset: %w", err)
	}

	clientset, _, err := NewClientSet()
	if err != nil {
		return fmt.Errorf("failed to create clientset: %w", err)
	}

	// Define the GroupVersionResource for Deployment (apps/v1)
	gvr := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}

	// Get the Deployment details
	deployment, err := dynamicClient.Resource(gvr).Namespace(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get deployment: %w", err)
	}

	// Extract the label selector from the deployment spec
	spec, ok := deployment.Object["spec"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to get spec from deployment")
	}

	selector, ok := spec["selector"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to get selector from spec")
	}

	matchLabels, ok := selector["matchLabels"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to get matchLabels from selector")
	}

	// Convert matchLabels to a label selector string
	labelSelector := buildLabelSelector(matchLabels)

	// List Pods associated with the deployment
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return fmt.Errorf("failed to list pods: %w", err)
	}

	// Check if any pods are found
	if len(pods.Items) == 0 {
		return fmt.Errorf("no pods found for deployment %s", deploymentName)
	}

	// Get logs for each pod and container (if specified)
	for _, pod := range pods.Items {
		if err := streamPodLogs(clientset, namespace, pod.Name, containerName, follow); err != nil {
			fmt.Printf("Error streaming logs for pod %s: %v\n", pod.Name, err)
		}
	}

	return nil
}

func buildLabelSelector(matchLabels map[string]interface{}) string {
	var labelSelector strings.Builder
	for key, value := range matchLabels {
		labelSelector.WriteString(fmt.Sprintf("%s=%s,", key, value))
	}
	return strings.TrimSuffix(labelSelector.String(), ",")
}

func streamPodLogs(clientset *kubernetes.Clientset, namespace, podName, containerName string, follow bool) error {
	// Prepare PodLogOptions (filter by container if needed)
	logOptions := &v1.PodLogOptions{
		Follow:    follow,
		TailLines: &LogLine,
	}
	if containerName != "" {
		logOptions.Container = containerName
	}

	// Get the logs for the pod (and container)
	req := clientset.CoreV1().Pods(namespace).GetLogs(podName, logOptions)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return fmt.Errorf("error retrieving logs for pod %s: %w", podName, err)
	}
	defer podLogs.Close()

	if _, err := io.Copy(os.Stdout, podLogs); err != nil && err != io.EOF {
		return fmt.Errorf("error streaming logs for pod %s: %w", podName, err)
	}

	return nil
}
