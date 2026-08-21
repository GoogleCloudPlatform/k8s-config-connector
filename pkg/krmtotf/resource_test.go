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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
)

func TestResource_GetImportID(t *testing.T) {
	tests := []struct {
		name                 string
		kind                 string
		apiVersion           string
		rc                   *v1alpha1.ResourceConfig
		metadataName         string
		spec                 map[string]interface{}
		status               map[string]interface{}
		referencedResources  []*unstructured.Unstructured
		expected             string
		assertGotExpectedErr func(t *testing.T, err error)
	}{
		{
			name: "with fields from spec and status",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "{{spec_field}}/{{status_field}}",
			},
			expected: "abc123/foobar",
		},
		{
			name: "with field from container",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "project/{{project}}",
				Containers: []v1alpha1.Container{
					{
						Type:    v1alpha1.ContainerTypeProject,
						TFField: "project",
					},
				},
			},
			expected: "project/my-project-1",
		},
		{
			name: "with reference's ID",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "bar/{{bar_ref}}",
				ResourceReferences: []v1alpha1.ReferenceConfig{
					{
						TFField: "bar_ref",
						TypeConfig: v1alpha1.TypeConfig{
							Key: "barRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
						},
					},
				},
			},
			spec: map[string]interface{}{
				"barRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-ref1", "", corev1.ConditionTrue),
			},
			expected: "bar/my-ref1",
		},
		{
			name: "with reference's ID, but referenced resource is not found",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "bar/{{bar_ref}}",
				ResourceReferences: []v1alpha1.ReferenceConfig{
					{
						TFField: "bar_ref",
						TypeConfig: v1alpha1.TypeConfig{
							Key: "barRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
						},
					},
				},
			},
			spec: map[string]interface{}{
				"barRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			assertGotExpectedErr: assertIsReferenceNotFoundError,
		},
		{
			name: "with reference's ID, but referenced resource is not ready",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "bar/{{bar_ref}}",
				ResourceReferences: []v1alpha1.ReferenceConfig{
					{
						TFField: "bar_ref",
						TypeConfig: v1alpha1.TypeConfig{
							Key: "barRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
						},
					},
				},
			},
			spec: map[string]interface{}{
				"barRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-ref1", "", corev1.ConditionFalse),
			},
			assertGotExpectedErr: assertIsReferenceNotReadyError,
		},
		{
			name: "with reference's status field",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "bar/{{bar_ref}}",
				ResourceReferences: []v1alpha1.ReferenceConfig{
					{
						TFField: "bar_ref",
						TypeConfig: v1alpha1.TypeConfig{
							Key: "barRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
							TargetField: "status_field",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"barRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-ref1", "", corev1.ConditionTrue),
			},
			expected: "bar/foobar",
		},
		{
			name: "with reference's status field, but referenced resource is not found",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "bar/{{bar_ref}}",
				ResourceReferences: []v1alpha1.ReferenceConfig{
					{
						TFField: "bar_ref",
						TypeConfig: v1alpha1.TypeConfig{
							Key: "barRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
							TargetField: "status_field",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"barRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			assertGotExpectedErr: assertIsReferenceNotFoundError,
		},
		{
			name: "with reference's status field, but referenced resource is not ready",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "bar/{{bar_ref}}",
				ResourceReferences: []v1alpha1.ReferenceConfig{
					{
						TFField: "bar_ref",
						TypeConfig: v1alpha1.TypeConfig{
							Key: "barRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
							TargetField: "status_field",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"barRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-ref1", "", corev1.ConditionFalse),
			},
			assertGotExpectedErr: assertIsReferenceNotReadyError,
		},
		{
			name: "with reference's spec field",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "bar/{{bar_ref}}",
				ResourceReferences: []v1alpha1.ReferenceConfig{
					{
						TFField: "bar_ref",
						TypeConfig: v1alpha1.TypeConfig{
							Key: "barRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
							TargetField: "spec_field",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"barRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-ref1", "", corev1.ConditionTrue),
			},
			expected: "bar/abc123",
		},
		{
			name: "with reference's spec field, but referenced resource is not found",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "bar/{{bar_ref}}",
				ResourceReferences: []v1alpha1.ReferenceConfig{
					{
						TFField: "bar_ref",
						TypeConfig: v1alpha1.TypeConfig{
							Key: "barRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
							TargetField: "spec_field",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"barRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			assertGotExpectedErr: assertIsReferenceNotFoundError,
		},
		{
			name: "with reference's spec field, but referenced resource is not ready",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "bar/{{bar_ref}}",
				ResourceReferences: []v1alpha1.ReferenceConfig{
					{
						TFField: "bar_ref",
						TypeConfig: v1alpha1.TypeConfig{
							Key: "barRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
							TargetField: "spec_field",
						},
					},
				},
			},
			spec: map[string]interface{}{
				"barRef": map[string]interface{}{
					"name": "my-ref1",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-ref1", "", corev1.ConditionFalse),
			},
			assertGotExpectedErr: assertIsReferenceNotReadyError,
		},
		{
			name: "server-generated ID from status",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "status_field",
			},
			expected: "foobar",
		},
		{
			name: "server-generated ID from status, but server-generated ID is not found",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "non_existent_status_field",
			},
			assertGotExpectedErr: assertIsServerGeneratedIDNotFoundError,
		},
		{
			name: "server-generated ID from status with template",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate:             "id/{{status_field}}",
				ServerGeneratedIDField: "status_field",
			},
			expected: "id/foobar",
		},
		{
			name: "server-generated ID from status with template, but server-generated ID is not found",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate:             "id/{{non_existent_status_field}}",
				ServerGeneratedIDField: "non_existent_status_field",
			},
			assertGotExpectedErr: assertIsServerGeneratedIDNotFoundError,
		},
		{
			name: "no template implies {{project}}/{{resource_id}}",
			rc: &v1alpha1.ResourceConfig{
				MetadataMapping: v1alpha1.MetadataMapping{
					Name: "resource_id",
				},
				Containers: []v1alpha1.Container{
					{
						Type:    v1alpha1.ContainerTypeProject,
						TFField: "project",
					},
				},
			},
			metadataName: "my-resource",
			expected:     "my-project-1/my-resource",
		},
		{
			name: "regional resources map location to region field",
			rc: &v1alpha1.ResourceConfig{
				Locationality: gcp.Regional,
				IDTemplate:    "regions/{{region}}",
			},
			expected: "regions/test-location",
		},
		{
			name: "zonal resources map location to zone field",
			rc: &v1alpha1.ResourceConfig{
				Locationality: gcp.Zonal,
				IDTemplate:    "zones/{{zone}}",
			},
			expected: "zones/test-location",
		},
		{
			name: "optional spec field should resolve when specified",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "id/{{barField?}}",
			},
			spec: map[string]interface{}{
				"barField": "bar-value",
			},
			expected: "id/bar-value",
		},
		{
			name: "optional spec field should not resolve when not specified",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "id/{{barField?}}/another-field",
			},
			spec:     map[string]interface{}{},
			expected: "id//another-field",
		},
		{
			name: "'or' should pick first value when second is not present",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "id/[text-with-{{field1}}|{{field2}}]/another-string",
			},
			spec: map[string]interface{}{
				"field1": "field1-value",
			},
			expected: "id/text-with-field1-value/another-string",
		},
		{
			name: "'or' should pick second value when first is not present",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "id/[text-with-{{field1}}|{{field2}}]/another-string",
			},
			spec: map[string]interface{}{
				"field2": "field2-value",
			},
			expected: "id/field2-value/another-string",
		},
		{
			name: "'or' should pick first value when first and second are present",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "id/[text-with-{{field1}}|{{field2}}]/another-string",
			},
			spec: map[string]interface{}{
				"field1": "field1-value",
				"field2": "field2-value",
			},
			expected: "id/text-with-field1-value/another-string",
		},
		{
			name: "server-generated resourceID supported",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "resource_id_field_in_status",
				ResourceID: v1alpha1.ResourceID{
					TargetField: "resource_id_field_in_status",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "id-in-spec",
			},
			expected: "id-in-spec",
		},
		{
			name: "server-generated resourceID supported, with a value template",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "resource_id_field_in_status",
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "resource_id_field_in_status",
					ValueTemplate: "values/{{value}}",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "id-in-spec",
			},
			expected: "values/id-in-spec",
		},
		{
			name: "server-generated resourceID supported, with an ID template",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate:             "id/{{resource_id_field_in_status}}",
				ServerGeneratedIDField: "resource_id_field_in_status",
				ResourceID: v1alpha1.ResourceID{
					TargetField: "resource_id_field_in_status",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "id-in-spec",
			},
			expected: "id/id-in-spec",
		},
		{
			name: "server-generated resourceID supported, with a value " +
				"template and an ID template",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate:             "parents/parent/{{resource_id_field_in_status}}",
				ServerGeneratedIDField: "resource_id_field_in_status",
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "resource_id_field_in_status",
					ValueTemplate: "groups/group/id/{{value}}",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "id-in-spec",
			},
			expected: "parents/parent/groups/group/id/id-in-spec",
		},
		{
			name: "server-generated resourceID supported, with a value " +
				"template that contains the container",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "resource_id_field_in_status",
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "resource_id_field_in_status",
					ValueTemplate: "projects/{{project}}/id/{{value}}",
				},
				Containers: []v1alpha1.Container{
					{
						Type:    v1alpha1.ContainerTypeProject,
						TFField: "project",
					},
				},
			},
			spec: map[string]interface{}{
				"resourceID": "id-in-spec",
			},
			expected: "projects/my-project-1/id/id-in-spec",
		},
		{
			name: "server-generated resourceID supported, with a value " +
				"template that contains a spec field",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "resource_id_field_in_status",
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "resource_id_field_in_status",
					ValueTemplate: "groups/{{spec_field}}/id/{{value}}",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "id-in-spec",
				"specField":  "specTestValue",
			},
			expected: "groups/specTestValue/id/id-in-spec",
		},
		{
			name: "server-generated resourceID supported, with a value " +
				"template that contains a status field",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "resource_id_field_in_status",
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "resource_id_field_in_status",
					ValueTemplate: "groups/{{status_field}}/id/{{value}}",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "id-in-spec",
			},
			expected: "groups/foobar/id/id-in-spec",
		},
		{
			name: "server-generated resourceID supported, with a value " +
				"template containing a reference field",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "resource_id_field_in_status",
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "resource_id_field_in_status",
					ValueTemplate: "bar/{{bar_name}}/id/{{value}}",
				},
				ResourceReferences: []v1alpha1.ReferenceConfig{
					{
						TFField: "bar_name",
						TypeConfig: v1alpha1.TypeConfig{
							Key: "barRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "test1.cnrm.cloud.google.com",
								Version: "v1alpha1",
								Kind:    "Test1Bar",
							},
							TargetField: "spec_field",
							Parent:      true,
						},
					},
				},
			},
			spec: map[string]interface{}{
				"resourceID": "id-in-spec",
				"barRef": map[string]interface{}{
					"name": "my-bar-ref",
				},
			},
			referencedResources: []*unstructured.Unstructured{
				test.NewBarUnstructured("my-bar-ref", "", corev1.ConditionTrue),
			},
			expected: "bar/abc123/id/id-in-spec",
		},
		{
			name: "server-generated resourceID supported, with spec.resourceID " +
				"unspecified but serverGeneratedIDField in status specified",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "status_field",
				ResourceID: v1alpha1.ResourceID{
					TargetField: "status_field",
				},
			},
			expected: "foobar",
		},
		{
			name: "server-generated resourceID supported, with both " +
				"spec.resourceID and serverGeneratedIDField in status unspecified",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "non_existent_resource_id",
				ResourceID: v1alpha1.ResourceID{
					TargetField: "non_existent_resource_id",
				},
			},
			assertGotExpectedErr: assertIsServerGeneratedIDNotFoundError,
		},
		{
			name: "user-specified resourceID supported",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "id/{{resource_id}}",
				MetadataMapping: v1alpha1.MetadataMapping{
					Name: "resource_id",
				},
				ResourceID: v1alpha1.ResourceID{
					TargetField: "resource_id",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "user-specified-id",
			},
			expected: "id/user-specified-id",
		},
		{
			name: "user-specified resourceID supported, with spec.resourceID " +
				"unspecified but metadata.name specified",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "{{resource_id}}",
				MetadataMapping: v1alpha1.MetadataMapping{
					Name: "resource_id",
				},
				ResourceID: v1alpha1.ResourceID{
					TargetField: "resource_id",
				},
			},
			metadataName: "default-id",
			expected:     "default-id",
		},
		{
			name: "user-specified resourceID supported, with both " +
				"spec.resourceID and metadata.name unspecified",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "id/{{resource_id}}",
				MetadataMapping: v1alpha1.MetadataMapping{
					Name: "resource_id",
				},
				ResourceID: v1alpha1.ResourceID{
					TargetField: "resource_id",
				},
			},
			assertGotExpectedErr: hasError,
		},
		{
			name:       "with field from status.observedState",
			kind:       "TestKind",
			apiVersion: "test.cnrm.cloud.google.com/v1beta1",
			rc: &v1alpha1.ResourceConfig{
				IDTemplate: "ids/{{status_field}}",
			},
			status: map[string]interface{}{
				"observedState": map[string]interface{}{
					"statusField": "id-value",
				},
			},
			expected: "ids/id-value",
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
			if tc.kind != "" {
				r.Kind = tc.kind
			}
			if tc.apiVersion != "" {
				r.APIVersion = tc.apiVersion
			}
			r.ResourceConfig = *tc.rc
			if tc.metadataName != "" {
				r.SetName(tc.metadataName)
			}
			r.SetNamespace(testID)
			testcontroller.EnsureNamespaceExistsT(t, c, testID)
			bar := test.NewBarUnstructured("my-resource", testID, corev1.ConditionTrue)
			r.SetAnnotations(bar.GetAnnotations())
			r.Spec = tc.spec
			r.Status = tc.status
			if r.Spec == nil {
				r.Spec = bar.Object["spec"].(map[string]interface{})
			}
			if r.Status == nil {
				r.Status = bar.Object["status"].(map[string]interface{})
			}
			for _, obj := range tc.referencedResources {
				obj.SetNamespace(testID)
			}
			test.EnsureObjectsExist(t, tc.referencedResources, c)
			actual, err := r.GetImportID(c, smLoader)
			if tc.assertGotExpectedErr != nil {
				tc.assertGotExpectedErr(t, err)
			} else if err != nil {
				t.Fatalf("error getting import ID: %v", err)
			}
			if tc.expected != actual {
				t.Fatalf("expected: %v, actual %v", tc.expected, actual)
			}
		})
	}
}

func assertIsReferenceNotFoundError(t *testing.T, err error) {
	if _, ok := k8s.AsReferenceNotFoundError(err); !ok {
		t.Fatalf("expected error that can be unwrapped as '%v', but got: '%v'", reflect.TypeOf(&k8s.ReferenceNotFoundError{}), err)
	}
}

func assertIsReferenceNotReadyError(t *testing.T, err error) {
	if _, ok := k8s.AsReferenceNotReadyError(err); !ok {
		t.Fatalf("expected error that can be unwrapped as '%v', but got: '%v'", reflect.TypeOf(&k8s.ReferenceNotReadyError{}), err)
	}
}

func assertIsServerGeneratedIDNotFoundError(t *testing.T, err error) {
	if _, ok := k8s.AsServerGeneratedIDNotFoundError(err); !ok {
		t.Fatalf("expected error that can be unwrapped as '%v', but got: %v", reflect.TypeOf(&k8s.ServerGeneratedIDNotFoundError{}), err)
	}
}

func hasError(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("got nil, want an error")
	}
}

func TestResource_ValidateResourceIDIfSupported(t *testing.T) {
	tests := []struct {
		name     string
		rc       *v1alpha1.ResourceConfig
		spec     map[string]interface{}
		hasError bool
	}{
		{
			name: "resourceID not supported",
			rc:   &v1alpha1.ResourceConfig{},
		},
		{
			name: "unspecified resourceID",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: v1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			spec: map[string]interface{}{},
		},
		{
			name: "empty resourceID",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: v1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "",
			},
			hasError: true,
		},
		{
			name: "nonempty resourceID",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: v1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "test-resource-id",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := resourceSkeleton()
			r.ResourceConfig = *tc.rc
			r.SetName("my-resource")
			r.Spec = tc.spec
			if r.Spec == nil {
				r.Spec = map[string]interface{}{}
			}

			err := r.ValidateResourceIDIfSupported()
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, want an error")
				}
				return
			} else if err != nil {
				t.Fatalf("error validating the resource ID if supported: %v", err)
			}
		})
	}
}

func TestResource_ConstructServerGeneratedIDInStatusFromResourceID(t *testing.T) {
	tests := []struct {
		name           string
		rc             *v1alpha1.ResourceConfig
		spec           map[string]interface{}
		status         map[string]interface{}
		expectedSpec   map[string]interface{}
		expectedStatus map[string]interface{}
		expectedResult string
		hasError       bool
	}{
		{
			name: "both resourceID field and server-generated ID field are" +
				"unspecified",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "server_generated_id",
				ResourceID: v1alpha1.ResourceID{
					TargetField: "server_generated_id",
				},
			},
			spec:           map[string]interface{}{},
			status:         map[string]interface{}{},
			expectedSpec:   map[string]interface{}{},
			expectedStatus: map[string]interface{}{},
			expectedResult: "",
		},
		{
			name: "resourceID field is unspecified and server-generated ID " +
				"field is non-empty",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "server_generated_id",
				ResourceID: v1alpha1.ResourceID{
					TargetField: "server_generated_id",
				},
			},
			spec: map[string]interface{}{},
			status: map[string]interface{}{
				"serverGeneratedId": "server-generated-id",
			},
			expectedSpec: map[string]interface{}{},
			expectedStatus: map[string]interface{}{
				"serverGeneratedId": "server-generated-id",
			},
			expectedResult: "",
		},
		{
			name: "resourceID field is empty",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "server_generated_id",
				ResourceID: v1alpha1.ResourceID{
					TargetField: "server_generated_id",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "",
			},
			status: map[string]interface{}{
				"serverGeneratedId": "server-generated-id",
			},
			hasError: true,
		},
		{
			name: "resourceID field and server-generated ID field have " +
				"different non-empty values",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "server_generated_id",
				ResourceID: v1alpha1.ResourceID{
					TargetField: "server_generated_id",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "non-empty-resource-id",
			},
			status: map[string]interface{}{
				"serverGeneratedId": "non-empty-id-in-status",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "non-empty-resource-id",
			},
			expectedStatus: map[string]interface{}{
				"serverGeneratedId": "non-empty-id-in-status",
			},
			expectedResult: "non-empty-resource-id",
		},
		{
			name: "resourceID field has a non-empty value and " +
				"server-generated ID field is unspecified",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "server_generated_id",
				ResourceID: v1alpha1.ResourceID{
					TargetField: "server_generated_id",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "non-empty-resource-id",
			},
			status: map[string]interface{}{},
			expectedSpec: map[string]interface{}{
				"resourceID": "non-empty-resource-id",
			},
			expectedStatus: map[string]interface{}{},
			expectedResult: "non-empty-resource-id",
		},
		{
			name: "resourceID field has a non-empty value and the nested " +
				"server-generated ID field is empty",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "parent_field.server_generated_id",
				ResourceID: v1alpha1.ResourceID{
					TargetField: "parent_field.server_generated_id",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "non-empty-resource-id",
			},
			status: map[string]interface{}{
				"parentField": map[string]interface{}{
					"serverGeneratedId": "",
				},
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "non-empty-resource-id",
			},
			expectedStatus: map[string]interface{}{
				"parentField": map[string]interface{}{
					"serverGeneratedId": "",
				},
			},
			expectedResult: "non-empty-resource-id",
		},

		{
			name: "with a value template, resourceID field has a " +
				"non-empty value and server-generated ID field is empty",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "server_generated_id",
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "server_generated_id",
					ValueTemplate: "id/{{value}}",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "non-empty-resource-id",
			},
			status: map[string]interface{}{
				"serverGeneratedId": "",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "non-empty-resource-id",
			},
			expectedStatus: map[string]interface{}{
				"serverGeneratedId": "",
			},
			expectedResult: "id/non-empty-resource-id",
		},
		{
			name: "with a complex value template, resourceID field " +
				"has a non-empty value and server-generated ID field is empty",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "server_generated_id",
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "server_generated_id",
					ValueTemplate: "projects/{{project}}/groups/{{group_id}}/values/{{value}}",
				},
				Containers: []v1alpha1.Container{
					{
						Type:    v1alpha1.ContainerTypeProject,
						TFField: "project",
					},
				},
			},
			spec: map[string]interface{}{
				"resourceID": "resource-id",
				"groupId":    "test-group",
			},
			status: map[string]interface{}{
				"serverGeneratedId": "",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"groupId":    "test-group",
			},
			expectedStatus: map[string]interface{}{
				"serverGeneratedId": "",
			},
			expectedResult: "projects/my-project-1/groups/test-group/values/resource-id",
		},
		{
			name: "with a value template containing parent field, the resourceID " +
				"field has a non-empty value and server-generated ID field is empty",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "server_generated_id",
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "server_generated_id",
					ValueTemplate: "{{parent}}/values/{{value}}",
				},
				ResourceReferences: []v1alpha1.ReferenceConfig{
					{
						TFField: "parent",
						TypeConfig: v1alpha1.TypeConfig{
							Key: "parentRef",
							GVK: k8sschema.GroupVersionKind{
								Group:   "datacatalog.cnrm.cloud.google.com",
								Version: "v1beta1",
								Kind:    "DataCatalogTaxonomy",
							},
							TargetField: "name",
							Parent:      true,
						},
					},
				},
			},
			spec: map[string]interface{}{
				"resourceID": "resource-id",
				"parentRef": map[string]interface{}{
					"external": "projects/project/locations/us/taxonomies/tid",
				},
			},
			status: map[string]interface{}{
				"serverGeneratedId": "",
			},
			expectedSpec: map[string]interface{}{
				"resourceID": "resource-id",
				"parentRef": map[string]interface{}{
					"external": "projects/project/locations/us/taxonomies/tid",
				},
			},
			expectedStatus: map[string]interface{}{
				"serverGeneratedId": "",
			},
			expectedResult: "projects/project/locations/us/taxonomies/tid/values/resource-id",
		},
		{
			name: "resourceID field has a non-empty value and status is nil",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "server_generated_id",
				ResourceID: v1alpha1.ResourceID{
					TargetField: "server_generated_id",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "non-empty-resource-id",
			},
			status: nil,
			expectedSpec: map[string]interface{}{
				"resourceID": "non-empty-resource-id",
			},
			expectedStatus: nil,
			expectedResult: "non-empty-resource-id",
		},
	}

	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			c := mgr.GetClient()
			r := resourceSkeleton()
			r.ResourceConfig = *tc.rc
			r.SetName("test-resource")
			bar := test.NewBarUnstructured("test", "", corev1.ConditionTrue)
			r.SetAnnotations(bar.GetAnnotations())
			r.Spec = tc.spec
			r.Status = tc.status

			result, err := r.ConstructServerGeneratedIDInStatusFromResourceID(c, smLoader)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, want an error")
				}
				return
			} else if err != nil {
				t.Fatalf("error syncing the server-generated ID in "+
					"status from resource ID: %v", err)
			}
			if got, want := result, tc.expectedResult; got != want {
				t.Fatalf("got: %v, want: %v", got, want)
			}
			if got, want := r.Spec, tc.expectedSpec; !test.Equals(t, got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
			if got, want := r.Status, tc.expectedStatus; !test.Equals(t, got, want) {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestResource_GetServerGeneratedID(t *testing.T) {
	tests := []struct {
		name                 string
		kind                 string
		apiVersion           string
		rc                   *v1alpha1.ResourceConfig
		spec                 map[string]interface{}
		status               map[string]interface{}
		expectedID           string
		assertGotExpectedErr func(t *testing.T, err error)
	}{
		{
			name: "get server-generated ID",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "test_field",
			},
			status: map[string]interface{}{
				"testField": "test-id",
			},
			expectedID: "test-id",
		},
		{
			name:       "get server-generated ID from observed state",
			kind:       "TestKind",
			apiVersion: "test.cnrm.cloud.google.com/v1beta1",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "test_field",
			},
			status: map[string]interface{}{
				"observedState": map[string]interface{}{
					"testField": "test-id",
				},
			},
			expectedID: "test-id",
		},
		{
			name: "get server-generated ID from nested field",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "nested_field.test_field",
			},
			status: map[string]interface{}{
				"nestedField": map[string]interface{}{
					"testField": "nested-id",
				},
			},
			expectedID: "nested-id",
		},
		{
			name: "server-generated ID not found",
			rc: &v1alpha1.ResourceConfig{
				ServerGeneratedIDField: "test_field",
			},
			status: map[string]interface{}{
				"otherStatusField": "testValue",
			},
			expectedID:           "",
			assertGotExpectedErr: assertIsServerGeneratedIDNotFoundError,
		},
		{
			name: "get server-generated ID from resource ID",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{
				"resourceID": "test-id-in-spec",
			},
			expectedID: "test-id-in-spec",
		},
		{
			name: "server-generated ID not set in spec.resourceID or in status",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{
				"otherSpecField": "testValue",
			},
			status: map[string]interface{}{
				"otherStatusField": "testValue",
			},
			expectedID:           "",
			assertGotExpectedErr: assertIsServerGeneratedIDNotFoundError,
		},
		{
			name: "server-generated ID with a valueTemplate not set in " +
				"spec.resourceID but in status ",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "test_field",
					ValueTemplate: "values/{{value}}",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{
				"otherSpecField": "testValue",
			},
			status: map[string]interface{}{
				"testField": "values/test-id",
			},
			expectedID: "test-id",
		},
		{
			name: "empty spec.resourceID",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{
				"resourceID": "",
			},
			status:               map[string]interface{}{},
			expectedID:           "",
			assertGotExpectedErr: hasError,
		},
		{
			name: "spec.resourceID unspecified and serverGeneratedIDField in " +
				"status empty",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{},
			status: map[string]interface{}{
				"testField": "",
			},
			expectedID:           "",
			assertGotExpectedErr: hasError,
		},
		{
			name: "spec.resourceID unspecified and serverGeneratedIDField in " +
				"status not matching value template",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "test_field",
					ValueTemplate: "values/{{value}}",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{},
			status: map[string]interface{}{
				"testField": "test-123",
			},
			expectedID:           "",
			assertGotExpectedErr: hasError,
		},
		{
			name: "spec.resourceID unspecified and resource ID extracted from " +
				"serverGeneratedIDField in status empty",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "test_field",
					ValueTemplate: "values/{{value}}",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{},
			status: map[string]interface{}{
				"testField": "values/",
			},
			expectedID:           "",
			assertGotExpectedErr: hasError,
		},
		{
			name: "spec.resourceID specified but does not map to server-generated ID",
			rc: &v1alpha1.ResourceConfig{
				MetadataMapping: v1alpha1.MetadataMapping{
					Name: "name",
				},
				ResourceID: v1alpha1.ResourceID{
					TargetField: "name",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{
				"resourceID": "test-id-in-spec",
			},
			status: map[string]interface{}{
				"testField": "test-id-in-status",
			},
			expectedID: "test-id-in-status",
		},
		{
			name: "spec.resourceID specified but does not map to server-generated ID, " +
				"and serverGeneratedIDField in status not found",
			rc: &v1alpha1.ResourceConfig{
				MetadataMapping: v1alpha1.MetadataMapping{
					Name: "name",
				},
				ResourceID: v1alpha1.ResourceID{
					TargetField: "name",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{
				"resourceID": "test-id-in-spec",
			},
			status: map[string]interface{}{
				"otherStatusField": "testValue",
			},
			expectedID:           "",
			assertGotExpectedErr: assertIsServerGeneratedIDNotFoundError,
		},
		{
			name: "spec.resourceID specified and supports a value template, " +
				"but does not map to server-generated ID",
			rc: &v1alpha1.ResourceConfig{
				MetadataMapping: v1alpha1.MetadataMapping{
					Name: "name",
				},
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "name",
					ValueTemplate: "values/{{value}}",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{
				"resourceID": "test-id-in-spec",
			},
			status: map[string]interface{}{
				"testField": "test-id-in-status",
			},
			expectedID: "test-id-in-status",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := resourceSkeleton()
			if tc.kind != "" {
				r.Kind = tc.kind
			}
			if tc.apiVersion != "" {
				r.APIVersion = tc.apiVersion
			}
			r.ResourceConfig = *tc.rc
			r.SetName("my-resource")
			r.Spec = tc.spec
			if r.Spec == nil {
				r.Spec = map[string]interface{}{}
			}
			r.Status = tc.status
			if r.Status == nil {
				r.Status = map[string]interface{}{}
			}

			id, err := r.GetServerGeneratedID()
			if tc.assertGotExpectedErr != nil {
				tc.assertGotExpectedErr(t, err)
			} else if err != nil {
				t.Fatalf("error getting server-generated ID: %v", err)
			}
			if got, want := id, tc.expectedID; got != want {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}

func TestResource_GetResourceID(t *testing.T) {
	tests := []struct {
		name         string
		rc           *v1alpha1.ResourceConfig
		metadataName string
		spec         map[string]interface{}
		status       map[string]interface{}
		expectedID   string
		hasError     bool
	}{
		{
			name: "empty resourceID",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "",
			},
			hasError: true,
		},
		{
			name: "non-empty resource ID",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
			},
			spec: map[string]interface{}{
				"resourceID": "test-id",
			},
			expectedID: "test-id",
		},
		{
			name: "user-specified ID from metadata.name",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: v1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			metadataName: "user-specified-id",
			spec:         map[string]interface{}{},
			expectedID:   "user-specified-id",
		},
		{
			name: "user-specified ID not found",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
				MetadataMapping: v1alpha1.MetadataMapping{
					Name: "test_field",
				},
			},
			spec:     map[string]interface{}{},
			hasError: true,
		},
		{
			name: "server-generated ID from status",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{},
			status: map[string]interface{}{
				"testField": "id-in-status",
			},
			expectedID: "id-in-status",
		},
		{
			name: "server-generated ID with a value template from status",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "test_field",
					ValueTemplate: "values/{{value}}",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{},
			status: map[string]interface{}{
				"testField": "values/id-with-template-in-status",
			},
			expectedID: "id-with-template-in-status",
		},
		{
			name: "server-generated ID not found",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField: "test_field",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec:     map[string]interface{}{},
			status:   map[string]interface{}{},
			hasError: true,
		},
		{
			name: "server-generated ID from status but it doesn't match the " +
				"value template",
			rc: &v1alpha1.ResourceConfig{
				ResourceID: v1alpha1.ResourceID{
					TargetField:   "test_field",
					ValueTemplate: "values/{{value}}",
				},
				ServerGeneratedIDField: "test_field",
			},
			spec: map[string]interface{}{},
			status: map[string]interface{}{
				"testField": "incorrectly-formatted-id",
			},
			hasError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := resourceSkeleton()
			r.ResourceConfig = *tc.rc
			if tc.metadataName != "" {
				r.SetName(tc.metadataName)
			}
			r.Spec = tc.spec
			if r.Spec == nil {
				r.Spec = map[string]interface{}{}
			}
			r.Status = tc.status
			if r.Status == nil {
				r.Status = map[string]interface{}{}
			}

			id, err := r.GetResourceID()
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, want an error")
				}
				return
			} else if err != nil {
				t.Fatalf("error getting resource ID: %v", err)
			}
			if got, want := id, tc.expectedID; got != want {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}
