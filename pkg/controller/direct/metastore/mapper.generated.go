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
// krm.group: metastore.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.metastore.v1

package metastore

import (
	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/metastore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func MetastoreBackupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.MetastoreBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreBackupObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ServiceRevision = MetastoreServiceSpec_FromProto(mapCtx, in.GetServiceRevision())
	out.RestoringServices = in.RestoringServices
	return out
}
func MetastoreBackupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.ServiceRevision = MetastoreServiceSpec_ToProto(mapCtx, in.ServiceRevision)
	out.RestoringServices = in.RestoringServices
	return out
}
func MetastoreBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.MetastoreBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.MetastoreBackupSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func MetastoreBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.MetastoreBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	return out
}

func MetadataExportObservedState_DestinationGcsUri_ToProto(mapCtx *direct.MapContext, in *string) *pb.MetadataExport_DestinationGcsUri {
	if in == nil {
		return nil
	}
	return &pb.MetadataExport_DestinationGcsUri{DestinationGcsUri: *in}
}

func MetadataImport_DatabaseDump_FromProto(mapCtx *direct.MapContext, in *pb.MetadataImport_DatabaseDump) *krmv1alpha1.MetadataImport_DatabaseDump {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MetadataImport_DatabaseDump{}
	out.DatabaseType = direct.Enum_FromProto(mapCtx, in.GetDatabaseType())
	out.GCSURI = direct.LazyPtr(in.GetGcsUri())
	out.SourceDatabase = direct.LazyPtr(in.GetSourceDatabase())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}

func MetadataImport_DatabaseDump_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MetadataImport_DatabaseDump) *pb.MetadataImport_DatabaseDump {
	if in == nil {
		return nil
	}
	out := &pb.MetadataImport_DatabaseDump{}
	out.DatabaseType = direct.Enum_ToProto[pb.MetadataImport_DatabaseDump_DatabaseType](mapCtx, in.DatabaseType)
	out.GcsUri = direct.ValueOf(in.GCSURI)
	out.SourceDatabase = direct.ValueOf(in.SourceDatabase)
	out.Type = direct.Enum_ToProto[pb.DatabaseDumpSpec_Type](mapCtx, in.Type)
	return out
}

func MetastoreMetadataImportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataImport) *krmv1alpha1.MetastoreMetadataImportObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MetastoreMetadataImportObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}

func MetastoreMetadataImportObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MetastoreMetadataImportObservedState) *pb.MetadataImport {
	if in == nil {
		return nil
	}
	out := &pb.MetadataImport{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.MetadataImport_State](mapCtx, in.State)
	return out
}

func MetastoreMetadataImportSpec_FromProto(mapCtx *direct.MapContext, in *pb.MetadataImport) *krmv1alpha1.MetastoreMetadataImportSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.MetastoreMetadataImportSpec{}
	out.DatabaseDump = MetadataImport_DatabaseDump_FromProto(mapCtx, in.GetDatabaseDump())
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}

func MetastoreMetadataImportSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.MetastoreMetadataImportSpec) *pb.MetadataImport {
	if in == nil {
		return nil
	}
	out := &pb.MetadataImport{}
	if oneof := MetadataImport_DatabaseDump_ToProto(mapCtx, in.DatabaseDump); oneof != nil {
		out.Metadata = &pb.MetadataImport_DatabaseDump_{DatabaseDump: oneof}
	}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	return out
}

func ScalingConfig_InstanceSize_ToProto(mapCtx *direct.MapContext, in *string) *pb.ScalingConfig_InstanceSize_ {
	if in == nil {
		return nil
	}
	return &pb.ScalingConfig_InstanceSize_{InstanceSize: direct.Enum_ToProto[pb.ScalingConfig_InstanceSize](mapCtx, in)}
}

func ScalingConfig_ScalingFactor_ToProto(mapCtx *direct.MapContext, in *float32) *pb.ScalingConfig_ScalingFactor {
	if in == nil {
		return nil
	}
	return &pb.ScalingConfig_ScalingFactor{ScalingFactor: *in}
}
