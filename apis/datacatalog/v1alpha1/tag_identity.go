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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TagIdentity defines the resource reference to DataCatalogTag, which "External" field
// holds the GCP identifier for the KRM object.
type TagIdentity struct {
	parent *TagParent
	id     string
}

func (i *TagIdentity) String() string {
	return i.parent.String() + "/tags/" + i.id
}

func (i *TagIdentity) ID() string {
	return i.id
}

func (i *TagIdentity) Parent() *TagParent {
	return i.parent
}

// No changes were needed to the TagParent struct.
type TagParent struct {
	ProjectID    string
	Location     string
	EntryGroupID string
	EntryID      string
}

func (p *TagParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/entryGroups/" + p.EntryGroupID + "/entries/" + p.EntryID
}

// New builds a TagIdentity from the Config Connector Tag object.
func NewTagIdentity(ctx context.Context, reader client.Reader, obj *DataCatalogTag) (*TagIdentity, error) {

	// Get Parent
	tagParent, err := ParseEntryExternal(obj.Spec.Parent.EntryRef.External)
	if err != nil {
		return nil, err
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
		actualParent, actualResourceID, err := ParseTagExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != tagParent.ProjectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, tagParent.ProjectID)
		}
		if actualParent.Location != tagParent.Location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, tagParent.Location)
		}
		if actualParent.EntryGroupID != tagParent.EntryGroupID {
			return nil, fmt.Errorf("spec.parent.entryGroupRef changed, expect %s, got %s", actualParent.EntryGroupID, tagParent.EntryGroupID)
		}
		if actualParent.EntryID != tagParent.EntryID {
			return nil, fmt.Errorf("spec.parent.entryRef changed, expect %s, got %s", actualParent.EntryID, tagParent.EntryID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &TagIdentity{
		parent: tagParent,
		id:     resourceID,
	}, nil
}

func ParseTagExternal(external string) (parent *TagParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "entryGroups" || tokens[6] != "entries" || tokens[8] != "tags" {
		return nil, "", fmt.Errorf("format of DataCatalogTag external=%q was not known (use projects/{{projectID}}/locations/{{location}}/entryGroups/{{entryGroupID}}/entries/{{entryID}}/tags/{{tagID}})", external)
	}
	parent = &TagParent{
		ProjectID:    tokens[1],
		Location:     tokens[3],
		EntryGroupID: tokens[5],
		EntryID:      tokens[7],
	}
	resourceID = tokens[9]
	return parent, resourceID, nil
}

// ParseEntryExternal parses the external reference to a DataCatalogEntry.
func ParseEntryExternal(external string) (parent *TagParent, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "entryGroups" || tokens[6] != "entries" {
		return nil, fmt.Errorf("format of DataCatalogEntry external=%q was not known (use projects/{{projectID}}/locations/{{location}}/entryGroups/{{entryGroupID}}/entries/{{entryID}})", external)
	}
	parent = &TagParent{
		ProjectID:    tokens[1],
		Location:     tokens[3],
		EntryGroupID: tokens[5],
		EntryID:      tokens[7],
	}
	return parent, nil
}
