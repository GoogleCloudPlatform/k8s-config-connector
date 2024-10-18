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

type DataExchangeRef struct {
	/* The DataExchange selfLink, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `DataExchange` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `DataExchange` resource. */
	Namespace string `json:"namespace,omitempty"`
}

type DataExchange struct {
	ProjectID      string
	Location       string
	DataExchangeID string
}

func (s *DataExchange) String() string {
	return "projects/" + s.ProjectID + "locations/" + s.Location + "/dataExchanges/" + s.DataExchangeID
}

func ResolveDataExchangeRef(ctx context.Context, reader client.Reader, obj client.Object, ref *DataExchangeRef) (*DataExchange, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on dataExchangeRef")
	}
	if ref.External != "" && ref.Name != "" {
		return nil, fmt.Errorf("cannot specify both spec.dataExchangeRef.name and spec.dataExchangeRef.external")
	}

	if ref.External != "" {
		// External should be in the `projects/[projectID]/locations/[Location]/dataExchanges/[dataExchangeID]` format.
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dataExchanges" {
			return &DataExchange{
				ProjectID:      tokens[1],
				Location:       tokens[3],
				DataExchangeID: tokens[5],
			}, nil
		}
		return nil, fmt.Errorf("format of DataExchange external=%q was not known (use projects/<projectId>/locations/<location>/dataExchanges/<dataExchangeID>)", ref.External)
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	exchange := &unstructured.Unstructured{}
	exchange.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "bigqueryanalyticshub.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "BigQueryAnalyticsHubDataExchange",
	})
	if err := reader.Get(ctx, key, exchange); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced DataExchange %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced DataExchange %v: %w", key, err)
	}

	resourceID, _, err := unstructured.NestedString(exchange.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from DataExchange %s/%s: %w", exchange.GetNamespace(), exchange.GetName(), err)
	}
	if resourceID == "" {
		resourceID = exchange.GetName()
	}

	location, _, err := unstructured.NestedString(exchange.Object, "spec", "location")
	if err != nil {
		return nil, fmt.Errorf("reading spec.region from DataExchange %s/%s: %w", exchange.GetNamespace(), exchange.GetName(), err)
	}

	projectID, err := ResolveProjectID(ctx, reader, exchange)
	if err != nil {
		return nil, err
	}

	return &DataExchange{
		ProjectID:      projectID,
		Location:       location,
		DataExchangeID: resourceID,
	}, nil
}
