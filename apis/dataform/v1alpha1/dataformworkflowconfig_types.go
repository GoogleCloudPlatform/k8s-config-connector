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

// API sources for DataformWorkflowConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/dataform/v1/dataform.proto
// +kcc:source:service-docs=https://cloud.google.com/dataform/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataform/reference/rest/v1/projects.locations.repositories.workflowConfigs

package v1alpha1

import (
	dataformv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataformWorkflowConfigGVK = GroupVersion.WithKind("DataformWorkflowConfig")

// DataformWorkflowConfigSpec defines the desired state of DataformWorkflowConfig
// +kcc:spec:proto=google.cloud.dataform.v1.WorkflowConfig
// +kcc:required-from-proto
type DataformWorkflowConfigSpec struct {
	// The DataformRepository this DataformWorkflowConfig is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *dataformv1beta1.DataformRepositoryRef `json:"parentRef"`

	// The DataformWorkflowConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The name of the release config whose release_compilation_result
	//  should be executed. Must be in the format
	//  `projects/*/locations/*/repositories/*/releaseConfigs/*`.
	// +kcc:proto:field=google.cloud.dataform.v1.WorkflowConfig.release_config
	// +required
	ReleaseConfigRef *DataformReleaseConfigRef `json:"releaseConfigRef"`

	// Optional. If left unset, a default InvocationConfig will be used.
	// +kcc:proto:field=google.cloud.dataform.v1.WorkflowConfig.invocation_config
	InvocationConfig *InvocationConfig `json:"invocationConfig,omitempty"`

	// Optional. Optional schedule (in cron format) for automatic execution of
	//  this workflow config.
	// +kcc:proto:field=google.cloud.dataform.v1.WorkflowConfig.cron_schedule
	CronSchedule *string `json:"cronSchedule,omitempty"`

	// Optional. Specifies the time zone to be used when interpreting
	//  cron_schedule. Must be a time zone name from the time zone database
	//  (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left
	//  unspecified, the default is UTC.
	// +kcc:proto:field=google.cloud.dataform.v1.WorkflowConfig.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Optional. Disables automatic creation of workflow invocations.
	// +kcc:proto:field=google.cloud.dataform.v1.WorkflowConfig.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.cloud.dataform.v1.InvocationConfig
type InvocationConfig struct {
	// Optional. The set of action identifiers to include.
	// +kcc:proto:field=google.cloud.dataform.v1.InvocationConfig.included_targets
	IncludedTargets []Target `json:"includedTargets,omitempty"`

	// Optional. The set of tags to include.
	// +kcc:proto:field=google.cloud.dataform.v1.InvocationConfig.included_tags
	IncludedTags []string `json:"includedTags,omitempty"`

	// Optional. When set to true, transitive dependencies of included actions
	//  will be executed.
	// +kcc:proto:field=google.cloud.dataform.v1.InvocationConfig.transitive_dependencies_included
	TransitiveDependenciesIncluded *bool `json:"transitiveDependenciesIncluded,omitempty"`

	// Optional. When set to true, transitive dependents of included actions will
	//  be executed.
	// +kcc:proto:field=google.cloud.dataform.v1.InvocationConfig.transitive_dependents_included
	TransitiveDependentsIncluded *bool `json:"transitiveDependentsIncluded,omitempty"`

	// Optional. When set to true, any incremental tables will be fully refreshed.
	// +kcc:proto:field=google.cloud.dataform.v1.InvocationConfig.fully_refresh_incremental_tables_enabled
	FullyRefreshIncrementalTablesEnabled *bool `json:"fullyRefreshIncrementalTablesEnabled,omitempty"`

	// Optional. The service account to run workflow invocations under.
	// +kcc:proto:field=google.cloud.dataform.v1.InvocationConfig.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. Specifies the priority for query execution in BigQuery.
	//  More information can be found at
	//  https://cloud.google.com/bigquery/docs/running-queries#queries.
	// +kcc:proto:field=google.cloud.dataform.v1.InvocationConfig.query_priority
	QueryPriority *string `json:"queryPriority,omitempty"`
}

// DataformWorkflowConfigStatus defines the config connector machine state of DataformWorkflowConfig
type DataformWorkflowConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataformWorkflowConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataformWorkflowConfigObservedState `json:"observedState,omitempty"`
}

// DataformWorkflowConfigObservedState is the state of the DataformWorkflowConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataform.v1.WorkflowConfig
type DataformWorkflowConfigObservedState struct {
	// Output only. Records of the 10 most recent scheduled execution attempts,
	//  ordered in descending order of `execution_time`. Updated whenever automatic
	//  creation of a workflow invocation is triggered by cron_schedule.
	// +kcc:proto:field=google.cloud.dataform.v1.WorkflowConfig.recent_scheduled_execution_records
	RecentScheduledExecutionRecords []WorkflowConfig_ScheduledExecutionRecordObservedState `json:"recentScheduledExecutionRecords,omitempty"`

	// Output only. The timestamp of when the WorkflowConfig was created.
	// +kcc:proto:field=google.cloud.dataform.v1.WorkflowConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp of when the WorkflowConfig was last updated.
	// +kcc:proto:field=google.cloud.dataform.v1.WorkflowConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. All the metadata information that is used internally to serve
	//  the resource. For example: timestamps, flags, status fields, etc. The
	//  format of this field is a JSON string.
	// +kcc:proto:field=google.cloud.dataform.v1.WorkflowConfig.internal_metadata
	InternalMetadata *string `json:"internalMetadata,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataformworkflowconfig;gcpdataformworkflowconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataformWorkflowConfig is the Schema for the DataformWorkflowConfig API
// +k8s:openapi-gen=true
type DataformWorkflowConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataformWorkflowConfigSpec   `json:"spec,omitempty"`
	Status DataformWorkflowConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataformWorkflowConfigList contains a list of DataformWorkflowConfig
type DataformWorkflowConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataformWorkflowConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataformWorkflowConfig{}, &DataformWorkflowConfigList{})
}
