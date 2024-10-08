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

type BigQueryDatasetRef struct {
	// If provided must be in the format `projects/[project_id]/datasets/[dataset_id]`.
	External string `json:"external,omitempty"`
	// The `metadata.name` field of a `BigQueryDataset` resource.
	Name string `json:"name,omitempty"`
	// The `metadata.namespace` field of a `BigQueryDataset` resource.
	Namespace string `json:"namespace,omitempty"`
}

type BigQueryDataset struct {
	projectID string
	datasetID string
}

func ResolveBigQueryDataset(ctx context.Context, reader client.Reader, src client.Object, ref *BigQueryDatasetRef) (*BigQueryDataset, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.Name == "" && ref.External == "" {
		return nil, fmt.Errorf("must specify either name or external on BigQueryDatasetRef")
	}
	if ref.Name != "" && ref.External != "" {
		return nil, fmt.Errorf("cannot specify both name and external on BigQueryDatasetRef")
	}

	// External is provided.
	if ref.External != "" {
		// External should be in the `projects/[project_id]/datasets/[dataset_id]` format.
		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "datasets" {
			return &BigQueryDataset{
				projectID: tokens[1],
				datasetID: tokens[3],
			}, nil
		}
		return nil, fmt.Errorf("format of BigQueryDatasetRef external=%q was not known (use projects/[project_id]/datasets/[dataset_id])", ref.External)

	}

	// Fetch BigQueryDataset object to construct the external form.
	dataset := &unstructured.Unstructured{}
	dataset.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "bigquery.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "BigQueryDataset",
	})
	nn := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if nn.Namespace == "" {
		nn.Namespace = src.GetNamespace()
	}
	if err := reader.Get(ctx, nn, dataset); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced BigQueryDataset %v not found", nn)
		}
		return nil, fmt.Errorf("error reading referenced BigQueryDataset %v: %w", nn, err)
	}
	projectID, err := ResolveProjectID(ctx, reader, dataset)
	if err != nil {
		return nil, err
	}
	datasetID, err := GetResourceID(dataset)
	if err != nil {
		return nil, err
	}
	return &BigQueryDataset{
		projectID: projectID,
		datasetID: datasetID,
	}, nil
}

func (d *BigQueryDataset) String() string {
	return fmt.Sprintf("projects/%s/datasets/%s", d.projectID, d.datasetID)
}

func (d *BigQueryDataset) GetDatasetID() string {
	return d.datasetID
}
