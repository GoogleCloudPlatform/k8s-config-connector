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

func TestNotebookInstanceV2Identity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     *NotebookInstanceV2Identity
		wantErr  bool
	}{
		{
			name:     "valid external",
			external: "projects/p1/locations/l1/instances/i1",
			want: &NotebookInstanceV2Identity{
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
			i := &NotebookInstanceV2Identity{}
			err := i.FromExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NotebookInstanceV2Identity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("NotebookInstanceV2Identity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if i.String() != tt.external {
					t.Errorf("NotebookInstanceV2Identity.String() = %q, want %q", i.String(), tt.external)
				}
				if i.ParentString() != "projects/p1/locations/l1" {
					t.Errorf("NotebookInstanceV2Identity.ParentString() = %q, want %q", i.ParentString(), "projects/p1/locations/l1")
				}
			}
		})
	}
}
