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


// +kcc:proto=google.cloud.dataplex.v1.Action
type Action struct {
	// The category of issue associated with the action.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.category
	Category *string `json:"category,omitempty"`

	// Detailed description of the issue requiring action.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.issue
	Issue *string `json:"issue,omitempty"`

	// The time that the issue was detected.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.detect_time
	DetectTime *string `json:"detectTime,omitempty"`

	// The list of data locations associated with this action. Cloud Storage
	//  locations are represented as URI paths(E.g.
	//  `gs://bucket/table1/year=2020/month=Jan/`). BigQuery locations refer to
	//  resource names(E.g.
	//  `bigquery.googleapis.com/projects/project-id/datasets/dataset-id`).
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.data_locations
	DataLocations []string `json:"dataLocations,omitempty"`

	// Details for issues related to invalid or unsupported data formats.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.invalid_data_format
	InvalidDataFormat *Action_InvalidDataFormat `json:"invalidDataFormat,omitempty"`

	// Details for issues related to incompatible schemas detected within data.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.incompatible_data_schema
	IncompatibleDataSchema *Action_IncompatibleDataSchema `json:"incompatibleDataSchema,omitempty"`

	// Details for issues related to invalid or unsupported data partition
	//  structure.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.invalid_data_partition
	InvalidDataPartition *Action_InvalidDataPartition `json:"invalidDataPartition,omitempty"`

	// Details for issues related to absence of data within managed resources.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.missing_data
	MissingData *Action_MissingData `json:"missingData,omitempty"`

	// Details for issues related to absence of a managed resource.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.missing_resource
	MissingResource *Action_MissingResource `json:"missingResource,omitempty"`

	// Details for issues related to lack of permissions to access data
	//  resources.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.unauthorized_resource
	UnauthorizedResource *Action_UnauthorizedResource `json:"unauthorizedResource,omitempty"`

	// Details for issues related to applying security policy.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.failed_security_policy_apply
	FailedSecurityPolicyApply *Action_FailedSecurityPolicyApply `json:"failedSecurityPolicyApply,omitempty"`

	// Details for issues related to invalid data arrangement.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.invalid_data_organization
	InvalidDataOrganization *Action_InvalidDataOrganization `json:"invalidDataOrganization,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Action.FailedSecurityPolicyApply
type Action_FailedSecurityPolicyApply struct {
	// Resource name of one of the assets with failing security policy
	//  application. Populated for a lake or zone resource only.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.FailedSecurityPolicyApply.asset
	Asset *string `json:"asset,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Action.IncompatibleDataSchema
type Action_IncompatibleDataSchema struct {
	// The name of the table containing invalid data.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.IncompatibleDataSchema.table
	Table *string `json:"table,omitempty"`

	// The existing and expected schema of the table. The schema is provided as
	//  a JSON formatted structure listing columns and data types.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.IncompatibleDataSchema.existing_schema
	ExistingSchema *string `json:"existingSchema,omitempty"`

	// The new and incompatible schema within the table. The schema is provided
	//  as a JSON formatted structured listing columns and data types.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.IncompatibleDataSchema.new_schema
	NewSchema *string `json:"newSchema,omitempty"`

	// The list of data locations sampled and used for format/schema
	//  inference.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.IncompatibleDataSchema.sampled_data_locations
	SampledDataLocations []string `json:"sampledDataLocations,omitempty"`

	// Whether the action relates to a schema that is incompatible or modified.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.IncompatibleDataSchema.schema_change
	SchemaChange *string `json:"schemaChange,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Action.InvalidDataFormat
type Action_InvalidDataFormat struct {
	// The list of data locations sampled and used for format/schema
	//  inference.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.InvalidDataFormat.sampled_data_locations
	SampledDataLocations []string `json:"sampledDataLocations,omitempty"`

	// The expected data format of the entity.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.InvalidDataFormat.expected_format
	ExpectedFormat *string `json:"expectedFormat,omitempty"`

	// The new unexpected data format within the entity.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.InvalidDataFormat.new_format
	NewFormat *string `json:"newFormat,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Action.InvalidDataOrganization
type Action_InvalidDataOrganization struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Action.InvalidDataPartition
type Action_InvalidDataPartition struct {
	// The issue type of InvalidDataPartition.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.InvalidDataPartition.expected_structure
	ExpectedStructure *string `json:"expectedStructure,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Action.MissingData
type Action_MissingData struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Action.MissingResource
type Action_MissingResource struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Action.UnauthorizedResource
type Action_UnauthorizedResource struct {
}

// +kcc:proto=google.cloud.dataplex.v1.Action
type ActionObservedState struct {
	// Output only. The relative resource name of the action, of the form:
	//  `projects/{project}/locations/{location}/lakes/{lake}/actions/{action}`
	//  `projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}/actions/{action}`
	//  `projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}/assets/{asset}/actions/{action}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.name
	Name *string `json:"name,omitempty"`

	// Output only. The relative resource name of the lake, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.lake
	Lake *string `json:"lake,omitempty"`

	// Output only. The relative resource name of the zone, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.zone
	Zone *string `json:"zone,omitempty"`

	// Output only. The relative resource name of the asset, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}/assets/{asset_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Action.asset
	Asset *string `json:"asset,omitempty"`
}
