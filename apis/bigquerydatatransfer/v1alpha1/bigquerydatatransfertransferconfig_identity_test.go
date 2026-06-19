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

func TestBigQueryDataTransferTransferConfigIdentity_FromExternal(t *testing.T) {
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
			ref:          "projects/my-project/locations/us-central1/transferConfigs/my-transfer-config",
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-transfer-config",
			wantErr:      false,
		},
		{
			name:         "full url",
			ref:          "https://bigquerydatatransfer.googleapis.com/projects/my-project/locations/us-central1/transferConfigs/my-transfer-config",
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-transfer-config",
			wantErr:      false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/transferConfigs/my-transfer-config",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BigQueryDataTransferTransferConfigIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("BigQueryDataTransferTransferConfigIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.Project != tt.wantProject {
					t.Errorf("BigQueryDataTransferTransferConfigIdentity.FromExternal() Project = %v, want %v", i.Project, tt.wantProject)
				}
				if i.Location != tt.wantLocation {
					t.Errorf("BigQueryDataTransferTransferConfigIdentity.FromExternal() Location = %v, want %v", i.Location, tt.wantLocation)
				}
				if i.Transfer_config != tt.wantID {
					t.Errorf("BigQueryDataTransferTransferConfigIdentity.FromExternal() Transfer_config = %v, want %v", i.Transfer_config, tt.wantID)
				}
				if got := i.String(); got != "projects/"+tt.wantProject+"/locations/"+tt.wantLocation+"/transferConfigs/"+tt.wantID {
					t.Errorf("BigQueryDataTransferTransferConfigIdentity.String() = %v", got)
				}
			}
		})
	}
}
