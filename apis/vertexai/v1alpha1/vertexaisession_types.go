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

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

var VertexAISessionGVK = GroupVersion.WithKind("VertexAISession")

// VertexAISessionSpec defines the desired state of VertexAISession
type VertexAISessionSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The parent ReasoningEngine of this resource.
	// +required
	ReasoningEngineRef *VertexAIReasoningEngineRef `json:"reasoningEngineRef"`

	// The VertexAISession name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Timestamp of when this session is considered expired.
	//  This is *always* provided on output, regardless of what was sent
	//  on input.
	//  The minimum value is 24 hours from the time of creation.
	ExpireTime *string `json:"expireTime,omitempty"`

	// Optional. Input only. The TTL for this session.
	//  The minimum value is 24 hours.
	TTL *string `json:"ttl,omitempty"`

	// Optional. The display name of the session.
	DisplayName *string `json:"displayName,omitempty"`

	// The labels with user-defined metadata to organize your Sessions.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Session specific memory which stores key conversation points.
	SessionState apiextensionsv1.JSON `json:"sessionState,omitempty"`

	// Required. Immutable. String id provided by the user
	UserID *string `json:"userID,omitempty"`
}

// VertexAISessionStatus defines the config connector machine state of VertexAISession
type VertexAISessionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAISession resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAISessionObservedState `json:"observedState,omitempty"`
}

// VertexAISessionObservedState is the state of the VertexAISession resource as most recently observed in GCP.
type VertexAISessionObservedState struct {
	// Output only. Timestamp when the session was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the session was updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaisession;gcpvertexaisessions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAISession is the Schema for the VertexAISession API
// +k8s:openapi-gen=true
type VertexAISession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAISessionSpec   `json:"spec,omitempty"`
	Status VertexAISessionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAISessionList contains a list of VertexAISession
type VertexAISessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAISession `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAISession{}, &VertexAISessionList{})
}
