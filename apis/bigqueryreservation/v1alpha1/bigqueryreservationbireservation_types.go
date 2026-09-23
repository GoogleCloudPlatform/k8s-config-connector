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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryReservationBiReservationGVK = GroupVersion.WithKind("BigQueryReservationBiReservation")

// BigQueryReservationBiReservationSpec defines the desired state of BigQueryReservationBiReservation
// +kcc:spec:proto=google.cloud.bigquery.reservation.v1.BiReservation
type BigQueryReservationBiReservationSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The BigQueryReservationBiReservation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Size of a reservation, in bytes.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.BiReservation.size
	Size *int64 `json:"size,omitempty"`

	// Preferred tables to use BI capacity for.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.BiReservation.preferred_tables
	PreferredTables []TableReference `json:"preferredTables,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.reservation.v1.TableReference
type TableReference struct {
	// The assigned project ID of the project.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.TableReference.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// The ID of the dataset in the above project.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.TableReference.dataset_id
	DatasetID *string `json:"datasetID,omitempty"`

	// The ID of the table in the above dataset.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.TableReference.table_id
	TableID *string `json:"tableID,omitempty"`
}

// BigQueryReservationBiReservationStatus defines the config connector machine state of BigQueryReservationBiReservation
type BigQueryReservationBiReservationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryReservationBiReservation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryReservationBiReservationObservedState `json:"observedState,omitempty"`
}

// BigQueryReservationBiReservationObservedState is the state of the BigQueryReservationBiReservation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.bigquery.reservation.v1.BiReservation
type BigQueryReservationBiReservationObservedState struct {
	// Output only. The last update timestamp of a reservation.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.BiReservation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryreservationbireservation;gcpbigqueryreservationbireservations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryReservationBiReservation is the Schema for the BigQueryReservationBiReservation API
// +k8s:openapi-gen=true
type BigQueryReservationBiReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryReservationBiReservationSpec   `json:"spec,omitempty"`
	Status BigQueryReservationBiReservationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryReservationBiReservationList contains a list of BigQueryReservationBiReservation
type BigQueryReservationBiReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryReservationBiReservation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryReservationBiReservation{}, &BigQueryReservationBiReservationList{})
}
