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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryReservationCapacityCommitmentGVK = GroupVersion.WithKind("BigQueryReservationCapacityCommitment")

// BigQueryReservationCapacityCommitmentSpec defines the desired state of BigQueryReservationCapacityCommitment
// +kcc:spec:proto=google.cloud.bigquery.reservation.v1.CapacityCommitment
type BigQueryReservationCapacityCommitmentSpec struct {
	// Immutable. The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS.
	// +optional
	Edition *string `json:"edition,omitempty"`

	// Immutable. If true, fail the request if another project in the organization has a capacity commitment.
	// +optional
	EnforceSingleAdminProjectPerOrg *string `json:"enforceSingleAdminProjectPerOrg,omitempty"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// Capacity commitment plan. Valid values are at https://cloud.google.com/bigquery/docs/reference/reservations/rpc/google.cloud.bigquery.reservation.v1#commitmentplan.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.plan
	Plan *string `json:"plan,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// The plan this capacity commitment is converted to after commitmentEndTime passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable some commitment plans.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.renewal_plan
	RenewalPlan *string `json:"renewalPlan,omitempty"`

	// The BigQueryReservationCapacityCommitment name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. Number of slots in this commitment.
	// +required
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.slot_count
	SlotCount *int64 `json:"slotCount,omitempty"`
}

// BigQueryReservationCapacityCommitmentStatus defines the config connector machine state of BigQueryReservationCapacityCommitment
// +kcc:status:proto=google.cloud.bigquery.reservation.v1.CapacityCommitment
type BigQueryReservationCapacityCommitmentStatus struct {
	/* Conditions represent the latest available observations of the
	   BigQueryReservationCapacityCommitment's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.commitment_end_time
	CommitmentEndTime *string `json:"commitmentEndTime,omitempty"`

	// The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.commitment_start_time
	CommitmentStartTime *string `json:"commitmentStartTime,omitempty"`

	// The resource name of the capacity commitment, e.g., projects/myproject/locations/US/capacityCommitments/123.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.name
	Name *string `json:"name,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// State of the commitment.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.CapacityCommitment.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryreservationcapacitycommitment;gcpbigqueryreservationcapacitycommitments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:storageversion

// BigQueryReservationCapacityCommitment is the Schema for the BigQueryReservationCapacityCommitment API
// +k8s:openapi-gen=true
type BigQueryReservationCapacityCommitment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryReservationCapacityCommitmentSpec   `json:"spec,omitempty"`
	Status BigQueryReservationCapacityCommitmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryReservationCapacityCommitmentList contains a list of BigQueryReservationCapacityCommitment
type BigQueryReservationCapacityCommitmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryReservationCapacityCommitment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryReservationCapacityCommitment{}, &BigQueryReservationCapacityCommitmentList{})
}
