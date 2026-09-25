// Copyright 2026 Google LLC
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

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackupDRBackupSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupDRBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupDRBackupSpec{}
	out.Labels = in.Labels
	out.EnforcedRetentionEndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEnforcedRetentionEndTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.BackupApplianceLocks = direct.Slice_FromProto(mapCtx, in.GetBackupApplianceLocks(), BackupLock_v1alpha1_FromProto)
	return out
}

func BackupDRBackupSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BackupDRBackupSpec) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Labels = in.Labels
	out.EnforcedRetentionEndTime = direct.StringTimestamp_ToProto(mapCtx, in.EnforcedRetentionEndTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.BackupApplianceLocks = direct.Slice_ToProto(mapCtx, in.BackupApplianceLocks, BackupLock_v1alpha1_ToProto)
	return out
}

func BackupLock_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.BackupLock) *krm.BackupLock {
	if in == nil {
		return nil
	}
	out := &krm.BackupLock{}
	out.LockUntilTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLockUntilTime())
	out.BackupApplianceLockInfo = BackupApplianceLockInfo_v1alpha1_FromProto(mapCtx, in.GetBackupApplianceLockInfo())
	return out
}

func BackupLock_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BackupLock) *pb.BackupLock {
	if in == nil {
		return nil
	}
	out := &pb.BackupLock{}
	out.LockUntilTime = direct.StringTimestamp_ToProto(mapCtx, in.LockUntilTime)
	if in.BackupApplianceLockInfo != nil {
		out.ClientLockInfo = &pb.BackupLock_BackupApplianceLockInfo{
			BackupApplianceLockInfo: BackupApplianceLockInfo_v1alpha1_ToProto(mapCtx, in.BackupApplianceLockInfo),
		}
	}
	return out
}

func BackupLockObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.BackupLock) *krm.BackupLockObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupLockObservedState{}
	out.ServiceLockInfo = ServiceLockInfoObservedState_v1alpha1_FromProto(mapCtx, in.GetServiceLockInfo())
	return out
}

func BackupLockObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BackupLockObservedState) *pb.BackupLock {
	if in == nil {
		return nil
	}
	out := &pb.BackupLock{}
	if in.ServiceLockInfo != nil {
		out.ClientLockInfo = &pb.BackupLock_ServiceLockInfo{
			ServiceLockInfo: ServiceLockInfoObservedState_v1alpha1_ToProto(mapCtx, in.ServiceLockInfo),
		}
	}
	return out
}

func ServiceLockInfoObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ServiceLockInfo) *krm.ServiceLockInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServiceLockInfoObservedState{}
	out.Operation = direct.LazyPtr(in.GetOperation())
	return out
}

func ServiceLockInfoObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ServiceLockInfoObservedState) *pb.ServiceLockInfo {
	if in == nil {
		return nil
	}
	out := &pb.ServiceLockInfo{}
	out.Operation = direct.ValueOf(in.Operation)
	return out
}

func BackupApplianceLockInfo_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.BackupApplianceLockInfo) *krm.BackupApplianceLockInfo {
	if in == nil {
		return nil
	}
	out := &krm.BackupApplianceLockInfo{}
	out.BackupApplianceID = direct.LazyPtr(in.GetBackupApplianceId())
	out.BackupApplianceName = direct.LazyPtr(in.GetBackupApplianceName())
	out.LockReason = direct.LazyPtr(in.GetLockReason())
	out.JobName = direct.LazyPtr(in.GetJobName())
	out.BackupImage = direct.LazyPtr(in.GetBackupImage())
	out.SlaID = direct.LazyPtr(in.GetSlaId())
	return out
}

func BackupApplianceLockInfo_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BackupApplianceLockInfo) *pb.BackupApplianceLockInfo {
	if in == nil {
		return nil
	}
	out := &pb.BackupApplianceLockInfo{}
	out.BackupApplianceId = direct.ValueOf(in.BackupApplianceID)
	out.BackupApplianceName = direct.ValueOf(in.BackupApplianceName)
	out.LockReason = direct.ValueOf(in.LockReason)
	if in.JobName != nil {
		out.LockSource = &pb.BackupApplianceLockInfo_JobName{
			JobName: direct.ValueOf(in.JobName),
		}
	}
	if in.BackupImage != nil {
		out.LockSource = &pb.BackupApplianceLockInfo_BackupImage{
			BackupImage: direct.ValueOf(in.BackupImage),
		}
	}
	if in.SlaID != nil {
		out.LockSource = &pb.BackupApplianceLockInfo_SlaId{
			SlaId: direct.ValueOf(in.SlaID),
		}
	}
	return out
}

func Backup_GcpBackupPlanInfo_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Backup_GCPBackupPlanInfo) *krm.Backup_GcpBackupPlanInfo {
	if in == nil {
		return nil
	}
	out := &krm.Backup_GcpBackupPlanInfo{}
	out.BackupPlan = direct.LazyPtr(in.GetBackupPlan())
	out.BackupPlanRuleID = direct.LazyPtr(in.GetBackupPlanRuleId())
	out.BackupPlanRevisionName = direct.LazyPtr(in.GetBackupPlanRevisionName())
	out.BackupPlanRevisionID = direct.LazyPtr(in.GetBackupPlanRevisionId())
	return out
}

func Backup_GcpBackupPlanInfo_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Backup_GcpBackupPlanInfo) *pb.Backup_GCPBackupPlanInfo {
	if in == nil {
		return nil
	}
	out := &pb.Backup_GCPBackupPlanInfo{}
	out.BackupPlan = direct.ValueOf(in.BackupPlan)
	out.BackupPlanRuleId = direct.ValueOf(in.BackupPlanRuleID)
	out.BackupPlanRevisionName = direct.ValueOf(in.BackupPlanRevisionName)
	out.BackupPlanRevisionId = direct.ValueOf(in.BackupPlanRevisionID)
	return out
}

func BackupDRBackupObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Backup) *krm.BackupDRBackupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupDRBackupObservedState{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ConsistencyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetConsistencyTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ServiceLocks = direct.Slice_FromProto(mapCtx, in.GetServiceLocks(), BackupLockObservedState_v1alpha1_FromProto)
	out.BackupApplianceLocks = direct.Slice_FromProto(mapCtx, in.GetBackupApplianceLocks(), BackupLockObservedState_v1alpha1_FromProto)
	out.BackupType = direct.Enum_FromProto(mapCtx, in.GetBackupType())
	out.GcpBackupPlanInfo = Backup_GcpBackupPlanInfo_v1alpha1_FromProto(mapCtx, in.GetGcpBackupPlanInfo())
	out.ResourceSizeBytes = direct.LazyPtr(in.GetResourceSizeBytes())
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}

func BackupDRBackupObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BackupDRBackupObservedState) *pb.Backup {
	if in == nil {
		return nil
	}
	out := &pb.Backup{}
	out.Description = in.Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ConsistencyTime = direct.StringTimestamp_ToProto(mapCtx, in.ConsistencyTime)
	out.State = direct.Enum_ToProto[pb.Backup_State](mapCtx, in.State)
	out.ServiceLocks = direct.Slice_ToProto(mapCtx, in.ServiceLocks, BackupLockObservedState_v1alpha1_ToProto)
	out.BackupApplianceLocks = direct.Slice_ToProto(mapCtx, in.BackupApplianceLocks, BackupLockObservedState_v1alpha1_ToProto)
	out.BackupType = direct.Enum_ToProto[pb.Backup_BackupType](mapCtx, in.BackupType)
	if in.GcpBackupPlanInfo != nil {
		out.PlanInfo = &pb.Backup_GcpBackupPlanInfo{
			GcpBackupPlanInfo: Backup_GcpBackupPlanInfo_v1alpha1_ToProto(mapCtx, in.GcpBackupPlanInfo),
		}
	}
	out.ResourceSizeBytes = direct.ValueOf(in.ResourceSizeBytes)
	out.SatisfiesPzs = in.SatisfiesPzs
	out.SatisfiesPzi = in.SatisfiesPzi
	out.Etag = in.Etag
	return out
}
