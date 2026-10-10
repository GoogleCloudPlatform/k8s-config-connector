// Copyright 2025 Google LLC
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

// API sources for DataformReleaseConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/dataform/v1/dataform.proto
// +kcc:source:service-docs=https://cloud.google.com/dataform/docs
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataform/docs/reference/rest/v1/projects.locations.repositories.releaseConfigs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataformReleaseConfigGVK = GroupVersion.WithKind("DataformReleaseConfig")

// DataformReleaseConfigSpec defines the desired state of DataformReleaseConfig
// +kcc:spec:proto=google.cloud.dataform.v1.ReleaseConfig
// +kcc:required-from-proto
type DataformReleaseConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The DataformReleaseConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Git commit/tag/branch name at which the repository should be
	//  compiled. Must exist in the remote repository. Examples:
	//  - a commit SHA: `12ade345`
	//  - a tag: `tag1`
	//  - a branch name: `branch1`
	// +kcc:proto:field=google.cloud.dataform.v1.ReleaseConfig.git_commitish
	// +required
	GitCommitish *string `json:"gitCommitish,omitempty"`

	// Optional. If set, fields of `code_compilation_config` override the default
	//  compilation settings that are specified in dataform.json.
	// +kcc:proto:field=google.cloud.dataform.v1.ReleaseConfig.code_compilation_config
	CodeCompilationConfig *CodeCompilationConfig `json:"codeCompilationConfig,omitempty"`

	// Optional. Optional schedule (in cron format) for automatic creation of
	//  compilation results.
	// +kcc:proto:field=google.cloud.dataform.v1.ReleaseConfig.cron_schedule
	CronSchedule *string `json:"cronSchedule,omitempty"`

	// Optional. Specifies the time zone to be used when interpreting
	//  cron_schedule. Must be a time zone name from the time zone database
	//  (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left
	//  unspecified, the default is UTC.
	// +kcc:proto:field=google.cloud.dataform.v1.ReleaseConfig.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Optional. The name of the currently released compilation result for this
	//  release config. This value is updated when a compilation result is
	//  automatically created from this release config (using cron_schedule), or
	//  when this resource is updated by API call (perhaps to roll back to an
	//  earlier release). The compilation result must have been created using this
	//  release config. Must be in the format
	//  `projects/*/locations/*/repositories/*/compilationResults/*`.
	// +kcc:proto:field=google.cloud.dataform.v1.ReleaseConfig.release_compilation_result
	ReleaseCompilationResult *string `json:"releaseCompilationResult,omitempty"`

	// Optional. Disables automatic creation of compilation results.
	// +kcc:proto:field=google.cloud.dataform.v1.ReleaseConfig.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// DataformReleaseConfigStatus defines the config connector machine state of DataformReleaseConfig
type DataformReleaseConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataformReleaseConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataformReleaseConfigObservedState `json:"observedState,omitempty"`
}

// DataformReleaseConfigObservedState is the state of the DataformReleaseConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataform.v1.ReleaseConfig
type DataformReleaseConfigObservedState struct {
	// Output only. Records of the 10 most recent scheduled release attempts,
	//  ordered in descending order of `release_time`. Updated whenever automatic
	//  creation of a compilation result is triggered by cron_schedule.
	// +kcc:proto:field=google.cloud.dataform.v1.ReleaseConfig.recent_scheduled_release_records
	RecentScheduledReleaseRecords []ReleaseConfig_ScheduledReleaseRecordObservedState `json:"recentScheduledReleaseRecords,omitempty"`

	// Output only. All the metadata information that is used internally to serve
	//  the resource. For example: timestamps, flags, status fields, etc. The
	//  format of this field is a JSON string.
	// +kcc:proto:field=google.cloud.dataform.v1.ReleaseConfig.internal_metadata
	InternalMetadata *string `json:"internalMetadata,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataformreleaseconfig;gcpdataformreleaseconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataformReleaseConfig is the Schema for the DataformReleaseConfig API
// +k8s:openapi-gen=true
type DataformReleaseConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataformReleaseConfigSpec   `json:"spec,omitempty"`
	Status DataformReleaseConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataformReleaseConfigList contains a list of DataformReleaseConfig
type DataformReleaseConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataformReleaseConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataformReleaseConfig{}, &DataformReleaseConfigList{})
}
