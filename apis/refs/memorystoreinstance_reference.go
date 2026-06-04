// Copyright 2025 Google LLC
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

package refs

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var memorystoreInstanceGVK = schema.GroupVersionKind{
	Group:   "memorystore.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "MemorystoreInstance",
}

var _ refs.Ref = &MemorystoreInstanceRef{}

// MemorystoreInstanceRef is a reference to a MemorystoreInstance.
type MemorystoreInstanceRef struct {
	// A reference to an externally managed MemorystoreInstance resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}".
	External string `json:"external,omitempty"`

	// The name of a MemorystoreInstance resource.
	Name string `json:"name,omitempty"`

	// The namespace of a MemorystoreInstance resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&MemorystoreInstanceRef{})
}

func (r *MemorystoreInstanceRef) GetGVK() schema.GroupVersionKind {
	return memorystoreInstanceGVK
}

func (r *MemorystoreInstanceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *MemorystoreInstanceRef) GetExternal() string {
	return r.External
}

func (r *MemorystoreInstanceRef) SetExternal(ref string) {
	r.External = ref
}

func (r *MemorystoreInstanceRef) ValidateExternal(ref string) error {
	id := &MemorystoreInstanceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *MemorystoreInstanceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &MemorystoreInstanceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *MemorystoreInstanceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := MemorystoreInstance_IdentityFromSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
