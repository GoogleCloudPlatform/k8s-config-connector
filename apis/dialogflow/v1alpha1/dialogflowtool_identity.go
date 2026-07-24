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
	_ identity.IdentityV2 = &DialogflowToolIdentity{}
	_ identity.Resource   = &DialogflowTool{}
)

var DialogflowToolIdentityFormat = gcpurls.Template[DialogflowToolIdentity]("dialogflow.googleapis.com", "projects/{project}/locations/{location}/agents/{agent}/tools/{tool}")

// +k8s:deepcopy-gen=false
// DialogflowToolIdentity is the identity of a GCP DialogflowTool resource.
type DialogflowToolIdentity struct {
	Project  string
	Location string
	Agent    string
	Tool     string
}

func (i *DialogflowToolIdentity) String() string {
	return DialogflowToolIdentityFormat.ToString(*i)
}

func (i *DialogflowToolIdentity) FromExternal(ref string) error {
	parsed, match, err := DialogflowToolIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DialogflowTool external=%q was not known (use %s): %w", ref, DialogflowToolIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DialogflowTool external=%q was not known (use %s)", ref, DialogflowToolIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DialogflowToolIdentity) Host() string {
	return DialogflowToolIdentityFormat.Host()
}

func getIdentityFromDialogflowToolSpec(ctx context.Context, reader client.Reader, obj *DialogflowTool) (*DialogflowToolIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	if obj.Spec.AgentRef == nil {
		return nil, fmt.Errorf("spec.agentRef is required")
	}
	agentRef := obj.Spec.AgentRef.DeepCopy()
	if err := agentRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("cannot normalize agentRef: %w", err)
	}
	agentIdentity, err := agentRef.ParseExternalToIdentity()
	if err != nil {
		return nil, fmt.Errorf("cannot parse agentRef: %w", err)
	}
	agentId, ok := agentIdentity.(*DialogflowCXAgentIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected agent identity type: %T", agentIdentity)
	}

	identity := &DialogflowToolIdentity{
		Project:  agentId.Project,
		Location: agentId.Location,
		Agent:    agentId.Agent,
		Tool:     resourceID,
	}
	return identity, nil
}

func (obj *DialogflowTool) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromDialogflowToolSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DialogflowToolIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change DialogflowTool identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
