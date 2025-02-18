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

type SQLInstanceRef struct {
	/* The SQLInstance selfLink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `SQLInstance` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `SQLInstance` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type SQLInstance struct {
	ProjectID       string
	Location        string
	SQLInstanceName string
}

func (s *SQLInstance) String() string {
	return "projects/" + s.ProjectID + "locations/" + s.Location + "/instances/" + s.SQLInstanceName
}

func (s *SQLInstance) ConnectionName() string {
	return s.ProjectID + ":" + s.Location + ":" + s.SQLInstanceName
}

func ResolveSQLInstanceRef(ctx context.Context, reader client.Reader, obj client.Object, ref *SQLInstanceRef) (*SQLInstance, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on instanceRef")
	}
	if ref.External != "" && ref.Name != "" {
		return nil, fmt.Errorf("cannot specify both spec.instanceRef.name and spec.instanceRef.external")
	}

	if ref.External != "" {
		// External should be in the `projects/[projectID]/locations/[Location]/instances/[instanceName]` format.
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
			return &SQLInstance{
				ProjectID:       tokens[1],
				Location:        tokens[3],
				SQLInstanceName: tokens[5],
			}, nil
		}
		return nil, fmt.Errorf("format of sqlinstance external=%q was not known (use projects/<projectId>/locations/[Location]/instances/<instanceName>)", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	sqlinstance := &unstructured.Unstructured{}
	sqlinstance.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "sql.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "SQLInstance",
	})
	if err := reader.Get(ctx, key, sqlinstance); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced SQLInstance %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced SQLInstance %v: %w", key, err)
	}
	resource, err := k8s.NewResource(sqlinstance)
	if err != nil {
		return nil, fmt.Errorf("error converting unstructured to resource: %w", err)
	}
	if !k8s.IsResourceReady(resource) {
		return nil, k8s.NewReferenceNotReadyError(sqlinstance.GroupVersionKind(), key)
	}

	resourceID, _, err := unstructured.NestedString(sqlinstance.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from SQLInstance %s/%s: %w", sqlinstance.GetNamespace(), sqlinstance.GetName(), err)
	}
	if resourceID == "" {
		resourceID = sqlinstance.GetName()
	}

	location, _, err := unstructured.NestedString(sqlinstance.Object, "spec", "region")
	if err != nil {
		return nil, fmt.Errorf("reading spec.region from SQLInstance %s/%s: %w", sqlinstance.GetNamespace(), sqlinstance.GetName(), err)
	}

	projectID, err := ResolveProjectID(ctx, reader, sqlinstance)
	if err != nil {
		return nil, err
	}

	return &SQLInstance{
		ProjectID:       projectID,
		Location:        location,
		SQLInstanceName: resourceID,
	}, nil
}

func ResolveSQLInstanceID(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (string, error) {
	instanceRefExternal, _, _ := unstructured.NestedString(obj.Object, "spec", "instanceRef", "external")
	if instanceRefExternal != "" {
		instanceRef := &SQLInstanceRef{
			External: instanceRefExternal,
		}
		instance, err := ResolveSQLInstanceRef(ctx, reader, obj, instanceRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse instanceRef.external %q in %v %v/%v: %w", instanceRefExternal, obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return instance.SQLInstanceName, nil
	}

	instanceRefName, _, _ := unstructured.NestedString(obj.Object, "spec", "instanceRef", "name")
	if instanceRefName != "" {
		namespace, _, _ := unstructured.NestedString(obj.Object, "spec", "instanceRef", "namespace")
		instanceRef := &SQLInstanceRef{
			Name:      instanceRefName,
			Namespace: namespace,
		}
		if instanceRef.Namespace == "" {
			instanceRef.Namespace = obj.GetNamespace()
		}
		instance, err := ResolveSQLInstanceRef(ctx, reader, obj, instanceRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse instanceRef.name %q in %v %v/%v: %w", instanceRefName, obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return instance.SQLInstanceName, nil
	}

	return "", fmt.Errorf("cannot find instance id for %v %v/%v", obj.GetKind(), obj.GetNamespace(), obj.GetName())
}
