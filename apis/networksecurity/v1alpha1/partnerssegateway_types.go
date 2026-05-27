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

var NetworkSecurityPartnerSSEGatewayGVK = GroupVersion.WithKind("NetworkSecurityPartnerSSEGateway")

// +kcc:proto=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.PartnerSSEGatewaySymantecOptions
type NetworkSecurityPartnerSSEGatewaySymantecOptions struct {

	// Optional. Target for the NCGs to send traffic to on the Symantec side.
	//  Only supports IP address today.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.PartnerSSEGatewaySymantecOptions.symantec_site_target_host
	SymantecSiteTargetHost *string `json:"symantecSiteTargetHost,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.PartnerSSEGatewaySymantecOptions
type NetworkSecurityPartnerSSEGatewaySymantecOptionsObservedState struct {
	// Output only. UUID of the Symantec Location created on the customer's
	//  behalf.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.PartnerSSEGatewaySymantecOptions.symantec_location_uuid
	SymantecLocationUuid *string `json:"symantecLocationUuid,omitempty"`

	// Output only. Symantec data center identifier that this SSEGW will connect
	//  to. Filled from the customer SSEGateway, and only for PartnerSSEGateways
	//  associated with Symantec today.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.PartnerSSEGatewaySymantecOptions.symantec_site
	SymantecSite *string `json:"symantecSite,omitempty"`
}

// NetworkSecurityPartnerSSEGatewaySpec defines the desired state of NetworkSecurityPartnerSSEGateway
// +kcc:spec:proto=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway
type NetworkSecurityPartnerSSEGatewaySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetworkSecurityPartnerSSEGateway name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels as key value pairs
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. ID of the SSEGatewayReference that pairs with this
	//  PartnerSSEGateway
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.sse_gateway_reference_id
	SseGatewayReferenceID *string `json:"sseGatewayReferenceID,omitempty"`

	// Optional. Subnet range of the partner_vpc
	//  This field is deprecated. Use partner_subnet_range instead.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.partner_vpc_subnet_range
	PartnerVPCSubnetRange *string `json:"partnerVPCSubnetRange,omitempty"`

	// Optional. Subnet range where SSE GW instances are deployed.
	//  Default value is set to "100.88.255.0/24".
	//  The CIDR suffix should be less than or equal to 25.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.sse_subnet_range
	SseSubnetRange *string `json:"sseSubnetRange,omitempty"`

	// Optional. Subnet range of the partner-owned subnet.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.partner_subnet_range
	PartnerSubnetRange *string `json:"partnerSubnetRange,omitempty"`

	// Optional. Virtual Network Identifier to use in NCG.
	//  Today the only partner that depends on it is Symantec.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.vni
	Vni *int32 `json:"vni,omitempty"`

	// Optional. Required iff Partner is Symantec.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.symantec_options
	SymantecOptions *NetworkSecurityPartnerSSEGatewaySymantecOptions `json:"symantecOptions,omitempty"`
}

// NetworkSecurityPartnerSSEGatewayObservedState is the state of the NetworkSecurityPartnerSSEGateway resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway
type NetworkSecurityPartnerSSEGatewayObservedState struct {
	// Output only. Create time stamp
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Subnet range of the subnet where partner traffic is routed.
	//  This field is deprecated. Use sse_subnet_range instead.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.sse_vpc_subnet_range
	SseVPCSubnetRange *string `json:"sseVPCSubnetRange,omitempty"`

	// Output only. This is the IP where the partner traffic should be routed to.
	//  This field is deprecated. Use sse_target_ip instead.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.sse_vpc_target_ip
	SseVPCTargetIP *string `json:"sseVPCTargetIP,omitempty"`

	// Output only. IP of SSE BGP
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.sse_bgp_ips
	SseBGPIps []string `json:"sseBGPIps,omitempty"`

	// Output only. ASN of SSE BGP
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.sse_bgp_asn
	SseBGPAsn *int32 `json:"sseBGPAsn,omitempty"`

	// Output only. name of PartnerSSERealm owning the PartnerSSEGateway
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.partner_sse_realm
	PartnerSseRealm *string `json:"partnerSseRealm,omitempty"`

	// Output only. Target IP that belongs to sse_subnet_range where partner
	//  should send the traffic to reach the customer networks.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.sse_target_ip
	SseTargetIP *string `json:"sseTargetIP,omitempty"`

	// Optional. Required iff Partner is Symantec.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.symantec_options
	SymantecOptions *NetworkSecurityPartnerSSEGatewaySymantecOptionsObservedState `json:"symantecOptions,omitempty"`

	// Output only. The project owning partner_facing_network. Only filled for
	//  PartnerSSEGateways associated with Symantec today.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.sse_project
	SseProject *string `json:"sseProject,omitempty"`

	// Output only. The ID of the network in sse_project containing
	//  sse_subnet_range. This is also known as the partnerFacingNetwork. Only
	//  filled for PartnerSSEGateways associated with Symantec today.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.sse_network
	SseNetwork *string `json:"sseNetwork,omitempty"`

	// Output only. Full URI of the partner environment this PartnerSSEGateway is
	//  connected to. Filled from the customer SSEGateway, and only for
	//  PartnerSSEGateways associated with Symantec today.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.partner_sse_environment
	PartnerSseEnvironment *string `json:"partnerSseEnvironment,omitempty"`

	// Output only. ISO-3166 alpha 2 country code used for localization.
	//  Filled from the customer SSEGateway, and only for PartnerSSEGateways
	//  associated with Symantec today.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.country
	Country *string `json:"country,omitempty"`

	// Output only. tzinfo identifier used for localization.
	//  Filled from the customer SSEGateway, and only for PartnerSSEGateways
	//  associated with Symantec today.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.timezone
	Timezone *string `json:"timezone,omitempty"`

	// Output only. Copied from the associated NCC resource in Symantec NCCGW
	//  flows. Used by Symantec API.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.capacity_bps
	CapacityBps *int64 `json:"capacityBps,omitempty"`

	// Output only. State of the gateway.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.state
	State *string `json:"state,omitempty"`

	// Output only. Subnet ranges for Google-issued probe packets.
	//  It's populated only for Prisma Access partners.
	// +kcc:proto:field=google.cloud.networksecurity.v1alpha1.PartnerSSEGateway.prober_subnet_ranges
	ProberSubnetRanges []string `json:"proberSubnetRanges,omitempty"`
}

// NetworkSecurityPartnerSSEGatewayStatus defines the config connector machine state of NetworkSecurityPartnerSSEGateway
type NetworkSecurityPartnerSSEGatewayStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityPartnerSSEGateway resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityPartnerSSEGatewayObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritypartnerssegateway;gcpnetworksecuritypartnerssegateways
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityPartnerSSEGateway is the Schema for the NetworkSecurityPartnerSSEGateway API
// +k8s:openapi-gen=true
type NetworkSecurityPartnerSSEGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityPartnerSSEGatewaySpec   `json:"spec,omitempty"`
	Status NetworkSecurityPartnerSSEGatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityPartnerSSEGatewayList contains a list of NetworkSecurityPartnerSSEGateway
type NetworkSecurityPartnerSSEGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityPartnerSSEGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityPartnerSSEGateway{}, &NetworkSecurityPartnerSSEGatewayList{})
}
