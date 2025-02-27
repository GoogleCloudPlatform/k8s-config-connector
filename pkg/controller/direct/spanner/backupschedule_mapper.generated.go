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

package spanner

import (
	pb "cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackupScheduleSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupScheduleSpec) *krm.BackupScheduleSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupScheduleSpec{}
	out.CronSpec = CrontabSpec_FromProto(mapCtx, in.GetCronSpec())
	return out
}
func BackupScheduleSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupScheduleSpec) *pb.BackupScheduleSpec {
	if in == nil {
		return nil
	}
	out := &pb.BackupScheduleSpec{}
	if oneof := CrontabSpec_ToProto(mapCtx, in.CronSpec); oneof != nil {
		out.ScheduleSpec = &pb.BackupScheduleSpec_CronSpec{CronSpec: oneof}
	}
	return out
}
func BackupScheduleSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupScheduleSpec) *krm.BackupScheduleSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupScheduleSpecObservedState{}
	out.CronSpec = CrontabSpecObservedState_FromProto(mapCtx, in.GetCronSpec())
	return out
}
func BackupScheduleSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupScheduleSpecObservedState) *pb.BackupScheduleSpec {
	if in == nil {
		return nil
	}
	out := &pb.BackupScheduleSpec{}
	if oneof := CrontabSpecObservedState_ToProto(mapCtx, in.CronSpec); oneof != nil {
		out.ScheduleSpec = &pb.BackupScheduleSpec_CronSpec{CronSpec: oneof}
	}
	return out
}
func FullBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.FullBackupSpec) *krm.FullBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.FullBackupSpec{}
	return out
}
func FullBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.FullBackupSpec) *pb.FullBackupSpec {
	if in == nil {
		return nil
	}
	out := &pb.FullBackupSpec{}
	return out
}
func IncrementalBackupSpec_FromProto(mapCtx *direct.MapContext, in *pb.IncrementalBackupSpec) *krm.IncrementalBackupSpec {
	if in == nil {
		return nil
	}
	out := &krm.IncrementalBackupSpec{}
	return out
}
func IncrementalBackupSpec_ToProto(mapCtx *direct.MapContext, in *krm.IncrementalBackupSpec) *pb.IncrementalBackupSpec {
	if in == nil {
		return nil
	}
	out := &pb.IncrementalBackupSpec{}
	return out
}
