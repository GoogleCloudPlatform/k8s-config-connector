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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &DialogflowWebhookIdentity{}
	_ identity.Resource   = &DialogflowWebhook{}
)

var DialogflowWebhookIdentityFormat = gcpurls.Template[DialogflowWebhookIdentity]("dialogflow.googleapis.com", "projects/{project}/locations/{location}/agents/{agent}/webhooks/{webhook}")

// DialogflowWebhookIdentity is the identity of a GCP DialogflowWebhook resource.
// +k8s:deepcopy-gen=false
type DialogflowWebhookIdentity struct {
	Project  string
	Location string
	Agent    string
	Webhook  string
}

func (i *DialogflowWebhookIdentity) String() string {
	return DialogflowWebhookIdentityFormat.ToString(*i)
}

func (i *DialogflowWebhookIdentity) Host() string {
	return DialogflowWebhookIdentityFormat.Host()
}

func (i *DialogflowWebhookIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s", i.Project, i.Location, i.Agent)
}

func (i *DialogflowWebhookIdentity) FromExternal(ref string) error {
	parsed, match, err := DialogflowWebhookIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DialogflowWebhook external=%q was not known (use %s): %w", ref, DialogflowWebhookIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DialogflowWebhook external=%q was not known (use %s)", ref, DialogflowWebhookIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func getIdentityFromDialogflowWebhookSpec(ctx context.Context, reader client.Reader, obj *DialogflowWebhook) (*DialogflowWebhookIdentity, error) {
	if obj.Spec.AgentRef == nil {
		return nil, fmt.Errorf("spec.agentRef is required")
	}

	if err := obj.Spec.AgentRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.agentRef: %w", err)
	}

	parentIdentity := &DialogflowCXAgentIdentity{}
	if err := parentIdentity.FromExternal(obj.Spec.AgentRef.GetExternal()); err != nil {
		return nil, fmt.Errorf("parsing agentRef.external=%q: %w", obj.Spec.AgentRef.GetExternal(), err)
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	return &DialogflowWebhookIdentity{
		Project:  parentIdentity.Project,
		Location: parentIdentity.Location,
		Agent:    parentIdentity.Agent,
		Webhook:  resourceID,
	}, nil
}

func (obj *DialogflowWebhook) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDialogflowWebhookSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &DialogflowWebhookIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DialogflowWebhook identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
