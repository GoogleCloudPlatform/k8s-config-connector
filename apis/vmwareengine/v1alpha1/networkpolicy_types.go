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

var VMwareEngineNetworkPolicyGVK = GroupVersion.WithKind("VMwareEngineNetworkPolicy")

// VMwareEngineNetworkPolicySpec defines the desired state of VMwareEngineNetworkPolicy
// +kcc:proto=google.cloud.vmwareengine.v1.NetworkPolicy
type VMwareEngineNetworkPolicySpec struct {
	// The VMwareEngineNetworkPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Network service that allows VMware workloads to access the internet.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.internet_access
	InternetAccess *NetworkPolicy_NetworkService `json:"internetAccess,omitempty"`

	// Network service that allows External IP addresses to be assigned to VMware
	//  workloads. This service can only be enabled when `internet_access` is also
	//  enabled.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.external_ip
	ExternalIP *NetworkPolicy_NetworkService `json:"externalIP,omitempty"`

	// Required. IP address range in CIDR notation used to create internet access
	//  and external IP access. An RFC 1918 CIDR block, with a "/26" prefix, is
	//  required. The range cannot overlap with any prefixes either in the consumer
	//  VPC network or in use by the private clouds attached to that VPC network.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.edge_services_cidr
	// +required
	EdgeServicesCIDR *string `json:"edgeServicesCIDR,omitempty"`

	// Optional. The relative resource name of the VMware Engine network.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.vmware_engine_network
	VMwareEngineNetworkRef *VmwareEngineNetworkRef `json:"vmwareEngineNetworkRef,omitempty"`

	// Optional. User-provided description for this network policy.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.description
	Description *string `json:"description,omitempty"`
}

// VMwareEngineNetworkPolicyStatus defines the config connector machine state of VMwareEngineNetworkPolicy
type VMwareEngineNetworkPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMwareEngineNetworkPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMwareEngineNetworkPolicyObservedState `json:"observedState,omitempty"`
}

// VMwareEngineNetworkPolicyObservedState is the state of the VMwareEngineNetworkPolicy resource as most recently observed in GCP.
// +kcc:proto=google.cloud.vmwareengine.v1.NetworkPolicy
type VMwareEngineNetworkPolicyObservedState struct {
	// Output only. The resource name of this network policy.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1/networkPolicies/my-network-policy`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Network service that allows VMware workloads to access the internet.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.internet_access
	InternetAccess *NetworkPolicy_NetworkServiceObservedState `json:"internetAccess,omitempty"`

	// Network service that allows External IP addresses to be assigned to VMware
	//  workloads. This service can only be enabled when `internet_access` is also
	//  enabled.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.external_ip
	ExternalIP *NetworkPolicy_NetworkServiceObservedState `json:"externalIP,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.uid
	UID *string `json:"uid,omitempty"`

	// Output only. The canonical name of the VMware Engine network in the form:
	//  `projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.NetworkPolicy.vmware_engine_network_canonical
	VMwareEngineNetworkCanonical *string `json:"vmwareEngineNetworkCanonical,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmwareenginenetworkpolicy;gcpvmwareenginenetworkpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMwareEngineNetworkPolicy is the Schema for the VMwareEngineNetworkPolicy API
// +k8s:openapi-gen=true
type VMwareEngineNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMwareEngineNetworkPolicySpec   `json:"spec,omitempty"`
	Status VMwareEngineNetworkPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMwareEngineNetworkPolicyList contains a list of VMwareEngineNetworkPolicy
type VMwareEngineNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMwareEngineNetworkPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMwareEngineNetworkPolicy{}, &VMwareEngineNetworkPolicyList{})
}
