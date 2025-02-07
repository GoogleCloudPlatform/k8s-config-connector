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


// +kcc:proto=google.analytics.admin.v1beta.FirebaseLink
type FirebaseLink struct {

	// Immutable. Firebase project resource name. When creating a FirebaseLink,
	//  you may provide this resource name using either a project number or project
	//  ID. Once this resource has been created, returned FirebaseLinks will always
	//  have a project_name that contains a project number.
	//
	//  Format: 'projects/{project number}'
	//  Example: 'projects/1234'
	// +kcc:proto:field=google.analytics.admin.v1beta.FirebaseLink.project
	Project *string `json:"project,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.FirebaseLink
type FirebaseLinkObservedState struct {
	// Output only. Example format: properties/1234/firebaseLinks/5678
	// +kcc:proto:field=google.analytics.admin.v1beta.FirebaseLink.name
	Name *string `json:"name,omitempty"`

	// Output only. Time when this FirebaseLink was originally created.
	// +kcc:proto:field=google.analytics.admin.v1beta.FirebaseLink.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
