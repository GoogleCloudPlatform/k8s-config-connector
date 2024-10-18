// Copyright 2024 Google LLC
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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VPCAccessConnectorGVK = GroupVersion.WithKind("VPCAccessConnector")

// VPCAccessConnectorSpec defines the desired state of VPCAccessConnector
// +kcc:proto=google.cloud.vpcaccess.v1.Connector
type VPCAccessConnectorSpec struct {
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	// Immutable.
	// The VPCAccessConnector name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. Name of a VPC network.  Required if ipCidrRange is set
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef"`

	// Immutable. The range of internal addresses that follows RFC 4632 notation.
	// Example: `10.132.0.0/28`.
	IPCidrRange *string `json:"ipCidrRange,omitempty"`

	// Immutable. Minimum throughput of the connector in Mbps. Default and min is 200.
	MinThroughput *int32 `json:"minThroughput,omitempty"`

	// Immutable. Maximum throughput of the connector in Mbps. Default is 300, max is 1000.
	MaxThroughput *int32 `json:"maxThroughput,omitempty"`

	// Immutable.  The subnet in which to house the VPC Access Connector.
	Subnet *ConnectorSubnet `json:"subnet,omitempty"`

	// Immutable. Machine type of VM Instance underlying connector. Default is e2-micro.
	MachineType *string `json:"machineType,omitempty"`

	// Immutable. Minimum value of instances in autoscaling group underlying the connector.
	MinInstances *int32 `json:"minInstances,omitempty"`

	// Immutable. Maximum value of instances in autoscaling group underlying the connector.
	MaxInstances *int32 `json:"maxInstances,omitempty"`

	// Location represents the geographical location of the VPCAccessConnector. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location"`

	// The project that this resource belongs to.
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`
}

type ConnectorSubnet struct {
	// Immutable. The subnet in which to house the connector.
	NameRef *refs.ComputeSubnetworkRef `json:"nameRef,omitempty"`

	// Immutable. Project in which the subnet exists. If
	// not set, this project is assumed to be the project for which
	// the connector create request was issued.
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`
}

// VPCAccessConnectorStatus defines the config connector machine state of VPCAccessConnector
type VPCAccessConnectorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* NOTYET
	// A unique specifier for the VPCAccessConnector resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VPCAccessConnectorObservedState `json:"observedState,omitempty"`
	*/

	// Output only. List of projects using the connector.
	ConnectedProjects []string `json:"connectedProjects,omitempty"`

	// The fully qualified name of this VPC connector.
	SelfLink string `json:"selfLink,omitempty"`

	// State of the VPC access connector.
	State string `json:"state,omitempty"`
}

// VPCAccessConnectorObservedState is the state of the VPCAccessConnector resource as most recently observed in GCP.
type VPCAccessConnectorObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpvpcaccessconnector;gcpvpcaccessconnectors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VPCAccessConnector is the Schema for the VPCAccessConnector API
// +k8s:openapi-gen=true
type VPCAccessConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VPCAccessConnectorSpec   `json:"spec,omitempty"`
	Status VPCAccessConnectorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VPCAccessConnectorList contains a list of VPCAccessConnector
type VPCAccessConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCAccessConnector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VPCAccessConnector{}, &VPCAccessConnectorList{})
}
