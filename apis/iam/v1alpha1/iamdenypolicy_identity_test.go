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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestIAMDenyPolicyIdentity(t *testing.T) {
	ctx := context.Background()
	reader := fake.NewClientBuilder().Build()

	tests := []struct {
		name        string
		obj         *IAMDenyPolicy
		expected    string
		expectError bool
	}{
		{
			name: "project ref",
			obj: &IAMDenyPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IAMDenyPolicySpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
				},
			},
			expected: "policies/cloudresourcemanager.googleapis.com%2Fprojects%2Fmy-project/denypolicies/my-policy",
		},
		{
			name: "folder ref",
			obj: &IAMDenyPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IAMDenyPolicySpec{
					FolderRef: &refs.FolderRef{
						External: "123456",
					},
				},
			},
			expected: "policies/cloudresourcemanager.googleapis.com%2Ffolders%2F123456/denypolicies/my-policy",
		},
		{
			name: "organization ref",
			obj: &IAMDenyPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IAMDenyPolicySpec{
					OrganizationRef: &refs.OrganizationRef{
						External: "789012",
					},
				},
			},
			expected: "policies/cloudresourcemanager.googleapis.com%2Forganizations%2F789012/denypolicies/my-policy",
		},
		{
			name: "project ref with resourceID",
			obj: &IAMDenyPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IAMDenyPolicySpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
					ResourceID: func() *string { s := "custom-policy-id"; return &s }(),
				},
			},
			expected: "policies/cloudresourcemanager.googleapis.com%2Fprojects%2Fmy-project/denypolicies/custom-policy-id",
		},
		{
			name: "project ref missing external, using name",
			obj: &IAMDenyPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IAMDenyPolicySpec{
					ProjectRef: &refs.ProjectRef{
						Name: "my-project-name",
					},
				},
			},
			expected: "policies/cloudresourcemanager.googleapis.com%2Fprojects%2Fmy-project-name/denypolicies/my-policy",
		},
		{
			name: "project ref missing external and name, fallback to namespace",
			obj: &IAMDenyPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IAMDenyPolicySpec{
					ProjectRef: &refs.ProjectRef{},
				},
			},
			expected: "policies/cloudresourcemanager.googleapis.com%2Fprojects%2Fmy-namespace/denypolicies/my-policy",
		},
		{
			name: "missing all refs",
			obj: &IAMDenyPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IAMDenyPolicySpec{},
			},
			expectError: true,
		},
		{
			name: "status mismatch",
			obj: &IAMDenyPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name:      "my-policy",
					Namespace: "my-namespace",
				},
				Spec: IAMDenyPolicySpec{
					ProjectRef: &refs.ProjectRef{
						External: "my-project",
					},
				},
				Status: IAMDenyPolicyStatus{
					ExternalRef: func() *string {
						s := "policies/cloudresourcemanager.googleapis.com%2Fprojects%2Fother-project/denypolicies/my-policy"
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
			parsed := &IAMDenyPolicyIdentity{}
			if err := parsed.FromExternal(tc.expected); err != nil {
				t.Fatalf("unexpected error parsing external: %v", err)
			}
			if parsed.String() != tc.expected {
				t.Errorf("FromExternal() -> String() mismatch: expected %q, got %q", tc.expected, parsed.String())
			}
		})
	}
}

func TestIAMDenyPolicyIdentityV2(t *testing.T) {
	var _ identity.IdentityV2 = &IAMDenyPolicyIdentity{}
	var _ identity.Resource = &IAMDenyPolicy{}
}
