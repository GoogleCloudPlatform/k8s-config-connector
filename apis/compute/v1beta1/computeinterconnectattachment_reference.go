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
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ComputeInterconnectAttachmentRef{}

// ComputeInterconnectAttachmentRef is a reference to a GCP ComputeInterconnectAttachment.
type ComputeInterconnectAttachmentRef struct {
	// A reference to an externally managed ComputeInterconnectAttachment resource.
	// Should be in the format "projects/{{projectID}}/regions/{{region}}/interconnectAttachments/{{interconnectAttachmentID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeInterconnectAttachment resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeInterconnectAttachment resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeInterconnectAttachmentRef{}, &ComputeInterconnectAttachment{})
}

func (r *ComputeInterconnectAttachmentRef) GetGVK() schema.GroupVersionKind {
	return ComputeInterconnectAttachmentGVK
}

func (r *ComputeInterconnectAttachmentRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeInterconnectAttachmentRef) GetExternal() string {
	return r.External
}

func (r *ComputeInterconnectAttachmentRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeInterconnectAttachmentRef) ValidateExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	id := &ComputeInterconnectAttachmentIdentity{}
	if err := id.FromExternal(trimmedRef); err != nil {
		return err
	}
	return nil
}

func (r *ComputeInterconnectAttachmentRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeInterconnectAttachmentIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeInterconnectAttachmentRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = apirefs.TrimComputeURIPrefix(r.External)
	}

	fallback := func(u *unstructured.Unstructured) string {
		// Get external from status.selfLink. This ensures backward compatibility for TF/DCL-based resources that lack status.externalRef.
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			trimmed := apirefs.TrimComputeURIPrefix(selfLink)
			id := &ComputeInterconnectAttachmentIdentity{}
			if err := id.FromExternal(trimmed); err == nil {
				return trimmed
			}
		}

		obj, err := common.ToStructuredType[*ComputeInterconnectAttachment](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeInterconnectAttachmentSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
