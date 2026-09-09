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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &VertexAIIndexRef{}
var VertexAIIndexGVK = schema.GroupVersionKind{Group: "vertexai.cnrm.cloud.google.com", Version: "v1beta1", Kind: "VertexAIIndex"}

// VertexAIIndexRef is a reference to a VertexAIIndex.
type VertexAIIndexRef struct {
	// A reference to an externally managed VertexAIIndex resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/indexes/{{indexID}}".
	External string `json:"external,omitempty"`

	/* NOTYET
	// The name of a VertexAIIndex resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIIndex resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

func init() {
	refs.Register(&VertexAIIndexRef{})
}

func (r *VertexAIIndexRef) GetGVK() schema.GroupVersionKind {
	return VertexAIIndexGVK
}

func (r *VertexAIIndexRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *VertexAIIndexRef) GetExternal() string {
	return r.External
}

func (r *VertexAIIndexRef) SetExternal(ref string) {
	r.External = ref
}

func (r *VertexAIIndexRef) ValidateExternal(ref string) error {
	id := &VertexAIIndexIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *VertexAIIndexRef) ParseExternalToIdentity() (identity.Identity, error) {
	if r.External == "" {
		return nil, fmt.Errorf("external is empty")
	}
	id := &VertexAIIndexIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAIIndexRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", VertexAIIndexGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}
