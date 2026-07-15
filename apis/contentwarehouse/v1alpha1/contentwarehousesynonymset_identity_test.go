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

func TestContentWarehouseSynonymSetIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     *ContentWarehouseSynonymSetIdentity
		wantErr  bool
	}{
		{
			name:     "canonical format",
			external: "projects/my-project/locations/us-central1/synonymSets/my-synonymset",
			want: &ContentWarehouseSynonymSetIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Context:  "my-synonymset",
			},
			wantErr: false,
		},
		{
			name:     "invalid format",
			external: "projects/my-project/locations/us-central1",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ContentWarehouseSynonymSetIdentity{}
			if err := i.FromExternal(tt.external); (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if i.String() != tt.external {
					t.Errorf("String() = %v, want %v", i.String(), tt.external)
				}
			}
		})
	}
}
