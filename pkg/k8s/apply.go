package k8s

import (
	"context"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/restmapper"
)

// ApplyYaml applies a given YAML string to a Kubernetes cluster
func ApplyYaml(yamlContent string) (bool, error) {
	client, discoveryClient, err := NewClient()
	if err != nil {
		return false, err
	}

	// Split the YAML into individual documents
	yamlDocuments := strings.Split(yamlContent, "---")

	// Create a RESTMapper to map resources to their GroupVersionKind and GroupVersionResource
	groupResources, err := restmapper.GetAPIGroupResources(discoveryClient)
	if err != nil {
		return false, fmt.Errorf("failed to get API group resources: %v", err)
	}
	restMapper := restmapper.NewDiscoveryRESTMapper(groupResources)

	for _, yamlDoc := range yamlDocuments {
		// Trim whitespace and ignore empty documents
		yamlDoc = strings.TrimSpace(yamlDoc)
		if len(yamlDoc) == 0 {
			continue
		}

		// Decode the YAML content into an unstructured object
		decUnstructured := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
		obj := &unstructured.Unstructured{}
		_, _, err = decUnstructured.Decode([]byte(yamlDoc), nil, obj)
		if err != nil {
			return false, fmt.Errorf("failed to decode YAML: %v", err)
		}

		// Map the object's GVK to its GVR (GroupVersionResource) dynamically
		gvk := obj.GroupVersionKind()
		mapping, err := restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
		if err != nil {
			return false, fmt.Errorf("failed to map GVK to GVR: %v", err)
		}

		// Determine if the resource is namespaced or cluster-scoped
		isNamespaced := isNamespacedResource(mapping.Resource.Resource)

		// Apply the resource to the cluster
		if isNamespaced {
			// Handle namespaced resources
			namespace := obj.GetNamespace()
			if namespace == "" {
				namespace = "default" // Set default namespace if not provided
			}
			_, err = client.Resource(mapping.Resource).Namespace(namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
		} else {
			// Handle cluster-scoped resources (like ClusterRole, ClusterRoleBinding)
			_, err = client.Resource(mapping.Resource).Create(context.TODO(), obj, metav1.CreateOptions{})
		}

		if err != nil {
			return false, fmt.Errorf("failed to apply resource: %v", err)
		}

	}

	return true, nil
}
