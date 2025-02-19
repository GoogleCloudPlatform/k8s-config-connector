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
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Api_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.Api {
	if in == nil {
		return nil
	}
	out := &krm.Api{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Documentation = Documentation_FromProto(mapCtx, in.GetDocumentation())
	out.Owner = Owner_FromProto(mapCtx, in.GetOwner())
	// MISSING: Versions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.TargetUser = AttributeValues_FromProto(mapCtx, in.GetTargetUser())
	out.Team = AttributeValues_FromProto(mapCtx, in.GetTeam())
	out.BusinessUnit = AttributeValues_FromProto(mapCtx, in.GetBusinessUnit())
	out.MaturityLevel = AttributeValues_FromProto(mapCtx, in.GetMaturityLevel())
	// MISSING: Attributes
	out.ApiStyle = AttributeValues_FromProto(mapCtx, in.GetApiStyle())
	out.SelectedVersion = direct.LazyPtr(in.GetSelectedVersion())
	return out
}
func Api_ToProto(mapCtx *direct.MapContext, in *krm.Api) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Documentation = Documentation_ToProto(mapCtx, in.Documentation)
	out.Owner = Owner_ToProto(mapCtx, in.Owner)
	// MISSING: Versions
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.TargetUser = AttributeValues_ToProto(mapCtx, in.TargetUser)
	out.Team = AttributeValues_ToProto(mapCtx, in.Team)
	out.BusinessUnit = AttributeValues_ToProto(mapCtx, in.BusinessUnit)
	out.MaturityLevel = AttributeValues_ToProto(mapCtx, in.MaturityLevel)
	// MISSING: Attributes
	out.ApiStyle = AttributeValues_ToProto(mapCtx, in.ApiStyle)
	out.SelectedVersion = direct.ValueOf(in.SelectedVersion)
	return out
}
func ApiObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.ApiObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApiObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: Owner
	out.Versions = in.Versions
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.TargetUser = AttributeValuesObservedState_FromProto(mapCtx, in.GetTargetUser())
	// MISSING: Team
	// MISSING: BusinessUnit
	// MISSING: MaturityLevel
	// MISSING: Attributes
	// MISSING: ApiStyle
	// MISSING: SelectedVersion
	return out
}
func ApiObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApiObservedState) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Documentation
	// MISSING: Owner
	out.Versions = in.Versions
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.TargetUser = AttributeValuesObservedState_ToProto(mapCtx, in.TargetUser)
	// MISSING: Team
	// MISSING: BusinessUnit
	// MISSING: MaturityLevel
	// MISSING: Attributes
	// MISSING: ApiStyle
	// MISSING: SelectedVersion
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
