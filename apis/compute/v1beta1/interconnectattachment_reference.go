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
	"fmt"

	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ComputeInterconnectAttachmentRef{}
var ComputeInterconnectAttachmentGVK = GroupVersion.WithKind("ComputeInterconnectAttachment")

type ComputeInterconnectAttachmentRef struct {
	// The value of an externally managed ComputeInterconnectAttachment resource.
	// Should be in the format "https://www.googleapis.com/compute/{{version}}/projects/{{projectId}}/regions/{{region}}/interconnectAttachments/{{name}}" or "projects/{{projectId}}/regions/{{region}}/interconnectAttachments/{{name}}"
	External string `json:"external,omitempty"`

	// The name of a ComputeInterconnectAttachment resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeInterconnectAttachment resource.
	Namespace string `json:"namespace,omitempty"`
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
}

func (r *ComputeInterconnectAttachmentRef) ValidateExternal(ref string) error {
	id := &ComputeInterconnectAttachmentIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ComputeInterconnectAttachmentRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		_, err := ParseComputeInterconnectAttachmentExternal(r.External)
		if err != nil {
			return err
		}
		external := common.FixStaleComputeExternalFormat(r.External)
		r.External = external
		return nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = defaultNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeInterconnectAttachmentGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return fmt.Errorf("reading referenced %s %s: %w", ComputeInterconnectAttachmentGVK, key, err)
	}

	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef != "" {
		r.External = actualExternalRef
		return nil
	}

	// Get external from status.selfLink. This ensures backward compatibility for TF/DCL-based resources that lack status.externalRef.
	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil {
		return fmt.Errorf("reading status.selfLink: %w", err)
	}
	if selfLink == "" {
		return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
	}

	external := common.FixStaleComputeExternalFormat(selfLink)
	r.External = external
	return nil
}
