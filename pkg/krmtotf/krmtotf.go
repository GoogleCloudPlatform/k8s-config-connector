// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package krmtotf

import (
	"fmt"
	"strings"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	tfresource "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/resource"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// KRMResourceToTFResourceConfig converts a KCC KRM resource to a Terraform
// resource config. Note: this function does not fully validate the input KRM
// config or output TF config to ensure that they correspond to valid GCP
// resources (e.g. if the input KRM config is missing a required field, the
// function won't complain and just output a TF config without that field).
// This function just converts one abstract data structure to another;
// validation of either the input KRM or output TF is left as the
// responsibility of other layers (e.g. webhooks, CRD schemas, GCP API, etc.)
func KRMResourceToTFResourceConfig(r *Resource, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (tfConfig *terraform.ResourceConfig, secretVersions map[string]string, err error) {
	return KRMResourceToTFResourceConfigFull(r, c, smLoader, nil, nil, true, label.GetDefaultLabels())
}

// KRMResourceToTFResourceConfigFull is a more flexible version of KRMResourceToTFResourceConfig,
// including the following additional flags:
//   - liveState: if set, these values will be used as the default values of the returned tfConfig, subject to
//     be overridden by r.spec, etc.
//   - jsonSchema: if set, externally managed fields will be populated.
//   - mustResolveSensitiveFields: if set, sensitive fields will be resolved.
//   - defaultLabels: if set, these labels will be added to tfConfig.
func KRMResourceToTFResourceConfigFull(r *Resource, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader,
	liveState *terraform.InstanceState, jsonSchema *apiextensions.JSONSchemaProps, mustResolveSensitiveFields bool, defaultLabels map[string]string) (tfConfig *terraform.ResourceConfig, secretVersions map[string]string, err error) {
	config := deepcopy.MapStringInterface(r.Spec)
	if config == nil {
		config = make(map[string]interface{})
	}
	if jsonSchema != nil {
		if err := ResolveLegacyGCPManagedFields(r, liveState, config); err != nil {
			return nil, nil, fmt.Errorf("error resolving legacy GCP-managed fields: %w", err)
		}
		config, err = resolveUnmanagedFields(config, r, liveState, jsonSchema)
		if err != nil {
			return nil, nil, fmt.Errorf("error resolving externally-managed fields: %w", err)
		}
	}
	if err := handleUserSpecifiedID(config, r, smLoader, c); err != nil {
		return nil, nil, err
	}
	if r.ResourceConfig.MetadataMapping.Labels != "" {
		path := text.SnakeCaseToLowerCamelCase(r.ResourceConfig.MetadataMapping.Labels)
		labels := label.NewGCPLabelsFromK8SLabels(r.GetLabels(), defaultLabels)
		if err := setValue(config, path, labels); err != nil {
			return nil, nil, fmt.Errorf("error mapping 'metadata.labels': %w", err)
		}
	}
	if r.ResourceConfig.Locationality != "" {
		switch r.ResourceConfig.Locationality {
		case gcp.Global:
			delete(config, "location")
		case gcp.Regional:
			config["region"] = config["location"]
			delete(config, "location")
		case gcp.Zonal:
			config["zone"] = config["location"]
			delete(config, "location")
		default:
			return nil, nil, fmt.Errorf("INTERNAL_ERROR: %v locationality is not supported", r.ResourceConfig.Locationality)
		}
	}
	for _, refConfig := range r.ResourceConfig.ResourceReferences {
		if err := handleResourceReference(config, refConfig, r, c, smLoader); err != nil {
			return nil, nil, err
		}
	}
	config, secretVersions, err = resolveSensitiveFields(config, r.TFResource, r.GetNamespace(), c, mustResolveSensitiveFields)
	if err != nil {
		return nil, nil, err
	}
	config, err = KRMObjectToTFObjectWithConfigurableFieldsOnly(config, r.TFResource)
	if err != nil {
		return nil, nil, fmt.Errorf("error converting to config: %w", err)
	}
	for _, d := range r.ResourceConfig.Directives {
		key := k8s.FormatAnnotation(text.SnakeCaseToKebabCase(d))
		if val, ok := k8s.GetAnnotation(key, r); ok {
			if val == "" {
				return nil, nil, fmt.Errorf("the value for directive '%v' must not be empty", key)
			}
			if err := setValue(config, d, val); err != nil {
				return nil, nil, fmt.Errorf("error mapping directive '%v': %w", d, err)
			}
		}
	}
	if err := resolveContainerValue(config, r, c, smLoader); err != nil {
		return nil, nil, fmt.Errorf("error resolving container value: %w", err)
	}
	config, err = withCustomFlatteners(config, r.Kind)
	if err != nil {
		return nil, nil, fmt.Errorf("error running custom flatteners: %w", err)
	}
	// Set desired state with default values.
	defaultingMap := map[string]map[string]string{
		"CloudBuildTrigger": {
			"location": "global",
		},
		"CloudIdentityGroup": {
			"initial_group_config": "EMPTY",
		},
		"FirestoreIndex": {
			"database": "(default)",
		},
	}
	if defaults, ok := defaultingMap[r.Kind]; ok {
		for field, value := range defaults {
			if v, ok := config[field]; !ok || v == "" {
				config[field] = value
			}
		}
	}
	state := InstanceStateToMap(r.TFResource, liveState)
	config, err = withResourceCustomResolvers(config, state, r.Kind, r.TFResource)
	if err != nil {
		return nil, nil, fmt.Errorf("error running resource custom resolver: %w", err)
	}
	return MapToResourceConfig(r.TFResource, config), secretVersions, nil
}

func KRMObjectToTFObject(obj map[string]interface{}, resource *tfschema.Resource) (map[string]interface{}, error) {
	return krmObjectToTFObject(obj, resource, false)
}

func KRMObjectToTFObjectWithConfigurableFieldsOnly(obj map[string]interface{}, resource *tfschema.Resource) (map[string]interface{}, error) {
	return krmObjectToTFObject(obj, resource, true)
}

func krmObjectToTFObject(obj map[string]interface{}, resource *tfschema.Resource, includeConfigurableFieldsOnly bool) (map[string]interface{}, error) {
	var err error
	if obj == nil {
		return nil, nil
	}
	ret := make(map[string]interface{})
	for k, v := range obj {
		tfKey := text.AsSnakeCase(k)
		schema, ok := resource.Schema[tfKey]
		if !ok {
			// TODO(b/239223470): We want to error out explicitly if certain field from spec
			// cannot be mapped to TFObject, instead of silently swallow the error.
			continue
		}
		if includeConfigurableFieldsOnly && !tfresource.IsConfigurableField(schema) {
			continue
		}
		ret[tfKey], err = convertToTF(v, schema, includeConfigurableFieldsOnly)
		if err != nil {
			return nil, fmt.Errorf("error converting '%v': %w", k, err)
		}
	}
	return ret, nil
}

func convertToTF(obj interface{}, schema *tfschema.Schema, includeConfigurableFieldsOnly bool) (interface{}, error) {
	switch schema.Type {
	case tfschema.TypeBool, tfschema.TypeFloat, tfschema.TypeString, tfschema.TypeInt:
		// Treat these values as primitives
		return obj, nil
	case tfschema.TypeMap:
		// Maps are kept identical to the input
		return deepcopy.DeepCopy(obj), nil
	case tfschema.TypeList, tfschema.TypeSet:
		items, err := toList(obj, schema)
		if err != nil {
			return nil, err
		}
		retList := make([]interface{}, 0)
		for _, item := range items {
			var processedItem interface{}
			switch elem := schema.Elem.(type) {
			case *tfschema.Schema:
				processedItem, err = convertToTF(item, elem, includeConfigurableFieldsOnly)
				if err != nil {
					return nil, fmt.Errorf("error converting list item: %w", err)
				}
			case *tfschema.Resource:
				itemAsMap, ok := item.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("expected list item to be map but was not")
				}
				processedItem, err = krmObjectToTFObject(itemAsMap, elem, includeConfigurableFieldsOnly)
				if err != nil {
					return nil, fmt.Errorf("error converting map list item: %w", err)
				}
			default:
				return nil, fmt.Errorf("unknown elem type")
			}
			retList = append(retList, processedItem)
		}
		return retList, nil
	case tfschema.TypeInvalid:
		return nil, fmt.Errorf("schema type is invalid")
	default:
		return nil, fmt.Errorf("unrecognized schema type %v", schema.Type)
	}
}

// handleUserSpecifiedID takes the resource's user-specified ID (if it supports
// one and has one) and places it into the config object. If the resource
// doesn't support user-specified IDs (e.g. supports server-generated IDs
// instead), then this function is a no-op. If the resource does support
// user-specified IDs, then this function tries to get it from the resource's
// spec.resourceID first if specified, and then metadata.name if specified.
func handleUserSpecifiedID(config map[string]interface{}, r *Resource, smLoader *servicemappingloader.ServiceMappingLoader, c client.Client) error {
	if SupportsResourceIDField(&r.ResourceConfig) && !IsResourceIDFieldServerGenerated(&r.ResourceConfig) && r.HasResourceIDField() {
		path := text.SnakeCaseToLowerCamelCase(r.ResourceConfig.ResourceID.TargetField)
		resourceID, err := resolveResourceID(r, c, smLoader)
		if err != nil {
			return fmt.Errorf("error resolving resource ID: %w", err)
		}
		if err := setValue(config, path, resourceID); err != nil {
			return fmt.Errorf("error mapping user-specified %v: %w", k8s.ResourceIDFieldPath, err)
		}
	} else if r.ResourceConfig.MetadataMapping.Name != "" && r.GetName() != "" {
		path := text.SnakeCaseToLowerCamelCase(r.ResourceConfig.MetadataMapping.Name)
		name, err := resolveNameMetadataMapping(r, c, smLoader)
		if err != nil {
			return fmt.Errorf("error resolving metadata.name mapping: %w", err)
		}
		if err := setValue(config, path, name); err != nil {
			return fmt.Errorf("error mapping metadata.name: %w", err)
		}
	}
	return nil
}

func resolveSensitiveFields(config map[string]interface{}, resource *tfschema.Resource, namespace string, c client.Client, mustResolveSensitiveFields bool) (resolvedConfig map[string]interface{}, secretVersions map[string]string, err error) {
	resolvedConfig = deepcopy.MapStringInterface(config)
	secretVersions = make(map[string]string)
	for k, v := range config {
		tfKey := text.AsSnakeCase(k)
		schema, ok := resource.Schema[tfKey]
		if !ok {
			continue
		}
		switch schema.Type {
		case tfschema.TypeString:
			if !tfresource.IsSensitiveConfigurableField(schema) {
				continue
			}

			field := corekccv1alpha1.SensitiveField{}
			if err := util.Marshal(v, &field); err != nil {
				return nil, nil, fmt.Errorf("error parsing %v onto a SensitiveField struct: %w", v, err)
			}

			if field.Value != nil {
				resolvedConfig[k] = *field.Value
				continue
			}

			secretKeyRef := field.ValueFrom.SecretKeyRef
			secretVal, secretVer, err := k8s.GetSecretVal(secretKeyRef, namespace, c)
			if err != nil {
				if mustResolveSensitiveFields {
					return nil, nil, err
				}
				delete(resolvedConfig, k)
				continue
			}
			resolvedConfig[k] = secretVal
			secretVersions[secretKeyRef.Name] = secretVer
		default:
			resolvedObj, secretVers, err := resolveSensitiveFieldsInObj(v, schema, namespace, c, mustResolveSensitiveFields)
			if err != nil {
				return nil, nil, err
			}
			resolvedConfig[k] = resolvedObj
			secretVersions = addToMap(secretVersions, secretVers)
		}
	}
	return resolvedConfig, secretVersions, nil
}

func resolveSensitiveFieldsInObj(obj interface{}, schema *tfschema.Schema, namespace string, c client.Client, mustResolveSensitiveFields bool) (resolvedObj interface{}, secretVersions map[string]string, err error) {
	secretVersions = make(map[string]string)
	switch schema.Type {
	case tfschema.TypeList, tfschema.TypeSet:
		items, err := toList(obj, schema)
		if err != nil {
			return nil, nil, err
		}
		resolvedItems := make([]interface{}, 0)
		for _, item := range items {
			var resolvedItem interface{}
			var secretVers map[string]string
			var err error

			switch elem := schema.Elem.(type) {
			case *tfschema.Schema:
				resolvedItem, secretVers, err = resolveSensitiveFieldsInObj(item, elem, namespace, c, mustResolveSensitiveFields)
				if err != nil {
					return nil, nil, err
				}
			case *tfschema.Resource:
				itemAsMap, ok := item.(map[string]interface{})
				if !ok {
					return nil, nil, fmt.Errorf("expected list item to be map but was not")
				}
				resolvedItem, secretVers, err = resolveSensitiveFields(itemAsMap, elem, namespace, c, mustResolveSensitiveFields)
				if err != nil {
					return nil, nil, err
				}
			}

			resolvedItems = append(resolvedItems, resolvedItem)
			secretVersions = addToMap(secretVersions, secretVers)
		}
		return resolvedItems, secretVersions, nil
	default:
		return obj, secretVersions, nil
	}
}

func toList(obj interface{}, schema *tfschema.Schema) ([]interface{}, error) {
	if obj == nil {
		return nil, nil
	}
	switch obj := obj.(type) {
	case []interface{}:
		return obj, nil
	case map[string]interface{}:
		// An object nested in a KRM field can be interpreted as a list if the
		// corresponding TF field is a list with MaxItems == 1. This is due to
		// limitations with TF schemas.
		if schema.MaxItems == 1 {
			return []interface{}{obj}, nil
		}
		return nil, fmt.Errorf("cannot interpret map as list without maxItems == 1")
	default:
		return nil, fmt.Errorf("cannot interpret non-list %T as list", obj)
	}
}

func setValue(m map[string]interface{}, path string, value interface{}) error {
	return unstructured.SetNestedField(m, value, strings.Split(path, ".")...)
}

// addToMap adds all the key-value pairs from the 'right' map onto the 'left'
// map. If the key already existed in the 'left' map, then it is overridden by
// the value in the 'right' map.
func addToMap(left map[string]string, right map[string]string) map[string]string {
	left = deepcopy.StringStringMap(left)
	for k, v := range right {
		left[k] = v
	}
	return left
}

func resolveContainerValue(config map[string]interface{}, r *Resource, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) error {
	if len(r.ResourceConfig.Containers) == 0 {
		return nil
	}
	if SupportsHierarchicalReferences(&r.ResourceConfig) {
		// If resource supports hierarchical references, use those references
		// instead to set the parent fields in the underlying resource.
		// TODO(b/193177782): Delete this function once all resources support
		// hierarchical references.
		return nil
	}
	for _, container := range r.ResourceConfig.Containers {
		val, ok := k8s.GetAnnotation(k8s.GetAnnotationForContainerType(container.Type), r)
		if !ok {
			continue
		}
		val, err := ResolveValueTemplate(container.ValueTemplate, val, r, c, smLoader)
		if err != nil {
			return fmt.Errorf("error resolving templated value: %w", err)
		}
		if err := setValue(config, container.TFField, val); err != nil {
			return fmt.Errorf("error setting container value: %w", err)
		}
		return nil
	}
	return fmt.Errorf("no annotation found that matches one of the required containers")
}

func resolveNameMetadataMapping(r *Resource, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (string, error) {
	name := r.GetName()
	if name == "" {
		return "", fmt.Errorf("invalid empty value for name")
	}
	return ResolveValueTemplate(r.ResourceConfig.MetadataMapping.NameValueTemplate, name, r, c, smLoader)
}

func resolveResourceID(r *Resource, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (string, error) {
	resourceID, err := r.GetResourceID()
	if err != nil {
		return "", err
	}

	return ResolveValueTemplate(r.ResourceConfig.ResourceID.ValueTemplate, resourceID, r, c, smLoader)
}
