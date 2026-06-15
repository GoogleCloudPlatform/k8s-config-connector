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
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestComputeSnapshotIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name       string
		ref        string
		wantErr    bool
		want       *ComputeSnapshotIdentity
		wantString string
		wantParent string
	}{
		{
			name: "valid global reference",
			ref:  "projects/my-project/global/snapshots/my-snapshot",
			want: &ComputeSnapshotIdentity{
				Project:  "my-project",
				Snapshot: "my-snapshot",
			},
			wantString: "projects/my-project/global/snapshots/my-snapshot",
			wantParent: "projects/my-project/global",
		},
		{
			name: "valid regional reference",
			ref:  "projects/my-project/regions/us-central1/snapshots/my-snapshot",
			want: &ComputeSnapshotIdentity{
				Project:  "my-project",
				Region:   "us-central1",
				Snapshot: "my-snapshot",
			},
			wantString: "projects/my-project/regions/us-central1/snapshots/my-snapshot",
			wantParent: "projects/my-project/regions/us-central1",
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full global url",
			ref:  "https://www.googleapis.com/compute/v1/projects/my-project/global/snapshots/my-snapshot",
			want: &ComputeSnapshotIdentity{
				Project:  "my-project",
				Snapshot: "my-snapshot",
			},
			wantString: "projects/my-project/global/snapshots/my-snapshot",
			wantParent: "projects/my-project/global",
		},
		{
			name: "full regional url",
			ref:  "https://www.googleapis.com/compute/v1/projects/my-project/regions/us-central1/snapshots/my-snapshot",
			want: &ComputeSnapshotIdentity{
				Project:  "my-project",
				Region:   "us-central1",
				Snapshot: "my-snapshot",
			},
			wantString: "projects/my-project/regions/us-central1/snapshots/my-snapshot",
			wantParent: "projects/my-project/regions/us-central1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ComputeSnapshotIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if i.String() != tt.wantString {
					t.Errorf("String() = %q, want %q", i.String(), tt.wantString)
				}
				if i.ParentString() != tt.wantParent {
					t.Errorf("ParentString() = %q, want %q", i.ParentString(), tt.wantParent)
				}
			}
		})
	}
}
