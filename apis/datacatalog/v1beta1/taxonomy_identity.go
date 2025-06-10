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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TaxonomyIdentity defines the resource reference to DataCatalogTaxonomy, which "External" field
// holds the GCP identifier for the KRM object.
type TaxonomyIdentity struct {
	parent *TaxonomyParent
	id     string
}

func (i *TaxonomyIdentity) String() string {
	return i.parent.String() + "/taxonomies/" + i.id
}

func (i *TaxonomyIdentity) ID() string {
	return i.id
}

func (i *TaxonomyIdentity) Parent() *TaxonomyParent {
	return i.parent
}

type TaxonomyParent struct {
	ProjectID string
	Region    string
}

func (p *TaxonomyParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Region
}

// New builds a TaxonomyIdentity from the Config Connector Taxonomy object.
func NewTaxonomyIdentity(ctx context.Context, reader client.Reader, obj *DataCatalogTaxonomy) (*TaxonomyIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := common.ValueOf(obj.Spec.Region)

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseTaxonomyExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Region != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Region, location)
		}
		if resourceID != "" && actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
		resourceID = actualResourceID
	}
	return &TaxonomyIdentity{
		parent: &TaxonomyParent{
			ProjectID: projectID,
			Region:    location,
		},
		id: resourceID,
	}, nil
}

func ParseTaxonomyExternal(external string) (parent *TaxonomyParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "taxonomies" {
		return nil, "", fmt.Errorf("format of DataCatalogTaxonomy external=%q was not known (use projects/{{projectID}}/locations/{{location}}/taxonomies/{{taxonomyID}})", external)
	}
	parent = &TaxonomyParent{
		ProjectID: tokens[1],
		Region:    tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
