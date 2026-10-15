// Copyright 2025 Google LLC
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

// API sources for CESTool, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/ces/v1beta/tool.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/docs/reference/rest/v1beta/projects.locations.apps.tools

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CESToolGVK = GroupVersion.WithKind("CESTool")

// CESToolSpec defines the desired state of CESTool
// +kcc:spec:proto=google.cloud.ces.v1beta.Tool
// +kcc:required-from-proto
type CESToolSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The CESTool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The client function.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.client_function
	ClientFunction *ClientFunction `json:"clientFunction,omitempty"`

	// Optional. The open API tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.open_api_tool
	OpenAPITool *OpenAPITool `json:"openAPITool,omitempty"`

	// Optional. The google search tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.google_search_tool
	GoogleSearchTool *GoogleSearchTool `json:"googleSearchTool,omitempty"`

	// Optional. The Integration Connector tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.connector_tool
	ConnectorTool *ConnectorTool `json:"connectorTool,omitempty"`

	// Optional. The data store tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.data_store_tool
	DataStoreTool *DataStoreTool `json:"dataStoreTool,omitempty"`

	// Optional. The python function tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.python_function
	PythonFunction *PythonFunction `json:"pythonFunction,omitempty"`

	// Optional. The MCP tool. An MCP tool cannot be created or updated directly
	//  and is managed by the MCP toolset.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.mcp_tool
	McpTool *McpTool `json:"mcpTool,omitempty"`

	// Optional. The file search tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.file_search_tool
	FileSearchTool *FileSearchTool `json:"fileSearchTool,omitempty"`

	// Optional. The system tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.system_tool
	SystemTool *SystemTool `json:"systemTool,omitempty"`

	// Optional. The agent tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.agent_tool
	AgentTool *AgentTool `json:"agentTool,omitempty"`

	// Optional. The widget tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.widget_tool
	WidgetTool *WidgetTool `json:"widgetTool,omitempty"`

	// Optional. The remote agent tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.remote_agent_tool
	RemoteAgentTool *RemoteAgentTool `json:"remoteAgentTool,omitempty"`

	// Optional. The execution type of the tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.execution_type
	ExecutionType *string `json:"executionType,omitempty"`

	// Optional. The timeout for the tool execution. If not set, the default
	//  timeout is 30 seconds for `SYNCHRONOUS` tools and 60 seconds for
	//  `ASYNCHRONOUS` tools.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Etag used to ensure the object hasn't changed during a read-modify-write
	//  operation. If the etag is empty, the update will overwrite any concurrent
	//  changes.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Configuration for tool behavior in fake mode.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.tool_fake_config
	ToolFakeConfig *ToolFakeConfig `json:"toolFakeConfig,omitempty"`
}

// CESToolStatus defines the config connector machine state of CESTool
type CESToolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CESTool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CESToolObservedState `json:"observedState,omitempty"`
}

// CESToolObservedState is the state of the CESTool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ces.v1beta.Tool
type CESToolObservedState struct {
	// Optional. The data store tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.data_store_tool
	DataStoreTool *DataStoreToolObservedState `json:"dataStoreTool,omitempty"`

	// Optional. The python function tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.python_function
	PythonFunction *PythonFunctionObservedState `json:"pythonFunction,omitempty"`

	// Optional. The MCP tool. An MCP tool cannot be created or updated directly
	//  and is managed by the MCP toolset.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.mcp_tool
	McpTool *McpToolObservedState `json:"mcpTool,omitempty"`

	// Optional. The system tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.system_tool
	SystemTool *SystemToolObservedState `json:"systemTool,omitempty"`

	// Optional. The widget tool.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.widget_tool
	WidgetTool *WidgetToolObservedState `json:"widgetTool,omitempty"`

	// Output only. The display name of the tool, derived based on the tool's
	//  type. For example, display name of a [ClientFunction][Tool.ClientFunction]
	//  is derived from its `name` property.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. Timestamp when the tool was created.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the tool was last updated.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. If the tool is generated by the LLM assistant, this field
	//  contains a descriptive summary of the generation.
	// +kcc:proto:field=google.cloud.ces.v1beta.Tool.generated_summary
	GeneratedSummary *string `json:"generatedSummary,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcestool;gcpcestools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CESTool is the Schema for the CESTool API
// +k8s:openapi-gen=true
type CESTool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CESToolSpec   `json:"spec,omitempty"`
	Status CESToolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CESToolList contains a list of CESTool
type CESToolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CESTool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CESTool{}, &CESToolList{})
}
