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

// API sources for EdgenetworkRouter, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/edgenetwork/v1/resources.proto
// +kcc:source:service-docs=https://cloud.google.com/distributed-cloud/edge/latest/docs/overview
// +kcc:source:resource-docs=https://docs.cloud.google.com/distributed-cloud/edge/latest/docs/reference/network/rest/v1/projects.locations.zones.routers

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EdgenetworkRouterGVK = GroupVersion.WithKind("EdgenetworkRouter")

// EdgenetworkRouterSpec defines the desired state of EdgenetworkRouter
// +kcc:spec:proto=google.cloud.edgenetwork.v1.Router
// +kcc:required-from-proto
type EdgenetworkRouterSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The EdgenetworkRouter name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Labels associated with this resource.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.description
	Description *string `json:"description,omitempty"`

	// Required. The canonical name of the network to which this router belongs.
	//  The name is in the form of
	//  `projects/{project}/locations/{location}/zones/{zone}/networks/{network}`.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.network
	// +required
	Network *string `json:"network,omitempty"`

	// Router interfaces.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.interface
	Interface []Router_Interface `json:"interface,omitempty"`

	// BGP peers.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.bgp_peer
	BGPPeer []Router_BGPPeer `json:"bgpPeer,omitempty"`

	// BGP information specific to this router.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.bgp
	BGP *Router_BGP `json:"bgp,omitempty"`

	// Optional. A list of CIDRs in IP/Length format to advertise northbound as
	//  static routes from this router.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.route_advertisements
	RouteAdvertisements []string `json:"routeAdvertisements,omitempty"`
}

// EdgenetworkRouterStatus defines the config connector machine state of EdgenetworkRouter
type EdgenetworkRouterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EdgenetworkRouter resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *EdgenetworkRouterObservedState `json:"observedState,omitempty"`
}

// EdgenetworkRouterObservedState is the state of the EdgenetworkRouter resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.edgenetwork.v1.Router
type EdgenetworkRouterObservedState struct {
	// Output only. The time when the router was created.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the router was last updated.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// BGP peers.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.bgp_peer
	BGPPeer []Router_BGPPeerObservedState `json:"bgpPeer,omitempty"`

	// Output only. Current stage of the resource to the device by config push.
	// +kcc:proto:field=google.cloud.edgenetwork.v1.Router.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpedgenetworkrouter;gcpedgenetworkrouters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EdgenetworkRouter is the Schema for the EdgenetworkRouter API
// +k8s:openapi-gen=true
type EdgenetworkRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   EdgenetworkRouterSpec   `json:"spec,omitempty"`
	Status EdgenetworkRouterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EdgenetworkRouterList contains a list of EdgenetworkRouter
type EdgenetworkRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgenetworkRouter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EdgenetworkRouter{}, &EdgenetworkRouterList{})
}
