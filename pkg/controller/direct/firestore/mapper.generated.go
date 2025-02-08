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
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func FirestoreIndexObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.FirestoreIndexObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreIndexObservedState{}
	// MISSING: Name
	// MISSING: QueryScope
	// MISSING: ApiScope
	// MISSING: Fields
	// MISSING: State
	return out
}
func FirestoreIndexObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreIndexObservedState) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	// MISSING: Name
	// MISSING: QueryScope
	// MISSING: ApiScope
	// MISSING: Fields
	// MISSING: State
	return out
}
func FirestoreIndexSpec_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.FirestoreIndexSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreIndexSpec{}
	// MISSING: Name
	// MISSING: QueryScope
	// MISSING: ApiScope
	// MISSING: Fields
	// MISSING: State
	return out
}
func FirestoreIndexSpec_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreIndexSpec) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	// MISSING: Name
	// MISSING: QueryScope
	// MISSING: ApiScope
	// MISSING: Fields
	// MISSING: State
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
