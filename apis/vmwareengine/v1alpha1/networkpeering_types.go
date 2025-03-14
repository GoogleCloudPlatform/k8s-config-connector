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

var VMwareEngineNetworkPeeringGVK = GroupVersion.WithKind("VMwareEngineNetworkPeering")

// VMwareEngineNetworkPeeringSpec defines the desired state of VMwareEngineNetworkPeering
// +kcc:proto=google.cloud.vmwareengine.v1.NetworkPeering
type VMwareEngineNetworkPeeringSpec struct {
	// The VMwareEngineNetworkPeering name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Required. The relative resource name of the network to peer with
	//  a standard VMware Engine network. The provided network can be a
	//  consumer VPC network or another standard VMware Engine network. If the
	//  `peer_network_type` is VMWARE_ENGINE_NETWORK, specify the name in the form:
	//  `projects/{project}/locations/global/vmwareEngineNetworks/{vmware_engine_network_id}`.
	//  Otherwise specify the name in the form:
	//  `projects/{project}/global/networks/{network_id}`, where
	//  `{project}` can either be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.peer_network
	// +required
	PeerNetwork *string `json:"peerNetwork,omitempty"`

	// Optional. True if custom routes are exported to the peered network;
	//  false otherwise. The default value is true.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.export_custom_routes
	ExportCustomRoutes *bool `json:"exportCustomRoutes,omitempty"`

	// Optional. True if custom routes are imported from the peered network;
	//  false otherwise. The default value is true.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.import_custom_routes
	ImportCustomRoutes *bool `json:"importCustomRoutes,omitempty"`

	// Optional. True if full mesh connectivity is created and managed
	//  automatically between peered networks; false otherwise. Currently this
	//  field is always true because Google Compute Engine automatically creates
	//  and manages subnetwork routes between two VPC networks when peering state
	//  is 'ACTIVE'.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.exchange_subnet_routes
	ExchangeSubnetRoutes *bool `json:"exchangeSubnetRoutes,omitempty"`

	// Optional. True if all subnet routes with a public IP address range are
	//  exported; false otherwise. The default value is true. IPv4 special-use
	//  ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always
	//  exported to peers and are not controlled by this field.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.export_custom_routes_with_public_ip
	ExportCustomRoutesWithPublicIP *bool `json:"exportCustomRoutesWithPublicIP,omitempty"`

	// Optional. True if all subnet routes with public IP address range are
	//  imported; false otherwise. The default value is true. IPv4 special-use
	//  ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always
	//  imported to peers and are not controlled by this field.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.import_custom_routes_with_public_ip
	ImportCustomRoutesWithPublicIP *bool `json:"importCustomRoutesWithPublicIP,omitempty"`

	// Optional. Maximum transmission unit (MTU) in bytes.
	//  The default value is `1500`. If a value of `0` is provided for this field,
	//  VMware Engine uses the default value instead.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.peer_mtu
	PeerMTU *int32 `json:"peerMTU,omitempty"`

	// Required. The type of the network to peer with the VMware Engine network.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.peer_network_type
	// +required
	PeerNetworkType *string `json:"peerNetworkType,omitempty"`

	// Required. The relative resource name of the VMware Engine network.
	//  Specify the name in the following form:
	//  `projects/{project}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`
	//  where `{project}` can either be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.vmware_engine_network
	// +required
	VmwareEngineNetwork *string `json:"vmwareEngineNetwork,omitempty"`

	// Optional. User-provided description for this network peering.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.description
	Description *string `json:"description,omitempty"`
}

// VMwareEngineNetworkPeeringStatus defines the config connector machine state of VMwareEngineNetworkPeering
type VMwareEngineNetworkPeeringStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMwareEngineNetworkPeering resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMwareEngineNetworkPeeringObservedState `json:"observedState,omitempty"`
}

// VMwareEngineNetworkPeeringObservedState is the state of the VMwareEngineNetworkPeering resource as most recently observed in GCP.
// +kcc:proto=google.cloud.vmwareengine.v1.NetworkPeering
type VMwareEngineNetworkPeeringObservedState struct {
	// Output only. The resource name of the network peering. NetworkPeering is a
	//  global resource and location can only be global. Resource names are
	//  scheme-less URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/global/networkPeerings/my-peering`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the network peering. This field
	//  has a value of 'ACTIVE' when there's a matching configuration in the peer
	//  network. New values may be added to this enum when appropriate.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.state
	State *string `json:"state,omitempty"`

	// Output only. Output Only. Details about the current state of the network
	//  peering.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.state_details
	StateDetails *string `json:"stateDetails,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPeering.uid
	UID *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmwareenginenetworkpeering;gcpvmwareenginenetworkpeerings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMwareEngineNetworkPeering is the Schema for the VMwareEngineNetworkPeering API
// +k8s:openapi-gen=true
type VMwareEngineNetworkPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMwareEngineNetworkPeeringSpec   `json:"spec,omitempty"`
	Status VMwareEngineNetworkPeeringStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMwareEngineNetworkPeeringList contains a list of VMwareEngineNetworkPeering
type VMwareEngineNetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMwareEngineNetworkPeering `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMwareEngineNetworkPeering{}, &VMwareEngineNetworkPeeringList{})
}
