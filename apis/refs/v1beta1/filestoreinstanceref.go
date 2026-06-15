// Copyright 2026 Google LLC
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

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FilestoreInstanceRef struct {
	/* The FilestoreInstance selfLink/name, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `FilestoreInstance` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `FilestoreInstance` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type FilestoreInstance struct {
	ProjectID string
	Location  string
	Name      string
}

func (s *FilestoreInstance) String() string {
	return "projects/" + s.ProjectID + "/locations/" + s.Location + "/instances/" + s.Name
}

func ResolveFilestoreInstanceRef(ctx context.Context, reader client.Reader, obj client.Object, ref *FilestoreInstanceRef) (*FilestoreInstance, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on filestoreInstanceRef")
	}
	if ref.External != "" && ref.Name != "" {
		return nil, fmt.Errorf("cannot specify both name and external on filestoreInstanceRef")
	}

	if ref.External != "" {
		// External should be in the `projects/[projectID]/locations/[location]/instances/[name]` format.
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
			return &FilestoreInstance{
				ProjectID: tokens[1],
				Location:  tokens[3],
				Name:      tokens[5],
			}, nil
		}
		return nil, fmt.Errorf("format of filestoreinstance external=%q was not known (use projects/<projectId>/locations/<location>/instances/<instanceName>)", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	instance := &unstructured.Unstructured{}
	instance.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "filestore.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "FilestoreInstance",
	})
	if err := reader.Get(ctx, key, instance); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced FilestoreInstance %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced FilestoreInstance %v: %w", key, err)
	}

	resourceID, _, err := unstructured.NestedString(instance.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from FilestoreInstance %s/%s: %w", instance.GetNamespace(), instance.GetName(), err)
	}
	if resourceID == "" {
		resourceID = instance.GetName()
	}

	location, _, err := unstructured.NestedString(instance.Object, "spec", "location")
	if err != nil {
		return nil, fmt.Errorf("reading spec.location from FilestoreInstance %s/%s: %w", instance.GetNamespace(), instance.GetName(), err)
	}

	projectID, err := ResolveProjectID(ctx, reader, instance)
	if err != nil {
		return nil, err
	}

	return &FilestoreInstance{
		ProjectID: projectID,
		Location:  location,
		Name:      resourceID,
	}, nil
}
