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


// +kcc:proto=google.cloud.video.stitcher.v1.Slate
type Slate struct {

	// The URI to fetch the source content for the slate. This URI must return an
	//  MP4 video with at least one audio track.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Slate.uri
	URI *string `json:"uri,omitempty"`

	// gam_slate has all the GAM-related attributes of slates.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Slate.gam_slate
	GamSlate *Slate_GamSlate `json:"gamSlate,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.Slate.GamSlate
type Slate_GamSlate struct {
	// Required. Ad Manager network code to associate with the live config.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Slate.GamSlate.network_code
	NetworkCode *string `json:"networkCode,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.Slate
type SlateObservedState struct {
	// Output only. The name of the slate, in the form of
	//  `projects/{project_number}/locations/{location}/slates/{id}`.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Slate.name
	Name *string `json:"name,omitempty"`

	// gam_slate has all the GAM-related attributes of slates.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Slate.gam_slate
	GamSlate *Slate_GamSlateObservedState `json:"gamSlate,omitempty"`
}

// +kcc:proto=google.cloud.video.stitcher.v1.Slate.GamSlate
type Slate_GamSlateObservedState struct {
	// Output only. The identifier generated for the slate by GAM.
	// +kcc:proto:field=google.cloud.video.stitcher.v1.Slate.GamSlate.gam_slate_id
	GamSlateID *int64 `json:"gamSlateID,omitempty"`
}
