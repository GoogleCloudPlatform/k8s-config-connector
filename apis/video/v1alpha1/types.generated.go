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


// +kcc:proto=google.cloud.video.stitcher.v1.Companion
type Companion struct {
	// The IFrame ad resource associated with the companion ad.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Companion.iframe_ad_resource
	IframeAdResource *IframeAdResource `json:"iframeAdResource,omitempty"`

	// The static ad resource associated with the companion ad.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Companion.static_ad_resource
	StaticAdResource *StaticAdResource `json:"staticAdResource,omitempty"`

	// The HTML ad resource associated with the companion ad.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Companion.html_ad_resource
	HTMLAdResource *HtmlAdResource `json:"htmlAdResource,omitempty"`

	// The API necessary to communicate with the creative if available.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Companion.api_framework
	ApiFramework *string `json:"apiFramework,omitempty"`

	// The pixel height of the placement slot for the intended creative.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Companion.height_px
	HeightPx *int32 `json:"heightPx,omitempty"`

	// The pixel width of the placement slot for the intended creative.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Companion.width_px
	WidthPx *int32 `json:"widthPx,omitempty"`

	// The pixel height of the creative.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Companion.asset_height_px
	AssetHeightPx *int32 `json:"assetHeightPx,omitempty"`

	// The maximum pixel height of the creative in its expanded state.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Companion.expanded_height_px
	ExpandedHeightPx *int32 `json:"expandedHeightPx,omitempty"`

	// The pixel width of the creative.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Companion.asset_width_px
	AssetWidthPx *int32 `json:"assetWidthPx,omitempty"`

	// The maximum pixel width of the creative in its expanded state.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Companion.expanded_width_px
	ExpandedWidthPx *int32 `json:"expandedWidthPx,omitempty"`

	// The ID used to identify the desired placement on a publisher's page.
	//  Values to be used should be discussed between publishers and
	//  advertisers.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Companion.ad_slot_id
	AdSlotID *string `json:"adSlotID,omitempty"`

	// The list of tracking events for the companion.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Companion.events
	Events []Event `json:"events,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.CompanionAds
type CompanionAds struct {
	// Indicates how many of the companions should be displayed with the ad.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.CompanionAds.display_requirement
	DisplayRequirement *string `json:"displayRequirement,omitempty"`

	// List of companion ads.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.CompanionAds.companions
	Companions []Companion `json:"companions,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.Event
type Event struct {
	// Describes the event that occurred.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Event.type
	Type *string `json:"type,omitempty"`

	// The URI to trigger for this event.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Event.uri
	URI *string `json:"uri,omitempty"`

	// The ID of the event.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Event.id
	ID *string `json:"id,omitempty"`

	// The offset in seconds if the event type is `PROGRESS`.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Event.offset
	Offset *string `json:"offset,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.HtmlAdResource
type HtmlAdResource struct {
	// The HTML to display for the ad resource.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.HtmlAdResource.html_source
	HTMLSource *string `json:"htmlSource,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.IframeAdResource
type IframeAdResource struct {
	// URI source for an IFrame to display for the ad resource.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.IframeAdResource.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.Interstitials
type Interstitials struct {
	// List of ad breaks ordered by time.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Interstitials.ad_breaks
	AdBreaks []VodSessionAdBreak `json:"adBreaks,omitempty"`

	// Information related to the content of the VOD session.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Interstitials.session_content
	SessionContent *VodSessionContent `json:"sessionContent,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.ManifestOptions
type ManifestOptions struct {
	// If specified, the output manifest will only return renditions matching the
	//  specified filters.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.ManifestOptions.include_renditions
	IncludeRenditions []RenditionFilter `json:"includeRenditions,omitempty"`

	// If specified, the output manifest will orders the video and muxed
	//  renditions by bitrate according to the ordering policy.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.ManifestOptions.bitrate_order
	BitrateOrder *string `json:"bitrateOrder,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.ProgressEvent
type ProgressEvent struct {
	// The time when the following tracking events occurs. The time is in
	//  seconds relative to the start of the VOD asset.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.ProgressEvent.time_offset
	TimeOffset *string `json:"timeOffset,omitempty"`

	// The list of progress tracking events for the ad break. These can be of
	//  the following IAB types: `BREAK_START`, `BREAK_END`, `IMPRESSION`,
	//  `CREATIVE_VIEW`, `START`, `FIRST_QUARTILE`, `MIDPOINT`, `THIRD_QUARTILE`,
	//  `COMPLETE`, `PROGRESS`.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.ProgressEvent.events
	Events []Event `json:"events,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.RenditionFilter
type RenditionFilter struct {
	// Bitrate in bits per second for the rendition. If set, only renditions with
	//  the exact bitrate will match.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.RenditionFilter.bitrate_bps
	BitrateBps *int32 `json:"bitrateBps,omitempty"`

	// Codecs for the rendition. If set, only renditions with the exact value
	//  will match.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.RenditionFilter.codecs
	Codecs *string `json:"codecs,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.StaticAdResource
type StaticAdResource struct {
	// URI to the static file for the ad resource.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.StaticAdResource.uri
	URI *string `json:"uri,omitempty"`

	// Describes the MIME type of the ad resource.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.StaticAdResource.creative_type
	CreativeType *string `json:"creativeType,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.VodSession
type VodSession struct {

	// URI of the media to stitch. For most use cases, you should create a
	//  [VodConfig][google.cloud.video.stitcher.v1.VodConfig] with this information
	//  rather than setting this field directly.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.source_uri
	SourceURI *string `json:"sourceURI,omitempty"`

	// Ad tag URI. For most use cases, you should create a
	//  [VodConfig][google.cloud.video.stitcher.v1.VodConfig] with this information
	//  rather than setting this field directly.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.ad_tag_uri
	AdTagURI *string `json:"adTagURI,omitempty"`

	// Key value pairs for ad tag macro replacement, only available for VOD
	//  sessions that do not implement Google Ad manager ad insertion. If the
	//  specified ad tag URI has macros, this field provides the mapping to the
	//  value that will replace the macro in the ad tag URI.
	//
	//  Macros are designated by square brackets, for example:
	//
	//    Ad tag URI: `"https://doubleclick.google.com/ad/1?geo_id=[geoId]"`
	//
	//    Ad tag macro map: `{"geoId": "123"}`
	//
	//    Fully qualified ad tag:
	//    `"https://doubleclick.google.com/ad/1?geo_id=123"`
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.ad_tag_macro_map
	AdTagMacroMap map[string]string `json:"adTagMacroMap,omitempty"`

	// Additional options that affect the output of the manifest.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.manifest_options
	ManifestOptions *ManifestOptions `json:"manifestOptions,omitempty"`

	// Required. Determines how the ad should be tracked.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.ad_tracking
	AdTracking *string `json:"adTracking,omitempty"`

	// This field should be set with appropriate values if GAM is being used for
	//  ads.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.gam_settings
	GamSettings *VodSession_GamSettings `json:"gamSettings,omitempty"`

	// The resource name of the VOD config for this session, in the form of
	//  `projects/{project}/locations/{location}/vodConfigs/{id}`.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.vod_config
	VodConfig *string `json:"vodConfig,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.VodSession.GamSettings
type VodSession_GamSettings struct {
	// Required. Ad Manager network code.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.GamSettings.network_code
	NetworkCode *string `json:"networkCode,omitempty"`

	// Required. The stream ID generated by Ad Manager.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.GamSettings.stream_id
	StreamID *string `json:"streamID,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.VodSessionAd
type VodSessionAd struct {
	// Duration in seconds of the ad.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSessionAd.duration
	Duration *string `json:"duration,omitempty"`

	// Metadata of companion ads associated with the ad.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSessionAd.companion_ads
	CompanionAds *CompanionAds `json:"companionAds,omitempty"`

	// The list of progress tracking events for the ad break. These can be of
	//  the following IAB types: `MUTE`, `UNMUTE`, `PAUSE`, `CLICK`,
	//  `CLICK_THROUGH`, `REWIND`, `RESUME`, `ERROR`, `FULLSCREEN`,
	//  `EXIT_FULLSCREEN`, `EXPAND`, `COLLAPSE`, `ACCEPT_INVITATION_LINEAR`,
	//  `CLOSE_LINEAR`, `SKIP`.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSessionAd.activity_events
	ActivityEvents []Event `json:"activityEvents,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.VodSessionAdBreak
type VodSessionAdBreak struct {
	// List of events that are expected to be triggered, ordered by time.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSessionAdBreak.progress_events
	ProgressEvents []ProgressEvent `json:"progressEvents,omitempty"`

	// Ordered list of ads stitched into the ad break.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSessionAdBreak.ads
	Ads []VodSessionAd `json:"ads,omitempty"`

	// Ad break end time in seconds relative to the start of the VOD asset.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSessionAdBreak.end_time_offset
	EndTimeOffset *string `json:"endTimeOffset,omitempty"`

	// Ad break start time in seconds relative to the start of the VOD asset.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSessionAdBreak.start_time_offset
	StartTimeOffset *string `json:"startTimeOffset,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.VodSessionContent
type VodSessionContent struct {
	// The total duration in seconds of the content including the ads stitched
	//  in.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSessionContent.duration
	Duration *string `json:"duration,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.VodSession
type VodSessionObservedState struct {
	// Output only. The name of the VOD session, in the form of
	//  `projects/{project_number}/locations/{location}/vodSessions/{id}`.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.name
	Name *string `json:"name,omitempty"`

	// Output only. Metadata of what was stitched into the content.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.interstitials
	Interstitials *Interstitials `json:"interstitials,omitempty"`

	// Output only. The playback URI of the stitched content.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.play_uri
	PlayURI *string `json:"playURI,omitempty"`

	// Output only. The generated ID of the VodSession's source media.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.VodSession.asset_id
	AssetID *string `json:"assetID,omitempty"`
}
