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


// +kcc:proto=google.cloud.video.stitcher.v1.AkamaiCdnKey
type AkamaiCdnKey struct {
	// Input only. Token key for the Akamai CDN edge configuration.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.AkamaiCdnKey.token_key
	TokenKey []byte `json:"tokenKey,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.CdnKey
type CdnKey struct {
	// The configuration for a Google Cloud CDN key.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.CdnKey.google_cdn_key
	GoogleCdnKey *GoogleCdnKey `json:"googleCdnKey,omitempty"`

	// The configuration for an Akamai CDN key.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.CdnKey.akamai_cdn_key
	AkamaiCdnKey *AkamaiCdnKey `json:"akamaiCdnKey,omitempty"`

	// The configuration for a Media CDN key.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.CdnKey.media_cdn_key
	MediaCdnKey *MediaCdnKey `json:"mediaCdnKey,omitempty"`

	// The resource name of the CDN key, in the form of
	//  `projects/{project}/locations/{location}/cdnKeys/{id}`.
	//  The name is ignored when creating a CDN key.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.CdnKey.name
	Name *string `json:"name,omitempty"`

	// The hostname this key applies to.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.CdnKey.hostname
	Hostname *string `json:"hostname,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.GoogleCdnKey
type GoogleCdnKey struct {
	// Input only. Secret for this Google Cloud CDN key.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.GoogleCdnKey.private_key
	PrivateKey []byte `json:"privateKey,omitempty"`

	// The public name of the Google Cloud CDN key.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.GoogleCdnKey.key_name
	KeyName *string `json:"keyName,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.MediaCdnKey
type MediaCdnKey struct {
	// Input only. 64-byte ed25519 private key for this Media CDN key.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.MediaCdnKey.private_key
	PrivateKey []byte `json:"privateKey,omitempty"`

	// The keyset name of the Media CDN key.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.MediaCdnKey.key_name
	KeyName *string `json:"keyName,omitempty"`

	// Optional. If set, the URL will be signed using the Media CDN token.
	//  Otherwise, the URL would be signed using the standard Media CDN signature.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.MediaCdnKey.token_config
	TokenConfig *MediaCdnKey_TokenConfig `json:"tokenConfig,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.MediaCdnKey.TokenConfig
type MediaCdnKey_TokenConfig struct {
	// Optional. The query parameter in which to find the token.
	//
	//  The name must be 1-64 characters long and match
	//  the regular expression `[a-zA-Z]([a-zA-Z0-9_-])*` which means the
	//  first character must be a letter, and all following characters
	//  must be a dash, underscore, letter or digit.
	//
	//  Defaults to `edge-cache-token`.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.MediaCdnKey.TokenConfig.query_parameter
	QueryParameter *string `json:"queryParameter,omitempty"`
}
