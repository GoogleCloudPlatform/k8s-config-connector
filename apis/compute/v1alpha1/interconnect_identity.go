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

// InterconnectIdentity defines the resource reference to ComputeInterconnect, which "External" field
// holds the GCP identifier for the KRM object.
type InterconnectIdentity struct {
	parent *InterconnectParent
	id     string
}

func (i *InterconnectIdentity) String() string {
	return i.parent.String() + "/interconnects/" + i.id
}

func (i *InterconnectIdentity) ID() string {
	return i.id
}

func (i *InterconnectIdentity) Parent() *InterconnectParent {
	return i.parent
}

// No changes needed to the InterconnectParent struct.
type InterconnectParent struct {
	ProjectID string
	Location  string
}

func (p *InterconnectParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a InterconnectIdentity from the Config Connector Interconnect object.
func NewInterconnectIdentity(ctx context.Context, reader client.Reader, obj *ComputeInterconnect) (*InterconnectIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.NewProjectRefFromResource(obj.ObjectMeta)
	if err != nil {
		return nil, fmt.Errorf("error getting project reference: %w", err)
	}
	projectID := projectRef.ProjectID
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
		actualParent, actualResourceID, err := ParseInterconnectExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}

		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &InterconnectIdentity{
		parent: &InterconnectParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseInterconnectExternal(external string) (parent *InterconnectParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "interconnects" {
		return nil, "", fmt.Errorf("format of ComputeInterconnect external=%q was not known (use projects/{{projectID}}/locations/{{location}}/interconnects/{{interconnectID}})", external)
	}
	parent = &InterconnectParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
