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


// +kcc:proto=google.cloud.contactcenterinsights.v1.View
type View struct {
	// Immutable. The resource name of the view.
	//  Format:
	//  projects/{project}/locations/{location}/views/{view}
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.View.name
	Name *string `json:"name,omitempty"`

	// The human-readable display name of the view.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.View.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// String with specific view properties, must be non-empty.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.View.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.contactcenterinsights.v1.View
type ViewObservedState struct {
	// Output only. The time at which this view was created.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.View.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time at which the view was updated.
	// +kcc:proto:field=google.cloud.contactcenterinsights.v1.View.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
