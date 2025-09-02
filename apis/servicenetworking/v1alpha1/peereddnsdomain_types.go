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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ServiceNetworkingPeeredDNSDomainGVK = GroupVersion.WithKind("ServiceNetworkingPeeredDNSDomain")

// ServiceNetworkingPeeredDNSDomainSpec defines the desired state of ServiceNetworkingPeeredDNSDomain
// +kcc:spec:proto=mockgcp.cloud.servicenetworking.v1.PeeredDnsDomain
type ServiceNetworkingPeeredDNSDomainSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`

	// The ServiceNetworkingPeeredDNSDomain name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The DNS domain name suffix e.g. `example.com.`. Cloud DNS requires that a DNS suffix ends with a trailing dot.
	// +kcc:proto:field=mockgcp.cloud.servicenetworking.v1.PeeredDnsDomain.dns_suffix
	DNSSuffix *string `json:"dnsSuffix,omitempty"`

	// // Required. User assigned name for this resource. Must be unique within the consumer network. The name must be 1-63 characters long, must begin with a letter, end with a letter or digit, and only contain lowercase letters, digits or dashes.
	// // +kcc:proto:field=mockgcp.cloud.servicenetworking.v1.PeeredDnsDomain.name
	// Name *string `json:"name,omitempty"`
}

// ServiceNetworkingPeeredDNSDomainStatus defines the config connector machine state of ServiceNetworkingPeeredDNSDomain
type ServiceNetworkingPeeredDNSDomainStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ServiceNetworkingPeeredDNSDomain resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ServiceNetworkingPeeredDNSDomainObservedState `json:"observedState,omitempty"`
}

// ServiceNetworkingPeeredDNSDomainObservedState is the state of the ServiceNetworkingPeeredDNSDomain resource as most recently observed in GCP.
// +kcc:observedstate:proto=mockgcp.cloud.servicenetworking.v1.PeeredDnsDomain
type ServiceNetworkingPeeredDNSDomainObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpservicenetworkingpeereddnsdomain;gcpservicenetworkingpeereddnsdomains
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ServiceNetworkingPeeredDNSDomain is the Schema for the ServiceNetworkingPeeredDNSDomain API
// +k8s:openapi-gen=true
type ServiceNetworkingPeeredDNSDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ServiceNetworkingPeeredDNSDomainSpec   `json:"spec,omitempty"`
	Status ServiceNetworkingPeeredDNSDomainStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ServiceNetworkingPeeredDNSDomainList contains a list of ServiceNetworkingPeeredDNSDomain
type ServiceNetworkingPeeredDNSDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceNetworkingPeeredDNSDomain `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceNetworkingPeeredDNSDomain{}, &ServiceNetworkingPeeredDNSDomainList{})
}
