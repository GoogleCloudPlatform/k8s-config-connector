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

// +kcc:proto=google.bigtable.admin.v2.GcRule
type GcRule struct {
	// Delete all cells in a column except the most recent N.
	MaxNumVersions *int32 `json:"maxNumVersions,omitempty"`

	// Delete cells in a column older than the given age.
	//  Values must be at least one millisecond, and will be truncated to
	//  microsecond granularity.
	MaxAge *string `json:"maxAge,omitempty"`

	// Delete cells that would be deleted by every nested rule.
	Intersection *GcRule_Intersection `json:"intersection,omitempty"`

	// Delete cells that would be deleted by any nested rule.
	Union *GcRule_Union `json:"union,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.GcRule.Intersection
type GcRule_Intersection struct {
	// Only delete cells which would be deleted by every element of `rules`.
	Rules []GcRule `json:"rules,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.GcRule.Union
type GcRule_Union struct {
	// Delete cells which would be deleted by any element of `rules`.
	Rules []GcRule `json:"rules,omitempty"`
}
