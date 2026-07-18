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

var VertexAICustomJobGVK = GroupVersion.WithKind("VertexAICustomJob")

// VertexAICustomJobSpec defines the desired state of VertexAICustomJob
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.CustomJob
type VertexAICustomJobSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The VertexAICustomJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the CustomJob.
	// The name can be up to 128 characters long and can consist of any UTF-8 characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Job spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.job_spec
	// +required
	JobSpec *CustomJobSpec `json:"jobSpec,omitempty"`

	// Optional. The labels with user-defined metadata to organize CustomJobs.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Customer-managed encryption key options for a CustomJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAICustomJobStatus defines the config connector machine state of VertexAICustomJob
type VertexAICustomJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAICustomJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAICustomJobObservedState `json:"observedState,omitempty"`
}

// VertexAICustomJobObservedState is the state of the VertexAICustomJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.CustomJob
type VertexAICustomJobObservedState struct {
	// Output only. Resource name of a CustomJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.name
	Name *string `json:"name,omitempty"`

	// Output only. The detailed state of the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.state
	State *string `json:"state,omitempty"`

	// Output only. Time when the CustomJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the CustomJob for the first time entered the `JOB_STATE_RUNNING` state.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the CustomJob entered any of the following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the CustomJob was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Only populated when job's state is `JOB_STATE_FAILED` or `JOB_STATE_CANCELLED`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. URIs for accessing interactive shells.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.web_access_uris
	WebAccessUris map[string]string `json:"webAccessUris,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CustomJob.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaicustomjob;gcpvertexaicustomjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAICustomJob is the Schema for the VertexAICustomJob API
// +k8s:openapi-gen=true
type VertexAICustomJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAICustomJobSpec   `json:"spec,omitempty"`
	Status VertexAICustomJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAICustomJobList contains a list of VertexAICustomJob
type VertexAICustomJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAICustomJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAICustomJob{}, &VertexAICustomJobList{})
}
