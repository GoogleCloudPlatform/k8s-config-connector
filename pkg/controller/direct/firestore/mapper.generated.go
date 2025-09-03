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
