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


// +kcc:proto=google.cloud.discoveryengine.v1beta.SampleQuery
type SampleQuery struct {
	// The query entry.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuery.query_entry
	QueryEntry *SampleQuery_QueryEntry `json:"queryEntry,omitempty"`

	// Identifier. The full resource name of the sample query, in the format of
	//  `projects/{project}/locations/{location}/sampleQuerySets/{sample_query_set}/sampleQueries/{sample_query}`.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 1024
	//  characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuery.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SampleQuery.QueryEntry
type SampleQuery_QueryEntry struct {
	// Required. The query.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuery.QueryEntry.query
	Query *string `json:"query,omitempty"`

	// List of targets for the query.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuery.QueryEntry.targets
	Targets []SampleQuery_QueryEntry_Target `json:"targets,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SampleQuery.QueryEntry.Target
type SampleQuery_QueryEntry_Target struct {
	// Expected uri of the target.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 2048
	//  characters.
	//
	//  Example of valid uris: `https://example.com/abc`,
	//  `gcs://example/example.pdf`.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuery.QueryEntry.Target.uri
	URI *string `json:"uri,omitempty"`

	// Expected page numbers of the target.
	//
	//  Each page number must be non negative.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuery.QueryEntry.Target.page_numbers
	PageNumbers []int32 `json:"pageNumbers,omitempty"`

	// Relevance score of the target.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuery.QueryEntry.Target.score
	Score *float64 `json:"score,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.SampleQuery
type SampleQueryObservedState struct {
	// Output only. Timestamp the
	//  [SampleQuery][google.cloud.discoveryengine.v1beta.SampleQuery] was created
	//  at.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuery.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
