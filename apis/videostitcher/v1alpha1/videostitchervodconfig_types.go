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

// API sources for VideoStitcherVODConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/video/stitcher/v1/vod_configs.proto
// +kcc:guess=source-link reason=verify-service-docs-link
// +kcc:source:service-docs=
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VideoStitcherVODConfigGVK = GroupVersion.WithKind("VideoStitcherVODConfig")

// VideoStitcherVODConfigSpec defines the desired state of VideoStitcherVODConfig
// +kcc:spec:proto=google.cloud.video.stitcher.v1.VodConfig
// +kcc:required-from-proto
type VideoStitcherVODConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VideoStitcherVODConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Source URI for the VOD stream manifest.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodConfig.source_uri
	// +required
	SourceURI *string `json:"sourceURI,omitempty"`

	// Required. The default ad tag associated with this VOD config.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodConfig.ad_tag_uri
	// +required
	AdTagURI *string `json:"adTagURI,omitempty"`

	// Optional. Google Ad Manager (GAM) metadata.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodConfig.gam_vod_config
	GamVodConfig *GamVodConfig `json:"gamVodConfig,omitempty"`

	// Options for fetching source manifests and segments.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodConfig.source_fetch_options
	SourceFetchOptions *FetchOptions `json:"sourceFetchOptions,omitempty"`
}

// VideoStitcherVODConfigStatus defines the config connector machine state of VideoStitcherVODConfig
type VideoStitcherVODConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VideoStitcherVODConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VideoStitcherVODConfigObservedState `json:"observedState,omitempty"`
}

// VideoStitcherVODConfigObservedState is the state of the VideoStitcherVODConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.video.stitcher.v1.VodConfig
type VideoStitcherVODConfigObservedState struct {
	// Output only. State of the VOD config.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodConfig.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvideostitchervodconfig;gcpvideostitchervodconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VideoStitcherVODConfig is the Schema for the VideoStitcherVODConfig API
// +k8s:openapi-gen=true
type VideoStitcherVODConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VideoStitcherVODConfigSpec   `json:"spec,omitempty"`
	Status VideoStitcherVODConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VideoStitcherVODConfigList contains a list of VideoStitcherVODConfig
type VideoStitcherVODConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VideoStitcherVODConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VideoStitcherVODConfig{}, &VideoStitcherVODConfigList{})
}
