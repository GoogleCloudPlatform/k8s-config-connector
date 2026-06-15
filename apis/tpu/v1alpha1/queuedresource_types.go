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

var TPUQueuedResourceGVK = GroupVersion.WithKind("TPUQueuedResource")

// TPUQueuedResourceSpec defines the desired state of TPUQueuedResource
// +kcc:spec:proto=google.cloud.tpu.v2.QueuedResource
type TPUQueuedResourceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The TPUQueuedResource name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Defines a TPU resource.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.tpu
	Tpu *QueuedResource_Tpu `json:"tpu,omitempty"`

	// Optional. The Spot tier.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.spot
	Spot *QueuedResource_Spot `json:"spot,omitempty"`

	// Optional. The Guaranteed tier.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.guaranteed
	Guaranteed *QueuedResource_Guaranteed `json:"guaranteed,omitempty"`

	// Optional. The queueing policy of the QueuedRequest.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.queueing_policy
	QueueingPolicy *QueuedResource_QueueingPolicy `json:"queueingPolicy,omitempty"`

	// Optional. Name of the reservation in which the resource should be provisioned.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.reservation_name
	ReservationName *string `json:"reservationName,omitempty"`
}

// TPUQueuedResourceStatus defines the config connector machine state of TPUQueuedResource
type TPUQueuedResourceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TPUQueuedResource resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TPUQueuedResourceObservedState `json:"observedState,omitempty"`
}

// TPUQueuedResourceObservedState is the state of the TPUQueuedResource resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.tpu.v2.QueuedResource
type TPUQueuedResourceObservedState struct {
	// Output only. The time when the QueuedResource was created.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. State of the QueuedResource request.
	// +kcc:proto:field=google.cloud.tpu.v2.QueuedResource.state
	State *QueuedResourceStateObservedState `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.tpu.v2.Node
type Node struct {
	// The user-supplied description of the TPU. Maximum of 512 characters.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.description
	Description *string `json:"description,omitempty"`

	// Optional. The type of hardware accelerators associated with this node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// Required. The runtime version running in the Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.runtime_version
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`

	// Network configurations for the TPU node. network_config and network_configs
	//  are mutually exclusive, you can only specify one of them. If both are
	//  specified, an error will be returned.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.network_config
	NetworkConfig *NetworkConfig `json:"networkConfig,omitempty"`

	// Optional. Repeated network configurations for the TPU node. This field is
	//  used to specify multiple networks configs for the TPU node. network_config
	//  and network_configs are mutually exclusive, you can only specify one of
	//  them. If both are specified, an error will be returned.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.network_configs
	NetworkConfigs []NetworkConfig `json:"networkConfigs,omitempty"`

	// The CIDR block that the TPU node will use when selecting an IP address.
	//  This CIDR block must be a /29 block; the Compute Engine networks API
	//  forbids a smaller block, and using a larger block would be wasteful (a
	//  node can only consume one IP address). Errors will occur if the CIDR block
	//  has already been used for a currently existing TPU node, the CIDR block
	//  conflicts with any subnetworks in the user's provided network, or the
	//  provided network is peered with another network that is using that CIDR
	//  block.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.cidr_block
	CIDRBlock *string `json:"cidrBlock,omitempty"`

	// The Google Cloud Platform Service Account to be used by the TPU node VMs.
	//  If None is specified, the default compute service account will be used.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.service_account
	ServiceAccount *ServiceAccount `json:"serviceAccount,omitempty"`

	// The scheduling options for this node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.scheduling_config
	SchedulingConfig *SchedulingConfig `json:"schedulingConfig,omitempty"`

	// Custom metadata to apply to the TPU Node.
	//  Can set startup-script and shutdown-script
	// +kcc:proto:field=google.cloud.tpu.v2.Node.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Tags to apply to the TPU Node. Tags are used to identify valid sources or
	//  targets for network firewalls.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.tags
	Tags []string `json:"tags,omitempty"`

	// The additional data disks for the Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.data_disks
	DataDisks []AttachedDisk `json:"dataDisks,omitempty"`

	// Shielded Instance options.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.shielded_instance_config
	ShieldedInstanceConfig *ShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	// The AccleratorConfig for the TPU Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.accelerator_config
	AcceleratorConfig *AcceleratorConfig `json:"acceleratorConfig,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.tpu.v2.Node
type NodeObservedState struct {
	// Output only. The name of the TPU node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.name
	Name *string `json:"name,omitempty"`

	// Output only. The current state for the TPU node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.state
	State *string `json:"state,omitempty"`

	// Output only. The health description for the TPU node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.health_description
	HealthDescription *string `json:"healthDescription,omitempty"`

	// Output only. The time when the TPU node was created.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The network endpoints of this TPU node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.network_endpoints
	NetworkEndpoints []NetworkEndpointObservedState `json:"networkEndpoints,omitempty"`

	// Output only. The unique identifier for the TPU node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.id
	Id *string `json:"id,omitempty"`

	// Output only. The API version that created this Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.api_version
	ApiVersion *string `json:"apiVersion,omitempty"`

	// Output only. The symptoms of the TPU node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.symptoms
	Symptoms []Symptom `json:"symptoms,omitempty"`

	// Output only. The qualified name of the QueuedResource that created this Node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.queued_resource
	QueuedResource *string `json:"queuedResource,omitempty"`

	// Output only. Whether the Node belongs to a Multislice group.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.multislice_node
	MultisliceNode *bool `json:"multisliceNode,omitempty"`

	// The health status of the TPU node.
	// +kcc:proto:field=google.cloud.tpu.v2.Node.health
	Health *string `json:"health,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptpuqueuedresource;gcptpuqueuedresources
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TPUQueuedResource is the Schema for the TPUQueuedResource API
// +k8s:openapi-gen=true
type TPUQueuedResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TPUQueuedResourceSpec   `json:"spec,omitempty"`
	Status TPUQueuedResourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TPUQueuedResourceList contains a list of TPUQueuedResource
type TPUQueuedResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TPUQueuedResource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TPUQueuedResource{}, &TPUQueuedResourceList{})
}
