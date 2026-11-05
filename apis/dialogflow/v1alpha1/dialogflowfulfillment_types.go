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

package v1alpha1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowFulfillmentGVK = GroupVersion.WithKind("DialogflowFulfillment")

// +kcc:proto=google.cloud.dialogflow.v2.Fulfillment.Feature
type Fulfillment_Feature struct {
	// The type of the feature that enabled for fulfillment.
	//  * SMALLTALK: Fulfillment is enabled for SmallTalk. Possible values: ["SMALLTALK"].
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.Feature.type
	// +required
	Type *string `json:"type"`
}

// +kcc:proto=google.cloud.dialogflow.v2.Fulfillment.GenericWebService
type Fulfillment_GenericWebService struct {
	// The password for HTTP Basic authentication.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.GenericWebService.password
	// +optional
	Password *string `json:"password,omitempty"`

	// The HTTP request headers to send together with fulfillment requests.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.GenericWebService.request_headers
	// +optional
	RequestHeaders map[string]string `json:"requestHeaders,omitempty"`

	// The fulfillment URI for receiving POST requests. It must use https protocol.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.GenericWebService.uri
	// +required
	URI *string `json:"uri"`

	// The user name for HTTP Basic authentication.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.GenericWebService.username
	// +optional
	Username *string `json:"username,omitempty"`
}

// +kcc:spec:proto=google.cloud.dialogflow.v2.Fulfillment
type DialogflowFulfillmentSpec struct {
	// The human-readable name of the fulfillment, unique within the agent.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.display_name
	// +required
	DisplayName *string `json:"displayName"`

	// Whether fulfillment is enabled.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.enabled
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	// The field defines whether the fulfillment is enabled for certain features.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.features
	// +optional
	Features []Fulfillment_Feature `json:"features,omitempty"`

	// Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.generic_web_service
	// +optional
	GenericWebService *Fulfillment_GenericWebService `json:"genericWebService,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type DialogflowFulfillmentStatus struct {
	/* Conditions represent the latest available observation of the
	   resource's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// The unique identifier of the fulfillment.
	// Format: projects/<Project ID>/agent/fulfillment - projects/<Project ID>/locations/<Location ID>/agent/fulfillment.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Fulfillment.name
	// +optional
	Name *string `json:"name,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowFulfillment resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`
}

// DialogflowFulfillmentObservedState is the state of the DialogflowFulfillment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2.Fulfillment
type DialogflowFulfillmentObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowfulfillment;gcpdialogflowfulfillments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowFulfillment is the Schema for the DialogflowFulfillment API
// +k8s:openapi-gen=true
type DialogflowFulfillment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowFulfillmentSpec   `json:"spec,omitempty"`
	Status DialogflowFulfillmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowFulfillmentList contains a list of DialogflowFulfillment
type DialogflowFulfillmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowFulfillment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowFulfillment{}, &DialogflowFulfillmentList{})
}
