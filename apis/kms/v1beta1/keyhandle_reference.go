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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &kmsKeyHandleRef{}

// kmsKeyHandleRef defines the resource reference to KMSKeyHandle
type kmsKeyHandleRef struct {
	// The name of a KMSKeyHandle resource.
	Name string `json:"name,omitempty"`

	// The namespace of a KMSKeyHandle resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on KMSKeyHandle.
// The "Name" and "Namespace" will be used to query the actual KMSKeyHandle object from the cluster.
func (r *kmsKeyHandleRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.Name == "" {
		return "", fmt.Errorf("use KMS autokey requires referring to the Config Connector `KMSKeyHanle` object. please provide the `name` of your  `KMSKeyHanle`.")
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
	// Use status.observedState.kmsKey instead of status.externalRef as the external resourceID of autoKey
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
