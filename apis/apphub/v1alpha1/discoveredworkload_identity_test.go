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
	"reflect"
	"testing"
)

func TestParseDiscoveredWorkloadExternal(t *testing.T) {
	tests := []struct {
		name           string
		external       string
		wantParent     *DiscoveredWorkloadParent
		wantResourceID string
		wantErr        bool
	}{
		{
			name:     "valid external",
			external: "projects/p1/locations/l1/discoveredworkloads/w1",
			wantParent: &DiscoveredWorkloadParent{
				ProjectID: "p1",
				Location:  "l1",
			},
			wantResourceID: "w1",
			wantErr:        false,
		},
		{
			name:     "invalid format",
			external: "invalid/format",
			wantErr:  true,
		},
		{
			name:     "wrong service prefix",
			external: "projects/p1/locations/l1/wrongresource/w1",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParent, gotResourceID, err := ParseDiscoveredWorkloadExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDiscoveredWorkloadExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(gotParent, tt.wantParent) {
					t.Errorf("ParseDiscoveredWorkloadExternal() gotParent = %v, want %v", gotParent, tt.wantParent)
				}
				if gotResourceID != tt.wantResourceID {
					t.Errorf("ParseDiscoveredWorkloadExternal() gotResourceID = %v, want %v", gotResourceID, tt.wantResourceID)
				}
			}
		})
	}
}

func TestDiscoveredWorkloadIdentity_Methods(t *testing.T) {
	parent := &DiscoveredWorkloadParent{
		ProjectID: "p1",
		Location:  "l1",
	}
	id := &DiscoveredWorkloadIdentity{
		parent: parent,
		id:     "w1",
	}

	if id.ID() != "w1" {
		t.Errorf("ID() = %v, want %v", id.ID(), "w1")
	}

	if !reflect.DeepEqual(id.Parent(), parent) {
		t.Errorf("Parent() = %v, want %v", id.Parent(), parent)
	}

	expectedString := "projects/p1/locations/l1/discoveredworkloads/w1"
	if id.String() != expectedString {
		t.Errorf("String() = %v, want %v", id.String(), expectedString)
	}
}
