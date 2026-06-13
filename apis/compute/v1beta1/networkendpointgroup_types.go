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

package v1beta1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeRegionNetworkEndpointGroupGVK = GroupVersion.WithKind("ComputeRegionNetworkEndpointGroup")

// +kcc:proto=google.cloud.compute.v1.NetworkEndpointGroupCloudFunction
type NetworkEndpointGroupCloudFunction struct {
	// A user-defined name of the Cloud Function. The function name is case-sensitive and must be 1-63 characters long. Example value: func1.
	// +optional
	FunctionRef *refsv1beta1.CloudFunctionsFunctionRef `json:"functionRef,omitempty"`

	// An URL mask is one of the main components of the Cloud Function. A template to parse function field from a request URL. URL mask allows for routing to multiple Cloud Functions without having to create multiple Network Endpoint Groups and backend services. For example, request URLs mydomain.com/function1 and mydomain.com/function2 can be backed by the same Serverless NEG with URL mask /<function>. The URL mask will parse them to { function = "function1" } and { function = "function2" } respectively.
	// +optional
	URLMask *string `json:"urlMask,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.NetworkEndpointGroupCloudRun
type NetworkEndpointGroupCloudRun struct {
	// Cloud Run service is the main resource of Cloud Run. The service must be 1-63 characters long, and comply with RFC1035. Example value: "run-service".
	// +optional
	ServiceRef *refsv1beta1.RunServiceRef `json:"serviceRef,omitempty"`

	// Optional Cloud Run tag represents the "named-revision" to provide additional fine-grained traffic routing information. The tag must be 1-63 characters long, and comply with RFC1035. Example value: "revision-0010".
	// +optional
	Tag *string `json:"tag,omitempty"`

	// An URL mask is one of the main components of the Cloud Function. A template to parse <service> and <tag> fields from a request URL. URL mask allows for routing to multiple Run services without having to create multiple network endpoint groups and backend services. For example, request URLs foo1.domain.com/bar1 and foo1.domain.com/bar2 can be backed by the same Serverless Network Endpoint Group (NEG) with URL mask <tag>.domain.com/<service>. The URL mask will parse them to { service="bar1", tag="foo1" } and { service="bar2", tag="foo2" } respectively.
	// +optional
	URLMask *string `json:"urlMask,omitempty"`
}

// ComputeRegionNetworkEndpointGroupSpec defines the desired state of ComputeRegionNetworkEndpointGroup
// +kcc:spec:proto=google.cloud.compute.v1.NetworkEndpointGroup
type ComputeRegionNetworkEndpointGroupSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ComputeRegionNetworkEndpointGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Only valid when networkEndpointType is SERVERLESS. Only one of cloudRun, appEngine or cloudFunction may be set.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.cloud_function
	CloudFunction *NetworkEndpointGroupCloudFunction `json:"cloudFunction,omitempty"`

	// Optional. Only valid when networkEndpointType is SERVERLESS. Only one of cloudRun, appEngine or cloudFunction may be set.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.cloud_run
	CloudRun *NetworkEndpointGroupCloudRun `json:"cloudRun,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.description
	Description *string `json:"description,omitempty"`

	// Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT, GCE_VM_IP_PORTMAP.
	//  Check the NetworkEndpointType enum for the list of possible values.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.network_endpoint_type
	NetworkEndpointType *string `json:"networkEndpointType,omitempty"`

	// The URL of the network to which all network endpoints in the NEG belong. Uses default project network if unspecified.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.network
	NetworkRef *ComputeNetworkRef `json:"networkRef,omitempty"`

	// The target service url used to set up private service connection to a Google API or a PSC Producer Service Attachment. An example value is: asia-northeast3-cloudkms.googleapis.com. Optional. Only valid when networkEndpointType is PRIVATE_SERVICE_CONNECT.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.psc_target_service
	PSCTargetService *string `json:"pscTargetService,omitempty"`

	// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.NetworkEndpointGroup.subnetwork
	SubnetworkRef *ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
}

// ComputeRegionNetworkEndpointGroupStatus defines the config connector machine state of ComputeRegionNetworkEndpointGroup
type ComputeRegionNetworkEndpointGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeRegionNetworkEndpointGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeRegionNetworkEndpointGroupObservedState `json:"observedState,omitempty"`

	// The self link of the ComputeRegionNetworkEndpointGroup.
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// ComputeRegionNetworkEndpointGroupObservedState is the state of the ComputeRegionNetworkEndpointGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.NetworkEndpointGroup
type ComputeRegionNetworkEndpointGroupObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeregionnetworkendpointgroup;gcpcomputeregionnetworkendpointgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeRegionNetworkEndpointGroup is the Schema for the ComputeRegionNetworkEndpointGroup API
// +k8s:openapi-gen=true
type ComputeRegionNetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeRegionNetworkEndpointGroupSpec   `json:"spec,omitempty"`
	Status ComputeRegionNetworkEndpointGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeRegionNetworkEndpointGroupList contains a list of ComputeRegionNetworkEndpointGroup
type ComputeRegionNetworkEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRegionNetworkEndpointGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRegionNetworkEndpointGroup{}, &ComputeRegionNetworkEndpointGroupList{})
}
