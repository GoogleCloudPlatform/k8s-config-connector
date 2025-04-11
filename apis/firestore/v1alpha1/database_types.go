// Copyright 2024 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FirestoreDatabaseGVK = GroupVersion.WithKind("FirestoreDatabase")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FirestoreDatabaseSpec defines the desired state of FirestoreDatabase
// +kcc:spec:proto=google.firestore.admin.v1.Database
type FirestoreDatabaseSpec struct {
	// Immutable. The Project that this resource belongs to.
	ProjectRef v1beta1.ProjectRef `json:"projectRef"`

	// The FirestoreDatabase name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The location of the database. Available locations are listed at
	//  https://cloud.google.com/firestore/docs/locations.
	LocationID *string `json:"locationID,omitempty"`

	// The concurrency control mode to use for this database.
	// See https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases#concurrencymode for more info.
	ConcurrencyMode *string `json:"concurrencyMode,omitempty"`

	// Whether to enable the PITR feature on this database.
	// See https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases#pointintimerecoveryenablement for more info.
	PointInTimeRecoveryEnablement *string `json:"pointInTimeRecoveryEnablement,omitempty"`
}

// FirestoreDatabaseStatus defines the config connector machine state of FirestoreDatabase
type FirestoreDatabaseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FirestoreDatabase resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *FirestoreDatabaseObservedState `json:"observedState,omitempty"`
}

// FirestoreDatabaseSpec defines the desired state of FirestoreDatabase
// +kcc:observedstate:proto=google.firestore.admin.v1.Database
type FirestoreDatabaseObservedState struct {
	// Output only. The system-generated UUID4 for this Database.
	Uid *string `json:"uid,omitempty"`

	// Output only. The timestamp at which this database was created. Databases
	//  created before 2016 do not populate create_time.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp at which this database was most recently
	//  updated. Note this only includes updates to the database resource and not
	//  data contained by the database.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The period during which past versions of data are retained in
	//  the database.
	//
	//  Any [read][google.firestore.v1.GetDocumentRequest.read_time]
	//  or [query][google.firestore.v1.ListDocumentsRequest.read_time] can specify
	//  a `read_time` within this window, and will read the state of the database
	//  at that time.
	//
	//  If the PITR feature is enabled, the retention period is 7 days. Otherwise,
	//  the retention period is 1 hour.
	VersionRetentionPeriod *string `json:"versionRetentionPeriod,omitempty"`

	// Output only. The earliest timestamp at which older versions of the data can
	//  be read from the database. See [version_retention_period] above; this field
	//  is populated with `now - version_retention_period`.
	//
	//  This value is continuously updated, and becomes stale the moment it is
	//  queried. If you are using this value to recover data, make sure to account
	//  for the time from the moment when the value is queried to the moment when
	//  you initiate the recovery.
	EarliestVersionTime *string `json:"earliestVersionTime,omitempty"`

	// Output only. The key_prefix for this database. This key_prefix is used, in
	//  combination with the project id ("<key prefix>~<project id>") to construct
	//  the application id that is returned from the Cloud Datastore APIs in Google
	//  App Engine first generation runtimes.
	//
	//  This value may be empty in which case the appid to use for URL-encoded keys
	//  is the project_id (eg: foo instead of v~foo).
	KeyPrefix *string `json:"keyPrefix,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfirestoredatabase;gcpfirestoredatabases
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FirestoreDatabase is the Schema for the FirestoreDatabase API
// +k8s:openapi-gen=true
type FirestoreDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirestoreDatabaseSpec   `json:"spec,omitempty"`
	Status FirestoreDatabaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FirestoreDatabaseList contains a list of FirestoreDatabase
type FirestoreDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirestoreDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FirestoreDatabase{}, &FirestoreDatabaseList{})
}
