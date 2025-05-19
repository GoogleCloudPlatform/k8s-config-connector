// Copyright 2024 Google LLC
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

package v1beta1

import (
	dataprocv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1beta1"
	metastorev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1alpha1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	spannerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryConnectionConnectionGVK = GroupVersion.WithKind("BigQueryConnectionConnection")

type Parent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// BigQueryConnectionConnectionSpec defines the desired state to connect BigQuery to external resources
// +kcc:spec:proto=google.cloud.bigquery.connection.v1.Connection
type BigQueryConnectionConnectionSpec struct {
	Parent `json:",inline"`

	// Immutable. Optional.
	// The BigQuery Connection ID used for resource creation or acquisition.
	// For creation: If specified, this value is used as the connection ID. If not provided, a UUID will be generated and assigned as the connection ID.
	// For acquisition: This field must be provided to identify the connection resource to acquire.
	ResourceID *string `json:"resourceID,omitempty"`

	// User provided display name for the connection.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// User provided description.
	Description *string `json:"description,omitempty"`

	// Cloud SQL properties.
	CloudSQLSpec *CloudSqlPropertiesSpec `json:"cloudSQL,omitempty"`

	// Amazon Web Services (AWS) properties.
	AwsSpec *AwsPropertiesSpec `json:"aws,omitempty"`

	// Azure properties.
	AzureSpec *AzurePropertiesSpec `json:"azure,omitempty"`

	/* NOTYET
	// Optional. Salesforce DataCloud properties. This field is intended for
	//  use only by Salesforce partner projects. This field contains properties
	//  for your Salesforce DataCloud connection.
	SalesforceDataCloud *SalesforceDataCloudProperties `json:"salesforceDataCloud,omitempty"`
	*/

	// Use Cloud Resource properties.
	CloudResourceSpec *CloudResourcePropertiesSpec `json:"cloudResource,omitempty"`

	// Cloud Spanner properties.
	CloudSpannerSpec *CloudSpannerPropertiesSpec `json:"cloudSpanner,omitempty"`

	// Spark properties.
	SparkSpec *SparkPropertiesSpec `json:"spark,omitempty"`
}

// BigQueryConnectionConnectionStatus defines the config connector machine state of BigQueryConnectionConnection
type BigQueryConnectionConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryConnectionConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryConnectionConnectionObservedState `json:"observedState,omitempty"`
}

// BigQueryConnectionConnectionSpec defines the desired state of BigQueryConnectionConnection
// +kcc:observedstate:proto=google.cloud.bigquery.connection.v1.Connection
type BigQueryConnectionConnectionObservedState struct {
	Aws *AwsPropertiesStatus `json:"aws,omitempty"`

	Azure *AzurePropertiesStatus `json:"azure,omitempty"`

	CloudResource *CloudResourcePropertiesStatus `json:"cloudResource,omitempty"`

	CloudSQL *CloudSqlPropertiesStatus `json:"cloudSQL,omitempty"`

	Spark *SparkPropertiesStatus `json:"spark,omitempty"`

	// The display name for the connection.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// The description for the connection.
	Description *string `json:"description,omitempty"`

	// Output only. True, if credential is configured for this connection.
	HasCredential *bool `json:"hasCredential,omitempty"`
}

type AwsPropertiesSpec struct {
	// Authentication using Google owned service account to assume into
	//  customer's AWS IAM Role.
	// +required
	AccessRole *AwsAccessRoleSpec `json:"accessRole,omitempty"`
}

type AwsAccessRoleSpec struct {
	// The user’s AWS IAM Role that trusts the Google-owned AWS IAM user
	//  Connection.
	// +required
	IamRoleID *string `json:"iamRoleID,omitempty"`
}

type AzurePropertiesSpec struct {
	// The id of customer's directory that host the data.
	// +required
	CustomerTenantID *string `json:"customerTenantID,omitempty"`

	// The client ID of the user's Azure Active Directory Application used for a
	//  federated connection.
	FederatedApplicationClientID *string `json:"federatedApplicationClientID,omitempty"`
}

type CloudResourcePropertiesSpec struct{}

type CloudSqlPropertiesSpec struct {
	// Reference to the Cloud SQL instance ID.
	// +required
	InstanceRef *refv1beta1.SQLInstanceRef `json:"instanceRef,omitempty"`

	// Reference to the SQL Database.
	// +required
	DatabaseRef *refv1beta1.SQLDatabaseRef `json:"databaseRef,omitempty"`

	// Type of the Cloud SQL database.
	// +required
	Type *string `json:"type,omitempty"`

	// Cloud SQL credential.
	// +required
	Credential *CloudSqlCredential `json:"credential,omitempty"`
}

type CloudSpannerPropertiesSpec struct {
	// Reference to a spanner database ID.
	// +required
	DatabaseRef *spannerv1beta1.SpannerDatabaseRef `json:"databaseRef,omitempty"`

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

type SparkPropertiesSpec struct {
	// Optional. Dataproc Metastore Service configuration for the connection.
	MetastoreService *MetastoreServiceConfigSpec `json:"metastoreService,omitempty"`

	// Optional. Spark History Server configuration for the connection.
	SparkHistoryServer *SparkHistoryServerConfigSpec `json:"sparkHistoryServer,omitempty"`
}

type MetastoreServiceConfigSpec struct {
	// Optional. Resource name of an existing Dataproc Metastore service.
	//
	//  Example:
	//
	//  * `projects/[project_id]/locations/[region]/services/[service_id]`
	MetastoreServiceRef *metastorev1alpha1.ServiceRef `json:"metastoreServiceRef,omitempty"`
}

type SparkHistoryServerConfigSpec struct {
	// Optional. Resource name of an existing Dataproc Cluster to act as a Spark
	//  History Server for the connection.
	//
	//  Example:
	//
	//  * `projects/[project_id]/regions/[region]/clusters/[cluster_name]`
	DataprocClusterRef *dataprocv1beta1.DataprocClusterRef `json:"dataprocClusterRef,omitempty"`
}

// +kcc:status:proto=google.cloud.bigquery.connection.v1.AwsProperties
type AwsPropertiesStatus struct {
	AccessRole *AwsAccessRoleStatus `json:"accessRole,omitempty"`
}

// +kcc:status:proto=google.cloud.bigquery.connection.v1.AwsAccessRole
type AwsAccessRoleStatus struct {
	// A unique Google-owned and Google-generated identity for the Connection.
	//  This identity will be used to access the user's AWS IAM Role.
	Identity *string `json:"identity,omitempty"`
}

// +kcc:status:proto=google.cloud.bigquery.connection.v1.AzureProperties
type AzurePropertiesStatus struct {
	// The name of the Azure Active Directory Application.
	Application *string `json:"application,omitempty"`

	// The client id of the Azure Active Directory Application.
	ClientID *string `json:"clientID,omitempty"`

	// The object id of the Azure Active Directory Application.
	ObjectID *string `json:"objectID,omitempty"`

	// The URL user will be redirected to after granting consent during connection
	//  setup.
	RedirectUri *string `json:"redirectUri,omitempty"`

	// A unique Google-owned and Google-generated identity for the
	//  Connection. This identity will be used to access the user's Azure Active
	//  Directory Application.
	Identity *string `json:"identity,omitempty"`
}

// +kcc:status:proto=google.cloud.bigquery.connection.v1.CloudSqlProperties
type CloudSqlPropertiesStatus struct {
	// The account ID of the service used for the purpose of this connection.
	//
	//  When the connection is used in the context of an operation in
	//  BigQuery, this service account will serve as the identity being used for
	//  connecting to the CloudSQL instance specified in this connection.
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

// +kcc:status:proto=google.cloud.bigquery.connection.v1.CloudResourceProperties
type CloudResourcePropertiesStatus struct {
	//  The account ID of the service created for the purpose of this
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

// +kcc:status:proto=google.cloud.bigquery.connection.v1.SparkProperties
type SparkPropertiesStatus struct {
	//  The account ID of the service created for the purpose of this
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
}

// +kcc:proto=google.cloud.bigquery.connection.v1.CloudSqlCredential
type CloudSqlCredential struct {
	// The Kubernetes Secret object that stores the "username" and "password" information.
	// The Secret type has to be `kubernetes.io/basic-auth`.
	SecretRef *refsv1beta1secret.BasicAuthSecretRef `json:"secretRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryconnectionconnection;gcpbigqueryconnectionconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:storageversion

// BigQueryConnectionConnection is the Schema for the BigQueryConnectionConnection API
// +k8s:openapi-gen=true
type BigQueryConnectionConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryConnectionConnectionSpec   `json:"spec,omitempty"`
	Status BigQueryConnectionConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryConnectionConnectionList contains a list of BigQueryConnectionConnection
type BigQueryConnectionConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryConnectionConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryConnectionConnection{}, &BigQueryConnectionConnectionList{})
}
