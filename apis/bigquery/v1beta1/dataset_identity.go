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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &DatasetIdentity{}
	_ identity.Resource   = &BigQueryDataset{}
)

var (
	DatasetIdentityFormat = gcpurls.Template[DatasetIdentity]("bigquery.googleapis.com", "projects/{project}/datasets/{dataset}")
)

// DatasetIdentity defines the resource reference to BigQueryDataset, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type DatasetIdentity struct {
	Project string
	Dataset string
}

func (i *DatasetIdentity) String() string {
	return DatasetIdentityFormat.ToString(*i)
}

func (i *DatasetIdentity) FromExternal(ref string) error {
	parsed, match, err := DatasetIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigQueryDataset external=%q was not known (use %s): %w", ref, DatasetIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigQueryDataset external=%q was not known (use %s)", ref, DatasetIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DatasetIdentity) Host() string {
	return DatasetIdentityFormat.Host()
}

func getIdentityFromBigQueryDatasetSpec(ctx context.Context, reader client.Reader, obj client.Object) (*DatasetIdentity, error) {
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	return &DatasetIdentity{
		Project: projectID,
		Dataset: resourceID,
	}, nil
}

func (obj *BigQueryDataset) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromBigQueryDatasetSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &DatasetIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change BigQueryDataset identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}
	return specIdentity, nil
}
