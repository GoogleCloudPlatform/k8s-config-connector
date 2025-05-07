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

var BigLakeDatabaseGVK = GroupVersion.WithKind("BigLakeDatabase")

// BigLakeDatabaseSpec defines the desired state of BigLakeDatabase
// +kcc:proto=google.cloud.bigquery.biglake.v1.Database
type BigLakeDatabaseSpec struct {
	// Required. Defines the parent path of the resource.
	*Parent `json:",inline"`

	// The BigLakeDatabase name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Options of a Hive database.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.hive_options
	HiveOptions *HiveDatabaseOptions `json:"hiveOptions,omitempty"`

	// The database type.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.type
	Type *string `json:"type,omitempty"`
}

// BigLakeDatabaseStatus defines the config connector machine state of BigLakeDatabase
type BigLakeDatabaseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigLakeDatabase resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigLakeDatabaseObservedState `json:"observedState,omitempty"`
}

// BigLakeDatabaseObservedState is the state of the BigLakeDatabase resource as most recently observed in GCP.
// +kcc:proto=google.cloud.bigquery.biglake.v1.Database
type BigLakeDatabaseObservedState struct {
	// Output only. The creation time of the database.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last modification time of the database.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The deletion time of the database. Only set after the database
	//  is deleted.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time when this database is considered expired. Only set
	//  after the database is deleted.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Database.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbiglakedatabase;gcpbiglakedatabases
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigLakeDatabase is the Schema for the BigLakeDatabase API
// +k8s:openapi-gen=true
type BigLakeDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigLakeDatabaseSpec   `json:"spec,omitempty"`
	Status BigLakeDatabaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigLakeDatabaseList contains a list of BigLakeDatabase
type BigLakeDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigLakeDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigLakeDatabase{}, &BigLakeDatabaseList{})
}
