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

// +kcc:proto=google.cloud.kms.v1.AutokeyConfig
type AutokeyConfig struct {
	// Identifier. Name of the [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig]
	//  resource, e.g. `folders/{FOLDER_NUMBER}/autokeyConfig`.
	Name *string `json:"name,omitempty"`

	// Optional. Name of the key project, e.g. `projects/{PROJECT_ID}` or
	//  `projects/{PROJECT_NUMBER}`, where Cloud KMS Autokey will provision a new
	//  [CryptoKey][google.cloud.kms.v1.CryptoKey] when a
	//  [KeyHandle][google.cloud.kms.v1.KeyHandle] is created. On
	//  [UpdateAutokeyConfig][google.cloud.kms.v1.AutokeyAdmin.UpdateAutokeyConfig],
	//  the caller will require `cloudkms.cryptoKeys.setIamPolicy` permission on
	//  this key project. Once configured, for Cloud KMS Autokey to function
	//  properly, this key project must have the Cloud KMS API activated and the
	//  Cloud KMS Service Agent for this key project must be granted the
	//  `cloudkms.admin` role (or pertinent permissions). A request with an empty
	//  key project field will clear the configuration.
	KeyProject *string `json:"keyProject,omitempty"`

	// Output only. The state for the AutokeyConfig.
	State *string `json:"state,omitempty"`
}
