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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var WorkstationsWorkstationClusterGVK = GroupVersion.WithKind("WorkstationsWorkstationCluster")

// Defines the desired state of WorkstationsWorkstationCluster
// +kcc:proto=google.cloud.workstations.v1.WorkstationCluster
type WorkstationsWorkstationClusterSpec struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. The location where the workstation cluster should reside. */
	// +required
	Location string `json:"location"`

	// The WorkstationCluster name. If not given, the metadata.name will be used.
	// + optional
	ResourceID *string `json:"resourceID,omitempty"`

	// // Full name of this workstation cluster.
	// Name *string `json:"name,omitempty"`

	// Optional. Human-readable name for this workstation cluster.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Client-specified annotations.
	Annotations map[string]string `json:"annotations,omitempty"`

	// NOTYET: Not dealing with labels yet
	// // Optional.
	// //  [Labels](https://cloud.google.com/workstations/docs/label-resources) that
	// //  are applied to the workstation cluster and that are also propagated to the
	// //  underlying Compute Engine resources.
	// Labels map[string]string `json:"labels,omitempty"`

	// Immutable. Name of the Compute Engine network in which instances associated
	//  with this workstation cluster will be created.
	// +required
	Network *string `json:"network,omitempty"`

	// Immutable. Name of the Compute Engine subnetwork in which instances
	//  associated with this workstation cluster will be created. Must be part of
	//  the subnetwork specified for this workstation cluster.
	// +required
	Subnetwork *string `json:"subnetwork,omitempty"`

	// Optional. Configuration for private workstation cluster.
	PrivateClusterConfig *WorkstationCluster_PrivateClusterConfig `json:"privateClusterConfig,omitempty"`
}

type WorkstationsWorkstationClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the WorkstationCluster resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *WorkstationsWorkstationClusterObservedState `json:"observedState,omitempty"`
}

// WorkstationsWorkstationClusterSpec defines the desired state of WorkstationCluster
// +kcc:proto=google.cloud.workstations.v1.WorkstationCluster
type WorkstationsWorkstationClusterObservedState struct {
	// Output only. A system-assigned unique identifier for this workstation
	//  cluster.
	Uid *string `json:"uid,omitempty"`

	// NOTYET: This may be better surfaced as status.conditions?
	// // Output only. Indicates whether this workstation cluster is currently being
	// //  updated to match its intended state.
	// Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. Time when this workstation cluster was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when this workstation cluster was most recently updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Time when this workstation cluster was soft-deleted.
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Optional. Checksum computed by the server. May be sent on update and delete
	//  requests to make sure that the client has an up-to-date value before
	//  proceeding.
	Etag *string `json:"etag,omitempty"`

	// NOTYET: This may be better surfaced as status.conditions?
	// // Output only. Whether this workstation cluster is in degraded mode, in which
	// //  case it may require user action to restore full functionality. Details can
	// //  be found in
	// //  [conditions][google.cloud.workstations.v1.WorkstationCluster.conditions].
	// Degraded *bool `json:"degraded,omitempty"`

	// NOTYET: This may be better surfaced as status.conditions?
	// // Output only. Status conditions describing the workstation cluster's current
	// //  state.
	// Conditions []google_rpc_Status `json:"conditions,omitempty"`

	// Output only. The private IP address of the control plane for this
	//  workstation cluster. Workstation VMs need access to this IP address to work
	//  with the service, so make sure that your firewall rules allow egress from
	//  the workstation VMs to this address.
	ControlPlaneIP *string `json:"controlPlaneIP,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpworkstationsworkstationcluster;gcpworkstationsworkstationclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// WorkstationsWorkstationCluster is the Schema for the Workstations WorkstationCluster API
// +k8s:openapi-gen=true
type WorkstationsWorkstationCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   WorkstationsWorkstationClusterSpec   `json:"spec,omitempty"`
	Status WorkstationsWorkstationClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// WorkstationsWorkstationClusterList contains a list of WorkstationsWorkstationCluster
type WorkstationsWorkstationClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkstationsWorkstationCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkstationsWorkstationCluster{}, &WorkstationsWorkstationClusterList{})
}
