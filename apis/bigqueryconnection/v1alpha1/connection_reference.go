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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func New(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (*BigQueryConnectionConnectionRef, error) {
	id := &BigQueryConnectionConnectionRef{}
	obj := &BigQueryConnectionConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

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
	// var id *BigQueryConnectionConnectionIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		tokens := strings.Split(externalRef, "/")

		if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "connections" {
			return nil, fmt.Errorf("externalRef should be projects/<project>/locations/<location>/connections/<Connection>, got %s", externalRef)
		}
		id.parent = "projects/" + tokens[1] + "/locations/" + tokens[3]
		id.serviceGeneratedID = tokens[5]

		if tokens[1] != projectID {
			return nil, fmt.Errorf("BigQueryConnectionConnection %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), tokens[1], projectID)
		}
		if tokens[3] != location {
			return nil, fmt.Errorf("BigQueryConnectionConnection %s/%s has spec.location changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), tokens[3], location)
		}

		if desiredServiceID != "" && tokens[5] != desiredServiceID {
			// Service generated ID shall not be reset in the same BigQueryConnectionConnection.
			// TODO: what if multiple BigQueryConnectionConnection points to the same GCP Connection?
			return nil, fmt.Errorf("cannot reset `spec.resourceID` to %s, since it has already acquired the Connection %s",
				desiredServiceID, tokens[5])
		}
	} else {
		id.parent = "projects/" + projectID + "/locations/" + location
		id.serviceGeneratedID = desiredServiceID
	}
	if id.serviceGeneratedID != "" {
		id.External = id.parent + "/connections/" + id.serviceGeneratedID
	}
	return id, nil
}

var _ refsv1beta1.Resolver = &BigQueryConnectionConnectionRef{}

type BigQueryConnectionConnectionRef struct {
	// A reference to an externally managed BigQueryConnectionConnection resource.
	// Should be in the format `projects/<projectID>/locations/<location>/connections/<connectionID>`.
	External string `json:"external,omitempty"`

	// The `name` of a `BigQueryConnectionConnection` resource.
	Name string `json:"name,omitempty"`
	// The `namespace` of a `BigQueryConnectionConnection` resource.
	Namespace string `json:"namespace,omitempty"`

	parent             string
	serviceGeneratedID string
}

func (r *BigQueryConnectionConnectionRef) ValidateExternal() error {
	if r.External == "" {
		return fmt.Errorf("external not specified")
	}
	r.External = strings.TrimPrefix(r.External, "/")
	tokens := strings.Split(r.External, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "connections" {
		return fmt.Errorf("format of BigQueryConnectionConnection external=%q was not known (use projects/<projectId>/locations/<location>/connections/<connectionID>)", r.External)
	}
	return nil
}

func (r *BigQueryConnectionConnectionRef) Resolve(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (string, error) {
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", fmt.Errorf("BigQueryConnectionConnection is not ready yet.")
	}
	return actualExternalRef, nil
}

func (r *BigQueryConnectionConnectionRef) Parent() string {
	return r.parent
}
