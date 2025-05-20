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
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NOTE: If references a KMSKey, use `KMSKeyRef_OneOf` instead!
// todo: use unexported variable kmsCryptoKeyRef to avoid referencing to the KMS crypto key.

type KMSCryptoKeyRef struct {
	// The `name` of a `KMSCryptoKey` resource.
	Name string `json:"name,omitempty"`

	// The `namespace` of a `KMSCryptoKey` resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal will resolve a KMSCryptoKeyRef to external.
func (ref *KMSCryptoKeyRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if ref == nil {
		return "", nil
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = otherNamespace
	}

	// Fetch object from k8s cluster to construct the external form
	kmsKey := &unstructured.Unstructured{}
	kmsKey.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "kms.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "KMSCryptoKey",
	})
	if err := reader.Get(ctx, key, kmsKey); err != nil {
		if apierrors.IsNotFound(err) {
			return "", fmt.Errorf("referenced KMSCryptoKey %v not found", key)
		}
		return "", fmt.Errorf("error reading referenced KMSCryptoKey %v: %w", key, err)
	}

	kmsKeyResourceID, err := GetResourceID(kmsKey)
	if err != nil {
		return "", err
	}

	kmsRing, err := ResolveKeyRingForObject(ctx, reader, kmsKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", kmsRing.ProjectID, kmsRing.Location, kmsRing.ResourceID, kmsKeyResourceID), nil
}

type KMSKeyRingRef struct {
	//  If provided must be in the format `projects/[kms_project_id]/locations/[region]/keyRings/[key_ring_id]`.
	External string `json:"external,omitempty"`
	// The `name` field of a `KMSKeyRing` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `KMSKeyRing` resource.
	Namespace string `json:"namespace,omitempty"`
}

type KMSKeyRing struct {
	Ref        *KMSKeyRingRef
	ProjectID  string
	ResourceID string
	Location   string
}

// ResolveKMSKeyRingRef will resolve a KMSKeyRingRef to a KMSKeyRing.
func ResolveKMSKeyRingRef(ctx context.Context, reader client.Reader, src client.Object, ref *KMSKeyRingRef) (*KMSKeyRing, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on KMSKeyRingRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on KMSKeyRingRef")
	}

	// External should be in the `projects/[kms_project_id]/locations/[region]/keyRings/[key_ring_id]` format
	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "keyRings" {
			ref = &KMSKeyRingRef{
				External: fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", tokens[1], tokens[3], tokens[5]),
			}
			return &KMSKeyRing{Ref: ref, ResourceID: tokens[5], Location: tokens[3]}, nil
		}
		return nil, fmt.Errorf("format of KMSKeyRingRef external=%q was not known (use projects/[kms_project_id]/locations/[region]/keyRings/[key_ring_id])", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	// Fetch object from k8s cluster to construct the external form
	kmsKeyRing := &unstructured.Unstructured{}
	kmsKeyRing.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "kms.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "KMSKeyRing",
	})
	if err := reader.Get(ctx, key, kmsKeyRing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced KMSKeyRing %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced KMSKeyRing %v: %w", key, err)
	}

	kmsKeyResourceID, err := GetResourceID(kmsKeyRing)
	if err != nil {
		return nil, err
	}

	projectID, err := ResolveProjectID(ctx, reader, kmsKeyRing)
	if err != nil {
		return nil, err
	}

	location, err := GetLocation(kmsKeyRing)
	if err != nil {
		return nil, err
	}

	ref = &KMSKeyRingRef{
		External: fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", projectID, location, kmsKeyResourceID),
	}

	return &KMSKeyRing{Ref: ref, ProjectID: projectID, ResourceID: kmsKeyResourceID, Location: location}, nil
}

func ResolveKeyRingForObject(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (*KMSKeyRing, error) {
	keyRingRefExternal, _, err := unstructured.NestedString(obj.Object, "spec", "keyRingRef", "external")
	if err != nil {
		return nil, fmt.Errorf("error fetching secretRef.external %w", err)
	}
	if keyRingRefExternal != "" {
		return ResolveKMSKeyRingRef(ctx, reader, obj, &KMSKeyRingRef{External: keyRingRefExternal})
	}

	keyRingRefName, _, err := unstructured.NestedString(obj.Object, "spec", "keyRingRef", "name")
	if err != nil {
		return nil, fmt.Errorf("error fetching keyRingRef.name %w", err)
	}
	if keyRingRefName != "" {
		keyRingRefNamespace, _, err := unstructured.NestedString(obj.Object, "spec", "keyRingRef", "namespace")
		if err != nil {
			return nil, fmt.Errorf("error fetching keyRingRef.namespace %w", err)
		}

		keyRingRef := KMSKeyRingRef{
			Name:      keyRingRefName,
			Namespace: keyRingRefNamespace,
		}
		if keyRingRef.Namespace == "" {
			keyRingRef.Namespace = obj.GetNamespace()
		}

		return ResolveKMSKeyRingRef(ctx, reader, obj, &keyRingRef)
	}

	return nil, fmt.Errorf("cannot find keyRingRef for %v %v/%v", obj.GetKind(), obj.GetNamespace(), obj.GetName())
}
