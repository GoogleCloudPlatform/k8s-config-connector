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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SpannerInstancePartitionGVK = GroupVersion.WithKind("SpannerInstancePartition")

// SpannerInstancePartitionSpec defines the desired state of SpannerInstancePartition
// +kcc:proto=google.spanner.admin.instance.v1.InstancePartition
type SpannerInstancePartitionSpec struct {

	// Immutable. The Project that this resource belongs to.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// The location of the cluster.
	Location string `json:"location,omitempty"`

	// The SpannerInstancePartition name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The reference to the instance partition's configuration.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.config
	Config *InstanceConfigRef `json:"configRef,omitempty"`

	// Required. The descriptive name for this instance partition as it appears in
	//  UIs. Must be unique per project and between 4 and 30 characters in length.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The number of nodes allocated to this instance partition.
	//
	//  Users can set the `node_count` field to specify the target number of
	//  nodes allocated to the instance partition.
	//
	//  This may be zero in API responses for instance partitions that are not
	//  yet in state `READY`.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// The number of processing units allocated to this instance partition.
	//
	//  Users can set the `processing_units` field to specify the target number
	//  of processing units allocated to the instance partition.
	//
	//  This might be zero in API responses for instance partitions that are not
	//  yet in the `READY` state.
	// +kcc:proto:field=google.spanner.admin.instance.v1.InstancePartition.processing_units
	ProcessingUnits *int32 `json:"processingUnits,omitempty"`
}

// SpannerInstancePartitionStatus defines the config connector machine state of SpannerInstancePartition
type SpannerInstancePartitionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SpannerInstancePartition resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SpannerInstancePartitionObservedState `json:"observedState,omitempty"`
}

// SpannerInstancePartitionSpec defines the desired state of SpannerInstancePartition
// +kcc:proto=google.spanner.admin.instance.v1.InstancePartition
// SpannerInstancePartitionObservedState is the state of the SpannerInstancePartition resource as most recently observed in GCP.
type SpannerInstancePartitionObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpspannerinstancepartition;gcpspannerinstancepartitions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SpannerInstancePartition is the Schema for the SpannerInstancePartition API
// +k8s:openapi-gen=true
type SpannerInstancePartition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SpannerInstancePartitionSpec   `json:"spec,omitempty"`
	Status SpannerInstancePartitionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SpannerInstancePartitionList contains a list of SpannerInstancePartition
type SpannerInstancePartitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpannerInstancePartition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpannerInstancePartition{}, &SpannerInstancePartitionList{})
}
