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

package refs

import (
	"testing"
)

func TestMemorystoreInstanceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     *MemorystoreInstanceIdentity
		wantErr  bool
	}{
		{
			name:     "valid external",
			external: "projects/my-project/locations/us-central1/instances/my-instance",
			want: &MemorystoreInstanceIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Instance: "my-instance",
			},
			wantErr: false,
		},
		{
			name:     "invalid external - missing parts",
			external: "projects/my-project/locations/us-central1/instances",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "invalid external - wrong format",
			external: "my-project/us-central1/my-instance",
			want:     nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &MemorystoreInstanceIdentity{}
			err := i.FromExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemorystoreInstanceIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project || i.Location != tt.want.Location || i.Instance != tt.want.Instance {
					t.Errorf("MemorystoreInstanceIdentity.FromExternal() = %v, want %v", i, tt.want)
				}
			}
		})
	}
}

func TestMemorystoreInstanceIdentity_String(t *testing.T) {
	i := &MemorystoreInstanceIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Instance: "my-instance",
	}
	want := "projects/my-project/locations/us-central1/instances/my-instance"
	if got := i.String(); got != want {
		t.Errorf("MemorystoreInstanceIdentity.String() = %v, want %v", got, want)
	}
}

func TestParseInstanceExternal(t *testing.T) {
	external := "projects/my-project/locations/us-central1/instances/my-instance"
	parent, id, err := ParseInstanceExternal(external)
	if err != nil {
		t.Fatalf("ParseInstanceExternal() error = %v", err)
	}
	if parent.ProjectID != "my-project" || parent.Location != "us-central1" {
		t.Errorf("ParseInstanceExternal() parent = %v, want project=my-project, location=us-central1", parent)
	}
	if id != "my-instance" {
		t.Errorf("ParseInstanceExternal() id = %v, want my-instance", id)
	}
}
