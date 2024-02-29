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
	"context"
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	. "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"

	"github.com/google/go-cmp/cmp"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var mgr manager.Manager

func resourceSkeleton() *Resource {
	return &Resource{
		TFInfo: &terraform.InstanceInfo{},
		TFResource: &tfschema.Resource{
			Schema: map[string]*tfschema.Schema{
				"int_key": {
					Type:     tfschema.TypeInt,
					Optional: true,
				},
				"float_key": {
					Type:     tfschema.TypeFloat,
					Optional: true,
				},
				"string_key": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
				"bool_key": {
					Type:     tfschema.TypeBool,
					Optional: true,
				},
				"nonconfigurable_string_key": {
					Type: tfschema.TypeString,
				},
				"directive_key": {
					Type: tfschema.TypeBool,
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
							"nested_simple_reference_key": {
								Type:     tfschema.TypeString,
								Optional: true,
							},
							"nested_complex_reference_key": {
								Type:     tfschema.TypeString,
								Optional: true,
							},
							"nested_sensitive_field_key": {
								Type:      tfschema.TypeString,
								Optional:  true,
								Sensitive: true,
							},
							"nested_map_key": {
								Type:     tfschema.TypeMap,
								Optional: true,
							},
							"nested_nonconfigurable_key": {
								Type: tfschema.TypeString,
							},
						},
					},
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
							"reference_nested_in_list_of_objects_key": {
								Type:     tfschema.TypeString,
								Optional: true,
							},
							"sensitive_field_nested_in_list_of_objects_key": {
								Type:      tfschema.TypeString,
								Optional:  true,
								Sensitive: true,
							},
							"nonconfigurable_field_nested_in_list_of_objects_key": {
								Type: tfschema.TypeString,
							},
						},
					},
				},
				"list_of_primitives_key": {
					Type:     tfschema.TypeList,
					Optional: true,
					Elem: &tfschema.Schema{
						Type: tfschema.TypeString,
					},
				},
				"reference_key": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
				"list_of_references_key": {
					Type:     tfschema.TypeList,
					Optional: true,
					Elem: &tfschema.Schema{
						Type: tfschema.TypeString,
					},
				},
				"complex_set_of_references_key": {
					Type:     tfschema.TypeSet,
					Optional: true,
					Elem: &tfschema.Schema{
						Type: tfschema.TypeString,
					},
				},
				"map_key": {
					Type:     tfschema.TypeMap,
					Optional: true,
				},
				"unused_key": {
					Type:     tfschema.TypeSet,
					Optional: true,
				},
				"primitive_set_key": {
					Type:     tfschema.TypeSet,
					Optional: true,
					Elem: &tfschema.Schema{
						Type: tfschema.TypeString,
					},
				},
				"object_set_key": {
					Type:     tfschema.TypeSet,
					Optional: true,
					Elem: &tfschema.Resource{
						Schema: map[string]*tfschema.Schema{
							"index": {
								Type:     tfschema.TypeInt,
								Optional: true,
							},
							"nested_bool_key": {
								Type:     tfschema.TypeBool,
								Optional: true,
							},
						},
					},
					Set: func(val interface{}) int {
						m := val.(map[string]interface{})
						return m["index"].(int)
					},
				},
				"parent_key": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
				"sensitive_field_key": {
					Type:      tfschema.TypeString,
					Optional:  true,
					Sensitive: true,
				},
				"complex_reference_key": {
					Type:     tfschema.TypeString,
					Optional: true,
				},
			},
		},
		ResourceConfig: corekccv1alpha1.ResourceConfig{
			Name: "google_foo",
			Kind: "Foo",
			ResourceReferences: []corekccv1alpha1.ReferenceConfig{
				{
					TFField: "nested_object_key.nested_simple_reference_key",
					TypeConfig: corekccv1alpha1.TypeConfig{
						Key: "nestedRef",
						GVK: k8sschema.GroupVersionKind{
							Group:   "test1.cnrm.cloud.google.com",
							Version: "v1alpha1",
							Kind:    "Test1Bar",
						},
					},
				},
				{
					TFField: "nested_object_key.nested_complex_reference_key",
					Types: []corekccv1alpha1.TypeConfig{
						{
							Key: "barRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
						},
						{
							Key:            "value",
							JSONSchemaType: "string",
						},
					},
				},
				{
					TFField: "list_of_objects_key.reference_nested_in_list_of_objects_key",
					TypeConfig: corekccv1alpha1.TypeConfig{
						Key: "nestedInListOfObjectsRef",
						GVK: k8sschema.GroupVersionKind{
							Group:   "test1.cnrm.cloud.google.com",
							Version: "v1alpha1",
							Kind:    "Test1Bar",
						},
					},
				},
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
				{
					TFField: "list_of_references_key",
					TypeConfig: corekccv1alpha1.TypeConfig{
						GVK: k8sschema.GroupVersionKind{
							Group:   "test1.cnrm.cloud.google.com",
							Version: "v1alpha1",
							Kind:    "Test1Bar",
						},
					},
				},
				{
					TFField: "complex_set_of_references_key",
					Types: []corekccv1alpha1.TypeConfig{
						{
							Key: "subKeyRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
						},
					},
				},
				{
					TFField: "complex_reference_key",
					Types: []corekccv1alpha1.TypeConfig{
						{
							Key:            "value",
							JSONSchemaType: "string",
						},
						{
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
			Directives: []string{
				"directive_key",
			},
		},
	}
}

func TestKRMResourceSpecsToTFConfig(t *testing.T) {
	tests := []struct {
		name                  string
		prevSpec              map[string]interface{}
		hasResourceReferences bool
		hasSecretReferences   bool
		expected              map[string]interface{}
	}{
		{
			name: "keys as constant",
			prevSpec: map[string]interface{}{
				"intKey":    "1",
				"floatKey":  "0.5",
				"stringKey": "StringVal",
				"boolKey":   "true",
			},
			expected: map[string]interface{}{
				"int_key":    1,
				"float_key":  0.5,
				"string_key": "StringVal",
				"bool_key":   true,
			},
		},
		{
			name: "list of primitives key",
			prevSpec: map[string]interface{}{
				"listOfPrimitivesKey": []interface{}{
					"myString1",
					"myString2",
				},
			},
			expected: map[string]interface{}{
				"list_of_primitives_key": []interface{}{
					"myString1",
					"myString2",
				},
			},
		},
		{
			name: "map key",
			prevSpec: map[string]interface{}{
				"mapKey": map[string]interface{}{
					"myMapKey1": "MyMapValue1",
					"myMapKey2": "MyMapValue2",
				},
			},
			expected: map[string]interface{}{
				"map_key": map[string]interface{}{
					"myMapKey1": "MyMapValue1",
					"myMapKey2": "MyMapValue2",
				},
			},
		},
		{
			name:                  "nested objects key",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedFloatKey": "0.5",
					"nestedRef": map[string]interface{}{
						"name": "my-ref1",
					},
					"nestedMapKey": map[string]interface{}{
						"name": "val",
					},
					"nestedComplexReferenceKey": map[string]interface{}{
						"value": "foobar",
					},
				},
			},
			expected: map[string]interface{}{
				"nested_object_key": []interface{}{
					map[string]interface{}{
						"nested_float_key":             0.5,
						"nested_simple_reference_key":  "my-ref1",
						"nested_complex_reference_key": "foobar",
						"nested_map_key": map[string]interface{}{
							"name": "val",
						},
					},
				},
			},
		},
		{
			name:                  "list of objects key",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"nestedIntKey": "2",
						"nestedInListOfObjectsRef": map[string]interface{}{
							"name": "my-ref1",
						},
					},
					map[string]interface{}{
						"nestedIntKey": "3",
						"nestedInListOfObjectsRef": map[string]interface{}{
							"name": "my-ref2",
						},
					},
				},
			},
			expected: map[string]interface{}{
				"list_of_objects_key": []interface{}{
					map[string]interface{}{
						"nested_int_key": 2,
						"reference_nested_in_list_of_objects_key": "my-ref1",
					},
					map[string]interface{}{
						"nested_int_key": 3,
						"reference_nested_in_list_of_objects_key": "my-ref2",
					},
				},
			},
		},
		{
			name:                  "simple reference key",
			hasResourceReferences: true,
			prevSpec: map[string]interface{}{
				"referenceRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			expected: map[string]interface{}{
				"reference_key": "my-ref1",
			},
		},
		{
			name: "sensitive field with simple value",
			prevSpec: map[string]interface{}{
				"sensitiveFieldKey": map[string]interface{}{
					"value": "val",
				},
			},
			expected: map[string]interface{}{
				"sensitive_field_key": "val",
			},
		},
		{
			name:                "sensitive field with value from secret ref",
			hasSecretReferences: true,
			prevSpec: map[string]interface{}{
				"sensitiveFieldKey": map[string]interface{}{
					"valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{
							"name": "secret1",
							"key":  "secret-key1",
						},
					},
				},
			},
			expected: map[string]interface{}{
				"sensitive_field_key": "secret-val1",
			},
		},
		{
			name: "nested sensitive field with simple value",
			prevSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedSensitiveFieldKey": map[string]interface{}{
						"value": "val",
					},
				},
			},
			expected: map[string]interface{}{
				"nested_object_key": []interface{}{
					map[string]interface{}{
						"nested_sensitive_field_key": "val",
					},
				},
			},
		},
		{
			name:                "nested sensitive field with value from secret ref",
			hasSecretReferences: true,
			prevSpec: map[string]interface{}{
				"nestedObjectKey": map[string]interface{}{
					"nestedSensitiveFieldKey": map[string]interface{}{
						"valueFrom": map[string]interface{}{
							"secretKeyRef": map[string]interface{}{
								"name": "secret1",
								"key":  "secret-key1",
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"nested_object_key": []interface{}{
					map[string]interface{}{
						"nested_sensitive_field_key": "secret-val1",
					},
				},
			},
		},
		{
			name: "sensitive field nested in list of objects with simple values",
			prevSpec: map[string]interface{}{
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
			expected: map[string]interface{}{
				"list_of_objects_key": []interface{}{
					map[string]interface{}{
						"sensitive_field_nested_in_list_of_objects_key": "val1",
					},
					map[string]interface{}{
						"sensitive_field_nested_in_list_of_objects_key": "val2",
					},
				},
			},
		},
		{
			name:                "sensitive field nested in list of objects with values from secret refs",
			hasSecretReferences: true,
			prevSpec: map[string]interface{}{
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret1",
									"key":  "secret-key1",
								},
							},
						},
					},
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
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
			expected: map[string]interface{}{
				"list_of_objects_key": []interface{}{
					map[string]interface{}{
						"sensitive_field_nested_in_list_of_objects_key": "secret-val1",
					},
					map[string]interface{}{
						"sensitive_field_nested_in_list_of_objects_key": "secret-val2",
					},
				},
			},
		},
		{
			name: "drop nonconfigurable fields",
			prevSpec: map[string]interface{}{
				"nonconfigurable_string_key": "true",
			},
			expected: map[string]interface{}{},
		},
		{
			name: "drop nonconfigurable nested fields",
			prevSpec: map[string]interface{}{
				"nested_object_key": []interface{}{
					map[string]interface{}{
						"nested_nonconfigurable_key": "value",
					},
				},
			},
			expected: map[string]interface{}{
				"nested_object_key": []interface{}{
					map[string]interface{}{},
				},
			},
		},
		{
			name: "drop nonconfigurable field nested in list of objects",
			prevSpec: map[string]interface{}{
				"list_of_objects_key": []interface{}{
					map[string]interface{}{
						"nonconfigurable_field_nested_in_list_of_objects_key": "value",
					},
				},
			},
			expected: map[string]interface{}{
				"list_of_objects_key": []interface{}{
					map[string]interface{}{},
				},
			},
		},
	}
	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			c := mgr.GetClient()
			testcontroller.EnsureNamespaceExistsT(t, c, testID)
			r := resourceSkeleton()
			r.SetNamespace(testID)
			r.Spec = tc.prevSpec
			if tc.hasResourceReferences {
				references := []*unstructured.Unstructured{
					test.NewBarUnstructured("my-ref1", testID, corev1.ConditionTrue),
					test.NewBarUnstructured("my-ref2", testID, corev1.ConditionTrue),
				}
				test.EnsureObjectsExist(t, references, c)
			}
			if tc.hasSecretReferences {
				secretsData := []map[string]interface{}{
					{
						"secret-key1": "secret-val1",
					},
					{
						"secret-key2": "secret-val2",
					},
				}
				secrets := []*unstructured.Unstructured{
					test.NewSecretUnstructured("secret1", testID, secretsData[0]),
					test.NewSecretUnstructured("secret2", testID, secretsData[1]),
				}
				test.EnsureObjectsExist(t, secrets, c)
			}
			actual, _, err := KRMResourceToTFResourceConfig(r, c, smLoader)
			if err != nil {
				t.Fatalf("error convert to TF resource config: %v", err)
			}
			if !test.Equals(t, tc.expected, actual.Raw) {
				t.Fatalf("expected: %v, actual %v", tc.expected, actual.Raw)
			}
		})
	}
}

func TestKRMResourceMetadataToTFConfig(t *testing.T) {
	tests := []struct {
		name           string
		rc             *corekccv1alpha1.ResourceConfig
		annotation     map[string]string
		metadataName   string
		expectedConfig map[string]interface{}
	}{
		{
			name: "directives",
			annotation: map[string]string{
				"cnrm.cloud.google.com/directive-key": "true",
			},
			expectedConfig: map[string]interface{}{
				"directive_key": true,
			},
		},
		{
			name:         "metadata.name",
			metadataName: "my-name",
			expectedConfig: map[string]interface{}{
				"string_key": "my-name",
			},
			rc: &corekccv1alpha1.ResourceConfig{
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "string_key",
				},
			},
		},
		{
			name:         "metadata.name with value templating",
			metadataName: "my-name",
			expectedConfig: map[string]interface{}{
				"string_key": "resources/my-name",
			},
			rc: &corekccv1alpha1.ResourceConfig{
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name:              "string_key",
					NameValueTemplate: "resources/{{value}}",
				},
			},
		},
		{
			// KRMResourceToTFResourceConfig should not error out if the
			// resource doesn't have a metadata.name. This can happen when
			// KRMResourceToTFResourceConfig is called with resource skeletons
			// (e.g. the config-connector CLI calls
			// KRMResourceToTFResourceConfig with Project skeletons that don't
			// have a metadata.name since a Project's status.number can also
			// serve as its identity).
			name:           "empty metadata.name",
			expectedConfig: map[string]interface{}{},
			rc: &corekccv1alpha1.ResourceConfig{
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "string_key",
				},
			},
		},
	}
	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			c := mgr.GetClient()
			r := resourceSkeleton()
			if tc.rc != nil {
				r.ResourceConfig = *tc.rc
			}
			r.SetName(tc.metadataName)
			r.SetNamespace(testID)
			r.SetAnnotations(tc.annotation)
			actual, _, err := KRMResourceToTFResourceConfig(r, c, smLoader)
			if err != nil {
				t.Fatalf("error convert to TF resource config: %v", err)
			}
			if !test.Equals(t, tc.expectedConfig, actual.Raw) {
				t.Fatalf("expected: %v, actual %v", tc.expectedConfig, actual.Raw)
			}
		})
	}
}

func TestKRMResourceToTFResourceHierarchicalReferencesAndContainers(t *testing.T) {
	tests := []struct {
		name                 string
		rc                   *corekccv1alpha1.ResourceConfig
		annotations          map[string]string
		spec                 map[string]interface{}
		hasResourceReference bool
		expectedConfig       map[string]interface{}
	}{
		{
			name: "container annotations are mapped to config if hierarchical references are not supported",
			rc: &corekccv1alpha1.ResourceConfig{
				Containers: []corekccv1alpha1.Container{
					{
						Type:    corekccv1alpha1.ContainerTypeProject,
						TFField: "parent_key",
					},
				},
			},
			annotations: map[string]string{
				"cnrm.cloud.google.com/project-id": "project-id-from-annotations",
			},
			expectedConfig: map[string]interface{}{
				"parent_key": "project-id-from-annotations",
			},
		},
		{
			name: "container annotations support value templating",
			rc: &corekccv1alpha1.ResourceConfig{
				Containers: []corekccv1alpha1.Container{
					{
						Type:          corekccv1alpha1.ContainerTypeProject,
						TFField:       "parent_key",
						ValueTemplate: "projects/{{value}}",
					},
				},
			},
			annotations: map[string]string{
				"cnrm.cloud.google.com/project-id": "project-id-from-annotation",
			},
			expectedConfig: map[string]interface{}{
				"parent_key": "projects/project-id-from-annotation",
			},
		},
		{
			name: "hierarchical references are mapped to config over container annotations",
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
			annotations: map[string]string{
				"cnrm.cloud.google.com/project-id": "project-id-from-annotations",
			},
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "my-ref",
				},
			},
			hasResourceReference: true,
			expectedConfig: map[string]interface{}{
				"parent_key": "my-ref",
			},
		},
		{
			name: "external hierarchical references are mapped to config over container annotations",
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
			annotations: map[string]string{
				"cnrm.cloud.google.com/project-id": "project-id-from-annotations",
			},
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project-id-from-spec",
				},
			},
			expectedConfig: map[string]interface{}{
				"parent_key": "project-id-from-spec",
			},
		},
		{
			name: "hierarchical references are mapped to config for resource that only supports hierarchical references",
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
			annotations: map[string]string{
				"cnrm.cloud.google.com/project-id": "project-id-from-annotations",
			},
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "my-ref",
				},
			},
			hasResourceReference: true,
			expectedConfig: map[string]interface{}{
				"parent_key": "my-ref",
			},
		},
		{
			name: "external hierarchical references are mapped to config for resource that only supports hierarchical references",
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
			annotations: map[string]string{
				"cnrm.cloud.google.com/project-id": "project-id-from-annotations",
			},
			spec: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project-id-from-spec",
				},
			},
			expectedConfig: map[string]interface{}{
				"parent_key": "project-id-from-spec",
			},
		},
		{
			// KRMResourceToTFResourceConfig should not error out if spec does
			// not contain a hierarchical reference. Even if we can expect the
			// resource to always contain a hierarchical reference if
			// KRMResourceToTFResourceConfig is called from the controller,
			// other components like the config-connector CLI can also call
			// KRMResourceToTFResourceConfig to perform GETs with resource
			// skeletons that may not have a hierarchical reference.
			name: "spec without hierarchical reference does not result in an error",
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
			spec:           map[string]interface{}{},
			expectedConfig: map[string]interface{}{},
		},
	}

	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			c := mgr.GetClient()
			testcontroller.EnsureNamespaceExistsT(t, c, testID)
			r := resourceSkeleton()
			if tc.rc != nil {
				r.ResourceConfig = *tc.rc
			}
			r.SetNamespace(testID)
			r.SetAnnotations(tc.annotations)
			r.Spec = tc.spec
			if tc.hasResourceReference {
				references := []*unstructured.Unstructured{
					test.NewBarUnstructured("my-ref", testID, corev1.ConditionTrue),
				}
				test.EnsureObjectsExist(t, references, c)
			}
			actual, _, err := KRMResourceToTFResourceConfig(r, c, smLoader)
			if err != nil {
				t.Fatalf("error converting to TF resource config: %v", err)
			}
			if !test.Equals(t, tc.expectedConfig, actual.Raw) {
				diff := cmp.Diff(tc.expectedConfig, actual.Raw)
				t.Fatalf("actual TF config did not match expected TF config; diff (-want +got):\n%v", diff)
			}
		})
	}
}

func TestKRMResourceToTFResourceConfigSecretVersions(t *testing.T) {
	tests := []struct {
		name                               string
		spec                               map[string]interface{}
		referencedSecrets                  []*unstructured.Unstructured
		secretVersionsShouldContainSecrets []string
	}{
		{
			name: "multiple Secret references",
			spec: map[string]interface{}{
				"sensitiveFieldKey": map[string]interface{}{
					"valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{
							"name": "secret1",
							"key":  "secret-key1",
						},
					},
				},
				"nestedObjectKey": map[string]interface{}{
					"nestedSensitiveFieldKey": map[string]interface{}{
						"valueFrom": map[string]interface{}{
							"secretKeyRef": map[string]interface{}{
								"name": "secret2",
								"key":  "secret-key2",
							},
						},
					},
				},
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret3",
									"key":  "secret-key3",
								},
							},
						},
					},
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret4",
									"key":  "secret-key4",
								},
							},
						},
					},
				},
			},
			referencedSecrets: []*unstructured.Unstructured{
				test.NewSecretUnstructured("secret1", "", map[string]interface{}{"secret-key1": "secret-val1"}),
				test.NewSecretUnstructured("secret2", "", map[string]interface{}{"secret-key2": "secret-val2"}),
				test.NewSecretUnstructured("secret3", "", map[string]interface{}{"secret-key3": "secret-val3"}),
				test.NewSecretUnstructured("secret4", "", map[string]interface{}{"secret-key4": "secret-val4"}),
			},
			secretVersionsShouldContainSecrets: []string{
				"secret1",
				"secret2",
				"secret3",
				"secret4",
			},
		},
		{
			name: "multiple Secret references, but shared Secret",
			spec: map[string]interface{}{
				"sensitiveFieldKey": map[string]interface{}{
					"valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{
							"name": "secret1",
							"key":  "secret-key1",
						},
					},
				},
				"nestedObjectKey": map[string]interface{}{
					"nestedSensitiveFieldKey": map[string]interface{}{
						"valueFrom": map[string]interface{}{
							"secretKeyRef": map[string]interface{}{
								"name": "secret1",
								"key":  "secret-key1",
							},
						},
					},
				},
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret1",
									"key":  "secret-key1",
								},
							},
						},
					},
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{
									"name": "secret1",
									"key":  "secret-key1",
								},
							},
						},
					},
				},
			},
			referencedSecrets: []*unstructured.Unstructured{
				test.NewSecretUnstructured("secret1", "", map[string]interface{}{"secret-key1": "secret-val1"}),
				test.NewSecretUnstructured("secret2", "", map[string]interface{}{"secret-key2": "secret-val2"}),
				test.NewSecretUnstructured("secret3", "", map[string]interface{}{"secret-key3": "secret-val3"}),
				test.NewSecretUnstructured("secret4", "", map[string]interface{}{"secret-key4": "secret-val4"}),
			},
			secretVersionsShouldContainSecrets: []string{
				"secret1",
			},
		},
		{
			name: "no Secret references",
			spec: map[string]interface{}{
				"sensitiveFieldKey": map[string]interface{}{
					"value": "val",
				},
				"nestedObjectKey": map[string]interface{}{
					"nestedSensitiveFieldKey": map[string]interface{}{
						"value": "val",
					},
				},
				"listOfObjectsKey": []interface{}{
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"value": "val",
						},
					},
					map[string]interface{}{
						"sensitiveFieldNestedInListOfObjectsKey": map[string]interface{}{
							"value": "val",
						},
					},
				},
			},
			referencedSecrets: []*unstructured.Unstructured{
				test.NewSecretUnstructured("secret1", "", map[string]interface{}{"secret-key1": "secret-val1"}),
				test.NewSecretUnstructured("secret2", "", map[string]interface{}{"secret-key2": "secret-val2"}),
				test.NewSecretUnstructured("secret3", "", map[string]interface{}{"secret-key3": "secret-val3"}),
				test.NewSecretUnstructured("secret4", "", map[string]interface{}{"secret-key4": "secret-val4"}),
			},
			secretVersionsShouldContainSecrets: []string{},
		},
		{
			name: "no sensitive fields",
			spec: map[string]interface{}{
				"stringKey": "StringVal",
			},
			referencedSecrets: []*unstructured.Unstructured{
				test.NewSecretUnstructured("secret1", "", map[string]interface{}{"secret-key1": "secret-val1"}),
				test.NewSecretUnstructured("secret2", "", map[string]interface{}{"secret-key2": "secret-val2"}),
				test.NewSecretUnstructured("secret3", "", map[string]interface{}{"secret-key3": "secret-val3"}),
				test.NewSecretUnstructured("secret4", "", map[string]interface{}{"secret-key4": "secret-val4"}),
			},
			secretVersionsShouldContainSecrets: []string{},
		},
	}
	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			c := mgr.GetClient()
			testcontroller.EnsureNamespaceExistsT(t, c, testID)
			r := resourceSkeleton()
			r.SetNamespace(testID)
			r.Spec = tc.spec
			for _, obj := range tc.referencedSecrets {
				obj.SetNamespace(testID)
			}
			test.EnsureObjectsExist(t, tc.referencedSecrets, c)
			expectedSecretVersions := make(map[string]string)
			for _, secretName := range tc.secretVersionsShouldContainSecrets {
				version, err := getResourceVersionOfSecret(secretName, testID, c)
				if err != nil {
					t.Fatalf("error determining version of Secret %v: %v", secretName, err)
				}
				expectedSecretVersions[secretName] = version
			}
			_, secretVersions, err := KRMResourceToTFResourceConfig(r, c, smLoader)
			if err != nil {
				t.Fatalf("error converting to TF resource config: %v", err)
			}
			if !test.Equals(t, expectedSecretVersions, secretVersions) {
				t.Fatalf("got: %v, wanted: %v", secretVersions, expectedSecretVersions)
			}
		})
	}
}

func getResourceVersionOfSecret(name, namespace string, c client.Client) (string, error) {
	nn := types.NamespacedName{
		Name:      name,
		Namespace: namespace,
	}
	secret := v1.Secret{}
	if err := c.Get(context.TODO(), nn, &secret); err != nil {
		return "", err
	}
	return secret.GetResourceVersion(), nil
}

func TestKRMResourceResourceIDToTFConfig(t *testing.T) {
	tests := []struct {
		name           string
		rc             *corekccv1alpha1.ResourceConfig
		metadataName   string
		prevSpec       map[string]interface{}
		expectedConfig map[string]interface{}
		hasError       bool
	}{
		{
			name: "nonempty user-specified resource ID",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "string_key",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "string_key",
				},
			},
			metadataName: "metadata-name",
			prevSpec: map[string]interface{}{
				"resourceID": "resource-id",
			},
			expectedConfig: map[string]interface{}{
				"string_key": "resource-id",
			},
		},
		{
			name: "nonempty server-generated resource ID",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "string_key",
				},
				ServerGeneratedIDField: "string_key",
			},
			prevSpec: map[string]interface{}{
				"resourceID": "resource-id",
			},
			expectedConfig: map[string]interface{}{},
		},
		{
			name: "resource ID with value template",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField:   "string_key",
					ValueTemplate: "values/{{value}}",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name:              "string_key",
					NameValueTemplate: "values/{{value}}",
				},
			},
			metadataName: "metadata-name",
			prevSpec: map[string]interface{}{
				"resourceID": "resource-id",
			},
			expectedConfig: map[string]interface{}{
				"string_key": "values/resource-id",
			},
		},
		{
			name: "empty resource ID",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "string_key",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "string_key",
				},
			},
			metadataName: "metadata-name",
			prevSpec: map[string]interface{}{
				"resourceID": "",
			},
			hasError: true,
		},
		{
			name: "unspecified resource ID with non-empty metadata.name",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "string_key",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "string_key",
				},
			},
			metadataName: "metadata-name",
			prevSpec:     map[string]interface{}{},
			expectedConfig: map[string]interface{}{
				"string_key": "metadata-name",
			},
		},
		{
			// KRMResourceToTFResourceConfig should not error out if the
			// resource has neither a metadata.name or spec.resourceID. This
			// can happen when KRMResourceToTFResourceConfig is called with
			// resource skeletons (e.g. the config-connector CLI calls
			// KRMResourceToTFResourceConfig with Project skeletons that don't
			// have a metadata.name or spec.resourceID since a Project's
			// status.number can also serve as its identity).
			name: "unspecified resource ID and unspecified metadata.name",
			rc: &corekccv1alpha1.ResourceConfig{
				ResourceID: corekccv1alpha1.ResourceID{
					TargetField: "string_key",
				},
				MetadataMapping: corekccv1alpha1.MetadataMapping{
					Name: "string_key",
				},
			},
			prevSpec:       map[string]interface{}{},
			expectedConfig: map[string]interface{}{},
		},
	}
	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			c := mgr.GetClient()
			r := resourceSkeleton()
			if tc.rc != nil {
				r.ResourceConfig = *tc.rc
			}
			if tc.metadataName != "" {
				r.SetName(tc.metadataName)
			}
			r.SetNamespace(testID)
			r.Spec = tc.prevSpec

			actual, _, err := KRMResourceToTFResourceConfig(r, c, smLoader)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, want an error")
				}
				return
			} else if err != nil {
				t.Fatalf("error converting KRM resource to TF "+
					"resource config: %v", err)
			}

			if got, want := actual.Raw, tc.expectedConfig; !test.Equals(t, got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	testmain.ForUnitTests(m, &mgr)
}
