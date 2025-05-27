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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// PolicyTagIdentity defines the resource reference to DataCatalogPolicyTag, which "External" field
// holds the GCP identifier for the KRM object.
type PolicyTagIdentity struct {
	parent *PolicyTagParent
	id     string
}

func (i *PolicyTagIdentity) String() string {
	return i.parent.String() + "/policytags/" + i.id
}

func (i *PolicyTagIdentity) ID() string {
	return i.id
}

func (i *PolicyTagIdentity) Parent() *PolicyTagParent {
	return i.parent
}

type PolicyTagParent struct {
	ProjectID  string
	Location   string
	TaxonomyID string
}

func (p *PolicyTagParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/taxonomies/" + p.TaxonomyID
}

// New builds a PolicyTagIdentity from the Config Connector PolicyTag object.
func NewPolicyTagIdentity(ctx context.Context, reader client.Reader, obj *DataCatalogPolicyTag) (*PolicyTagIdentity, error) {

	// // Get Parent
	// projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)

	taxonomy, err := obj.Spec.TaxonomyRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}

	taxonomyParent, taxonomyID, err := ParseTaxonomyExternal(taxonomy)
	if err != nil {
		return nil, err
	}
	projectID := taxonomyParent.ProjectID
	location := taxonomyParent.Region

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
		actualParent, actualResourceID, err := ParsePolicyTagExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("Project ID changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("Location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualParent.TaxonomyID != taxonomyID {
			return nil, fmt.Errorf("Taxonomy Id changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &PolicyTagIdentity{
		parent: &PolicyTagParent{
			ProjectID:  projectID,
			Location:   location,
			TaxonomyID: taxonomyID,
		},
		id: resourceID,
	}, nil
}

func ParsePolicyTagExternal(external string) (parent *PolicyTagParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "taxonomies" || tokens[6] != "policytags" {
		return nil, "", fmt.Errorf("format of DataCatalogPolicyTag external=%q was not known (use projects/{{projectID}}/locations/{{location}}/taxonomies/{{taxonomyID}}/policytags/{{policytagID}})", external)
	}
	parent = &PolicyTagParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
