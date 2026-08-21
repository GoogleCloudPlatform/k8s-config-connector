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

package test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"

	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const Namespace = "namespace-1"

func FakeCRDs() []*apiextensions.CustomResourceDefinition {
	return []*apiextensions.CustomResourceDefinition{
		CRDForGVK(schema.GroupVersionKind{
			Group:   "test1.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "Test1Foo",
		}),
		CRDForGVK(schema.GroupVersionKind{
			Group:   "test1.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "Test1Bar",
		}),
		CRDForGVK(schema.GroupVersionKind{
			// Unique group
			Group:   "test2.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "Test2Baz",
		}),
		CRDForGVK(schema.GroupVersionKind{
			Group:   "test3.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "Test3UserSpecifiedResourceIDKind",
		}),
		CRDForGVK(schema.GroupVersionKind{
			Group:   "test3.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "Test3ServerGeneratedResourceIDKind",
		}),
		CRDForGVK(schema.GroupVersionKind{
			Group:   "test4.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "Test4DCLResourceServerGeneratedResourceIDKind",
		}),
		CRDForGVK(schema.GroupVersionKind{
			Group:   "test4.cnrm.cloud.google.com",
			Version: "v1alpha1",
			Kind:    "Test4DCLResourceUserSpecifiedResourceIDKind",
		}),
		CRDForGVK(schema.GroupVersionKind{
			Group:   "test5.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "TestKindWithObservedState",
		}),
	}
}

// FakeCRDsWithHierarchicalResources returns a CRD list which includes
// hierarchical resources to allow for the testing of resources that reference
// hierarchical resources (e.g. "Project")
func FakeCRDsWithHierarchicalResources() []*apiextensions.CustomResourceDefinition {
	return append(FakeCRDs(),
		CRDForGVK(schema.GroupVersionKind{
			Group:   "resourcemanager.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "Project",
		}),
		CRDForGVK(schema.GroupVersionKind{
			Group:   "resourcemanager.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "Folder",
		}),
	)
}

func FakeServiceMappings() []v1alpha1.ServiceMapping {
	var test1FooReconciliationIntervalInSeconds uint32 = 100
	return []v1alpha1.ServiceMapping{
		{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "cnrm-system",
				Name:      "test1.cnrm.cloud.google.com",
			},
			Spec: v1alpha1.ServiceMappingSpec{
				Name:            "test1",
				ServiceHostName: "test1",
				Version:         "v1alpha1",
				Resources: []v1alpha1.ResourceConfig{
					{
						Name:                            "foo",
						Kind:                            "Test1Foo",
						ReconciliationIntervalInSeconds: &test1FooReconciliationIntervalInSeconds,
					},
					{
						Name: "bar",
						Kind: "Test1Bar",
					},
					{
						Name: "fake_tf_based_resource",
						Kind: "Test1FakeTFBasedResource",
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "cnrm-system",
				Name:      "test2.cnrm.cloud.google.com",
			},
			Spec: v1alpha1.ServiceMappingSpec{
				Name:            "test2",
				ServiceHostName: "test2",
				Version:         "v1alpha1",
				Resources: []v1alpha1.ResourceConfig{
					{
						Name: "baz",
						Kind: "Test2Baz",
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "cnrm-system",
				Name:      "test3.cnrm.cloud.google.com",
			},
			Spec: v1alpha1.ServiceMappingSpec{
				Name:            "test3",
				ServiceHostName: "test3",
				Version:         "v1alpha1",
				Resources: []v1alpha1.ResourceConfig{
					{
						Name: "user_specified_resource_id_kind",
						Kind: "Test3UserSpecifiedResourceIDKind",
						ResourceID: v1alpha1.ResourceID{
							TargetField: "resource_id_field",
						},
						MetadataMapping: v1alpha1.MetadataMapping{
							Name: "resource_id_field",
						},
						IDTemplate: "{{resource_id_field}}",
					},
					{
						Name: "server_generated_resource_id_kind_with_value_template",
						Kind: "Test3ServerGeneratedResourceIDKind",
						ResourceID: v1alpha1.ResourceID{
							TargetField:   "resource_id_field",
							ValueTemplate: "values/{{value}}",
						},
						ServerGeneratedIDField: "resource_id_field",
						IDTemplate:             "{{resource_id_field}}",
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "cnrm-system",
				Name:      "test4.cnrm.cloud.google.com",
			},
			Spec: v1alpha1.ServiceMappingSpec{
				Name:            "test4",
				ServiceHostName: "test4",
				Version:         "v1alpha1",
				Resources: []v1alpha1.ResourceConfig{
					{
						Kind: "Test4ProjectScopedResource",
						Containers: []v1alpha1.Container{
							{Type: v1alpha1.ContainerTypeProject},
						},
						HierarchicalReferences: []v1alpha1.HierarchicalReference{
							{
								Type: v1alpha1.HierarchicalReferenceTypeProject,
								Key:  "projectRef",
							},
						},
					},
					{
						Kind: "Test4ProjectScopedResourceWithOnlyContainerSupport",
						Containers: []v1alpha1.Container{
							{Type: v1alpha1.ContainerTypeProject},
						},
					},
					{
						Kind: "Test4ProjectScopedResourceWithOnlyHierarchicalReferenceSupport",
						Containers: []v1alpha1.Container{
							{Type: v1alpha1.ContainerTypeProject},
						},
						HierarchicalReferences: []v1alpha1.HierarchicalReference{
							{
								Type: v1alpha1.HierarchicalReferenceTypeProject,
								Key:  "projectRef",
							},
						},
					},
					{
						Kind: "Test4MultiParentResource",
						Containers: []v1alpha1.Container{
							{Type: v1alpha1.ContainerTypeFolder},
							{Type: v1alpha1.ContainerTypeOrganization},
						},
						HierarchicalReferences: []v1alpha1.HierarchicalReference{
							{
								Type: v1alpha1.HierarchicalReferenceTypeFolder,
								Key:  "folderRef",
							},
							{
								Type: v1alpha1.HierarchicalReferenceTypeOrganization,
								Key:  "organizationRef",
							},
						},
					},
					{
						Kind: "Test4MultiParentResourceWithOnlyContainerSupport",
						Containers: []v1alpha1.Container{
							{Type: v1alpha1.ContainerTypeFolder},
							{Type: v1alpha1.ContainerTypeOrganization},
						},
					},
					{
						Kind: "Test4MultiParentResourceWithOnlyHierarchicalReferenceSupport",
						HierarchicalReferences: []v1alpha1.HierarchicalReference{
							{
								Type: v1alpha1.HierarchicalReferenceTypeFolder,
								Key:  "folderRef",
							},
							{
								Type: v1alpha1.HierarchicalReferenceTypeOrganization,
								Key:  "organizationRef",
							},
						},
					},
					{
						Kind: "Test4NoParentResource",
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "cnrm-system",
				Name:      "test5.cnrm.cloud.google.com",
			},
			Spec: v1alpha1.ServiceMappingSpec{
				Name:            "test5",
				ServiceHostName: "test5",
				Version:         "v1beta1",
				Resources: []v1alpha1.ResourceConfig{
					{
						Kind: "TestKindWithObservedState",
					},
				},
			},
		},
	}
}

// FakeServiceMappingsWithHierarchicalResources returns a ServiceMapping list
// which includes hierarchical resources to allow for the testing of resources
// that reference hierarchical resources (e.g. "Project")
func FakeServiceMappingsWithHierarchicalResources() []v1alpha1.ServiceMapping {
	return append(FakeServiceMappings(),
		v1alpha1.ServiceMapping{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "cnrm-system",
				Name:      "resourcemanager.cnrm.cloud.google.com",
			},
			Spec: v1alpha1.ServiceMappingSpec{
				Name:            "ResourceManager",
				Version:         "v1beta1",
				ServiceHostName: "cloudresourcemanager.googleapis.com",
				Resources: []v1alpha1.ResourceConfig{
					{
						Kind: "Project",
						MetadataMapping: v1alpha1.MetadataMapping{
							Name: "project_id",
						},
						ResourceID: v1alpha1.ResourceID{
							TargetField: "project_id",
						},
					},
					{
						Kind: "Folder",
					},
				},
			},
		})
}

func NewBarUnstructured(name, ns string, readyStatus corev1.ConditionStatus) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
			"kind":       "Test1Bar",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					k8s.ProjectIDAnnotation: "my-project-1",
				},
				"name":      name,
				"namespace": ns,
			},
			"spec": map[string]interface{}{
				"location":  "test-location",
				"specField": "abc123",
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": readyStatus,
					},
				},
				"statusField": "foobar",
			},
		},
	}
}

func NewUnstructuredWithObservedState(name, ns string, readyStatus corev1.ConditionStatus) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "test5.cnrm.cloud.google.com/v1beta1",
			"kind":       "TestKindWithObservedState",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					k8s.ProjectIDAnnotation: "my-project-1",
				},
				"name":      name,
				"namespace": ns,
			},
			"spec": map[string]interface{}{
				"location":  "test-location",
				"specField": "abc123",
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": readyStatus,
					},
				},
				"observedState": map[string]interface{}{
					"statusField":          "foobar",
					"referenceTargetField": "reference-value",
				},
			},
		},
	}
}

func NewIAMServiceAccountUnstructured(name, namespace string) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
			"kind":       "IAMServiceAccount",
			"metadata": map[string]interface{}{
				"name":      name,
				"namespace": namespace,
			},
		},
	}

}

func NewProjectUnstructured(name, projectID string, readyStatus corev1.ConditionStatus) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "resourcemanager.cnrm.cloud.google.com/v1beta1",
			"kind":       "Project",
			"metadata": map[string]interface{}{
				"name": name,
			},
			"spec": map[string]interface{}{
				"resourceID": projectID,
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": readyStatus,
					},
				},
			},
		},
	}
}

func NewFolderUnstructured(name, folderID string, readyStatus corev1.ConditionStatus) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "resourcemanager.cnrm.cloud.google.com/v1beta1",
			"kind":       "Folder",
			"metadata": map[string]interface{}{
				"name": name,
			},
			"spec": map[string]interface{}{
				"resourceID": folderID,
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": readyStatus,
					},
				},
				"folderId": folderID,
			},
		},
	}
}

func NewSecretUnstructured(name, ns string, stringData map[string]interface{}) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"name":      name,
				"namespace": ns,
			},
			"stringData": stringData,
		},
	}
}

func EnsureObjectsExist(t *testing.T, objs []*unstructured.Unstructured, c client.Client) {
	t.Helper()
	for _, obj := range objs {
		EnsureObjectExists(t, obj, c)
	}
}

func EnsureObjectExists(t *testing.T, obj *unstructured.Unstructured, c client.Client) {
	if err := c.Create(context.Background(), obj); err != nil {
		if !errors.IsAlreadyExists(err) {
			t.Errorf("error creating resource %v %v/%v: %v",
				obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
	}
}

func CRDForGVK(gvk schema.GroupVersionKind) *apiextensions.CustomResourceDefinition {
	singular := strings.ToLower(gvk.Kind)
	plural := text.Pluralize(singular)
	preserveUnknownFields := true
	return &apiextensions.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apiextensions.k8s.io/v1",
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%v.%v", plural, gvk.Group),
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group: gvk.Group,
			Names: apiextensions.CustomResourceDefinitionNames{
				Plural:   plural,
				Singular: singular,
				Kind:     gvk.Kind,
			},
			Scope: apiextensions.NamespaceScoped,
			Versions: []apiextensions.CustomResourceDefinitionVersion{
				{
					Name:    gvk.Version,
					Storage: true,
					Served:  true,
					Schema: &apiextensions.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
							Type:                   "object",
							XPreserveUnknownFields: &preserveUnknownFields,
						},
					},
				},
			},
		},
	}
}
