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

// +kcc:proto=google.logging.v2.Settings
type Settings struct {

	// Optional. The resource name for the configured Cloud KMS key.
	//
	//  KMS key name format:
	//
	//      "projects/[PROJECT_ID]/locations/[LOCATION]/keyRings/[KEYRING]/cryptoKeys/[KEY]"
	//
	//  For example:
	//
	//    `"projects/my-project/locations/us-central1/keyRings/my-ring/cryptoKeys/my-key"`
	//
	//
	//
	//  To enable CMEK for the Log Router, set this field to a valid
	//  `kms_key_name` for which the associated service account has the required
	//  `roles/cloudkms.cryptoKeyEncrypterDecrypter` role assigned for the key.
	//
	//  The Cloud KMS key used by the Log Router can be updated by changing the
	//  `kms_key_name` to a new valid key name. Encryption operations that are in
	//  progress will be completed with the key that was in use when they started.
	//  Decryption operations will be completed using the key that was used at the
	//  time of encryption unless access to that key has been revoked.
	//
	//  To disable CMEK for the Log Router, set this field to an empty string.
	//
	//  See [Enabling CMEK for Log
	//  Router](https://cloud.google.com/logging/docs/routing/managed-encryption)
	//  for more information.
	// +kcc:proto:field=google.logging.v2.Settings.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Optional. The Cloud region that will be used for _Default and _Required log
	//  buckets for newly created projects and folders. For example `europe-west1`.
	//  This setting does not affect the location of custom log buckets.
	// +kcc:proto:field=google.logging.v2.Settings.storage_location
	StorageLocation *string `json:"storageLocation,omitempty"`

	// Optional. If set to true, the _Default sink in newly created projects and
	//  folders will created in a disabled state. This can be used to automatically
	//  disable log ingestion if there is already an aggregated sink configured in
	//  the hierarchy. The _Default sink can be re-enabled manually if needed.
	// +kcc:proto:field=google.logging.v2.Settings.disable_default_sink
	DisableDefaultSink *bool `json:"disableDefaultSink,omitempty"`
}

// +kcc:proto=google.logging.v2.Settings
type SettingsObservedState struct {
	// Output only. The resource name of the settings.
	// +kcc:proto:field=google.logging.v2.Settings.name
	Name *string `json:"name,omitempty"`

	// Output only. The service account that will be used by the Log Router to
	//  access your Cloud KMS key.
	//
	//  Before enabling CMEK for Log Router, you must first assign the role
	//  `roles/cloudkms.cryptoKeyEncrypterDecrypter` to the service account that
	//  the Log Router will use to access your Cloud KMS key. Use
	//  [GetSettings][google.logging.v2.ConfigServiceV2.GetSettings] to
	//  obtain the service account ID.
	//
	//  See [Enabling CMEK for Log
	//  Router](https://cloud.google.com/logging/docs/routing/managed-encryption)
	//  for more information.
	// +kcc:proto:field=google.logging.v2.Settings.kms_service_account_id
	KMSServiceAccountID *string `json:"kmsServiceAccountID,omitempty"`
}
