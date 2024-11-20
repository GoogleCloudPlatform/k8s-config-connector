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
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	. "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestResolveResourceReferenceToTFResource(t *testing.T) {
	testID := testvariable.NewUniqueID()
	c := mgr.GetClient()
	gvk := schema.GroupVersionKind{Group: "test1.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "Foo"}
	ns := testID
	testcontroller.EnsureNamespaceExistsT(t, c, ns)
	resource := &Resource{}
	resource.SetGroupVersionKind(gvk)
	resource.SetNamespace(ns)
	tfField := "bar_field"
	tests := []struct {
		name                string
		config              map[string]interface{}
		referencedResources []*unstructured.Unstructured
		refConfig           v1alpha1.ReferenceConfig
		expectedFinalConfig map[string]interface{}
		shouldError         bool
	}{
		// Embedded TypeConfig
		{
			name: "embedded type, found",
			config: map[string]interface{}{
				"key1": "val1",
				"barRef": map[string]interface{}{
					"name": "my-bar",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-bar", ns, corev1.ConditionTrue),
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1":     "val1",
				"barField": "my-bar",
			},
		},
		{
			name: "embedded type, not found",
			config: map[string]interface{}{
				"key1": "val1",
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Kind: "Bar",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1": "val1",
			},
		},
		{
			name: "embedded type, ValueTemplate, found",
			config: map[string]interface{}{
				"key1": "val1",
				"barRef": map[string]interface{}{
					"name": "my-bar",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-bar", ns, corev1.ConditionTrue),
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					ValueTemplate: "user:{{value}}",
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1":     "val1",
				"barField": "user:my-bar",
			},
		},

		//Nested Types
		{
			name: "nested type, referenced object, found",
			config: map[string]interface{}{
				"key1": "val1",
				"barField": map[string]interface{}{
					"barRef": map[string]interface{}{
						"name": "my-bar",
					},
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-bar", ns, corev1.ConditionTrue),
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				Types: []v1alpha1.TypeConfig{
					{
						Key: "barRef",
						GVK: schema.GroupVersionKind{
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
			expectedFinalConfig: map[string]interface{}{
				"key1":     "val1",
				"barField": "my-bar",
			},
		},
		{
			name: "nested type, JSONSchemaType value, found",
			config: map[string]interface{}{
				"key1": "val1",
				"barField": map[string]interface{}{
					"value": "bar-value",
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				Types: []v1alpha1.TypeConfig{
					{
						Key: "barRef",
						GVK: schema.GroupVersionKind{
							Kind: "Bar",
						},
					},
					{
						Key:            "value",
						JSONSchemaType: "string",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1":     "val1",
				"barField": "bar-value",
			},
		},
		{
			name: "nested type, JSONSchemaType, ValueTemplate, found",
			config: map[string]interface{}{
				"key1": "val1",
				"barField": map[string]interface{}{
					"user": "fooser",
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				Types: []v1alpha1.TypeConfig{
					{
						Key: "barRef",
						GVK: schema.GroupVersionKind{
							Kind: "Bar",
						},
					},
					{
						Key:            "user",
						JSONSchemaType: "string",
						ValueTemplate:  "user:{{value}}",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1":     "val1",
				"barField": "user:fooser",
			},
		},
		{
			name: "nested type, not found",
			config: map[string]interface{}{
				"key1": "val1",
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				Types: []v1alpha1.TypeConfig{
					{
						Key: "barRef",
						GVK: schema.GroupVersionKind{
							Kind: "Bar",
						},
					},
					{
						Key:            "value",
						JSONSchemaType: "string",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1": "val1",
			},
		},
		{
			name: "list of objects with complex reference fields",
			config: map[string]interface{}{
				"key1": "val1",
				"foo": []interface{}{
					map[string]interface{}{
						"barField": map[string]interface{}{
							"barRef": map[string]interface{}{
								"name": "my-bar1",
							},
						},
					},
					map[string]interface{}{
						"barField": map[string]interface{}{
							"barRef": map[string]interface{}{
								"name": "my-bar2",
							},
						},
					},
					map[string]interface{}{
						"barField": map[string]interface{}{
							"group": "foobaz",
						},
					},
					map[string]interface{}{
						"barField": map[string]interface{}{
							"fooBarRef": map[string]interface{}{
								"name": "my-bar3",
							},
						},
					},
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-bar1", ns, corev1.ConditionTrue),
				test.NewBarUnstructured("my-bar2", ns, corev1.ConditionTrue),
				test.NewBarUnstructured("my-bar3", ns, corev1.ConditionTrue),
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: "foo." + tfField,
				Types: []v1alpha1.TypeConfig{
					{
						Key: "barRef",
						GVK: schema.GroupVersionKind{
							Group:   "test1.cnrm.cloud.google.com",
							Version: "v1alpha1",
							Kind:    "Test1Bar",
						},
					},
					{
						Key:            "group",
						JSONSchemaType: "string",
						ValueTemplate:  "group:{{value}}",
					},
					{
						Key: "fooBarRef",
						GVK: schema.GroupVersionKind{
							Group:   "test1.cnrm.cloud.google.com",
							Version: "v1alpha1",
							Kind:    "Test1Bar",
						},
						ValueTemplate: "fooRef: {{value}}",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1": "val1",
				"foo": []interface{}{
					map[string]interface{}{
						"barField": "my-bar1",
					},
					map[string]interface{}{
						"barField": "my-bar2",
					},
					map[string]interface{}{
						"barField": "group:foobaz",
					},
					map[string]interface{}{
						"barField": "fooRef: my-bar3",
					},
				},
			},
		},
		{
			name: "list of references",
			config: map[string]interface{}{
				"key1": "val1",
				"barField": []interface{}{
					map[string]interface{}{
						"name": "my-bar1",
					},
					map[string]interface{}{
						"name": "my-bar2",
					},
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-bar1", ns, corev1.ConditionTrue),
				test.NewBarUnstructured("my-bar2", ns, corev1.ConditionTrue),
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1": "val1",
				"barField": []interface{}{
					"my-bar1",
					"my-bar2",
				},
			},
		},
		{
			name: "reference not ready",
			config: map[string]interface{}{
				"key1": "val1",
				"barRef": map[string]interface{}{
					"name": "my-bar-not-ready",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-bar-not-ready", ns, corev1.ConditionFalse),
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1":     "val1",
				"barField": "my-bar-not-ready",
			},
			shouldError: true,
		},
		{
			name: "reference doesn't exist",
			config: map[string]interface{}{
				"key1": "val1",
				"barRef": map[string]interface{}{
					"name": "nonexistent-bar",
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
			shouldError: true,
		},

		// External references
		{
			name: "external reference",
			config: map[string]interface{}{
				"key1": "val1",
				"barRef": map[string]interface{}{
					"external": "my-bar",
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1":     "val1",
				"barField": "my-bar",
			},
		},
		{
			name: "list of objects with external references",
			config: map[string]interface{}{
				"key1": "val1",
				"foo": []interface{}{
					map[string]interface{}{
						"barRef": map[string]interface{}{
							"external": "my-bar1",
						},
					},
					map[string]interface{}{
						"barRef": map[string]interface{}{
							"external": "my-bar2",
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: "foo." + tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1": "val1",
				"foo": []interface{}{
					map[string]interface{}{
						"barField": "my-bar1",
					},
					map[string]interface{}{
						"barField": "my-bar2",
					},
				},
			},
		},
		{
			name: "list of external references",
			config: map[string]interface{}{
				"key1": "val1",
				"barField": []interface{}{
					map[string]interface{}{
						"external": "my-bar1",
					},
					map[string]interface{}{
						"external": "my-bar2",
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1": "val1",
				"barField": []interface{}{
					"my-bar1",
					"my-bar2",
				},
			},
		},
		{
			name: "mixed list of regular and external references",
			config: map[string]interface{}{
				"key1": "val1",
				"barField": []interface{}{
					map[string]interface{}{
						"name": "my-bar1",
					},
					map[string]interface{}{
						"external": "my-bar2",
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1": "val1",
				"barField": []interface{}{
					"my-bar1",
					"my-bar2",
				},
			},
		},
		{
			name: "reference with arbitrary target field from status",
			config: map[string]interface{}{
				"key1": "val1",
				"barRef": map[string]interface{}{
					"name": "my-bar",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-bar", ns, corev1.ConditionTrue),
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					TargetField: "status_field",
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1":     "val1",
				"barField": "foobar",
			},
		},
		{
			name: "reference with arbitrary target field from spec",
			config: map[string]interface{}{
				"key1": "val1",
				"barRef": map[string]interface{}{
					"name": "my-bar",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-bar", ns, corev1.ConditionTrue),
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					TargetField: "spec_field",
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1":     "val1",
				"barField": "abc123",
			},
		},
		{
			name: "reference with value template",
			config: map[string]interface{}{
				"key1": "val1",
				"barRef": map[string]interface{}{
					"name": "my-bar",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-bar", ns, corev1.ConditionTrue),
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					TargetField:   "status_field",
					ValueTemplate: "location/{{location}}/bars/{{value}}",
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1":     "val1",
				"barField": "location/test-location/bars/foobar",
			},
		},

		// With observed state
		{
			name: "target field is computed field under observedState, found",
			config: map[string]interface{}{
				"key1": "val1",
				"testFieldRef": map[string]interface{}{
					"name": "test-cr",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewUnstructuredWithObservedState("test-cr", ns, corev1.ConditionTrue),
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: "test_field",
				TypeConfig: v1alpha1.TypeConfig{
					Key: "testFieldRef",
					GVK: schema.GroupVersionKind{
						Group:   "test5.cnrm.cloud.google.com",
						Version: "v1beta1",
						Kind:    "TestKindWithObservedState",
					},
					TargetField: "reference_target_field",
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1":      "val1",
				"testField": "reference-value",
			},
		},
	}
	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			test.EnsureObjectsExist(t, tc.referencedResources, c)
			config := tc.config
			path := strings.Split(tc.refConfig.TFField, ".")
			if err := ResolveResourceReference(path, config, tc.refConfig, resource, c, smLoader); err != nil {
				if tc.shouldError {
					return
				}
				t.Errorf("error resolving: %v", err)
				return
			}
			if !reflect.DeepEqual(tc.expectedFinalConfig, config) {
				t.Errorf("expected config: %v, actual config: %v", tc.expectedFinalConfig, config)
			}
		})
	}
}

func TestResolveResourceReferenceToTFResource_deleting(t *testing.T) {
	testID := testvariable.NewUniqueID()
	c := mgr.GetClient()
	gvk := schema.GroupVersionKind{Group: "test1.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "Foo"}
	ns := testID
	testcontroller.EnsureNamespaceExistsT(t, c, ns)
	resource := &Resource{}
	resource.SetGroupVersionKind(gvk)
	resource.SetNamespace(ns)
	resource.SetDeletionTimestamp(&metav1.Time{Time: time.Now()})
	tfField := "bar_field"
	tests := []struct {
		name                string
		config              map[string]interface{}
		referencedResources []*unstructured.Unstructured
		refConfig           v1alpha1.ReferenceConfig
		expectedFinalConfig map[string]interface{}
		shouldError         bool
	}{
		{
			name: "deleting a resource when its reference is ready",
			config: map[string]interface{}{
				"key1": "val1",
				"barRef": map[string]interface{}{
					"name": "my-bar",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-bar", ns, corev1.ConditionTrue),
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1":     "val1",
				"barField": "my-bar",
			},
			shouldError: false,
		},
		{
			name: "deleting a resource when its reference not ready",
			config: map[string]interface{}{
				"key1": "val1",
				"barRef": map[string]interface{}{
					"name": "my-bar-not-ready",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-bar-not-ready", ns, corev1.ConditionFalse),
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"key1":     "val1",
				"barField": "my-bar-not-ready",
			},
			shouldError: false,
		},
	}
	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			test.EnsureObjectsExist(t, tc.referencedResources, c)
			config := tc.config
			path := strings.Split(tc.refConfig.TFField, ".")
			if err := ResolveResourceReference(path, config, tc.refConfig, resource, c, smLoader); err != nil {
				if tc.shouldError {
					return
				}
				t.Errorf("error resolving: %v", err)
				return
			}
			if !reflect.DeepEqual(tc.expectedFinalConfig, config) {
				t.Errorf("expected config: %v, actual config: %v", tc.expectedFinalConfig, config)
			}
		})
	}
}

func TestResolveResourceReferenceToDCLResourceWithResourceID(t *testing.T) {
	testID := testvariable.NewUniqueID()
	c := mgr.GetClient()
	gvk := schema.GroupVersionKind{
		Group:   "test1.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "Foo",
	}
	ns := testID
	if err := testcontroller.EnsureNamespaceExists(c, ns); err != nil {
		t.Fatal(err)
	}

	resource := &Resource{}
	resource.SetGroupVersionKind(gvk)
	resource.SetNamespace(ns)
	tfField := "bar_field"
	tests := []struct {
		name                string
		config              map[string]interface{}
		referencedResources []*unstructured.Unstructured
		refConfig           v1alpha1.ReferenceConfig
		expectedFinalConfig map[string]interface{}
		hasError            bool
	}{
		{
			name: "refResource with spec.resourceID",
			config: map[string]interface{}{
				"userSpecifiedResourceIDKindRef": map[string]interface{}{
					"name": "ref-resource-with-id",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
						"kind":       "Test4DCLResourceUserSpecifiedResourceIDKind",
						"metadata": map[string]interface{}{
							"name":      "ref-resource-with-id",
							"namespace": ns,
						},
						"spec": map[string]interface{}{
							"resourceID": "user-specified-id",
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": corev1.ConditionTrue,
								},
							},
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "userSpecifiedResourceIDKindRef",
					GVK: schema.GroupVersionKind{
						Group:   "test4.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test4DCLResourceUserSpecifiedResourceIDKind",
					},
					DCLBasedResource: true,
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"barField": "user-specified-id",
			},
		},
		{
			name: "reference config has a value template",
			config: map[string]interface{}{
				"userSpecifiedResourceIDKindRef": map[string]interface{}{
					"name": "ref-resource-with-id",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
						"kind":       "Test4DCLResourceUserSpecifiedResourceIDKind",
						"metadata": map[string]interface{}{
							"name":      "ref-resource-with-id",
							"namespace": ns,
						},
						"spec": map[string]interface{}{
							"resourceID": "user-specified-id",
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": corev1.ConditionTrue,
								},
							},
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "userSpecifiedResourceIDKindRef",
					GVK: schema.GroupVersionKind{
						Group:   "test4.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test4DCLResourceUserSpecifiedResourceIDKind",
					},
					ValueTemplate:    "user:{{value}}",
					DCLBasedResource: true,
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"barField": "user:user-specified-id",
			},
		},
		{
			name: "uses the unresolved value of spec.resourceID when " +
				"reference config has a value template referencing the " +
				"target field via '{{value}}'",
			config: map[string]interface{}{
				"serverGeneratedResourceIDKindRef": map[string]interface{}{
					"name": "ref-resource-with-value-template",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
						"kind":       "Test4DCLResourceServerGeneratedResourceIDKind",
						"metadata": map[string]interface{}{
							"name":      "ref-resource-with-value-template",
							"namespace": ns,
						},
						"spec": map[string]interface{}{
							"resourceID": "server-generated-id",
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": corev1.ConditionTrue,
								},
							},
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "serverGeneratedResourceIDKindRef",
					GVK: schema.GroupVersionKind{
						Group:   "test4.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test4DCLResourceServerGeneratedResourceIDKind",
					},
					ValueTemplate:    "user:{{value}}",
					DCLBasedResource: true,
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"barField": "user:server-generated-id",
			},
		},
		{
			name: "uses the target field rather than the resolved value of " +
				"spec.resourceID when reference config has a value template " +
				"explicitly referencing the target field of resource ID",
			config: map[string]interface{}{
				"serverGeneratedResourceIDKindRef": map[string]interface{}{
					"name": "ref-resource-with-status-field",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
						"kind":       "Test4DCLResourceServerGeneratedResourceIDKind",
						"metadata": map[string]interface{}{
							"name":      "ref-resource-with-status-field",
							"namespace": ns,
						},
						"spec": map[string]interface{}{
							"resourceID": "server-generated-id",
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": corev1.ConditionTrue,
								},
							},
							"resourceIdField": "the-correct-id",
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "serverGeneratedResourceIDKindRef",
					GVK: schema.GroupVersionKind{
						Group:   "test4.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test4DCLResourceServerGeneratedResourceIDKind",
					},
					ValueTemplate:    "user:{{resource_id_field}}",
					DCLBasedResource: true,
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"barField": "user:the-correct-id",
			},
		},
		{
			name: "refResource empty spec.resourceID",
			config: map[string]interface{}{
				"userSpecifiedResourceIDKindRef": map[string]interface{}{
					"name": "ref-resource-with-empty-id",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
						"kind":       "Test4DCLResourceUserSpecifiedResourceIDKind",
						"metadata": map[string]interface{}{
							"name":      "ref-resource-with-empty-id",
							"namespace": ns,
						},
						"spec": map[string]interface{}{
							"resourceID": "",
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": corev1.ConditionTrue,
								},
							},
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "userSpecifiedResourceIDKindRef",
					GVK: schema.GroupVersionKind{
						Group:   "test4.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test4DCLResourceUserSpecifiedResourceIDKind",
					},
					DCLBasedResource: true,
				},
			},
			hasError: true,
		},
	}
	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			test.EnsureObjectsExist(t, tc.referencedResources, c)
			config := tc.config
			path := strings.Split(tc.refConfig.TFField, ".")
			err := ResolveResourceReference(path, config, tc.refConfig,
				resource, c, smLoader)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, want an error")
				}
				return
			} else if err != nil {
				t.Fatalf("error resolving resource references that "+
					"support resource ID field: %v", err)
				return
			}
			if got, want := config, tc.expectedFinalConfig; !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestResolveResourceReferenceToTFResourceWithResourceID(t *testing.T) {
	testID := testvariable.NewUniqueID()
	c := mgr.GetClient()
	gvk := schema.GroupVersionKind{
		Group:   "test1.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "Foo",
	}
	ns := testID
	if err := testcontroller.EnsureNamespaceExists(c, ns); err != nil {
		t.Fatal(err)
	}

	resource := &Resource{}
	resource.SetGroupVersionKind(gvk)
	resource.SetNamespace(ns)
	tfField := "bar_field"
	tests := []struct {
		name                string
		config              map[string]interface{}
		referencedResources []*unstructured.Unstructured
		refConfig           v1alpha1.ReferenceConfig
		expectedFinalConfig map[string]interface{}
		hasError            bool
	}{
		{
			name: "refResource with spec.resourceID",
			config: map[string]interface{}{
				"userSpecifiedResourceIDKindRef": map[string]interface{}{
					"name": "ref-resource-with-id",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "test3.cnrm.cloud.google.com/v1alpha1",
						"kind":       "Test3UserSpecifiedResourceIDKind",
						"metadata": map[string]interface{}{
							"name":      "ref-resource-with-id",
							"namespace": ns,
						},
						"spec": map[string]interface{}{
							"resourceID": "user-specified-id",
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": corev1.ConditionTrue,
								},
							},
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "userSpecifiedResourceIDKindRef",
					GVK: schema.GroupVersionKind{
						Group:   "test3.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test3UserSpecifiedResourceIDKind",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"barField": "user-specified-id",
			},
		},
		{
			name: "refResource with user-specified spec.resourceID unset and " +
				"metadata.name set",
			config: map[string]interface{}{
				"userSpecifiedResourceIDKindRef": map[string]interface{}{
					"name": "ref-resource-with-unspecified-id",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "test3.cnrm.cloud.google.com/v1alpha1",
						"kind":       "Test3UserSpecifiedResourceIDKind",
						"metadata": map[string]interface{}{
							"name":      "ref-resource-with-unspecified-id",
							"namespace": ns,
						},
						"spec": map[string]interface{}{},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": corev1.ConditionTrue,
								},
							},
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "userSpecifiedResourceIDKindRef",
					GVK: schema.GroupVersionKind{
						Group:   "test3.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test3UserSpecifiedResourceIDKind",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"barField": "ref-resource-with-unspecified-id",
			},
		},
		{
			name: "reference config has a value template",
			config: map[string]interface{}{
				"userSpecifiedResourceIDKindRef": map[string]interface{}{
					"name": "ref-resource-with-id",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "test3.cnrm.cloud.google.com/v1alpha1",
						"kind":       "Test3UserSpecifiedResourceIDKind",
						"metadata": map[string]interface{}{
							"name":      "ref-resource-with-id",
							"namespace": ns,
						},
						"spec": map[string]interface{}{
							"resourceID": "user-specified-id",
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": corev1.ConditionTrue,
								},
							},
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "userSpecifiedResourceIDKindRef",
					GVK: schema.GroupVersionKind{
						Group:   "test3.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test3UserSpecifiedResourceIDKind",
					},
					ValueTemplate: "user:{{value}}",
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"barField": "user:user-specified-id",
			},
		},
		{
			name: "refResource with server-generated spec.resourceID unset " +
				"and the server-generated ID in status set",
			config: map[string]interface{}{
				"serverGeneratedResourceIDKindRef": map[string]interface{}{
					"name": "ref-resource-with-server-generated-id",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "test3.cnrm.cloud.google.com/v1alpha1",
						"kind":       "Test3ServerGeneratedResourceIDKind",
						"metadata": map[string]interface{}{
							"name":      "ref-resource-with-server-generated-id",
							"namespace": ns,
						},
						"spec": map[string]interface{}{},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": corev1.ConditionTrue,
								},
							},
							"resourceIdField": "values/id-in-status",
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "serverGeneratedResourceIDKindRef",
					GVK: schema.GroupVersionKind{
						Group:   "test3.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test3ServerGeneratedResourceIDKind",
					},
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"barField": "id-in-status",
			},
		},
		{
			name: "uses the unresolved value of spec.resourceID when " +
				"reference config has a value template referencing the " +
				"target field via '{{value}}'",
			config: map[string]interface{}{
				"serverGeneratedResourceIDKindRef": map[string]interface{}{
					"name": "ref-resource-with-value-template",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "test3.cnrm.cloud.google.com/v1alpha1",
						"kind":       "Test3ServerGeneratedResourceIDKind",
						"metadata": map[string]interface{}{
							"name":      "ref-resource-with-value-template",
							"namespace": ns,
						},
						"spec": map[string]interface{}{
							"resourceID": "server-generated-id",
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": corev1.ConditionTrue,
								},
							},
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "serverGeneratedResourceIDKindRef",
					GVK: schema.GroupVersionKind{
						Group:   "test3.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test3ServerGeneratedResourceIDKind",
					},
					ValueTemplate: "user:{{value}}",
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"barField": "user:server-generated-id",
			},
		},
		{
			name: "uses the resolved value of spec.resourceID when reference " +
				"config has a value template explicitly referencing the " +
				"target field of resource ID",
			config: map[string]interface{}{
				"serverGeneratedResourceIDKindRef": map[string]interface{}{
					"name": "ref-resource-with-status-field",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "test3.cnrm.cloud.google.com/v1alpha1",
						"kind":       "Test3ServerGeneratedResourceIDKind",
						"metadata": map[string]interface{}{
							"name":      "ref-resource-with-status-field",
							"namespace": ns,
						},
						"spec": map[string]interface{}{
							"resourceID": "server-generated-id",
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": corev1.ConditionTrue,
								},
							},
							"resourceIdField": "a-wrong-id",
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "serverGeneratedResourceIDKindRef",
					GVK: schema.GroupVersionKind{
						Group:   "test3.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test3ServerGeneratedResourceIDKind",
					},
					ValueTemplate: "user:{{resource_id_field}}",
				},
			},
			expectedFinalConfig: map[string]interface{}{
				"barField": "user:values/server-generated-id",
			},
		},
		{
			name: "refResource empty spec.resourceID",
			config: map[string]interface{}{
				"userSpecifiedResourceIDKindRef": map[string]interface{}{
					"name": "ref-resource-with-empty-id",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "test3.cnrm.cloud.google.com/v1alpha1",
						"kind":       "Test3UserSpecifiedResourceIDKind",
						"metadata": map[string]interface{}{
							"name":      "ref-resource-with-empty-id",
							"namespace": ns,
						},
						"spec": map[string]interface{}{
							"resourceID": "",
						},
						"status": map[string]interface{}{
							"conditions": []interface{}{
								map[string]interface{}{
									"type":   "Ready",
									"status": corev1.ConditionTrue,
								},
							},
						},
					},
				},
			},
			refConfig: v1alpha1.ReferenceConfig{
				TFField: tfField,
				TypeConfig: v1alpha1.TypeConfig{
					Key: "userSpecifiedResourceIDKindRef",
					GVK: schema.GroupVersionKind{
						Group:   "test3.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test3UserSpecifiedResourceIDKind",
					},
				},
			},
			hasError: true,
		},
	}
	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			test.EnsureObjectsExist(t, tc.referencedResources, c)
			config := tc.config
			path := strings.Split(tc.refConfig.TFField, ".")
			err := ResolveResourceReference(path, config, tc.refConfig,
				resource, c, smLoader)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, want an error")
				}
				return
			} else if err != nil {
				t.Fatalf("error resolving resource references that "+
					"support resource ID field: %v", err)
				return
			}
			if got, want := config, tc.expectedFinalConfig; !reflect.DeepEqual(got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestResource_GetReferencedDCLResource(t *testing.T) {
	testID := testvariable.NewUniqueID()
	c := mgr.GetClient()
	gvk := schema.GroupVersionKind{Group: "test1.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "Foo"}
	ns1 := testID + "-1"
	ns2 := testID + "-2"
	testcontroller.EnsureNamespaceExistsT(t, c, ns1)
	testcontroller.EnsureNamespaceExistsT(t, c, ns2)
	resource := &Resource{}
	resource.SetGroupVersionKind(gvk)
	resource.SetNamespace(ns1)
	tests := []struct {
		name               string
		referencedResource *unstructured.Unstructured
		reference          *v1alpha1.ResourceReference
		refConfig          v1alpha1.ReferenceConfig
		shouldError        bool
	}{
		{
			name: "reference in same namespace",
			referencedResource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"name":      "reference-in-same-namespace",
						"namespace": ns1,
					},
				},
			},
			reference: &v1alpha1.ResourceReference{
				Name: "reference-in-same-namespace",
			},
			refConfig: v1alpha1.ReferenceConfig{
				TypeConfig: v1alpha1.TypeConfig{
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					DCLBasedResource: true,
				},
			},
		},
		{
			name: "reference in different namespace",
			referencedResource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"name":      "reference-in-different-namespace",
						"namespace": ns2,
					},
				},
			},
			reference: &v1alpha1.ResourceReference{
				Name:      "reference-in-different-namespace",
				Namespace: ns2,
			},
			refConfig: v1alpha1.ReferenceConfig{
				TypeConfig: v1alpha1.TypeConfig{
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					DCLBasedResource: true,
				},
			},
		},
		{
			name: "cross-group reference",
			referencedResource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test2.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test2Baz",
					"metadata": map[string]interface{}{
						"name":      "cross-group-reference",
						"namespace": ns1,
					},
				},
			},
			reference: &v1alpha1.ResourceReference{
				Name: "cross-group-reference",
			},
			refConfig: v1alpha1.ReferenceConfig{
				TypeConfig: v1alpha1.TypeConfig{
					GVK: schema.GroupVersionKind{
						Group:   "test2.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test2Baz",
					},
					DCLBasedResource: true,
				},
			},
		},
		{
			name: "external reference",
			reference: &v1alpha1.ResourceReference{
				External: "external-reference",
			},
			refConfig: v1alpha1.ReferenceConfig{
				TypeConfig: v1alpha1.TypeConfig{
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
					DCLBasedResource: true,
				},
			},
			shouldError: true,
		},
	}
	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.referencedResource != nil {
				if err := c.Create(context.Background(), tc.referencedResource); err != nil {
					t.Fatalf("error creating referenced resource in API server: %v", err)
				}
			}
			rsrc, err := GetReferencedResource(resource, tc.refConfig.TypeConfig, tc.reference, c, smLoader)
			if err != nil {
				if tc.shouldError {
					return
				}
				t.Fatalf("error getting referenced resource: %v", err)
			}
			if rsrc.GetName() != tc.referencedResource.GetName() {
				t.Fatalf("got unexpected referenced resource %v", rsrc.GetName())
			}
			if rsrc.ResourceConfig.Kind != "" {
				t.Fatalf("got the ResourceConfig Kind, but wanted an empty string")
			}
		})
	}
}

func TestResource_GetReferencedTFResource(t *testing.T) {
	testID := testvariable.NewUniqueID()
	c := mgr.GetClient()
	gvk := schema.GroupVersionKind{Group: "test1.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "Foo"}
	ns1 := testID + "-1"
	ns2 := testID + "-2"
	testcontroller.EnsureNamespaceExistsT(t, c, ns1)
	testcontroller.EnsureNamespaceExistsT(t, c, ns2)
	resource := &Resource{}
	resource.SetGroupVersionKind(gvk)
	resource.SetNamespace(ns1)
	tests := []struct {
		name               string
		referencedResource *unstructured.Unstructured
		reference          *v1alpha1.ResourceReference
		refConfig          v1alpha1.ReferenceConfig
		shouldError        bool
	}{
		{
			name: "reference in same namespace",
			referencedResource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"name":      "reference-in-same-namespace",
						"namespace": ns1,
					},
				},
			},
			reference: &v1alpha1.ResourceReference{
				Name: "reference-in-same-namespace",
			},
			refConfig: v1alpha1.ReferenceConfig{
				TypeConfig: v1alpha1.TypeConfig{
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
		},
		{
			name: "reference in different namespace",
			referencedResource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"name":      "reference-in-different-namespace",
						"namespace": ns2,
					},
				},
			},
			reference: &v1alpha1.ResourceReference{
				Name:      "reference-in-different-namespace",
				Namespace: ns2,
			},
			refConfig: v1alpha1.ReferenceConfig{
				TypeConfig: v1alpha1.TypeConfig{
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
		},
		{
			name: "cross-group reference",
			referencedResource: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test2.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test2Baz",
					"metadata": map[string]interface{}{
						"name":      "cross-group-reference",
						"namespace": ns1,
					},
				},
			},
			reference: &v1alpha1.ResourceReference{
				Name: "cross-group-reference",
			},
			refConfig: v1alpha1.ReferenceConfig{
				TypeConfig: v1alpha1.TypeConfig{
					GVK: schema.GroupVersionKind{
						Group:   "test2.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test2Baz",
					},
				},
			},
		},
		{
			name: "external reference",
			reference: &v1alpha1.ResourceReference{
				External: "external-reference",
			},
			refConfig: v1alpha1.ReferenceConfig{
				TypeConfig: v1alpha1.TypeConfig{
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Bar",
					},
				},
			},
			shouldError: true,
		},
	}
	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.referencedResource != nil {
				if err := c.Create(context.Background(), tc.referencedResource); err != nil {
					t.Fatalf("error creating referenced resource in API server: %v", err)
				}
			}
			rsrc, err := GetReferencedResource(resource, tc.refConfig.TypeConfig, tc.reference, c, smLoader)
			if err != nil {
				if tc.shouldError {
					return
				}
				t.Fatalf("error getting referenced resource: %v", err)
			}
			if rsrc.GetName() != tc.referencedResource.GetName() {
				t.Fatalf("got unexpected referenced resource %v", rsrc.GetName())
			}
			if rsrc.ResourceConfig.Kind != tc.refConfig.GVK.Kind {
				t.Fatalf("got the ResourceConfig for kind %v, but want %v", rsrc.ResourceConfig.Kind, tc.refConfig.GVK.Kind)
			}
		})
	}
}

func TestIsHierarchicalReference(t *testing.T) {
	tests := []struct {
		name             string
		ref              corekccv1alpha1.ReferenceConfig
		hierarchicalRefs []corekccv1alpha1.HierarchicalReference

		isHierarchicalRef bool
	}{
		{
			name: "resource reference is one of hierarchical references",
			ref: corekccv1alpha1.ReferenceConfig{
				TFField: "project",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key: "projectRef",
					GVK: schema.GroupVersionKind{
						Group:   "resourcemanager.cnrm.cloud.google.com",
						Version: "v1beta1",
						Kind:    "Project",
					},
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
					Key:  "organizationRef",
				},
			},
			isHierarchicalRef: true,
		},
		{
			name: "resource reference is not one of hierarchical references",
			ref: corekccv1alpha1.ReferenceConfig{
				TFField: "billing",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key: "billingRef",
					GVK: schema.GroupVersionKind{
						Group:   "resourcemanager.cnrm.cloud.google.com",
						Version: "v1beta1",
						Kind:    "BillingAccount",
					},
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
					Key:  "organizationRef",
				},
			},
			isHierarchicalRef: false,
		},
		{
			name: "resource reference is not a root-level field",
			ref: corekccv1alpha1.ReferenceConfig{
				TFField: "service_projects.project",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key: "projectRef",
					GVK: schema.GroupVersionKind{
						Group:   "resourcemanager.cnrm.cloud.google.com",
						Version: "v1beta1",
						Kind:    "Project",
					},
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
					Key:  "organizationRef",
				},
			},
			isHierarchicalRef: false,
		},
		{
			name: "resource reference looks like a hierarchical reference but is not one of hierarchical references",
			ref: corekccv1alpha1.ReferenceConfig{
				TFField: "project",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key: "projectRef",
					GVK: schema.GroupVersionKind{
						Group:   "resourcemanager.cnrm.cloud.google.com",
						Version: "v1beta1",
						Kind:    "Project",
					},
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "hostProjectRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
					Key:  "organizationRef",
				},
			},
			isHierarchicalRef: false,
		},
		{
			name: "resource reference is the only hierarchical reference",
			ref: corekccv1alpha1.ReferenceConfig{
				TFField: "project",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key: "projectRef",
					GVK: schema.GroupVersionKind{
						Group:   "resourcemanager.cnrm.cloud.google.com",
						Version: "v1beta1",
						Kind:    "Project",
					},
				},
			},
			hierarchicalRefs: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
			isHierarchicalRef: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := krmtotf.IsHierarchicalReference(tc.ref, tc.hierarchicalRefs)
			if result != tc.isHierarchicalRef {
				t.Fatalf("got %v, want %v", result, tc.isHierarchicalRef)
			}
		})
	}
}

func TestIsRequiredParentReference(t *testing.T) {
	tests := []struct {
		name        string
		ref         corekccv1alpha1.ReferenceConfig
		resource    *Resource
		isParentRef bool
	}{
		{
			name: "configured parent references",
			ref: corekccv1alpha1.ReferenceConfig{
				TFField: "foo",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key:    "fooRef",
					Parent: true,
				},
			},
			isParentRef: true,
			resource: &Resource{
				Resource: k8s.Resource{
					TypeMeta: metav1.TypeMeta{
						Kind: "Bar",
					},
				},
			},
		},
		{
			name: "regular references",
			ref: corekccv1alpha1.ReferenceConfig{
				TFField: "foo",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key: "fooRef",
				},
			},
			resource: &Resource{
				Resource: k8s.Resource{
					TypeMeta: metav1.TypeMeta{
						Kind: "Bar",
					},
				},
			},
			isParentRef: false,
		},
		{
			name: "project's hierarchical references are not parent references",
			ref: corekccv1alpha1.ReferenceConfig{
				TFField: "folder",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key: "folderRef",
					GVK: schema.GroupVersionKind{
						Group:   "resourcemanager.cnrm.cloud.google.com",
						Version: "v1beta1",
						Kind:    "Folder",
					},
				},
			},
			resource: &Resource{
				Resource: k8s.Resource{
					TypeMeta: metav1.TypeMeta{
						Kind: "Project",
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{
					HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
						{
							Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
							Key:  "folderRef",
						},
						{
							Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
							Key:  "organizationRef",
						},
					},
				},
			},
			isParentRef: false,
		},
		{
			name: "folder's hierarchical references are not parent references",
			ref: corekccv1alpha1.ReferenceConfig{
				TFField: "organization",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key: "organizationRef",
					GVK: schema.GroupVersionKind{
						Group:   "resourcemanager.cnrm.cloud.google.com",
						Version: "v1beta1",
						Kind:    "Organization",
					},
				},
			},
			resource: &Resource{
				Resource: k8s.Resource{
					TypeMeta: metav1.TypeMeta{
						Kind: "Folder",
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{
					HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
						{
							Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
							Key:  "folderRef",
						},
						{
							Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
							Key:  "organizationRef",
						},
					},
				},
			},
			isParentRef: false,
		},
		{
			name: "hierarchical references should be treated as parent references for other resource kinds",
			ref: corekccv1alpha1.ReferenceConfig{
				TFField: "project",
				TypeConfig: corekccv1alpha1.TypeConfig{
					Key: "projectRef",
					GVK: schema.GroupVersionKind{
						Group:   "resourcemanager.cnrm.cloud.google.com",
						Version: "v1beta1",
						Kind:    "Project",
					},
				},
			},
			resource: &Resource{
				Resource: k8s.Resource{
					TypeMeta: metav1.TypeMeta{
						Kind: "Bar",
					},
				},
				ResourceConfig: v1alpha1.ResourceConfig{
					HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
						{
							Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
							Key:  "projectRef",
						},
						{
							Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
							Key:  "folderRef",
						},
						{
							Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
							Key:  "organizationRef",
						},
					},
				},
			},
			isParentRef: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := krmtotf.IsRequiredParentReference(tc.ref, tc.resource)
			if result != tc.isParentRef {
				t.Fatalf("got %v, want %v", result, tc.isParentRef)
			}
		})
	}
}
