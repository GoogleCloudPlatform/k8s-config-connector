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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DatastreamConnectionProfileGVK = GroupVersion.WithKind("DatastreamConnectionProfile")

// DatastreamConnectionProfileSpec defines the desired state of DatastreamConnectionProfile
// +kcc:proto=google.cloud.datastream.v1.ConnectionProfile
type DatastreamConnectionProfileSpec struct {
	// The DatastreamConnectionProfile name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Labels.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Display name.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Oracle ConnectionProfile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.oracle_profile
	OracleProfile *OracleProfile `json:"oracleProfile,omitempty"`

	// Cloud Storage ConnectionProfile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.gcs_profile
	GCSProfile *GCSProfile `json:"gcsProfile,omitempty"`

	// MySQL ConnectionProfile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.mysql_profile
	MySQLProfile *MysqlProfile `json:"mySQLProfile,omitempty"`

	// BigQuery Connection Profile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.bigquery_profile
	BigQueryProfile *BigQueryProfile `json:"bigQueryProfile,omitempty"`

	// PostgreSQL Connection Profile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.postgresql_profile
	// NOTYET: this field is not implemented
	// PostgreSQLProfile *PostgresqlProfile `json:"postgreSQLProfile,omitempty"`

	// SQLServer Connection Profile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.sql_server_profile
	SQLServerProfile *SQLServerProfile `json:"sqlServerProfile,omitempty"`

	// Static Service IP connectivity.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.static_service_ip_connectivity
	StaticServiceIPConnectivity *StaticServiceIPConnectivity `json:"staticServiceIPConnectivity,omitempty"`

	// Forward SSH tunnel connectivity.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.forward_ssh_connectivity
	ForwardSSHConnectivity *ForwardSSHTunnelConnectivity `json:"forwardSSHConnectivity,omitempty"`

	// Private connectivity.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.private_connectivity
	PrivateConnectivity *PrivateConnectivity `json:"privateConnectivity,omitempty"`
}

// DatastreamConnectionProfileStatus defines the config connector machine state of DatastreamConnectionProfile
type DatastreamConnectionProfileStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DatastreamConnectionProfile resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DatastreamConnectionProfileObservedState `json:"observedState,omitempty"`
}

// DatastreamConnectionProfileObservedState is the state of the DatastreamConnectionProfile resource as most recently observed in GCP.
// +kcc:proto=google.cloud.datastream.v1.ConnectionProfile
type DatastreamConnectionProfileObservedState struct {
	// Output only. The resource's name.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. The create time of the resource.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time of the resource.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Oracle ConnectionProfile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.oracle_profile
	OracleProfile *OracleProfileObservedState `json:"oracleProfile,omitempty"`

	// MySQL ConnectionProfile configuration.
	// +kcc:proto:field=google.cloud.datastream.v1.ConnectionProfile.mysql_profile
	MySQLProfile *MysqlProfileObservedState `json:"mySQLProfile,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatastreamconnectionprofile;gcpdatastreamconnectionprofiles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DatastreamConnectionProfile is the Schema for the DatastreamConnectionProfile API
// +k8s:openapi-gen=true
type DatastreamConnectionProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DatastreamConnectionProfileSpec   `json:"spec,omitempty"`
	Status DatastreamConnectionProfileStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DatastreamConnectionProfileList contains a list of DatastreamConnectionProfile
type DatastreamConnectionProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatastreamConnectionProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatastreamConnectionProfile{}, &DatastreamConnectionProfileList{})
}

// +kcc:proto=google.cloud.datastream.v1.OracleAsmConfig
type OracleAsmConfig struct {
	// Required. Hostname for the Oracle ASM connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.hostname
	// +required
	Hostname *string `json:"hostname,omitempty"`

	// Required. Port for the Oracle ASM connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.port
	// +required
	Port *int32 `json:"port,omitempty"`

	// Required. Username for the Oracle ASM connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.username
	// +required
	// NOTYET: this field is replaced by the secretRef field
	// Username *string `json:"username,omitempty"`

	// Required. Password for the Oracle ASM connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.password
	// +required
	// NOTYET: this field is replaced by the secretRef field
	// Password *string `json:"password,omitempty"`

	// The Kubernetes Secret object that stores the "username" and "password" information for the Oracle ASM connection.
	// The Secret type has to be `kubernetes.io/basic-auth`.
	// +required
	SecretRef *refsv1beta1secret.BasicAuthSecretRef `json:"secretRef,omitempty"`

	// Required. ASM service name for the Oracle ASM connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.asm_service
	// +required
	ASMService *string `json:"asmService,omitempty"`

	// Optional. Connection string attributes
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.connection_attributes
	ConnectionAttributes map[string]string `json:"connectionAttributes,omitempty"`

	// Optional. SSL configuration for the Oracle connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleAsmConfig.oracle_ssl_config
	OracleSSLConfig *OracleSSLConfig `json:"oracleSSLConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.ForwardSshTunnelConnectivity
type ForwardSSHTunnelConnectivity struct {
	// Required. Hostname for the SSH tunnel.
	// +kcc:proto:field=google.cloud.datastream.v1.ForwardSshTunnelConnectivity.hostname
	// +required
	Hostname *string `json:"hostname,omitempty"`

	// Port for the SSH tunnel, default value is 22.
	// +kcc:proto:field=google.cloud.datastream.v1.ForwardSshTunnelConnectivity.port
	Port *int32 `json:"port,omitempty"`

	// Required. Username for the SSH tunnel.
	// +kcc:proto:field=google.cloud.datastream.v1.ForwardSshTunnelConnectivity.username
	// +required
	// NOTYET: this field is replaced by the secretRef field
	// Username *string `json:"username,omitempty"`

	// Input only. SSH password.
	// +kcc:proto:field=google.cloud.datastream.v1.ForwardSshTunnelConnectivity.password
	// NOTYET: this field is replaced by the secretRef field
	// Password *string `json:"password,omitempty"`

	// The Kubernetes Secret object that stores the "username" and "password" information for the SSH tunnel.
	// The Secret type has to be `kubernetes.io/basic-auth`.
	// +required
	SecretRef *refsv1beta1secret.BasicAuthSecretRef `json:"secretRef,omitempty"`

	// Input only. SSH private key.
	// +kcc:proto:field=google.cloud.datastream.v1.ForwardSshTunnelConnectivity.private_key
	PrivateKey *string `json:"privateKey,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.GcsProfile
type GCSProfile struct {
	// Required. The Cloud Storage bucket name.
	// +kcc:proto:field=google.cloud.datastream.v1.GcsProfile.bucket
	// +required
	Bucket *string `json:"bucket,omitempty"`

	// The root path inside the Cloud Storage bucket.
	// +kcc:proto:field=google.cloud.datastream.v1.GcsProfile.root_path
	RootPath *string `json:"rootPath,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlProfile
type MysqlProfile struct {
	// Required. Hostname for the MySQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlProfile.hostname
	// +required
	Hostname *string `json:"hostname,omitempty"`

	// Port for the MySQL connection, default value is 3306.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlProfile.port
	Port *int32 `json:"port,omitempty"`

	// Required. Username for the MySQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlProfile.username
	// +required
	// NOTYET: this field is replaced by the secretRef field
	// Username *string `json:"username,omitempty"`

	// Optional. Input only. Password for the MySQL connection. Mutually exclusive
	//  with the `secret_manager_stored_password` field.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlProfile.password
	// NOTYET: this field is replaced by the secretRef field
	// Password *string `json:"password,omitempty"`

	// The Kubernetes Secret object that stores the "username" and "password" information for the MySQL connection.
	// The Secret type has to be `kubernetes.io/basic-auth`.
	// +required
	SecretRef *refsv1beta1secret.BasicAuthSecretRef `json:"secretRef,omitempty"`

	// SSL configuration for the MySQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlProfile.ssl_config
	SSLConfig *MysqlSSLConfig `json:"sslConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleProfile
type OracleProfile struct {
	// Required. Hostname for the Oracle connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.hostname
	// +required
	Hostname *string `json:"hostname,omitempty"`

	// Port for the Oracle connection, default value is 1521.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.port
	Port *int32 `json:"port,omitempty"`

	// Required. Username for the Oracle connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.username
	// +required
	// NOTYET: this field is replaced by the secretRef field
	// Username *string `json:"username,omitempty"`

	// Optional. Password for the Oracle connection. Mutually exclusive with the
	//  `secret_manager_stored_password` field.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.password
	// NOTYET: this field is replaced by the secretRef field
	// Password *string `json:"password,omitempty"`

	// The Kubernetes Secret object that stores the "username" and "password" information for the Oracle connection.
	// The Secret type has to be `kubernetes.io/basic-auth`.
	// +required
	SecretRef *refsv1beta1secret.BasicAuthSecretRef `json:"secretRef,omitempty"`

	// Required. Database for the Oracle connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.database_service
	// +required
	DatabaseService *string `json:"databaseService,omitempty"`

	// Connection string attributes
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.connection_attributes
	ConnectionAttributes map[string]string `json:"connectionAttributes,omitempty"`

	// Optional. SSL configuration for the Oracle connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.oracle_ssl_config
	OracleSSLConfig *OracleSSLConfig `json:"oracleSSLConfig,omitempty"`

	// Optional. Configuration for Oracle ASM connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.oracle_asm_config
	OracleASMConfig *OracleAsmConfig `json:"oracleASMConfig,omitempty"`

	// Optional. A reference to a Secret Manager resource name storing the Oracle
	//  connection password. Mutually exclusive with the `secretRef` field.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.secret_manager_stored_password
	SecreteManagerSecretRef *refsv1beta1.SecretManagerSecretRef `json:"secretManagerSecretRef,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PostgresqlProfile
type PostgresqlProfile struct {
	// Required. Hostname for the PostgreSQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlProfile.hostname
	// +required
	Hostname *string `json:"hostname,omitempty"`

	// Port for the PostgreSQL connection, default value is 5432.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlProfile.port
	Port *int32 `json:"port,omitempty"`

	// Required. Username for the PostgreSQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlProfile.username
	// +required
	// NOTYET: this field is replaced by the secretRef field
	// Username *string `json:"username,omitempty"`

	// Optional. Password for the PostgreSQL connection. Mutually exclusive with
	//  the `secret_manager_stored_password` field.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlProfile.password
	// NOTYET: this field is replaced by the secretRef field
	// Password *string `json:"password,omitempty"`

	// The Kubernetes Secret object that stores the "username" and "password" information for the PostgreSQL connection.
	// The Secret type has to be `kubernetes.io/basic-auth`.
	// +required
	SecretRef *refsv1beta1secret.BasicAuthSecretRef `json:"secretRef,omitempty"`

	// Required. Database for the PostgreSQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.PostgresqlProfile.database
	// +required
	Database *string `json:"database,omitempty"`

	// TODO: ssl_config proto field is not generated
}

// +kcc:proto=google.cloud.datastream.v1.SqlServerProfile
type SQLServerProfile struct {
	// Required. Hostname for the SQLServer connection.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerProfile.hostname
	// +required
	Hostname *string `json:"hostname,omitempty"`

	// Port for the SQLServer connection, default value is 1433.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerProfile.port
	Port *int32 `json:"port,omitempty"`

	// Required. Username for the SQLServer connection.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerProfile.username
	// +required
	// NOTYET: this field is replaced by the secretRef field
	// Username *string `json:"username,omitempty"`

	// Optional. Password for the SQLServer connection. Mutually exclusive with
	//  the `secret_manager_stored_password` field.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerProfile.password
	// NOTYET: this field is replaced by the secretRef field
	// Password *string `json:"password,omitempty"`

	// The Kubernetes Secret object that stores the "username" and "password" information for the SQLServer connection.
	// The Secret type has to be `kubernetes.io/basic-auth`.
	// +required
	SecretRef *refsv1beta1secret.BasicAuthSecretRef `json:"secretRef,omitempty"`

	// Required. Database for the SQLServer connection.
	// +kcc:proto:field=google.cloud.datastream.v1.SqlServerProfile.database
	// +required
	Database *string `json:"database,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.PrivateConnectivity
type PrivateConnectivity struct {
	// Required. A reference to a private connection resource.
	// +kcc:proto:field=google.cloud.datastream.v1.PrivateConnectivity.private_connection
	// +required
	PrivateConnectionRef *PrivateConnectionRef `json:"privateConnectionRef,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.MysqlProfile
type MysqlProfileObservedState struct {
	// SSL configuration for the MySQL connection.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlProfile.ssl_config
	SSLConfig *MysqlSSLConfigObservedState `json:"sslConfig,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleProfile
type OracleProfileObservedState struct {
	// Optional. SSL configuration for the Oracle connection.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleProfile.oracle_ssl_config
	OracleSSLConfig *OracleSSLConfigObservedState `json:"oracleSSLConfig,omitempty"`
}
