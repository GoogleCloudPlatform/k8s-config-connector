// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
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

var ComputeInstanceGroupNamedPortGVK = GroupVersion.WithKind("ComputeInstanceGroupNamedPort")

// ComputeInstanceGroupNamedPortSpec defines the desired state of ComputeInstanceGroupNamedPort
// +kcc:spec:proto=google.cloud.compute.v1.NamedPort
type ComputeInstanceGroupNamedPortSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The zone of the instance group.
	// +required
	Zone string `json:"zone"`

	// Reference to the ComputeInstanceGroup resource.
	// +required
	GroupRef *ComputeInstanceGroupRef `json:"groupRef"`

	// Immutable. The port number, which can be a value between 1 and 65535.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.NamedPort.port
	Port *int32 `json:"port"`

	// The ComputeInstanceGroupNamedPort name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeInstanceGroupNamedPortStatus defines the config connector machine state of ComputeInstanceGroupNamedPort
type ComputeInstanceGroupNamedPortStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeInstanceGroupNamedPort resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeInstanceGroupNamedPortObservedState `json:"observedState,omitempty"`
}

// ComputeInstanceGroupNamedPortObservedState is the state of the ComputeInstanceGroupNamedPort resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.NamedPort
type ComputeInstanceGroupNamedPortObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinstancegroupnamedport;gcpcomputeinstancegroupnamedports
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeInstanceGroupNamedPort is the Schema for the ComputeInstanceGroupNamedPort API
// +k8s:openapi-gen=true
type ComputeInstanceGroupNamedPort struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeInstanceGroupNamedPortSpec   `json:"spec,omitempty"`
	Status ComputeInstanceGroupNamedPortStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeInstanceGroupNamedPortList contains a list of ComputeInstanceGroupNamedPort
type ComputeInstanceGroupNamedPortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstanceGroupNamedPort `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstanceGroupNamedPort{}, &ComputeInstanceGroupNamedPortList{})
}
