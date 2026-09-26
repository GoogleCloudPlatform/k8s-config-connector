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

package v1beta1

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ComputeFirewallPolicyAssociationRef{}

// ComputeFirewallPolicyAssociationRef is a reference to a GCP ComputeFirewallPolicyAssociation.
type ComputeFirewallPolicyAssociationRef struct {
	// A reference to an externally managed ComputeFirewallPolicyAssociation resource.
	// Should be in the format "locations/global/firewallPolicies/{{firewallPolicy}}/associations/{{association}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeFirewallPolicyAssociation resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeFirewallPolicyAssociation resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeFirewallPolicyAssociationRef{}, &ComputeFirewallPolicyAssociation{})
}

func (r *ComputeFirewallPolicyAssociationRef) GetGVK() schema.GroupVersionKind {
	return ComputeFirewallPolicyAssociationGVK
}

func (r *ComputeFirewallPolicyAssociationRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeFirewallPolicyAssociationRef) GetExternal() string {
	return r.External
}

func (r *ComputeFirewallPolicyAssociationRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeFirewallPolicyAssociationRef) ValidateExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	id := &ComputeFirewallPolicyAssociationIdentity{}
	if err := id.FromExternal(trimmedRef); err != nil {
		return err
	}
	return nil
}

func (r *ComputeFirewallPolicyAssociationRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeFirewallPolicyAssociationIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeFirewallPolicyAssociationRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = apirefs.TrimComputeURIPrefix(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		// First check if the referenced resource is ready
		ready := false
		conditions, found, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
		if found {
			for _, condObj := range conditions {
				cond, ok := condObj.(map[string]interface{})
				if !ok {
					continue
				}
				if cond["type"] == "Ready" && cond["status"] == "True" {
					ready = true
					break
				}
			}
		}

		if !ready {
			return ""
		}

		// Convert to structured type to get spec fields
		obj, err := common.ToStructuredType[*ComputeFirewallPolicyAssociation](u)
		if err != nil {
			return ""
		}

		id, err := getIdentityFromComputeFirewallPolicyAssociationSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}

		return id.String()
	}

	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
