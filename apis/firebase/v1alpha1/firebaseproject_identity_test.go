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

func TestFirebaseProjectIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *FirebaseProjectIdentity
		wantErr bool
	}{
		{
			name: "full url with scheme and host",
			ref:  "https://firebase.googleapis.com/projects/my-project",
			want: &FirebaseProjectIdentity{
				Project: "my-project",
			},
		},
		{
			name: "url without scheme",
			ref:  "//firebase.googleapis.com/projects/my-project",
			want: &FirebaseProjectIdentity{
				Project: "my-project",
			},
		},
		{
			name: "relative path",
			ref:  "projects/my-project",
			want: &FirebaseProjectIdentity{
				Project: "my-project",
			},
		},
		{
			name:    "invalid format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "too few segments",
			ref:     "projects",
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &FirebaseProjectIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if i.String() != tt.want.String() {
					t.Errorf("String() = %q, want %q", i.String(), tt.want.String())
				}
			}
		})
	}
}

func TestFirebaseProjectRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name: "valid relative external ref",
			ref:  "projects/my-project",
		},
		{
			name: "valid full external ref",
			ref:  "//firebase.googleapis.com/projects/my-project",
		},
		{
			name:    "invalid external ref",
			ref:     "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &FirebaseProjectRef{
				External: tt.ref,
			}
			err := r.ValidateExternal(r.External)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
