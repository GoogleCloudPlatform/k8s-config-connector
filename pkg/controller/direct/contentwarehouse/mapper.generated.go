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

package contentwarehouse

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/contentwarehouse/apiv1/contentwarehousepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ContentwarehouseDocumentSchemaObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema) *krm.ContentwarehouseDocumentSchemaObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContentwarehouseDocumentSchemaObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: PropertyDefinitions
	// MISSING: DocumentIsFolder
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: Description
	return out
}
func ContentwarehouseDocumentSchemaObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContentwarehouseDocumentSchemaObservedState) *pb.DocumentSchema {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: PropertyDefinitions
	// MISSING: DocumentIsFolder
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: Description
	return out
}
func ContentwarehouseDocumentSchemaSpec_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema) *krm.ContentwarehouseDocumentSchemaSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContentwarehouseDocumentSchemaSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: PropertyDefinitions
	// MISSING: DocumentIsFolder
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: Description
	return out
}
func ContentwarehouseDocumentSchemaSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContentwarehouseDocumentSchemaSpec) *pb.DocumentSchema {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: PropertyDefinitions
	// MISSING: DocumentIsFolder
	// MISSING: UpdateTime
	// MISSING: CreateTime
	// MISSING: Description
	return out
}
func DateTimeTypeOptions_FromProto(mapCtx *direct.MapContext, in *pb.DateTimeTypeOptions) *krm.DateTimeTypeOptions {
	if in == nil {
		return nil
	}
	out := &krm.DateTimeTypeOptions{}
	return out
}
func DateTimeTypeOptions_ToProto(mapCtx *direct.MapContext, in *krm.DateTimeTypeOptions) *pb.DateTimeTypeOptions {
	if in == nil {
		return nil
	}
	out := &pb.DateTimeTypeOptions{}
	return out
}
func DocumentSchema_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema) *krm.DocumentSchema {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.PropertyDefinitions = direct.Slice_FromProto(mapCtx, in.PropertyDefinitions, PropertyDefinition_FromProto)
	out.DocumentIsFolder = direct.LazyPtr(in.GetDocumentIsFolder())
	// MISSING: UpdateTime
	// MISSING: CreateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func DocumentSchema_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema) *pb.DocumentSchema {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.PropertyDefinitions = direct.Slice_ToProto(mapCtx, in.PropertyDefinitions, PropertyDefinition_ToProto)
	out.DocumentIsFolder = direct.ValueOf(in.DocumentIsFolder)
	// MISSING: UpdateTime
	// MISSING: CreateTime
	out.Description = direct.ValueOf(in.Description)
	return out
}
func DocumentSchemaObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema) *krm.DocumentSchemaObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchemaObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: PropertyDefinitions
	// MISSING: DocumentIsFolder
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Description
	return out
}
func DocumentSchemaObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchemaObservedState) *pb.DocumentSchema {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: PropertyDefinitions
	// MISSING: DocumentIsFolder
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Description
	return out
}
func EnumTypeOptions_FromProto(mapCtx *direct.MapContext, in *pb.EnumTypeOptions) *krm.EnumTypeOptions {
	if in == nil {
		return nil
	}
	out := &krm.EnumTypeOptions{}
	out.PossibleValues = in.PossibleValues
	out.ValidationCheckDisabled = direct.LazyPtr(in.GetValidationCheckDisabled())
	return out
}
func EnumTypeOptions_ToProto(mapCtx *direct.MapContext, in *krm.EnumTypeOptions) *pb.EnumTypeOptions {
	if in == nil {
		return nil
	}
	out := &pb.EnumTypeOptions{}
	out.PossibleValues = in.PossibleValues
	out.ValidationCheckDisabled = direct.ValueOf(in.ValidationCheckDisabled)
	return out
}
func FloatTypeOptions_FromProto(mapCtx *direct.MapContext, in *pb.FloatTypeOptions) *krm.FloatTypeOptions {
	if in == nil {
		return nil
	}
	out := &krm.FloatTypeOptions{}
	return out
}
func FloatTypeOptions_ToProto(mapCtx *direct.MapContext, in *krm.FloatTypeOptions) *pb.FloatTypeOptions {
	if in == nil {
		return nil
	}
	out := &pb.FloatTypeOptions{}
	return out
}
func IntegerTypeOptions_FromProto(mapCtx *direct.MapContext, in *pb.IntegerTypeOptions) *krm.IntegerTypeOptions {
	if in == nil {
		return nil
	}
	out := &krm.IntegerTypeOptions{}
	return out
}
func IntegerTypeOptions_ToProto(mapCtx *direct.MapContext, in *krm.IntegerTypeOptions) *pb.IntegerTypeOptions {
	if in == nil {
		return nil
	}
	out := &pb.IntegerTypeOptions{}
	return out
}
func MapTypeOptions_FromProto(mapCtx *direct.MapContext, in *pb.MapTypeOptions) *krm.MapTypeOptions {
	if in == nil {
		return nil
	}
	out := &krm.MapTypeOptions{}
	return out
}
func MapTypeOptions_ToProto(mapCtx *direct.MapContext, in *krm.MapTypeOptions) *pb.MapTypeOptions {
	if in == nil {
		return nil
	}
	out := &pb.MapTypeOptions{}
	return out
}
func PropertyDefinition_FromProto(mapCtx *direct.MapContext, in *pb.PropertyDefinition) *krm.PropertyDefinition {
	if in == nil {
		return nil
	}
	out := &krm.PropertyDefinition{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.IsRepeatable = direct.LazyPtr(in.GetIsRepeatable())
	out.IsFilterable = direct.LazyPtr(in.GetIsFilterable())
	out.IsSearchable = direct.LazyPtr(in.GetIsSearchable())
	out.IsMetadata = direct.LazyPtr(in.GetIsMetadata())
	out.IsRequired = direct.LazyPtr(in.GetIsRequired())
	out.RetrievalImportance = direct.Enum_FromProto(mapCtx, in.GetRetrievalImportance())
	out.IntegerTypeOptions = IntegerTypeOptions_FromProto(mapCtx, in.GetIntegerTypeOptions())
	out.FloatTypeOptions = FloatTypeOptions_FromProto(mapCtx, in.GetFloatTypeOptions())
	out.TextTypeOptions = TextTypeOptions_FromProto(mapCtx, in.GetTextTypeOptions())
	out.PropertyTypeOptions = PropertyTypeOptions_FromProto(mapCtx, in.GetPropertyTypeOptions())
	out.EnumTypeOptions = EnumTypeOptions_FromProto(mapCtx, in.GetEnumTypeOptions())
	out.DateTimeTypeOptions = DateTimeTypeOptions_FromProto(mapCtx, in.GetDateTimeTypeOptions())
	out.MapTypeOptions = MapTypeOptions_FromProto(mapCtx, in.GetMapTypeOptions())
	out.TimestampTypeOptions = TimestampTypeOptions_FromProto(mapCtx, in.GetTimestampTypeOptions())
	out.SchemaSources = direct.Slice_FromProto(mapCtx, in.SchemaSources, PropertyDefinition_SchemaSource_FromProto)
	return out
}
func PropertyDefinition_ToProto(mapCtx *direct.MapContext, in *krm.PropertyDefinition) *pb.PropertyDefinition {
	if in == nil {
		return nil
	}
	out := &pb.PropertyDefinition{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IsRepeatable = direct.ValueOf(in.IsRepeatable)
	out.IsFilterable = direct.ValueOf(in.IsFilterable)
	out.IsSearchable = direct.ValueOf(in.IsSearchable)
	out.IsMetadata = direct.ValueOf(in.IsMetadata)
	out.IsRequired = direct.ValueOf(in.IsRequired)
	out.RetrievalImportance = direct.Enum_ToProto[pb.PropertyDefinition_RetrievalImportance](mapCtx, in.RetrievalImportance)
	if oneof := IntegerTypeOptions_ToProto(mapCtx, in.IntegerTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_IntegerTypeOptions{IntegerTypeOptions: oneof}
	}
	if oneof := FloatTypeOptions_ToProto(mapCtx, in.FloatTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_FloatTypeOptions{FloatTypeOptions: oneof}
	}
	if oneof := TextTypeOptions_ToProto(mapCtx, in.TextTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_TextTypeOptions{TextTypeOptions: oneof}
	}
	if oneof := PropertyTypeOptions_ToProto(mapCtx, in.PropertyTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_PropertyTypeOptions{PropertyTypeOptions: oneof}
	}
	if oneof := EnumTypeOptions_ToProto(mapCtx, in.EnumTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_EnumTypeOptions{EnumTypeOptions: oneof}
	}
	if oneof := DateTimeTypeOptions_ToProto(mapCtx, in.DateTimeTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_DateTimeTypeOptions{DateTimeTypeOptions: oneof}
	}
	if oneof := MapTypeOptions_ToProto(mapCtx, in.MapTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_MapTypeOptions{MapTypeOptions: oneof}
	}
	if oneof := TimestampTypeOptions_ToProto(mapCtx, in.TimestampTypeOptions); oneof != nil {
		out.ValueTypeOptions = &pb.PropertyDefinition_TimestampTypeOptions{TimestampTypeOptions: oneof}
	}
	out.SchemaSources = direct.Slice_ToProto(mapCtx, in.SchemaSources, PropertyDefinition_SchemaSource_ToProto)
	return out
}
func PropertyDefinition_SchemaSource_FromProto(mapCtx *direct.MapContext, in *pb.PropertyDefinition_SchemaSource) *krm.PropertyDefinition_SchemaSource {
	if in == nil {
		return nil
	}
	out := &krm.PropertyDefinition_SchemaSource{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ProcessorType = direct.LazyPtr(in.GetProcessorType())
	return out
}
func PropertyDefinition_SchemaSource_ToProto(mapCtx *direct.MapContext, in *krm.PropertyDefinition_SchemaSource) *pb.PropertyDefinition_SchemaSource {
	if in == nil {
		return nil
	}
	out := &pb.PropertyDefinition_SchemaSource{}
	out.Name = direct.ValueOf(in.Name)
	out.ProcessorType = direct.ValueOf(in.ProcessorType)
	return out
}
func PropertyTypeOptions_FromProto(mapCtx *direct.MapContext, in *pb.PropertyTypeOptions) *krm.PropertyTypeOptions {
	if in == nil {
		return nil
	}
	out := &krm.PropertyTypeOptions{}
	out.PropertyDefinitions = direct.Slice_FromProto(mapCtx, in.PropertyDefinitions, PropertyDefinition_FromProto)
	return out
}
func PropertyTypeOptions_ToProto(mapCtx *direct.MapContext, in *krm.PropertyTypeOptions) *pb.PropertyTypeOptions {
	if in == nil {
		return nil
	}
	out := &pb.PropertyTypeOptions{}
	out.PropertyDefinitions = direct.Slice_ToProto(mapCtx, in.PropertyDefinitions, PropertyDefinition_ToProto)
	return out
}
func TextTypeOptions_FromProto(mapCtx *direct.MapContext, in *pb.TextTypeOptions) *krm.TextTypeOptions {
	if in == nil {
		return nil
	}
	out := &krm.TextTypeOptions{}
	return out
}
func TextTypeOptions_ToProto(mapCtx *direct.MapContext, in *krm.TextTypeOptions) *pb.TextTypeOptions {
	if in == nil {
		return nil
	}
	out := &pb.TextTypeOptions{}
	return out
}
func TimestampTypeOptions_FromProto(mapCtx *direct.MapContext, in *pb.TimestampTypeOptions) *krm.TimestampTypeOptions {
	if in == nil {
		return nil
	}
	out := &krm.TimestampTypeOptions{}
	return out
}
func TimestampTypeOptions_ToProto(mapCtx *direct.MapContext, in *krm.TimestampTypeOptions) *pb.TimestampTypeOptions {
	if in == nil {
		return nil
	}
	out := &pb.TimestampTypeOptions{}
	return out
}
