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

var TelcoAutomationEdgeSlmGVK = GroupVersion.WithKind("TelcoAutomationEdgeSlm")

// TelcoAutomationEdgeSlmSpec defines the desired state of TelcoAutomationEdgeSlm
// +kcc:spec:proto=google.cloud.telcoautomation.v1.EdgeSlm
type TelcoAutomationEdgeSlmSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The TelcoAutomationEdgeSlm name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. Reference to the orchestration cluster on which templates for
	//  this resources will be applied.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.orchestration_cluster
	OrchestrationClusterRef *TelcoAutomationOrchestrationClusterRef `json:"orchestrationClusterRef,omitempty"`

	// Optional. Labels as key value pairs. The key and value should contain
	//  characters which are UTF-8 compliant and less than 50 characters.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Type of workload cluster for which an EdgeSLM resource is
	//  created.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.workload_cluster_type
	WorkloadClusterType *string `json:"workloadClusterType,omitempty"`
}

// TelcoAutomationEdgeSlmStatus defines the config connector machine state of TelcoAutomationEdgeSlm
type TelcoAutomationEdgeSlmStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TelcoAutomationEdgeSlm resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TelcoAutomationEdgeSlmObservedState `json:"observedState,omitempty"`
}

// TelcoAutomationEdgeSlmObservedState is the state of the TelcoAutomationEdgeSlm resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.telcoautomation.v1.EdgeSlm
type TelcoAutomationEdgeSlmObservedState struct {
	// Output only. Create time stamp.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Provides the active TNA version for this resource.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.tna_version
	TnaVersion *string `json:"tnaVersion,omitempty"`

	// Output only. State of the EdgeSlm resource.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.EdgeSlm.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptelcoautomationedgeslm;gcptelcoautomationedgeslms
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TelcoAutomationEdgeSlm is the Schema for the TelcoAutomationEdgeSlm API
// +k8s:openapi-gen=true
type TelcoAutomationEdgeSlm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TelcoAutomationEdgeSlmSpec   `json:"spec,omitempty"`
	Status TelcoAutomationEdgeSlmStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TelcoAutomationEdgeSlmList contains a list of TelcoAutomationEdgeSlm
type TelcoAutomationEdgeSlmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TelcoAutomationEdgeSlm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TelcoAutomationEdgeSlm{}, &TelcoAutomationEdgeSlmList{})
}
