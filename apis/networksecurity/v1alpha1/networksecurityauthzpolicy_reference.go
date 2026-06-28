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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &NetworkSecurityAuthzPolicyRef{}

// NetworkSecurityAuthzPolicyRef defines the resource reference to NetworkSecurityAuthzPolicy, which "External" field
type NetworkSecurityAuthzPolicyRef struct {
	// A reference to an externally managed NetworkSecurityAuthzPolicy resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/authzPolicies/{{authzpolicy}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkSecurityAuthzPolicy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkSecurityAuthzPolicy resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&NetworkSecurityAuthzPolicyRef{})
}

func (r *NetworkSecurityAuthzPolicyRef) GetGVK() schema.GroupVersionKind {
	return NetworkSecurityAuthzPolicyGVK
}

func (r *NetworkSecurityAuthzPolicyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkSecurityAuthzPolicyRef) GetExternal() string {
	return r.External
}

func (r *NetworkSecurityAuthzPolicyRef) SetExternal(external string) {
	r.External = external
}

func (r *NetworkSecurityAuthzPolicyRef) ValidateExternal(external string) error {
	id := &NetworkSecurityAuthzPolicyIdentity{}
	return id.FromExternal(external)
}

func (r *NetworkSecurityAuthzPolicyRef) ParseExternalToIdentity(external string) (interface{}, error) {
	id := &NetworkSecurityAuthzPolicyIdentity{}
	err := id.FromExternal(external)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetworkSecurityAuthzPolicyRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromNetworkSecurityAuthzPolicySpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	})
}
