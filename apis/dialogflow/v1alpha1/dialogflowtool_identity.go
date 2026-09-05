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

// DialogflowToolIdentity is the identity of a GCP DialogflowTool resource.
// +k8s:deepcopy-gen=false
type DialogflowToolIdentity struct {
	Project  string
	Location string
	Agent    string
	Tool     string
}

func (i *DialogflowToolIdentity) String() string {
	return DialogflowToolIdentityFormat.ToString(*i)
}

func (i *DialogflowToolIdentity) Host() string {
	return DialogflowToolIdentityFormat.Host()
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

func (i *DialogflowToolIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/agents/%s", i.Project, i.Location, i.Agent)
}

func getIdentityFromDialogflowToolSpec(ctx context.Context, reader client.Reader, obj *DialogflowTool) (*DialogflowToolIdentity, error) {
	// 1. Resolve agent reference
	agentID, err := ResolveDialogflowAgent(ctx, reader, obj, obj.Spec.AgentRef)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve agentRef: %w", err)
	}
	if agentID == nil {
		return nil, fmt.Errorf("spec.agentRef is required and must be resolvable")
	}

	// 2. Resolve resource ID
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	identity := &DialogflowToolIdentity{
		Project:  agentID.Project,
		Location: agentID.Location,
		Agent:    agentID.Agent,
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
		if *specIdentity != *statusIdentity {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, specIdentity.String())
		}
	}

	return specIdentity, nil
}
