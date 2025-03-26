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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkManagementConnectivityTestGVK = GroupVersion.WithKind("NetworkManagementConnectivityTest")

// NetworkManagementConnectivityTestSpec defines the desired state of NetworkManagementConnectivityTest
// +kcc:proto=google.cloud.networkmanagement.v1.ConnectivityTest
type NetworkManagementConnectivityTestSpec struct {
	// The project that this resource belongs to. If not provided, the provider project is used.
	ProjectRef refsv1beta1.Reference `json:"projectRef"`
	// The NetworkManagementConnectivityTest name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Identifier. Unique name of the resource using the form:
	//      `projects/{project_id}/locations/global/connectivityTests/{test_id}`
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.name
	Name *string `json:"name,omitempty"`

	// The user-supplied description of the Connectivity Test.
	//  Maximum of 512 characters.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.description
	Description *string `json:"description,omitempty"`

	// Required. Source specification of the Connectivity Test.
	//
	//  You can use a combination of source IP address, virtual machine
	//  (VM) instance, or Compute Engine network to uniquely identify
	//  the source location.
	//
	//  Examples:
	//  If the source IP address is an internal IP address within a Google Cloud
	//  Virtual Private Cloud (VPC) network, then you must also specify the VPC
	//  network. Otherwise, specify the VM instance, which already contains its
	//  internal IP address and VPC network information.
	//
	//  If the source of the test is within an on-premises network, then you must
	//  provide the destination VPC network.
	//
	//  If the source endpoint is a Compute Engine VM instance with multiple
	//  network interfaces, the instance itself is not sufficient to identify the
	//  endpoint. So, you must also specify the source IP address or VPC network.
	//
	//  A reachability analysis proceeds even if the source location is
	//  ambiguous. However, the test result may include endpoints that you don't
	//  intend to test.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.source
	Source *Endpoint `json:"source,omitempty"`

	// Required. Destination specification of the Connectivity Test.
	//
	//  You can use a combination of destination IP address, Compute Engine
	//  VM instance, or VPC network to uniquely identify the destination
	//  location.
	//
	//  Even if the destination IP address is not unique, the source IP
	//  location is unique. Usually, the analysis can infer the destination
	//  endpoint from route information.
	//
	//  If the destination you specify is a VM instance and the instance has
	//  multiple network interfaces, then you must also specify either
	//  a destination IP address  or VPC network to identify the destination
	//  interface.
	//
	//  A reachability analysis proceeds even if the destination location is
	//  ambiguous. However, the result can include endpoints that you don't
	//  intend to test.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.destination
	Destination *Endpoint `json:"destination,omitempty"`

	// IP Protocol of the test. When not provided, "TCP" is assumed.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.protocol
	Protocol *string `json:"protocol,omitempty"`

	// Other projects that may be relevant for reachability analysis.
	//  This is applicable to scenarios where a test can cross project boundaries.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.related_projects
	RelatedProjects []string `json:"relatedProjects,omitempty"`

	// Resource labels to represent user-provided metadata.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Whether run analysis for the return path from destination to source.
	//  Default value is false.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.round_trip
	RoundTrip *bool `json:"roundTrip,omitempty"`

	// Whether the test should skip firewall checking.
	//  If not provided, we assume false.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.bypass_firewall_checks
	BypassFirewallChecks *bool `json:"bypassFirewallChecks,omitempty"`
}

// NetworkManagementConnectivityTestStatus defines the config connector machine state of NetworkManagementConnectivityTest
type NetworkManagementConnectivityTestStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkManagementConnectivityTest resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkManagementConnectivityTestObservedState `json:"observedState,omitempty"`
}

// NetworkManagementConnectivityTestObservedState is the state of the NetworkManagementConnectivityTest resource as most recently observed in GCP.
// +kcc:proto=google.cloud.networkmanagement.v1.ConnectivityTest
type NetworkManagementConnectivityTestObservedState struct {
	// Required. Source specification of the Connectivity Test.
	//
	//  You can use a combination of source IP address, virtual machine
	//  (VM) instance, or Compute Engine network to uniquely identify
	//  the source location.
	//
	//  Examples:
	//  If the source IP address is an internal IP address within a Google Cloud
	//  Virtual Private Cloud (VPC) network, then you must also specify the VPC
	//  network. Otherwise, specify the VM instance, which already contains its
	//  internal IP address and VPC network information.
	//
	//  If the source of the test is within an on-premises network, then you must
	//  provide the destination VPC network.
	//
	//  If the source endpoint is a Compute Engine VM instance with multiple
	//  network interfaces, the instance itself is not sufficient to identify the
	//  endpoint. So, you must also specify the source IP address or VPC network.
	//
	//  A reachability analysis proceeds even if the source location is
	//  ambiguous. However, the test result may include endpoints that you don't
	//  intend to test.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.source
	Source *EndpointObservedState `json:"source,omitempty"`

	// Output only. The display name of a Connectivity Test.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The time the test was created.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the test's configuration was updated.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The reachability details of this test from the latest run.
	//  The details are updated when creating a new test, updating an
	//  existing test, or triggering a one-time rerun of an existing test.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.reachability_details
	ReachabilityDetails *ReachabilityDetails `json:"reachabilityDetails,omitempty"`

	// Output only. The probing details of this test from the latest run, present
	//  for applicable tests only. The details are updated when creating a new
	//  test, updating an existing test, or triggering a one-time rerun of an
	//  existing test.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.probing_details
	ProbingDetails *ProbingDetails `json:"probingDetails,omitempty"`

	// Output only. The reachability details of this test from the latest run for
	//  the return path. The details are updated when creating a new test,
	//  updating an existing test, or triggering a one-time rerun of an existing
	//  test.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.return_reachability_details
	ReturnReachabilityDetails *ReachabilityDetails `json:"returnReachabilityDetails,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkmanagementconnectivitytest;gcpnetworkmanagementconnectivitytests
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkManagementConnectivityTest is the Schema for the NetworkManagementConnectivityTest API
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NetworkManagementConnectivityTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkManagementConnectivityTestSpec   `json:"spec,omitempty"`
	Status NetworkManagementConnectivityTestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkManagementConnectivityTestList contains a list of NetworkManagementConnectivityTest
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NetworkManagementConnectivityTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkManagementConnectivityTest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkManagementConnectivityTest{}, &NetworkManagementConnectivityTestList{})
}
