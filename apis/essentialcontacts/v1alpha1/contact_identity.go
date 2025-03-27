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
}

func (p *ContactParent) String() string {
	if p.ProjectID != "" {
		return "projects/" + p.ProjectID
	} else if p.OrganizationID != "" {
		return "organizations/" + p.OrganizationID
	} else if p.FolderID != "" {
		return "folders/" + p.FolderID
	}
	return ""
}

// New builds a ContactIdentity from the Config Connector Contact object.
func NewContactIdentity(ctx context.Context, reader client.Reader, obj *EssentialContactsContact) (*ContactIdentity, error) {

	// Get Parent
	var projectID, organizationID, folderID string

	if obj.Spec.ProjectRef != nil {
		projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.Parent.ProjectRef)
		if err != nil {
			return nil, err
		}
		projectID = projectRef.ProjectID
		if projectID == "" {
			return nil, fmt.Errorf("cannot resolve project")
		}
	} else if obj.Spec.OrganizationRef != nil {
		organization, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, obj.Spec.Parent.OrganizationRef)
		if err != nil {
			return nil, fmt.Errorf("error getting organization: %w", err)
		}
		organizationID = organization.OrganizationID
	} else if obj.Spec.FolderRef != nil {
		if obj.Spec.FolderRef.External != "" {
			folderID = obj.Spec.FolderRef.External
		} else {
			folder, err := refsv1beta1.ResolveFolder(ctx, reader, obj, obj.Spec.Parent.FolderRef)
			if err != nil {
				return nil, err
			}
			folderID = folder.FolderID
		}
		if folderID == "" {
			return nil, fmt.Errorf("cannot resolve folder")
		}
	} else {
		return nil, fmt.Errorf("cannot resolve parent for contact")
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
		},
		id: resourceID,
	}, nil
}

func ParseContactExternal(external string) (parent *ContactParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) == 4 && tokens[2] == "contacts" {
		switch tokens[0] {
		case "projects":
			parent = &ContactParent{
				ProjectID: tokens[1],
			}
			resourceID = tokens[3]
		case "organizations":
			parent = &ContactParent{
				OrganizationID: tokens[1],
			}
			resourceID = tokens[3]
		case "folders":
			parent = &ContactParent{
				FolderID: tokens[1],
			}
			resourceID = tokens[3]
		default:
			return nil, "", fmt.Errorf("format of EssentialContactsContact external=%q was not known (use projects/{{projectID}}/contacts/{{contactID}} or organizations/{{organizationID}}/contacts/{{contactID}} or folders/{{folderID}}/contacts/{{contactID}})", external)
		}
		return parent, resourceID, nil
	}
	return nil, "", fmt.Errorf("format of EssentialContactsContact external=%q was not known (use projects/{{projectID}}/contacts/{{contactID}} or organizations/{{organizationID}}/contacts/{{contactID}} or folders/{{folderID}}/contacts/{{contactID}})", external)
}
