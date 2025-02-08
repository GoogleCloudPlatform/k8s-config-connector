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

// +kcc:proto=google.logging.v2.Link
type Link struct {
	// The resource name of the link. The name can have up to 100 characters.
	//  A valid link id (at the end of the link name) must only have alphanumeric
	//  characters and underscores within it.
	//
	//      "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]"
	//      "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]"
	//      "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]"
	//      "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]"
	//
	//  For example:
	//
	//    `projects/my-project/locations/global/buckets/my-bucket/links/my_link
	// +kcc:proto:field=google.logging.v2.Link.name
	Name *string `json:"name,omitempty"`

	// Describes this link.
	//
	//  The maximum length of the description is 8000 characters.
	// +kcc:proto:field=google.logging.v2.Link.description
	Description *string `json:"description,omitempty"`

	// The information of a BigQuery Dataset. When a link is created, a BigQuery
	//  dataset is created along with it, in the same project as the LogBucket it's
	//  linked to. This dataset will also have BigQuery Views corresponding to the
	//  LogViews in the bucket.
	// +kcc:proto:field=google.logging.v2.Link.bigquery_dataset
	BigqueryDataset *BigQueryDataset `json:"bigqueryDataset,omitempty"`
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

// +kcc:proto=google.logging.v2.Link
type LinkObservedState struct {
	// Output only. The creation timestamp of the link.
	// +kcc:proto:field=google.logging.v2.Link.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The resource lifecycle state.
	// +kcc:proto:field=google.logging.v2.Link.lifecycle_state
	LifecycleState *string `json:"lifecycleState,omitempty"`

	// The information of a BigQuery Dataset. When a link is created, a BigQuery
	//  dataset is created along with it, in the same project as the LogBucket it's
	//  linked to. This dataset will also have BigQuery Views corresponding to the
	//  LogViews in the bucket.
	// +kcc:proto:field=google.logging.v2.Link.bigquery_dataset
	BigqueryDataset *BigQueryDatasetObservedState `json:"bigqueryDataset,omitempty"`
}
