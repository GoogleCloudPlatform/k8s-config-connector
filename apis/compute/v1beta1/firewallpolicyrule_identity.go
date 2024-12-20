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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FirewallPolicyRuleIdentity struct {
	id     int64
	parent *FirewallPolicyRuleParent
}

func (i *FirewallPolicyRuleIdentity) String() string {
	p := strconv.Itoa(int(i.id))
	return i.parent.String() + "/rules/" + p
}

func (i *FirewallPolicyRuleIdentity) Parent() *FirewallPolicyRuleParent {
	return i.parent
}

func (i *FirewallPolicyRuleIdentity) ID() int64 {
	return i.id
}

type FirewallPolicyRuleParent struct {
	FirewallPolicy string
}

func (p *FirewallPolicyRuleParent) String() string {
	return "locations/global/firewallPolicies/" + p.FirewallPolicy
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
	priority := obj.Spec.Priority

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actualIdentity, err := parseFirewallPolicyRuleExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualIdentity.parent.FirewallPolicy != firewallPolicy {
			return nil, fmt.Errorf("spec.FirewallPolicyRef changed, expect %s, got %s", actualIdentity.parent.FirewallPolicy, firewallPolicy)
		}
		if actualIdentity.id != priority {
			return nil, fmt.Errorf("cannot reset `spec.priority` to %d, since it has already assigned to %d",
				priority, actualIdentity.id)
		}
	}

	return &FirewallPolicyRuleIdentity{
		parent: &FirewallPolicyRuleParent{FirewallPolicy: firewallPolicy},
		id:     priority,
	}, nil
}
