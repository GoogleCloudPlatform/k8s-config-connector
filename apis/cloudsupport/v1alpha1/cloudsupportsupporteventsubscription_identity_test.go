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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCloudSupportSupportEventSubscriptionIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *CloudSupportSupportEventSubscriptionIdentity
		hasError bool
	}{
		{
			name:  "Full resource name",
			input: "organizations/123456789/supportEventSubscriptions/my-subscription",
			expected: &CloudSupportSupportEventSubscriptionIdentity{
				Organization:             "123456789",
				SupportEventSubscription: "my-subscription",
			},
			hasError: false,
		},
		{
			name:  "Full resource name with host",
			input: "cloudsupport.googleapis.com/organizations/123456789/supportEventSubscriptions/my-subscription",
			expected: &CloudSupportSupportEventSubscriptionIdentity{
				Organization:             "123456789",
				SupportEventSubscription: "my-subscription",
			},
			hasError: false,
		},
		{
			name:     "Invalid format",
			input:    "organizations/123456789/invalid/my-subscription",
			expected: nil,
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &CloudSupportSupportEventSubscriptionIdentity{}
			err := id.FromExternal(tc.input)
			if tc.hasError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.expected, id); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCloudSupportSupportEventSubscriptionIdentity_StringAndParentString(t *testing.T) {
	id := &CloudSupportSupportEventSubscriptionIdentity{
		Organization:             "123456789",
		SupportEventSubscription: "my-subscription",
	}

	expectedStr := "organizations/123456789/supportEventSubscriptions/my-subscription"
	if got := id.String(); got != expectedStr {
		t.Errorf("String() = %q, want %q", got, expectedStr)
	}

	expectedParentStr := "organizations/123456789"
	if got := id.ParentString(); got != expectedParentStr {
		t.Errorf("ParentString() = %q, want %q", got, expectedParentStr)
	}

	// Verify FromExternal round-trip
	parsed := &CloudSupportSupportEventSubscriptionIdentity{}
	if err := parsed.FromExternal(expectedStr); err != nil {
		t.Fatalf("unexpected error parsing %q: %v", expectedStr, err)
	}
	if parsed.String() != expectedStr {
		t.Errorf("round-trip failed: got %q, want %q", parsed.String(), expectedStr)
	}
}

func TestCloudSupportSupportEventSubscription_GetIdentity(t *testing.T) {
	ctx := context.Background()

	// Test case 1: ResourceID is set explicitly
	customID := "my-custom-subscription"
	subscriptionObjWithCustomID := &CloudSupportSupportEventSubscription{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-subscription-k8s-name",
			Namespace: "default",
		},
		Spec: CloudSupportSupportEventSubscriptionSpec{
			OrganizationRef: &refsv1beta1.OrganizationRef{
				External: "organizations/123456789",
			},
			ResourceID: &customID,
		},
	}

	id, err := subscriptionObjWithCustomID.GetIdentity(ctx, nil)
	if err != nil {
		t.Fatalf("unexpected error getting identity: %v", err)
	}
	expectedID := "organizations/123456789/supportEventSubscriptions/my-custom-subscription"
	if id.String() != expectedID {
		t.Errorf("GetIdentity() = %q, want %q", id.String(), expectedID)
	}

	// Test case 2: ResourceID is not set (should default to k8s metadata.name)
	subscriptionObjWithDefaultID := &CloudSupportSupportEventSubscription{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-subscription-k8s-name",
			Namespace: "default",
		},
		Spec: CloudSupportSupportEventSubscriptionSpec{
			OrganizationRef: &refsv1beta1.OrganizationRef{
				External: "organizations/123456789",
			},
		},
	}

	id, err = subscriptionObjWithDefaultID.GetIdentity(ctx, nil)
	if err != nil {
		t.Fatalf("unexpected error getting identity: %v", err)
	}
	expectedDefaultID := "organizations/123456789/supportEventSubscriptions/my-subscription-k8s-name"
	if id.String() != expectedDefaultID {
		t.Errorf("GetIdentity() = %q, want %q", id.String(), expectedDefaultID)
	}
}
