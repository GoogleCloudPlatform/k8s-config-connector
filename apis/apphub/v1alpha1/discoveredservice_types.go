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

var AppHubDiscoveredServiceGVK = GroupVersion.WithKind("AppHubDiscoveredService")

// AppHubDiscoveredServiceSpec defines the desired state of AppHubDiscoveredService
// +kcc:spec:proto=google.cloud.apphub.v1.DiscoveredService
type AppHubDiscoveredServiceSpec struct {
	Parent `json:",inline"`
	// The AppHubDiscoveredService name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// AppHubDiscoveredServiceStatus defines the config connector machine state of AppHubDiscoveredService
type AppHubDiscoveredServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AppHubDiscoveredService resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AppHubDiscoveredServiceObservedState `json:"observedState,omitempty"`
}

// +kcc:proto=google.cloud.apphub.v1.ServiceProperties
type ServiceProperties struct {
}

// +kcc:proto=google.cloud.apphub.v1.ServiceReference
type ServiceReference struct {
}

// +kcc:observedstate:proto=google.cloud.apphub.v1.ServiceProperties
type ServicePropertiesObservedState struct {
	// Output only. The service project identifier that the underlying cloud
	//  resource resides in.
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceProperties.gcp_project
	GcpProject *string `json:"gcpProject,omitempty"`

	// Output only. The location that the underlying resource resides in, for
	//  example, us-west1.
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceProperties.location
	Location *string `json:"location,omitempty"`

	// Output only. The location that the underlying resource resides in if it is
	//  zonal, for example, us-west1-a).
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceProperties.zone
	Zone *string `json:"zone,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.apphub.v1.ServiceReference
type ServiceReferenceObservedState struct {
	// Output only. The underlying resource URI (For example, URI of Forwarding
	//  Rule, URL Map, and Backend Service).
	// +kcc:proto:field=google.cloud.apphub.v1.ServiceReference.uri
	URI *string `json:"uri,omitempty"`
}

// AppHubDiscoveredServiceObservedState is the state of the AppHubDiscoveredService resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apphub.v1.DiscoveredService
type AppHubDiscoveredServiceObservedState struct {
	// Output only. Reference to an underlying networking resource that can
	//  comprise a Service. These are immutable.
	// +kcc:proto:field=google.cloud.apphub.v1.DiscoveredService.service_reference
	ServiceReference *ServiceReferenceObservedState `json:"serviceReference,omitempty"`

	// Output only. Properties of an underlying compute resource that can comprise
	//  a Service. These are immutable.
	// +kcc:proto:field=google.cloud.apphub.v1.DiscoveredService.service_properties
	ServiceProperties *ServicePropertiesObservedState `json:"serviceProperties,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpapphubdiscoveredservice;gcpapphubdiscoveredservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AppHubDiscoveredService is the Schema for the AppHubDiscoveredService API
// +k8s:openapi-gen=true
type AppHubDiscoveredService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AppHubDiscoveredServiceSpec   `json:"spec,omitempty"`
	Status AppHubDiscoveredServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AppHubDiscoveredServiceList contains a list of AppHubDiscoveredService
type AppHubDiscoveredServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppHubDiscoveredService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppHubDiscoveredService{}, &AppHubDiscoveredServiceList{})
}
