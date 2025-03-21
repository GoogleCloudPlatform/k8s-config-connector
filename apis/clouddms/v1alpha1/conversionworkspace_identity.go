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

// ConversionWorkspaceIdentity defines the resource reference to CloudDMSConversionWorkspace, which "External" field
// holds the GCP identifier for the KRM object.
type ConversionWorkspaceIdentity struct {
	parent *ConversionWorkspaceParent
	id     string
}

func (i *ConversionWorkspaceIdentity) String() string {
	return i.parent.String() + "/conversionworkspaces/" + i.id
}

func (i *ConversionWorkspaceIdentity) ID() string {
	return i.id
}

func (i *ConversionWorkspaceIdentity) Parent() *ConversionWorkspaceParent {
	return i.parent
}

type ConversionWorkspaceParent struct {
	ProjectID string
	Location  string
}

func (p *ConversionWorkspaceParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a ConversionWorkspaceIdentity from the Config Connector ConversionWorkspace object.
func NewConversionWorkspaceIdentity(ctx context.Context, reader client.Reader, obj *CloudDMSConversionWorkspace) (*ConversionWorkspaceIdentity, error) {

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
		actualParent, actualResourceID, err := ParseConversionWorkspaceExternal(externalRef)
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
	return &ConversionWorkspaceIdentity{
		parent: &ConversionWorkspaceParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseConversionWorkspaceExternal(external string) (parent *ConversionWorkspaceParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "conversionworkspaces" {
		return nil, "", fmt.Errorf("format of CloudDMSConversionWorkspace external=%q was not known (use projects/{{projectID}}/locations/{{location}}/conversionworkspaces/{{conversionworkspaceID}})", external)
	}
	parent = &ConversionWorkspaceParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
