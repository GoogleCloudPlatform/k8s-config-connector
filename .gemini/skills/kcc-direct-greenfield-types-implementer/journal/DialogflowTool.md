# DialogflowTool Journal

## Service and Proto Details
- Proto package: `google.cloud.dialogflow.cx.v3beta1`
- Target Resource: `DialogflowTool:Tool`
- Resource URI format: `projects/{project}/locations/{location}/agents/{agent}/tools/{tool}`

## Multi-service generation in `apis/dialogflow/generate.sh`
`apis/dialogflow` hosts multiple Dialogflow API versions:
- `google.cloud.dialogflow.v2` -> `types.generated.go`
- `google.cloud.dialogflow.cx.v3` -> `securitysettings_types.generated.go`
- `google.cloud.dialogflow.v2beta1` -> `siptrunk_types.generated.go`
- `google.cloud.dialogflow.cx.v3beta1` -> `tool_types.generated.go`

## Empty message `FallbackPrompt`
`Tool.DataStoreTool.FallbackPrompt` is defined as an empty message (`message FallbackPrompt {}`) in proto. In KRM, empty Go structs generate empty schema objects lacking `properties`, which fails Kubernetes OpenAPI validation (`TestCRDObjectTypes`). To resolve this, `Tool_DataStoreTool_FallbackPrompt` was defined with:
```go
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:validation:XPreserveUnknownFields
// +kcc:proto=google.cloud.dialogflow.cx.v3beta1.Tool.DataStoreTool.FallbackPrompt
type Tool_DataStoreTool_FallbackPrompt struct {
}
```

## References
1. `AgentRef`: References the parent `DialogflowCXAgent` (`projects/{project}/locations/{location}/agents/{agent}`). Implemented external-only `DialogflowCXAgentRef` and `DialogflowCXAgentIdentity`.
2. `ServiceRef` in `Tool_ServiceDirectoryConfig`: References `servicedirectoryv1beta1.ServiceDirectoryServiceRef`.
3. `DataStoreRef` in `DataStoreConnection`: References `discoveryenginev1alpha1.DiscoveryEngineDataStoreRef`.
4. `ConnectionRef` in `Tool_ConnectorTool`: References `connectorsv1alpha1.ConnectorsConnectionRef`.
5. `ExtensionRef` in `Tool_ExtensionTool`: References `vertexaiv1alpha1.VertexAIExtensionRef`. Implemented external-only `VertexAIExtensionRef` and `VertexAIExtensionIdentity`.

## Acronym Casing
The OpenAPI tool specification field requires `openAPISpec` in JSON tags (`json:"openAPISpec,omitempty"`) to adhere to KRM acronym capitalization rules.
