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

func TestAIPlatformModelIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name      string
		external  string
		wantErr   bool
		wantProj  string
		wantLoc   string
		wantModel string
	}{
		{
			name:      "valid model external",
			external:  "projects/my-project/locations/us-central1/models/my-model",
			wantProj:  "my-project",
			wantLoc:   "us-central1",
			wantModel: "my-model",
		},
		{
			name:     "invalid prefix",
			external: "project/my-project/locations/us-central1/models/my-model",
			wantErr:  true,
		},
		{
			name:     "invalid format too short",
			external: "projects/my-project/locations/us-central1/models",
			wantErr:  true,
		},
		{
			name:     "invalid format too long",
			external: "projects/my-project/locations/us-central1/models/my-model/extra",
			wantErr:  true,
		},
		{
			name:     "invalid keyword location",
			external: "projects/my-project/location/us-central1/models/my-model",
			wantErr:  true,
		},
		{
			name:     "invalid keyword models",
			external: "projects/my-project/locations/us-central1/model/my-model",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := &AIPlatformModelIdentity{}
			err := identity.FromExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if identity.Project != tt.wantProj {
					t.Errorf("Project = %v, want %v", identity.Project, tt.wantProj)
				}
				if identity.Location != tt.wantLoc {
					t.Errorf("Location = %v, want %v", identity.Location, tt.wantLoc)
				}
				if identity.Model != tt.wantModel {
					t.Errorf("Model = %v, want %v", identity.Model, tt.wantModel)
				}
			}
		})
	}
}

func TestAIPlatformModelIdentity_String(t *testing.T) {
	identity := &AIPlatformModelIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Model:    "my-model",
	}

	want := "projects/my-project/locations/us-central1/models/my-model"
	if got := identity.String(); got != want {
		t.Errorf("ModelIdentity.String() = %q, want %q", got, want)
	}
}
