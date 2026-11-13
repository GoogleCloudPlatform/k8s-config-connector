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

	vertexaiv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAIPlatformContext_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		obj         *AIPlatformContext
		expected    string
		expectError bool
	}{
		{
			name: "resolved identity with external parentRef and resourceID",
			obj: &AIPlatformContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-context",
					Namespace: "default",
				},
				Spec: AIPlatformContextSpec{
					ParentRef: &vertexaiv1beta1.MetadataStoreRef{
						External: "projects/my-project/locations/us-central1/metadataStores/my-store",
					},
					ResourceID: direct.LazyPtr("my-context"),
				},
			},
			expected: "projects/my-project/locations/us-central1/metadataStores/my-store/contexts/my-context",
		},
		{
			name: "resolved identity with external parentRef and fallback to metadata.name",
			obj: &AIPlatformContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-context",
					Namespace: "default",
				},
				Spec: AIPlatformContextSpec{
					ParentRef: &vertexaiv1beta1.MetadataStoreRef{
						External: "projects/my-project/locations/us-central1/metadataStores/my-store",
					},
				},
			},
			expected: "projects/my-project/locations/us-central1/metadataStores/my-store/contexts/k8s-context",
		},
		{
			name: "resolved identity matching status.externalRef",
			obj: &AIPlatformContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-context",
					Namespace: "default",
				},
				Spec: AIPlatformContextSpec{
					ParentRef: &vertexaiv1beta1.MetadataStoreRef{
						External: "projects/my-project/locations/us-central1/metadataStores/my-store",
					},
					ResourceID: direct.LazyPtr("my-context"),
				},
				Status: AIPlatformContextStatus{
					ExternalRef: direct.LazyPtr("projects/my-project/locations/us-central1/metadataStores/my-store/contexts/my-context"),
				},
			},
			expected: "projects/my-project/locations/us-central1/metadataStores/my-store/contexts/my-context",
		},
		{
			name: "error on status.externalRef mismatch",
			obj: &AIPlatformContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-context",
					Namespace: "default",
				},
				Spec: AIPlatformContextSpec{
					ParentRef: &vertexaiv1beta1.MetadataStoreRef{
						External: "projects/my-project/locations/us-central1/metadataStores/my-store",
					},
					ResourceID: direct.LazyPtr("my-context"),
				},
				Status: AIPlatformContextStatus{
					ExternalRef: direct.LazyPtr("projects/other-project/locations/us-central1/metadataStores/my-store/contexts/my-context"),
				},
			},
			expectError: true,
		},
		{
			name: "error when parentRef is nil",
			obj: &AIPlatformContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-context",
					Namespace: "default",
				},
				Spec: AIPlatformContextSpec{
					ResourceID: direct.LazyPtr("my-context"),
				},
			},
			expectError: true,
		},
		{
			name: "error when parentRef has invalid format",
			obj: &AIPlatformContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-context",
					Namespace: "default",
				},
				Spec: AIPlatformContextSpec{
					ParentRef: &vertexaiv1beta1.MetadataStoreRef{
						External: "invalid/parent/format",
					},
					ResourceID: direct.LazyPtr("my-context"),
				},
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.obj.GetIdentity(ctx, nil)
			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.expected, id.String()); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
