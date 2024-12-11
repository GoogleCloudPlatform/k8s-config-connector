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
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &ComputeFirewallPolicyRef{}
var ComputeFirewallPolicyGVK = GroupVersion.WithKind("ComputeFirewallPolicy")

// ComputeFirewallPolicyRef defines the resource reference to ComputeFirewallPolicy, which "External" field
// holds the GCP identifier for the KRM object.
type ComputeFirewallPolicyRef struct {
	// A reference to an externally managed ComputeFirewallPolicy resource.
	// Should be in the format "locations/global/firewallPolicies/{{firewallPolicy}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeFirewallPolicy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeFirewallPolicy resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on ComputeFirewallPolicy.
// If the "External" is given in the other resource's spec.ComputeFirewallPolicyRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ComputeFirewallPolicy object from the cluster.
func (r *ComputeFirewallPolicyRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ComputeFirewallPolicyGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, err := parseComputeFirewallPolicyExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(ComputeFirewallPolicyGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ComputeFirewallPolicyGVK, key, err)
	}

	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, found, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("error getting status.externalRef for %s %s/%s: %w", u.GetKind(), u.GetNamespace(), u.GetName(), err)
	}

	// If status.externalRef does not exist, it's created by legacy controller. Get values from target field.
	if !found {
		resourceID, _, err := unstructured.NestedString(u.Object, "spec", "resourceID")
		if err != nil {
			return "", fmt.Errorf("reading spec.resourceID from %v %v/%v: %w", u.GroupVersionKind().Kind, u.GetNamespace(), u.GetName(), err)
		}
		if resourceID == "" {
			resourceID = u.GetName()
		}
		r.External = resourceID
	} else {
		r.External = actualExternalRef
	}

	return r.External, nil
}

func parseComputeFirewallPolicyExternal(external string) (firewallPolicy string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) == 4 && tokens[0] == "locations" && tokens[1] == "global" && tokens[2] == "firewallPolicies" {
		return tokens[3], nil
	}
	return "", fmt.Errorf("format of ComputeFirewallPolicy external=%q was not known (use locations/global/firewallPolicies/{{firewallPolicy}})", external)
}
