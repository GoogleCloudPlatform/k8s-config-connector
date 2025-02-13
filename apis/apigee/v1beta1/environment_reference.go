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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.ExternalNormalizer = &EnvironmentRef{}

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
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ApigeeEnvironmentGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, err := ParseEnvironmentExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(ApigeeEnvironmentGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ApigeeEnvironmentGVK, key, err)
	}

	/* TODO: Use status.externalRef once direct controller is implemented
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	*/

	// TODO: Use status.externalRef once direct controller is implemented.
	// For now, we can try to build it using spec fields.
	// Build OrganizationRef
	orgName, _, err := unstructured.NestedString(u.Object, "spec", "apigeeOrganizationRef", "name")
	if err != nil {
		return "", fmt.Errorf("reading spec.apigeeOrganizationRef.name: %w", err)
	}
	orgNamespace, _, err := unstructured.NestedString(u.Object, "spec", "apigeeOrganizationRef", "namespace")
	if err != nil {
		return "", fmt.Errorf("reading spec.apigeeOrganizationRef.namespace: %w", err)
	}
	orgExternal, _, err := unstructured.NestedString(u.Object, "spec", "apigeeOrganizationRef", "external")
	if err != nil {
		return "", fmt.Errorf("reading spec.apigeeOrganizationRef.external: %w", err)
	}
	// Normalize OrganizationRef
	orgRef := OrganizationRef{
		Name:      orgName,
		Namespace: orgNamespace,
		External:  orgExternal,
	}
	orgID, err := orgRef.NormalizedExternal(ctx, reader, otherNamespace)
	if err != nil {
		return "", fmt.Errorf("failed to normalize org ref: %w", err)
	}
	if orgID == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	// Build EnvironmentID
	resourceID, _, err := unstructured.NestedString(u.Object, "spec", "resourceID")
	if err != nil {
		return "", fmt.Errorf("reading spec.resourceID: %w", err)
	}
	metadataName, _, err := unstructured.NestedString(u.Object, "metadata", "name")
	if err != nil {
		return "", fmt.Errorf("reading metadata.name: %w", err)
	}
	envID := resourceID
	if envID == "" {
		envID = metadataName
	}
	if envID == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	// Build Environment external ref format
	r.External = orgID + "/environments/" + envID

	return r.External, nil
}
