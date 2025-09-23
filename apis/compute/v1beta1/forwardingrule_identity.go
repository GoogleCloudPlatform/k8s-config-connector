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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ForwardingRuleIdentity defines the resource reference to ComputeForwardingRule, which "External" field
// holds the GCP identifier for the KRM object.
type ForwardingRuleIdentity struct {
	parent *ForwardingRuleParent
	id     string
}

func (i *ForwardingRuleIdentity) String() string {
	return i.parent.String() + "/forwardingRules/" + i.id
}

func (i *ForwardingRuleIdentity) ID() string {
	return i.id
}

func (i *ForwardingRuleIdentity) Parent() *ForwardingRuleParent {
	return i.parent
}

type ForwardingRuleParent struct {
	ProjectID string
	Region    string
}

func (p *ForwardingRuleParent) String() string {
	return "projects/" + p.ProjectID + "/region/" + p.Region
}

// New builds a ForwardingRuleIdentity from the Config Connector ForwardingRule object.
func NewForwardingRuleIdentity(ctx context.Context, reader client.Reader, obj *ComputeForwardingRule) (*ForwardingRuleIdentity, error) {
	// Get Parent
	projectID, err := refsv1beta1.ResolveProjectFromAnnotation(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	region := obj.Spec.Location

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
		actualParent, actualResourceID, err := ParseForwardingRuleExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.ProjectID != projectID.ProjectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualParent.ProjectID, projectID)
		}
		if actualParent.Region != region {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualParent.Region, region)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &ForwardingRuleIdentity{
		parent: &ForwardingRuleParent{
			ProjectID: projectID.ProjectID,
			Region:    region,
		},
		id: resourceID,
	}, nil
}

func ParseForwardingRuleExternal(external string) (parent *ForwardingRuleParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "regions" || tokens[4] != "forwardingRules" {
		return nil, "", fmt.Errorf("format of ComputeForwardingRule external=%q was not known (use projects/{{projectID}}/regions/{{region}}/forwardingRules/{{forwardingruleID}})", external)
	}
	parent = &ForwardingRuleParent{
		ProjectID: tokens[1],
		Region:    tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
