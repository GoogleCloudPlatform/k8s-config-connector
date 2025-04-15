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
	workflow "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflows/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ExecutionIdentity defines the resource reference to WorkflowsExecution, which "External" field
// holds the GCP identifier for the KRM object.
type ExecutionIdentity struct {
	parent *ExecutionParent
	id     string
}

func (i *ExecutionIdentity) String() string {
	return i.parent.String() + "/executions/" + i.id
}

func (i *ExecutionIdentity) ID() string {
	return i.id
}

func (i *ExecutionIdentity) Parent() *ExecutionParent {
	return i.parent
}

type ExecutionParent struct {
	ProjectID string
	Location  string
	Workflow  string
}

func (p *ExecutionParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/workflows/" + p.Workflow
}

// New builds a ExecutionIdentity from the Config Connector Execution object.
func NewExecutionIdentity(ctx context.Context, reader client.Reader, obj *WorkflowsExecution) (*ExecutionIdentity, error) {

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
	_, workflow, err := workflow.ParseWorkflowsWorkflowExternal(obj.Spec.WorkflowRef.External)
	if err != nil {
		return nil, err
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
		actualParent, actualResourceID, err := ParseExecutionExternal(externalRef)
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
	return &ExecutionIdentity{
		parent: &ExecutionParent{
			ProjectID: projectID,
			Location:  location,
			Workflow:  workflow,
		},
		id: resourceID,
	}, nil
}

func ParseExecutionExternal(external string) (parent *ExecutionParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "workflows" || tokens[6] != "executions" {
		return nil, "", fmt.Errorf("format of WorkflowsExecution external=%q was not known (use projects/{{projectID}}/locations/{{location}}/workflows/{{workflow}}/executions/{{executionID}})", external)
	}
	parent = &ExecutionParent{
		ProjectID: tokens[1],
		Location:  tokens[3],
		Workflow:  tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
