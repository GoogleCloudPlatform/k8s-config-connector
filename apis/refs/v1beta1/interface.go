// Copyright 2024 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ExternalNormalizer interface {
	// NormalizedExternal expects the implemented struct has a "External" field, and this function
	// assigns a value to the "External" field if it is empty.
	// In general, it retrieves the corresponding ConfigConnector object from the cluster, using
	// the `status.externalRef` or other field as the "External" value
	NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error)
}

type Ref interface {
	// GetGVK returns the schema.GroupVersionKind of the reference type
	GetGVK() schema.GroupVersionKind

	// GetNamespacedName returns the types.NamespacedName of a given reference
	GetNamespacedName() types.NamespacedName

	// GetExternal returns the external reference string (if set) of the reference.
	GetExternal() string

	// SetExternal sets the external reference string for a reference.
	SetExternal(ref string)

	// ValidateExternal returns nil if the given external reference string has a valid format for the reference.
	// Otherwise, it returns an error.
	ValidateExternal(ref string) error

	// Normalize ensures the "External" reference (in string format) is
	// set for a given Ref, and that it has the correct format.
	//
	// If "External" is already set, the format will be validated.
	//
	// If "External" is not set, the NamespacedName will be used to query the
	// referenced object from the K8s API and fetch it's external reference
	// value. If "Namespace" is not specified in the reference, the
	// `defaultNamespace“ will be used instead.
	Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error
}

// Normalize is a general-purpose reference resolver that can be used to
// implement the "Normalize" interface method for most Ref types.
func Normalize(ctx context.Context, reader client.Reader, ref Ref, defaultNamespace string) error {
	if ref.GetExternal() == "" {
		key := ref.GetNamespacedName()
		if key.Namespace == "" {
			key.Namespace = defaultNamespace
		}
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(ref.GetGVK())
		if err := reader.Get(ctx, key, u); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
			}
			return fmt.Errorf("reading referenced %s %s: %w", ref.GetGVK(), key, err)
		}
		// Get external from status.externalRef. This is the most trustworthy place.
		externalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
		if err != nil {
			return fmt.Errorf("reading status.externalRef: %w", err)
		}
		if externalRef == "" {
			return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
		}
		ref.SetExternal(externalRef)
	}

	return ref.ValidateExternal(ref.GetExternal())
}
