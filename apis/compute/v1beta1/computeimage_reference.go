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

var ComputeImageGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeImage",
}

var _ refs.Ref = &ComputeImageRef{}

// ComputeImageRef is a reference to a ComputeImage.
type ComputeImageRef struct {
	// A reference to an externally managed ComputeImage resource.
	// Should be in the format "projects/{{project}}/global/images/{{name}}" or "projects/{{project}}/global/images/family/{{family}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeImage resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeImage resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeImageRef{}, &ComputeImage{})
}

func (r *ComputeImageRef) GetGVK() schema.GroupVersionKind {
	return ComputeImageGVK
}

func (r *ComputeImageRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeImageRef) GetExternal() string {
	return r.External
}

func (r *ComputeImageRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeImageRef) ValidateExternal(ref string) error {
	id := &ComputeImageIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeImageRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeImageIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeImageRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*ComputeImage](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeImageSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
