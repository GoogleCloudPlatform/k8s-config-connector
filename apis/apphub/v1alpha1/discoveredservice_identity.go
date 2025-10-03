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

// DiscoveredServiceIdentity defines the resource reference to AppHubDiscoveredService, which "External" field
// holds the GCP identifier for the KRM object.
type DiscoveredServiceIdentity struct {
	parent *DiscoveredServiceParent
	id     string
}

func (i *DiscoveredServiceIdentity) String() string {
	return i.parent.String() + "/discoveredservices/" + i.id
}

func (i *DiscoveredServiceIdentity) ID() string {
	return i.id
}

func (i *DiscoveredServiceIdentity) Parent() *DiscoveredServiceParent {
	return i.parent
}

// No changes were needed to the DiscoveredServiceParent struct, String() method, or ParseDiscoveredServiceExternal function.
// No changes were needed to the DiscoveredServiceParent struct, String() method, or ParseDiscoveredServiceExternal function.
type DiscoveredServiceParent struct {
	ProjectID string
	Location  string
}

func (p *DiscoveredServiceParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a DiscoveredServiceIdentity from the Config Connector DiscoveredService object.
func NewDiscoveredServiceIdentity(ctx context.Context, reader client.Reader, obj *AppHubDiscoveredService) (*DiscoveredServiceIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.Parent.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
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
		actualParent, actualResourceID, err := ParseDiscoveredServiceExternal(externalRef)
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
	return &DiscoveredServiceIdentity{
		parent: &DiscoveredServiceParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseDiscoveredServiceExternal(external string) (parent *DiscoveredServiceParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "discoveredservices" {
		return nil, "", fmt.Errorf("format of AppHubDiscoveredService external=%q was not known (use projects/{{projectID}}/locations/{{location}}/discoveredservices/{{discoveredserviceID}})", external)
	}
	parent = &DiscoveredServiceParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
