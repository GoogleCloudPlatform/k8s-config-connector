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

var VMwareEngineNetworkGVK = GroupVersion.WithKind("VMwareEngineNetwork")

type Parent struct {
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// VMwareEngineNetworkSpec defines the desired state of VMwareEngineNetwork
// +kcc:proto=google.cloud.vmwareengine.v1.VmwareEngineNetwork
type VMwareEngineNetworkSpec struct {
	// The VMwareEngineNetwork name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	Parent `json:",inline"`

	// User-provided description for this VMware Engine network.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.description
	Description *string `json:"description,omitempty"`

	// Required. VMware Engine network type.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.type
	// +required
	Type *string `json:"type,omitempty"`

	// Checksum that may be sent on update and delete requests to ensure that the
	//  user-provided value is up to date before the server processes a request.
	//  The server computes checksums based on the value of other fields in the
	//  request.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.etag
	Etag *string `json:"etag,omitempty"`
}

// VMwareEngineNetworkStatus defines the config connector machine state of VMwareEngineNetwork
type VMwareEngineNetworkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMwareEngineNetwork resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMwareEngineNetworkObservedState `json:"observedState,omitempty"`
}

// VMwareEngineNetworkObservedState is the state of the VMwareEngineNetwork resource as most recently observed in GCP.
// +kcc:proto=google.cloud.vmwareengine.v1.VmwareEngineNetwork
type VMwareEngineNetworkObservedState struct {
	// Output only. The resource name of the VMware Engine network.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/global/vmwareEngineNetworks/my-network`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. VMware Engine service VPC networks that provide connectivity
	//  from a private cloud to customer projects, the internet, and other Google
	//  Cloud services.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.vpc_networks
	VPCNetworks []VmwareEngineNetwork_VpcNetworkObservedState `json:"vpcNetworks,omitempty"`

	// Output only. State of the VMware Engine network.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.state
	State *string `json:"state,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.VmwareEngineNetwork.uid
	UID *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmwareenginenetwork;gcpvmwareenginenetworks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMwareEngineNetwork is the Schema for the VMwareEngineNetwork API
// +k8s:openapi-gen=true
type VMwareEngineNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMwareEngineNetworkSpec   `json:"spec,omitempty"`
	Status VMwareEngineNetworkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMwareEngineNetworkList contains a list of VMwareEngineNetwork
type VMwareEngineNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMwareEngineNetwork `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMwareEngineNetwork{}, &VMwareEngineNetworkList{})
}
