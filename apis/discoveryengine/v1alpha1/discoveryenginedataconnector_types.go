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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineDataConnectorGVK = GroupVersion.WithKind("DiscoveryEngineDataConnector")

// +kcc:proto=google.cloud.discoveryengine.v1alpha.ActionConfig
type ActionConfig struct {
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.ActionConfig.service_name
	ServiceName *string `json:"serviceName,omitempty"`

	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.ActionConfig.action_params
	ActionParams apiextensionsv1.JSON `json:"actionParams,omitempty"`

	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.ActionConfig.create_bap_connection
	CreateBapConnection *bool `json:"createBapConnection,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.ConnectorMetadata
type ConnectorMetadata struct {
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.ConnectorMetadata.title
	Title *string `json:"title,omitempty"`

	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.ConnectorMetadata.short_description
	ShortDescription *string `json:"shortDescription,omitempty"`

	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.ConnectorMetadata.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.ConnectorMetadata.author
	Author *string `json:"author,omitempty"`

	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.ConnectorMetadata.note
	Note *string `json:"note,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.ConnectorRun
type ConnectorRun struct {
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.ConnectorRun.name
	Name *string `json:"name,omitempty"`

	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.ConnectorRun.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DataProtectionPolicy
type DataProtectionPolicy struct {
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataProtectionPolicy.policy_type
	PolicyType *string `json:"policyType,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DestinationConfig
type DestinationConfig struct {
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DestinationConfig.destination_type
	DestinationType *string `json:"destinationType,omitempty"`

	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DestinationConfig.params
	Params apiextensionsv1.JSON `json:"params,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.RealtimeSyncConfig
type RealtimeSyncConfig struct {
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.RealtimeSyncConfig.streaming_error
	StreamingError *string `json:"streamingError,omitempty"`
}

// DiscoveryEngineDataConnectorSpec defines the desired state of DiscoveryEngineDataConnector
// +kcc:spec:proto=google.cloud.discoveryengine.v1alpha.DataConnector
type DiscoveryEngineDataConnectorSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// +required
	Location string `json:"location"`

	// Immutable. The collection for the DataConnector.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Collection field is immutable"
	// +required
	Collection string `json:"collection"`

	// The DiscoveryEngineDataConnector name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The data source type.
	// +required
	DataSource string `json:"dataSource"`

	// Optional. Connector-specific parameters in structured JSON format.
	Params apiextensionsv1.JSON `json:"params,omitempty"`

	// Optional. The interval at which the connector should run a full sync.
	RefreshInterval *string `json:"refreshInterval,omitempty"`

	// Optional. If true, the connector will not run automatically.
	AutoRunDisabled *bool `json:"autoRunDisabled,omitempty"`

	// Optional. Configuration to support actions on the connector.
	ActionConfig *ActionConfig `json:"actionConfig,omitempty"`

	// Optional. List of destination configurations for the ingested data.
	DestinationConfigs []DestinationConfig `json:"destinationConfigs,omitempty"`

	// Optional. The sync mode.
	SyncMode *string `json:"syncMode,omitempty"`

	// Optional. If true, incremental sync is disabled.
	IncrementalSyncDisabled *bool `json:"incrementalSyncDisabled,omitempty"`

	// Optional. The interval for incremental syncs.
	IncrementalRefreshInterval *string `json:"incrementalRefreshInterval,omitempty"`

	// Optional. Specifies the data protection policy for the connector.
	DataProtectionPolicy *DataProtectionPolicy `json:"dataProtectionPolicy,omitempty"`

	// Optional. User-facing metadata.
	ConnectorMetadata *ConnectorMetadata `json:"connectorMetadata,omitempty"`

	// Optional. Configuration for real-time data synchronization.
	RealtimeSyncConfig *RealtimeSyncConfig `json:"realtimeSyncConfig,omitempty"`

	// Optional. The KMS key used for encryption.
	KMSKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}

// DiscoveryEngineDataConnectorStatus defines the config connector machine state of DiscoveryEngineDataConnector
type DiscoveryEngineDataConnectorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineDataConnector resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineDataConnectorObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineDataConnectorObservedState is the state of the DiscoveryEngineDataConnector resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1alpha.DataConnector
type DiscoveryEngineDataConnectorObservedState struct {
	// BlockingReasons represents the reasons why the connector is blocked from running.
	BlockingReasons []string `json:"blockingReasons,omitempty"`

	// LastRun represents the details about the last connector run.
	LastRun *ConnectorRun `json:"lastRun,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginedataconnector;gcpdiscoveryenginedataconnectors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineDataConnector is the Schema for the DiscoveryEngineDataConnector API
// +k8s:openapi-gen=true
type DiscoveryEngineDataConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineDataConnectorSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineDataConnectorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineDataConnectorList contains a list of DiscoveryEngineDataConnector
type DiscoveryEngineDataConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineDataConnector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineDataConnector{}, &DiscoveryEngineDataConnectorList{})
}
