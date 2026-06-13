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
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VMwareEnginePrivateConnectionGVK = GroupVersion.WithKind("VMwareEnginePrivateConnection")

// VMwareEnginePrivateConnectionSpec defines the desired state of VMwareEnginePrivateConnection
// +kcc:spec:proto=google.cloud.vmwareengine.v1.PrivateConnection
type VMwareEnginePrivateConnectionSpec struct {
	// The VMwareEnginePrivateConnection name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// Optional. User-provided description for this private connection.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.description
	Description *string `json:"description,omitempty"`

	// Required. Reference to the VMware Engine network.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.vmware_engine_network
	// +required
	VMwareEngineNetworkRef *VmwareEngineNetworkRef `json:"vmwareEngineNetworkRef"`

	// Required. Private connection type.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.type
	// +kubebuilder:validation:Enum=PRIVATE_SERVICE_ACCESS;NETAPP_CLOUD_VOLUME;DELL_POWERSCALE;THIRD_PARTY_SERVICE
	// +required
	Type *string `json:"type"`

	// Optional. Routing Mode.
	//  Default value is set to GLOBAL.
	//  For type = PRIVATE_SERVICE_ACCESS, this field can be set to GLOBAL or
	//  REGIONAL, for other types only GLOBAL is supported.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.routing_mode
	// +kubebuilder:validation:Enum=GLOBAL;REGIONAL
	RoutingMode *string `json:"routingMode,omitempty"`

	// Required. Reference to the service network to create private connection.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.service_network
	// +required
	ServiceNetworkRef *computerefs.ComputeNetworkRef `json:"serviceNetworkRef"`
}

// VMwareEnginePrivateConnectionStatus defines the config connector machine state of VMwareEnginePrivateConnection
type VMwareEnginePrivateConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMwareEnginePrivateConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMwareEnginePrivateConnectionObservedState `json:"observedState,omitempty"`
}

// VMwareEnginePrivateConnectionObservedState is the state of the VMwareEnginePrivateConnection resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vmwareengine.v1.PrivateConnection
type VMwareEnginePrivateConnectionObservedState struct {
	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the private connection.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.state
	State *string `json:"state,omitempty"`

	// Output only. The canonical name of the VMware Engine network in the form:
	//  `projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.vmware_engine_network_canonical
	VmwareEngineNetworkCanonical *string `json:"vmwareEngineNetworkCanonical,omitempty"`

	// Output only. VPC network peering id between given network VPC and
	//  VMwareEngineNetwork.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.peering_id
	PeeringID *string `json:"peeringID,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Peering state between service network and VMware Engine
	//  network.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.PrivateConnection.peering_state
	PeeringState *string `json:"peeringState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmwareengineprivateconnection;gcpvmwareengineprivateconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMwareEnginePrivateConnection is the Schema for the VMwareEnginePrivateConnection API
// +k8s:openapi-gen=true
type VMwareEnginePrivateConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMwareEnginePrivateConnectionSpec   `json:"spec,omitempty"`
	Status VMwareEnginePrivateConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMwareEnginePrivateConnectionList contains a list of VMwareEnginePrivateConnection
type VMwareEnginePrivateConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMwareEnginePrivateConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMwareEnginePrivateConnection{}, &VMwareEnginePrivateConnectionList{})
}
