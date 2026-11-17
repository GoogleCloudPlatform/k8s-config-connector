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

// API sources for AgentIdentityAuthProvider, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/2c4dd791cc44104f1fdc49b4cbdadb160772b11e/google/cloud/agentidentity/v1/auth_provider_service.proto
// +kcc:source:service-docs=https://cloud.google.com/iam/docs/

package v1alpha1

import (
	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AgentIdentityAuthProviderGVK = GroupVersion.WithKind("AgentIdentityAuthProvider")

// AgentIdentityAuthProviderSpec defines the desired state of AgentIdentityAuthProvider
// +kcc:spec:proto=google.cloud.agentidentity.v1.AuthProvider
// +kcc:required-from-proto
type AgentIdentityAuthProviderSpec struct {
	*parent.ProjectAndLocationRef `json:",inline"`

	// The AgentIdentityAuthProvider name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. AuthProvider type specific parameters.
	// +kcc:proto:field=google.cloud.agentidentity.v1.AuthProvider.auth_provider_type_params
	// +required
	AuthProviderTypeParams *AuthProvider_AuthProviderTypeParams `json:"authProviderTypeParams,omitempty"`

	// Optional. List of scopes that are allowed to be requested for this
	//  auth_provider. If this list is non-empty, only scopes within this list may
	//  be requested. If this list is empty, all scopes may be requested. Scopes
	//  appearing in `blocked_scopes` are disallowed even if they appear in
	//  `allowed_scopes`.
	//  The number of allowed scopes is limited to 200.
	// +kcc:proto:field=google.cloud.agentidentity.v1.AuthProvider.allowed_scopes
	AllowedScopes []string `json:"allowedScopes,omitempty"`

	// Optional. List of scopes that are blocked from being requested for this
	//  auth_provider. If a scope appears in this list, it will not be requested,
	//  even if it also appears in `allowed_scopes`. `blocked_scopes` takes
	//  precedence over `allowed_scopes`. The number of blocked scopes is limited
	//  to 200.
	// +kcc:proto:field=google.cloud.agentidentity.v1.AuthProvider.blocked_scopes
	BlockedScopes []string `json:"blockedScopes,omitempty"`

	// Optional. Description of the resource.
	//  Must be less than 256 characters.
	// +kcc:proto:field=google.cloud.agentidentity.v1.AuthProvider.description
	Description *string `json:"description,omitempty"`

	// Optional. Input only. Represents the workload identity in IAM
	//  `principal://` format of the agent(s) that will use this AuthProvider.
	//  Example:
	//  `principal://agents.global.org-${ORG_ID}.system.id.goog/resources/aiplatform/projects/{PROJECT_ID}/locations/{LOCATIONS}/reasoningEngines/{ID}`
	// +kcc:proto:field=google.cloud.agentidentity.v1.AuthProvider.workload_ids
	WorkloadIDs []string `json:"workloadIDs,omitempty"`
}

// AgentIdentityAuthProviderStatus defines the config connector machine state of AgentIdentityAuthProvider
type AgentIdentityAuthProviderStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AgentIdentityAuthProvider resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AgentIdentityAuthProviderObservedState `json:"observedState,omitempty"`
}

// AgentIdentityAuthProviderObservedState is the state of the AgentIdentityAuthProvider resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.agentidentity.v1.AuthProvider
type AgentIdentityAuthProviderObservedState struct {
	// Output only. [Output only] Create time stamp
	// +kcc:proto:field=google.cloud.agentidentity.v1.AuthProvider.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. [Output only] Update time stamp
	// +kcc:proto:field=google.cloud.agentidentity.v1.AuthProvider.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Required. AuthProvider type specific parameters.
	// +kcc:proto:field=google.cloud.agentidentity.v1.AuthProvider.auth_provider_type_params
	AuthProviderTypeParams *AuthProvider_AuthProviderTypeParamsObservedState `json:"authProviderTypeParams,omitempty"`

	// Output only. This is set to true if the auth_provider is deleted.
	// +kcc:proto:field=google.cloud.agentidentity.v1.AuthProvider.deleted
	Deleted *bool `json:"deleted,omitempty"`

	// Output only. The time when the auth_provider will expire.
	// +kcc:proto:field=google.cloud.agentidentity.v1.AuthProvider.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. The state of the auth_provider.
	// +kcc:proto:field=google.cloud.agentidentity.v1.AuthProvider.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpagentidentityauthprovider;gcpagentidentityauthproviders
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AgentIdentityAuthProvider is the Schema for the AgentIdentityAuthProvider API
// +k8s:openapi-gen=true
type AgentIdentityAuthProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AgentIdentityAuthProviderSpec   `json:"spec,omitempty"`
	Status AgentIdentityAuthProviderStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AgentIdentityAuthProviderList contains a list of AgentIdentityAuthProvider
type AgentIdentityAuthProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentIdentityAuthProvider `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AgentIdentityAuthProvider{}, &AgentIdentityAuthProviderList{})
}
