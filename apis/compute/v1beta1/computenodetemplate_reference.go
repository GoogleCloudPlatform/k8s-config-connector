// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ComputeNodeTemplateRef{}

func init() {
	refsv1beta1.Register(&ComputeNodeTemplateRef{})
}

// ComputeNodeTemplateRef is a reference to a ComputeNodeTemplate.
type ComputeNodeTemplateRef struct {
	// A reference to an externally managed ComputeNodeTemplate resource. Should be in the format "projects/{{projectID}}/regions/{{region}}/nodeTemplates/{{nodeTemplateID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeNodeTemplate resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeNodeTemplate resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *ComputeNodeTemplateRef) GetGVK() schema.GroupVersionKind {
	return ComputeNodeTemplateGVK
}

func (r *ComputeNodeTemplateRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeNodeTemplateRef) GetExternal() string {
	return r.External
}

func (r *ComputeNodeTemplateRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeNodeTemplateRef) ValidateExternal(ref string) error {
	id := &ComputeNodeTemplateIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeNodeTemplateRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeNodeTemplateIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeNodeTemplateRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = apirefs.TrimComputeURIPrefix(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		// Get external from status.selfLink or status.externalRef. This ensures backward compatibility for TF/DCL-based resources that lack status.externalRef.
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			trimmed := apirefs.TrimComputeURIPrefix(selfLink)
			id := &ComputeNodeTemplateIdentity{}
			if err := id.FromExternal(trimmed); err == nil {
				return trimmed
			}
		}

		externalRef, _, _ := unstructured.NestedString(u.Object, "status", "externalRef")
		if externalRef != "" {
			trimmed := apirefs.TrimComputeURIPrefix(externalRef)
			id := &ComputeNodeTemplateIdentity{}
			if err := id.FromExternal(trimmed); err == nil {
				return trimmed
			}
		}

		obj, err := common.ToStructuredType[*ComputeNodeTemplate](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeNodeTemplateSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
