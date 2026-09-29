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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NetworkServicesAgentGatewayIdentity{}
	_ identity.Resource   = &NetworkServicesAgentGateway{}
)

var NetworkServicesAgentGatewayIdentityFormat = gcpurls.Template[NetworkServicesAgentGatewayIdentity]("networkservices.googleapis.com", "projects/{project}/locations/{location}/agentGateways/{agentGateway}")

// NetworkServicesAgentGatewayIdentity is the identity of a GCP NetworkServicesAgentGateway resource.
// +k8s:deepcopy-gen=false
type NetworkServicesAgentGatewayIdentity struct {
	Project      string
	Location     string
	AgentGateway string
}

func (i *NetworkServicesAgentGatewayIdentity) String() string {
	return NetworkServicesAgentGatewayIdentityFormat.ToString(*i)
}

func (i *NetworkServicesAgentGatewayIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkServicesAgentGatewayIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkServicesAgentGateway external=%q was not known (use %s): %w", ref, NetworkServicesAgentGatewayIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkServicesAgentGateway external=%q was not known (use %s)", ref, NetworkServicesAgentGatewayIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkServicesAgentGatewayIdentity) Host() string {
	return NetworkServicesAgentGatewayIdentityFormat.Host()
}

func (i *NetworkServicesAgentGatewayIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func NewNetworkServicesAgentGatewayIdentity(ctx context.Context, reader client.Reader, obj *NetworkServicesAgentGateway) (*NetworkServicesAgentGatewayIdentity, error) {
	resourceID, err := refsv1beta1.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location, err := refsv1beta1.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &NetworkServicesAgentGatewayIdentity{
		Project:      projectID,
		Location:     location,
		AgentGateway: resourceID,
	}
	return identity, nil
}

func (obj *NetworkServicesAgentGateway) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := NewNetworkServicesAgentGatewayIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	return specIdentity, nil
}
