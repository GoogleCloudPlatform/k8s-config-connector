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

	sqlv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
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
			name: "External: //sqladmin.googleapis.com/projects/{project}/instances/{instance} with kind=SQLInstance",
			initial: TagsTagBindingParentRef{
				Kind:     "SQLInstance",
				External: "//sqladmin.googleapis.com/projects/my-project/instances/my-instance",
			},
			expectedExternal: "//sqladmin.googleapis.com/projects/my-project/instances/my-instance",
		},
		{
			name: "External: projects/{project}/instances/{instance} with kind=SQLInstance",
			initial: TagsTagBindingParentRef{
				Kind:     "SQLInstance",
				External: "projects/my-project/instances/my-instance",
			},
			expectedExternal: "//sqladmin.googleapis.com/projects/my-project/instances/my-instance",
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
			name: "External: projects/{project}/instances/{instance} with kind=Project (should fail)",
			initial: TagsTagBindingParentRef{
				Kind:     "Project",
				External: "projects/my-project/instances/my-instance",
			},
			wantErr:       true,
			wantErrSubstr: "unknown format for a Project reference",
		},
		{
			name: "External: projects/{project}/instances/{instance} with no kind (defaults to Project, should fail)",
			initial: TagsTagBindingParentRef{
				External: "projects/my-project/instances/my-instance",
			},
			wantErr:       true,
			wantErrSubstr: "unknown format for a Project reference",
		},
		{
			name: "External: //sqladmin.googleapis.com/projects/{project}/instances/{instance} with kind=Project (should fail)",
			initial: TagsTagBindingParentRef{
				Kind:     "Project",
				External: "//sqladmin.googleapis.com/projects/my-project/instances/my-instance",
			},
			wantErr:       true,
			wantErrSubstr: "unknown format for a Project reference",
		},
		{
			name: "External: //sqladmin.googleapis.com/projects/{project}/instances/{instance} with no kind (defaults to Project, should fail)",
			initial: TagsTagBindingParentRef{
				External: "//sqladmin.googleapis.com/projects/my-project/instances/my-instance",
			},
			wantErr:       true,
			wantErrSubstr: "unknown format for a Project reference",
		},
		{
			name: "External: invalid format with kind=SQLInstance (should fail)",
			initial: TagsTagBindingParentRef{
				Kind:     "SQLInstance",
				External: "projects/my-project/locations/us-central1/instances/my-instance",
			},
			wantErr: true,
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

func TestTagsTagBindingParentRef_Normalize_ByName(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	if err := sqlv1beta1.AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme() error = %v", err)
	}

	sqlInstanceReady := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": sqlv1beta1.SQLInstanceGVK.GroupVersion().String(),
			"kind":       sqlv1beta1.SQLInstanceGVK.Kind,
			"metadata": map[string]interface{}{
				"name":      "ready-instance",
				"namespace": "test-ns",
				"annotations": map[string]interface{}{
					"cnrm.cloud.google.com/project-id": "my-project",
				},
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": "True",
					},
				},
			},
		},
	}

	sqlInstanceNotReady := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": sqlv1beta1.SQLInstanceGVK.GroupVersion().String(),
			"kind":       sqlv1beta1.SQLInstanceGVK.Kind,
			"metadata": map[string]interface{}{
				"name":      "not-ready-instance",
				"namespace": "test-ns",
				"annotations": map[string]interface{}{
					"cnrm.cloud.google.com/project-id": "my-project",
				},
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": "False",
					},
				},
			},
		},
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(sqlInstanceReady, sqlInstanceNotReady).Build()

	tests := []struct {
		name             string
		initial          TagsTagBindingParentRef
		defaultNS        string
		expectedExternal string
		wantErr          bool
	}{
		{
			name: "SQLInstance by name (ready) with explicit kind",
			initial: TagsTagBindingParentRef{
				Kind:      "SQLInstance",
				Name:      "ready-instance",
				Namespace: "test-ns",
			},
			defaultNS:        "test-ns",
			expectedExternal: "//sqladmin.googleapis.com/projects/my-project/instances/ready-instance",
		},
		{
			name: "SQLInstance by name (ready) using default namespace",
			initial: TagsTagBindingParentRef{
				Kind: "SQLInstance",
				Name: "ready-instance",
			},
			defaultNS:        "test-ns",
			expectedExternal: "//sqladmin.googleapis.com/projects/my-project/instances/ready-instance",
		},
		{
			name: "SQLInstance by name (not ready) should fail",
			initial: TagsTagBindingParentRef{
				Kind:      "SQLInstance",
				Name:      "not-ready-instance",
				Namespace: "test-ns",
			},
			defaultNS: "test-ns",
			wantErr:   true,
		},
		{
			name: "SQLInstance by name (not found) should fail",
			initial: TagsTagBindingParentRef{
				Kind:      "SQLInstance",
				Name:      "non-existent-instance",
				Namespace: "test-ns",
			},
			defaultNS: "test-ns",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.initial
			err := r.Normalize(ctx, fakeClient, tt.defaultNS)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if r.External != tt.expectedExternal {
				t.Errorf("Normalize() external = %v, want %v", r.External, tt.expectedExternal)
			}
			if r.Name != "" || r.Namespace != "" {
				t.Errorf("Normalize() name/namespace not cleared: name=%q, namespace=%q", r.Name, r.Namespace)
			}
		})
	}
}
