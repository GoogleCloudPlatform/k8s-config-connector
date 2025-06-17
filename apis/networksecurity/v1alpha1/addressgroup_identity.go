// Copyright 2025 Google LLC
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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AddressGroupIdentity defines the resource reference to NetworkSecurityAddressGroup, which "External" field
// holds the GCP identifier for the KRM object.
type AddressGroupIdentity struct {
	parent *AddressGroupParent
	id     string
}

func (i *AddressGroupIdentity) String() string {
	return i.parent.String() + "/addressGroups/" + i.id
}

func (i *AddressGroupIdentity) ID() string {
	return i.id
}

func (i *AddressGroupIdentity) Parent() *AddressGroupParent {
	return i.parent
}

type AddressGroupParent struct {
	OrganizationID string
	ProjectID      string
	Location       string
}

func (p *AddressGroupParent) String() string {
	if p.OrganizationID != "" {
		return "organizations/" + p.OrganizationID + "/locations/" + p.Location
	} else if p.ProjectID != "" {
		return "projects/" + p.ProjectID + "/locations/" + p.Location
	}
	return ""
}

// NewAddressGroupIdentity builds a AddressGroupIdentity from the Config Connector AddressGroup object.
func NewAddressGroupIdentity(ctx context.Context, reader client.Reader, obj *NetworkSecurityAddressGroup) (*AddressGroupIdentity, error) {
	if obj.Spec.OrganizationRef != nil && obj.Spec.ProjectRef != nil {
		return nil, fmt.Errorf("organization and project cannot be defined at the same time")
	}
	if obj.Spec.OrganizationRef == nil && obj.Spec.ProjectRef == nil {
		return nil, fmt.Errorf("one of organization and project must be defined")
	}
	// Get Parent
	var parentID string
	if obj.Spec.OrganizationRef != nil {
		organization, err := refsv1beta1.ResolveOrganization(ctx, reader, nil, obj.Spec.OrganizationRef)
		if err != nil {
			return nil, err
		}
		parentID = organization.OrganizationID
	} else {
		project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
		if err != nil {
			return nil, err
		}
		parentID = project.ProjectID
	}

	if obj.Spec.Location == "" {
		return nil, fmt.Errorf("location must be defined")
	}
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
		actualParent, actualResourceID, err := ParseAddressGroupExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.OrganizationID != "" {
			if actualParent.OrganizationID != parentID {
				return nil, fmt.Errorf("spec.organizationID changed, expect %s, got %s", actualParent.OrganizationID, parentID)
			}
		} else if actualParent.ProjectID != "" {
			if actualParent.ProjectID != parentID {
				return nil, fmt.Errorf("spec.projectID changed, expect %s, got %s", actualParent.ProjectID, parentID)
			}
		}

		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	if obj.Spec.OrganizationRef != nil {
		return &AddressGroupIdentity{
			parent: &AddressGroupParent{
				OrganizationID: parentID,
				Location:       location,
			},
			id: resourceID,
		}, nil
	}
	return &AddressGroupIdentity{
		parent: &AddressGroupParent{
			ProjectID: parentID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseAddressGroupExternal(external string) (parent *AddressGroupParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) == 6 && tokens[0] == "organizations" && tokens[2] == "locations" && tokens[4] == "addressGroups" {
		parent = &AddressGroupParent{
			OrganizationID: tokens[1],
			Location:       tokens[3],
		}
		resourceID = tokens[5]
		return parent, resourceID, nil
	} else if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "addressGroups" {
		parent = &AddressGroupParent{
			ProjectID: tokens[1],
			Location:  tokens[3],
		}
		resourceID = tokens[5]
		return parent, resourceID, nil
	}
	return nil, "", fmt.Errorf("format of NetworkSecurityAddressGroup external=%q was not known (use projects/{{projectID}}/locations/{{location}}/addressGroups/{{addressgroupID}} or organizations/{{organizationsID}}/locations/{{location}}/addressGroups/{{addressgroupID}})", external)

}
