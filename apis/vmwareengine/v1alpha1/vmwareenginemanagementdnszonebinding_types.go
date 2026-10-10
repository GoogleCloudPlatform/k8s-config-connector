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

// API sources for VMwareEngineManagementDNSZoneBinding, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/vmwareengine/v1/vmwareengine_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/solutions/vmware-as-a-service
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/solutions/vmware-as-a-service/docs/reference/rest/v1/projects.locations.privateClouds.managementDnsZoneBindings

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VMwareEngineManagementDNSZoneBindingGVK = GroupVersion.WithKind("VMwareEngineManagementDNSZoneBinding")

// VMwareEngineManagementDNSZoneBindingSpec defines the desired state of VMwareEngineManagementDNSZoneBinding
// +kcc:spec:proto=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding
// +kcc:required-from-proto
type VMwareEngineManagementDNSZoneBindingSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// A reference to the projects/{project}/locations/{location}/privateClouds/{private_cloud} this resource belongs to.
	// +kcc:guess
	// PrivateCloudRef *PrivateCloudRef `json:"privateCloudRef,omitempty"`

	// The VMwareEngineManagementDNSZoneBinding name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// User-provided description for this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.description
	Description *string `json:"description,omitempty"`

	// Network to bind is a standard consumer VPC.
	//  Specify the name in the following form for consumer
	//  VPC network: `projects/{project}/global/networks/{network_id}`.
	//  `{project}` can either be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.vpc_network
	VPCNetwork *string `json:"vpcNetwork,omitempty"`

	// Network to bind is a VMware Engine network.
	//  Specify the name in the following form for VMware engine network:
	//  `projects/{project}/locations/global/vmwareEngineNetworks/{vmware_engine_network_id}`.
	//  `{project}` can either be a project number or a project ID.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.vmware_engine_network
	VmwareEngineNetwork *string `json:"vmwareEngineNetwork,omitempty"`
}

// VMwareEngineManagementDNSZoneBindingStatus defines the config connector machine state of VMwareEngineManagementDNSZoneBinding
type VMwareEngineManagementDNSZoneBindingStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMwareEngineManagementDNSZoneBinding resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMwareEngineManagementDNSZoneBindingObservedState `json:"observedState,omitempty"`
}

// VMwareEngineManagementDNSZoneBindingObservedState is the state of the VMwareEngineManagementDNSZoneBinding resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding
type VMwareEngineManagementDNSZoneBindingObservedState struct {
	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.state
	State *string `json:"state,omitempty"`

	// Output only. System-generated unique identifier for the resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.ManagementDnsZoneBinding.uid
	Uid *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmwareenginemanagementdnszonebinding;gcpvmwareenginemanagementdnszonebindings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMwareEngineManagementDNSZoneBinding is the Schema for the VMwareEngineManagementDNSZoneBinding API
// +k8s:openapi-gen=true
type VMwareEngineManagementDNSZoneBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMwareEngineManagementDNSZoneBindingSpec   `json:"spec,omitempty"`
	Status VMwareEngineManagementDNSZoneBindingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMwareEngineManagementDNSZoneBindingList contains a list of VMwareEngineManagementDNSZoneBinding
type VMwareEngineManagementDNSZoneBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMwareEngineManagementDNSZoneBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMwareEngineManagementDNSZoneBinding{}, &VMwareEngineManagementDNSZoneBindingList{})
}
