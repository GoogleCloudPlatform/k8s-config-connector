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

// API sources for DiscoveryEngineDataConnector, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/4d40dcbbad29b0ad0927d748796c6bb207a277de/google/cloud/discoveryengine/v1/data_connector.proto
// +kcc:source:service-docs=https://cloud.google.com/generative-ai-app-builder/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/generative-ai-app-builder/docs/reference/rest/v1/projects.locations.collections.dataConnector

package v1alpha1

import (
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineDataConnectorGVK = GroupVersion.WithKind("DiscoveryEngineDataConnector")

// DiscoveryEngineDataConnectorSpec defines the desired state of DiscoveryEngineDataConnector
// +kcc:spec:proto=google.cloud.discoveryengine.v1.DataConnector
// +kcc:required-from-proto
type DiscoveryEngineDataConnectorSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DiscoveryEngineDataConnector name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required data connector parameters in structured json format.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.params
	Params apiextensionsv1.JSON `json:"params,omitempty"`

	// Required data connector parameters in json string format.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.json_params
	JsonParams *string `json:"jsonParams,omitempty"`

	// Required. The identifier for the data source.
	//  For the full, up-to-date list of supported connectors and their values,
	//  see
	//  [Connect a third-party data
	//  source](https://docs.cloud.google.com/gemini/enterprise/docs/connectors/connect-third-party-data-source#sources-by-launch-stage).
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.data_source
	// +required
	DataSource *string `json:"dataSource,omitempty"`

	// Optional. Specifies keys to be removed from the 'params' field.
	//  This is only active when 'params' is included in the 'update_mask'
	//  in an UpdateDataConnectorRequest.
	//  Deletion takes precedence if a key is both in 'remove_param_keys'
	//  and present in the 'params' field of the request.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.remove_param_keys
	RemoveParamKeys []string `json:"removeParamKeys,omitempty"`

	// Required. The refresh interval for data sync. If duration is set to 0, the
	//  data will be synced in real time. The streaming feature is not supported
	//  yet. The minimum is 30 minutes and maximum is 7 days. When the refresh
	//  interval is set to the same value as the incremental refresh interval,
	//  incremental sync will be disabled.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.refresh_interval
	// +required
	RefreshInterval *string `json:"refreshInterval,omitempty"`

	// Defines the scheduled time for the next data synchronization. This field
	//  requires hour , minute, and time_zone from the
	//  [IANA Time Zone Database](https://www.iana.org/time-zones). This is
	//  utilized when the data connector has a refresh interval greater than 1 day.
	//  When the hours or minutes are not specified, we will assume a sync time of
	//  0:00. The user must provide a time zone to avoid ambiguity.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.next_sync_time
	NextSyncTime *DateTime `json:"nextSyncTime,omitempty"`

	// List of entities from the connected data source to ingest.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.entities
	Entities []DataConnector_SourceEntity `json:"entities,omitempty"`

	// The data synchronization mode supported by the data connector.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.sync_mode
	SyncMode *string `json:"syncMode,omitempty"`

	// Input only. The KMS key to be used to protect the DataStores managed by
	//  this connector.
	//
	//  Must be set for requests that need to comply with CMEK Org Policy
	//  protections.
	//
	//  If this field is set and processed successfully, the DataStores created by
	//  this connector will be protected by the KMS key.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Optional. Any target destinations used to connect to third-party services.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.destination_configs
	DestinationConfigs []DestinationConfig `json:"destinationConfigs,omitempty"`

	// Optional. Action configurations to make the connector support actions.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.action_config
	ActionConfig *ActionConfig `json:"actionConfig,omitempty"`

	// Optional. Any params and credentials used specifically for EUA connectors.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.end_user_config
	EndUserConfig *DataConnector_EndUserConfig `json:"endUserConfig,omitempty"`

	// Optional. Any params and credentials used specifically for hybrid
	//  connectors supporting FEDERATED mode. This field should only be set if the
	//  connector is a hybrid connector and we want to enable FEDERATED mode.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.federated_config
	FederatedConfig *DataConnector_FederatedConfig `json:"federatedConfig,omitempty"`

	// Optional. The modes enabled for this connector. Default state is
	//  CONNECTOR_MODE_UNSPECIFIED.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.connector_modes
	ConnectorModes []string `json:"connectorModes,omitempty"`

	// Optional. Whether customer has enabled static IP addresses for this
	//  connector.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.static_ip_enabled
	StaticIPEnabled *bool `json:"staticIPEnabled,omitempty"`

	// The configuration for the identity data synchronization runs. This contains
	//  the refresh interval to sync the Access Control List information
	//  for the documents ingested by this connector.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.identity_schedule_config
	IdentityScheduleConfig *IdentityScheduleConfig `json:"identityScheduleConfig,omitempty"`

	// Optional. The configuration for establishing a BAP connection.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.bap_config
	BapConfig *BapConfig `json:"bapConfig,omitempty"`

	// Optional. The connector level alert config.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.alert_policy_configs
	AlertPolicyConfigs []AlertPolicyConfig `json:"alertPolicyConfigs,omitempty"`

	// Optional. The configuration for realtime sync.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.realtime_sync_config
	RealtimeSyncConfig *DataConnector_RealtimeSyncConfig `json:"realtimeSyncConfig,omitempty"`

	// Optional. Whether the connector will be created with an ACL config.
	//  Currently this field only affects Cloud Storage and BigQuery connectors.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.acl_enabled
	AclEnabled *bool `json:"aclEnabled,omitempty"`

	// Optional. Whether the END USER AUTHENTICATION connector is created in SaaS.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.create_eua_saas
	CreateEuaSaas *bool `json:"createEuaSaas,omitempty"`

	// Optional. Indicates whether the connector is disabled for auto run. It can
	//  be used to pause periodical and real time sync. Update: with the
	//  introduction of incremental_sync_disabled, auto_run_disabled is used to
	//  pause/disable only full syncs
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.auto_run_disabled
	AutoRunDisabled *bool `json:"autoRunDisabled,omitempty"`

	// Optional. Indicates whether incremental syncs are paused for this
	//  connector. This is independent of auto_run_disabled. Applicable to only 3P
	//  connectors. When the refresh interval is set to the same value as the
	//  incremental refresh interval, incremental sync will be disabled, i.e. set
	//  to true.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.incremental_sync_disabled
	IncrementalSyncDisabled *bool `json:"incrementalSyncDisabled,omitempty"`

	// Optional. The refresh interval specifically for incremental data syncs. If
	//  unset, incremental syncs will use the default from env, set to 3hrs. The
	//  minimum is 30 minutes and maximum is 7 days. Applicable to only 3P
	//  connectors. When the refresh interval is
	//  set to the same value as the incremental refresh interval, incremental
	//  sync will be disabled.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.incremental_refresh_interval
	IncrementalRefreshInterval *string `json:"incrementalRefreshInterval,omitempty"`

	// Optional. If the connector is a hybrid connector, determines whether
	//  ingestion is enabled and appropriate resources are provisioned during
	//  connector creation. If the connector is not a hybrid connector, this field
	//  is ignored.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.hybrid_ingestion_disabled
	HybridIngestionDisabled *bool `json:"hybridIngestionDisabled,omitempty"`

	// Optional. Specifies the data protection policy for the connector.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.data_protection_policy
	DataProtectionPolicy *DataProtectionPolicy `json:"dataProtectionPolicy,omitempty"`

	// Optional. The configuration for establishing a CLI connection.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.cli_config
	CliConfig *CliConfig `json:"cliConfig,omitempty"`

	// Optional. If set, this value instead of `data_source` is used to fetch the
	//  corresponding connector source.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.connector_source_id
	ConnectorSourceID *string `json:"connectorSourceID,omitempty"`

	// Optional. Immutable. User-facing, version-independent label for this
	//  connector. May be shared by multiple connectors under the same
	//  (project, location, collection, data_source); tag-based lookup returns
	//  the one with the greatest create_time.
	//
	//  Optional at Create time. If the caller omits `tag`, the server
	//  auto-derives one from the collection_id (falling back to data_source,
	//  else a time-based `t-<micros>` sentinel). The auto-derived tag is
	//  subject to the same immutability guarantee as a caller-supplied one,
	//  so callers who care about the exact tag value should provide it
	//  explicitly rather than relying on the server default.
	//
	//  Agent Designer resolves connectors via (data_source, tag) when set,
	//  falling back to the legacy resource-name lookup when unset, so
	//  connectors created before the tag-write launch continue to work
	//  without a backfill.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.tag
	Tag *string `json:"tag,omitempty"`

	// Optional. User-facing metadata for the connector. Populated from the
	//  connector's generated metadata / registry `ConnectorSource`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.metadata
	Metadata *DataConnector_ConnectorMetadata `json:"metadata,omitempty"`
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
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.DataConnector
type DiscoveryEngineDataConnectorObservedState struct {
	// Output only. State of the connector.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.state
	State *string `json:"state,omitempty"`

	// List of entities from the connected data source to ingest.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.entities
	Entities []DataConnector_SourceEntityObservedState `json:"entities,omitempty"`

	// Output only. Timestamp the
	//  [DataConnector][google.cloud.discoveryengine.v1.DataConnector] was created
	//  at.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent timestamp when this
	//  [DataConnector][google.cloud.discoveryengine.v1.DataConnector] was paused,
	//  affecting all functionalities such as data synchronization.
	//  Pausing a connector has the following effects:
	//
	//    - All functionalities, including data synchronization, are halted.
	//    - Any ongoing data synchronization job will be canceled.
	//    - No future data synchronization runs will be scheduled nor can be
	//    triggered.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.latest_pause_time
	LatestPauseTime *string `json:"latestPauseTime,omitempty"`

	// Output only. For periodic connectors only, the last time a data sync was
	//  completed.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.last_sync_time
	LastSyncTime *string `json:"lastSyncTime,omitempty"`

	// Output only. Timestamp the
	//  [DataConnector][google.cloud.discoveryengine.v1.DataConnector] was last
	//  updated.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The errors from initialization or from the latest connector
	//  run.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.errors
	Errors []common.Status `json:"errors,omitempty"`

	// Output only. State of the action connector. This reflects whether the
	//  action connector is initializing, active or has encountered errors.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.action_state
	ActionState *string `json:"actionState,omitempty"`

	// Optional. Action configurations to make the connector support actions.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.action_config
	ActionConfig *ActionConfigObservedState `json:"actionConfig,omitempty"`

	// Output only. The static IP addresses used by this connector.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.static_ip_addresses
	StaticIPAddresses []string `json:"staticIPAddresses,omitempty"`

	// Output only. The static IP addresses used by this connector for OAuth APIs
	//  (e.g. end user authentication). These are surfaced separately from
	//  `static_ip_addresses` so that customers can apply granular firewall
	//  settings for OAuth endpoints. Only populated for connectors that have
	//  static IP enabled and are used for actions and/or federated search.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.oauth_static_ip_addresses
	OauthStaticIPAddresses []string `json:"oauthStaticIPAddresses,omitempty"`

	// Output only. User actions that must be completed before the connector can
	//  start syncing data.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.blocking_reasons
	BlockingReasons []string `json:"blockingReasons,omitempty"`

	// Output only. The tenant project ID associated with private connectivity
	//  connectors. This project must be allowlisted by in order for the connector
	//  to function.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.private_connectivity_project_id
	PrivateConnectivityProjectID *string `json:"privateConnectivityProjectID,omitempty"`

	// Output only. The type of connector. Each source can only map to one type.
	//   For example, salesforce, confluence and jira have THIRD_PARTY connector
	//   type. It is not mutable once set by system.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.connector_type
	ConnectorType *string `json:"connectorType,omitempty"`

	// Output only. real-time sync state
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.realtime_state
	RealtimeState *string `json:"realtimeState,omitempty"`

	// Output only. The list of FQDNs of the data connector can egress to. This
	//  includes both FQDN derived from the customer provided instance URL and
	//  default per connector type FQDNs.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.egress_fqdns
	EgressFQDNs []string `json:"egressFQDNs,omitempty"`

	// Output only. Whether the connector is created with VPC-SC enabled.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.vpcsc_enabled
	VpcscEnabled *bool `json:"vpcscEnabled,omitempty"`

	// Output only. The dynamic tools fetched for this connector.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.DataConnector.dynamic_tools
	DynamicTools []DynamicToolObservedState `json:"dynamicTools,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginedataconnector;gcpdiscoveryenginedataconnectors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
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
