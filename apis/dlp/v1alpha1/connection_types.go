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

var DLPConnectionGVK = GroupVersion.WithKind("DLPConnection")

// DLPConnectionSpec defines the desired state of DLPConnection
// +kcc:spec:proto=google.privacy.dlp.v2.Connection
type DLPConnectionSpec struct {
	// The Project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The DLPConnection name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The connection's state in its lifecycle.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=CONNECTION_STATE_UNSPECIFIED;MISSING_CREDENTIALS;AVAILABLE;ERROR
	// +kcc:proto:field=google.privacy.dlp.v2.Connection.state
	State *string `json:"state,omitempty"`

	// Connect to a Cloud SQL instance.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.privacy.dlp.v2.Connection.cloud_sql
	CloudSQL *CloudSQLProperties `json:"cloudSQL,omitempty"`
}

// +kcc:proto=google.privacy.dlp.v2.CloudSqlProperties
type CloudSQLProperties struct {
	// Optional. Immutable. The Cloud SQL instance for which the connection is
	//  defined. Only one connection per instance is allowed. This can only be set
	//  at creation time, and cannot be updated.
	//
	//  It is an error to use a connection_name from different project or region
	//  than the one that holds the connection.
	//  For example, a Connection resource for Cloud SQL connection_name
	//  `project-id:us-central1:sql-instance`
	//  must be created under the parent
	//  `projects/project-id/locations/us-central1`
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.privacy.dlp.v2.CloudSqlProperties.connection_name
	ConnectionName *string `json:"connectionName,omitempty"`

	// A username and password stored in Secret Manager.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.privacy.dlp.v2.CloudSqlProperties.username_password
	UsernamePassword *SecretManagerCredential `json:"usernamePassword,omitempty"`

	// Built-in IAM authentication (must be configured in Cloud SQL).
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.privacy.dlp.v2.CloudSqlProperties.cloud_sql_iam
	CloudSQLIAM *CloudSQLIAMCredential `json:"cloudSQLIAM,omitempty"`

	// Required. The DLP API will limit its connections to max_connections.
	//  Must be 2 or greater.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.privacy.dlp.v2.CloudSqlProperties.max_connections
	MaxConnections *int32 `json:"maxConnections,omitempty"`

	// Required. The database engine used by the Cloud SQL instance that this
	//  connection configures.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.privacy.dlp.v2.CloudSqlProperties.database_engine
	DatabaseEngine *string `json:"databaseEngine,omitempty"`
}

// +kcc:proto=google.privacy.dlp.v2.SecretManagerCredential
type SecretManagerCredential struct {
	// Required. The username.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.privacy.dlp.v2.SecretManagerCredential.username
	Username *string `json:"username,omitempty"`

	// Required. The name of the Secret Manager resource that stores the password,
	//  in the form `projects/project-id/secrets/secret-name/versions/version`.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.privacy.dlp.v2.SecretManagerCredential.password_secret_version_name
	PasswordSecretVersionName *string `json:"passwordSecretVersionName,omitempty"`
}

// +kcc:proto=google.privacy.dlp.v2.CloudSqlIamCredential
type CloudSQLIAMCredential struct {
}

// DLPConnectionStatus defines the config connector machine state of DLPConnection
type DLPConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DLPConnection resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *DLPConnectionObservedState `json:"observedState,omitempty"`
}

// DLPConnectionObservedState is the state of the DLPConnection resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.privacy.dlp.v2.Connection
type DLPConnectionObservedState struct {
	// Output only. Name of the connection:
	//  `projects/{project}/locations/{location}/connections/{name}`.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.privacy.dlp.v2.Connection.name
	Name *string `json:"name,omitempty"`

	// Output only. Set if status == ERROR, to provide additional details. Will
	//  store the last 10 errors sorted with the most recent first.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.privacy.dlp.v2.Connection.errors
	Errors []Error `json:"errors,omitempty"`
}

// +kcc:proto=google.privacy.dlp.v2.Error
type Error struct {
	// Detailed error codes and messages.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.privacy.dlp.v2.Error.details
	Details *common.Status `json:"details,omitempty"`

	// The times the error occurred. List includes the oldest timestamp and the
	//  last 9 timestamps.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.privacy.dlp.v2.Error.timestamps
	Timestamps []string `json:"timestamps,omitempty"`

	// Additional information about the error.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.privacy.dlp.v2.Error.extra_info
	ExtraInfo *string `json:"extraInfo,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdlpconnection;gcpdlpconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DLPConnection is the Schema for the DLPConnection API
// +k8s:openapi-gen=true
type DLPConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DLPConnectionSpec   `json:"spec,omitempty"`
	Status DLPConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DLPConnectionList contains a list of DLPConnection
type DLPConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DLPConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DLPConnection{}, &DLPConnectionList{})
}
