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

func TestParseAPIGatewayAPIConfigExternal(t *testing.T) {
	tests := []struct {
		name       string
		external   string
		wantParent string
		wantID     string
		wantErr    bool
	}{
		{
			name:       "valid",
			external:   "projects/my-project/locations/global/apis/my-api/configs/my-config",
			wantParent: "projects/my-project/locations/global/apis/my-api",
			wantID:     "my-config",
		},
		{
			name:     "invalid format",
			external: "projects/my-project/locations/global/apis/my-api",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParent, gotID, err := ParseAPIGatewayAPIConfigExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAPIGatewayAPIConfigExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if gotParent.String() != tt.wantParent {
				t.Errorf("ParseAPIGatewayAPIConfigExternal() gotParent = %v, want %v", gotParent.String(), tt.wantParent)
			}
			if gotID != tt.wantID {
				t.Errorf("ParseAPIGatewayAPIConfigExternal() gotID = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}
