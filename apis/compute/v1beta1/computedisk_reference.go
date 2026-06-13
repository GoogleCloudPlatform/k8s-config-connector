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
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ComputeDiskGVK = GroupVersion.WithKind("ComputeDisk")
)

var _ refsv1beta1.Ref = &ComputeDiskRef{}

// ComputeDiskRef is a reference to a ComputeDisk resource.
type ComputeDiskRef struct {
	// A reference to an externally managed ComputeDisk resource.
	// Should be in the format "projects/{{project}}/zones/{{zone}}/disks/{{name}}" or "projects/{{project}}/regions/{{region}}/disks/{{name}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeDisk resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeDisk resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&ComputeDiskRef{})
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
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ComputeDiskRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	err := refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
	if r.GetExternal() != "" {
		return err
	}

	key := r.GetNamespacedName()
	if key.Namespace == "" {
		key.Namespace = defaultNamespace
	}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(r.GetGVK())
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return fmt.Errorf("reading referenced %s %s: %w", r.GetGVK(), key, err)
	}
	// `status.externalRef` is the preferred field to store externalRef. Since ComputeDisk is not yet migrated to
	// direct controller, so this field does not exist. We will use `.status.selfLink` instead.
	externalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil || externalRef == "" {
		if externalRef, _, err = unstructured.NestedString(u.Object, "status", "selfLink"); err != nil {
			return err
		}
		if externalRef == "" {
			return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
		}
		externalRef = strings.TrimPrefix(externalRef, "https://www.googleapis.com/compute/v1/")
	}
	r.SetExternal(externalRef)
	return r.ValidateExternal(externalRef)
}
