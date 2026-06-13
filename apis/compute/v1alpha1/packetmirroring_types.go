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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputePacketMirroringGVK = GroupVersion.WithKind("ComputePacketMirroring")

// ComputePacketMirroringSpec defines the desired state of ComputePacketMirroring
// +kcc:spec:proto=google.cloud.compute.v1.PacketMirroring
type ComputePacketMirroringSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The ComputePacketMirroring name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// The Forwarding Rule resource of type loadBalancingScheme=INTERNAL that will be used as collector for mirrored traffic. The specified forwarding rule must have isMirroringCollector set to true.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.collector_ilb
	CollectorIlb PacketMirroringCollectorIlb `json:"collectorIlb"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.description
	Description *string `json:"description,omitempty"`

	// Indicates whether or not this packet mirroring takes effect. If set to FALSE, this packet mirroring policy will not be enforced on the network. The default is TRUE.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.enable
	Enable *string `json:"enable,omitempty"`

	// Filter for mirrored traffic. If unspecified, all IPv4 traffic is mirrored.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.filter
	Filter *PacketMirroringFilter `json:"filter,omitempty"`

	// PacketMirroring mirroredResourceInfos. MirroredResourceInfo specifies a set of mirrored VM instances, subnetworks and/or tags for which traffic from/to all VM instances will be mirrored.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.mirrored_resources
	MirroredResources PacketMirroringMirroredResources `json:"mirroredResources"`

	// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in the given network. All mirrored subnetworks should belong to the given network.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.network
	Network PacketMirroringNetwork `json:"network"`

	// The priority of applying this configuration. Priority is used to break ties in cases where there is more than one matching rule. In the case of two rules that apply for a given Instance, the one with the lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.priority
	Priority *uint32 `json:"priority,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.PacketMirroringForwardingRuleInfo
type PacketMirroringCollectorIlb struct {
	// Reference to a ComputeForwardingRule resource.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringForwardingRuleInfo.url
	URLRef *computev1beta1.ForwardingRuleRef `json:"urlRef"`
}

// +kcc:proto=google.cloud.compute.v1.PacketMirroringNetworkInfo
type PacketMirroringNetwork struct {
	// URL of the network resource.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringNetworkInfo.url
	URLRef *computev1beta1.ComputeNetworkRef `json:"urlRef"`
}

// +kcc:proto=google.cloud.compute.v1.PacketMirroringFilter
type PacketMirroringFilter struct {
	// Protocols that apply as filter on mirrored traffic. If no protocols are specified, all traffic that matches the specified CIDR ranges is mirrored. If neither cidrRanges nor IPProtocols is specified, all IPv4 traffic is mirrored.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringFilter.I_p_protocols
	IPProtocols []string `json:"ipProtocols,omitempty"`

	// One or more IPv4 or IPv6 CIDR ranges that apply as filters on the source (ingress) or destination (egress) IP in the IP header. If no ranges are specified, all IPv4 traffic that matches the specified IPProtocols is mirrored. If neither cidrRanges nor IPProtocols is specified, all IPv4 traffic is mirrored. To mirror all IPv4 and IPv6 traffic, use "0.0.0.0/0,::/0".
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringFilter.cidr_ranges
	CIDRRanges []string `json:"cidrRanges,omitempty"`

	// Direction of traffic to mirror, either INGRESS, EGRESS, or BOTH. The default is BOTH.
	//  Check the Direction enum for the list of possible values.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringFilter.direction
	Direction *string `json:"direction,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.PacketMirroringMirroredResourceInfo
type PacketMirroringMirroredResources struct {
	// A set of virtual machine instances that are being mirrored. They must live in zones contained in the same region as this packetMirroring. Note that this config will apply only to those network interfaces of the Instances that belong to the network specified in this packetMirroring. You may specify a maximum of 50 Instances.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringMirroredResourceInfo.instances
	Instances []PacketMirroringMirroredResourceInfoInstance `json:"instances,omitempty"`

	// A set of subnetworks for which traffic from/to all VM instances will be mirrored. They must live in the same region as this packetMirroring. You may specify a maximum of 5 subnetworks.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringMirroredResourceInfo.subnetworks
	Subnetworks []PacketMirroringMirroredResourceInfoSubnet `json:"subnetworks,omitempty"`

	// A set of mirrored tags. Traffic from/to all VM instances that have one or more of these tags will be mirrored.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringMirroredResourceInfo.tags
	Tags []string `json:"tags,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.PacketMirroringMirroredResourceInfoInstanceInfo
type PacketMirroringMirroredResourceInfoInstance struct {
	// Resource URL to the virtual machine instance which is being mirrored.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringMirroredResourceInfoInstanceInfo.url
	URLRef *computev1beta1.InstanceRef `json:"urlRef"`
}

// +kcc:proto=google.cloud.compute.v1.PacketMirroringMirroredResourceInfoSubnetInfo
type PacketMirroringMirroredResourceInfoSubnet struct {
	// Resource URL to the subnetwork for which traffic from/to all VM instances will be mirrored.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringMirroredResourceInfoSubnetInfo.url
	URLRef *computev1beta1.ComputeSubnetworkRef `json:"urlRef"`
}

// ComputePacketMirroringStatus defines the config connector machine state of ComputePacketMirroring
type ComputePacketMirroringStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputePacketMirroring resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputePacketMirroringObservedState `json:"observedState,omitempty"`
}

// ComputePacketMirroringObservedState is the state of the ComputePacketMirroring resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.PacketMirroring
type ComputePacketMirroringObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of the resource. Always compute#packetMirroring for packet mirrorings.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] URI of the region where the packetMirroring resides.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.region
	Region *string `json:"region,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// The Forwarding Rule resource of type loadBalancingScheme=INTERNAL that will be used as collector for mirrored traffic. The specified forwarding rule must have isMirroringCollector set to true.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.collector_ilb
	CollectorIlb *PacketMirroringCollectorIlbObservedState `json:"collectorIlb,omitempty"`

	// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in the given network. All mirrored subnetworks should belong to the given network.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.network
	Network *PacketMirroringNetworkObservedState `json:"network,omitempty"`

	// PacketMirroring mirroredResourceInfos. MirroredResourceInfo specifies a set of mirrored VM instances, subnetworks and/or tags for which traffic from/to all VM instances will be mirrored.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroring.mirrored_resources
	MirroredResources *PacketMirroringMirroredResourcesObservedState `json:"mirroredResources,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.PacketMirroringForwardingRuleInfo
type PacketMirroringCollectorIlbObservedState struct {
	// [Output Only] Unique identifier for the forwarding rule; defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringForwardingRuleInfo.canonical_url
	CanonicalURL *string `json:"canonicalURL,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.PacketMirroringNetworkInfo
type PacketMirroringNetworkObservedState struct {
	// [Output Only] Unique identifier for the network; defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringNetworkInfo.canonical_url
	CanonicalURL *string `json:"canonicalURL,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.PacketMirroringMirroredResourceInfo
type PacketMirroringMirroredResourcesObservedState struct {
	// A set of virtual machine instances that are being mirrored. They must live in zones contained in the same region as this packetMirroring. Note that this config will apply only to those network interfaces of the Instances that belong to the network specified in this packetMirroring. You may specify a maximum of 50 Instances.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringMirroredResourceInfo.instances
	Instances []PacketMirroringMirroredResourceInfoInstanceInfoObservedState `json:"instances,omitempty"`

	// A set of subnetworks for which traffic from/to all VM instances will be mirrored. They must live in the same region as this packetMirroring. You may specify a maximum of 5 subnetworks.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringMirroredResourceInfo.subnetworks
	Subnetworks []PacketMirroringMirroredResourceInfoSubnetInfoObservedState `json:"subnetworks,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.PacketMirroringMirroredResourceInfoInstanceInfo
type PacketMirroringMirroredResourceInfoInstanceInfoObservedState struct {
	// [Output Only] Unique identifier for the instance; defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringMirroredResourceInfoInstanceInfo.canonical_url
	CanonicalURL *string `json:"canonicalURL,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.PacketMirroringMirroredResourceInfoSubnetInfo
type PacketMirroringMirroredResourceInfoSubnetInfoObservedState struct {
	// [Output Only] Unique identifier for the subnetwork; defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.PacketMirroringMirroredResourceInfoSubnetInfo.canonical_url
	CanonicalURL *string `json:"canonicalURL,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputepacketmirroring;gcpcomputepacketmirrorings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputePacketMirroring is the Schema for the ComputePacketMirroring API
// +k8s:openapi-gen=true
type ComputePacketMirroring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputePacketMirroringSpec   `json:"spec,omitempty"`
	Status ComputePacketMirroringStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputePacketMirroringList contains a list of ComputePacketMirroring
type ComputePacketMirroringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputePacketMirroring `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputePacketMirroring{}, &ComputePacketMirroringList{})
}
