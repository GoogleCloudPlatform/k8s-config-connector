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


// +kcc:proto=google.firestore.admin.v1.Backup
type Backup struct {
}

// +kcc:proto=google.firestore.admin.v1.Backup.Stats
type Backup_Stats struct {
}

// +kcc:proto=google.firestore.admin.v1.Backup
type BackupObservedState struct {
	// Output only. The unique resource name of the Backup.
	//
	//  Format is `projects/{project}/locations/{location}/backups/{backup}`.
	// +kcc:proto:field=google.firestore.admin.v1.Backup.name
	Name *string `json:"name,omitempty"`

	// Output only. Name of the Firestore database that the backup is from.
	//
	//  Format is `projects/{project}/databases/{database}`.
	// +kcc:proto:field=google.firestore.admin.v1.Backup.database
	Database *string `json:"database,omitempty"`

	// Output only. The system-generated UUID4 for the Firestore database that the
	//  backup is from.
	// +kcc:proto:field=google.firestore.admin.v1.Backup.database_uid
	DatabaseUid *string `json:"databaseUid,omitempty"`

	// Output only. The backup contains an externally consistent copy of the
	//  database at this time.
	// +kcc:proto:field=google.firestore.admin.v1.Backup.snapshot_time
	SnapshotTime *string `json:"snapshotTime,omitempty"`

	// Output only. The timestamp at which this backup expires.
	// +kcc:proto:field=google.firestore.admin.v1.Backup.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Statistics about the backup.
	//
	//  This data only becomes available after the backup is fully materialized to
	//  secondary storage. This field will be empty till then.
	// +kcc:proto:field=google.firestore.admin.v1.Backup.stats
	Stats *Backup_Stats `json:"stats,omitempty"`

	// Output only. The current state of the backup.
	// +kcc:proto:field=google.firestore.admin.v1.Backup.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.firestore.admin.v1.Backup.Stats
type Backup_StatsObservedState struct {
	// Output only. Summation of the size of all documents and index entries in
	//  the backup, measured in bytes.
	// +kcc:proto:field=google.firestore.admin.v1.Backup.Stats.size_bytes
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	// Output only. The total number of documents contained in the backup.
	// +kcc:proto:field=google.firestore.admin.v1.Backup.Stats.document_count
	DocumentCount *int64 `json:"documentCount,omitempty"`

	// Output only. The total number of index entries contained in the backup.
	// +kcc:proto:field=google.firestore.admin.v1.Backup.Stats.index_count
	IndexCount *int64 `json:"indexCount,omitempty"`
}
