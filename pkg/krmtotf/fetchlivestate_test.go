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
	"encoding/json"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	"github.com/google/go-cmp/cmp"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestWithFieldsPresetForRead(t *testing.T) {
	nowTime := metav1.Now()
	tests := []struct {
		name        string
		imported    map[string]interface{}
		resource    *krmtotf.Resource
		expectedRet map[string]interface{}
	}{
		{
			name: "immutable fields",
			imported: map[string]interface{}{
				"imported_field": "imported_val",
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"primitiveField": "primitive_val",
						"listOfPrimitivesField": []interface{}{
							"list_of_primitives_val_0",
						},
						"mapField": map[string]interface{}{
							"map_key_a": "map_val_a",
						},
						"nestedObjectField": map[string]interface{}{
							"immutableField": "immutable_val",
							"mutableField":   "mutable_val",
						},
						"listOfObjectsField": []interface{}{
							map[string]interface{}{
								"immutableFieldA": "immutable_val_a",
								"immutableFieldB": "immutable_val_b",
							},
							map[string]interface{}{
								"immutableFieldA": "immutable_val_a",
								"mutableField":    "mutable_val",
							},
						},
					},
				},
				TFResource: &tfschema.Resource{
					Schema: map[string]*tfschema.Schema{
						"imported_field": {
							Type:     tfschema.TypeString,
							Optional: true,
						},
						"primitive_field": {
							Type:     tfschema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"list_of_primitives_field": {
							Type:     tfschema.TypeList,
							Optional: true,
							Elem: &tfschema.Schema{
								Type: tfschema.TypeString,
							},
							ForceNew: true,
						},
						"map_field": {
							Type:     tfschema.TypeMap,
							Optional: true,
							ForceNew: true,
						},
						"nested_object_field": {
							Type:     tfschema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &tfschema.Resource{
								Schema: map[string]*tfschema.Schema{
									"immutable_field": {
										Type:     tfschema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"mutable_field": {
										Type:     tfschema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"list_of_objects_field": {
							Type:     tfschema.TypeList,
							Optional: true,
							Elem: &tfschema.Resource{
								Schema: map[string]*tfschema.Schema{
									"immutable_field_a": {
										Type:     tfschema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"immutable_field_b": {
										Type:     tfschema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"mutable_field": {
										Type:     tfschema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			expectedRet: map[string]interface{}{
				"imported_field":  "imported_val",
				"primitive_field": "primitive_val",
				"list_of_primitives_field": []interface{}{
					"list_of_primitives_val_0",
				},
				"map_field": map[string]interface{}{
					"map_key_a": "map_val_a",
				},
				"nested_object_field": []interface{}{
					map[string]interface{}{
						"immutable_field": "immutable_val",
					},
				},
				"list_of_objects_field": []interface{}{
					map[string]interface{}{
						"immutable_field_a": "immutable_val_a",
						"immutable_field_b": "immutable_val_b",
					},
					map[string]interface{}{
						"immutable_field_a": "immutable_val_a",
						"immutable_field_b": "",
					},
				},
			},
		},
		{
			name: "mutable-but-unreadable fields; none set in annotation",
			imported: map[string]interface{}{
				"imported_field": "imported_val",
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							k8s.MutableButUnreadableFieldsAnnotation: `{}`,
						},
					},
				},
				TFResource: &tfschema.Resource{
					Schema: map[string]*tfschema.Schema{
						"imported_field": {
							Type:     tfschema.TypeString,
							Optional: true,
						},
						"primitive_field": {
							Type:     tfschema.TypeString,
							Optional: true,
						},
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{
					MutableButUnreadableFields: []string{
						"primitive_field",
					},
				},
			},
			expectedRet: map[string]interface{}{
				"imported_field": "imported_val",
			},
		},
		{
			name: "mutable-but-unreadable fields",
			imported: map[string]interface{}{
				"imported_field": "imported_val",
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							k8s.MutableButUnreadableFieldsAnnotation: `{"spec":{"primitiveField":"primitive_val","listOfPrimitivesField":["list_of_primitives_val_0"],"mapField":{"map_key_a":"map_val_a"},"nestedObjectField":{"mutableButUnreadableField":"mutable_but_unreadable_val"}}}`,
						},
					},
				},
				TFResource: &tfschema.Resource{
					Schema: map[string]*tfschema.Schema{
						"imported_field": {
							Type:     tfschema.TypeString,
							Optional: true,
						},
						"primitive_field": {
							Type:     tfschema.TypeString,
							Optional: true,
						},
						"list_of_primitives_field": {
							Type:     tfschema.TypeList,
							Optional: true,
							Elem: &tfschema.Schema{
								Type: tfschema.TypeString,
							},
						},
						"map_field": {
							Type:     tfschema.TypeMap,
							Optional: true,
						},
						"nested_object_field": {
							Type:     tfschema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &tfschema.Resource{
								Schema: map[string]*tfschema.Schema{
									"mutable_but_unreadable_field": {
										Type:     tfschema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{
					MutableButUnreadableFields: []string{
						"primitive_field",
						"list_of_primitives_field",
						"map_field",
						"nested_object_field.mutable_but_unreadable_field",
					},
				},
			},
			expectedRet: map[string]interface{}{
				"imported_field":  "imported_val",
				"primitive_field": "primitive_val",
				"list_of_primitives_field": []interface{}{
					"list_of_primitives_val_0",
				},
				"map_field": map[string]interface{}{
					"map_key_a": "map_val_a",
				},
				"nested_object_field": []interface{}{
					map[string]interface{}{
						"mutable_but_unreadable_field": "mutable_but_unreadable_val",
					},
				},
			},
		},
		{
			name: "mutable-but-unreadable fields; annotation values differ from spec values",
			imported: map[string]interface{}{
				"imported_field": "imported_val",
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							k8s.MutableButUnreadableFieldsAnnotation: `{"spec":{"primitiveField":"primitive_val","listOfPrimitivesField":["list_of_primitives_val_0"],"mapField":{"map_key_a":"map_val_a"},"nestedObjectField":{"mutableButUnreadableField":"mutable_but_unreadable_val"}}}`,
						},
					},
					Spec: map[string]interface{}{
						"primitiveField": "primitive_val_from_spec",
						"listOfPrimitivesField": []interface{}{
							"list_of_primitives_val_0_from_spec",
						},
						"mapField": map[string]interface{}{
							"map_key_a": "map_val_a_from_spec",
						},
						"nestedObjectField": map[string]interface{}{
							"mutableButUnreadableField": "mutable_but_unreadable_val_from_spec",
						},
					},
				},
				TFResource: &tfschema.Resource{
					Schema: map[string]*tfschema.Schema{
						"imported_field": {
							Type:     tfschema.TypeString,
							Optional: true,
						},
						"primitive_field": {
							Type:     tfschema.TypeString,
							Optional: true,
						},
						"list_of_primitives_field": {
							Type:     tfschema.TypeList,
							Optional: true,
							Elem: &tfschema.Schema{
								Type: tfschema.TypeString,
							},
						},
						"map_field": {
							Type:     tfschema.TypeMap,
							Optional: true,
						},
						"nested_object_field": {
							Type:     tfschema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &tfschema.Resource{
								Schema: map[string]*tfschema.Schema{
									"mutable_but_unreadable_field": {
										Type:     tfschema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{
					MutableButUnreadableFields: []string{
						"primitive_field",
						"list_of_primitives_field",
						"map_field",
						"nested_object_field.mutable_but_unreadable_field",
					},
				},
			},
			expectedRet: map[string]interface{}{
				"imported_field":  "imported_val",
				"primitive_field": "primitive_val",
				"list_of_primitives_field": []interface{}{
					"list_of_primitives_val_0",
				},
				"map_field": map[string]interface{}{
					"map_key_a": "map_val_a",
				},
				"nested_object_field": []interface{}{
					map[string]interface{}{
						"mutable_but_unreadable_field": "mutable_but_unreadable_val",
					},
				},
			},
		},
		{
			name: "mutable-but-unreadable fields; field in nested object that is set in imported state",
			imported: map[string]interface{}{
				"nested_object": []interface{}{
					map[string]interface{}{
						"field": "val",
					},
				},
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							k8s.MutableButUnreadableFieldsAnnotation: `{"spec":{"nestedObject":{"mutableButUnreadableField":"mutable_but_unreadable_val"}}}`,
						},
					},
				},
				TFResource: &tfschema.Resource{
					Schema: map[string]*tfschema.Schema{
						"nested_object": {
							Type:     tfschema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &tfschema.Resource{
								Schema: map[string]*tfschema.Schema{
									"field": {
										Type:     tfschema.TypeString,
										Optional: true,
									},
									"mutable_but_unreadable_field": {
										Type:     tfschema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{
					MutableButUnreadableFields: []string{
						"nested_object.mutable_but_unreadable_field",
					},
				},
			},
			expectedRet: map[string]interface{}{
				"nested_object": []interface{}{
					map[string]interface{}{
						"field":                        "val",
						"mutable_but_unreadable_field": "mutable_but_unreadable_val",
					},
				},
			},
		},
		{
			name: "mutable-but-unreadable fields; annotation not set (new KCC resource or backwards compatibility scenario); should get mutable-but-unreadable fields from spec",
			imported: map[string]interface{}{
				"imported_field": "imported_val",
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"primitiveField": "primitive_val",
						"listOfPrimitivesField": []interface{}{
							"list_of_primitives_val_0",
						},
						"mapField": map[string]interface{}{
							"map_key_a": "map_val_a",
						},
						"nestedObjectField": map[string]interface{}{
							"mutableButUnreadableField": "mutable_but_unreadable_val",
						},
					},
				},
				TFResource: &tfschema.Resource{
					Schema: map[string]*tfschema.Schema{
						"imported_field": {
							Type:     tfschema.TypeString,
							Optional: true,
						},
						"primitive_field": {
							Type:     tfschema.TypeString,
							Optional: true,
						},
						"list_of_primitives_field": {
							Type:     tfschema.TypeList,
							Optional: true,
							Elem: &tfschema.Schema{
								Type: tfschema.TypeString,
							},
						},
						"map_field": {
							Type:     tfschema.TypeMap,
							Optional: true,
						},
						"nested_object_field": {
							Type:     tfschema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &tfschema.Resource{
								Schema: map[string]*tfschema.Schema{
									"mutable_but_unreadable_field": {
										Type:     tfschema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{
					MutableButUnreadableFields: []string{
						"primitive_field",
						"list_of_primitives_field",
						"map_field",
						"nested_object_field.mutable_but_unreadable_field",
					},
				},
			},
			expectedRet: map[string]interface{}{
				"imported_field":  "imported_val",
				"primitive_field": "primitive_val",
				"list_of_primitives_field": []interface{}{
					"list_of_primitives_val_0",
				},
				"map_field": map[string]interface{}{
					"map_key_a": "map_val_a",
				},
				"nested_object_field": []interface{}{
					map[string]interface{}{
						"mutable_but_unreadable_field": "mutable_but_unreadable_val",
					},
				},
			},
		},
		{
			name: "directives",
			imported: map[string]interface{}{
				"imported_field": "imported_val",
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							k8s.FormatAnnotation("directive-field-b"): "directive_val_b",
						},
					},
				},
				TFResource: &tfschema.Resource{
					Schema: map[string]*tfschema.Schema{
						"imported_field": {
							Type:     tfschema.TypeString,
							Optional: true,
						},
						"directive_field_b": {
							Type:     tfschema.TypeString,
							Optional: true,
						},
						"directive_field_c": {
							Type:     tfschema.TypeString,
							Optional: true,
							Default:  "directive_val_c",
						},
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{
					Directives: []string{
						"directive_field_b",
						"directive_field_c",
					},
				},
			},
			expectedRet: map[string]interface{}{
				"imported_field":    "imported_val",
				"directive_field_b": "directive_val_b",
				"directive_field_c": "directive_val_c",
			},
		},
		{
			name: "status fields",
			imported: map[string]interface{}{
				"imported_field": "imported_val",
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					Status: map[string]interface{}{
						"primitiveField": "val_b",
						"listOfPrimitivesField": []interface{}{
							"list_of_primitives_val_0",
						},
						"mapField": map[string]interface{}{
							"map_key_a": "map_val_a",
						},
						"nestedObjectField": map[string]interface{}{
							"field": "val",
						},
						"listOfObjectsField": []interface{}{
							map[string]interface{}{
								"fieldA": "val_a",
							},
							map[string]interface{}{
								"fieldB": "val_b",
							},
						},
					},
				},
				TFResource: &tfschema.Resource{
					Schema: map[string]*tfschema.Schema{
						"imported_field": {
							Type: tfschema.TypeString,
						},
						"primitive_field": {
							Type: tfschema.TypeString,
						},
						"list_of_primitives_field": {
							Type: tfschema.TypeList,
							Elem: &tfschema.Schema{
								Type: tfschema.TypeString,
							},
						},
						"map_field": {
							Type: tfschema.TypeMap,
						},
						"nested_object_field": {
							Type:     tfschema.TypeList,
							MaxItems: 1,
							Elem: &tfschema.Resource{
								Schema: map[string]*tfschema.Schema{
									"field": {
										Type: tfschema.TypeString,
									},
								},
							},
						},
						"list_of_objects_field": {
							Type: tfschema.TypeList,
							Elem: &tfschema.Resource{
								Schema: map[string]*tfschema.Schema{
									"field_a": {
										Type: tfschema.TypeString,
									},
									"field_b": {
										Type: tfschema.TypeString,
									},
								},
							},
						},
					},
				},
			},
			expectedRet: map[string]interface{}{
				"imported_field":  "imported_val",
				"primitive_field": "val_b",
				"list_of_primitives_field": []interface{}{
					"list_of_primitives_val_0",
				},
				"map_field": map[string]interface{}{
					"map_key_a": "map_val_a",
				},
				"nested_object_field": []interface{}{
					map[string]interface{}{
						"field": "val",
					},
				},
				"list_of_objects_field": []interface{}{
					map[string]interface{}{
						"field_a": "val_a",
					},
					map[string]interface{}{
						"field_b": "val_b",
					},
				},
			},
		},
		{
			name: "computed fields under status.observedState",
			imported: map[string]interface{}{
				"imported_field": "imported_val",
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					TypeMeta: metav1.TypeMeta{
						Kind:       "TestKind",
						APIVersion: "test.cnrm.cloud.google.com/v1beta1",
					},
					Status: map[string]interface{}{
						"observedState": map[string]interface{}{
							"primitiveField": "val_b",
							"listOfPrimitivesField": []interface{}{
								"list_of_primitives_val_0",
							},
							"mapField": map[string]interface{}{
								"map_key_a": "map_val_a",
							},
							"nestedObjectField": map[string]interface{}{
								"field": "val",
							},
							"listOfObjectsField": []interface{}{
								map[string]interface{}{
									"fieldA": "val_a",
								},
								map[string]interface{}{
									"fieldB": "val_b",
								},
							},
						},
					},
				},
				TFResource: &tfschema.Resource{
					Schema: map[string]*tfschema.Schema{
						"imported_field": {
							Type: tfschema.TypeString,
						},
						"primitive_field": {
							Type: tfschema.TypeString,
						},
						"list_of_primitives_field": {
							Type: tfschema.TypeList,
							Elem: &tfschema.Schema{
								Type: tfschema.TypeString,
							},
						},
						"map_field": {
							Type: tfschema.TypeMap,
						},
						"nested_object_field": {
							Type:     tfschema.TypeList,
							MaxItems: 1,
							Elem: &tfschema.Resource{
								Schema: map[string]*tfschema.Schema{
									"field": {
										Type: tfschema.TypeString,
									},
								},
							},
						},
						"list_of_objects_field": {
							Type: tfschema.TypeList,
							Elem: &tfschema.Resource{
								Schema: map[string]*tfschema.Schema{
									"field_a": {
										Type: tfschema.TypeString,
									},
									"field_b": {
										Type: tfschema.TypeString,
									},
								},
							},
						},
					},
				},
			},
			expectedRet: map[string]interface{}{
				"imported_field":  "imported_val",
				"primitive_field": "val_b",
				"list_of_primitives_field": []interface{}{
					"list_of_primitives_val_0",
				},
				"map_field": map[string]interface{}{
					"map_key_a": "map_val_a",
				},
				"nested_object_field": []interface{}{
					map[string]interface{}{
						"field": "val",
					},
				},
				"list_of_objects_field": []interface{}{
					map[string]interface{}{
						"field_a": "val_a",
					},
					map[string]interface{}{
						"field_b": "val_b",
					},
				},
			},
		},
		{
			name: "if the object is marked for deletion, withPresetFieldsForRead should not return an error when a referenced secret does not exist",
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						DeletionTimestamp: &nowTime,
					},
					Spec: map[string]interface{}{
						"primitiveField": "primitive_val",
						"sensitiveField": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret1",
									"key":  "secret-key1",
								},
							},
						},
					},
				},
				TFResource: &tfschema.Resource{
					Schema: map[string]*tfschema.Schema{
						"sensitive_field": {
							Type:      tfschema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"primitive_field": {
							Type:     tfschema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			expectedRet: map[string]interface{}{
				"primitive_field": "primitive_val",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			c := mgr.GetClient()

			testcontroller.EnsureNamespaceExistsT(t, c, testID)
			tc.resource.SetNamespace(testID)
			ret, err := krmtotf.WithFieldsPresetForRead(tc.imported, tc.resource, mgr.GetClient(), nil)
			if err != nil {
				t.Fatal(err)
			}
			if !test.Equals(t, tc.expectedRet, ret) {
				diff := cmp.Diff(tc.expectedRet, ret)
				t.Fatalf("actual result did not match expected result; diff (-expected +actual):\n%v", diff)
			}
		})
	}
}

func TestWithFieldsPresetForReadMutableUnreadableSensitiveFields(t *testing.T) {
	// Variables common across test cases
	importedState := map[string]interface{}{
		"imported_field": "imported_val",
	}
	tfResource := &tfschema.Resource{
		Schema: map[string]*tfschema.Schema{
			"imported_field": {
				Type:     tfschema.TypeString,
				Optional: true,
			},
			"sensitive_field": {
				Type:      tfschema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"nested_object_field": {
				Type:     tfschema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &tfschema.Resource{
					Schema: map[string]*tfschema.Schema{
						"sensitive_field": {
							Type:      tfschema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
					},
				},
			},
		},
	}
	resourceConfig := v1alpha1.ResourceConfig{
		MutableButUnreadableFields: []string{
			"sensitive_field",
			"nested_object_field.sensitive_field",
		},
	}

	type versionStatus int
	const (
		upToDate versionStatus = iota
		outdated
		notFound
	)
	tests := []struct {
		name     string
		imported map[string]interface{}

		// Used to generate an "observed-secret-versions" annotation value for
		// `resource` if specified. If unspecified, the annotation will be left
		// unset. This annotation can't be configured accurately for testing
		// purposes before the Secrets are actually created since that is when
		// the versions of Secrets are generated. Each Secret can be specified
		// as:
		// * UP_TO_DATE: Secret's version in the annotation matches its current version
		// * OUTDATED: Secret's version in the annotation does not match its current version
		// * NOT_FOUND: Secret has a version in the annotation, but the Secret itself does not exist
		observedSecretVersions map[string]versionStatus

		resource          *krmtotf.Resource
		referencedSecrets []*unstructured.Unstructured
		expectedRet       map[string]interface{}
	}{
		{
			name:     "sensitive fields with values from Secrets in mutable-but-unreadable-fields, and observed-secret-versions is up-to-date",
			imported: importedState,
			observedSecretVersions: map[string]versionStatus{
				"secret1": upToDate,
				"secret2": upToDate,
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							k8s.MutableButUnreadableFieldsAnnotation: `{"spec":{"sensitiveField":{"valueFrom":{"secretKeyRef":{"name":"secret1","key":"secret-key1"}}},"nestedObjectField":{"sensitiveField":{"valueFrom":{"secretKeyRef":{"name":"secret2","key":"secret-key2"}}}}}}`,
						},
					},
				},
				TFResource:     tfResource,
				ResourceConfig: resourceConfig,
			},

			referencedSecrets: []*unstructured.Unstructured{
				test.NewSecretUnstructured("secret1", "", map[string]interface{}{"secret-key1": "secret-val1"}),
				test.NewSecretUnstructured("secret2", "", map[string]interface{}{"secret-key2": "secret-val2"}),
			},
			expectedRet: map[string]interface{}{
				"imported_field":  "imported_val",
				"sensitive_field": "secret-val1",
				"nested_object_field": []interface{}{
					map[string]interface{}{
						"sensitive_field": "secret-val2",
					},
				},
			},
		},
		{
			name:     "sensitive fields with values from Secrets in mutable-but-unreadable-fields, but observed-secret-versions is outdated",
			imported: importedState,
			observedSecretVersions: map[string]versionStatus{
				"secret1": outdated,
				"secret2": outdated,
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							k8s.MutableButUnreadableFieldsAnnotation: `{"spec":{"sensitiveField":{"valueFrom":{"secretKeyRef":{"name":"secret1","key":"secret-key1"}}},"nestedObjectField":{"sensitiveField":{"valueFrom":{"secretKeyRef":{"name":"secret2","key":"secret-key2"}}}}}}`,
						},
					},
				},
				TFResource:     tfResource,
				ResourceConfig: resourceConfig,
			},
			referencedSecrets: []*unstructured.Unstructured{
				test.NewSecretUnstructured("secret1", "", map[string]interface{}{"secret-key1": "secret-val1"}),
				test.NewSecretUnstructured("secret2", "", map[string]interface{}{"secret-key2": "secret-val2"}),
			},
			expectedRet: map[string]interface{}{
				"imported_field": "imported_val",
			},
		},
		{
			name:     "sensitive fields with values from Secrets in mutable-but-unreadable-fields, but Secrets are not found",
			imported: importedState,
			observedSecretVersions: map[string]versionStatus{
				"secret1": notFound,
				"secret2": notFound,
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							k8s.MutableButUnreadableFieldsAnnotation: `{"spec":{"sensitiveField":{"valueFrom":{"secretKeyRef":{"name":"secret1","key":"secret-key1"}}},"nestedObjectField":{"sensitiveField":{"valueFrom":{"secretKeyRef":{"name":"secret2","key":"secret-key2"}}}}}}`,
						},
					},
				},
				TFResource:     tfResource,
				ResourceConfig: resourceConfig,
			},
			expectedRet: map[string]interface{}{
				"imported_field": "imported_val",
			},
		},
		{
			name:     "sensitive fields with values from Secrets in mutable-but-unreadable-fields, and observed-secret-versions is up-to-date, but keys can't be found in Secrets",
			imported: importedState,
			observedSecretVersions: map[string]versionStatus{
				"secret1": upToDate,
				"secret2": upToDate,
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							k8s.MutableButUnreadableFieldsAnnotation: `{"spec":{"sensitiveField":{"valueFrom":{"secretKeyRef":{"name":"secret1","key":"secret-key1"}}},"nestedObjectField":{"sensitiveField":{"valueFrom":{"secretKeyRef":{"name":"secret2","key":"secret-key2"}}}}}}`,
						},
					},
				},
				TFResource:     tfResource,
				ResourceConfig: resourceConfig,
			},
			referencedSecrets: []*unstructured.Unstructured{
				test.NewSecretUnstructured("secret1", "", map[string]interface{}{"unused-secret-key1": "secret-val1"}),
				test.NewSecretUnstructured("secret2", "", map[string]interface{}{"unused-secret-key2": "secret-val2"}),
			},
			expectedRet: map[string]interface{}{
				"imported_field": "imported_val",
			},
		},
		{
			name:     "sensitive fields with simple values in mutable-but-unreadable-fields",
			imported: importedState,
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							k8s.MutableButUnreadableFieldsAnnotation: `{"spec":{"sensitiveField":{"value":"sensitive_val"},"nestedObjectField":{"sensitiveField":{"value":"nested_sensitive_val"}}}}`,
						},
					},
				},
				TFResource:     tfResource,
				ResourceConfig: resourceConfig,
			},
			expectedRet: map[string]interface{}{
				"imported_field":  "imported_val",
				"sensitive_field": "sensitive_val",
				"nested_object_field": []interface{}{
					map[string]interface{}{
						"sensitive_field": "nested_sensitive_val",
					},
				},
			},
		},
		{
			name:     "mutable-but-unreadable-fields and observed-secret-versions annotations not set (new KCC resource or backwards compatibility scenario); should get sensitive, mutable, unreadable fields from spec (where sensitive fields are from Secrets)",
			imported: importedState,
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"sensitiveField": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret1",
									"key":  "secret-key1",
								},
							},
						},
						"nestedObjectField": map[string]interface{}{
							"sensitiveField": map[string]interface{}{
								"valueFrom": map[string]interface{}{
									"secretKeyRef": map[string]interface{}{
										"name": "secret2",
										"key":  "secret-key2",
									},
								},
							},
						},
					},
				},
				TFResource:     tfResource,
				ResourceConfig: resourceConfig,
			},
			referencedSecrets: []*unstructured.Unstructured{
				test.NewSecretUnstructured("secret1", "", map[string]interface{}{"secret-key1": "secret-val1"}),
				test.NewSecretUnstructured("secret2", "", map[string]interface{}{"secret-key2": "secret-val2"}),
			},
			expectedRet: map[string]interface{}{
				"imported_field":  "imported_val",
				"sensitive_field": "secret-val1",
				"nested_object_field": []interface{}{
					map[string]interface{}{
						"sensitive_field": "secret-val2",
					},
				},
			},
		},
		{
			name:     "mutable-but-unreadable-fields and observed-secret-versions annotations not set (new KCC resource or backwards compatibility scenario); should get sensitive, mutable, unreadable fields from spec (where sensitive fields are simple values)",
			imported: importedState,
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					Spec: map[string]interface{}{
						"sensitiveField": map[string]interface{}{
							"value": "sensitive_val",
						},
						"nestedObjectField": map[string]interface{}{
							"sensitiveField": map[string]interface{}{
								"value": "nested_sensitive_val",
							},
						},
					},
				},
				TFResource:     tfResource,
				ResourceConfig: resourceConfig,
			},
			expectedRet: map[string]interface{}{
				"imported_field":  "imported_val",
				"sensitive_field": "sensitive_val",
				"nested_object_field": []interface{}{
					map[string]interface{}{
						"sensitive_field": "nested_sensitive_val",
					},
				},
			},
		},
		{
			name:     "mutable-but-unreadable-fields annotation values differ from spec values (where sensitive fields in annotation are from Secrets, but sensitive fields in Spec are simple values)",
			imported: importedState,
			observedSecretVersions: map[string]versionStatus{
				"secret1": upToDate,
				"secret2": upToDate,
			},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							k8s.MutableButUnreadableFieldsAnnotation: `{"spec":{"sensitiveField":{"valueFrom":{"secretKeyRef":{"name":"secret1","key":"secret-key1"}}},"nestedObjectField":{"sensitiveField":{"valueFrom":{"secretKeyRef":{"name":"secret2","key":"secret-key2"}}}}}}`,
						},
					},
					Spec: map[string]interface{}{
						"sensitiveField": map[string]interface{}{
							"value": "sensitive_val",
						},
						"nestedObjectField": map[string]interface{}{
							"sensitiveField": map[string]interface{}{
								"value": "nested_sensitive_val",
							},
						},
					},
				},
				TFResource:     tfResource,
				ResourceConfig: resourceConfig,
			},
			referencedSecrets: []*unstructured.Unstructured{
				test.NewSecretUnstructured("secret1", "", map[string]interface{}{"secret-key1": "secret-val1"}),
				test.NewSecretUnstructured("secret2", "", map[string]interface{}{"secret-key2": "secret-val2"}),
			},
			expectedRet: map[string]interface{}{
				"imported_field":  "imported_val",
				"sensitive_field": "secret-val1",
				"nested_object_field": []interface{}{
					map[string]interface{}{
						"sensitive_field": "secret-val2",
					},
				},
			},
		},
		{
			name:                   "mutable-but-unreadable-fields annotation values differ from spec values (where sensitive fields in annotation are simple values, but sensitive fields in Spec are from Secrets)",
			imported:               importedState,
			observedSecretVersions: map[string]versionStatus{},
			resource: &krmtotf.Resource{
				Resource: k8s.Resource{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							k8s.MutableButUnreadableFieldsAnnotation: `{"spec":{"sensitiveField":{"value":"sensitive_val"},"nestedObjectField":{"sensitiveField":{"value":"nested_sensitive_val"}}}}`,
						},
					},
					Spec: map[string]interface{}{
						"sensitiveField": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret1",
									"key":  "secret-key1",
								},
							},
						},
						"nestedObjectField": map[string]interface{}{
							"sensitiveField": map[string]interface{}{
								"valueFrom": map[string]interface{}{
									"secretKeyRef": map[string]interface{}{
										"name": "secret2",
										"key":  "secret-key2",
									},
								},
							},
						},
					},
				},
				TFResource:     tfResource,
				ResourceConfig: resourceConfig,
			},
			referencedSecrets: []*unstructured.Unstructured{
				test.NewSecretUnstructured("secret1", "", map[string]interface{}{"secret-key1": "secret-val1"}),
				test.NewSecretUnstructured("secret2", "", map[string]interface{}{"secret-key2": "secret-val2"}),
			},
			expectedRet: map[string]interface{}{
				"imported_field":  "imported_val",
				"sensitive_field": "sensitive_val",
				"nested_object_field": []interface{}{
					map[string]interface{}{
						"sensitive_field": "nested_sensitive_val",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			c := mgr.GetClient()

			testcontroller.EnsureNamespaceExistsT(t, c, testID)
			tc.resource.SetNamespace(testID)
			for _, obj := range tc.referencedSecrets {
				obj.SetNamespace(testID)
			}
			test.EnsureObjectsExist(t, tc.referencedSecrets, c)

			// Generate value for observed-secret-versions annotation if needed
			if tc.observedSecretVersions != nil {
				secretVersions := make(map[string]string)
				for secretName, status := range tc.observedSecretVersions {
					if status == notFound {
						secretVersions[secretName] = "12345"
						continue
					}
					version, err := getResourceVersionOfSecret(secretName, testID, c)
					if err != nil {
						t.Fatalf("error determining version of Secret %v: %v", secretName, err)
					}
					switch status {
					case upToDate:
						secretVersions[secretName] = version
					case outdated:
						secretVersions[secretName] = version + "0"
					}
				}
				b, err := json.Marshal(secretVersions)
				if err != nil {
					t.Fatalf("error marshalling secret versions map: %v", err)
				}
				k8s.SetAnnotation(k8s.ObservedSecretVersionsAnnotation, string(b), tc.resource)
			}

			ret, err := krmtotf.WithFieldsPresetForRead(tc.imported, tc.resource, mgr.GetClient(), nil)
			if err != nil {
				t.Fatal(err)
			}
			if !test.Equals(t, tc.expectedRet, ret) {
				diff := cmp.Diff(tc.expectedRet, ret)
				t.Fatalf("actual result did not match expected result; diff (-expected +actual):\n%v", diff)
			}
		})
	}
}
