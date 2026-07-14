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
	discoveryenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = apiextensionsv1.JSON{}

var DialogflowToolGVK = GroupVersion.WithKind("DialogflowTool")

// DialogflowToolSpec defines the desired state of DialogflowTool
// +kcc:spec:proto=google.cloud.dialogflow.cx.v3beta1.Tool
type DialogflowToolSpec struct {
	// The DialogflowCXAgent that this resource belongs to.
	// +required
	AgentRef *DialogflowCXAgentRef `json:"agentRef"`

	// The DialogflowTool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The human-readable name of the Tool, unique within an agent.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.display_name
	DisplayName *string `json:"displayName"`

	// Required. High level description of the Tool and its usage.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.description
	Description *string `json:"description"`

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

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool
type Tool_DataStoreTool struct {
	// Required. List of data stores to search.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool.data_store_connections
	DataStoreConnections []DataStoreConnection `json:"dataStoreConnections,omitempty"`

	// Required. Fallback prompt configurations to use.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool.fallback_prompt
	FallbackPrompt *Tool_DataStoreTool_FallbackPrompt `json:"fallbackPrompt,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection
type DataStoreConnection struct {
	// The type of the connected data store.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection.data_store_type
	DataStoreType *string `json:"dataStoreType,omitempty"`

	// The referenced DiscoveryEngineDataStore resource.
	DataStoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`

	// The document processing mode for the data store connection. Should only be
	//  set for PUBLIC_WEB and UNSTRUCTURED data stores. If not set it is
	//  considered as DOCUMENTS, as this is the legacy mode.
	// +kcc:proto:field=google.cloud.dialogflow.cx.v3beta1.DataStoreConnection.document_processing_mode
	DocumentProcessingMode *string `json:"documentProcessingMode,omitempty"`
}

// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.ExtensionTool
type Tool_ExtensionTool struct {
	// The referenced VertexAIExtension resource.
	ExtensionRef *VertexAIExtensionRef `json:"extensionRef,omitempty"`
}
