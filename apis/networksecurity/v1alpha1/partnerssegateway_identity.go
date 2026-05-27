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
	_ identity.IdentityV2 = &NetworkSecurityPartnerSSEGatewayIdentity{}
	_ identity.Resource   = &NetworkSecurityPartnerSSEGateway{}
)

var NetworkSecurityPartnerSSEGatewayIdentityFormat = gcpurls.Template[NetworkSecurityPartnerSSEGatewayIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/partnerSSEGateways/{partner_sse_gateway}")

// +k8s:deepcopy-gen=false
type NetworkSecurityPartnerSSEGatewayIdentity struct {
	Project           string
	Location          string
	PartnerSSEGateway string
}

func (i *NetworkSecurityPartnerSSEGatewayIdentity) String() string {
	return NetworkSecurityPartnerSSEGatewayIdentityFormat.ToString(*i)
}

func (i *NetworkSecurityPartnerSSEGatewayIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkSecurityPartnerSSEGatewayIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkSecurityPartnerSSEGateway external=%q was not known (use %s): %w", ref, NetworkSecurityPartnerSSEGatewayIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecurityPartnerSSEGateway external=%q was not known (use %s)", ref, NetworkSecurityPartnerSSEGatewayIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkSecurityPartnerSSEGatewayIdentity) Host() string {
	return NetworkSecurityPartnerSSEGatewayIdentityFormat.Host()
}

func getIdentityFromNetworkSecurityPartnerSSEGatewaySpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecurityPartnerSSEGatewayIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &NetworkSecurityPartnerSSEGatewayIdentity{
		Project:           projectID,
		Location:          location,
		PartnerSSEGateway: resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityPartnerSSEGateway) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityPartnerSSEGatewaySpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NetworkSecurityPartnerSSEGatewayIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkSecurityPartnerSSEGateway identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
