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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var LoggingLogScopeGVK = GroupVersion.WithKind("LoggingLogScope")

// LoggingLogScopeSpec defines the desired state of LoggingLogScope
type LoggingLogScopeSpec struct {
	// The parent project that this resource belongs to.
	// +optional
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// The location of the resource. Currently only 'global' is supported.
	// +required
	Location string `json:"location"`

	// The LoggingLogScope name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Describes this log scope. The maximum length of the description is 8000 characters.
	// +optional
	Description *string `json:"description,omitempty"`

	// Required. Names of one or more parent resources:
	// - projects/[PROJECT_ID]
	// May alternatively be one or more views:
	// - projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/views/[VIEW_ID]
	// A log scope can include a maximum of 50 projects and a maximum of 100 resources in total.
	// +required
	ResourceNames []string `json:"resourceNames"`
}

// LoggingLogScopeObservedState defines the observed state of LoggingLogScope
type LoggingLogScopeObservedState struct {
	// Output only. The creation timestamp of the log scope.
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of the log scope.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// LoggingLogScopeStatus defines the config connector machine state of LoggingLogScope
type LoggingLogScopeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// ObservedState contains the observed state of the resource in GCP.
	// +optional
	ObservedState *LoggingLogScopeObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplogginglogscope;gcplogginglogscopes
// +kubebuilder:subresource:status
// +kubebuilder:object:root=true
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LoggingLogScope is the Schema for the logging API
// +k8s:openapi-gen=true
type LoggingLogScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec LoggingLogScopeSpec `json:"spec"`

	Status LoggingLogScopeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoggingLogScopeList contains a list of LoggingLogScope
type LoggingLogScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingLogScope `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoggingLogScope{}, &LoggingLogScopeList{})
}
