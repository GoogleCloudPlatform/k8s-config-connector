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

var BigQueryReservationReservationGroupGVK = GroupVersion.WithKind("BigQueryReservationReservationGroup")

// BigQueryReservationReservationGroupSpec defines the desired state of BigQueryReservationReservationGroup
// +kcc:spec:proto=google.cloud.bigquery.reservation.v1.ReservationGroup
type BigQueryReservationReservationGroupSpec struct {
	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Optional. The parent reservation group of the reservation group.
	// +optional
	ParentGroupRef *BigQueryReservationReservationGroupRef `json:"parentGroupRef,omitempty"`

	// The BigQueryReservationReservationGroup name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// BigQueryReservationReservationGroupStatus defines the config connector machine state of BigQueryReservationReservationGroup
// +kcc:status:proto=google.cloud.bigquery.reservation.v1.ReservationGroup
type BigQueryReservationReservationGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   BigQueryReservationReservationGroup's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *BigQueryReservationReservationGroupObservedState `json:"observedState,omitempty"`

	// ExternalRef is a unique specifier for the BigQueryReservationReservationGroup resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`
}

// BigQueryReservationReservationGroupObservedState is the state of the BigQueryReservationReservationGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.bigquery.reservation.v1.ReservationGroup
type BigQueryReservationReservationGroupObservedState struct {
	// Output only. Creation time of the reservation group.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.ReservationGroup.creation_time
	CreationTime *string `json:"creationTime,omitempty"`

	// Output only. Last update time of the reservation group.
	// +optional
	// +kcc:proto:field=google.cloud.bigquery.reservation.v1.ReservationGroup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryreservationreservationgroup;gcpbigqueryreservationreservationgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:storageversion

// BigQueryReservationReservationGroup is the Schema for the BigQueryReservationReservationGroup API
// +k8s:openapi-gen=true
type BigQueryReservationReservationGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigQueryReservationReservationGroupSpec   `json:"spec,omitempty"`
	Status BigQueryReservationReservationGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryReservationReservationGroupList contains a list of BigQueryReservationReservationGroup
type BigQueryReservationReservationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryReservationReservationGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryReservationReservationGroup{}, &BigQueryReservationReservationGroupList{})
}
