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

var _ refs.Ref = &APIHubAttributeRef{}

// APIHubAttributeRef is a reference to a GCP APIHubAttribute.
type APIHubAttributeRef struct {
	// A reference to an externally managed APIHubAttribute resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/attributes/{{attribute}}"
	External string `json:"external,omitempty"`

	// The name of an APIHubAttribute resource.
	Name string `json:"name,omitempty"`

	// The namespace of an APIHubAttribute resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&APIHubAttributeRef{})
}

func (r *APIHubAttributeRef) GetGVK() schema.GroupVersionKind {
	return APIHubAttributeGVK
}

func (r *APIHubAttributeRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *APIHubAttributeRef) GetExternal() string {
	return r.External
}

func (r *APIHubAttributeRef) SetExternal(external string) {
	r.External = external
}

func (r *APIHubAttributeRef) ValidateExternal(ref string) error {
	id := &APIHubAttributeIdentity{}
	return id.FromExternal(ref)
}

func (r *APIHubAttributeRef) ParseExternalToIdentity() (any, error) {
	id := &APIHubAttributeIdentity{}
	err := id.FromExternal(r.External)
	return id, err
}

func (r *APIHubAttributeRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
