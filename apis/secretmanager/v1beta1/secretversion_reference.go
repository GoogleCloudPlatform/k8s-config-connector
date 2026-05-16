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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &SecretVersionRef{}

// SecretVersionRef defines the resource reference to SecretManagerSecretVersion, which "External" field
// holds the GCP identifier for the KRM object.
type SecretVersionRef struct {
	// A reference to an externally managed SecretManagerSecretVersion resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/secretversions/{{secretversionID}}".
	External string `json:"external,omitempty"`

	// The name of a SecretManagerSecretVersion resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SecretManagerSecretVersion resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on SecretManagerSecretVersion.
// If the "External" is given in the other resource's spec.SecretVersionRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual SecretManagerSecretVersion object from the cluster.
func (r *SecretVersionRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", SecretManagerSecretVersionGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, err := ParseSecretVersionExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(SecretManagerSecretVersionGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", SecretManagerSecretVersionGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil || actualExternalRef == "" {
		// Backward compatible to Terraform/DCL based resource, which does not have status.externalRef.
		actualExternalRef, _, err = unstructured.NestedString(u.Object, "status", "name")
		if err != nil {
			return "", err
		}
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}

// ParseSecretVersionExternal parses an external reference to a Secret Manager secret version.
// Supports both global secret versions (projects/{{project}}/secrets/{{secretID}}/versions/{{versionID}})
// and regional secret versions (projects/{{project}}/locations/{{location}}/secrets/{{secretID}}/versions/{{versionID}})
func ParseSecretVersionExternal(external string) (*SecretVersionIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("missing external value")
	}
	external = strings.TrimPrefix(external, "/")
	tokens := strings.Split(external, "/")

	// Check for global secret version format: projects/{{projectId}}/secrets/{{secretID}}/versions/{{versionID}}
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "secrets" && tokens[4] == "versions" {
		secretParent := &SecretParent{
			ProjectID: tokens[1],
			Location:  "global",
		}
		secretIdentity := &SecretIdentity{
			parent: secretParent,
			id:     tokens[3],
		}
		return &SecretVersionIdentity{
			parent: secretIdentity,
			id:     tokens[5],
		}, nil
	}

	// Check for regional secret version format: projects/{{projectId}}/locations/{{location}}/secrets/{{secretID}}/versions/{{versionID}}
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" &&
		tokens[4] == "secrets" && tokens[6] == "versions" {
		secretParent := &SecretParent{
			ProjectID: tokens[1],
			Location:  tokens[3],
		}
		secretIdentity := &SecretIdentity{
			parent: secretParent,
			id:     tokens[5],
		}
		return &SecretVersionIdentity{
			parent: secretIdentity,
			id:     tokens[7],
		}, nil
	}

	return nil, fmt.Errorf("format of SecretManagerSecretVersion external=%q was not known (use projects/{{projectId}}/secrets/{{secretID}}/versions/{{versionID}} or projects/{{projectId}}/locations/{{location}}/secrets/{{secretID}}/versions/{{versionID}})", external)
}
