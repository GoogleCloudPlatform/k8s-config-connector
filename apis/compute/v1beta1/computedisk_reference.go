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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ComputeDiskGVK = GroupVersion.WithKind("ComputeDisk")

var _ refs.Ref = &ComputeDiskRef{}

// ComputeDiskRef is a reference to a ComputeDisk.
type ComputeDiskRef struct {
	// A reference to an externally managed ComputeDisk resource. Should be in the format "projects/{{projectID}}/zones/{{zone}}/disks/{{diskID}}" or "projects/{{projectID}}/regions/{{region}}/disks/{{diskID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeDisk resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeDisk resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeDiskRef{})
}

func (r *ComputeDiskRef) GetGVK() schema.GroupVersionKind {
	return ComputeDiskGVK
}

func (r *ComputeDiskRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeDiskRef) GetExternal() string {
	return r.External
}

func (r *ComputeDiskRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeDiskRef) ValidateExternal(ref string) error {
	id := &ComputeDiskIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeDiskRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeDiskIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeDiskRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*ComputeDisk](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeDiskSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
