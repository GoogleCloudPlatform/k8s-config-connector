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
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestDNSResponsePolicyRuleIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DNSResponsePolicyRuleIdentity
	}{
		{
			name: "valid reference (with location)",
			ref:  "projects/my-project/locations/global/responsePolicies/my-responsepolicy/rules/my-rule",
			want: &DNSResponsePolicyRuleIdentity{
				Project:        "my-project",
				Location:       "global",
				ResponsePolicy: "my-responsepolicy",
				Rule:           "my-rule",
			},
		},
		{
			name: "valid reference (without location)",
			ref:  "projects/my-project/responsePolicies/my-responsepolicy/rules/my-rule",
			want: &DNSResponsePolicyRuleIdentity{
				Project:        "my-project",
				Location:       "",
				ResponsePolicy: "my-responsepolicy",
				Rule:           "my-rule",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url (with location)",
			ref:  "https://dns.googleapis.com/projects/my-project/locations/global/responsePolicies/my-responsepolicy/rules/my-rule",
			want: &DNSResponsePolicyRuleIdentity{
				Project:        "my-project",
				Location:       "global",
				ResponsePolicy: "my-responsepolicy",
				Rule:           "my-rule",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DNSResponsePolicyRuleIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

type mockReader struct {
	client.Reader
}

func TestGetIdentityFromDNSResponsePolicyRuleSpec(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *DNSResponsePolicyRule
		want    *DNSResponsePolicyRuleIdentity
		wantErr string
	}{
		{
			name: "invalid spec with simple responsePolicy name",
			obj: &DNSResponsePolicyRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-rule",
					Namespace: "my-namespace",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: DNSResponsePolicyRuleSpec{
					ResponsePolicy: "my-policy",
				},
			},
			wantErr: "invalid responsePolicy",
		},
		{
			name: "valid spec with external responsePolicy path",
			obj: &DNSResponsePolicyRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-rule",
					Namespace: "my-namespace",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: DNSResponsePolicyRuleSpec{
					ResponsePolicy: "projects/my-project/locations/global/responsePolicies/my-policy",
				},
			},
			want: &DNSResponsePolicyRuleIdentity{
				Project:        "my-project",
				Location:       "global",
				ResponsePolicy: "my-policy",
				Rule:           "my-rule",
			},
		},
		{
			name: "invalid spec with malformed external responsePolicy path (contains slashes)",
			obj: &DNSResponsePolicyRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-rule",
					Namespace: "my-namespace",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: DNSResponsePolicyRuleSpec{
					ResponsePolicy: "projects/my-project/invalid-format",
				},
			},
			wantErr: "invalid responsePolicy",
		},
		{
			name: "invalid spec with project mismatch in external responsePolicy",
			obj: &DNSResponsePolicyRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-rule",
					Namespace: "my-namespace",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: DNSResponsePolicyRuleSpec{
					ResponsePolicy: "projects/other-project/locations/global/responsePolicies/my-policy",
				},
			},
			wantErr: "responsePolicy project \"other-project\" must match project \"my-project\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIdentityFromDNSResponsePolicyRuleSpec(ctx, &mockReader{}, tt.obj)
			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("getIdentityFromDNSResponsePolicyRuleSpec() expected error, got nil identity")
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("getIdentityFromDNSResponsePolicyRuleSpec() error = %v, want error containing %q", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Fatalf("getIdentityFromDNSResponsePolicyRuleSpec() unexpected error: %v", err)
				}
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("getIdentityFromDNSResponsePolicyRuleSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
