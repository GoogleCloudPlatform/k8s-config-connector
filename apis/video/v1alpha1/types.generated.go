// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.video.stitcher.v1.FetchOptions
type FetchOptions struct {
	// Custom headers to pass into fetch request.
	//  Headers must have a maximum of 3 key value pairs.
	//  Each key value pair must have a maximum of 256 characters per key and 256
	//  characters per value.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.FetchOptions.headers
	Headers map[string]string `json:"headers,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.GamLiveConfig
type GamLiveConfig struct {
	// Required. Ad Manager network code to associate with the live config.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.GamLiveConfig.network_code
	NetworkCode *string `json:"networkCode,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.LiveConfig
type LiveConfig struct {

	// Required. Source URI for the live stream manifest.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.source_uri
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

// +kcc:proto=google.cloud.video.stitcher.v1.PrefetchConfig
type PrefetchConfig struct {
	// Required. Indicates whether the option to prefetch ad requests is enabled.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.PrefetchConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// The duration in seconds of the part of the break to be prefetched.
	//  This field is only relevant if prefetch is enabled.
	//  You should set this duration to as long as possible to increase the
	//  benefits of prefetching, but not longer than the shortest ad break
	//  expected. For example, for a live event with 30s and 60s ad breaks, the
	//  initial duration should be set to 30s.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.PrefetchConfig.initial_ad_request_duration
	InitialAdRequestDuration *string `json:"initialAdRequestDuration,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.GamLiveConfig
type GamLiveConfigObservedState struct {
	// Output only. The asset key identifier generated for the live config.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.GamLiveConfig.asset_key
	AssetKey *string `json:"assetKey,omitempty"`

	// Output only. The custom asset key identifier generated for the live config.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.GamLiveConfig.custom_asset_key
	CustomAssetKey *string `json:"customAssetKey,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.LiveConfig
type LiveConfigObservedState struct {
	// Output only. The resource name of the live config, in the form of
	//  `projects/{project}/locations/{location}/liveConfigs/{id}`.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.name
	Name *string `json:"name,omitempty"`

	// Additional metadata used to register a live stream with Google Ad Manager
	//  (GAM)
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.gam_live_config
	GamLiveConfig *GamLiveConfigObservedState `json:"gamLiveConfig,omitempty"`

	// Output only. State of the live config.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveConfig.state
	State *string `json:"state,omitempty"`
}
