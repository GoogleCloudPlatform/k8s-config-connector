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

// FeaturestoreIdentity defines the resource reference to VertexAIFeaturestore, which "External" field
// holds the GCP identifier for the KRM object.
type FeaturestoreIdentity struct {
	parent *FeaturestoreParent
	id     string
}

func (i *FeaturestoreIdentity) String() string {
	return i.parent.String() + "/featurestores/" + i.id
}

func (i *FeaturestoreIdentity) ID() string {
	return i.id
}

func (i *FeaturestoreIdentity) Parent() *FeaturestoreParent {
	return i.parent
}

type FeaturestoreParent struct {
	ProjectID string
	Location  string
}

func (p *FeaturestoreParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a FeaturestoreIdentity from the Config Connector Featurestore object.
func NewFeaturestoreIdentity(ctx context.Context, reader client.Reader, obj *VertexAIFeaturestore) (*FeaturestoreIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
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
		actualParent, actualResourceID, err := ParseFeaturestoreExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &FeaturestoreIdentity{
		parent: &FeaturestoreParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseFeaturestoreExternal(external string) (parent *FeaturestoreParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "featurestores" {
		return nil, "", fmt.Errorf("format of VertexAIFeaturestore external=%q was not known (use projects/{{projectID}}/locations/{{location}}/featurestores/{{featurestoreID}})", external)
	}
	parent = &FeaturestoreParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
