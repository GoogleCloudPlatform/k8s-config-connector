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

var _ refsv1beta1.ExternalNormalizer = &OrganizationRef{}

// OrganizationRef defines the resource reference to ApigeeOrganization, which "External" field
// holds the GCP identifier for the KRM object.
type OrganizationRef struct {
	// A reference to an externally managed ApigeeOrganization resource.
	// Should be in the format "organizations/{{organizationID}}".
	External string `json:"external,omitempty"`

	// The name of a ApigeeOrganization resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ApigeeOrganization resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on ApigeeOrganization.
// If the "External" is given in the other resource's spec.ApigeeOrganizationRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ApigeeOrganization object from the cluster.
func (r *OrganizationRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ApigeeOrganizationGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, err := ParseOrganizationExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(ApigeeOrganizationGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ApigeeOrganizationGVK, key, err)
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
	// For now, we can use status.projectID.
	// BUT only construct the reference if the resource is ready
	resource, err := k8s.NewResource(u)
	if err != nil {
		return "", fmt.Errorf("error converting unstructured to resource: %w", err)
	}
	if !k8s.IsResourceReady(resource) {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}

	projectID, _, err := unstructured.NestedString(u.Object, "status", "projectId")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if projectID == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = "organizations/" + projectID

	return r.External, nil
}
