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

	"github.com/google/go-cmp/cmp"
)

func TestDataProductIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name            string
		ref             string
		wantProject     string
		wantLocation    string
		wantDataProduct string
		wantErr         bool
	}{
		{
			name:            "valid external ref",
			ref:             "projects/my-project/locations/us-central1/dataProducts/my-data-product",
			wantProject:     "my-project",
			wantLocation:    "us-central1",
			wantDataProduct: "my-data-product",
			wantErr:         false,
		},
		{
			name:    "invalid external ref (missing location)",
			ref:     "projects/my-project/dataProducts/my-data-product",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DataProductIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Fatalf("DataProductIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				got := &DataProductIdentity{
					Project:     i.Project,
					Location:    i.Location,
					DataProduct: i.DataProduct,
				}
				want := &DataProductIdentity{
					Project:     tt.wantProject,
					Location:    tt.wantLocation,
					DataProduct: tt.wantDataProduct,
				}
				if diff := cmp.Diff(want, got); diff != "" {
					t.Errorf("DataProductIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if gotStr := i.String(); gotStr != tt.ref {
					t.Errorf("DataProductIdentity.String() = %v, want %v", gotStr, tt.ref)
				}
			}
		})
	}
}
