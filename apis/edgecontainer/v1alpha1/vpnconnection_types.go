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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EdgeContainerVpnConnectionGVK = GroupVersion.WithKind("EdgeContainerVpnConnection")

// EdgeContainerVpnConnectionSpec defines the desired state of EdgeContainerVpnConnection
// +kcc:proto=google.cloud.edgecontainer.v1.VpnConnection
type EdgeContainerVpnConnectionSpec struct {
	// The EdgeContainerVpnConnection name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// EdgeContainerVpnConnectionStatus defines the config connector machine state of EdgeContainerVpnConnection
type EdgeContainerVpnConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EdgeContainerVpnConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *EdgeContainerVpnConnectionObservedState `json:"observedState,omitempty"`
}

// EdgeContainerVpnConnectionObservedState is the state of the EdgeContainerVpnConnection resource as most recently observed in GCP.
// +kcc:proto=google.cloud.edgecontainer.v1.VpnConnection
type EdgeContainerVpnConnectionObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpedgecontainervpnconnection;gcpedgecontainervpnconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EdgeContainerVpnConnection is the Schema for the EdgeContainerVpnConnection API
// +k8s:openapi-gen=true
type EdgeContainerVpnConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   EdgeContainerVpnConnectionSpec   `json:"spec,omitempty"`
	Status EdgeContainerVpnConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EdgeContainerVpnConnectionList contains a list of EdgeContainerVpnConnection
type EdgeContainerVpnConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeContainerVpnConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EdgeContainerVpnConnection{}, &EdgeContainerVpnConnectionList{})
}

// +kcc:proto=google.cloud.edgecontainer.v1.VpnConnection
type VpnConnection struct {
	// Required. The resource name of VPN connection
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.name
	Name *string `json:"name,omitempty"`

	// Labels associated with this resource.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.labels
	Labels map[string]string `json:"labels,omitempty"`

	// NAT gateway IP, or WAN IP address. If a customer has multiple NAT IPs, the
	//  customer needs to configure NAT such that only one external IP maps to the
	//  GMEC Anthos cluster. This is empty if NAT is not used.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.nat_gateway_ip
	NATGatewayIP *string `json:"natGatewayIP,omitempty"`

	// Dynamic routing mode of the VPC network, `regional` or `global`.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.bgp_routing_mode
	BGPRoutingMode *string `json:"bgpRoutingMode,omitempty"`

	// The canonical Cluster name to connect to. It is in the form of
	//  projects/{project}/locations/{location}/clusters/{cluster}.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.cluster
	Cluster *string `json:"cluster,omitempty"`

	// The network ID of VPC to connect to.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.vpc
	Vpc *string `json:"vpc,omitempty"`

	// Optional. Project detail of the VPC network. Required if VPC is in a
	//  different project than the cluster project.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.vpc_project
	VpcProject *VpnConnection_VpcProject `json:"vpcProject,omitempty"`

	// Whether this VPN connection has HA enabled on cluster side. If enabled,
	//  when creating VPN connection we will attempt to use 2 ANG floating IPs.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.enable_high_availability
	EnableHighAvailability *bool `json:"enableHighAvailability,omitempty"`

	// Optional. The VPN connection Cloud Router name.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.router
	Router *string `json:"router,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.VpnConnection.Details
type VpnConnection_Details struct {
	// The state of this connection.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.Details.state
	State *string `json:"state,omitempty"`

	// The error message. This is only populated when state=ERROR.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.Details.error
	Error *string `json:"error,omitempty"`

	// The Cloud Router info.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.Details.cloud_router
	CloudRouter *VpnConnection_Details_CloudRouter `json:"cloudRouter,omitempty"`

	// Each connection has multiple Cloud VPN gateways.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.Details.cloud_vpns
	CloudVpns []VpnConnection_Details_CloudVpn `json:"cloudVpns,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.VpnConnection.Details.CloudRouter
type VpnConnection_Details_CloudRouter struct {
	// The associated Cloud Router name.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.Details.CloudRouter.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.VpnConnection.Details.CloudVpn
type VpnConnection_Details_CloudVpn struct {
	// The created Cloud VPN gateway name.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.Details.CloudVpn.gateway
	Gateway *string `json:"gateway,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.VpnConnection.VpcProject
type VpnConnection_VpcProject struct {
	// The project of the VPC to connect to. If not specified, it is the same as
	//  the cluster project.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.VpcProject.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Optional. Deprecated: do not use.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.VpcProject.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.VpnConnection
type VpnConnectionObservedState struct {
	// Output only. The time when the VPN connection was created.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the VPN connection was last updated.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The created connection details.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.VpnConnection.details
	Details *VpnConnection_Details `json:"details,omitempty"`
}
