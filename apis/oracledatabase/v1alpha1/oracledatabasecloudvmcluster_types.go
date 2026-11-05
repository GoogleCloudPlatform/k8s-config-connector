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

// API sources for OracleDatabaseCloudVMCluster, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/054c87e45c4683ea5430ab609cbd4a445c4c4c71/google/cloud/oracledatabase/v1/vm_cluster.proto
// +kcc:source:service-docs=https://cloud.google.com/oracle/database/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/oracle/database/docs/reference/rest/v1/projects.locations.cloudVmClusters

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var OracleDatabaseCloudVMClusterGVK = GroupVersion.WithKind("OracleDatabaseCloudVMCluster")

// OracleDatabaseCloudVMClusterSpec defines the desired state of OracleDatabaseCloudVMCluster
// +kcc:spec:proto=google.cloud.oracledatabase.v1.CloudVmCluster
// +kcc:required-from-proto
type OracleDatabaseCloudVMClusterSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The OracleDatabaseCloudVMCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The name of the Exadata Infrastructure resource on which VM
	//  cluster resource is created, in the following format:
	//  projects/{project}/locations/{region}/cloudExadataInfrastuctures/{cloud_extradata_infrastructure}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.exadata_infrastructure
	// +required
	ExadataInfrastructure *string `json:"exadataInfrastructure,omitempty"`

	// Optional. User friendly name for this resource.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Various properties of the VM Cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.properties
	Properties *CloudVMClusterProperties `json:"properties,omitempty"`

	// Optional. Labels or tags associated with the VM Cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Network settings. CIDR to use for cluster IP allocation.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.cidr
	CIDR *string `json:"cidr,omitempty"`

	// Optional. CIDR range of the backup subnet.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.backup_subnet_cidr
	BackupSubnetCIDR *string `json:"backupSubnetCIDR,omitempty"`

	// Optional. The name of the VPC network.
	//  Format: projects/{project}/global/networks/{network}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.network
	Network *string `json:"network,omitempty"`

	// Optional. The name of the OdbNetwork associated with the VM Cluster.
	//  Format:
	//  projects/{project}/locations/{location}/odbNetworks/{odb_network}
	//  It is optional but if specified, this should match the parent ODBNetwork of
	//  the odb_subnet and backup_odb_subnet.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.odb_network
	OdbNetwork *string `json:"odbNetwork,omitempty"`

	// Optional. The name of the OdbSubnet associated with the VM Cluster for
	//  IP allocation. Format:
	//  projects/{project}/locations/{location}/odbNetworks/{odb_network}/odbSubnets/{odb_subnet}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.odb_subnet
	OdbSubnet *string `json:"odbSubnet,omitempty"`

	// Optional. The name of the backup OdbSubnet associated with the VM Cluster.
	//  Format:
	//  projects/{project}/locations/{location}/odbNetworks/{odb_network}/odbSubnets/{odb_subnet}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.backup_odb_subnet
	BackupOdbSubnet *string `json:"backupOdbSubnet,omitempty"`

	// Optional. The name of ExascaleDbStorageVault associated with the VM
	//  Cluster. Format:
	//  projects/{project}/locations/{location}/exascaleDbStorageVaults/{exascale_db_storage_vault}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.exascale_db_storage_vault
	ExascaleDbStorageVault *string `json:"exascaleDbStorageVault,omitempty"`
}

// OracleDatabaseCloudVMClusterStatus defines the config connector machine state of OracleDatabaseCloudVMCluster
type OracleDatabaseCloudVMClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the OracleDatabaseCloudVMCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *OracleDatabaseCloudVMClusterObservedState `json:"observedState,omitempty"`
}

// OracleDatabaseCloudVMClusterObservedState is the state of the OracleDatabaseCloudVMCluster resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.oracledatabase.v1.CloudVmCluster
type OracleDatabaseCloudVMClusterObservedState struct {
	// Optional. Various properties of the VM Cluster.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.properties
	Properties *CloudVMClusterPropertiesObservedState `json:"properties,omitempty"`

	// Output only. The date and time that the VM cluster was created.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The GCP Oracle zone where Oracle CloudVmCluster is hosted.
	//  This will be the same as the gcp_oracle_zone of the
	//  CloudExadataInfrastructure. Example: us-east4-b-r2.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.gcp_oracle_zone
	GcpOracleZone *string `json:"gcpOracleZone,omitempty"`

	// Output only. The identity connector details which will allow OCI to
	//  securely access the resources in the customer project.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudVmCluster.identity_connector
	IdentityConnector *IdentityConnectorObservedState `json:"identityConnector,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcporacledatabasecloudvmcluster;gcporacledatabasecloudvmclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// OracleDatabaseCloudVMCluster is the Schema for the OracleDatabaseCloudVMCluster API
// +k8s:openapi-gen=true
type OracleDatabaseCloudVMCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   OracleDatabaseCloudVMClusterSpec   `json:"spec,omitempty"`
	Status OracleDatabaseCloudVMClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// OracleDatabaseCloudVMClusterList contains a list of OracleDatabaseCloudVMCluster
type OracleDatabaseCloudVMClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OracleDatabaseCloudVMCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OracleDatabaseCloudVMCluster{}, &OracleDatabaseCloudVMClusterList{})
}
