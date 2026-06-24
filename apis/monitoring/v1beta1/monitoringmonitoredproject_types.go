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

var MonitoringMonitoredProjectGVK = GroupVersion.WithKind("MonitoringMonitoredProject")

// MonitoringMonitoredProjectSpec defines the desired state of MonitoringMonitoredProject
// +kcc:spec:proto=google.monitoring.metricsscope.v1.MonitoredProject
type MonitoringMonitoredProjectSpec struct {
	/* Immutable. Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER} */
	MetricsScope string `json:"metricsScope"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// MonitoringMonitoredProjectStatus defines the config connector machine state of MonitoringMonitoredProject
type MonitoringMonitoredProjectStatus struct {
	/* Conditions represent the latest available observations of the
	   MonitoringMonitoredProject's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* Output only. The time when this `MonitoredProject` was created. */
	// +optional
	// +kubebuilder:validation:Format=date-time
	CreateTime *string `json:"createTime,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringmonitoredproject;gcpmonitoringmonitoredprojects
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringMonitoredProject is the Schema for the monitoring API
// +k8s:openapi-gen=true
type MonitoringMonitoredProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MonitoringMonitoredProjectSpec   `json:"spec,omitempty"`
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
