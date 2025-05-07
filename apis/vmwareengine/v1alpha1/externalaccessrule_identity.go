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

// ExternalAccessRuleIdentity defines the resource reference to VMwareEngineExternalAccessRule, which "External" field
// holds the GCP identifier for the KRM object.
type ExternalAccessRuleIdentity struct {
	parent *ExternalAccessRuleParent
	id     string
}

func (i *ExternalAccessRuleIdentity) String() string {
	return i.parent.String() + "/externalAccessRules/" + i.id
}

func (i *ExternalAccessRuleIdentity) ID() string {
	return i.id
}

func (i *ExternalAccessRuleIdentity) Parent() *ExternalAccessRuleParent {
	return i.parent
}

type ExternalAccessRuleParent struct {
	NetworkPolicy string
}

func (p *ExternalAccessRuleParent) String() string {
	return p.NetworkPolicy
}

// New builds a ExternalAccessRuleIdentity from the Config Connector ExternalAccessRule object.
func NewExternalAccessRuleIdentity(ctx context.Context, reader client.Reader, obj *VMwareEngineExternalAccessRule) (*ExternalAccessRuleIdentity, error) {
	// Get Parent
	networkPolicyRef := obj.Spec.NetworkPolicyRef
	networkPolicy, err := networkPolicyRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
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
		actualParent, actualResourceID, err := ParseExternalAccessRuleExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.NetworkPolicy != networkPolicy {
			return nil, fmt.Errorf("spec.networkPolicy changed, expect %s, got %s", actualParent.NetworkPolicy, networkPolicy)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &ExternalAccessRuleIdentity{
		parent: &ExternalAccessRuleParent{
			NetworkPolicy: networkPolicy,
		},
		id: resourceID,
	}, nil
}

func ParseExternalAccessRuleExternal(external string) (parent *ExternalAccessRuleParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "networkPolicies" || tokens[6] != "externalAccessRules" {
		return nil, "", fmt.Errorf("format of VMwareEngineExternalAccessRule external=%q was not known (use projects/{{projectID}}/locations/{{location}}/networkPolicies/{{networkpolicyID}}/externalAccessRules/{{externalaccessruleID}})", external)
	}
	networkPolicy := strings.Join(tokens[:len(tokens)-2], "/")
	parent = &ExternalAccessRuleParent{
		NetworkPolicy: networkPolicy,
	}
	resourceID = tokens[7]
	return parent, resourceID, nil
}
