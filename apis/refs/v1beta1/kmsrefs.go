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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NOTE: If references a KMSKey, use `KMSKeyRef_OneOf` instead!
// todo: use unexported variable kmsCryptoKeyRef to avoid referencing to the KMS crypto key.

type KMSCryptoKeyRef struct {
	// A reference to an externally managed KMSCryptoKey.
	// Should be in the format `projects/[kms_project_id]/locations/[region]/keyRings/[key_ring_id]/cryptoKeys/[key]`.
	External string `json:"external,omitempty"`

	// The `name` of a `KMSCryptoKey` resource.
	Name string `json:"name,omitempty"`

	// The `namespace` of a `KMSCryptoKey` resource.
	Namespace string `json:"namespace,omitempty"`
}

// ResolveKMSCryptoKeyRef will resolve a KMSCryptoKeyRef to a KMSCryptoKey.
func ResolveKMSCryptoKeyRef(ctx context.Context, reader client.Reader, src client.Object, ref *KMSCryptoKeyRef) (*KMSCryptoKeyRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on KMSCryptoKeyRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on KMSCryptoKeyRef")
	}

	// External should be in the `projects/[kms_project_id]/locations/[region]/keyRings/[key_ring_id]/cryptoKeys/[key]` format
	if ref.External != "" {
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "keyRings" && tokens[6] == "cryptoKeys" {
			ref = &KMSCryptoKeyRef{
				External: fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", tokens[1], tokens[3], tokens[5], tokens[7]),
			}
			return ref, nil
		}
		return nil, fmt.Errorf("format of KMSCryptoKeyRef external=%q was not known (use projects/[kms_project_id]/locations/[region]/keyRings/[key_ring_id]/cryptoKeys/[key])", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
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
			return nil, fmt.Errorf("referenced KMSCryptoKey %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced KMSCryptoKey %v: %w", key, err)
	}

	// Use status.selfLink as the external value of cryptokey
	selfLink, _, err := unstructured.NestedString(kmsKey.Object, "status", "selfLink")
	if err != nil {
		return nil, fmt.Errorf("reading status.selfLink: %w", err)
	}
	if selfLink == "" {
		return nil, k8s.NewReferenceNotReadyError(kmsKey.GroupVersionKind(), key)
	}
	ref.External = selfLink
	return ref, nil
}
