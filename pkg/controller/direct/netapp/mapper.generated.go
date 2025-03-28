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

package netapp

import (
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetAppBackupPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPolicy) *krm.NetAppBackupPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.NetAppBackupPolicySpec{}
	out.DailyBackupLimit = in.DailyBackupLimit
	out.WeeklyBackupLimit = in.WeeklyBackupLimit
	out.MonthlyBackupLimit = in.MonthlyBackupLimit
	out.Description = in.Description
	out.Enabled = in.Enabled

	// out.Labels = in.Labels // NOT YET

	return out
}
func NetAppBackupPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.NetAppBackupPolicySpec) *pb.BackupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackupPolicy{}
	out.DailyBackupLimit = in.DailyBackupLimit
	out.WeeklyBackupLimit = in.WeeklyBackupLimit
	out.MonthlyBackupLimit = in.MonthlyBackupLimit
	out.Description = in.Description
	out.Enabled = in.Enabled
	// out.Labels = in.Labels // NOT YET
	return out
}
func NetAppBackupPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupPolicy) *krm.NetAppBackupPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetAppBackupPolicyObservedState{}

	out.AssignedVolumeCount = in.AssignedVolumeCount
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func NetAppBackupPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetAppBackupPolicyObservedState) *pb.BackupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackupPolicy{}

	out.AssignedVolumeCount = in.AssignedVolumeCount
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.State = direct.Enum_ToProto[pb.BackupPolicy_State](mapCtx, in.State)
	return out
}
func BackupVault_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupVaultSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupVaultSpec{}
	// observed state: State
	// observed state: CreateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	// out.Labels = in.Labels
	return out
}
func BackupVault_ToProto(mapCtx *direct.MapContext, in *krm.BackupVaultSpec) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// observed state: State
	// observed state: CreateTime
	out.Description = direct.ValueOf(in.Description)

	return out
}
func BackupVaultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupVaultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupVaultObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Description
	// MISSING: Labels
	return out
}
func BackupVaultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupVaultObservedState) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.BackupVault_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Description
	// MISSING: Labels
	return out
}
