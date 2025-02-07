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
func AttributeValuesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AttributeValues) *krm.AttributeValuesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AttributeValuesObservedState{}
	// MISSING: EnumValues
	// MISSING: StringValues
	// MISSING: JsonValues
	out.Attribute = direct.LazyPtr(in.GetAttribute())
	return out
}
func AttributeValuesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AttributeValuesObservedState) *pb.AttributeValues {
	if in == nil {
		return nil
	}
	out := &pb.AttributeValues{}
	// MISSING: EnumValues
	// MISSING: StringValues
	// MISSING: JsonValues
	out.Attribute = direct.ValueOf(in.Attribute)
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
func Issue_FromProto(mapCtx *direct.MapContext, in *pb.Issue) *krm.Issue {
	if in == nil {
		return nil
	}
	out := &krm.Issue{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Path = in.Path
	out.Message = direct.LazyPtr(in.GetMessage())
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	out.Range = Range_FromProto(mapCtx, in.GetRange())
	return out
}
func Issue_ToProto(mapCtx *direct.MapContext, in *krm.Issue) *pb.Issue {
	if in == nil {
		return nil
	}
	out := &pb.Issue{}
	out.Code = direct.ValueOf(in.Code)
	out.Path = in.Path
	out.Message = direct.ValueOf(in.Message)
	out.Severity = direct.Enum_ToProto[pb.Severity](mapCtx, in.Severity)
	out.Range = Range_ToProto(mapCtx, in.Range)
	return out
}
func OpenApiSpecDetails_FromProto(mapCtx *direct.MapContext, in *pb.OpenApiSpecDetails) *krm.OpenApiSpecDetails {
	if in == nil {
		return nil
	}
	out := &krm.OpenApiSpecDetails{}
	// MISSING: Format
	// MISSING: Version
	// MISSING: Owner
	return out
}
func OpenApiSpecDetails_ToProto(mapCtx *direct.MapContext, in *krm.OpenApiSpecDetails) *pb.OpenApiSpecDetails {
	if in == nil {
		return nil
	}
	out := &pb.OpenApiSpecDetails{}
	// MISSING: Format
	// MISSING: Version
	// MISSING: Owner
	return out
}
func OpenApiSpecDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OpenApiSpecDetails) *krm.OpenApiSpecDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OpenApiSpecDetailsObservedState{}
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Owner = Owner_FromProto(mapCtx, in.GetOwner())
	return out
}
func OpenApiSpecDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OpenApiSpecDetailsObservedState) *pb.OpenApiSpecDetails {
	if in == nil {
		return nil
	}
	out := &pb.OpenApiSpecDetails{}
	out.Format = direct.Enum_ToProto[pb.OpenApiSpecDetails_Format](mapCtx, in.Format)
	out.Version = direct.ValueOf(in.Version)
	out.Owner = Owner_ToProto(mapCtx, in.Owner)
	return out
}
func Owner_FromProto(mapCtx *direct.MapContext, in *pb.Owner) *krm.Owner {
	if in == nil {
		return nil
	}
	out := &krm.Owner{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Email = direct.LazyPtr(in.GetEmail())
	return out
}
func Owner_ToProto(mapCtx *direct.MapContext, in *krm.Owner) *pb.Owner {
	if in == nil {
		return nil
	}
	out := &pb.Owner{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Email = direct.ValueOf(in.Email)
	return out
}
func Point_FromProto(mapCtx *direct.MapContext, in *pb.Point) *krm.Point {
	if in == nil {
		return nil
	}
	out := &krm.Point{}
	out.Line = direct.LazyPtr(in.GetLine())
	out.Character = direct.LazyPtr(in.GetCharacter())
	return out
}
func Point_ToProto(mapCtx *direct.MapContext, in *krm.Point) *pb.Point {
	if in == nil {
		return nil
	}
	out := &pb.Point{}
	out.Line = direct.ValueOf(in.Line)
	out.Character = direct.ValueOf(in.Character)
	return out
}
func Range_FromProto(mapCtx *direct.MapContext, in *pb.Range) *krm.Range {
	if in == nil {
		return nil
	}
	out := &krm.Range{}
	out.Start = Point_FromProto(mapCtx, in.GetStart())
	out.End = Point_FromProto(mapCtx, in.GetEnd())
	return out
}
func Range_ToProto(mapCtx *direct.MapContext, in *krm.Range) *pb.Range {
	if in == nil {
		return nil
	}
	out := &pb.Range{}
	out.Start = Point_ToProto(mapCtx, in.Start)
	out.End = Point_ToProto(mapCtx, in.End)
	return out
}
func Spec_FromProto(mapCtx *direct.MapContext, in *pb.Spec) *krm.Spec {
	if in == nil {
		return nil
	}
	out := &krm.Spec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.SpecType = AttributeValues_FromProto(mapCtx, in.GetSpecType())
	out.Contents = SpecContents_FromProto(mapCtx, in.GetContents())
	// MISSING: Details
	out.SourceURI = direct.LazyPtr(in.GetSourceUri())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.LintResponse = LintResponse_FromProto(mapCtx, in.GetLintResponse())
	// MISSING: Attributes
	out.Documentation = Documentation_FromProto(mapCtx, in.GetDocumentation())
	out.ParsingMode = direct.Enum_FromProto(mapCtx, in.GetParsingMode())
	return out
}
func Spec_ToProto(mapCtx *direct.MapContext, in *krm.Spec) *pb.Spec {
	if in == nil {
		return nil
	}
	out := &pb.Spec{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.SpecType = AttributeValues_ToProto(mapCtx, in.SpecType)
	out.Contents = SpecContents_ToProto(mapCtx, in.Contents)
	// MISSING: Details
	out.SourceUri = direct.ValueOf(in.SourceURI)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.LintResponse = LintResponse_ToProto(mapCtx, in.LintResponse)
	// MISSING: Attributes
	out.Documentation = Documentation_ToProto(mapCtx, in.Documentation)
	out.ParsingMode = direct.Enum_ToProto[pb.Spec_ParsingMode](mapCtx, in.ParsingMode)
	return out
}
func SpecContents_FromProto(mapCtx *direct.MapContext, in *pb.SpecContents) *krm.SpecContents {
	if in == nil {
		return nil
	}
	out := &krm.SpecContents{}
	out.Contents = in.GetContents()
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	return out
}
func SpecContents_ToProto(mapCtx *direct.MapContext, in *krm.SpecContents) *pb.SpecContents {
	if in == nil {
		return nil
	}
	out := &pb.SpecContents{}
	out.Contents = in.Contents
	out.MimeType = direct.ValueOf(in.MimeType)
	return out
}
func SpecDetails_FromProto(mapCtx *direct.MapContext, in *pb.SpecDetails) *krm.SpecDetails {
	if in == nil {
		return nil
	}
	out := &krm.SpecDetails{}
	// MISSING: OpenApiSpecDetails
	// MISSING: Description
	return out
}
func SpecDetails_ToProto(mapCtx *direct.MapContext, in *krm.SpecDetails) *pb.SpecDetails {
	if in == nil {
		return nil
	}
	out := &pb.SpecDetails{}
	// MISSING: OpenApiSpecDetails
	// MISSING: Description
	return out
}
func SpecDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SpecDetails) *krm.SpecDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpecDetailsObservedState{}
	out.OpenApiSpecDetails = OpenApiSpecDetails_FromProto(mapCtx, in.GetOpenApiSpecDetails())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func SpecDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpecDetailsObservedState) *pb.SpecDetails {
	if in == nil {
		return nil
	}
	out := &pb.SpecDetails{}
	if oneof := OpenApiSpecDetails_ToProto(mapCtx, in.OpenApiSpecDetails); oneof != nil {
		out.Details = &pb.SpecDetails_OpenApiSpecDetails{OpenApiSpecDetails: oneof}
	}
	out.Description = direct.ValueOf(in.Description)
	return out
}
func SpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Spec) *krm.SpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpecObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.SpecType = AttributeValuesObservedState_FromProto(mapCtx, in.GetSpecType())
	// MISSING: Contents
	out.Details = SpecDetails_FromProto(mapCtx, in.GetDetails())
	// MISSING: SourceURI
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: LintResponse
	// MISSING: Attributes
	// MISSING: Documentation
	// MISSING: ParsingMode
	return out
}
func SpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpecObservedState) *pb.Spec {
	if in == nil {
		return nil
	}
	out := &pb.Spec{}
	// MISSING: Name
	// MISSING: DisplayName
	out.SpecType = AttributeValuesObservedState_ToProto(mapCtx, in.SpecType)
	// MISSING: Contents
	out.Details = SpecDetails_ToProto(mapCtx, in.Details)
	// MISSING: SourceURI
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: LintResponse
	// MISSING: Attributes
	// MISSING: Documentation
	// MISSING: ParsingMode
	return out
}
