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

// NotebookRuntimeTemplateIdentity defines the resource reference to ColabRuntimeTemplate, which "External" field
// holds the GCP identifier for the KRM object.
type NotebookRuntimeTemplateIdentity struct {
	parent *NotebookRuntimeTemplateParent
	id     string
}

func (i *NotebookRuntimeTemplateIdentity) String() string {
	return i.parent.String() + "/notebookRuntimeTemplates/" + i.id
}

func (i *NotebookRuntimeTemplateIdentity) ID() string {
	return i.id
}

func (i *NotebookRuntimeTemplateIdentity) Parent() *NotebookRuntimeTemplateParent {
	return i.parent
}

type NotebookRuntimeTemplateParent struct {
	ProjectID string
	Location  string
}

func (p *NotebookRuntimeTemplateParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location
}

// New builds a NotebookRuntimeTemplateIdentity from the Config Connector NotebookRuntimeTemplate object.
func NewNotebookRuntimeTemplateIdentity(ctx context.Context, reader client.Reader, obj *ColabRuntimeTemplate) (*NotebookRuntimeTemplateIdentity, error) {

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
		actualID, err := ParseNotebookRuntimeTemplateExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualID.Parent().ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualID.Parent().ProjectID, projectID)
		}
		if actualID.Parent().Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualID.Parent().Location, location)
		}
		if actualID.ID() != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualID.ID())
		}
	}
	return &NotebookRuntimeTemplateIdentity{
		parent: &NotebookRuntimeTemplateParent{
			ProjectID: projectID,
			Location:  location,
		},
		id: resourceID,
	}, nil
}

func ParseNotebookRuntimeTemplateExternal(external string) (*NotebookRuntimeTemplateIdentity, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "notebookRuntimeTemplates" {
		return nil, fmt.Errorf("format of ColabRuntimeTemplate external=%q was not known (use projects/{{projectID}}/locations/{{location}}/notebookRuntimeTemplates/{{notebookruntimetemplateID}})", external)
	}
	parent := &NotebookRuntimeTemplateParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
	}
	resourceID := tokens[5]

	return &NotebookRuntimeTemplateIdentity{
		parent: parent,
		id:     resourceID,
	}, nil
}
