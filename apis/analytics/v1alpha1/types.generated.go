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


// +kcc:proto=google.analytics.admin.v1beta.DataRetentionSettings
type DataRetentionSettings struct {

	// The length of time that event-level data is retained.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataRetentionSettings.event_data_retention
	EventDataRetention *string `json:"eventDataRetention,omitempty"`

	// If true, reset the retention period for the user identifier with every
	//  event from that user.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataRetentionSettings.reset_user_data_on_new_activity
	ResetUserDataOnNewActivity *bool `json:"resetUserDataOnNewActivity,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.DataRetentionSettings
type DataRetentionSettingsObservedState struct {
	// Output only. Resource name for this DataRetentionSetting resource.
	//  Format: properties/{property}/dataRetentionSettings
	// +kcc:proto:field=google.analytics.admin.v1beta.DataRetentionSettings.name
	Name *string `json:"name,omitempty"`
}
