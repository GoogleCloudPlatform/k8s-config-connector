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
)

func DataCatalogTagObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Tag) *krm.DataCatalogTagObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogTagObservedState{}
	// MISSING: Name
	out.TemplateDisplayName = direct.LazyPtr(in.GetTemplateDisplayName())
	// TODO: map type string message for field Fields
	out.DataplexTransferStatus = direct.Enum_FromProto(mapCtx, in.GetDataplexTransferStatus())
	return out
}
func DataCatalogTagObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogTagObservedState) *pb.Tag {
	if in == nil {
		return nil
	}
	out := &pb.Tag{}
	// MISSING: Name
	out.TemplateDisplayName = direct.ValueOf(in.TemplateDisplayName)
	// TODO: map type string message for field Fields
	out.DataplexTransferStatus = direct.Enum_ToProto[pb.TagTemplate_DataplexTransferStatus](mapCtx, in.DataplexTransferStatus)
	return out
}
func DataCatalogTagSpec_FromProto(mapCtx *direct.MapContext, in *pb.Tag) *krm.DataCatalogTagSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogTagSpec{}
	// MISSING: Name
	out.Template = direct.LazyPtr(in.GetTemplate())
	out.Column = direct.LazyPtr(in.GetColumn())
	// TODO: map type string message for field Fields
	return out
}
func DataCatalogTagSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogTagSpec) *pb.Tag {
	if in == nil {
		return nil
	}
	out := &pb.Tag{}
	// MISSING: Name
	out.Template = DataCatalogTagSpec_Template_ToProto(mapCtx, in.Template)
	if oneof := DataCatalogTagSpec_Column_ToProto(mapCtx, in.Column); oneof != nil {
		out.Scope = oneof
	}
	// TODO: map type string message for field Fields
	return out
}
func Parent_FromProto(mapCtx *direct.MapContext, in *pb.Tag) *krm.Parent {
	if in == nil {
		return nil
	}
	out := &krm.Parent{}
	// MISSING: Name
	// MISSING: Template
	// MISSING: TemplateDisplayName
	// MISSING: Column
	// MISSING: Fields
	// MISSING: DataplexTransferStatus
	return out
}
func Parent_ToProto(mapCtx *direct.MapContext, in *krm.Parent) *pb.Tag {
	if in == nil {
		return nil
	}
	out := &pb.Tag{}
	// MISSING: Name
	// MISSING: Template
	// MISSING: TemplateDisplayName
	// MISSING: Column
	// MISSING: Fields
	// MISSING: DataplexTransferStatus
	return out
}
func TagField_FromProto(mapCtx *direct.MapContext, in *pb.TagField) *krm.TagField {
	if in == nil {
		return nil
	}
	out := &krm.TagField{}
	// MISSING: DisplayName
	out.DoubleValue = direct.LazyPtr(in.GetDoubleValue())
	out.StringValue = direct.LazyPtr(in.GetStringValue())
	out.BoolValue = direct.LazyPtr(in.GetBoolValue())
	out.TimestampValue = direct.StringTimestamp_FromProto(mapCtx, in.GetTimestampValue())
	out.EnumValue = TagField_EnumValue_FromProto(mapCtx, in.GetEnumValue())
	out.RichtextValue = direct.LazyPtr(in.GetRichtextValue())
	// MISSING: Order
	return out
}
func TagField_ToProto(mapCtx *direct.MapContext, in *krm.TagField) *pb.TagField {
	if in == nil {
		return nil
	}
	out := &pb.TagField{}
	// MISSING: DisplayName
	if oneof := TagField_DoubleValue_ToProto(mapCtx, in.DoubleValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := TagField_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := TagField_BoolValue_ToProto(mapCtx, in.BoolValue); oneof != nil {
		out.Kind = oneof
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.TimestampValue); oneof != nil {
		out.Kind = &pb.TagField_TimestampValue{TimestampValue: oneof}
	}
	if oneof := TagField_EnumValue_ToProto(mapCtx, in.EnumValue); oneof != nil {
		out.Kind = &pb.TagField_EnumValue_{EnumValue: oneof}
	}
	if oneof := TagField_RichtextValue_ToProto(mapCtx, in.RichtextValue); oneof != nil {
		out.Kind = oneof
	}
	// MISSING: Order
	return out
}
func TagFieldObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TagField) *krm.TagFieldObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TagFieldObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: DoubleValue
	// MISSING: StringValue
	// MISSING: BoolValue
	// MISSING: TimestampValue
	// MISSING: EnumValue
	// MISSING: RichtextValue
	// MISSING: Order
	return out
}
func TagFieldObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TagFieldObservedState) *pb.TagField {
	if in == nil {
		return nil
	}
	out := &pb.TagField{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: DoubleValue
	// MISSING: StringValue
	// MISSING: BoolValue
	// MISSING: TimestampValue
	// MISSING: EnumValue
	// MISSING: RichtextValue
	// MISSING: Order
	return out
}
func TagField_EnumValue_FromProto(mapCtx *direct.MapContext, in *pb.TagField_EnumValue) *krm.TagField_EnumValue {
	if in == nil {
		return nil
	}
	out := &krm.TagField_EnumValue{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func TagField_EnumValue_ToProto(mapCtx *direct.MapContext, in *krm.TagField_EnumValue) *pb.TagField_EnumValue {
	if in == nil {
		return nil
	}
	out := &pb.TagField_EnumValue{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
