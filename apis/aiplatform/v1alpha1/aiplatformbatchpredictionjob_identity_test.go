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
	"testing"
)

func TestAIPlatformBatchPredictionJobIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name        string
		external    string
		wantProject string
		wantLoc     string
		wantJob     string
		wantErr     bool
	}{
		{
			name:        "valid fully qualified",
			external:    "projects/my-project/locations/us-central1/batchPredictionJobs/my-job",
			wantProject: "my-project",
			wantLoc:     "us-central1",
			wantJob:     "my-job",
		},
		{
			name:        "valid short format with host",
			external:    "//aiplatform.googleapis.com/projects/my-project/locations/us-central1/batchPredictionJobs/my-job",
			wantProject: "my-project",
			wantLoc:     "us-central1",
			wantJob:     "my-job",
		},
		{
			name:     "invalid format",
			external: "projects/my-project/locations/us-central1/jobs/my-job",
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &AIPlatformBatchPredictionJobIdentity{}
			err := id.FromExternal(tc.external)
			if (err != nil) != tc.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tc.wantErr)
			}
			if err == nil {
				if id.Project != tc.wantProject {
					t.Errorf("Project = %v, want %v", id.Project, tc.wantProject)
				}
				if id.Location != tc.wantLoc {
					t.Errorf("Location = %v, want %v", id.Location, tc.wantLoc)
				}
				if id.BatchPredictionJob != tc.wantJob {
					t.Errorf("BatchPredictionJob = %v, want %v", id.BatchPredictionJob, tc.wantJob)
				}
			}
		})
	}
}
