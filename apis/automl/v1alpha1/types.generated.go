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


// +kcc:proto=google.cloud.automl.v1beta1.BigQuerySource
type BigQuerySource struct {
	// Required. BigQuery URI to a table, up to 2000 characters long.
	//  Accepted forms:
	//  *  BigQuery path e.g. bq://projectId.bqDatasetId.bqTableId
	// +kcc:proto:field=google.cloud.automl.v1beta1.BigQuerySource.input_uri
	InputURI *string `json:"inputURI,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.GcsSource
type GcsSource struct {
	// Required. Google Cloud Storage URIs to input files, up to 2000 characters
	//  long. Accepted forms:
	//  * Full object path, e.g. gs://bucket/directory/object.csv
	// +kcc:proto:field=google.cloud.automl.v1beta1.GcsSource.input_uris
	InputUris []string `json:"inputUris,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.InputConfig
type InputConfig struct {
	// The Google Cloud Storage location for the input content.
	//  In ImportData, the gcs_source points to a csv with structure described in
	//  the comment.
	// +kcc:proto:field=google.cloud.automl.v1beta1.InputConfig.gcs_source
	GcsSource *GcsSource `json:"gcsSource,omitempty"`

	// The BigQuery location for the input content.
	// +kcc:proto:field=google.cloud.automl.v1beta1.InputConfig.bigquery_source
	BigquerySource *BigQuerySource `json:"bigquerySource,omitempty"`

	// Additional domain-specific parameters describing the semantic of the
	//  imported data, any string must be up to 25000
	//  characters long.
	//
	//  *  For Tables:
	//     `schema_inference_version` - (integer) Required. The version of the
	//         algorithm that should be used for the initial inference of the
	//         schema (columns' DataTypes) of the table the data is being imported
	//         into. Allowed values: "1".
	// +kcc:proto:field=google.cloud.automl.v1beta1.InputConfig.params
	Params map[string]string `json:"params,omitempty"`
}

// +kcc:proto=google.cloud.automl.v1beta1.TableSpec
type TableSpec struct {
	// Output only. The resource name of the table spec.
	//  Form:
	//
	//  `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/tableSpecs/{table_spec_id}`
	// +kcc:proto:field=google.cloud.automl.v1beta1.TableSpec.name
	Name *string `json:"name,omitempty"`

	// column_spec_id of the time column. Only used if the parent dataset's
	//  ml_use_column_spec_id is not set. Used to split rows into TRAIN, VALIDATE
	//  and TEST sets such that oldest rows go to TRAIN set, newest to TEST, and
	//  those in between to VALIDATE.
	//  Required type: TIMESTAMP.
	//  If both this column and ml_use_column are not set, then ML use of all rows
	//  will be assigned by AutoML. NOTE: Updates of this field will instantly
	//  affect any other users concurrently working with the dataset.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TableSpec.time_column_spec_id
	TimeColumnSpecID *string `json:"timeColumnSpecID,omitempty"`

	// Output only. The number of rows (i.e. examples) in the table.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TableSpec.row_count
	RowCount *int64 `json:"rowCount,omitempty"`

	// Output only. The number of valid rows (i.e. without values that don't match
	//  DataType-s of their columns).
	// +kcc:proto:field=google.cloud.automl.v1beta1.TableSpec.valid_row_count
	ValidRowCount *int64 `json:"validRowCount,omitempty"`

	// Output only. The number of columns of the table. That is, the number of
	//  child ColumnSpec-s.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TableSpec.column_count
	ColumnCount *int64 `json:"columnCount,omitempty"`

	// Output only. Input configs via which data currently residing in the table
	//  had been imported.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TableSpec.input_configs
	InputConfigs []InputConfig `json:"inputConfigs,omitempty"`

	// Used to perform consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.automl.v1beta1.TableSpec.etag
	Etag *string `json:"etag,omitempty"`
}
