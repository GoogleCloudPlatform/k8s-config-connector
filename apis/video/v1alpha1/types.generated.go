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

// +kcc:proto=google.cloud.video.stitcher.v1.GamVodConfig
type GamVodConfig struct {
	// Required. Ad Manager network code to associate with the VOD config.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.GamVodConfig.network_code
	NetworkCode *string `json:"networkCode,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.VodConfig
type VodConfig struct {

	// Required. Source URI for the VOD stream manifest.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodConfig.source_uri
	SourceURI *string `json:"sourceURI,omitempty"`

	// Required. The default ad tag associated with this VOD config.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodConfig.ad_tag_uri
	AdTagURI *string `json:"adTagURI,omitempty"`

	// Optional. Google Ad Manager (GAM) metadata.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodConfig.gam_vod_config
	GamVodConfig *GamVodConfig `json:"gamVodConfig,omitempty"`

	// Options for fetching source manifests and segments.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodConfig.source_fetch_options
	SourceFetchOptions *FetchOptions `json:"sourceFetchOptions,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.VodConfig
type VodConfigObservedState struct {
	// Output only. The resource name of the VOD config, in the form of
	//  `projects/{project}/locations/{location}/vodConfigs/{id}`.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodConfig.name
	Name *string `json:"name,omitempty"`

	// Output only. State of the VOD config.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodConfig.state
	State *string `json:"state,omitempty"`
}
