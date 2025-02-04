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

package v1alpha1

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

var _ refsv1beta1.ExternalNormalizer = &EnvironmentRef{}

var EnvironmentGVK = GroupVersion.WithKind("ApigeeEnvironment") // todo acpana house this in environment_types.go eventually

// EnvironmentRef defines the resource reference to ApigeeEnvironment, which "External" field
// holds the GCP identifier for the KRM object.
type EnvironmentRef struct {
	// A reference to an externally managed ApigeeEnvironment resource.
	// Should be in the format "organizations/{{organizationID}}/environments/{{environmentID}}".
	External string `json:"external,omitempty"`

	// The name of a ApigeeEnvironment resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ApigeeEnvironment resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on ApigeeEnvironment.
// If the "External" is given in the other resource's spec.ApigeeEnvironmentRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ApigeeEnvironment object from the cluster.
func (r *EnvironmentRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", EnvironmentGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := ParseEnvironmentExternalRef(r.External); err != nil {
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
	u.SetGroupVersionKind(EnvironmentGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", EnvironmentGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, found, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if !found {
		// todo acpana careful to deal with this see:
		// https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/3619
		// for more context.
		return "", fmt.Errorf("status.externalRef is not found")
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}

// todo acpana this should be in a _identity.go file eventually
func ParseEnvironmentExternalRef(external string) (org string, environment string, err error) {
	tokens := strings.Split(external, "/")

	if len(tokens) != 4 || tokens[0] != "organizations" || tokens[2] != "environments" {
		return "", "", fmt.Errorf("external should be organizations/{{organizationID}}/environments/{{environmentID}}, got: %s", external)
	}

	return tokens[1], tokens[3], nil
}
