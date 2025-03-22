// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigqueryTableGVK = GroupVersion.WithKind("BigqueryTable")

// BigqueryTableSpec defines the desired state of BigqueryTable
// +kcc:proto=google.cloud.bigquery.v2.Table
type BigqueryTableSpec struct {
	// The GCP resource identifier. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The dataset the table belongs to.
	DatasetRef *refsv1beta1.DatasetRef `json:"datasetRef,omitempty"`

	// The type of resource ID.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.kind
	Kind *string `json:"kind,omitempty"`

	// Optional. A descriptive name for this table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.friendly_name
	FriendlyName *string `json:"friendlyName,omitempty"`

	// Optional. A user-friendly description of this table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.description
	Description *string `json:"description,omitempty"`

	// The labels associated with this table. You can use these to organize and
	//  group your tables. Label keys and values can be no longer than 63
	//  characters, can only contain lowercase letters, numeric characters,
	//  underscores and dashes. International characters are allowed. Label values
	//  are optional. Label keys must start with a letter and each label in the
	//  list must have a different key.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Describes the schema of this table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.schema
	Schema *TableSchema `json:"schema,omitempty"`

	// If specified, configures time-based partitioning for this table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.time_partitioning
	TimePartitioning *TimePartitioning `json:"timePartitioning,omitempty"`

	// If specified, configures range partitioning for this table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.range_partitioning
	RangePartitioning *RangePartitioning `json:"rangePartitioning,omitempty"`

	// Clustering specification for the table. Must be specified with time-based
	//  partitioning, data in the table will be first partitioned and subsequently
	//  clustered.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.clustering
	Clustering *Clustering `json:"clustering,omitempty"`

	// Optional. If set to true, queries over this table require
	//  a partition filter that can be used for partition elimination to be
	//  specified.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.require_partition_filter
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty"`

	// Optional. The partition information for all table formats, including
	//  managed partitioned tables, hive partitioned tables, iceberg partitioned,
	//  and metastore partitioned tables. This field is only populated for
	//  metastore partitioned tables. For other table formats, this is an output
	//  only field.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.partition_definition
	PartitionDefinition *PartitioningDefinition `json:"partitionDefinition,omitempty"`

	// Optional. The time when this table expires, in milliseconds since the
	//  epoch. If not present, the table will persist indefinitely. Expired tables
	//  will be deleted and their storage reclaimed.  The defaultTableExpirationMs
	//  property of the encapsulating dataset can be used to set a default
	//  expirationTime on newly created tables.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.expiration_time
	ExpirationTime *int64 `json:"expirationTime,omitempty"`

	// Optional. The view definition.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.view
	View *ViewDefinition `json:"view,omitempty"`

	// Optional. The materialized view definition.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.materialized_view
	MaterializedView *MaterializedViewDefinition `json:"materializedView,omitempty"`

	// Optional. Describes the data format, location, and other properties of
	//  a table stored outside of BigQuery. By defining these properties, the data
	//  source can then be queried as if it were a standard BigQuery table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.external_data_configuration
	ExternalDataConfiguration *ExternalDataConfiguration `json:"externalDataConfiguration,omitempty"`

	// Optional. Specifies the configuration of a BigLake managed table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.biglake_configuration
	BiglakeConfiguration *BigLakeConfiguration `json:"biglakeConfiguration,omitempty"`

	// Custom encryption configuration (e.g., Cloud KMS keys).
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.encryption_configuration
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`

	// Optional. Defines the default collation specification of new STRING fields
	//  in the table. During table creation or update, if a STRING field is added
	//  to this table without explicit collation specified, then the table inherits
	//  the table default collation. A change to this field affects only fields
	//  added afterwards, and does not alter the existing fields.
	//  The following values are supported:
	//
	//  * 'und:ci': undetermined locale, case insensitive.
	//  * '': empty string. Default to case-sensitive behavior.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.default_collation
	DefaultCollation *string `json:"defaultCollation,omitempty"`

	// Optional. Defines the default rounding mode specification of new decimal
	//  fields (NUMERIC OR BIGNUMERIC) in the table. During table creation or
	//  update, if a decimal field is added to this table without an explicit
	//  rounding mode specified, then the field inherits the table default
	//  rounding mode. Changing this field doesn't affect existing fields.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.default_rounding_mode
	DefaultRoundingMode *string `json:"defaultRoundingMode,omitempty"`

	// Optional. The maximum staleness of data that could be returned when the
	//  table (or stale MV) is queried. Staleness encoded as a string encoding
	//  of sql IntervalValue type.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.max_staleness
	MaxStaleness *string `json:"maxStaleness,omitempty"`

	// Optional. Tables Primary Key and Foreign Key information
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.table_constraints
	TableConstraints *TableConstraints `json:"tableConstraints,omitempty"`

	// Optional. The [tags](https://cloud.google.com/bigquery/docs/tags) attached
	//  to this table. Tag keys are globally unique. Tag key is expected to be in
	//  the namespaced format, for example "123456789012/environment" where
	//  123456789012 is the ID of the parent organization or project resource for
	//  this tag key. Tag value is expected to be the short name, for example
	//  "Production". See [Tag
	//  definitions](https://cloud.google.com/iam/docs/tags-access-control#definitions)
	//  for more details.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.resource_tags
	ResourceTags map[string]string `json:"resourceTags,omitempty"`

	// Optional. Table replication info for table created `AS REPLICA` DDL like:
	//  `CREATE MATERIALIZED VIEW mv1 AS REPLICA OF src_mv`
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.table_replication_info
	TableReplicationInfo *TableReplicationInfo `json:"tableReplicationInfo,omitempty"`

	// Optional. Options defining open source compatible table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.external_catalog_table_options
	ExternalCatalogTableOptions *ExternalCatalogTableOptions `json:"externalCatalogTableOptions,omitempty"`
}

// BigqueryTableStatus defines the config connector machine state of BigqueryTable
type BigqueryTableStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigqueryTable resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigqueryTableObservedState `json:"observedState,omitempty"`
}

// BigqueryTableObservedState is the state of the BigqueryTable resource as most recently observed in GCP.
// +kcc:proto=google.cloud.bigquery.v2.Table
type BigqueryTableObservedState struct {
	// Output only. A hash of this resource.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. An opaque ID uniquely identifying the table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.id
	ID *string `json:"id,omitempty"`

	// Output only. A URL that can be used to access this resource again.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. The size of this table in logical bytes, excluding any data in
	//  the streaming buffer.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_bytes
	NumBytes *int64 `json:"numBytes,omitempty"`

	// Output only. The physical size of this table in bytes. This includes
	//  storage used for time travel.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_physical_bytes
	NumPhysicalBytes *int64 `json:"numPhysicalBytes,omitempty"`

	// Output only. The number of logical bytes in the table that are considered
	//  "long-term storage".
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_long_term_bytes
	NumLongTermBytes *int64 `json:"numLongTermBytes,omitempty"`

	// Output only. The number of rows of data in this table, excluding any data
	//  in the streaming buffer.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_rows
	NumRows *UInt64Value `json:"numRows,omitempty"`

	// Output only. The time when this table was created, in milliseconds since
	//  the epoch.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.creation_time
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Output only. The time when this table was last modified, in milliseconds
	//  since the epoch.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.last_modified_time
	LastModifiedTime *uint64 `json:"lastModifiedTime,omitempty"`

	// Output only. Describes the table type. The following values are supported:
	//
	//  * `TABLE`: A normal BigQuery table.
	//  * `VIEW`: A virtual table defined by a SQL query.
	//  * `EXTERNAL`: A table that references data stored in an external storage
	//    system, such as Google Cloud Storage.
	//  * `MATERIALIZED_VIEW`: A precomputed view defined by a SQL query.
	//  * `SNAPSHOT`: An immutable BigQuery table that preserves the contents of a
	//    base table at a particular time. See additional information on
	//    [table
	//    snapshots](https://cloud.google.com/bigquery/docs/table-snapshots-intro).
	//
	//  The default value is `TABLE`.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.type
	Type *string `json:"type,omitempty"`

	// Optional. The view definition.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.view
	View *ViewDefinitionObservedState `json:"view,omitempty"`

	// Optional. The materialized view definition.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.materialized_view
	MaterializedView *MaterializedViewDefinitionObservedState `json:"materializedView,omitempty"`

	// Output only. The materialized view status.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.materialized_view_status
	MaterializedViewStatus *MaterializedViewStatus `json:"materializedViewStatus,omitempty"`

	// Optional. Describes the data format, location, and other properties of
	//  a table stored outside of BigQuery. By defining these properties, the data
	//  source can then be queried as if it were a standard BigQuery table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.external_data_configuration
	ExternalDataConfiguration *ExternalDataConfigurationObservedState `json:"externalDataConfiguration,omitempty"`

	// Output only. The geographic location where the table resides. This value
	//  is inherited from the dataset.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.location
	Location *string `json:"location,omitempty"`

	// Output only. Contains information regarding this table's streaming buffer,
	//  if one is present. This field will be absent if the table is not being
	//  streamed to or if there is no data in the streaming buffer.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.streaming_buffer
	StreamingBuffer *Streamingbuffer `json:"streamingBuffer,omitempty"`

	// Output only. Contains information about the snapshot. This value is set via
	//  snapshot creation.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.snapshot_definition
	SnapshotDefinition *SnapshotDefinition `json:"snapshotDefinition,omitempty"`

	// Output only. Contains information about the clone. This value is set via
	//  the clone operation.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.clone_definition
	CloneDefinition *CloneDefinition `json:"cloneDefinition,omitempty"`

	// Output only. Number of physical bytes used by time travel storage (deleted
	//  or changed data). This data is not kept in real time, and might be delayed
	//  by a few seconds to a few minutes.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_time_travel_physical_bytes
	NumTimeTravelPhysicalBytes *int64 `json:"numTimeTravelPhysicalBytes,omitempty"`

	// Output only. Total number of logical bytes in the table or materialized
	//  view.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_total_logical_bytes
	NumTotalLogicalBytes *int64 `json:"numTotalLogicalBytes,omitempty"`

	// Output only. Number of logical bytes that are less than 90 days old.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_active_logical_bytes
	NumActiveLogicalBytes *int64 `json:"numActiveLogicalBytes,omitempty"`

	// Output only. Number of logical bytes that are more than 90 days old.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_long_term_logical_bytes
	NumLongTermLogicalBytes *int64 `json:"numLongTermLogicalBytes,omitempty"`

	// Output only. Number of physical bytes used by current live data storage.
	//  This data is not kept in real time, and might be delayed by a few seconds
	//  to a few minutes.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_current_physical_bytes
	NumCurrentPhysicalBytes *int64 `json:"numCurrentPhysicalBytes,omitempty"`

	// Output only. The physical size of this table in bytes. This also includes
	//  storage used for time travel. This data is not kept in real time, and might
	//  be delayed by a few seconds to a few minutes.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_total_physical_bytes
	NumTotalPhysicalBytes *int64 `json:"numTotalPhysicalBytes,omitempty"`

	// Output only. Number of physical bytes less than 90 days old. This data is
	//  not kept in real time, and might be delayed by a few seconds to a few
	//  minutes.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_active_physical_bytes
	NumActivePhysicalBytes *int64 `json:"numActivePhysicalBytes,omitempty"`

	// Output only. Number of physical bytes more than 90 days old.
	//  This data is not kept in real time, and might be delayed by a few seconds
	//  to a few minutes.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_long_term_physical_bytes
	NumLongTermPhysicalBytes *int64 `json:"numLongTermPhysicalBytes,omitempty"`

	// Output only. The number of partitions present in the table or materialized
	//  view. This data is not kept in real time, and might be delayed by a few
	//  seconds to a few minutes.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.num_partitions
	NumPartitions *int64 `json:"numPartitions,omitempty"`

	// Optional. Output only. Restriction config for table. If set, restrict
	//  certain accesses on the table based on the config. See [Data
	//  egress](https://cloud.google.com/bigquery/docs/analytics-hub-introduction#data_egress)
	//  for more details.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.restrictions
	Restrictions *RestrictionConfig `json:"restrictions,omitempty"`

	// Optional. Table replication info for table created `AS REPLICA` DDL like:
	//  `CREATE MATERIALIZED VIEW mv1 AS REPLICA OF src_mv`
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.table_replication_info
	TableReplicationInfo *TableReplicationInfoObservedState `json:"tableReplicationInfo,omitempty"`

	// Optional. Output only. Table references of all replicas currently active on
	//  the table.
	// +kcc:proto:field=google.cloud.bigquery.v2.Table.replicas
	Replicas []TableReference `json:"replicas,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpbigquerytable;gcpbigquerytables
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigqueryTable is the Schema for the BigqueryTable API
// +k8s:openapi-gen=true
type BigqueryTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigqueryTableSpec   `json:"spec,omitempty"`
	Status BigqueryTableStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigqueryTableList contains a list of BigqueryTable
type BigqueryTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigqueryTable `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigqueryTable{}, &BigqueryTableList{})
}
