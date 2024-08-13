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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var WorkstationsWorkstationGVK = GroupVersion.WithKind("WorkstationsWorkstation")

// WorkstationsWorkstationSpec defines the desired state of WorkstationsWorkstation
// +kcc:proto=google.cloud.workstations.v1.Workstation
type WorkstationsWorkstationSpec struct {
	// The WorkstationsWorkstation name. If not given, the metadata.name will be used.
	// + optional
	ResourceID *string `json:"resourceID,omitempty"`

	// // Full name of this workstation.
	// Name *string `json:"name,omitempty"`

	// Optional. Human-readable name for this workstation.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Client-specified annotations.
	Annotations map[string]string `json:"annotations,omitempty"`

	// NOTYET: Not dealing with labels yet
	// // Optional.
	// //  [Labels](https://cloud.google.com/workstations/docs/label-resources) that
	// //  are applied to the workstation and that are also propagated to the
	// //  underlying Compute Engine resources.
	// Labels map[string]string `json:"labels,omitempty"`

}

// WorkstationsWorkstationStatus defines the config connector machine state of WorkstationsWorkstation
type WorkstationsWorkstationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the WorkstationsWorkstation resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *WorkstationsWorkstationObservedState `json:"observedState,omitempty"`
}

// WorkstationsWorkstationSpec defines the desired state of WorkstationsWorkstation
// +kcc:proto=google.cloud.workstations.v1.Workstation
type WorkstationsWorkstationObservedState struct {
	// Output only. A system-assigned unique identifier for this workstation.
	Uid *string `json:"uid,omitempty"`

	// NOTYET: This may be better surfaced as status.conditions?
	// // Output only. Indicates whether this workstation is currently being updated
	// //  to match its intended state.
	// Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. Time when this workstation was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when this workstation was most recently updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Time when this workstation was most recently successfully
	//  started, regardless of the workstation's initial state.
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when this workstation was soft-deleted.
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Optional. Checksum computed by the server. May be sent on update and delete
	//  requests to make sure that the client has an up-to-date value before
	//  proceeding.
	Etag *string `json:"etag,omitempty"`

	// Output only. Current state of the workstation.
	State *string `json:"state,omitempty"`

	// Output only. Host to which clients can send HTTPS traffic that will be
	//  received by the workstation. Authorized traffic will be received to the
	//  workstation as HTTP on port 80. To send traffic to a different port,
	//  clients may prefix the host with the destination port in the format
	//  `{port}-{host}`.
	Host *string `json:"host,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// WorkstationsWorkstation is the Schema for the WorkstationsWorkstation API
// +k8s:openapi-gen=true
type WorkstationsWorkstation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkstationsWorkstationSpec   `json:"spec,omitempty"`
	Status WorkstationsWorkstationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// WorkstationsWorkstationList contains a list of WorkstationsWorkstation
type WorkstationsWorkstationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkstationsWorkstation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkstationsWorkstation{}, &WorkstationsWorkstationList{})
}
