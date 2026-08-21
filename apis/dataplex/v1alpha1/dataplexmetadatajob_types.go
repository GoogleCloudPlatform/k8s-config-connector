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

var DataplexMetadataJobGVK = GroupVersion.WithKind("DataplexMetadataJob")

// DataplexMetadataJobSpec defines the desired state of DataplexMetadataJob
// +kcc:spec:proto=google.cloud.dataplex.v1.MetadataJob
type DataplexMetadataJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The DataplexMetadataJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User-defined labels.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Metadata job type.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.type
	Type *string `json:"type,omitempty"`

	// Import job specification.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.import_spec
	ImportSpec *MetadataJobImportJobSpec `json:"importSpec,omitempty"`

	// Export job specification.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.export_spec
	ExportSpec *MetadataJobExportJobSpec `json:"exportSpec,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec
type MetadataJobImportJobSpec struct {
	// Optional. The URI of a Cloud Storage bucket or folder (beginning with
	//  `gs://` and ending with `/`) that contains the metadata import files for
	//  this job.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.source_storage_uri
	SourceStorageURI *string `json:"sourceStorageURI,omitempty"`

	// Optional. The time when the process that created the metadata import
	//  files began.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.source_create_time
	SourceCreateTime *string `json:"sourceCreateTime,omitempty"`

	// Required. A boundary on the scope of impact that the metadata import job
	//  can have.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.scope
	Scope *MetadataJobImportJobSpecScope `json:"scope,omitempty"`

	// Required. The sync mode for entries.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.entry_sync_mode
	EntrySyncMode *string `json:"entrySyncMode,omitempty"`

	// Required. The sync mode for aspects.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.aspect_sync_mode
	AspectSyncMode *string `json:"aspectSyncMode,omitempty"`

	// Optional. The level of logs to write to Cloud Logging for this job.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.log_level
	LogLevel *string `json:"logLevel,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.ImportJobScope
type MetadataJobImportJobSpecScope struct {
	// Required. The entry group that is in scope for the import job.
	//  Must contain exactly one element. The entry group and the job
	//  must be in the same location.
	// +kubebuilder:validation:Required
	EntryGroupRefs []EntryGroupRef `json:"entryGroupRefs,omitempty"`

	// Required. The entry types that are in scope for the import job.
	//  The job modifies only the entries and aspects that belong to these
	//  entry types.
	// +kubebuilder:validation:Required
	EntryTypeRefs []EntryTypeRef `json:"entryTypeRefs,omitempty"`

	// Optional. The aspect types that are in scope for the import job.
	//  The job modifies only the aspects that belong to these aspect types.
	AspectTypeRefs []AspectTypeRef `json:"aspectTypeRefs,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.MetadataJob.ExportJobSpec
type MetadataJobExportJobSpec struct {
	// Required. The scope of the export job.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ExportJobSpec.scope
	Scope *MetadataJobExportJobSpecScope `json:"scope,omitempty"`

	// Required. The root path of the Cloud Storage bucket to export the
	//  metadata to, in the format `gs://{bucket}/`.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ExportJobSpec.output_path
	OutputPath *string `json:"outputPath,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.MetadataJob.ExportJobSpec.ExportJobScope
type MetadataJobExportJobSpecScope struct {
	// Optional. Whether the metadata export job is an organization-level export job.
	OrganizationLevel *bool `json:"organizationLevel,omitempty"`

	// Optional. The projects whose metadata you want to export.
	ProjectRefs []refsv1beta1.ProjectRef `json:"projectRefs,omitempty"`

	// Optional. The entry groups whose metadata you want to export.
	EntryGroupRefs []EntryGroupRef `json:"entryGroupRefs,omitempty"`

	// Optional. The entry types that are in scope for the export job.
	EntryTypeRefs []EntryTypeRef `json:"entryTypeRefs,omitempty"`

	// Optional. The aspect types that are in scope for the export job.
	AspectTypeRefs []AspectTypeRef `json:"aspectTypeRefs,omitempty"`
}

// DataplexMetadataJobStatus defines the config connector machine state of DataplexMetadataJob
type DataplexMetadataJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexMetadataJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexMetadataJobObservedState `json:"observedState,omitempty"`
}

// DataplexMetadataJobObservedState is the state of the DataplexMetadataJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.MetadataJob
type DataplexMetadataJobObservedState struct {
	// Output only. The time when the job was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the job was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. System generated globally unique ID for the job.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Results from the import job.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.import_result
	ImportResult *MetadataJobImportJobResultObservedState `json:"importResult,omitempty"`

	// Output only. Results from the export job.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.export_result
	ExportResult *MetadataJobExportJobResultObservedState `json:"exportResult,omitempty"`

	// Output only. Status of the job.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.status
	Status *MetadataJobStatusObservedState `json:"status,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataplex.v1.MetadataJob.ImportJobResult
type MetadataJobImportJobResultObservedState struct {
	// Output only. The total number of entries that were deleted.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobResult.deleted_entries
	DeletedEntries *int64 `json:"deletedEntries,omitempty"`

	// Output only. The total number of entries that were updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobResult.updated_entries
	UpdatedEntries *int64 `json:"updatedEntries,omitempty"`

	// Output only. The total number of entries that were created.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobResult.created_entries
	CreatedEntries *int64 `json:"createdEntries,omitempty"`

	// Output only. The total number of entries that were unchanged.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobResult.unchanged_entries
	UnchangedEntries *int64 `json:"unchangedEntries,omitempty"`

	// Output only. The total number of entries that were recreated.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobResult.recreated_entries
	RecreatedEntries *int64 `json:"recreatedEntries,omitempty"`

	// Output only. The time when the status was updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobResult.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataplex.v1.MetadataJob.ExportJobResult
type MetadataJobExportJobResultObservedState struct {
	// Output only. The number of entries that were exported.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ExportJobResult.exported_entries
	ExportedEntries *int64 `json:"exportedEntries,omitempty"`

	// Output only. The error message if the metadata export job failed.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ExportJobResult.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.dataplex.v1.MetadataJob.Status
type MetadataJobStatusObservedState struct {
	// Output only. State of the metadata job.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.Status.state
	State *string `json:"state,omitempty"`

	// Output only. Message relating to the progression of a metadata job.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.Status.message
	Message *string `json:"message,omitempty"`

	// Output only. Progress tracking.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.Status.completion_percent
	CompletionPercent *int32 `json:"completionPercent,omitempty"`

	// Output only. The time when the status was updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.Status.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexmetadatajob;gcpdataplexmetadatajobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexMetadataJob is the Schema for the DataplexMetadataJob API
// +k8s:openapi-gen=true
type DataplexMetadataJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexMetadataJobSpec   `json:"spec,omitempty"`
	Status DataplexMetadataJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexMetadataJobList contains a list of DataplexMetadataJob
type DataplexMetadataJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexMetadataJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexMetadataJob{}, &DataplexMetadataJobList{})
}
