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

// +kcc:proto=google.cloud.netapp.v1.BackupVault
type BackupVaultObservedState struct {
	// Output only. The backup vault state.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupVault.state
	State *string `json:"state,omitempty"`

	// Output only. Create time of the backup vault.
	// +kcc:proto:field=google.cloud.netapp.v1.BackupVault.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
