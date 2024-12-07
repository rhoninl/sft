package k8s

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// RestartDeployment restarts a Kubernetes deployment by scaling it down to zero and then back to its original size.
func RestartDeployment(deploymentName, namespace string) error {
	clientset, _, err := NewClientSet()
	if err != nil {
		return fmt.Errorf("failed to create clientset: %w", err)
	}

	// Annotate the deployment to trigger a restart
	patchData := []byte(`{"spec": {"template": {"metadata": {"annotations": {"kubectl.kubernetes.io/restartedAt": "` + metav1.Now().Format("2006-01-02T15:04:05Z07:00") + `"}}}}}`)
	_, err = clientset.AppsV1().Deployments(namespace).Patch(context.TODO(), deploymentName, types.MergePatchType, patchData, metav1.PatchOptions{})
	if err != nil {
		return fmt.Errorf("failed to restart deployment: %w", err)
	}

	return nil
}
