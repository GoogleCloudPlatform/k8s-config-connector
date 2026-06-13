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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NetworkSecuritySecurityProfileGroupIdentity{}
	_ identity.Resource   = &NetworkSecuritySecurityProfileGroup{}
)

var NetworkSecuritySecurityProfileGroupProjectIdentityFormat = gcpurls.Template[NetworkSecuritySecurityProfileGroupIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/securityProfileGroups/{securityprofilegroup}")
var NetworkSecuritySecurityProfileGroupOrgIdentityFormat = gcpurls.Template[NetworkSecuritySecurityProfileGroupIdentity]("networksecurity.googleapis.com", "organizations/{organization}/locations/{location}/securityProfileGroups/{securityprofilegroup}")

// NetworkSecuritySecurityProfileGroupIdentity is the identity of a GCP NetworkSecuritySecurityProfileGroup resource.
// +k8s:deepcopy-gen=false
type NetworkSecuritySecurityProfileGroupIdentity struct {
	Project              string
	Organization         string
	Location             string
	SecurityProfileGroup string
}

func (i *NetworkSecuritySecurityProfileGroupIdentity) String() string {
	if i.Organization != "" {
		return NetworkSecuritySecurityProfileGroupOrgIdentityFormat.ToString(*i)
	}
	return NetworkSecuritySecurityProfileGroupProjectIdentityFormat.ToString(*i)
}

func (i *NetworkSecuritySecurityProfileGroupIdentity) FromExternal(ref string) error {
	if strings.Contains(ref, "organizations/") {
		parsed, match, err := NetworkSecuritySecurityProfileGroupOrgIdentityFormat.Parse(ref)
		if err != nil {
			return fmt.Errorf("format of NetworkSecuritySecurityProfileGroup external=%q was not known (use %s): %w", ref, NetworkSecuritySecurityProfileGroupOrgIdentityFormat.CanonicalForm(), err)
		}
		if !match {
			return fmt.Errorf("format of NetworkSecuritySecurityProfileGroup external=%q was not known (use %s)", ref, NetworkSecuritySecurityProfileGroupOrgIdentityFormat.CanonicalForm())
		}
		*i = *parsed
		return nil
	}

	parsed, match, err := NetworkSecuritySecurityProfileGroupProjectIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkSecuritySecurityProfileGroup external=%q was not known (use %s): %w", ref, NetworkSecuritySecurityProfileGroupProjectIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkSecuritySecurityProfileGroup external=%q was not known (use %s)", ref, NetworkSecuritySecurityProfileGroupProjectIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkSecuritySecurityProfileGroupIdentity) Host() string {
	return NetworkSecuritySecurityProfileGroupProjectIdentityFormat.Host()
}

func getIdentityFromNetworkSecuritySecurityProfileGroupSpec(ctx context.Context, reader client.Reader, obj *NetworkSecuritySecurityProfileGroup) (*NetworkSecuritySecurityProfileGroupIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	identity := &NetworkSecuritySecurityProfileGroupIdentity{
		Location:             location,
		SecurityProfileGroup: resourceID,
	}

	if obj.Spec.ProjectRef != nil {
		projectID, err := refs.ResolveProjectID(ctx, reader, obj)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve project: %w", err)
		}
		identity.Project = projectID
	} else if obj.Spec.OrganizationRef != nil {
		org, err := refs.ResolveOrganization(ctx, reader, obj, obj.Spec.OrganizationRef)
		if err != nil {
			return nil, fmt.Errorf("cannot resolve organization: %w", err)
		}
		identity.Organization = org.OrganizationID
	} else {
		return nil, fmt.Errorf("either projectRef or organizationRef must be set")
	}

	return identity, nil
}

func (obj *NetworkSecuritySecurityProfileGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecuritySecurityProfileGroupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NetworkSecuritySecurityProfileGroupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NetworkSecuritySecurityProfileGroup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}

// ExternalIdentifier returns the GCP external identifier (the GCP URL).
func (obj *NetworkSecuritySecurityProfileGroup) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
