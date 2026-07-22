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
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAITrainingPipelineGVK = GroupVersion.WithKind("VertexAITrainingPipeline")

// VertexAITrainingPipelineSpec defines the desired state of VertexAITrainingPipeline
// +kcc:spec:proto=google.cloud.aiplatform.v1.TrainingPipeline
type VertexAITrainingPipelineSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAITrainingPipeline name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The user-defined name of this TrainingPipeline.
	// +required
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Specifies Vertex AI owned input data that may be used for training the
	//  Model.
	// +optional
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.input_data_config
	InputDataConfig *InputDataConfig `json:"inputDataConfig,omitempty"`

	// Required. A Google Cloud Storage path to the YAML file that defines the
	//  training task which is responsible for producing the model artifact, and
	//  may also include additional auxiliary work. The definition files that can
	//  be used here are found in
	//  gs://google-cloud-aiplatform/schema/trainingjob/definition/.
	// +required
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.training_task_definition
	TrainingTaskDefinition *string `json:"trainingTaskDefinition,omitempty"`

	// Required. The training task's parameter(s), as specified in the
	//  trainingTaskDefinition's inputs.
	// +required
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.training_task_inputs
	TrainingTaskInputs *Value `json:"trainingTaskInputs,omitempty"`

	// Describes the Model that may be uploaded by this TrainingPipeline.
	// +optional
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.model_to_upload
	ModelToUpload *AIPlatformModelSpec `json:"modelToUpload,omitempty"`

	// Optional. The ID to use for the uploaded Model, which will become the final
	//  component of the model resource name.
	// +optional
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.model_id
	ModelID *string `json:"modelID,omitempty"`

	// Optional. When specify this field, the `model_to_upload` will not be
	//  uploaded as a new model, instead, it will become a new version of this
	//  `parent_model`.
	// +optional
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.parent_model
	ModelRef *AIPlatformModelRef `json:"modelRef,omitempty"`

	// The labels with user-defined metadata to organize TrainingPipelines.
	// +optional
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Customer-managed encryption key spec for a TrainingPipeline. If set, this
	//  TrainingPipeline will be secured by this key.
	// +optional
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAITrainingPipelineStatus defines the config connector machine state of VertexAITrainingPipeline
type VertexAITrainingPipelineStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAITrainingPipeline resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAITrainingPipelineObservedState `json:"observedState,omitempty"`
}

// VertexAITrainingPipelineObservedState is the state of the VertexAITrainingPipeline resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.TrainingPipeline
type VertexAITrainingPipelineObservedState struct {
	// Output only. Resource name of the TrainingPipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.name
	Name *string `json:"name,omitempty"`

	// Output only. The metadata information as specified in the
	//  trainingTaskDefinition's metadata. This metadata is an auxiliary runtime
	//  and final information about the training task. While the pipeline is
	//  running this information is populated only at a best effort basis.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.training_task_metadata
	TrainingTaskMetadata *Value `json:"trainingTaskMetadata,omitempty"`

	// Describes the Model that may be uploaded by this TrainingPipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.model_to_upload
	ModelToUpload *AIPlatformModelObservedState `json:"modelToUpload,omitempty"`

	// Output only. The detailed state of the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.state
	State *string `json:"state,omitempty"`

	// Output only. Only populated when the pipeline's state is
	//  PIPELINE_STATE_FAILED or PIPELINE_STATE_CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. Time when the TrainingPipeline was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the TrainingPipeline for the first time entered the
	//  PIPELINE_STATE_RUNNING state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the TrainingPipeline entered any of the following
	//  states: PIPELINE_STATE_SUCCEEDED, PIPELINE_STATE_FAILED,
	//  PIPELINE_STATE_CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Time when the TrainingPipeline was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.TrainingPipeline.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaitrainingpipeline;gcpvertexaitrainingpipelines
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAITrainingPipeline is the Schema for the VertexAITrainingPipeline API
// +k8s:openapi-gen=true
type VertexAITrainingPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAITrainingPipelineSpec   `json:"spec,omitempty"`
	Status VertexAITrainingPipelineStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAITrainingPipelineList contains a list of VertexAITrainingPipeline
type VertexAITrainingPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAITrainingPipeline `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAITrainingPipeline{}, &VertexAITrainingPipelineList{})
}
