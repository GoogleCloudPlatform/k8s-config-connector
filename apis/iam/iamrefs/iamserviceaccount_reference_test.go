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

package iamrefs

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestIAMServiceAccountRefNormalize(t *testing.T) {
	tests := []struct {
		name         string
		ref          *IAMServiceAccountRef
		unstructured *unstructured.Unstructured
		want         *IAMServiceAccountRef
		wantErr      bool
	}{
		{
			name: "external only",
			ref: &IAMServiceAccountRef{
				External: "sa@project.iam.gserviceaccount.com",
			},
			want: &IAMServiceAccountRef{
				External: "sa@project.iam.gserviceaccount.com",
			},
		},
		{
			name: "k8s object reference with status.email",
			ref: &IAMServiceAccountRef{
				Name:      "my-sa",
				Namespace: "ns1",
			},
			unstructured: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{
						"email": "my-sa@my-project.iam.gserviceaccount.com",
					},
				},
			},
			want: &IAMServiceAccountRef{
				External: "my-sa@my-project.iam.gserviceaccount.com",
			},
		},
		{
			name: "k8s object reference not found",
			ref: &IAMServiceAccountRef{
				Name:      "missing-sa",
				Namespace: "ns1",
			},
			wantErr: true,
		},
		{
			name: "k8s object missing status.email",
			ref: &IAMServiceAccountRef{
				Name:      "not-ready-sa",
				Namespace: "ns1",
			},
			unstructured: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"status": map[string]interface{}{},
				},
			},
			wantErr: true,
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
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if diff := cmp.Diff(tc.want, tc.ref); diff != "" {
				t.Errorf("Normalize() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIAMServiceAccountIdentityParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      *IAMServiceAccountIdentity
		wantError bool
	}{
		{
			name:  "Normal parse with account ID",
			input: "projects/myProject/serviceAccounts/my-sa",
			want: &IAMServiceAccountIdentity{
				Project: "myProject",
				Account: "my-sa",
			},
			wantError: false,
		},
		{
			name:      "Parse with email (invalid)",
			input:     "projects/myProject/serviceAccounts/my-sa@myProject.iam.gserviceaccount.com",
			want:      nil,
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &IAMServiceAccountIdentity{}
			err := id.FromExternal(tc.input)
			if tc.wantError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, id); diff != "" {
				t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
