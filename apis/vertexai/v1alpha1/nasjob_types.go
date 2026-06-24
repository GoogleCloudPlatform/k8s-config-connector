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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAINasJobGVK = GroupVersion.WithKind("VertexAINasJob")

// VertexAINasJobSpec defines the desired state of VertexAINasJob
// +kcc:spec:proto=google.cloud.aiplatform.v1.NasJob
type VertexAINasJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The VertexAINasJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the NasJob.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The specification of a NasJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.nas_job_spec
	NasJobSpec *NasJobSpec `json:"nasJobSpec,omitempty"`

	// The labels with user-defined metadata to organize NasJobs.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key options for a NasJob.
	// If this is set, then all resources created by the NasJob
	// will be encrypted with the provided encryption key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Optional. Enable a separation of Custom model training
	// and restricted image training for tenant project.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.enable_restricted_image_training
	EnableRestrictedImageTraining *bool `json:"enableRestrictedImageTraining,omitempty"`
}

// VertexAINasJobStatus defines the config connector machine state of VertexAINasJob
type VertexAINasJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAINasJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAINasJobObservedState `json:"observedState,omitempty"`
}

// VertexAINasJobObservedState is the state of the VertexAINasJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.NasJob
type VertexAINasJobObservedState struct {
	// Output only. Resource name of the NasJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.name
	Name *string `json:"name,omitempty"`

	// Output only. Output of the NasJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.nas_job_output
	NasJobOutput *NasJobOutputObservedState `json:"nasJobOutput,omitempty"`

	// Output only. The detailed state of the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.state
	State *string `json:"state,omitempty"`

	// Output only. Time when the NasJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the NasJob for the first time entered the
	// `JOB_STATE_RUNNING` state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the NasJob entered any of the following states:
	// `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the NasJob was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Only populated when job's state is JOB_STATE_FAILED or
	// JOB_STATE_CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexainasjob;gcpvertexainasjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAINasJob is the Schema for the VertexAINasJob API
// +k8s:openapi-gen=true
type VertexAINasJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAINasJobSpec   `json:"spec,omitempty"`
	Status VertexAINasJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAINasJobList contains a list of VertexAINasJob
type VertexAINasJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAINasJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAINasJob{}, &VertexAINasJobList{})
}
