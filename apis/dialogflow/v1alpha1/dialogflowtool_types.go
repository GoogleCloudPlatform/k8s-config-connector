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

// API sources for DialogflowTool, recorded by generate-types:
// +kcc:source:proto=google.cloud.dialogflow.cx.v3beta1.Tool
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3beta1/projects.locations.agents.tools

package v1alpha1

import (
	connectorsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connectors/v1alpha1"
	discoveryenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	servicedirectoryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowToolGVK = GroupVersion.WithKind("DialogflowTool")

// DialogflowToolSpec defines the desired state of DialogflowTool
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3beta1.Tool
// +kcc:required-from-proto
type DialogflowToolSpec struct {
	// The DialogflowCXAgent this DialogflowTool belongs to.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *DialogflowCXAgentRef `json:"parentRef"`

	// The DialogflowTool name. If not given, the server will generate a unique identifier.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The human-readable name of the Tool, unique within an agent.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. High level description of the Tool and its usage.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.description
	// +required
	Description *string `json:"description,omitempty"`

	// OpenAPI specification of the Tool.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.open_api_spec
	OpenAPISpec *Tool_OpenAPITool `json:"openAPISpec,omitempty"`

	// Data store search tool specification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.data_store_spec
	DataStoreSpec *Tool_DataStoreTool `json:"dataStoreSpec,omitempty"`

	// Vertex extension tool specification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.extension_spec
	ExtensionSpec *Tool_ExtensionTool `json:"extensionSpec,omitempty"`

	// Client side executed function specification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.function_spec
	FunctionSpec *Tool_FunctionTool `json:"functionSpec,omitempty"`

	// Integration connectors tool specification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.connector_spec
	ConnectorSpec *Tool_ConnectorTool `json:"connectorSpec,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.OpenApiTool
type Tool_OpenAPITool struct {
	// Required. The OpenAPI schema specified as a text.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.OpenApiTool.text_schema
	// +required
	TextSchema *string `json:"textSchema,omitempty"`

	// Optional. Authentication information required by the API.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.OpenApiTool.authentication
	Authentication *Tool_Authentication `json:"authentication,omitempty"`

	// Optional. TLS configuration for the HTTPS verification.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.OpenApiTool.tls_config
	TLSConfig *Tool_TLSConfig `json:"tlsConfig,omitempty"`

	// Optional. Service Directory configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.OpenApiTool.service_directory_config
	ServiceDirectoryConfig *Tool_ServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.ServiceDirectoryConfig
type Tool_ServiceDirectoryConfig struct {
	// Required. The name of Service Directory service.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.ServiceDirectoryConfig.service
	// +required
	ServiceRef *servicedirectoryv1beta1.ServiceDirectoryServiceRef `json:"serviceRef,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool
type Tool_DataStoreTool struct {
	// Required. List of data stores to search.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool.data_store_connections
	DataStoreConnections []Tool_DataStoreTool_DataStoreConnection `json:"dataStoreConnections,omitempty"`

	// Optional. Fallback prompt config.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool.fallback_prompt
	FallbackPrompt *Tool_DataStoreTool_FallbackPrompt `json:"fallbackPrompt,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection
type Tool_DataStoreTool_DataStoreConnection struct {
	// Optional. The type of the data store.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection.data_store_type
	DataStoreType *string `json:"dataStoreType,omitempty"`

	// Optional. The full name of the referenced data store.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection.data_store
	DataStoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`

	// Optional. Document processing mode.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection.document_processing_mode
	DocumentProcessingMode *string `json:"documentProcessingMode,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool.FallbackPrompt
type Tool_DataStoreTool_FallbackPrompt struct {
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.ExtensionTool
type Tool_ExtensionTool struct {
	// Required. The full resource name of the referenced extension.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.ExtensionTool.name
	ExtensionRef *VertexAIExtensionRef `json:"extensionRef,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.ConnectorTool
type Tool_ConnectorTool struct {
	// Required. The full resource name of the Integration Connectors Connection.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.ConnectorTool.name
	ConnectionRef *connectorsv1alpha1.ConnectorsConnectionRef `json:"connectionRef,omitempty"`

	// Optional. Actions supported by the Connector tool.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.ConnectorTool.actions
	Actions []Tool_ConnectorTool_Action `json:"actions,omitempty"`

	// Optional. End user authentication configuration.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.ConnectorTool.end_user_auth_config
	EndUserAuthConfig *Tool_EndUserAuthConfig `json:"endUserAuthConfig,omitempty"`
}

// DialogflowToolStatus defines the config connector machine state of DialogflowTool
type DialogflowToolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowTool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowToolObservedState `json:"observedState,omitempty"`
}

// DialogflowToolObservedState is the state of the DialogflowTool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3beta1.Tool
type DialogflowToolObservedState struct {
	// Output only. The tool type.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.tool_type
	ToolType *string `json:"toolType,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowtool;gcpdialogflowtools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowTool is the Schema for the DialogflowTool API
// +k8s:openapi-gen=true
type DialogflowTool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowToolSpec   `json:"spec,omitempty"`
	Status DialogflowToolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowToolList contains a list of DialogflowTool
type DialogflowToolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowTool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowTool{}, &DialogflowToolList{})
}
