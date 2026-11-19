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

// API sources for CESTool, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1526e545e9d26f23b9c5d0f04af17297def8d045/google/cloud/ces/v1/tool.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps
// +kcc:source:resource-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/reference/rest/v1/projects.locations.apps.tools

package v1alpha1

import (
	aiplatformv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	connectorsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connectors/v1alpha1"
	discoveryenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	servicedirectoryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CESToolGVK = GroupVersion.WithKind("CESTool")

// CESToolSpec defines the desired state of CESTool
// +kcc:spec:proto=google.cloud.ces.v1.Tool
// +kcc:required-from-proto
type CESToolSpec struct {
	// The CESApp this CESTool is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *CESAppRef `json:"parentRef"`

	// The CESTool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. The client function.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.client_function
	ClientFunction *ClientFunction `json:"clientFunction,omitempty"`

	// Optional. The open API tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.open_api_tool
	OpenAPITool *OpenAPITool `json:"openAPITool,omitempty"`

	// Optional. The google search tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.google_search_tool
	GoogleSearchTool *GoogleSearchTool `json:"googleSearchTool,omitempty"`

	// Optional. The Integration Connector tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.connector_tool
	ConnectorTool *ConnectorTool `json:"connectorTool,omitempty"`

	// Optional. The data store tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.data_store_tool
	DataStoreTool *DataStoreTool `json:"dataStoreTool,omitempty"`

	// Optional. The python function tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.python_function
	PythonFunction *PythonFunction `json:"pythonFunction,omitempty"`

	// Optional. The MCP tool. An MCP tool cannot be created or updated directly
	//  and is managed by the MCP toolset.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.mcp_tool
	McpTool *McpTool `json:"mcpTool,omitempty"`

	// Optional. The file search tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.file_search_tool
	FileSearchTool *FileSearchTool `json:"fileSearchTool,omitempty"`

	// Optional. The system tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.system_tool
	SystemTool *SystemTool `json:"systemTool,omitempty"`

	// Optional. The agent tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.agent_tool
	AgentTool *AgentTool `json:"agentTool,omitempty"`

	// Optional. The widget tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.widget_tool
	WidgetTool *WidgetTool `json:"widgetTool,omitempty"`

	// Optional. The execution type of the tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.execution_type
	ExecutionType *string `json:"executionType,omitempty"`

	// Optional. Configuration for tool behavior in fake mode.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.tool_fake_config
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
// +kcc:observedstate:proto=google.cloud.ces.v1.Tool
type CESToolObservedState struct {
	// Optional. The data store tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.data_store_tool
	DataStoreTool *DataStoreToolObservedState `json:"dataStoreTool,omitempty"`

	// Optional. The python function tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.python_function
	PythonFunction *PythonFunctionObservedState `json:"pythonFunction,omitempty"`

	// Optional. The system tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.system_tool
	SystemTool *SystemToolObservedState `json:"systemTool,omitempty"`

	// Optional. The widget tool.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.widget_tool
	WidgetTool *WidgetToolObservedState `json:"widgetTool,omitempty"`

	// Output only. The display name of the tool, derived based on the tool's
	//  type. For example, display name of a [ClientFunction][Tool.ClientFunction]
	//  is derived from its `name` property.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. Timestamp when the tool was created.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the tool was last updated.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Etag used to ensure the object hasn't changed during a read-modify-write
	//  operation.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. If the tool is generated by the LLM assistant, this field
	//  contains a descriptive summary of the generation.
	// +kcc:proto:field=google.cloud.ces.v1.Tool.generated_summary
	GeneratedSummary *string `json:"generatedSummary,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.AgentTool
type AgentTool struct {
	// Required. The name of the agent tool.
	// +kcc:proto:field=google.cloud.ces.v1.AgentTool.name
	// +required
	Name *string `json:"name,omitempty"`

	// Optional. Description of the tool's purpose.
	// +kcc:proto:field=google.cloud.ces.v1.AgentTool.description
	Description *string `json:"description,omitempty"`

	// Optional. The resource name of the root agent that is the entry point of
	//  the tool. Format: `projects/{project}/locations/{location}/agents/{agent}`
	// +kcc:proto:field=google.cloud.ces.v1.AgentTool.root_agent
	RootAgentRef *CesAgentRef `json:"rootAgentRef,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.ApiKeyConfig
type APIKeyConfig struct {
	// Required. The parameter name or the header name of the API key.
	//  E.g., If the API request is "https://example.com/act?X-Api-Key=<API
	//  KEY>", "X-Api-Key" would be the parameter name.
	// +kcc:proto:field=google.cloud.ces.v1.ApiKeyConfig.key_name
	// +required
	KeyName *string `json:"keyName,omitempty"`

	// Required. The name of the SecretManager secret version resource storing the
	//  API key. Format: `projects/{project}/secrets/{secret}/versions/{version}`
	//
	//  Note: You should grant `roles/secretmanager.secretAccessor` role to the CES
	//  service agent
	//  `service-<PROJECT-NUMBER>@gcp-sa-ces.iam.gserviceaccount.com`.
	// +kcc:proto:field=google.cloud.ces.v1.ApiKeyConfig.api_key_secret_version
	// +required
	APIKeySecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"apiKeySecretVersionRef,omitempty"`

	// Required. Key location in the request.
	// +kcc:proto:field=google.cloud.ces.v1.ApiKeyConfig.request_location
	// +required
	RequestLocation *string `json:"requestLocation,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.ConnectorTool
type ConnectorTool struct {
	// Required. The full resource name of the referenced Integration Connectors
	//  Connection.
	//  Format:
	//  `projects/{project}/locations/{location}/connections/{connection}`
	// +kcc:proto:field=google.cloud.ces.v1.ConnectorTool.connection
	// +required
	ConnectionRef *connectorsv1alpha1.ConnectorsConnectionRef `json:"connectionRef,omitempty"`

	// Required. Action for the tool to use.
	// +kcc:proto:field=google.cloud.ces.v1.ConnectorTool.action
	// +required
	Action *Action `json:"action,omitempty"`

	// Optional. Configures how authentication is handled in Integration
	//  Connectors. By default, an admin authentication is passed in the
	//  Integration Connectors API requests. You can override it with a different
	//  end-user authentication config.
	//  **Note**: The Connection must have authentication override enabled in
	//  order to specify an EUC configuration here - otherwise, the ConnectorTool
	//  creation will fail. See
	//  https://cloud.google.com/application-integration/docs/configure-connectors-task#configure-authentication-override
	//  for details.
	// +kcc:proto:field=google.cloud.ces.v1.ConnectorTool.auth_config
	AuthConfig *EndUserAuthConfig `json:"authConfig,omitempty"`

	// Optional. The name of the tool that can be used by the Agent to decide
	//  whether to call this ConnectorTool.
	// +kcc:proto:field=google.cloud.ces.v1.ConnectorTool.name
	Name *string `json:"name,omitempty"`

	// Optional. The description of the tool that can be used by the Agent to
	//  decide whether to call this ConnectorTool.
	// +kcc:proto:field=google.cloud.ces.v1.ConnectorTool.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.DataStore
type DataStore struct {
	// Required. Full resource name of the DataStore.
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}`
	// +kcc:proto:field=google.cloud.ces.v1.DataStore.name
	// +required
	DataStoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"dataStoreRef,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.DataStoreTool.BoostSpecs
type DataStoreTool_BoostSpecs struct {
	// Required. The Data Store where the boosting configuration is applied.
	//  Full resource name of DataStore, such as
	//  projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}.
	// +kcc:proto:field=google.cloud.ces.v1.DataStoreTool.BoostSpecs.data_stores
	// +required
	DataStoreRefs []discoveryenginev1alpha1.DiscoveryEngineDataStoreRef `json:"dataStoreRefs,omitempty"`

	// Required. A list of boosting specifications.
	// +kcc:proto:field=google.cloud.ces.v1.DataStoreTool.BoostSpecs.spec
	// +required
	Spec []DataStoreTool_BoostSpec `json:"spec,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.DataStoreTool.EngineSource
type DataStoreTool_EngineSource struct {
	// Required. Full resource name of the Engine.
	//  Format:
	//  `projects/{project}/locations/{location}/collections/{collection}/engines/{engine}`
	// +kcc:proto:field=google.cloud.ces.v1.DataStoreTool.EngineSource.engine
	// +required
	EngineRef *discoveryenginev1alpha1.DiscoveryEngineEngineRef `json:"engineRef,omitempty"`

	// Optional. Use to target specific DataStores within the Engine.
	//  If empty, the search applies to all DataStores associated with the
	//  Engine.
	// +kcc:proto:field=google.cloud.ces.v1.DataStoreTool.EngineSource.data_store_sources
	DataStoreSources []DataStoreTool_DataStoreSource `json:"dataStoreSources,omitempty"`

	// Optional. A filter applied to the search across the Engine. Not relevant
	//  and not used if 'data_store_sources' is provided. See:
	//  https://cloud.google.com/generative-ai-app-builder/docs/filter-search-metadata
	// +kcc:proto:field=google.cloud.ces.v1.DataStoreTool.EngineSource.filter
	Filter *string `json:"filter,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.FileSearchTool
type FileSearchTool struct {
	// Optional. The type of the corpus. Default is FULLY_MANAGED.
	// +kcc:proto:field=google.cloud.ces.v1.FileSearchTool.corpus_type
	CorpusType *string `json:"corpusType,omitempty"`

	// Required. The tool name.
	// +kcc:proto:field=google.cloud.ces.v1.FileSearchTool.name
	// +required
	Name *string `json:"name,omitempty"`

	// Optional. The tool description.
	// +kcc:proto:field=google.cloud.ces.v1.FileSearchTool.description
	Description *string `json:"description,omitempty"`

	// Optional. The corpus where files are stored.
	//  Format:
	//  projects/{project}/locations/{location}/ragCorpora/{rag_corpus}
	// +kcc:proto:field=google.cloud.ces.v1.FileSearchTool.file_corpus
	FileCorpusRef *aiplatformv1alpha1.AIPlatformRAGCorpusRef `json:"fileCorpusRef,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.OAuthConfig
type OAuthConfig struct {
	// Required. OAuth grant types.
	// +kcc:proto:field=google.cloud.ces.v1.OAuthConfig.oauth_grant_type
	// +required
	OauthGrantType *string `json:"oauthGrantType,omitempty"`

	// Required. The client ID from the OAuth provider.
	// +kcc:proto:field=google.cloud.ces.v1.OAuthConfig.client_id
	// +required
	ClientID *string `json:"clientID,omitempty"`

	// Required. The name of the SecretManager secret version resource storing the
	//  client secret.
	//  Format: `projects/{project}/secrets/{secret}/versions/{version}`
	//
	//  Note: You should grant `roles/secretmanager.secretAccessor` role to the CES
	//  service agent
	//  `service-<PROJECT-NUMBER>@gcp-sa-ces.iam.gserviceaccount.com`.
	// +kcc:proto:field=google.cloud.ces.v1.OAuthConfig.client_secret_version
	// +required
	ClientSecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef `json:"clientSecretVersionRef,omitempty"`

	// Required. The token endpoint in the OAuth provider to exchange for an
	//  access token.
	// +kcc:proto:field=google.cloud.ces.v1.OAuthConfig.token_endpoint
	// +required
	TokenEndpoint *string `json:"tokenEndpoint,omitempty"`

	// Optional. The OAuth scopes to grant.
	// +kcc:proto:field=google.cloud.ces.v1.OAuthConfig.scopes
	Scopes []string `json:"scopes,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.ServiceAccountAuthConfig
type ServiceAccountAuthConfig struct {
	// Required. The email address of the service account used for authentication.
	//  CES uses this service account to exchange an access token and the access
	//  token is then sent in the `Authorization` header of the request.
	//
	//  The service account must have the
	//  `roles/iam.serviceAccountTokenCreator` role granted to the
	//  CES service agent
	//  `service-<PROJECT-NUMBER>@gcp-sa-ces.iam.gserviceaccount.com`.
	// +kcc:proto:field=google.cloud.ces.v1.ServiceAccountAuthConfig.service_account
	// +required
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// Optional. The OAuth scopes to grant. If not specified, the default scope
	//  `https://www.googleapis.com/auth/cloud-platform` is used.
	// +kcc:proto:field=google.cloud.ces.v1.ServiceAccountAuthConfig.scopes
	Scopes []string `json:"scopes,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.ServiceDirectoryConfig
type ServiceDirectoryConfig struct {
	// Required. The name of [Service
	//  Directory](https://cloud.google.com/service-directory) service.
	//  Format:
	//  `projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}`.
	//  Location of the service directory must be the same as the location of the
	//  app.
	// +kcc:proto:field=google.cloud.ces.v1.ServiceDirectoryConfig.service
	// +required
	ServiceRef *servicedirectoryv1beta1.ServiceDirectoryServiceRef `json:"serviceRef,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.WidgetTool.DataMapping
type WidgetTool_DataMapping struct {
	// Optional. The resource name of the tool that provides the data for the
	//  widget (e.g., a search tool or a custom function). Format:
	//  `projects/{project}/locations/{location}/agents/{agent}/tools/{tool}`
	// +kcc:proto:field=google.cloud.ces.v1.WidgetTool.DataMapping.source_tool_name
	SourceToolName *string `json:"sourceToolName,omitempty"`

	// Optional. A map of widget input parameter fields to the corresponding
	//  output fields of the source tool.
	// +kcc:proto:field=google.cloud.ces.v1.WidgetTool.DataMapping.field_mappings
	FieldMappings map[string]string `json:"fieldMappings,omitempty"`

	// Optional. Configuration for a Python function used to transform the
	//  source tool's output into the widget's input format.
	// +kcc:proto:field=google.cloud.ces.v1.WidgetTool.DataMapping.python_function
	PythonFunction *PythonFunction `json:"pythonFunction,omitempty"`

	// Optional. The mode of the data mapping.
	// +kcc:proto:field=google.cloud.ces.v1.WidgetTool.DataMapping.mode
	Mode *string `json:"mode,omitempty"`
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
