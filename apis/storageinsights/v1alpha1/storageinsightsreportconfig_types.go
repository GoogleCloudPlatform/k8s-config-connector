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

// API sources for StorageInsightsReportConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/storageinsights/v1/storageinsights.proto
// +kcc:source:service-docs=https://cloud.google.com/storage/docs/insights/storage-insights
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/storage/docs/reference/rest/v1/projects.locations.reportConfigs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var StorageInsightsReportConfigGVK = GroupVersion.WithKind("StorageInsightsReportConfig")

// StorageInsightsReportConfigSpec defines the desired state of StorageInsightsReportConfig
// +kcc:spec:proto=google.cloud.storageinsights.v1.ReportConfig
type StorageInsightsReportConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The StorageInsightsReportConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The frequency of report generation.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.frequency_options
	FrequencyOptions *FrequencyOptions `json:"frequencyOptions,omitempty"`

	// Options for CSV formatted reports.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.csv_options
	CsvOptions *CsvOptions `json:"csvOptions,omitempty"`

	// Options for Parquet formatted reports.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.parquet_options
	ParquetOptions *ParquetOptions `json:"parquetOptions,omitempty"`

	// Report for exporting object metadata.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.object_metadata_report_options
	ObjectMetadataReportOptions *ObjectMetadataReportOptions `json:"objectMetadataReportOptions,omitempty"`

	// Labels as key value pairs
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	// User provided display name which can be empty and limited to 256 characters
	//  that is editable.
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// StorageInsightsReportConfigStatus defines the config connector machine state of StorageInsightsReportConfig
type StorageInsightsReportConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the StorageInsightsReportConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *StorageInsightsReportConfigObservedState `json:"observedState,omitempty"`
}

// StorageInsightsReportConfigObservedState is the state of the StorageInsightsReportConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.storageinsights.v1.ReportConfig
type StorageInsightsReportConfigObservedState struct {
	// Output only. [Output only] Create time stamp
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp
	// +kcc:proto:field=google.cloud.storageinsights.v1.ReportConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpstorageinsightsreportconfig;gcpstorageinsightsreportconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// StorageInsightsReportConfig is the Schema for the StorageInsightsReportConfig API
// +k8s:openapi-gen=true
type StorageInsightsReportConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   StorageInsightsReportConfigSpec   `json:"spec,omitempty"`
	Status StorageInsightsReportConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// StorageInsightsReportConfigList contains a list of StorageInsightsReportConfig
type StorageInsightsReportConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageInsightsReportConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageInsightsReportConfig{}, &StorageInsightsReportConfigList{})
}
