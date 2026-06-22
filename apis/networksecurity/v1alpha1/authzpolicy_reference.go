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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &NetworkSecurityAuthzPolicyRef{}

// NetworkSecurityAuthzPolicyRef is a reference to a NetworkSecurityAuthzPolicy resource.
type NetworkSecurityAuthzPolicyRef struct {
	// A reference to an externally managed NetworkSecurityAuthzPolicy resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/authzPolicies/{{authzPolicy}}".
	// +optional
	External string `json:"external,omitempty"`

	// The name of a NetworkSecurityAuthzPolicy resource.
	// +optional
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkSecurityAuthzPolicy resource.
	// +optional
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
		Namespace: r.Namespace,
		Name:      r.Name,
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

func (r *NetworkSecurityAuthzPolicyRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetworkSecurityAuthzPolicyIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetworkSecurityAuthzPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		id, err := getIdentityFromNetworkSecurityAuthzPolicySpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
