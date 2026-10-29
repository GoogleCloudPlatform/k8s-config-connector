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

// API sources for VideoStitcherLiveConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/video/stitcher/v1/live_configs.proto
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

var VideoStitcherLiveConfigGVK = GroupVersion.WithKind("VideoStitcherLiveConfig")

// VideoStitcherLiveConfigSpec defines the desired state of VideoStitcherLiveConfig
// +kcc:spec:proto=google.cloud.video.stitcher.v1.LiveConfig
// +kcc:required-from-proto
type VideoStitcherLiveConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VideoStitcherLiveConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Source URI for the live stream manifest.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.source_uri
	// +required
	SourceURI *string `json:"sourceURI,omitempty"`

	// The default ad tag associated with this live stream config.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.ad_tag_uri
	AdTagURI *string `json:"adTagURI,omitempty"`

	// Additional metadata used to register a live stream with Google Ad Manager
	//  (GAM)
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.gam_live_config
	GamLiveConfig *GamLiveConfig `json:"gamLiveConfig,omitempty"`

	// Required. Determines how the ads are tracked.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.ad_tracking
	// +required
	AdTracking *string `json:"adTracking,omitempty"`

	// This must refer to a slate in the same
	//  project. If Google Ad Manager (GAM) is used for ads, this string sets the
	//  value of `slateCreativeId` in
	//  https://developers.google.com/ad-manager/api/reference/v202211/LiveStreamEventService.LiveStreamEvent#slateCreativeId
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.default_slate
	DefaultSlate *string `json:"defaultSlate,omitempty"`

	// Defines the stitcher behavior in case an ad does not align exactly with
	//  the ad break boundaries. If not specified, the default is `CUT_CURRENT`.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.stitching_policy
	StitchingPolicy *string `json:"stitchingPolicy,omitempty"`

	// The configuration for prefetching ads.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.prefetch_config
	PrefetchConfig *PrefetchConfig `json:"prefetchConfig,omitempty"`

	// Options for fetching source manifests and segments.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.source_fetch_options
	SourceFetchOptions *FetchOptions `json:"sourceFetchOptions,omitempty"`
}

// VideoStitcherLiveConfigStatus defines the config connector machine state of VideoStitcherLiveConfig
type VideoStitcherLiveConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VideoStitcherLiveConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VideoStitcherLiveConfigObservedState `json:"observedState,omitempty"`
}

// VideoStitcherLiveConfigObservedState is the state of the VideoStitcherLiveConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.video.stitcher.v1.LiveConfig
type VideoStitcherLiveConfigObservedState struct {
	// Additional metadata used to register a live stream with Google Ad Manager
	//  (GAM)
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.gam_live_config
	GamLiveConfig *GamLiveConfigObservedState `json:"gamLiveConfig,omitempty"`

	// Output only. State of the live config.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvideostitcherliveconfig;gcpvideostitcherliveconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VideoStitcherLiveConfig is the Schema for the VideoStitcherLiveConfig API
// +k8s:openapi-gen=true
type VideoStitcherLiveConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VideoStitcherLiveConfigSpec   `json:"spec,omitempty"`
	Status VideoStitcherLiveConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VideoStitcherLiveConfigList contains a list of VideoStitcherLiveConfig
type VideoStitcherLiveConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VideoStitcherLiveConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VideoStitcherLiveConfig{}, &VideoStitcherLiveConfigList{})
}
