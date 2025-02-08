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


// +kcc:proto=google.cloud.migrationcenter.v1.ReportConfig
type ReportConfig struct {

	// User-friendly display name. Maximum length is 63 characters.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Free-text description.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.description
	Description *string `json:"description,omitempty"`

	// Required. Collection of combinations of groups and preference sets.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.group_preferenceset_assignments
	GroupPreferencesetAssignments []ReportConfig_GroupPreferenceSetAssignment `json:"groupPreferencesetAssignments,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportConfig.GroupPreferenceSetAssignment
type ReportConfig_GroupPreferenceSetAssignment struct {
	// Required. Name of the group.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.GroupPreferenceSetAssignment.group
	Group *string `json:"group,omitempty"`

	// Required. Name of the Preference Set.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.GroupPreferenceSetAssignment.preference_set
	PreferenceSet *string `json:"preferenceSet,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.ReportConfig
type ReportConfigObservedState struct {
	// Output only. Name of resource.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.name
	Name *string `json:"name,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was last updated.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.ReportConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
