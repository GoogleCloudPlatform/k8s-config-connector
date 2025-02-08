// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.dataplex.v1.MetadataJob
type MetadataJob struct {

	// Optional. User-defined labels.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Metadata job type.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.type
	Type *string `json:"type,omitempty"`

	// Import job specification.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.import_spec
	ImportSpec *MetadataJob_ImportJobSpec `json:"importSpec,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.MetadataJob.ImportJobResult
type MetadataJob_ImportJobResult struct {
}

// +kcc:proto=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec
type MetadataJob_ImportJobSpec struct {
	// Optional. The URI of a Cloud Storage bucket or folder (beginning with
	//  `gs://` and ending with `/`) that contains the metadata import files for
	//  this job.
	//
	//  A metadata import file defines the values to set for each of the entries
	//  and aspects in a metadata job. For more information about how to create a
	//  metadata import file and the file requirements, see [Metadata import
	//  file](https://cloud.google.com/dataplex/docs/import-metadata#metadata-import-file).
	//
	//  You can provide multiple metadata import files in the same metadata job.
	//  The bucket or folder must contain at least one metadata import file, in
	//  JSON Lines format (either `.json` or `.jsonl` file extension).
	//
	//  In `FULL` entry sync mode, don't save the metadata import file in a
	//  folder named `SOURCE_STORAGE_URI/deletions/`.
	//
	//  **Caution**: If the metadata import file contains no data, all entries
	//  and aspects that belong to the job's scope are deleted.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.source_storage_uri
	SourceStorageURI *string `json:"sourceStorageURI,omitempty"`

	// Optional. The time when the process that created the metadata import
	//  files began.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.source_create_time
	SourceCreateTime *string `json:"sourceCreateTime,omitempty"`

	// Required. A boundary on the scope of impact that the metadata import job
	//  can have.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.scope
	Scope *MetadataJob_ImportJobSpec_ImportJobScope `json:"scope,omitempty"`

	// Required. The sync mode for entries.
	//  Only `FULL` mode is supported for entries. All entries in the job's scope
	//  are modified. If an entry exists in Dataplex but isn't included in the
	//  metadata import file, the entry is deleted when you run the metadata job.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.entry_sync_mode
	EntrySyncMode *string `json:"entrySyncMode,omitempty"`

	// Required. The sync mode for aspects.
	//  Only `INCREMENTAL` mode is supported for aspects. An aspect is modified
	//  only if the metadata import file includes a reference to the aspect in
	//  the `update_mask` field and the `aspect_keys` field.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.aspect_sync_mode
	AspectSyncMode *string `json:"aspectSyncMode,omitempty"`

	// Optional. The level of logs to write to Cloud Logging for this job.
	//
	//  Debug-level logs provide highly-detailed information for
	//  troubleshooting, but their increased verbosity could incur [additional
	//  costs](https://cloud.google.com/stackdriver/pricing) that might not be
	//  merited for all jobs.
	//
	//  If unspecified, defaults to `INFO`.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.log_level
	LogLevel *string `json:"logLevel,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.ImportJobScope
type MetadataJob_ImportJobSpec_ImportJobScope struct {
	// Required. The entry group that is in scope for the import job,
	//  specified as a relative resource name in the format
	//  `projects/{project_number_or_id}/locations/{location_id}/entryGroups/{entry_group_id}`.
	//  Only entries that belong to the specified entry group are affected by
	//  the job.
	//
	//  Must contain exactly one element. The entry group and the job
	//  must be in the same location.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.ImportJobScope.entry_groups
	EntryGroups []string `json:"entryGroups,omitempty"`

	// Required. The entry types that are in scope for the import job,
	//  specified as relative resource names in the format
	//  `projects/{project_number_or_id}/locations/{location_id}/entryTypes/{entry_type_id}`.
	//  The job modifies only the entries that belong to these entry types.
	//
	//  If the metadata import file attempts to modify an entry whose type
	//  isn't included in this list, the import job is halted before modifying
	//  any entries or aspects.
	//
	//  The location of an entry type must either match the location of the
	//  job, or the entry type must be global.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.ImportJobScope.entry_types
	EntryTypes []string `json:"entryTypes,omitempty"`

	// Optional. The aspect types that are in scope for the import job,
	//  specified as relative resource names in the format
	//  `projects/{project_number_or_id}/locations/{location_id}/aspectTypes/{aspect_type_id}`.
	//  The job modifies only the aspects that belong to these aspect types.
	//
	//  If the metadata import file attempts to modify an aspect whose type
	//  isn't included in this list, the import job is halted before modifying
	//  any entries or aspects.
	//
	//  The location of an aspect type must either match the location of the
	//  job, or the aspect type must be global.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.ImportJobSpec.ImportJobScope.aspect_types
	AspectTypes []string `json:"aspectTypes,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.MetadataJob.Status
type MetadataJob_Status struct {
}

// +kcc:proto=google.cloud.dataplex.v1.MetadataJob
type MetadataJobObservedState struct {
	// Output only. Identifier. The name of the resource that the configuration is
	//  applied to, in the format
	//  `projects/{project_number}/locations/{location_id}/metadataJobs/{metadata_job_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.name
	Name *string `json:"name,omitempty"`

	// Output only. A system-generated, globally unique ID for the metadata job.
	//  If the metadata job is deleted and then re-created with the same name, this
	//  ID is different.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the metadata job was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the metadata job was updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Import job result.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.import_result
	ImportResult *MetadataJob_ImportJobResult `json:"importResult,omitempty"`

	// Output only. Metadata job status.
	// +kcc:proto:field=google.cloud.dataplex.v1.MetadataJob.status
	Status *MetadataJob_Status `json:"status,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.MetadataJob.ImportJobResult
type MetadataJob_ImportJobResultObservedState struct {
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

// +kcc:proto=google.cloud.dataplex.v1.MetadataJob.Status
type MetadataJob_StatusObservedState struct {
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
