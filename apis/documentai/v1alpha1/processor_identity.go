// Copyright 2024 Google LLC
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

// ProcessorIdentity defines the resource reference to DocumentAIProcessor, which "External" field
// holds the GCP identifier for the KRM object.
type ProcessorIdentity struct {
	parent *ProcessorParent
	id     string
}

func (i *ProcessorIdentity) String() string {
	return i.parent.String() + "/processors/" + i.id
}

func (i *ProcessorIdentity) ID() string {
	return i.id
}

func (i *ProcessorIdentity) Parent() *ProcessorParent {
	return i.parent
}

type ProcessorParent struct {
	ProjectID     string
	ProjectNumber string // we need project number because the service generated ID uses project number instead of project ID
	Location      string
}

func (p *ProcessorParent) String() string {
	if p.ProjectNumber != "" {
		// if project number is available, means service has already generated the ID,
		// use project number instead of project ID because GCP API only accepts project number once the resource is created
		return "projects/" + p.ProjectNumber + "/locations/" + p.Location
	}
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a ProcessorIdentity from the Config Connector Processor object.
func NewProcessorIdentity(ctx context.Context, reader client.Reader, obj *DocumentAIProcessor) (*ProcessorIdentity, error) {
	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	projectNumber := parseProjectNumberFromExternalRef(common.ValueOf(obj.Status.ExternalRef))
	location := obj.Spec.Location

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		// get resource ID from status.ExternalRef (service generated ID) if it is available
		resourceID = parseProcessorIDFromExternalRef(common.ValueOf(obj.Status.ExternalRef))
	}
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseProcessorExternal(externalRef)
		if err != nil {
			return nil, err
		}
		/*
			// skip checking projectID because the service generated externalRef uses project number instead of project ID
			if actualParent.ProjectID != projectID {
				return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
			}
		*/
		if actualParent.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Location, location)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &ProcessorIdentity{
		parent: &ProcessorParent{
			ProjectID:     projectID,
			ProjectNumber: projectNumber,
			Location:      location,
		},
		id: resourceID,
	}, nil
}

func ParseProcessorExternal(external string) (parent *ProcessorParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "processors" {
		return nil, "", fmt.Errorf("format of DocumentAIProcessor external=%q was not known (use projects/{{projectNumber}}/locations/{{location}}/processors/{{processorID}})", external)
	}
	parent = &ProcessorParent{
		ProjectNumber: tokens[1],
		Location:      tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}

func parseProcessorIDFromExternalRef(externalRef string) string {
	_, id, err := ParseProcessorExternal(externalRef)
	if err != nil {
		return ""
	}
	return id
}

func parseProjectNumberFromExternalRef(externalRef string) string {
	parent, _, err := ParseProcessorExternal(externalRef)
	if err != nil {
		return ""
	}
	return parent.ProjectNumber
}
