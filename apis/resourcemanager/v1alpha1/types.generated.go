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


// +kcc:proto=google.cloud.resourcemanager.v3.TagValue
type TagValue struct {
	// Immutable. Resource name for TagValue in the format `tagValues/456`.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.name
	Name *string `json:"name,omitempty"`

	// Immutable. The resource name of the new TagValue's parent TagKey.
	//  Must be of the form `tagKeys/{tag_key_id}`.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.parent
	Parent *string `json:"parent,omitempty"`

	// Required. Immutable. User-assigned short name for TagValue. The short name
	//  should be unique for TagValues within the same parent TagKey.
	//
	//  The short name must be 63 characters or less, beginning and ending with
	//  an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_),
	//  dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.short_name
	ShortName *string `json:"shortName,omitempty"`

	// Optional. User-assigned description of the TagValue.
	//  Must not exceed 256 characters.
	//
	//  Read-write.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.description
	Description *string `json:"description,omitempty"`

	// Optional. Entity tag which users can pass to prevent race conditions. This
	//  field is always set in server responses. See UpdateTagValueRequest for
	//  details.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.resourcemanager.v3.TagValue
type TagValueObservedState struct {
	// Output only. The namespaced name of the TagValue. Can be in the form
	//  `{organization_id}/{tag_key_short_name}/{tag_value_short_name}` or
	//  `{project_id}/{tag_key_short_name}/{tag_value_short_name}` or
	//  `{project_number}/{tag_key_short_name}/{tag_value_short_name}`.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.namespaced_name
	NamespacedName *string `json:"namespacedName,omitempty"`

	// Output only. Creation time.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagValue.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
