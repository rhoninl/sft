package k8s

import (
	"encoding/json"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

// ToYaml converts a Kubernetes resource to YAML string
func ToYaml(obj interface{}) (string, error) {
	// First convert to JSON
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	// Convert JSON to YAML
	yamlData, err := yaml.JSONToYAML(jsonData)
	if err != nil {
		return "", err
	}

	return string(yamlData), nil
}

// ToUnstructured converts a Kubernetes resource to Unstructured
func ToUnstructured(obj runtime.Object) (map[string]interface{}, error) {
	// Convert to JSON
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// Convert JSON to map[string]interface{}
	var unstructured map[string]interface{}
	if err := json.Unmarshal(jsonData, &unstructured); err != nil {
		return nil, err
	}

	return unstructured, nil
} 