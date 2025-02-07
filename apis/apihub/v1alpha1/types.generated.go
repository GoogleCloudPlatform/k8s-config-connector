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


// +kcc:proto=google.cloud.apihub.v1.StyleGuide
type StyleGuide struct {
	// Identifier. The name of the style guide.
	//
	//  Format:
	//  `projects/{project}/locations/{location}/plugins/{plugin}/styleGuide`
	// +kcc:proto:field=google.cloud.apihub.v1.StyleGuide.name
	Name *string `json:"name,omitempty"`

	// Required. Target linter for the style guide.
	// +kcc:proto:field=google.cloud.apihub.v1.StyleGuide.linter
	Linter *string `json:"linter,omitempty"`

	// Required. Input only. The contents of the uploaded style guide.
	// +kcc:proto:field=google.cloud.apihub.v1.StyleGuide.contents
	Contents *StyleGuideContents `json:"contents,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.StyleGuideContents
type StyleGuideContents struct {
	// Required. The contents of the style guide.
	// +kcc:proto:field=google.cloud.apihub.v1.StyleGuideContents.contents
	Contents []byte `json:"contents,omitempty"`

	// Required. The mime type of the content.
	// +kcc:proto:field=google.cloud.apihub.v1.StyleGuideContents.mime_type
	MimeType *string `json:"mimeType,omitempty"`
}
