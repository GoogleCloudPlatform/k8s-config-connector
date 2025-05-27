// Copyright 2024 Google LLC
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
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ManagedKafkaClusterGVK = GroupVersion.WithKind("ManagedKafkaCluster")

// +kcc:proto=google.cloud.managedkafka.v1.AccessConfig
type AccessConfig struct {
	// Required. Virtual Private Cloud (VPC) networks that must be granted direct
	//  access to the Kafka cluster. Minimum of 1 network is required. Maximum 10
	//  networks can be specified.
	// +kcc:proto:field=google.cloud.managedkafka.v1.AccessConfig.network_configs
	// +required
	NetworkConfigs []NetworkConfig `json:"networkConfigs,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.CapacityConfig
type CapacityConfig struct {
	// Required. The number of vCPUs to provision for the cluster. Minimum: 3.
	// +kcc:proto:field=google.cloud.managedkafka.v1.CapacityConfig.vcpu_count
	// +required
	VcpuCount *int64 `json:"vcpuCount,omitempty"`

	// Required. The memory to provision for the cluster in bytes.
	//  The CPU:memory ratio (vCPU:GiB) must be between 1:1 and 1:8.
	//  Minimum: 3221225472 (3 GiB).
	// +kcc:proto:field=google.cloud.managedkafka.v1.CapacityConfig.memory_bytes
	// +required
	MemoryBytes *int64 `json:"memoryBytes,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.NetworkConfig
type NetworkConfig struct {
	// Required. Reference to the VPC subnet in which to create Private Service Connect
	//  (PSC) endpoints for the Kafka brokers and bootstrap address.
	//
	//  The subnet must be located in the same region as the Kafka cluster. The
	//  project may differ. Multiple subnets from the same parent network must not
	//  be specified.
	//
	//  The CIDR range of the subnet must be within the IPv4 address ranges for
	//  private networks, as specified in RFC 1918.
	// +kcc:proto:field=google.cloud.managedkafka.v1.NetworkConfig.subnet
	// +required
	SubnetworkRef *refs.ComputeSubnetworkRef `json:"subnetworkRef"`
}

// +kcc:proto=google.cloud.managedkafka.v1.GcpConfig
type GcpConfig struct {
	// Required. Access configuration for the Kafka cluster.
	// +kcc:proto:field=google.cloud.managedkafka.v1.GcpConfig.access_config
	// +required
	AccessConfig *AccessConfig `json:"accessConfig,omitempty"`

	// Optional. Immutable. The Cloud KMS Key name to use for encryption. The key
	//  must be located in the same region as the cluster and cannot be changed.
	// +kcc:proto:field=google.cloud.managedkafka.v1.GcpConfig.kms_key
	KMSKeyRef *kmsv1beta1.KMSKeyRef_OneOf `json:"kmsKeyRef,omitempty"`
}

// ManagedKafkaClusterSpec defines the desired state of ManagedKafkaCluster
// +kcc:spec:proto=google.cloud.managedkafka.v1.Cluster
type ManagedKafkaClusterSpec struct {
	commonv1alpha1.CommonSpec `json:",inline"`

	// +required
	Location string `json:"location"`

	// The ManagedKafkaCluster name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Configuration properties for a Kafka cluster deployed to Google
	//  Cloud Platform.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.gcp_config
	// +required
	GcpConfig *GcpConfig `json:"gcpConfig,omitempty"`

	// Optional. Labels as key value pairs.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Capacity configuration for the Kafka cluster.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.capacity_config
	// +required
	CapacityConfig *CapacityConfig `json:"capacityConfig,omitempty"`

	// Optional. Rebalance configuration for the Kafka cluster.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.rebalance_config
	RebalanceConfig *RebalanceConfig `json:"rebalanceConfig,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.RebalanceConfig
type RebalanceConfig struct {
	// Optional. The rebalance behavior for the cluster.
	//  When not specified, defaults to `NO_REBALANCE`.
	// +kcc:proto:field=google.cloud.managedkafka.v1.RebalanceConfig.mode
	Mode *string `json:"mode,omitempty"`
}

// ManagedKafkaClusterStatus defines the config connector machine state of ManagedKafkaCluster
type ManagedKafkaClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ManagedKafkaCluster resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ManagedKafkaClusterObservedState `json:"observedState,omitempty"`
}

// ManagedKafkaClusterObservedState is the state of the ManagedKafkaCluster resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.managedkafka.v1.Cluster
type ManagedKafkaClusterObservedState struct {
	// Output only. The time when the cluster was created.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the cluster was last updated.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The current state of the cluster.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.state
	State *string `json:"state,omitempty"`

	// NOTYET
	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.satisfies_pzi
	// SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// NOTYET
	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.managedkafka.v1.Cluster.satisfies_pzs
	// SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmanagedkafkacluster;gcpmanagedkafkaclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ManagedKafkaCluster is the Schema for the ManagedKafkaCluster API
// +k8s:openapi-gen=true
type ManagedKafkaCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ManagedKafkaClusterSpec   `json:"spec,omitempty"`
	Status ManagedKafkaClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ManagedKafkaClusterList contains a list of ManagedKafkaCluster
type ManagedKafkaClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedKafkaCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedKafkaCluster{}, &ManagedKafkaClusterList{})
}
