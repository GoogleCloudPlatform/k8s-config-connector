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

// API sources for NetappSnapshot, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e09e85d32ca349e1b205514a412817a7692595c6/google/cloud/netapp/v1/snapshot.proto
// +kcc:source:service-docs=https://cloud.google.com/netapp/
// +kcc:source:resource-docs=https://docs.cloud.google.com/netapp/volumes/docs/reference/rest/v1/projects.locations.volumes.snapshots

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetappSnapshotGVK = GroupVersion.WithKind("NetappSnapshot")

// NetappSnapshotSpec defines the desired state of NetappSnapshot
// +kcc:spec:proto=google.cloud.netapp.v1.Snapshot
// +kcc:required-from-proto
type NetappSnapshotSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetappSnapshot name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// A description of the snapshot with 2048 characters or less.
	//  Requests with longer descriptions will be rejected.
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.description
	Description *string `json:"description,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// NetappSnapshotStatus defines the config connector machine state of NetappSnapshot
type NetappSnapshotStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetappSnapshot resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetappSnapshotObservedState `json:"observedState,omitempty"`
}

// NetappSnapshotObservedState is the state of the NetappSnapshot resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.netapp.v1.Snapshot
type NetappSnapshotObservedState struct {
	// Output only. The snapshot state.
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.state
	State *string `json:"state,omitempty"`

	// Output only. State details of the storage pool
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.state_details
	StateDetails *string `json:"stateDetails,omitempty"`

	// Output only. Current storage usage for the snapshot in bytes.
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.used_bytes
	UsedBytes *float64 `json:"usedBytes,omitempty"`

	// Output only. The time when the snapshot was created.
	// +kcc:proto:field=google.cloud.netapp.v1.Snapshot.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetappsnapshot;gcpnetappsnapshots
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetappSnapshot is the Schema for the NetappSnapshot API
// +k8s:openapi-gen=true
type NetappSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetappSnapshotSpec   `json:"spec,omitempty"`
	Status NetappSnapshotStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetappSnapshotList contains a list of NetappSnapshot
type NetappSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetappSnapshot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetappSnapshot{}, &NetappSnapshotList{})
}
