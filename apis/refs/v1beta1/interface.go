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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ExternalNormalizer is a deprecated interface.
// Deprecated: Use Ref instead.
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

type ExternalRef interface {
	Ref
	// ParseExternalToIdentity parses the External field to an Identity.
	// Normalize should be called first to ensure that External is populated.
	ParseExternalToIdentity() (identity.Identity, error)
}

// Normalize is a general-purpose reference resolver that can be used to
// implement the "Normalize" interface method for most Ref types.
// Use Normalize when the referenced resource is a direct controller or standard
// resource where the canonical, fully-resolved GCP resource URI/identifier is
// guaranteed to be populated inside "status.externalRef".
func Normalize(ctx context.Context, reader client.Reader, ref Ref, defaultNamespace string) error {
	return NormalizeWithFallback(ctx, reader, ref, defaultNamespace, nil)
}

// NormalizeWithFallback extends Normalize by allowing a fallback function to be provided
// for obtaining the external reference if it is not found in status.externalRef.
//
// Use NormalizeWithFallback only for older/legacy resources (such as Terraform or DCL resources)
// that store the resolved external reference in a different status field (e.g., status.id, status.selfLink, or status.observedState).
//
// Crucial Guidelines for the Fallback Function:
// 1. The fallback function MUST ONLY read from status fields (e.g. status.selfLink, status.observedState) to determine if the resource is fully reconciled and ready.
// 2. The fallback function MUST NOT read from the resource spec (e.g., spec.resourceID, spec.location, spec.projectRef). Reading from spec is unsafe because it would successfully resolve identities for resources that are not yet created or ready in GCP, causing subsequent controller reconciliations to fail or behave incorrectly.
func NormalizeWithFallback(ctx context.Context, reader client.Reader, ref Ref, defaultNamespace string, fallback func(u *unstructured.Unstructured) string) error {
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
		if externalRef == "" && fallback != nil {
			externalRef = fallback(u)
		}
		if externalRef == "" {
			return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
		}
		ref.SetExternal(externalRef)
	}

	return ref.ValidateExternal(ref.GetExternal())
}
