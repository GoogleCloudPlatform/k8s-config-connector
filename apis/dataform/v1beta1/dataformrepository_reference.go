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

var _ refs.Ref = &DataformRepositoryRef{}

// DataformRepositoryRef is a reference to a DataformRepository.
type DataformRepositoryRef struct {
	// A reference to an externally managed DataformRepository resource.
	// Should be in the format "projects/{project}/locations/{location}/repositories/{repository}".
	External string `json:"external,omitempty"`

	// The name of a DataformRepository resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DataformRepository resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&DataformRepositoryRef{}, &DataformRepository{})
}

func (r *DataformRepositoryRef) GetGVK() schema.GroupVersionKind {
	return DataformRepositoryGVK
}

func (r *DataformRepositoryRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DataformRepositoryRef) GetExternal() string {
	return r.External
}

func (r *DataformRepositoryRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *DataformRepositoryRef) ValidateExternal(ref string) error {
	id := &DataformRepositoryIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DataformRepositoryRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DataformRepositoryIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DataformRepositoryRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromDataformRepositorySpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
