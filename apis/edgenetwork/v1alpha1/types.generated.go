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


// +kcc:proto=google.cloud.edgenetwork.v1.Zone
type Zone struct {
	// Required. The resource name of the zone.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Zone.name
	Name *string `json:"name,omitempty"`

	// Deprecated: not implemented.
	//  Labels as key value pairs.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Zone.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Deprecated: not implemented.
	//  The deployment layout type.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Zone.layout_name
	LayoutName *string `json:"layoutName,omitempty"`
}

// +kcc:proto=google.cloud.edgenetwork.v1.Zone
type ZoneObservedState struct {
	// Output only. The time when the zone was created.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Zone.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the zone was last updated.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Zone.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
