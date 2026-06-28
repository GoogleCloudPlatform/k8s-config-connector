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

var _ refs.Ref = &ComputeAutoscalerRef{}

// ComputeAutoscalerRef is a reference to a ComputeAutoscaler.
type ComputeAutoscalerRef struct {
	// A reference to an externally managed ComputeAutoscaler resource.
	// Should be in the format "projects/{{projectID}}/zones/{{zone}}/autoscalers/{{autoscaler}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeAutoscaler resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeAutoscaler resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeAutoscalerRef{}, &ComputeAutoscaler{})
}

func (r *ComputeAutoscalerRef) GetGVK() schema.GroupVersionKind {
	return ComputeAutoscalerGVK
}

func (r *ComputeAutoscalerRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeAutoscalerRef) GetExternal() string {
	return r.External
}

func (r *ComputeAutoscalerRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeAutoscalerRef) ValidateExternal(ref string) error {
	trimmedRef := refs.TrimComputeURIPrefix(ref)
	id := &ComputeAutoscalerIdentity{}
	if err := id.FromExternal(trimmedRef); err != nil {
		return err
	}
	return nil
}

func (r *ComputeAutoscalerRef) ParseExternalToIdentity() (identity.Identity, error) {
	trimmedRef := refs.TrimComputeURIPrefix(r.External)
	id := &ComputeAutoscalerIdentity{}
	if err := id.FromExternal(trimmedRef); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeAutoscalerRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = refs.TrimComputeURIPrefix(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		// Get external from status.selfLink. This ensures backward compatibility for TF/DCL-based resources that lack status.externalRef.
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			trimmed := refs.TrimComputeURIPrefix(selfLink)
			id := &ComputeAutoscalerIdentity{}
			if err := id.FromExternal(trimmed); err == nil {
				return trimmed
			}
		}

		obj, err := common.ToStructuredType[*ComputeAutoscaler](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeAutoscalerSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
