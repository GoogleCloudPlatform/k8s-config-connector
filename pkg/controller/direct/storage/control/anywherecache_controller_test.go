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

package storagecontrol

import (
	"testing"

	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestIsStateChangeRequested(t *testing.T) {
	tests := []struct {
		name         string
		actualState  string
		desiredState *string // Use pointer to simulate nil for default
		expected     bool
	}{
		{
			name:         "states are different (running to paused)",
			actualState:  anywhereCacheStateRunning,
			desiredState: direct.LazyPtr(anywhereCacheStatePaused),
			expected:     true,
		},
		{
			name:         "states are different (paused to running)",
			actualState:  anywhereCacheStatePaused,
			desiredState: direct.LazyPtr(anywhereCacheStateRunning),
			expected:     true,
		},
		{
			name:         "states are different (disabled to running)",
			actualState:  anywhereCacheStateDisabled,
			desiredState: direct.LazyPtr(anywhereCacheStateRunning),
			expected:     true,
		},
		{
			name:         "states are the same (running to running)",
			actualState:  anywhereCacheStateRunning,
			desiredState: direct.LazyPtr(anywhereCacheStateRunning),
			expected:     false,
		},
		{
			name:         "states are the same (paused to paused)",
			actualState:  anywhereCacheStatePaused,
			desiredState: direct.LazyPtr(anywhereCacheStatePaused),
			expected:     false,
		},
		{
			name:         "desired state is nil (defaults to running) and actual is running",
			actualState:  anywhereCacheStateRunning,
			desiredState: nil,
			expected:     false,
		},
		{
			name:         "desired state is nil (defaults to running) and actual is paused",
			actualState:  anywhereCacheStatePaused,
			desiredState: nil,
			expected:     true,
		},
		{
			name:         "desired state is nil (defaults to running) and actual is creating",
			actualState:  anywhereCacheStateCreating,
			desiredState: nil,
			expected:     true,
		},
		{
			name:         "actual state is creating, desired is running",
			actualState:  anywhereCacheStateCreating,
			desiredState: direct.LazyPtr(anywhereCacheStateRunning),
			expected:     true,
		},
		{
			name:         "actual state is running, desired is creating (invalid but tests comparison)",
			actualState:  anywhereCacheStateRunning,
			desiredState: direct.LazyPtr(anywhereCacheStateCreating),
			expected:     true,
		},
		{
			name:         "actual state is invalid, desired is running",
			actualState:  anywhereCacheStateInvalid,
			desiredState: direct.LazyPtr(anywhereCacheStateRunning),
			expected:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			adapter := &AnywhereCacheAdapter{
				actual: &pb.AnywhereCache{
					State: tc.actualState,
				},
				desired: &krm.StorageAnywhereCache{
					Spec: krm.StorageAnywhereCacheSpec{
						DesiredState: tc.desiredState,
					},
				},
			}

			if got := adapter.IsStateChangeRequested(); got != tc.expected {
				t.Errorf("IsStateChangeRequested() got = %v, want %v for actualState '%s' and desiredState '%v'", got, tc.expected, tc.actualState, tc.desiredState)
			}
		})
	}
}
