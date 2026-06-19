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
	_ identity.IdentityV2 = &NetworkSecurityInterceptEndpointGroupAssociationIdentity{}
	_ identity.Resource   = &NetworkSecurityInterceptEndpointGroupAssociation{}
)

var (
	ProjectInterceptEndpointGroupAssociationIdentityFormat      = gcpurls.Template[NetworkSecurityInterceptEndpointGroupAssociationIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/interceptEndpointGroupAssociations/{interceptEndpointGroupAssociation}")
	OrganizationInterceptEndpointGroupAssociationIdentityFormat = gcpurls.Template[NetworkSecurityInterceptEndpointGroupAssociationIdentity]("networksecurity.googleapis.com", "organizations/{organization}/locations/{location}/interceptEndpointGroupAssociations/{interceptEndpointGroupAssociation}")
)

// +k8s:deepcopy-gen=false
type NetworkSecurityInterceptEndpointGroupAssociationIdentity struct {
	Project                           string
	Organization                      string
	Location                          string
	InterceptEndpointGroupAssociation string
}

func (i *NetworkSecurityInterceptEndpointGroupAssociationIdentity) String() string {
	if i.Project != "" {
		return ProjectInterceptEndpointGroupAssociationIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return OrganizationInterceptEndpointGroupAssociationIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *NetworkSecurityInterceptEndpointGroupAssociationIdentity) FromExternal(ref string) error {
	if parsed, match, _ := ProjectInterceptEndpointGroupAssociationIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := OrganizationInterceptEndpointGroupAssociationIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of NetworkSecurityInterceptEndpointGroupAssociation external=%q was not known (use %s or %s)", ref, ProjectInterceptEndpointGroupAssociationIdentityFormat.CanonicalForm(), OrganizationInterceptEndpointGroupAssociationIdentityFormat.CanonicalForm())
}

func (i *NetworkSecurityInterceptEndpointGroupAssociationIdentity) Host() string {
	return "networksecurity.googleapis.com"
}

func getIdentityFromNetworkSecurityInterceptEndpointGroupAssociationSpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecurityInterceptEndpointGroupAssociationIdentity, error) {
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

	identity := &NetworkSecurityInterceptEndpointGroupAssociationIdentity{
		Project:                           projectID,
		Location:                          location,
		InterceptEndpointGroupAssociation: resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityInterceptEndpointGroupAssociation) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityInterceptEndpointGroupAssociationSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &NetworkSecurityInterceptEndpointGroupAssociationIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkSecurityInterceptEndpointGroupAssociation identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the GCP external identifier (the GCP URL).
func (obj *NetworkSecurityInterceptEndpointGroupAssociation) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
