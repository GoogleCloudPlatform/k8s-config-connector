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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
)

var VertexAINasJobGVK = GroupVersion.WithKind("VertexAINasJob")

// VertexAINasJobSpec defines the desired state of VertexAINasJob
type VertexAINasJobSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAINasJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the NasJob.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The specification of a NasJob.
	NasJobSpec *NasJobSpec `json:"nasJobSpec,omitempty"`

	// The labels with user-defined metadata to organize NasJobs.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key options for a NasJob.
	//  If this is set, then all resources created by the NasJob
	//  will be encrypted with the provided encryption key.
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Optional. Enable a separation of Custom model training
	//  and restricted image training for tenant project.
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
type VertexAINasJobObservedState struct {
	// Output only. Resource name of the NasJob.
	Name *string `json:"name,omitempty"`

	// Output only. Output of the NasJob.
	NasJobOutput *NasJobOutputObservedState `json:"nasJobOutput,omitempty"`

	// Output only. The detailed state of the job.
	State *string `json:"state,omitempty"`

	// Output only. Time when the NasJob was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the NasJob for the first time entered the
	//  `JOB_STATE_RUNNING` state.
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the NasJob entered any of the following states:
	//  `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the NasJob was most recently updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Only populated when job's state is JOB_STATE_FAILED or
	//  JOB_STATE_CANCELLED.
	Error *common.Status `json:"error,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
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
