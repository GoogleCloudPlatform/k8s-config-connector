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


// +kcc:proto=google.cloud.bigquery.connection.v1.AwsAccessRole
type AwsAccessRole struct {
	// The user’s AWS IAM Role that trusts the Google-owned AWS IAM user
	//  Connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AwsAccessRole.iam_role_id
	IamRoleID *string `json:"iamRoleID,omitempty"`

	// A unique Google-owned and Google-generated identity for the Connection.
	//  This identity will be used to access the user's AWS IAM Role.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AwsAccessRole.identity
	Identity *string `json:"identity,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.AwsCrossAccountRole
type AwsCrossAccountRole struct {
	// The user’s AWS IAM Role that trusts the Google-owned AWS IAM user
	//  Connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AwsCrossAccountRole.iam_role_id
	IamRoleID *string `json:"iamRoleID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.AwsProperties
type AwsProperties struct {
	// Authentication using Google owned AWS IAM user's access key to assume
	//  into customer's AWS IAM Role.
	//  Deprecated, do not use.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AwsProperties.cross_account_role
	CrossAccountRole *AwsCrossAccountRole `json:"crossAccountRole,omitempty"`

	// Authentication using Google owned service account to assume into
	//  customer's AWS IAM Role.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AwsProperties.access_role
	AccessRole *AwsAccessRole `json:"accessRole,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.AzureProperties
type AzureProperties struct {

	// The id of customer's directory that host the data.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AzureProperties.customer_tenant_id
	CustomerTenantID *string `json:"customerTenantID,omitempty"`

	// The URL user will be redirected to after granting consent during connection
	//  setup.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AzureProperties.redirect_uri
	RedirectURI *string `json:"redirectURI,omitempty"`

	// The client ID of the user's Azure Active Directory Application used for a
	//  federated connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AzureProperties.federated_application_client_id
	FederatedApplicationClientID *string `json:"federatedApplicationClientID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.CloudResourceProperties
type CloudResourceProperties struct {
}

// +kcc:proto=google.cloud.bigquery.connection.v1.CloudSpannerProperties
type CloudSpannerProperties struct {
	// Cloud Spanner database in the form `project/instance/database'
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSpannerProperties.database
	Database *string `json:"database,omitempty"`

	// If parallelism should be used when reading from Cloud Spanner
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSpannerProperties.use_parallelism
	UseParallelism *bool `json:"useParallelism,omitempty"`

	// Allows setting max parallelism per query when executing on Spanner
	//  independent compute resources. If unspecified, default values of
	//  parallelism are chosen that are dependent on the Cloud Spanner instance
	//  configuration.
	//
	//  REQUIRES: `use_parallelism` must be set.
	//  REQUIRES: Either `use_data_boost` or `use_serverless_analytics` must be
	//  set.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSpannerProperties.max_parallelism
	MaxParallelism *int32 `json:"maxParallelism,omitempty"`

	// If the serverless analytics service should be used to read data from Cloud
	//  Spanner.
	//  Note: `use_parallelism` must be set when using serverless analytics.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSpannerProperties.use_serverless_analytics
	UseServerlessAnalytics *bool `json:"useServerlessAnalytics,omitempty"`

	// If set, the request will be executed via Spanner independent compute
	//  resources.
	//  REQUIRES: `use_parallelism` must be set.
	//
	//  NOTE: `use_serverless_analytics` will be deprecated. Prefer
	//  `use_data_boost` over `use_serverless_analytics`.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSpannerProperties.use_data_boost
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
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSpannerProperties.database_role
	DatabaseRole *string `json:"databaseRole,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.CloudSqlCredential
type CloudSqlCredential struct {
	// The username for the credential.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSqlCredential.username
	Username *string `json:"username,omitempty"`

	// The password for the credential.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSqlCredential.password
	Password *string `json:"password,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.CloudSqlProperties
type CloudSqlProperties struct {
	// Cloud SQL instance ID in the form `project:location:instance`.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSqlProperties.instance_id
	InstanceID *string `json:"instanceID,omitempty"`

	// Database name.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSqlProperties.database
	Database *string `json:"database,omitempty"`

	// Type of the Cloud SQL database.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSqlProperties.type
	Type *string `json:"type,omitempty"`

	// Input only. Cloud SQL credential.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSqlProperties.credential
	Credential *CloudSqlCredential `json:"credential,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.Connection
type Connection struct {
	// The resource name of the connection in the form of:
	//  `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.name
	Name *string `json:"name,omitempty"`

	// User provided display name for the connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.friendly_name
	FriendlyName *string `json:"friendlyName,omitempty"`

	// User provided description.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.description
	Description *string `json:"description,omitempty"`

	// Cloud SQL properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.cloud_sql
	CloudSql *CloudSqlProperties `json:"cloudSql,omitempty"`

	// Amazon Web Services (AWS) properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.aws
	Aws *AwsProperties `json:"aws,omitempty"`

	// Azure properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.azure
	Azure *AzureProperties `json:"azure,omitempty"`

	// Cloud Spanner properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.cloud_spanner
	CloudSpanner *CloudSpannerProperties `json:"cloudSpanner,omitempty"`

	// Cloud Resource properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.cloud_resource
	CloudResource *CloudResourceProperties `json:"cloudResource,omitempty"`

	// Spark properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.spark
	Spark *SparkProperties `json:"spark,omitempty"`

	// Optional. Salesforce DataCloud properties. This field is intended for
	//  use only by Salesforce partner projects. This field contains properties
	//  for your Salesforce DataCloud connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.salesforce_data_cloud
	SalesforceDataCloud *SalesforceDataCloudProperties `json:"salesforceDataCloud,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.MetastoreServiceConfig
type MetastoreServiceConfig struct {
	// Optional. Resource name of an existing Dataproc Metastore service.
	//
	//  Example:
	//
	//  * `projects/[project_id]/locations/[region]/services/[service_id]`
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.MetastoreServiceConfig.metastore_service
	MetastoreService *string `json:"metastoreService,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.SalesforceDataCloudProperties
type SalesforceDataCloudProperties struct {
	// The URL to the user's Salesforce DataCloud instance.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.SalesforceDataCloudProperties.instance_uri
	InstanceURI *string `json:"instanceURI,omitempty"`

	// The ID of the user's Salesforce tenant.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.SalesforceDataCloudProperties.tenant_id
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
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.SparkHistoryServerConfig.dataproc_cluster
	DataprocCluster *string `json:"dataprocCluster,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.SparkProperties
type SparkProperties struct {

	// Optional. Dataproc Metastore Service configuration for the connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.SparkProperties.metastore_service_config
	MetastoreServiceConfig *MetastoreServiceConfig `json:"metastoreServiceConfig,omitempty"`

	// Optional. Spark History Server configuration for the connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.SparkProperties.spark_history_server_config
	SparkHistoryServerConfig *SparkHistoryServerConfig `json:"sparkHistoryServerConfig,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.AwsCrossAccountRole
type AwsCrossAccountRoleObservedState struct {
	// Output only. Google-owned AWS IAM User for a Connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AwsCrossAccountRole.iam_user_id
	IamUserID *string `json:"iamUserID,omitempty"`

	// Output only. A Google-generated id for representing Connection’s identity
	//  in AWS. External Id is also used for preventing the Confused Deputy
	//  Problem. See
	//  https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AwsCrossAccountRole.external_id
	ExternalID *string `json:"externalID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.AwsProperties
type AwsPropertiesObservedState struct {
	// Authentication using Google owned AWS IAM user's access key to assume
	//  into customer's AWS IAM Role.
	//  Deprecated, do not use.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AwsProperties.cross_account_role
	CrossAccountRole *AwsCrossAccountRoleObservedState `json:"crossAccountRole,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.AzureProperties
type AzurePropertiesObservedState struct {
	// Output only. The name of the Azure Active Directory Application.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AzureProperties.application
	Application *string `json:"application,omitempty"`

	// Output only. The client id of the Azure Active Directory Application.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AzureProperties.client_id
	ClientID *string `json:"clientID,omitempty"`

	// Output only. The object id of the Azure Active Directory Application.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AzureProperties.object_id
	ObjectID *string `json:"objectID,omitempty"`

	// Output only. A unique Google-owned and Google-generated identity for the
	//  Connection. This identity will be used to access the user's Azure Active
	//  Directory Application.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.AzureProperties.identity
	Identity *string `json:"identity,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.CloudResourceProperties
type CloudResourcePropertiesObservedState struct {
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
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudResourceProperties.service_account_id
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.CloudSqlProperties
type CloudSqlPropertiesObservedState struct {
	// Output only. The account ID of the service used for the purpose of this
	//  connection.
	//
	//  When the connection is used in the context of an operation in
	//  BigQuery, this service account will serve as the identity being used for
	//  connecting to the CloudSQL instance specified in this connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.CloudSqlProperties.service_account_id
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.Connection
type ConnectionObservedState struct {
	// Cloud SQL properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.cloud_sql
	CloudSql *CloudSqlPropertiesObservedState `json:"cloudSql,omitempty"`

	// Amazon Web Services (AWS) properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.aws
	Aws *AwsPropertiesObservedState `json:"aws,omitempty"`

	// Azure properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.azure
	Azure *AzurePropertiesObservedState `json:"azure,omitempty"`

	// Cloud Resource properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.cloud_resource
	CloudResource *CloudResourcePropertiesObservedState `json:"cloudResource,omitempty"`

	// Spark properties.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.spark
	Spark *SparkPropertiesObservedState `json:"spark,omitempty"`

	// Optional. Salesforce DataCloud properties. This field is intended for
	//  use only by Salesforce partner projects. This field contains properties
	//  for your Salesforce DataCloud connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.salesforce_data_cloud
	SalesforceDataCloud *SalesforceDataCloudPropertiesObservedState `json:"salesforceDataCloud,omitempty"`

	// Output only. The creation timestamp of the connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.creation_time
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Output only. The last update timestamp of the connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.last_modified_time
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`

	// Output only. True, if credential is configured for this connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.Connection.has_credential
	HasCredential *bool `json:"hasCredential,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.SalesforceDataCloudProperties
type SalesforceDataCloudPropertiesObservedState struct {
	// Output only. A unique Google-owned and Google-generated service account
	//  identity for the connection.
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.SalesforceDataCloudProperties.identity
	Identity *string `json:"identity,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.connection.v1.SparkProperties
type SparkPropertiesObservedState struct {
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
	// +kcc:proto:field=google.cloud.bigquery.connection.v1.SparkProperties.service_account_id
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}
