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

package webhook

import (
	"testing"

	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	testdclschemaloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/dclschemaloader"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	testservicemetadataloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemetadataloader"

	"github.com/google/go-cmp/cmp"
	"github.com/nasa9084/go-openapi"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestHandleContainerAnnotationsForDCLBasedResources(t *testing.T) {
	tests := []struct {
		name   string
		obj    *unstructured.Unstructured
		newObj *unstructured.Unstructured
		ns     *corev1.Namespace
		schema *openapi.Schema
		denied bool
	}{
		{
			name: "no defaulting if resource supports neither containers nor hierarchical references",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6NoContainerOrHierarchicalRef",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6NoContainerOrHierarchicalRef",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
			},
		},

		// Test defaulting of container annotations for resources which only
		// support container annotations.
		// TODO(b/186159460): Delete the following tests once all resources
		// support hierarchical references.
		{
			name: "no defaulting if resource already has annotation (project-scoped resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyContainer",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyContainer",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
		},
		{
			name: "default resource-level annotation from namespace annotation (project-scoped resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyContainer",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyContainer",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
		},
		{
			name: "default resource-level annotation from namespace annotation (folder-scoped resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyContainer",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyContainer",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "folder",
				},
			},
		},
		{
			name: "default resource-level annotation from namespace annotation (organization-scoped resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyContainer",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyContainer",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.OrgIDAnnotation: "organization-id-from-namespace-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.OrgIDAnnotation: "organization-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "organization",
				},
			},
		},
		{
			name: "default resource-level annotation from namespace name if namespace has no annotation (project-scoped resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyContainer",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyContainer",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "namespace-name",
						},
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
		},
		{
			name: "deny resource if no namespace annotation (non-project-scoped resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyContainer",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "folder",
				},
			},
			denied: true,
		},

		// Test defaulting of hierarchical references for resources which
		// support both container annotations and hierarchical references.
		{
			name: "no defaulting if resource already has reference (project-scoped resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6BothContainerAndHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"name": "project-id-from-spec",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6BothContainerAndHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"name": "project-id-from-spec",
						},
					},
				},
			},
			schema: &openapi.Schema{
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
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from resource annotation (project-scoped resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6BothContainerAndHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6BothContainerAndHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project-id-from-resource-annotation",
						},
					},
				},
			},
			schema: &openapi.Schema{
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
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from namespace annotation (project-scoped resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6BothContainerAndHierarchicalRef",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6BothContainerAndHierarchicalRef",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project-id-from-namespace-annotation",
						},
					},
				},
			},
			schema: &openapi.Schema{
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
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from namespace name (project-scoped resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6BothContainerAndHierarchicalRef",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6BothContainerAndHierarchicalRef",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "namespace-name",
						},
					},
				},
			},
			schema: &openapi.Schema{
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
				Extension: map[string]interface{}{
					"x-dcl-parent-container": "project",
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from resource annotation (project-scoped resource which supports both container annotations and hierarchical references, but doesn't have a x-dcl-parent-container extension)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6BothContainerAndHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6BothContainerAndHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project-id-from-resource-annotation",
						},
					},
				},
			},
			schema: &openapi.Schema{
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
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},

		// Test defaulting of hierarchical references for resources which
		// only support hierarchical references.
		{
			name: "no defaulting if resource already has reference (project-scoped resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"name": "project-id-from-spec",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"name": "project-id-from-spec",
						},
					},
				},
			},
			schema: &openapi.Schema{
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
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from namespace annotation (project-scoped resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project-id-from-namespace-annotation",
						},
					},
				},
			},
			schema: &openapi.Schema{
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
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from namespace name (project-scoped resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "namespace-name",
						},
					},
				},
			},
			schema: &openapi.Schema{
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
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from namespace annotation (multi-parent resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"external": "folder-id-from-namespace-annotation",
						},
					},
				},
			},
			schema: &openapi.Schema{
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
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from namespace name (multi-parent resource which can be project-scoped and which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "namespace-name",
						},
					},
				},
			},
			schema: &openapi.Schema{
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
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
		},
		{
			name: "deny resource if more than one namespace annotation (multi-parent resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			schema: &openapi.Schema{
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
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
						k8s.OrgIDAnnotation:    "org-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			denied: true,
		},
		{
			name: "deny resource if no namespace annotation (non-project-scoped resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			schema: &openapi.Schema{
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
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
			denied: true,
		},
		{
			name: "allow resource even if no annotations found on resource/namespace as long as resource has reference (multi-parent resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"name": "folder-id-from-spec",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test6.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test6OnlyHierarchicalRef",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"name": "folder-id-from-spec",
						},
					},
				},
			},
			schema: &openapi.Schema{
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
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
		},
	}

	smLoader := dclmetadata.NewFromServiceList(testservicemetadataloader.FakeServiceMetadataWithHierarchicalResources())
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dclSchemaKey := testdclschemaloader.DCLSchemaKeyForGVK(t, tc.obj.GroupVersionKind(), smLoader)
			dclSchemaMap := map[string]*openapi.Schema{
				dclSchemaKey: tc.schema,

				// Add the following to the list of fake DCL schemas to allow for our
				// test to test resources that reference hierarchical resources
				// (e.g. "Cloudresourcemanager/Project").
				"cloudresourcemanager_ga_project": {},
				"cloudresourcemanager_ga_folder":  {},
			}
			dclSchemaLoader := testdclschemaloader.New(dclSchemaMap)
			_, err := handleContainerAnnotationsForDCLBasedResources(tc.obj, tc.ns, dclSchemaLoader, smLoader)
			if tc.denied {
				if err == nil {
					t.Fatalf("expected request to be denied, but was allowed. Response:\n%v", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("request was unexpectedly denied. Response:\n%v", err)
			}
			diff := cmp.Diff(tc.newObj, tc.obj)
			if diff != "" {
				t.Fatalf("unexpected diff in the response (-want +got):\n%v", diff)
			}
		})
	}
}

func TestHandleContainerAnnotationsForTFBasedResources(t *testing.T) {
	tests := []struct {
		name string
		obj  *unstructured.Unstructured
		ns   *corev1.Namespace

		newObj *unstructured.Unstructured
		denied bool
	}{
		{
			name: "no defaulting if resource supports neither containers nor hierarchical references",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4NoParentResource",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4NoParentResource",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},

		// Test defaulting of container annotations for resources which only
		// support container annotations.
		// TODO(b/193177782): Delete the following tests once all resources
		// support hierarchical references.
		{
			name: "no defaulting if resource already has annotation (project-scoped resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResourceWithOnlyContainerSupport",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResourceWithOnlyContainerSupport",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default resource-level annotation from namespace annotation (project-scoped resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResourceWithOnlyContainerSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResourceWithOnlyContainerSupport",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default resource-level annotation from namespace name if namespace has no annotation (project-scoped resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResourceWithOnlyContainerSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResourceWithOnlyContainerSupport",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "namespace-name",
						},
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
		},
		{
			name: "no defaulting if resource already has annotation (multi-parent resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResourceWithOnlyContainerSupport",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.FolderIDAnnotation: "folder-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResourceWithOnlyContainerSupport",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.FolderIDAnnotation: "folder-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default resource-level annotation from namespace annotation (multi-parent resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResourceWithOnlyContainerSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResourceWithOnlyContainerSupport",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "deny resource if more than one namespace annotation (multi-parent resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResourceWithOnlyContainerSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
						k8s.OrgIDAnnotation:    "org-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			denied: true,
		},
		{
			name: "deny resource if no namespace annotation (non-project-scoped resource which only supports container annotations)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResourceWithOnlyContainerSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
			denied: true,
		},

		// Test defaulting of hierarchical references for resources which
		// support both container annotations and hierarchical references.
		{
			name: "no defaulting if resource already has reference (project-scoped resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResource",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"name": "project-id-from-spec",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResource",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"name": "project-id-from-spec",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from resource annotation (project-scoped resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResource",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResource",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.ProjectIDAnnotation: "project-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project-id-from-resource-annotation",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from namespace annotation (project-scoped resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResource",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResource",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project-id-from-namespace-annotation",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from namespace name (project-scoped resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResource",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResource",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "namespace-name",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from resource annotation (multi-parent resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResource",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.FolderIDAnnotation: "folder-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResource",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.FolderIDAnnotation: "folder-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"external": "folder-id-from-resource-annotation",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "deny resource if more than one resource annotation (multi-parent resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResource",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							k8s.FolderIDAnnotation: "folder-id-from-resource-annotation",
							k8s.OrgIDAnnotation:    "org-id-from-resource-annotation",
						},
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
						k8s.OrgIDAnnotation:    "org-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			denied: true,
		},
		{
			name: "default reference from namespace annotation (multi-parent resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResource",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResource",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"external": "folder-id-from-namespace-annotation",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "deny resource if more than one namespace annotation (multi-parent resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResource",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
						k8s.OrgIDAnnotation:    "org-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			denied: true,
		},
		{
			name: "deny resource if no namespace annotation (non-project-scoped resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResource",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
			denied: true,
		},
		{
			name: "allow resource even if no annotations found on resource/namespace as long as resource has reference (multi-parent resource which supports both container annotations and hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResource",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"name": "project-id-from-spec",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResource",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"name": "project-id-from-spec",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
		},

		// Test defaulting of hierarchical references for resources which
		// only support hierarchical references.
		{
			name: "no defaulting if resource already has reference (project-scoped resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResourceWithOnlyHierarchicalReferenceSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"name": "project-id-from-spec",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResourceWithOnlyHierarchicalReferenceSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"name": "project-id-from-spec",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from namespace annotation (project-scoped resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResourceWithOnlyHierarchicalReferenceSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResourceWithOnlyHierarchicalReferenceSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "project-id-from-namespace-annotation",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.ProjectIDAnnotation: "project-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from namespace name (project-scoped resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResourceWithOnlyHierarchicalReferenceSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4ProjectScopedResourceWithOnlyHierarchicalReferenceSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "namespace-name",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
		},
		{
			name: "default reference from namespace annotation (multi-parent resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResourceWithOnlyHierarchicalReferenceSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResourceWithOnlyHierarchicalReferenceSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"external": "folder-id-from-namespace-annotation",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
		},
		{
			name: "deny resource if more than one namespace annotation (multi-parent resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResourceWithOnlyHierarchicalReferenceSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						k8s.FolderIDAnnotation: "folder-id-from-namespace-annotation",
						k8s.OrgIDAnnotation:    "org-id-from-namespace-annotation",
					},
					Name: "namespace-name",
				},
			},
			denied: true,
		},
		{
			name: "deny resource if no namespace annotation (non-project-scoped resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResourceWithOnlyHierarchicalReferenceSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
			denied: true,
		},
		{
			name: "allow resource even if no annotations found on resource/namespace as long as resource has reference (multi-parent resource which only supports hierarchical references)",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResourceWithOnlyHierarchicalReferenceSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"name": "folder-id-from-spec",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test4.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test4MultiParentResourceWithOnlyHierarchicalReferenceSupport",
					"metadata": map[string]interface{}{
						"name": "resource-name",
					},
					"spec": map[string]interface{}{
						"folderRef": map[string]interface{}{
							"name": "folder-id-from-spec",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "namespace-name",
				},
			},
		},
	}

	smLoader := testservicemappingloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := handleContainerAnnotationsForTFBasedResources(tc.obj, tc.ns, smLoader)
			if tc.denied {
				if err == nil {
					t.Fatalf("expected request to be denied, but was allowed. Response:\n%v", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("request was unexpectedly denied. Response:\n%v", err)
			}
			diff := cmp.Diff(tc.newObj, tc.obj)
			if diff != "" {
				t.Fatalf("unexpected diff in the response (-want +got):\n%v", diff)
			}
		})
	}
}
