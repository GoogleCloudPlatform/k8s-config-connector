// Copyright 2024 Google LLC
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

package bigtable

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

func BigtableInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.BigtableInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableInstanceSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: State
	// MISSING: Type
	out.InstanceType = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: SatisfiesPzs
	return out
}
func BigtableInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigtableInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.DisplayName = ValueOf(in.DisplayName)
	out.Type = direct.Enum_ToProto[pb.Instance_Type](mapCtx, in.InstanceType)
	// MISSING: State
	// MISSING: Type
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: SatisfiesPzs
	return out
}

func BackupInfo_StartTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func BackupInfo_StartTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}
func BackupInfo_EndTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func BackupInfo_EndTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Backup_StartTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Backup_StartTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Backup_EndTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Backup_EndTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Backup_ExpireTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Backup_ExpireTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}
func HotTablet_StartTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func HotTablet_StartTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func HotTablet_EndTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func HotTablet_EndTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Instance_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Instance_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func OperationProgress_StartTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func OperationProgress_StartTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func OperationProgress_EndTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func OperationProgress_EndTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Snapshot_CreateTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Snapshot_CreateTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func Snapshot_DeleteTime_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	mapCtx.NotImplemented()
	return nil
}
func Snapshot_DeleteTime_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	mapCtx.NotImplemented()
	return nil
}

func AutomatedBackupPolicy_RetentionPeriod_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	mapCtx.NotImplemented()
	return nil
}
func AutomatedBackupPolicy_RetentionPeriod_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	mapCtx.NotImplemented()
	return nil
}

func AutomatedBackupPolicy_Frequency_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	mapCtx.NotImplemented()
	return nil
}
func AutomatedBackupPolicy_Frequency_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	mapCtx.NotImplemented()
	return nil
}

func ChangeStreamConfig_RetentionPeriod_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	mapCtx.NotImplemented()
	return nil
}
func ChangeStreamConfig_RetentionPeriod_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	mapCtx.NotImplemented()
	return nil
}

func GcRule_MaxAge_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	mapCtx.NotImplemented()
	return nil
}
func GcRule_MaxAge_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	mapCtx.NotImplemented()
	return nil
}

func EncryptionInfo_EncryptionStatus_FromProto(mapCtx *direct.MapContext, in *status.Status) *string {
	mapCtx.NotImplemented()
	return nil
}
func EncryptionInfo_EncryptionStatus_ToProto(mapCtx *direct.MapContext, in *string) *status.Status {
	mapCtx.NotImplemented()
	return nil
}

func GcRule_MaxNumVersions_ToProto(mapCtx *direct.MapContext, in *int32) *pb.GcRule_MaxNumVersions {
	mapCtx.NotImplemented()
	return nil
}

func Instance_SatisfiesPzs_ToProto(mapCtx *direct.MapContext, in *bool) *bool {
	mapCtx.NotImplemented()
	return nil
}

func AppProfile_DataBoostIsolationReadOnly_ComputeBillingOwner_ToProto(mapCtx *direct.MapContext, in *string) *pb.AppProfile_DataBoostIsolationReadOnly_ComputeBillingOwner {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.AppProfile_DataBoostIsolationReadOnly_ComputeBillingOwner](mapCtx, in)
	return &v
}

func AppProfile_Priority_ToProto(mapCtx *direct.MapContext, in *string) *pb.AppProfile_Priority_ {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.AppProfile_Priority](mapCtx, in)
	return &pb.AppProfile_Priority_{Priority: v}
}
