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


// +kcc:proto=google.cloud.bigquery.connection.v1beta1.CloudSqlCredential
type CloudSqlCredential struct {
	// The username for the credential.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.CloudSqlCredential.username
	Username *string `json:"username,omitempty"`

	// The password for the credential.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.CloudSqlCredential.password
	Password *string `json:"password,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1beta1.CloudSqlProperties
type CloudSqlProperties struct {
	// Cloud SQL instance ID in the form `project:location:instance`.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.CloudSqlProperties.instance_id
	InstanceID *string `json:"instanceID,omitempty"`

	// Database name.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.CloudSqlProperties.database
	Database *string `json:"database,omitempty"`

	// Type of the Cloud SQL database.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.CloudSqlProperties.type
	Type *string `json:"type,omitempty"`

	// Input only. Cloud SQL credential.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.CloudSqlProperties.credential
	Credential *CloudSqlCredential `json:"credential,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1beta1.Connection
type Connection struct {
	// The resource name of the connection in the form of:
	//  `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.Connection.name
	Name *string `json:"name,omitempty"`

	// User provided display name for the connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.Connection.friendly_name
	FriendlyName *string `json:"friendlyName,omitempty"`

	// User provided description.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.Connection.description
	Description *string `json:"description,omitempty"`

	// Cloud SQL properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.Connection.cloud_sql
	CloudSql *CloudSqlProperties `json:"cloudSql,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1beta1.CloudSqlProperties
type CloudSqlPropertiesObservedState struct {
	// Output only. The account ID of the service used for the purpose of this
	//  connection.
	//
	//  When the connection is used in the context of an operation in
	//  BigQuery, this service account will serve as the identity being used for
	//  connecting to the CloudSQL instance specified in this connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.CloudSqlProperties.service_account_id
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1beta1.Connection
type ConnectionObservedState struct {
	// Cloud SQL properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.Connection.cloud_sql
	CloudSql *CloudSqlPropertiesObservedState `json:"cloudSql,omitempty"`

	// Output only. The creation timestamp of the connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.Connection.creation_time
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Output only. The last update timestamp of the connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.Connection.last_modified_time
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`

	// Output only. True, if credential is configured for this connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1beta1.Connection.has_credential
	HasCredential *bool `json:"hasCredential,omitempty"`
}
