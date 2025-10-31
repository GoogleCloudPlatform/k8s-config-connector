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

package v1beta1

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FirewallPolicyRuleIdentity struct {
	resourceID int32
	Parent     *FirewallPolicyIdentity
}

func (i *FirewallPolicyRuleIdentity) String() string {
	p := strconv.Itoa(int(i.resourceID))
	return i.Parent.String() + "/rules/" + p
}

func NewFirewallPolicyRuleIdentity(ctx context.Context, reader client.Reader, obj *ComputeFirewallPolicyRule) (*FirewallPolicyRuleIdentity, error) {
	//Get parent
	firewallPolicyRef, err := refsv1beta1.ResolveComputeFirewallPolicy(ctx, reader, obj, obj.Spec.FirewallPolicyRef)
	if err != nil {
		return nil, err
	}
	firewallPolicy := firewallPolicyRef.External
	if firewallPolicy == "" {
		return nil, fmt.Errorf("cannot resolve firewallPolicy")
	}

	// Get priority. Priority is a required field
	priority := common.ValueOf(obj.Spec.Priority)

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actualIdentity, err := parseFirewallPolicyRuleExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualIdentity.Parent.ResourceID != firewallPolicy {
			return nil, fmt.Errorf("spec.FirewallPolicyRef changed, expect %s, got %s", actualIdentity.Parent.String(), firewallPolicy)
		}
		if actualIdentity.resourceID != priority {
			return nil, fmt.Errorf("cannot reset `spec.priority` to %d, since it has already assigned to %d",
				priority, actualIdentity.resourceID)
		}
	}

	return &FirewallPolicyRuleIdentity{
		Parent:     &FirewallPolicyIdentity{ResourceID: firewallPolicy},
		resourceID: priority,
	}, nil
}

func parseFirewallPolicyRuleExternal(external string) (*FirewallPolicyRuleIdentity, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "locations" || tokens[2] != "firewallPolicies" || tokens[4] != "rules" {
		return nil, fmt.Errorf("format of ComputeFirewallPolicyRule external=%q was not known (use firewallPolicies/{{firewallPolicy}}/rules/{{priority}})", external)
	}
	p, err := strconv.ParseInt(tokens[5], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("error convert priority %s of ComputeFirewallPolicyRule external=%q to an integer: %w", tokens[5], external, err)
	}
	return &FirewallPolicyRuleIdentity{Parent: &FirewallPolicyIdentity{ResourceID: tokens[3]}, resourceID: int32(p)}, nil
}
