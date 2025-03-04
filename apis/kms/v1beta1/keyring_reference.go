// Copyright 2025 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &KMSCryptoKeyRef{}

// KMSKeyRingRef defines the resource reference to KMSKeyRing, which "External" field
// holds the GCP identifier for the KRM object.
type KMSKeyRingRef struct {
	// A reference to an externally managed KMSKeyRing.
	// Should be in the format `projects/{{kms_project_id}}/locations/{{region}}/keyRings/{{key_ring_id}}`.
	External string `json:"external,omitempty"`

	// The `name` of a `KMSKeyRing` resource.
	Name string `json:"name,omitempty"`
	// The `namespace` of a `KMSKeyRing` resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on KMSKeyRingRef.
// If the "External" is given in the other resource's spec.KMSKeyRingRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual KMSKeyRingRef object from the cluster.
func (r *KMSKeyRingRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", KMSKeyRingGVK.Kind)
	}
	// From given External
	// External should be in the `projects/{{kms_project_id}}/locations/{{region}}/keyRings/{{key_ring_id}}` format
	if r.External != "" {
		if _, err := ParseKMSKeyRingExternal(r.External); err != nil {
			return "", err
		}
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(KMSKeyRingGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", KMSKeyRingGVK, key, err)
	}

	// todo: use externalRef for resource that managed by direct controller
	resourceID, err := refsv1beta1.GetResourceID(u)
	if err != nil {
		return "", err
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return "", err
	}

	location, err := refsv1beta1.GetLocation(u)
	if err != nil {
		return "", err
	}

	r.External = fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", projectID, location, resourceID)
	return r.External, nil
}

func ParseKMSKeyRingExternal(external string) (*KMSKeyRingIdentity, error) {
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	// projects/{{kms_project_id}}/locations/{{region}}/keyRings/{{key_ring_id}}
	if len(tokens) == 6 {
		p, err := parent.ParseProjectAndLocationParent(strings.Join(tokens[:len(tokens)-2], "/"))
		if err != nil {
			return nil, err
		}
		if tokens[len(tokens)-2] == "keyRings" {
			return &KMSKeyRingIdentity{parent: p, id: tokens[len(tokens)-1]}, nil
		}
	}
	return nil, fmt.Errorf("format of KMSKeyRing external=%q was not known (use projects/{{kms_project_id}}/locations/{{region}}/keyRings/{{key_ring_id}})", external)

}
