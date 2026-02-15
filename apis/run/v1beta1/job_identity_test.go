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
	"testing"
)

func TestJobIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name         string
		ref          string
		wantProject  string
		wantLocation string
		wantID       string
		wantErr      bool
	}{
		{
			name:         "valid external ref",
			ref:          "projects/my-project/locations/us-central1/jobs/my-job",
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-job",
			wantErr:      false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/jobs/my-job",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &JobIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("JobIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.ProjectID != tt.wantProject {
					t.Errorf("JobIdentity.FromExternal() ProjectID = %v, want %v", i.ProjectID, tt.wantProject)
				}
				if i.Location != tt.wantLocation {
					t.Errorf("JobIdentity.FromExternal() Location = %v, want %v", i.Location, tt.wantLocation)
				}
				if i.JobID != tt.wantID {
					t.Errorf("JobIdentity.FromExternal() ID = %v, want %v", i.JobID, tt.wantID)
				}
				if got := i.String(); got != tt.ref {
					t.Errorf("JobIdentity.String() = %v, want %v", got, tt.ref)
				}
			}
		})
	}
}
