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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &VertexAIExampleStoreRef{}

// VertexAIExampleStoreRef defines the resource reference to VertexAIExampleStore.
type VertexAIExampleStoreRef struct {
	// A reference to an externally managed VertexAIExampleStore resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/exampleStores/{{examplestore}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAIExampleStore resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIExampleStore resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&VertexAIExampleStoreRef{})
}

func (r *VertexAIExampleStoreRef) GetGVK() schema.GroupVersionKind {
	return VertexAIExampleStoreGVK
}

func (r *VertexAIExampleStoreRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAIExampleStoreRef) GetExternal() string {
	return r.External
}

func (r *VertexAIExampleStoreRef) SetExternal(external string) {
	r.External = external
}

func (r *VertexAIExampleStoreRef) ValidateExternal(ref string) error {
	id := &VertexAIExampleStoreIdentity{}
	return id.FromExternal(ref)
}

func (r *VertexAIExampleStoreRef) ParseExternalToIdentity() (identity.Identity, error) {
	if r.External == "" {
		return nil, fmt.Errorf("external is empty")
	}
	id := &VertexAIExampleStoreIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAIExampleStoreRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		id, err := getIdentityFromVertexAIExampleStoreSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, fallback)
}
