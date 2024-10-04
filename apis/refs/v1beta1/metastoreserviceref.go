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

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type MetastoreServiceRef struct {
	/* The self-link of an existing Dataproc Metastore service , when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a Dataproc Metastore service. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a Dataproc Metastore service. */
	Namespace string `json:"namespace,omitempty"`
}

type MetastoreService struct {
	ProjectID string
	Location  string
	ServiceID string
}

func (s *MetastoreService) String() string {
	return "projects/" + s.ProjectID + "/locations/" + s.Location + "/services/" + s.ServiceID
}

func ResolveMetastoreServiceRef(ctx context.Context, reader client.Reader, obj client.Object, ref *MetastoreServiceRef) (*MetastoreService, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on MetastoreServiceRef")
	}
	if ref.External != "" && ref.Name != "" {
		return nil, fmt.Errorf("cannot specify both spec.metastoreServiceRef.name and spec.metastoreServiceRef.external")
	}

	if ref.External != "" {
		// External must be in form `projects/<projectID>/locations/<location>/services/<ServiceID>`.
		// see https://cloud.google.com/dataproc-metastore/docs/reference/rest/v1beta/projects.locations.services/get
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "services" {
			return &MetastoreService{
				ProjectID: tokens[1],
				Location:  tokens[3],
				ServiceID: tokens[5],
			}, nil
		}
		return nil, fmt.Errorf("format of MetastoreService external=%q was not known (use projects/<projectID>/locations/<location>/services/<ServiceID>)", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	// NOTYET: Currently there is no MetaStoreService KCC resource
	service := &unstructured.Unstructured{}
	service.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "metastore.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "MetastoreService",
	})
	if err := reader.Get(ctx, key, service); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced MetastoreService %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced MetastoreService %v: %w", key, err)
	}

	resourceID, _, err := unstructured.NestedString(service.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from MetastoreService %s/%s: %w", service.GetNamespace(), service.GetName(), err)
	}
	if resourceID == "" {
		resourceID = service.GetName()
	}

	projectID, err := ResolveProjectID(ctx, reader, service)
	if err != nil {
		return nil, err
	}

	location, _, err := unstructured.NestedString(service.Object, "spec", "region")
	if err != nil {
		return nil, fmt.Errorf("reading spec.region from DataprocCluster %s/%s: %w", service.GetNamespace(), service.GetName(), err)
	}

	return &MetastoreService{
		ProjectID: projectID,
		Location:  location,
		ServiceID: resourceID,
	}, nil
}
