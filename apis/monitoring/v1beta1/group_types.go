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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MonitoringGroupGVK = GroupVersion.WithKind("MonitoringGroup")

// MonitoringGroupSpec defines the desired state of MonitoringGroup
// +kcc:spec:proto=google.monitoring.v3.Group
type MonitoringGroupSpec struct {
	// A user-assigned name for this group, used only for display purposes.
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// The filter used to determine which monitored resources belong to this group.
	// +required
	Filter *string `json:"filter,omitempty"`

	// If true, the members of this group are considered to be a cluster. The system can perform additional analysis on groups that are clusters.
	// +optional
	IsCluster *bool `json:"isCluster,omitempty"`

	// +optional
	ParentRef *MonitoringGroupRef `json:"parentRef,omitempty"`

	// Immutable. The Project that this resource belongs to.
	// +optional
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// MonitoringGroupStatus defines the config connector machine state of MonitoringGroup
type MonitoringGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MonitoringGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MonitoringGroupObservedState `json:"observedState,omitempty"`
}

// MonitoringGroupObservedState is the state of the MonitoringGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.monitoring.v3.Group
type MonitoringGroupObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringgroup;gcpmonitoringgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringGroup is the Schema for the MonitoringGroup API
// +k8s:openapi-gen=true
type MonitoringGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MonitoringGroupSpec   `json:"spec,omitempty"`
	Status MonitoringGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MonitoringGroupList contains a list of MonitoringGroup
type MonitoringGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringGroup{}, &MonitoringGroupList{})
}
