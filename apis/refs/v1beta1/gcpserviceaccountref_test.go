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
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestServiceAccountRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid relative resource name",
			ref:     "projects/test-project/serviceAccounts/test-sa@test-project.iam.gserviceaccount.com",
			wantErr: false,
		},
		{
			name:    "valid legacy email format",
			ref:     "test-sa@test-project.iam.gserviceaccount.com",
			wantErr: false,
		},
		{
			name:    "invalid format empty",
			ref:     "",
			wantErr: true,
		},
		{
			name:    "invalid format random text",
			ref:     "invalid-format",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ServiceAccountRef{}
			err := r.ValidateExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal(%q) error = %v, wantErr %v", tt.ref, err, tt.wantErr)
			}
		})
	}
}

func TestServiceAccountRef_Normalize(t *testing.T) {
	tests := []struct {
		name         string
		ref          *ServiceAccountRef
		unstructured *unstructured.Unstructured
		want         *ServiceAccountRef
		wantErr      bool
	}{
		{
			name: "external already set with relative resource name",
			ref: &ServiceAccountRef{
				External: "projects/test-project/serviceAccounts/test-sa@test-project.iam.gserviceaccount.com",
			},
			want: &ServiceAccountRef{
				External: "projects/test-project/serviceAccounts/test-sa@test-project.iam.gserviceaccount.com",
			},
			wantErr: false,
		},
		{
			name: "external already set with legacy email",
			ref: &ServiceAccountRef{
				External: "test-sa@test-project.iam.gserviceaccount.com",
			},
			want: &ServiceAccountRef{
				External: "test-sa@test-project.iam.gserviceaccount.com",
			},
			wantErr: false,
		},
		{
			name: "resolve from cluster using status.name",
			ref: &ServiceAccountRef{
				Name:      "test-sa",
				Namespace: "test-ns",
			},
			unstructured: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{
						"name": "projects/test-project/serviceAccounts/test-sa@test-project.iam.gserviceaccount.com",
					},
				},
			},
			want: &ServiceAccountRef{
				External: "projects/test-project/serviceAccounts/test-sa@test-project.iam.gserviceaccount.com",
			},
			wantErr: false,
		},
		{
			name: "resolve from cluster using status.email fallback",
			ref: &ServiceAccountRef{
				Name:      "test-sa2",
				Namespace: "test-ns",
			},
			unstructured: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": map[string]interface{}{
						"projectRef": map[string]interface{}{
							"external": "test-project",
						},
					},
					"status": map[string]interface{}{
						"email": "test-sa2@test-project.iam.gserviceaccount.com",
					},
				},
			},
			want: &ServiceAccountRef{
				External: "projects/test-project/serviceAccounts/test-sa2@test-project.iam.gserviceaccount.com",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()

			var objs []unstructured.Unstructured
			if tc.unstructured != nil {
				tc.unstructured.SetName(tc.ref.Name)
				tc.unstructured.SetNamespace(tc.ref.Namespace)
				tc.unstructured.SetGroupVersionKind(IAMServiceAccountGVK)
				objs = append(objs, *tc.unstructured)
			}

			s := fake.NewClientBuilder().WithLists(&unstructured.UnstructuredList{Items: objs}).Build()

			err := tc.ref.Normalize(ctx, s, tc.ref.Namespace)
			if (err != nil) != tc.wantErr {
				t.Fatalf("unexpected error result: %v, wantErr %v", err, tc.wantErr)
			}

			if !tc.wantErr {
				if diff := cmp.Diff(tc.ref, tc.want); diff != "" {
					t.Errorf("Normalize() mismatch (-got +want):\n%s", diff)
				}
			}
		})
	}
}
