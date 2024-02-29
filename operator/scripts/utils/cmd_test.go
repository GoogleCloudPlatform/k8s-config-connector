// Copyright 2024 Google LLC
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

package utils

import "testing"

func Test_parseKustomizeVersion(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantVersion string
		wantErr     bool
	}{
		{
			name:        "valid version",
			input:       "{kustomize/v3.5.4  2020-01-11T03:12:59Z  }",
			wantVersion: "v3.5.4",
		},
		{
			name:    "invalid version",
			input:   "",
			wantErr: true,
		},
		{
			name:        "newer valid version",
			input:       "{v5.3.0  2023-12-07T10:45:14Z   }", // notice the lack of "kustomize/"
			wantVersion: "v5.3.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			version, err := parseKustomizeVersion(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("expected nil error got %v", err)
			}
			if version != tt.wantVersion {
				t.Fatalf("failed parsing want: %s, got: %s", tt.wantVersion, version)
			}
		})
	}
}
