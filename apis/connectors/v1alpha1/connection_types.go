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
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ConnectorsConnectionGVK = GroupVersion.WithKind("ConnectorsConnection")

// ConnectorsConnectionSpec defines the desired state of ConnectorsConnection
// +kcc:spec:proto=google.cloud.connectors.v1.Connection
type ConnectorsConnectionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ConnectorsConnection name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Resource labels to represent user-provided metadata.
	//  Refer to cloud documentation on labels for more details.
	//  https://cloud.google.com/compute/docs/labeling-resources
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the resource.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.description
	Description *string `json:"description,omitempty"`

	// Required. Connector version on which the connection is created.
	//  The format is:
	//  projects/-*-/locations/-*-/providers/-*-/connectors/-*-/versions/-*
	//  Only global location is supported for ConnectorVersion resource.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.connector_version
	// +kubebuilder:validation:Required
	ConnectorVersionRef *ConnectorsConnectorVersionRef `json:"connectorVersionRef,omitempty"`

	// Optional. Configuration for configuring the connection with an external
	//  system.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.config_variables
	ConfigVariables []ConfigVariable `json:"configVariables,omitempty"`

	// Optional. Configuration for establishing the connection's authentication
	//  with an external system.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.auth_config
	AuthConfig *AuthConfig `json:"authConfig,omitempty"`

	// Optional. Configuration that indicates whether or not the Connection can be
	//  edited.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.lock_config
	LockConfig *LockConfig `json:"lockConfig,omitempty"`

	// Optional. Configuration of the Connector's destination. Only accepted for
	//  Connectors that accepts user defined destination(s).
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.destination_configs
	DestinationConfigs []DestinationConfig `json:"destinationConfigs,omitempty"`

	// Optional. Service account needed for runtime plane to access GCP resources.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. Suspended indicates if a user has suspended a connection or not.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.suspended
	Suspended *bool `json:"suspended,omitempty"`

	// Optional. Node configuration for the connection.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.node_config
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	// Optional. Ssl config of a connection
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.ssl_config
	SSLConfig *SSLConfig `json:"sslConfig,omitempty"`
}

// ConnectorsConnectionStatus defines the config connector machine state of ConnectorsConnection
type ConnectorsConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ConnectorsConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ConnectorsConnectionObservedState `json:"observedState,omitempty"`
}

// ConnectorsConnectionObservedState is the state of the ConnectorsConnection resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.connectors.v1.Connection
type ConnectorsConnectionObservedState struct {

	// Output only. Resource name of the Connection.
	//  Format: projects/{project}/locations/{location}/connections/{connection}
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.name
	Name *string `json:"name,omitempty"`

	// Output only. Created time.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Updated time.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current status of the connection.
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.status
	Status *ConnectionStatus `json:"status,omitempty"`

	// Output only. GCR location where the runtime image is stored.
	//  formatted like: gcr.io/{bucketName}/{imageName}
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.image_location
	ImageLocation *string `json:"imageLocation,omitempty"`

	// Output only. The name of the Service Directory service name. Used for
	//  Private Harpoon to resolve the ILB address.
	//  e.g.
	//  "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.service_directory
	ServiceDirectory *string `json:"serviceDirectory,omitempty"`

	// Output only. GCR location where the envoy image is stored.
	//  formatted like: gcr.io/{bucketName}/{imageName}
	// +kcc:proto:field=google.cloud.connectors.v1.Connection.envoy_image_location
	EnvoyImageLocation *string `json:"envoyImageLocation,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpconnectorsconnection;gcpconnectorsconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ConnectorsConnection is the Schema for the ConnectorsConnection API
// +k8s:openapi-gen=true
type ConnectorsConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ConnectorsConnectionSpec   `json:"spec,omitempty"`
	Status ConnectorsConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ConnectorsConnectionList contains a list of ConnectorsConnection
type ConnectorsConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectorsConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConnectorsConnection{}, &ConnectorsConnectionList{})
}

// +kcc:proto=google.cloud.connectors.v1.Destination
type Destination struct {
	// PSC service attachments.
	//  Format: projects/*/regions/*/serviceAttachments/*
	// +kcc:proto:field=google.cloud.connectors.v1.Destination.service_attachment
	ServiceAttachmentRef *refsv1beta1.ComputeServiceAttachmentRef `json:"serviceAttachmentRef,omitempty"`

	// For publicly routable host.
	// +kcc:proto:field=google.cloud.connectors.v1.Destination.host
	Host *string `json:"host,omitempty"`

	// The port is the target port number that is accepted by the destination.
	// +kcc:proto:field=google.cloud.connectors.v1.Destination.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.AuthConfig.UserPassword
type AuthConfig_UserPassword struct {
	// Username.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.UserPassword.username
	Username *string `json:"username,omitempty"`

	// Secret version reference containing the password.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfig.UserPassword.password
	SecretRef *secretmanagerv1beta1.SecretRef `json:"secretRef,omitempty"`
}
