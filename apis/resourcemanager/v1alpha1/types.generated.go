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


// +kcc:proto=google.cloud.resourcemanager.v3.TagBinding
type TagBinding struct {

	// The full resource name of the resource the TagValue is bound to.
	//  E.g. `//cloudresourcemanager.googleapis.com/projects/123`
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagBinding.parent
	Parent *string `json:"parent,omitempty"`

	// The TagValue of the TagBinding.
	//  Must be of the form `tagValues/456`.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagBinding.tag_value
	TagValue *string `json:"tagValue,omitempty"`

	// The namespaced name for the TagValue of the TagBinding.
	//  Must be in the format
	//  `{parent_id}/{tag_key_short_name}/{short_name}`.
	//
	//  For methods that support TagValue namespaced name, only one of
	//  tag_value_namespaced_name or tag_value may be filled. Requests with both
	//  fields will be rejected.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagBinding.tag_value_namespaced_name
	TagValueNamespacedName *string `json:"tagValueNamespacedName,omitempty"`
}

// +kcc:proto=google.cloud.resourcemanager.v3.TagBinding
type TagBindingObservedState struct {
	// Output only. The name of the TagBinding. This is a String of the form:
	//  `tagBindings/{full-resource-name}/{tag-value-name}` (e.g.
	//  `tagBindings/%2F%2Fcloudresourcemanager.googleapis.com%2Fprojects%2F123/tagValues/456`).
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagBinding.name
	Name *string `json:"name,omitempty"`
}
