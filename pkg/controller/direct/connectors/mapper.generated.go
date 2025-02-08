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

package connectors

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/connectors/apiv1/connectorspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connectors/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AuthConfigTemplate_FromProto(mapCtx *direct.MapContext, in *pb.AuthConfigTemplate) *krm.AuthConfigTemplate {
	if in == nil {
		return nil
	}
	out := &krm.AuthConfigTemplate{}
	out.AuthType = direct.Enum_FromProto(mapCtx, in.GetAuthType())
	out.ConfigVariableTemplates = direct.Slice_FromProto(mapCtx, in.ConfigVariableTemplates, ConfigVariableTemplate_FromProto)
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func AuthConfigTemplate_ToProto(mapCtx *direct.MapContext, in *krm.AuthConfigTemplate) *pb.AuthConfigTemplate {
	if in == nil {
		return nil
	}
	out := &pb.AuthConfigTemplate{}
	out.AuthType = direct.Enum_ToProto[pb.AuthType](mapCtx, in.AuthType)
	out.ConfigVariableTemplates = direct.Slice_ToProto(mapCtx, in.ConfigVariableTemplates, ConfigVariableTemplate_ToProto)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func AuthorizationCodeLink_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizationCodeLink) *krm.AuthorizationCodeLink {
	if in == nil {
		return nil
	}
	out := &krm.AuthorizationCodeLink{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.Scopes = in.Scopes
	out.ClientID = direct.LazyPtr(in.GetClientId())
	out.EnablePkce = direct.LazyPtr(in.GetEnablePkce())
	return out
}
func AuthorizationCodeLink_ToProto(mapCtx *direct.MapContext, in *krm.AuthorizationCodeLink) *pb.AuthorizationCodeLink {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizationCodeLink{}
	out.Uri = direct.ValueOf(in.URI)
	out.Scopes = in.Scopes
	out.ClientId = direct.ValueOf(in.ClientID)
	out.EnablePkce = direct.ValueOf(in.EnablePkce)
	return out
}
func ConfigVariableTemplate_FromProto(mapCtx *direct.MapContext, in *pb.ConfigVariableTemplate) *krm.ConfigVariableTemplate {
	if in == nil {
		return nil
	}
	out := &krm.ConfigVariableTemplate{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.ValueType = direct.Enum_FromProto(mapCtx, in.GetValueType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ValidationRegex = direct.LazyPtr(in.GetValidationRegex())
	out.Required = direct.LazyPtr(in.GetRequired())
	out.RoleGrant = RoleGrant_FromProto(mapCtx, in.GetRoleGrant())
	out.EnumOptions = direct.Slice_FromProto(mapCtx, in.EnumOptions, EnumOption_FromProto)
	out.AuthorizationCodeLink = AuthorizationCodeLink_FromProto(mapCtx, in.GetAuthorizationCodeLink())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.IsAdvanced = direct.LazyPtr(in.GetIsAdvanced())
	return out
}
func ConfigVariableTemplate_ToProto(mapCtx *direct.MapContext, in *krm.ConfigVariableTemplate) *pb.ConfigVariableTemplate {
	if in == nil {
		return nil
	}
	out := &pb.ConfigVariableTemplate{}
	out.Key = direct.ValueOf(in.Key)
	out.ValueType = direct.Enum_ToProto[pb.ConfigVariableTemplate_ValueType](mapCtx, in.ValueType)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.ValidationRegex = direct.ValueOf(in.ValidationRegex)
	out.Required = direct.ValueOf(in.Required)
	out.RoleGrant = RoleGrant_ToProto(mapCtx, in.RoleGrant)
	out.EnumOptions = direct.Slice_ToProto(mapCtx, in.EnumOptions, EnumOption_ToProto)
	out.AuthorizationCodeLink = AuthorizationCodeLink_ToProto(mapCtx, in.AuthorizationCodeLink)
	out.State = direct.Enum_ToProto[pb.ConfigVariableTemplate_State](mapCtx, in.State)
	out.IsAdvanced = direct.ValueOf(in.IsAdvanced)
	return out
}
func ConnectorVersion_FromProto(mapCtx *direct.MapContext, in *pb.ConnectorVersion) *krm.ConnectorVersion {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: LaunchStage
	// MISSING: ReleaseVersion
	// MISSING: AuthConfigTemplates
	// MISSING: ConfigVariableTemplates
	// MISSING: SupportedRuntimeFeatures
	// MISSING: DisplayName
	// MISSING: EgressControlConfig
	// MISSING: RoleGrants
	// MISSING: RoleGrant
	// MISSING: SslConfigTemplate
	return out
}
func ConnectorVersion_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorVersion) *pb.ConnectorVersion {
	if in == nil {
		return nil
	}
	out := &pb.ConnectorVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: LaunchStage
	// MISSING: ReleaseVersion
	// MISSING: AuthConfigTemplates
	// MISSING: ConfigVariableTemplates
	// MISSING: SupportedRuntimeFeatures
	// MISSING: DisplayName
	// MISSING: EgressControlConfig
	// MISSING: RoleGrants
	// MISSING: RoleGrant
	// MISSING: SslConfigTemplate
	return out
}
func ConnectorVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectorVersion) *krm.ConnectorVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorVersionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Labels = in.Labels
	out.LaunchStage = direct.Enum_FromProto(mapCtx, in.GetLaunchStage())
	out.ReleaseVersion = direct.LazyPtr(in.GetReleaseVersion())
	out.AuthConfigTemplates = direct.Slice_FromProto(mapCtx, in.AuthConfigTemplates, AuthConfigTemplate_FromProto)
	out.ConfigVariableTemplates = direct.Slice_FromProto(mapCtx, in.ConfigVariableTemplates, ConfigVariableTemplate_FromProto)
	out.SupportedRuntimeFeatures = SupportedRuntimeFeatures_FromProto(mapCtx, in.GetSupportedRuntimeFeatures())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.EgressControlConfig = EgressControlConfig_FromProto(mapCtx, in.GetEgressControlConfig())
	out.RoleGrants = direct.Slice_FromProto(mapCtx, in.RoleGrants, RoleGrant_FromProto)
	out.RoleGrant = RoleGrant_FromProto(mapCtx, in.GetRoleGrant())
	out.SslConfigTemplate = SslConfigTemplate_FromProto(mapCtx, in.GetSslConfigTemplate())
	return out
}
func ConnectorVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorVersionObservedState) *pb.ConnectorVersion {
	if in == nil {
		return nil
	}
	out := &pb.ConnectorVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Labels = in.Labels
	out.LaunchStage = direct.Enum_ToProto[pb.LaunchStage](mapCtx, in.LaunchStage)
	out.ReleaseVersion = direct.ValueOf(in.ReleaseVersion)
	out.AuthConfigTemplates = direct.Slice_ToProto(mapCtx, in.AuthConfigTemplates, AuthConfigTemplate_ToProto)
	out.ConfigVariableTemplates = direct.Slice_ToProto(mapCtx, in.ConfigVariableTemplates, ConfigVariableTemplate_ToProto)
	out.SupportedRuntimeFeatures = SupportedRuntimeFeatures_ToProto(mapCtx, in.SupportedRuntimeFeatures)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.EgressControlConfig = EgressControlConfig_ToProto(mapCtx, in.EgressControlConfig)
	out.RoleGrants = direct.Slice_ToProto(mapCtx, in.RoleGrants, RoleGrant_ToProto)
	out.RoleGrant = RoleGrant_ToProto(mapCtx, in.RoleGrant)
	out.SslConfigTemplate = SslConfigTemplate_ToProto(mapCtx, in.SslConfigTemplate)
	return out
}
func ConnectorsConnectorVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectorVersion) *krm.ConnectorsConnectorVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsConnectorVersionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: LaunchStage
	// MISSING: ReleaseVersion
	// MISSING: AuthConfigTemplates
	// MISSING: ConfigVariableTemplates
	// MISSING: SupportedRuntimeFeatures
	// MISSING: DisplayName
	// MISSING: EgressControlConfig
	// MISSING: RoleGrants
	// MISSING: RoleGrant
	// MISSING: SslConfigTemplate
	return out
}
func ConnectorsConnectorVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorsConnectorVersionObservedState) *pb.ConnectorVersion {
	if in == nil {
		return nil
	}
	out := &pb.ConnectorVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: LaunchStage
	// MISSING: ReleaseVersion
	// MISSING: AuthConfigTemplates
	// MISSING: ConfigVariableTemplates
	// MISSING: SupportedRuntimeFeatures
	// MISSING: DisplayName
	// MISSING: EgressControlConfig
	// MISSING: RoleGrants
	// MISSING: RoleGrant
	// MISSING: SslConfigTemplate
	return out
}
func ConnectorsConnectorVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConnectorVersion) *krm.ConnectorsConnectorVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConnectorsConnectorVersionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: LaunchStage
	// MISSING: ReleaseVersion
	// MISSING: AuthConfigTemplates
	// MISSING: ConfigVariableTemplates
	// MISSING: SupportedRuntimeFeatures
	// MISSING: DisplayName
	// MISSING: EgressControlConfig
	// MISSING: RoleGrants
	// MISSING: RoleGrant
	// MISSING: SslConfigTemplate
	return out
}
func ConnectorsConnectorVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConnectorsConnectorVersionSpec) *pb.ConnectorVersion {
	if in == nil {
		return nil
	}
	out := &pb.ConnectorVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: LaunchStage
	// MISSING: ReleaseVersion
	// MISSING: AuthConfigTemplates
	// MISSING: ConfigVariableTemplates
	// MISSING: SupportedRuntimeFeatures
	// MISSING: DisplayName
	// MISSING: EgressControlConfig
	// MISSING: RoleGrants
	// MISSING: RoleGrant
	// MISSING: SslConfigTemplate
	return out
}
func EgressControlConfig_FromProto(mapCtx *direct.MapContext, in *pb.EgressControlConfig) *krm.EgressControlConfig {
	if in == nil {
		return nil
	}
	out := &krm.EgressControlConfig{}
	out.Backends = direct.LazyPtr(in.GetBackends())
	out.ExtractionRules = ExtractionRules_FromProto(mapCtx, in.GetExtractionRules())
	return out
}
func EgressControlConfig_ToProto(mapCtx *direct.MapContext, in *krm.EgressControlConfig) *pb.EgressControlConfig {
	if in == nil {
		return nil
	}
	out := &pb.EgressControlConfig{}
	if oneof := EgressControlConfig_Backends_ToProto(mapCtx, in.Backends); oneof != nil {
		out.OneofBackends = oneof
	}
	if oneof := ExtractionRules_ToProto(mapCtx, in.ExtractionRules); oneof != nil {
		out.OneofBackends = &pb.EgressControlConfig_ExtractionRules{ExtractionRules: oneof}
	}
	return out
}
func EnumOption_FromProto(mapCtx *direct.MapContext, in *pb.EnumOption) *krm.EnumOption {
	if in == nil {
		return nil
	}
	out := &krm.EnumOption{}
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func EnumOption_ToProto(mapCtx *direct.MapContext, in *krm.EnumOption) *pb.EnumOption {
	if in == nil {
		return nil
	}
	out := &pb.EnumOption{}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func ExtractionRule_FromProto(mapCtx *direct.MapContext, in *pb.ExtractionRule) *krm.ExtractionRule {
	if in == nil {
		return nil
	}
	out := &krm.ExtractionRule{}
	out.Source = ExtractionRule_Source_FromProto(mapCtx, in.GetSource())
	out.ExtractionRegex = direct.LazyPtr(in.GetExtractionRegex())
	return out
}
func ExtractionRule_ToProto(mapCtx *direct.MapContext, in *krm.ExtractionRule) *pb.ExtractionRule {
	if in == nil {
		return nil
	}
	out := &pb.ExtractionRule{}
	out.Source = ExtractionRule_Source_ToProto(mapCtx, in.Source)
	out.ExtractionRegex = direct.ValueOf(in.ExtractionRegex)
	return out
}
func ExtractionRule_Source_FromProto(mapCtx *direct.MapContext, in *pb.ExtractionRule_Source) *krm.ExtractionRule_Source {
	if in == nil {
		return nil
	}
	out := &krm.ExtractionRule_Source{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	out.FieldID = direct.LazyPtr(in.GetFieldId())
	return out
}
func ExtractionRule_Source_ToProto(mapCtx *direct.MapContext, in *krm.ExtractionRule_Source) *pb.ExtractionRule_Source {
	if in == nil {
		return nil
	}
	out := &pb.ExtractionRule_Source{}
	out.SourceType = direct.Enum_ToProto[pb.ExtractionRule_SourceType](mapCtx, in.SourceType)
	out.FieldId = direct.ValueOf(in.FieldID)
	return out
}
func ExtractionRules_FromProto(mapCtx *direct.MapContext, in *pb.ExtractionRules) *krm.ExtractionRules {
	if in == nil {
		return nil
	}
	out := &krm.ExtractionRules{}
	out.ExtractionRule = direct.Slice_FromProto(mapCtx, in.ExtractionRule, ExtractionRule_FromProto)
	return out
}
func ExtractionRules_ToProto(mapCtx *direct.MapContext, in *krm.ExtractionRules) *pb.ExtractionRules {
	if in == nil {
		return nil
	}
	out := &pb.ExtractionRules{}
	out.ExtractionRule = direct.Slice_ToProto(mapCtx, in.ExtractionRule, ExtractionRule_ToProto)
	return out
}
func RoleGrant_FromProto(mapCtx *direct.MapContext, in *pb.RoleGrant) *krm.RoleGrant {
	if in == nil {
		return nil
	}
	out := &krm.RoleGrant{}
	out.Principal = direct.Enum_FromProto(mapCtx, in.GetPrincipal())
	out.Roles = in.Roles
	out.Resource = RoleGrant_Resource_FromProto(mapCtx, in.GetResource())
	out.HelperTextTemplate = direct.LazyPtr(in.GetHelperTextTemplate())
	return out
}
func RoleGrant_ToProto(mapCtx *direct.MapContext, in *krm.RoleGrant) *pb.RoleGrant {
	if in == nil {
		return nil
	}
	out := &pb.RoleGrant{}
	out.Principal = direct.Enum_ToProto[pb.RoleGrant_Principal](mapCtx, in.Principal)
	out.Roles = in.Roles
	out.Resource = RoleGrant_Resource_ToProto(mapCtx, in.Resource)
	out.HelperTextTemplate = direct.ValueOf(in.HelperTextTemplate)
	return out
}
func RoleGrant_Resource_FromProto(mapCtx *direct.MapContext, in *pb.RoleGrant_Resource) *krm.RoleGrant_Resource {
	if in == nil {
		return nil
	}
	out := &krm.RoleGrant_Resource{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.PathTemplate = direct.LazyPtr(in.GetPathTemplate())
	return out
}
func RoleGrant_Resource_ToProto(mapCtx *direct.MapContext, in *krm.RoleGrant_Resource) *pb.RoleGrant_Resource {
	if in == nil {
		return nil
	}
	out := &pb.RoleGrant_Resource{}
	out.Type = direct.Enum_ToProto[pb.RoleGrant_Resource_Type](mapCtx, in.Type)
	out.PathTemplate = direct.ValueOf(in.PathTemplate)
	return out
}
func SslConfigTemplate_FromProto(mapCtx *direct.MapContext, in *pb.SslConfigTemplate) *krm.SslConfigTemplate {
	if in == nil {
		return nil
	}
	out := &krm.SslConfigTemplate{}
	out.SslType = direct.Enum_FromProto(mapCtx, in.GetSslType())
	out.IsTlsMandatory = direct.LazyPtr(in.GetIsTlsMandatory())
	out.ServerCertType = direct.EnumSlice_FromProto(mapCtx, in.ServerCertType)
	out.ClientCertType = direct.EnumSlice_FromProto(mapCtx, in.ClientCertType)
	out.AdditionalVariables = direct.Slice_FromProto(mapCtx, in.AdditionalVariables, ConfigVariableTemplate_FromProto)
	return out
}
func SslConfigTemplate_ToProto(mapCtx *direct.MapContext, in *krm.SslConfigTemplate) *pb.SslConfigTemplate {
	if in == nil {
		return nil
	}
	out := &pb.SslConfigTemplate{}
	out.SslType = direct.Enum_ToProto[pb.SslType](mapCtx, in.SslType)
	out.IsTlsMandatory = direct.ValueOf(in.IsTlsMandatory)
	out.ServerCertType = direct.EnumSlice_ToProto[pb.CertType](mapCtx, in.ServerCertType)
	out.ClientCertType = direct.EnumSlice_ToProto[pb.CertType](mapCtx, in.ClientCertType)
	out.AdditionalVariables = direct.Slice_ToProto(mapCtx, in.AdditionalVariables, ConfigVariableTemplate_ToProto)
	return out
}
func SupportedRuntimeFeatures_FromProto(mapCtx *direct.MapContext, in *pb.SupportedRuntimeFeatures) *krm.SupportedRuntimeFeatures {
	if in == nil {
		return nil
	}
	out := &krm.SupportedRuntimeFeatures{}
	out.EntityApis = direct.LazyPtr(in.GetEntityApis())
	out.ActionApis = direct.LazyPtr(in.GetActionApis())
	out.SqlQuery = direct.LazyPtr(in.GetSqlQuery())
	return out
}
func SupportedRuntimeFeatures_ToProto(mapCtx *direct.MapContext, in *krm.SupportedRuntimeFeatures) *pb.SupportedRuntimeFeatures {
	if in == nil {
		return nil
	}
	out := &pb.SupportedRuntimeFeatures{}
	out.EntityApis = direct.ValueOf(in.EntityApis)
	out.ActionApis = direct.ValueOf(in.ActionApis)
	out.SqlQuery = direct.ValueOf(in.SqlQuery)
	return out
}
