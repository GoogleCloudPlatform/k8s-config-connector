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

package v1beta1

import (
	"reflect"
	"testing"
)

func TestParseInstanceExternal(t *testing.T) {
	tests := []struct {
		name           string
		external       string
		wantParent     *InstanceParent
		wantResourceID string
		wantErr        bool
	}{
		{
			name:     "valid external",
			external: "projects/p1/locations/l1/instances/i1",
			wantParent: &InstanceParent{
				ProjectID: "p1",
				Location:  "l1",
			},
			wantResourceID: "i1",
			wantErr:        false,
		},
		{
			name:     "invalid format",
			external: "invalid/format",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParent, gotResourceID, err := ParseInstanceExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInstanceExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(gotParent, tt.wantParent) {
					t.Errorf("ParseInstanceExternal() gotParent = %v, want %v", gotParent, tt.wantParent)
				}
				if gotResourceID != tt.wantResourceID {
					t.Errorf("ParseInstanceExternal() gotResourceID = %v, want %v", gotResourceID, tt.wantResourceID)
				}
			}
		})
	}
}

func TestInstanceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     *InstanceIdentity
		wantErr  bool
	}{
		{
			name:     "valid external",
			external: "projects/p1/locations/l1/instances/i1",
			want: &InstanceIdentity{
				Project:  "p1",
				Location: "l1",
				Instance: "i1",
			},
			wantErr: false,
		},
		{
			name:     "invalid format",
			external: "invalid/format",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InstanceIdentity{}
			if err := i.FromExternal(tt.external); (err != nil) != tt.wantErr {
				t.Errorf("InstanceIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(i, tt.want) {
					t.Errorf("InstanceIdentity.FromExternal() = %v, want %v", i, tt.want)
				}
			}
		})
	}
}
