// Copyright 2025 Google LLC
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

package sql

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestDetermineMaintenanceVersion(t *testing.T) {
	tests := []struct {
		name                 string
		userSpecifiedVersion *string
		actualVersion        string
		availableVersions    []string
		expectedVersion      string
		description          string
	}{
		{
			name:                 "user specifies version",
			userSpecifiedVersion: direct.LazyPtr("user-version"),
			actualVersion:        "actual-version",
			availableVersions:    []string{"available-1", "available-2"},
			expectedVersion:      "user-version",
			description:          "User specified version should be used regardless of actual/available",
		},
		{
			name:                 "available versions not populated - assume actual is valid",
			userSpecifiedVersion: nil,
			actualVersion:        "actual-version",
			availableVersions:    []string{},
			expectedVersion:      "actual-version",
			description:          "When AvailableMaintenanceVersions is empty, assume actual version is valid",
		},
		{
			name:                 "actual version is still valid",
			userSpecifiedVersion: nil,
			actualVersion:        "valid-version",
			availableVersions:    []string{"valid-version", "other-version"},
			expectedVersion:      "valid-version",
			description:          "When actual version is in available list, use it",
		},
		{
			name:                 "actual version is retired - pick first available",
			userSpecifiedVersion: nil,
			actualVersion:        "retired-version",
			availableVersions:    []string{"available-1", "available-2"},
			expectedVersion:      "available-1",
			description:          "When actual version is retired, pick first available version",
		},
		{
			name:                 "no actual version",
			userSpecifiedVersion: nil,
			actualVersion:        "",
			availableVersions:    []string{},
			expectedVersion:      "",
			description:          "When no actual version, don't set maintenance version",
		},
		{
			name:                 "empty actual version with available versions",
			userSpecifiedVersion: nil,
			actualVersion:        "",
			availableVersions:    []string{"available-1", "available-2"},
			expectedVersion:      "",
			description:          "When actual version is empty, don't set maintenance version even if available versions exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := determineMaintenanceVersion(tt.userSpecifiedVersion, tt.actualVersion, tt.availableVersions)

			if result != tt.expectedVersion {
				t.Errorf("MaintenanceVersion mismatch:\nExpected: %s\nGot: %s\nDescription: %s",
					tt.expectedVersion, result, tt.description)
			}
		})
	}
}
