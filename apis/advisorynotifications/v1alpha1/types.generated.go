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

// +kcc:proto=google.cloud.advisorynotifications.v1.NotificationSettings
type NotificationSettings struct {
	// Whether the associated NotificationType is enabled.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.NotificationSettings.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.advisorynotifications.v1.Settings
type Settings struct {
	// Identifier. The resource name of the settings to retrieve.
	//  Format:
	//  organizations/{organization}/locations/{location}/settings or
	//  projects/{projects}/locations/{location}/settings.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Settings.name
	Name *string `json:"name,omitempty"`

	// TODO: unsupported map type with key string and value message

	// Required. Fingerprint for optimistic concurrency returned in Get requests.
	//  Must be provided for Update requests. If the value provided does not match
	//  the value known to the server, ABORTED will be thrown, and the client
	//  should retry the read-modify-write cycle.
	// +kcc:proto:field=google.cloud.advisorynotifications.v1.Settings.etag
	Etag *string `json:"etag,omitempty"`
}
