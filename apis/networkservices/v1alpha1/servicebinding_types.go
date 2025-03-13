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

var NetworkServicesServiceBindingGVK = GroupVersion.WithKind("NetworkServicesServiceBinding")

// NetworkServicesServiceBindingSpec defines the desired state of NetworkServicesServiceBinding
// +kcc:proto=google.cloud.networkservices.v1.ServiceBinding
type NetworkServicesServiceBindingSpec struct {

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceBinding.description
	Description *string `json:"description,omitempty"`

	// Required. The full service directory service name of the format
	//  /projects/*/locations/*/namespaces/*/services/*
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceBinding.service
	Service *string `json:"service,omitempty"`

	// Optional. Set of label tags associated with the ServiceBinding resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceBinding.labels
	Labels map[string]string `json:"labels,omitempty"`

	*Parent `json:",inline"`

	// The NetworkServicesServiceBinding name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

type Parent struct {
	// Required. The location of the application.
	Location string `json:"location,omitempty"`

	// Required. The host project of the application.
	ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// NetworkServicesServiceBindingStatus defines the config connector machine state of NetworkServicesServiceBinding
type NetworkServicesServiceBindingStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkServicesServiceBinding resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkServicesServiceBindingObservedState `json:"observedState,omitempty"`
}

// NetworkServicesServiceBindingObservedState is the state of the NetworkServicesServiceBinding resource as most recently observed in GCP.
// +kcc:proto=google.cloud.networkservices.v1.ServiceBinding
type NetworkServicesServiceBindingObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceBinding.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.ServiceBinding.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkservicesservicebinding;gcpnetworkservicesservicebindings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkServicesServiceBinding is the Schema for the NetworkServicesServiceBinding API
// +k8s:openapi-gen=true
type NetworkServicesServiceBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkServicesServiceBindingSpec   `json:"spec,omitempty"`
	Status NetworkServicesServiceBindingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkServicesServiceBindingList contains a list of NetworkServicesServiceBinding
type NetworkServicesServiceBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesServiceBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesServiceBinding{}, &NetworkServicesServiceBindingList{})
}
