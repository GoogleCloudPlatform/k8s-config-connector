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


// +kcc:proto=google.ai.generativelanguage.v1beta3.Permission
type Permission struct {

	// Required. Immutable. The type of the grantee.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.Permission.grantee_type
	GranteeType *string `json:"granteeType,omitempty"`

	// Optional. Immutable. The email address of the user of group which this
	//  permission refers. Field is not set when permission's grantee type is
	//  EVERYONE.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.Permission.email_address
	EmailAddress *string `json:"emailAddress,omitempty"`

	// Required. The role granted by this permission.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.Permission.role
	Role *string `json:"role,omitempty"`
}

// +kcc:proto=google.ai.generativelanguage.v1beta3.Permission
type PermissionObservedState struct {
	// Output only. The permission name. A unique name will be generated on
	//  create. Example: tunedModels/{tuned_model}permssions/{permission} Output
	//  only.
	// +kcc:proto:field=google.ai.generativelanguage.v1beta3.Permission.name
	Name *string `json:"name,omitempty"`
}
