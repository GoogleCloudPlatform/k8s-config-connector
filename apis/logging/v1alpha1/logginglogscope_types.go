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

// API sources for LoggingLogScope, recorded by generate-types:
// +kcc:source:proto=https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/apis/google/logging/v2/logging_config.proto
// +kcc:source:service-docs=https://cloud.google.com/logging/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/logging/docs/reference/rest/v2/projects.locations.logScopes

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var LoggingLogScopeGVK = GroupVersion.WithKind("LoggingLogScope")

// LoggingLogScopeSpec defines the desired state of LoggingLogScope
// +kcc:spec:proto=google.logging.v2.LogScope
// +kcc:required-from-proto
type LoggingLogScopeSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The LoggingLogScope name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Describes this log scope.
	// +kcc:proto:field=google.logging.v2.LogScope.description
	Description *string `json:"description,omitempty"`

	// Required. Names of one or more parent resources:
	//    * `projects/[PROJECT_ID]`
	//  May alternatively be one or more views:
	//    * `projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/views/[VIEW_ID]`
	//  A log scope can include a maximum of 50 projects and a maximum of 100 resources in total.
	// +kcc:proto:field=google.logging.v2.LogScope.resource_names
	// +required
	ResourceNames []string `json:"resourceNames,omitempty"`
}

// LoggingLogScopeStatus defines the config connector machine state of LoggingLogScope
type LoggingLogScopeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the LoggingLogScope resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *LoggingLogScopeObservedState `json:"observedState,omitempty"`
}

// LoggingLogScopeObservedState is the state of the LoggingLogScope resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.logging.v2.LogScope
type LoggingLogScopeObservedState struct {
	// Output only. The creation timestamp of the log scope.
	// +kcc:proto:field=google.logging.v2.LogScope.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of the log scope.
	// +kcc:proto:field=google.logging.v2.LogScope.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplogginglogscope;gcplogginglogscopes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LoggingLogScope is the Schema for the LoggingLogScope API
// +k8s:openapi-gen=true
type LoggingLogScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LoggingLogScopeSpec   `json:"spec,omitempty"`
	Status LoggingLogScopeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LoggingLogScopeList contains a list of LoggingLogScope
type LoggingLogScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingLogScope `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoggingLogScope{}, &LoggingLogScopeList{})
}
