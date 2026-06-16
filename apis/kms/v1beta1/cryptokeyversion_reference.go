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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &KMSCryptoKeyVersionRef{}

var KMSCryptoKeyVersionGVK = schema.GroupVersionKind{
	Group:   "kms.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "KMSCryptoKeyVersion",
}

// KMSCryptoKeyVersionRef is a reference to a KMSCryptoKeyVersion.
type KMSCryptoKeyVersionRef struct {
	// A reference to an externally managed cryptoKeyVersion.
	// Should be in the format `projects/{{kms_project_id}}/locations/{{region}}/keyRings/{{key_ring_id}}/cryptoKeys/{{key}}/cryptoKeyVersions/{{version}}`.
	External string `json:"external,omitempty"`

	// The name of a KMSCryptoKeyVersion resource.
	Name string `json:"name,omitempty"`

	// The namespace of a KMSCryptoKeyVersion resource.
	Namespace string `json:"namespace,omitempty"`
}

// Deprecated: NormalizedExternal is kept for backwards compatibility with existing controllers/callers.
// Prefer implementing and using refs.Ref on newer resources instead.
func (r *KMSCryptoKeyVersionRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", KMSCryptoKeyVersionGVK.Kind)
	}

	// From given External
	if r.External != "" {
		return r.External, nil
	}

	// From the Config Connector object
	if r.Name == "" {
		return "", fmt.Errorf("either external or name must be specified on %s reference", KMSCryptoKeyVersionGVK.Kind)
	}

	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}

	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(KMSCryptoKeyVersionGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", KMSCryptoKeyVersionGVK.Kind, key, err)
	}

	// Get external from status.name. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "name")
	if err != nil {
		return "", fmt.Errorf("reading status.name: %w", err)
	}
	if actualExternalRef != "" {
		return actualExternalRef, nil
	}

	return "", fmt.Errorf("referenced %s %s does not have status.name", KMSCryptoKeyVersionGVK.Kind, key)
}
