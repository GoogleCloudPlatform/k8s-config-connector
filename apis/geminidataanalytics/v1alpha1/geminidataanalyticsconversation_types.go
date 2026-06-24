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
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var GeminiDataAnalyticsConversationGVK = GroupVersion.WithKind("GeminiDataAnalyticsConversation")

var _ refsv1beta1.ExternalNormalizer = &GeminiDataAnalyticsDataAgentRef{}

// GeminiDataAnalyticsDataAgentRef is a reference to a DataAgent.
type GeminiDataAnalyticsDataAgentRef struct {
	// A reference to an externally managed DataAgent resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/dataAgents/{{dataAgentID}}".
	External string `json:"external,omitempty"`
}

func (r *GeminiDataAnalyticsDataAgentRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External == "" {
		return "", fmt.Errorf("external is required on DataAgent reference")
	}
	// Validate format: projects/{project}/locations/{location}/dataAgents/{dataAgentID}
	if !strings.HasPrefix(r.External, "projects/") || !strings.Contains(r.External, "/locations/") || !strings.Contains(r.External, "/dataAgents/") {
		return "", fmt.Errorf("invalid DataAgent external reference format: %q", r.External)
	}
	return r.External, nil
}

// GeminiDataAnalyticsConversationSpec defines the desired state of GeminiDataAnalyticsConversation
// +kcc:spec:proto=google.cloud.geminidataanalytics.v1beta.Conversation
type GeminiDataAnalyticsConversationSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The GeminiDataAnalyticsConversation name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Agent(s) in the conversation.
	// Currently, only one agent is supported. This field is repeated to allow
	// for future support of multiple agents in a conversation.
	// +required
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1beta.Conversation.agents
	AgentRefs []GeminiDataAnalyticsDataAgentRef `json:"agentRefs"`

	// Optional. Open-ended and user-defined labels that can be set by the client
	// to tag a conversation.
	// +optional
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1beta.Conversation.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// GeminiDataAnalyticsConversationStatus defines the config connector machine state of GeminiDataAnalyticsConversation
type GeminiDataAnalyticsConversationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GeminiDataAnalyticsConversation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GeminiDataAnalyticsConversationObservedState `json:"observedState,omitempty"`
}

// GeminiDataAnalyticsConversationObservedState is the state of the GeminiDataAnalyticsConversation resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.geminidataanalytics.v1beta.Conversation
type GeminiDataAnalyticsConversationObservedState struct {
	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1beta.Conversation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp of the last used conversation.
	// +kcc:proto:field=google.cloud.geminidataanalytics.v1beta.Conversation.last_used_time
	LastUsedTime *string `json:"lastUsedTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgeminidataanalyticsconversation;gcpgeminidataanalyticsconversations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GeminiDataAnalyticsConversation is the Schema for the GeminiDataAnalyticsConversation API
// +k8s:openapi-gen=true
type GeminiDataAnalyticsConversation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GeminiDataAnalyticsConversationSpec   `json:"spec,omitempty"`
	Status GeminiDataAnalyticsConversationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GeminiDataAnalyticsConversationList contains a list of GeminiDataAnalyticsConversation
type GeminiDataAnalyticsConversationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GeminiDataAnalyticsConversation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GeminiDataAnalyticsConversation{}, &GeminiDataAnalyticsConversationList{})
}
