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

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataplexAspectTypeSpec_FromProto(mapCtx *direct.MapContext, in *pb.AspectType) *krm.AspectType {
	if in == nil {
		return nil
	}
	out := &krm.AspectType{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Authorization = AspectType_Authorization_FromProto(mapCtx, in.GetAuthorization())
	out.MetadataTemplate = AspectType_MetadataTemplate_FromProto(mapCtx, in.GetMetadataTemplate())
	// MISSING: TransferStatus
	return out
}
func DataplexAspectTypeSpec_ToProto(mapCtx *direct.MapContext, in *krm.AspectType) *pb.AspectType {
	if in == nil {
		return nil
	}
	out := &pb.AspectType{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	out.Etag = direct.ValueOf(in.Etag)
	out.Authorization = AspectType_Authorization_ToProto(mapCtx, in.Authorization)
	out.MetadataTemplate = AspectType_MetadataTemplate_ToProto(mapCtx, in.MetadataTemplate)
	// MISSING: TransferStatus
	return out
}
func DataplexAspectTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AspectType) *krm.AspectTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AspectTypeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Authorization
	// MISSING: MetadataTemplate
	out.TransferStatus = direct.Enum_FromProto(mapCtx, in.GetTransferStatus())
	return out
}
func DataplexAspectTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AspectTypeObservedState) *pb.AspectType {
	if in == nil {
		return nil
	}
	out := &pb.AspectType{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: DisplayName
	// MISSING: Labels
	// MISSING: Etag
	// MISSING: Authorization
	// MISSING: MetadataTemplate
	out.TransferStatus = direct.Enum_ToProto[pb.TransferStatus](mapCtx, in.TransferStatus)
	return out
}
func AspectType_Authorization_FromProto(mapCtx *direct.MapContext, in *pb.AspectType_Authorization) *krm.AspectType_Authorization {
	if in == nil {
		return nil
	}
	out := &krm.AspectType_Authorization{}
	out.AlternateUsePermission = direct.LazyPtr(in.GetAlternateUsePermission())
	return out
}
func AspectType_Authorization_ToProto(mapCtx *direct.MapContext, in *krm.AspectType_Authorization) *pb.AspectType_Authorization {
	if in == nil {
		return nil
	}
	out := &pb.AspectType_Authorization{}
	out.AlternateUsePermission = direct.ValueOf(in.AlternateUsePermission)
	return out
}
func AspectType_MetadataTemplate_FromProto(mapCtx *direct.MapContext, in *pb.AspectType_MetadataTemplate) *krm.AspectType_MetadataTemplate {
	if in == nil {
		return nil
	}
	out := &krm.AspectType_MetadataTemplate{}
	out.Index = direct.LazyPtr(in.GetIndex())
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	out.RecordFields = direct.Slice_FromProto(mapCtx, in.RecordFields, AspectType_MetadataTemplate_FromProto)
	out.EnumValues = direct.Slice_FromProto(mapCtx, in.EnumValues, AspectType_MetadataTemplate_EnumValue_FromProto)
	out.MapItems = AspectType_MetadataTemplate_FromProto(mapCtx, in.GetMapItems())
	out.ArrayItems = AspectType_MetadataTemplate_FromProto(mapCtx, in.GetArrayItems())
	out.TypeID = direct.LazyPtr(in.GetTypeId())
	out.TypeRef = direct.LazyPtr(in.GetTypeRef())
	out.Constraints = AspectType_MetadataTemplate_Constraints_FromProto(mapCtx, in.GetConstraints())
	out.Annotations = AspectType_MetadataTemplate_Annotations_FromProto(mapCtx, in.GetAnnotations())
	return out
}
func AspectType_MetadataTemplate_ToProto(mapCtx *direct.MapContext, in *krm.AspectType_MetadataTemplate) *pb.AspectType_MetadataTemplate {
	if in == nil {
		return nil
	}
	out := &pb.AspectType_MetadataTemplate{}
	out.Index = direct.ValueOf(in.Index)
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)
	out.RecordFields = direct.Slice_ToProto(mapCtx, in.RecordFields, AspectType_MetadataTemplate_ToProto)
	out.EnumValues = direct.Slice_ToProto(mapCtx, in.EnumValues, AspectType_MetadataTemplate_EnumValue_ToProto)
	out.MapItems = AspectType_MetadataTemplate_ToProto(mapCtx, in.MapItems)
	out.ArrayItems = AspectType_MetadataTemplate_ToProto(mapCtx, in.ArrayItems)
	out.TypeId = direct.ValueOf(in.TypeID)
	out.TypeRef = direct.ValueOf(in.TypeRef)
	out.Constraints = AspectType_MetadataTemplate_Constraints_ToProto(mapCtx, in.Constraints)
	out.Annotations = AspectType_MetadataTemplate_Annotations_ToProto(mapCtx, in.Annotations)
	return out
}
func AspectType_MetadataTemplate_Annotations_FromProto(mapCtx *direct.MapContext, in *pb.AspectType_MetadataTemplate_Annotations) *krm.AspectType_MetadataTemplate_Annotations {
	if in == nil {
		return nil
	}
	out := &krm.AspectType_MetadataTemplate_Annotations{}
	out.Deprecated = direct.LazyPtr(in.GetDeprecated())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayOrder = direct.LazyPtr(in.GetDisplayOrder())
	out.StringType = direct.LazyPtr(in.GetStringType())
	out.StringValues = in.StringValues
	return out
}
func AspectType_MetadataTemplate_Annotations_ToProto(mapCtx *direct.MapContext, in *krm.AspectType_MetadataTemplate_Annotations) *pb.AspectType_MetadataTemplate_Annotations {
	if in == nil {
		return nil
	}
	out := &pb.AspectType_MetadataTemplate_Annotations{}
	out.Deprecated = direct.ValueOf(in.Deprecated)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.DisplayOrder = direct.ValueOf(in.DisplayOrder)
	out.StringType = direct.ValueOf(in.StringType)
	out.StringValues = in.StringValues
	return out
}
func AspectType_MetadataTemplate_Constraints_FromProto(mapCtx *direct.MapContext, in *pb.AspectType_MetadataTemplate_Constraints) *krm.AspectType_MetadataTemplate_Constraints {
	if in == nil {
		return nil
	}
	out := &krm.AspectType_MetadataTemplate_Constraints{}
	out.Required = direct.LazyPtr(in.GetRequired())
	return out
}
func AspectType_MetadataTemplate_Constraints_ToProto(mapCtx *direct.MapContext, in *krm.AspectType_MetadataTemplate_Constraints) *pb.AspectType_MetadataTemplate_Constraints {
	if in == nil {
		return nil
	}
	out := &pb.AspectType_MetadataTemplate_Constraints{}
	out.Required = direct.ValueOf(in.Required)
	return out
}
func AspectType_MetadataTemplate_EnumValue_FromProto(mapCtx *direct.MapContext, in *pb.AspectType_MetadataTemplate_EnumValue) *krm.AspectType_MetadataTemplate_EnumValue {
	if in == nil {
		return nil
	}
	out := &krm.AspectType_MetadataTemplate_EnumValue{}
	out.Index = direct.LazyPtr(in.GetIndex())
	out.Name = direct.LazyPtr(in.GetName())
	out.Deprecated = direct.LazyPtr(in.GetDeprecated())
	return out
}
func AspectType_MetadataTemplate_EnumValue_ToProto(mapCtx *direct.MapContext, in *krm.AspectType_MetadataTemplate_EnumValue) *pb.AspectType_MetadataTemplate_EnumValue {
	if in == nil {
		return nil
	}
	out := &pb.AspectType_MetadataTemplate_EnumValue{}
	out.Index = direct.ValueOf(in.Index)
	out.Name = direct.ValueOf(in.Name)
	out.Deprecated = direct.ValueOf(in.Deprecated)
	return out
}
