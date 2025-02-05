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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SpannerBackupGVK = GroupVersion.WithKind("SpannerBackup")

// SpannerBackupSpec defines the desired state of SpannerBackup
// +kcc:proto=google.spanner.admin.database.v1.Backup
type SpannerBackupSpec struct {
	*InstanceParent `json:",inline"`

	// The SpannerBackup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// reference to the database from which this backup was created. This
	//  needs to be in the same instance as the backup.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.database
	DatabaseRef *DatabaseRef `json:"databaseRef,omitempty"`

	// The backup will contain an externally consistent copy of the database at
	//  the timestamp specified by `version_time`. If `version_time` is not
	//  specified, the system will set `version_time` to the `create_time` of the
	//  backup.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.version_time
	VersionTime *string `json:"versionTime,omitempty"`

	// Required for the
	//  [CreateBackup][google.spanner.admin.database.v1.DatabaseAdmin.CreateBackup]
	//  operation. The expiration time of the backup, with microseconds
	//  granularity that must be at least 6 hours and at most 366 days
	//  from the time the CreateBackup request is processed. Once the `expire_time`
	//  has passed, the backup is eligible to be automatically deleted by Cloud
	//  Spanner to free the resources used by the backup.
	// +kcc:proto:field=google.spanner.admin.database.v1.Backup.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// SpannerBackupStatus defines the config connector machine state of SpannerBackup
type SpannerBackupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SpannerBackup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SpannerBackupObservedState `json:"observedState,omitempty"`
}

// SpannerBackupSpec defines the desired state of SpannerBackup
// +kcc:proto=google.spanner.admin.database.v1.Backup
// SpannerBackupObservedState is the state of the SpannerBackup resource as most recently observed in GCP.
type SpannerBackupObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpspannerbackup;gcpspannerbackups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SpannerBackup is the Schema for the SpannerBackup API
// +k8s:openapi-gen=true
type SpannerBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SpannerBackupSpec   `json:"spec,omitempty"`
	Status SpannerBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SpannerBackupList contains a list of SpannerBackup
type SpannerBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpannerBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpannerBackup{}, &SpannerBackupList{})
}
