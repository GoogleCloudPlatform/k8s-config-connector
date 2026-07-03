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

var (
	ProjectNetworkSecurityFirewallEndpointIdentityFormat      = gcpurls.Template[NetworkSecurityFirewallEndpointIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/firewallEndpoints/{firewallendpoint}")
	OrganizationNetworkSecurityFirewallEndpointIdentityFormat = gcpurls.Template[NetworkSecurityFirewallEndpointIdentity]("networksecurity.googleapis.com", "organizations/{organization}/locations/{location}/firewallEndpoints/{firewallendpoint}")
)

// NetworkSecurityFirewallEndpointIdentity is the identity of a GCP NetworkSecurityFirewallEndpoint resource.
// +k8s:deepcopy-gen=false
type NetworkSecurityFirewallEndpointIdentity struct {
	Project          string
	Organization     string
	Location         string
	Firewallendpoint string
}

func (i *NetworkSecurityFirewallEndpointIdentity) String() string {
	if i.Project != "" {
		return ProjectNetworkSecurityFirewallEndpointIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return OrganizationNetworkSecurityFirewallEndpointIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *NetworkSecurityFirewallEndpointIdentity) FromExternal(ref string) error {
	if parsed, match, err := ProjectNetworkSecurityFirewallEndpointIdentityFormat.Parse(ref); err == nil && match {
		*i = *parsed
		return nil
	}
	if parsed, match, err := OrganizationNetworkSecurityFirewallEndpointIdentityFormat.Parse(ref); err == nil && match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of NetworkSecurityFirewallEndpoint external=%q was not known (use %s or %s)", ref, ProjectNetworkSecurityFirewallEndpointIdentityFormat.CanonicalForm(), OrganizationNetworkSecurityFirewallEndpointIdentityFormat.CanonicalForm())
}

func (i *NetworkSecurityFirewallEndpointIdentity) Host() string {
	return ProjectNetworkSecurityFirewallEndpointIdentityFormat.Host()
}

func (i *NetworkSecurityFirewallEndpointIdentity) ParentString() string {
	if i.Project != "" {
		return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
	}
	if i.Organization != "" {
		return fmt.Sprintf("organizations/%s/locations/%s", i.Organization, i.Location)
	}
	return ""
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

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &NetworkSecurityFirewallEndpointIdentity{
		Project:          projectID,
		Location:         location,
		Firewallendpoint: resourceID,
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

// ExternalIdentifier returns the GCP external identifier (the GCP URL).
func (obj *NetworkSecurityFirewallEndpoint) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
