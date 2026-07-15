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

func TestParallelstoreInstanceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *ParallelstoreInstanceIdentity
		wantErr bool
	}{
		{
			name: "valid external ref",
			ref:  "projects/my-project/locations/us-central1/instances/my-instance",
			want: &ParallelstoreInstanceIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Instance: "my-instance",
			},
			wantErr: false,
		},
		{
			name: "full url",
			ref:  "https://parallelstore.googleapis.com/projects/my-project/locations/us-central1/instances/my-instance",
			want: &ParallelstoreInstanceIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Instance: "my-instance",
			},
			wantErr: false,
		},
		{
			name:    "invalid external ref missing segments",
			ref:     "projects/my-project/instances/my-instance",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ParallelstoreInstanceIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParallelstoreInstanceIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("ParallelstoreInstanceIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
				gotStr := i.String()
				wantStr := "projects/" + tt.want.Project + "/locations/" + tt.want.Location + "/instances/" + tt.want.Instance
				if gotStr != wantStr {
					t.Errorf("ParallelstoreInstanceIdentity.String() = %v, want %v", gotStr, wantStr)
				}
			}
		})
	}
}
