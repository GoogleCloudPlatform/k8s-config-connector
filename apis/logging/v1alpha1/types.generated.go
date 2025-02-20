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

// +kcc:proto=google.logging.v2.LogExclusion
type LogExclusion struct {
	// Required. A client-assigned identifier, such as
	//  `"load-balancer-exclusion"`. Identifiers are limited to 100 characters and
	//  can include only letters, digits, underscores, hyphens, and periods. First
	//  character has to be alphanumeric.
	// +kcc:proto:field=google.logging.v2.LogExclusion.name
	Name *string `json:"name,omitempty"`

	// Optional. A description of this exclusion.
	// +kcc:proto:field=google.logging.v2.LogExclusion.description
	Description *string `json:"description,omitempty"`

	// Required. An [advanced logs
	//  filter](https://cloud.google.com/logging/docs/view/advanced-queries) that
	//  matches the log entries to be excluded. By using the [sample
	//  function](https://cloud.google.com/logging/docs/view/advanced-queries#sample),
	//  you can exclude less than 100% of the matching log entries.
	//
	//  For example, the following query matches 99% of low-severity log entries
	//  from Google Cloud Storage buckets:
	//
	//    `resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)`
	// +kcc:proto:field=google.logging.v2.LogExclusion.filter
	Filter *string `json:"filter,omitempty"`

	// Optional. If set to True, then this exclusion is disabled and it does not
	//  exclude any log entries. You can [update an
	//  exclusion][google.logging.v2.ConfigServiceV2.UpdateExclusion] to change the
	//  value of this field.
	// +kcc:proto:field=google.logging.v2.LogExclusion.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.logging.v2.LogExclusion
type LogExclusionObservedState struct {
	// Output only. The creation timestamp of the exclusion.
	//
	//  This field may not be present for older exclusions.
	// +kcc:proto:field=google.logging.v2.LogExclusion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of the exclusion.
	//
	//  This field may not be present for older exclusions.
	// +kcc:proto:field=google.logging.v2.LogExclusion.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
