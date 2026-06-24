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

func TestAspectTypeIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name           string
		ref            string
		wantProject    string
		wantLocation   string
		wantAspectType string
		wantErr        bool
	}{
		{
			name:           "valid external ref",
			ref:            "projects/my-project/locations/us-central1/aspectTypes/my-aspect-type",
			wantProject:    "my-project",
			wantLocation:   "us-central1",
			wantAspectType: "my-aspect-type",
			wantErr:        false,
		},
		{
			name:    "invalid external ref (missing location)",
			ref:     "projects/my-project/aspectTypes/my-aspect-type",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AspectTypeIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Fatalf("AspectTypeIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				got := &AspectTypeIdentity{
					Project:    i.Project,
					Location:   i.Location,
					AspectType: i.AspectType,
				}
				want := &AspectTypeIdentity{
					Project:    tt.wantProject,
					Location:   tt.wantLocation,
					AspectType: tt.wantAspectType,
				}
				if diff := cmp.Diff(want, got); diff != "" {
					t.Errorf("AspectTypeIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if gotStr := i.String(); gotStr != tt.ref {
					t.Errorf("AspectTypeIdentity.String() = %v, want %v", gotStr, tt.ref)
				}
			}
		})
	}
}
