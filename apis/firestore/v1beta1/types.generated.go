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
// resource: FirestoreIndex:Index

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

// +kcc:proto=google.firestore.admin.v1.Index.IndexField.VectorConfig
type Index_IndexField_VectorConfig struct {
	// Required. The vector dimension this configuration applies to.
	//
	//  The resulting index will only include vectors of this dimension, and
	//  can be used for vector search with the same dimension.
	// +kcc:proto:field=google.firestore.admin.v1.Index.IndexField.VectorConfig.dimension
	Dimension *int32 `json:"dimension,omitempty"`

	// Indicates the vector index is a flat index.
	// +kcc:proto:field=google.firestore.admin.v1.Index.IndexField.VectorConfig.flat
	Flat *Index_IndexField_VectorConfig_FlatIndex `json:"flat,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Index.IndexField.VectorConfig.FlatIndex
type Index_IndexField_VectorConfig_FlatIndex struct {
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
