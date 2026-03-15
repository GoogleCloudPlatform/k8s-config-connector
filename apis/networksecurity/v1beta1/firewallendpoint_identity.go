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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// FirewallEndpointIdentity defines the resource reference to NetworkSecurityFirewallEndpoint, which "External" field
// holds the GCP identifier for the KRM object.
type FirewallEndpointIdentity struct {
	parent *FirewallEndpointParent
	id     string
}

func (i *FirewallEndpointIdentity) String() string {
	return i.parent.String() + "/firewallEndpoints/" + i.id
}

func (i *FirewallEndpointIdentity) ID() string {
	return i.id
}

func (i *FirewallEndpointIdentity) Parent() *FirewallEndpointParent {
	return i.parent
}

type FirewallEndpointParent struct {
	OrganizationID string
	Location       string
}

func (p *FirewallEndpointParent) String() string {
	return "organizations/" + p.OrganizationID + "/locations/" + p.Location
}

// NewFirewallEndpointIdentity builds a FirewallEndpointIdentity from the Config Connector FirewallEndpoint object.
func NewFirewallEndpointIdentity(ctx context.Context, reader client.Reader, obj *NetworkSecurityFirewallEndpoint) (*FirewallEndpointIdentity, error) {

	// Get Parent
	orgID, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
	if err != nil {
		return nil, err
	}
	if orgID == nil {
		return nil, fmt.Errorf("organizationRef is required")
	}
	organizationID := orgID.OrganizationID
	location := obj.Spec.Location

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseFirewallEndpointExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.OrganizationID != organizationID {
			return nil, fmt.Errorf("spec.organizationRef changed, expect %s, got %s", actualParent.OrganizationID, organizationID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &FirewallEndpointIdentity{
		parent: &FirewallEndpointParent{
			OrganizationID: organizationID,
			Location:       location,
		},
		id: resourceID,
	}, nil
}

func ParseFirewallEndpointExternal(external string) (parent *FirewallEndpointParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "organizations" || tokens[2] != "locations" || tokens[4] != "firewallEndpoints" {
		return nil, "", fmt.Errorf("format of NetworkSecurityFirewallEndpoint external=%q was not known (use organizations/{{organizationID}}/locations/{{location}}/firewallEndpoints/{{firewallEndpointID}})", external)
	}
	parent = &FirewallEndpointParent{
		OrganizationID: tokens[1],
		Location:       tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
