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


// +kcc:proto=google.cloud.video.stitcher.v1.AdRequest
type AdRequest struct {
	// The ad tag URI processed with integrated macros.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.AdRequest.uri
	URI *string `json:"uri,omitempty"`

	// The request metadata used to make the ad request.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.AdRequest.request_metadata
	RequestMetadata *RequestMetadata `json:"requestMetadata,omitempty"`

	// The response metadata received from the ad request.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.AdRequest.response_metadata
	ResponseMetadata *ResponseMetadata `json:"responseMetadata,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.LiveAdTagDetail
type LiveAdTagDetail struct {
	// The resource name in the form of
	//  `projects/{project}/locations/{location}/liveSessions/{live_session}/liveAdTagDetails/{id}`.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveAdTagDetail.name
	Name *string `json:"name,omitempty"`

	// A list of ad requests.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.LiveAdTagDetail.ad_requests
	AdRequests []AdRequest `json:"adRequests,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.RequestMetadata
type RequestMetadata struct {
	// The HTTP headers of the ad request.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.RequestMetadata.headers
	Headers map[string]string `json:"headers,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.ResponseMetadata
type ResponseMetadata struct {
	// Error message received when making the ad request.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.ResponseMetadata.error
	Error *string `json:"error,omitempty"`

	// Headers from the response.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.ResponseMetadata.headers
	Headers map[string]string `json:"headers,omitempty"`

	// Status code for the response.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.ResponseMetadata.status_code
	StatusCode *string `json:"statusCode,omitempty"`

	// Size in bytes of the response.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.ResponseMetadata.size_bytes
	SizeBytes *int32 `json:"sizeBytes,omitempty"`

	// Total time elapsed for the response.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.ResponseMetadata.duration
	Duration *string `json:"duration,omitempty"`

	// The body of the response.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.ResponseMetadata.body
	Body *string `json:"body,omitempty"`
}
