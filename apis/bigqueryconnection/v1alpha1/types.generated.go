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

package v1alpha1

// +kcc:proto=google.cloud.bigquery.connection.v1.AwsAccessRole
type AwsAccessRole struct {
	// The user’s AWS IAM Role that trusts the Google-owned AWS IAM user
	//  Connection.
	IamRoleID *string `json:"iamRoleID,omitempty"`

	// A unique Google-owned and Google-generated identity for the Connection.
	//  This identity will be used to access the user's AWS IAM Role.
	Identity *string `json:"identity,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.AwsCrossAccountRole
type AwsCrossAccountRole struct {
	// The user’s AWS IAM Role that trusts the Google-owned AWS IAM user
	//  Connection.
	IamRoleID *string `json:"iamRoleID,omitempty"`

	// Output only. Google-owned AWS IAM User for a Connection.
	IamUserID *string `json:"iamUserID,omitempty"`

	// Output only. A Google-generated id for representing Connection’s identity
	//  in AWS. External Id is also used for preventing the Confused Deputy
	//  Problem. See
	//  https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html
	ExternalID *string `json:"externalID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.AwsProperties
type AwsProperties struct {
	// Authentication using Google owned AWS IAM user's access key to assume
	//  into customer's AWS IAM Role.
	//  Deprecated, do not use.
	CrossAccountRole *AwsCrossAccountRole `json:"crossAccountRole,omitempty"`

	// Authentication using Google owned service account to assume into
	//  customer's AWS IAM Role.
	AccessRole *AwsAccessRole `json:"accessRole,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.AzureProperties
type AzureProperties struct {
	// Output only. The name of the Azure Active Directory Application.
	Application *string `json:"application,omitempty"`

	// Output only. The client id of the Azure Active Directory Application.
	ClientID *string `json:"clientID,omitempty"`

	// Output only. The object id of the Azure Active Directory Application.
	ObjectID *string `json:"objectID,omitempty"`

	// The id of customer's directory that host the data.
	CustomerTenantID *string `json:"customerTenantID,omitempty"`

	// The URL user will be redirected to after granting consent during connection
	//  setup.
	RedirectUri *string `json:"redirectUri,omitempty"`

	// The client ID of the user's Azure Active Directory Application used for a
	//  federated connection.
	FederatedApplicationClientID *string `json:"federatedApplicationClientID,omitempty"`

	// Output only. A unique Google-owned and Google-generated identity for the
	//  Connection. This identity will be used to access the user's Azure Active
	//  Directory Application.
	Identity *string `json:"identity,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.CloudResourceProperties
type CloudResourceProperties struct {
	// Output only. The account ID of the service created for the purpose of this
	//  connection.
	//
	//  The service account does not have any permissions associated with it
	//  when it is created. After creation, customers delegate permissions
	//  to the service account. When the connection is used in the context of an
	//  operation in BigQuery, the service account will be used to connect to the
	//  desired resources in GCP.
	//
	//  The account ID is in the form of:
	//    <service-1234>@gcp-sa-bigquery-cloudresource.iam.gserviceaccount.com
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.CloudSpannerProperties
type CloudSpannerProperties struct {
	// Cloud Spanner database in the form `project/instance/database'
	Database *string `json:"database,omitempty"`

	// If parallelism should be used when reading from Cloud Spanner
	UseParallelism *bool `json:"useParallelism,omitempty"`

	// Allows setting max parallelism per query when executing on Spanner
	//  independent compute resources. If unspecified, default values of
	//  parallelism are chosen that are dependent on the Cloud Spanner instance
	//  configuration.
	//
	//  REQUIRES: `use_parallelism` must be set.
	//  REQUIRES: Either `use_data_boost` or `use_serverless_analytics` must be
	//  set.
	MaxParallelism *int32 `json:"maxParallelism,omitempty"`

	// If the serverless analytics service should be used to read data from Cloud
	//  Spanner.
	//  Note: `use_parallelism` must be set when using serverless analytics.
	UseServerlessAnalytics *bool `json:"useServerlessAnalytics,omitempty"`

	// If set, the request will be executed via Spanner independent compute
	//  resources.
	//  REQUIRES: `use_parallelism` must be set.
	//
	//  NOTE: `use_serverless_analytics` will be deprecated. Prefer
	//  `use_data_boost` over `use_serverless_analytics`.
	UseDataBoost *bool `json:"useDataBoost,omitempty"`

	// Optional. Cloud Spanner database role for fine-grained access control.
	//  The Cloud Spanner admin should have provisioned the database role with
	//  appropriate permissions, such as `SELECT` and `INSERT`. Other users should
	//  only use roles provided by their Cloud Spanner admins.
	//
	//  For more details, see [About fine-grained access control]
	//  (https://cloud.google.com/spanner/docs/fgac-about).
	//
	//  REQUIRES: The database role name must start with a letter, and can only
	//  contain letters, numbers, and underscores.
	DatabaseRole *string `json:"databaseRole,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.CloudSqlProperties
type CloudSqlProperties struct {
	// Cloud SQL instance ID in the form `project:location:instance`.
	InstanceID *string `json:"instanceID,omitempty"`

	// Database name.
	Database *string `json:"database,omitempty"`

	// Type of the Cloud SQL database.
	Type *string `json:"type,omitempty"`

	// Input only. Cloud SQL credential.
	Credential *CloudSqlCredential `json:"credential,omitempty"`

	// Output only. The account ID of the service used for the purpose of this
	//  connection.
	//
	//  When the connection is used in the context of an operation in
	//  BigQuery, this service account will serve as the identity being used for
	//  connecting to the CloudSQL instance specified in this connection.
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.Connection
type Connection struct {
	// The resource name of the connection in the form of:
	//  `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
	Name *string `json:"name,omitempty"`

	// User provided display name for the connection.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// User provided description.
	Description *string `json:"description,omitempty"`

	// Cloud SQL properties.
	CloudSql *CloudSqlProperties `json:"cloudSql,omitempty"`

	// Amazon Web Services (AWS) properties.
	Aws *AwsProperties `json:"aws,omitempty"`

	// Azure properties.
	Azure *AzureProperties `json:"azure,omitempty"`

	// Cloud Spanner properties.
	CloudSpanner *CloudSpannerProperties `json:"cloudSpanner,omitempty"`

	// Cloud Resource properties.
	CloudResource *CloudResourceProperties `json:"cloudResource,omitempty"`

	// Spark properties.
	Spark *SparkProperties `json:"spark,omitempty"`

	// Optional. Salesforce DataCloud properties. This field is intended for
	//  use only by Salesforce partner projects. This field contains properties
	//  for your Salesforce DataCloud connection.
	SalesforceDataCloud *SalesforceDataCloudProperties `json:"salesforceDataCloud,omitempty"`

	// Output only. The creation timestamp of the connection.
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Output only. The last update timestamp of the connection.
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`

	// Output only. True, if credential is configured for this connection.
	HasCredential *bool `json:"hasCredential,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.MetastoreServiceConfig
type MetastoreServiceConfig struct {
	// Optional. Resource name of an existing Dataproc Metastore service.
	//
	//  Example:
	//
	//  * `projects/[project_id]/locations/[region]/services/[service_id]`
	MetastoreService *string `json:"metastoreService,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.SalesforceDataCloudProperties
type SalesforceDataCloudProperties struct {
	// The URL to the user's Salesforce DataCloud instance.
	InstanceUri *string `json:"instanceUri,omitempty"`

	// Output only. A unique Google-owned and Google-generated service account
	//  identity for the connection.
	Identity *string `json:"identity,omitempty"`

	// The ID of the user's Salesforce tenant.
	TenantID *string `json:"tenantID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.SparkHistoryServerConfig
type SparkHistoryServerConfig struct {
	// Optional. Resource name of an existing Dataproc Cluster to act as a Spark
	//  History Server for the connection.
	//
	//  Example:
	//
	//  * `projects/[project_id]/regions/[region]/clusters/[cluster_name]`
	DataprocCluster *string `json:"dataprocCluster,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.SparkProperties
type SparkProperties struct {
	// Output only. The account ID of the service created for the purpose of this
	//  connection.
	//
	//  The service account does not have any permissions associated with it when
	//  it is created. After creation, customers delegate permissions to the
	//  service account. When the connection is used in the context of a stored
	//  procedure for Apache Spark in BigQuery, the service account is used to
	//  connect to the desired resources in Google Cloud.
	//
	//  The account ID is in the form of:
	//  bqcx-<projectnumber>-<uniqueid>@gcp-sa-bigquery-consp.iam.gserviceaccount.com
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`

	// Optional. Dataproc Metastore Service configuration for the connection.
	MetastoreServiceConfig *MetastoreServiceConfig `json:"metastoreServiceConfig,omitempty"`

	// Optional. Spark History Server configuration for the connection.
	SparkHistoryServerConfig *SparkHistoryServerConfig `json:"sparkHistoryServerConfig,omitempty"`
}
