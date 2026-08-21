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
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TranscoderJobGVK = GroupVersion.WithKind("TranscoderJob")

// TranscoderJobSpec defines the desired state of TranscoderJob
// +kcc:spec:proto=google.cloud.video.transcoder.v1.Job
type TranscoderJobSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The TranscoderJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Input only. Specify the `input_uri` to populate empty `uri` fields in each
	// element of `Job.config.inputs` or `JobTemplate.config.inputs` when using
	// template. URI of the media. Input files must be at least 5 seconds in
	// duration and stored in Cloud Storage (for example,
	// `gs://bucket/inputs/file.mp4`). See [Supported input and output
	// formats](https://cloud.google.com/transcoder/docs/concepts/supported-input-and-output-formats).
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.input_uri
	InputURI *string `json:"inputURI,omitempty"`

	// Input only. Specify the `output_uri` to populate an empty
	// `Job.config.output.uri` or `JobTemplate.config.output.uri` when using
	// template. URI for the output file(s). For example,
	// `gs://my-bucket/outputs/`. See [Supported input and output
	// formats](https://cloud.google.com/transcoder/docs/concepts/supported-input-and-output-formats).
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.output_uri
	OutputURI *string `json:"outputURI,omitempty"`

	// Input only. Specify the `template_id` to use for populating `Job.config`.
	// The default is `preset/web-hd`, which is the only supported preset.
	//
	// User defined JobTemplate: `{job_template_id}`
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.template_id
	TemplateID *string `json:"templateID,omitempty"`

	// The configuration for this job.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.config
	Config *JobConfig `json:"config,omitempty"`

	// Job time to live value in days, which will be effective after job
	// completion. Job should be deleted automatically after the given TTL. Enter
	// a value between 1 and 90. The default is 30.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.ttl_after_completion_days
	TTLAfterCompletionDays *int32 `json:"ttlAfterCompletionDays,omitempty"`

	// The processing mode of the job.
	// The default is `PROCESSING_MODE_INTERACTIVE`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.mode
	Mode *string `json:"mode,omitempty"`

	// The processing priority of a batch job.
	// This field can only be set for batch mode jobs. The default value is 0.
	// This value cannot be negative. Higher values correspond to higher
	// priorities for the job.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.batch_mode_priority
	BatchModePriority *int32 `json:"batchModePriority,omitempty"`

	// Optional. The optimization strategy of the job. The default is
	// `AUTODETECT`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.optimization
	Optimization *string `json:"optimization,omitempty"`

	// Optional. Insert silence and duplicate frames when timestamp gaps are
	// detected in a given stream.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.fill_content_gaps
	FillContentGaps *bool `json:"fillContentGaps,omitempty"`
}

// TranscoderJobStatus defines the config connector machine state of TranscoderJob
type TranscoderJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TranscoderJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TranscoderJobObservedState `json:"observedState,omitempty"`
}

// TranscoderJobObservedState is the state of the TranscoderJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.video.transcoder.v1.Job
type TranscoderJobObservedState struct {
	// The configuration for this job.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.config
	Config *JobConfigObservedState `json:"config,omitempty"`

	// Output only. The current state of the job.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.state
	State *string `json:"state,omitempty"`

	// Output only. The time the job was created.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the transcoding started.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The time the transcoding finished.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. An error object that describes the reason for the failure.
	// This property is always present when
	// [ProcessingState][google.cloud.video.transcoder.v1.Job.ProcessingState] is
	// `FAILED`.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Job.error
	Error *common.Status `json:"error,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.PubsubDestination
type PubsubDestination struct {
	// TopicRef is a reference to a PubSubTopic.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.PubsubDestination.topic
	TopicRef *pubsubv1beta1.PubSubTopicRef `json:"topicRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptranscoderjob;gcptranscoderjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TranscoderJob is the Schema for the TranscoderJob API
// +k8s:openapi-gen=true
type TranscoderJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TranscoderJobSpec   `json:"spec,omitempty"`
	Status TranscoderJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TranscoderJobList contains a list of TranscoderJob
type TranscoderJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TranscoderJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TranscoderJob{}, &TranscoderJobList{})
}
