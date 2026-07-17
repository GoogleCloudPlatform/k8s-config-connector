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
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeInstanceGroupGVK = GroupVersion.WithKind("ComputeInstanceGroup")

// ComputeInstanceGroupSpec defines the desired state of ComputeInstanceGroup
// +kcc:spec:proto=google.cloud.compute.v1.InstanceGroup
type ComputeInstanceGroupSpec struct {
	// Immutable. An optional textual description of the instance group.
	// +optional
	Description *string `json:"description,omitempty"`

	// +optional
	Instances []InstanceRef `json:"instances,omitempty"`

	// The named port configuration.
	// +optional
	NamedPorts []InstanceGroupNamedPort `json:"namedPort,omitempty"`

	// +optional
	NetworkRef *computerefs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Immutable. Optional. The name of the resource. Used for
	// creation and acquisition. When unset, the value of `metadata.name`
	// is used as the default.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The zone that this instance group should be created in.
	Zone string `json:"zone"`
}

// +kcc:proto=google.cloud.compute.v1.NamedPort
type InstanceGroupNamedPort struct {
	// The name which the port will be mapped to.
	Name string `json:"name"`

	// The port number to map the name to.
	Port int32 `json:"port"`
}

// ComputeInstanceGroupStatus defines the config connector machine state of ComputeInstanceGroup
type ComputeInstanceGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The URI of the created resource.
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	// The number of instances in the group.
	// +optional
	Size *int64 `json:"size,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinstancegroup;gcpcomputeinstancegroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeInstanceGroup is the Schema for the ComputeInstanceGroup API
// +k8s:openapi-gen=true
type ComputeInstanceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeInstanceGroupSpec   `json:"spec,omitempty"`
	Status ComputeInstanceGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeInstanceGroupList contains a list of ComputeInstanceGroup
type ComputeInstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstanceGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstanceGroup{}, &ComputeInstanceGroupList{})
}
