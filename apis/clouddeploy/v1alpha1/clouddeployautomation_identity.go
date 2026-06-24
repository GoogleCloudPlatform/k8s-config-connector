// Copyright 2026 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &CloudDeployAutomationIdentity{}

type CloudDeployAutomationIdentity struct {
	parent *AutomationParent
	id     string
}

func (i *CloudDeployAutomationIdentity) String() string {
	return i.parent.String() + "/automations/" + i.id
}

func (i *CloudDeployAutomationIdentity) ID() string {
	return i.id
}

func (i *CloudDeployAutomationIdentity) Parent() *AutomationParent {
	return i.parent
}

func (i *CloudDeployAutomationIdentity) FromExternal(external string) error {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "deliveryPipelines" || tokens[6] != "automations" {
		return fmt.Errorf("format of CloudDeployAutomation external=%q was not known (use projects/{{projectID}}/locations/{{location}}/deliveryPipelines/{{deliveryPipelineID}}/automations/{{automationID}})", external)
	}
	i.parent = &AutomationParent{
		ProjectID:          tokens[1],
		Location:           tokens[3],
		DeliveryPipelineID: tokens[5],
	}
	i.id = tokens[7]
	return nil
}

type AutomationParent struct {
	ProjectID          string
	Location           string
	DeliveryPipelineID string
}

func (p *AutomationParent) String() string {
	return "projects/" + p.ProjectID + "/locations/" + p.Location + "/deliveryPipelines/" + p.DeliveryPipelineID
}

var _ identity.Resource = &CloudDeployAutomation{}

func (obj *CloudDeployAutomation) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get Project
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	// Get Location
	location := common.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	// Get DeliveryPipeline
	pipelineExternal, err := obj.Spec.DeliveryPipelineRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	pipelineTokens := strings.Split(pipelineExternal, "/")
	if len(pipelineTokens) != 6 || pipelineTokens[0] != "projects" || pipelineTokens[2] != "locations" || pipelineTokens[4] != "deliveryPipelines" {
		return nil, fmt.Errorf("format of DeliveryPipeline external=%q was not known", pipelineExternal)
	}
	pipelineID := pipelineTokens[5]

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
		actualIdentity := &CloudDeployAutomationIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.Parent().ProjectID != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.Parent().ProjectID, projectID)
		}
		if actualIdentity.Parent().Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualIdentity.Parent().Location, location)
		}
		if actualIdentity.Parent().DeliveryPipelineID != pipelineID {
			return nil, fmt.Errorf("spec.deliveryPipelineRef changed, expect %s, got %s", actualIdentity.Parent().DeliveryPipelineID, pipelineID)
		}
		if actualIdentity.ID() != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.ID())
		}
	}

	return &CloudDeployAutomationIdentity{
		parent: &AutomationParent{
			ProjectID:          projectID,
			Location:           location,
			DeliveryPipelineID: pipelineID,
		},
		id: resourceID,
	}, nil
}
