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

package documentai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/documentai/apiv1beta3/documentaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DatasetSchema_FromProto(mapCtx *direct.MapContext, in *pb.DatasetSchema) *krm.DatasetSchema {
	if in == nil {
		return nil
	}
	out := &krm.DatasetSchema{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DocumentSchema = DocumentSchema_FromProto(mapCtx, in.GetDocumentSchema())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DatasetSchema_ToProto(mapCtx *direct.MapContext, in *krm.DatasetSchema) *pb.DatasetSchema {
	if in == nil {
		return nil
	}
	out := &pb.DatasetSchema{}
	out.Name = direct.ValueOf(in.Name)
	out.DocumentSchema = DocumentSchema_ToProto(mapCtx, in.DocumentSchema)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DatasetSchemaObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DatasetSchema) *krm.DatasetSchemaObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatasetSchemaObservedState{}
	// MISSING: Name
	// MISSING: DocumentSchema
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func DatasetSchemaObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatasetSchemaObservedState) *pb.DatasetSchema {
	if in == nil {
		return nil
	}
	out := &pb.DatasetSchema{}
	// MISSING: Name
	// MISSING: DocumentSchema
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func DocumentSchema_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema) *krm.DocumentSchema {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.EntityTypes = direct.Slice_FromProto(mapCtx, in.EntityTypes, DocumentSchema_EntityType_FromProto)
	out.Metadata = DocumentSchema_Metadata_FromProto(mapCtx, in.GetMetadata())
	return out
}
func DocumentSchema_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema) *pb.DocumentSchema {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.EntityTypes = direct.Slice_ToProto(mapCtx, in.EntityTypes, DocumentSchema_EntityType_ToProto)
	out.Metadata = DocumentSchema_Metadata_ToProto(mapCtx, in.Metadata)
	return out
}
func DocumentSchema_EntityType_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema_EntityType) *krm.DocumentSchema_EntityType {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema_EntityType{}
	out.EnumValues = DocumentSchema_EntityType_EnumValues_FromProto(mapCtx, in.GetEnumValues())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.BaseTypes = in.BaseTypes
	out.Properties = direct.Slice_FromProto(mapCtx, in.Properties, DocumentSchema_EntityType_Property_FromProto)
	out.EntityTypeMetadata = EntityTypeMetadata_FromProto(mapCtx, in.GetEntityTypeMetadata())
	return out
}
func DocumentSchema_EntityType_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema_EntityType) *pb.DocumentSchema_EntityType {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema_EntityType{}
	if oneof := DocumentSchema_EntityType_EnumValues_ToProto(mapCtx, in.EnumValues); oneof != nil {
		out.ValueSource = &pb.DocumentSchema_EntityType_EnumValues_{EnumValues: oneof}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.BaseTypes = in.BaseTypes
	out.Properties = direct.Slice_ToProto(mapCtx, in.Properties, DocumentSchema_EntityType_Property_ToProto)
	out.EntityTypeMetadata = EntityTypeMetadata_ToProto(mapCtx, in.EntityTypeMetadata)
	return out
}
func DocumentSchema_EntityType_EnumValues_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema_EntityType_EnumValues) *krm.DocumentSchema_EntityType_EnumValues {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema_EntityType_EnumValues{}
	out.Values = in.Values
	return out
}
func DocumentSchema_EntityType_EnumValues_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema_EntityType_EnumValues) *pb.DocumentSchema_EntityType_EnumValues {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema_EntityType_EnumValues{}
	out.Values = in.Values
	return out
}
func DocumentSchema_EntityType_Property_FromProto(mapCtx *direct.MapContext, in *pb.DocumentSchema_EntityType_Property) *krm.DocumentSchema_EntityType_Property {
	if in == nil {
		return nil
	}
	out := &krm.DocumentSchema_EntityType_Property{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ValueType = direct.LazyPtr(in.GetValueType())
	out.OccurrenceType = direct.Enum_FromProto(mapCtx, in.GetOccurrenceType())
	out.PropertyMetadata = PropertyMetadata_FromProto(mapCtx, in.GetPropertyMetadata())
	return out
}
func DocumentSchema_EntityType_Property_ToProto(mapCtx *direct.MapContext, in *krm.DocumentSchema_EntityType_Property) *pb.DocumentSchema_EntityType_Property {
	if in == nil {
		return nil
	}
	out := &pb.DocumentSchema_EntityType_Property{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ValueType = direct.ValueOf(in.ValueType)
	out.OccurrenceType = direct.Enum_ToProto[pb.DocumentSchema_EntityType_Property_OccurrenceType](mapCtx, in.OccurrenceType)
	out.PropertyMetadata = PropertyMetadata_ToProto(mapCtx, in.PropertyMetadata)
	return out
}
func DocumentaiDatasetSchemaObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DatasetSchema) *krm.DocumentaiDatasetSchemaObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DocumentaiDatasetSchemaObservedState{}
	// MISSING: Name
	// MISSING: DocumentSchema
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DocumentaiDatasetSchemaObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DocumentaiDatasetSchemaObservedState) *pb.DatasetSchema {
	if in == nil {
		return nil
	}
	out := &pb.DatasetSchema{}
	// MISSING: Name
	// MISSING: DocumentSchema
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DocumentaiDatasetSchemaSpec_FromProto(mapCtx *direct.MapContext, in *pb.DatasetSchema) *krm.DocumentaiDatasetSchemaSpec {
	if in == nil {
		return nil
	}
	out := &krm.DocumentaiDatasetSchemaSpec{}
	// MISSING: Name
	// MISSING: DocumentSchema
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DocumentaiDatasetSchemaSpec_ToProto(mapCtx *direct.MapContext, in *krm.DocumentaiDatasetSchemaSpec) *pb.DatasetSchema {
	if in == nil {
		return nil
	}
	out := &pb.DatasetSchema{}
	// MISSING: Name
	// MISSING: DocumentSchema
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func SummaryOptions_FromProto(mapCtx *direct.MapContext, in *pb.SummaryOptions) *krm.SummaryOptions {
	if in == nil {
		return nil
	}
	out := &krm.SummaryOptions{}
	out.Length = direct.Enum_FromProto(mapCtx, in.GetLength())
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	return out
}
func SummaryOptions_ToProto(mapCtx *direct.MapContext, in *krm.SummaryOptions) *pb.SummaryOptions {
	if in == nil {
		return nil
	}
	out := &pb.SummaryOptions{}
	out.Length = direct.Enum_ToProto[pb.SummaryOptions_Length](mapCtx, in.Length)
	out.Format = direct.Enum_ToProto[pb.SummaryOptions_Format](mapCtx, in.Format)
	return out
}
