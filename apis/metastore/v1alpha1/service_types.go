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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MetastoreServiceGVK = GroupVersion.WithKind("MetastoreService")

// MetastoreServiceSpec defines the desired state of MetastoreService
// +kcc:proto=google.cloud.metastore.v1.Service
type MetastoreServiceSpec struct {
	// The MetastoreService name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Configuration information specific to running Hive metastore
	//  software as the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.hive_metastore_config
	HiveMetastoreConfig *HiveMetastoreConfig `json:"hiveMetastoreConfig,omitempty"`

	// User-defined labels for the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The relative resource name of the VPC network on which the
	//  instance can be accessed. It is specified in the following form:
	//
	//  `projects/{project_number}/global/networks/{network_id}`.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.network
	// +k8s:config:google.com/references: "ComputeNetwork" // Assuming it references ComputeNetwork, adjust if needed
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	// The TCP port at which the metastore service is reached. Default: 9083.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.port
	Port *int32 `json:"port,omitempty"`

	// The tier of the service.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.tier
	Tier *string `json:"tier,omitempty"`

	// The one hour maintenance window of the metastore service. This specifies
	//  when the service can be restarted for maintenance purposes in UTC time.
	//  Maintenance window is not needed for services with the SPANNER
	//  database type.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.maintenance_window
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`

	// Immutable. The release channel of the service.
	//  If unspecified, defaults to `STABLE`.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.release_channel
	ReleaseChannel *string `json:"releaseChannel,omitempty"`

	// Immutable. Information used to configure the Dataproc Metastore service to
	//  encrypt customer data at rest. Cannot be updated.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.encryption_config
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// The configuration specifying the network settings for the
	//  Dataproc Metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.network_config
	NetworkConfig *NetworkConfig `json:"networkConfig,omitempty"`

	// Immutable. The database type that the Metastore service stores its data.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.database_type
	DatabaseType *string `json:"databaseType,omitempty"`

	// The configuration specifying telemetry settings for the Dataproc Metastore
	//  service. If unspecified defaults to `JSON`.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.telemetry_config
	TelemetryConfig *TelemetryConfig `json:"telemetryConfig,omitempty"`

	// Scaling configuration of the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.scaling_config
	ScalingConfig *ScalingConfig `json:"scalingConfig,omitempty"`
}

// MetastoreServiceStatus defines the config connector machine state of MetastoreService
type MetastoreServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MetastoreService resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MetastoreServiceObservedState `json:"observedState,omitempty"`
}

// MetastoreServiceObservedState is the state of the MetastoreService resource as most recently observed in GCP.
// +kcc:proto=google.cloud.metastore.v1.Service
type MetastoreServiceObservedState struct {
	// Output only. The time when the metastore service was created.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the metastore service was last updated.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The URI of the endpoint used to access the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.endpoint_uri
	EndpointURI *string `json:"endpointURI,omitempty"`

	// Output only. The current state of the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the current state of the
	//  metastore service, if available.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. A Cloud Storage URI (starting with `gs://`) that specifies
	//  where artifacts related to the metastore service are stored.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.artifact_gcs_uri
	ArtifactGCSURI *string `json:"artifactGCSURI,omitempty"`

	// Output only. The globally unique resource identifier of the metastore
	//  service.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The metadata management activities of the metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.metadata_management_activity
	MetadataManagementActivity *MetadataManagementActivityObservedState `json:"metadataManagementActivity,omitempty"` // Adjusted to use observed state type

	// The configuration specifying the network settings for the
	//  Dataproc Metastore service.
	// +kcc:proto:field=google.cloud.metastore.v1.Service.network_config
	NetworkConfig *NetworkConfigObservedState `json:"networkConfig,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmetastoreservice;gcpmetastoreservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MetastoreService is the Schema for the MetastoreService API
// +k8s:openapi-gen=true
type MetastoreService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MetastoreServiceSpec   `json:"spec,omitempty"`
	Status MetastoreServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MetastoreServiceList contains a list of MetastoreService
type MetastoreServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetastoreService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MetastoreService{}, &MetastoreServiceList{})
}
