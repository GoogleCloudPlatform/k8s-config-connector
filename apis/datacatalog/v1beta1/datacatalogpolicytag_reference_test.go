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
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDataCatalogPolicyTagRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/locations/us-central1/taxonomies/my-taxonomy/policyTags/my-policytag",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/locations/us-central1/taxonomies/my-taxonomy/policyTags/my-policytag",
			wantErr: true,
		},
		{
			name:    "missing location",
			ref:     "projects/my-project/taxonomies/my-taxonomy/policyTags/my-policytag",
			wantErr: true,
		},
		{
			name:    "missing policytag",
			ref:     "projects/my-project/locations/us-central1/taxonomies/my-taxonomy",
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
			r := &DataCatalogPolicyTagRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("DataCatalogPolicyTagRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataCatalogPolicyTagRef_Normalize(t *testing.T) {
	testCases := []struct {
		name           string
		ref            *DataCatalogPolicyTagRef
		otherNamespace string
		objects        []runtime.Object
		wantExternal   string
		wantErr        string
	}{
		{
			name: "external with valid format",
			ref: &DataCatalogPolicyTagRef{
				External: "projects/my-project/locations/us-central1/taxonomies/my-taxonomy/policyTags/my-policytag",
			},
			wantExternal: "projects/my-project/locations/us-central1/taxonomies/my-taxonomy/policyTags/my-policytag",
		},
		{
			name: "external with invalid format",
			ref: &DataCatalogPolicyTagRef{
				External: "invalid-format",
			},
			wantErr: `format of DataCatalogPolicyTag external="invalid-format" was not known (use projects/{project}/locations/{location}/taxonomies/{taxonomy}/policyTags/{policyTag})`,
		},
		{
			name: "name specified, with status.externalRef",
			ref: &DataCatalogPolicyTagRef{
				Name:      "test-policytag",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "datacatalog.cnrm.cloud.google.com/v1beta1",
						"kind":       "DataCatalogPolicyTag",
						"metadata": map[string]interface{}{
							"name":      "test-policytag",
							"namespace": "my-namespace",
						},
						"status": map[string]interface{}{
							"externalRef": "projects/my-project/locations/us-central1/taxonomies/my-taxonomy/policyTags/test-policytag",
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/us-central1/taxonomies/my-taxonomy/policyTags/test-policytag",
		},
		{
			name: "name specified, with status.externalRef missing but fallback works",
			ref: &DataCatalogPolicyTagRef{
				Name:      "test-policytag",
				Namespace: "my-namespace",
			},
			objects: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"apiVersion": "datacatalog.cnrm.cloud.google.com/v1beta1",
						"kind":       "DataCatalogPolicyTag",
						"metadata": map[string]interface{}{
							"name":      "test-policytag",
							"namespace": "my-namespace",
						},
						"spec": map[string]interface{}{
							"taxonomyRef": map[string]interface{}{
								"external": "projects/my-project/locations/us-central1/taxonomies/my-taxonomy",
							},
						},
					},
				},
			},
			wantExternal: "projects/my-project/locations/us-central1/taxonomies/my-taxonomy/policyTags/test-policytag",
		},
		{
			name: "name specified, resource not found",
			ref: &DataCatalogPolicyTagRef{
				Name:      "test-policytag",
				Namespace: "my-namespace",
			},
			wantErr: `reference DataCatalogPolicyTag my-namespace/test-policytag is not found`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := runtime.NewScheme()
			s.AddKnownTypes(schema.GroupVersion{Group: "datacatalog.cnrm.cloud.google.com", Version: "v1beta1"}, &unstructured.Unstructured{})
			fakeClient := fake.NewClientBuilder().WithScheme(s).WithRuntimeObjects(tc.objects...).Build()

			err := tc.ref.Normalize(context.TODO(), fakeClient, tc.otherNamespace)
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("got nil error, want %q", tc.wantErr)
				}
				if !cmp.Equal(err.Error(), tc.wantErr) {
					t.Errorf("got error %q, want %q", err.Error(), tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			gotExternal := tc.ref.External
			if gotExternal != tc.wantExternal {
				t.Errorf("got external %q, want %q", gotExternal, tc.wantExternal)
			}
		})
	}
}

func TestDataCatalogPolicyTagRef_ParseExternalToIdentity(t *testing.T) {
	ref := &DataCatalogPolicyTagRef{
		External: "projects/my-project/locations/us-central1/taxonomies/my-taxonomy/policyTags/my-policytag",
	}
	id, err := ref.ParseExternalToIdentity()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	concreteID, ok := id.(*DataCatalogPolicyTagIdentity)
	if !ok {
		t.Fatalf("expected *DataCatalogPolicyTagIdentity, got %T", id)
	}
	if concreteID.Project != "my-project" || concreteID.Location != "us-central1" || concreteID.Taxonomy != "my-taxonomy" || concreteID.PolicyTag != "my-policytag" {
		t.Errorf("parsed identity fields got %+v, want expected values", concreteID)
	}
}
