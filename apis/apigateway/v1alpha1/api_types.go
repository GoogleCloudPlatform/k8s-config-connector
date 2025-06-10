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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var APIGatewayAPIGVK = GroupVersion.WithKind("APIGatewayAPI")

// APIGatewayAPISpec defines the desired state of APIGatewayAPI
// +kcc:spec:proto=google.cloud.apigateway.v1.Api
type APIGatewayAPISpec struct {
	// The APIGatewayAPI name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Resource labels to represent user-provided metadata.
	//  Refer to cloud documentation on labels for more details.
	//  https://cloud.google.com/compute/docs/labeling-resources
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Display name.
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Immutable. The name of a Google Managed Service (
	//  https://cloud.google.com/service-infrastructure/docs/glossary#managed). If
	//  not specified, a new Service will automatically be created in the same
	//  project as this API.
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.managed_service
	ManagedService *string `json:"managedService,omitempty"` //  TODO: Before promote to v1beta. Check if this field should be a ref to  https://cloud.google.com/service-infrastructure/docs/service-management/reference/rest/v1/services#ManagedService.

	//  Optional. The project that this resource belongs to.
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`
}

// APIGatewayAPIStatus defines the config connector machine state of APIGatewayAPI
type APIGatewayAPIStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIGatewayAPI resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *APIGatewayAPIObservedState `json:"observedState,omitempty"`
}

// APIGatewayAPIObservedState is the state of the APIGatewayAPI resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apigateway.v1.Api
type APIGatewayAPIObservedState struct {

	// Resource name of the API.
	//  Format: projects/{project}/locations/global/apis/{api}
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.name
	Name *string `json:"name,omitempty"`

	// Created time.
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Updated time.
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	//  State of the API.
	// +kcc:proto:field=google.cloud.apigateway.v1.Api.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigatewayapi;gcpapigatewayapis
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIGatewayAPI is the Schema for the APIGatewayAPI API
// +k8s:openapi-gen=true
type APIGatewayAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIGatewayAPISpec   `json:"spec,omitempty"`
	Status APIGatewayAPIStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIGatewayAPIList contains a list of APIGatewayAPI
type APIGatewayAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIGatewayAPI `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIGatewayAPI{}, &APIGatewayAPIList{})
}
