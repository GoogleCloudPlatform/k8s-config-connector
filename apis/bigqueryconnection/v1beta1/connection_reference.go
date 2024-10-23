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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/google/uuid"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewBigQueryConnectionConnectionRef builds a BigQueryConnectionConnectionRef from the ConfigConnector BigQueryConnectionConnection object.
func NewBigQueryConnectionConnectionRef(ctx context.Context, reader client.Reader, obj *BigQueryConnectionConnection) (*BigQueryConnectionConnectionRef, error) {
	id := &BigQueryConnectionConnectionRef{}

	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	// Get location
	location := obj.Spec.Location

	// Get desired service-generated ID from spec
	desiredServiceID := direct.ValueOf(obj.Spec.ResourceID)
	if desiredServiceID != "" {
		if _, err := uuid.Parse(desiredServiceID); err != nil {
			return nil, fmt.Errorf("spec.resourceID should be in a UUID format, got %s ", desiredServiceID)
		}
	}

	// Get externalReference
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		tokens := strings.Split(externalRef, "/")

		if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "connections" {
			return nil, fmt.Errorf("externalRef should be projects/<project>/locations/<location>/connections/<Connection>, got %s", externalRef)
		}
		id.parent = "projects/" + tokens[1] + "/locations/" + tokens[3]

		// Validate spec parent and resourceID field if the resource is already reconcilied with a GCP Connection resource.
		if tokens[1] != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s",
				tokens[1], projectID)
		}
		if tokens[3] != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s",
				tokens[3], location)
		}
		if desiredServiceID != "" && tokens[5] != desiredServiceID {
			// Service generated ID shall not be reset in the same BigQueryConnectionConnection.
			// TODO: what if multiple BigQueryConnectionConnection points to the same GCP Connection?
			return nil, fmt.Errorf("cannot reset `spec.resourceID` to %s, since it has already acquired the Connection %s",
				desiredServiceID, tokens[5])
		}
		id.External = externalRef
		return id, nil
	}
	id.parent = "projects/" + projectID + "/locations/" + location
	if desiredServiceID != "" {
		id.External = id.parent + "/connections/" + desiredServiceID
	}
	return id, nil
}

var _ refsv1beta1.ExternalNormalizer = &BigQueryConnectionConnectionRef{}

// BigQueryConnectionConnectionRef defines the resource reference to BigQueryConnectionConnection, which "External" field
// holds the GCP identifier for the KRM object.
type BigQueryConnectionConnectionRef struct {
	// A reference to an externally managed BigQueryConnectionConnection resource.
	// Should be in the format `projects/<projectID>/locations/<location>/connections/<connectionID>`.
	External string `json:"external,omitempty"`

	// The `name` of a `BigQueryConnectionConnection` resource.
	Name string `json:"name,omitempty"`
	// The `namespace` of a `BigQueryConnectionConnection` resource.
	Namespace string `json:"namespace,omitempty"`

	parent string
}

func (r *BigQueryConnectionConnectionRef) Parent() (string, error) {
	if r.parent != "" {
		return r.parent, nil
	}
	if r.External != "" {
		r.External = strings.TrimPrefix(r.External, "/")
		tokens := strings.Split(r.External, "/")
		if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "connections" {
			return "", fmt.Errorf("format of BigQueryConnectionConnection external=%q was not known (use projects/<projectId>/locations/<location>/connections/<connectionID>)", r.External)
		}
		r.parent = "projects/" + tokens[1] + "/locations/" + tokens[3]
		return r.parent, nil
	}
	return "", fmt.Errorf("BigQueryConnectionConnectionRef not normalized to External form or not created from `New()`")
}

// NormalizedExternal provision the "External" value.
// If the "External" comes from the ConfigConnector object, it has to acquire or reconcile with the GCP resource already.
func (r *BigQueryConnectionConnectionRef) NormalizedExternal(ctx context.Context, reader client.Reader, othernamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", BigQueryConnectionConnectionGVK.Kind)
	}
	if r.External != "" {
		r.External = strings.TrimPrefix(r.External, "/")
		tokens := strings.Split(r.External, "/")
		if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "connections" {
			return "", fmt.Errorf("format of BigQueryConnectionConnection external=%q was not known (use projects/<projectId>/locations/<location>/connections/<connectionID>)", r.External)
		}
		return r.External, nil
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(BigQueryConnectionConnectionGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", BigQueryConnectionConnectionGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", fmt.Errorf("BigQueryConnectionConnection is not ready yet.")
	}
	r.External = actualExternalRef
	return r.External, nil
}
