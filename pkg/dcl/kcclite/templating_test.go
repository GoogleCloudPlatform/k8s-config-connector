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

package kcclite_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/kcclite"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testdclschemaloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/dclschemaloader"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	testservicemetadataloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemetadataloader"
	"github.com/nasa9084/go-openapi"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	hierarchicalRefProject = corekccv1alpha1.HierarchicalReference{
		Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
		Key:  "projectRef",
	}
	hierarchicalRefFolder = corekccv1alpha1.HierarchicalReference{
		Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
		Key:  "folderRef",
	}
	hierarchicalRefOrganization = corekccv1alpha1.HierarchicalReference{
		Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
		Key:  "organizationRef",
	}
	hierarchicalRefBillingAccount = corekccv1alpha1.HierarchicalReference{
		Type: corekccv1alpha1.HierarchicalReferenceTypeBillingAccount,
		Key:  "billingAccountRef",
	}

	projectRefConfig = corekccv1alpha1.ReferenceConfig{
		TFField: "project",
		TypeConfig: corekccv1alpha1.TypeConfig{
			Key: "projectRef",
			GVK: schema.GroupVersionKind{
				Group:   "resourcemanager.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "Project",
			},
		},
	}
	folderRefConfig = corekccv1alpha1.ReferenceConfig{
		TFField: "folder",
		TypeConfig: corekccv1alpha1.TypeConfig{
			Key: "folderRef",
			GVK: schema.GroupVersionKind{
				Group:   "resourcemanager.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "Folder",
			},
			TargetField: "folder_id",
		},
	}
	organizationRefConfig = corekccv1alpha1.ReferenceConfig{
		TFField: "organization",
		TypeConfig: corekccv1alpha1.TypeConfig{
			Key: "organizationRef",
			GVK: schema.GroupVersionKind{
				Group:   "resourcemanager.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "Organization",
			},
		},
	}
	billingAccountRefConfig = corekccv1alpha1.ReferenceConfig{
		TFField: "billing_account",
		TypeConfig: corekccv1alpha1.TypeConfig{
			Key: "billingAccountRef",
			GVK: schema.GroupVersionKind{
				Group:   "billing.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "BillingAccount",
			},
		},
	}
)

func TestCanonicalizeReferencedResourceName(t *testing.T) {
	tests := []struct {
		name                 string
		template             string
		refResource          *k8s.Resource
		refResourceSchema    *openapi.Schema
		refResourceReference *unstructured.Unstructured
		expectedCanonName    string
		errorCheckFunc       func(t *testing.T, err error)
	}{
		{
			name:     "template requires just {{name}}",
			template: "{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test1.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test1Foo",
				},
			},
			expectedCanonName: "name",
		},
		{
			name:     "template requires top-level spec field",
			template: "fields/{{field}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test1.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test1Foo",
				},
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"field": {
						Type: "string",
					},
				},
			},
			expectedCanonName: "fields/val/names/name",
		},
		{
			name:     "template requires top-level spec field, but referenced resource doesn't have field in spec",
			template: "fields/{{field}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test1.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test1Foo",
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"field": {
						Type: "string",
					},
				},
			},
			errorCheckFunc: hasErrorCheckFunc,
		},
		{
			name:     "template requires parent of single-parent resource which only supports container annotations",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6OnlyContainer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project_id",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			expectedCanonName: "projects/project_id/names/name",
		},
		{
			name:     "template requires parent of single-parent resource with no hierarchical reference in spec",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6BothContainerAndHierarchicalRef",
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			errorCheckFunc: hasErrorCheckFunc,
		},
		{
			name:     "template requires parent of single-parent resource with external project reference",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6BothContainerAndHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"external": "project_id",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			expectedCanonName: "projects/project_id/names/name",
		},
		{
			name:     "template requires parent of single-parent resource with external project reference that is set to a path",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6BothContainerAndHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"external": "projects/project_id",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			expectedCanonName: "projects/project_id/names/name",
		},
		{
			name:     "template requires parent of single-parent resource with project reference to nonexistent Project",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6BothContainerAndHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"name": "project-name",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			errorCheckFunc: transDepNotFoundErrorCheckFunc,
		},
		{
			name:     "template requires parent of single-parent resource with project reference to non-ready Project",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6BothContainerAndHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"name": "project-name",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			refResourceReference: test.NewProjectUnstructured("project-name", "project_id", corev1.ConditionFalse),
			errorCheckFunc:       transDepNotReadyErrorCheckFunc,
		},
		{
			name:     "template requires parent of single-parent resource with project reference",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6BothContainerAndHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"name": "project-name",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			refResourceReference: test.NewProjectUnstructured("project-name", "project_id", corev1.ConditionTrue),
			expectedCanonName:    "projects/project_id/names/name",
		},
		{
			name:     "template requires parent of single-parent resource with project reference that resolved to a path",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6BothContainerAndHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"name": "project-name",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			refResourceReference: test.NewProjectUnstructured("project-name", "projects/project_id", corev1.ConditionTrue),
			expectedCanonName:    "projects/project_id/names/name",
		},
		{
			name:     "template requires parent of multi-parent resource with no hierarchical reference in spec",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6OnlyHierarchicalRef",
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
			errorCheckFunc: hasErrorCheckFunc,
		},
		{
			name:     "template requires parent of multi-parent resource with external folder reference",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6OnlyHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"external": "folder_id",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
			expectedCanonName: "folders/folder_id/names/name",
		},
		{
			name:     "template requires parent of multi-parent resource with external folder reference that is set to a path",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6OnlyHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"external": "folders/folder_id",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
			expectedCanonName: "folders/folder_id/names/name",
		},
		{
			name:     "template requires parent of multi-parent resource with folder reference to nonexistent Folder",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6OnlyHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-name",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
			errorCheckFunc: transDepNotFoundErrorCheckFunc,
		},
		{
			name:     "template requires parent of multi-parent resource with folder reference to non-ready Folder",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6OnlyHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-name",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
			refResourceReference: test.NewFolderUnstructured("folder-name", "folder_id", corev1.ConditionFalse),
			errorCheckFunc:       transDepNotReadyErrorCheckFunc,
		},
		{
			name:     "template requires parent of multi-parent resource with folder reference",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6OnlyHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-name",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
			refResourceReference: test.NewFolderUnstructured("folder-name", "folder_id", corev1.ConditionTrue),
			expectedCanonName:    "folders/folder_id/names/name",
		},
		{
			name:     "template requires parent of multi-parent resource with folder reference that resolves to a path",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6OnlyHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-name",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
			refResourceReference: test.NewFolderUnstructured("folder-name", "folders/folder_id", corev1.ConditionTrue),
			expectedCanonName:    "folders/folder_id/names/name",
		},
		{
			name:     "template requires parent of multi-parent resource with external billing account reference that is set to a path",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "test6.cnrm.cloud.google.com/v1alpha1",
					Kind:       "Test6OnlyHierarchicalRef",
				},
				Spec: map[string]interface{}{
					"billingAccountRef": map[string]interface{}{
						"external": "billingAccounts/billing_account_id",
					},
				},
			},
			refResourceSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
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
								map[interface{}]interface{}{
									"field":    "name",
									"parent":   true,
									"resource": "Cloudresourcemanager/BillingAccount",
								},
							},
						},
					},
				},
			},
			expectedCanonName: "billingAccounts/billing_account_id/names/name",
		},
	}

	smLoader := dclmetadata.NewFromServiceList(testservicemetadataloader.FakeServiceMetadataWithHierarchicalResources())
	serviceMappingLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			c := mgr.GetClient()
			if err := testcontroller.EnsureNamespaceExists(c, testID); err != nil {
				t.Fatal(err)
			}
			tc.refResource.SetNamespace(testID)
			if tc.refResourceReference != nil {
				tc.refResourceReference.SetNamespace(testID)
				test.EnsureObjectExists(t, tc.refResourceReference, c)
			}

			schemaKey := testdclschemaloader.DCLSchemaKeyForGVK(t, tc.refResource.GroupVersionKind(), smLoader)
			schemaMap := map[string]*openapi.Schema{
				schemaKey: tc.refResourceSchema,
			}
			schemaLoader := testdclschemaloader.New(schemaMap)

			canonName, err := kcclite.CanonicalizeReferencedResourceName("name", tc.template, tc.refResource, smLoader, schemaLoader, serviceMappingLoader, c)
			if tc.errorCheckFunc != nil {
				tc.errorCheckFunc(t, err)
				return
			}
			if err != nil {
				t.Fatalf("got error, wanted none: %v", err)
			}
			if canonName != tc.expectedCanonName {
				t.Fatalf("got %v, want %v", canonName, tc.expectedCanonName)
			}
		})
	}
}

func TestCanonicalizeReferencedResourceNameForTFBasedResource(t *testing.T) {
	tests := []struct {
		name                 string
		template             string
		refResource          *k8s.Resource
		refResourceConfig    corekccv1alpha1.ResourceConfig
		refResourceReference *unstructured.Unstructured
		expectedCanonName    string
		errorCheckFunc       func(t *testing.T, err error)
	}{
		{
			name:     "template requires just {{name}}",
			template: "{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
			},
			expectedCanonName: "name",
		},
		{
			name:     "template requires top-level spec field",
			template: "fields/{{field}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"field": "val",
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
			},
			expectedCanonName: "fields/val/names/name",
		},
		{
			name:     "template requires top-level spec field, but referenced resource doesn't have field in spec",
			template: "fields/{{field}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
			},
			errorCheckFunc: hasErrorCheckFunc,
		},
		{
			name:     "template requires parent of single-parent resource which only supports container annotations",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project_id",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
				},
			},
			expectedCanonName: "projects/project_id/names/name",
		},
		{
			name:     "template requires parent of single-parent resource with no hierarchical reference in spec",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
				},
			},
			errorCheckFunc: hasErrorCheckFunc,
		},
		{
			name:     "template requires parent of single-parent resource with external project reference",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"external": "project_id",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
				},
			},
			expectedCanonName: "projects/project_id/names/name",
		},
		{
			name:     "template requires parent of single-parent resource with external project reference that is set to a path",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"external": "projects/project_id",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
				},
			},
			expectedCanonName: "projects/project_id/names/name",
		},
		{
			name:     "template requires parent of single-parent resource with project reference to nonexistent Project",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"name": "project-name",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
				},
			},
			errorCheckFunc: transDepNotFoundErrorCheckFunc,
		},
		{
			name:     "template requires parent of single-parent resource with project reference to non-ready Project",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"name": "project-name",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
				},
			},
			refResourceReference: test.NewProjectUnstructured("project-name", "project_id", corev1.ConditionFalse),
			errorCheckFunc:       transDepNotReadyErrorCheckFunc,
		},
		{
			name:     "template requires parent of single-parent resource with project reference",
			template: "projects/{{project}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"projectRef": map[string]interface{}{
						"name": "project-name",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
				},
			},
			refResourceReference: test.NewProjectUnstructured("project-name", "project_id", corev1.ConditionTrue),
			expectedCanonName:    "projects/project_id/names/name",
		},
		{
			name:     "template requires parent of single-parent resource with folder reference that resolved to a path",
			template: "folders/{{folder}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-name",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeFolder},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefFolder,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					folderRefConfig,
				},
			},
			refResourceReference: test.NewFolderUnstructured("folder-name", "folders/folder_id", corev1.ConditionTrue),
			expectedCanonName:    "folders/folder_id/names/name",
		},
		{
			name:     "template requires parent of multi-parent resource with no hierarchical reference in spec",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
					{Type: corekccv1alpha1.ContainerTypeFolder},
					{Type: corekccv1alpha1.ContainerTypeOrganization},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
					hierarchicalRefFolder,
					hierarchicalRefOrganization,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
					folderRefConfig,
					organizationRefConfig,
				},
			},
			errorCheckFunc: hasErrorCheckFunc,
		},
		{
			name:     "template requires parent of multi-parent resource with external folder reference",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"external": "folder_id",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
					{Type: corekccv1alpha1.ContainerTypeFolder},
					{Type: corekccv1alpha1.ContainerTypeOrganization},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
					hierarchicalRefFolder,
					hierarchicalRefOrganization,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
					folderRefConfig,
					organizationRefConfig,
				},
			},
			expectedCanonName: "folders/folder_id/names/name",
		},
		{
			name:     "template requires parent of multi-parent resource with external folder reference that is set to a path",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"external": "folders/folder_id",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
					{Type: corekccv1alpha1.ContainerTypeFolder},
					{Type: corekccv1alpha1.ContainerTypeOrganization},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
					hierarchicalRefFolder,
					hierarchicalRefOrganization,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
					folderRefConfig,
					organizationRefConfig,
				},
			},
			expectedCanonName: "folders/folder_id/names/name",
		},
		{
			name:     "template requires parent of multi-parent resource with folder reference to nonexistent Folder",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-name",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
					{Type: corekccv1alpha1.ContainerTypeFolder},
					{Type: corekccv1alpha1.ContainerTypeOrganization},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
					hierarchicalRefFolder,
					hierarchicalRefOrganization,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
					folderRefConfig,
					organizationRefConfig,
				},
			},
			errorCheckFunc: transDepNotFoundErrorCheckFunc,
		},
		{
			name:     "template requires parent of multi-parent resource with folder reference to non-ready Folder",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-name",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
					{Type: corekccv1alpha1.ContainerTypeFolder},
					{Type: corekccv1alpha1.ContainerTypeOrganization},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
					hierarchicalRefFolder,
					hierarchicalRefOrganization,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
					folderRefConfig,
					organizationRefConfig,
				},
			},
			refResourceReference: test.NewFolderUnstructured("folder-name", "folder_id", corev1.ConditionFalse),
			errorCheckFunc:       transDepNotReadyErrorCheckFunc,
		},
		{
			name:     "template requires parent of multi-parent resource with folder reference",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-name",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
					{Type: corekccv1alpha1.ContainerTypeFolder},
					{Type: corekccv1alpha1.ContainerTypeOrganization},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
					hierarchicalRefFolder,
					hierarchicalRefOrganization,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
					folderRefConfig,
					organizationRefConfig,
				},
			},
			refResourceReference: test.NewFolderUnstructured("folder-name", "folder_id", corev1.ConditionTrue),
			expectedCanonName:    "folders/folder_id/names/name",
		},
		{
			name:     "template requires parent of multi-parent resource with folder reference that resolves to a path",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"folderRef": map[string]interface{}{
						"name": "folder-name",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
					{Type: corekccv1alpha1.ContainerTypeFolder},
					{Type: corekccv1alpha1.ContainerTypeOrganization},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
					hierarchicalRefFolder,
					hierarchicalRefOrganization,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
					folderRefConfig,
					organizationRefConfig,
				},
			},
			refResourceReference: test.NewFolderUnstructured("folder-name", "folders/folder_id", corev1.ConditionTrue),
			expectedCanonName:    "folders/folder_id/names/name",
		},
		{
			name:     "template requires parent of multi-parent resource with external billing account reference that is set to a path",
			template: "{{parent}}/names/{{name}}",
			refResource: &k8s.Resource{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "tfonly.cnrm.cloud.google.com/v1alpha1",
					Kind:       "TFOnlyKind",
				},
				Spec: map[string]interface{}{
					"billingAccountRef": map[string]interface{}{
						"external": "billingAccounts/billing_account_id",
					},
				},
			},
			refResourceConfig: corekccv1alpha1.ResourceConfig{
				Name: "tf_only_kind",
				Kind: "TFOnlyKind",
				Containers: []corekccv1alpha1.Container{
					{Type: corekccv1alpha1.ContainerTypeProject},
					{Type: corekccv1alpha1.ContainerTypeFolder},
					{Type: corekccv1alpha1.ContainerTypeOrganization},
				},
				HierarchicalReferences: []corekccv1alpha1.HierarchicalReference{
					hierarchicalRefProject,
					hierarchicalRefFolder,
					hierarchicalRefOrganization,
					hierarchicalRefBillingAccount,
				},
				ResourceReferences: []corekccv1alpha1.ReferenceConfig{
					projectRefConfig,
					folderRefConfig,
					organizationRefConfig,
					billingAccountRefConfig,
				},
			},
			expectedCanonName: "billingAccounts/billing_account_id/names/name",
		},
	}

	smLoader := testservicemetadataloader.NewForUnitTest()
	schemaLoader := testdclschemaloader.New(make(map[string]*openapi.Schema))
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testID := testvariable.NewUniqueID()
			c := mgr.GetClient()
			if err := testcontroller.EnsureNamespaceExists(c, testID); err != nil {
				t.Fatal(err)
			}
			tc.refResource.SetNamespace(testID)
			if tc.refResourceReference != nil {
				tc.refResourceReference.SetNamespace(testID)
				test.EnsureObjectExists(t, tc.refResourceReference, c)
			}

			// Define a custom ServiceMapping for this test to contain TF-only
			// resources (i.e. kinds only defined in the TF service mappings
			// and not in the DCL service metadata).
			serviceMapping := corekccv1alpha1.ServiceMapping{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "cnrm-system",
					Name:      "tfonly.cnrm.cloud.google.com",
				},
				Spec: corekccv1alpha1.ServiceMappingSpec{
					Name:            "tfonly",
					ServiceHostName: "tfonly",
					Version:         "v1alpha1",
					Resources: []v1alpha1.ResourceConfig{
						tc.refResourceConfig,
					},
				},
			}
			serviceMappings := append(test.FakeServiceMappingsWithHierarchicalResources(), serviceMapping)
			serviceMappingLoader := servicemappingloader.NewFromServiceMappings(serviceMappings)

			canonName, err := kcclite.CanonicalizeReferencedResourceName("name", tc.template, tc.refResource, smLoader, schemaLoader, serviceMappingLoader, c)
			if tc.errorCheckFunc != nil {
				tc.errorCheckFunc(t, err)
				return
			}
			if err != nil {
				t.Fatalf("got error, wanted none: %v", err)
			}
			if canonName != tc.expectedCanonName {
				t.Fatalf("got %v, want %v", canonName, tc.expectedCanonName)
			}
		})
	}
}
