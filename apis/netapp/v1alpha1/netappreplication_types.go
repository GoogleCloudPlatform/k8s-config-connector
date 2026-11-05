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

// API sources for NetappReplication, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e09e85d32ca349e1b205514a412817a7692595c6/google/cloud/netapp/v1/replication.proto
// +kcc:source:service-docs=https://cloud.google.com/netapp/
// +kcc:source:resource-docs=https://docs.cloud.google.com/netapp/volumes/docs/reference/rest/v1/projects.locations.volumes.replications

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetappReplicationGVK = GroupVersion.WithKind("NetappReplication")

// NetappReplicationSpec defines the desired state of NetappReplication
// +kcc:spec:proto=google.cloud.netapp.v1.Replication
// +kcc:required-from-proto
type NetappReplicationSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetappReplication name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Indicates the schedule for replication.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.replication_schedule
	// +required
	ReplicationSchedule *string `json:"replicationSchedule,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.labels
	Labels map[string]string `json:"labels,omitempty"`

	// A description about this replication relationship.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.description
	Description *string `json:"description,omitempty"`

	// Required. Input only. Destination volume parameters
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.destination_volume_parameters
	// +required
	DestinationVolumeParameters *DestinationVolumeParameters `json:"destinationVolumeParameters,omitempty"`

	// Optional. Location of the user cluster.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.cluster_location
	ClusterLocation *string `json:"clusterLocation,omitempty"`
}

// NetappReplicationStatus defines the config connector machine state of NetappReplication
type NetappReplicationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetappReplication resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetappReplicationObservedState `json:"observedState,omitempty"`
}

// NetappReplicationObservedState is the state of the NetappReplication resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.netapp.v1.Replication
type NetappReplicationObservedState struct {
	// Output only. State of the replication.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.state
	State *string `json:"state,omitempty"`

	// Output only. State details of the replication.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.state_details
	StateDetails *string `json:"stateDetails,omitempty"`

	// Output only. Indicates whether this points to source or destination.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.role
	Role *string `json:"role,omitempty"`

	// Output only. Indicates the state of mirroring.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.mirror_state
	MirrorState *string `json:"mirrorState,omitempty"`

	// Output only. Condition of the relationship. Can be one of the following:
	//  - true: The replication relationship is healthy. It has not missed the most
	//  recent scheduled transfer.
	//  - false: The replication relationship is not healthy. It has missed the
	//  most recent scheduled transfer.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.healthy
	Healthy *bool `json:"healthy,omitempty"`

	// Output only. Replication create time.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Full name of destination volume resource.
	//  Example : "projects/{project}/locations/{location}/volumes/{volume_id}"
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.destination_volume
	DestinationVolume *string `json:"destinationVolume,omitempty"`

	// Output only. Replication transfer statistics.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.transfer_stats
	TransferStats *TransferStats `json:"transferStats,omitempty"`

	// Output only. Full name of source volume resource.
	//  Example : "projects/{project}/locations/{location}/volumes/{volume_id}"
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.source_volume
	SourceVolume *string `json:"sourceVolume,omitempty"`

	// Output only. Hybrid peering details.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.hybrid_peering_details
	HybridPeeringDetails *HybridPeeringDetailsObservedState `json:"hybridPeeringDetails,omitempty"`

	// Output only. Type of the hybrid replication.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.hybrid_replication_type
	HybridReplicationType *string `json:"hybridReplicationType,omitempty"`

	// Output only. Copy pastable snapmirror commands to be executed on onprem
	//  cluster by the customer.
	// +kcc:proto:field=google.cloud.netapp.v1.Replication.hybrid_replication_user_commands
	HybridReplicationUserCommands *UserCommandsObservedState `json:"hybridReplicationUserCommands,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetappreplication;gcpnetappreplications
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetappReplication is the Schema for the NetappReplication API
// +k8s:openapi-gen=true
type NetappReplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetappReplicationSpec   `json:"spec,omitempty"`
	Status NetappReplicationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetappReplicationList contains a list of NetappReplication
type NetappReplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetappReplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetappReplication{}, &NetappReplicationList{})
}
