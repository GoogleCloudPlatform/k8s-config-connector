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

package dataform

import (
	"testing"

	dataformpb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
)

func TestGitRemoteSettingsEqual(t *testing.T) {
	tests := []struct {
		name    string
		desired *dataformpb.Repository_GitRemoteSettings
		actual  *dataformpb.Repository_GitRemoteSettings
		want    bool
	}{
		{
			name:    "both nil",
			desired: nil,
			actual:  nil,
			want:    true,
		},
		{
			name: "equal when actual has output-only TokenStatus populated",
			desired: &dataformpb.Repository_GitRemoteSettings{
				Url:           "https://github.com/example/repo.git",
				DefaultBranch: "main",
			},
			actual: &dataformpb.Repository_GitRemoteSettings{
				Url:           "https://github.com/example/repo.git",
				DefaultBranch: "main",
				TokenStatus:   dataformpb.Repository_GitRemoteSettings_VALID,
			},
			want: true,
		},
		{
			name: "not equal when default_branch differs",
			desired: &dataformpb.Repository_GitRemoteSettings{
				Url:           "https://github.com/example/repo.git",
				DefaultBranch: "develop",
			},
			actual: &dataformpb.Repository_GitRemoteSettings{
				Url:           "https://github.com/example/repo.git",
				DefaultBranch: "main",
				TokenStatus:   dataformpb.Repository_GitRemoteSettings_VALID,
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := gitRemoteSettingsEqual(tc.desired, tc.actual); got != tc.want {
				t.Errorf("gitRemoteSettingsEqual() = %v, want %v", got, tc.want)
			}
		})
	}
}
