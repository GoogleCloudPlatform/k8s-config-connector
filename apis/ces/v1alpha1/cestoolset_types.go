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

// API sources for CESToolset, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1526e545e9d26f23b9c5d0f04af17297def8d045/google/cloud/ces/v1/toolset.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps
// +kcc:source:resource-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/reference/rest/v1/projects.locations.apps.toolsets

package v1alpha1

import (
	connectorsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connectors/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CESToolsetGVK = GroupVersion.WithKind("CESToolset")

// CESToolsetSpec defines the desired state of CESToolset
// +kcc:spec:proto=google.cloud.ces.v1.Toolset
// +kcc:required-from-proto
type CESToolsetSpec struct {
	// The CESApp this CESToolset is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *CESAppRef `json:"parentRef"`

	// The CESToolset name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. A toolset that contains a list of tools that are offered by the
	//  MCP server.
	// +kcc:proto:field=google.cloud.ces.v1.Toolset.mcp_toolset
	McpToolset *McpToolset `json:"mcpToolset,omitempty"`

	// Optional. A toolset that contains a list of tools that are defined by an
	//  OpenAPI schema.
	// +kcc:proto:field=google.cloud.ces.v1.Toolset.open_api_toolset
	OpenAPIToolset *OpenAPIToolset `json:"openAPIToolset,omitempty"`

	// Optional. A toolset that generates tools from an Integration Connectors
	//  Connection.
	// +kcc:proto:field=google.cloud.ces.v1.Toolset.connector_toolset
	ConnectorToolset *ConnectorToolset `json:"connectorToolset,omitempty"`

	// Optional. The display name of the toolset. Must be unique within the same
	//  app.
	// +kcc:proto:field=google.cloud.ces.v1.Toolset.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the toolset.
	// +kcc:proto:field=google.cloud.ces.v1.Toolset.description
	Description *string `json:"description,omitempty"`

	// Optional. The execution type of the tools in the toolset.
	// +kcc:proto:field=google.cloud.ces.v1.Toolset.execution_type
	ExecutionType *string `json:"executionType,omitempty"`

	// Optional. Configuration for tools behavior in fake mode.
	// +kcc:proto:field=google.cloud.ces.v1.Toolset.tool_fake_config
	ToolFakeConfig *ToolFakeConfig `json:"toolFakeConfig,omitempty"`
}

// +kcc:proto=google.cloud.ces.v1.ConnectorToolset
type ConnectorToolset struct {
	// Required. The full resource name of the referenced Integration Connectors
	//  Connection.
	//  Format:
	//  `projects/{project}/locations/{location}/connections/{connection}`
	// +kcc:proto:field=google.cloud.ces.v1.ConnectorToolset.connection
	// +required
	ConnectionRef *connectorsv1alpha1.ConnectorsConnectionRef `json:"connectionRef,omitempty"`

	// Optional. Configures how authentication is handled in Integration
	//  Connectors. By default, an admin authentication is passed in the
	//  Integration Connectors API requests. You can override it with a different
	//  end-user authentication config.
	//  **Note**: The Connection must have authentication override enabled in
	//  order to specify an EUC configuration here - otherwise, the Toolset
	//  creation will fail. See:
	//  https://cloud.google.com/application-integration/docs/configure-connectors-task#configure-authentication-override
	// +kcc:proto:field=google.cloud.ces.v1.ConnectorToolset.auth_config
	AuthConfig *EndUserAuthConfig `json:"authConfig,omitempty"`

	// Required. The list of connector actions/entity operations to generate tools
	//  for.
	// +kcc:proto:field=google.cloud.ces.v1.ConnectorToolset.connector_actions
	// +required
	ConnectorActions []Action `json:"connectorActions,omitempty"`
}

// CESToolsetStatus defines the config connector machine state of CESToolset
type CESToolsetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CESToolset resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CESToolsetObservedState `json:"observedState,omitempty"`
}

// CESToolsetObservedState is the state of the CESToolset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ces.v1.Toolset
type CESToolsetObservedState struct {
	// Output only. Timestamp when the toolset was created.
	// +kcc:proto:field=google.cloud.ces.v1.Toolset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the toolset was last updated.
	// +kcc:proto:field=google.cloud.ces.v1.Toolset.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// ETag used to ensure the object hasn't changed during a read-modify-write
	//  operation. If the etag is empty, the update will overwrite any concurrent
	//  changes.
	// +kcc:proto:field=google.cloud.ces.v1.Toolset.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcestoolset;gcpcestoolsets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CESToolset is the Schema for the CESToolset API
// +k8s:openapi-gen=true
type CESToolset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CESToolsetSpec   `json:"spec,omitempty"`
	Status CESToolsetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CESToolsetList contains a list of CESToolset
type CESToolsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CESToolset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CESToolset{}, &CESToolsetList{})
}
