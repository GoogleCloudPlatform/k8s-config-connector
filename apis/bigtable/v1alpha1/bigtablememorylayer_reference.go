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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &BigtableMemoryLayerRef{}

// BigtableMemoryLayerRef is a reference to a GCP BigtableMemoryLayer.
type BigtableMemoryLayerRef struct {
	// A reference to an externally managed BigtableMemoryLayer resource. Should be in the format "projects/{{projectID}}/instances/{{instanceID}}/clusters/{{clusterID}}/memoryLayer".
	External string `json:"external,omitempty"`

	// The name of a BigtableMemoryLayer resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigtableMemoryLayer resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&BigtableMemoryLayerRef{}, &BigtableMemoryLayer{})
}

func (r *BigtableMemoryLayerRef) GetGVK() schema.GroupVersionKind {
	return BigtableMemoryLayerGVK
}

func (r *BigtableMemoryLayerRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BigtableMemoryLayerRef) GetExternal() string {
	return r.External
}

func (r *BigtableMemoryLayerRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *BigtableMemoryLayerRef) ValidateExternal(ref string) error {
	id := &BigtableMemoryLayerIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *BigtableMemoryLayerRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &BigtableMemoryLayerIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *BigtableMemoryLayerRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

// NormalizedExternal is kept for backwards compatibility of existing callers.
func (r *BigtableMemoryLayerRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
