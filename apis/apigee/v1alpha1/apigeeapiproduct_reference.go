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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ApigeeAPIProductRef{}

// ApigeeAPIProductRef is a reference to an ApigeeAPIProduct.
type ApigeeAPIProductRef struct {
	// A reference to an externally managed ApigeeAPIProduct resource.
	// Should be in the format "organizations/{{organizationID}}/apiproducts/{{apiproductID}}".
	External string `json:"external,omitempty"`

	// The name of a ApigeeAPIProduct resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ApigeeAPIProduct resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ApigeeAPIProductRef{})
}

func (r *ApigeeAPIProductRef) GetGVK() schema.GroupVersionKind {
	return ApigeeAPIProductGVK
}

func (r *ApigeeAPIProductRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ApigeeAPIProductRef) GetExternal() string {
	return r.External
}

func (r *ApigeeAPIProductRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ApigeeAPIProductRef) ValidateExternal(ref string) error {
	id := &ApigeeAPIProductIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ApigeeAPIProductRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ApigeeAPIProductIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ApigeeAPIProductRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		obj := &ApigeeAPIProduct{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
			return ""
		}
		id, err := getIdentityFromApigeeAPIProductSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
