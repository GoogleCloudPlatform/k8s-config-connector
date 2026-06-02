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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NetworkSecuritySecurityProfileGroupIdentity{}
	_ identity.Resource   = &NetworkSecuritySecurityProfileGroup{}
)

var NetworkSecuritySecurityProfileGroupProjectIdentityFormat = gcpurls.Template[NetworkSecuritySecurityProfileGroupIdentity]("networksecurity.googleapis.com", "projects/{project}/locations/{location}/securityProfileGroups/{securityprofilegroup}")
var NetworkSecuritySecurityProfileGroupOrganizationIdentityFormat = gcpurls.Template[NetworkSecuritySecurityProfileGroupIdentity]("networksecurity.googleapis.com", "organizations/{organization}/locations/{location}/securityProfileGroups/{securityprofilegroup}")

// +k8s:deepcopy-gen=false
type NetworkSecuritySecurityProfileGroupIdentity struct {
	Project              string
	Organization         string
	Location             string
	SecurityProfileGroup string `gcpurls:"security_profile_group"`
}

func (i *NetworkSecuritySecurityProfileGroupIdentity) String() string {
	if i.Organization != "" {
		return NetworkSecuritySecurityProfileGroupOrganizationIdentityFormat.ToString(*i)
	}
	return NetworkSecuritySecurityProfileGroupProjectIdentityFormat.ToString(*i)
}

func (i *NetworkSecuritySecurityProfileGroupIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkSecuritySecurityProfileGroupProjectIdentityFormat.Parse(ref)
	if err != nil {
		return err
	}
	if match {
		*i = *parsed
		return nil
	}

	parsed, match, err = NetworkSecuritySecurityProfileGroupOrganizationIdentityFormat.Parse(ref)
	if err != nil {
		return err
	}
	if match {
		*i = *parsed
		return nil
	}

	return fmt.Errorf("format of NetworkSecuritySecurityProfileGroup external=%q was not known (use %s or %s)", ref, NetworkSecuritySecurityProfileGroupProjectIdentityFormat.CanonicalForm(), NetworkSecuritySecurityProfileGroupOrganizationIdentityFormat.CanonicalForm())
}

func (i *NetworkSecuritySecurityProfileGroupIdentity) Host() string {
	return NetworkSecuritySecurityProfileGroupProjectIdentityFormat.Host()
}

func (i *NetworkSecuritySecurityProfileGroupIdentity) ID() string {
	return i.SecurityProfileGroup
}

func (i *NetworkSecuritySecurityProfileGroupIdentity) ParentString() string {
	if i.Organization != "" {
		return fmt.Sprintf("organizations/%s/locations/%s", i.Organization, i.Location)
	}
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromNetworkSecuritySecurityProfileGroupSpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkSecuritySecurityProfileGroupIdentity, error) {
	var projectID, organizationID string
	var location string

	if typed, ok := obj.(*NetworkSecuritySecurityProfileGroup); ok {
		if typed.Spec.ProjectRef == nil && typed.Spec.OrganizationRef == nil {
			return nil, fmt.Errorf("one of projectRef or organizationRef must be set")
		}
		if typed.Spec.ProjectRef != nil && typed.Spec.OrganizationRef != nil {
			return nil, fmt.Errorf("only one of projectRef or organizationRef can be set")
		}

		if typed.Spec.ProjectRef != nil {
			projectRef, err := refs.ResolveProject(ctx, reader, typed.GetNamespace(), typed.Spec.ProjectRef)
			if err != nil {
				return nil, err
			}
			projectID = projectRef.ProjectID
			if projectID == "" {
				return nil, fmt.Errorf("cannot resolve project")
			}
		} else if typed.Spec.OrganizationRef != nil {
			organizationRef, err := refs.ResolveOrganization(ctx, reader, typed, typed.Spec.OrganizationRef)
			if err != nil {
				return nil, err
			}
			organizationID = organizationRef.OrganizationID
			if organizationID == "" {
				return nil, fmt.Errorf("cannot resolve organization")
			}
		}
		location = common.ValueOf(typed.Spec.Location)
	} else if u, ok := obj.(*unstructured.Unstructured); ok {
		hasProjectRef, _, _ := unstructured.NestedFieldNoCopy(u.Object, "spec", "projectRef")
		hasOrganizationRef, _, _ := unstructured.NestedFieldNoCopy(u.Object, "spec", "organizationRef")

		if hasProjectRef == nil && hasOrganizationRef == nil {
			return nil, fmt.Errorf("one of projectRef or organizationRef must be set")
		}
		if hasProjectRef != nil && hasOrganizationRef != nil {
			return nil, fmt.Errorf("only one of projectRef or organizationRef can be set")
		}

		if hasProjectRef != nil {
			var err error
			projectID, err = refs.ResolveProjectID(ctx, reader, u)
			if err != nil {
				return nil, err
			}
			if projectID == "" {
				return nil, fmt.Errorf("cannot resolve project")
			}
		} else if hasOrganizationRef != nil {
			var err error
			organizationID, err = refs.ResolveOrganizationID(ctx, reader, u)
			if err != nil {
				return nil, err
			}
			if organizationID == "" {
				return nil, fmt.Errorf("cannot resolve organization")
			}
		}
		location, _, _ = unstructured.NestedString(u.Object, "spec", "location")
	} else {
		return nil, fmt.Errorf("unexpected type %T", obj)
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	return &NetworkSecuritySecurityProfileGroupIdentity{
		Project:              projectID,
		Organization:         organizationID,
		Location:             location,
		SecurityProfileGroup: resourceID,
	}, nil
}

func (obj *NetworkSecuritySecurityProfileGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNetworkSecuritySecurityProfileGroupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
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

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (obj *NetworkSecuritySecurityProfileGroup) ExternalIdentifier() *string {
	return obj.Status.ExternalRef
}
