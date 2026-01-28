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

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
