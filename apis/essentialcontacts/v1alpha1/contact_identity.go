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

// ContactIdentity defines the resource reference to EssentialContactsContact, which "External" field
// holds the GCP identifier for the KRM object.
type ContactIdentity struct {
	parent *ContactParent
	id     string
}

func (i *ContactIdentity) String() string {
	return i.parent.String() + "/contacts/" + i.id
}

func (i *ContactIdentity) ID() string {
	return i.id
}

func (i *ContactIdentity) Parent() *ContactParent {
	return i.parent
}

// ContactParent defines the parent for the contact.
type ContactParent struct {
	ProjectID      string
	OrganizationID string
	FolderID       string
	Location       string
}

func (p *ContactParent) String() string {
	if p.ProjectID != "" {
		return "projects/" + p.ProjectID + "/locations/" + p.Location
	} else if p.OrganizationID != "" {
		return "organizations/" + p.OrganizationID + "/locations/" + p.Location
	} else if p.FolderID != "" {
		return "folders/" + p.FolderID + "/locations/" + p.Location
	}
	return ""
}

// New builds a ContactIdentity from the Config Connector Contact object.
func NewContactIdentity(ctx context.Context, reader client.Reader, obj *EssentialContactsContact) (*ContactIdentity, error) {

	// Get Parent
	var projectID, organizationID, folderID string
	location := obj.Spec.Location

	if obj.Spec.ProjectRef != nil {
		projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), *obj.Spec.ProjectRef)
		if err != nil {
			return nil, err
		}
		projectID = projectRef.ProjectID
		if projectID == "" {
			return nil, fmt.Errorf("cannot resolve project")
		}
	} else if obj.Spec.OrganizationRef != nil {
		organization, err := refsv1beta1.GetOrganization(ctx, obj.Spec.OrganizationRef.Name, obj.Spec.OrganizationRef.Namespace, reader)
		if err != nil {
			return nil, fmt.Errorf("error getting organization: %w", err)
		}
		organizationID = organization
	} else if obj.Spec.FolderRef != nil {
		if obj.Spec.FolderRef.External != "" {
			folderID = obj.Spec.FolderRef.External
		} else {
			folder, err := refsv1beta1.GetFolderID(ctx, reader, obj.Spec.FolderRef.Name, obj.Namespace)
			if err != nil {
				return nil, err
			}
			folderID = folder
		}
		if folderID == "" {
			return nil, fmt.Errorf("cannot resolve folder")
		}
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
		actualParent, actualResourceID, err := ParseContactExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID && actualParent.OrganizationID != organizationID && actualParent.FolderID != folderID {
			if projectID != "" {
				return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
			} else if organizationID != "" {
				return nil, fmt.Errorf("spec.organizationRef changed, expect %s, got %s", actualParent.OrganizationID, organizationID)
			} else if folderID != "" {
				return nil, fmt.Errorf("spec.folderRef changed, expect %s, got %s", actualParent.FolderID, folderID)
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
	return &ContactIdentity{
		parent: &ContactParent{
			ProjectID:      projectID,
			OrganizationID: organizationID,
			FolderID:       folderID,
			Location:       location,
		},
		id: resourceID,
	}, nil
}

func ParseContactExternal(external string) (parent *ContactParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) == 6 && tokens[2] == "locations" && tokens[4] == "contacts" {
		switch tokens[0] {
		case "projects":
			parent = &ContactParent{
				ProjectID: tokens[1],
				Location:  tokens[3],
			}
			resourceID = tokens[5]
			return parent, resourceID, nil
		case "organizations":
			parent = &ContactParent{
				OrganizationID: tokens[1],
				Location:       tokens[3],
			}
			resourceID = tokens[5]
			return parent, resourceID, nil
		case "folders":
			parent = &ContactParent{
				FolderID: tokens[1],
				Location: tokens[3],
			}
			resourceID = tokens[5]
			return parent, resourceID, nil
		}
	}
	return nil, "", fmt.Errorf("format of EssentialContactsContact external=%q was not known (use projects/{{projectID}}/locations/{{location}}/contacts/{{contactID}} or organizations/{{organizationID}}/locations/{{location}}/contacts/{{contactID}} or folders/{{folderID}}/locations/{{location}}/contacts/{{contactID}})", external)
}
