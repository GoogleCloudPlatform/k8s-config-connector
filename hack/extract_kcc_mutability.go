package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	tpgProvider "github.com/hashicorp/terraform-provider-google-beta/google-beta/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v3"
)

type ResourceInfo struct {
	ControllerType  string `json:"controller_type"`
	ResourceName    string `json:"resource_name"`
	ServiceHostName string `json:"service_hostname,omitempty"`
	Kind            string `json:"kind"`
	Group           string `json:"group"`
}

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	smLoader, err := servicemappingloader.New()
	if err != nil {
		fmt.Printf("Error initializing servicemappingloader: %v\n", err)
		os.Exit(1)
	}

	resourcesFile := filepath.Join(rootDir, "resources.json")
	content, err := ioutil.ReadFile(resourcesFile)
	if err != nil {
		fmt.Printf("Error reading resources.json: %v\n", err)
		os.Exit(1)
	}

	var resources map[string]ResourceInfo
	if err := json.Unmarshal(content, &resources); err != nil {
		fmt.Printf("Error unmarshaling resources.json: %v\n", err)
		os.Exit(1)
	}

	mutabilityData := make(map[string][]string)
	tfProvider := tpgProvider.Provider()

	for kind, info := range resources {
		var immutableFields []string

		if info.ControllerType == "TF" {
			sm, _ := smLoader.GetServiceMapping(info.Group)
			immutableFields = extractTFMutability(tfProvider, info.ResourceName, kind, sm)
		} else if info.ControllerType == "DCL" {
			immutableFields = extractDCLMutability(rootDir, info.ServiceHostName, info.ResourceName)
		}

		if immutableFields == nil {
			immutableFields = []string{}
		}
		sort.Strings(immutableFields)
		mutabilityData[kind] = immutableFields
	}

	outFile := filepath.Join(rootDir, "kcc_mutability.json")
	outData, err := json.MarshalIndent(mutabilityData, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling mutability data: %v\n", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(outFile, outData, 0644); err != nil {
		fmt.Printf("Error writing kcc_mutability.json: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully extracted mutability for %d resources to kcc_mutability.json\n", len(mutabilityData))
}

func extractTFMutability(tfProvider *schema.Provider, tfResourceName string, kind string, sm *v1alpha1.ServiceMapping) []string {
	var immutableFields []string

	tfRes, ok := tfProvider.ResourcesMap[tfResourceName]
	if !ok {
		return nil
	}

	// Build a map of references
	refs := make(map[string]string)
	if sm != nil {
		for _, r := range sm.Spec.Resources {
			if r.Kind == kind {
				for _, ref := range r.ResourceReferences {
					key := ref.Key
					if key == "" {
						// Fallback if Key is empty, KCC generates it, but usually Key is set.
						key = text.SnakeCaseToLowerCamelCase(ref.TFField) + "Ref"
					}
					refs[ref.TFField] = key
				}
				for _, hRef := range r.HierarchicalReferences {
					if hRef.Type == "project" {
						refs["project"] = "projectRef"
					} else if hRef.Type == "folder" {
						refs["folder"] = "folderRef"
					} else if hRef.Type == "organization" {
						refs["organization"] = "organizationRef"
					} else if hRef.Type == "billingAccount" {
						refs["billing_account"] = "billingAccountRef"
					}
				}
				// Also handle containers if present
				for _, c := range r.Containers {
					refs[c.TFField] = text.SnakeCaseToLowerCamelCase(c.TFField) + "Ref"
				}
			}
		}
	}

	var walk func(s map[string]*schema.Schema, tfPath []string, krmPath []string)
	walk = func(s map[string]*schema.Schema, tfPath []string, krmPath []string) {
		for fieldName, fieldSchema := range s {
			currentTFPath := append([]string{}, tfPath...)
			currentTFPath = append(currentTFPath, fieldName)
			fullTFField := strings.Join(currentTFPath, ".")

			currentKRMPath := append([]string{}, krmPath...)
			
			// Check if this full TF path is a reference
			if key, isRef := refs[fullTFField]; isRef {
				// Replace the last segment with the reference key
				currentKRMPath = append(currentKRMPath, key)
			} else {
				currentKRMPath = append(currentKRMPath, text.SnakeCaseToLowerCamelCase(fieldName))
			}

			fullKRMField := strings.Join(currentKRMPath, ".")

			if fieldSchema.ForceNew {
				immutableFields = append(immutableFields, fullKRMField)
			}

			if elem, ok := fieldSchema.Elem.(*schema.Resource); ok {
				walk(elem.Schema, currentTFPath, currentKRMPath)
			}
		}
	}

	walk(tfRes.Schema, []string{}, []string{})
	return immutableFields
}

func extractDCLMutability(rootDir, serviceName, resourceName string) []string {
	var immutableFields []string
	fileName := toSnakeCase(resourceName) + ".yaml"
	baseDir := filepath.Join(rootDir, "third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google", serviceName)
	pathsToCheck := []string{
		filepath.Join(baseDir, "beta", fileName),
		filepath.Join(baseDir, fileName),
		filepath.Join(baseDir, "alpha", fileName),
	}

	var yamlContent []byte
	var err error
	for _, p := range pathsToCheck {
		yamlContent, err = ioutil.ReadFile(p)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil
	}

	var schemaData map[string]interface{}
	if err := yaml.Unmarshal(yamlContent, &schemaData); err != nil {
		return nil
	}

	components, ok := schemaData["components"].(map[string]interface{})
	if !ok { return nil }
	schemas, ok := components["schemas"].(map[string]interface{})
	if !ok { return nil }
	resSchema, ok := schemas[resourceName].(map[string]interface{})
	if !ok { return nil }
	properties, ok := resSchema["properties"].(map[string]interface{})
	if !ok { return nil }

	var walk func(props map[string]interface{}, krmPath []string)
	walk = func(props map[string]interface{}, krmPath []string) {
		for fieldName, fieldInfoInter := range props {
			fieldInfo, ok := fieldInfoInter.(map[string]interface{})
			if !ok { continue }

			// KCC uses camelCase for DCL fields (already true in DCL YAML).
			// If it's a reference, append "Ref".
			krmSegment := fieldName
			if _, hasRef := fieldInfo["x-dcl-references"]; hasRef {
				krmSegment = krmSegment + "Ref"
			}
			
			// Hierarchical fields in DCL
			if fieldName == "project" || fieldName == "folder" || fieldName == "organization" {
				krmSegment = fieldName + "Ref"
			}

			currentKRMPath := append([]string{}, krmPath...)
			currentKRMPath = append(currentKRMPath, krmSegment)
			fullKRMField := strings.Join(currentKRMPath, ".")

			if immutable, ok := fieldInfo["x-kubernetes-immutable"].(bool); ok && immutable {
				immutableFields = append(immutableFields, fullKRMField)
			}

			if nestedProps, ok := fieldInfo["properties"].(map[string]interface{}); ok {
				walk(nestedProps, currentKRMPath)
			}
			
			if items, ok := fieldInfo["items"].(map[string]interface{}); ok {
				if nestedProps, ok := items["properties"].(map[string]interface{}); ok {
					walk(nestedProps, currentKRMPath)
				}
			}
		}
	}

	walk(properties, []string{})
	return immutableFields
}

func toSnakeCase(s string) string {
	switch s {
	case "TenantOAuthIdpConfig": return "tenant_oauth_idp_config"
	case "OAuthIdpConfig": return "oauth_idp_config"
	case "OSPolicyAssignment": return "os_policy_assignment"
	}

	var res string
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				res += "_"
			}
			res += strings.ToLower(string(r))
		} else {
			res += string(r)
		}
	}
	return res
}