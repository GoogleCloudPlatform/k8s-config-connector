// Copyright 2025 Google LLC
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
	TagsTagValueGVK = GroupVersion.WithKind("TagsTagValue")
)

var _ refsv1beta1.Ref = &TagsTagValueRef{}

// TagsTagValueRef is a reference to a TagsTagValue resource.
// +kcc:ref=TagsTagValue
type TagsTagValueRef struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`

	// Allowed value: string of the format `tagValues/{{value}}`,
	// where {{value}} is the `name` field of a `TagsTagValue` resource.
	External string `json:"external,omitempty"`
}

// GetGVK returns the GroupVersionKind of the referenced resource.
func (r *TagsTagValueRef) GetGVK() schema.GroupVersionKind {
	return TagsTagValueGVK
}

// GetNamespacedName returns the NamespacedName of the referenced resource.
func (r *TagsTagValueRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

// GetExternal returns the external reference string of the referenced resource.
func (r *TagsTagValueRef) GetExternal() string {
	return r.External
}

// SetExternal sets the external reference string of the referenced resource.
func (r *TagsTagValueRef) SetExternal(external string) {
	r.External = external
}

// ValidateExternal validates the external reference string of the referenced resource.
func (r *TagsTagValueRef) ValidateExternal(external string) error {
	if strings.HasPrefix(external, "tagValues/") {
		return fmt.Errorf("format of TagsTagValue external=%q was not known, missing prefix tagValues/", external)
	}
	return nil
}

func (r *TagsTagValueRef) GetExternalFromCustomFields() []string {
	return []string{"status", "name"}
}

// Normalize resolves the reference to an external resource string.
func (r *TagsTagValueRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.GetExternal() == "" {
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
			// Try to get external from legacy fields if status.externalRef is not set.
			fieldPaths := r.GetExternalFromCustomFields()
			if fieldPaths == nil {
				return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
			}
			if externalRef, _, err = unstructured.NestedString(u.Object, fieldPaths...); err != nil {
				return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
			}
		}
		r.SetExternal(externalRef)
	}

	return r.ValidateExternal(r.GetExternal())
}
