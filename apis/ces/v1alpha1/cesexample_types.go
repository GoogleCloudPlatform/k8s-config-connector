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

// API sources for CESExample, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/ces/v1beta/example.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/docs/reference/rest/v1beta/projects.locations.apps.examples

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CESExampleGVK = GroupVersion.WithKind("CESExample")

// CESExampleSpec defines the desired state of CESExample
// +kcc:spec:proto=google.cloud.ces.v1beta.Example
// +kcc:required-from-proto
type CESExampleSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The CESExample name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Display name of the example.
	// +kcc:proto:field=google.cloud.ces.v1beta.Example.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Human-readable description of the example.
	// +kcc:proto:field=google.cloud.ces.v1beta.Example.description
	Description *string `json:"description,omitempty"`

	// Optional. The agent that initially handles the conversation. If not
	//  specified, the example represents a conversation that is handled by the
	//  root agent. Format:
	//  `projects/{project}/locations/{location}/apps/{app}/agents/{agent}`
	// +kcc:proto:field=google.cloud.ces.v1beta.Example.entry_agent
	EntryAgent *string `json:"entryAgent,omitempty"`

	// Optional. The collection of messages that make up the conversation.
	// +kcc:proto:field=google.cloud.ces.v1beta.Example.messages
	Messages []Message `json:"messages,omitempty"`

	// Etag used to ensure the object hasn't changed during a read-modify-write
	//  operation. If the etag is empty, the update will overwrite any concurrent
	//  changes.
	// +kcc:proto:field=google.cloud.ces.v1beta.Example.etag
	Etag *string `json:"etag,omitempty"`
}

// CESExampleStatus defines the config connector machine state of CESExample
type CESExampleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CESExample resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CESExampleObservedState `json:"observedState,omitempty"`
}

// CESExampleObservedState is the state of the CESExample resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ces.v1beta.Example
type CESExampleObservedState struct {
	// Optional. The collection of messages that make up the conversation.
	// +kcc:proto:field=google.cloud.ces.v1beta.Example.messages
	Messages []MessageObservedState `json:"messages,omitempty"`

	// Output only. Timestamp when the example was created.
	// +kcc:proto:field=google.cloud.ces.v1beta.Example.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the example was last updated.
	// +kcc:proto:field=google.cloud.ces.v1beta.Example.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The example may become invalid if referencing resources are
	//  deleted. Invalid examples will not be used as few-shot examples.
	// +kcc:proto:field=google.cloud.ces.v1beta.Example.invalid
	Invalid *bool `json:"invalid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcesexample;gcpcesexamples
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CESExample is the Schema for the CESExample API
// +k8s:openapi-gen=true
type CESExample struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CESExampleSpec   `json:"spec,omitempty"`
	Status CESExampleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CESExampleList contains a list of CESExample
type CESExampleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CESExample `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CESExample{}, &CESExampleList{})
}
