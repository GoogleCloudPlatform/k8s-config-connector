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

var AppHubApplicationGVK = GroupVersion.WithKind("AppHubApplication")

// AppHubApplicationSpec defines the desired state of AppHubApplication
// +kcc:proto=google.cloud.apphub.v1.Application
type AppHubApplicationSpec struct {
	// Identifier. The resource name of an Application. Format:
	//  "projects/{host-project-id}/locations/{location}/applications/{application-id}"
	// +kcc:proto:field=google.cloud.apphub.v1.Application.name
	Name *string `json:"name,omitempty"`

	// Optional. User-defined name for the Application.
	//  Can have a maximum length of 63 characters.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined description of an Application.
	//  Can have a maximum length of 2048 characters.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.description
	Description *string `json:"description,omitempty"`

	// Optional. Consumer provided attributes.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.attributes
	Attributes *Attributes `json:"attributes,omitempty"`

	// Required. Immutable. Defines what data can be included into this
	//  Application. Limits which Services and Workloads can be registered.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.scope
	Scope *Scope `json:"scope,omitempty"`

	// The AppHubApplication name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// AppHubApplicationStatus defines the config connector machine state of AppHubApplication
type AppHubApplicationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AppHubApplication resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AppHubApplicationObservedState `json:"observedState,omitempty"`
}

// AppHubApplicationObservedState is the state of the AppHubApplication resource as most recently observed in GCP.
// +kcc:proto=google.cloud.apphub.v1.Application
type AppHubApplicationObservedState struct {
	// Output only. Create time.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A universally unique identifier (in UUID4 format) for the
	//  `Application`.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Application state.
	// +kcc:proto:field=google.cloud.apphub.v1.Application.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpapphubapplication;gcpapphubapplications
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AppHubApplication is the Schema for the AppHubApplication API
// +k8s:openapi-gen=true
type AppHubApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AppHubApplicationSpec   `json:"spec,omitempty"`
	Status AppHubApplicationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AppHubApplicationList contains a list of AppHubApplication
type AppHubApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppHubApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppHubApplication{}, &AppHubApplicationList{})
}
