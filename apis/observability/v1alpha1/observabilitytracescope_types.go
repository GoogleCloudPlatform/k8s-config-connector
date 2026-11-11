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

// API sources for ObservabilityTraceScope, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/observability/v1/observability.proto
// +kcc:source:service-docs=https://cloud.google.com/stackdriver/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/stackdriver/docs/reference/observability/api/rest/v1/projects.locations.traceScopes

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ObservabilityTraceScopeGVK = GroupVersion.WithKind("ObservabilityTraceScope")

// ObservabilityTraceScopeSpec defines the desired state of ObservabilityTraceScope
// +kcc:spec:proto=google.cloud.observability.v1.TraceScope
// +kcc:required-from-proto
type ObservabilityTraceScopeSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ObservabilityTraceScope name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Names of the projects that are included in this trace scope.
	//
	//  *  `projects/[PROJECT_ID]`
	//
	//  A trace scope can include a maximum of 20 projects.
	// +kcc:proto:field=google.cloud.observability.v1.TraceScope.resource_names
	// +required
	ResourceNames []string `json:"resourceNames,omitempty"`

	// Optional. Describes this trace scope.
	//
	//  The maximum length of the description is 8000 characters.
	// +kcc:proto:field=google.cloud.observability.v1.TraceScope.description
	Description *string `json:"description,omitempty"`
}

// ObservabilityTraceScopeStatus defines the config connector machine state of ObservabilityTraceScope
type ObservabilityTraceScopeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ObservabilityTraceScope resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ObservabilityTraceScopeObservedState `json:"observedState,omitempty"`
}

// ObservabilityTraceScopeObservedState is the state of the ObservabilityTraceScope resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.observability.v1.TraceScope
type ObservabilityTraceScopeObservedState struct {
	// Output only. The creation timestamp of the trace scope.
	// +kcc:proto:field=google.cloud.observability.v1.TraceScope.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of the trace scope.
	// +kcc:proto:field=google.cloud.observability.v1.TraceScope.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpobservabilitytracescope;gcpobservabilitytracescopes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ObservabilityTraceScope is the Schema for the ObservabilityTraceScope API
// +k8s:openapi-gen=true
type ObservabilityTraceScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ObservabilityTraceScopeSpec   `json:"spec,omitempty"`
	Status ObservabilityTraceScopeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ObservabilityTraceScopeList contains a list of ObservabilityTraceScope
type ObservabilityTraceScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObservabilityTraceScope `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ObservabilityTraceScope{}, &ObservabilityTraceScopeList{})
}
