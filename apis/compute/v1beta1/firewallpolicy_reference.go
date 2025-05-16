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
	External string `json:"external,omitempty"`

	// The name of a ComputeFirewallPolicyRule resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeFirewallPolicyRule resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ComputeFirewallPolicyRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ComputeFirewallPolicyGVK.Kind)
	}

	// From given External
	// For backward compatibility, we are not enforcing the external format.
	// todo: validate external when it's referenced by a pure direct resource
	if r.External != "" {
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
		return "", fmt.Errorf("error reading referenced ComputeFirewallPolicy %v: %w", key, err)
	}

	externalRef, _, _ := unstructured.NestedString(u.Object, "status", "externalRef")
	if externalRef != "" {
		return externalRef, nil
	}

	selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
	if selfLink == "" {
		return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
	}

	external := fixStaleExternalFormat(selfLink)
	return external, nil
}

func ParseFirewallPolicyExternal(external string) (*FirewallPolicyIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeFirewallPolicy external value")
	}
	external = fixStaleExternalFormat(external)
	tokens := strings.Split(external, "/")
	if len(tokens) == 4 && tokens[0] == "locations" && tokens[1] == "global" && tokens[2] == "firewallPolicies" {
		return &FirewallPolicyIdentity{ID: tokens[3]}, nil
	}

	return nil, fmt.Errorf("format of ComputeFirewallPolicy external=%q was not known (use locations/global/firewallPolicies/{{firewallPolicy}})", external)

}
