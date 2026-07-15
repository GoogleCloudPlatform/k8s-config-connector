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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AgentRegistryServiceGVK = GroupVersion.WithKind("AgentRegistryService")

// Declaring dummy variable to keep the unused import of apiextensionsv1 if types.generated.go is compiled separately.
var _ = apiextensionsv1.JSON{}

// +kcc:proto=google.cloud.agentregistry.v1.Interface
type Interface struct {
	// Required. The destination URL.
	// +required
	URL *string `json:"url,omitempty"`

	// Required. The protocol binding of the interface.
	// +required
	ProtocolBinding *string `json:"protocolBinding,omitempty"`
}

// +kcc:proto=google.cloud.agentregistry.v1.Service.AgentSpec
type Service_AgentSpec struct {
	// Required. The type of the agent spec content.
	// +required
	// +kubebuilder:validation:Enum=TYPE_UNSPECIFIED;NO_SPEC;A2A_AGENT_CARD
	Type *string `json:"type,omitempty"`

	// Optional. The content of the Agent spec in the JSON format.
	// This payload is validated against the schema for the specified type.
	// The content size is limited to 10KB.
	// +optional
	Content apiextensionsv1.JSON `json:"content,omitempty"`
}

// +kcc:proto=google.cloud.agentregistry.v1.Service.EndpointSpec
type Service_EndpointSpec struct {
	// Required. The type of the endpoint spec content.
	// +required
	// +kubebuilder:validation:Enum=TYPE_UNSPECIFIED;NO_SPEC
	Type *string `json:"type,omitempty"`

	// Optional. The content of the endpoint spec.
	// Reserved for future use.
	// +optional
	Content apiextensionsv1.JSON `json:"content,omitempty"`
}

// +kcc:proto=google.cloud.agentregistry.v1.Service.McpServerSpec
type Service_McpServerSpec struct {
	// Required. The type of the MCP Server spec content.
	// +required
	// +kubebuilder:validation:Enum=TYPE_UNSPECIFIED;NO_SPEC;TOOL_SPEC
	Type *string `json:"type,omitempty"`

	// Optional. The content of the MCP Server spec.
	// This payload is validated against the schema for the specified type.
	// The content size is limited to 10KB.
	// +optional
	Content apiextensionsv1.JSON `json:"content,omitempty"`
}

// AgentRegistryServiceSpec defines the desired state of AgentRegistryService
// +kcc:spec:proto=google.cloud.agentregistry.v1.Service
type AgentRegistryServiceSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	Location string `json:"location"`

	// Optional. The spec of the Agent. When `agentSpec` is set, the type of
	// the service is Agent.
	// +optional
	AgentSpec *Service_AgentSpec `json:"agentSpec,omitempty"`

	// Optional. The spec of the MCP Server. When `mcpServerSpec` is set, the
	// type of the service is MCP Server.
	// +optional
	McpServerSpec *Service_McpServerSpec `json:"mcpServerSpec,omitempty"`

	// Optional. The spec of the Endpoint. When `endpointSpec` is set, the type
	// of the service is Endpoint.
	// +optional
	EndpointSpec *Service_EndpointSpec `json:"endpointSpec,omitempty"`

	// Optional. User-defined display name for the Service.
	// Can have a maximum length of 63 characters.
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined description of a Service.
	// Can have a maximum length of 2048 characters.
	// +optional
	Description *string `json:"description,omitempty"`

	// Optional. The connection details for the Service.
	// +optional
	Interfaces []Interface `json:"interfaces,omitempty"`

	// The AgentRegistryService name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// AgentRegistryServiceStatus defines the config connector machine state of AgentRegistryService
type AgentRegistryServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AgentRegistryService resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AgentRegistryServiceObservedState `json:"observedState,omitempty"`
}

// AgentRegistryServiceObservedState is the state of the AgentRegistryService resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.agentregistry.v1.Service
type AgentRegistryServiceObservedState struct {
	// Output only. The resource name of the resulting Agent, MCP Server, or
	// Endpoint. Format:
	// * `projects/{project}/locations/{location}/mcpServers/{mcp_server}`
	// * `projects/{project}/locations/{location}/agents/{agent}`
	// * `projects/{project}/locations/{location}/endpoints/{endpoint}`
	// +optional
	RegistryResource *string `json:"registryResource,omitempty"`

	// Output only. Create time.
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpagentregistryservice;gcpagentregistryservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AgentRegistryService is the Schema for the AgentRegistryService API
// +k8s:openapi-gen=true
type AgentRegistryService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AgentRegistryServiceSpec   `json:"spec,omitempty"`
	Status AgentRegistryServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AgentRegistryServiceList contains a list of AgentRegistryService
type AgentRegistryServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentRegistryService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AgentRegistryService{}, &AgentRegistryServiceList{})
}
