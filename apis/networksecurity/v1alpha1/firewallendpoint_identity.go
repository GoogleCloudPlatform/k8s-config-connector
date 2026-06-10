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
	_ identity.IdentityV2 = &NetworkSecurityFirewallEndpointIdentity{}
	_ identity.Resource   = &NetworkSecurityFirewallEndpoint{}
)

var NetworkSecurityFirewallEndpointIdentityFormatProject = gcpurls.Template[NetworkSecurityFirewallEndpointIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/firewallEndpoints/{firewallendpoint}")
var NetworkSecurityFirewallEndpointIdentityFormatOrganization = gcpurls.Template[NetworkSecurityFirewallEndpointIdentity]("networksecurity.googleapis.com", "organizations/{organization}/locations/{location}/firewallEndpoints/{firewallendpoint}")

// +k8s:deepcopy-gen=false
type NetworkSecurityFirewallEndpointIdentity struct {
	Project          string
	Organization     string
	Location         string
	Firewallendpoint string
}

func (i *NetworkSecurityFirewallEndpointIdentity) String() string {
	if i.Organization != "" {
		return NetworkSecurityFirewallEndpointIdentityFormatOrganization.ToString(*i)
	}
	return NetworkSecurityFirewallEndpointIdentityFormatProject.ToString(*i)
}

func (i *NetworkSecurityFirewallEndpointIdentity) FromExternal(ref string) error {
	if parsed, match, _ := NetworkSecurityFirewallEndpointIdentityFormatOrganization.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := NetworkSecurityFirewallEndpointIdentityFormatProject.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of NetworkSecurityFirewallEndpoint external=%q was not known (use %s or %s)", ref, NetworkSecurityFirewallEndpointIdentityFormatProject.CanonicalForm(), NetworkSecurityFirewallEndpointIdentityFormatOrganization.CanonicalForm())
}

func (i *NetworkSecurityFirewallEndpointIdentity) Host() string {
	return NetworkSecurityFirewallEndpointIdentityFormatProject.Host()
}

func getIdentityFromNetworkSecurityFirewallEndpointSpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecurityFirewallEndpointIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	identity := &NetworkSecurityFirewallEndpointIdentity{
		Location:         location,
		Firewallendpoint: resourceID,
	}

	// The Endpoint can be scoped to either a project or an organization.
	// We attempt to resolve the project first, and if that fails, try the organization.
	if p, ok := obj.(*NetworkSecurityFirewallEndpoint); ok && p.Spec.OrganizationRef != nil {
		orgIdentity, err := refs.ResolveOrganization(ctx, reader, obj, p.Spec.OrganizationRef)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve organization: %w", err)
		}
		identity.Organization = orgIdentity.OrganizationID
	} else {
		projectID, err := refs.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		identity.Project = projectID
	}

	return identity, nil
}

func (obj *NetworkSecurityFirewallEndpoint) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityFirewallEndpointSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NetworkSecurityFirewallEndpointIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkSecurityFirewallEndpoint identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

func (obj *NetworkSecurityFirewallEndpoint) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
