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

package registration

import (
	"context"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestCheckResourceKindsDisabled(t *testing.T) {
	tests := []struct {
		name     string
		rk       *operatorv1beta1.ResourceKinds
		gvk      schema.GroupVersionKind
		expected bool
	}{
		{
			name:     "nil ResourceKinds",
			rk:       nil,
			gvk:      schema.GroupVersionKind{Kind: "StorageBucket"},
			expected: false,
		},
		{
			name: "Kind in denylist",
			rk: &operatorv1beta1.ResourceKinds{
				Denylist: []string{"StorageBucket"},
			},
			gvk:      schema.GroupVersionKind{Kind: "StorageBucket"},
			expected: true,
		},
		{
			name: "Kind in allowlist",
			rk: &operatorv1beta1.ResourceKinds{
				Allowlist: []string{"StorageBucket"},
			},
			gvk:      schema.GroupVersionKind{Kind: "StorageBucket"},
			expected: false,
		},
		{
			name: "Kind not in allowlist",
			rk: &operatorv1beta1.ResourceKinds{
				Allowlist: []string{"StorageBucket"},
			},
			gvk:      schema.GroupVersionKind{Kind: "PubSubTopic"},
			expected: true,
		},
		{
			name: "Kind in both allowlist and denylist",
			rk: &operatorv1beta1.ResourceKinds{
				Allowlist: []string{"StorageBucket"},
				Denylist:  []string{"StorageBucket"},
			},
			gvk:      schema.GroupVersionKind{Kind: "StorageBucket"},
			expected: true,
		},
		{
			name: "Empty allowlist and denylist",
			rk: &operatorv1beta1.ResourceKinds{
				Allowlist: []string{},
				Denylist:  []string{},
			},
			gvk:      schema.GroupVersionKind{Kind: "StorageBucket"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkResourceKindsDisabled(tt.rk, tt.gvk)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsResourceDisabledWithResourceKinds(t *testing.T) {
	gvk := schema.GroupVersionKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageBucket"}
	ctx := context.Background()

	t.Run("Disabled by CC ResourceKinds", func(t *testing.T) {
		ccRK := &operatorv1beta1.ResourceKinds{
			Denylist: []string{"StorageBucket"},
		}
		if !isResourceDisabled(ctx, gvk, "", nil, nil, nil, ccRK) {
			t.Errorf("expected StorageBucket to be disabled by CC ResourceKinds")
		}
	})

	t.Run("Disabled by CCC ResourceKinds", func(t *testing.T) {
		cccRK := &operatorv1beta1.ResourceKinds{
			Denylist: []string{"StorageBucket"},
		}
		if !isResourceDisabled(ctx, gvk, "test-ns", nil, nil, cccRK, nil) {
			t.Errorf("expected StorageBucket to be disabled by CCC ResourceKinds")
		}
	})

	t.Run("Allowed by CC ResourceKinds", func(t *testing.T) {
		ccRK := &operatorv1beta1.ResourceKinds{
			Allowlist: []string{"StorageBucket"},
		}
		if isResourceDisabled(ctx, gvk, "", nil, nil, nil, ccRK) {
			t.Errorf("expected StorageBucket to be allowed by CC ResourceKinds")
		}
	})

	t.Run("Denied by CCC even if allowed by CC", func(t *testing.T) {
		ccRK := &operatorv1beta1.ResourceKinds{
			Allowlist: []string{"StorageBucket"},
		}
		cccRK := &operatorv1beta1.ResourceKinds{
			Denylist: []string{"StorageBucket"},
		}
		if !isResourceDisabled(ctx, gvk, "test-ns", nil, nil, cccRK, ccRK) {
			t.Errorf("expected StorageBucket to be disabled by CCC Denylist even if in CC Allowlist")
		}
	})
}
