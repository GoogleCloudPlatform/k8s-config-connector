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

package container_test

import (
	"reflect"
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension/container"
	testdclschemaloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/dclschemaloader"
	testservicemetadataloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemetadataloader"

	"github.com/google/go-cmp/cmp"
	"github.com/nasa9084/go-openapi"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
)

func TestGetContainersForGVK(t *testing.T) {
	tests := []struct {
		name      string
		gvk       k8sschema.GroupVersionKind
		expected  []corekccv1alpha1.Container
		shouldErr bool
	}{
		{
			name: "resource with project container",
			gvk: k8sschema.GroupVersionKind{
				Group:   "test4.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test4ProjectContainer",
			},
			expected: []corekccv1alpha1.Container{
				{
					TFField: "project",
					Type:    corekccv1alpha1.ContainerTypeProject,
				},
			},
		},
		{
			name: "resource with organization container",
			gvk: k8sschema.GroupVersionKind{
				Group:   "test4.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test4OrganizationContainer",
			},
			expected: []corekccv1alpha1.Container{
				{
					TFField: "organization",
					Type:    corekccv1alpha1.ContainerTypeOrganization,
				},
			},
		},
		{
			name: "resource with no x-dcl-parent-container extension",
			gvk: k8sschema.GroupVersionKind{
				Group:   "test4.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test4NoContainer",
			},
			expected: nil,
		},
		{
			name: "resource with no x-dcl-parent-container extension but should support container annotations",
			gvk: k8sschema.GroupVersionKind{
				Group:   "test4.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test4NoContainerExtensionButSupportsContainers",
			},
			expected: []corekccv1alpha1.Container{
				{
					TFField: "project",
					Type:    corekccv1alpha1.ContainerTypeProject,
				},
			},
		},
		{
			name: "resource with no x-dcl-parent-container extension, and should support container annotations, but has no hierarchical references",
			gvk: k8sschema.GroupVersionKind{
				Group:   "test4.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "Test4NoContainerExtensionAndSupportsContainersButHasNoHierarchicalRefs",
			},
			shouldErr: true,
		},
	}

	schemaLoader := testdclschemaloader.New(dclSchemaMap)
	smLoader := testservicemetadataloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := container.GetContainersForGVK(tc.gvk, smLoader, schemaLoader)
			if tc.shouldErr {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Fatalf("got unexpected containers (-want +got): \n%v", cmp.Diff(tc.expected, actual))
			}
		})
	}
}

var dclSchemaMap = map[string]*openapi.Schema{
	"test4_beta_projectcontainer": {
		Type: "object",
		Properties: map[string]*openapi.Schema{
			"project": {
				Type: "string",
			},
			"name": {
				Type: "string",
			},
		},
		Extension: map[string]interface{}{
			"x-dcl-parent-container": "project",
		},
	},
	"test4_beta_organizationcontainer": {
		Type: "object",
		Properties: map[string]*openapi.Schema{
			"organization": {
				Type: "string",
			},
			"name": {
				Type: "string",
			},
		},
		Extension: map[string]interface{}{
			"x-dcl-parent-container": "organization",
		},
	},
	"test4_beta_nocontainer": {
		Type: "object",
		Properties: map[string]*openapi.Schema{
			"project": {
				Type: "string",
			},
			"name": {
				Type: "string",
			},
		},
	},
	"test4_beta_nocontainerextensionbutsupportscontainers": {
		Type: "object",
		Properties: map[string]*openapi.Schema{
			"project": {
				Type: "string",
			},
			"name": {
				Type: "string",
			},
		},
	},
	"test4_beta_nocontainerextensionandsupportscontainersbuthasnohierarchicalrefs": {
		Type: "object",
		Properties: map[string]*openapi.Schema{
			"name": {
				Type: "string",
			},
		},
	},
}
