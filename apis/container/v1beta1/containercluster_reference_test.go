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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
)

func TestContainerClusterRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid regional reference",
			ref:     "projects/my-project/locations/us-central1/clusters/my-cluster",
			wantErr: false,
		},
		{
			name:    "valid zonal reference",
			ref:     "projects/my-project/zones/us-central1-a/clusters/my-cluster",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/locations/us-central1/clusters/my-cluster",
			wantErr: true,
		},
		{
			name:    "missing location/zone",
			ref:     "projects/my-project/clusters/my-cluster",
			wantErr: true,
		},
		{
			name:    "missing cluster name",
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
			r := &v1beta1.ContainerClusterRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ContainerClusterRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
