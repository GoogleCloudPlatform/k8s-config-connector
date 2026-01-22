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

package v1beta1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var TableIdentityFormat = gcpurls.Template[TableIdentity]("bigquery.googleapis.com", "projects/{project}/datasets/{dataset}/tables/{table}")

var _ identity.Identity = &TableIdentity{}

// TableIdentity defines the resource reference to BigQueryTable, which "External" field
// holds the GCP identifier for the KRM object.
type TableIdentity struct {
	Project string
	Dataset string
	Table   string
}

func (i *TableIdentity) String() string {
	return TableIdentityFormat.ToString(*i)
}

func (i *TableIdentity) ID() string {
	return i.Table
}

func (i *TableIdentity) Parent() *TableParent {
	return &TableParent{
		ProjectID: i.Project,
		DatasetID: i.Dataset,
	}
}

func (i *TableIdentity) FromExternal(ref string) error {
	parsed, match, err := TableIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of Table external=%q was not known (use %s): %w", ref, TableIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of Table external=%q was not known (use %s)", ref, TableIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *TableIdentity) Host() string {
	return TableIdentityFormat.Host()
}

type TableParent struct {
	ProjectID string
	DatasetID string
}

func (p *TableParent) String() string {
	return "projects/" + p.ProjectID + "/datasets/" + p.DatasetID
}

// New builds a TableIdentity from the Config Connector Table object.
func NewTableIdentity(ctx context.Context, reader client.Reader, obj *BigQueryTable) (*TableIdentity, error) {
	if err := obj.Spec.DatasetRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	datasetExternalRef := obj.Spec.DatasetRef.External
	datasetParent, dataset, err := ParseDatasetExternal(datasetExternalRef)
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
		actualIdentity := &TableIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.Project != datasetParent.ProjectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.Project, projectID)
		}
		if actualIdentity.Dataset != dataset {
			return nil, fmt.Errorf("spec.datasetRef changed, expect %s, got %s", actualIdentity.Dataset, dataset)
		}
		if actualIdentity.Table != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.Table)
		}
	}
	return &TableIdentity{
		Project: projectID,
		Dataset: dataset,
		Table:   resourceID,
	}, nil
}
