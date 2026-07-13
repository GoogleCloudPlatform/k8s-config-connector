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

var AgentRegistryBindingGVK = GroupVersion.WithKind("AgentRegistryBinding")

// AgentRegistryBindingSpec defines the desired state of AgentRegistryBinding
// +kcc:spec:proto=google.cloud.agentregistry.v1.Binding
type AgentRegistryBindingSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The AgentRegistryBinding name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. User-defined display name for the Binding.
	// Can have a maximum length of 63 characters.
	// +optional
	// +kcc:proto:field=google.cloud.agentregistry.v1.Binding.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined description of a Binding.
	// Can have a maximum length of 2048 characters.
	// +optional
	// +kcc:proto:field=google.cloud.agentregistry.v1.Binding.description
	Description *string `json:"description,omitempty"`

	// The binding for AuthProvider.
	// +optional
	// +kcc:proto:field=google.cloud.agentregistry.v1.Binding.auth_provider_binding
	AuthProviderBinding *AgentRegistryBindingAuthProviderBinding `json:"authProviderBinding,omitempty"`

	// Required. The target Agent of the Binding.
	// +required
	// +kcc:proto:field=google.cloud.agentregistry.v1.Binding.source
	Source *AgentRegistryBindingSource `json:"source,omitempty"`

	// Required. The target Agent Registry Resource of the Binding.
	// +required
	// +kcc:proto:field=google.cloud.agentregistry.v1.Binding.target
	Target *AgentRegistryBindingTarget `json:"target,omitempty"`
}

// +kcc:proto=google.cloud.agentregistry.v1.Binding.AuthProviderBinding
type AgentRegistryBindingAuthProviderBinding struct {
	// Required. The resource name of the target AuthProvider.
	// Format: `projects/{project}/locations/{location}/authProviders/{auth_provider}`
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.agentregistry.v1.Binding.AuthProviderBinding.auth_provider
	AuthProvider *string `json:"authProvider"`

	// Optional. The list of OAuth2 scopes of the AuthProvider.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.agentregistry.v1.Binding.AuthProviderBinding.scopes
	Scopes []string `json:"scopes,omitempty"`

	// Optional. The continue URI of the AuthProvider.
	// The URI is used to reauthenticate the user and finalize the managed OAuth
	// flow.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.agentregistry.v1.Binding.AuthProviderBinding.continue_uri
	ContinueURI *string `json:"continueURI,omitempty"`
}

// +kcc:proto=google.cloud.agentregistry.v1.Binding.Source
type AgentRegistryBindingSource struct {
	// The identifier of the source Agent.
	// Format: `urn:agent:{publisher}:{namespace}:{name}`
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.agentregistry.v1.Binding.Source.identifier
	Identifier *string `json:"identifier"`
}

// +kcc:proto=google.cloud.agentregistry.v1.Binding.Target
type AgentRegistryBindingTarget struct {
	// The identifier of the target Agent, MCP Server, or Endpoint.
	// Format:
	// * `urn:agent:{publisher}:{namespace}:{name}`
	// * `urn:mcp:{publisher}:{namespace}:{name}`
	// * `urn:endpoint:{publisher}:{namespace}:{name}`
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.agentregistry.v1.Binding.Target.identifier
	Identifier *string `json:"identifier"`
}

// AgentRegistryBindingStatus defines the config connector machine state of AgentRegistryBinding
type AgentRegistryBindingStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AgentRegistryBinding resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AgentRegistryBindingObservedState `json:"observedState,omitempty"`
}

// AgentRegistryBindingObservedState is the state of the AgentRegistryBinding resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.agentregistry.v1.Binding
type AgentRegistryBindingObservedState struct {
	// Output only. Timestamp when this binding was created.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Binding.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this binding was last updated.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Binding.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpagentregistrybinding;gcpagentregistrybindings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AgentRegistryBinding is the Schema for the AgentRegistryBinding API
// +k8s:openapi-gen=true
type AgentRegistryBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AgentRegistryBindingSpec   `json:"spec,omitempty"`
	Status AgentRegistryBindingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AgentRegistryBindingList contains a list of AgentRegistryBinding
type AgentRegistryBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentRegistryBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AgentRegistryBinding{}, &AgentRegistryBindingList{})
}
