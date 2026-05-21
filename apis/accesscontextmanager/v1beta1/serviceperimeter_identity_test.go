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
)

func TestServicePerimeterIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     *ServicePerimeterIdentity
		wantStr  string
		wantErr  bool
	}{
		{
			name:     "canonical",
			external: "accessPolicies/12345/servicePerimeters/my-perimeter",
			want: &ServicePerimeterIdentity{
				AccessPolicy:     "12345",
				ServicePerimeter: "my-perimeter",
			},
			wantStr: "accessPolicies/12345/servicePerimeters/my-perimeter",
		},
		{
			name:     "full URL",
			external: "//accesscontextmanager.googleapis.com/accessPolicies/12345/servicePerimeters/my-perimeter",
			want: &ServicePerimeterIdentity{
				AccessPolicy:     "12345",
				ServicePerimeter: "my-perimeter",
			},
			wantStr: "accessPolicies/12345/servicePerimeters/my-perimeter",
		},
		{
			name:     "invalid format",
			external: "projects/my-project/locations/global/servicePerimeters/my-perimeter",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ServicePerimeterIdentity{}
			err := i.FromExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if i.AccessPolicy != tt.want.AccessPolicy || i.ServicePerimeter != tt.want.ServicePerimeter {
				t.Errorf("FromExternal() got = %v, want %v", i, tt.want)
			}
			if i.String() != tt.wantStr {
				t.Errorf("String() got = %v, want %v", i.String(), tt.wantStr)
			}
		})
	}
}
