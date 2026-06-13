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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MonitoringServiceLevelObjectiveGVK = GroupVersion.WithKind("MonitoringServiceLevelObjective")

type MonitoringServiceRef struct {
	/* The name of a `MonitoringService` resource. */
	External string `json:"external,omitempty"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`
	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.monitoring.v3.BasicSli.AvailabilityCriteria
// +kubebuilder:validation:XPreserveUnknownFields
type BasicSli_AvailabilityCriteria struct {
}

// +kubebuilder:validation:XPreserveUnknownFields
type BasicSli_OperationAvailability struct {
}

// +kcc:proto=google.monitoring.v3.BasicSli.LatencyCriteria
type BasicSli_LatencyCriteria struct {
	// +optional
	Experience *string `json:"experience,omitempty"`
	// +optional
	Threshold *string `json:"threshold,omitempty"`
}

type BasicSli_OperationLatency struct {
	// +optional
	Experience *string `json:"experience,omitempty"`
	// +optional
	Threshold *string `json:"threshold,omitempty"`
}

// +kcc:proto=google.monitoring.v3.BasicSli
type BasicSli struct {
	// +optional
	Method []string `json:"method,omitempty"`
	// +optional
	Location []string `json:"location,omitempty"`
	// +optional
	Version []string `json:"version,omitempty"`
	// +optional
	Availability *BasicSli_AvailabilityCriteria `json:"availability,omitempty"`
	// +optional
	Latency *BasicSli_LatencyCriteria `json:"latency,omitempty"`
	// +optional
	OperationAvailability *BasicSli_OperationAvailability `json:"operationAvailability,omitempty"`
	// +optional
	OperationLatency *BasicSli_OperationLatency `json:"operationLatency,omitempty"`
}

// +kcc:proto=google.monitoring.v3.Range
type Range struct {
	// Range minimum.
	// +optional
	// +kubebuilder:validation:Format=double
	Min *float64 `json:"min,omitempty"`

	// Range maximum.
	// +optional
	// +kubebuilder:validation:Format=double
	Max *float64 `json:"max,omitempty"`
}

// +kcc:proto=google.monitoring.v3.WindowsBasedSli.PerformanceThreshold
type WindowsBasedSli_PerformanceThreshold struct {
	// +optional
	Performance *RequestBasedSli `json:"performance,omitempty"`

	// +optional
	BasicSliPerformance *BasicSli `json:"basicSliPerformance,omitempty"`

	// +optional
	// +kubebuilder:validation:Format=double
	Threshold *float64 `json:"threshold,omitempty"`
}

// MonitoringServiceLevelObjectiveSpec defines the desired state of MonitoringServiceLevelObjective
// +kcc:spec:proto=google.monitoring.v3.ServiceLevelObjective
type MonitoringServiceLevelObjectiveSpec struct {
	// The project that this resource belongs to.
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// Immutable.
	ServiceRef MonitoringServiceRef `json:"serviceRef"`

	/* A calendar period, semantically "since the start of the current ``". At this time, only `DAY`, `WEEK`, `FORTNIGHT`, and `MONTH` are supported. Possible values: CALENDAR_PERIOD_UNSPECIFIED, DAY, WEEK, FORTNIGHT, MONTH, QUARTER, HALF, YEAR */
	// +optional
	CalendarPeriod *string `json:"calendarPeriod,omitempty"`

	/* Name used for UI elements listing this SLO. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* The fraction of service that must be good in order for this objective to be met. `0 < goal <= 0.999`. */
	// +kubebuilder:validation:Format=double
	Goal float64 `json:"goal"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* A rolling time period, semantically "in the past ``". Must be an integer multiple of 1 day no larger than 30 days. */
	// +optional
	RollingPeriod *string `json:"rollingPeriod,omitempty"`

	/* The definition of good service, used to measure and calculate the quality of the `Service`'s performance with respect to a single aspect of service quality. */
	// +optional
	ServiceLevelIndicator *ServiceLevelIndicator `json:"serviceLevelIndicator,omitempty"`
}

// MonitoringServiceLevelObjectiveStatus defines the config connector machine state of MonitoringServiceLevelObjective
type MonitoringServiceLevelObjectiveStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* Time stamp of the `Create` or most recent `Update` command on this `Slo`. */
	// +optional
	// +kubebuilder:validation:Format=date-time
	CreateTime *string `json:"createTime,omitempty"`

	/* Time stamp of the `Update` or `Delete` command that made this no longer a current `Slo`. This field is not populated in `ServiceLevelObjective`s returned from calls to `GetServiceLevelObjective` and `ListServiceLevelObjectives`, because it is always empty in the current version. It is populated in `ServiceLevelObjective`s representing previous versions in the output of `ListServiceLevelObjectiveVersions`. Because all old configuration versions are stored, `Update` operations mark the obsoleted version as deleted. */
	// +optional
	// +kubebuilder:validation:Format=date-time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Output only. If set, this SLO is managed at the [Service Management](https://cloud.google.com/service-management/overview) level. Therefore the service yaml file is the source of truth for this SLO, and API `Update` and `Delete` operations are forbidden. */
	// +optional
	ServiceManagementOwned *bool `json:"serviceManagementOwned,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringservicelevelobjective;gcpmonitoringservicelevelobjectives
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
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
