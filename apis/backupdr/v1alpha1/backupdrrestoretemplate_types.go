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

// API sources for BackupDRRestoreTemplate, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/backupdr/v1/restore_template.proto
// +kcc:source:service-docs=https://cloud.google.com/backup-disaster-recovery
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/backup-disaster-recovery/docs/reference/rest/v1/projects.locations.restoreTemplates

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BackupDRRestoreTemplateGVK = GroupVersion.WithKind("BackupDRRestoreTemplate")

// BackupDRRestoreTemplateSpec defines the desired state of BackupDRRestoreTemplate
// +kcc:spec:proto=google.cloud.backupdr.v1.RestoreTemplate
// +kcc:required-from-proto
type BackupDRRestoreTemplateSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The BackupDRRestoreTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. User-provided description of the restore template.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. The specific type of the workload being restored.
	//  The format is "{service_name}/{resource_kind}", for example,
	//  "compute.googleapis.com/Instance".
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.resource_type
	// +required
	ResourceType *string `json:"resourceType,omitempty"`

	// Optional. User-defined labels.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Backup selection configuration.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.backup_selection_config
	// +required
	BackupSelectionConfig *BackupSelectionConfig `json:"backupSelectionConfig,omitempty"`

	// Optional. Pre-restore configuration.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.pre_restore_config
	PreRestoreConfig *PreRestoreConfig `json:"preRestoreConfig,omitempty"`

	// Optional. Post-restore configuration.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.post_restore_config
	PostRestoreConfig *PostRestoreConfig `json:"postRestoreConfig,omitempty"`

	// Optional. Verification configuration.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.verification_config
	VerificationConfig *VerificationConfig `json:"verificationConfig,omitempty"`

	// Required. Defines target environments where the resource will be restored.
	//
	//  The target resource name in `restore_properties` supports the
	//  `{EXECUTION_ID}` token to dynamically generate unique resource names for
	//  each execution. During restore execution, every occurrence of
	//  `{EXECUTION_ID}` in the target resource name is replaced with the
	//  8-character `RestoreTemplateExecution` ID (for example,
	//  `"restored-resource-{EXECUTION_ID}"` or
	//  `"restored-{EXECUTION_ID}-resource"` resolves to
	//  `"restored-resource-5d437b06"` or `"restored-5d437b06-resource"`).
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.restore_properties
	// +required
	RestoreProperties apiextensionsv1.JSON `json:"restoreProperties,omitempty"`

	// Optional. Etag for concurrency control.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.etag
	Etag *string `json:"etag,omitempty"`
}

// BackupDRRestoreTemplateStatus defines the config connector machine state of BackupDRRestoreTemplate
type BackupDRRestoreTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BackupDRRestoreTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BackupDRRestoreTemplateObservedState `json:"observedState,omitempty"`
}

// BackupDRRestoreTemplateObservedState is the state of the BackupDRRestoreTemplate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.backupdr.v1.RestoreTemplate
type BackupDRRestoreTemplateObservedState struct {
	// Output only. The time when the restore template was created.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the restore template was last updated.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the restore template.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.state
	State *string `json:"state,omitempty"`

	// Output only. Summary of the last execution.
	// +kcc:proto:field=google.cloud.backupdr.v1.RestoreTemplate.execution_summary
	ExecutionSummary *ExecutionSummaryObservedState `json:"executionSummary,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbackupdrrestoretemplate;gcpbackupdrrestoretemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BackupDRRestoreTemplate is the Schema for the BackupDRRestoreTemplate API
// +k8s:openapi-gen=true
type BackupDRRestoreTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BackupDRRestoreTemplateSpec   `json:"spec,omitempty"`
	Status BackupDRRestoreTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BackupDRRestoreTemplateList contains a list of BackupDRRestoreTemplate
type BackupDRRestoreTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupDRRestoreTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupDRRestoreTemplate{}, &BackupDRRestoreTemplateList{})
}
