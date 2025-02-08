// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dialogflow

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DataStoreConnection_FromProto(mapCtx *direct.MapContext, in *pb.DataStoreConnection) *krm.DataStoreConnection {
	if in == nil {
		return nil
	}
	out := &krm.DataStoreConnection{}
	out.DataStoreType = direct.Enum_FromProto(mapCtx, in.GetDataStoreType())
	out.DataStore = direct.LazyPtr(in.GetDataStore())
	return out
}
func DataStoreConnection_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreConnection) *pb.DataStoreConnection {
	if in == nil {
		return nil
	}
	out := &pb.DataStoreConnection{}
	out.DataStoreType = direct.Enum_ToProto[pb.DataStoreType](mapCtx, in.DataStoreType)
	out.DataStore = direct.ValueOf(in.DataStore)
	return out
}
func DialogflowToolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Tool) *krm.DialogflowToolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowToolObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: OpenApiSpec
	// MISSING: DataStoreSpec
	// MISSING: ExtensionSpec
	// MISSING: FunctionSpec
	// MISSING: ToolType
	return out
}
func DialogflowToolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowToolObservedState) *pb.Tool {
	if in == nil {
		return nil
	}
	out := &pb.Tool{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: OpenApiSpec
	// MISSING: DataStoreSpec
	// MISSING: ExtensionSpec
	// MISSING: FunctionSpec
	// MISSING: ToolType
	return out
}
func DialogflowToolSpec_FromProto(mapCtx *direct.MapContext, in *pb.Tool) *krm.DialogflowToolSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowToolSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: OpenApiSpec
	// MISSING: DataStoreSpec
	// MISSING: ExtensionSpec
	// MISSING: FunctionSpec
	// MISSING: ToolType
	return out
}
func DialogflowToolSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowToolSpec) *pb.Tool {
	if in == nil {
		return nil
	}
	out := &pb.Tool{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: OpenApiSpec
	// MISSING: DataStoreSpec
	// MISSING: ExtensionSpec
	// MISSING: FunctionSpec
	// MISSING: ToolType
	return out
}
func Tool_FromProto(mapCtx *direct.MapContext, in *pb.Tool) *krm.Tool {
	if in == nil {
		return nil
	}
	out := &krm.Tool{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.OpenApiSpec = Tool_OpenApiTool_FromProto(mapCtx, in.GetOpenApiSpec())
	out.DataStoreSpec = Tool_DataStoreTool_FromProto(mapCtx, in.GetDataStoreSpec())
	out.ExtensionSpec = Tool_ExtensionTool_FromProto(mapCtx, in.GetExtensionSpec())
	out.FunctionSpec = Tool_FunctionTool_FromProto(mapCtx, in.GetFunctionSpec())
	// MISSING: ToolType
	return out
}
func Tool_ToProto(mapCtx *direct.MapContext, in *krm.Tool) *pb.Tool {
	if in == nil {
		return nil
	}
	out := &pb.Tool{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	if oneof := Tool_OpenApiTool_ToProto(mapCtx, in.OpenApiSpec); oneof != nil {
		out.Specification = &pb.Tool_OpenApiSpec{OpenApiSpec: oneof}
	}
	if oneof := Tool_DataStoreTool_ToProto(mapCtx, in.DataStoreSpec); oneof != nil {
		out.Specification = &pb.Tool_DataStoreSpec{DataStoreSpec: oneof}
	}
	if oneof := Tool_ExtensionTool_ToProto(mapCtx, in.ExtensionSpec); oneof != nil {
		out.Specification = &pb.Tool_ExtensionSpec{ExtensionSpec: oneof}
	}
	if oneof := Tool_FunctionTool_ToProto(mapCtx, in.FunctionSpec); oneof != nil {
		out.Specification = &pb.Tool_FunctionSpec{FunctionSpec: oneof}
	}
	// MISSING: ToolType
	return out
}
func ToolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Tool) *krm.ToolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ToolObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: OpenApiSpec
	// MISSING: DataStoreSpec
	// MISSING: ExtensionSpec
	// MISSING: FunctionSpec
	out.ToolType = direct.Enum_FromProto(mapCtx, in.GetToolType())
	return out
}
func ToolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ToolObservedState) *pb.Tool {
	if in == nil {
		return nil
	}
	out := &pb.Tool{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: OpenApiSpec
	// MISSING: DataStoreSpec
	// MISSING: ExtensionSpec
	// MISSING: FunctionSpec
	out.ToolType = direct.Enum_ToProto[pb.Tool_ToolType](mapCtx, in.ToolType)
	return out
}
func Tool_Authentication_FromProto(mapCtx *direct.MapContext, in *pb.Tool_Authentication) *krm.Tool_Authentication {
	if in == nil {
		return nil
	}
	out := &krm.Tool_Authentication{}
	out.ApiKeyConfig = Tool_Authentication_ApiKeyConfig_FromProto(mapCtx, in.GetApiKeyConfig())
	out.OauthConfig = Tool_Authentication_OAuthConfig_FromProto(mapCtx, in.GetOauthConfig())
	out.ServiceAgentAuthConfig = Tool_Authentication_ServiceAgentAuthConfig_FromProto(mapCtx, in.GetServiceAgentAuthConfig())
	out.BearerTokenConfig = Tool_Authentication_BearerTokenConfig_FromProto(mapCtx, in.GetBearerTokenConfig())
	return out
}
func Tool_Authentication_ToProto(mapCtx *direct.MapContext, in *krm.Tool_Authentication) *pb.Tool_Authentication {
	if in == nil {
		return nil
	}
	out := &pb.Tool_Authentication{}
	if oneof := Tool_Authentication_ApiKeyConfig_ToProto(mapCtx, in.ApiKeyConfig); oneof != nil {
		out.AuthConfig = &pb.Tool_Authentication_ApiKeyConfig_{ApiKeyConfig: oneof}
	}
	if oneof := Tool_Authentication_OAuthConfig_ToProto(mapCtx, in.OauthConfig); oneof != nil {
		out.AuthConfig = &pb.Tool_Authentication_OauthConfig{OauthConfig: oneof}
	}
	if oneof := Tool_Authentication_ServiceAgentAuthConfig_ToProto(mapCtx, in.ServiceAgentAuthConfig); oneof != nil {
		out.AuthConfig = &pb.Tool_Authentication_ServiceAgentAuthConfig_{ServiceAgentAuthConfig: oneof}
	}
	if oneof := Tool_Authentication_BearerTokenConfig_ToProto(mapCtx, in.BearerTokenConfig); oneof != nil {
		out.AuthConfig = &pb.Tool_Authentication_BearerTokenConfig_{BearerTokenConfig: oneof}
	}
	return out
}
func Tool_Authentication_ApiKeyConfig_FromProto(mapCtx *direct.MapContext, in *pb.Tool_Authentication_ApiKeyConfig) *krm.Tool_Authentication_ApiKeyConfig {
	if in == nil {
		return nil
	}
	out := &krm.Tool_Authentication_ApiKeyConfig{}
	out.KeyName = direct.LazyPtr(in.GetKeyName())
	out.ApiKey = direct.LazyPtr(in.GetApiKey())
	out.RequestLocation = direct.Enum_FromProto(mapCtx, in.GetRequestLocation())
	return out
}
func Tool_Authentication_ApiKeyConfig_ToProto(mapCtx *direct.MapContext, in *krm.Tool_Authentication_ApiKeyConfig) *pb.Tool_Authentication_ApiKeyConfig {
	if in == nil {
		return nil
	}
	out := &pb.Tool_Authentication_ApiKeyConfig{}
	out.KeyName = direct.ValueOf(in.KeyName)
	out.ApiKey = direct.ValueOf(in.ApiKey)
	out.RequestLocation = direct.Enum_ToProto[pb.Tool_Authentication_RequestLocation](mapCtx, in.RequestLocation)
	return out
}
func Tool_Authentication_BearerTokenConfig_FromProto(mapCtx *direct.MapContext, in *pb.Tool_Authentication_BearerTokenConfig) *krm.Tool_Authentication_BearerTokenConfig {
	if in == nil {
		return nil
	}
	out := &krm.Tool_Authentication_BearerTokenConfig{}
	out.Token = direct.LazyPtr(in.GetToken())
	return out
}
func Tool_Authentication_BearerTokenConfig_ToProto(mapCtx *direct.MapContext, in *krm.Tool_Authentication_BearerTokenConfig) *pb.Tool_Authentication_BearerTokenConfig {
	if in == nil {
		return nil
	}
	out := &pb.Tool_Authentication_BearerTokenConfig{}
	out.Token = direct.ValueOf(in.Token)
	return out
}
func Tool_Authentication_OAuthConfig_FromProto(mapCtx *direct.MapContext, in *pb.Tool_Authentication_OAuthConfig) *krm.Tool_Authentication_OAuthConfig {
	if in == nil {
		return nil
	}
	out := &krm.Tool_Authentication_OAuthConfig{}
	out.OauthGrantType = direct.Enum_FromProto(mapCtx, in.GetOauthGrantType())
	out.ClientID = direct.LazyPtr(in.GetClientId())
	out.ClientSecret = direct.LazyPtr(in.GetClientSecret())
	out.TokenEndpoint = direct.LazyPtr(in.GetTokenEndpoint())
	out.Scopes = in.Scopes
	return out
}
func Tool_Authentication_OAuthConfig_ToProto(mapCtx *direct.MapContext, in *krm.Tool_Authentication_OAuthConfig) *pb.Tool_Authentication_OAuthConfig {
	if in == nil {
		return nil
	}
	out := &pb.Tool_Authentication_OAuthConfig{}
	out.OauthGrantType = direct.Enum_ToProto[pb.Tool_Authentication_OAuthConfig_OauthGrantType](mapCtx, in.OauthGrantType)
	out.ClientId = direct.ValueOf(in.ClientID)
	out.ClientSecret = direct.ValueOf(in.ClientSecret)
	out.TokenEndpoint = direct.ValueOf(in.TokenEndpoint)
	out.Scopes = in.Scopes
	return out
}
func Tool_Authentication_ServiceAgentAuthConfig_FromProto(mapCtx *direct.MapContext, in *pb.Tool_Authentication_ServiceAgentAuthConfig) *krm.Tool_Authentication_ServiceAgentAuthConfig {
	if in == nil {
		return nil
	}
	out := &krm.Tool_Authentication_ServiceAgentAuthConfig{}
	out.ServiceAgentAuth = direct.Enum_FromProto(mapCtx, in.GetServiceAgentAuth())
	return out
}
func Tool_Authentication_ServiceAgentAuthConfig_ToProto(mapCtx *direct.MapContext, in *krm.Tool_Authentication_ServiceAgentAuthConfig) *pb.Tool_Authentication_ServiceAgentAuthConfig {
	if in == nil {
		return nil
	}
	out := &pb.Tool_Authentication_ServiceAgentAuthConfig{}
	out.ServiceAgentAuth = direct.Enum_ToProto[pb.Tool_Authentication_ServiceAgentAuthConfig_ServiceAgentAuth](mapCtx, in.ServiceAgentAuth)
	return out
}
func Tool_DataStoreTool_FromProto(mapCtx *direct.MapContext, in *pb.Tool_DataStoreTool) *krm.Tool_DataStoreTool {
	if in == nil {
		return nil
	}
	out := &krm.Tool_DataStoreTool{}
	out.DataStoreConnections = direct.Slice_FromProto(mapCtx, in.DataStoreConnections, DataStoreConnection_FromProto)
	out.FallbackPrompt = Tool_DataStoreTool_FallbackPrompt_FromProto(mapCtx, in.GetFallbackPrompt())
	return out
}
func Tool_DataStoreTool_ToProto(mapCtx *direct.MapContext, in *krm.Tool_DataStoreTool) *pb.Tool_DataStoreTool {
	if in == nil {
		return nil
	}
	out := &pb.Tool_DataStoreTool{}
	out.DataStoreConnections = direct.Slice_ToProto(mapCtx, in.DataStoreConnections, DataStoreConnection_ToProto)
	out.FallbackPrompt = Tool_DataStoreTool_FallbackPrompt_ToProto(mapCtx, in.FallbackPrompt)
	return out
}
func Tool_DataStoreTool_FallbackPrompt_FromProto(mapCtx *direct.MapContext, in *pb.Tool_DataStoreTool_FallbackPrompt) *krm.Tool_DataStoreTool_FallbackPrompt {
	if in == nil {
		return nil
	}
	out := &krm.Tool_DataStoreTool_FallbackPrompt{}
	return out
}
func Tool_DataStoreTool_FallbackPrompt_ToProto(mapCtx *direct.MapContext, in *krm.Tool_DataStoreTool_FallbackPrompt) *pb.Tool_DataStoreTool_FallbackPrompt {
	if in == nil {
		return nil
	}
	out := &pb.Tool_DataStoreTool_FallbackPrompt{}
	return out
}
func Tool_ExtensionTool_FromProto(mapCtx *direct.MapContext, in *pb.Tool_ExtensionTool) *krm.Tool_ExtensionTool {
	if in == nil {
		return nil
	}
	out := &krm.Tool_ExtensionTool{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Tool_ExtensionTool_ToProto(mapCtx *direct.MapContext, in *krm.Tool_ExtensionTool) *pb.Tool_ExtensionTool {
	if in == nil {
		return nil
	}
	out := &pb.Tool_ExtensionTool{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func Tool_FunctionTool_FromProto(mapCtx *direct.MapContext, in *pb.Tool_FunctionTool) *krm.Tool_FunctionTool {
	if in == nil {
		return nil
	}
	out := &krm.Tool_FunctionTool{}
	out.InputSchema = InputSchema_FromProto(mapCtx, in.GetInputSchema())
	out.OutputSchema = OutputSchema_FromProto(mapCtx, in.GetOutputSchema())
	return out
}
func Tool_FunctionTool_ToProto(mapCtx *direct.MapContext, in *krm.Tool_FunctionTool) *pb.Tool_FunctionTool {
	if in == nil {
		return nil
	}
	out := &pb.Tool_FunctionTool{}
	out.InputSchema = InputSchema_ToProto(mapCtx, in.InputSchema)
	out.OutputSchema = OutputSchema_ToProto(mapCtx, in.OutputSchema)
	return out
}
func Tool_OpenApiTool_FromProto(mapCtx *direct.MapContext, in *pb.Tool_OpenApiTool) *krm.Tool_OpenApiTool {
	if in == nil {
		return nil
	}
	out := &krm.Tool_OpenApiTool{}
	out.TextSchema = direct.LazyPtr(in.GetTextSchema())
	out.Authentication = Tool_Authentication_FromProto(mapCtx, in.GetAuthentication())
	out.TlsConfig = Tool_TLSConfig_FromProto(mapCtx, in.GetTlsConfig())
	out.ServiceDirectoryConfig = Tool_ServiceDirectoryConfig_FromProto(mapCtx, in.GetServiceDirectoryConfig())
	return out
}
func Tool_OpenApiTool_ToProto(mapCtx *direct.MapContext, in *krm.Tool_OpenApiTool) *pb.Tool_OpenApiTool {
	if in == nil {
		return nil
	}
	out := &pb.Tool_OpenApiTool{}
	if oneof := Tool_OpenApiTool_TextSchema_ToProto(mapCtx, in.TextSchema); oneof != nil {
		out.Schema = oneof
	}
	out.Authentication = Tool_Authentication_ToProto(mapCtx, in.Authentication)
	out.TlsConfig = Tool_TLSConfig_ToProto(mapCtx, in.TlsConfig)
	out.ServiceDirectoryConfig = Tool_ServiceDirectoryConfig_ToProto(mapCtx, in.ServiceDirectoryConfig)
	return out
}
func Tool_ServiceDirectoryConfig_FromProto(mapCtx *direct.MapContext, in *pb.Tool_ServiceDirectoryConfig) *krm.Tool_ServiceDirectoryConfig {
	if in == nil {
		return nil
	}
	out := &krm.Tool_ServiceDirectoryConfig{}
	out.Service = direct.LazyPtr(in.GetService())
	return out
}
func Tool_ServiceDirectoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.Tool_ServiceDirectoryConfig) *pb.Tool_ServiceDirectoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Tool_ServiceDirectoryConfig{}
	out.Service = direct.ValueOf(in.Service)
	return out
}
func Tool_TLSConfig_FromProto(mapCtx *direct.MapContext, in *pb.Tool_TLSConfig) *krm.Tool_TLSConfig {
	if in == nil {
		return nil
	}
	out := &krm.Tool_TLSConfig{}
	out.CaCerts = direct.Slice_FromProto(mapCtx, in.CaCerts, Tool_TLSConfig_CACert_FromProto)
	return out
}
func Tool_TLSConfig_ToProto(mapCtx *direct.MapContext, in *krm.Tool_TLSConfig) *pb.Tool_TLSConfig {
	if in == nil {
		return nil
	}
	out := &pb.Tool_TLSConfig{}
	out.CaCerts = direct.Slice_ToProto(mapCtx, in.CaCerts, Tool_TLSConfig_CACert_ToProto)
	return out
}
func Tool_TLSConfig_CACert_FromProto(mapCtx *direct.MapContext, in *pb.Tool_TLSConfig_CACert) *krm.Tool_TLSConfig_CACert {
	if in == nil {
		return nil
	}
	out := &krm.Tool_TLSConfig_CACert{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Cert = in.GetCert()
	return out
}
func Tool_TLSConfig_CACert_ToProto(mapCtx *direct.MapContext, in *krm.Tool_TLSConfig_CACert) *pb.Tool_TLSConfig_CACert {
	if in == nil {
		return nil
	}
	out := &pb.Tool_TLSConfig_CACert{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Cert = in.Cert
	return out
}
