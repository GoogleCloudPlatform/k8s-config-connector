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
	"errors"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &SecretRef{}

// SecretRef defines the resource reference to SecretManagerSecret, which "External" field
// holds the GCP identifier for the KRM object.
type SecretRef struct {
	// A reference to an externally managed SecretManagerSecret resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/secrets/{{secretID}}".
	External string `json:"external,omitempty"`

	// The name of a SecretManagerSecret resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SecretManagerSecret resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on SecretManagerSecret.
// If the "External" is given in the other resource's spec.SecretManagerSecretRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual SecretManagerSecret object from the cluster.
func (r *SecretRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", SecretManagerSecretGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, err := ParseSecretExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(SecretManagerSecretGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", SecretManagerSecretGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err1 := unstructured.NestedString(u.Object, "status", "externalRef")
	if err1 != nil {
		err1 = fmt.Errorf("SecretManagerSecret `status.externalRef` not configured: %w", err1)
		// Backward compatible to Terraform/DCL based resource, which does not have status.externalRef.
		var err2 error
		actualExternalRef, _, err2 = unstructured.NestedString(u.Object, "status", "name")
		if err2 != nil {
			return "", errors.Join(err1, err2)
		}
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}

func ParseSecretExternal(external string) (*SecretIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("missing external value")
	}
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "secrets" {
		return nil, fmt.Errorf("format of SecretManagerSecret external=%q was not known (use projects/{{projectId}}/secrets/{{secretID}})", external)
	}
	return &SecretIdentity{
		parent: &SecretParent{ProjectID: tokens[1]},
		id:     tokens[3],
	}, nil
}
