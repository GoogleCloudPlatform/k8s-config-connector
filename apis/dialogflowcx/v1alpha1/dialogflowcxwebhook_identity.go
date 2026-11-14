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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.Resource = &DialogflowCXWebhook{}
)

func getIdentityFromDialogflowCXWebhookSpec(ctx context.Context, reader client.Reader, obj *DialogflowCXWebhook) (*DialogflowCXWebhookIdentity, error) {
	if obj.Spec.Parent == nil || *obj.Spec.Parent == "" {
		return nil, fmt.Errorf("spec.parent is required")
	}

	parent, match, err := agentParentFormat.Parse(*obj.Spec.Parent)
	if err != nil {
		return nil, fmt.Errorf("parsing spec.parent %q: %w", *obj.Spec.Parent, err)
	}
	if !match {
		return nil, fmt.Errorf("format of DialogflowCXWebhook spec.parent=%q was not known (use %s)", *obj.Spec.Parent, agentParentFormat.CanonicalForm())
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)

	return &DialogflowCXWebhookIdentity{
		Project:  parent.Project,
		Location: parent.Location,
		Agent:    parent.Agent,
		Webhook:  resourceID,
	}, nil
}

func (obj *DialogflowCXWebhook) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDialogflowCXWebhookSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		externalRef = common.ValueOf(obj.Status.Name)
	}
	if externalRef != "" {
		// To preserve backward compatibility with Terraform, status.Name can be just the webhook ID
		// (e.g. "4150c119-76bc-4592-b12d-8d65e0b22606") or the full path. If it doesn't contain a "/", it's
		// the short ID; we can verify it directly against specIdentity.Webhook.
		if !strings.Contains(externalRef, "/") {
			if specIdentity.Webhook == "" {
				specIdentity.Webhook = externalRef
			} else if specIdentity.Webhook != externalRef {
				return nil, fmt.Errorf("cannot change DialogflowCXWebhook identity (old=%q, new=%q)", externalRef, specIdentity.Webhook)
			}
		} else {
			statusIdentity := &DialogflowCXWebhookIdentity{}
			if err := statusIdentity.FromExternal(externalRef); err != nil {
				return nil, err
			}

			if specIdentity.Webhook == "" {
				specIdentity.Webhook = statusIdentity.Webhook
			}

			if statusIdentity.String() != specIdentity.String() {
				return nil, fmt.Errorf("cannot change DialogflowCXWebhook identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
			}
		}
	}

	return specIdentity, nil
}
