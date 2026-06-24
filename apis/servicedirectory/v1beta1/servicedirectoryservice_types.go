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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ServiceDirectoryServiceGVK = GroupVersion.WithKind("ServiceDirectoryService")

// ServiceDirectoryServiceSpec defines the desired state of ServiceDirectoryService
// +kcc:spec:proto=google.cloud.servicedirectory.v1beta1.Service
type ServiceDirectoryServiceSpec struct {
	// The ServiceDirectoryNamespace that this service belongs to.
	NamespaceRef ServiceDirectoryNamespaceRef `json:"namespaceRef"`

	// Immutable. Optional. The serviceId of the resource.
	// Used for creation and acquisition. When unset, the value of `metadata.name`
	// is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`
}

// ServiceDirectoryServiceStatus defines the config connector machine state of ServiceDirectoryService
type ServiceDirectoryServiceStatus struct {
	// Conditions represent the latest available observations of the
	// object's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The resource name for the service
	// in the format 'projects/*/locations/*/namespaces/*/services/*'.
	Name *string `json:"name,omitempty"`
}

// ServiceDirectoryServiceObservedState is the state of the ServiceDirectoryService resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.servicedirectory.v1beta1.Service
type ServiceDirectoryServiceObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpservicedirectoryservice;gcpservicedirectoryservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ServiceDirectoryService is the Schema for the ServiceDirectoryService API
// +k8s:openapi-gen=true
type ServiceDirectoryService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ServiceDirectoryServiceSpec   `json:"spec,omitempty"`
	Status ServiceDirectoryServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ServiceDirectoryServiceList contains a list of ServiceDirectoryService
type ServiceDirectoryServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDirectoryService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceDirectoryService{}, &ServiceDirectoryServiceList{})
}
