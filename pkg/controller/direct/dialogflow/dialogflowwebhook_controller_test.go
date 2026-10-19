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

package dialogflow

import (
	"context"
	"testing"

	cxpb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestDialogflowWebhook_AdapterForURL(t *testing.T) {
	ctx := context.Background()
	model := &webhookModel{
		config: config.ControllerConfig{},
	}

	// Test invalid URL returns nil, nil
	adapter, err := model.AdapterForURL(ctx, "invalid-url")
	if err != nil {
		t.Fatalf("unexpected error for invalid url: %v", err)
	}
	if adapter != nil {
		t.Fatalf("expected nil adapter for invalid url, got: %v", adapter)
	}
}

func TestDialogflowWebhook_Export(t *testing.T) {
	ctx := context.Background()
	adapter := &webhookAdapter{
		id: &krm.DialogflowWebhookIdentity{
			Project:  "test-project",
			Location: "global",
			Agent:    "test-agent",
			Webhook:  "test-webhook",
		},
		actual: &cxpb.Webhook{
			Name:        "projects/test-project/locations/global/agents/test-agent/webhooks/test-webhook",
			DisplayName: "Test Webhook",
			Webhook: &cxpb.Webhook_GenericWebService_{
				GenericWebService: &cxpb.Webhook_GenericWebService{
					Uri: "https://example.com/webhook",
				},
			},
		},
	}

	u, err := adapter.Export(ctx)
	if err != nil {
		t.Fatalf("Export() returned error: %v", err)
	}

	if u.GetName() != "test-webhook" {
		t.Errorf("Export() name = %q, want %q", u.GetName(), "test-webhook")
	}
	if u.GroupVersionKind() != krm.DialogflowWebhookGVK {
		t.Errorf("Export() GVK = %v, want %v", u.GroupVersionKind(), krm.DialogflowWebhookGVK)
	}

	var obj krm.DialogflowWebhook
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		t.Fatalf("failed to convert unstructured to DialogflowWebhook: %v", err)
	}

	if obj.Spec.AgentRef == nil || obj.Spec.AgentRef.External != "projects/test-project/locations/global/agents/test-agent" {
		t.Errorf("Export() AgentRef.External = %v, want projects/test-project/locations/global/agents/test-agent", obj.Spec.AgentRef)
	}
	if obj.Spec.ResourceID == nil || *obj.Spec.ResourceID != "test-webhook" {
		t.Errorf("Export() ResourceID = %v, want test-webhook", obj.Spec.ResourceID)
	}
	if obj.Spec.DisplayName == nil || *obj.Spec.DisplayName != "Test Webhook" {
		t.Errorf("Export() DisplayName = %v, want 'Test Webhook'", obj.Spec.DisplayName)
	}
	if obj.Spec.GenericWebService == nil || obj.Spec.GenericWebService.URI == nil || *obj.Spec.GenericWebService.URI != "https://example.com/webhook" {
		t.Errorf("Export() GenericWebService.URI = %v, want 'https://example.com/webhook'", obj.Spec.GenericWebService)
	}

	// Verify Find() not called error
	uninitializedAdapter := &webhookAdapter{}
	if _, err := uninitializedAdapter.Export(ctx); err == nil {
		t.Errorf("expected error when actual is nil")
	}
}
