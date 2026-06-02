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

func TestCloudRunInstanceIdentity_FromExternal(t *testing.T) {
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
			ref:          "projects/my-project/locations/us-central1/instances/my-instance",
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-instance",
			wantErr:      false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/instances/my-instance",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CloudRunInstanceIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("CloudRunInstanceIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.Project != tt.wantProject {
					t.Errorf("CloudRunInstanceIdentity.FromExternal() project = %v, want %v", i.Project, tt.wantProject)
				}
				if i.Location != tt.wantLocation {
					t.Errorf("CloudRunInstanceIdentity.FromExternal() location = %v, want %v", i.Location, tt.wantLocation)
				}
				if i.Instance != tt.wantID {
					t.Errorf("CloudRunInstanceIdentity.FromExternal() instance = %v, want %v", i.Instance, tt.wantID)
				}
				if got := i.String(); got != tt.ref {
					t.Errorf("CloudRunInstanceIdentity.String() = %v, want %v", got, tt.ref)
				}
			}
		})
	}
}
