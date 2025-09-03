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

// +generated:types
// krm.group: firestore.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.firestore.admin.v1
// resource: FirestoreDatabase:Database

package v1beta1

// +kcc:proto=google.firestore.admin.v1.Database.CmekConfig
type Database_CmekConfig struct {
	// Required. Only keys in the same location as this database are allowed to
	//  be used for encryption.
	//
	//  For Firestore's nam5 multi-region, this corresponds to Cloud KMS
	//  multi-region us. For Firestore's eur3 multi-region, this corresponds to
	//  Cloud KMS multi-region europe. See
	//  https://cloud.google.com/kms/docs/locations.
	//
	//  The expected format is
	//  `projects/{project_id}/locations/{kms_location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	// +kcc:proto:field=google.firestore.admin.v1.Database.CmekConfig.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Database.SourceInfo
type Database_SourceInfo struct {
	// If set, this database was restored from the specified backup (or a
	//  snapshot thereof).
	// +kcc:proto:field=google.firestore.admin.v1.Database.SourceInfo.backup
	Backup *Database_SourceInfo_BackupSource `json:"backup,omitempty"`

	// The associated long-running operation. This field may not be set after
	//  the operation has completed. Format:
	//  `projects/{project}/databases/{database}/operations/{operation}`.
	// +kcc:proto:field=google.firestore.admin.v1.Database.SourceInfo.operation
	Operation *string `json:"operation,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Database.SourceInfo.BackupSource
type Database_SourceInfo_BackupSource struct {
	// The resource name of the backup that was used to restore this
	//  database. Format:
	//  `projects/{project}/locations/{location}/backups/{backup}`.
	// +kcc:proto:field=google.firestore.admin.v1.Database.SourceInfo.BackupSource.backup
	Backup *string `json:"backup,omitempty"`
}

// +kcc:observedstate:proto=google.firestore.admin.v1.Database.CmekConfig
type Database_CmekConfigObservedState struct {
	// Output only. Currently in-use [KMS key
	//  versions](https://cloud.google.com/kms/docs/resource-hierarchy#key_versions).
	//  During [key rotation](https://cloud.google.com/kms/docs/key-rotation),
	//  there can be multiple in-use key versions.
	//
	//  The expected format is
	//  `projects/{project_id}/locations/{kms_location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{key_version}`.
	// +kcc:proto:field=google.firestore.admin.v1.Database.CmekConfig.active_key_version
	ActiveKeyVersion []string `json:"activeKeyVersion,omitempty"`
}
