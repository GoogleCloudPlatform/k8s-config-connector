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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &ComputeFirewallPolicyRuleRef{}

// ComputeFirewallPolicyRuleRef defines the resource reference to ComputeFirewallPolicyRule, which "External" field
// holds the GCP identifier for the KRM object.
type ComputeFirewallPolicyRuleRef struct {
	// A reference to an externally managed ComputeFirewallPolicyRule resource.
	// Should be in the format "locations/global/firewallPolicies/{{firewallPolicy}}/rules/{{priority}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeFirewallPolicyRule resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeFirewallPolicyRule resource.
	Namespace string `json:"namespace,omitempty"`

	parent *ComputeFirewallPolicyRuleParent
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeFirewallPolicyRule.
// If the "External" is given in the other resource's spec.ComputeFirewallPolicyRuleRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ComputeFirewallPolicyRule object from the cluster.
func (r *ComputeFirewallPolicyRuleRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ComputeFirewallPolicyRuleGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := parseComputeFirewallPolicyRuleExternal(r.External); err != nil {
			return "", err
		}
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeFirewallPolicyRuleGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ComputeFirewallPolicyRuleGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", fmt.Errorf("ComputeFirewallPolicyRule is not ready yet")
	}
	r.External = actualExternalRef
	return r.External, nil
}

// New builds a NewComputeFirewallPolicyRuleRef from the Config Connector ComputeFirewallPolicyRule object.
func NewComputeFirewallPolicyRuleRef(ctx context.Context, reader client.Reader, obj *ComputeFirewallPolicyRule) (*ComputeFirewallPolicyRuleRef, error) {
	id := &ComputeFirewallPolicyRuleRef{}

	firewallPolicyRef, err := refsv1beta1.ResolveComputeFirewallPolicy(ctx, reader, obj, obj.Spec.FirewallPolicyRef)
	if err != nil {
		return nil, err
	}
	firewallPolicy := firewallPolicyRef.External
	if firewallPolicy == "" {
		return nil, fmt.Errorf("cannot resolve firewallPolicy")
	}

	id.parent = &ComputeFirewallPolicyRuleParent{FirewallPolicy: firewallPolicy}

	// Get priority. Priority is a required field
	priority := obj.Spec.Priority

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id.External = asComputeFirewallPolicyRuleExternal(id.parent, priority)
		return id, nil
	}

	// Validate desired with actual
	actualParent, actualPriority, err := parseComputeFirewallPolicyRuleExternal(externalRef)
	if err != nil {
		return nil, err
	}
	if actualParent.FirewallPolicy != firewallPolicy {
		return nil, fmt.Errorf("spec.firewallPolicyRef changed, expect %s, got %s", actualParent.FirewallPolicy, firewallPolicy)
	}
	if actualPriority != priority {
		return nil, fmt.Errorf("cannot reset `spec.priority` to %d, since it has already assigned to %d",
			priority, actualPriority)
	}
	id.External = externalRef
	id.parent = &ComputeFirewallPolicyRuleParent{FirewallPolicy: firewallPolicy}
	return id, nil
}

func (r *ComputeFirewallPolicyRuleRef) Parent() (*ComputeFirewallPolicyRuleParent, error) {
	if r.parent != nil {
		return r.parent, nil
	}
	if r.External != "" {
		parent, _, err := parseComputeFirewallPolicyRuleExternal(r.External)
		if err != nil {
			return nil, err
		}
		return parent, nil
	}
	return nil, fmt.Errorf("ComputeFirewallPolicyRule not initialized from `NewComputeFirewallPolicyRuleRef` or `NormalizedExternal`")
}

type ComputeFirewallPolicyRuleParent struct {
	FirewallPolicy string
}

func (p *ComputeFirewallPolicyRuleParent) String() string {
	return "locations/global/firewallPolicies/" + p.FirewallPolicy
}

func asComputeFirewallPolicyRuleExternal(parent *ComputeFirewallPolicyRuleParent, priority int64) (external string) {
	p := strconv.Itoa(int(priority))
	return parent.String() + "/rules/" + p
}

func parseComputeFirewallPolicyRuleExternal(external string) (parent *ComputeFirewallPolicyRuleParent, priority int64, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "locations" || tokens[2] != "firewallPolicies" || tokens[4] != "rules" {
		return nil, -1, fmt.Errorf("format of ComputeFirewallPolicyRule external=%q was not known (use firewallPolicies/{{firewallPolicy}}/rules/{{priority}})", external)
	}
	parent = &ComputeFirewallPolicyRuleParent{
		FirewallPolicy: tokens[3],
	}
	p, err := strconv.ParseInt(tokens[5], 10, 32)
	if err != nil {
		return nil, -1, fmt.Errorf("error convert priority %s of ComputeFirewallPolicyRule external=%q to an integer: %w", tokens[5], external, err)
	}
	priority = p
	return parent, priority, nil
}
