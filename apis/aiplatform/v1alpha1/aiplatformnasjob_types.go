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

// API sources for AIPlatformNasJob, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/aiplatform/v1/nas_job.proto
// +kcc:source:service-docs=https://cloud.google.com/vertex-ai/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.nasJobs

package v1alpha1

import (
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformNasJobGVK = GroupVersion.WithKind("AIPlatformNasJob")

// AIPlatformNasJobSpec defines the desired state of AIPlatformNasJob
// +kcc:spec:proto=google.cloud.aiplatform.v1.NasJob
// +kcc:required-from-proto
type AIPlatformNasJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AIPlatformNasJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name of the NasJob.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The specification of a NasJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.nas_job_spec
	// +required
	NasJobSpec *NasJobSpec `json:"nasJobSpec,omitempty"`

	// The labels with user-defined metadata to organize NasJobs.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key options for a NasJob.
	//  If this is set, then all resources created by the NasJob
	//  will be encrypted with the provided encryption key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// AIPlatformNasJobStatus defines the config connector machine state of AIPlatformNasJob
type AIPlatformNasJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformNasJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformNasJobObservedState `json:"observedState,omitempty"`
}

// AIPlatformNasJobObservedState is the state of the AIPlatformNasJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.NasJob
type AIPlatformNasJobObservedState struct {
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
	//  `JOB_STATE_RUNNING` state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the NasJob entered any of the following states:
	//  `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the NasJob was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.NasJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Only populated when job's state is JOB_STATE_FAILED or
	//  JOB_STATE_CANCELLED.
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
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformnasjob;gcpaiplatformnasjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformNasJob is the Schema for the AIPlatformNasJob API
// +k8s:openapi-gen=true
type AIPlatformNasJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformNasJobSpec   `json:"spec,omitempty"`
	Status AIPlatformNasJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformNasJobList contains a list of AIPlatformNasJob
type AIPlatformNasJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformNasJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformNasJob{}, &AIPlatformNasJobList{})
}
