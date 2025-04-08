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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DatastreamStreamGVK = GroupVersion.WithKind("DatastreamStream")

// DatastreamStreamSpec defines the desired state of DatastreamStream
// +kcc:proto=google.cloud.datastream.v1.Stream
type DatastreamStreamSpec struct {
	// The DatastreamStream name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Labels.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Display name.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Source connection profile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.source_config
	SourceConfig *SourceConfig `json:"sourceConfig,omitempty"`

	// Required. Destination connection profile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.destination_config
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty"`

	// The state of the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.state
	State *string `json:"state,omitempty"`

	// Automatically backfill objects included in the stream source
	//  configuration. Specific objects can be excluded.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.backfill_all
	BackfillAll *Stream_BackfillAllStrategy `json:"backfillAll,omitempty"`

	// Do not automatically backfill any objects.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.backfill_none
	BackfillNone *Stream_BackfillNoneStrategy `json:"backfillNone,omitempty"`

	// Immutable. A reference to a KMS encryption key.
	//  If provided, it will be used to encrypt the data.
	//  If left blank, data will be encrypted using an internal Stream-specific
	//  encryption key provisioned through KMS.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.customer_managed_encryption_key
	CustomerManagedEncryptionKey *string `json:"customerManagedEncryptionKey,omitempty"`
}

// DatastreamStreamStatus defines the config connector machine state of DatastreamStream
type DatastreamStreamStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DatastreamStream resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DatastreamStreamObservedState `json:"observedState,omitempty"`
}

// DatastreamStreamObservedState is the state of the DatastreamStream resource as most recently observed in GCP.
// +kcc:proto=google.cloud.datastream.v1.Stream
type DatastreamStreamObservedState struct {
	// Output only. The stream's name.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. The creation time of the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update time of the stream.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Errors on the Stream.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.errors
	Errors []Error `json:"errors,omitempty"`

	// Output only. If the stream was recovered, the time of the last recovery.
	//  Note: This field is currently experimental.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.last_recovery_time
	LastRecoveryTime *string `json:"lastRecoveryTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatastreamstream;gcpdatastreamstreams
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DatastreamStream is the Schema for the DatastreamStream API
// +k8s:openapi-gen=true
type DatastreamStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DatastreamStreamSpec   `json:"spec,omitempty"`
	Status DatastreamStreamStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DatastreamStreamList contains a list of DatastreamStream
type DatastreamStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatastreamStream `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatastreamStream{}, &DatastreamStreamList{})
}

// +kcc:proto=google.cloud.datastream.v1.SourceConfig
type SourceConfig struct {
	// Required. Source connection profile resoource.
	//  Format: `projects/{project}/locations/{location}/connectionProfiles/{name}`
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.source_connection_profile
	SourceConnectionProfile *string `json:"sourceConnectionProfile,omitempty"`

	// Oracle data source configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.oracle_source_config
	OracleSourceConfig *OracleSourceConfig `json:"oracleSourceConfig,omitempty"`

	// MySQL data source configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.mysql_source_config
	MySQLSourceConfig *MysqlSourceConfig `json:"mySQLSourceConfig,omitempty"`

	// PostgreSQL data source configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.postgresql_source_config
	PostgreSQLSourceConfig *PostgresqlSourceConfig `json:"postgreSQLSourceConfig,omitempty"`

	// SQLServer data source configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceConfig.sql_server_source_config
	SQLServerSourceConfig *SQLServerSourceConfig `json:"sqlServerSourceConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.Stream.BackfillAllStrategy
type Stream_BackfillAllStrategy struct {
	// Oracle data source objects to avoid backfilling.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.BackfillAllStrategy.oracle_excluded_objects
	OracleExcludedObjects *OracleRdbms `json:"oracleExcludedObjects,omitempty"`

	// MySQL data source objects to avoid backfilling.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.BackfillAllStrategy.mysql_excluded_objects
	MySQLExcludedObjects *MysqlRdbms `json:"mySQLExcludedObjects,omitempty"`

	// PostgreSQL data source objects to avoid backfilling.
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.BackfillAllStrategy.postgresql_excluded_objects
	PostgreSQLExcludedObjects *PostgresqlRdbms `json:"postgreSQLExcludedObjects,omitempty"`

	// SQLServer data source objects to avoid backfilling
	// +kcc:proto:field=google.cloud.datastream.v1.Stream.BackfillAllStrategy.sql_server_excluded_objects
	SQLServerExcludedObjects *SQLServerRdbms `json:"sqlServerExcludedObjects,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlRdbms
type MysqlRdbms struct {
	// Mysql databases on the server
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlRdbms.mysql_databases
	MySQLDatabases []MySQLDatabase `json:"mySQLDatabases,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlDatabase
type MySQLDatabase struct {
	// Database name.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlDatabase.database
	Database *string `json:"database,omitempty"`

	// Tables in the database.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlDatabase.mysql_tables
	MySQLTables []MySQLTable `json:"mySQLTables,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlTable
type MySQLTable struct {
	// Table name.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlTable.table
	Table *string `json:"table,omitempty"`

	// MySQL columns in the database.
	//  When unspecified as part of include/exclude objects, includes/excludes
	//  everything.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlTable.mysql_columns
	MySQLColumns []MySQLColumn `json:"mySQLColumns,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlColumn
type MySQLColumn struct {
	// Column name.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.column
	Column *string `json:"column,omitempty"`

	// The MySQL data type. Full data types list can be found here:
	//  https://dev.mysql.com/doc/refman/8.0/en/data-types.html
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.data_type
	DataType *string `json:"dataType,omitempty"`

	// Column length.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.length
	Length *int32 `json:"length,omitempty"`

	// Column collation.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.collation
	Collation *string `json:"collation,omitempty"`

	// Whether or not the column represents a primary key.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.primary_key
	PrimaryKey *bool `json:"primaryKey,omitempty"`

	// Whether or not the column can accept a null value.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.nullable
	Nullable *bool `json:"nullable,omitempty"`

	// The ordinal position of the column in the table.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.ordinal_position
	OrdinalPosition *int32 `json:"ordinalPosition,omitempty"`

	// Column precision.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.precision
	Precision *int32 `json:"precision,omitempty"`

	// Column scale.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlColumn.scale
	Scale *int32 `json:"scale,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlRdbms
type PostgresqlRdbms struct {
	// PostgreSQL schemas in the database server.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlRdbms.postgresql_schemas
	PostgreSQLSchemas []PostgreSQLSchema `json:"postgreSQLSchemas,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlSchema
type PostgreSQLSchema struct {
	// Schema name.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSchema.schema
	Schema *string `json:"schema,omitempty"`

	// Tables in the schema.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlSchema.postgresql_tables
	PostgreSQLTables []PostgreSQLTable `json:"postgreSQLTables,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlTable
type PostgreSQLTable struct {
	// Table name.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlTable.table
	Table *string `json:"table,omitempty"`

	// PostgreSQL columns in the schema.
	//  When unspecified as part of include/exclude objects,
	//  includes/excludes everything.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlTable.postgresql_columns
	PostgreSQLColumns []PostgreSQLColumn `json:"postgreSQLColumns,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlColumn
type PostgreSQLColumn struct {
	// Column name.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.column
	Column *string `json:"column,omitempty"`

	// The PostgreSQL data type.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.data_type
	DataType *string `json:"dataType,omitempty"`

	// Column length.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.length
	Length *int32 `json:"length,omitempty"`

	// Column precision.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.precision
	Precision *int32 `json:"precision,omitempty"`

	// Column scale.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.scale
	Scale *int32 `json:"scale,omitempty"`

	// Whether or not the column represents a primary key.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.primary_key
	PrimaryKey *bool `json:"primaryKey,omitempty"`

	// Whether or not the column can accept a null value.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.nullable
	Nullable *bool `json:"nullable,omitempty"`

	// The ordinal position of the column in the table.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlColumn.ordinal_position
	OrdinalPosition *int32 `json:"ordinalPosition,omitempty"`
}
