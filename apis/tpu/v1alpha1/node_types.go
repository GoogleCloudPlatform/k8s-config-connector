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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TPUNodeGVK = GroupVersion.WithKind("TPUNode")

// TPUNodeSpec defines the desired state of TPUNode
// +kcc:proto=google.cloud.tpu.v1.Node
type TPUNodeSpec struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable.  The GCP location for the TPU. */
	// +required
	Zone *string `json:"location"`

	// The TPUNode name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The user-supplied description of the TPU. Maximum of 512 characters.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.description
	Description *string `json:"description,omitempty"`

	// Required. The type of hardware accelerators associated with this node.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.accelerator_type
	// +required
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// Required. The version of Tensorflow running in the Node.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.tensorflow_version
	// +required
	TensorflowVersion *string `json:"tensorflowVersion,omitempty"`

	// The name of a network they wish to peer the TPU node to. It must be a
	//  preexisting Compute Engine network inside of the project on which this API
	//  has been activated. If none is provided, "default" will be used.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.network
	Network *string `json:"network,omitempty"`

	// The CIDR block that the TPU node will use when selecting an IP address.
	//  This CIDR block must be a /29 block; the Compute Engine networks API
	//  forbids a smaller block, and using a larger block would be wasteful (a
	//  node can only consume one IP address). Errors will occur if the CIDR block
	//  has already been used for a currently existing TPU node, the CIDR block
	//  conflicts with any subnetworks in the user's provided network, or the
	//  provided network is peered with another network that is using that CIDR
	//  block.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.cidr_block
	CIDRBlock *string `json:"cidrBlock,omitempty"`

	// The scheduling options for this node.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.scheduling_config
	SchedulingConfig *SchedulingConfig `json:"schedulingConfig,omitempty"`

	/* NOTYET-LABELS
	// Resource labels to represent user-provided metadata.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.labels
	Labels map[string]string `json:"labels,omitempty"`
	*/

	// Whether the VPC peering for the node is set up through Service Networking
	//  API. The VPC Peering should be set up before provisioning the node.
	//  If this field is set, cidr_block field should not be specified. If the
	//  network, that you want to peer the TPU Node to, is Shared VPC networks,
	//  the node must be created with this this field enabled.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.use_service_networking
	UseServiceNetworking *bool `json:"useServiceNetworking,omitempty"`
}

// TPUNodeStatus defines the config connector machine state of TPUNode
type TPUNodeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TPUNode resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TPUNodeObservedState `json:"observedState,omitempty"`
}

// TPUNodeObservedState is the state of the TPUNode resource as most recently observed in GCP.
// +kcc:proto=google.cloud.tpu.v1.Node
type TPUNodeObservedState struct {

	/* DEPRECATED
	// Output only. DEPRECATED! Use network_endpoints instead.
	//  The network address for the TPU Node as visible to Compute Engine
	//  instances.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`
	*/

	/* DEPRECATED
	// Output only. DEPRECATED! Use network_endpoints instead.
	//  The network port for the TPU Node as visible to Compute Engine instances.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.port
	Port *string `json:"port,omitempty"`
	*/

	// Output only. The current state for the TPU Node.
	State *string `json:"state,omitempty"`

	// The health status of the TPU node.
	// +kcc:proto:field=google.cloud.tpu.v1.Node.health
	Health *string `json:"health,omitempty"`

	// Output only. If this field is populated, it contains a description of why
	// the TPU Node is unhealthy.
	HealthDescription *string `json:"healthDescription,omitempty"`

	// Output only. The service account used to run the tensor flow services
	// within the node. To share resources, including Google Cloud Storage data,
	// with the Tensorflow job running in the Node, this account must have
	// permissions to that data.
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	/* NOTYET-VOLATILE
	// Output only. The time when the node was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	*/

	// Output only. The network endpoints where TPU workers can be accessed and
	// sent work. It is recommended that Tensorflow clients of the node reach out
	// to the 0th entry in this map first.
	NetworkEndpoints []NetworkEndpoint `json:"networkEndpoints,omitempty"`

	// Output only. The API version that created this Node.
	APIVersion *string `json:"apiVersion,omitempty"`

	// Output only. The Symptoms that have occurred to the TPU Node.
	Symptoms []Symptom `json:"symptoms,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcptpunode;gcptpunodes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TPUNode is the Schema for the TPUNode API
// +k8s:openapi-gen=true
type TPUNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TPUNodeSpec   `json:"spec,omitempty"`
	Status TPUNodeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TPUNodeList contains a list of TPUNode
type TPUNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TPUNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TPUNode{}, &TPUNodeList{})
}
