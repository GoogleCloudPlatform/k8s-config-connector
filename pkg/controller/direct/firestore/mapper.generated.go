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

package firestore

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
)
func Field_FromProto(mapCtx *direct.MapContext, in *pb.Field) *krm.Field {
	if in == nil {
		return nil
	}
	out := &krm.Field{}
	out.Name = direct.LazyPtr(in.GetName())
	out.IndexConfig = Field_IndexConfig_FromProto(mapCtx, in.GetIndexConfig())
	out.TtlConfig = Field_TtlConfig_FromProto(mapCtx, in.GetTtlConfig())
	return out
}
func Field_ToProto(mapCtx *direct.MapContext, in *krm.Field) *pb.Field {
	if in == nil {
		return nil
	}
	out := &pb.Field{}
	out.Name = direct.ValueOf(in.Name)
	out.IndexConfig = Field_IndexConfig_ToProto(mapCtx, in.IndexConfig)
	out.TtlConfig = Field_TtlConfig_ToProto(mapCtx, in.TtlConfig)
	return out
}
func FieldObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Field) *krm.FieldObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FieldObservedState{}
	// MISSING: Name
	// MISSING: IndexConfig
	out.TtlConfig = Field_TtlConfigObservedState_FromProto(mapCtx, in.GetTtlConfig())
	return out
}
func FieldObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FieldObservedState) *pb.Field {
	if in == nil {
		return nil
	}
	out := &pb.Field{}
	// MISSING: Name
	// MISSING: IndexConfig
	out.TtlConfig = Field_TtlConfigObservedState_ToProto(mapCtx, in.TtlConfig)
	return out
}
func Field_IndexConfig_FromProto(mapCtx *direct.MapContext, in *pb.Field_IndexConfig) *krm.Field_IndexConfig {
	if in == nil {
		return nil
	}
	out := &krm.Field_IndexConfig{}
	out.Indexes = direct.Slice_FromProto(mapCtx, in.Indexes, Index_FromProto)
	out.UsesAncestorConfig = direct.LazyPtr(in.GetUsesAncestorConfig())
	out.AncestorField = direct.LazyPtr(in.GetAncestorField())
	out.Reverting = direct.LazyPtr(in.GetReverting())
	return out
}
func Field_IndexConfig_ToProto(mapCtx *direct.MapContext, in *krm.Field_IndexConfig) *pb.Field_IndexConfig {
	if in == nil {
		return nil
	}
	out := &pb.Field_IndexConfig{}
	out.Indexes = direct.Slice_ToProto(mapCtx, in.Indexes, Index_ToProto)
	out.UsesAncestorConfig = direct.ValueOf(in.UsesAncestorConfig)
	out.AncestorField = direct.ValueOf(in.AncestorField)
	out.Reverting = direct.ValueOf(in.Reverting)
	return out
}
func Field_TtlConfig_FromProto(mapCtx *direct.MapContext, in *pb.Field_TtlConfig) *krm.Field_TtlConfig {
	if in == nil {
		return nil
	}
	out := &krm.Field_TtlConfig{}
	// MISSING: State
	return out
}
func Field_TtlConfig_ToProto(mapCtx *direct.MapContext, in *krm.Field_TtlConfig) *pb.Field_TtlConfig {
	if in == nil {
		return nil
	}
	out := &pb.Field_TtlConfig{}
	// MISSING: State
	return out
}
func Field_TtlConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Field_TtlConfig) *krm.Field_TtlConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Field_TtlConfigObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func Field_TtlConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Field_TtlConfigObservedState) *pb.Field_TtlConfig {
	if in == nil {
		return nil
	}
	out := &pb.Field_TtlConfig{}
	out.State = direct.Enum_ToProto[pb.Field_TtlConfig_State](mapCtx, in.State)
	return out
}
func FirestoreFieldObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Field) *krm.FirestoreFieldObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreFieldObservedState{}
	// MISSING: Name
	// MISSING: IndexConfig
	// MISSING: TtlConfig
	return out
}
func FirestoreFieldObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreFieldObservedState) *pb.Field {
	if in == nil {
		return nil
	}
	out := &pb.Field{}
	// MISSING: Name
	// MISSING: IndexConfig
	// MISSING: TtlConfig
	return out
}
func FirestoreFieldSpec_FromProto(mapCtx *direct.MapContext, in *pb.Field) *krm.FirestoreFieldSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreFieldSpec{}
	// MISSING: Name
	// MISSING: IndexConfig
	// MISSING: TtlConfig
	return out
}
func FirestoreFieldSpec_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreFieldSpec) *pb.Field {
	if in == nil {
		return nil
	}
	out := &pb.Field{}
	// MISSING: Name
	// MISSING: IndexConfig
	// MISSING: TtlConfig
	return out
}
func Index_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.Index {
	if in == nil {
		return nil
	}
	out := &krm.Index{}
	out.Name = direct.LazyPtr(in.GetName())
	out.QueryScope = direct.Enum_FromProto(mapCtx, in.GetQueryScope())
	out.ApiScope = direct.Enum_FromProto(mapCtx, in.GetApiScope())
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, Index_IndexField_FromProto)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func Index_ToProto(mapCtx *direct.MapContext, in *krm.Index) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	out.Name = direct.ValueOf(in.Name)
	out.QueryScope = direct.Enum_ToProto[pb.Index_QueryScope](mapCtx, in.QueryScope)
	out.ApiScope = direct.Enum_ToProto[pb.Index_ApiScope](mapCtx, in.ApiScope)
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, Index_IndexField_ToProto)
	out.State = direct.Enum_ToProto[pb.Index_State](mapCtx, in.State)
	return out
}
func Index_IndexField_FromProto(mapCtx *direct.MapContext, in *pb.Index_IndexField) *krm.Index_IndexField {
	if in == nil {
		return nil
	}
	out := &krm.Index_IndexField{}
	out.FieldPath = direct.LazyPtr(in.GetFieldPath())
	out.Order = direct.Enum_FromProto(mapCtx, in.GetOrder())
	out.ArrayConfig = direct.Enum_FromProto(mapCtx, in.GetArrayConfig())
	out.VectorConfig = Index_IndexField_VectorConfig_FromProto(mapCtx, in.GetVectorConfig())
	return out
}
func Index_IndexField_ToProto(mapCtx *direct.MapContext, in *krm.Index_IndexField) *pb.Index_IndexField {
	if in == nil {
		return nil
	}
	out := &pb.Index_IndexField{}
	out.FieldPath = direct.ValueOf(in.FieldPath)
	if oneof := Index_IndexField_Order_ToProto(mapCtx, in.Order); oneof != nil {
		out.ValueMode = oneof
	}
	if oneof := Index_IndexField_ArrayConfig_ToProto(mapCtx, in.ArrayConfig); oneof != nil {
		out.ValueMode = oneof
	}
	if oneof := Index_IndexField_VectorConfig_ToProto(mapCtx, in.VectorConfig); oneof != nil {
		out.ValueMode = &pb.Index_IndexField_VectorConfig_{VectorConfig: oneof}
	}
	return out
}
func Index_IndexField_VectorConfig_FromProto(mapCtx *direct.MapContext, in *pb.Index_IndexField_VectorConfig) *krm.Index_IndexField_VectorConfig {
	if in == nil {
		return nil
	}
	out := &krm.Index_IndexField_VectorConfig{}
	out.Dimension = direct.LazyPtr(in.GetDimension())
	out.Flat = Index_IndexField_VectorConfig_FlatIndex_FromProto(mapCtx, in.GetFlat())
	return out
}
func Index_IndexField_VectorConfig_ToProto(mapCtx *direct.MapContext, in *krm.Index_IndexField_VectorConfig) *pb.Index_IndexField_VectorConfig {
	if in == nil {
		return nil
	}
	out := &pb.Index_IndexField_VectorConfig{}
	out.Dimension = direct.ValueOf(in.Dimension)
	if oneof := Index_IndexField_VectorConfig_FlatIndex_ToProto(mapCtx, in.Flat); oneof != nil {
		out.Type = &pb.Index_IndexField_VectorConfig_Flat{Flat: oneof}
	}
	return out
}
func Index_IndexField_VectorConfig_FlatIndex_FromProto(mapCtx *direct.MapContext, in *pb.Index_IndexField_VectorConfig_FlatIndex) *krm.Index_IndexField_VectorConfig_FlatIndex {
	if in == nil {
		return nil
	}
	out := &krm.Index_IndexField_VectorConfig_FlatIndex{}
	return out
}
func Index_IndexField_VectorConfig_FlatIndex_ToProto(mapCtx *direct.MapContext, in *krm.Index_IndexField_VectorConfig_FlatIndex) *pb.Index_IndexField_VectorConfig_FlatIndex {
	if in == nil {
		return nil
	}
	out := &pb.Index_IndexField_VectorConfig_FlatIndex{}
	return out
}
