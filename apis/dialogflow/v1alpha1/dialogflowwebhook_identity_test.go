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

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDialogflowWebhookIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DialogflowWebhookIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/agents/my-agent/webhooks/my-webhook",
			want: &DialogflowWebhookIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Webhook:  "my-webhook",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/locations/us-central1/agents/my-agent/webhooks/my-webhook",
			want: &DialogflowWebhookIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Webhook:  "my-webhook",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DialogflowWebhookIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got, want := i.String(), "projects/my-project/locations/us-central1/agents/my-agent/webhooks/my-webhook"; got != want {
					t.Errorf("String() = %v, want %v", got, want)
				}
				if got, want := i.ParentString(), "projects/my-project/locations/us-central1/agents/my-agent"; got != want {
					t.Errorf("ParentString() = %v, want %v", got, want)
				}
			}
		})
	}
}

func TestDialogflowWebhook_GetIdentity(t *testing.T) {
	resourceID := "my-webhook"
	obj := &DialogflowWebhook{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-webhook",
			Namespace: "default",
		},
		Spec: DialogflowWebhookSpec{
			AgentRef: &DialogflowCXAgentRef{
				External: "projects/my-project/locations/us-central1/agents/my-agent",
			},
			ResourceID: &resourceID,
		},
	}

	fakeClient := fake.NewFakeClient()
	id, err := obj.GetIdentity(context.Background(), fakeClient)
	if err != nil {
		t.Fatalf("GetIdentity() error = %v", err)
	}

	webhookID, ok := id.(*DialogflowWebhookIdentity)
	if !ok {
		t.Fatalf("unexpected identity type: %T", id)
	}

	want := &DialogflowWebhookIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Agent:    "my-agent",
		Webhook:  "my-webhook",
	}

	if diff := cmp.Diff(want, webhookID); diff != "" {
		t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
	}
}
