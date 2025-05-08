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

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var WorkflowsWorkflowGVK = GroupVersion.WithKind("WorkflowsWorkflow")

type WorkflowsWorkflow_StateError struct {
	// Provides specifics about the error.
	Details *string `json:"details,omitempty"`

	// The type of this state error.
	Type *string `json:"type,omitempty"`
}

type Parent struct {
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// WorkflowsWorkflowSpec defines the desired state of Workflow
// +kcc:proto=google.cloud.workflows.v1.Workflow
type WorkflowsWorkflowSpec struct {
	Parent `json:",inline"`

	// Description of the workflow provided by the user.
	// Must be at most 1000 unicode characters long.
	Description *string `json:"description,omitempty"`

	// Labels associated with this workflow.
	// Labels can contain at most 64 entries. Keys and values can be no longer
	// than 63 characters and can only contain lowercase letters, numeric
	// characters, underscores, and dashes. Label keys must start with a letter.
	// International characters are allowed.
	Labels map[string]string `json:"labels,omitempty"`

	// The service account associated with the latest workflow version.
	// This service account represents the identity of the workflow and determines
	// what permissions the workflow has.
	// If not provided, workflow will use the project's default service account.
	// Modifying this field for an existing workflow results in a new workflow
	// revision.
	ServiceAccountRef *refs.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Required. Workflow code to be executed. The size limit is 128KB.
	// +required
	SourceContents *string `json:"sourceContents,omitempty"`

	// Optional. The resource name of a KMS crypto key used to encrypt or decrypt
	// the data associated with the workflow.
	// If not provided, data associated with the workflow will not be
	// CMEK-encrypted.
	// +optional
	KMSCryptoKeyRef *refs.KMSCryptoKeyRef `json:"kmsCryptoKeyRef,omitempty"`

	// Optional. Describes the level of platform logging to apply to calls and
	// call responses during executions of this workflow. If both the workflow and
	// the execution specify a logging level, the execution level takes
	// precedence.
	//+optional
	CallLogLevel *string `json:"callLogLevel,omitempty"`

	// Optional.User-defined environment variables associated with this workflow
	// revision. This map has a maximum length of 20. Each string can take up to
	// 40KiB. Keys cannot be empty strings and cannot start with “GOOGLE” or
	// “WORKFLOWS".
	// +optional
	UserEnvVars map[string]string `json:"userEnvVars,omitempty"`

	// Optional. Describes the execution history level to apply to this workflow.
	ExecutionHistoryLevel *string `json:"executionHistoryLevel,omitempty"`

	// Optional. Input only. Immutable. Tags associated with this workflow.
	Tags map[string]string `json:"tags,omitempty"`

	// The Workflow name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// WorkflowsWorkflowStatus defines the config connector machine state of Workflow
type WorkflowsWorkflowStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the Workflow resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *WorkflowsWorkflowObservedState `json:"observedState,omitempty"`
}

// WorkflowsWorkflowObservedState is the state of the Workflow resource as most recently observed in GCP.
// +kcc:proto=google.cloud.workflows.v1.Workflow
type WorkflowsWorkflowObservedState struct {
	// State of the workflow deployment.
	State *string `json:"state,omitempty"`

	// The revision of the workflow.
	// A new revision of a workflow is created as a result of updating the
	// following properties of a workflow:
	// - service_account
	// - source_content
	// The format is "000001-a4d", where the first six characters define
	// the zero-padded revision ordinal number. They are followed by a hyphen and
	// three hexadecimal random characters.
	RevisionId *string `json:"revisionId,omitempty"`

	// The timestamp for when the workflow was created.
	CreateTime *string `json:"createTime,omitempty"`

	// The timestamp for when the workflow was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	//  The timestamp for the latest revision of the workflow's
	// creation.
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Error regarding the state of the workflow. For example, this
	// field will have error details if the execution data is unavailable due to
	// revoked KMS key permissions.
	StateError *WorkflowsWorkflow_StateError `json:"stateError,omitempty"`

	// Output only. A list of all KMS crypto keys used to encrypt or decrypt the
	// data associated with the workflow.
	AllKmsKeys []refs.KMSCryptoKeyRef `json:"allKmsKeys,omitempty"`

	// Output only. A list of all KMS crypto key versions used to encrypt or
	// decrypt the data associated with the workflow.
	AllKmsKeysVersions []string `json:"allKmsKeysVersions,omitempty"`

	// Output only. The resource name of a KMS crypto key version used to encrypt
	// or decrypt the data associated with the workflow.
	CryptoKeyVersion *string `json:"cryptoKeyVersion,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpworkflowsworkflow;gcpworkflowsworkflows
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// Workflow is the Schema for the Workflow API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type WorkflowsWorkflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   WorkflowsWorkflowSpec   `json:"spec,omitempty"`
	Status WorkflowsWorkflowStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// WorkflowsWorkflowList contains a list of WorkflowsWorkflow
type WorkflowsWorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkflowsWorkflow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkflowsWorkflow{}, &WorkflowsWorkflowList{})
}
