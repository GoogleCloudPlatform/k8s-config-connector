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

var _ refs.Ref = &DataCatalogPolicyTagRef{}

// DataCatalogPolicyTagRef is a reference to a DataCatalogPolicyTag resource.
type DataCatalogPolicyTagRef struct {
	// A reference to an externally managed DataCatalogPolicyTag resource.
	// Should be in the format "projects/{project}/locations/{location}/taxonomies/{taxonomy}/policyTags/{policyTag}".
	External string `json:"external,omitempty"`

	// The name of a DataCatalogPolicyTag resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DataCatalogPolicyTag resource.
	Namespace string `json:"namespace,omitempty"`
}

// PolicyTagRef is an alias for DataCatalogPolicyTagRef for backward compatibility.
// +k8s:deepcopy-gen=false
type PolicyTagRef = DataCatalogPolicyTagRef

func init() {
	refs.Register(&DataCatalogPolicyTagRef{})
}

func (r *DataCatalogPolicyTagRef) GetGVK() schema.GroupVersionKind {
	return DataCatalogPolicyTagGVK
}

func (r *DataCatalogPolicyTagRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DataCatalogPolicyTagRef) GetExternal() string {
	return r.External
}

func (r *DataCatalogPolicyTagRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DataCatalogPolicyTagRef) ValidateExternal(ref string) error {
	id := &DataCatalogPolicyTagIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DataCatalogPolicyTagRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DataCatalogPolicyTagIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DataCatalogPolicyTagRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromDataCatalogPolicyTagSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
