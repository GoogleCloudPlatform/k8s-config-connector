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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineDataConnectorGVK = GroupVersion.WithKind("DiscoveryEngineDataConnector")

// DiscoveryEngineDataConnectorSpec defines the desired state of DiscoveryEngineDataConnector
// +kcc:spec:proto=google.cloud.discoveryengine.v1alpha.DataConnector
type DiscoveryEngineDataConnectorSpec struct {
	// The DiscoveryEngineDataConnector name. Since it is a singleton resource under the collection, this name should be exactly "dataConnector" or is ignored.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Immutable. The collection ID for the DataConnector.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="CollectionID field is immutable"
	// +required
	CollectionID *string `json:"collectionID,omitempty"`

	// Required. Immutable. The location for the resource.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// +required
	Location *string `json:"location,omitempty"`

	// Immutable. The Project that this resource belongs to.
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// Required. The identifier for the data source (e.g., 'bigquery', 'google_drive', 'onedrive', 'jira', 'salesforce').
	// +required
	DataSource *string `json:"dataSource,omitempty"`

	// Required. The refresh interval for data sync. Minimum 30 minutes, maximum 7 days.
	// +required
	RefreshInterval *string `json:"refreshInterval,omitempty"`

	// Action configuration for the data connector.
	ActionConfig *DataConnector_ActionConfig `json:"actionConfig,omitempty"`

	// Indicates whether full syncs are paused for this connector.
	AutoRunDisabled *bool `json:"autoRunDisabled,omitempty"`

	// BAP configuration for the data connector.
	BapConfig *DataConnector_BapConfig `json:"bapConfig,omitempty"`

	// The modes enabled for this connector (e.g., 'DATA_INGESTION', 'ACTIONS').
	ConnectorModes []string `json:"connectorModes,omitempty"`

	// The version of the data source.
	DataSourceVersion *int64 `json:"dataSourceVersion,omitempty"`

	// Destination connector configurations for the data connector.
	DestinationConfigs []DataConnector_DestinationConfig `json:"destinationConfigs,omitempty"`

	// List of entities from the connected data source to ingest.
	Entities []DataConnector_SourceEntity `json:"entities,omitempty"`

	// The refresh interval specifically for incremental data syncs.
	IncrementalRefreshInterval *string `json:"incrementalRefreshInterval,omitempty"`

	// Indicates whether incremental syncs are paused for this connector.
	IncrementalSyncDisabled *bool `json:"incrementalSyncDisabled,omitempty"`

	// Params needed to access the source in JSON string format.
	JsonParams *string `json:"jsonParams,omitempty"`

	// The KMS key to be used to protect the DataStores managed by this connector.
	KmsKeyName *string `json:"kmsKeyName,omitempty"`

	// Params needed to access the source in String-to-String format.
	Params map[string]string `json:"params,omitempty"`

	// Whether customer has enabled static IP addresses for this connector.
	StaticIPEnabled *bool `json:"staticIPEnabled,omitempty"`

	// The data synchronization mode supported by the data connector.
	SyncMode *string `json:"syncMode,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DataConnector.ActionConfig
type DataConnector_ActionConfig struct {
	// Params needed to configure the actions in String-to-String format.
	ActionParams map[string]string `json:"actionParams,omitempty"`

	// Whether to create a BAP (Business Application Platform) connection.
	CreateBapConnection *bool `json:"createBapConnection,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DataConnector.BapConfig
type DataConnector_BapConfig struct {
	// The list of enabled actions for this connector.
	EnabledActions []string `json:"enabledActions,omitempty"`

	// The connector modes supported by the BAP configuration.
	SupportedConnectorModes []string `json:"supportedConnectorModes,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DataConnector.DestinationConfig
type DataConnector_DestinationConfig struct {
	// The list of destinations for this configuration.
	Destinations []DataConnector_Destination `json:"destinations,omitempty"`

	// The key of the destination configuration.
	Key *string `json:"key,omitempty"`

	// Additional parameters for this destination config in structured JSON format.
	Params *string `json:"params,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DataConnector.Destination
type DataConnector_Destination struct {
	// The host of the destination.
	Host *string `json:"host,omitempty"`

	// Target port number accepted by the destination.
	Port *int64 `json:"port,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DataConnector.SourceEntity
type DataConnector_SourceEntity struct {
	// The name of the entity.
	EntityName *string `json:"entityName,omitempty"`

	// Attributes for indexing (e.g., mapping fields to 'title' or 'description').
	KeyPropertyMappings map[string]string `json:"keyPropertyMappings,omitempty"`

	// The parameters for the entity to facilitate data ingestion.
	Params *string `json:"params,omitempty"`
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
	// Output only. State of the action connector.
	ActionState *string `json:"actionState,omitempty"`

	// Output only. User actions that must be completed before the connector can start syncing data.
	BlockingReasons []string `json:"blockingReasons,omitempty"`

	// Output only. The type of connector.
	ConnectorType *string `json:"connectorType,omitempty"`

	// Output only. Timestamp when the DataConnector was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The errors from initialization or from the latest connector run.
	Errors []DataConnector_Error `json:"errors,omitempty"`

	// Output only. For periodic connectors only, the last time a data sync was completed.
	LastSyncTime *string `json:"lastSyncTime,omitempty"`

	// Output only. The most recent timestamp when this DataConnector was paused.
	LatestPauseTime *string `json:"latestPauseTime,omitempty"`

	// Output only. The full resource name of the Data Connector.
	Name *string `json:"name,omitempty"`

	// Output only. The tenant project ID associated with private connectivity connectors.
	PrivateConnectivityProjectID *string `json:"privateConnectivityProjectID,omitempty"`

	// Output only. The real-time sync state.
	RealtimeState *string `json:"realtimeState,omitempty"`

	// Output only. The state of the connector.
	State *string `json:"state,omitempty"`

	// Output only. The static IP addresses used by this connector.
	StaticIPAddresses []string `json:"staticIPAddresses,omitempty"`

	// Output only. Timestamp when the DataConnector was updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Whether the action connector is fully configured.
	IsActionConfigured *bool `json:"isActionConfigured,omitempty"`

	// Output only. Entities statuses.
	EntitiesStatus []DataConnector_SourceEntityStatus `json:"entitiesStatus,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DataConnector.Error
type DataConnector_Error struct {
	// The status code (enum value of google.rpc.Code).
	Code *int64 `json:"code,omitempty"`

	// A developer-facing error message.
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DataConnector.SourceEntityStatus
type DataConnector_SourceEntityStatus struct {
	// The name of the entity.
	EntityName *string `json:"entityName,omitempty"`

	// The full resource name of the associated data store for the source entity.
	DataStore *string `json:"dataStore,omitempty"`
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
