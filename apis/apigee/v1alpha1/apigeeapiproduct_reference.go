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

var _ refs.Ref = &ApigeeApiProductRef{}

// ApigeeApiProductRef is a reference to an ApigeeApiProduct.
type ApigeeApiProductRef struct {
	// A reference to an externally managed ApigeeApiProduct resource.
	// Should be in the format "organizations/{{organizationID}}/apiproducts/{{apiproductID}}".
	External string `json:"external,omitempty"`

	// The name of a ApigeeApiProduct resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ApigeeApiProduct resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ApigeeApiProductRef{})
}

func (r *ApigeeApiProductRef) GetGVK() schema.GroupVersionKind {
	return ApigeeApiProductGVK
}

func (r *ApigeeApiProductRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ApigeeApiProductRef) GetExternal() string {
	return r.External
}

func (r *ApigeeApiProductRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ApigeeApiProductRef) ValidateExternal(ref string) error {
	id := &ApigeeApiProductIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ApigeeApiProductRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ApigeeApiProductIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ApigeeApiProductRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		obj := &ApigeeApiProduct{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
			return ""
		}
		id, err := getIdentityFromApigeeApiProductSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
