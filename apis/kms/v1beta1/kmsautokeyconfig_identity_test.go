// Copyright 2026 Google LLC
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

package v1beta1

import (
	"context"
	"strings"
	"testing"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestNewAutokeyConfigIdentity(t *testing.T) {
	ctx := context.Background()

	// Set up mock K8s client with a Folder and a Project
	folderObj := &unstructured.Unstructured{}
	folderObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Folder",
	})
	folderObj.SetName("my-folder-name")
	folderObj.SetNamespace("my-namespace")
	if err := unstructured.SetNestedField(folderObj.Object, "12345", "spec", "resourceID"); err != nil {
		t.Fatalf("failed to set nested field: %v", err)
	}

	projectObj := &unstructured.Unstructured{}
	projectObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Project",
	})
	projectObj.SetName("my-project-name")
	projectObj.SetNamespace("my-namespace")
	if err := unstructured.SetNestedField(projectObj.Object, "my-project-id", "spec", "resourceID"); err != nil {
		t.Fatalf("failed to set nested field: %v", err)
	}

	scheme := runtime.NewScheme()
	scheme.AddKnownTypeWithName(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "FolderList",
	}, &unstructured.UnstructuredList{})
	scheme.AddKnownTypeWithName(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ProjectList",
	}, &unstructured.UnstructuredList{})

	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(folderObj, projectObj).Build()

	tests := []struct {
		name        string
		obj         *KMSAutokeyConfig
		expected    string
		expectError string
	}{
		{
			name: "no externalRef, folderRef external",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					FolderRef: &refsv1beta1.FolderRef{
						External: "folders/12345",
					},
				},
			},
			expected: "folders/12345/autokeyConfig",
		},
		{
			name: "no externalRef, folderRef name (resolved)",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					FolderRef: &refsv1beta1.FolderRef{
						Name: "my-folder-name",
					},
				},
			},
			expected: "folders/12345/autokeyConfig",
		},
		{
			name: "no externalRef, projectRef external",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project-id",
					},
				},
			},
			expected: "projects/my-project-id/autokeyConfig",
		},
		{
			name: "no externalRef, projectRef name (resolved)",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						Name: "my-project-name",
					},
				},
			},
			expected: "projects/my-project-id/autokeyConfig",
		},
		{
			name: "no externalRef, both folderRef and projectRef set",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					FolderRef: &refsv1beta1.FolderRef{
						External: "folders/12345",
					},
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project-id",
					},
				},
			},
			expectError: "only one of spec.folderRef or spec.projectRef can be specified",
		},
		{
			name: "no externalRef, neither set",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{},
			},
			expectError: "either spec.folderRef or spec.projectRef must be specified",
		},
		{
			name: "externalRef (folder) matches spec folderRef external",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					FolderRef: &refsv1beta1.FolderRef{
						External: "folders/12345",
					},
				},
				Status: KMSAutokeyConfigStatus{
					ExternalRef: func() *string { s := "folders/12345/autokeyConfig"; return &s }(),
				},
			},
			expected: "folders/12345/autokeyConfig",
		},
		{
			name: "externalRef (folder) matches spec folderRef name (resolved)",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					FolderRef: &refsv1beta1.FolderRef{
						Name: "my-folder-name",
					},
				},
				Status: KMSAutokeyConfigStatus{
					ExternalRef: func() *string { s := "folders/12345/autokeyConfig"; return &s }(),
				},
			},
			expected: "folders/12345/autokeyConfig",
		},
		{
			name: "externalRef (folder) spec folderRef external changed",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					FolderRef: &refsv1beta1.FolderRef{
						External: "folders/54321",
					},
				},
				Status: KMSAutokeyConfigStatus{
					ExternalRef: func() *string { s := "folders/12345/autokeyConfig"; return &s }(),
				},
			},
			expectError: "parent changed, expect folders/12345, got 54321",
		},
		{
			name: "externalRef (folder) spec changed to projectRef",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project-id",
					},
				},
				Status: KMSAutokeyConfigStatus{
					ExternalRef: func() *string { s := "folders/12345/autokeyConfig"; return &s }(),
				},
			},
			expectError: "parent changed, expect folders/12345, got my-project-id",
		},
		{
			name: "externalRef (project) matches spec projectRef external",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/my-project-id",
					},
				},
				Status: KMSAutokeyConfigStatus{
					ExternalRef: func() *string { s := "projects/my-project-id/autokeyConfig"; return &s }(),
				},
			},
			expected: "projects/my-project-id/autokeyConfig",
		},
		{
			name: "externalRef (project) matches spec projectRef name (resolved)",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						Name: "my-project-name",
					},
				},
				Status: KMSAutokeyConfigStatus{
					ExternalRef: func() *string { s := "projects/my-project-id/autokeyConfig"; return &s }(),
				},
			},
			expected: "projects/my-project-id/autokeyConfig",
		},
		{
			name: "externalRef (project) spec projectRef external changed",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "projects/other-project-id",
					},
				},
				Status: KMSAutokeyConfigStatus{
					ExternalRef: func() *string { s := "projects/my-project-id/autokeyConfig"; return &s }(),
				},
			},
			expectError: "parent changed, expect projects/my-project-id, got other-project-id",
		},
		{
			name: "externalRef (project) spec changed to folderRef",
			obj: &KMSAutokeyConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-autokeyconfig",
					Namespace: "my-namespace",
				},
				Spec: KMSAutokeyConfigSpec{
					FolderRef: &refsv1beta1.FolderRef{
						External: "folders/12345",
					},
				},
				Status: KMSAutokeyConfigStatus{
					ExternalRef: func() *string { s := "projects/my-project-id/autokeyConfig"; return &s }(),
				},
			},
			expectError: "parent changed, expect projects/my-project-id, got 12345",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			identity, err := NewAutokeyConfigIdentity(ctx, reader, tc.obj)
			if tc.expectError != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tc.expectError)
				}
				if !strings.Contains(err.Error(), tc.expectError) {
					t.Errorf("expected error containing %q, got: %v", tc.expectError, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if identity.String() != tc.expected {
				t.Errorf("expected identity %q, got: %q", tc.expected, identity.String())
			}
		})
	}
}
