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


// +kcc:proto=google.cloud.datacatalog.v1beta1.PolicyTag
type PolicyTag struct {
	// Identifier. Resource name of this policy tag, whose format is:
	//  "projects/{project_number}/locations/{location_id}/taxonomies/{taxonomy_id}/policyTags/{id}".
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.PolicyTag.name
	Name *string `json:"name,omitempty"`

	// Required. User defined name of this policy tag. It must: be unique within
	//  the parent taxonomy; contain only unicode letters, numbers, underscores,
	//  dashes and spaces; not start or end with spaces; and be at most 200 bytes
	//  long when encoded in UTF-8.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.PolicyTag.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Description of this policy tag. It must: contain only unicode characters,
	//  tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
	//  long when encoded in UTF-8. If not set, defaults to an empty description.
	//  If not set, defaults to an empty description.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.PolicyTag.description
	Description *string `json:"description,omitempty"`

	// Resource name of this policy tag's parent policy tag (e.g. for the
	//  "LatLong" policy tag in the example above, this field contains the
	//  resource name of the "Geolocation" policy tag). If empty, it means this
	//  policy tag is a top level policy tag (e.g. this field is empty for the
	//  "Geolocation" policy tag in the example above). If not set, defaults to an
	//  empty string.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.PolicyTag.parent_policy_tag
	ParentPolicyTag *string `json:"parentPolicyTag,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1beta1.PolicyTag
type PolicyTagObservedState struct {
	// Output only. Resource names of child policy tags of this policy tag.
	// +kcc:proto:field=google.cloud.datacatalog.v1beta1.PolicyTag.child_policy_tags
	ChildPolicyTags []string `json:"childPolicyTags,omitempty"`
}
