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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestParseTaskExternal(t *testing.T) {
	tests := []struct {
		name          string
		external      string
		wantProject   string
		wantLocation  string
		wantJobID     string
		wantTaskGroup string
		wantID        string
		wantErr       bool
	}{
		{
			name:         "valid external 6-token ref",
			external:     "projects/test-project/locations/us-central1/tasks/task-01",
			wantProject:  "test-project",
			wantLocation: "us-central1",
			wantID:       "task-01",
			wantErr:      false,
		},
		{
			name:          "valid external 10-token ref",
			external:      "projects/test-project/locations/us-central1/jobs/job-01/taskGroups/group-01/tasks/task-01",
			wantProject:   "test-project",
			wantLocation:  "us-central1",
			wantJobID:     "job-01",
			wantTaskGroup: "group-01",
			wantID:        "task-01",
			wantErr:       false,
		},
		{
			name:     "invalid prefix",
			external: "project/test-project/locations/us-central1/tasks/task-01",
			wantErr:  true,
		},
		{
			name:     "missing tasks token",
			external: "projects/test-project/locations/us-central1/job/task-01",
			wantErr:  true,
		},
		{
			name:     "wrong number of tokens",
			external: "projects/test-project/locations/us-central1/tasks",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parent, resourceID, err := ParseTaskExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseTaskExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if parent.ProjectID != tt.wantProject {
					t.Errorf("expected ProjectID %q, got %q", tt.wantProject, parent.ProjectID)
				}
				if parent.Location != tt.wantLocation {
					t.Errorf("expected Location %q, got %q", tt.wantLocation, parent.Location)
				}
				if parent.JobID != tt.wantJobID {
					t.Errorf("expected JobID %q, got %q", tt.wantJobID, parent.JobID)
				}
				if parent.TaskGroup != tt.wantTaskGroup {
					t.Errorf("expected TaskGroup %q, got %q", tt.wantTaskGroup, parent.TaskGroup)
				}
				if resourceID != tt.wantID {
					t.Errorf("expected resourceID %q, got %q", tt.wantID, resourceID)
				}
			}
		})
	}
}

func TestNewTaskIdentity(t *testing.T) {
	tests := []struct {
		name       string
		obj        *BatchTask
		wantString string
		wantErr    bool
	}{
		{
			name: "valid construction 6-token",
			obj: &BatchTask{
				Spec: BatchTaskSpec{
					Parent: Parent{
						Location: direct.PtrTo("us-central1"),
						ProjectRef: &refsv1beta1.ProjectRef{
							External: "test-project",
						},
					},
					ResourceID: direct.PtrTo("task-01"),
				},
			},
			wantString: "projects/test-project/locations/us-central1/tasks/task-01",
			wantErr:    false,
		},
		{
			name: "valid construction 10-token",
			obj: &BatchTask{
				Spec: BatchTaskSpec{
					Parent: Parent{
						Location: direct.PtrTo("us-central1"),
						ProjectRef: &refsv1beta1.ProjectRef{
							External: "test-project",
						},
					},
					ResourceID: direct.PtrTo("task-01"),
				},
				Status: BatchTaskStatus{
					ExternalRef: direct.PtrTo("projects/test-project/locations/us-central1/jobs/job-01/taskGroups/group-01/tasks/task-01"),
				},
			},
			wantString: "projects/test-project/locations/us-central1/jobs/job-01/taskGroups/group-01/tasks/task-01",
			wantErr:    false,
		},
		{
			name: "missing project",
			obj: &BatchTask{
				Spec: BatchTaskSpec{
					Parent: Parent{
						Location: direct.PtrTo("us-central1"),
					},
					ResourceID: direct.PtrTo("task-01"),
				},
			},
			wantErr: true,
		},
		{
			name: "invalid actual externalRef mismatch project",
			obj: &BatchTask{
				Spec: BatchTaskSpec{
					Parent: Parent{
						Location: direct.PtrTo("us-central1"),
						ProjectRef: &refsv1beta1.ProjectRef{
							External: "test-project",
						},
					},
					ResourceID: direct.PtrTo("task-01"),
				},
				Status: BatchTaskStatus{
					ExternalRef: direct.PtrTo("projects/another-project/locations/us-central1/tasks/task-01"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			id, err := NewTaskIdentity(ctx, nil, tt.obj)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewTaskIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if id.String() != tt.wantString {
					t.Errorf("expected identity %q, got %q", tt.wantString, id.String())
				}
			}
		})
	}
}
