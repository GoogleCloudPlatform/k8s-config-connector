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

// +kcc:proto=google.logging.v2.BigQueryDataset
type BigQueryDataset struct {
}

// +kcc:proto=google.logging.v2.BigQueryDataset
type BigQueryDatasetObservedState struct {
	// Output only. The full resource name of the BigQuery dataset. The DATASET_ID
	//  will match the ID of the link, so the link must match the naming
	//  restrictions of BigQuery datasets (alphanumeric characters and underscores
	//  only).
	//
	//  The dataset will have a resource path of
	//    "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET_ID]"
	// +kcc:proto:field=google.logging.v2.BigQueryDataset.dataset_id
	DatasetID *string `json:"datasetID,omitempty"`
}
