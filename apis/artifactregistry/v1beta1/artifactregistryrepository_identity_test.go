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
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestArtifactRegistryRepositoryIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ArtifactRegistryRepositoryIdentity
	}{
		{
			name: "valid_reference",
			ref:  "projects/my-project/locations/us-central1/repositories/my-repo",
			want: &ArtifactRegistryRepositoryIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Repository: "my-repo",
			},
		},
		{
			name:    "invalid_reference_format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full_url",
			ref:  "https://artifactregistry.googleapis.com/projects/my-project/locations/us-central1/repositories/my-repo",
			want: &ArtifactRegistryRepositoryIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Repository: "my-repo",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			i := &ArtifactRegistryRepositoryIdentity{}
			err := i.FromExternal(tc.ref)

			if (err != nil) != tc.wantErr {
				t.Fatalf("FromExternal(%q) error = %v, wantErr %v", tc.ref, err, tc.wantErr)
			}
			if !tc.wantErr {
				if diff := cmp.Diff(tc.want, i); diff != "" {
					t.Errorf("FromExternal(%q) mismatch (-want +got):\n%s", tc.ref, diff)
				}
			}
		})
	}
}

func TestArtifactRegistryRepository_GetIdentity(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme() unexpected error: %v", err)
	}
	reader := fake.NewClientBuilder().WithScheme(scheme).Build()

	strPtr := func(s string) *string { return &s }

	tests := []struct {
		name          string
		obj           *ArtifactRegistryRepository
		want          *ArtifactRegistryRepositoryIdentity
		wantErrSubstr string
	}{
		{
			name: "valid_spec_and_matching_short_status_name",
			obj: &ArtifactRegistryRepository{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-repo",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: ArtifactRegistryRepositorySpec{
					Location:   "us-central1",
					ResourceID: strPtr("real-repo"),
				},
				Status: ArtifactRegistryRepositoryStatus{
					Name: strPtr("real-repo"),
				},
			},
			want: &ArtifactRegistryRepositoryIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Repository: "real-repo",
			},
		},
		{
			name: "rejects_package_path_traversal_in_resource_id",
			obj: &ArtifactRegistryRepository{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "bad-id",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: ArtifactRegistryRepositorySpec{
					Location:   "us-west1",
					ResourceID: strPtr("real-repo/packages/my-pkg"),
				},
			},
			wantErrSubstr: "invalid value",
		},
		{
			name: "rejects_dot_dot_segment_in_resource_id",
			obj: &ArtifactRegistryRepository{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "dot-dot",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: ArtifactRegistryRepositorySpec{
					Location:   "us-west1",
					ResourceID: strPtr(".."),
				},
			},
			wantErrSubstr: "invalid value",
		},
		{
			name: "rejects_slash_in_project_id",
			obj: &ArtifactRegistryRepository{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-repo",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project/locations/us",
					},
				},
				Spec: ArtifactRegistryRepositorySpec{
					Location: "us-west1",
				},
			},
			wantErrSubstr: "invalid value",
		},
		{
			name: "reports_location_error_when_spec_location_invalid",
			obj: &ArtifactRegistryRepository{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "bad-loc",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: ArtifactRegistryRepositorySpec{
					Location: "us_central1",
				},
			},
			wantErrSubstr: "cannot resolve location",
		},
		{
			name: "reports_project_error_when_project_missing",
			obj: &ArtifactRegistryRepository{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "missing-proj",
					Namespace: "default",
				},
				Spec: ArtifactRegistryRepositorySpec{
					Location: "us-central1",
				},
			},
			wantErrSubstr: "cannot resolve project",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.obj.GetIdentity(t.Context(), reader)

			if tc.wantErrSubstr != "" {
				if err == nil {
					t.Fatalf("GetIdentity(%+v) = %+v, want error containing %q", tc.obj, got, tc.wantErrSubstr)
				}
				if !strings.Contains(err.Error(), tc.wantErrSubstr) {
					t.Errorf("GetIdentity() error = %q, want substring %q", err.Error(), tc.wantErrSubstr)
				}
				return
			}
			if err != nil {
				t.Fatalf("GetIdentity() unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
