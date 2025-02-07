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


// +kcc:proto=google.api.cloudquotas.v1beta.QuotaAdjusterSettings
type QuotaAdjusterSettings struct {
	// Identifier. Name of the config would be of the format:
	//    projects/12345/locations/global/quotaAdjusterSettings
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaAdjusterSettings.name
	Name *string `json:"name,omitempty"`

	// Required. The configured value of the enablement at the given resource.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaAdjusterSettings.enablement
	Enablement *string `json:"enablement,omitempty"`

	// Optional. The current etag of the QuotaAdjusterSettings. If an etag is
	//  provided on update and does not match the current server's etag of the
	//  QuotaAdjusterSettings, the request will be blocked and an ABORTED error
	//  will be returned. See https://google.aip.dev/134#etags for more details on
	//  etags.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaAdjusterSettings.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.api.cloudquotas.v1beta.QuotaAdjusterSettings
type QuotaAdjusterSettingsObservedState struct {
	// Output only. The timestamp when the QuotaAdjusterSettings was last updated.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaAdjusterSettings.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
