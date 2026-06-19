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

func TestParseDiscoveryEngineEngineExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    string
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/global/collections/default_collection/engines/my-engine",
			want: "projects/my-project/locations/global/collections/default_collection/engines/my-engine",
		},
		{
			name: "valid reference with api prefix",
			ref:  "//discoveryengine.googleapis.com/projects/my-project/locations/global/collections/default_collection/engines/my-engine",
			want: "projects/my-project/locations/global/collections/default_collection/engines/my-engine",
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "missing engines",
			ref:     "projects/my-project/locations/global/collections/default_collection/my-engine",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := ParseDiscoveryEngineEngineExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDiscoveryEngineEngineExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if id.String() != tt.want {
					t.Errorf("id.String() = %v, want %v", id.String(), tt.want)
				}
				if id.ProjectID != "my-project" {
					t.Errorf("ProjectID = %v, want my-project", id.ProjectID)
				}
				if id.Location != "global" {
					t.Errorf("Location = %v, want global", id.Location)
				}
				if id.Collection != "default_collection" {
					t.Errorf("Collection = %v, want default_collection", id.Collection)
				}
				if id.DataStore != "my-engine" {
					t.Errorf("DataStore = %v, want my-engine", id.DataStore)
				}
			}
		})
	}
}
