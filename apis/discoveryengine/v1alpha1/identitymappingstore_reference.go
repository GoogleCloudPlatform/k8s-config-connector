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
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &DiscoveryEngineIdentityMappingStoreRef{}

type DiscoveryEngineIdentityMappingStoreRef struct {
	// A reference to an externally managed DiscoveryEngineIdentityMappingStore resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/identityMappingStores/{{identitymappingstore}}"
	External string `json:"external,omitempty"`

	// The name of a DiscoveryEngineIdentityMappingStore resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DiscoveryEngineIdentityMappingStore resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&DiscoveryEngineIdentityMappingStoreRef{})
}

func (r *DiscoveryEngineIdentityMappingStoreRef) GetGVK() schema.GroupVersionKind {
	return DiscoveryEngineIdentityMappingStoreGVK
}

func (r *DiscoveryEngineIdentityMappingStoreRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DiscoveryEngineIdentityMappingStoreRef) GetExternal() string {
	return r.External
}

func (r *DiscoveryEngineIdentityMappingStoreRef) SetExternal(external string) {
	r.External = external
}

func (r *DiscoveryEngineIdentityMappingStoreRef) ValidateExternal() error {
	id := &DiscoveryEngineIdentityMappingStoreIdentity{}
	return id.FromExternal(r.External)
}

func (r *DiscoveryEngineIdentityMappingStoreRef) ParseExternalToIdentity() (any, error) {
	id := &DiscoveryEngineIdentityMappingStoreIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DiscoveryEngineIdentityMappingStoreRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) (string, error) {
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) (string, error) {
		id, err := getIdentityFromDiscoveryEngineIdentityMappingStoreSpec(ctx, reader, u)
		if err != nil {
			return "", fmt.Errorf("failed to get identity from spec: %w", err)
		}
		return id.String(), nil
	})
}
