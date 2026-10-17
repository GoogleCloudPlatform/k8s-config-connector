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

// API sources for AgentregistryService, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/agentregistry/v1/service.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/agent-registry/overview
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/agent-registry/overview/docs/reference/rest/v1/projects.locations.services

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AgentregistryServiceGVK = GroupVersion.WithKind("AgentregistryService")

// AgentregistryServiceSpec defines the desired state of AgentregistryService
// +kcc:spec:proto=google.cloud.agentregistry.v1.Service
// +kcc:required-from-proto
type AgentregistryServiceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The AgentregistryService name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The spec of the Agent. When `agent_spec` is set, the type of
	//  the service is Agent.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Service.agent_spec
	AgentSpec *Service_AgentSpec `json:"agentSpec,omitempty"`

	// Optional. The spec of the MCP Server. When `mcp_server_spec` is set, the
	//  type of the service is MCP Server.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Service.mcp_server_spec
	McpServerSpec *Service_McpServerSpec `json:"mcpServerSpec,omitempty"`

	// Optional. The spec of the Endpoint. When `endpoint_spec` is set, the type
	//  of the service is Endpoint.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Service.endpoint_spec
	EndpointSpec *Service_EndpointSpec `json:"endpointSpec,omitempty"`

	// Optional. User-defined display name for the Service.
	//  Can have a maximum length of `63` characters.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Service.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined description of an Service.
	//  Can have a maximum length of `2048` characters.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Service.description
	Description *string `json:"description,omitempty"`

	// Optional. The connection details for the Service.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Service.interfaces
	Interfaces []Interface `json:"interfaces,omitempty"`
}

// AgentregistryServiceStatus defines the config connector machine state of AgentregistryService
type AgentregistryServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AgentregistryService resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AgentregistryServiceObservedState `json:"observedState,omitempty"`
}

// AgentregistryServiceObservedState is the state of the AgentregistryService resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.agentregistry.v1.Service
type AgentregistryServiceObservedState struct {
	// Output only. The resource name of the resulting Agent, MCP Server, or
	//  Endpoint. Format:
	//
	//  * `projects/{project}/locations/{location}/mcpServers/{mcp_server}`
	//  * `projects/{project}/locations/{location}/agents/{agent}`
	//  * `projects/{project}/locations/{location}/endpoints/{endpoint}`
	// +kcc:proto:field=google.cloud.agentregistry.v1.Service.registry_resource
	RegistryResource *string `json:"registryResource,omitempty"`

	// Output only. Create time.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Service.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time.
	// +kcc:proto:field=google.cloud.agentregistry.v1.Service.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpagentregistryservice;gcpagentregistryservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AgentregistryService is the Schema for the AgentregistryService API
// +k8s:openapi-gen=true
type AgentregistryService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AgentregistryServiceSpec   `json:"spec,omitempty"`
	Status AgentregistryServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AgentregistryServiceList contains a list of AgentregistryService
type AgentregistryServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AgentregistryService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AgentregistryService{}, &AgentregistryServiceList{})
}
