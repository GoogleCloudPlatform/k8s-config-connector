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

// SavedQueryIdentity defines the resource reference to AssetSavedQuery, which "External" field
// holds the GCP identifier for the KRM object.
type SavedQueryIdentity struct {
	parent *SavedQueryParent
	id     string
}

func (i *SavedQueryIdentity) String() string {
	return i.parent.String() + "/savedQueries/" + i.id
}

func (i *SavedQueryIdentity) ID() string {
	return i.id
}

func (i *SavedQueryIdentity) Parent() *SavedQueryParent {
	return i.parent
}

// No changes were needed.
type SavedQueryParent struct {
	ProjectID      string
	FolderID       string
	OrganizationID string
}

func (p *SavedQueryParent) String() string {
	if p.ProjectID != "" {
		return "projects/" + p.ProjectID
	}
	if p.FolderID != "" {
		return "folders/" + p.FolderID
	}
	return "organizations/" + p.OrganizationID
}

// New builds a SavedQueryIdentity from the Config Connector SavedQuery object.
func NewSavedQueryIdentity(ctx context.Context, reader client.Reader, obj *AssetSavedQuery) (*SavedQueryIdentity, error) {
	var projectID, folderID, organizationID string
	// Get Parent
	if obj.Spec.Parent.ProjectRef != nil {
		projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.Parent.ProjectRef)
		if err != nil {
			return nil, err
		}
		projectID = projectRef.ProjectID
		if projectID == "" {
			return nil, fmt.Errorf("cannot resolve project")
		}
	} else if obj.Spec.Parent.FolderRef != nil {
		folderRef, err := refsv1beta1.ResolveFolder(ctx, reader, obj, obj.Spec.Parent.FolderRef)
		if err != nil {
			return nil, err
		}
		folderID = folderRef.FolderID
		if folderID == "" {
			return nil, fmt.Errorf("cannot resolve folder")
		}
	} else if obj.Spec.Parent.OrganizationRef != nil {
		organizationRef, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, obj.Spec.Parent.OrganizationRef)
		if err != nil {
			return nil, err
		}
		organizationID = organizationRef.OrganizationID
		if organizationID == "" {
			return nil, fmt.Errorf("cannot resolve organization")
		}
	} else {
		return nil, fmt.Errorf("one of spec.parent.projectRef, spec.parent.folderRef, or spec.parent.organizationRef must be set")
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
		actualParent, actualResourceID, err := ParseSavedQueryExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.FolderID != folderID {
			return nil, fmt.Errorf("spec.folderRef changed, expect %s, got %s", actualParent.FolderID, folderID)
		}
		if actualParent.OrganizationID != organizationID {
			return nil, fmt.Errorf("spec.organizationRef changed, expect %s, got %s", actualParent.OrganizationID, organizationID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &SavedQueryIdentity{
		parent: &SavedQueryParent{
			ProjectID:      projectID,
			FolderID:       folderID,
			OrganizationID: organizationID,
		},
		id: resourceID,
	}, nil
}

func ParseSavedQueryExternal(external string) (parent *SavedQueryParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) == 4 && tokens[2] == "savedQueries" {
		parent = &SavedQueryParent{}
		switch tokens[0] {
		case "projects":
			parent.ProjectID = tokens[1]
		case "folders":
			parent.FolderID = tokens[1]
		case "organizations":
			parent.OrganizationID = tokens[1]
		default:
			return nil, "", fmt.Errorf("format of AssetSavedQuery external=%q was not known (use projects/{{projectID}}/savedQueries/{{savedqueryID}})", external)
		}
		resourceID = tokens[3]
		return parent, resourceID, nil
	}
	return nil, "", fmt.Errorf("format of AssetSavedQuery external=%q was not known (use projects/{{projectID}}/savedQueries/{{savedqueryID}})", external)
}
