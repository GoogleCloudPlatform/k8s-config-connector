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

var _ refs.Ref = &ComputeMachineImageRef{}

// ComputeMachineImageRef is a reference to a ComputeMachineImage.
type ComputeMachineImageRef struct {
	// A reference to an externally managed ComputeMachineImage resource.
	// Should be in the format "projects/{{projectID}}/global/machineImages/{{machineImage}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeMachineImage resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeMachineImage resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeMachineImageRef{}, &ComputeMachineImage{})
}

func (r *ComputeMachineImageRef) GetGVK() schema.GroupVersionKind {
	return ComputeMachineImageGVK
}

func (r *ComputeMachineImageRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeMachineImageRef) GetExternal() string {
	return r.External
}

func (r *ComputeMachineImageRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeMachineImageRef) ValidateExternal(ref string) error {
	trimmedRef := refs.TrimComputeURIPrefix(ref)
	id := &ComputeMachineImageIdentity{}
	if err := id.FromExternal(trimmedRef); err != nil {
		return err
	}
	return nil
}

func (r *ComputeMachineImageRef) ParseExternalToIdentity() (identity.Identity, error) {
	trimmedRef := refs.TrimComputeURIPrefix(r.External)
	id := &ComputeMachineImageIdentity{}
	if err := id.FromExternal(trimmedRef); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeMachineImageRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = refs.TrimComputeURIPrefix(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		// Get external from status.selfLink. This ensures backward compatibility for TF/DCL-based resources that lack status.externalRef.
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			trimmed := refs.TrimComputeURIPrefix(selfLink)
			id := &ComputeMachineImageIdentity{}
			if err := id.FromExternal(trimmed); err == nil {
				return trimmed
			}
		}

		obj, err := common.ToStructuredType[*ComputeMachineImage](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeMachineImageSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
