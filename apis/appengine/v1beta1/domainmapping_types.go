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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AppEngineDomainMappingGVK = GroupVersion.WithKind("AppEngineDomainMapping")

// AppEngineDomainMappingSpec defines the desired state of AppEngineDomainMapping
// +kcc:spec:proto=google.appengine.v1.DomainMapping
type AppEngineDomainMappingSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The AppEngineDomainMapping name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	// +kcc:proto:field=google.appengine.v1.DomainMapping.ssl_settings
	SslSettings *SSLSettings `json:"sslSettings,omitempty"`
}

// AppEngineDomainMappingStatus defines the config connector machine state of AppEngineDomainMapping
type AppEngineDomainMappingStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AppEngineDomainMapping resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AppEngineDomainMappingObservedState `json:"observedState,omitempty"`
}

// AppEngineDomainMappingObservedState is the state of the AppEngineDomainMapping resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.appengine.v1.DomainMapping
type AppEngineDomainMappingObservedState struct {
	// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.
	// +kcc:proto:field=google.appengine.v1.DomainMapping.name
	Name *string `json:"name,omitempty"`

	// The resource records required to configure this domain mapping. These records must be added
	// to the domain's DNS configuration in order to serve the application via this domain mapping.
	// +kcc:proto:field=google.appengine.v1.DomainMapping.resource_records
	ResourceRecords []ResourceRecord `json:"resourceRecords,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpappenginedomainmapping;gcpappenginedomainmappings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AppEngineDomainMapping is the Schema for the AppEngineDomainMapping API
// +k8s:openapi-gen=true
type AppEngineDomainMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AppEngineDomainMappingSpec   `json:"spec,omitempty"`
	Status AppEngineDomainMappingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AppEngineDomainMappingList contains a list of AppEngineDomainMapping
type AppEngineDomainMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppEngineDomainMapping `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppEngineDomainMapping{}, &AppEngineDomainMappingList{})
}
