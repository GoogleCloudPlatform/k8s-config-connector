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
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestIamAccessPolicyIdentity(t *testing.T) {
	ctx := context.Background()
	reader := fake.NewClientBuilder().Build()

	tests := []struct {
		name        string
		obj         *IamAccessPolicy
		expected    string
		expectError bool
	}{
		{
			name: "organization ref with location and resourceID",
			obj: &IamAccessPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IamAccessPolicySpec{
					OrganizationRef: &refsv1beta1.OrganizationRef{
						External: "organizations/123456",
					},
					Location:   "global",
					ResourceID: func() *string { s := "custom-policy"; return &s }(),
				},
			},
			expected: "organizations/123456/locations/global/accessPolicies/custom-policy",
		},
		{
			name: "organization ref with location, using name",
			obj: &IamAccessPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IamAccessPolicySpec{
					OrganizationRef: &refsv1beta1.OrganizationRef{
						External: "organizations/123456",
					},
					Location: "global",
				},
			},
			expected: "organizations/123456/locations/global/accessPolicies/my-policy",
		},
		{
			name: "missing organizationRef",
			obj: &IamAccessPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IamAccessPolicySpec{
					Location: "global",
				},
			},
			expectError: true,
		},
		{
			name: "missing location",
			obj: &IamAccessPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IamAccessPolicySpec{
					OrganizationRef: &refsv1beta1.OrganizationRef{
						External: "organizations/123456",
					},
				},
			},
			expectError: true,
		},
		{
			name: "status mismatch",
			obj: &IamAccessPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IamAccessPolicySpec{
					OrganizationRef: &refsv1beta1.OrganizationRef{
						External: "organizations/123456",
					},
					Location: "global",
				},
				Status: IamAccessPolicyStatus{
					ExternalRef: func() *string {
						s := "organizations/123456/locations/global/accessPolicies/other-policy"
						return &s
					}(),
				},
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.obj.GetIdentity(ctx, reader)
			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if id.String() != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, id.String())
			}

			// test FromExternal
			parsed := &IamAccessPolicyIdentity{}
			if err := parsed.FromExternal(tc.expected); err != nil {
				t.Fatalf("unexpected error parsing external: %v", err)
			}
			if parsed.String() != tc.expected {
				t.Errorf("FromExternal() -> String() mismatch: expected %q, got %q", tc.expected, parsed.String())
			}
		})
	}
}

func TestIamAccessPolicyIdentityV2(t *testing.T) {
	var _ identity.IdentityV2 = &IamAccessPolicyIdentity{}
	var _ identity.Resource = &IamAccessPolicy{}
}
