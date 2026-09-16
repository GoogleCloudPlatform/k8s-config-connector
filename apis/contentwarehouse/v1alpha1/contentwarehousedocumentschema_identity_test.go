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

package v1alpha1_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
	"github.com/google/go-cmp/cmp"
)

func TestContentWarehouseDocumentSchemaIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *v1alpha1.ContentWarehouseDocumentSchemaIdentity
		wantErr bool
	}{
		{
			name: "valid external ref",
			ref:  "projects/my-project/locations/us-central1/documentSchemas/my-documentschema",
			want: &v1alpha1.ContentWarehouseDocumentSchemaIdentity{
				Project:        "my-project",
				Location:       "us-central1",
				DocumentSchema: "my-documentschema",
			},
			wantErr: false,
		},
		{
			name: "full url",
			ref:  "https://contentwarehouse.googleapis.com/projects/my-project/locations/us-central1/documentSchemas/my-documentschema",
			want: &v1alpha1.ContentWarehouseDocumentSchemaIdentity{
				Project:        "my-project",
				Location:       "us-central1",
				DocumentSchema: "my-documentschema",
			},
			wantErr: false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/documentSchemas/my-documentschema",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &v1alpha1.ContentWarehouseDocumentSchemaIdentity{}
			err := got.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ContentWarehouseDocumentSchemaIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("ContentWarehouseDocumentSchemaIdentity.FromExternal() mismatch (-want +got):\n%s", diff)
				}
				expectedString := "projects/" + tt.want.Project + "/locations/" + tt.want.Location + "/documentSchemas/" + tt.want.DocumentSchema
				if str := got.String(); str != expectedString {
					t.Errorf("ContentWarehouseDocumentSchemaIdentity.String() = %v, want %v", str, expectedString)
				}
			}
		})
	}
}
