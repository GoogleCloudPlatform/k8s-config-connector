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

var ComputeResourcePolicyGVK = GroupVersion.WithKind("ComputeResourcePolicy")

// ComputeResourcePolicyParent holds the fields describing the parent of the ComputeResourcePolicy resource.
type ComputeResourcePolicyParent struct {
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`
	// +required
	Location string `json:"location"`
}

// ComputeResourcePolicySpec defines the desired state of ComputeResourcePolicy
// +kcc:spec:proto=google.cloud.compute.v1.ResourcePolicy
type ComputeResourcePolicySpec struct {
	// Parent reference.
	ComputeResourcePolicyParent `json:",inline"`

	// The ComputeResourcePolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.description
	Description *string `json:"description,omitempty"`

	// Immutable. Replication consistency group for asynchronous disk replication.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.disk_consistency_group_policy
	DiskConsistencyGroupPolicy *ResourcePolicyDiskConsistencyGroupPolicy `json:"diskConsistencyGroupPolicy,omitempty"`

	// Immutable. Resource policy for instances used for placement configuration.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.group_placement_policy
	GroupPlacementPolicy *ResourcePolicyGroupPlacementPolicy `json:"groupPlacementPolicy,omitempty"`

	// Immutable. Resource policy for scheduling instance operations.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.instance_schedule_policy
	InstanceSchedulePolicy *ResourcePolicyInstanceSchedulePolicy `json:"instanceSchedulePolicy,omitempty"`

	// Immutable. Policy for creating snapshots of persistent disks.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.snapshot_schedule_policy
	SnapshotSchedulePolicy *ResourcePolicySnapshotSchedulePolicy `json:"snapshotSchedulePolicy,omitempty"`
}

// ComputeResourcePolicyStatus defines the config connector machine state of ComputeResourcePolicy
type ComputeResourcePolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeResourcePolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeResourcePolicyObservedState `json:"observedState,omitempty"`
}

// ComputeResourcePolicyObservedState is the state of the ComputeResourcePolicy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.ResourcePolicy
type ComputeResourcePolicyObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] The status of resource policy creation.
	// +kcc:proto:field=google.cloud.compute.v1.ResourcePolicy.status
	Status *string `json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeresourcepolicy;gcpcomputeresourcepolicys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeResourcePolicy is the Schema for the ComputeResourcePolicy API
// +k8s:openapi-gen=true
type ComputeResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeResourcePolicySpec   `json:"spec,omitempty"`
	Status ComputeResourcePolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeResourcePolicyList contains a list of ComputeResourcePolicy
type ComputeResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeResourcePolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeResourcePolicy{}, &ComputeResourcePolicyList{})
}
