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

// API sources for AppHubService, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/apphub/v1/service.proto
// +kcc:source:service-docs=https://cloud.google.com/app-hub/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/app-hub/docs/reference/rest/v1/projects.locations.applications.services

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AppHubServiceGVK = GroupVersion.WithKind("AppHubService")

// AppHubServiceSpec defines the desired state of AppHubService
// +kcc:spec:proto=google.cloud.apphub.v1.Service
// +kcc:required-from-proto
type AppHubServiceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The AppHubService name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. User-defined name for the Service.
	//  Can have a maximum length of 63 characters.
	// +kcc:proto:field=google.cloud.apphub.v1.Service.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined description of a Service.
	//  Can have a maximum length of 2048 characters.
	// +kcc:proto:field=google.cloud.apphub.v1.Service.description
	Description *string `json:"description,omitempty"`

	// Optional. Consumer provided attributes.
	// +kcc:proto:field=google.cloud.apphub.v1.Service.attributes
	Attributes *Attributes `json:"attributes,omitempty"`

	// Required. Immutable. The resource name of the original discovered service.
	// +kcc:guess=possible-reference target=AppHubDiscoveredService
	// +kcc:proto:field=google.cloud.apphub.v1.Service.discovered_service
	// +required
	DiscoveredService *string `json:"discoveredService,omitempty"`
}

// AppHubServiceStatus defines the config connector machine state of AppHubService
type AppHubServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AppHubService resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AppHubServiceObservedState `json:"observedState,omitempty"`
}

// AppHubServiceObservedState is the state of the AppHubService resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apphub.v1.Service
type AppHubServiceObservedState struct {
	// Output only. Reference to an underlying networking resource that can
	//  comprise a Service. These are immutable.
	// +kcc:proto:field=google.cloud.apphub.v1.Service.service_reference
	ServiceReference *ServiceReferenceObservedState `json:"serviceReference,omitempty"`

	// Output only. Properties of an underlying compute resource that can comprise
	//  a Service. These are immutable.
	// +kcc:proto:field=google.cloud.apphub.v1.Service.service_properties
	ServiceProperties *ServicePropertiesObservedState `json:"serviceProperties,omitempty"`

	// Output only. Create time.
	// +kcc:proto:field=google.cloud.apphub.v1.Service.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time.
	// +kcc:proto:field=google.cloud.apphub.v1.Service.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A universally unique identifier (UUID) for the `Service` in
	//  the UUID4 format.
	// +kcc:proto:field=google.cloud.apphub.v1.Service.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Service state.
	// +kcc:proto:field=google.cloud.apphub.v1.Service.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapphubservice;gcpapphubservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AppHubService is the Schema for the AppHubService API
// +k8s:openapi-gen=true
type AppHubService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AppHubServiceSpec   `json:"spec,omitempty"`
	Status AppHubServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AppHubServiceList contains a list of AppHubService
type AppHubServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppHubService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppHubService{}, &AppHubServiceList{})
}
