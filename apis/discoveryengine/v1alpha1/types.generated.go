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


// +kcc:proto=google.cloud.discoveryengine.v1beta.SampleQuerySet
type SampleQuerySet struct {
	// Identifier. The full resource name of the
	//  [SampleQuerySet][google.cloud.discoveryengine.v1beta.SampleQuerySet], in
	//  the format of
	//  `projects/{project}/locations/{location}/sampleQuerySets/{sample_query_set}`.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 1024
	//  characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuerySet.name
	Name *string `json:"name,omitempty"`

	// Required. The sample query set display name.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 128
	//  characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuerySet.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The description of the
	//  [SampleQuerySet][google.cloud.discoveryengine.v1beta.SampleQuerySet].
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuerySet.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SampleQuerySet
type SampleQuerySetObservedState struct {
	// Output only. Timestamp the
	//  [SampleQuerySet][google.cloud.discoveryengine.v1beta.SampleQuerySet] was
	//  created at.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuerySet.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
