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
	"fmt"
	"strings"
)

// DatabaseIdentity defines the resource reference to BigLakeDatabase, which "External" field
// holds the GCP identifier for the KRM object.
type DatabaseIdentity struct {
	parent *DatabaseParent
	id     string
}

func (i *DatabaseIdentity) String() string {
	return i.parent.String() + "/databases/" + i.id
}

func (i *DatabaseIdentity) ID() string {
	return i.id
}

func (i *DatabaseIdentity) Parent() *DatabaseParent {
	return i.parent
}

type DatabaseParent struct {
	ProjectID string
	Location  string
	CatalogID string
}

func (p *DatabaseParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/catalogs/" + p.CatalogID
}

// NOT YET
// // New builds a DatabaseIdentity from the Config Connector Database object.
// func NewDatabaseIdentity(ctx context.Context, reader client.Reader, obj *BigLakeDatabase) (*DatabaseIdentity, error) {

// 	// Get Parent
// 	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
// 	if err != nil {
// 		return nil, err
// 	}
// 	projectID := projectRef.ProjectID
// 	if projectID == "" {
// 		return nil, fmt.Errorf("cannot resolve project")
// 	}
// 	location := obj.Spec.Location

// 	// Get desired ID
// 	resourceID := common.ValueOf(obj.Spec.ResourceID)
// 	if resourceID == "" {
// 		resourceID = obj.GetName()
// 	}
// 	if resourceID == "" {
// 		return nil, fmt.Errorf("cannot resolve resource ID")
// 	}

// 	// Use approved External
// 	externalRef := common.ValueOf(obj.Status.ExternalRef)
// 	if externalRef != "" {
// 		// Validate desired with actual
// 		actualParent, actualResourceID, err := ParseDatabaseExternal(externalRef)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if actualParent.ProjectID != projectID {
// 			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
// 		}
// 		if actualParent.Location != location {
// 			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
// 		}
// 		if actualResourceID != resourceID {
// 			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
// 				resourceID, actualResourceID)
// 		}
// 	}
// 	return &DatabaseIdentity{
// 		parent: &DatabaseParent{
// 			ProjectID: projectID,
// 			Location:  location,
// 		},
// 		id: resourceID,
// 	}, nil
// }

func ParseDatabaseExternal(external string) (parent *DatabaseParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "catalogs" || tokens[6] != "databases" {
		return nil, "", fmt.Errorf("format of BigLakeDatabase external=%q was not known (use projects/{{projectID}}/locations/{{location}}/databases/{{databaseID}})", external)
	}
	parent = &DatabaseParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
		CatalogID: tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
