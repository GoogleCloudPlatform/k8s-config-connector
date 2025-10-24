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

// +generated:mapper
// krm.group: apihub.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.apihub.v1

package apihub

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func APIHubAPIObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.APIHubAPIObservedState {
	if in == nil {
		return nil
	}
	out := &krm.APIHubAPIObservedState{}
	// MISSING: Name
	out.Versions = in.Versions
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.TargetUser = AttributeValuesObservedState_FromProto(mapCtx, in.GetTargetUser())
	// MISSING: Attributes
	return out
}
func APIHubAPIObservedState_ToProto(mapCtx *direct.MapContext, in *krm.APIHubAPIObservedState) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	// MISSING: Name
	out.Versions = in.Versions
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.TargetUser = AttributeValuesObservedState_ToProto(mapCtx, in.TargetUser)
	// MISSING: Attributes
	return out
}
func APIHubAPISpec_FromProto(mapCtx *direct.MapContext, in *pb.Api) *krm.APIHubAPISpec {
	if in == nil {
		return nil
	}
	out := &krm.APIHubAPISpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Documentation = Documentation_FromProto(mapCtx, in.GetDocumentation())
	out.Owner = Owner_FromProto(mapCtx, in.GetOwner())
	out.TargetUser = AttributeValues_FromProto(mapCtx, in.GetTargetUser())
	out.Team = AttributeValues_FromProto(mapCtx, in.GetTeam())
	out.BusinessUnit = AttributeValues_FromProto(mapCtx, in.GetBusinessUnit())
	out.MaturityLevel = AttributeValues_FromProto(mapCtx, in.GetMaturityLevel())
	// MISSING: Attributes
	out.APIStyle = AttributeValues_FromProto(mapCtx, in.GetApiStyle())
	out.SelectedVersion = direct.LazyPtr(in.GetSelectedVersion())
	return out
}
func APIHubAPISpec_ToProto(mapCtx *direct.MapContext, in *krm.APIHubAPISpec) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Documentation = Documentation_ToProto(mapCtx, in.Documentation)
	out.Owner = Owner_ToProto(mapCtx, in.Owner)
	out.TargetUser = AttributeValues_ToProto(mapCtx, in.TargetUser)
	out.Team = AttributeValues_ToProto(mapCtx, in.Team)
	out.BusinessUnit = AttributeValues_ToProto(mapCtx, in.BusinessUnit)
	out.MaturityLevel = AttributeValues_ToProto(mapCtx, in.MaturityLevel)
	// MISSING: Attributes
	out.ApiStyle = AttributeValues_ToProto(mapCtx, in.APIStyle)
	out.SelectedVersion = direct.ValueOf(in.SelectedVersion)
	return out
}
func AttributeValues_FromProto(mapCtx *direct.MapContext, in *pb.AttributeValues) *krm.AttributeValues {
	if in == nil {
		return nil
	}
	out := &krm.AttributeValues{}
	// MISSING: Attribute
	switch v := in.Value.(type) {
	case *pb.AttributeValues_EnumValues:
		out.EnumValues = AttributeValues_EnumAttributeValues_FromProto(mapCtx, v.EnumValues)
	case *pb.AttributeValues_StringValues:
		out.StringValues = AttributeValues_StringAttributeValues_FromProto(mapCtx, v.StringValues)
	case *pb.AttributeValues_JsonValues:
		out.JsonValues = AttributeValues_StringAttributeValues_FromProto(mapCtx, v.JsonValues)
	}
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
	out.Attribute = direct.LazyPtr(in.GetAttribute())
	switch v := in.Value.(type) {
	case *pb.AttributeValues_EnumValues:
		// MISSING "EnumValues"
	case *pb.AttributeValues_StringValues:
		// MISSING "StringValues"
	case *pb.AttributeValues_JsonValues:
		// MISSING "JsonValues"
	}
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
