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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

// +kcc:proto=google.cloud.bigquery.v2.Access
type Access struct {
	// An IAM role ID that should be granted to the user, group,
	//  or domain specified in this access entry.
	//  The following legacy mappings will be applied:
	//
	//  * `OWNER`: `roles/bigquery.dataOwner`
	//  * `WRITER`: `roles/bigquery.dataEditor`
	//  * `READER`: `roles/bigquery.dataViewer`
	//
	//  This field will accept any of the above formats, but will return only
	//  the legacy format. For example, if you set this field to
	//  "roles/bigquery.dataOwner", it will be returned back as "OWNER".
	Role *string `json:"role,omitempty"`

	// [Pick one] An email address of a user to grant access to. For example:
	//  fred@example.com. Maps to IAM policy member "user:EMAIL" or
	//  "serviceAccount:EMAIL".
	UserByEmail *string `json:"userByEmail,omitempty"`

	// [Pick one] An email address of a Google Group to grant access to.
	//  Maps to IAM policy member "group:GROUP".
	GroupByEmail *string `json:"groupByEmail,omitempty"`

	// [Pick one] A domain to grant access to. Any users signed in with the domain
	//  specified will be granted the specified access. Example: "example.com".
	//  Maps to IAM policy member "domain:DOMAIN".
	Domain *string `json:"domain,omitempty"`

	// [Pick one] A special group to grant access to. Possible values include:
	//
	//    * projectOwners: Owners of the enclosing project.
	//    * projectReaders: Readers of the enclosing project.
	//    * projectWriters: Writers of the enclosing project.
	//    * allAuthenticatedUsers: All authenticated BigQuery users.
	//
	//  Maps to similarly-named IAM members.
	SpecialGroup *string `json:"specialGroup,omitempty"`

	// [Pick one] Some other type of member that appears in the IAM Policy but
	//  isn't a user, group, domain, or special group.
	IamMember *string `json:"iamMember,omitempty"`

	// [Pick one] A view from a different dataset to grant access to. Queries
	//  executed against that view will have read access to views/tables/routines
	//  in this dataset.
	//  The role field is not required when this field is set. If that view is
	//  updated by any user, access to the view needs to be granted again via an
	//  update operation.
	View *TableReference `json:"view,omitempty"`

	// [Pick one] A routine from a different dataset to grant access to. Queries
	//  executed against that routine will have read access to
	//  views/tables/routines in this dataset. Only UDF is supported for now.
	//  The role field is not required when this field is set. If that routine is
	//  updated by any user, access to the routine needs to be granted again via
	//  an update operation.
	Routine *RoutineReference `json:"routine,omitempty"`

	// [Pick one] A grant authorizing all resources of a particular type in a
	//  particular dataset access to this dataset. Only views are supported for
	//  now. The role field is not required when this field is set. If that dataset
	//  is deleted and re-created, its access needs to be granted again via an
	//  update operation.
	Dataset *DatasetAccessEntry `json:"dataset,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.Dataset
type Dataset struct {
	// Output only. The resource type.
	Kind *string `json:"kind,omitempty"`

	// Output only. A hash of the resource.
	Etag *string `json:"etag,omitempty"`

	// Output only. The fully-qualified unique name of the dataset in the format
	//  projectId:datasetId. The dataset name without the project name is given in
	//  the datasetId field. When creating a new dataset, leave this field blank,
	//  and instead specify the datasetId field.
	ID *string `json:"id,omitempty"`

	// Output only. A URL that can be used to access the resource again. You can
	//  use this URL in Get or Update requests to the resource.
	SelfLink *string `json:"selfLink,omitempty"`

	// Required. A reference that identifies the dataset.
	DatasetReference *DatasetReference `json:"datasetReference,omitempty"`

	// Optional. A descriptive name for the dataset.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// Optional. A user-friendly description of the dataset.
	Description *string `json:"description,omitempty"`

	// Optional. The default lifetime of all tables in the dataset, in
	//  milliseconds. The minimum lifetime value is 3600000 milliseconds (one
	//  hour). To clear an existing default expiration with a PATCH request, set to
	//  0. Once this property is set, all newly-created tables in the dataset will
	//  have an expirationTime property set to the creation time plus the value in
	//  this property, and changing the value will only affect new tables, not
	//  existing ones. When the expirationTime for a given table is reached, that
	//  table will be deleted automatically.
	//  If a table's expirationTime is modified or removed before the table
	//  expires, or if you provide an explicit expirationTime when creating a
	//  table, that value takes precedence over the default expiration time
	//  indicated by this property.
	DefaultTableExpirationMs *int64 `json:"defaultTableExpirationMs,omitempty"`

	// This default partition expiration, expressed in milliseconds.
	//
	//  When new time-partitioned tables are created in a dataset where this
	//  property is set, the table will inherit this value, propagated as the
	//  `TimePartitioning.expirationMs` property on the new table.  If you set
	//  `TimePartitioning.expirationMs` explicitly when creating a table,
	//  the `defaultPartitionExpirationMs` of the containing dataset is ignored.
	//
	//  When creating a partitioned table, if `defaultPartitionExpirationMs`
	//  is set, the `defaultTableExpirationMs` value is ignored and the table
	//  will not be inherit a table expiration deadline.
	DefaultPartitionExpirationMs *int64 `json:"defaultPartitionExpirationMs,omitempty"`

	// The labels associated with this dataset. You can use these
	//  to organize and group your datasets.
	//  You can set this property when inserting or updating a dataset.
	//  See [Creating and Updating Dataset
	//  Labels](https://cloud.google.com/bigquery/docs/creating-managing-labels#creating_and_updating_dataset_labels)
	//  for more information.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. An array of objects that define dataset access for one or more
	//  entities. You can set this property when inserting or updating a dataset in
	//  order to control who is allowed to access the data. If unspecified at
	//  dataset creation time, BigQuery adds default dataset access for the
	//  following entities: access.specialGroup: projectReaders; access.role:
	//  READER; access.specialGroup: projectWriters; access.role: WRITER;
	//  access.specialGroup: projectOwners; access.role: OWNER;
	//  access.userByEmail: [dataset creator email]; access.role: OWNER;
	//  If you patch a dataset, then this field is overwritten by the patched
	//  dataset's access field. To add entities, you must supply the entire
	//  existing access array in addition to any new entities that you want to add.
	Access []Access `json:"access,omitempty"`

	// Output only. The time when this dataset was created, in milliseconds since
	//  the epoch.
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Output only. The date when this dataset was last modified, in milliseconds
	//  since the epoch.
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`

	// The geographic location where the dataset should reside. See
	//  https://cloud.google.com/bigquery/docs/locations for supported
	//  locations.
	Location *string `json:"location,omitempty"`

	// The default encryption key for all tables in the dataset.
	//  After this property is set, the encryption key of all newly-created tables
	//  in the dataset is set to this value unless the table creation request or
	//  query explicitly overrides the key.
	DefaultEncryptionConfiguration *EncryptionConfiguration `json:"defaultEncryptionConfiguration,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *BoolValue `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzi *BoolValue `json:"satisfiesPzi,omitempty"`

	// Output only. Same as `type` in `ListFormatDataset`.
	//  The type of the dataset, one of:
	//
	//  * DEFAULT - only accessible by owner and authorized accounts,
	//  * PUBLIC - accessible by everyone,
	//  * LINKED - linked dataset,
	//  * EXTERNAL - dataset with definition in external metadata catalog.
	Type *string `json:"type,omitempty"`

	// Optional. The source dataset reference when the dataset is of type LINKED.
	//  For all other dataset types it is not set. This field cannot be updated
	//  once it is set. Any attempt to update this field using Update and Patch API
	//  Operations will be ignored.
	LinkedDatasetSource *LinkedDatasetSource `json:"linkedDatasetSource,omitempty"`

	// Output only. Metadata about the LinkedDataset. Filled out when the dataset
	//  type is LINKED.
	LinkedDatasetMetadata *LinkedDatasetMetadata `json:"linkedDatasetMetadata,omitempty"`

	// Optional. Reference to a read-only external dataset defined in data
	//  catalogs outside of BigQuery. Filled out when the dataset type is EXTERNAL.
	ExternalDatasetReference *ExternalDatasetReference `json:"externalDatasetReference,omitempty"`

	// Optional. Options defining open source compatible datasets living in the
	//  BigQuery catalog. Contains metadata of open source database, schema or
	//  namespace represented by the current dataset.
	ExternalCatalogDatasetOptions *ExternalCatalogDatasetOptions `json:"externalCatalogDatasetOptions,omitempty"`

	// Optional. TRUE if the dataset and its table names are case-insensitive,
	//  otherwise FALSE. By default, this is FALSE, which means the dataset and its
	//  table names are case-sensitive. This field does not affect routine
	//  references.
	IsCaseInsensitive *BoolValue `json:"isCaseInsensitive,omitempty"`

	// Optional. Defines the default collation specification of future tables
	//  created in the dataset. If a table is created in this dataset without
	//  table-level default collation, then the table inherits the dataset default
	//  collation, which is applied to the string fields that do not have explicit
	//  collation specified. A change to this field affects only tables created
	//  afterwards, and does not alter the existing tables.
	//  The following values are supported:
	//
	//  * 'und:ci': undetermined locale, case insensitive.
	//  * '': empty string. Default to case-sensitive behavior.
	DefaultCollation *string `json:"defaultCollation,omitempty"`

	// Optional. Defines the default rounding mode specification of new tables
	//  created within this dataset. During table creation, if this field is
	//  specified, the table within this dataset will inherit the default rounding
	//  mode of the dataset. Setting the default rounding mode on a table overrides
	//  this option. Existing tables in the dataset are unaffected.
	//  If columns are defined during that table creation,
	//  they will immediately inherit the table's default rounding mode,
	//  unless otherwise specified.
	DefaultRoundingMode *string `json:"defaultRoundingMode,omitempty"`

	// Optional. Defines the time travel window in hours. The value can be from 48
	//  to 168 hours (2 to 7 days). The default value is 168 hours if this is not
	//  set.
	MaxTimeTravelHours *int64 `json:"maxTimeTravelHours,omitempty"`

	// Output only. Tags for the dataset. To provide tags as inputs, use the
	//  `resourceTags` field.
	Tags []GcpTag `json:"tags,omitempty"`

	// Optional. Updates storage_billing_model for the dataset.
	StorageBillingModel *string `json:"storageBillingModel,omitempty"`

	// Optional. Output only. Restriction config for all tables and dataset. If
	//  set, restrict certain accesses on the dataset and all its tables based on
	//  the config. See [Data
	//  egress](https://cloud.google.com/bigquery/docs/analytics-hub-introduction#data_egress)
	//  for more details.
	Restrictions *RestrictionConfig `json:"restrictions,omitempty"`

	// Optional. The [tags](https://cloud.google.com/bigquery/docs/tags) attached
	//  to this dataset. Tag keys are globally unique. Tag key is expected to be in
	//  the namespaced format, for example "123456789012/environment" where
	//  123456789012 is the ID of the parent organization or project resource for
	//  this tag key. Tag value is expected to be the short name, for example
	//  "Production". See [Tag
	//  definitions](https://cloud.google.com/iam/docs/tags-access-control#definitions)
	//  for more details.
	ResourceTags map[string]string `json:"resourceTags,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DatasetAccessEntry
type DatasetAccessEntry struct {
	// The dataset this entry applies to.
	// +required
	Dataset *DatasetReference `json:"dataset,omitempty"`

	// Which resources in the dataset this entry applies to. Currently, only
	//  views are supported, but additional target types may be added in the
	//  future.
	// +required
	TargetTypes []string `json:"targetTypes,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.DatasetReference
type DatasetReference struct {
	// Required. A unique Id for this dataset, without the project name. The Id
	//  must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	//  The maximum length is 1,024 characters.
	DatasetId *string `json:"datasetId,omitempty"`

	// Required. The Id of the project containing this dataset.
	ProjectId *string `json:"projectId,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.EncryptionConfiguration
type EncryptionConfiguration struct {
	// Optional. Describes the Cloud KMS encryption key that will be used to
	//  protect destination BigQuery table. The BigQuery Service Account associated
	//  with your project requires access to this encryption key.
	KmsKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ErrorProto
type ErrorProto struct {
	// A short error code that summarizes the error.
	Reason *string `json:"reason,omitempty"`

	// Specifies where the error occurred, if present.
	Location *string `json:"location,omitempty"`

	// Debugging information. This property is internal to Google and should not
	//  be used.
	DebugInfo *string `json:"debugInfo,omitempty"`

	// A human-readable description of the error.
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExplainQueryStage
type ExplainQueryStage struct {
	// Human-readable name for the stage.
	Name *string `json:"name,omitempty"`

	// Unique ID for the stage within the plan.
	ID *int64 `json:"id,omitempty"`

	// Stage start time represented as milliseconds since the epoch.
	StartMs *int64 `json:"startMs,omitempty"`

	// Stage end time represented as milliseconds since the epoch.
	EndMs *int64 `json:"endMs,omitempty"`

	// IDs for stages that are inputs to this stage.
	InputStages []int64 `json:"inputStages,omitempty"`

	// Relative amount of time the average shard spent waiting to be
	//  scheduled.
	WaitRatioAvg *float64 `json:"waitRatioAvg,omitempty"`

	// Milliseconds the average shard spent waiting to be scheduled.
	WaitMsAvg *int64 `json:"waitMsAvg,omitempty"`

	// Relative amount of time the slowest shard spent waiting to be
	//  scheduled.
	WaitRatioMax *float64 `json:"waitRatioMax,omitempty"`

	// Milliseconds the slowest shard spent waiting to be scheduled.
	WaitMsMax *int64 `json:"waitMsMax,omitempty"`

	// Relative amount of time the average shard spent reading input.
	ReadRatioAvg *float64 `json:"readRatioAvg,omitempty"`

	// Milliseconds the average shard spent reading input.
	ReadMsAvg *int64 `json:"readMsAvg,omitempty"`

	// Relative amount of time the slowest shard spent reading input.
	ReadRatioMax *float64 `json:"readRatioMax,omitempty"`

	// Milliseconds the slowest shard spent reading input.
	ReadMsMax *int64 `json:"readMsMax,omitempty"`

	// Relative amount of time the average shard spent on CPU-bound tasks.
	ComputeRatioAvg *float64 `json:"computeRatioAvg,omitempty"`

	// Milliseconds the average shard spent on CPU-bound tasks.
	ComputeMsAvg *int64 `json:"computeMsAvg,omitempty"`

	// Relative amount of time the slowest shard spent on CPU-bound tasks.
	ComputeRatioMax *float64 `json:"computeRatioMax,omitempty"`

	// Milliseconds the slowest shard spent on CPU-bound tasks.
	ComputeMsMax *int64 `json:"computeMsMax,omitempty"`

	// Relative amount of time the average shard spent on writing output.
	WriteRatioAvg *float64 `json:"writeRatioAvg,omitempty"`

	// Milliseconds the average shard spent on writing output.
	WriteMsAvg *int64 `json:"writeMsAvg,omitempty"`

	// Relative amount of time the slowest shard spent on writing output.
	WriteRatioMax *float64 `json:"writeRatioMax,omitempty"`

	// Milliseconds the slowest shard spent on writing output.
	WriteMsMax *int64 `json:"writeMsMax,omitempty"`

	// Total number of bytes written to shuffle.
	ShuffleOutputBytes *int64 `json:"shuffleOutputBytes,omitempty"`

	// Total number of bytes written to shuffle and spilled to disk.
	ShuffleOutputBytesSpilled *int64 `json:"shuffleOutputBytesSpilled,omitempty"`

	// Number of records read into the stage.
	RecordsRead *int64 `json:"recordsRead,omitempty"`

	// Number of records written by the stage.
	RecordsWritten *int64 `json:"recordsWritten,omitempty"`

	// Number of parallel input segments to be processed
	ParallelInputs *int64 `json:"parallelInputs,omitempty"`

	// Number of parallel input segments completed.
	CompletedParallelInputs *int64 `json:"completedParallelInputs,omitempty"`

	// Current status for this stage.
	Status *string `json:"status,omitempty"`

	// List of operations within the stage in dependency order (approximately
	//  chronological).
	Steps []ExplainQueryStep `json:"steps,omitempty"`

	// Slot-milliseconds used by the stage.
	SlotMs *int64 `json:"slotMs,omitempty"`

	// Output only. Compute mode for this stage.
	ComputeMode *string `json:"computeMode,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExplainQueryStep
type ExplainQueryStep struct {
	// Machine-readable operation type.
	Kind *string `json:"kind,omitempty"`

	// Human-readable description of the step(s).
	Substeps []string `json:"substeps,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExportDataStatistics
type ExportDataStatistics struct {
	// Number of destination files generated in case of EXPORT DATA
	//  statement only.
	FileCount *int64 `json:"fileCount,omitempty"`

	// [Alpha] Number of destination rows generated in case of EXPORT DATA
	//  statement only.
	RowCount *int64 `json:"rowCount,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExternalCatalogDatasetOptions
type ExternalCatalogDatasetOptions struct {
	// Optional. A map of key value pairs defining the parameters and properties
	//  of the open source schema. Maximum size of 2Mib.
	Parameters map[string]string `json:"parameters,omitempty"`

	// Optional. The storage location URI for all tables in the dataset.
	//  Equivalent to hive metastore's database locationUri. Maximum length of 1024
	//  characters.
	DefaultStorageLocationUri *string `json:"defaultStorageLocationUri,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.ExternalDatasetReference
type ExternalDatasetReference struct {
	// Required. External source that backs this dataset.
	ExternalSource *string `json:"externalSource,omitempty"`

	// Required. The connection id that is used to access the external_source.
	//
	//  Format:
	//    projects/{project_id}/locations/{location_id}/connections/{connection_id}
	Connection *string `json:"connection,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.GcpTag
type GcpTag struct {
	// Required. The namespaced friendly name of the tag key, e.g.
	//  "12345/environment" where 12345 is org id.
	TagKey *string `json:"tagKey,omitempty"`

	// Required. The friendly short name of the tag value, e.g. "production".
	TagValue *string `json:"tagValue,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.LinkedDatasetSource
type LinkedDatasetSource struct {
	// The source dataset reference contains project numbers and not project ids.
	SourceDataset *DatasetReference `json:"sourceDataset,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RestrictionConfig
type RestrictionConfig struct {
	// Output only. Specifies the type of dataset/table restriction.
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.RoutineReference
type RoutineReference struct {
	// Required. The Id of the project containing this routine.
	ProjectId *string `json:"projectId,omitempty"`

	// Required. The Id of the dataset containing this routine.
	DatasetId *string `json:"datasetId,omitempty"`

	// Required. The Id of the routine. The Id must contain only
	//  letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	//  length is 256 characters.
	RoutineId *string `json:"routineId,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.v2.TableReference
type TableReference struct {
	// Required. The Id of the project containing this table.
	ProjectId *string `json:"projectId,omitempty"`

	// Required. The Id of the dataset containing this table.
	DatasetId *string `json:"datasetId,omitempty"`

	// Required. The Id of the table. The Id can contain Unicode characters in
	//  category L (letter), M (mark), N (number), Pc (connector, including
	//  underscore), Pd (dash), and Zs (space). For more information, see [General
	//  Category](https://wikipedia.org/wiki/Unicode_character_property#General_Category).
	//  The maximum length is 1,024 characters.  Certain operations allow suffixing
	//  of the table Id with a partition decorator, such as
	//  `sample_table$20190123`.
	TableId *string `json:"tableId,omitempty"`
}

// +kcc:proto=google.protobuf.BoolValue
type BoolValue struct {
	// The bool value.
	Value *bool `json:"value,omitempty"`
}
