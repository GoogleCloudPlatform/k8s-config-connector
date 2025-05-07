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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EdgeContainerMachineGVK = GroupVersion.WithKind("EdgeContainerMachine")

// EdgeContainerMachineSpec defines the desired state of EdgeContainerMachine
// +kcc:proto=google.cloud.edgecontainer.v1.Machine
type EdgeContainerMachineSpec struct {
	// Labels associated with this resource.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The Google Distributed Cloud Edge zone of this machine.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.zone
	Zone *string `json:"zone,omitempty"`

	*Parent `json:",inline"`

	// The EdgeContainerMachine name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

type Parent struct {
	// Required. The location of the machine.
	Location string `json:"location,omitempty"`

	// Required. The host project of the machine.
	ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// EdgeContainerMachineStatus defines the config connector machine state of EdgeContainerMachine
type EdgeContainerMachineStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EdgeContainerMachine resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *EdgeContainerMachineObservedState `json:"observedState,omitempty"`
}

// EdgeContainerMachineObservedState is the state of the EdgeContainerMachine resource as most recently observed in GCP.
// +kcc:proto=google.cloud.edgecontainer.v1.Machine
type EdgeContainerMachineObservedState struct {
	// Output only. The time when the node pool was created.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the node pool was last updated.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The software version of the machine.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.version
	Version *string `json:"version,omitempty"`

	// Output only. Whether the machine is disabled. If disabled, the machine is
	//  unable to enter service.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// Canonical resource name of the node that this machine is responsible for
	//  hosting e.g.
	//  projects/{project}/locations/{location}/clusters/{cluster_id}/nodePools/{pool_id}/{node},
	//  Or empty if the machine is not assigned to assume the role of a node.
	//
	//  For control plane nodes hosted on edge machines, this will return
	//  the following format:
	//    "projects/{project}/locations/{location}/clusters/{cluster_id}/controlPlaneNodes/{node}".
	// +kcc:proto:field=google.cloud.edgecontainer.v1.Machine.hosted_node
	HostedNode *string `json:"hostedNode,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpedgecontainermachine;gcpedgecontainermachines
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EdgeContainerMachine is the Schema for the EdgeContainerMachine API
// +k8s:openapi-gen=true
type EdgeContainerMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   EdgeContainerMachineSpec   `json:"spec,omitempty"`
	Status EdgeContainerMachineStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EdgeContainerMachineList contains a list of EdgeContainerMachine
type EdgeContainerMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeContainerMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EdgeContainerMachine{}, &EdgeContainerMachineList{})
}
