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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeInstanceGroupNamedPortGVK = GroupVersion.WithKind("ComputeInstanceGroupNamedPort")

type ComputeInstanceGroupRef struct {
	/* The `name` field of a `ComputeInstanceGroup` resource. */
	Name string `json:"name,omitempty"`

	/* The `namespace` field of a `ComputeInstanceGroup` resource. */
	Namespace string `json:"namespace,omitempty"`

	/* Allowed value: The `name` field of a `ComputeInstanceGroup` resource. */
	External string `json:"external,omitempty"`
}

// ComputeInstanceGroupNamedPortSpec defines the desired state of ComputeInstanceGroupNamedPort
// +kcc:spec:proto=google.cloud.compute.v1.NamedPort
type ComputeInstanceGroupNamedPortSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. The zone of the instance group.
	Zone string `json:"zone"`

	// The instance group that this named port belongs to.
	GroupRef *ComputeInstanceGroupRef `json:"groupRef"`

	// Immutable. The port number, which can be a value between 1 and 65535.
	Port int64 `json:"port"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ComputeInstanceGroupNamedPortStatus defines the config connector machine state of ComputeInstanceGroupNamedPort
type ComputeInstanceGroupNamedPortStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
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
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
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
