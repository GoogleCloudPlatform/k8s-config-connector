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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MonitoringMonitoredProjectGVK = GroupVersion.WithKind("MonitoringMonitoredProject")

// MonitoringMonitoredProjectSpec defines the desired state of MonitoringMonitoredProject
// +kcc:proto=google.monitoring.metricsscope.v1.MonitoredProject
type MonitoringMonitoredProjectSpec struct {
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	// Immutable.
	// The MonitoringMonitoredProject name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. Required.
	// The resource name of the existing Metrics Scope that will monitor this project.
	// Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}
	// +required
	MetricsScope *string `json:"metricsScope,omitempty"`

	/* NOTYET
	// Immutable. The resource name of the `MonitoredProject`. On input, the resource name
	//  includes the scoping project ID and monitored project ID. On output, it
	//  contains the equivalent project numbers.
	//  Example:
	//  `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
	Name *string `json:"name,omitempty"`
	*/
}

// MonitoringMonitoredProjectStatus defines the config connector machine state of MonitoringMonitoredProject
// +kcc:proto=google.monitoring.metricsscope.v1.MonitoredProject
type MonitoringMonitoredProjectStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* NOTYET
	// A unique specifier for the MonitoringMonitoredProject resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
	*/

	/* NOTYET
	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MonitoringMonitoredProjectObservedState `json:"observedState,omitempty"`
	*/

	// Output only. The time when this `MonitoredProject` was created.
	// +kubebuilder:validation:Format=date-time
	CreateTime *string `json:"createTime,omitempty"`
}

// MonitoringMonitoredProjectSpec defines the desired state of MonitoringMonitoredProject
// +kcc:proto=google.monitoring.metricsscope.v1.MonitoredProject
type MonitoringMonitoredProjectObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringmonitoredproject;gcpmonitoringmonitoredprojects
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/dcl2crd=true";"cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringMonitoredProject is the Schema for the MonitoringMonitoredProject API
// +k8s:openapi-gen=true
type MonitoringMonitoredProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec MonitoringMonitoredProjectSpec `json:"spec,omitempty"`

	Status MonitoringMonitoredProjectStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MonitoringMonitoredProjectList contains a list of MonitoringMonitoredProject
type MonitoringMonitoredProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringMonitoredProject `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringMonitoredProject{}, &MonitoringMonitoredProjectList{})
}
