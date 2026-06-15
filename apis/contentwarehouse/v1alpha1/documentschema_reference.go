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

var _ refs.Ref = &ContentWarehouseSchemaRef{}

// ContentWarehouseSchemaRef defines the resource reference to ContentWarehouseSchema, which "External" field
// holds the GCP identifier for the KRM object.
type ContentWarehouseSchemaRef struct {
	// A reference to an externally managed ContentWarehouseSchema resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/documentSchemas/{{document_schema}}"
	External string `json:"external,omitempty"`

	// The name of a ContentWarehouseSchema resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ContentWarehouseSchema resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ContentWarehouseSchemaRef{}, &ContentWarehouseSchema{})
}

func (r *ContentWarehouseSchemaRef) GetGVK() schema.GroupVersionKind {
	return ContentWarehouseSchemaGVK
}

func (r *ContentWarehouseSchemaRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ContentWarehouseSchemaRef) GetExternal() string {
	return r.External
}

func (r *ContentWarehouseSchemaRef) SetExternal(external string) {
	r.External = external
}

func (r *ContentWarehouseSchemaRef) ValidateExternal(ref string) error {
	identity := &ContentWarehouseSchemaIdentity{}
	return identity.FromExternal(ref)
}
func (r *ContentWarehouseSchemaRef) ParseExternalToIdentity() (identity.Identity, error) {
	identity := &ContentWarehouseSchemaIdentity{}
	err := identity.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return identity, nil
}

func (r *ContentWarehouseSchemaRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromContentWarehouseSchemaSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, fallback)
}
func (r *ContentWarehouseSchemaRef) Resolve(ctx context.Context, reader client.Reader, src client.Object) (string, error) {
	if r == nil {
		return "", nil
	}
	err := r.Normalize(ctx, reader, src.GetNamespace())
	if err != nil {
		return "", err
	}
	return r.External, nil
}
