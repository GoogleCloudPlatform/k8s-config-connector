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
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var StorageInsightsReportConfigGVK = GroupVersion.WithKind("StorageInsightsReportConfig")

// +kcc:proto=google.cloud.storageinsights.v1.CloudStorageFilters
type CloudStorageFilters struct {
	// Optional. Bucket for which the report will be generated.
	// +optional
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.CloudStorageDestinationOptions
type CloudStorageDestinationOptions struct {
	// Optional. Destination bucket.
	// +optional
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// Optional. Destination path is the path in the bucket where the report should be generated.
	// +optional
	DestinationPath *string `json:"destinationPath,omitempty"`
}

// +kcc:proto=google.cloud.storageinsights.v1.ParquetOptions
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:Schemaless
type ParquetOptions struct {
}

// StorageInsightsReportConfigSpec defines the desired state of StorageInsightsReportConfig
// +kcc:spec:proto=google.cloud.storageinsights.v1.ReportConfig
type StorageInsightsReportConfigSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location"`

	// The StorageInsightsReportConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The frequency of report generation.
	// +optional
	FrequencyOptions *FrequencyOptions `json:"frequencyOptions,omitempty"`

	// Optional. Options for CSV formatted reports.
	// +optional
	CsvOptions *CsvOptions `json:"csvOptions,omitempty"`

	// Optional. Options for Parquet formatted reports.
	// +optional
	ParquetOptions *ParquetOptions `json:"parquetOptions,omitempty"`

	// Optional. Report for exporting object metadata.
	// +optional
	ObjectMetadataReportOptions *ObjectMetadataReportOptions `json:"objectMetadataReportOptions,omitempty"`

	// Optional. Labels as key value pairs.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. User provided display name which can be empty and limited to 256 characters that is editable.
	// +optional
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
	// Output only. Create time stamp.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpstorageinsightsreportconfig;gcpstorageinsightsreportconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
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
