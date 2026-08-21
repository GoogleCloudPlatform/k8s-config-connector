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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeNodeTemplateGVK is the GroupVersionKind for the ComputeNodeTemplate resource.
var ComputeNodeTemplateGVK = GroupVersion.WithKind("ComputeNodeTemplate")

// +kcc:proto=google.cloud.compute.v1.NodeTemplateNodeTypeFlexibility
type NodeTemplateNodeTypeFlexibility struct {
	/* Immutable. Number of virtual CPUs to use. */
	// +kcc:proto:field=google.cloud.compute.v1.NodeTemplateNodeTypeFlexibility.cpus
	Cpus *string `json:"cpus,omitempty"`

	/* Use local SSD. */
	// +kcc:proto:field=google.cloud.compute.v1.NodeTemplateNodeTypeFlexibility.local_ssd
	LocalSsd *string `json:"localSsd,omitempty"`

	/* Immutable. Physical memory available to the node, defined in MB. */
	// +kcc:proto:field=google.cloud.compute.v1.NodeTemplateNodeTypeFlexibility.memory
	Memory *string `json:"memory,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.ServerBinding
type NodeTemplateServerBinding struct {
	/* Immutable. Type of server binding policy. If 'RESTART_NODE_ON_ANY_SERVER',
	nodes using this template will restart on any physical server
	following a maintenance event.

	If 'RESTART_NODE_ON_MINIMAL_SERVER', nodes using this template
	will restart on the same physical server following a maintenance
	event, instead of being live migrated to or restarted on a new
	physical server. This option may be useful if you are using
	software licenses tied to the underlying server characteristics
	such as physical sockets or cores, to avoid the need for
	additional licenses when maintenance occurs. However, VMs on such
	nodes will experience outages while maintenance is applied. Possible values: ["RESTART_NODE_ON_ANY_SERVER", "RESTART_NODE_ON_MINIMAL_SERVERS"]. */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.ServerBinding.type
	Type *string `json:"type"`
}

// ComputeNodeTemplateSpec defines the desired state of ComputeNodeTemplate
// +kcc:spec:proto=google.cloud.compute.v1.NodeTemplate
type ComputeNodeTemplateSpec struct {
	/* Immutable. CPU overcommit. Default value: "NONE" Possible
	values: ["ENABLED", "NONE"]. */
	// +kcc:proto:field=google.cloud.compute.v1.NodeTemplate.cpu_overcommit_type
	CPUOvercommitType *string `json:"cpuOvercommitType,omitempty"`

	/* Immutable. An optional textual description of the resource. */
	// +kcc:proto:field=google.cloud.compute.v1.NodeTemplate.description
	Description *string `json:"description,omitempty"`

	/* Immutable. Node type to use for nodes group that are created from this template.
	Only one of nodeTypeFlexibility and nodeType can be specified. */
	// +kcc:proto:field=google.cloud.compute.v1.NodeTemplate.node_type
	NodeType *string `json:"nodeType,omitempty"`

	/* Immutable. Flexible properties for the desired node type. Node groups that
	use this node template will create nodes of a type that matches
	these properties. Only one of nodeTypeFlexibility and nodeType can
	be specified. */
	// +kcc:proto:field=google.cloud.compute.v1.NodeTemplate.node_type_flexibility
	NodeTypeFlexibility *NodeTemplateNodeTypeFlexibility `json:"nodeTypeFlexibility,omitempty"`

	/* Immutable. Region where nodes using the node template will be created. */
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.NodeTemplate.region
	Region *string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for
	creation and acquisition. When unset, the value of `metadata.name`
	is used as the default. */
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The server binding policy for nodes using this template. Determines
	where the nodes should restart following a maintenance event. */
	// +kcc:proto:field=google.cloud.compute.v1.NodeTemplate.server_binding
	ServerBinding *NodeTemplateServerBinding `json:"serverBinding,omitempty"`
}

// ComputeNodeTemplateStatus defines the config connector machine state of ComputeNodeTemplate
// +kcc:status:proto=google.cloud.compute.v1.NodeTemplate
type ComputeNodeTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +kubebuilder:validation:Format=""
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* A unique specifier for the ComputeNodeTemplate resource in GCP. */
	ExternalRef *string `json:"externalRef,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	// +kcc:proto:field=google.cloud.compute.v1.NodeTemplate.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NodeTemplate.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenodetemplate;gcpcomputenodetemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeNodeTemplate is the Schema for the ComputeNodeTemplate API
// +k8s:openapi-gen=true
type ComputeNodeTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeNodeTemplateSpec   `json:"spec,omitempty"`
	Status ComputeNodeTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeNodeTemplateList contains a list of ComputeNodeTemplate
type ComputeNodeTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNodeTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNodeTemplate{}, &ComputeNodeTemplateList{})
}
