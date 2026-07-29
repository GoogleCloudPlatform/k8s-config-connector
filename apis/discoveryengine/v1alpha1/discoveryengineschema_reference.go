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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &DiscoveryEngineSchemaRef{}

// DiscoveryEngineSchemaRef is a reference to a GCP DiscoveryEngineSchema.
type DiscoveryEngineSchemaRef struct {
	// A reference to an externally managed DiscoveryEngineSchema resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/dataStores/{{dataStore}}/schemas/{{schema}}".
	External string `json:"external,omitempty"`

	// The name of a DiscoveryEngineSchema resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DiscoveryEngineSchema resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&DiscoveryEngineSchemaRef{}, &DiscoveryEngineSchema{})
}

func (r *DiscoveryEngineSchemaRef) GetGVK() schema.GroupVersionKind {
	return DiscoveryEngineSchemaGVK
}

func (r *DiscoveryEngineSchemaRef) GetNamespacedName() client.ObjectKey {
	return client.ObjectKey{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *DiscoveryEngineSchemaRef) GetExternal() string {
	return r.External
}

func (r *DiscoveryEngineSchemaRef) SetExternal(external string) {
	r.External = external
}

func (r *DiscoveryEngineSchemaRef) ValidateExternal(external string) error {
	identity := &DiscoveryEngineSchemaIdentity{}
	return identity.FromExternal(external)
}

func (r *DiscoveryEngineSchemaRef) ParseExternalToIdentity(external string) (any, error) {
	identity := &DiscoveryEngineSchemaIdentity{}
	if err := identity.FromExternal(external); err != nil {
		return nil, err
	}
	return identity, nil
}

func (r *DiscoveryEngineSchemaRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
