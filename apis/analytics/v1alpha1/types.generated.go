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


// +kcc:proto=google.analytics.admin.v1beta.MeasurementProtocolSecret
type MeasurementProtocolSecret struct {

	// Required. Human-readable display name for this secret.
	// +kcc:proto:field=google.analytics.admin.v1beta.MeasurementProtocolSecret.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.MeasurementProtocolSecret
type MeasurementProtocolSecretObservedState struct {
	// Output only. Resource name of this secret. This secret may be a child of
	//  any type of stream. Format:
	//  properties/{property}/dataStreams/{dataStream}/measurementProtocolSecrets/{measurementProtocolSecret}
	// +kcc:proto:field=google.analytics.admin.v1beta.MeasurementProtocolSecret.name
	Name *string `json:"name,omitempty"`

	// Output only. The measurement protocol secret value. Pass this value to the
	//  api_secret field of the Measurement Protocol API when sending hits to this
	//  secret's parent property.
	// +kcc:proto:field=google.analytics.admin.v1beta.MeasurementProtocolSecret.secret_value
	SecretValue *string `json:"secretValue,omitempty"`
}
