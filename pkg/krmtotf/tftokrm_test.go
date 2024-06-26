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

package krmtotf_test

import (
	"reflect"
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	. "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testk8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/k8s"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

func TestConvertTFObjToKCCObj(t *testing.T) {
	tests := []struct {
		name           string
		rc             *corekccv1alpha1.ResourceConfig
		state          map[string]interface{}
		schemaOverride map[string]*tfschema.Schema
		prevSpec       map[string]interface{}
		managedFields  *fieldpath.Set
		expected       map[string]interface{}
	}{
		{
			name: "defaulted non-zero primitive values are set",
			state: map[string]interface{}{
				"int_key":    float64(1),
				"float_key":  float64(0.5),
				"string_key": "my-string",
				"bool_key":   true,
				"map_key": map[string]interface{}{
					"foo": "bar",
				},
				"list_of_primitives_key": []interface{}{
					"element_1",
					"element_2",
				},
			},
			prevSpec: map[string]interface{}{},
			expected: map[string]interface{}{
				"intKey":    float64(1),
				"floatKey":  float64(0.5),
				"stringKey": "my-string",
				"boolKey":   true,
				"mapKey": map[string]interface{}{
					"foo": "bar",
				},
				"listOfPrimitivesKey": []interface{}{
					"element_1",
					"element_2",
				},
			},
		},
		{
			name: "defaulted zero-value data structures are pruned",
			state: map[string]interface{}{
				"map_key":                map[string]interface{}{},
				"list_of_primitives_key": []interface{}{},
				"nested_object_key":      []interface{}{},
				"list_of_objects_key":    []interface{}{},
			},
			prevSpec: nil,
			expected: nil,
		},
		// handle an edge case where some values are not zero-value by default
		// if these are not read in the resource could be accidentally be re-applied
		// and change state (e.g. false value omitted will be defaulted to true)
		{
			name: "non-string primitives are not pruned if differ from default value",
			state: map[string]interface{}{
				"int_key":    float64(0),
				"float_key":  float64(0.0),
				"string_key": "",
				"bool_key":   false,
			},
			prevSpec: nil,
			expected: map[string]interface{}{
				"intKey":   float64(0),
				"floatKey": float64(0.0),
				"boolKey":  false,
			},
			schemaOverride: map[string]*tfschema.Schema{
				"int_key": {
					Type:     tfschema.TypeInt,
					Optional: true,
					Default:  1,
				},
				"float_key": {
					Type:     tfschema.TypeFloat,
					Optional: true,
					Default:  1,
				},
				"string_key": {
					Type:     tfschema.TypeString,
					Optional: true,
					Default:  "foo",
				},
				"bool_key": {
					Type:     tfschema.TypeBool,
					Optional: true,
					Default:  true,
				},
			},
		},
		{
			name: "lists of objects are set",
			state: map[string]interface{}{
				"list_of_objects_key": []interface{}{
					map[string]interface{}{
						"nested_int_key": float64(1),
					},
				},
			},
			prevSpec: map[string]interface{}{},
			expected: map[string]interface{}{
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"nestedIntKey": float64(1),
					},
				},
			},
		},
		{
			name: "nested objects are converted to maps and set",
			state: map[string]interface{}{
				"nested_object_key": []interface{}{
					map[string]interface{}{
						"nested_float_key": float64(0.5),
					},
				},
			},
			prevSpec: map[string]interface{}{},
			expected: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedFloatKey": float64(0.5),
				},
			},
		},
		{
			name: "individual resource references are preserved",
			state: map[string]interface{}{
				"reference_key": "ref-val",
			},
			prevSpec: map[string]interface{}{
				"referenceRef": map[string]interface{}{
					"name": "my-reference",
				},
			},
			expected: map[string]interface{}{
				"referenceRef": map[string]interface{}{
					"name": "my-reference",
				},
			},
		},
		{
			name: "lists of resource references are preserved",
			state: map[string]interface{}{
				"list_of_references_key": []interface{}{
					"ref1",
					"ref2",
				},
			},
			prevSpec: map[string]interface{}{
				"listOfReferencesKey": []interface{}{
					map[string]interface{}{
						"name": "my-reference1",
					},
					map[string]interface{}{
						"name": "my-reference2",
					},
				},
			},
			expected: map[string]interface{}{
				"listOfReferencesKey": []interface{}{
					map[string]interface{}{
						"name": "my-reference1",
					},
					map[string]interface{}{
						"name": "my-reference2",
					},
				},
			},
		},
		{
			name: "resource references nested in lists of objects are preserved",
			state: map[string]interface{}{
				"list_of_objects_key": []interface{}{
					map[string]interface{}{
						"reference_nested_in_list_of_objects_key": "ref-val1",
					},
					map[string]interface{}{
						"reference_nested_in_list_of_objects_key": "ref-val2",
					},
				},
			},
			prevSpec: map[string]interface{}{
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"nestedInListOfObjectsRef": map[string]interface{}{
							"name": "my-reference1",
						},
					},
					map[string]interface{}{
						"nestedInListOfObjectsRef": map[string]interface{}{
							"name": "my-reference2",
						},
					},
				},
			},
			expected: map[string]interface{}{
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"nestedInListOfObjectsRef": map[string]interface{}{
							"name": "my-reference1",
						},
					},
					map[string]interface{}{
						"nestedInListOfObjectsRef": map[string]interface{}{
							"name": "my-reference2",
						},
					},
				},
			},
		},
		{
			name: "external resource references are preserved",
			state: map[string]interface{}{
				"reference_key": "ref-val",
			},
			prevSpec: map[string]interface{}{
				"referenceRef": map[string]interface{}{
					"external": "my-reference",
				},
			},
			expected: map[string]interface{}{
				"referenceRef": map[string]interface{}{
					"external": "my-reference",
				},
			},
		},
		{
			name: "external resource reference set if no reference defined by spec",
			state: map[string]interface{}{
				"reference_key": "ref-val",
			},
			prevSpec: map[string]interface{}{},
			expected: map[string]interface{}{
				"referenceRef": map[string]interface{}{
					"external": "ref-val",
				},
			},
		},
		{
			name: "list of external resource references set if no list defined by spec",
			state: map[string]interface{}{
				"list_of_references_key": []interface{}{
					"ref-val-1",
					"ref-val-2",
				},
			},
			prevSpec: map[string]interface{}{},
			expected: map[string]interface{}{
				"listOfReferencesKey": []interface{}{
					map[string]interface{}{
						"external": "ref-val-1",
					},
					map[string]interface{}{
						"external": "ref-val-2",
					},
				},
			},
		},
		{
			name: "set of external resource references with complex key set if no set defined by spec",
			state: map[string]interface{}{
				"complex_set_of_references_key": []interface{}{
					"ref-val-1",
					"ref-val-2",
				},
			},
			prevSpec: map[string]interface{}{},
			expected: map[string]interface{}{
				"complexSetOfReferencesKey": []interface{}{
					map[string]interface{}{
						"subKeyRef": map[string]interface{}{
							"external": "ref-val-1",
						},
					},
					map[string]interface{}{
						"subKeyRef": map[string]interface{}{
							"external": "ref-val-2",
						},
					},
				},
			},
		},
		{
			name: "spec-defined values are preserved",
			state: map[string]interface{}{
				"string_key": "fully-expanded-string-value",
			},
			prevSpec: map[string]interface{}{
				"stringKey": "short-string-val",
			},
			expected: map[string]interface{}{
				"stringKey": "short-string-val",
			},
		},
		{
			name: "primitive set ordering is kept consistent",
			state: map[string]interface{}{
				"primitive_set_key": []interface{}{
					"a",
					"b",
					"c",
					"d",
				},
			},
			prevSpec: map[string]interface{}{
				"primitiveSetKey": []interface{}{
					"b",
					"a",
					"c",
				},
			},
			expected: map[string]interface{}{
				"primitiveSetKey": []interface{}{
					"b",
					"a",
					"c",
					"d",
				},
			},
		},
		{
			name: "object set ordering is kept consistent",
			state: map[string]interface{}{
				"object_set_key": []interface{}{
					map[string]interface{}{
						"index": float64(0),
					},
					map[string]interface{}{
						"index": float64(1),
					},
					map[string]interface{}{
						"index": float64(3),
					},
				},
			},
			prevSpec: map[string]interface{}{
				"objectSetKey": []interface{}{
					map[string]interface{}{
						"index": float64(1),
					},
					map[string]interface{}{
						"index": float64(0),
					},
				},
			},
			expected: map[string]interface{}{
				"objectSetKey": []interface{}{
					map[string]interface{}{
						"index": float64(1),
					},
					map[string]interface{}{
						"index": float64(0),
					},
					map[string]interface{}{
						"index": float64(3),
					},
				},
			},
		},
		{
			name: "defaulting is applied to the correct object in the set",
			state: map[string]interface{}{
				"object_set_key": []interface{}{
					map[string]interface{}{
						"index":           float64(0),
						"nested_bool_key": true,
					},
					map[string]interface{}{
						"index":           float64(1),
						"nested_bool_key": false,
					},
				},
			},
			prevSpec: map[string]interface{}{
				"objectSetKey": []interface{}{
					map[string]interface{}{
						"index": float64(1),
					},
					map[string]interface{}{
						"index": float64(0),
					},
				},
			},
			expected: map[string]interface{}{
				"objectSetKey": []interface{}{
					map[string]interface{}{
						"index": float64(1),
					},
					map[string]interface{}{
						"index":         float64(0),
						"nestedBoolKey": true,
					},
				},
			},
		},
		{
			name: "defaulting is applied to correct complex resource reference type",
			state: map[string]interface{}{
				"complex_reference_key": "ref-val",
			},
			prevSpec: map[string]interface{}{},
			expected: map[string]interface{}{
				"complexReferenceKey": map[string]interface{}{
					"value": "ref-val",
				},
			},
		},
		{
			name: "parent values are filtered out of result if resource only supports container annotations",
			rc: &corekccv1alpha1.ResourceConfig{
				Containers: []corekccv1alpha1.Container{
					{
						Type:    corekccv1alpha1.ContainerTypeFolder,
						TFField: "parent_key",
					},
				},
			},
			state: map[string]interface{}{
				"parent_key": "project-id-from-tf-state",
				"string_key": "string-val",
			},
			prevSpec: map[string]interface{}{
				"stringKey": "string-val",
			},
			expected: map[string]interface{}{
				"stringKey": "string-val",
			},
		},
		{
			name: "parent values are set as external hierarchical references if resource supports hierarchical references",
			rc: &corekccv1alpha1.ResourceConfig{
				Containers: []corekccv1alpha1.Container{
					{
						Type:    corekccv1alpha1.ContainerTypeProject,
						TFField: "parent_key",
					},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					{
						Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
						Key:  "projectRef",
					},
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "parent_key",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "projectRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
						},
					},
				},
			},
			state: map[string]interface{}{
				"parent_key": "project-id-from-tf-state",
				"string_key": "string-val",
			},
			prevSpec: map[string]interface{}{
				"stringKey": "string-val",
			},
			expected: map[string]interface{}{
				"stringKey": "string-val",
				"projectRef": map[string]interface{}{
					"external": "project-id-from-tf-state",
				},
			},
		},
		{
			name: "parent values are set as external hierarchical references if resource only supports hierarchical references",
			rc: &corekccv1alpha1.ResourceConfig{
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					{
						Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
						Key:  "projectRef",
					},
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "parent_key",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "projectRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
						},
					},
				},
			},
			state: map[string]interface{}{
				"parent_key": "project-id-from-tf-state",
				"string_key": "string-val",
			},
			prevSpec: map[string]interface{}{
				"stringKey": "string-val",
			},
			expected: map[string]interface{}{
				"stringKey": "string-val",
				"projectRef": map[string]interface{}{
					"external": "project-id-from-tf-state",
				},
			},
		},
		{
			name: "hierarchical references are preserved",
			rc: &corekccv1alpha1.ResourceConfig{
				Containers: []corekccv1alpha1.Container{
					{
						Type:    corekccv1alpha1.ContainerTypeProject,
						TFField: "parent_key",
					},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					{
						Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
						Key:  "projectRef",
					},
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "parent_key",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "projectRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
						},
					},
				},
			},
			state: map[string]interface{}{
				"parent_key": "project-id-from-tf-state",
				"string_key": "string-val",
			},
			prevSpec: map[string]interface{}{
				"stringKey": "string-val",
				"projectRef": map[string]interface{}{
					"name": "my-ref",
				},
			},
			expected: map[string]interface{}{
				"stringKey": "string-val",
				"projectRef": map[string]interface{}{
					"name": "my-ref",
				},
			},
		},
		{
			name: "external hierarchical references are preserved",
			rc: &corekccv1alpha1.ResourceConfig{
				Containers: []corekccv1alpha1.Container{
					{
						Type:    corekccv1alpha1.ContainerTypeProject,
						TFField: "parent_key",
					},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					{
						Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
						Key:  "projectRef",
					},
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "parent_key",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "projectRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
						},
					},
				},
			},
			state: map[string]interface{}{
				"parent_key": "project-id-from-tf-state",
				"string_key": "string-val",
			},
			prevSpec: map[string]interface{}{
				"stringKey": "string-val",
				"projectRef": map[string]interface{}{
					"external": "my-ref",
				},
			},
			expected: map[string]interface{}{
				"stringKey": "string-val",
				"projectRef": map[string]interface{}{
					"external": "my-ref",
				},
			},
		},
		{
			name: "sensitive fields with simple values are preserved",
			state: map[string]interface{}{
				"sensitive_field_key": "val",
			},
			prevSpec: map[string]interface{}{
				"sensitiveFieldKey": map[string]interface{}{
					"value": "old-val",
				},
			},
			expected: map[string]interface{}{
				"sensitiveFieldKey": map[string]interface{}{
					"value": "old-val",
				},
			},
		},
		{
			name: "sensitive fields with values from secret refs are preserved",
			state: map[string]interface{}{
				"sensitive_field_key": "val",
			},
			prevSpec: map[string]interface{}{
				"sensitiveFieldKey": map[string]interface{}{
					"valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{
							"name": "secret1",
							"key":  "key1",
						},
					},
				},
			},
			expected: map[string]interface{}{
				"sensitiveFieldKey": map[string]interface{}{
					"valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{
							"name": "secret1",
							"key":  "key1",
						},
					},
				},
			},
		},
		{
			name: "sensitive fields nested in lists of objects are preserved",
			state: map[string]interface{}{
				"list_of_objects_key": []interface{}{
					map[string]interface{}{
						"sensitive_field_nested_in_list_of_objects_key": "val1",
					},
					map[string]interface{}{
						"sensitive_field_nested_in_list_of_objects_key": "val2",
					},
				},
			},
			prevSpec: map[string]interface{}{
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret1",
									"key":  "key1",
								},
							},
						},
					},
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret2",
									"key":  "key2",
								},
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret1",
									"key":  "key1",
								},
							},
						},
					},
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret2",
									"key":  "key2",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "sensitive fields nested in objects are preserved",
			state: map[string]interface{}{
				"nested_object_key": []interface{}{
					map[string]interface{}{
						"nested_sensitive_field_key": "val",
					},
				},
			},
			prevSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedSensitiveFieldKey": map[string]interface{}{
						"valueFrom": map[string]interface{}{
							"secretKeyRef": map[string]interface{}{
								"name": "secret1",
								"key":  "key1",
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedSensitiveFieldKey": map[string]interface{}{
						"valueFrom": map[string]interface{}{
							"secretKeyRef": map[string]interface{}{
								"name": "secret1",
								"key":  "key1",
							},
						},
					},
				},
			},
		},
		{
			name: "sensitive fields set with simple value if not specified",
			state: map[string]interface{}{
				"sensitive_field_key": "val",
			},
			prevSpec: map[string]interface{}{},
			expected: map[string]interface{}{
				"sensitiveFieldKey": map[string]interface{}{
					"value": "val",
				},
			},
		},
		{
			name: "sensitive fields nested in lists of objects set with simple value if not specified",
			state: map[string]interface{}{
				"list_of_objects_key": []interface{}{
					map[string]interface{}{
						"sensitive_field_nested_in_list_of_objects_key": "val1",
					},
					map[string]interface{}{
						"sensitive_field_nested_in_list_of_objects_key": "val2",
					},
				},
			},
			prevSpec: map[string]interface{}{},
			expected: map[string]interface{}{
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"value": "val1",
						},
					},
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"value": "val2",
						},
					},
				},
			},
		},
		{
			name: "sensitive fields nested in objects set with simple value if not specified",
			state: map[string]interface{}{
				"nested_object_key": []interface{}{
					map[string]interface{}{
						"nested_sensitive_field_key": "val",
					},
				},
			},
			prevSpec: map[string]interface{}{},
			expected: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedSensitiveFieldKey": map[string]interface{}{
						"value": "val",
					},
				},
			},
		},
		{
			name: "maps are treated as atomic when specified by user",
			state: map[string]interface{}{
				"map_key": map[string]interface{}{
					"foo": "bar",
					"baz": "abc",
				},
			},
			prevSpec: map[string]interface{}{
				"mapKey": map[string]interface{}{
					"foo": "bar",
				},
			},
			expected: map[string]interface{}{
				"mapKey": map[string]interface{}{
					"foo": "bar",
				},
			},
		},
		// Tests surrounding managed fields
		{
			name: "values are sourced from live state when not in managed fields set",
			state: map[string]interface{}{
				"int_key":    float64(2),
				"float_key":  float64(1.5),
				"string_key": "external",
				"bool_key":   true,
				"nested_object_key": []interface{}{
					map[string]interface{}{
						"nested_float_key": float64(1.5),
					},
				},
			},
			prevSpec: map[string]interface{}{
				"intKey":    float64(1),
				"floatKey":  float64(0.5),
				"stringKey": "k8s",
				"boolKey":   false,
				"nestedObjectKey": map[string]interface{}{
					"nestedFloatKey": float64(0.5),
				},
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:unrelated": emptyObject,
			}),
			expected: map[string]interface{}{
				"intKey":    float64(2),
				"floatKey":  float64(1.5),
				"stringKey": "external",
				"boolKey":   true,
				"nestedObjectKey": map[string]interface{}{
					"nestedFloatKey": float64(1.5),
				},
			},
		},
		{
			name: "values are sourced from spec when in managed fields set",
			state: map[string]interface{}{
				"int_key":    float64(2),
				"float_key":  float64(1.5),
				"string_key": "external",
				"bool_key":   true,
				"nested_object_key": []interface{}{
					map[string]interface{}{
						"nested_float_key": float64(1.5),
					},
				},
			},
			prevSpec: map[string]interface{}{
				"intKey":    float64(1),
				"floatKey":  float64(0.5),
				"stringKey": "k8s",
				"boolKey":   false,
				"nestedObjectKey": map[string]interface{}{
					"nestedFloatKey": float64(0.5),
				},
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:intKey":    emptyObject,
				"f:floatKey":  emptyObject,
				"f:stringKey": emptyObject,
				"f:boolKey":   emptyObject,
				"f:nestedObjectKey": map[string]interface{}{
					"f:nestedFloatKey": emptyObject,
				},
			}),
			expected: map[string]interface{}{
				"intKey":    float64(1),
				"floatKey":  float64(0.5),
				"stringKey": "k8s",
				"boolKey":   false,
				"nestedObjectKey": map[string]interface{}{
					"nestedFloatKey": float64(0.5),
				},
			},
		},
		{
			name: "maps are treated as atomic when k8s-managed",
			state: map[string]interface{}{
				"map_key": map[string]interface{}{
					"foo": "bar",
					"baz": "abc",
				},
			},
			prevSpec: map[string]interface{}{
				"mapKey": map[string]interface{}{
					"foo": "bar",
				},
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:mapKey": map[string]interface{}{
					"f:foo": emptyObject,
				},
			}),
			expected: map[string]interface{}{
				"mapKey": map[string]interface{}{
					"foo": "bar",
				},
			},
		},
		// TODO(kcc-eng): The following list behavior is required today to keep
		//  consistent with the existing behavior that defaults values in lists.
		//  This will be modified to be more advanced as part of the externally-
		//  managed list merging implementation.
		{
			name: "values in lists of objects ignore managed fields",
			state: map[string]interface{}{
				"list_of_objects_key": []interface{}{
					map[string]interface{}{
						"nested_int_key": float64(1),
					},
				},
			},
			prevSpec: map[string]interface{}{
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"nestedIntKey": float64(2),
					},
				},
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:unrelated": emptyObject,
			}),
			expected: map[string]interface{}{
				// reflects the traditional fully-k8s-managed overlay of
				// the spec list on the live state list
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"nestedIntKey": float64(2),
					},
				},
			},
		},
		{
			name: "values in primitive lists ignore managed fields",
			state: map[string]interface{}{
				"list_of_primitives_key": []interface{}{
					"element_1",
					"element_2",
				},
			},
			prevSpec: map[string]interface{}{
				"listOfPrimitivesKey": []interface{}{
					"element1",
				},
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:listOfPrimitivesKey": emptyObject,
			}),
			expected: map[string]interface{}{
				// reflects solely the live state
				"listOfPrimitivesKey": []interface{}{
					"element_1",
					"element_2",
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := resourceSkeleton()
			if tc.rc != nil {
				r.ResourceConfig = *tc.rc
			}
			if tc.schemaOverride != nil {
				r.TFResource.Schema = tc.schemaOverride
			}
			r.SetNamespace(test.Namespace)
			actual, _ := ConvertTFObjToKCCObj(tc.state, tc.prevSpec, r.TFResource.Schema, &r.ResourceConfig, "", tc.managedFields)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Fatalf("expected: %v, actual: %v", tc.expected, actual)
			}
		})

	}
}

func TestGetLabelsFromState(t *testing.T) {
	tests := []struct {
		name         string
		rc           *corekccv1alpha1.ResourceConfig
		tfAttributes map[string]string
		expected     map[string]string
	}{
		{
			name: "empty labels should resolve",
			rc: &corekccv1alpha1.ResourceConfig{
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Labels: "map_key",
				},
			},
			tfAttributes: map[string]string{
				"map_key.%": "0",
			},
			expected: map[string]string{},
		},
		{
			name: "simple labels should resolve",
			rc: &corekccv1alpha1.ResourceConfig{
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Labels: "map_key",
				},
			},
			tfAttributes: map[string]string{
				"map_key.%":    "2",
				"map_key.key1": "val1",
				"map_key.key2": "val2",
			},
			expected: map[string]string{
				"key1": "val1",
				"key2": "val2",
			},
		},
		{
			name: "nested labels should resolve",
			rc: &corekccv1alpha1.ResourceConfig{
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Labels: "nested_object_key.nested_map_key",
				},
			},
			tfAttributes: map[string]string{
				"nested_object_key.#":                     "1",
				"nested_object_key.0.nested_map_key.%":    "2",
				"nested_object_key.0.nested_map_key.key1": "val1",
				"nested_object_key.0.nested_map_key.key2": "val2",
			},
			expected: map[string]string{
				"key1": "val1",
				"key2": "val2",
			},
		},
		{
			name: "nested labels with nil should resolve",
			rc: &corekccv1alpha1.ResourceConfig{
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Labels: "nested_object_key.nested_map_key",
				},
			},
			tfAttributes: map[string]string{
				"nested_object_key.#": "1",
			},
			expected: map[string]string{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			r := resourceSkeleton()
			if tc.rc != nil {
				r.ResourceConfig = *tc.rc
			}
			rawState := terraform.InstanceState{
				Attributes: tc.tfAttributes,
			}
			labels := GetLabelsFromState(r, &rawState)
			if !reflect.DeepEqual(tc.expected, labels) {
				t.Fatalf("expected: %v, actual: %v", tc.expected, labels)
			}
		})
	}
}

func TestResolveSpecAndStatusWithResourceID_WithDesiredStateInSpecAndObservedStateInStatus(t *testing.T) {
	tests := []struct {
		name           string
		rc             *corekccv1alpha1.ResourceConfig
		metadataName   string
		prevSpec       map[string]interface{}
		prevStatus     map[string]interface{}
		tfResource     *tfschema.Resource
		tfAttributes   map[string]string
		expectedSpec   map[string]interface{}
		expectedStatus map[string]interface{}
		managedFields  *fieldpath.Set
	}{
		{
			name: "only persist specified fields in spec",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "test_field",
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					{
						Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
						Key:  "projectRef",
					},
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "parent_key",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "projectRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
						},
					},
				},
			},
			prevSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"intKey":     int64(1),
				"floatKey":   0.5,
				"projectRef": map[string]interface{}{
					"name": "my-ref",
				},
				"sensitiveFieldKey": map[string]interface{}{
					"valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{
							"name": "secret1",
							"key":  "key1",
						},
					},
				},
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"parent_key": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"int_key": {
						Type:     tfschema.TypeInt,
						Optional: true,
					},
					"float_key": {
						Type:     tfschema.TypeFloat,
						Optional: true,
					},
					"bool_key": {
						Type:     tfschema.TypeBool,
						Optional: true,
					},
					"sensitive_field_key": {
						Type:      tfschema.TypeString,
						Required:  true,
						Sensitive: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":          "resource-id",
				"int_key":             "1",
				"float_key":           "0.5",
				"bool_key":            "false",
				"parent_key":          "project-id-from-tf-state",
				"sensitive_field_key": "val",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"intKey":     int64(1),
				"floatKey":   0.5,
				"projectRef": map[string]interface{}{
					"name": "my-ref",
				},
				"sensitiveFieldKey": map[string]interface{}{
					"valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{
							"name": "secret1",
							"key":  "key1",
						},
					},
				},
			},
			expectedStatus: nil,
		},
		{
			name:       "observed state for output-only fields are persisted in status",
			prevSpec:   map[string]interface{}{},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"status_field": {
						Type:     tfschema.TypeString,
						Computed: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"status_field": "strVal",
			},
			expectedSpec: map[string]interface{}{},
			expectedStatus: map[string]interface{}{
				"statusField": "strVal",
			},
		},
		{
			name: "persist desired state in spec and output-only observed state in status",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			prevSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"intKey":     int64(1),
				"floatKey":   0.5,
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"int_key": {
						Type:     tfschema.TypeInt,
						Optional: true,
					},
					"float_key": {
						Type:     tfschema.TypeFloat,
						Optional: true,
					},
					"bool_key": {
						Type:     tfschema.TypeBool,
						Optional: true,
					},
					"status_field": {
						Type:     tfschema.TypeString,
						Computed: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":   "resource-id",
				"int_key":      "1",
				"float_key":    "0.5",
				"bool_key":     "false",
				"status_field": "strVal",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"intKey":     int64(1),
				"floatKey":   0.5,
			},
			expectedStatus: map[string]interface{}{
				"statusField": "strVal",
			},
		},
		{
			name: "only persist specified nested fields in spec",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			prevSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"nestedObjectKey": map[string]interface{}{
					"nestedKey1": "val1",
				},
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"nested_object_key": {
						Type:     tfschema.TypeList,
						MaxItems: 1,
						Optional: true,
						Elem: &tfschema.Resource{
							Schema: map[string]*tfschema.Schema{
								"nested_key1": {
									Type:     tfschema.TypeString,
									Optional: true,
								},
								"nested_key2": {
									Type:     tfschema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":                      "resource-id",
				"nested_object_key.#":             "1",
				"nested_object_key.0.nested_key1": "val1",
				"nested_object_key.0.nested_key2": "val2",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"nestedObjectKey": map[string]interface{}{
					"nestedKey1": "val1",
				},
			},
		},
		{
			name: "preserve lists of objects unmodified in spec if specified",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			prevSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"field1": 0.5,
						"field2": "strval1",
					},
					map[string]interface{}{
						"field1": 0.7,
					},
				},
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"list_of_objects_key": {
						Type:     tfschema.TypeList,
						Optional: true,
						Elem: &tfschema.Resource{
							Schema: map[string]*tfschema.Schema{
								"field1": {
									Type:     tfschema.TypeFloat,
									Optional: true,
								},
								"field2": {
									Type:     tfschema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":                   "resource-id",
				"list_of_objects_key.#":        "2",
				"list_of_objects_key.0.field1": "0.5",
				"list_of_objects_key.0.field2": "strval1",
				"list_of_objects_key.1.field1": "0.7",
				"list_of_objects_key.1.field2": "strval2",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"field1": 0.5,
						"field2": "strval1",
					},
					map[string]interface{}{
						"field1": 0.7,
					},
				},
			},
		},
		{
			name: "primitive lists are preserved with specified values",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			prevSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"listOfPrimitivesKey": []interface{}{
					"element_1",
				},
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"list_of_primitives_key": {
						Type:     tfschema.TypeList,
						Optional: true,
						Elem: &tfschema.Schema{
							Type: tfschema.TypeString,
						},
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":               "resource-id",
				"list_of_primitives_key.#": "2",
				"list_of_primitives_key.0": "element_1",
				"list_of_primitives_key.1": "element_2",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"listOfPrimitivesKey": []interface{}{
					"element_1",
				},
			},
		},
		{
			name: "server-generated id is retrieved from state and persisted",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				ServerGeneratedIDField: "test_field",
			},
			prevSpec:   map[string]interface{}{},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Computed: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "new-server-generated-id",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "new-server-generated-id",
			},
			expectedStatus: map[string]interface{}{
				"testField": "new-server-generated-id",
			},
		},
		{
			name: "fields in spec are persisted even if they not in managed fields set",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			prevSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"intKey":     int64(1),
				"floatKey":   0.5,
				"boolKey":    false,
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"int_key": {
						Type:     tfschema.TypeInt,
						Optional: true,
					},
					"float_key": {
						Type:     tfschema.TypeFloat,
						Optional: true,
					},
					"bool_key": {
						Type:     tfschema.TypeBool,
						Optional: true,
					},
					"list_of_objects_key": {
						Type:     tfschema.TypeList,
						Optional: true,
						Elem: &tfschema.Resource{
							Schema: map[string]*tfschema.Schema{
								"field1": {
									Type:     tfschema.TypeFloat,
									Optional: true,
								},
								"field2": {
									Type:     tfschema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":                   "resource-id",
				"int_key":                      "1",
				"float_key":                    "0.5",
				"bool_key":                     "false",
				"list_of_objects_key.#":        "2",
				"list_of_objects_key.0.field1": "0.5",
				"list_of_objects_key.0.field2": "strval1",
				"list_of_objects_key.1.field1": "0.7",
				"list_of_objects_key.1.field2": "strval2",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"intKey":     int64(1),
				"floatKey":   0.5,
				"boolKey":    false,
			},
			managedFields: testk8s.MapToFieldPathSet(t, map[string]interface{}{
				"f:unrelated": emptyObject,
			}),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			r := resourceSkeleton()
			if tc.metadataName != "" {
				r.SetName(tc.metadataName)
			}
			r.Spec = tc.prevSpec
			r.Status = tc.prevStatus
			r.TFResource = tc.tfResource
			r.ManagedFields = tc.managedFields
			if tc.rc != nil {
				r.ResourceConfig = *tc.rc
			}
			state := terraform.InstanceState{
				Attributes: tc.tfAttributes,
			}
			k8s.SetAnnotation(k8s.StateIntoSpecAnnotation, k8s.StateAbsentInSpec, r)
			spec, status := ResolveSpecAndStatusWithResourceID(r, &state)
			if got, want := spec, tc.expectedSpec; !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
			if got, want := status, tc.expectedStatus; !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestResolveSpecAndStatusWithResourceID(t *testing.T) {
	tests := []struct {
		name           string
		kind           string
		apiVersion     string
		rc             *corekccv1alpha1.ResourceConfig
		metadataName   string
		prevSpec       map[string]interface{}
		prevStatus     map[string]interface{}
		tfResource     *tfschema.Resource
		tfAttributes   map[string]string
		expectedSpec   map[string]interface{}
		expectedStatus map[string]interface{}
	}{
		{
			name: "with existing user-specified resource ID",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			prevSpec: map[string]interface{}{
				"resourceID": "resource-id",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "resource-id",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "resource-id",
			},
			expectedStatus: nil,
		},
		{
			name: "with empty user-specified resource ID",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			prevSpec: map[string]interface{}{
				"resourceID": "",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "metadata-name-value",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "",
			},
			expectedStatus: nil,
		},
		{
			name: "with user-specified resource ID unset and metadata.name set",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			metadataName: "default-id",
			prevSpec:     map[string]interface{}{},
			prevStatus:   map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "metadata-name-value",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "default-id",
			},
			expectedStatus: nil,
		},
		{
			name: "specifying server-generated resource ID for the first time",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				ServerGeneratedIDField: "test_field",
			},
			prevSpec:   map[string]interface{}{},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Computed: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "new-server-generated-id",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "new-server-generated-id",
			},
			expectedStatus: map[string]interface{}{
				"testField": "new-server-generated-id",
			},
		},
		{
			name: "specifying server-generated resource ID in the observed " +
				"state for the first time",
			kind:       "TestKind",
			apiVersion: "test.cnrm.cloud.google.com/v1beta1",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				ServerGeneratedIDField: "test_field",
			},
			prevSpec:   map[string]interface{}{},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Computed: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "new-server-generated-id",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "new-server-generated-id",
			},
			expectedStatus: map[string]interface{}{
				"observedState": map[string]interface{}{
					"testField": "new-server-generated-id",
				},
			},
		},
		{
			name: "specifying server-generated resource ID with a value " +
				"template for the first time",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField:   "test_field",
					ValueTemplate: "id/{{value}}",
				},
				ServerGeneratedIDField: "test_field",
			},
			prevSpec:   map[string]interface{}{},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Computed: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "id/id-with-value-template",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "id-with-value-template",
			},
			expectedStatus: map[string]interface{}{
				"testField": "id/id-with-value-template",
			},
		},
		{
			name: "specifying server-generated resource ID after it is " +
				"supported in resource config",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				ServerGeneratedIDField: "test_field",
			},
			prevSpec: map[string]interface{}{},
			prevStatus: map[string]interface{}{
				"testField": "existing-server-generated-id",
			},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Computed: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "existing-server-generated-id",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "existing-server-generated-id",
			},
			expectedStatus: map[string]interface{}{
				"testField": "existing-server-generated-id",
			},
		},
		{
			name: "with server-generated resource ID already set",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				ServerGeneratedIDField: "test_field",
			},
			prevSpec: map[string]interface{}{
				"resourceID": "existing-server-generated-id",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Computed: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "existing-server-generated-id",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "existing-server-generated-id",
			},
			expectedStatus: map[string]interface{}{
				"testField": "existing-server-generated-id",
			},
		},
		{
			name: "with resource ID not supported",
			rc:   &corekccv1alpha1.ResourceConfig{},
			prevSpec: map[string]interface{}{
				"testField": "testValue",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Optional: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "testValue",
			},
			expectedSpec: map[string]interface{}{
				"testField": "testValue",
			},
			expectedStatus: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			r := resourceSkeleton()
			if tc.kind != "" {
				r.Kind = tc.kind
			}
			if tc.apiVersion != "" {
				r.APIVersion = tc.apiVersion
			}
			if tc.metadataName != "" {
				r.SetName(tc.metadataName)
			}
			r.Spec = tc.prevSpec
			r.Status = tc.prevStatus
			r.TFResource = tc.tfResource
			if tc.rc != nil {
				r.ResourceConfig = *tc.rc
			}
			state := terraform.InstanceState{
				Attributes: tc.tfAttributes,
			}
			spec, status := ResolveSpecAndStatusWithResourceID(r, &state)
			if got, want := spec, tc.expectedSpec; !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
			if got, want := status, tc.expectedStatus; !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestResolveSpecAndStatusWithFieldRenaming(t *testing.T) {
	tests := []struct {
		name           string
		rc             *corekccv1alpha1.ResourceConfig
		tfResource     *tfschema.Resource
		tfAttributes   map[string]string
		expectedSpec   map[string]interface{}
		expectedStatus map[string]interface{}
	}{
		{
			name: "status fields that collide with reserved status fields are renamed",
			rc: &corekccv1alpha1.ResourceConfig{
				Name: "test-tf-resource-name",
			},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"generation": { // computed field maps to KRM status field
						Type:     tfschema.TypeString,
						Computed: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"generation": "testValue1",
			},
			expectedStatus: map[string]interface{}{
				"resourceGeneration": "testValue1",
			},
		},
		{
			name: "spec fields that collide with reserved status fields are not renamed",
			rc: &corekccv1alpha1.ResourceConfig{
				Name: "test-tf-resource-name",
			},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"generation": {
						Type:     tfschema.TypeString,
						Optional: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"generation": "testValue1",
			},
			expectedSpec: map[string]interface{}{
				"generation": "testValue1",
			},
		},
		{
			name: "status fields that collide with reserved status fields are not renamed if resource is in the exclude list",
			rc: &corekccv1alpha1.ResourceConfig{
				Name: "google_storage_default_object_access_control", // this TF resource is in the exclude list
			},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"generation": {
						Type:     tfschema.TypeString,
						Computed: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"generation": "testValue1",
			},
			expectedStatus: map[string]interface{}{
				"generation": "testValue1",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			r := resourceSkeleton()
			r.TFResource = tc.tfResource
			if tc.rc != nil {
				r.ResourceConfig = *tc.rc
			}
			state := terraform.InstanceState{
				Attributes: tc.tfAttributes,
			}
			spec, status := ResolveSpecAndStatus(r, &state)
			t.Logf("spec = %v\nstatus = %v\n", spec, status)
			if got, want := spec, tc.expectedSpec; !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
			if got, want := status, tc.expectedStatus; !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestResolveSpecAndStatusWithResourceIDPanic(t *testing.T) {
	tests := []struct {
		name         string
		rc           *corekccv1alpha1.ResourceConfig
		metadataName string
		prevSpec     map[string]interface{}
		prevStatus   map[string]interface{}
		tfResource   *tfschema.Resource
		tfAttributes map[string]string
	}{
		{
			name: "with user-specified resource ID unset and metadata.name unset",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			prevSpec:   map[string]interface{}{},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "metadata-name-value",
			},
		},
		{
			name: "with server-generated resource ID not found",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "test_field",
				},
				ServerGeneratedIDField: "test_field",
			},
			prevSpec:   map[string]interface{}{},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Computed: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"different_field": "new-server-generated-id",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			r := resourceSkeleton()
			r.Spec = tc.prevSpec
			r.Status = tc.prevStatus
			r.TFResource = tc.tfResource
			if tc.rc != nil {
				r.ResourceConfig = *tc.rc
			}
			state := terraform.InstanceState{
				Attributes: tc.tfAttributes,
			}

			assertGetSpecAndStatusFromStateWithResourceIDPanic(t, r, &state)
		})
	}
}

func assertGetSpecAndStatusFromStateWithResourceIDPanic(t *testing.T, resource *Resource, state *terraform.InstanceState) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("GetSpecAndStatusFromState should have panicked")
		}
	}()
	ResolveSpecAndStatusWithResourceID(resource, state)
}

func TestResolveSpecAndStatusContainingObservedState(t *testing.T) {
	tests := []struct {
		name           string
		rc             *corekccv1alpha1.ResourceConfig
		kind           string
		apiVersion     string
		annotations    map[string]string
		prevSpec       map[string]interface{}
		prevStatus     map[string]interface{}
		tfResource     *tfschema.Resource
		tfAttributes   map[string]string
		expectedSpec   map[string]interface{}
		expectedStatus map[string]interface{}
		shouldPanic    bool
	}{
		{
			name: "with string observed fields and state-into-status absent",
			rc: &corekccv1alpha1.ResourceConfig{
				ObservedFields: &[]string{
					"nested_object_key.nested_float_key",
					"string_key",
				},
			},
			annotations: map[string]string{
				"cnrm.cloud.google.com/state-into-spec": "absent",
			},
			prevSpec: map[string]interface{}{
				"testField": "desired-value",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"string_key": {
						Type:     tfschema.TypeString,
						Optional: true,
					},
					"nested_object_key": {
						Type:     tfschema.TypeList,
						MaxItems: 1,
						Optional: true,
						Elem: &tfschema.Resource{
							Schema: map[string]*tfschema.Schema{
								"nested_float_key": {
									Type:     tfschema.TypeFloat,
									Optional: true,
								},
								"nested_string_key": {
									Type:     tfschema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":                            "desired-value",
				"nested_object_key.#":                   "1",
				"nested_object_key.0.nested_float_key":  "123",
				"nested_object_key.0.nested_string_key": "not-in-observed-state",
				"string_key":                            "test-observed-field",
			},
			expectedSpec: map[string]interface{}{
				"testField": "desired-value",
			},
			expectedStatus: map[string]interface{}{
				"observedState": map[string]interface{}{
					"nestedObjectKey": map[string]interface{}{
						"nestedFloatKey": float64(123),
					},
					"stringKey": "test-observed-field",
				},
			},
		},
		{
			name: "with string observed fields",
			rc: &corekccv1alpha1.ResourceConfig{
				ObservedFields: &[]string{
					"nested_object_key.nested_float_key",
					"string_key",
				},
			},
			prevSpec: map[string]interface{}{
				"testField": "desired-value",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"string_key": {
						Type:     tfschema.TypeString,
						Optional: true,
					},
					"nested_object_key": {
						Type:     tfschema.TypeList,
						MaxItems: 1,
						Optional: true,
						Elem: &tfschema.Resource{
							Schema: map[string]*tfschema.Schema{
								"nested_float_key": {
									Type:     tfschema.TypeFloat,
									Optional: true,
								},
								"nested_string_key": {
									Type:     tfschema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":                            "desired-value",
				"nested_object_key.#":                   "1",
				"nested_object_key.0.nested_float_key":  "123",
				"nested_object_key.0.nested_string_key": "not-in-observed-state",
				"string_key":                            "test-observed-field",
			},
			expectedSpec: map[string]interface{}{
				"testField": "desired-value",
				"nestedObjectKey": map[string]interface{}{
					"nestedFloatKey":  float64(123),
					"nestedStringKey": "not-in-observed-state",
				},
				"stringKey": "test-observed-field",
			},
			expectedStatus: map[string]interface{}{
				"observedState": map[string]interface{}{
					"nestedObjectKey": map[string]interface{}{
						"nestedFloatKey": float64(123),
					},
					"stringKey": "test-observed-field",
				},
			},
		},
		{
			name: "nested observed field not exist in the returned state but " +
				"its parent field and sibling field do",
			rc: &corekccv1alpha1.ResourceConfig{
				ObservedFields: &[]string{
					"nested_object_key.nested_float_key",
				},
			},
			prevSpec: map[string]interface{}{
				"testField": "desired-value",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"nested_object_key": {
						Type:     tfschema.TypeList,
						MaxItems: 1,
						Optional: true,
						Elem: &tfschema.Resource{
							Schema: map[string]*tfschema.Schema{
								"nested_float_key": {
									Type:     tfschema.TypeFloat,
									Optional: true,
								},
								"nested_string_key": {
									Type:     tfschema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":                            "desired-value",
				"nested_object_key.#":                   "1",
				"nested_object_key.0.nested_string_key": "not-in-observed-state",
			},
			expectedSpec: map[string]interface{}{
				"testField": "desired-value",
				"nestedObjectKey": map[string]interface{}{
					"nestedStringKey": "not-in-observed-state",
				},
			},
			expectedStatus: nil,
		},
		{
			name: "nested observed field and its parent not exist in the " +
				"returned state",
			rc: &corekccv1alpha1.ResourceConfig{
				ObservedFields: &[]string{
					"nested_object_key.nested_float_key",
					"string_key",
				},
			},
			prevSpec: map[string]interface{}{
				"testField": "desired-value",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"string_key": {
						Type:     tfschema.TypeString,
						Optional: true,
					},
					"nested_object_key": {
						Type:     tfschema.TypeList,
						MaxItems: 1,
						Optional: true,
						Elem: &tfschema.Resource{
							Schema: map[string]*tfschema.Schema{
								"nested_float_key": {
									Type:     tfschema.TypeFloat,
									Optional: true,
								},
								"nested_string_key": {
									Type:     tfschema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "desired-value",
				"string_key": "test-observed-field",
			},
			expectedSpec: map[string]interface{}{
				"testField": "desired-value",
				"stringKey": "test-observed-field",
			},
			expectedStatus: map[string]interface{}{
				"observedState": map[string]interface{}{
					"stringKey": "test-observed-field",
				},
			},
		},
		{
			name:       "with computed fields under observed state",
			rc:         &corekccv1alpha1.ResourceConfig{},
			kind:       "TestKind",
			apiVersion: "test.cnrm.cloud.google.com/v1beta1",
			prevSpec: map[string]interface{}{
				"requiredField": "desired-value",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"required_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"computed_string": {
						Type:     tfschema.TypeString,
						Computed: true,
					},
					"computed_object": {
						Type:     tfschema.TypeList,
						Computed: true,
						MaxItems: 1,
						Elem: &tfschema.Resource{
							Schema: map[string]*tfschema.Schema{
								"computed_bool": {
									Type:     tfschema.TypeBool,
									Computed: true,
								},
							},
						},
					},
				},
			},
			tfAttributes: map[string]string{
				"required_field":                  "desired-value",
				"computed_string":                 "computed-status",
				"computed_object.#":               "1",
				"computed_object.0.computed_bool": "true",
			},
			expectedSpec: map[string]interface{}{
				"requiredField": "desired-value",
			},
			expectedStatus: map[string]interface{}{
				"observedState": map[string]interface{}{
					"computedString": "computed-status",
					"computedObject": map[string]interface{}{
						"computedBool": true,
					},
				},
			},
		},
		{
			name: "with observed field and computed field under observed state",
			rc: &corekccv1alpha1.ResourceConfig{
				ObservedFields: &[]string{
					"optional_and_computed_field",
				},
			},
			kind:       "TestKind",
			apiVersion: "test.cnrm.cloud.google.com/v1beta1",
			prevSpec: map[string]interface{}{
				"requiredField": "desired-value",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"required_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"optional_and_computed_field": {
						Type:     tfschema.TypeString,
						Optional: true,
						Computed: true,
					},
					"computed_string": {
						Type:     tfschema.TypeString,
						Computed: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"required_field":              "desired-value",
				"optional_and_computed_field": "observed-field",
				"computed_string":             "computed-status",
			},
			expectedSpec: map[string]interface{}{
				"requiredField":            "desired-value",
				"optionalAndComputedField": "observed-field",
			},
			expectedStatus: map[string]interface{}{
				"observedState": map[string]interface{}{
					"computedString":           "computed-status",
					"optionalAndComputedField": "observed-field",
				},
			},
		},
		{
			name: "panic with observed reference fields",
			rc: &corekccv1alpha1.ResourceConfig{
				ObservedFields: &[]string{
					"reference_key",
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					{
						TFField: "reference_key",
						TypeConfig: corekccv1alpha1.TypeConfig{
							Key: "referenceRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
						},
					},
				},
			},
			prevSpec: map[string]interface{}{
				"testField": "desired-value",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"reference_key": {
						Type:     tfschema.TypeString,
						Optional: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":    "desired-value",
				"reference_key": "invalid-observed-field",
			},
			expectedSpec:   nil,
			expectedStatus: nil,
			shouldPanic:    true,
		},
		{
			name: "panic with observed labels fields",
			rc: &corekccv1alpha1.ResourceConfig{
				ObservedFields: &[]string{
					"string_key",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Labels: "string_key",
				},
			},
			prevSpec: map[string]interface{}{
				"testField": "desired-value",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"string_key": {
						Type:     tfschema.TypeString,
						Optional: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field": "desired-value",
				"string_key": "invalid-observed-field",
			},
			expectedSpec:   nil,
			expectedStatus: nil,
			shouldPanic:    true,
		},
		{
			name: "panic with observed user-specified name fields",
			rc: &corekccv1alpha1.ResourceConfig{
				ObservedFields: &[]string{
					"user_specified_name_key",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "user_specified_name_key",
				},
			},
			prevSpec: map[string]interface{}{
				"testField": "desired-value",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"user_specified_name_key": {
						Type:     tfschema.TypeString,
						Required: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":              "desired-value",
				"user_specified_name_key": "invalid-observed-field",
			},
			expectedSpec:   nil,
			expectedStatus: nil,
			shouldPanic:    true,
		},
		{
			name: "panic with observed server-generated name fields",
			rc: &corekccv1alpha1.ResourceConfig{
				ObservedFields: &[]string{
					"server_generated_name_key",
				},
				ServerGeneratedIDField: "server_generated_name_key",
			},
			prevSpec: map[string]interface{}{
				"testField": "desired-value",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"server_generated_name_key": {
						Type:     tfschema.TypeString,
						Optional: true,
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":                "desired-value",
				"server_generated_name_key": "invalid-observed-field",
			},
			expectedSpec:   nil,
			expectedStatus: nil,
			shouldPanic:    true,
		},
		{
			name: "panic with observed fields under array",
			rc: &corekccv1alpha1.ResourceConfig{
				ObservedFields: &[]string{
					"list_of_objects_key.nested_int_key",
				},
			},
			prevSpec: map[string]interface{}{
				"testField": "desired-value",
			},
			prevStatus: map[string]interface{}{},
			tfResource: &tfschema.Resource{
				Schema: map[string]*tfschema.Schema{
					"test_field": {
						Type:     tfschema.TypeString,
						Required: true,
					},
					"list_of_objects_key": {
						Type:     tfschema.TypeList,
						Optional: true,
						Elem: &tfschema.Resource{
							Schema: map[string]*tfschema.Schema{
								"nested_int_key": {
									Type:     tfschema.TypeInt,
									Optional: true,
								},
								"sensitive_field_nested_in_list_of_objects_key": {
									Type:      tfschema.TypeString,
									Optional:  true,
									Sensitive: true,
								},
							},
						},
					},
				},
			},
			tfAttributes: map[string]string{
				"test_field":                           "desired-value",
				"list_of_objects_key.#":                "2",
				"list_of_objects_key.0.nested_int_key": "invalid-observed-field-1",
				"list_of_objects_key.1.nested_int_key": "invalid-observed-field-2",
			},
			expectedSpec:   nil,
			expectedStatus: nil,
			shouldPanic:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			r := resourceSkeleton()
			if tc.kind != "" {
				r.Kind = tc.kind
			}
			if tc.apiVersion != "" {
				r.APIVersion = tc.apiVersion
			}
			if len(tc.annotations) > 0 {
				r.SetAnnotations(tc.annotations)
			}
			r.Spec = tc.prevSpec
			r.Status = tc.prevStatus
			r.TFResource = tc.tfResource
			if tc.rc != nil {
				r.ResourceConfig = *tc.rc
			}
			state := terraform.InstanceState{
				Attributes: tc.tfAttributes,
			}
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic but it should")
					}
				}()
				ResolveSpecAndStatusWithResourceID(r, &state)
				return
			}
			spec, status := ResolveSpecAndStatusWithResourceID(r, &state)
			if got, want := spec, tc.expectedSpec; !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
			if got, want := status, tc.expectedStatus; !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}
