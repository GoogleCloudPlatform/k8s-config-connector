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

package v1alpha1

import (
	"context"
	"testing"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

func TestVertexAIStudyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *VertexAIStudyIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/studies/my-study",
			want: &VertexAIStudyIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Study:    "my-study",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://aiplatform.googleapis.com/projects/my-project/locations/us-central1/studies/my-study",
			want: &VertexAIStudyIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Study:    "my-study",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &VertexAIStudyIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Location != tt.want.Location {
					t.Errorf("Location = %v, want %v", i.Location, tt.want.Location)
				}
				if i.Study != tt.want.Study {
					t.Errorf("Study = %v, want %v", i.Study, tt.want.Study)
				}
			}
		})
	}
}

func TestVertexAIStudy_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		obj         *VertexAIStudy
		expected    string
		expectError bool
	}{
		{
			name: "resolved identity without externalRef",
			obj: &VertexAIStudy{
				Spec: VertexAIStudySpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
					Location:   stringPtr("us-central1"),
					ResourceID: stringPtr("my-study"),
				},
			},
			expected: "projects/my-project/locations/us-central1/studies/my-study",
		},
		{
			name: "resolved identity matching externalRef but numeric project number",
			obj: &VertexAIStudy{
				Spec: VertexAIStudySpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
					Location:   stringPtr("us-central1"),
					ResourceID: stringPtr("my-study"),
				},
				Status: VertexAIStudyStatus{
					ExternalRef: stringPtr("projects/1234567890/locations/us-central1/studies/2654607894077"),
				},
			},
			expected: "projects/1234567890/locations/us-central1/studies/2654607894077",
		},
		{
			name: "resolved identity mismatching location in externalRef",
			obj: &VertexAIStudy{
				Spec: VertexAIStudySpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
					Location:   stringPtr("us-central1"),
					ResourceID: stringPtr("my-study"),
				},
				Status: VertexAIStudyStatus{
					ExternalRef: stringPtr("projects/1234567890/locations/us-east1/studies/2654607894077"),
				},
			},
			expectError: true,
		},
		{
			name: "resolved identity mismatching project in externalRef (both alphanumeric)",
			obj: &VertexAIStudy{
				Spec: VertexAIStudySpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
					Location:   stringPtr("us-central1"),
					ResourceID: stringPtr("my-study"),
				},
				Status: VertexAIStudyStatus{
					ExternalRef: stringPtr("projects/other-project/locations/us-central1/studies/2654607894077"),
				},
			},
			expectError: true,
		},
		{
			name: "resolved identity mismatching project in externalRef (both numeric)",
			obj: &VertexAIStudy{
				Spec: VertexAIStudySpec{
					ProjectRef: &refs.ProjectRef{
						External: "1234567890",
					},
					Location:   stringPtr("us-central1"),
					ResourceID: stringPtr("my-study"),
				},
				Status: VertexAIStudyStatus{
					ExternalRef: stringPtr("projects/9876543210/locations/us-central1/studies/2654607894077"),
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, nil)
			if (err != nil) != tt.expectError {
				t.Errorf("GetIdentity() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError {
				if got.String() != tt.expected {
					t.Errorf("GetIdentity() = %q, expected %q", got.String(), tt.expected)
				}
			}
		})
	}
}

func TestIsProjectIDMatch(t *testing.T) {
	tests := []struct {
		p1, p2 string
		want   bool
	}{
		{"", "", false},
		{"a", "", false},
		{"", "b", false},
		{"my-project", "my-project", true},
		{"my-project", "other-project", false},
		{"123456", "123456", true},
		{"123456", "654321", false},
		{"123456", "my-project", true}, // alphanumeric vs numeric should skip strict check
		{"my-project", "123456", true}, // alphanumeric vs numeric should skip strict check
	}

	for _, tt := range tests {
		t.Run(tt.p1+"_vs_"+tt.p2, func(t *testing.T) {
			got := IsProjectIDMatch(tt.p1, tt.p2)
			if got != tt.want {
				t.Errorf("IsProjectIDMatch(%q, %q) = %v; want %v", tt.p1, tt.p2, got, tt.want)
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}
