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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &BigQueryTableIdentity{}
	_ identity.Resource   = &BigQueryTable{}
)

var BigQueryTableIdentityFormat = gcpurls.Template[BigQueryTableIdentity]("bigquery.googleapis.com", "projects/{project}/datasets/{dataset}/tables/{table}")

// BigQueryTableIdentity defines the resource identifier for BigQueryTable.
// +k8s:deepcopy-gen=false
type BigQueryTableIdentity struct {
	Project string
	Dataset string
	Table   string
}

func (i *BigQueryTableIdentity) String() string {
	return BigQueryTableIdentityFormat.ToString(*i)
}

func (i *BigQueryTableIdentity) FromExternal(ref string) error {
	parsed, match, err := BigQueryTableIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigQueryTable external=%q was not known (use %s): %w", ref, BigQueryTableIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigQueryTable external=%q was not known (use %s)", ref, BigQueryTableIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *BigQueryTableIdentity) Host() string {
	return BigQueryTableIdentityFormat.Host()
}

func getIdentityFromBigQueryTableSpec(ctx context.Context, reader client.Reader, obj client.Object) (*BigQueryTableIdentity, error) {
	bigQueryTable, ok := obj.(*BigQueryTable)
	if !ok {
		u, ok := obj.(*unstructured.Unstructured)
		if !ok {
			return nil, fmt.Errorf("expected *BigQueryTable or *unstructured.Unstructured, got %T", obj)
		}
		bigQueryTable = &BigQueryTable{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, bigQueryTable); err != nil {
			return nil, fmt.Errorf("error converting unstructured to BigQueryTable: %w", err)
		}
	}

	if err := bigQueryTable.Spec.DatasetRef.Normalize(ctx, reader, bigQueryTable.GetNamespace()); err != nil {
		return nil, err
	}
	datasetExternalRef := bigQueryTable.Spec.DatasetRef.External
	datasetID := &DatasetIdentity{}
	if err := datasetID.FromExternal(datasetExternalRef); err != nil {
		return nil, err
	}

	projectID := datasetID.Project
	dataset := datasetID.Dataset

	// Get desired ID
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}

	return &BigQueryTableIdentity{
		Project: projectID,
		Dataset: dataset,
		Table:   resourceID,
	}, nil
}

// GetIdentity builds a BigQueryTableIdentity from the Config Connector BigQueryTable object.
func (obj *BigQueryTable) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigQueryTableSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualIdentity := &BigQueryTableIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.Project != specIdentity.Project {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.Project, specIdentity.Project)
		}
		if actualIdentity.Dataset != specIdentity.Dataset {
			return nil, fmt.Errorf("spec.datasetRef changed, expect %s, got %s", actualIdentity.Dataset, specIdentity.Dataset)
		}
		if actualIdentity.Table != specIdentity.Table {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				specIdentity.Table, actualIdentity.Table)
		}
	}
	return specIdentity, nil
}
