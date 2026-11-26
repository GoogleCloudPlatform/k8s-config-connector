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

// API sources for LiveStreamClip, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/video/livestream/v1/resources.proto
// +kcc:guess=source-link reason=verify-service-docs-link
// +kcc:source:service-docs=
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=

package v1alpha1

import (
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var LiveStreamClipGVK = GroupVersion.WithKind("LiveStreamClip")

// LiveStreamClipSpec defines the desired state of LiveStreamClip
// +kcc:spec:proto=google.cloud.video.livestream.v1.Clip
// +kcc:required-from-proto
type LiveStreamClipSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/channels/{channel} this resource belongs to.
	// +kcc:guess
	// ChannelRef *LiveStreamChannelRef `json:"channelRef,omitempty"`

	// The LiveStreamClip name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The labels associated with this resource. Each label is a key-value pair.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Specify the `output_uri` to determine where to place the clip segments and
	//  clip manifest files in Cloud Storage. The manifests specified in
	//  `clip_manifests` fields will be placed under this URI. The exact URI of the
	//  generated manifests will be provided in `clip_manifests.output_uri` for
	//  each manifest.
	//  Example:
	//  "output_uri": "gs://my-bucket/clip-outputs"
	//  "clip_manifests.output_uri": "gs://my-bucket/clip-outputs/main.m3u8"
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.output_uri
	OutputURI *string `json:"outputURI,omitempty"`

	// The specified ranges of segments to generate a clip.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.slices
	Slices []Clip_Slice `json:"slices,omitempty"`

	// Required. A list of clip manifests. Currently only one clip manifest is
	//  allowed.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.clip_manifests
	// +required
	ClipManifests []Clip_ClipManifest `json:"clipManifests,omitempty"`

	// Optional. OutputType of the clip. If not specified, the default value is
	//  MANIFEST.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.output_type
	OutputType *string `json:"outputType,omitempty"`
}

// LiveStreamClipStatus defines the config connector machine state of LiveStreamClip
type LiveStreamClipStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the LiveStreamClip resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *LiveStreamClipObservedState `json:"observedState,omitempty"`
}

// LiveStreamClipObservedState is the state of the LiveStreamClip resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.video.livestream.v1.Clip
type LiveStreamClipObservedState struct {
	// Output only. The creation timestamp of the clip resource.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the clip request starts to be processed.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. The update timestamp of the clip resource.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the clip.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.state
	State *string `json:"state,omitempty"`

	// Output only. An error object that describes the reason for the failure.
	//  This property only presents when `state` is `FAILED`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.error
	Error *common.Status `json:"error,omitempty"`

	// Required. A list of clip manifests. Currently only one clip manifest is
	//  allowed.
	// +kcc:proto:field=google.cloud.video.livestream.v1.Clip.clip_manifests
	ClipManifests []Clip_ClipManifestObservedState `json:"clipManifests,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplivestreamclip;gcplivestreamclips
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LiveStreamClip is the Schema for the LiveStreamClip API
// +k8s:openapi-gen=true
type LiveStreamClip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LiveStreamClipSpec   `json:"spec,omitempty"`
	Status LiveStreamClipStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LiveStreamClipList contains a list of LiveStreamClip
type LiveStreamClipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiveStreamClip `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LiveStreamClip{}, &LiveStreamClipList{})
}
