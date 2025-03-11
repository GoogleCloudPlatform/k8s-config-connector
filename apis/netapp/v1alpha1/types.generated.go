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

// +kcc:proto=google.cloud.netapp.v1.ActiveDirectory
type ActiveDirectory struct {
	// Identifier. The resource name of the active directory.
	//  Format:
	//  `projects/{project_number}/locations/{location_id}/activeDirectories/{active_directory_id}`.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.name
	Name *string `json:"name,omitempty"`

	// Required. Name of the Active Directory domain
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.domain
	Domain *string `json:"domain,omitempty"`

	// The Active Directory site the service will limit Domain Controller
	//  discovery too.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.site
	Site *string `json:"site,omitempty"`

	// Required. Comma separated list of DNS server IP addresses for the Active
	//  Directory domain.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.dns
	DNS *string `json:"dns,omitempty"`

	// Required. NetBIOSPrefix is used as a prefix for SMB server name.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.net_bios_prefix
	NetBiosPrefix *string `json:"netBiosPrefix,omitempty"`

	// The Organizational Unit (OU) within the Windows Active Directory the user
	//  belongs to.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.organizational_unit
	OrganizationalUnit *string `json:"organizationalUnit,omitempty"`

	// If enabled, AES encryption will be enabled for SMB communication.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.aes_encryption
	AesEncryption *bool `json:"aesEncryption,omitempty"`

	// Required. Username of the Active Directory domain administrator.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.username
	Username *string `json:"username,omitempty"`

	// Required. Password of the Active Directory domain administrator.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.password
	Password *string `json:"password,omitempty"`

	// Optional. Users to be added to the Built-in Backup Operator active
	//  directory group.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.backup_operators
	BackupOperators []string `json:"backupOperators,omitempty"`

	// Optional. Users to be added to the Built-in Admininstrators group.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.administrators
	Administrators []string `json:"administrators,omitempty"`

	// Optional. Domain users to be given the SeSecurityPrivilege.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.security_operators
	SecurityOperators []string `json:"securityOperators,omitempty"`

	// Name of the active directory machine. This optional parameter is used only
	//  while creating kerberos volume
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.kdc_hostname
	KdcHostname *string `json:"kdcHostname,omitempty"`

	// KDC server IP address for the active directory machine.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.kdc_ip
	KdcIP *string `json:"kdcIP,omitempty"`

	// If enabled, will allow access to local users and LDAP users. If access is
	//  needed for only LDAP users, it has to be disabled.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.nfs_users_with_ldap
	NfsUsersWithLdap *bool `json:"nfsUsersWithLdap,omitempty"`

	// Description of the active directory.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.description
	Description *string `json:"description,omitempty"`

	// Specifies whether or not the LDAP traffic needs to be signed.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.ldap_signing
	LdapSigning *bool `json:"ldapSigning,omitempty"`

	// If enabled, traffic between the SMB server to Domain Controller (DC) will
	//  be encrypted.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.encrypt_dc_connections
	EncryptDcConnections *bool `json:"encryptDcConnections,omitempty"`

	// Labels for the active directory.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.netapp.v1.ActiveDirectory
type ActiveDirectoryObservedState struct {
	// Output only. Create time of the active directory.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The state of the AD.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.state
	State *string `json:"state,omitempty"`

	// Output only. The state details of the Active Directory.
	// +kcc:proto:field=google.cloud.netapp.v1.ActiveDirectory.state_details
	StateDetails *string `json:"stateDetails,omitempty"`
}
