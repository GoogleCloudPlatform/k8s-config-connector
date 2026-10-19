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

// API sources for AIPlatformTensorboardTimeSeries, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/aiplatform/v1/tensorboard_time_series.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.tensorboards.experiments.runs.timeSeries

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformTensorboardTimeSeriesGVK = GroupVersion.WithKind("AIPlatformTensorboardTimeSeries")

// AIPlatformTensorboardTimeSeriesSpec defines the desired state of AIPlatformTensorboardTimeSeries
// +kcc:spec:proto=google.cloud.aiplatform.v1.TensorboardTimeSeries
// +kcc:required-from-proto
type AIPlatformTensorboardTimeSeriesSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AIPlatformTensorboardTimeSeries name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. User provided name of this TensorboardTimeSeries.
	//  This value should be unique among all TensorboardTimeSeries resources
	//  belonging to the same TensorboardRun resource (parent resource).
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardTimeSeries.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Description of this TensorboardTimeSeries.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardTimeSeries.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. Type of TensorboardTimeSeries value.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardTimeSeries.value_type
	// +required
	ValueType *string `json:"valueType,omitempty"`

	// Used to perform a consistent read-modify-write updates. If not set, a blind
	//  "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardTimeSeries.etag
	Etag *string `json:"etag,omitempty"`

	// Immutable. Name of the plugin this time series pertain to. Such as Scalar,
	//  Tensor, Blob
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardTimeSeries.plugin_name
	PluginName *string `json:"pluginName,omitempty"`

	// Data of the current plugin, with the size limited to 65KB.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardTimeSeries.plugin_data
	PluginData []byte `json:"pluginData,omitempty"`
}

// AIPlatformTensorboardTimeSeriesStatus defines the config connector machine state of AIPlatformTensorboardTimeSeries
type AIPlatformTensorboardTimeSeriesStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformTensorboardTimeSeries resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformTensorboardTimeSeriesObservedState `json:"observedState,omitempty"`
}

// AIPlatformTensorboardTimeSeriesObservedState is the state of the AIPlatformTensorboardTimeSeries resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.TensorboardTimeSeries
type AIPlatformTensorboardTimeSeriesObservedState struct {
	// Output only. Timestamp when this TensorboardTimeSeries was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardTimeSeries.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this TensorboardTimeSeries was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardTimeSeries.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Scalar, Tensor, or Blob metadata for this
	//  TensorboardTimeSeries.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TensorboardTimeSeries.metadata
	Metadata *TensorboardTimeSeries_MetadataObservedState `json:"metadata,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformtensorboardtimeseries;gcpaiplatformtensorboardtimeseries
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformTensorboardTimeSeries is the Schema for the AIPlatformTensorboardTimeSeries API
// +k8s:openapi-gen=true
type AIPlatformTensorboardTimeSeries struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformTensorboardTimeSeriesSpec   `json:"spec,omitempty"`
	Status AIPlatformTensorboardTimeSeriesStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformTensorboardTimeSeriesList contains a list of AIPlatformTensorboardTimeSeries
type AIPlatformTensorboardTimeSeriesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformTensorboardTimeSeries `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformTensorboardTimeSeries{}, &AIPlatformTensorboardTimeSeriesList{})
}
