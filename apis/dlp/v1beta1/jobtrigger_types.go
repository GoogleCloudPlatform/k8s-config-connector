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
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DLPJobTriggerGVK = GroupVersion.WithKind("DLPJobTrigger")

// DLPJobTriggerSpec defines the desired state of DLPJobTrigger
// +kcc:spec:proto=google.privacy.dlp.v2.JobTrigger
type DLPJobTriggerSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The DLPJobTrigger name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Display name (max 100 chars)
	// +kcc:proto:field=google.privacy.dlp.v2.JobTrigger.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// User provided description (max 256 chars)
	// +kcc:proto:field=google.privacy.dlp.v2.JobTrigger.description
	Description *string `json:"description,omitempty"`

	// +required
	// For inspect jobs, a snapshot of the configuration.
	// +kcc:proto:field=google.privacy.dlp.v2.JobTrigger.inspect_job
	InspectJob *InspectJobConfig `json:"inspectJob,omitempty"`

	// +required
	// A list of triggers which will be OR'ed together. Only one in the list
	//  needs to trigger for a job to be started. The list may contain only
	//  a single Schedule trigger and must have at least one object.
	// +kcc:proto:field=google.privacy.dlp.v2.JobTrigger.triggers
	Triggers []JobTrigger_Trigger `json:"triggers,omitempty"`

	// +required
	// Required. A status for this trigger.
	// +kcc:proto:field=google.privacy.dlp.v2.JobTrigger.status
	Status *string `json:"status,omitempty"`
}

// +kcc:proto=google.privacy.dlp.v2.InspectJobConfig
type InspectJobConfig struct {
	// +required
	// The data to scan.
	// +kcc:proto:field=google.privacy.dlp.v2.InspectJobConfig.storage_config
	StorageConfig *StorageConfig `json:"storageConfig,omitempty"`

	// How and what to scan for.
	// +kcc:proto:field=google.privacy.dlp.v2.InspectJobConfig.inspect_config
	InspectConfig *InspectConfig `json:"inspectConfig,omitempty"`

	// If provided, will be used as the default for all values in InspectConfig.
	//  `inspect_config` will be merged into the values persisted as part of the
	//  template.
	// +kcc:proto:field=google.privacy.dlp.v2.InspectJobConfig.inspect_template_name
	InspectTemplateRef *DLPInspectTemplateRef `json:"inspectTemplateRef,omitempty"`

	// Actions to execute at the completion of the job.
	// +kcc:proto:field=google.privacy.dlp.v2.InspectJobConfig.actions
	Actions []Action `json:"actions,omitempty"`
}

type DLPInspectTemplateRef struct {
	/* The self-link of the DLPInspectTemplate resource. */
	External string `json:"external,omitempty"`
	/* The name of the DLPInspectTemplate resource. */
	Name string `json:"name,omitempty"`
	/* The namespace of the DLPInspectTemplate resource. */
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.privacy.dlp.v2.DatastoreOptions
type DatastoreOptions struct {
	// A partition ID identifies a grouping of entities. The grouping is always
	//  by project and namespace, however the namespace ID may be empty.
	// +kcc:proto:field=google.privacy.dlp.v2.DatastoreOptions.partition_id
	PartitionId *PartitionId `json:"partitionId,omitempty"`

	// The kind to process.
	// +kcc:proto:field=google.privacy.dlp.v2.DatastoreOptions.kind
	Kind *KindExpression `json:"kind,omitempty"`
}

// +kcc:proto=google.privacy.dlp.v2.PartitionId
type PartitionId struct {
	// The ID of the project to which the entities belong.
	// +kcc:proto:field=google.privacy.dlp.v2.PartitionId.project_id
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// If not empty, the ID of the namespace to which the entities belong.
	// +kcc:proto:field=google.privacy.dlp.v2.PartitionId.namespace_id
	NamespaceId *string `json:"namespaceId,omitempty"`
}

// +kcc:proto=google.privacy.dlp.v2.StorageConfig.TimespanConfig
type StorageConfig_TimespanConfig struct {
	// Exclude files, tables, or rows older than this value. If not set, no lower
	//  time limit is applied.
	// +kcc:proto:field=google.privacy.dlp.v2.StorageConfig.TimespanConfig.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Exclude files, tables, or rows newer than this value. If not set, no upper
	//  time limit is applied.
	// +kcc:proto:field=google.privacy.dlp.v2.StorageConfig.TimespanConfig.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Specification of the field containing the timestamp of scanned items.
	//  Used for data sources like Datastore and BigQuery. For BigQuery: If this
	//  value is not specified and the table was modified between the given
	//  start and end times, the entire table will be scanned. If this value is
	//  specified, then rows are filtered based on the given start and end
	//  times. Rows with a `NULL` value in the provided BigQuery column are
	//  skipped. Valid data types of the provided BigQuery column are:
	//  `INTEGER`, `DATE`, `TIMESTAMP`, and `DATETIME`. For Datastore: If this
	//  value is specified, then entities are filtered based on the given start
	//  and end times. If an entity does not contain the provided timestamp
	//  property or contains empty or invalid values, then it is included.
	//  Valid data types of the provided timestamp property are: `TIMESTAMP`.
	// +kcc:proto:field=google.privacy.dlp.v2.StorageConfig.TimespanConfig.timestamp_field
	TimestampField *FieldID `json:"timestampField,omitempty"`

	// When the job is started by a JobTrigger we will automatically figure out
	//  a valid start_time to avoid scanning files that have not been modified
	//  since the last time the JobTrigger executed. This will be based on the
	//  time of the execution of the last run of the JobTrigger.
	// +kcc:proto:field=google.privacy.dlp.v2.StorageConfig.TimespanConfig.enable_auto_population_of_timespan_config
	EnableAutoPopulationOfTimespanConfig *bool `json:"enableAutoPopulationOfTimespanConfig,omitempty"`
}

// +kcc:proto=google.privacy.dlp.v2.Action.JobNotificationEmails
type Action_JobNotificationEmails struct {
}

// +kcc:proto=google.privacy.dlp.v2.BigQueryOptions
type BigQueryOptions struct {
	// +required
	// Complete identifier of the table and project ID, dataset ID, and table ID
	//  must be set.
	// +kcc:proto:field=google.privacy.dlp.v2.BigQueryOptions.table_reference
	TableReference *BigQueryTable `json:"tableReference,omitempty"`

	// References to fields excluded from scanning. This allows you to skip
	//  inspection of entire columns which you know have no findings. When
	//  inspecting a table, we recommend that you inspect all columns.
	//  Otherwise, findings might be affected because hints from excluded
	//  columns will not be used.
	// +kcc:proto:field=google.privacy.dlp.v2.BigQueryOptions.excluded_fields
	ExcludedFields []FieldID `json:"excludedFields,omitempty"`

	// Table fields that may uniquely identify a row within the table. When
	//  `actions.saveFindings.outputConfig.table` is specified, the values of
	//  columns specified here are available in the output table under
	//  `location.content_locations.record_location.record_key.id_values`. Nested
	//  fields such as `person.birthdate.year` are allowed.
	// +kcc:proto:field=google.privacy.dlp.v2.BigQueryOptions.identifying_fields
	IdentifyingFields []FieldID `json:"identifyingFields,omitempty"`

	// Max number of rows to scan. If the table has more rows than this value,
	//  the rest of the rows are omitted. If not set, or if set to 0, all rows
	//  will be scanned. Only one of rows_limit and rows_limit_percent can be
	//  specified. Cannot be used in conjunction with TimespanConfig.
	// +kcc:proto:field=google.privacy.dlp.v2.BigQueryOptions.rows_limit
	RowsLimit *int64 `json:"rowsLimit,omitempty"`

	// Max percentage of rows to scan. The rest are omitted. The number of rows
	//  scanned is rounded down. Must be between 0 and 100, inclusively. Both 0
	//  and 100 means no limit. Defaults to 0. Only one of rows_limit and
	//  rows_limit_percent can be specified. Cannot be used in conjunction with
	//  TimespanConfig.
	//
	//  Caution: A [known
	//  issue](https://cloud.google.com/sensitive-data-protection/docs/known-issues#bq-sampling)
	//  is causing the `rowsLimitPercent` field to behave unexpectedly. We
	//  recommend using `rowsLimit` instead.
	// +kcc:proto:field=google.privacy.dlp.v2.BigQueryOptions.rows_limit_percent
	RowsLimitPercent *int32 `json:"rowsLimitPercent,omitempty"`

	// How to sample the data.
	// +kcc:proto:field=google.privacy.dlp.v2.BigQueryOptions.sample_method
	SampleMethod *string `json:"sampleMethod,omitempty"`

	// Limit scanning only to these fields. When inspecting a table, we
	//  recommend that you inspect all columns. Otherwise, findings might be
	//  affected because hints from excluded columns will not be used.
	// +kcc:proto:field=google.privacy.dlp.v2.BigQueryOptions.included_fields
	IncludedFields []FieldID `json:"includedFields,omitempty"`
}

// +kcc:proto=google.privacy.dlp.v2.CloudStorageRegexFileSet
type CloudStorageRegexFileSet struct {
	// +required
	// The name of a Cloud Storage bucket. Required.
	// +kcc:proto:field=google.privacy.dlp.v2.CloudStorageRegexFileSet.bucket_name
	BucketRef *storagev1beta1.StorageBucketRef `json:"bucketRef,omitempty"`

	// A list of regular expressions matching file paths to include. All files in
	//  the bucket that match at least one of these regular expressions will be
	//  included in the set of files, except for those that also match an item in
	//  `exclude_regex`. Leaving this field empty will match all files by default
	//  (this is equivalent to including `.*` in the list).
	//
	//  Regular expressions use RE2
	//  [syntax](https://github.com/google/re2/wiki/Syntax); a guide can be found
	//  under the google/re2 repository on GitHub.
	// +kcc:proto:field=google.privacy.dlp.v2.CloudStorageRegexFileSet.include_regex
	IncludeRegex []string `json:"includeRegex,omitempty"`

	// A list of regular expressions matching file paths to exclude. All files in
	//  the bucket that match at least one of these regular expressions will be
	//  excluded from the scan.
	//
	//  Regular expressions use RE2
	//  [syntax](https://github.com/google/re2/wiki/Syntax); a guide can be found
	//  under the google/re2 repository on GitHub.
	// +kcc:proto:field=google.privacy.dlp.v2.CloudStorageRegexFileSet.exclude_regex
	ExcludeRegex []string `json:"excludeRegex,omitempty"`
}

// DLPJobTriggerStatus defines the config connector machine state of DLPJobTrigger
// +kcc:status:proto=google.privacy.dlp.v2.JobTrigger
type DLPJobTriggerStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DLPJobTrigger resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Output only. The creation timestamp of a triggeredJob.
	// +kcc:proto:field=google.privacy.dlp.v2.JobTrigger.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. A stream of errors encountered when the trigger was activated.
	//  Repeated errors may result in the JobTrigger automatically being paused.
	//  Will return the last 100 errors. Whenever the JobTrigger is modified
	//  this list will be cleared.
	// +kcc:proto:field=google.privacy.dlp.v2.JobTrigger.errors
	Errors []Error `json:"errors,omitempty"`

	// Output only. The timestamp of the last time this trigger executed.
	// +kcc:proto:field=google.privacy.dlp.v2.JobTrigger.last_run_time
	LastRunTime *string `json:"lastRunTime,omitempty"`

	// Output only. The geographic location where this resource is stored.
	LocationId *string `json:"locationId,omitempty"`

	// Output only. The last update timestamp of a triggeredJob.
	// +kcc:proto:field=google.privacy.dlp.v2.JobTrigger.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdlpjobtrigger;gcpdlpjobtriggers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DLPJobTrigger is the Schema for the DLPJobTrigger API
// +k8s:openapi-gen=true
type DLPJobTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DLPJobTriggerSpec   `json:"spec,omitempty"`
	Status DLPJobTriggerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DLPJobTriggerList contains a list of DLPJobTrigger
type DLPJobTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DLPJobTrigger `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DLPJobTrigger{}, &DLPJobTriggerList{})
}
