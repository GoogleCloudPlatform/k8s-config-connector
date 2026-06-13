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

var MonitoringServiceLevelObjectiveGVK = GroupVersion.WithKind("MonitoringServiceLevelObjective")

// MonitoringServiceLevelObjectiveSpec defines the desired state of MonitoringServiceLevelObjective
// +kcc:spec:proto=google.monitoring.v3.ServiceLevelObjective
type MonitoringServiceLevelObjectiveSpec struct {
	// Immutable. The Project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The Service that this resource belongs to.
	// +required
	ServiceRef *MonitoringServiceRef `json:"serviceRef"`

	// Name used for UI elements listing this SLO.
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// The definition of good service, used to measure and calculate the quality
	//  of the `Service`'s performance with respect to a single aspect of service
	//  quality.
	// +optional
	ServiceLevelIndicator *ServiceLevelIndicator `json:"serviceLevelIndicator,omitempty"`

	// The fraction of service that must be good in order for this objective to be
	//  met. `0 < goal <= 0.9999`.
	// +required
	Goal *float64 `json:"goal"`

	// A rolling time period, semantically "in the past `<rolling_period>`".
	//  Must be an integer multiple of 1 day no larger than 30 days.
	// +optional
	RollingPeriod *string `json:"rollingPeriod,omitempty"`

	// A calendar period, semantically "since the start of the current
	//  `<calendar_period>`". At this time, only `DAY`, `WEEK`, `FORTNIGHT`, and
	//  `MONTH` are supported.
	// +optional
	CalendarPeriod *string `json:"calendarPeriod,omitempty"`

	// Labels which have been used to annotate the service-level objective. Label
	//  keys must start with a letter. Label keys and values may contain lowercase
	//  letters, numbers, underscores, and dashes. Label keys and values have a
	//  maximum length of 63 characters, and must be less than 128 bytes in size.
	//  Up to 64 label entries may be stored. For labels which do not have a
	//  semantic value, the empty string may be supplied for the label value.
	// +optional
	UserLabels map[string]string `json:"userLabels,omitempty"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// MonitoringServiceLevelObjectiveStatus defines the config connector machine state of MonitoringServiceLevelObjective
type MonitoringServiceLevelObjectiveStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MonitoringServiceLevelObjective resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MonitoringServiceLevelObjectiveObservedState `json:"observedState,omitempty"`
}

// MonitoringServiceLevelObjectiveObservedState is the state of the MonitoringServiceLevelObjective resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.monitoring.v3.ServiceLevelObjective
type MonitoringServiceLevelObjectiveObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringservicelevelobjective;gcpmonitoringservicelevelobjectives
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringServiceLevelObjective is the Schema for the MonitoringServiceLevelObjective API
// +k8s:openapi-gen=true
type MonitoringServiceLevelObjective struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MonitoringServiceLevelObjectiveSpec   `json:"spec,omitempty"`
	Status MonitoringServiceLevelObjectiveStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MonitoringServiceLevelObjectiveList contains a list of MonitoringServiceLevelObjective
type MonitoringServiceLevelObjectiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringServiceLevelObjective `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringServiceLevelObjective{}, &MonitoringServiceLevelObjectiveList{})
}
