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

var _ refs.Ref = &BlockchainNodeEngineBlockchainNodeRef{}

// BlockchainNodeEngineBlockchainNodeRef defines the resource reference to BlockchainNodeEngineBlockchainNode, which "External" field
// holds the GCP identifier for the KRM object.
type BlockchainNodeEngineBlockchainNodeRef struct {
	// A reference to an externally managed BlockchainNodeEngineBlockchainNode resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/blockchainNodes/{{blockchainNodeID}}".
	External string `json:"external,omitempty"`

	// The name of a BlockchainNodeEngineBlockchainNode resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BlockchainNodeEngineBlockchainNode resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&BlockchainNodeEngineBlockchainNodeRef{})
}

func (r *BlockchainNodeEngineBlockchainNodeRef) GetGVK() schema.GroupVersionKind {
	return BlockchainNodeEngineBlockchainNodeGVK
}

func (r *BlockchainNodeEngineBlockchainNodeRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BlockchainNodeEngineBlockchainNodeRef) GetExternal() string {
	return r.External
}

func (r *BlockchainNodeEngineBlockchainNodeRef) SetExternal(ref string) {
	r.External = ref
}

func (r *BlockchainNodeEngineBlockchainNodeRef) ValidateExternal(ref string) error {
	id := &BlockchainNodeEngineBlockchainNodeIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *BlockchainNodeEngineBlockchainNodeRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &BlockchainNodeEngineBlockchainNodeIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *BlockchainNodeEngineBlockchainNodeRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromBlockchainNodeEngineBlockchainNodeSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
