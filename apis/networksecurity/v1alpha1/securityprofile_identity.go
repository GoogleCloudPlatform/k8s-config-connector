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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.IdentityV2 = &SecurityProfileIdentity{}

var securityProfileProjectURL = gcpurls.Template[SecurityProfileIdentity](
	"networksecurity.googleapis.com",
	"projects/{project}/locations/{location}/securityProfiles/{securityProfile}",
)

var securityProfileOrganizationURL = gcpurls.Template[SecurityProfileIdentity](
	"networksecurity.googleapis.com",
	"organizations/{organization}/locations/{location}/securityProfiles/{securityProfile}",
)

// SecurityProfileIdentity defines the resource reference to NetworkSecuritySecurityProfile, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type SecurityProfileIdentity struct {
	Project         string
	Organization    string
	Location        string
	SecurityProfile string
}

func (i *SecurityProfileIdentity) Parent() string {
	if i.Organization != "" {
		return "organizations/" + i.Organization + "/locations/" + i.Location
	}
	return "projects/" + i.Project + "/locations/" + i.Location
}

func (i *SecurityProfileIdentity) ID() string {
	return i.SecurityProfile
}

func (i *SecurityProfileIdentity) FromExternal(ref string) error {
	out, match, err := securityProfileProjectURL.Parse(ref)
	if err != nil {
		return err
	}
	if match {
		*i = *out
		return nil
	}

	out, match, err = securityProfileOrganizationURL.Parse(ref)
	if err != nil {
		return err
	}
	if match {
		*i = *out
		return nil
	}

	return fmt.Errorf("format of NetworkSecuritySecurityProfile external=%q was not known", ref)
}

func (i *SecurityProfileIdentity) String() string {
	if i.Organization != "" {
		return securityProfileOrganizationURL.ToString(*i)
	}
	return securityProfileProjectURL.ToString(*i)
}

func (i *SecurityProfileIdentity) Host() string {
	return securityProfileProjectURL.Host()
}

// GetIdentity builds a SecurityProfileIdentity from the Config Connector NetworkSecuritySecurityProfile object.
func (obj *NetworkSecuritySecurityProfile) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return NewSecurityProfileIdentity(ctx, reader, obj)
}

// NewSecurityProfileIdentity builds a SecurityProfileIdentity from the Config Connector NetworkSecuritySecurityProfile object.
func NewSecurityProfileIdentity(ctx context.Context, reader client.Reader, obj *NetworkSecuritySecurityProfile) (*SecurityProfileIdentity, error) {
	var projectID, organizationID string
	if obj.Spec.ProjectRef != nil {
		projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
		if err != nil {
			return nil, err
		}
		projectID = projectRef.ProjectID
		if projectID == "" {
			return nil, fmt.Errorf("cannot resolve project")
		}
	} else if obj.Spec.OrganizationRef != nil {
		organizationID = obj.Spec.OrganizationRef.External
		if organizationID == "" {
			return nil, fmt.Errorf("cannot resolve organization")
		}
	} else {
		// Default to project resolution
		projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), &refsv1beta1.ProjectRef{})
		if err != nil {
			return nil, err
		}
		projectID = projectRef.ProjectID
		if projectID == "" {
			return nil, fmt.Errorf("cannot resolve project")
		}
	}

	location := common.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

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
		actualIdentity := &SecurityProfileIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if projectID != "" && actualIdentity.Project != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.Project, projectID)
		}
		if organizationID != "" && actualIdentity.Organization != organizationID {
			return nil, fmt.Errorf("spec.organizationRef changed, expect %s, got %s", actualIdentity.Organization, organizationID)
		}
		if actualIdentity.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualIdentity.Location, location)
		}
		if actualIdentity.SecurityProfile != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.SecurityProfile)
		}
	}
	return &SecurityProfileIdentity{
		Project:         projectID,
		Organization:    organizationID,
		Location:        location,
		SecurityProfile: resourceID,
	}, nil
}

func (obj *NetworkSecuritySecurityProfile) ExternalIdentifier() *string {
	if obj.Status.ExternalRef != nil {
		return obj.Status.ExternalRef
	}
	return nil
}
