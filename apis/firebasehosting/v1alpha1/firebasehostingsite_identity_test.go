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

func TestFirebaseHostingSiteIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *FirebaseHostingSiteIdentity
	}{
		{
			name: "valid relative reference",
			ref:  "projects/my-project/sites/my-site",
			want: &FirebaseHostingSiteIdentity{
				Project: "my-project",
				Site:    "my-site",
			},
		},
		{
			name: "valid full url with https",
			ref:  "https://firebasehosting.googleapis.com/projects/my-project/sites/my-site",
			want: &FirebaseHostingSiteIdentity{
				Project: "my-project",
				Site:    "my-site",
			},
		},
		{
			name: "valid full url with //",
			ref:  "//firebasehosting.googleapis.com/projects/my-project/sites/my-site",
			want: &FirebaseHostingSiteIdentity{
				Project: "my-project",
				Site:    "my-site",
			},
		},
		{
			name:    "invalid format",
			ref:     "invalid/format",
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
			i := &FirebaseHostingSiteIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Site != tt.want.Site {
					t.Errorf("Site = %v, want %v", i.Site, tt.want.Site)
				}
				if got := i.String(); got != "projects/my-project/sites/my-site" {
					t.Errorf("String() = %v, want %v", got, "projects/my-project/sites/my-site")
				}
			}
		})
	}
}
