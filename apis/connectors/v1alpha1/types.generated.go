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


// +kcc:proto=google.cloud.connectors.v1.Settings
type Settings struct {

	// Optional. Flag indicates whether vpc-sc is enabled.
	// +kcc:proto:field=google.cloud.connectors.v1.Settings.vpcsc
	Vpcsc *bool `json:"vpcsc,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.Settings
type SettingsObservedState struct {
	// Output only. Resource name of the Connection.
	//  Format: projects/{project}/locations/global/settings}
	// +kcc:proto:field=google.cloud.connectors.v1.Settings.name
	Name *string `json:"name,omitempty"`

	// Output only. Flag indicates if user is in PayG model
	// +kcc:proto:field=google.cloud.connectors.v1.Settings.payg
	Payg *bool `json:"payg,omitempty"`
}
