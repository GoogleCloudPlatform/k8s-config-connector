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

var _ refsv1beta1.ExternalNormalizer = &VPCAccessConnectorRef{}

// VPCAccessConnectorRef defines the resource reference to VPCAccessConnector, which "External" field
// holds the GCP identifier for the KRM object.
type VPCAccessConnectorRef struct {
	// A reference to an externally managed VPCAccessConnector resource.
	// Should be in the format `projects/{project_id}/locations/{location}/connectors/{connector_id}`
	External string `json:"external,omitempty"`

	// The name of a VPCAccessConnector resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VPCAccessConnector resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *VPCAccessConnectorRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", VPCAccessConnectorGVK.Kind)
	}
	// From given External
	if r.External != "" {
		if _, _, _, err := ParseVPCAccessConnectorExternal(r.External); err != nil {
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
	u.SetGroupVersionKind(VPCAccessConnectorGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", VPCAccessConnectorGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		// It's possible the referenced VPCAccessConnector is a legacy one and doesn't
		// have `status.externalRef`.
		ready, err := isResourceReady(u)
		if err != nil {
			return "", fmt.Errorf("checking if referenced %s %s is ready: %w", VPCAccessConnectorGVK, key, err)
		}
		if !ready {
			return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
		}
		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return "", err
		}
		location, err := refsv1beta1.GetLocation(u)
		if err != nil {
			return "", err
		}
		resourceID, err := refsv1beta1.GetResourceID(u)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("projects/%s/locations/%s/connectors/%s", projectID, location, resourceID), nil
	}

	r.External = actualExternalRef
	return r.External, nil
}

func ParseVPCAccessConnectorExternal(external string) (projectID, location, name string, err error) {
	parts := strings.Split(external, "/")
	if len(parts) != 6 {
		return "", "", "", fmt.Errorf("unexpected format for external %q, expected projects/{project_id}/locations/{location}/connectors/{name}", external)
	}
	if parts[0] != "projects" {
		return "", "", "", fmt.Errorf("unexpected format for external %q, expected projects/{project_id}/locations/{location}/connectors/{name}", external)
	}
	if parts[2] != "locations" {
		return "", "", "", fmt.Errorf("unexpected format for external %q, expected projects/{project_id}/locations/{location}/connectors/{name}", external)
	}
	if parts[4] != "connectors" {
		return "", "", "", fmt.Errorf("unexpected format for external %q, expected projects/{project_id}/locations/{location}/connectors/{name}", external)
	}
	return parts[1], parts[3], parts[5], nil
}

// isResourceReady checks if the given unstructured resource is ready.
// This is a simplified check and might need to be adjusted based on the actual resource's status conditions.
func isResourceReady(u *unstructured.Unstructured) (bool, error) {
	conditions, found, err := unstructured.NestedSlice(u.Object, "status", "conditions")
	if err != nil {
		return false, fmt.Errorf("getting status.conditions: %w", err)
	}
	if !found {
		return false, nil
	}
	for _, c := range conditions {
		condition, ok := c.(map[string]interface{})
		if !ok {
			continue
		}
		t, ok := condition["type"].(string)
		if !ok {
			continue
		}
		status, ok := condition["status"].(string)
		if !ok {
			continue
		}
		if t == "Ready" && status == "True" {
			return true, nil
		}
	}
	return false, nil
}
