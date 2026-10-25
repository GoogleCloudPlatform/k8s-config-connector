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

// API sources for TranscoderJobTemplate, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/video/transcoder/v1/resources.proto
// +kcc:source:service-docs=https://cloud.google.com/transcoder/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/transcoder/docs/reference/rest/v1/projects.locations.jobTemplates

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TranscoderJobTemplateGVK = GroupVersion.WithKind("TranscoderJobTemplate")

// TranscoderJobTemplateSpec defines the desired state of TranscoderJobTemplate
// +kcc:spec:proto=google.cloud.video.transcoder.v1.JobTemplate
// +kcc:required-from-proto
type TranscoderJobTemplateSpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// The TranscoderJobTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The configuration for this template.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobTemplate.config
	Config *TranscoderJobTemplate_JobConfig `json:"config,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.JobConfig
type TranscoderJobTemplate_JobConfig struct {
	// List of input assets stored in Cloud Storage.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.inputs
	Inputs []Input `json:"inputs,omitempty"`

	// List of edit atoms. Defines the ultimate timeline of the resulting
	//  file or manifest.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.edit_list
	EditList []EditAtom `json:"editList,omitempty"`

	// List of elementary streams.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.elementary_streams
	ElementaryStreams []ElementaryStreamWithRequired `json:"elementaryStreams,omitempty"`

	// List of multiplexing settings for output streams.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.mux_streams
	MuxStreams []MuxStreamWithRequired `json:"muxStreams,omitempty"`

	// List of output manifests.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.manifests
	Manifests []ManifestWithRequired `json:"manifests,omitempty"`

	// Output configuration.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.output
	Output *Output `json:"output,omitempty"`

	// List of ad breaks. Specifies where to insert ad break tags in the output
	//  manifests.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.ad_breaks
	AdBreaks []AdBreak `json:"adBreaks,omitempty"`

	// Destination on Pub/Sub.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.pubsub_destination
	PubsubDestination *PubsubDestination `json:"pubsubDestination,omitempty"`

	// List of output sprite sheets.
	//  Spritesheets require at least one VideoStream in the Jobconfig.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.sprite_sheets
	SpriteSheets []SpriteSheetWithRequired `json:"spriteSheets,omitempty"`

	// List of overlays on the output video, in descending Z-order.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.overlays
	Overlays []OverlayWithRequired `json:"overlays,omitempty"`

	// List of encryption configurations for the content.
	//  Each configuration has an ID. Specify this ID in the
	//  [MuxStream.encryption_id][google.cloud.video.transcoder.v1.MuxStream.encryption_id]
	//  field to indicate the configuration to use for that `MuxStream` output.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobConfig.encryptions
	Encryptions []TranscoderJobTemplate_Encryption `json:"encryptions,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Encryption
type TranscoderJobTemplate_Encryption struct {
	// Required. Identifier for this set of encryption options.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.id
	// +required
	ID *string `json:"id,omitempty"`

	// Configuration for AES-128 encryption.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.aes_128
	Aes128 *Encryption_Aes128Encryption `json:"aes128,omitempty"`

	// Configuration for SAMPLE-AES encryption.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.sample_aes
	SampleAes *Encryption_SampleAesEncryption `json:"sampleAes,omitempty"`

	// Configuration for MPEG Common Encryption (MPEG-CENC).
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.mpeg_cenc
	MpegCenc *Encryption_MpegCommonEncryptionWithRequired `json:"mpegCenc,omitempty"`

	// Keys are stored in Google Secret Manager.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.secret_manager_key_source
	SecretManagerKeySource *TranscoderJobTemplate_Encryption_SecretManagerSource `json:"secretManagerKeySource,omitempty"`

	// Required. DRM system(s) to use; at least one must be specified. If a
	//  DRM system is omitted, it is considered disabled.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.drm_systems
	// +required
	DrmSystems *Encryption_DrmSystems `json:"drmSystems,omitempty"`
}

// +kcc:proto=google.cloud.video.transcoder.v1.Encryption.SecretManagerSource
type TranscoderJobTemplate_Encryption_SecretManagerSource struct {
	// Required. The name of the Secret Version containing the encryption key in
	//  the following format:
	//  `projects/{project}/secrets/{secret_id}/versions/{version_number}`
	//
	//  Note that only numbered versions are supported. Aliases like "latest" are
	//  not supported.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.Encryption.SecretManagerSource.secret_version
	// +required
	SecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"secretVersionRef,omitempty"`
}

// TranscoderJobTemplateStatus defines the config connector machine state of TranscoderJobTemplate
type TranscoderJobTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TranscoderJobTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TranscoderJobTemplateObservedState `json:"observedState,omitempty"`
}

// TranscoderJobTemplateObservedState is the state of the TranscoderJobTemplate resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.video.transcoder.v1.JobTemplate
type TranscoderJobTemplateObservedState struct {
	// The configuration for this template.
	// +kcc:proto:field=google.cloud.video.transcoder.v1.JobTemplate.config
	Config *JobConfigObservedState `json:"config,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptranscoderjobtemplate;gcptranscoderjobtemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TranscoderJobTemplate is the Schema for the TranscoderJobTemplate API
// +k8s:openapi-gen=true
type TranscoderJobTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TranscoderJobTemplateSpec   `json:"spec,omitempty"`
	Status TranscoderJobTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TranscoderJobTemplateList contains a list of TranscoderJobTemplate
type TranscoderJobTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TranscoderJobTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TranscoderJobTemplate{}, &TranscoderJobTemplateList{})
}
