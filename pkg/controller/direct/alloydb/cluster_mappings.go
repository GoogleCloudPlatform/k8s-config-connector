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

package alloydb

import (
	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	"google.golang.org/genproto/googleapis/type/dayofweek"
	"google.golang.org/genproto/googleapis/type/timeofday"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackupSourceObservedStateArray_FromProto(mapCtx *direct.MapContext, in *pb.BackupSource) []*krm.BackupSourceObservedState {
	backupSourceInfo := BackupSourceObservedState_FromProto(mapCtx, in)
	if backupSourceInfo == nil {
		return nil
	}
	return []*krm.BackupSourceObservedState{backupSourceInfo}
}
func ContinuousBackupInfoObservedStateArray_FromProto(mapCtx *direct.MapContext, in *pb.ContinuousBackupInfo) []*krm.ContinuousBackupInfoObservedState {
	continuousBackupInfo := ContinuousBackupInfoObservedState_FromProto(mapCtx, in)
	if continuousBackupInfo == nil {
		return nil
	}
	return []*krm.ContinuousBackupInfoObservedState{continuousBackupInfo}
}
func EncryptionInfoObservedStateArray_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionInfo) []*krm.EncryptionInfoObservedState {
	encryptionInfo := EncryptionInfoObservedState_FromProto(mapCtx, in)
	if encryptionInfo == nil {
		return nil
	}
	return []*krm.EncryptionInfoObservedState{encryptionInfo}
}
func MigrationSourceObservedStateArray_FromProto(mapCtx *direct.MapContext, in *pb.MigrationSource) []*krm.MigrationSourceObservedState {
	migrationSource := MigrationSourceObservedState_FromProto(mapCtx, in)
	if migrationSource == nil {
		return nil
	}
	return []*krm.MigrationSourceObservedState{migrationSource}
}
func AlloyDBClusterStatus_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.AlloyDBClusterStatus {
	if in == nil {
		return nil
	}
	out := &krm.AlloyDBClusterStatus{}
	out.BackupSource = BackupSourceObservedStateArray_FromProto(mapCtx, in.GetBackupSource())
	out.MigrationSource = MigrationSourceObservedStateArray_FromProto(mapCtx, in.GetMigrationSource())
	// MISSING: CloudsqlBackupRunSource
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	out.DatabaseVersion = direct.Enum_FromProto(mapCtx, in.GetDatabaseVersion())
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SslConfig
	// MISSING: EncryptionConfig
	out.EncryptionInfo = EncryptionInfoObservedStateArray_FromProto(mapCtx, in.GetEncryptionInfo())
	// MISSING: ContinuousBackupConfig
	out.ContinuousBackupInfo = ContinuousBackupInfoObservedStateArray_FromProto(mapCtx, in.GetContinuousBackupInfo())
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PscConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	return out
}
func AlloyDBClusterStatus_ToProto(mapCtx *direct.MapContext, in *krm.AlloyDBClusterStatus) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	if len(in.BackupSource) == 1 {
		if oneof := BackupSourceObservedState_ToProto(mapCtx, in.BackupSource[0]); oneof != nil {
			out.Source = &pb.Cluster_BackupSource{BackupSource: oneof}
		}
	}
	if len(in.MigrationSource) == 1 {
		if oneof := MigrationSourceObservedState_ToProto(mapCtx, in.MigrationSource[0]); oneof != nil {
			out.Source = &pb.Cluster_MigrationSource{MigrationSource: oneof}
		}
	}
	// MISSING: CloudsqlBackupRunSource
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: ClusterType
	out.DatabaseVersion = direct.Enum_ToProto[pb.DatabaseVersion](mapCtx, in.DatabaseVersion)
	// MISSING: NetworkConfig
	// MISSING: Network
	// MISSING: Etag
	// MISSING: Annotations
	// MISSING: Reconciling
	// MISSING: InitialUser
	// MISSING: AutomatedBackupPolicy
	// MISSING: SslConfig
	// MISSING: EncryptionConfig
	if len(in.EncryptionInfo) == 1 {
		out.EncryptionInfo = EncryptionInfoObservedState_ToProto(mapCtx, in.EncryptionInfo[0])
	}
	// MISSING: ContinuousBackupConfig
	if len(in.ContinuousBackupInfo) == 1 {
		out.ContinuousBackupInfo = ContinuousBackupInfoObservedState_ToProto(mapCtx, in.ContinuousBackupInfo[0])
	}
	// MISSING: SecondaryConfig
	// MISSING: PrimaryConfig
	// MISSING: SatisfiesPzs
	// MISSING: PscConfig
	// MISSING: MaintenanceUpdatePolicy
	// MISSING: MaintenanceSchedule
	// MISSING: GeminiConfig
	// MISSING: SubscriptionType
	// MISSING: TrialMetadata
	// MISSING: Tags
	return out
}
func AutomatedBackupPolicy_WeeklySchedule_FromProto(mapCtx *direct.MapContext, in *pb.AutomatedBackupPolicy_WeeklySchedule) *krm.AutomatedBackupPolicy_WeeklySchedule {
	if in == nil {
		return nil
	}
	out := &krm.AutomatedBackupPolicy_WeeklySchedule{}
	out.StartTimes = direct.Slice_FromProto(mapCtx, in.StartTimes, TimeOfDay_FromProto)
	out.DaysOfWeek = direct.EnumSlice_FromProto(mapCtx, in.DaysOfWeek)
	return out
}
func AutomatedBackupPolicy_WeeklySchedule_ToProto(mapCtx *direct.MapContext, in *krm.AutomatedBackupPolicy_WeeklySchedule) *pb.AutomatedBackupPolicy_WeeklySchedule {
	if in == nil {
		return nil
	}
	out := &pb.AutomatedBackupPolicy_WeeklySchedule{}
	out.StartTimes = direct.Slice_ToProto(mapCtx, in.StartTimes, TimeOfDay_ToProto)
	out.DaysOfWeek = direct.EnumSlice_ToProto[dayofweek.DayOfWeek](mapCtx, in.DaysOfWeek)
	return out
}
func ContinuousBackupInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ContinuousBackupInfo) *krm.ContinuousBackupInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContinuousBackupInfoObservedState{}
	out.EncryptionInfo = EncryptionInfoObservedStateArray_FromProto(mapCtx, in.GetEncryptionInfo())
	out.EnabledTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEnabledTime())
	out.Schedule = direct.EnumSlice_FromProto(mapCtx, in.Schedule)
	out.EarliestRestorableTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEarliestRestorableTime())
	return out
}
func ContinuousBackupInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContinuousBackupInfoObservedState) *pb.ContinuousBackupInfo {
	if in == nil {
		return nil
	}
	out := &pb.ContinuousBackupInfo{}
	if len(in.EncryptionInfo) == 1 {
		out.EncryptionInfo = EncryptionInfoObservedState_ToProto(mapCtx, in.EncryptionInfo[0])
	}
	out.EnabledTime = direct.StringTimestamp_ToProto(mapCtx, in.EnabledTime)
	out.Schedule = direct.EnumSlice_ToProto[dayofweek.DayOfWeek](mapCtx, in.Schedule)
	out.EarliestRestorableTime = direct.StringTimestamp_ToProto(mapCtx, in.EarliestRestorableTime)
	return out
}
func TimeOfDay_FromProto(mapCtx *direct.MapContext, in *timeofday.TimeOfDay) *krm.TimeOfDay {
	if in == nil {
		return nil
	}
	out := &krm.TimeOfDay{}
	out.Hours = direct.PtrTo(in.GetHours())
	out.Minutes = direct.PtrTo(in.GetMinutes())
	out.Seconds = direct.PtrTo(in.GetSeconds())
	out.Nanos = direct.PtrTo(in.GetNanos())
	return out
}
func TimeOfDay_ToProto(mapCtx *direct.MapContext, in *krm.TimeOfDay) *timeofday.TimeOfDay {
	if in == nil {
		return nil
	}
	out := &timeofday.TimeOfDay{}
	out.Hours = direct.ValueOf(in.Hours)
	out.Minutes = direct.ValueOf(in.Minutes)
	out.Seconds = direct.ValueOf(in.Seconds)
	out.Nanos = direct.ValueOf(in.Nanos)
	return out
}
func MaintenanceUpdatePolicy_MaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceUpdatePolicy_MaintenanceWindow) *krm.MaintenanceUpdatePolicy_MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krm.MaintenanceUpdatePolicy_MaintenanceWindow{}
	out.Day = direct.Enum_FromProto(mapCtx, in.GetDay())
	out.StartTime = TimeOfDay_FromProto(mapCtx, in.GetStartTime())
	return out
}
func MaintenanceUpdatePolicy_MaintenanceWindow_ToProto(mapCtx *direct.MapContext, in *krm.MaintenanceUpdatePolicy_MaintenanceWindow) *pb.MaintenanceUpdatePolicy_MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.MaintenanceUpdatePolicy_MaintenanceWindow{}
	out.Day = direct.Enum_ToProto[dayofweek.DayOfWeek](mapCtx, in.Day)
	out.StartTime = TimeOfDay_ToProto(mapCtx, in.StartTime)
	return out
}

func UserPassword_FromProto(mapCtx *direct.MapContext, in *pb.UserPassword) *krm.UserPassword {
	if in == nil {
		return nil
	}
	out := &krm.UserPassword{}
	out.User = direct.LazyPtr(in.GetUser())
	// Password is unreadable.
	// MISSING: Password
	return out
}
func UserPassword_ToProto(mapCtx *direct.MapContext, in *krm.UserPassword) *pb.UserPassword {
	if in == nil {
		return nil
	}
	out := &pb.UserPassword{}
	out.User = direct.ValueOf(in.User)
	out.Password = UserPassword_Password_ToProto(mapCtx, in.Password)
	return out
}

func UserPassword_Password_ToProto(mapCtx *direct.MapContext, in *refsv1beta1secret.Legacy) string {
	if in == nil {
		return ""
	}
	return direct.ValueOf(in.Value)
}

func Cluster_SecondaryConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_SecondaryConfig) *krm.Cluster_SecondaryConfig {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_SecondaryConfig{}
	if in.GetPrimaryClusterName() != "" {
		out.PrimaryClusterNameRef = &krm.ClusterRef{External: in.GetPrimaryClusterName()}
	}
	return out
}
func Cluster_SecondaryConfig_ToProto(mapCtx *direct.MapContext, in *krm.Cluster_SecondaryConfig) *pb.Cluster_SecondaryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_SecondaryConfig{}
	if in.PrimaryClusterNameRef != nil {
		out.PrimaryClusterName = in.PrimaryClusterNameRef.External
	}
	return out
}
