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
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// mockReader implements client.Reader for testing
type mockReader struct {
	client.Reader
}

func (m *mockReader) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	return nil
}

func TestApigeeAPIProductIdentity(t *testing.T) {
	ctx := context.TODO()
	reader := &mockReader{}

	tests := []struct {
		name     string
		obj      *v1alpha1.ApigeeAPIProduct
		expected string
		wantErr  bool
	}{
		{
			name: "basic format",
			obj: &v1alpha1.ApigeeAPIProduct{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-apiproduct",
					Namespace: "test-namespace",
				},
				Spec: v1alpha1.ApigeeAPIProductSpec{
					OrganizationRef: &apigeev1beta1.ApigeeOrganizationRef{
						External: "organizations/test-org",
					},
				},
			},
			expected: "organizations/test-org/apiproducts/test-apiproduct",
		},
		{
			name: "with project id format",
			obj: &v1alpha1.ApigeeAPIProduct{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-apiproduct",
					Namespace: "test-namespace",
				},
				Spec: v1alpha1.ApigeeAPIProductSpec{
					OrganizationRef: &apigeev1beta1.ApigeeOrganizationRef{
						External: "projects/test-project",
					},
				},
			},
			expected: "organizations/test-project/apiproducts/test-apiproduct",
		},
		{
			name: "with resource id",
			obj: &v1alpha1.ApigeeAPIProduct{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-apiproduct-krm-name",
					Namespace: "test-namespace",
				},
				Spec: v1alpha1.ApigeeAPIProductSpec{
					OrganizationRef: &apigeev1beta1.ApigeeOrganizationRef{
						External: "organizations/test-org",
					},
					ResourceID: ptr("test-apiproduct"),
				},
			},
			expected: "organizations/test-org/apiproducts/test-apiproduct",
		},
		{
			name: "missing organization ref",
			obj: &v1alpha1.ApigeeAPIProduct{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-apiproduct",
					Namespace: "test-namespace",
				},
				Spec: v1alpha1.ApigeeAPIProductSpec{},
			},
			wantErr: true,
		},
		{
			name: "cross-check with status external ref success",
			obj: &v1alpha1.ApigeeAPIProduct{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-apiproduct",
					Namespace: "test-namespace",
				},
				Spec: v1alpha1.ApigeeAPIProductSpec{
					OrganizationRef: &apigeev1beta1.ApigeeOrganizationRef{
						External: "organizations/test-org",
					},
				},
				Status: v1alpha1.ApigeeAPIProductStatus{
					ExternalRef: ptr("organizations/test-org/apiproducts/test-apiproduct"),
				},
			},
			expected: "organizations/test-org/apiproducts/test-apiproduct",
		},
		{
			name: "cross-check with status external ref failure",
			obj: &v1alpha1.ApigeeAPIProduct{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-apiproduct",
					Namespace: "test-namespace",
				},
				Spec: v1alpha1.ApigeeAPIProductSpec{
					OrganizationRef: &apigeev1beta1.ApigeeOrganizationRef{
						External: "organizations/test-org",
					},
				},
				Status: v1alpha1.ApigeeAPIProductStatus{
					ExternalRef: ptr("organizations/test-org/apiproducts/different-apiproduct"),
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.obj.GetIdentity(ctx, reader)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if id.String() != tc.expected {
				t.Errorf("expected identity %q, got %q", tc.expected, id.String())
			}
		})
	}
}

func TestApigeeAPIProductIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		ref      string
		expected v1alpha1.ApigeeAPIProductIdentity
		wantErr  bool
	}{
		{
			name: "valid ref",
			ref:  "organizations/test-org/apiproducts/test-apiproduct",
			expected: v1alpha1.ApigeeAPIProductIdentity{
				Organization: "test-org",
				Apiproduct:   "test-apiproduct",
			},
		},
		{
			name:    "invalid format",
			ref:     "projects/test-project/locations/us-central1/apiproducts/test-apiproduct",
			wantErr: true,
		},
		{
			name:    "missing apiproduct",
			ref:     "organizations/test-org/apiproducts/",
			wantErr: true,
		},
		{
			name:    "missing organization",
			ref:     "organizations//apiproducts/test-apiproduct",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &v1alpha1.ApigeeAPIProductIdentity{}
			err := id.FromExternal(tc.ref)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if *id != tc.expected {
				t.Errorf("expected %+v, got %+v", tc.expected, *id)
			}
		})
	}
}
func ptr(s string) *string {
	return &s
}
