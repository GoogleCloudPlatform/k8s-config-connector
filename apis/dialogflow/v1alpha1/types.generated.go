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


// +kcc:proto=google.cloud.dialogflow.v2beta1.EncryptionSpec
type EncryptionSpec struct {
	// Immutable. The resource name of the encryption key specification resource.
	//  Format:
	//  projects/{project}/locations/{location}/encryptionSpec
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.EncryptionSpec.name
	Name *string `json:"name,omitempty"`

	// Required. The name of customer-managed encryption key that is used to
	//  secure a resource and its sub-resources. If empty, the resource is secured
	//  by the default Google encryption key. Only the key in the same location as
	//  this resource is allowed to be used for encryption. Format:
	//  `projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{key}`
	// +kcc:proto:field=google.cloud.dialogflow.v2beta1.EncryptionSpec.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}
