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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &BigtableInstanceRef{}

type BigtableInstanceRef struct {
	// A reference to an externally managed BigtableInstance resource. Should be in the format "projects/{{projectID}}/instances/{{instanceID}}"
	External string `json:"external,omitempty"`

	// The name of a BigtableInstance resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigtableInstance resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&BigtableInstanceRef{})
}

func (r *BigtableInstanceRef) GetGVK() schema.GroupVersionKind {
	return BigtableInstanceGVK
}

func (r *BigtableInstanceRef) GetNamespacedName() client.ObjectKey {
	return client.ObjectKey{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *BigtableInstanceRef) GetExternal() string {
	return r.External
}

func (r *BigtableInstanceRef) SetExternal(external string) {
	r.External = external
}

func (r *BigtableInstanceRef) ValidateExternal(external string) error {
	return (&BigtableInstanceIdentity{}).FromExternal(external)
}

func (r *BigtableInstanceRef) ParseExternalToIdentity() (identity.IdentityV2, error) {
	id := &BigtableInstanceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *BigtableInstanceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		if id, err := getIdentityFromBigtableInstanceSpec(ctx, reader, u); err == nil {
			return id.String()
		}
		return ""
	})
}
