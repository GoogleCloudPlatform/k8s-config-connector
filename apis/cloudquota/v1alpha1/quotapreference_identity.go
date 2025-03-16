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

// QuotaPreferenceIdentity defines the resource reference to APIQuotaPreference, which "External" field
// holds the GCP identifier for the KRM object.
type QuotaPreferenceIdentity struct {
	parent *QuotaPreferenceParent
	id     string
}

func (i *QuotaPreferenceIdentity) String() string {
	return i.parent.String() + "/quotaPreferences/" + i.id
}

func (i *QuotaPreferenceIdentity) ID() string {
	return i.id
}

func (i *QuotaPreferenceIdentity) Parent() *QuotaPreferenceParent {
	return i.parent
}

// QuotaPreferenceParent defines the GCP project, folder or organization, and location.
type QuotaPreferenceParent struct {
	ProjectID      string
	OrganizationID string
	FolderID       string
	Location       string
}

func (p *QuotaPreferenceParent) String() string {
	prefix := ""
	if p.ProjectID != "" {
		prefix = "projects/" + p.ProjectID
	} else if p.OrganizationID != "" {
		prefix = "organizations/" + p.OrganizationID
	} else if p.FolderID != "" {
		prefix = "folders/" + p.FolderID
	} else {
		return ""
	}
	return prefix + "/locations/" + p.Location
}

// ParseQuotaPreferenceExternal parses the external quota preference name.
func ParseQuotaPreferenceExternal(external string) (*QuotaPreferenceParent, string, error) {
	parts := strings.Split(external, "/")
	if len(parts) != 6 {
		return nil, "", fmt.Errorf("invalid external quota preference format: %s", external)
	}
	// projects/{project}/locations/{location}/quotaPreferences/{quotaPreference}
	// organizations/{organization}/locations/{location}/quotaPreferences/{quotaPreference}
	// folders/{folder}/locations/{location}/quotaPreferences/{quotaPreference}
	if parts[2] != "locations" || parts[4] != "quotaPreferences" {
		return nil, "", fmt.Errorf("invalid external quota preference format: %s", external)
	}
	parent := &QuotaPreferenceParent{
		Location: parts[3],
	}
	switch parts[0] {
	case "projects":
		parent.ProjectID = parts[1]
	case "organizations":
		parent.OrganizationID = parts[1]
	case "folders":
		parent.FolderID = parts[1]
	default:
		return nil, "", fmt.Errorf("invalid external quota preference format: %s", external)
	}
	return parent, parts[5], nil
}

// New builds a QuotaPreferenceIdentity from the Config Connector QuotaPreference object.
func NewQuotaPreferenceIdentity(ctx context.Context, reader client.Reader, obj *APIQuotaPreference) (*QuotaPreferenceIdentity, error) {
	var projectID, organizationID, folderID string
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
	} else if obj.Spec.Parent.OrganizationRef != nil {
		organizationRef, err := refsv1beta1.ResolveOrganization(ctx, reader, obj, obj.Spec.Parent.OrganizationRef)
		if err != nil {
			return nil, err
		}
		organizationID = organizationRef.OrganizationID
		if organizationID == "" {
			return nil, fmt.Errorf("cannot resolve organization")
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
	} else {
		return nil, fmt.Errorf("one of spec.projectRef, spec.organizationRef, or spec.folderRef must be set")
	}
	location := obj.Spec.Parent.Location

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
		actualParent, actualResourceID, err := ParseQuotaPreferenceExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != "" && actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.OrganizationID != "" && actualParent.OrganizationID != organizationID {
			return nil, fmt.Errorf("spec.organizationRef changed, expect %s, got %s", actualParent.OrganizationID, organizationID)
		}
		if actualParent.FolderID != "" && actualParent.FolderID != folderID {
			return nil, fmt.Errorf("spec.folderRef changed, expect %s, got %s", actualParent.FolderID, folderID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &QuotaPreferenceIdentity{
		parent: &QuotaPreferenceParent{
			ProjectID:      projectID,
			OrganizationID: organizationID,
			FolderID:       folderID,
			Location:       location,
		},
		id: resourceID,
	}, nil
}
