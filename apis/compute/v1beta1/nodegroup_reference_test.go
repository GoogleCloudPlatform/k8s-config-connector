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

package v1beta1_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
)

func TestComputeNodeGroupRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid relative path",
			ref:     "projects/my-project/zones/us-central1-a/nodeGroups/my-node-group",
			wantErr: false,
		},
		{
			name:    "valid full URL",
			ref:     "https://www.googleapis.com/compute/v1/projects/my-project/zones/us-central1-a/nodeGroups/my-node-group",
			wantErr: false,
		},
		{
			name:    "valid full URL with beta",
			ref:     "https://www.googleapis.com/compute/beta/projects/my-project/zones/us-central1-a/nodeGroups/my-node-group",
			wantErr: false,
		},
		{
			name:    "invalid format",
			ref:     "projects/my-project/nodeGroups/my-node-group",
			wantErr: true,
		},
		{
			name:    "missing zone",
			ref:     "projects/my-project/zones/nodeGroups/my-node-group",
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
			r := &v1beta1.ComputeNodeGroupRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeNodeGroupRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
