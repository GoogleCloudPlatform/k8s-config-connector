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

func TestServiceAccountKeyParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      *ServiceAccountKeyIdentity
		wantError bool
	}{
		{
			name:  "Normal parse",
			input: "projects/myProject/serviceAccounts/my-sa@myProject.iam.gserviceaccount.com/keys/mykey123",
			want: &ServiceAccountKeyIdentity{
				ServiceAccountIdentity: ServiceAccountIdentity{
					Project: "myProject",
					Account: "my-sa@myProject.iam.gserviceaccount.com",
				},
				Id: "mykey123",
			},
			wantError: false,
		},
		{
			name:  "Parse with leading slash",
			input: "/projects/p1/serviceAccounts/sa@p1.iam.gserviceaccount.com/keys/k1",
			want: &ServiceAccountKeyIdentity{
				ServiceAccountIdentity: ServiceAccountIdentity{
					Project: "p1",
					Account: "sa@p1.iam.gserviceaccount.com",
				},
				Id: "k1",
			},
			wantError: false,
		},
		{
			name:  "Parse with domain",
			input: "//iam.googleapis.com/projects/first/serviceAccounts/second@first.iam.gserviceaccount.com/keys/third",
			want: &ServiceAccountKeyIdentity{
				ServiceAccountIdentity: ServiceAccountIdentity{
					Project: "first",
					Account: "second@first.iam.gserviceaccount.com",
				},
				Id: "third",
			},
			wantError: false,
		},
		{
			name:      "Empty string",
			input:     "",
			want:      nil,
			wantError: true,
		},
		{
			name:      "Wrong format - missing keys",
			input:     "projects/myProject/serviceAccounts/my-sa@myProject.iam.gserviceaccount.com",
			want:      nil,
			wantError: true,
		},
		{
			name:      "Wrong format - wrong project key",
			input:     "orgs/myProject/serviceAccounts/my-sa@myProject.iam.gserviceaccount.com/keys/mykey123",
			want:      nil,
			wantError: true,
		},
		{
			name:      "Wrong format - wrong serviceAccounts key",
			input:     "projects/myProject/users/my-sa@myProject.iam.gserviceaccount.com/keys/mykey123",
			want:      nil,
			wantError: true,
		},
		{
			name:      "Wrong format - wrong keys key",
			input:     "projects/myProject/serviceAccounts/my-sa@myProject.iam.gserviceaccount.com/tokens/mykey123",
			want:      nil,
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := &ServiceAccountKeyIdentity{}
			err := got.FromExternal(tc.input)
			// got, err := ParseServiceAccountKeyExternal(tc.input)
			if tc.wantError {
				if err == nil {
					t.Errorf("FromExternal(%q) expected error but got none", tc.input)
				}
				return
			}
			if err != nil {
				t.Errorf("FromExternal(%q) unexpected error: %v", tc.input, err)
				return
			}
			if got == nil {
				t.Errorf("FromExternal(%q) returned nil", tc.input)
				return
			}
			if got.ServiceAccountIdentity.Project != tc.want.ServiceAccountIdentity.Project {
				t.Errorf("FromExternal(%q).ServiceAccountIdentity.Project = %q, want %q",
					tc.input, got.ServiceAccountIdentity.Project, tc.want.ServiceAccountIdentity.Project)
			}
			if got.ServiceAccountIdentity.Account != tc.want.ServiceAccountIdentity.Account {
				t.Errorf("FromExternal(%q).ServiceAccountIdentity.Account = %q, want %q",
					tc.input, got.ServiceAccountIdentity.Account, tc.want.ServiceAccountIdentity.Account)
			}
			if got.Id != tc.want.Id {
				t.Errorf("FromExternal(%q).Id = %q, want %q", tc.input, got.Id, tc.want.Id)
			}
		})
	}
}
