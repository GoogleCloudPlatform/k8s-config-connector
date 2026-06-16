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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ServiceDirectoryEndpointGVK = GroupVersion.WithKind("ServiceDirectoryEndpoint")

// ServiceDirectoryEndpointSpec defines the desired state of ServiceDirectoryEndpoint
// +kcc:spec:proto=google.cloud.servicedirectory.v1beta1.Endpoint
type ServiceDirectoryEndpointSpec struct {
	// Required. The ServiceDirectoryService that this endpoint belongs to.
	// +kcc:proto:field=google.cloud.servicedirectory.v1beta1.Endpoint.service
	ServiceRef ServiceDirectoryServiceRef `json:"serviceRef"`

	// Optional. The ComputeAddress resource whose IP address this endpoint uses.
	// +kcc:proto:field=google.cloud.servicedirectory.v1beta1.Endpoint.address
	AddressRef *computev1beta1.ComputeAddressRef `json:"addressRef,omitempty"`

	// Optional. The ComputeNetwork resource representing the GCE network of this endpoint.
	// +kcc:proto:field=google.cloud.servicedirectory.v1beta1.Endpoint.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Port that the endpoint is running on, must be in the
	// range of [0, 65535]. If unspecified, the default is 0.
	// +kcc:proto:field=google.cloud.servicedirectory.v1beta1.Endpoint.port
	Port *int32 `json:"port,omitempty"`

	// Immutable. Optional. The endpointId of the resource.
	// Used for creation and acquisition. When unset, the value of `metadata.name`
	// is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ServiceDirectoryEndpointStatus defines the config connector machine state of ServiceDirectoryEndpoint
type ServiceDirectoryEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The resource name for the endpoint in the format
	// 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
	Name *string `json:"name,omitempty"`
}

// ServiceDirectoryEndpointObservedState is the state of the ServiceDirectoryEndpoint resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.servicedirectory.v1beta1.Endpoint
type ServiceDirectoryEndpointObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpservicedirectoryendpoint;gcpservicedirectoryendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ServiceDirectoryEndpoint is the Schema for the ServiceDirectoryEndpoint API
// +k8s:openapi-gen=true
type ServiceDirectoryEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ServiceDirectoryEndpointSpec   `json:"spec,omitempty"`
	Status ServiceDirectoryEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ServiceDirectoryEndpointList contains a list of ServiceDirectoryEndpoint
type ServiceDirectoryEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDirectoryEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceDirectoryEndpoint{}, &ServiceDirectoryEndpointList{})
}
