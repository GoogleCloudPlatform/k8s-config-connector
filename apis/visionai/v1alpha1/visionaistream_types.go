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

// API sources for VisionAIStream, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/visionai/v1/streams_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/vision-ai/docs
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/vision-ai/docs/reference/rest/v1/projects.locations.clusters.streams

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VisionAIStreamGVK = GroupVersion.WithKind("VisionAIStream")

// VisionAIStreamSpec defines the desired state of VisionAIStream
// +kcc:spec:proto=google.cloud.visionai.v1.Stream
// +kcc:required-from-proto
type VisionAIStreamSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VisionAIStream name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations to allow clients to store small amounts of arbitrary data.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// The display name for the stream resource.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Whether to enable the HLS playback service on this stream.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.enable_hls_playback
	EnableHlsPlayback *bool `json:"enableHlsPlayback,omitempty"`

	// The name of the media warehouse asset for long term storage of stream data.
	//  Format: projects/${p_id}/locations/${l_id}/corpora/${c_id}/assets/${a_id}
	//  Remain empty if the media warehouse storage is not needed for the stream.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.media_warehouse_asset
	MediaWarehouseAsset *string `json:"mediaWarehouseAsset,omitempty"`
}

// VisionAIStreamStatus defines the config connector machine state of VisionAIStream
type VisionAIStreamStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VisionAIStream resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VisionAIStreamObservedState `json:"observedState,omitempty"`
}

// VisionAIStreamObservedState is the state of the VisionAIStream resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.visionai.v1.Stream
type VisionAIStreamObservedState struct {
	// Output only. The create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.Stream.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvisionaistream;gcpvisionaistreams
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VisionAIStream is the Schema for the VisionAIStream API
// +k8s:openapi-gen=true
type VisionAIStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VisionAIStreamSpec   `json:"spec,omitempty"`
	Status VisionAIStreamStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VisionAIStreamList contains a list of VisionAIStream
type VisionAIStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VisionAIStream `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VisionAIStream{}, &VisionAIStreamList{})
}
