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


// +kcc:proto=google.cloud.backupdr.v1.BackupVault
type BackupVault struct {

	// Optional. The description of the BackupVault instance (2048 characters or
	//  less).
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.description
	Description *string `json:"description,omitempty"`

	// Optional. Resource labels to represent user provided metadata.
	//  No labels currently defined:
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The default and minimum enforced retention for each backup within
	//  the backup vault.  The enforced retention for each backup can be extended.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.backup_minimum_enforced_retention_duration
	BackupMinimumEnforcedRetentionDuration *string `json:"backupMinimumEnforcedRetentionDuration,omitempty"`

	// Optional. Server specified ETag for the backup vault resource to
	//  prevent simultaneous updates from overwiting each other.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Time after which the BackupVault resource is locked.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.effective_time
	EffectiveTime *string `json:"effectiveTime,omitempty"`

	// Optional. User annotations. See https://google.aip.dev/128#annotations
	//  Stores small amounts of arbitrary data.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Note: This field is added for future use case and will not be
	//  supported in the current release.
	//
	//  Access restriction for the backup vault.
	//  Default value is WITHIN_ORGANIZATION if not provided during creation.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.access_restriction
	AccessRestriction *string `json:"accessRestriction,omitempty"`
}

// +kcc:proto=google.cloud.backupdr.v1.BackupVault
type BackupVaultObservedState struct {
	// Output only. Identifier. Name of the backup vault to create. It must have
	//  the
	//  format`"projects/{project}/locations/{location}/backupVaults/{backupvault}"`.
	//  `{backupvault}` cannot be changed after creation. It must be between 3-63
	//  characters long and must be unique within the project and location.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.name
	Name *string `json:"name,omitempty"`

	// Output only. The time when the instance was created.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the instance was updated.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Set to true when there are no backups nested under this
	//  resource.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.deletable
	Deletable *bool `json:"deletable,omitempty"`

	// Output only. The BackupVault resource instance state.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.state
	State *string `json:"state,omitempty"`

	// Output only. The number of backups in this backup vault.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.backup_count
	BackupCount *int64 `json:"backupCount,omitempty"`

	// Output only. Service account used by the BackupVault Service for this
	//  BackupVault.  The user should grant this account permissions in their
	//  workload project to enable the service to run backups and restores there.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. Total size of the storage used by all backup resources.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.total_stored_bytes
	TotalStoredBytes *int64 `json:"totalStoredBytes,omitempty"`

	// Output only. Immutable after resource creation until resource deletion.
	// +kcc:proto:field=google.cloud.backupdr.v1.BackupVault.uid
	Uid *string `json:"uid,omitempty"`
}
