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

package v1beta1

// +kcc:proto=google.bigtable.admin.v2.EncryptionInfo
type EncryptionInfo struct {
	// Output only. The type of encryption used to protect this resource.
	EncryptionType *string `json:"encryptionType,omitempty"`

	/* NOTYET
	// Output only. The status of encrypt/decrypt calls on underlying data for
	//  this resource. Regardless of status, the existing data is always encrypted
	//  at rest.
	EncryptionStatus *google_rpc_Status `json:"encryptionStatus,omitempty"`
	*/

	// Output only. The version of the Cloud KMS key specified in the parent
	//  cluster that is in use for the data underlying this table.
	KmsKeyVersion *string `json:"kmsKeyVersion,omitempty"`
}
