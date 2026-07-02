// Copyright 2026 Google LLC
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
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1alpha1"
	tablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigLakeLockGVK = GroupVersion.WithKind("BigLakeLock")

// BigLakeLockSpec defines the desired state of BigLakeLock
// +kcc:spec:proto=google.cloud.bigquery.biglake.v1alpha1.Lock
type BigLakeLockSpec struct {
	// Required. The parent resource where this lock will be created.
	// Format:
	// projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}/databases/{database_id}
	// +required
	ParentRef *krmv1alpha1.BigQueryBigLakeDatabaseRef `json:"parentDatabaseRef,omitempty"`

	// The BigLakeLock name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The table ID (not fully qualified name) in the same database that the
	// lock will be created on. The table must exist.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1alpha1.Lock.table_id
	TableRef *tablev1beta1.BigQueryBigLakeTableRef `json:"tableRef,omitempty"`

	// The lock type.
	// +kubebuilder:validation:Enum=EXCLUSIVE
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1alpha1.Lock.type
	Type *string `json:"type,omitempty"`
}

// BigLakeLockStatus defines the config connector machine state of BigLakeLock
type BigLakeLockStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigLakeLock resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigLakeLockObservedState `json:"observedState,omitempty"`
}

// BigLakeLockObservedState is the state of the BigLakeLock resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.bigquery.biglake.v1alpha1.Lock
type BigLakeLockObservedState struct {
	// Output only. The resource name.
	// Format:
	// projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}/databases/{database_id}/locks/{lock_id}
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1alpha1.Lock.name
	Name *string `json:"name,omitempty"`

	// Output only. The creation time of the lock.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1alpha1.Lock.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The lock state.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1alpha1.Lock.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbiglakelock;gcpbiglakelocks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigLakeLock is the Schema for the BigLakeLock API
// +k8s:openapi-gen=true
type BigLakeLock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigLakeLockSpec   `json:"spec,omitempty"`
	Status BigLakeLockStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigLakeLockList contains a list of BigLakeLock
type BigLakeLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigLakeLock `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigLakeLock{}, &BigLakeLockList{})
}
