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

// TableIdentity defines the resource reference to BigLakeTable, which "External" field
// holds the GCP identifier for the KRM object.
type TableIdentity struct {
	parent *TableParent
	id     string
}

func (i *TableIdentity) String() string {
	return i.parent.String() + "/tables/" + i.id
}

func (i *TableIdentity) ID() string {
	return i.id
}

func (i *TableIdentity) Parent() *TableParent {
	return i.parent
}

type TableParent struct {
	ProjectID  string
	Location   string
	CatalogID  string
	DatabaseID string
}

func (p *TableParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/catalogs/" + p.CatalogID + "/databases/" + p.DatabaseID
}

// New builds a TableIdentity from the Config Connector Table object.
func NewTableIdentity(ctx context.Context, reader client.Reader, obj *BigLakeTable) (*TableIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := common.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	catalogRef := obj.Spec.CatalogRef
	if catalogRef == nil {
		return nil, fmt.Errorf("cannot resolve catalog")
	}
	catalogExternal, err := catalogRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, fmt.Errorf("cannot normalize catalog ref: %w", err)
	}
	if catalogExternal == "" {
		return nil, fmt.Errorf("could not resolve external catalog")
	}
	_, catalogID, err := ParseCatalogExternal(catalogExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse catalog external: %w", err)
	}

	databaseRef := obj.Spec.DatabaseRef
	if databaseRef == nil {
		return nil, fmt.Errorf("cannot resolve database")
	}
	databaseExternal, err := databaseRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, fmt.Errorf("cannot normalize database ref: %w", err)
	}
	if databaseExternal == "" {
		return nil, fmt.Errorf("could not resolve external database")
	}
	_, databaseID, err := ParseDatabaseExternal(databaseExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse database external: %w", err)
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
		actualParent, actualResourceID, err := ParseTableExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualParent.CatalogID != catalogID {
			return nil, fmt.Errorf("spec.catalogRef changed, expect %s, got %s", actualParent.CatalogID, catalogID)
		}
		if actualParent.DatabaseID != databaseID {
			return nil, fmt.Errorf("spec.databaseRef changed, expect %s, got %s", actualParent.DatabaseID, databaseID)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &TableIdentity{
		parent: &TableParent{
			ProjectID:  projectID,
			Location:   location,
			CatalogID:  catalogID,
			DatabaseID: databaseID,
		},
		id: resourceID,
	}, nil
}

func ParseTableExternal(external string) (parent *TableParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 10 ||
		tokens[0] != "projects" ||
		tokens[2] != "locations" ||
		tokens[4] != "catalogs" ||
		tokens[6] != "databases" ||
		tokens[8] != "tables" {
		return nil, "", fmt.Errorf("format of BigLakeTable external=%q was not known (use projects/{{projectID}}/locations/{{location}}/catalogs/{{catalog}}/databases/{{database}}/tables/{{tableID}})", external)
	}
	parent = &TableParent{
		ProjectID:  tokens[1],
		Location:   tokens[3],
		CatalogID:  tokens[5],
		DatabaseID: tokens[7],
	}
	resourceID = tokens[9]
	return parent, resourceID, nil
}
