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

	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestVertexAIDeploymentResourcePoolRef_GetGVK(t *testing.T) {
	tests := []struct {
		name string
		want schema.GroupVersionKind
	}{
		{
			name: "success",
			want: schema.GroupVersionKind{
				Group:   "vertexai.cnrm.cloud.google.com",
				Version: "v1alpha1",
				Kind:    "VertexAIDeploymentResourcePool",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &VertexAIDeploymentResourcePoolRef{}
			if got := r.GetGVK(); got != tt.want {
				t.Errorf("GetGVK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVertexAIDeploymentResourcePoolRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/locations/us-central1/deploymentResourcePools/my-pool",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/locations/us-central1/deploymentResourcePools/my-pool",
			wantErr: true,
		},
		{
			name:    "missing location",
			ref:     "projects/my-project/deploymentResourcePools/my-pool",
			wantErr: true,
		},
		{
			name:    "missing deploymentResourcePool",
			ref:     "projects/my-project/locations/us-central1",
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
			r := &VertexAIDeploymentResourcePoolRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("VertexAIDeploymentResourcePoolRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
