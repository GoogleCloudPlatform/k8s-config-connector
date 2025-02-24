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

// +kcc:proto=google.logging.v2.CmekSettings
type CmekSettings struct {

	// The resource name for the configured Cloud KMS key.
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
	//  cloudkms.cryptoKeyEncrypterDecrypter roles assigned for the key.
	//
	//  The Cloud KMS key used by the Log Router can be updated by changing the
	//  `kms_key_name` to a new valid key name or disabled by setting the key name
	//  to an empty string. Encryption operations that are in progress will be
	//  completed with the key that was in use when they started. Decryption
	//  operations will be completed using the key that was used at the time of
	//  encryption unless access to that key has been revoked.
	//
	//  To disable CMEK for the Log Router, set this field to an empty string.
	//
	//  See [Enabling CMEK for Log
	//  Router](https://cloud.google.com/logging/docs/routing/managed-encryption)
	//  for more information.
	// +kcc:proto:field=google.logging.v2.CmekSettings.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// The CryptoKeyVersion resource name for the configured Cloud KMS key.
	//
	//  KMS key name format:
	//
	//      "projects/[PROJECT_ID]/locations/[LOCATION]/keyRings/[KEYRING]/cryptoKeys/[KEY]/cryptoKeyVersions/[VERSION]"
	//
	//  For example:
	//
	//    `"projects/my-project/locations/us-central1/keyRings/my-ring/cryptoKeys/my-key/cryptoKeyVersions/1"`
	//
	//  This is a read-only field used to convey the specific configured
	//  CryptoKeyVersion of `kms_key` that has been configured. It will be
	//  populated in cases where the CMEK settings are bound to a single key
	//  version.
	//
	//  If this field is populated, the `kms_key` is tied to a specific
	//  CryptoKeyVersion.
	// +kcc:proto:field=google.logging.v2.CmekSettings.kms_key_version_name
	KMSKeyVersionName *string `json:"kmsKeyVersionName,omitempty"`
}

// +kcc:proto=google.logging.v2.IndexConfig
type IndexConfig struct {
	// Required. The LogEntry field path to index.
	//
	//  Note that some paths are automatically indexed, and other paths are not
	//  eligible for indexing. See [indexing documentation](
	//  https://cloud.google.com/logging/docs/view/advanced-queries#indexed-fields)
	//  for details.
	//
	//  For example: `jsonPayload.request.status`
	// +kcc:proto:field=google.logging.v2.IndexConfig.field_path
	FieldPath *string `json:"fieldPath,omitempty"`

	// Required. The type of data in this index.
	// +kcc:proto:field=google.logging.v2.IndexConfig.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.logging.v2.CmekSettings
type CmekSettingsObservedState struct {
	// Output only. The resource name of the CMEK settings.
	// +kcc:proto:field=google.logging.v2.CmekSettings.name
	Name *string `json:"name,omitempty"`

	// Output only. The service account that will be used by the Log Router to
	//  access your Cloud KMS key.
	//
	//  Before enabling CMEK for Log Router, you must first assign the
	//  cloudkms.cryptoKeyEncrypterDecrypter role to the service account that
	//  the Log Router will use to access your Cloud KMS key. Use
	//  [GetCmekSettings][google.logging.v2.ConfigServiceV2.GetCmekSettings] to
	//  obtain the service account ID.
	//
	//  See [Enabling CMEK for Log
	//  Router](https://cloud.google.com/logging/docs/routing/managed-encryption)
	//  for more information.
	// +kcc:proto:field=google.logging.v2.CmekSettings.service_account_id
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

// +kcc:proto=google.logging.v2.IndexConfig
type IndexConfigObservedState struct {
	// Output only. The timestamp when the index was last modified.
	//
	//  This is used to return the timestamp, and will be ignored if supplied
	//  during update.
	// +kcc:proto:field=google.logging.v2.IndexConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
