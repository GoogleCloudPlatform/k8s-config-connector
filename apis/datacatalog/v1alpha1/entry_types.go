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
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kcc:proto=google.cloud.datacatalog.v1.DataSource
type DataSourceObservedState struct {
	// Output only. Data Catalog entry name, if applicable.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataSource.source_entry
	SourceEntry *string `json:"sourceEntry,omitempty"`

	// Service that physically stores the data.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataSource.service
	Service *string `json:"service,omitempty"`

	// Full name of a resource as defined by the service. For example:
	//
	//  `//bigquery.googleapis.com/projects/{PROJECT_ID}/locations/{LOCATION}/datasets/{DATASET_ID}/tables/{TABLE_ID}`
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataSource.resource
	Resource *string `json:"resource,omitempty"`

	// Detailed properties of the underlying storage.
	// +kcc:proto:field=google.cloud.datacatalog.v1.DataSource.storage_properties
	StorageProperties *StorageProperties `json:"storageProperties,omitempty"`
}

// Parent defines the potential parent resources for a DataCatalogEntry.
type DataCatalogEntryParent struct {
	// Optional. Reference to the entry group that contains the entry.
	// +optional
	EntryGroupRef *EntryGroupRef `json:"entryGroupRef,omitempty"`

	// Optional. The location for the entry. Cannot be specified if entryGroupRef is specified.
	// +optional
	Location *string `json:"location,omitempty"`

	// Optional. Reference to the project that contains the entry. Cannot be specified if entryGroupRef is specified.
	// +optional
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef,omitempty"`
}

var DataCatalogEntryGVK = GroupVersion.WithKind("DataCatalogEntry")

// Copying over since we are editing the generated types

// +kcc:proto=google.cloud.datacatalog.v1.ColumnSchema
type ColumnSchema struct {
	// Required. Name of the column.
	//
	// Required. Name of the column.
	//
	//  Must be a UTF-8 string without dots (.).
	//  The maximum size is 64 bytes.
	// +required
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.column
	Column *string `json:"column,omitempty"`

	// Required. Type of the column.
	//
	//  Must be a UTF-8 string with the maximum size of 128 bytes.
	// +required
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.type
	Type *string `json:"type,omitempty"`

	// Optional. Description of the column. Default value is an empty string.
	//
	//  The description must be a UTF-8 string with the maximum size of 2000
	//  bytes.
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.description
	Description *string `json:"description,omitempty"`

	// Optional. A column's mode indicates whether values in this column are
	//  required, nullable, or repeated.
	//
	//  Only `NULLABLE`, `REQUIRED`, and `REPEATED` values are supported.
	//  Default mode is `NULLABLE`.
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.mode
	Mode *string `json:"mode,omitempty"`

	// Optional. Default value for the column.
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.default_value
	DefaultValue *string `json:"defaultValue,omitempty"`

	// Optional. Ordinal position
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.ordinal_position
	OrdinalPosition *int32 `json:"ordinalPosition,omitempty"`

	// Optional. Most important inclusion of this column.
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.highest_indexing_type
	HighestIndexingType *string `json:"highestIndexingType,omitempty"`

	// TODO: Known issue: recursive types are tripping CRD generation
	//  https://github.com/kubernetes-sigs/controller-tools/issues/489
	//  https://github.com/kubernetes-sigs/controller-tools/issues/585#issuecomment-968354281

	// Optional. Schema of sub-columns. A column can have zero or more
	//  sub-columns.
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.subcolumns
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Subcolumns []ColumnSchema `json:"subcolumns,omitempty"`

	// Looker specific column info of this column.
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.looker_column_spec
	LookerColumnSpec *ColumnSchema_LookerColumnSpec `json:"lookerColumnSpec,omitempty"`

	// Optional. The subtype of the RANGE, if the type of this field is RANGE. If
	//  the type is RANGE, this field is required. Possible values for the field
	//  element type of a RANGE include:
	//  * DATE
	//  * DATETIME
	//  * TIMESTAMP
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.range_element_type
	RangeElementType *ColumnSchema_FieldElementType `json:"rangeElementType,omitempty"`

	// Optional. Garbage collection policy for the column or column family.
	//  Applies to systems like Cloud Bigtable.
	// +kcc:proto:field=google.cloud.datacatalog.v1.ColumnSchema.gc_rule
	GcRule *string `json:"gcRule,omitempty"`
}

// DataCatalogEntrySpec defines the desired state of DataCatalogEntry
// +kcc:proto=google.cloud.datacatalog.v1.Entry
type DataCatalogEntrySpec struct {
	DataCatalogEntryParent `json:",inline"`

	// The DataCatalogEntry name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The resource this metadata entry refers to.
	//
	//  For Google Cloud Platform resources, `linked_resource` is the
	//  [Full Resource Name]
	//  (https://cloud.google.com/apis/design/resource_names#full_resource_name).
	//
	//  Output only when the entry is one of the types in the `EntryType` enum.
	//
	//  For entries with a `user_specified_type`, this field is optional and
	//  defaults to an empty string.
	//
	//  The resource string must contain only letters (a-z, A-Z), numbers (0-9),
	//  underscores (_), periods (.), colons (:), slashes (/), dashes (-),
	//  and hashes (#).
	//  The maximum size is 200 bytes when encoded in UTF-8.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.linked_resource
	LinkedResource *string `json:"linkedResource,omitempty"`

	// [Fully Qualified Name
	//  (FQN)](https://cloud.google.com//data-catalog/docs/fully-qualified-names)
	//  of the resource. Set automatically for entries representing resources from
	//  synced systems. Settable only during creation, and read-only later. Can
	//  be used for search and lookup of the entries.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.fully_qualified_name
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty"`

	// The type of the entry.
	//
	//  For details, see [`EntryType`](#entrytype).
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.type
	Type *string `json:"type,omitempty"`

	// Custom entry type that doesn't match any of the values allowed for input
	//  and listed in the `EntryType` enum.
	//
	//  When creating an entry, first check the type values in the enum.
	//  If there are no appropriate types for the new entry,
	//  provide a custom value, for example, `my_special_type`.
	//
	//  The `user_specified_type` string has the following limitations:
	//
	//  * Is case insensitive.
	//  * Must begin with a letter or underscore.
	//  * Can only contain letters, numbers, and underscores.
	//  * Must be at least 1 character and at most 64 characters long.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.user_specified_type
	UserSpecifiedType *string `json:"userSpecifiedType,omitempty"`

	// Indicates the entry's source system that Data Catalog doesn't
	//  automatically integrate with.
	//
	//  The `user_specified_system` string has the following limitations:
	//
	//  * Is case insensitive.
	//  * Must begin with a letter or underscore.
	//  * Can only contain letters, numbers, and underscores.
	//  * Must be at least 1 character and at most 64 characters long.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.user_specified_system
	UserSpecifiedSystem *string `json:"userSpecifiedSystem,omitempty"`

	// Specification that applies to a relational database system. Only settable
	//  when `user_specified_system` is equal to `SQL_DATABASE`
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.sql_database_system_spec
	SQLDatabaseSystemSpec *SQLDatabaseSystemSpec `json:"sqlDatabaseSystemSpec,omitempty"`

	// Specification that applies to Looker sysstem. Only settable when
	//  `user_specified_system` is equal to `LOOKER`
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.looker_system_spec
	LookerSystemSpec *LookerSystemSpec `json:"lookerSystemSpec,omitempty"`

	// Specification that applies to Cloud Bigtable system. Only settable when
	//  `integrated_system` is equal to `CLOUD_BIGTABLE`
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.cloud_bigtable_system_spec
	CloudBigtableSystemSpec *CloudBigtableSystemSpec `json:"cloudBigtableSystemSpec,omitempty"`

	// Specification that applies to a Cloud Storage fileset. Valid only
	//  for entries with the `FILESET` type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.gcs_fileset_spec
	GCSFilesetSpec *GCSFilesetSpec `json:"gcsFilesetSpec,omitempty"`

	// Specification that applies to a table resource. Valid only
	//  for entries with the `TABLE` or `EXPLORE` type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.database_table_spec
	DatabaseTableSpec *DatabaseTableSpec `json:"databaseTableSpec,omitempty"`

	// Specification that applies to a data source connection. Valid only
	//  for entries with the `DATA_SOURCE_CONNECTION` type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.data_source_connection_spec
	DataSourceConnectionSpec *DataSourceConnectionSpec `json:"dataSourceConnectionSpec,omitempty"`

	// Specification that applies to a user-defined function or procedure. Valid
	//  only for entries with the `ROUTINE` type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.routine_spec
	RoutineSpec *RoutineSpec `json:"routineSpec,omitempty"`

	// Specification that applies to a dataset.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.dataset_spec
	DatasetSpec *DatasetSpec `json:"datasetSpec,omitempty"`

	// Specification that applies to a fileset resource. Valid only
	//  for entries with the `FILESET` type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.fileset_spec
	FilesetSpec *FilesetSpec `json:"filesetSpec,omitempty"`

	// Specification that applies to a Service resource.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.service_spec
	ServiceSpec *ServiceSpec `json:"serviceSpec,omitempty"`

	// Model specification.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.model_spec
	ModelSpec *ModelSpec `json:"modelSpec,omitempty"`

	// FeatureonlineStore spec for Vertex AI Feature Store.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.feature_online_store_spec
	FeatureOnlineStoreSpec *FeatureOnlineStoreSpec `json:"featureOnlineStoreSpec,omitempty"`

	// Display name of an entry.
	//
	//  The maximum size is 500 bytes when encoded in UTF-8.
	//  Default value is an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Entry description that can consist of several sentences or paragraphs
	//  that describe entry contents.
	//
	//  The description must not contain Unicode non-characters as well as C0
	//  and C1 control codes except tabs (HT), new lines (LF), carriage returns
	//  (CR), and page breaks (FF).
	//  The maximum size is 2000 bytes when encoded in UTF-8.
	//  Default value is an empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.description
	Description *string `json:"description,omitempty"`

	// Business Context of the entry. Not supported for BigQuery datasets
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.business_context
	BusinessContext *BusinessContext `json:"businessContext,omitempty"`

	// Schema of the entry. An entry might not have any schema attached to it.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.schema
	Schema *Schema `json:"schema,omitempty"`

	// Timestamps from the underlying resource, not from the Data Catalog
	//  entry.
	//
	//  Output only when the entry has a system listed in the `IntegratedSystem`
	//  enum. For entries with `user_specified_system`, this field is optional
	//  and defaults to an empty timestamp.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.source_system_timestamps
	SourceSystemTimestamps *SystemTimestamps `json:"sourceSystemTimestamps,omitempty"`

	// Resource usage statistics.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.usage_signal
	UsageSignal *UsageSignal `json:"usageSignal,omitempty"`

	// Cloud labels attached to the entry.
	//
	//  In Data Catalog, you can create and modify labels attached only to custom
	//  entries. Synced entries have unmodifiable labels that come from the source
	//  system.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// DataCatalogEntryStatus defines the config connector machine state of DataCatalogEntry
type DataCatalogEntryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataCatalogEntry resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataCatalogEntryObservedState `json:"observedState,omitempty"`
}

// DataCatalogEntryObservedState is the state of the DataCatalogEntry resource as most recently observed in GCP.
// +kcc:proto=google.cloud.datacatalog.v1.Entry
type DataCatalogEntryObservedState struct {
	// Output only. Identifier. The resource name of an entry in URL format.
	//
	//  Note: The entry itself and its child resources might not be
	//  stored in the location specified in its name.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.name
	Name *string `json:"name,omitempty"`

	// Output only. Indicates the entry's source system that Data Catalog
	//  integrates with, such as BigQuery, Pub/Sub, or Dataproc Metastore.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.integrated_system
	IntegratedSystem *string `json:"integratedSystem,omitempty"`

	// Specification that applies to a Cloud Storage fileset. Valid only
	//  for entries with the `FILESET` type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.gcs_fileset_spec
	GCSFilesetSpec *GCSFilesetSpecObservedState `json:"gcsFilesetSpec,omitempty"`

	// Output only. Specification that applies to a BigQuery table. Valid only
	//  for entries with the `TABLE` type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.bigquery_table_spec
	BigqueryTableSpec *BigQueryTableSpecObservedState `json:"bigqueryTableSpec,omitempty"`

	// Output only. Specification for a group of BigQuery tables with
	//  the `[prefix]YYYYMMDD` name pattern.
	//
	//  For more information, see [Introduction to partitioned tables]
	//  (https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding).
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.bigquery_date_sharded_spec
	BigqueryDateShardedSpec *BigQueryDateShardedSpecObservedState `json:"bigqueryDateShardedSpec,omitempty"`

	// Specification that applies to a table resource. Valid only
	//  for entries with the `TABLE` or `EXPLORE` type.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.database_table_spec
	DatabaseTableSpec *DatabaseTableSpecObservedState `json:"databaseTableSpec,omitempty"`

	// FeatureonlineStore spec for Vertex AI Feature Store.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.feature_online_store_spec
	FeatureOnlineStoreSpec *FeatureOnlineStoreSpecObservedState `json:"featureOnlineStoreSpec,omitempty"`

	// Resource usage statistics.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.usage_signal
	UsageSignal *UsageSignalObservedState `json:"usageSignal,omitempty"`

	// Output only. Physical location of the entry.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.data_source
	DataSource *DataSourceObservedState `json:"dataSource,omitempty"`

	// Output only. Additional information related to the entry. Private to the
	//  current user.
	// +kcc:proto:field=google.cloud.datacatalog.v1.Entry.personal_details
	PersonalDetails *PersonalDetails `json:"personalDetails,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatacatalogentry;gcpdatacatalogentries
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataCatalogEntry is the Schema for the DataCatalogEntry API
// +k8s:openapi-gen=true
type DataCatalogEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataCatalogEntrySpec   `json:"spec,omitempty"`
	Status DataCatalogEntryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataCatalogEntryList contains a list of DataCatalogEntry
type DataCatalogEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalogEntry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataCatalogEntry{}, &DataCatalogEntryList{})
}
