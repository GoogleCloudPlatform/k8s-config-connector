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

package metastore

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1alpha1"
)
func MetadataImport_FromProto(mapCtx *direct.MapContext, in *pb.MetadataImport) *krm.MetadataImport {
	if in == nil {
		return nil
	}
	out := &krm.MetadataImport{}
	out.DatabaseDump = MetadataImport_DatabaseDump_FromProto(mapCtx, in.GetDatabaseDump())
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EndTime
	// MISSING: State
	return out
}
func MetadataImport_ToProto(mapCtx *direct.MapContext, in *krm.MetadataImport) *pb.MetadataImport {
	if in == nil {
		return nil
	}
	out := &pb.MetadataImport{}
	if oneof := MetadataImport_DatabaseDump_ToProto(mapCtx, in.DatabaseDump); oneof != nil {
		out.Metadata = &pb.MetadataImport_DatabaseDump_{DatabaseDump: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EndTime
	// MISSING: State
	return out
}
func MetadataImportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataImport) *krm.MetadataImportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetadataImportObservedState{}
	// MISSING: DatabaseDump
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func MetadataImportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetadataImportObservedState) *pb.MetadataImport {
	if in == nil {
		return nil
	}
	out := &pb.MetadataImport{}
	// MISSING: DatabaseDump
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.MetadataImport_State](mapCtx, in.State)
	return out
}
func MetadataImport_DatabaseDump_FromProto(mapCtx *direct.MapContext, in *pb.MetadataImport_DatabaseDump) *krm.MetadataImport_DatabaseDump {
	if in == nil {
		return nil
	}
	out := &krm.MetadataImport_DatabaseDump{}
	out.DatabaseType = direct.Enum_FromProto(mapCtx, in.GetDatabaseType())
	out.GcsURI = direct.LazyPtr(in.GetGcsUri())
	out.SourceDatabase = direct.LazyPtr(in.GetSourceDatabase())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func MetadataImport_DatabaseDump_ToProto(mapCtx *direct.MapContext, in *krm.MetadataImport_DatabaseDump) *pb.MetadataImport_DatabaseDump {
	if in == nil {
		return nil
	}
	out := &pb.MetadataImport_DatabaseDump{}
	out.DatabaseType = direct.Enum_ToProto[pb.MetadataImport_DatabaseDump_DatabaseType](mapCtx, in.DatabaseType)
	out.GcsUri = direct.ValueOf(in.GcsURI)
	out.SourceDatabase = direct.ValueOf(in.SourceDatabase)
	out.Type = direct.Enum_ToProto[pb.DatabaseDumpSpec_Type](mapCtx, in.Type)
	return out
}
func MetastoreMetadataImportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataImport) *krm.MetastoreMetadataImportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreMetadataImportObservedState{}
	// MISSING: DatabaseDump
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EndTime
	// MISSING: State
	return out
}
func MetastoreMetadataImportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreMetadataImportObservedState) *pb.MetadataImport {
	if in == nil {
		return nil
	}
	out := &pb.MetadataImport{}
	// MISSING: DatabaseDump
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EndTime
	// MISSING: State
	return out
}
func MetastoreMetadataImportSpec_FromProto(mapCtx *direct.MapContext, in *pb.MetadataImport) *krm.MetastoreMetadataImportSpec {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreMetadataImportSpec{}
	// MISSING: DatabaseDump
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EndTime
	// MISSING: State
	return out
}
func MetastoreMetadataImportSpec_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreMetadataImportSpec) *pb.MetadataImport {
	if in == nil {
		return nil
	}
	out := &pb.MetadataImport{}
	// MISSING: DatabaseDump
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: EndTime
	// MISSING: State
	return out
}
