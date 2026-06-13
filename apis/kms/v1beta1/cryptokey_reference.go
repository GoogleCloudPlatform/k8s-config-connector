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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var KMSCryptoKeyGVK = GroupVersion.WithKind("KMSCryptoKey")

var _ refs.Ref = &KMSCryptoKeyRef{}

// KMSCryptoKeyRef is a reference to a KMSCryptoKey.
type KMSCryptoKeyRef struct {
	// A reference to an externally managed KMSCryptoKey resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptokey}}".
	External string `json:"external,omitempty"`

	// The name of a KMSCryptoKey resource.
	Name string `json:"name,omitempty"`

	// The namespace of a KMSCryptoKey resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&KMSCryptoKeyRef{})
}

func (r *KMSCryptoKeyRef) GetGVK() schema.GroupVersionKind {
	return KMSCryptoKeyGVK
}

func (r *KMSCryptoKeyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *KMSCryptoKeyRef) GetExternal() string {
	return r.External
}

func (r *KMSCryptoKeyRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *KMSCryptoKeyRef) ValidateExternal(ref string) error {
	id := &KMSCryptoKeyIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *KMSCryptoKeyRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &KMSCryptoKeyIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *KMSCryptoKeyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*KMSCryptoKey](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromKMSCryptoKeySpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// kmsCryptoKeyRef defines the resource reference to KMSCryptoKey.
// This unexported/legacy type is retained for backward compatibility where embedded inside KMSKeyRef_OneOf.
type kmsCryptoKeyRef struct {
	// The `name` of a `KMSCryptoKey` resource.
	Name string `json:"name,omitempty"`
	// The `namespace` of a `KMSCryptoKey` resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on KMSCryptoKeyRef.
// The "Name" and "Namespace" will be used to query the actual KMSCryptoKeyRef object from the cluster.
func (r *kmsCryptoKeyRef) normalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.Name == "" {
		return "", fmt.Errorf("`name` of `KMSCryptoKey` must be set")
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(KMSCryptoKeyGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", KMSCryptoKeyGVK, key, err)
	}

	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef != "" {
		return actualExternalRef, nil
	}

	// Backward compatible for resources still managed by the legacy controller and without the status.externalRef
	// Use status.selfLink as the external value of cryptokey
	// status.selfLink: projects/${projectId}/locations/us-central1/keyRings/kmscryptokey-${uniqueId}/cryptoKeys/kmscryptokey-${uniqueId}
	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil {
		return "", fmt.Errorf("reading status.selfLink: %w", err)
	}
	if selfLink == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	return selfLink, nil
}
