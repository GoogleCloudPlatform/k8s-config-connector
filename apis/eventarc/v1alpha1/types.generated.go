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


// +kcc:proto=google.cloud.eventarc.v1.GoogleApiSource
type GoogleApiSource struct {
	// Identifier. Resource name of the form
	//  projects/{project}/locations/{location}/googleApiSources/{google_api_source}
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleApiSource.name
	Name *string `json:"name,omitempty"`

	// Optional. Resource labels.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleApiSource.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Resource annotations.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleApiSource.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Resource display name.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleApiSource.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Destination is the message bus that the GoogleApiSource is
	//  delivering to. It must be point to the full resource name of a MessageBus.
	//  Format:
	//  "projects/{PROJECT_ID}/locations/{region}/messagesBuses/{MESSAGE_BUS_ID)
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleApiSource.destination
	Destination *string `json:"destination,omitempty"`

	// Optional. Resource name of a KMS crypto key (managed by the user) used to
	//  encrypt/decrypt their event data.
	//
	//  It must match the pattern
	//  `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleApiSource.crypto_key_name
	CryptoKeyName *string `json:"cryptoKeyName,omitempty"`

	// Optional. Config to control Platform logging for the GoogleApiSource.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleApiSource.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.LoggingConfig
type LoggingConfig struct {
	// Optional. The minimum severity of logs that will be sent to
	//  Stackdriver/Platform Telemetry. Logs at severitiy ≥ this value will be
	//  sent, unless it is NONE.
	// +kcc:proto:field=google.cloud.eventarc.v1.LoggingConfig.log_severity
	LogSeverity *string `json:"logSeverity,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.GoogleApiSource
type GoogleApiSourceObservedState struct {
	// Output only. Server assigned unique identifier for the channel. The value
	//  is a UUID4 string and guaranteed to remain unchanged until the resource is
	//  deleted.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleApiSource.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields, and might be sent only on update and delete requests to
	//  ensure that the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleApiSource.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleApiSource.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.eventarc.v1.GoogleApiSource.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
