// Copyright 2026 Google LLC
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

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDMSConnectionProfileGVK = GroupVersion.WithKind("CloudDMSConnectionProfile")

// CloudDMSConnectionProfileSpec defines the desired state of CloudDMSConnectionProfile
// +kcc:spec:proto=google.cloud.clouddms.v1.ConnectionProfile
type CloudDMSConnectionProfileSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The CloudDMSConnectionProfile name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The connection profile display name.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The resource labels for connection profile to use to annotate any related
	// underlying resources such as Compute Engine VMs. An object containing a
	// list of "key": "value" pairs.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.labels
	Labels map[string]string `json:"labels,omitempty"`

	// A MySQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.mysql
	Mysql *MySQLConnectionProfile `json:"mysql,omitempty"`

	// A PostgreSQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.postgresql
	Postgresql *PostgreSQLConnectionProfile `json:"postgresql,omitempty"`

	// An Oracle database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.oracle
	Oracle *OracleConnectionProfile `json:"oracle,omitempty"`

	// A CloudSQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.cloudsql
	Cloudsql *CloudSQLConnectionProfile `json:"cloudsql,omitempty"`

	// An AlloyDB cluster connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.alloydb
	Alloydb *AlloyDbConnectionProfile `json:"alloydb,omitempty"`

	// The database provider.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.provider
	Provider *string `json:"provider,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.AlloyDbConnectionProfile
type AlloyDbConnectionProfile struct {
	// Required. The AlloyDB cluster ID that this connection profile is associated with.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbConnectionProfile.cluster_id
	ClusterRef *refsv1beta1.AlloyDBClusterTypeRef `json:"clusterRef,omitempty"`

	// Immutable. Metadata used to create the destination AlloyDB cluster.
	// +kcc:proto:field=google.cloud.clouddms.v1.AlloyDbConnectionProfile.settings
	Settings *AlloyDbSettings `json:"settings,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.OracleConnectionProfile
type OracleConnectionProfile struct {
	// Required. The IP or hostname of the source Oracle database.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.host
	Host *string `json:"host,omitempty"`

	// Required. The network port of the source Oracle database.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.port
	Port *int32 `json:"port,omitempty"`

	// Required. The username that Database Migration Service will use to connect
	// to the database. The value is encrypted when stored in Database Migration
	// Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.username
	Username *string `json:"username,omitempty"`

	// Required. Input only. The password for the user that Database Migration
	// Service will be using to connect to the database. This field is not
	// returned on request, and the value is encrypted when stored in Database
	// Migration Service.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.password
	Password *string `json:"password,omitempty"`

	// Required. Database service for the Oracle connection.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.database_service
	DatabaseService *string `json:"databaseService,omitempty"`

	// SSL configuration for the destination to connect to the source database.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.ssl
	SSL *SSLConfig `json:"ssl,omitempty"`

	// Static ip connectivity data (default, no additional details needed).
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.static_service_ip_connectivity
	StaticServiceIPConnectivity *StaticServiceIPConnectivity `json:"staticServiceIPConnectivity,omitempty"`

	// Forward SSH tunnel connectivity.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.forward_ssh_connectivity
	ForwardSSHConnectivity *ForwardSSHTunnelConnectivity `json:"forwardSSHConnectivity,omitempty"`

	// Private connectivity.
	// +kcc:proto:field=google.cloud.clouddms.v1.OracleConnectionProfile.private_connectivity
	PrivateConnectivity *PrivateConnectivity `json:"privateConnectivity,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.PrivateConnectivity
type PrivateConnectivity struct {
	// Required. The resource name (URI) of the private connection.
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateConnectivity.private_connection
	PrivateConnectionRef *PrivateConnectionRef `json:"privateConnectionRef,omitempty"`
}

// CloudDMSConnectionProfileStatus defines the config connector machine state of CloudDMSConnectionProfile
type CloudDMSConnectionProfileStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDMSConnectionProfile resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDMSConnectionProfileObservedState `json:"observedState,omitempty"`
}

// CloudDMSConnectionProfileObservedState is the state of the CloudDMSConnectionProfile resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.clouddms.v1.ConnectionProfile
type CloudDMSConnectionProfileObservedState struct {
	// Output only. The create time of the resource.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update time of the resource.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.state
	State *string `json:"state,omitempty"`

	// Output only. The error details in case of state FAILED.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.error
	Error *Status `json:"error,omitempty"`

	// A MySQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.mysql
	Mysql *MySQLConnectionProfileObservedState `json:"mysql,omitempty"`

	// A PostgreSQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.postgresql
	Postgresql *PostgreSQLConnectionProfileObservedState `json:"postgresql,omitempty"`

	// A CloudSQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.cloudsql
	Cloudsql *CloudSQLConnectionProfileObservedState `json:"cloudsql,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddmsconnectionprofile;gcpclouddmsconnectionprofiles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDMSConnectionProfile is the Schema for the CloudDMSConnectionProfile API
// +k8s:openapi-gen=true
type CloudDMSConnectionProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudDMSConnectionProfileSpec   `json:"spec,omitempty"`
	Status CloudDMSConnectionProfileStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDMSConnectionProfileList contains a list of CloudDMSConnectionProfile
type CloudDMSConnectionProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDMSConnectionProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDMSConnectionProfile{}, &CloudDMSConnectionProfileList{})
}

// +kcc:proto=google.cloud.clouddms.v1.StaticIpConnectivity
// +kubebuilder:validation:XPreserveUnknownFields
type StaticIPConnectivity struct {
}

// +kcc:proto=google.cloud.clouddms.v1.StaticServiceIpConnectivity
// +kubebuilder:validation:XPreserveUnknownFields
type StaticServiceIPConnectivity struct {
}
