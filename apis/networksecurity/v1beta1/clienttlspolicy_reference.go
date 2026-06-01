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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &NetworkSecurityClientTLSPolicyRef{}

// NetworkSecurityClientTLSPolicyRef defines the resource reference to NetworkSecurityClientTLSPolicy, which "External" field
// holds the GCP identifier for the KRM object.
type NetworkSecurityClientTLSPolicyRef struct {
	// A reference to an externally managed NetworkSecurityClientTLSPolicy resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/clientTlsPolicies/{{clientTlsPolicyID}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkSecurityClientTLSPolicy resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkSecurityClientTLSPolicy resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&NetworkSecurityClientTLSPolicyRef{})
}

func (r *NetworkSecurityClientTLSPolicyRef) GetGVK() schema.GroupVersionKind {
	return NetworkSecurityClientTLSPolicyGVK
}

func (r *NetworkSecurityClientTLSPolicyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkSecurityClientTLSPolicyRef) GetExternal() string {
	return r.External
}

func (r *NetworkSecurityClientTLSPolicyRef) SetExternal(ref string) {
	r.External = ref
}

func (r *NetworkSecurityClientTLSPolicyRef) ValidateExternal(ref string) error {
	id := &NetworkSecurityClientTLSPolicyIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *NetworkSecurityClientTLSPolicyRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetworkSecurityClientTLSPolicyIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetworkSecurityClientTLSPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromNetworkSecurityClientTLSPolicySpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
