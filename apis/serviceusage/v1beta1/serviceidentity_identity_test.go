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

func TestServiceIdentityIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *ServiceIdentityIdentity
	}{
		{
			name: "valid reference (new format)",
			ref:  "projects/my-project/services/pubsub.googleapis.com/identity",
			want: &ServiceIdentityIdentity{
				Project: "my-project",
				Service: "pubsub.googleapis.com",
			},
		},
		{
			name: "full url (new format)",
			ref:  "https://serviceusage.googleapis.com/projects/my-project/services/pubsub.googleapis.com/identity",
			want: &ServiceIdentityIdentity{
				Project: "my-project",
				Service: "pubsub.googleapis.com",
			},
		},
		{
			name: "legacy format for backward compatibility",
			ref:  "projects/my-project/locations/global/serviceidentitys/pubsub.googleapis.com",
			want: &ServiceIdentityIdentity{
				Project: "my-project",
				Service: "pubsub.googleapis.com",
			},
		},
		{
			name: "legacy format full url",
			ref:  "https://serviceusage.googleapis.com/projects/my-project/locations/global/serviceidentitys/pubsub.googleapis.com",
			want: &ServiceIdentityIdentity{
				Project: "my-project",
				Service: "pubsub.googleapis.com",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ServiceIdentityIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
