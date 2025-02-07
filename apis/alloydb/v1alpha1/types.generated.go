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


// +kcc:proto=google.cloud.alloydb.v1.User
type User struct {

	// Input only. Password for the user.
	// +kcc:proto:field=google.cloud.alloydb.v1.User.password
	Password *string `json:"password,omitempty"`

	// Optional. List of database roles this user has.
	//  The database role strings are subject to the PostgreSQL naming conventions.
	// +kcc:proto:field=google.cloud.alloydb.v1.User.database_roles
	DatabaseRoles []string `json:"databaseRoles,omitempty"`

	// Optional. Type of this user.
	// +kcc:proto:field=google.cloud.alloydb.v1.User.user_type
	UserType *string `json:"userType,omitempty"`

	// Input only. If the user already exists and it has additional roles, keep
	//  them granted.
	// +kcc:proto:field=google.cloud.alloydb.v1.User.keep_extra_roles
	KeepExtraRoles *bool `json:"keepExtraRoles,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1.User
type UserObservedState struct {
	// Output only. Name of the resource in the form of
	//  projects/{project}/locations/{location}/cluster/{cluster}/users/{user}.
	// +kcc:proto:field=google.cloud.alloydb.v1.User.name
	Name *string `json:"name,omitempty"`
}
