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

var _ refs.Ref = &KMSKeyHandleRef{}

// KMSKeyHandleRef is a reference to a Google Cloud KMSKeyHandle.
type KMSKeyHandleRef struct {
	// A reference to an externally managed KMSKeyHandle resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/keyHandles/{{keyHandleID}}".
	External string `json:"external,omitempty"`

	// The name of a KMSKeyHandle resource.
	Name string `json:"name,omitempty"`

	// The namespace of a KMSKeyHandle resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&KMSKeyHandleRef{})
}

func (r *KMSKeyHandleRef) GetGVK() schema.GroupVersionKind {
	return KMSKeyHandleGVK
}

func (r *KMSKeyHandleRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *KMSKeyHandleRef) GetExternal() string {
	return r.External
}

func (r *KMSKeyHandleRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *KMSKeyHandleRef) ValidateExternal(ref string) error {
	id := &KMSKeyHandleIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *KMSKeyHandleRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &KMSKeyHandleIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *KMSKeyHandleRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*KMSKeyHandle](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromKMSKeyHandleSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// kmsKeyHandleRef defines the resource reference to KMSKeyHandle.
// This unexported/legacy type is retained for backward compatibility where embedded inside KMSKeyRef_OneOf.
type kmsKeyHandleRef struct {
	// The name of a KMSKeyHandle resource.
	Name string `json:"name,omitempty"`

	// The namespace of a KMSKeyHandle resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on KMSKeyHandle.
// The "Name" and "Namespace" will be used to query the actual KMSKeyHandle object from the cluster.
func (r *kmsKeyHandleRef) normalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.Name == "" {
		return "", fmt.Errorf("name` of `KMSKeyHandle` must be set")
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(KMSKeyHandleGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", KMSKeyHandleGVK, key, err)
	}
	// Use status.observedState.kmsKey instead of status.externalRef as the external value of autoKey
	// status.externalRef: projects/${projectId}/locations/us-central1/keyHandles/1a1a1a-222b-3cc3-d444-e555ee555555
	// status.observedState.kmsKey: projects/${key_project}/locations/us-central1/keyRings/autokey/cryptoKeys/${projectNumber}-compute-disk-${generated-id}
	kmsKey, _, err := unstructured.NestedString(u.Object, "status", "observedState", "kmsKey")
	if err != nil {
		return "", fmt.Errorf("reading status.observedState.kmsKey: %w", err)
	}
	if kmsKey == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	return kmsKey, nil
}
