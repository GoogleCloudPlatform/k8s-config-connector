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

// API sources for NetworkConnectivityPolicyBasedRoute, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/networkconnectivity/v1/policy_based_routing.proto
// +kcc:source:service-docs=https://cloud.google.com/network-connectivity/docs/reference/networkconnectivity/rest
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/network-connectivity/docs/reference/networkconnectivity/rest/reference/rest/v1/projects.locations.global.PolicyBasedRoutes

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkConnectivityPolicyBasedRouteGVK = GroupVersion.WithKind("NetworkConnectivityPolicyBasedRoute")

// NetworkConnectivityPolicyBasedRouteSpec defines the desired state of NetworkConnectivityPolicyBasedRoute
// +kcc:spec:proto=google.cloud.networkconnectivity.v1.PolicyBasedRoute
// +kcc:required-from-proto
type NetworkConnectivityPolicyBasedRouteSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The NetworkConnectivityPolicyBasedRoute name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. VM instances that this policy-based route applies to.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.virtual_machine
	VirtualMachine *PolicyBasedRoute_VirtualMachine `json:"virtualMachine,omitempty"`

	// Optional. The interconnect attachments that this policy-based route
	//  applies to.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.interconnect_attachment
	InterconnectAttachment *PolicyBasedRoute_InterconnectAttachment `json:"interconnectAttachment,omitempty"`

	// Optional. The IP address of a global-access-enabled L4 ILB that is the
	//  next hop for matching packets. For this version, only nextHopIlbIp is
	//  supported.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.next_hop_ilb_ip
	NextHopIlbIP *string `json:"nextHopIlbIP,omitempty"`

	// Optional. Other routes that will be referenced to determine the next hop
	//  of the packet.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.next_hop_other_routes
	NextHopOtherRoutes *string `json:"nextHopOtherRoutes,omitempty"`

	// User-defined labels.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. An optional description of this resource. Provide this field when
	//  you create the resource.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.description
	Description *string `json:"description,omitempty"`

	// Required. Fully-qualified URL of the network that this route applies to,
	//  for example: projects/my-project/global/networks/my-network.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.network
	// +required
	Network *string `json:"network,omitempty"`

	// Required. The filter to match L4 traffic.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.filter
	// +required
	Filter *PolicyBasedRoute_Filter `json:"filter,omitempty"`

	// Optional. The priority of this policy-based route. Priority is used to
	//  break ties in cases where there are more than one matching policy-based
	//  routes found. In cases where multiple policy-based routes are matched, the
	//  one with the lowest-numbered priority value wins. The default value is
	//  1000. The priority value must be from 1 to 65535, inclusive.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.priority
	Priority *int32 `json:"priority,omitempty"`
}

// NetworkConnectivityPolicyBasedRouteStatus defines the config connector machine state of NetworkConnectivityPolicyBasedRoute
type NetworkConnectivityPolicyBasedRouteStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkConnectivityPolicyBasedRoute resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkConnectivityPolicyBasedRouteObservedState `json:"observedState,omitempty"`
}

// NetworkConnectivityPolicyBasedRouteObservedState is the state of the NetworkConnectivityPolicyBasedRoute resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networkconnectivity.v1.PolicyBasedRoute
type NetworkConnectivityPolicyBasedRouteObservedState struct {
	// Output only. Time when the policy-based route was created.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the policy-based route was updated.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. If potential misconfigurations are detected for this route,
	//  this field will be populated with warning messages.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.warnings
	Warnings []PolicyBasedRoute_WarningsObservedState `json:"warnings,omitempty"`

	// Output only. Server-defined fully-qualified URL for this resource.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. Type of this resource. Always
	//  networkconnectivity#policyBasedRoute for policy-based Route resources.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.PolicyBasedRoute.kind
	Kind *string `json:"kind,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkconnectivitypolicybasedroute;gcpnetworkconnectivitypolicybasedroutes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkConnectivityPolicyBasedRoute is the Schema for the NetworkConnectivityPolicyBasedRoute API
// +k8s:openapi-gen=true
type NetworkConnectivityPolicyBasedRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkConnectivityPolicyBasedRouteSpec   `json:"spec,omitempty"`
	Status NetworkConnectivityPolicyBasedRouteStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkConnectivityPolicyBasedRouteList contains a list of NetworkConnectivityPolicyBasedRoute
type NetworkConnectivityPolicyBasedRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConnectivityPolicyBasedRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkConnectivityPolicyBasedRoute{}, &NetworkConnectivityPolicyBasedRouteList{})
}
