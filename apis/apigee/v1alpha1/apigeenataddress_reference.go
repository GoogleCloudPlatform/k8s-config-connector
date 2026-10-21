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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ApigeeNATAddressRef{}

// ApigeeNATAddressRef is a reference to a GCP ApigeeNATAddress.
type ApigeeNATAddressRef struct {
	// A reference to an externally managed ApigeeNATAddress resource.
	// Should be in the format "organizations/{{organizationID}}/instances/{{instanceID}}/natAddresses/{{natAddressID}}".
	External string `json:"external,omitempty"`

	// The name of an ApigeeNATAddress resource.
	Name string `json:"name,omitempty"`

	// The namespace of an ApigeeNATAddress resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ApigeeNATAddressRef{}, &ApigeeNATAddress{})
}

func (r *ApigeeNATAddressRef) GetGVK() schema.GroupVersionKind {
	return ApigeeNATAddressGVK
}

func (r *ApigeeNATAddressRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ApigeeNATAddressRef) GetExternal() string {
	return r.External
}

func (r *ApigeeNATAddressRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ApigeeNATAddressRef) ValidateExternal(ref string) error {
	id := &ApigeeNATAddressIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ApigeeNATAddressRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ApigeeNATAddressIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ApigeeNATAddressRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*ApigeeNATAddress](u)
		if err != nil {
			return ""
		}
		ready := false
		for _, cond := range obj.Status.Conditions {
			if cond.Type == "Ready" && cond.Status == "True" {
				ready = true
				break
			}
		}
		if !ready {
			return ""
		}
		id, err := getIdentityFromApigeeNATAddressSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return id.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
