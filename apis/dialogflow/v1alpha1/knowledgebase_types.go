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

var DialogflowKnowledgeBaseGVK = GroupVersion.WithKind("DialogflowKnowledgeBase")

// DialogflowKnowledgeBaseSpec defines the desired state of DialogflowKnowledgeBase
// +kcc:spec:proto=google.cloud.dialogflow.v2.KnowledgeBase
type DialogflowKnowledgeBaseSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location,omitempty"`

	// The DialogflowKnowledgeBase name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +required
	// Required. The display name of the knowledge base. The name must be 1024
	//  bytes or less; otherwise, the creation request fails.
	// +kcc:proto:field=google.cloud.dialogflow.v2.KnowledgeBase.display_name
	DisplayName *string `json:"displayName"`

	// Language which represents the KnowledgeBase. When the KnowledgeBase is
	//  created/updated, expect this to be present for non en-us languages. When
	//  unspecified, the default language code en-us applies.
	// +kcc:proto:field=google.cloud.dialogflow.v2.KnowledgeBase.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// DialogflowKnowledgeBaseStatus defines the config connector machine state of DialogflowKnowledgeBase
type DialogflowKnowledgeBaseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowKnowledgeBase resource in Google Cloud.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	ObservedState *DialogflowKnowledgeBaseObservedState `json:"observedState,omitempty"`
}

// DialogflowKnowledgeBaseObservedState is the state of the DialogflowKnowledgeBase resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2.KnowledgeBase
type DialogflowKnowledgeBaseObservedState struct {
	// The knowledge base resource name.
	//  The name must be empty when creating a knowledge base.
	//  Format: `projects/<Project ID>/locations/<Location ID>/knowledgeBases/<Knowledge Base ID>`.
	// +kcc:proto:field=google.cloud.dialogflow.v2.KnowledgeBase.name
	Name *string `json:"name,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowknowledgebase;gcpdialogflowknowledgebases
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowKnowledgeBase is the Schema for the DialogflowKnowledgeBase API
// +k8s:openapi-gen=true
type DialogflowKnowledgeBase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowKnowledgeBaseSpec   `json:"spec,omitempty"`
	Status DialogflowKnowledgeBaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowKnowledgeBaseList contains a list of DialogflowKnowledgeBase
type DialogflowKnowledgeBaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowKnowledgeBase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowKnowledgeBase{}, &DialogflowKnowledgeBaseList{})
}
