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

package refs

import (
	"context"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDataprocMetastoreServiceRef_Basic(t *testing.T) {
	ref := &DataprocMetastoreServiceRef{
		Name:      "my-service",
		Namespace: "my-namespace",
	}

	if ref.GetGVK() != DataprocMetastoreServiceGVK {
		t.Errorf("GetGVK() = %v, want %v", ref.GetGVK(), DataprocMetastoreServiceGVK)
	}

	expectedNamespacedName := types.NamespacedName{
		Name:      "my-service",
		Namespace: "my-namespace",
	}
	if ref.GetNamespacedName() != expectedNamespacedName {
		t.Errorf("GetNamespacedName() = %v, want %v", ref.GetNamespacedName(), expectedNamespacedName)
	}

	externalRef := "projects/my-project/locations/us-central1/services/my-service"
	ref.SetExternal(externalRef)
	if ref.GetExternal() != externalRef {
		t.Errorf("GetExternal() = %v, want %v", ref.GetExternal(), externalRef)
	}

	parsedIdentity, err := ref.ParseExternalToIdentity()
	if err != nil {
		t.Errorf("ParseExternalToIdentity() unexpected error: %v", err)
	}

	expectedIdentity := &DataprocMetastoreServiceIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Service:  "my-service",
	}

	actualIdentity, ok := parsedIdentity.(*DataprocMetastoreServiceIdentity)
	if !ok {
		t.Fatalf("ParseExternalToIdentity() returned wrong type %T, want *DataprocMetastoreServiceIdentity", parsedIdentity)
	}

	if actualIdentity.Project != expectedIdentity.Project ||
		actualIdentity.Location != expectedIdentity.Location ||
		actualIdentity.Service != expectedIdentity.Service {
		t.Errorf("ParseExternalToIdentity() = %+v, want %+v", actualIdentity, expectedIdentity)
	}
}

func TestDataprocMetastoreServiceRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		wantErr  bool
	}{
		{
			name:     "valid external relative path",
			external: "projects/my-project/locations/us-central1/services/my-service",
			wantErr:  false,
		},
		{
			name:     "valid external relative path with trailing slash",
			external: "projects/my-project/locations/us-central1/services/my-service/",
			wantErr:  false,
		},
		{
			name:     "valid external host prefixed path",
			external: "metastore.googleapis.com/projects/my-project/locations/us-central1/services/my-service",
			wantErr:  false,
		},
		{
			name:     "valid external scheme and host prefixed path",
			external: "https://metastore.googleapis.com/projects/my-project/locations/us-central1/services/my-service",
			wantErr:  false,
		},
		{
			name:     "valid external double slash and host prefixed path",
			external: "//metastore.googleapis.com/projects/my-project/locations/us-central1/services/my-service",
			wantErr:  false,
		},
		{
			name:     "invalid external - missing service name",
			external: "projects/my-project/locations/us-central1/services/",
			wantErr:  true,
		},
		{
			name:     "invalid external - missing services segment",
			external: "projects/my-project/locations/us-central1/",
			wantErr:  true,
		},
		{
			name:     "invalid external - empty location",
			external: "projects/my-project/locations//services/my-service",
			wantErr:  true,
		},
		{
			name:     "invalid external - empty project",
			external: "projects//locations/us-central1/services/my-service",
			wantErr:  true,
		},
		{
			name:     "invalid external - too many segments",
			external: "projects/my-project/locations/us-central1/services/my-service/extra",
			wantErr:  true,
		},
		{
			name:     "invalid external - wrong prefix",
			external: "other/my-project/locations/us-central1/services/my-service",
			wantErr:  true,
		},
		{
			name:     "invalid external - wrong host",
			external: "https://other.googleapis.com/projects/my-project/locations/us-central1/services/my-service",
			wantErr:  true,
		},
		{
			name:     "invalid external - empty string",
			external: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DataprocMetastoreServiceRef{}
			err := r.ValidateExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataprocMetastoreServiceRef_Normalize(t *testing.T) {
	s := runtime.NewScheme()

	service := &unstructured.Unstructured{}
	service.SetGroupVersionKind(DataprocMetastoreServiceGVK)
	service.SetName("my-service")
	service.SetNamespace("my-ns")
	service.Object["spec"] = map[string]interface{}{
		"resourceID": "my-service-id",
		"location":   "us-central1",
		"projectRef": map[string]interface{}{
			"external": "my-project",
		},
	}

	reader := fake.NewClientBuilder().WithScheme(s).WithObjects(service).Build()

	tests := []struct {
		name             string
		ref              *DataprocMetastoreServiceRef
		defaultNamespace string
		want             string
		wantErr          bool
	}{
		{
			name: "external reference",
			ref: &DataprocMetastoreServiceRef{
				External: "projects/my-project/locations/us-central1/services/my-service",
			},
			want: "projects/my-project/locations/us-central1/services/my-service",
		},
		{
			name: "internal reference with namespace",
			ref: &DataprocMetastoreServiceRef{
				Name:      "my-service",
				Namespace: "my-ns",
			},
			want: "projects/my-project/locations/us-central1/services/my-service-id",
		},
		{
			name: "internal reference with default namespace",
			ref: &DataprocMetastoreServiceRef{
				Name: "my-service",
			},
			defaultNamespace: "my-ns",
			want:             "projects/my-project/locations/us-central1/services/my-service-id",
		},
		{
			name: "internal reference not found",
			ref: &DataprocMetastoreServiceRef{
				Name:      "non-existent",
				Namespace: "my-ns",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(context.Background(), reader, tt.defaultNamespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if tt.ref.External != tt.want {
					t.Errorf("Normalize() got = %v, want %v", tt.ref.External, tt.want)
				}
			}
		})
	}
}
