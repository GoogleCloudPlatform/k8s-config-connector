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

package apihub

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ApihubApiHubInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApiHubInstance) *krm.ApihubApiHubInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApihubApiHubInstanceObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	// MISSING: Labels
	// MISSING: Description
	return out
}
func ApihubApiHubInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApihubApiHubInstanceObservedState) *pb.ApiHubInstance {
	if in == nil {
		return nil
	}
	out := &pb.ApiHubInstance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	// MISSING: Labels
	// MISSING: Description
	return out
}
func ApihubApiHubInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.ApiHubInstance) *krm.ApihubApiHubInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApihubApiHubInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	// MISSING: Labels
	// MISSING: Description
	return out
}
func ApihubApiHubInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApihubApiHubInstanceSpec) *pb.ApiHubInstance {
	if in == nil {
		return nil
	}
	out := &pb.ApiHubInstance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: Config
	// MISSING: Labels
	// MISSING: Description
	return out
}
func ApihubApiObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.ApihubApiObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApihubApiObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: Owner
	// MISSING: Versions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: TargetUser
	// MISSING: Team
	// MISSING: BusinessUnit
	// MISSING: MaturityLevel
	// MISSING: Attributes
	// MISSING: ApiStyle
	// MISSING: SelectedVersion
	return out
}
func ApihubApiObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApihubApiObservedState) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: Owner
	// MISSING: Versions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: TargetUser
	// MISSING: Team
	// MISSING: BusinessUnit
	// MISSING: MaturityLevel
	// MISSING: Attributes
	// MISSING: ApiStyle
	// MISSING: SelectedVersion
	return out
}
func ApihubApiOperationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ApiOperation) *krm.ApihubApiOperationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApihubApiOperationObservedState{}
	// MISSING: Name
	// MISSING: Spec
	// MISSING: Details
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Attributes
	return out
}
func ApihubApiOperationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApihubApiOperationObservedState) *pb.ApiOperation {
	if in == nil {
		return nil
	}
	out := &pb.ApiOperation{}
	// MISSING: Name
	// MISSING: Spec
	// MISSING: Details
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Attributes
	return out
}
func ApihubApiOperationSpec_FromProto(mapCtx *direct.MapContext, in *pb.ApiOperation) *krm.ApihubApiOperationSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApihubApiOperationSpec{}
	// MISSING: Name
	// MISSING: Spec
	// MISSING: Details
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Attributes
	return out
}
func ApihubApiOperationSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApihubApiOperationSpec) *pb.ApiOperation {
	if in == nil {
		return nil
	}
	out := &pb.ApiOperation{}
	// MISSING: Name
	// MISSING: Spec
	// MISSING: Details
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Attributes
	return out
}
func ApihubApiSpec_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.ApihubApiSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApihubApiSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: Owner
	// MISSING: Versions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: TargetUser
	// MISSING: Team
	// MISSING: BusinessUnit
	// MISSING: MaturityLevel
	// MISSING: Attributes
	// MISSING: ApiStyle
	// MISSING: SelectedVersion
	return out
}
func ApihubApiSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApihubApiSpec) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: Owner
	// MISSING: Versions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: TargetUser
	// MISSING: Team
	// MISSING: BusinessUnit
	// MISSING: MaturityLevel
	// MISSING: Attributes
	// MISSING: ApiStyle
	// MISSING: SelectedVersion
	return out
}
func ApihubAttributeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Attribute) *krm.ApihubAttributeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApihubAttributeObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefinitionType
	// MISSING: Scope
	// MISSING: DataType
	// MISSING: AllowedValues
	// MISSING: Cardinality
	// MISSING: Mandatory
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ApihubAttributeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApihubAttributeObservedState) *pb.Attribute {
	if in == nil {
		return nil
	}
	out := &pb.Attribute{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefinitionType
	// MISSING: Scope
	// MISSING: DataType
	// MISSING: AllowedValues
	// MISSING: Cardinality
	// MISSING: Mandatory
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ApihubAttributeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Attribute) *krm.ApihubAttributeSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApihubAttributeSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefinitionType
	// MISSING: Scope
	// MISSING: DataType
	// MISSING: AllowedValues
	// MISSING: Cardinality
	// MISSING: Mandatory
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ApihubAttributeSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApihubAttributeSpec) *pb.Attribute {
	if in == nil {
		return nil
	}
	out := &pb.Attribute{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DefinitionType
	// MISSING: Scope
	// MISSING: DataType
	// MISSING: AllowedValues
	// MISSING: Cardinality
	// MISSING: Mandatory
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ApihubDefinitionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Definition) *krm.ApihubDefinitionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApihubDefinitionObservedState{}
	// MISSING: Schema
	// MISSING: Name
	// MISSING: Spec
	// MISSING: Type
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Attributes
	return out
}
func ApihubDefinitionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApihubDefinitionObservedState) *pb.Definition {
	if in == nil {
		return nil
	}
	out := &pb.Definition{}
	// MISSING: Schema
	// MISSING: Name
	// MISSING: Spec
	// MISSING: Type
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Attributes
	return out
}
func ApihubDefinitionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Definition) *krm.ApihubDefinitionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApihubDefinitionSpec{}
	// MISSING: Schema
	// MISSING: Name
	// MISSING: Spec
	// MISSING: Type
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Attributes
	return out
}
func ApihubDefinitionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApihubDefinitionSpec) *pb.Definition {
	if in == nil {
		return nil
	}
	out := &pb.Definition{}
	// MISSING: Schema
	// MISSING: Name
	// MISSING: Spec
	// MISSING: Type
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Attributes
	return out
}
func ApihubDependencyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Dependency) *krm.ApihubDependencyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApihubDependencyObservedState{}
	// MISSING: Name
	// MISSING: Consumer
	// MISSING: Supplier
	// MISSING: State
	// MISSING: Description
	// MISSING: DiscoveryMode
	// MISSING: ErrorDetail
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Attributes
	return out
}
func ApihubDependencyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApihubDependencyObservedState) *pb.Dependency {
	if in == nil {
		return nil
	}
	out := &pb.Dependency{}
	// MISSING: Name
	// MISSING: Consumer
	// MISSING: Supplier
	// MISSING: State
	// MISSING: Description
	// MISSING: DiscoveryMode
	// MISSING: ErrorDetail
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Attributes
	return out
}
func ApihubDependencySpec_FromProto(mapCtx *direct.MapContext, in *pb.Dependency) *krm.ApihubDependencySpec {
	if in == nil {
		return nil
	}
	out := &krm.ApihubDependencySpec{}
	// MISSING: Name
	// MISSING: Consumer
	// MISSING: Supplier
	// MISSING: State
	// MISSING: Description
	// MISSING: DiscoveryMode
	// MISSING: ErrorDetail
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Attributes
	return out
}
func ApihubDependencySpec_ToProto(mapCtx *direct.MapContext, in *krm.ApihubDependencySpec) *pb.Dependency {
	if in == nil {
		return nil
	}
	out := &pb.Dependency{}
	// MISSING: Name
	// MISSING: Consumer
	// MISSING: Supplier
	// MISSING: State
	// MISSING: Description
	// MISSING: DiscoveryMode
	// MISSING: ErrorDetail
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Attributes
	return out
}
func ApihubDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.ApihubDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApihubDeploymentObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: DeploymentType
	// MISSING: ResourceURI
	// MISSING: Endpoints
	// MISSING: ApiVersions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Slo
	// MISSING: Environment
	// MISSING: Attributes
	return out
}
func ApihubDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApihubDeploymentObservedState) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: DeploymentType
	// MISSING: ResourceURI
	// MISSING: Endpoints
	// MISSING: ApiVersions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Slo
	// MISSING: Environment
	// MISSING: Attributes
	return out
}
func ApihubDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.ApihubDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApihubDeploymentSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: DeploymentType
	// MISSING: ResourceURI
	// MISSING: Endpoints
	// MISSING: ApiVersions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Slo
	// MISSING: Environment
	// MISSING: Attributes
	return out
}
func ApihubDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApihubDeploymentSpec) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: DeploymentType
	// MISSING: ResourceURI
	// MISSING: Endpoints
	// MISSING: ApiVersions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Slo
	// MISSING: Environment
	// MISSING: Attributes
	return out
}
func ApihubExternalApiObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExternalApi) *krm.ApihubExternalApiObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApihubExternalApiObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Endpoints
	// MISSING: Paths
	// MISSING: Documentation
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ApihubExternalApiObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApihubExternalApiObservedState) *pb.ExternalApi {
	if in == nil {
		return nil
	}
	out := &pb.ExternalApi{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Endpoints
	// MISSING: Paths
	// MISSING: Documentation
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ApihubExternalApiSpec_FromProto(mapCtx *direct.MapContext, in *pb.ExternalApi) *krm.ApihubExternalApiSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApihubExternalApiSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Endpoints
	// MISSING: Paths
	// MISSING: Documentation
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ApihubExternalApiSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApihubExternalApiSpec) *pb.ExternalApi {
	if in == nil {
		return nil
	}
	out := &pb.ExternalApi{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Endpoints
	// MISSING: Paths
	// MISSING: Documentation
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ApihubSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Spec) *krm.ApihubSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApihubSpecObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SpecType
	// MISSING: Contents
	// MISSING: Details
	// MISSING: SourceURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LintResponse
	// MISSING: Attributes
	// MISSING: Documentation
	// MISSING: ParsingMode
	return out
}
func ApihubSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApihubSpecObservedState) *pb.Spec {
	if in == nil {
		return nil
	}
	out := &pb.Spec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SpecType
	// MISSING: Contents
	// MISSING: Details
	// MISSING: SourceURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LintResponse
	// MISSING: Attributes
	// MISSING: Documentation
	// MISSING: ParsingMode
	return out
}
func ApihubSpecSpec_FromProto(mapCtx *direct.MapContext, in *pb.Spec) *krm.ApihubSpecSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApihubSpecSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SpecType
	// MISSING: Contents
	// MISSING: Details
	// MISSING: SourceURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LintResponse
	// MISSING: Attributes
	// MISSING: Documentation
	// MISSING: ParsingMode
	return out
}
func ApihubSpecSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApihubSpecSpec) *pb.Spec {
	if in == nil {
		return nil
	}
	out := &pb.Spec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: SpecType
	// MISSING: Contents
	// MISSING: Details
	// MISSING: SourceURI
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LintResponse
	// MISSING: Attributes
	// MISSING: Documentation
	// MISSING: ParsingMode
	return out
}
func ApihubVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.ApihubVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApihubVersionObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: Specs
	// MISSING: ApiOperations
	// MISSING: Definitions
	// MISSING: Deployments
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Lifecycle
	// MISSING: Compliance
	// MISSING: Accreditation
	// MISSING: Attributes
	// MISSING: SelectedDeployment
	return out
}
func ApihubVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApihubVersionObservedState) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: Specs
	// MISSING: ApiOperations
	// MISSING: Definitions
	// MISSING: Deployments
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Lifecycle
	// MISSING: Compliance
	// MISSING: Accreditation
	// MISSING: Attributes
	// MISSING: SelectedDeployment
	return out
}
func ApihubVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.ApihubVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApihubVersionSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: Specs
	// MISSING: ApiOperations
	// MISSING: Definitions
	// MISSING: Deployments
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Lifecycle
	// MISSING: Compliance
	// MISSING: Accreditation
	// MISSING: Attributes
	// MISSING: SelectedDeployment
	return out
}
func ApihubVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApihubVersionSpec) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: Specs
	// MISSING: ApiOperations
	// MISSING: Definitions
	// MISSING: Deployments
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Lifecycle
	// MISSING: Compliance
	// MISSING: Accreditation
	// MISSING: Attributes
	// MISSING: SelectedDeployment
	return out
}
func AttributeValues_FromProto(mapCtx *direct.MapContext, in *pb.AttributeValues) *krm.AttributeValues {
	if in == nil {
		return nil
	}
	out := &krm.AttributeValues{}
	out.EnumValues = AttributeValues_EnumAttributeValues_FromProto(mapCtx, in.GetEnumValues())
	out.StringValues = AttributeValues_StringAttributeValues_FromProto(mapCtx, in.GetStringValues())
	out.JsonValues = AttributeValues_StringAttributeValues_FromProto(mapCtx, in.GetJsonValues())
	// MISSING: Attribute
	return out
}
func AttributeValues_ToProto(mapCtx *direct.MapContext, in *krm.AttributeValues) *pb.AttributeValues {
	if in == nil {
		return nil
	}
	out := &pb.AttributeValues{}
	if oneof := AttributeValues_EnumAttributeValues_ToProto(mapCtx, in.EnumValues); oneof != nil {
		out.Value = &pb.AttributeValues_EnumValues{EnumValues: oneof}
	}
	if oneof := AttributeValues_StringAttributeValues_ToProto(mapCtx, in.StringValues); oneof != nil {
		out.Value = &pb.AttributeValues_StringValues{StringValues: oneof}
	}
	if oneof := AttributeValues_StringAttributeValues_ToProto(mapCtx, in.JsonValues); oneof != nil {
		out.Value = &pb.AttributeValues_JsonValues{JsonValues: oneof}
	}
	// MISSING: Attribute
	return out
}
func AttributeValues_EnumAttributeValues_FromProto(mapCtx *direct.MapContext, in *pb.AttributeValues_EnumAttributeValues) *krm.AttributeValues_EnumAttributeValues {
	if in == nil {
		return nil
	}
	out := &krm.AttributeValues_EnumAttributeValues{}
	out.Values = direct.Slice_FromProto(mapCtx, in.Values, Attribute_AllowedValue_FromProto)
	return out
}
func AttributeValues_EnumAttributeValues_ToProto(mapCtx *direct.MapContext, in *krm.AttributeValues_EnumAttributeValues) *pb.AttributeValues_EnumAttributeValues {
	if in == nil {
		return nil
	}
	out := &pb.AttributeValues_EnumAttributeValues{}
	out.Values = direct.Slice_ToProto(mapCtx, in.Values, Attribute_AllowedValue_ToProto)
	return out
}
func AttributeValues_StringAttributeValues_FromProto(mapCtx *direct.MapContext, in *pb.AttributeValues_StringAttributeValues) *krm.AttributeValues_StringAttributeValues {
	if in == nil {
		return nil
	}
	out := &krm.AttributeValues_StringAttributeValues{}
	out.Values = in.Values
	return out
}
func AttributeValues_StringAttributeValues_ToProto(mapCtx *direct.MapContext, in *krm.AttributeValues_StringAttributeValues) *pb.AttributeValues_StringAttributeValues {
	if in == nil {
		return nil
	}
	out := &pb.AttributeValues_StringAttributeValues{}
	out.Values = in.Values
	return out
}
func Attribute_AllowedValue_FromProto(mapCtx *direct.MapContext, in *pb.Attribute_AllowedValue) *krm.Attribute_AllowedValue {
	if in == nil {
		return nil
	}
	out := &krm.Attribute_AllowedValue{}
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Immutable = direct.LazyPtr(in.GetImmutable())
	return out
}
func Attribute_AllowedValue_ToProto(mapCtx *direct.MapContext, in *krm.Attribute_AllowedValue) *pb.Attribute_AllowedValue {
	if in == nil {
		return nil
	}
	out := &pb.Attribute_AllowedValue{}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Immutable = direct.ValueOf(in.Immutable)
	return out
}
func Documentation_FromProto(mapCtx *direct.MapContext, in *pb.Documentation) *krm.Documentation {
	if in == nil {
		return nil
	}
	out := &krm.Documentation{}
	out.ExternalURI = direct.LazyPtr(in.GetExternalUri())
	return out
}
func Documentation_ToProto(mapCtx *direct.MapContext, in *krm.Documentation) *pb.Documentation {
	if in == nil {
		return nil
	}
	out := &pb.Documentation{}
	out.ExternalUri = direct.ValueOf(in.ExternalURI)
	return out
}
func ExternalApi_FromProto(mapCtx *direct.MapContext, in *pb.ExternalApi) *krm.ExternalApi {
	if in == nil {
		return nil
	}
	out := &krm.ExternalApi{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Endpoints = in.Endpoints
	out.Paths = in.Paths
	out.Documentation = Documentation_FromProto(mapCtx, in.GetDocumentation())
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ExternalApi_ToProto(mapCtx *direct.MapContext, in *krm.ExternalApi) *pb.ExternalApi {
	if in == nil {
		return nil
	}
	out := &pb.ExternalApi{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Endpoints = in.Endpoints
	out.Paths = in.Paths
	out.Documentation = Documentation_ToProto(mapCtx, in.Documentation)
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ExternalApiObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExternalApi) *krm.ExternalApiObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ExternalApiObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Endpoints
	// MISSING: Paths
	// MISSING: Documentation
	// MISSING: Attributes
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func ExternalApiObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ExternalApiObservedState) *pb.ExternalApi {
	if in == nil {
		return nil
	}
	out := &pb.ExternalApi{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Endpoints
	// MISSING: Paths
	// MISSING: Documentation
	// MISSING: Attributes
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
