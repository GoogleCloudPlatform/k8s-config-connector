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

var MonitoringMetricDescriptorGVK = GroupVersion.WithKind("MonitoringMetricDescriptor")

// MonitoringMetricDescriptorSpec defines the desired state of MonitoringMetricDescriptor
// +kcc:spec:proto=google.api.MetricDescriptor
type MonitoringMetricDescriptorSpec struct {
	// Immutable. The Project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The metric type, including its DNS name prefix. The type is not URL-encoded.
	// All user-defined metric types have the DNS name `custom.googleapis.com` or `external.googleapis.com`.
	// Metric types should use a natural hierarchical grouping.
	// +required
	Type *string `json:"type,omitempty"`

	// Immutable. Whether the metric records instantaneous values, changes to a value, etc.
	// Some combinations of `metricKind` and `valueType` might not be supported.
	// Possible values: METRIC_KIND_UNSPECIFIED, GAUGE, DELTA, CUMULATIVE
	// +required
	MetricKind *string `json:"metricKind,omitempty"`

	// Immutable. Whether the measurement is an integer, a floating-point number, etc.
	// Some combinations of `metricKind` and `valueType` might not be supported.
	// Possible values: STRING, BOOL, INT64, DOUBLE, DISTRIBUTION
	// +required
	ValueType *string `json:"valueType,omitempty"`

	// Immutable. A detailed description of the metric, which can be used in documentation.
	Description *string `json:"description,omitempty"`

	// Immutable. A concise name for the metric, which can be displayed in user interfaces.
	DisplayName *string `json:"displayName,omitempty"`

	// Immutable. The set of labels that can be used to describe a specific instance of this metric type.
	Labels []LabelDescriptor `json:"labels,omitempty"`

	// Immutable. The launch stage of the metric definition.
	// Possible values: LAUNCH_STAGE_UNSPECIFIED, UNIMPLEMENTED, PRELAUNCH, EARLY_ACCESS, ALPHA, BETA, GA, DEPRECATED
	LaunchStage *string `json:"launchStage,omitempty"`

	// Immutable. Metadata which can be used to guide usage of the metric.
	Metadata *MetricDescriptor_MetricDescriptorMetadata `json:"metadata,omitempty"`

	// Immutable. The units in which the metric value is reported. It is only applicable
	// if the `valueType` is `INT64`, `DOUBLE`, or `DISTRIBUTION`.
	Unit *string `json:"unit,omitempty"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`
}

// MonitoringMetricDescriptorStatus defines the config connector machine state of MonitoringMetricDescriptor
type MonitoringMetricDescriptorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MonitoringMetricDescriptor resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *MonitoringMetricDescriptorObservedState `json:"observedState,omitempty"`

	// Redundant fields for backwards compatibility with DCL controller
	MonitoredResourceTypes []string `json:"monitoredResourceTypes,omitempty"`
	SelfLink               *string  `json:"selfLink,omitempty"`
}

// MonitoringMetricDescriptorObservedState is the state of the MonitoringMetricDescriptor resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.api.MetricDescriptor
type MonitoringMetricDescriptorObservedState struct {
	// The resource name of the metric descriptor.
	Name *string `json:"name,omitempty"`

	// Read-only. If present, then a time series, which is identified partially by
	//  a metric type and a MonitoredResourceDescriptor, that is associated with
	//  this metric type can only be associated with one of the monitored
	//  resource types listed here.
	MonitoredResourceTypes []string `json:"monitoredResourceTypes,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringmetricdescriptor;gcpmonitoringmetricdescriptors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringMetricDescriptor is the Schema for the MonitoringMetricDescriptor API
// +k8s:openapi-gen=true
type MonitoringMetricDescriptor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MonitoringMetricDescriptorSpec   `json:"spec,omitempty"`
	Status MonitoringMetricDescriptorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MonitoringMetricDescriptorList contains a list of MonitoringMetricDescriptor
type MonitoringMetricDescriptorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringMetricDescriptor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringMetricDescriptor{}, &MonitoringMetricDescriptorList{})
}
