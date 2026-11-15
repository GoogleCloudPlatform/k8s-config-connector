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

package cloudsupport

import (
	"context"
	"net/http"
	"testing"

	supportpb "cloud.google.com/go/support/apiv2/supportpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudsupport/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestCloudSupportSupportEventSubscription_AdapterForURL(t *testing.T) {
	ctx := context.Background()
	m := &modelSupportEventSubscription{
		config: config.ControllerConfig{
			HTTPClient: &http.Client{},
		},
	}

	t.Run("valid URL with host prefix", func(t *testing.T) {
		url := "//cloudsupport.googleapis.com/organizations/123456789/supportEventSubscriptions/sub-123"
		adapter, err := m.AdapterForURL(ctx, url)
		if err != nil {
			t.Fatalf("AdapterForURL() unexpected error: %v", err)
		}
		if adapter == nil {
			t.Fatalf("AdapterForURL() returned nil adapter")
		}

		subAdapter, ok := adapter.(*SupportEventSubscriptionAdapter)
		if !ok {
			t.Fatalf("expected *SupportEventSubscriptionAdapter, got %T", adapter)
		}
		if subAdapter.id.Organization != "123456789" || subAdapter.id.SupportEventSubscription != "sub-123" {
			t.Errorf("unexpected identity parsed: %+v", subAdapter.id)
		}
	})

	t.Run("URL without host prefix returns nil without error", func(t *testing.T) {
		url := "organizations/123456789/supportEventSubscriptions/sub-123"
		adapter, err := m.AdapterForURL(ctx, url)
		if err != nil {
			t.Fatalf("AdapterForURL() unexpected error: %v", err)
		}
		if adapter != nil {
			t.Errorf("expected nil adapter for url without host prefix, got %v", adapter)
		}
	})

	t.Run("invalid URL returns nil without error", func(t *testing.T) {
		url := "https://compute.googleapis.com/compute/v1/projects/p/zones/z/instances/i"
		adapter, err := m.AdapterForURL(ctx, url)
		if err != nil {
			t.Fatalf("AdapterForURL() unexpected error: %v", err)
		}
		if adapter != nil {
			t.Errorf("expected nil adapter for invalid url, got %v", adapter)
		}
	})
}

func TestCloudSupportSupportEventSubscription_Export(t *testing.T) {
	ctx := context.Background()
	adapter := &SupportEventSubscriptionAdapter{
		id: &krm.CloudSupportSupportEventSubscriptionIdentity{
			Organization:             "123456789",
			SupportEventSubscription: "sub-123",
		},
		actual: &supportpb.SupportEventSubscription{
			Name:        "organizations/123456789/supportEventSubscriptions/sub-123",
			PubSubTopic: "projects/test-project/topics/test-topic",
		},
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export() returned error: %v", err)
	}

	if u.GetName() != "sub-123" {
		t.Errorf("Export() name = %q, want %q", u.GetName(), "sub-123")
	}
	if u.GroupVersionKind() != krm.CloudSupportSupportEventSubscriptionGVK {
		t.Errorf("Export() GVK = %v, want %v", u.GroupVersionKind(), krm.CloudSupportSupportEventSubscriptionGVK)
	}

	var obj krm.CloudSupportSupportEventSubscription
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		t.Fatalf("failed to convert unstructured to CloudSupportSupportEventSubscription: %v", err)
	}

	if obj.Spec.OrganizationRef == nil || obj.Spec.OrganizationRef.External != "organizations/123456789" {
		t.Errorf("Export() OrganizationRef = %+v, want external organizations/123456789", obj.Spec.OrganizationRef)
	}
	if obj.Spec.PubSubTopicRef == nil || obj.Spec.PubSubTopicRef.External != "projects/test-project/topics/test-topic" {
		t.Errorf("Export() PubSubTopicRef = %+v, want external projects/test-project/topics/test-topic", obj.Spec.PubSubTopicRef)
	}

	// Verify Find() not called error
	uninitializedAdapter := &SupportEventSubscriptionAdapter{}
	if _, err := uninitializedAdapter.Export(ctx); err == nil {
		t.Errorf("expected error when actual is nil")
	}
}
