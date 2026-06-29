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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowSipTrunkGVK = GroupVersion.WithKind("DialogflowSipTrunk")

// DialogflowSipTrunkSpec defines the desired state of DialogflowSipTrunk
// +kcc:spec:proto=google.cloud.dialogflow.v2beta1.SipTrunk
type DialogflowSipTrunkSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	Location *string `json:"location"`

	// The DialogflowSipTrunk name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The expected hostnames in the peer certificate from partner that is used for TLS authentication.
	// +kubebuilder:validation:Required
	ExpectedHostname []string `json:"expectedHostname"`

	// Optional. Human readable alias for this trunk.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty"`
}

// DialogflowSipTrunkStatus defines the config connector machine state of DialogflowSipTrunk
type DialogflowSipTrunkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowSipTrunk resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowSipTrunkObservedState `json:"observedState,omitempty"`
}

// DialogflowSipTrunkObservedState is the state of the DialogflowSipTrunk resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2beta1.SipTrunk
type DialogflowSipTrunkObservedState struct {
	// Output only. Connections of the SIP trunk.
	Connections []ConnectionObservedState `json:"connections,omitempty"`
}

type ConnectionObservedState struct {
	// Output only. The unique identifier of the SIP Trunk connection.
	ConnectionID *string `json:"connectionID,omitempty"`

	// Output only. State of the connection.
	State *string `json:"state,omitempty"`

	// Output only. When the connection status changed.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The error details for the connection. Only populated when authentication errors occur.
	ErrorDetails *Connection_ErrorDetailsObservedState `json:"errorDetails,omitempty"`
}

type Connection_ErrorDetailsObservedState struct {
	// Output only. The status of the certificate authentication.
	CertificateState *string `json:"certificateState,omitempty"`

	// The error message provided from SIP trunking auth service.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowsiptrunk;gcpdialogflowsiptrunks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowSipTrunk is the Schema for the DialogflowSipTrunk API
// +k8s:openapi-gen=true
type DialogflowSipTrunk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowSipTrunkSpec   `json:"spec,omitempty"`
	Status DialogflowSipTrunkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowSipTrunkList contains a list of DialogflowSipTrunk
type DialogflowSipTrunkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowSipTrunk `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowSipTrunk{}, &DialogflowSipTrunkList{})
}
