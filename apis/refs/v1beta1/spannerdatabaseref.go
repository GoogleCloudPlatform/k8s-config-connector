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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SpannerDatabaseRef struct {
	/* The Spanner Database selfLink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `SpannerDatabase` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `SpannerDatabase` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type SpannerDatabase struct {
	ProjectID  string
	InstanceID string
	DatabaseID string
}

func (s *SpannerDatabase) String() string {
	return "projects/" + s.ProjectID + "/instances/" + s.InstanceID + "/databases/" + s.DatabaseID
}

func ResolveSpannerDatabaseRef(ctx context.Context, reader client.Reader, obj client.Object, ref *SpannerDatabaseRef) (*SpannerDatabase, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on databaseRef")
	}
	if ref.External != "" && ref.Name != "" {
		return nil, fmt.Errorf("cannot specify both spec.databaseRef.name and spec.databaseRef.external")
	}

	if ref.External != "" {
		// External must be in form `projects/<projectID>/instances/<instanceName>/databases/<databaseName>`.
		// see https://cloud.google.com/spanner/docs/reference/rest#rest-resource:-v1.projects.instances.databases
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "databases" {
			return &SpannerDatabase{
				ProjectID:  tokens[1],
				InstanceID: tokens[3],
				DatabaseID: tokens[5],
			}, nil
		}
		return nil, fmt.Errorf("format of sqlinstance external=%q was not known (use projects/<projectID>/instances/<instanceName>/databases/<databaseName>)", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	spannerdatabase := &unstructured.Unstructured{}
	spannerdatabase.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "spanner.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "SpannerDatabase",
	})
	if err := reader.Get(ctx, key, spannerdatabase); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced Spanner databse %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced Spanner databse %v: %w", key, err)
	}
	resource, err := k8s.NewResource(spannerdatabase)
	if err != nil {
		return nil, fmt.Errorf("error converting unstructured to resource: %w", err)
	}
	if !k8s.IsResourceReady(resource) {
		return nil, k8s.NewReferenceNotReadyError(spannerdatabase.GroupVersionKind(), key)
	}

	resourceID, _, err := unstructured.NestedString(spannerdatabase.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from Spanner Database %s/%s: %w", spannerdatabase.GetNamespace(), spannerdatabase.GetName(), err)
	}
	if resourceID == "" {
		resourceID = spannerdatabase.GetName()
	}

	projectID, err := ResolveProjectID(ctx, reader, spannerdatabase)
	if err != nil {
		return nil, err
	}

	instanceID, err := ResolveSpannerInstanceID(ctx, reader, spannerdatabase)
	if err != nil {
		return nil, err
	}

	return &SpannerDatabase{
		ProjectID:  projectID,
		InstanceID: instanceID,
		DatabaseID: resourceID,
	}, nil
}
