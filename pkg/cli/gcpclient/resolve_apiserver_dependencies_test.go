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

package gcpclient_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	testservicemapping "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemapping"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	kubeschema "k8s.io/apimachinery/pkg/runtime/schema"
)

/*
 * The functions in this file enable one to pass in a set of unstructureds and remove all API server dependencies, i.e. the following is done:
 * 1. All resource references are externalized
 * 2. Secrets are resolved to their values.
 */

func resolveAPIServerDependenciesIfKCCManaged(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *schema.Provider,
	resources []*unstructured.Unstructured, resource *unstructured.Unstructured) {
	if !k8s.IsManagedByKCC(resource.GroupVersionKind()) {
		return
	}
	externalizeReferences(t, smLoader, resources, resource)
	resolveSensitiveFields(t, smLoader, tfProvider, resources, resource)
}

// this function changes a SensitiveField to a "ValueFrom" to a "Value" so that the 'krmtotf' package will not try to resolve
// the values using the client.Client
func resolveSensitiveFields(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *schema.Provider,
	resources []*unstructured.Unstructured, resource *unstructured.Unstructured) {
	rc := testservicemapping.GetResourceConfig(t, smLoader, resource)
	spec, ok := resource.Object["spec"]
	if !ok {
		return
	}
	tfResource := tfProvider.ResourcesMap[rc.Name]
	tfSchema := schema.Schema{
		Elem: tfResource,
	}
	convertSensitiveFieldsInObjectToValue(t, resources, &tfSchema, spec)
}

func convertSensitiveFieldsInObjectToValue(t *testing.T, resources []*unstructured.Unstructured, tfSchema *schema.Schema, obj interface{}) {
	switch elem := tfSchema.Elem.(type) {
	case *schema.Schema:
		convertSensitiveFieldsToValue(t, resources, elem, obj)
	case *schema.Resource:
		for k, v := range elem.Schema {
			mapValue, ok := obj.(map[string]interface{})
			if !ok {
				t.Fatalf("expected resource type to have an obj type of map[string]interface, instead was '%v'", reflect.TypeOf(obj).Name())
			}
			k8sKey := text.SnakeCaseToLowerCamelCase(k)
			val, ok := mapValue[k8sKey]
			if !ok {
				continue
			}
			convertSensitiveFieldsToValue(t, resources, v, val)
		}
	}
}

func convertSensitiveFieldsToValue(t *testing.T, resources []*unstructured.Unstructured, tfSchema *schema.Schema, obj interface{}) {
	switch tfSchema.Type {
	case schema.TypeString:
		if !(tfSchema.Sensitive && isConfigurableField(tfSchema)) {
			return
		}
		var field v1alpha1.SensitiveField
		if err := util.Marshal(obj, &field); err != nil {
			t.Fatalf("error marshalling '%v' to a SensitiveField: %v", obj, err)
		}
		if field.Value != nil {
			return
		}
		secretKeyRef := field.ValueFrom.SecretKeyRef
		secretValue := findResource(t, resources, secretKeyRef.Name, kubeschema.GroupVersionKind{Group: "", Version: "v1", Kind: "Secret"})
		value, err := getSecretValue(secretValue, secretKeyRef.Key)
		if err != nil {
			t.Fatalf("error getting secret value: %v", err)
		}
		field.ValueFrom = nil
		field.Value = &value
		var newValue map[string]interface{}
		if err := util.Marshal(field, &newValue); err != nil {
			t.Fatalf("error marshalling SensitiveField '%v' to a object map: %v", field, err)
		}
		mapValue := obj.(map[string]interface{})
		for k := range mapValue {
			delete(mapValue, k)
		}
		for k, v := range newValue {
			mapValue[k] = v
		}
	case schema.TypeList, schema.TypeSet:
		values, err := toList(obj, tfSchema)
		if err != nil {
			t.Fatalf("error converting to list: %v", err)
		}
		for _, v := range values {
			convertSensitiveFieldsInObjectToValue(t, resources, tfSchema, v)
		}
	}
}

func getSecretValue(secretResource *unstructured.Unstructured, key string) (string, error) {
	var secret v1.Secret
	if err := util.Marshal(secretResource, &secret); err != nil {
		return "", fmt.Errorf("error marshalling '%v' to a Secret: %w", secretResource, err)
	}
	if secret.Data != nil {
		value, ok := secret.Data[key]
		if ok {
			return string(value), nil
		}
	}
	if secret.StringData != nil {
		value, ok := secret.StringData[key]
		if ok {
			return value, nil
		}
	}
	return "", fmt.Errorf("unexpected missing value for key '%v' in secret data", key)
}

func toList(obj interface{}, schema *schema.Schema) ([]interface{}, error) {
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

func isConfigurableField(schema *schema.Schema) bool {
	return schema.Required || schema.Optional
}

func externalizeReferences(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader, resources []*unstructured.Unstructured, resource *unstructured.Unstructured) {
	rc := testservicemapping.GetResourceConfig(t, smLoader, resource)
	for _, ref := range rc.ResourceReferences {
		// TODO: can we make this resolution of the reference key smoother across the codebase?
		// There is overlap with how we are resolving the key here and with krmttotf.GetPathToReferenceKey(...)
		// However there are subtle differences in that
		tfFields := strings.Split(ref.TFField, ".")
		rootPath := []string{"spec"}
		for _, field := range tfFields {
			rootPath = append(rootPath, text.SnakeCaseToLowerCamelCase(field))
		}
		typeConfigs := ref.Types
		if len(typeConfigs) == 0 {
			typeConfigs = []v1alpha1.TypeConfig{ref.TypeConfig}
			// The path to a field in the 'spec' is the same as the path in the terraform resource with snake case converted
			// to camel case. However, in the case of references, the last portion of the path is replaced with the value in 'Key'.
			// For example, in ComputeDisk, the reference with tf field 'disk_encryption_key.kms_key_self_link' converts to
			// 'diskEncryption.kmsKeySelfLink' in the spec, but then the 'kmsKeySelfLink' is replaced with 'kmsKeyRef' as
			// that is filled into the 'Key' field. For that reason, we remove the last value of the path here.
			//
			// When there is more than one possible reference value, then ref.Types will contain a non-empty list. For those
			// cases the full ref.TFField converted to camel case is used in the path.
			if ref.TypeConfig.Key != "" {
				rootPath = rootPath[:len(rootPath)-1]
			}
		}
		for _, tc := range typeConfigs {
			findMatchingReferencesAndExternalize(t, rootPath, tc, resources, resource)
		}
	}
}

func findMatchingReferencesAndExternalize(t *testing.T, rootPath []string, tc v1alpha1.TypeConfig, resources []*unstructured.Unstructured, resource *unstructured.Unstructured) {
	keyFields := rootPath
	if tc.Key != "" {
		keyFields = append(rootPath, tc.Key)
	}
	rawValue, ok, err := getValueInSliceAwareNestedMap(resource.Object, keyFields)
	if err != nil {
		t.Fatalf("error retrieving field '%v': %v", tc.Key, err)
	}
	if !ok {
		return
	}
	switch value := rawValue.(type) {
	case map[string]interface{}:
		externalizeReference(t, value, tc, resources, resource)
	case []interface{}:
		for _, v := range value {
			mapValue, ok := v.(map[string]interface{})
			if !ok {
				t.Fatalf("unexpected non-map value in resource reference list value at '%v': %v",
					strings.Join(rootPath, "."), v)
			}
			externalizeReference(t, mapValue, tc, resources, resource)
		}
	}
}

// Walks the given path through the object. It is assumed the object is either of type map[string]interface{} or
// []interface{} and the same is true for each sub-key
//
// The reason why the map contains slices in the middle of paths is a terraform oddity where often times there can be a
// list of size 1 for various reasons
func getValueInSliceAwareNestedMap(object interface{}, path []string) (interface{}, bool, error) {
	if len(path) == 0 {
		return object, true, nil
	}
	switch t := object.(type) {
	case map[string]interface{}:
		value, ok := t[path[0]]
		if !ok {
			return nil, false, nil
		}
		return getValueInSliceAwareNestedMap(value, path[1:])
	case []interface{}:
		for _, subValue := range t {
			value, ok, err := getValueInSliceAwareNestedMap(subValue, path)
			if err != nil {
				return nil, false, err
			}
			if ok {
				return value, ok, nil
			}
		}
		return nil, false, nil
	}
	return nil, false, fmt.Errorf("error, unknown type: %v", reflect.TypeOf(object).Name())
}

func findResource(t *testing.T, resources []*unstructured.Unstructured, name string, gvk kubeschema.GroupVersionKind) *unstructured.Unstructured {
	var matches []*unstructured.Unstructured
	for _, r := range resources {
		if r.GetName() != name {
			continue
		}
		if r.GroupVersionKind() != gvk {
			continue
		}
		matches = append(matches, r)
	}
	if len(matches) == 0 {
		t.Fatalf("unable to find resource matching %v/%v", gvk, name)
	}
	if len(matches) > 1 {
		t.Fatalf("found multiple resources matching %v/%v", gvk, name)
	}

	return matches[0]
}

func externalizeReference(t *testing.T, value map[string]interface{}, tc v1alpha1.TypeConfig, resources []*unstructured.Unstructured, resource *unstructured.Unstructured) {
	resourceReference := toResourceReference(t, value)
	if resourceReference.External != "" {
		return
	}

	refResource := findResource(t, resources, resourceReference.Name, tc.GVK)
	externalVal := resolveTargetFieldValue(t, refResource, tc)
	// Do basic value template expansion, as we don't have *krmtotf.Resource,
	// *client.Client, or *servicemappingloader.ServiceMappingLoader present in
	// order to do full expansion.
	if tc.ValueTemplate != "" {
		externalVal = strings.ReplaceAll(tc.ValueTemplate, "{{value}}", externalVal)
		if strings.Contains(externalVal, "{{project}}") {
			project, found := k8s.GetAnnotation(k8s.ProjectIDAnnotation, refResource)
			if !found {
				t.Fatalf("unsupported project value template resolution")
			}
			externalVal = strings.ReplaceAll(externalVal, "{{project}}", project)
		}
		if strings.Contains(externalVal, "{{location}}") {
			location, found, _ := unstructured.NestedString(resource.Object, strings.Split("spec.location", ".")...)
			if !found {
				t.Fatalf("unsupported location value template resolution")
			}
			externalVal = strings.ReplaceAll(externalVal, "{{location}}", location)
		}
		if strings.Contains(externalVal, "{{") {
			t.Fatalf("test does not support template expansion required for value template '%v'", tc.ValueTemplate)
		}
	}
	newResourceReference := v1alpha1.ResourceReference{
		External: externalVal,
	}
	newValue := resourceReferenceToMap(t, newResourceReference)
	// replace the fields in the 'old' map, the map itself may be nested under a slice, so it is important
	// to modify that value or write a complicated function which traverses the object structure resolving
	// through slices.
	for k := range value {
		delete(value, k)
	}
	for k, v := range newValue {
		if v != "" {
			value[k] = v
		}
	}
}

// TODO: this is copied from references.go and is changed to use unstructured, can we consolidate?
func resolveTargetFieldValue(t *testing.T, resource *unstructured.Unstructured, tc v1alpha1.TypeConfig) string {
	key := text.SnakeCaseToLowerCamelCase(tc.TargetField)
	switch key {
	case "":
		return resource.GetName()
	default:
		if val, exists, _ := unstructured.NestedString(resource.Object, strings.Split(fmt.Sprintf("spec.%v", key), ".")...); exists {
			return val
		}
		if val, exists, _ := unstructured.NestedString(resource.Object, strings.Split(fmt.Sprintf("status.%v", key), ".")...); exists {
			return val
		}
		// For now, we do not support recursive target field resolution (i.e. targeting a field in
		// the referenced resource that itself is a reference to a third resource, which would require
		// its own target field resolution).
		t.Fatalf("referenced resource's target field %v is unsupported", tc.TargetField)
	}
	panic("got here which should be impossible")
}

func toResourceReference(t *testing.T, value map[string]interface{}) v1alpha1.ResourceReference {
	bytes, err := yaml.Marshal(value)
	if err != nil {
		t.Fatalf("error marshalling '%v' to YAML: %v", value, err)
	}
	var resourceReference v1alpha1.ResourceReference
	if err := yaml.Unmarshal(bytes, &resourceReference); err != nil {
		t.Fatalf("error unmarshalling to resource reference; %v", err)
	}
	return resourceReference
}

func resourceReferenceToMap(t *testing.T, resourceReference v1alpha1.ResourceReference) map[string]interface{} {
	bytes, err := yaml.Marshal(resourceReference)
	if err != nil {
		t.Fatalf("error marshalling '%v' to YAML: %v", resourceReference, err)
	}
	var value map[string]interface{}
	if err := yaml.Unmarshal(bytes, &value); err != nil {
		t.Fatalf("error unmarshalling to resource reference; %v", err)
	}
	return value
}
