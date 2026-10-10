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

// API sources for LiveStreamDVRSession, recorded by generate-types:
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

var LiveStreamDVRSessionGVK = GroupVersion.WithKind("LiveStreamDVRSession")

// LiveStreamDVRSessionSpec defines the desired state of LiveStreamDVRSession
// +kcc:spec:proto=google.cloud.video.livestream.v1.DvrSession
// +kcc:required-from-proto
type LiveStreamDVRSessionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/channels/{channel} this resource belongs to.
	// +kcc:guess
	// ChannelRef *LiveStreamChannelRef `json:"channelRef,omitempty"`

	// The LiveStreamDVRSession name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. User-defined key/value metadata.
	// +kcc:proto:field=google.cloud.video.livestream.v1.DvrSession.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. A list of DVR manifests. Currently only one DVR manifest is
	//  allowed.
	// +kcc:proto:field=google.cloud.video.livestream.v1.DvrSession.dvr_manifests
	// +required
	DvrManifests []DvrSession_DvrManifest `json:"dvrManifests,omitempty"`

	// Required. The specified ranges of segments to generate a DVR recording.
	// +kcc:proto:field=google.cloud.video.livestream.v1.DvrSession.dvr_windows
	// +required
	DvrWindows []DvrSession_DvrWindow `json:"dvrWindows,omitempty"`
}

// LiveStreamDVRSessionStatus defines the config connector machine state of LiveStreamDVRSession
type LiveStreamDVRSessionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the LiveStreamDVRSession resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *LiveStreamDVRSessionObservedState `json:"observedState,omitempty"`
}

// LiveStreamDVRSessionObservedState is the state of the LiveStreamDVRSession resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.video.livestream.v1.DvrSession
type LiveStreamDVRSessionObservedState struct {
	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.DvrSession.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time.
	// +kcc:proto:field=google.cloud.video.livestream.v1.DvrSession.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the clip.
	// +kcc:proto:field=google.cloud.video.livestream.v1.DvrSession.state
	State *string `json:"state,omitempty"`

	// Output only. An error object that describes the reason for the failure.
	//  This property only presents when `state` is `FAILED`.
	// +kcc:proto:field=google.cloud.video.livestream.v1.DvrSession.error
	Error *common.Status `json:"error,omitempty"`

	// Required. A list of DVR manifests. Currently only one DVR manifest is
	//  allowed.
	// +kcc:proto:field=google.cloud.video.livestream.v1.DvrSession.dvr_manifests
	DvrManifests []DvrSession_DvrManifestObservedState `json:"dvrManifests,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplivestreamdvrsession;gcplivestreamdvrsessions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LiveStreamDVRSession is the Schema for the LiveStreamDVRSession API
// +k8s:openapi-gen=true
type LiveStreamDVRSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LiveStreamDVRSessionSpec   `json:"spec,omitempty"`
	Status LiveStreamDVRSessionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LiveStreamDVRSessionList contains a list of LiveStreamDVRSession
type LiveStreamDVRSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiveStreamDVRSession `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LiveStreamDVRSession{}, &LiveStreamDVRSessionList{})
}
