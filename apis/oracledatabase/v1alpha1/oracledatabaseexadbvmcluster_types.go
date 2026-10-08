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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var OracleDatabaseExadbVMClusterGVK = GroupVersion.WithKind("OracleDatabaseExadbVMCluster")

// OracleDatabaseExadbVMClusterSpec defines the desired state of OracleDatabaseExadbVMCluster
// +kcc:spec:proto=google.cloud.oracledatabase.v1.ExadbVmCluster
type OracleDatabaseExadbVMClusterSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The OracleDatabaseExadbVMCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Immutable. The display name for the ExadbVmCluster. The name does
	//  not have to be unique within your project. The name must be 1-255
	//  characters long and can only contain alphanumeric characters.
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The properties of the ExadbVmCluster.
	// +required
	Properties *ExadbVMClusterProperties `json:"properties,omitempty"`

	// Optional. The labels or tags associated with the ExadbVmCluster.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Immutable. The OdbNetwork associated with the ExadbVmCluster.
	// +optional
	OdbNetworkRef *OracleDatabaseODBNetworkRef `json:"odbNetworkRef,omitempty"`

	// Required. Immutable. The OdbSubnet associated with the ExadbVmCluster for IP allocation.
	// +required
	OdbSubnetRef *OracleDatabaseODBSubnetRef `json:"odbSubnetRef,omitempty"`

	// Required. Immutable. The backup OdbSubnet associated with the ExadbVmCluster.
	// +required
	BackupOdbSubnetRef *OracleDatabaseODBSubnetRef `json:"backupOdbSubnetRef,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.ExadbVmClusterProperties
type ExadbVMClusterProperties struct {
	// Optional. Immutable. The cluster name for Exascale vm cluster. The cluster
	//  name must begin with an alphabetic character and may contain hyphens(-) but
	//  can not contain underscores(_). It should be not more than 11 characters
	//  and is not case sensitive. OCI Cluster name.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.cluster_name
	ClusterName *string `json:"clusterName,omitempty"`

	// Required. Immutable. Grid Infrastructure Version.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.grid_image_id
	GridImageID *string `json:"gridImageID,omitempty"`

	// Required. The number of nodes/VMs in the ExadbVmCluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Required. Immutable. The number of ECPUs enabled per node for an exadata vm
	//  cluster on exascale infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.enabled_ecpu_count_per_node
	EnabledEcpuCountPerNode *int32 `json:"enabledEcpuCountPerNode,omitempty"`

	// Optional. Immutable. The number of additional ECPUs per node for an Exadata
	//  VM cluster on exascale infrastructure.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.additional_ecpu_count_per_node
	AdditionalEcpuCountPerNode *int32 `json:"additionalEcpuCountPerNode,omitempty"`

	// Required. Immutable. Total storage details for the ExadbVmCluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.vm_file_system_storage
	VMFileSystemStorage *ExadbVMClusterStorageDetails `json:"vmFileSystemStorage,omitempty"`

	// Optional. Immutable. The license type of the ExadbVmCluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.license_model
	LicenseModel *string `json:"licenseModel,omitempty"`

	// Required. Immutable. The ExascaleDbStorageVault associated with the ExadbVmCluster.
	ExascaleDbStorageVaultRef *OracleDatabaseExascaleDBStorageVaultRef `json:"exascaleDbStorageVaultRef,omitempty"`

	// Required. Immutable. Prefix for VM cluster host names.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.hostname_prefix
	HostnamePrefix *string `json:"hostnamePrefix,omitempty"`

	// Required. Immutable. The SSH public keys for the ExadbVmCluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.ssh_public_keys
	SSHPublicKeys []string `json:"sshPublicKeys,omitempty"`

	// Optional. Immutable. Indicates user preference for data collection options.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.data_collection_options
	DataCollectionOptions *DataCollectionOptionsCommon `json:"dataCollectionOptions,omitempty"`

	// Optional. Immutable. The time zone of the ExadbVmCluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.time_zone
	TimeZone *TimeZone `json:"timeZone,omitempty"`

	// Required. Immutable. The shape attribute of the VM cluster. The type of
	//  Exascale storage used for Exadata VM cluster. The default is SMART_STORAGE
	//  which supports Oracle Database 23ai and later
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.shape_attribute
	ShapeAttribute *string `json:"shapeAttribute,omitempty"`

	// Optional. Immutable. SCAN listener port - TCP
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.scan_listener_port_tcp
	ScanListenerPortTCP *int32 `json:"scanListenerPortTCP,omitempty"`
}

// +kcc:proto=google.cloud.oracledatabase.v1.ExadbVmClusterStorageDetails
type ExadbVMClusterStorageDetails struct {
	// Required. The storage allocation for the exadbvmcluster per node, in
	//  gigabytes (GB). This field is used to calculate the total storage
	//  allocation for the exadbvmcluster.
	SizeInGBsPerNode *int32 `json:"sizeInGBsPerNode,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.oracledatabase.v1.ExadbVmClusterProperties
type ExadbVMClusterPropertiesObservedState struct {
	// Output only. The hostname of the ExadbVmCluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.hostname
	Hostname *string `json:"hostname,omitempty"`

	// Output only. State of the cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.lifecycle_state
	LifecycleState *string `json:"lifecycleState,omitempty"`

	// Output only. Memory per VM (GB) (Read-only): Shows the amount of memory
	//  allocated to each VM. Memory is calculated based on 2.75 GB per Total
	//  ECPUs.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.memory_size_gb
	MemorySizeGB *int32 `json:"memorySizeGB,omitempty"`

	// Output only. Deep link to the OCI console to view this resource.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.oci_uri
	OciURI *string `json:"ociURI,omitempty"`

	// Output only. The Oracle Grid Infrastructure (GI) software version.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.ExadbVmClusterProperties.gi_version
	GiVersion *string `json:"giVersion,omitempty"`
}

// OracleDatabaseExadbVMClusterStatus defines the config connector machine state of OracleDatabaseExadbVMCluster
type OracleDatabaseExadbVMClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the OracleDatabaseExadbVMCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *OracleDatabaseExadbVMClusterObservedState `json:"observedState,omitempty"`
}

// OracleDatabaseExadbVMClusterObservedState is the state of the OracleDatabaseExadbVMCluster resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.oracledatabase.v1.ExadbVmCluster
type OracleDatabaseExadbVMClusterObservedState struct {
	// Required. The properties of the ExadbVmCluster.
	Properties *ExadbVMClusterPropertiesObservedState `json:"properties,omitempty"`

	// Output only. Immutable. The GCP Oracle zone where Oracle ExadbVmCluster is
	//  hosted. Example: us-east4-b-r2. During creation, the system will pick the
	//  zone assigned to the ExascaleDbStorageVault.
	GcpOracleZone *string `json:"gcpOracleZone,omitempty"`

	// Output only. The date and time that the ExadbVmCluster was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The ID of the subscription entitlement associated with the
	//  ExadbVmCluster.
	EntitlementID *string `json:"entitlementID,omitempty"`

	// Output only. The identity connector details which will allow OCI to
	//  securely access the resources in the customer project.
	IdentityConnector *IdentityConnectorObservedState `json:"identityConnector,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcporacledatabaseexadbvmcluster;gcporacledatabaseexadbvmclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// OracleDatabaseExadbVMCluster is the Schema for the OracleDatabaseExadbVMCluster API
// +k8s:openapi-gen=true
type OracleDatabaseExadbVMCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   OracleDatabaseExadbVMClusterSpec   `json:"spec,omitempty"`
	Status OracleDatabaseExadbVMClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// OracleDatabaseExadbVMClusterList contains a list of OracleDatabaseExadbVMCluster
type OracleDatabaseExadbVMClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OracleDatabaseExadbVMCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OracleDatabaseExadbVMCluster{}, &OracleDatabaseExadbVMClusterList{})
}
