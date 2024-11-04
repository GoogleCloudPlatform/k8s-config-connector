// Copyright 2024 Google LLC
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

// +kcc:proto=google.cloud.kms.v1.KeyHandle
type KeyHandle struct {
	// Identifier. Name of the [KeyHandle][google.cloud.kms.v1.KeyHandle]
	//  resource, e.g.
	//  `projects/{PROJECT_ID}/locations/{LOCATION}/keyHandles/{KEY_HANDLE_ID}`.
	Name *string `json:"name,omitempty"`

	// Output only. Name of a [CryptoKey][google.cloud.kms.v1.CryptoKey] that has
	//  been provisioned for Customer Managed Encryption Key (CMEK) use in the
	//  [KeyHandle][google.cloud.kms.v1.KeyHandle] project and location for the
	//  requested resource type. The [CryptoKey][google.cloud.kms.v1.CryptoKey]
	//  project will reflect the value configured in the
	//  [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig] on the resource
	//  project's ancestor folder at the time of the
	//  [KeyHandle][google.cloud.kms.v1.KeyHandle] creation. If more than one
	//  ancestor folder has a configured
	//  [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig], the nearest of these
	//  configurations is used.
	KmsKey *string `json:"kmsKey,omitempty"`

	// Required. Indicates the resource type that the resulting
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey] is meant to protect, e.g.
	//  `{SERVICE}.googleapis.com/{TYPE}`. See documentation for supported resource
	//  types.
	ResourceTypeSelector *string `json:"resourceTypeSelector,omitempty"`
}
