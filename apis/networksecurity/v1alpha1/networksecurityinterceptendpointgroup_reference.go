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

var _ refs.Ref = &NetworkSecurityInterceptEndpointGroupRef{}

// NetworkSecurityInterceptEndpointGroupRef is a reference to a NetworkSecurityInterceptEndpointGroup.
type NetworkSecurityInterceptEndpointGroupRef struct {
	// A reference to an externally managed NetworkSecurityInterceptEndpointGroup resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/interceptEndpointGroups/{{intercept_endpoint_group}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkSecurityInterceptEndpointGroup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkSecurityInterceptEndpointGroup resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&NetworkSecurityInterceptEndpointGroupRef{}, &NetworkSecurityInterceptEndpointGroup{})
}

func (r *NetworkSecurityInterceptEndpointGroupRef) GetGVK() schema.GroupVersionKind {
	return NetworkSecurityInterceptEndpointGroupGVK
}

func (r *NetworkSecurityInterceptEndpointGroupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkSecurityInterceptEndpointGroupRef) GetExternal() string {
	return r.External
}

func (r *NetworkSecurityInterceptEndpointGroupRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *NetworkSecurityInterceptEndpointGroupRef) ValidateExternal(ref string) error {
	id := &NetworkSecurityInterceptEndpointGroupIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *NetworkSecurityInterceptEndpointGroupRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetworkSecurityInterceptEndpointGroupIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetworkSecurityInterceptEndpointGroupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromNetworkSecurityInterceptEndpointGroupSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
