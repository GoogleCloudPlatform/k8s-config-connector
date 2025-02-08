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


// +kcc:proto=google.cloud.datacatalog.v1.PolicyTag
type PolicyTag struct {
	// Identifier. Resource name of this policy tag in the URL format.
	//
	//  The policy tag manager generates unique taxonomy IDs and policy tag IDs.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PolicyTag.name
	Name *string `json:"name,omitempty"`

	// Required. User-defined name of this policy tag.
	//
	//  The name can't start or end with spaces and must be unique within the
	//  parent taxonomy, contain only Unicode letters, numbers, underscores, dashes
	//  and spaces, and be at most 200 bytes long when encoded in UTF-8.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PolicyTag.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Description of this policy tag. If not set, defaults to empty.
	//
	//  The description must contain only Unicode characters,
	//  tabs, newlines, carriage returns and page breaks, and be at most 2000 bytes
	//  long when encoded in UTF-8.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PolicyTag.description
	Description *string `json:"description,omitempty"`

	// Resource name of this policy tag's parent policy tag. If empty, this is a
	//  top level tag. If not set, defaults to an empty string.
	//
	//  For example, for the "LatLong" policy tag in the example above, this field
	//  contains the resource name of the "Geolocation" policy tag, and, for
	//  "Geolocation", this field is empty.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PolicyTag.parent_policy_tag
	ParentPolicyTag *string `json:"parentPolicyTag,omitempty"`
}

// +kcc:proto=google.cloud.datacatalog.v1.PolicyTag
type PolicyTagObservedState struct {
	// Output only. Resource names of child policy tags of this policy tag.
	// +kcc:proto:field=google.cloud.datacatalog.v1.PolicyTag.child_policy_tags
	ChildPolicyTags []string `json:"childPolicyTags,omitempty"`
}
