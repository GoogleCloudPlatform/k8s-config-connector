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
// krm.group: firestore.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.firestore.admin.v1

package firestore

import (
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Database_CmekConfig_FromProto(mapCtx *direct.MapContext, in *pb.Database_CmekConfig) *krm.Database_CmekConfig {
	if in == nil {
		return nil
	}
	out := &krm.Database_CmekConfig{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	// MISSING: ActiveKeyVersion
	return out
}
func Database_CmekConfig_ToProto(mapCtx *direct.MapContext, in *krm.Database_CmekConfig) *pb.Database_CmekConfig {
	if in == nil {
		return nil
	}
	out := &pb.Database_CmekConfig{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	// MISSING: ActiveKeyVersion
	return out
}
func Database_CmekConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Database_CmekConfig) *krm.Database_CmekConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Database_CmekConfigObservedState{}
	// MISSING: KMSKeyName
	out.ActiveKeyVersion = in.ActiveKeyVersion
	return out
}
func Database_CmekConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Database_CmekConfigObservedState) *pb.Database_CmekConfig {
	if in == nil {
		return nil
	}
	out := &pb.Database_CmekConfig{}
	// MISSING: KMSKeyName
	out.ActiveKeyVersion = in.ActiveKeyVersion
	return out
}
func Database_SourceInfo_FromProto(mapCtx *direct.MapContext, in *pb.Database_SourceInfo) *krm.Database_SourceInfo {
	if in == nil {
		return nil
	}
	out := &krm.Database_SourceInfo{}
	out.Backup = Database_SourceInfo_BackupSource_FromProto(mapCtx, in.GetBackup())
	out.Operation = direct.LazyPtr(in.GetOperation())
	return out
}
func Database_SourceInfo_ToProto(mapCtx *direct.MapContext, in *krm.Database_SourceInfo) *pb.Database_SourceInfo {
	if in == nil {
		return nil
	}
	out := &pb.Database_SourceInfo{}
	if oneof := Database_SourceInfo_BackupSource_ToProto(mapCtx, in.Backup); oneof != nil {
		out.Source = &pb.Database_SourceInfo_Backup{Backup: oneof}
	}
	out.Operation = direct.ValueOf(in.Operation)
	return out
}
func Database_SourceInfo_BackupSource_FromProto(mapCtx *direct.MapContext, in *pb.Database_SourceInfo_BackupSource) *krm.Database_SourceInfo_BackupSource {
	if in == nil {
		return nil
	}
	out := &krm.Database_SourceInfo_BackupSource{}
	out.Backup = direct.LazyPtr(in.GetBackup())
	return out
}
func Database_SourceInfo_BackupSource_ToProto(mapCtx *direct.MapContext, in *krm.Database_SourceInfo_BackupSource) *pb.Database_SourceInfo_BackupSource {
	if in == nil {
		return nil
	}
	out := &pb.Database_SourceInfo_BackupSource{}
	out.Backup = direct.ValueOf(in.Backup)
	return out
}
func FirestoreIndexSpec_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.FirestoreIndexSpec {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreIndexSpec{}
	// MISSING: Name
	out.QueryScope = direct.Enum_FromProto(mapCtx, in.GetQueryScope())
	// MISSING: APIScope
	out.Fields = direct.Slice_FromProto(mapCtx, in.Fields, IndexFields_FromProto)
	// MISSING: State
	// MISSING: Density
	// MISSING: Multikey
	// MISSING: ShardCount
	return out
}
func FirestoreIndexSpec_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreIndexSpec) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	// MISSING: Name
	out.QueryScope = direct.Enum_ToProto[pb.Index_QueryScope](mapCtx, in.QueryScope)
	// MISSING: APIScope
	out.Fields = direct.Slice_ToProto(mapCtx, in.Fields, IndexFields_ToProto)
	// MISSING: State
	// MISSING: Density
	// MISSING: Multikey
	// MISSING: ShardCount
	return out
}
func FirestoreIndexStatus_FromProto(mapCtx *direct.MapContext, in *pb.Index) *krm.FirestoreIndexStatus {
	if in == nil {
		return nil
	}
	out := &krm.FirestoreIndexStatus{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: QueryScope
	// MISSING: APIScope
	// MISSING: Fields
	// MISSING: State
	// MISSING: Density
	// MISSING: Multikey
	// MISSING: ShardCount
	return out
}
func FirestoreIndexStatus_ToProto(mapCtx *direct.MapContext, in *krm.FirestoreIndexStatus) *pb.Index {
	if in == nil {
		return nil
	}
	out := &pb.Index{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: QueryScope
	// MISSING: APIScope
	// MISSING: Fields
	// MISSING: State
	// MISSING: Density
	// MISSING: Multikey
	// MISSING: ShardCount
	return out
}
func IndexFields_FromProto(mapCtx *direct.MapContext, in *pb.Index_IndexField) *krm.IndexFields {
	if in == nil {
		return nil
	}
	out := &krm.IndexFields{}
	out.FieldPath = direct.LazyPtr(in.GetFieldPath())
	out.Order = direct.Enum_FromProto(mapCtx, in.GetOrder())
	out.ArrayConfig = direct.Enum_FromProto(mapCtx, in.GetArrayConfig())
	// MISSING: VectorConfig
	return out
}
func IndexFields_ToProto(mapCtx *direct.MapContext, in *krm.IndexFields) *pb.Index_IndexField {
	if in == nil {
		return nil
	}
	out := &pb.Index_IndexField{}
	out.FieldPath = direct.ValueOf(in.FieldPath)
	if oneof := IndexFields_Order_ToProto(mapCtx, in.Order); oneof != nil {
		out.ValueMode = oneof
	}
	if oneof := IndexFields_ArrayConfig_ToProto(mapCtx, in.ArrayConfig); oneof != nil {
		out.ValueMode = oneof
	}
	// MISSING: VectorConfig
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
