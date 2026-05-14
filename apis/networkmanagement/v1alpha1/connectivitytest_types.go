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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	container "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	run "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkManagementConnectivityTestGVK = GroupVersion.WithKind("NetworkManagementConnectivityTest")

// NetworkManagementConnectivityTestSpec defines the desired state of NetworkManagementConnectivityTest
// +kcc:proto=google.cloud.networkmanagement.v1.ConnectivityTest
type NetworkManagementConnectivityTestSpec struct {
	// The project that this resource belongs to. If not provided, the provider project is used.
	// Resource name is in the format of `projects/{project_id}/locations/global/connectivityTests/{test_id}`.

	// Immutable. The Project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The name of the location where the RuntimeTemplate will be created.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location is immutable."
	Location string `json:"location"`

	// The NetworkManagementConnectivityTest name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

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
	Source *Endpoint `json:"source"`

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
	Destination *Endpoint `json:"destination"`

	// IP Protocol of the test. When not provided, "TCP" is assumed.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.protocol
	Protocol *string `json:"protocol,omitempty"`

	// Other projects that may be relevant for reachability analysis.
	//  This is applicable to scenarios where a test can cross project boundaries.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.related_projects
	RelatedProjects []refsv1beta1.ProjectRef `json:"relatedProjects,omitempty"`

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

// +kcc:spec:proto=google.cloud.networkmanagement.v1.Endpoint
type Endpoint struct {
	// The IP address of the endpoint, which can be an external or internal IP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// The IP protocol port of the endpoint.
	//  Only applicable when protocol is TCP or UDP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.port
	Port *int32 `json:"port,omitempty"`

	// A Compute Engine instance URI.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.instance
	ComputeInstanceRef *computev1beta1.InstanceRef `json:"computeInstanceRef,omitempty"`

	// TODO: Should be reference.

	// A forwarding rule and its corresponding IP address represent the frontend
	//  configuration of a Google Cloud load balancer. Forwarding rules are also
	//  used for protocol forwarding, Private Service Connect and other network
	//  services to provide forwarding information in the control plane. Format:
	//   projects/{project}/global/forwardingRules/{id} or
	//   projects/{project}/regions/{region}/forwardingRules/{id}
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.forwarding_rule
	ComputeForwardingRuleRef *string `json:"computeForwardingRuleRef,omitempty"`

	// A cluster URI for [Google Kubernetes Engine cluster control
	//  plane](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-architecture).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.gke_master_cluster
	ContainerClusterRef *container.ContainerClusterRef `json:"containerClusterRef,omitempty"`

	// DNS endpoint of [Google Kubernetes Engine cluster control
	//  plane](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-architecture).
	//  Requires gke_master_cluster to be set, can't be used simultaneoulsly with
	//  ip_address or network. Applicable only to destination endpoint.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.fqdn
	FQDN *string `json:"fqdn,omitempty"`

	// A [Cloud SQL](https://cloud.google.com/sql) instance URI.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.cloud_sql_instance
	SQLInstanceRef *refsv1beta1.SQLInstanceRef `json:"sqlInstance,omitempty"`

	// TODO: Should be reference.

	// A [Redis Instance](https://cloud.google.com/memorystore/docs/redis)
	//  URI.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.redis_instance
	RedisInstance *string `json:"redisInstance,omitempty"`

	// TODO: Should be reference.

	// A [Redis Cluster](https://cloud.google.com/memorystore/docs/cluster)
	//  URI.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.redis_cluster
	RedisCluster *string `json:"redisCluster,omitempty"`

	// TODO: Should be reference.

	// A [Cloud Function](https://cloud.google.com/functions).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.cloud_function
	CloudFunction *Endpoint_CloudFunctionEndpoint `json:"cloudFunction,omitempty"`

	// An [App Engine](https://cloud.google.com/appengine) [service
	//  version](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.app_engine_version
	AppEngineVersion *Endpoint_AppEngineVersionEndpoint `json:"appEngineVersion,omitempty"`

	// A [Cloud Run](https://cloud.google.com/run)
	//  [revision](https://cloud.google.com/run/docs/reference/rest/v1/namespaces.revisions/get)
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.cloud_run_revision
	CloudRunRevision *Endpoint_CloudRunRevisionEndpoint `json:"cloudRunRevision,omitempty"`

	// A Compute Engine network URI.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.network
	ComputeNetworkRef *computev1beta1.ComputeNetworkRef `json:"computeNetworkRef,omitempty"`

	// Type of the network where the endpoint is located.
	//  Applicable only to source endpoint, as destination network type can be
	//  inferred from the source.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.network_type
	NetworkType *string `json:"networkType,omitempty"`

	// Project ID where the endpoint is located.
	//  The Project ID can be derived from the URI if you provide a VM instance or
	//  network URI.
	//  The following are two cases where you must provide the project ID:
	//  1. Only the IP address is specified, and the IP address is within a Google
	//  Cloud project.
	//  2. When you are using Shared VPC and the IP address that you provide is
	//  from the service project. In this case, the network that the IP address
	//  resides in is defined in the host project.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.project_id
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// +kcc:spec:proto=google.cloud.networkmanagement.v1.Endpoint.CloudRunRevisionEndpoint
type Endpoint_CloudRunRevisionEndpoint struct {
	// A [Cloud Run](https://cloud.google.com/run)
	//  [revision](https://cloud.google.com/run/docs/reference/rest/v1/namespaces.revisions/get)
	//  URI. The format is:
	//  projects/{project}/locations/{location}/revisions/{revision}
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.CloudRunRevisionEndpoint.uri
	RunRevisionRef *run.RevisionRef `json:"runRevisionRef,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.networkmanagement.v1.Endpoint.CloudRunRevisionEndpoint
type Endpoint_CloudRunRevisionEndpointObservedState struct {
	// Output only. The URI of the Cloud Run service that the revision belongs
	//  to. The format is:
	//  `projects/{project}/locations/{location}/services/{service}`
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.CloudRunRevisionEndpoint.service_uri
	ServiceURI *string `json:"serviceURI,omitempty"`
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
	// Source specification of the Connectivity Test.
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

	// Destination specification of the Connectivity Test.
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
	Destination *EndpointObservedState `json:"destination,omitempty"`

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
	ReachabilityDetails *ReachabilityDetailsObservedState `json:"reachabilityDetails,omitempty"`

	// Output only. The probing details of this test from the latest run, present
	//  for applicable tests only. The details are updated when creating a new
	//  test, updating an existing test, or triggering a one-time rerun of an
	//  existing test.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.probing_details
	ProbingDetails *ProbingDetailsObservedState `json:"probingDetails,omitempty"`

	// Output only. The reachability details of this test from the latest run for
	//  the return path. The details are updated when creating a new test,
	//  updating an existing test, or triggering a one-time rerun of an existing
	//  test.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ConnectivityTest.return_reachability_details
	ReturnReachabilityDetails *ReachabilityDetailsObservedState `json:"returnReachabilityDetails,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.Endpoint
type EndpointObservedState struct {
	// Output only. Specifies the type of the target of the forwarding rule.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.forwarding_rule_target
	ForwardingRuleTarget *string `json:"forwardingRuleTarget,omitempty"`

	// Output only. ID of the load balancer the forwarding rule points to. Empty
	//  for forwarding rules not related to load balancers.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.load_balancer_id
	LoadBalancerID *string `json:"loadBalancerID,omitempty"`

	// Output only. Type of the load balancer the forwarding rule points to.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.load_balancer_type
	LoadBalancerType *string `json:"loadBalancerType,omitempty"`
}

// +kcc:proto=google.rpc.Status
type StatusObservedState struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// Commented out: We don't want to try to expose the Any type.
	// // A list of messages that carry the error details.  There is a common set of
	// //  message types for APIs to use.
	// // +kcc:proto:field=google.rpc.Status.details
	// Details []Any `json:"details,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkmanagementconnectivitytest;gcpnetworkmanagementconnectivitytests
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
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
