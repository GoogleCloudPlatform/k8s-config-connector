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


// +kcc:proto=google.logging.v2.LogView
type LogView struct {
	// The resource name of the view.
	//
	//  For example:
	//
	//    `projects/my-project/locations/global/buckets/my-bucket/views/my-view`
	// +kcc:proto:field=google.logging.v2.LogView.name
	Name *string `json:"name,omitempty"`

	// Describes this view.
	// +kcc:proto:field=google.logging.v2.LogView.description
	Description *string `json:"description,omitempty"`

	// Filter that restricts which log entries in a bucket are visible in this
	//  view.
	//
	//  Filters are restricted to be a logical AND of ==/!= of any of the
	//  following:
	//
	//    - originating project/folder/organization/billing account.
	//    - resource type
	//    - log id
	//
	//  For example:
	//
	//    SOURCE("projects/myproject") AND resource.type = "gce_instance"
	//                                 AND LOG_ID("stdout")
	// +kcc:proto:field=google.logging.v2.LogView.filter
	Filter *string `json:"filter,omitempty"`
}

// +kcc:proto=google.logging.v2.LogView
type LogViewObservedState struct {
	// Output only. The creation timestamp of the view.
	// +kcc:proto:field=google.logging.v2.LogView.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of the view.
	// +kcc:proto:field=google.logging.v2.LogView.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
