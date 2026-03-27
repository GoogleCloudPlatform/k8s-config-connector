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

func TestComputeNodeTemplateRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/regions/us-central1/nodeTemplates/my-node-template",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/regions/us-central1/nodeTemplates/my-node-template",
			wantErr: true,
		},
		{
			name:    "missing region",
			ref:     "projects/my-project/nodeTemplates/my-node-template",
			wantErr: true,
		},
		{
			name:    "missing nodeTemplate",
			ref:     "projects/my-project/regions/us-central1",
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
			r := &ComputeNodeTemplateRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ComputeNodeTemplateRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
