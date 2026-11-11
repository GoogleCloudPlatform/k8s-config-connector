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

	"github.com/google/go-cmp/cmp"
)

func TestCESAppIdentity_FromExternal(t *testing.T) {
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
			ref:          "projects/my-project/locations/us-central1/apps/my-app",
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-app",
			wantErr:      false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/apps/my-app",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CESAppIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("CESAppIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.Project != tt.wantProject {
					t.Errorf("CESAppIdentity.FromExternal() Project = %v, want %v", i.Project, tt.wantProject)
				}
				if i.Location != tt.wantLocation {
					t.Errorf("CESAppIdentity.FromExternal() Location = %v, want %v", i.Location, tt.wantLocation)
				}
				if i.App != tt.wantID {
					t.Errorf("CESAppIdentity.FromExternal() App = %v, want %v", i.App, tt.wantID)
				}
				if got := i.String(); got != tt.ref {
					t.Errorf("CESAppIdentity.String() = %v, want %v", got, tt.ref)
				}
			}
		})
	}
}

func TestCesToolIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name         string
		ref          string
		wantIdentity *CesToolIdentity
		wantErr      bool
	}{
		{
			name: "valid external ref",
			ref:  "projects/my-project/locations/us-central1/apps/my-app/tools/my-tool",
			wantIdentity: &CesToolIdentity{
				Project:  "my-project",
				Location: "us-central1",
				App:      "my-app",
				Tool:     "my-tool",
			},
			wantErr: false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/apps/my-app/tools/my-tool",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CesToolIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Fatalf("CesToolIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.wantIdentity, i); diff != "" {
					t.Errorf("CesToolIdentity mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != tt.ref {
					t.Errorf("CesToolIdentity.String() = %v, want %v", got, tt.ref)
				}
			}
		})
	}
}

func TestCesToolsetIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name         string
		ref          string
		wantIdentity *CesToolsetIdentity
		wantErr      bool
	}{
		{
			name: "valid external ref",
			ref:  "projects/my-project/locations/us-central1/apps/my-app/toolsets/my-toolset",
			wantIdentity: &CesToolsetIdentity{
				Project:  "my-project",
				Location: "us-central1",
				App:      "my-app",
				Toolset:  "my-toolset",
			},
			wantErr: false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/apps/my-app/toolsets/my-toolset",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CesToolsetIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Fatalf("CesToolsetIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.wantIdentity, i); diff != "" {
					t.Errorf("CesToolsetIdentity mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != tt.ref {
					t.Errorf("CesToolsetIdentity.String() = %v, want %v", got, tt.ref)
				}
			}
		})
	}
}
