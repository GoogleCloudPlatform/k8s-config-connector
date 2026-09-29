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
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	computev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkServicesAgentGatewayGVK = GroupVersion.WithKind("NetworkServicesAgentGateway")

// NetworkServicesAgentGatewaySpec defines the desired state of NetworkServicesAgentGateway
// +kcc:spec:proto=google.cloud.networkservices.v1.AgentGateway
type NetworkServicesAgentGatewaySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetworkServicesAgentGateway name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Proxy is orchestrated and managed by GoogleCloud in a tenant
	//  project.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.google_managed
	GoogleManaged *AgentGateway_GoogleManaged `json:"googleManaged,omitempty"`

	// Optional. Attach to existing Application Load Balancers or Secure Web
	//  Proxies.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.self_managed
	SelfManaged *AgentGateway_SelfManaged `json:"selfManaged,omitempty"`

	// Optional. Set of label tags associated with the AgentGateway resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.description
	Description *string `json:"description,omitempty"`

	// Optional. A list of Agent registries containing the agents, MCP servers and
	//  tools governed by the Agent Gateway. Note: Currently limited to
	//  project-scoped registries Must be of format
	//  `//agentregistry.googleapis.com/projects/{project}/locations/{location}/`
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.registries
	Registries []string `json:"registries,omitempty"`

	// Optional. Network configuration for the AgentGateway.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.network_config
	NetworkConfig *AgentGateway_NetworkConfig `json:"networkConfig,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.AgentGateway.SelfManaged
type AgentGateway_SelfManaged struct {
	// Optional. A supported Google Cloud networking proxy in the Project and
	//  Location
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.SelfManaged.resource_uri
	ResourceURI *string `json:"resourceURI,omitempty"`

	// Optional. List of supported Google Cloud networking proxies in the Project and
	//  Location.
	//  resource_uris is mutually exclusive with resource_uri.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.SelfManaged.resource_uris
	ResourceURIs []string `json:"resourceURIs,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.AgentGateway.NetworkConfig.DnsPeeringConfig
type AgentGateway_NetworkConfig_DNSPeeringConfig struct {
	// Required. Domain names for which DNS queries should be forwarded to the
	//  target network.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.NetworkConfig.DnsPeeringConfig.domains
	Domains []string `json:"domains,omitempty"`

	// Required. Target project ID to which DNS queries should be forwarded
	//  to. This can be the same project that contains the AgentGateway or a
	//  different project.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.NetworkConfig.DnsPeeringConfig.target_project
	TargetProject *string `json:"targetProject,omitempty"`

	// Required. Target network in 'target project' to which DNS queries
	//  should be forwarded to. Must be in format of
	//  `projects/{project}/global/networks/{network}`.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.NetworkConfig.DnsPeeringConfig.target_network
	TargetNetworkRef *computerefs.ComputeNetworkRef `json:"targetNetworkRef,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.AgentGateway.NetworkConfig.Egress
type AgentGateway_NetworkConfig_Egress struct {
	// Optional. The URI of the Network Attachment resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.NetworkConfig.Egress.network_attachment
	NetworkAttachmentRef *computev1alpha1.ComputeNetworkAttachmentRef `json:"networkAttachmentRef,omitempty"`

	// Optional. TrustConfig defines the trust configuration for egress.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.NetworkConfig.Egress.trust_config
	TrustConfig *AgentGateway_NetworkConfig_Egress_TrustConfig `json:"trustConfig,omitempty"`
}

// NetworkServicesAgentGatewayStatus defines the config connector machine state of NetworkServicesAgentGateway
type NetworkServicesAgentGatewayStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkServicesAgentGateway resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkServicesAgentGatewayObservedState `json:"observedState,omitempty"`
}

// NetworkServicesAgentGatewayObservedState is the state of the NetworkServicesAgentGateway resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networkservices.v1.AgentGateway
type NetworkServicesAgentGatewayObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Field for populated AgentGateway card.
	// +kcc:proto:field=google.cloud.networkservices.v1.AgentGateway.agent_gateway_card
	AgentGatewayCard *AgentGateway_AgentGatewayOutputCardObservedState `json:"agentGatewayCard,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkservicesagentgateway;gcpnetworkservicesagentgateways
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkServicesAgentGateway is the Schema for the NetworkServicesAgentGateway API
// +k8s:openapi-gen=true
type NetworkServicesAgentGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkServicesAgentGatewaySpec   `json:"spec,omitempty"`
	Status NetworkServicesAgentGatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkServicesAgentGatewayList contains a list of NetworkServicesAgentGateway
type NetworkServicesAgentGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesAgentGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesAgentGateway{}, &NetworkServicesAgentGatewayList{})
}
