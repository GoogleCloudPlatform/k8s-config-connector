// Copyright 2026 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DatabaseMigrationConnectionProfileSpec defines the desired state of DatabaseMigrationConnectionProfile
// +kcc:spec:proto=google.cloud.clouddms.v1.ConnectionProfile
type DatabaseMigrationConnectionProfileSpec struct {
	// The DatabaseMigrationConnectionProfile name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The resource labels for connection profile to use to annotate any related underlying resources.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The connection profile display name.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.display_name
	DisplayName *string `json:"displayName,omitempty"`

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

// DatabaseMigrationConnectionProfileStatus defines the config connector machine state of DatabaseMigrationConnectionProfile
type DatabaseMigrationConnectionProfileStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DatabaseMigrationConnectionProfile resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DatabaseMigrationConnectionProfileObservedState `json:"observedState,omitempty"`
}

// DatabaseMigrationConnectionProfileObservedState is the state of the DatabaseMigrationConnectionProfile resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.clouddms.v1.ConnectionProfile
type DatabaseMigrationConnectionProfileObservedState struct {
	// Output only. The timestamp when the resource was created.
	//  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	//  Example: "2014-10-02T15:01:23.045123456Z".
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was last updated.
	//  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	//  Example: "2014-10-02T15:01:23.045123456Z".
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// A MySQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.mysql
	Mysql *MySQLConnectionProfileObservedState `json:"mysql,omitempty"`

	// A PostgreSQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.postgresql
	Postgresql *PostgreSQLConnectionProfileObservedState `json:"postgresql,omitempty"`

	// An Oracle database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.oracle
	Oracle *OracleConnectionProfileObservedState `json:"oracle,omitempty"`

	// A CloudSQL database connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.cloudsql
	Cloudsql *CloudSQLConnectionProfileObservedState `json:"cloudsql,omitempty"`

	// An AlloyDB cluster connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.alloydb
	Alloydb *AlloyDbConnectionProfileObservedState `json:"alloydb,omitempty"`

	// Output only. The error details in case of state FAILED.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. The current connection profile state (e.g. DRAFT, READY, or FAILED).
	// +kcc:proto:field=google.cloud.clouddms.v1.ConnectionProfile.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatabasemigrationconnectionprofile;gcpdatabasemigrationconnectionprofiles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DatabaseMigrationConnectionProfile is the Schema for the DatabaseMigrationConnectionProfile API
// +k8s:openapi-gen=true
type DatabaseMigrationConnectionProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DatabaseMigrationConnectionProfileSpec   `json:"spec,omitempty"`
	Status DatabaseMigrationConnectionProfileStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DatabaseMigrationConnectionProfileList contains a list of DatabaseMigrationConnectionProfile
type DatabaseMigrationConnectionProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseMigrationConnectionProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseMigrationConnectionProfile{}, &DatabaseMigrationConnectionProfileList{})
}

// +kcc:proto=google.cloud.clouddms.v1.StaticServiceIpConnectivity
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:Schemaless
type StaticServiceIPConnectivity struct {
}
