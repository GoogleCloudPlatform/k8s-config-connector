// Copyright 2025 Google LLC
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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TableIdentity defines the resource reference to BigqueryTable, which "External" field
// holds the GCP identifier for the KRM object.
type TableIdentity struct {
	parent *TableParent
	id     string
}

func (i *TableIdentity) String() string {
	return i.parent.String() + "/tables/" + i.id
}

func (i *TableIdentity) ID() string {
	return i.id
}

func (i *TableIdentity) Parent() *TableParent {
	return i.parent
}

type TableParent struct {
	ProjectID string
	DatasetID string
}

func (p *TableParent) String() string {
	return "projects/" + p.ProjectID + "/datasets/" + p.DatasetID
}

// New builds a TableIdentity from the Config Connector Table object.
func NewTableIdentity(ctx context.Context, reader client.Reader, obj *BigqueryTable) (*TableIdentity, error) {
	datasetExternalRef, err := obj.Spec.DatasetRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	datasetParent, dataset, err := refsv1beta1.ParseDatasetExternal(datasetExternalRef)
	if err != nil {
		return nil, err
	}
	projectID := datasetParent.ProjectID

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualParent, actualResourceID, err := ParseTableExternal(externalRef)
		if err != nil {
			return nil, err
		}
		if actualParent.DatasetID != dataset {
			return nil, fmt.Errorf("spec.datasetRef changed, expect %s, got %s", actualParent.DatasetID, dataset)
		}
		if actualResourceID != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualResourceID)
		}
	}
	return &TableIdentity{
		parent: &TableParent{
			ProjectID: projectID,
			DatasetID: dataset,
		},
		id: resourceID,
	}, nil
}

func ParseTableExternal(external string) (parent *TableParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "datasets" || tokens[4] != "tables" {
		return nil, "", fmt.Errorf("format of BigqueryTable external=%q was not known (use projects/{{projectID}}/datasets/{{location}}/tables/{{tableID}})", external)
	}
	parent = &TableParent{
		ProjectID: tokens[1],
		DatasetID: tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
