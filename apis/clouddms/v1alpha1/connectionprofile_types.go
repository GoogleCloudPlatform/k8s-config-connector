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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDMSConnectionProfileGVK = GroupVersion.WithKind("CloudDMSConnectionProfile")

// CloudDMSConnectionProfileSpec defines the desired state of CloudDMSConnectionProfile
// +kcc:spec:proto=google.cloud.clouddms.v1.ConnectionProfile
type CloudDMSConnectionProfileSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The CloudDMSConnectionProfile name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The connection profile display name.
	// +kcc:proto=display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A MySQL database connection profile.
	// +kcc:proto=mysql
	Mysql *MySQLConnectionProfile `json:"mysql,omitempty"`

	// A PostgreSQL database connection profile.
	// +kcc:proto=postgresql
	Postgresql *PostgreSQLConnectionProfile `json:"postgresql,omitempty"`

	// An Oracle database connection profile.
	// +kcc:proto=oracle
	Oracle *OracleConnectionProfile `json:"oracle,omitempty"`

	// A CloudSQL database connection profile.
	// +kcc:proto=cloudsql
	Cloudsql *CloudSQLConnectionProfile `json:"cloudsql,omitempty"`

	// An AlloyDB cluster connection profile.
	// +kcc:proto=alloydb
	Alloydb *AlloyDbConnectionProfile `json:"alloydb,omitempty"`

	// The database provider.
	// +kcc:proto=provider
	Provider *string `json:"provider,omitempty"`
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
	// The name of this connection profile resource in the form of
	//  projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
	// +kcc:proto=name
	Name *string `json:"name,omitempty"`

	// The current connection profile state (e.g. DRAFT, READY, or FAILED).
	// +kcc:proto=state
	State *string `json:"state,omitempty"`

	// Output only. The timestamp when the resource was created.
	//  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	//  Example: "2014-10-02T15:01:23.045123456Z".
	// +kcc:proto=create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was last updated.
	//  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	//  Example: "2014-10-02T15:01:23.045123456Z".
	// +kcc:proto=update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The error details in case of state FAILED.
	// +kcc:proto=error
	Error *Status `json:"error,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddmsconnectionprofile;gcpclouddmsconnectionprofiles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
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
