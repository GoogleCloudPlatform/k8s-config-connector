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

var AppOptimizeReportGVK = GroupVersion.WithKind("AppOptimizeReport")

// AppOptimizeReportSpec defines the desired state of AppOptimizeReport
// +kcc:spec:proto=google.cloud.appoptimize.v1beta.Report
type AppOptimizeReportSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location"`

	// The AppOptimizeReport name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. A list of dimensions to include in the report. Supported values:
	//   * `project`
	//   * `application`
	//   * `service_or_workload`
	//   * `resource`
	//   * `resource_type`
	//   * `location`
	//   * `product_display_name`
	//   * `sku`
	//   * `month`
	//   * `day`
	//   * `hour`
	// +required
	Dimensions []string `json:"dimensions"`

	// Required. A list of metrics to include in the report. Supported values:
	//   * `cost`
	//   * `cpu_mean_utilization`
	//   * `cpu_usage_core_seconds`
	//   * `cpu_allocation_core_seconds`
	//   * `cpu_p95_utilization`
	//   * `memory_mean_utilization`
	//   * `memory_usage_byte_seconds`
	//   * `memory_allocation_byte_seconds`
	//   * `memory_p95_utilization`
	// +required
	Metrics []string `json:"metrics"`

	// Optional. The resource containers for which to fetch data. Default is the
	// project specified in the report's parent.
	// +optional
	Scopes []Scope `json:"scopes,omitempty"`

	// Optional. A Common Expression Language (CEL) expression used to filter the
	// data for the report.
	// +optional
	Filter *string `json:"filter,omitempty"`
}

// AppOptimizeReportStatus defines the config connector machine state of AppOptimizeReport
type AppOptimizeReportStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AppOptimizeReport resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AppOptimizeReportObservedState `json:"observedState,omitempty"`
}

// AppOptimizeReportObservedState is the state of the AppOptimizeReport resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.appoptimize.v1beta.Report
type AppOptimizeReportObservedState struct {
	// Output only. Timestamp in UTC of when this report expires. Once the
	// report expires, it will no longer be accessible and the report's
	// underlying data will be deleted.
	// +optional
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpappoptimizereport;gcpappoptimizereports
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AppOptimizeReport is the Schema for the AppOptimizeReport API
// +k8s:openapi-gen=true
type AppOptimizeReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AppOptimizeReportSpec   `json:"spec,omitempty"`
	Status AppOptimizeReportStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AppOptimizeReportList contains a list of AppOptimizeReport
type AppOptimizeReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppOptimizeReport `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppOptimizeReport{}, &AppOptimizeReportList{})
}
