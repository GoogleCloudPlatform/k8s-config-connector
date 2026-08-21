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

package dcl_test

import (
	"reflect"
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	testdclschemaloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/dclschemaloader"
	testservicemetadataloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemetadataloader"

	"github.com/google/go-cmp/cmp"
	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestToTypeConfig(t *testing.T) {
	tests := []struct {
		name                string
		dclRefExtensionElem map[interface{}]interface{}
		tc                  *corekccv1alpha1.TypeConfig
		hasError            bool
	}{
		{
			name: "non-parent referenced resource type",
			dclRefExtensionElem: map[interface{}]interface{}{
				"resource": "Test1/Foo",
				"field":    "name",
			},
			tc: &corekccv1alpha1.TypeConfig{
				TargetField: "name",
				GVK: schema.GroupVersionKind{
					Group:   "test1.cnrm.cloud.google.com",
					Version: "v1alpha1",
					Kind:    "Test1Foo",
				},
				Key: "fooRef",
			},
		},
		{
			name: "parent referenced resource type",
			dclRefExtensionElem: map[interface{}]interface{}{
				"resource": "Test1/Foo",
				"field":    "name",
				"parent":   true,
			},
			tc: &corekccv1alpha1.TypeConfig{
				TargetField: "name",
				GVK: schema.GroupVersionKind{
					Group:   "test1.cnrm.cloud.google.com",
					Version: "v1alpha1",
					Kind:    "Test1Foo",
				},
				Key:    "fooRef",
				Parent: true,
			},
		},
		{
			name: "the referenced resource type has an unrecognized service",
			dclRefExtensionElem: map[interface{}]interface{}{
				"resource": "SomeUnknownService/Foo",
				"field":    "name",
				"parent":   true,
			},
			hasError: true,
		},
		{
			name: "required 'field' attribute is not defined",
			dclRefExtensionElem: map[interface{}]interface{}{
				"resource": "Test1/Foo",
			},
			hasError: true,
		},
		{
			name: "required 'resource' attribute is not defined",
			dclRefExtensionElem: map[interface{}]interface{}{
				"field": "name",
			},
			hasError: true,
		},
	}
	smLoader := testservicemetadataloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := dcl.ToTypeConfig(tc.dclRefExtensionElem, smLoader)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, but expect to get an error on converting to TypeConfig")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.tc, res) {
				t.Fatalf("unexpected TypeConfig diff (-want +got): \n%v", cmp.Diff(tc.tc, res))
			}
		})
	}
}

func TestGetReferenceTypeConfigs(t *testing.T) {
	tests := []struct {
		name        string
		schema      *openapi.Schema
		typeConfigs []corekccv1alpha1.TypeConfig
	}{
		{
			name: "one non-parent type config",
			schema: &openapi.Schema{
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"resource": "Test1/Foo",
							"field":    "name",
						},
					},
				},
			},
			typeConfigs: []corekccv1alpha1.TypeConfig{
				{
					TargetField: "name",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Foo",
					},
					Key: "fooRef",
				},
			},
		},
		{
			name: "one parent type config",
			schema: &openapi.Schema{
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"resource": "Test1/Foo",
							"field":    "name",
							"parent":   true,
						},
					},
				},
			},
			typeConfigs: []corekccv1alpha1.TypeConfig{
				{
					TargetField: "name",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Foo",
					},
					Key:    "fooRef",
					Parent: true,
				},
			},
		},
		{
			name: "multiple non-parent type configs",
			schema: &openapi.Schema{
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"resource": "Test1/Foo",
							"field":    "name",
						},
						map[interface{}]interface{}{
							"resource": "Test2/Baz",
							"field":    "id",
						},
					},
				},
			},
			typeConfigs: []corekccv1alpha1.TypeConfig{
				{
					TargetField: "name",
					GVK: schema.GroupVersionKind{
						Group:   "test1.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test1Foo",
					},
					Key: "fooRef",
				},
				{
					TargetField: "id",
					GVK: schema.GroupVersionKind{
						Group:   "test2.cnrm.cloud.google.com",
						Version: "v1alpha1",
						Kind:    "Test2Baz",
					},
					Key: "bazRef",
				},
			},
		},
	}

	smLoader := testservicemetadataloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := dcl.GetReferenceTypeConfigs(tc.schema, smLoader)
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.typeConfigs, res) {
				t.Fatalf("unexpected diff in result (-want +got): \n%v", cmp.Diff(tc.typeConfigs, res))
			}
		})
	}
}

func TestGetHierarchicalReferencesForGVK(t *testing.T) {
	tests := []struct {
		name     string
		gvk      schema.GroupVersionKind
		expected []corekccv1alpha1.HierarchicalReference
	}{
		{
			name: "resource with no hierarchical reference",
			gvk: schema.GroupVersionKind{
				Group:   "test5.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test5NoHierarchicalRef",
			},
			expected: nil,
		},
		{
			name: "resource with project reference",
			gvk: schema.GroupVersionKind{
				Group:   "test5.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test5ProjectRef",
			},
			expected: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
					Key:  "projectRef",
				},
			},
		},
		{
			name: "resource with folder reference",
			gvk: schema.GroupVersionKind{
				Group:   "test5.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test5FolderRef",
			},
			expected: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
					Key:  "folderRef",
				},
			},
		},
		{
			name: "resource with organization reference",
			gvk: schema.GroupVersionKind{
				Group:   "test5.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test5OrganizationRef",
			},
			expected: []corekccv1alpha1.HierarchicalReference{
				{
					Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
					Key:  "organizationRef",
				},
			},
		},
		{
			name: "resource with multiple references",
			gvk: schema.GroupVersionKind{
				Group:   "test5.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test5MultipleRefs",
			},
			expected: []corekccv1alpha1.HierarchicalReference{
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
		{
			name: "resource with multiple references, but only two",
			gvk: schema.GroupVersionKind{
				Group:   "test5.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test5TwoRefs",
			},
			expected: []corekccv1alpha1.HierarchicalReference{
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
	}

	smLoader := dclmetadata.NewFromServiceList(testservicemetadataloader.FakeServiceMetadataWithHierarchicalResources())
	dclSchemaLoader := testdclschemaloader.New(dclSchemaMap())
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := dcl.GetHierarchicalReferencesForGVK(tc.gvk, smLoader, dclSchemaLoader)
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Fatalf("unexpected diff in result (-want +got): \n%v", cmp.Diff(tc.expected, actual))
			}
		})
	}
}

func TestGetHierarchicalReferenceConfigForMultiParentResource(t *testing.T) {
	tests := []struct {
		name      string
		schema    *openapi.Schema
		expected  []corekccv1alpha1.HierarchicalReference
		shouldErr bool
	}{
		{
			name: "no parent field",
			schema: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"field": {Type: "string"},
				},
			},
			shouldErr: true,
		},
		{
			name: "has project field, but no parent field",
			schema: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"field": {Type: "string"},
					"project": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Project",
								},
							},
						},
					},
				},
			},
			shouldErr: true,
		},
		{
			name: "has parent field",
			schema: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"field": {Type: "string"},
					"parent": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Project",
								},
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Folder",
								},
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Organization",
								},
							},
						},
					},
				},
			},
			expected: []corekccv1alpha1.HierarchicalReference{
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
		{
			name: "has parent field that can reference only two types of resources",
			schema: &openapi.Schema{
				Properties: map[string]*openapi.Schema{
					"field": {Type: "string"},
					"parent": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Folder",
								},
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/Organization",
								},
							},
						},
					},
				},
			},
			expected: []corekccv1alpha1.HierarchicalReference{
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
	}

	smLoader := dclmetadata.NewFromServiceList(testservicemetadataloader.FakeServiceMetadataWithHierarchicalResources())
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := dcl.GetHierarchicalReferenceConfigForMultiParentResource(tc.schema, smLoader)
			if tc.shouldErr {
				if err == nil {
					t.Fatalf("got no error, but wanted one")
				}
				return
			}
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Fatalf("unexpected diff in result (-want +got): \n%v", cmp.Diff(tc.expected, actual))
			}
		})
	}
}

func TestGetHierarchicalRefFromConfigForMultiParentResource(t *testing.T) {
	tests := []struct {
		name               string
		schema             *openapi.Schema
		config             map[string]interface{}
		expectedVal        interface{}
		expectedTypeConfig *corekccv1alpha1.TypeConfig
		hasError           bool
	}{
		{
			name: "no hierarchical reference in config (empty config)",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Project",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Folder",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Organization",
						},
					},
				},
			},
			expectedVal:        nil,
			expectedTypeConfig: nil,
		},
		{
			name: "no hierarchical reference in config",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Project",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Folder",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Organization",
						},
					},
				},
			},
			config: map[string]interface{}{
				"field": "val",
			},
			expectedVal:        nil,
			expectedTypeConfig: nil,
		},
		{
			name: "project reference in config",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Project",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Folder",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Organization",
						},
					},
				},
			},
			config: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"name": "project",
				},
			},
			expectedVal: map[string]interface{}{
				"name": "project",
			},
			expectedTypeConfig: &corekccv1alpha1.TypeConfig{
				TargetField: "name",
				GVK: schema.GroupVersionKind{
					Group:   "resourcemanager.cnrm.cloud.google.com",
					Version: "v1beta1",
					Kind:    "Project",
				},
				Key:    "projectRef",
				Parent: true,
			},
		},
		{
			name: "external project reference in config",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Project",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Folder",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Organization",
						},
					},
				},
			},
			config: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"external": "project",
				},
			},
			expectedVal: map[string]interface{}{
				"external": "project",
			},
			expectedTypeConfig: &corekccv1alpha1.TypeConfig{
				TargetField: "name",
				GVK: schema.GroupVersionKind{
					Group:   "resourcemanager.cnrm.cloud.google.com",
					Version: "v1beta1",
					Kind:    "Project",
				},
				Key:    "projectRef",
				Parent: true,
			},
		},
		{
			name: "namespaced project reference in config",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Project",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Folder",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Organization",
						},
					},
				},
			},
			config: map[string]interface{}{
				"projectRef": map[string]interface{}{
					"namespace": "project-namespace",
					"name":      "project",
				},
			},
			expectedVal: map[string]interface{}{
				"namespace": "project-namespace",
				"name":      "project",
			},
			expectedTypeConfig: &corekccv1alpha1.TypeConfig{
				TargetField: "name",
				GVK: schema.GroupVersionKind{
					Group:   "resourcemanager.cnrm.cloud.google.com",
					Version: "v1beta1",
					Kind:    "Project",
				},
				Key:    "projectRef",
				Parent: true,
			},
		},
		{
			name: "folder reference in config",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Project",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Folder",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Organization",
						},
					},
				},
			},
			config: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"name": "folder",
				},
			},
			expectedVal: map[string]interface{}{
				"name": "folder",
			},
			expectedTypeConfig: &corekccv1alpha1.TypeConfig{
				TargetField: "name",
				GVK: schema.GroupVersionKind{
					Group:   "resourcemanager.cnrm.cloud.google.com",
					Version: "v1beta1",
					Kind:    "Folder",
				},
				Key:    "folderRef",
				Parent: true,
			},
		},
		{
			name: "organization reference in config",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Project",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Folder",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Organization",
						},
					},
				},
			},
			config: map[string]interface{}{
				"organizationRef": map[string]interface{}{
					"external": "organization",
				},
			},
			expectedVal: map[string]interface{}{
				"external": "organization",
			},
			expectedTypeConfig: &corekccv1alpha1.TypeConfig{
				TargetField: "name",
				GVK: schema.GroupVersionKind{
					Group:   "resourcemanager.cnrm.cloud.google.com",
					Version: "v1beta1",
					Kind:    "Organization",
				},
				Key:    "organizationRef",
				Parent: true,
			},
		},
		{
			name: "multiple hierarchical references in config",
			schema: &openapi.Schema{
				Type: "string",
				Extension: map[string]interface{}{
					"x-dcl-references": []interface{}{
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Project",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Folder",
						},
						map[interface{}]interface{}{
							"field":    "name",
							"parent":   true,
							"resource": "Cloudresourcemanager/Organization",
						},
					},
				},
			},
			config: map[string]interface{}{
				"folderRef": map[string]interface{}{
					"external": "folder",
				},
				"organizationRef": map[string]interface{}{
					"external": "organization",
				},
			},
			hasError: true,
		},
	}

	smLoader := dclmetadata.NewFromServiceList(testservicemetadataloader.FakeServiceMetadataWithHierarchicalResources())
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			val, typeConfig, err := dcl.GetHierarchicalRefFromConfigForMultiParentResource(tc.config, tc.schema, smLoader)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got no error, but wanted one")
				}
				return
			}
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.expectedVal, val) {
				t.Fatalf("unexpected diff in resulting val (-want +got): \n%v", cmp.Diff(tc.expectedVal, val))
			}
			if !reflect.DeepEqual(tc.expectedTypeConfig, typeConfig) {
				t.Fatalf("unexpected diff in resulting type config (-want +got): \n%v", cmp.Diff(tc.expectedTypeConfig, typeConfig))
			}
		})
	}
}

func dclSchemaMap() map[string]*openapi.Schema {
	return map[string]*openapi.Schema{
		"test5_beta_nohierarchicalref": &openapi.Schema{
			Type: "object",
			Properties: map[string]*openapi.Schema{
				"name": &openapi.Schema{
					Type: "string",
				},
			},
		},
		"test5_beta_projectref": &openapi.Schema{
			Type: "object",
			Properties: map[string]*openapi.Schema{
				"project": &openapi.Schema{
					Type: "string",
					Extension: map[string]interface{}{
						"x-dcl-references": []interface{}{
							map[interface{}]interface{}{
								"field":    "name",
								"parent":   true,
								"resource": "Cloudresourcemanager/Project",
							},
						},
					},
				},
			},
		},
		"test5_beta_folderref": &openapi.Schema{
			Type: "object",
			Properties: map[string]*openapi.Schema{
				"folder": &openapi.Schema{
					Type: "string",
					Extension: map[string]interface{}{
						"x-dcl-references": []interface{}{
							map[interface{}]interface{}{
								"field":    "name",
								"parent":   true,
								"resource": "Cloudresourcemanager/Folder",
							},
						},
					},
				},
			},
		},
		"test5_beta_organizationref": &openapi.Schema{
			Type: "object",
			Properties: map[string]*openapi.Schema{
				"organization": &openapi.Schema{
					Type: "string",
					Extension: map[string]interface{}{
						"x-dcl-references": []interface{}{
							map[interface{}]interface{}{
								"field":    "name",
								"parent":   true,
								"resource": "Cloudresourcemanager/Organization",
							},
						},
					},
				},
			},
		},
		"test5_beta_multiplerefs": &openapi.Schema{
			Type: "object",
			Properties: map[string]*openapi.Schema{
				"parent": &openapi.Schema{
					Type: "string",
					Extension: map[string]interface{}{
						"x-dcl-references": []interface{}{
							map[interface{}]interface{}{
								"field":    "name",
								"parent":   true,
								"resource": "Cloudresourcemanager/Project",
							},
							map[interface{}]interface{}{
								"field":    "name",
								"parent":   true,
								"resource": "Cloudresourcemanager/Folder",
							},
							map[interface{}]interface{}{
								"field":    "name",
								"parent":   true,
								"resource": "Cloudresourcemanager/Organization",
							},
						},
					},
				},
			},
		},
		"test5_beta_tworefs": &openapi.Schema{
			Type: "object",
			Properties: map[string]*openapi.Schema{
				"parent": &openapi.Schema{
					Type: "string",
					Extension: map[string]interface{}{
						"x-dcl-references": []interface{}{
							map[interface{}]interface{}{
								"field":    "name",
								"parent":   true,
								"resource": "Cloudresourcemanager/Folder",
							},
							map[interface{}]interface{}{
								"field":    "name",
								"parent":   true,
								"resource": "Cloudresourcemanager/Organization",
							},
						},
					},
				},
			},
		},

		// Add the following to the list of fake DCL schemas to allow for our
		// test to test resources that reference hierarchical resources
		// (e.g. "Cloudresourcemanager/Project").
		"cloudresourcemanager_ga_project": &openapi.Schema{},
		"cloudresourcemanager_ga_folder":  &openapi.Schema{},
	}
}
