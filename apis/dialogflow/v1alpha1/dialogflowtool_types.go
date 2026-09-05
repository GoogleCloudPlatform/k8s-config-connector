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
	aiplatformv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	discoveryenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	servicedirectoryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowToolGVK = GroupVersion.WithKind("DialogflowTool")

// DialogflowToolSpec defines the desired state of DialogflowTool
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3beta1.Tool
type DialogflowToolSpec struct {
	// The agent that this resource belongs to.
	// +required
	AgentRef *DialogflowAgentRef `json:"agentRef"`

	// Required. The human-readable name of the Tool, unique within an agent.
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. High level description of the Tool and its usage.
	// +required
	Description *string `json:"description,omitempty"`

	// OpenAPI specification of the Tool.
	// +optional
	OpenAPISpec *Tool_OpenAPITool `json:"openAPISpec,omitempty"`

	// Data store search tool specification.
	// +optional
	DataStoreSpec *Tool_DataStoreTool `json:"dataStoreSpec,omitempty"`

	// Vertex extension tool specification.
	// +optional
	ExtensionSpec *Tool_ExtensionTool `json:"extensionSpec,omitempty"`

	// Client side executed function specification.
	// +optional
	FunctionSpec *Tool_FunctionTool `json:"functionSpec,omitempty"`

	// Integration connectors tool specification.
	// +optional
	ConnectorSpec *Tool_ConnectorTool `json:"connectorSpec,omitempty"`

	// The DialogflowTool name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.OpenApiTool
type Tool_OpenAPITool struct {
	// Required. The OpenAPI schema specified as a text.
	TextSchema *string `json:"textSchema,omitempty"`

	// Optional. Authentication information required by the API.
	Authentication *Tool_Authentication `json:"authentication,omitempty"`

	// Optional. TLS configuration for the HTTPS verification.
	TLSConfig *Tool_TLSConfig `json:"tlsConfig,omitempty"`

	// Optional. Service Directory configuration.
	ServiceDirectoryConfig *Tool_ServiceDirectoryConfig `json:"serviceDirectoryConfig,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.ServiceDirectoryConfig
type Tool_ServiceDirectoryConfig struct {
	// The name of [Service Directory](https://cloud.google.com/service-directory) service.
	ServiceRef *servicedirectoryv1beta1.ServiceDirectoryServiceRef `json:"serviceRef,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool
type Tool_DataStoreTool struct {
	// Required. List of Web Queries database connection configs.
	DataStoreConnections []DataStoreConnection `json:"dataStoreConnections,omitempty"`

	// Optional. Fallback prompt config when no content is matched.
	FallbackPrompt *Tool_DataStoreTool_FallbackPrompt `json:"fallbackPrompt,omitempty"`
}

// +kubebuilder:validation:XPreserveUnknownFields
// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool.FallbackPrompt
type Tool_DataStoreTool_FallbackPrompt struct {
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection
type DataStoreConnection struct {
	// The type of the connected data store.
	DataStoreType *string `json:"dataStoreType,omitempty"`

	// The full name of the referenced data store.
	DataStoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`

	// The document processing mode for the data store connection. Should only be
	//  set for PUBLIC_WEB and UNSTRUCTURED data stores. If not set it is
	//  considered as DOCUMENTS, as this is the legacy mode.
	DocumentProcessingMode *string `json:"documentProcessingMode,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.FunctionTool
type Tool_FunctionTool struct {
	// Optional. The JSON schema is encapsulated in a
	//  [google.protobuf.Struct][google.protobuf.Struct] to describe the input of
	//  the function. This input is a JSON object that contains the function's
	//  parameters as properties of the object.
	InputSchema apiextensionsv1.JSON `json:"inputSchema,omitempty"`

	// Optional. The JSON schema is encapsulated in a
	//  [google.protobuf.Struct][google.protobuf.Struct] to describe the output
	//  of the function. This output is a JSON object that contains the
	//  function's parameters as properties of the object.
	OutputSchema apiextensionsv1.JSON `json:"outputSchema,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.ExtensionTool
type Tool_ExtensionTool struct {
	// Required. The referenced VertexAIExtension resource.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.ExtensionTool.name
	ExtensionRef *aiplatformv1alpha1.VertexAIExtensionRef `json:"extensionRef,omitempty"`
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
