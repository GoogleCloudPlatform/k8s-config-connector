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
	"math"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

func TestAppEngineFirewallRuleIdentity_SpecProjectPrecedence(t *testing.T) {
	ctx := context.Background()
	obj := &AppEngineFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name: "200",
			Annotations: map[string]string{
				"cnrm.cloud.google.com/project-id": "project-from-annotation",
			},
		},
		Spec: AppEngineFirewallRuleSpec{
			Project: ptr.To("project-from-spec"),
		},
	}

	id, err := getIdentityFromAppEngineFirewallRuleSpec(ctx, nil, obj)

	if err != nil {
		t.Fatalf("unexpected error resolving identity: %v", err)
	}
	if id.App != "project-from-spec" {
		t.Errorf("got app %q, want %q", id.App, "project-from-spec")
	}
	if id.IngressRule != "200" {
		t.Errorf("got ingressRule %q, want %q", id.IngressRule, "200")
	}
}

func TestAppEngineFirewallRuleIdentity_ResolveProjectIDFallback(t *testing.T) {
	ctx := context.Background()
	obj := &AppEngineFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name: "rule-cr",
			Annotations: map[string]string{
				"cnrm.cloud.google.com/project-id": "project-from-annotation",
			},
		},
		Spec: AppEngineFirewallRuleSpec{
			ResourceID: ptr.To("100"),
		},
	}

	id, err := getIdentityFromAppEngineFirewallRuleSpec(ctx, nil, obj)

	if err != nil {
		t.Fatalf("unexpected error resolving identity: %v", err)
	}
	if id.App != "project-from-annotation" {
		t.Errorf("got app %q, want %q", id.App, "project-from-annotation")
	}
	if id.IngressRule != "100" {
		t.Errorf("got ingressRule %q, want %q", id.IngressRule, "100")
	}
}

func TestAppEngineFirewallRuleIdentity_ParseAndCanonicalize(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name        string
		metaName    string
		resourceID  *string
		wantRuleID  string
		expectError bool
	}{
		{
			name:       "minimum valid priority 1",
			metaName:   "1",
			wantRuleID: "1",
		},
		{
			name:       "leading zeros canonicalized",
			metaName:   "rule-cr",
			resourceID: ptr.To("0100"),
			wantRuleID: "100",
		},
		{
			name:       "default rule priority 2147483647",
			metaName:   "default-rule",
			resourceID: ptr.To("2147483647"),
			wantRuleID: "2147483647",
		},
		{
			name:       "zero priority resolves in identity so existing CRs can be deleted",
			metaName:   "rule-zero",
			resourceID: ptr.To("0"),
			wantRuleID: "0",
		},
		{
			name:       "negative priority resolves in identity so existing CRs can be deleted",
			metaName:   "rule-neg",
			resourceID: ptr.To("-10"),
			wantRuleID: "-10",
		},
		{
			name:        "overflow above MaxInt32 rejected",
			metaName:    "rule-overflow",
			resourceID:  ptr.To("2147483648"),
			expectError: true,
		},
		{
			name:        "non-integer metadata.name without resourceID rejected",
			metaName:    "allow-corp-cidr",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			obj := &AppEngineFirewallRule{
				ObjectMeta: metav1.ObjectMeta{
					Name: tc.metaName,
				},
				Spec: AppEngineFirewallRuleSpec{
					Project:    ptr.To("test-project"),
					ResourceID: tc.resourceID,
				},
			}

			id, err := getIdentityFromAppEngineFirewallRuleSpec(ctx, nil, obj)

			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error, got identity %+v", id)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if id.IngressRule != tc.wantRuleID {
				t.Errorf("got ingressRule %q, want %q", id.IngressRule, tc.wantRuleID)
			}
		})
	}
}

func TestAppEngineFirewallRuleIdentity_Priority(t *testing.T) {
	tests := []struct {
		name         string
		identity     *AppEngineFirewallRuleIdentity
		wantPriority int32
		expectError  bool
	}{
		{
			name:         "priority 1 valid",
			identity:     &AppEngineFirewallRuleIdentity{App: "p", IngressRule: "1"},
			wantPriority: 1,
		},
		{
			name:         "default rule priority valid",
			identity:     &AppEngineFirewallRuleIdentity{App: "p", IngressRule: "2147483647"},
			wantPriority: math.MaxInt32,
		},
		{
			name:        "priority 0 rejected for actuation",
			identity:    &AppEngineFirewallRuleIdentity{App: "p", IngressRule: "0"},
			expectError: true,
		},
		{
			name:        "negative priority rejected for actuation",
			identity:    &AppEngineFirewallRuleIdentity{App: "p", IngressRule: "-10"},
			expectError: true,
		},
		{
			name:        "non-integer rejected",
			identity:    &AppEngineFirewallRuleIdentity{App: "p", IngressRule: "allow-corp"},
			expectError: true,
		},
		{
			name:        "nil identity rejected",
			identity:    nil,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.identity.Priority()

			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error, got priority %d", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.wantPriority {
				t.Errorf("got priority %d, want %d", got, tc.wantPriority)
			}
		})
	}
}

func TestAppEngineFirewallRuleIdentity_ErrorWhenProjectUnset(t *testing.T) {
	ctx := context.Background()
	obj := &AppEngineFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name: "200",
		},
	}

	_, err := getIdentityFromAppEngineFirewallRuleSpec(ctx, nil, obj)

	if err == nil {
		t.Fatalf("expected error resolving identity when project is unset, got nil")
	}
}
