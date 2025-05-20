// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SpannerBackupScheduleGVK = GroupVersion.WithKind("SpannerBackupSchedule")

// SpannerBackupScheduleSpec defines the desired state of SpannerBackupSchedule
// +kcc:spec:proto=google.spanner.admin.database.v1.BackupSchedule
type SpannerBackupScheduleSpec struct {
	// Optional. The schedule specification based on which the backup creations
	//  are triggered.
	// +kcc:proto:field=google.spanner.admin.database.v1.BackupSchedule.spec
	Spec *BackupScheduleSpec `json:"spec,omitempty"`

	// Optional. The retention duration of a backup that must be at least 6 hours
	//  and at most 366 days. The backup is eligible to be automatically deleted
	//  once the retention period has elapsed.
	// +kcc:proto:field=google.spanner.admin.database.v1.BackupSchedule.retention_duration
	RetentionDuration *string `json:"retentionDuration,omitempty"`

	// Optional. The encryption configuration that will be used to encrypt the
	//  backup. If this field is not specified, the backup will use the same
	//  encryption configuration as the database.
	// +kcc:proto:field=google.spanner.admin.database.v1.BackupSchedule.encryption_config
	EncryptionConfig *CreateBackupEncryptionConfig `json:"encryptionConfig,omitempty"`

	// TODO: Confirm whether FullBackupSpec and IncrementalBackupSpec should be references before promoting to beta.

	// The schedule creates only full backups.
	// +kcc:proto:field=google.spanner.admin.database.v1.BackupSchedule.full_backup_spec
	FullBackupSpec *FullBackupSpec `json:"fullBackupSpec,omitempty"`

	// The schedule creates incremental backup chains.
	// +kcc:proto:field=google.spanner.admin.database.v1.BackupSchedule.incremental_backup_spec
	IncrementalBackupSpec *IncrementalBackupSpec `json:"incrementalBackupSpec,omitempty"`

	// Required. The spanner database that this backup applies to.
	DatabaseRef *refs.SpannerDatabaseRef `json:"spannerDatabaseRef,omitempty"`

	// The SpannerBackupSchedule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
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
	KMSKeyRef *kmsv1beta1.KMSKeyRef_OneOf `json:"kmsKeyRef,omitempty"`

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
	KMSKeyRefs []*kmsv1beta1.KMSKeyRef_OneOf `json:"kmsKeyRefs,omitempty"`
}

// SpannerBackupScheduleStatus defines the config connector machine state of SpannerBackupSchedule
type SpannerBackupScheduleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SpannerBackupSchedule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SpannerBackupScheduleObservedState `json:"observedState,omitempty"`
}

// SpannerBackupScheduleObservedState is the state of the SpannerBackupSchedule resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.spanner.admin.database.v1.BackupSchedule
type SpannerBackupScheduleObservedState struct {

	// Output only. The timestamp at which the schedule was last updated.
	//  If the schedule has never been updated, this field contains the timestamp
	//  when the schedule was first created.
	// +kcc:proto:field=google.spanner.admin.database.v1.BackupSchedule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Cron style schedule specification.
	// +kcc:proto:field=google.spanner.admin.database.v1.BackupScheduleSpec.cron_spec
	CronSpec *CrontabSpecObservedState `json:"cronSpec,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpspannerbackupschedule;gcpspannerbackupschedules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SpannerBackupSchedule is the Schema for the SpannerBackupSchedule API
// +k8s:openapi-gen=true
type SpannerBackupSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SpannerBackupScheduleSpec   `json:"spec,omitempty"`
	Status SpannerBackupScheduleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SpannerBackupScheduleList contains a list of SpannerBackupSchedule
type SpannerBackupScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpannerBackupSchedule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpannerBackupSchedule{}, &SpannerBackupScheduleList{})
}
