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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AutomationIdentity defines the resource reference to CloudDeployAutomation, which "External" field
// holds the GCP identifier for the KRM object.
type AutomationIdentity struct {
	parent *DeliveryPipelineIdentity
	id     string
}

func (i *AutomationIdentity) String() string {
	return i.parent.String() + "/automations/" + i.id
}

func (i *AutomationIdentity) ID() string {
	return i.id
}

func (i *AutomationIdentity) Parent() *DeliveryPipelineIdentity {
	return i.parent
}

// New builds a AutomationIdentity from the Config Connector Automation object.
func NewAutomationIdentity(ctx context.Context, reader client.Reader, obj *CloudDeployAutomation) (*AutomationIdentity, error) {
	parent, err := NewDeliveryPipelineIdentityFromRef(ctx, reader, obj.GetNamespace(), &obj.Spec.DeliveryPipelineRef)
	if err != nil {
		return nil, fmt.Errorf("resolving parent %v: %w", obj.Spec.DeliveryPipelineRef, err)
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
		actualParent, actualResourceID, err := ParseAutomationExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.String() != parent.String() {
			return nil, fmt.Errorf("parent changed, expect %s, got %s", actualParent.String(), parent.String())
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &AutomationIdentity{
		parent: parent,
		id:     resourceID,
	}, nil
}

func ParseAutomationExternal(external string) (parent *DeliveryPipelineIdentity, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "deliveryPipelines" || tokens[6] != "automations" {
		return nil, "", fmt.Errorf("format of CloudDeployAutomation external=%q was not known (use projects/{{projectID}}/locations/{{location}}/deliveryPipelines/{{deliveryPipelineID}}/automations/{{automationID}})", external)
	}
	parent = &DeliveryPipelineIdentity{
		parent: &DeliveryPipelineParent{
			ProjectID: tokens[1],
			Location:  tokens[3],
		},
		id: tokens[5],
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
