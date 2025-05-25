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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SQLDatabaseRef struct {
	/* The SQL Database name, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `SQLDatabase` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `SQLDatabase` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type SQLDatabase struct {
	ProjectID  string
	InstanceID string
	DatabaseID string
}

func (s *SQLDatabase) Name() string {
	return s.DatabaseID
}

func (s *SQLDatabase) String() string {
	if s.ProjectID != "" && s.InstanceID != "" {
		return "projects/" + s.ProjectID + "/instances/" + s.InstanceID + "/databases/" + s.DatabaseID
	}
	return s.DatabaseID
}

func ResolveSQLDatabaseRef(ctx context.Context, reader client.Reader, obj client.Object, ref *SQLDatabaseRef) (*SQLDatabase, error) {
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
		// External is the name of the sql database
		return &SQLDatabase{
			DatabaseID: ref.External,
		}, nil
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	sqldatabase := &unstructured.Unstructured{}
	sqldatabase.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "sql.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "SQLDatabase",
	})
	if err := reader.Get(ctx, key, sqldatabase); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced SQL databse %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced SQL databse %v: %w", key, err)
	}
	resource, err := k8s.NewResource(sqldatabase)
	if err != nil {
		return nil, fmt.Errorf("error converting unstructured to resource: %w", err)
	}
	if !k8s.IsResourceReady(resource) {
		return nil, k8s.NewReferenceNotReadyError(sqldatabase.GroupVersionKind(), key)
	}

	resourceID, _, err := unstructured.NestedString(sqldatabase.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from Sql Database %s/%s: %w", sqldatabase.GetNamespace(), sqldatabase.GetName(), err)
	}
	if resourceID == "" {
		resourceID = sqldatabase.GetName()
	}

	projectID, err := ResolveProjectID(ctx, reader, sqldatabase)
	if err != nil {
		return nil, err
	}

	instanceID, err := ResolveSQLInstanceID(ctx, reader, sqldatabase)
	if err != nil {
		return nil, err
	}

	return &SQLDatabase{
		ProjectID:  projectID,
		InstanceID: instanceID,
		DatabaseID: resourceID,
	}, nil
}
