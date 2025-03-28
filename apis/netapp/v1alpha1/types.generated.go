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

// +kcc:proto=google.cloud.netapp.v1.BackupPolicy
type BackupPolicy struct {
	// Identifier. The resource name of the backup policy.
	//  Format:
	//  `projects/{project_id}/locations/{location}/backupPolicies/{backup_policy_id}`.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.name
	Name *string `json:"name,omitempty"`

	// Number of daily backups to keep. Note that the minimum daily backup limit
	//  is 2.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.daily_backup_limit
	DailyBackupLimit *int32 `json:"dailyBackupLimit,omitempty"`

	// Number of weekly backups to keep. Note that the sum of daily, weekly and
	//  monthly backups should be greater than 1.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.weekly_backup_limit
	WeeklyBackupLimit *int32 `json:"weeklyBackupLimit,omitempty"`

	// Number of monthly backups to keep. Note that the sum of daily, weekly and
	//  monthly backups should be greater than 1.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.monthly_backup_limit
	MonthlyBackupLimit *int32 `json:"monthlyBackupLimit,omitempty"`

	// Description of the backup policy.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.description
	Description *string `json:"description,omitempty"`

	// If enabled, make backups automatically according to the schedules.
	//  This will be applied to all volumes that have this policy attached and
	//  enforced on volume level. If not specified, default is true.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.BackupPolicy
type BackupPolicyObservedState struct {
	// Output only. The total number of volumes assigned by this backup policy.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.assigned_volume_count
	AssignedVolumeCount *int32 `json:"assignedVolumeCount,omitempty"`

	// Output only. The time when the backup policy was created.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The backup policy state.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupPolicy.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.BackupVault
type BackupVault struct {
	// Identifier. The resource name of the backup vault.
	//  Format:
	//  `projects/{project_id}/locations/{location}/backupVaults/{backup_vault_id}`.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupVault.name
	Name *string `json:"name,omitempty"`

	// Description of the backup vault.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupVault.description
	Description *string `json:"description,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupVault.labels
	Labels map[string]string `json:"labels,omitempty"`
}
