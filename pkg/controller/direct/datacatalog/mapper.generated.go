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

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func DatacatalogTagTemplateFieldObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TagTemplateField) *krm.DatacatalogTagTemplateFieldObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatacatalogTagTemplateFieldObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Type
	// MISSING: IsRequired
	// MISSING: Description
	// MISSING: Order
	return out
}
func DatacatalogTagTemplateFieldObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatacatalogTagTemplateFieldObservedState) *pb.TagTemplateField {
	if in == nil {
		return nil
	}
	out := &pb.TagTemplateField{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Type
	// MISSING: IsRequired
	// MISSING: Description
	// MISSING: Order
	return out
}
func DatacatalogTagTemplateFieldSpec_FromProto(mapCtx *direct.MapContext, in *pb.TagTemplateField) *krm.DatacatalogTagTemplateFieldSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatacatalogTagTemplateFieldSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Type
	// MISSING: IsRequired
	// MISSING: Description
	// MISSING: Order
	return out
}
func DatacatalogTagTemplateFieldSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatacatalogTagTemplateFieldSpec) *pb.TagTemplateField {
	if in == nil {
		return nil
	}
	out := &pb.TagTemplateField{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Type
	// MISSING: IsRequired
	// MISSING: Description
	// MISSING: Order
	return out
}
func FieldType_FromProto(mapCtx *direct.MapContext, in *pb.FieldType) *krm.FieldType {
	if in == nil {
		return nil
	}
	out := &krm.FieldType{}
	out.PrimitiveType = direct.Enum_FromProto(mapCtx, in.GetPrimitiveType())
	out.EnumType = FieldType_EnumType_FromProto(mapCtx, in.GetEnumType())
	return out
}
func FieldType_ToProto(mapCtx *direct.MapContext, in *krm.FieldType) *pb.FieldType {
	if in == nil {
		return nil
	}
	out := &pb.FieldType{}
	if oneof := FieldType_PrimitiveType_ToProto(mapCtx, in.PrimitiveType); oneof != nil {
		out.TypeDecl = oneof
	}
	if oneof := FieldType_EnumType_ToProto(mapCtx, in.EnumType); oneof != nil {
		out.TypeDecl = &pb.FieldType_EnumType_{EnumType: oneof}
	}
	return out
}
func FieldType_EnumType_FromProto(mapCtx *direct.MapContext, in *pb.FieldType_EnumType) *krm.FieldType_EnumType {
	if in == nil {
		return nil
	}
	out := &krm.FieldType_EnumType{}
	out.AllowedValues = direct.Slice_FromProto(mapCtx, in.AllowedValues, FieldType_EnumType_EnumValue_FromProto)
	return out
}
func FieldType_EnumType_ToProto(mapCtx *direct.MapContext, in *krm.FieldType_EnumType) *pb.FieldType_EnumType {
	if in == nil {
		return nil
	}
	out := &pb.FieldType_EnumType{}
	out.AllowedValues = direct.Slice_ToProto(mapCtx, in.AllowedValues, FieldType_EnumType_EnumValue_ToProto)
	return out
}
func FieldType_EnumType_EnumValue_FromProto(mapCtx *direct.MapContext, in *pb.FieldType_EnumType_EnumValue) *krm.FieldType_EnumType_EnumValue {
	if in == nil {
		return nil
	}
	out := &krm.FieldType_EnumType_EnumValue{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func FieldType_EnumType_EnumValue_ToProto(mapCtx *direct.MapContext, in *krm.FieldType_EnumType_EnumValue) *pb.FieldType_EnumType_EnumValue {
	if in == nil {
		return nil
	}
	out := &pb.FieldType_EnumType_EnumValue{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func TagTemplateField_FromProto(mapCtx *direct.MapContext, in *pb.TagTemplateField) *krm.TagTemplateField {
	if in == nil {
		return nil
	}
	out := &krm.TagTemplateField{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Type = FieldType_FromProto(mapCtx, in.GetType())
	out.IsRequired = direct.LazyPtr(in.GetIsRequired())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Order = direct.LazyPtr(in.GetOrder())
	return out
}
func TagTemplateField_ToProto(mapCtx *direct.MapContext, in *krm.TagTemplateField) *pb.TagTemplateField {
	if in == nil {
		return nil
	}
	out := &pb.TagTemplateField{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Type = FieldType_ToProto(mapCtx, in.Type)
	out.IsRequired = direct.ValueOf(in.IsRequired)
	out.Description = direct.ValueOf(in.Description)
	out.Order = direct.ValueOf(in.Order)
	return out
}
