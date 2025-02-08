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


// +kcc:proto=google.cloud.datastream.v1.BackfillJob
type BackfillJob struct {

	// Backfill job's triggering reason.
	// +kcc:proto:field=google.cloud.datastream.v1.BackfillJob.trigger
	Trigger *string `json:"trigger,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.Error
type Error struct {
	// A title that explains the reason for the error.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.reason
	Reason *string `json:"reason,omitempty"`

	// A unique identifier for this specific error,
	//  allowing it to be traced throughout the system in logs and API responses.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.error_uuid
	ErrorUuid *string `json:"errorUuid,omitempty"`

	// A message containing more information about the error that occurred.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.message
	Message *string `json:"message,omitempty"`

	// The time when the error occurred.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.error_time
	ErrorTime *string `json:"errorTime,omitempty"`

	// Additional information about the error.
	// +kcc:proto:field=google.cloud.datastream.v1.Error.details
	Details map[string]string `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SourceObjectIdentifier
type SourceObjectIdentifier struct {
	// Oracle data source object identifier.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceObjectIdentifier.oracle_identifier
	OracleIdentifier *SourceObjectIdentifier_OracleObjectIdentifier `json:"oracleIdentifier,omitempty"`

	// Mysql data source object identifier.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceObjectIdentifier.mysql_identifier
	MysqlIdentifier *SourceObjectIdentifier_MysqlObjectIdentifier `json:"mysqlIdentifier,omitempty"`

	// PostgreSQL data source object identifier.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceObjectIdentifier.postgresql_identifier
	PostgresqlIdentifier *SourceObjectIdentifier_PostgresqlObjectIdentifier `json:"postgresqlIdentifier,omitempty"`

	// SQLServer data source object identifier.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceObjectIdentifier.sql_server_identifier
	SqlServerIdentifier *SourceObjectIdentifier_SqlServerObjectIdentifier `json:"sqlServerIdentifier,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SourceObjectIdentifier.MysqlObjectIdentifier
type SourceObjectIdentifier_MysqlObjectIdentifier struct {
	// Required. The database name.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceObjectIdentifier.MysqlObjectIdentifier.database
	Database *string `json:"database,omitempty"`

	// Required. The table name.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceObjectIdentifier.MysqlObjectIdentifier.table
	Table *string `json:"table,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SourceObjectIdentifier.OracleObjectIdentifier
type SourceObjectIdentifier_OracleObjectIdentifier struct {
	// Required. The schema name.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceObjectIdentifier.OracleObjectIdentifier.schema
	Schema *string `json:"schema,omitempty"`

	// Required. The table name.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceObjectIdentifier.OracleObjectIdentifier.table
	Table *string `json:"table,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SourceObjectIdentifier.PostgresqlObjectIdentifier
type SourceObjectIdentifier_PostgresqlObjectIdentifier struct {
	// Required. The schema name.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceObjectIdentifier.PostgresqlObjectIdentifier.schema
	Schema *string `json:"schema,omitempty"`

	// Required. The table name.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceObjectIdentifier.PostgresqlObjectIdentifier.table
	Table *string `json:"table,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.SourceObjectIdentifier.SqlServerObjectIdentifier
type SourceObjectIdentifier_SqlServerObjectIdentifier struct {
	// Required. The schema name.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceObjectIdentifier.SqlServerObjectIdentifier.schema
	Schema *string `json:"schema,omitempty"`

	// Required. The table name.
	// +kcc:proto:field=google.cloud.datastream.v1.SourceObjectIdentifier.SqlServerObjectIdentifier.table
	Table *string `json:"table,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.StreamObject
type StreamObject struct {

	// Required. Display name.
	// +kcc:proto:field=google.cloud.datastream.v1.StreamObject.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The latest backfill job that was initiated for the stream object.
	// +kcc:proto:field=google.cloud.datastream.v1.StreamObject.backfill_job
	BackfillJob *BackfillJob `json:"backfillJob,omitempty"`

	// The object identifier in the data source.
	// +kcc:proto:field=google.cloud.datastream.v1.StreamObject.source_object
	SourceObject *SourceObjectIdentifier `json:"sourceObject,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.BackfillJob
type BackfillJobObservedState struct {
	// Output only. Backfill job state.
	// +kcc:proto:field=google.cloud.datastream.v1.BackfillJob.state
	State *string `json:"state,omitempty"`

	// Output only. Backfill job's start time.
	// +kcc:proto:field=google.cloud.datastream.v1.BackfillJob.last_start_time
	LastStartTime *string `json:"lastStartTime,omitempty"`

	// Output only. Backfill job's end time.
	// +kcc:proto:field=google.cloud.datastream.v1.BackfillJob.last_end_time
	LastEndTime *string `json:"lastEndTime,omitempty"`

	// Output only. Errors which caused the backfill job to fail.
	// +kcc:proto:field=google.cloud.datastream.v1.BackfillJob.errors
	Errors []Error `json:"errors,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.StreamObject
type StreamObjectObservedState struct {
	// Output only. The object resource's name.
	// +kcc:proto:field=google.cloud.datastream.v1.StreamObject.name
	Name *string `json:"name,omitempty"`

	// Output only. The creation time of the object.
	// +kcc:proto:field=google.cloud.datastream.v1.StreamObject.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update time of the object.
	// +kcc:proto:field=google.cloud.datastream.v1.StreamObject.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Active errors on the object.
	// +kcc:proto:field=google.cloud.datastream.v1.StreamObject.errors
	Errors []Error `json:"errors,omitempty"`

	// The latest backfill job that was initiated for the stream object.
	// +kcc:proto:field=google.cloud.datastream.v1.StreamObject.backfill_job
	BackfillJob *BackfillJobObservedState `json:"backfillJob,omitempty"`
}
