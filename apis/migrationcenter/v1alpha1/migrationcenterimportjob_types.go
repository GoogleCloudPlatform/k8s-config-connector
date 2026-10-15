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

// API sources for MigrationCenterImportJob, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/migrationcenter/v1/migrationcenter.proto
// +kcc:source:service-docs=https://cloud.google.com/migration-center
// +kcc:source:resource-docs=https://docs.cloud.google.com/migration-center/docs/reference/rest/v1/projects.locations.importJobs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MigrationCenterImportJobGVK = GroupVersion.WithKind("MigrationCenterImportJob")

// MigrationCenterImportJobSpec defines the desired state of MigrationCenterImportJob
// +kcc:spec:proto=google.cloud.migrationcenter.v1.ImportJob
// +kcc:required-from-proto
type MigrationCenterImportJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The MigrationCenterImportJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// User-friendly display name. Maximum length is 63 characters.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Reference to a source.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.asset_source
	// +required
	AssetSource *string `json:"assetSource,omitempty"`
}

// MigrationCenterImportJobStatus defines the config connector machine state of MigrationCenterImportJob
type MigrationCenterImportJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MigrationCenterImportJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MigrationCenterImportJobObservedState `json:"observedState,omitempty"`
}

// MigrationCenterImportJobObservedState is the state of the MigrationCenterImportJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.migrationcenter.v1.ImportJob
type MigrationCenterImportJobObservedState struct {
	// Output only. The timestamp when the import job was created.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the import job was last updated.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The timestamp when the import job was completed.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.complete_time
	CompleteTime *string `json:"completeTime,omitempty"`

	// Output only. The state of the import job.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.state
	State *string `json:"state,omitempty"`

	// Output only. The report with the validation results of the import job.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.validation_report
	ValidationReport *ValidationReport `json:"validationReport,omitempty"`

	// Output only. The report with the results of running the import job.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ImportJob.execution_report
	ExecutionReport *ExecutionReportObservedState `json:"executionReport,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmigrationcenterimportjob;gcpmigrationcenterimportjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MigrationCenterImportJob is the Schema for the MigrationCenterImportJob API
// +k8s:openapi-gen=true
type MigrationCenterImportJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MigrationCenterImportJobSpec   `json:"spec,omitempty"`
	Status MigrationCenterImportJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MigrationCenterImportJobList contains a list of MigrationCenterImportJob
type MigrationCenterImportJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MigrationCenterImportJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MigrationCenterImportJob{}, &MigrationCenterImportJobList{})
}
