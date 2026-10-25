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

package stream

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/randomid"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	tfresource "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/resource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const namePath = "metadata.name"

type FixupOptions struct {
	FilterSensitiveFields bool
	SMLoader              *servicemappingloader.ServiceMappingLoader
	TFProvider            *schema.Provider
}

// UnstructuredResourceFixupStream will make any necessary edits to a resource to make sure it will work with K8s and KCC
// For example, ensuring the metadata.name is a K8s legal value (some GCP resource's allow values in 'name' with
// characters that are not allowed by K8s in metadata.name)
type UnstructuredResourceFixupStream struct {
	unstructStream UnstructuredStream
	opts           FixupOptions
}

func NewUnstructuredResourceFixupStream(unstructuredStream UnstructuredStream) *UnstructuredResourceFixupStream {
	return NewUnstructuredResourceFixupStreamWithOptions(unstructuredStream, FixupOptions{})
}

func NewUnstructuredResourceFixupStreamWithOptions(unstructuredStream UnstructuredStream, opts FixupOptions) *UnstructuredResourceFixupStream {
	return &UnstructuredResourceFixupStream{
		unstructStream: unstructuredStream,
		opts:           opts,
	}
}

func (s *UnstructuredResourceFixupStream) Next(ctx context.Context) (*unstructured.Unstructured, error) {
	u, err := s.unstructStream.Next(ctx)
	if err != nil {
		return nil, err
	}
	if err := defaultNameIfNotPresent(u); err != nil {
		return nil, err
	}
	if err := ensureNameIsK8sLegal(u); err != nil {
		return nil, err
	}
	if s.opts.FilterSensitiveFields {
		filterSensitiveFields(u, s.opts.SMLoader, s.opts.TFProvider)
	}
	return u, nil
}

func filterSensitiveFields(u *unstructured.Unstructured, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *schema.Provider) {
	if u == nil || u.Object == nil {
		return
	}
	spec, ok := u.Object["spec"].(map[string]interface{})
	if !ok {
		return
	}
	if smLoader != nil && tfProvider != nil {
		if rc, err := smLoader.GetResourceConfig(u); err == nil && rc != nil {
			if tfResource, ok := tfProvider.ResourcesMap[rc.Name]; ok && tfResource != nil {
				tfSchema := &schema.Schema{Elem: tfResource}
				filterSensitiveFieldsUsingTFSchema(tfSchema, spec)
			}
		}
	}
	filterSensitiveFieldStructsInMap(spec)
}

func filterSensitiveFieldsUsingTFSchema(tfSchema *schema.Schema, obj interface{}) {
	if tfSchema == nil || obj == nil {
		return
	}
	switch elem := tfSchema.Elem.(type) {
	case *schema.Schema:
		filterSensitiveFieldValueUsingTFSchema(elem, obj)
	case *schema.Resource:
		mapValue, ok := obj.(map[string]interface{})
		if !ok {
			return
		}
		for k, v := range elem.Schema {
			k8sKey := text.SnakeCaseToLowerCamelCase(k)
			val, exists := mapValue[k8sKey]
			if !exists {
				continue
			}
			if tfresource.IsSensitiveConfigurableField(v) {
				if valMap, isMap := val.(map[string]interface{}); isMap {
					delete(valMap, "value")
				} else {
					mapValue[k8sKey] = ""
				}
				continue
			}
			filterSensitiveFieldValueUsingTFSchema(v, val)
		}
	}
}

func filterSensitiveFieldValueUsingTFSchema(tfSchema *schema.Schema, obj interface{}) {
	if tfSchema == nil || obj == nil {
		return
	}
	switch tfSchema.Type {
	case schema.TypeList, schema.TypeSet:
		switch typedObj := obj.(type) {
		case []interface{}:
			for _, item := range typedObj {
				filterSensitiveFieldsUsingTFSchema(tfSchema, item)
			}
		case map[string]interface{}:
			if tfSchema.MaxItems == 1 {
				filterSensitiveFieldsUsingTFSchema(tfSchema, typedObj)
			}
		}
	}
}

func filterSensitiveFieldStructsInMap(m map[string]interface{}) {
	for _, val := range m {
		switch typedVal := val.(type) {
		case map[string]interface{}:
			if isSensitiveFieldMap(typedVal) {
				delete(typedVal, "value")
			} else {
				filterSensitiveFieldStructsInMap(typedVal)
			}
		case []interface{}:
			filterSensitiveFieldStructsInSlice(typedVal)
		}
	}
}

func filterSensitiveFieldStructsInSlice(s []interface{}) {
	for _, item := range s {
		switch typedItem := item.(type) {
		case map[string]interface{}:
			filterSensitiveFieldStructsInMap(typedItem)
		case []interface{}:
			filterSensitiveFieldStructsInSlice(typedItem)
		}
	}
}

func isSensitiveFieldMap(m map[string]interface{}) bool {
	if len(m) == 0 || len(m) > 2 {
		return false
	}
	val, hasValue := m["value"]
	if !hasValue {
		return false
	}
	if _, isStr := val.(string); !isStr {
		return false
	}
	if len(m) == 1 {
		return true
	}
	_, hasValueFrom := m["valueFrom"]
	return hasValueFrom
}

// if the resource does not have a name, fill it with the value in the resourceID field
// Not all resources support resourceID yet, temporarily fill in such resources with a generated name until the feature
// is completely rolled out
func defaultNameIfNotPresent(u *unstructured.Unstructured) error {
	_, ok, err := unstructured.NestedString(u.Object, strings.Split(namePath, ".")...)
	if err != nil {
		return fmt.Errorf("error retrieving %v from unstruct: %w", namePath, err)
	}
	if ok {
		return nil
	}
	val, ok, err := unstructured.NestedString(u.Object, strings.Split(k8s.ResourceIDFieldPath, ".")...)
	if err != nil {
		return fmt.Errorf("error retrieving %v from unstruct: %w", k8s.ResourceIDFieldPath, err)
	}
	if !ok {
		val = randomid.New().String()
	}
	if err := unstructured.SetNestedField(u.Object, val, strings.Split(namePath, ".")...); err != nil {
		return fmt.Errorf("error setting %v to '%v': %w", namePath, val, err)
	}
	return nil
}

// GCP resources names are not always valid K8s resource names, this function modifies the metadata.name to ensure it
// is always K8s legal.
func ensureNameIsK8sLegal(u *unstructured.Unstructured) error {
	name, ok, err := unstructured.NestedString(u.Object, strings.Split(namePath, ".")...)
	if err != nil {
		return fmt.Errorf("error retrieving %v from unstruct: %w", namePath, err)
	}
	if !ok {
		return fmt.Errorf("unexpected empty value for %v in resource", namePath)
	}
	// From https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#customresourcedefinitions,
	// a custom resource's name must be a valid DNS subdomain name
	k8sName := k8s.ValueToDNSSubdomainName(name)
	if k8sName == name {
		return nil
	}
	if err := unstructured.SetNestedField(u.Object, k8sName, strings.Split(namePath, ".")...); err != nil {
		return fmt.Errorf("error setting %v to '%v': %w", namePath, k8sName, err)
	}
	return nil
}
