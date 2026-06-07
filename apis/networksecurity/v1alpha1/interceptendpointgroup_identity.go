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
	_ identity.IdentityV2 = &NetworkSecurityInterceptEndpointGroupIdentity{}
	_ identity.Resource   = &NetworkSecurityInterceptEndpointGroup{}
)

var (
	ProjectInterceptEndpointGroupIdentityFormat      = gcpurls.Template[NetworkSecurityInterceptEndpointGroupIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/interceptEndpointGroups/{interceptendpointgroup}")
	OrganizationInterceptEndpointGroupIdentityFormat = gcpurls.Template[NetworkSecurityInterceptEndpointGroupIdentity]("networksecurity.googleapis.com", "organizations/{organization}/locations/{location}/interceptEndpointGroups/{interceptendpointgroup}")
)

// +k8s:deepcopy-gen=false
type NetworkSecurityInterceptEndpointGroupIdentity struct {
	Project                string
	Organization           string
	Location               string
	Interceptendpointgroup string
}

func (i *NetworkSecurityInterceptEndpointGroupIdentity) String() string {
	if i.Project != "" {
		return ProjectInterceptEndpointGroupIdentityFormat.ToString(*i)
	}
	if i.Organization != "" {
		return OrganizationInterceptEndpointGroupIdentityFormat.ToString(*i)
	}
	return ""
}

func (i *NetworkSecurityInterceptEndpointGroupIdentity) FromExternal(ref string) error {
	if parsed, match, _ := ProjectInterceptEndpointGroupIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := OrganizationInterceptEndpointGroupIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of NetworkSecurityInterceptEndpointGroup external=%q was not known (use %s or %s)", ref, ProjectInterceptEndpointGroupIdentityFormat.CanonicalForm(), OrganizationInterceptEndpointGroupIdentityFormat.CanonicalForm())
}

func (i *NetworkSecurityInterceptEndpointGroupIdentity) Host() string {
	return "networksecurity.googleapis.com"
}

func getIdentityFromNetworkSecurityInterceptEndpointGroupSpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecurityInterceptEndpointGroupIdentity, error) {
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

	identity := &NetworkSecurityInterceptEndpointGroupIdentity{
		Project:                projectID,
		Location:               location,
		Interceptendpointgroup: resourceID,
	}
	return identity, nil
}

func (obj *NetworkSecurityInterceptEndpointGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecurityInterceptEndpointGroupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NetworkSecurityInterceptEndpointGroupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkSecurityInterceptEndpointGroup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the GCP external identifier (the GCP URL).
func (obj *NetworkSecurityInterceptEndpointGroup) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
