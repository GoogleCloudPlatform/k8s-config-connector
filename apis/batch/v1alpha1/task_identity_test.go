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
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestTaskIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name          string
		external      string
		wantProject   string
		wantLocation  string
		wantJob       string
		wantTaskGroup string
		wantTask      string
		wantErr       bool
	}{
		{
			name:          "valid external 10-token ref",
			external:      "projects/test-project/locations/us-central1/jobs/job-01/taskGroups/group-01/tasks/task-01",
			wantProject:   "test-project",
			wantLocation:  "us-central1",
			wantJob:       "job-01",
			wantTaskGroup: "group-01",
			wantTask:      "task-01",
			wantErr:       false,
		},
		{
			name:     "invalid prefix",
			external: "project/test-project/locations/us-central1/tasks/task-01",
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
			id := &TaskIdentity{}
			err := id.FromExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if id.Project != tt.wantProject {
					t.Errorf("expected Project %q, got %q", tt.wantProject, id.Project)
				}
				if id.Location != tt.wantLocation {
					t.Errorf("expected Location %q, got %q", tt.wantLocation, id.Location)
				}
				if id.Job != tt.wantJob {
					t.Errorf("expected Job %q, got %q", tt.wantJob, id.Job)
				}
				if id.TaskGroup != tt.wantTaskGroup {
					t.Errorf("expected TaskGroup %q, got %q", tt.wantTaskGroup, id.TaskGroup)
				}
				if id.Task != tt.wantTask {
					t.Errorf("expected Task %q, got %q", tt.wantTask, id.Task)
				}
			}
		})
	}
}

func TestBatchTask_GetIdentity(t *testing.T) {
	tests := []struct {
		name       string
		obj        *BatchTask
		wantString string
		wantErr    bool
	}{
		{
			name: "valid construction",
			obj: &BatchTask{
				Spec: BatchTaskSpec{
					Parent: Parent{
						Location: "us-central1",
						ProjectRef: &refsv1beta1.ProjectRef{
							External: "test-project",
						},
						JobRef: &BatchJobRef{
							External: "projects/test-project/locations/us-central1/jobs/job-01",
						},
						TaskGroup: "group-01",
					},
					ResourceID: direct.PtrTo("task-01"),
				},
			},
			wantString: "projects/test-project/locations/us-central1/jobs/job-01/taskGroups/group-01/tasks/task-01",
			wantErr:    false,
		},
		{
			name: "invalid actual externalRef mismatch project",
			obj: &BatchTask{
				Spec: BatchTaskSpec{
					Parent: Parent{
						Location: "us-central1",
						ProjectRef: &refsv1beta1.ProjectRef{
							External: "test-project",
						},
						JobRef: &BatchJobRef{
							External: "projects/test-project/locations/us-central1/jobs/job-01",
						},
						TaskGroup: "group-01",
					},
					ResourceID: direct.PtrTo("task-01"),
				},
				Status: BatchTaskStatus{
					ExternalRef: direct.PtrTo("projects/another-project/locations/us-central1/jobs/job-01/taskGroups/group-01/tasks/task-01"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			reader := fake.NewClientBuilder().Build()
			id, err := tt.obj.GetIdentity(ctx, reader)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if id.String() != tt.wantString {
					t.Errorf("expected identity %q, got %q", tt.wantString, id.String())
				}
			}
		})
	}
}
