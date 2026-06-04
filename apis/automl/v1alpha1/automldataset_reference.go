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

var _ refs.Ref = &AutoMLDatasetRef{}

// AutoMLDatasetRef defines the resource reference to AutoMLDataset.
type AutoMLDatasetRef struct {
	// A reference to an externally managed AutoMLDataset resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/datasets/{{dataset}}".
	External string `json:"external,omitempty"`

	// The name of an AutoMLDataset resource.
	Name string `json:"name,omitempty"`

	// The namespace of an AutoMLDataset resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&AutoMLDatasetRef{})
}

func (r *AutoMLDatasetRef) GetGVK() schema.GroupVersionKind {
	return AutoMLDatasetGVK
}

func (r *AutoMLDatasetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *AutoMLDatasetRef) GetExternal() string {
	return r.External
}

func (r *AutoMLDatasetRef) SetExternal(external string) {
	r.External = external
}

func (r *AutoMLDatasetRef) ValidateExternal(external string) error {
	id := &AutoMLDatasetIdentity{}
	return id.FromExternal(external)
}

func (r *AutoMLDatasetRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &AutoMLDatasetIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *AutoMLDatasetRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		id, err := getIdentityFromAutoMLDatasetSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
