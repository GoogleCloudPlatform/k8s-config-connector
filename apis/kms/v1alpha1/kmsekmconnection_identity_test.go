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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestKMSEKMConnectionIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		obj         *KMSEKMConnection
		reader      client.Reader
		expected    *KMSEKMConnectionIdentity
		expectError bool
	}{
		{
			name: "basic",
			obj: &KMSEKMConnection{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-ekmconnection",
				},
				Spec: KMSEKMConnectionSpec{
					ProjectRef: &v1beta1.ProjectRef{
						External: "my-project",
					},
					Location: func(s string) *string { return &s }("us-central1"),
				},
			},
			expected: &KMSEKMConnectionIdentity{
				Project:       "my-project",
				Location:      "us-central1",
				EkmConnection: "my-ekmconnection",
			},
		},
		{
			name: "with resourceID",
			obj: &KMSEKMConnection{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-ekmconnection",
				},
				Spec: KMSEKMConnectionSpec{
					ProjectRef: &v1beta1.ProjectRef{
						External: "my-project",
					},
					Location:   func(s string) *string { return &s }("us-central1"),
					ResourceID: func(s string) *string { return &s }("my-actual-ekmconnection"),
				},
			},
			expected: &KMSEKMConnectionIdentity{
				Project:       "my-project",
				Location:      "us-central1",
				EkmConnection: "my-actual-ekmconnection",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.obj.GetIdentity(ctx, tc.reader)
			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			actualIdentity, ok := actual.(*KMSEKMConnectionIdentity)
			if !ok {
				t.Fatalf("expected KMSEKMConnectionIdentity, got %T", actual)
			}
			if *actualIdentity != *tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actualIdentity)
			}

			// Test string parsing
			s := actualIdentity.String()
			parsed := &KMSEKMConnectionIdentity{}
			if err := parsed.FromExternal(s); err != nil {
				t.Fatalf("FromExternal failed: %v", err)
			}
			if *parsed != *actualIdentity {
				t.Errorf("parsed %v != actual %v", parsed, actualIdentity)
			}

			var _ identity.IdentityV2 = parsed
		})
	}
}
