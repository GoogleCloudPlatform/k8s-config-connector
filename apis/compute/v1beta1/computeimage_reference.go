// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not not use this file except in compliance with the License.
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
	ComputeImageGVK = GroupVersion.WithKind("ComputeImage")
)
var _ refsv1beta1.Ref = &ComputeImageRef{}

// ComputeImageRef is a reference to a ComputeImage resource.
type ComputeImageRef struct {
	// A reference to an externally managed ComputeImage resource.
	// Should be in the format "projects/{{project}}/global/images/{{name}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeImage resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeImage resource.
	Namespace string `json:"namespace,omitempty"`
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
}

func (r *ComputeImageRef) ValidateExternal(ref string) error {
	id := &ComputeImageIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *ComputeImageRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	err := refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
	if r.GetExternal() != "" {
		// TODO: validate the external for legacy selfLink.
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
	// Get external from status.externalRef. This is the most trustworthy place.
	externalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return fmt.Errorf("reading status.externalRef: %w", err)
	}
	if externalRef == "" {
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
