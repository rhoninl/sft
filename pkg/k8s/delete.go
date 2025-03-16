package k8s

import (
	"context"
	"fmt"
	"strings"

	"github.com/rhoninl/sft/pkg/utils/logger"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/restmapper"
)

// DeleteYaml deletes resources defined in a given YAML string from a Kubernetes cluster
func DeleteYaml(yamlContent string, ignoreIfNotExists bool) (bool, error) {
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

		// Get the resource's name (required for delete operation)
		resourceName := obj.GetName()
		if resourceName == "" {
			return false, fmt.Errorf("resource name not found in YAML document")
		}

		// Delete the resource from the cluster
		if isNamespaced {
			// Handle namespaced resources
			namespace := obj.GetNamespace()
			if namespace == "" {
				namespace = "default" // Set default namespace if not provided
			}
			err = client.Resource(mapping.Resource).Namespace(namespace).Delete(context.TODO(), resourceName, metav1.DeleteOptions{})
		} else {
			// Handle cluster-scoped resources (like ClusterRole, ClusterRoleBinding)
			err = client.Resource(mapping.Resource).Delete(context.TODO(), resourceName, metav1.DeleteOptions{})
		}

		if err != nil {
			if ignoreIfNotExists && errors.IsNotFound(err) {
				logger.Debugf(logger.Verbose, "resource not found: %s/%s", gvk.Kind, resourceName)
			} else {
				return false, fmt.Errorf("failed to delete resource: %v", err)
			}
		}

		logger.Debugf(logger.Verbose, "deleted resource: %s/%s", gvk.Kind, resourceName)
	}

	return true, nil
}
