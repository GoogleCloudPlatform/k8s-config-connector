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

var NetworkConnectivityInternalRangeGVK = GroupVersion.WithKind("NetworkConnectivityInternalRange")

// NetworkConnectivityInternalRangeSpec defines the desired state of NetworkConnectivityInternalRange
// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.InternalRange
type NetworkConnectivityInternalRangeSpec struct {
	// The NetworkConnectivityInternalRange name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`

	// A description of this resource.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.description
	Description *string `json:"description,omitempty"`

	// The IP range that this internal range defines. NOTE: IPv6 ranges are limited to usage=EXTERNAL_TO_VPC and peering=FOR_SELF. NOTE: For IPv6 Ranges this field is compulsory, i.e. the address range must be specified explicitly.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.ip_cidr_range
	IPCIDRRange *string `json:"ipCIDRRange,omitempty"`

	// User-defined labels.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Must be present if usage is set to FOR_MIGRATION. This field is for internal use.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.migration
	Migration *Migration `json:"migration,omitempty"`

	// Immutable. The name of an internal range. Format: projects/{project}/locations/{location}/internalRanges/{internal_range} See: https://google.aip.dev/122#fields-representing-resource-names
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.name
	Name *string `json:"name,omitempty"`

	// The network in which to reserve the internal range. The network cannot be deleted if there are any reserved internal ranges referring to it. Legacy networks are not supported. For example: https://www.googleapis.com/compute/v1/projects/{project}/locations/global/networks/{network} projects/{project}/locations/global/networks/{network} {network}
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.network
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Types of resources that are allowed to overlap with the current internal range.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.overlaps
	Overlaps []string `json:"overlaps,omitempty"`

	// The type of peering set for this internal range.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.peering
	Peering *string `json:"peering,omitempty"`

	// An alternate to ip_cidr_range. Can be set when trying to create an IPv4 reservation that automatically finds a free range of the given size. If both ip_cidr_range and prefix_length are set, there is an error if the range sizes do not match. Can also be used during updates to change the range size. NOTE: For IPv6 this field only works if ip_cidr_range is set as well, and both fields must match. In other words, with IPv6 this field only works as a redundant parameter.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.prefix_length
	PrefixLength *int32 `json:"prefixLength,omitempty"`

	// Optional. Can be set to narrow down or pick a different address space while searching for a free range. If not set, defaults to the "10.0.0.0/8" address space. This can be used to search in other rfc-1918 address spaces like "172.16.0.0/12" and "192.168.0.0/16" or non-rfc-1918 address spaces used in the VPC.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.target_cidr_range
	TargetCIDRRange []string `json:"targetCIDRRange,omitempty"`

	// The type of usage set for this InternalRange.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.usage
	Usage *string `json:"usage,omitempty"`
}

// NetworkConnectivityInternalRangeStatus defines the config connector machine state of NetworkConnectivityInternalRange
type NetworkConnectivityInternalRangeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkConnectivityInternalRange resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkConnectivityInternalRangeObservedState `json:"observedState,omitempty"`
}

// NetworkConnectivityInternalRangeObservedState is the state of the NetworkConnectivityInternalRange resource as most recently observed in GCP.
// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.InternalRange
type NetworkConnectivityInternalRangeObservedState struct {
	// Output only. The list of resources that refer to this internal range. Resources that use the internal range for their range allocation are referred to as users of the range. Other resources mark themselves as users while doing so by creating a reference to this internal range. Having a user, based on this reference, prevents deletion of the internal range referred to. Can be empty.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.users
	Users []string `json:"users,omitempty"`

	// Time when the internal range was created.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Time when the internal range was updated.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.InternalRange.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkconnectivityinternalrange;gcpnetworkconnectivityinternalranges
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkConnectivityInternalRange is the Schema for the NetworkConnectivityInternalRange API
// +k8s:openapi-gen=true
type NetworkConnectivityInternalRange struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkConnectivityInternalRangeSpec   `json:"spec,omitempty"`
	Status NetworkConnectivityInternalRangeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkConnectivityInternalRangeList contains a list of NetworkConnectivityInternalRange
type NetworkConnectivityInternalRangeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConnectivityInternalRange `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkConnectivityInternalRange{}, &NetworkConnectivityInternalRangeList{})
}
