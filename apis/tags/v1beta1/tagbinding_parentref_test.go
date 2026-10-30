// Copyright 2025 Google LLC
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

	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestTagsTagBindingParentRef_Normalize(t *testing.T) {
	tests := []struct {
		name             string
		initial          TagsTagBindingParentRef
		expectedExternal string
		wantErr          bool
		wantErrSubstr    string
	}{
		{
			name: "External: projects/{project} with kind=Project",
			initial: TagsTagBindingParentRef{
				Kind:     "Project",
				External: "projects/my-project",
			},
			expectedExternal: "//cloudresourcemanager.googleapis.com/projects/my-project",
		},
		{
			name: "External: projects/{project} with no kind",
			initial: TagsTagBindingParentRef{
				External: "projects/my-project",
			},
			expectedExternal: "//cloudresourcemanager.googleapis.com/projects/my-project",
		},
		{
			name: "External: //cloudresourcemanager.googleapis.com/projects/{project} with kind=Project",
			initial: TagsTagBindingParentRef{
				Kind:     "Project",
				External: "//cloudresourcemanager.googleapis.com/projects/my-project",
			},
			expectedExternal: "//cloudresourcemanager.googleapis.com/projects/my-project",
		},
		{
			name: "External: //cloudresourcemanager.googleapis.com/projects/{project} with no kind",
			initial: TagsTagBindingParentRef{
				External: "//cloudresourcemanager.googleapis.com/projects/my-project",
			},
			expectedExternal: "//cloudresourcemanager.googleapis.com/projects/my-project",
		},
		{
			name: "External: projects/{project}/buckets/{bucket} with kind=StorageBucket",
			initial: TagsTagBindingParentRef{
				Kind:     "StorageBucket",
				External: "projects/my-project/buckets/my-bucket",
			},
			expectedExternal: "//storage.googleapis.com/projects/my-project/buckets/my-bucket",
		},
		{
			name: "External: projects/{project}/secrets/{secret} with kind=SecretManagerSecret",
			initial: TagsTagBindingParentRef{
				Kind:     "SecretManagerSecret",
				External: "projects/my-project/secrets/my-secret",
			},
			expectedExternal: "//secretmanager.googleapis.com/projects/my-project/secrets/my-secret",
		},
		{
			name: "External: /projects/{project}/secrets/{secret} with kind=SecretManagerSecret",
			initial: TagsTagBindingParentRef{
				Kind:     "SecretManagerSecret",
				External: "/projects/my-project/secrets/my-secret",
			},
			expectedExternal: "//secretmanager.googleapis.com/projects/my-project/secrets/my-secret",
		},
		{
			name: "External: //secretmanager.googleapis.com/projects/{project}/secrets/{secret} with kind=SecretManagerSecret",
			initial: TagsTagBindingParentRef{
				Kind:     "SecretManagerSecret",
				External: "//secretmanager.googleapis.com/projects/my-project/secrets/my-secret",
			},
			expectedExternal: "//secretmanager.googleapis.com/projects/my-project/secrets/my-secret",
		},
		{
			name: "External: //storage.googleapis.com/projects/{project}/buckets/{bucket} with kind=StorageBucket",
			initial: TagsTagBindingParentRef{
				Kind:     "StorageBucket",
				External: "//storage.googleapis.com/projects/my-project/buckets/my-bucket",
			},
			expectedExternal: "//storage.googleapis.com/projects/my-project/buckets/my-bucket",
		},
		{
			name: "External: //storage.googleapis.com/buckets/{bucket} with kind=StorageBucket (invalid without projects)",
			initial: TagsTagBindingParentRef{
				Kind:     "StorageBucket",
				External: "//storage.googleapis.com/buckets/my-bucket",
			},
			wantErr: true,
		},
		{
			name: "External: //storage.googleapis.com/projects/_/buckets/somebucket with kind=StorageBucket",
			initial: TagsTagBindingParentRef{
				Kind:     "StorageBucket",
				External: "//storage.googleapis.com/projects/_/buckets/my-bucket",
			},
			expectedExternal: "//storage.googleapis.com/projects/_/buckets/my-bucket",
		},
		{
			name: "External: {project} (project ID) with kind=Project",
			initial: TagsTagBindingParentRef{
				Kind:     "Project",
				External: "my-project",
			},
			expectedExternal: "//cloudresourcemanager.googleapis.com/projects/my-project",
		},
		{
			name: "External: {project} (project ID) with no kind",
			initial: TagsTagBindingParentRef{
				External: "my-project",
			},
			expectedExternal: "//cloudresourcemanager.googleapis.com/projects/my-project",
		},
		{
			name: "External: projects/{project}/buckets/{bucket} with kind=Project (should fail)",
			initial: TagsTagBindingParentRef{
				Kind:     "Project",
				External: "projects/my-project/buckets/my-bucket",
			},
			wantErr:       true,
			wantErrSubstr: "unknown format for a Project reference",
		},
		{
			name: "External: projects/{project}/buckets/{bucket} with no kind (defaults to Project, should fail)",
			initial: TagsTagBindingParentRef{
				External: "projects/my-project/buckets/my-bucket",
			},
			wantErr:       true,
			wantErrSubstr: "unknown format for a Project reference",
		},
		{
			name: "External: //storage.googleapis.com/projects/{project}/buckets/{bucket} with kind=Project (should fail)",
			initial: TagsTagBindingParentRef{
				Kind:     "Project",
				External: "//storage.googleapis.com/projects/my-project/buckets/my-bucket",
			},
			wantErr:       true,
			wantErrSubstr: "unknown format for a Project reference",
		},
		{
			name: "External: //storage.googleapis.com/projects/{project}/buckets/{bucket} with no kind (defaults to Project, should fail)",
			initial: TagsTagBindingParentRef{
				External: "//storage.googleapis.com/projects/my-project/buckets/my-bucket",
			},
			wantErr:       true,
			wantErrSubstr: "unknown format for a Project reference",
		},
		{
			name: "External: organizations/{organization} with kind=Organization",
			initial: TagsTagBindingParentRef{
				Kind:     "Organization",
				External: "organizations/123456789",
			},
			expectedExternal: "//cloudresourcemanager.googleapis.com/organizations/123456789",
		},
		{
			name: "External: //cloudresourcemanager.googleapis.com/organizations/{organization} with kind=Organization",
			initial: TagsTagBindingParentRef{
				Kind:     "Organization",
				External: "//cloudresourcemanager.googleapis.com/organizations/123456789",
			},
			expectedExternal: "//cloudresourcemanager.googleapis.com/organizations/123456789",
		},
		{
			name: "External: organizations/{organization} with no kind (should infer Organization)",
			initial: TagsTagBindingParentRef{
				External: "organizations/123456789",
			},
			expectedExternal: "//cloudresourcemanager.googleapis.com/organizations/123456789",
		},
		{
			name: "External: //cloudresourcemanager.googleapis.com/organizations/{organization} with no kind (should infer Organization)",
			initial: TagsTagBindingParentRef{
				External: "//cloudresourcemanager.googleapis.com/organizations/123456789",
			},
			expectedExternal: "//cloudresourcemanager.googleapis.com/organizations/123456789",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.initial
			// We pass a nil reader as we are only testing External normalization which shouldn't require reading k8s objects
			var reader client.Reader
			err := r.Normalize(context.Background(), reader, "default")
			if (err != nil) != tt.wantErr {
				t.Errorf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.wantErrSubstr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErrSubstr) {
					t.Errorf("Normalize() error = %v, want substring %q", err, tt.wantErrSubstr)
				}
			}
			if !tt.wantErr && r.External != tt.expectedExternal {
				t.Errorf("Normalize() external = %v, want %v", r.External, tt.expectedExternal)
			}
		})
	}
}

func TestTagsTagBindingParentRef_NormalizeSecretManagerSecretByName(t *testing.T) {
	tests := []struct {
		name                string
		statusExternalRef   string
		expectedExternalRef string
	}{
		{
			name:                "unqualified status external reference",
			statusExternalRef:   "projects/my-project/secrets/my-secret",
			expectedExternalRef: "//secretmanager.googleapis.com/projects/my-project/secrets/my-secret",
		},
		{
			name:                "host-qualified status external reference",
			statusExternalRef:   "//secretmanager.googleapis.com/projects/my-project/secrets/my-secret",
			expectedExternalRef: "//secretmanager.googleapis.com/projects/my-project/secrets/my-secret",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			scheme := runtime.NewScheme()
			scheme.AddKnownTypeWithName(secretmanagerv1beta1.SecretManagerSecretGVK, &unstructured.Unstructured{})

			secret := &unstructured.Unstructured{Object: map[string]any{
				"apiVersion": "secretmanager.cnrm.cloud.google.com/v1beta1",
				"kind":       "SecretManagerSecret",
				"metadata": map[string]any{
					"name":      "my-secret",
					"namespace": "default",
				},
				"status": map[string]any{
					"externalRef": tc.statusExternalRef,
				},
			}}
			reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(secret).Build()

			ref := &TagsTagBindingParentRef{Kind: "SecretManagerSecret", Name: "my-secret"}
			if err := ref.Normalize(context.Background(), reader, "default"); err != nil {
				t.Fatalf("Normalize() returned error: %v", err)
			}
			if got := ref.External; got != tc.expectedExternalRef {
				t.Errorf("Normalize() external = %q, want %q", got, tc.expectedExternalRef)
			}
			if ref.Name != "" || ref.Namespace != "" {
				t.Errorf("Normalize() did not clear name reference: name=%q namespace=%q", ref.Name, ref.Namespace)
			}
			if err := ref.Normalize(context.Background(), reader, "default"); err != nil {
				t.Fatalf("second Normalize() returned error: %v", err)
			}
		})
	}
}
