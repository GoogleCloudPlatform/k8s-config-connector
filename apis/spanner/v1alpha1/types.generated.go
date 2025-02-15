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

// +kcc:proto=google.spanner.admin.database.v1.BackupScheduleSpec
type BackupScheduleSpec struct {
	// Cron style schedule specification.
	// +kcc:proto:field=google.spanner.admin.database.v1.BackupScheduleSpec.cron_spec
	CronSpec *CrontabSpec `json:"cronSpec,omitempty"`
}

// +kcc:proto=google.spanner.admin.database.v1.CreateBackupEncryptionConfig
type CreateBackupEncryptionConfig struct {
	// Required. The encryption type of the backup.
	// +kcc:proto:field=google.spanner.admin.database.v1.CreateBackupEncryptionConfig.encryption_type
	EncryptionType *string `json:"encryptionType,omitempty"`

	// Optional. The Cloud KMS key that will be used to protect the backup.
	//  This field should be set only when
	//  [encryption_type][google.spanner.admin.database.v1.CreateBackupEncryptionConfig.encryption_type]
	//  is `CUSTOMER_MANAGED_ENCRYPTION`. Values are of the form
	//  `projects/<project>/locations/<location>/keyRings/<key_ring>/cryptoKeys/<kms_key_name>`.
	// +kcc:proto:field=google.spanner.admin.database.v1.CreateBackupEncryptionConfig.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Optional. Specifies the KMS configuration for the one or more keys used to
	//  protect the backup. Values are of the form
	//  `projects/<project>/locations/<location>/keyRings/<key_ring>/cryptoKeys/<kms_key_name>`.
	//
	//  The keys referenced by kms_key_names must fully cover all
	//  regions of the backup's instance configuration. Some examples:
	//  * For single region instance configs, specify a single regional
	//  location KMS key.
	//  * For multi-regional instance configs of type GOOGLE_MANAGED,
	//  either specify a multi-regional location KMS key or multiple regional
	//  location KMS keys that cover all regions in the instance config.
	//  * For an instance config of type USER_MANAGED, please specify only
	//  regional location KMS keys to cover each region in the instance config.
	//  Multi-regional location KMS keys are not supported for USER_MANAGED
	//  instance configs.
	// +kcc:proto:field=google.spanner.admin.database.v1.CreateBackupEncryptionConfig.kms_key_names
	KMSKeyNames []string `json:"kmsKeyNames,omitempty"`
}

// +kcc:proto=google.spanner.admin.database.v1.CrontabSpec
type CrontabSpec struct {
	// Required. Textual representation of the crontab. User can customize the
	//  backup frequency and the backup version time using the cron
	//  expression. The version time must be in UTC timzeone.
	//
	//  The backup will contain an externally consistent copy of the
	//  database at the version time. Allowed frequencies are 12 hour, 1 day,
	//  1 week and 1 month. Examples of valid cron specifications:
	//    * `0 2/12 * * * ` : every 12 hours at (2, 14) hours past midnight in UTC.
	//    * `0 2,14 * * * ` : every 12 hours at (2,14) hours past midnight in UTC.
	//    * `0 2 * * * `    : once a day at 2 past midnight in UTC.
	//    * `0 2 * * 0 `    : once a week every Sunday at 2 past midnight in UTC.
	//    * `0 2 8 * * `    : once a month on 8th day at 2 past midnight in UTC.
	// +kcc:proto:field=google.spanner.admin.database.v1.CrontabSpec.text
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.spanner.admin.database.v1.FullBackupSpec
type FullBackupSpec struct {
}

// +kcc:proto=google.spanner.admin.database.v1.IncrementalBackupSpec
type IncrementalBackupSpec struct {
}

// +kcc:proto=google.spanner.admin.database.v1.BackupScheduleSpec
type BackupScheduleSpecObservedState struct {
	// Cron style schedule specification.
	// +kcc:proto:field=google.spanner.admin.database.v1.BackupScheduleSpec.cron_spec
	CronSpec *CrontabSpecObservedState `json:"cronSpec,omitempty"`
}

// +kcc:proto=google.spanner.admin.database.v1.CrontabSpec
type CrontabSpecObservedState struct {
	// Output only. The time zone of the times in `CrontabSpec.text`. Currently
	//  only UTC is supported.
	// +kcc:proto:field=google.spanner.admin.database.v1.CrontabSpec.time_zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Output only. Schedule backups will contain an externally consistent copy
	//  of the database at the version time specified in
	//  `schedule_spec.cron_spec`. However, Spanner may not initiate the creation
	//  of the scheduled backups at that version time. Spanner will initiate
	//  the creation of scheduled backups within the time window bounded by the
	//  version_time specified in `schedule_spec.cron_spec` and version_time +
	//  `creation_window`.
	// +kcc:proto:field=google.spanner.admin.database.v1.CrontabSpec.creation_window
	CreationWindow *string `json:"creationWindow,omitempty"`
}

// +kcc:proto=google.spanner.admin.instance.v1.ReplicaInfo
type ReplicaInfo struct {
	// The location of the serving resources, e.g., "us-central1".
	// +kcc:proto:field=google.spanner.admin.instance.v1.ReplicaInfo.location
	Location *string `json:"location,omitempty"`

	// The type of replica.
	// +kcc:proto:field=google.spanner.admin.instance.v1.ReplicaInfo.type
	Type *string `json:"type,omitempty"`

	// If true, this location is designated as the default leader location where
	//  leader replicas are placed. See the [region types
	//  documentation](https://cloud.google.com/spanner/docs/instances#region_types)
	//  for more details.
	// +kcc:proto:field=google.spanner.admin.instance.v1.ReplicaInfo.default_leader_location
	DefaultLeaderLocation *bool `json:"defaultLeaderLocation,omitempty"`
}
