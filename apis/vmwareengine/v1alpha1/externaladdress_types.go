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

var VMwareEngineExternalAddressGVK = GroupVersion.WithKind("VMwareEngineExternalAddress")

// VMwareEngineExternalAddressSpec defines the desired state of VMwareEngineExternalAddress
// +kcc:spec:proto=google.cloud.vmwareengine.v1.ExternalAddress
type VMwareEngineExternalAddressSpec struct {
	// The VMwareEngineExternalAddress name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The resource name of the private cloud to create a new external IP address in.
	// +required
	PrivateCloudRef *PrivateCloudRef `json:"privateCloudRef,omitempty"`

	// The internal IP address of a workload VM.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.internal_ip
	InternalIP *string `json:"internalIP,omitempty"`

	// User-provided description for this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.description
	Description *string `json:"description,omitempty"`
}

// VMwareEngineExternalAddressStatus defines the config connector machine state of VMwareEngineExternalAddress
type VMwareEngineExternalAddressStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMwareEngineExternalAddress resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMwareEngineExternalAddressObservedState `json:"observedState,omitempty"`
}

// VMwareEngineExternalAddressObservedState is the state of the VMwareEngineExternalAddress resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vmwareengine.v1.ExternalAddress
type VMwareEngineExternalAddressObservedState struct {
	// Output only. The resource name of this external IP address.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1-a/privateClouds/my-cloud/externalAddresses/my-address`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The external IP address of a workload VM.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.external_ip
	ExternalIP *string `json:"externalIP,omitempty"`

	// Output only. The state of the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.state
	State *string `json:"state,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ExternalAddress.uid
	UID *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmwareengineexternaladdress;gcpvmwareengineexternaladdresses
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMwareEngineExternalAddress is the Schema for the VMwareEngineExternalAddress API
// +k8s:openapi-gen=true
type VMwareEngineExternalAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMwareEngineExternalAddressSpec   `json:"spec,omitempty"`
	Status VMwareEngineExternalAddressStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMwareEngineExternalAddressList contains a list of VMwareEngineExternalAddress
type VMwareEngineExternalAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMwareEngineExternalAddress `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMwareEngineExternalAddress{}, &VMwareEngineExternalAddressList{})
}
