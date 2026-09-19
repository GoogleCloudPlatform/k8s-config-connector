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
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &NetworkConnectivityServiceConnectionMapRef{}

// NetworkConnectivityServiceConnectionMapRef is a reference to a GCP NetworkConnectivityServiceConnectionMap.
type NetworkConnectivityServiceConnectionMapRef struct {
	// A reference to an externally managed NetworkConnectivityServiceConnectionMap resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/serviceConnectionMaps/{{serviceConnectionMap}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkConnectivityServiceConnectionMap resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkConnectivityServiceConnectionMap resource.
	Namespace string `json:"namespace,omitempty"`
}

// GetGVK returns the GroupVersionKind for NetworkConnectivityServiceConnectionMap
func (r *NetworkConnectivityServiceConnectionMapRef) GetGVK() schema.GroupVersionKind {
	return NetworkConnectivityServiceConnectionMapGVK
}

// GetNamespacedName returns the NamespacedName for NetworkConnectivityServiceConnectionMap
func (r *NetworkConnectivityServiceConnectionMapRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

// GetExternal returns the External string for NetworkConnectivityServiceConnectionMap
func (r *NetworkConnectivityServiceConnectionMapRef) GetExternal() string {
	return r.External
}

// SetExternal sets the External string for NetworkConnectivityServiceConnectionMap
func (r *NetworkConnectivityServiceConnectionMapRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

// ValidateExternal checks if the External string is valid
func (r *NetworkConnectivityServiceConnectionMapRef) ValidateExternal(external string) error {
	identity := &NetworkConnectivityServiceConnectionMapIdentity{}
	return identity.FromExternal(external)
}

// ParseExternalToIdentity parses the External string to an identity
func (r *NetworkConnectivityServiceConnectionMapRef) ParseExternalToIdentity() (any, error) {
	identity := &NetworkConnectivityServiceConnectionMapIdentity{}
	err := identity.FromExternal(r.External)
	return identity, err
}

// Normalize normalizes the reference
func (r *NetworkConnectivityServiceConnectionMapRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

func init() {
	refs.Register(&NetworkConnectivityServiceConnectionMapRef{}, &NetworkConnectivityServiceConnectionMap{})
}
