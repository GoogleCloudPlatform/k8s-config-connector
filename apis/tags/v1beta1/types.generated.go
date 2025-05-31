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

package v1beta1

// +kcc:proto=google.cloud.resourcemanager.v3.TagKey
type TagKey struct {
	// Immutable. The resource name for a TagKey. Must be in the format
	//  `tagKeys/{tag_key_id}`, where `tag_key_id` is the generated numeric id for
	//  the TagKey.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.name
	Name *string `json:"name,omitempty"`

	// Immutable. The resource name of the TagKey's parent. A TagKey can be
	//  parented by an Organization or a Project. For a TagKey parented by an
	//  Organization, its parent must be in the form `organizations/{org_id}`. For
	//  a TagKey parented by a Project, its parent can be in the form
	//  `projects/{project_id}` or `projects/{project_number}`.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.parent
	Parent *string `json:"parent,omitempty"`

	// Required. Immutable. The user friendly name for a TagKey. The short name
	//  should be unique for TagKeys within the same tag namespace.
	//
	//  The short name must be 1-63 characters, beginning and ending with
	//  an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_),
	//  dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.short_name
	ShortName *string `json:"shortName,omitempty"`

	// Optional. User-assigned description of the TagKey. Must not exceed 256
	//  characters.
	//
	//  Read-write.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.description
	Description *string `json:"description,omitempty"`

	// Optional. Entity tag which users can pass to prevent race conditions. This
	//  field is always set in server responses. See UpdateTagKeyRequest for
	//  details.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. A purpose denotes that this Tag is intended for use in policies
	//  of a specific policy engine, and will involve that policy engine in
	//  management operations involving this Tag. A purpose does not grant a
	//  policy engine exclusive rights to the Tag, and it may be referenced by
	//  other policy engines.
	//
	//  A purpose cannot be changed once set.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.purpose
	Purpose *string `json:"purpose,omitempty"`

	// Optional. Purpose data corresponds to the policy system that the tag is
	//  intended for. See documentation for `Purpose` for formatting of this field.
	//
	//  Purpose data cannot be changed once set.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.purpose_data
	PurposeData map[string]string `json:"purposeData,omitempty"`
}

// +kcc:proto=google.cloud.resourcemanager.v3.TagKey
type TagKeyObservedState struct {
	// Output only. Immutable. Namespaced name of the TagKey.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.namespaced_name
	NamespacedName *string `json:"namespacedName,omitempty"`

	// Output only. Creation time.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagKey.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
