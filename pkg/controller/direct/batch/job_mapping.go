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

package batch

import (
	pb "cloud.google.com/go/batch/apiv1/batchpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/batch/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BatchJobStatus_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.BatchJobStatus {
	if in == nil {
		return nil
	}
	out := &krm.BatchJobStatus{}
	out.ObservedState = BatchJobObservedState_FromProto(mapCtx, in)
	return out
}
func BatchJobStatus_ToProto(mapCtx *direct.MapContext, in *krm.BatchJobStatus) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out = BatchJobObservedState_ToProto(mapCtx, in.ObservedState)
	return out
}
func BatchJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.BatchJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BatchJobObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.TaskGroups = direct.Slice_FromProto(mapCtx, in.TaskGroups, TaskGroupObservedState_FromProto)
	out.Status = JobStatus_FromProto(mapCtx, in.GetStatus())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func BatchJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BatchJobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	out.TaskGroups = direct.Slice_ToProto(mapCtx, in.TaskGroups, TaskGroupObservedState_ToProto)
	out.Status = JobStatus_ToProto(mapCtx, in.Status)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func BatchJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.BatchJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.BatchJobSpec{}
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.TaskGroups = direct.Slice_FromProto(mapCtx, in.TaskGroups, TaskGroup_FromProto)
	out.AllocationPolicy = AllocationPolicy_FromProto(mapCtx, in.GetAllocationPolicy())
	out.Labels = in.Labels
	out.LogsPolicy = LogsPolicy_FromProto(mapCtx, in.GetLogsPolicy())
	out.Notifications = direct.Slice_FromProto(mapCtx, in.Notifications, JobNotification_FromProto)
	return out
}
func BatchJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.BatchJobSpec) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Priority = direct.ValueOf(in.Priority)
	out.TaskGroups = direct.Slice_ToProto(mapCtx, in.TaskGroups, TaskGroup_ToProto)
	out.AllocationPolicy = AllocationPolicy_ToProto(mapCtx, in.AllocationPolicy)
	out.Labels = in.Labels
	out.LogsPolicy = LogsPolicy_ToProto(mapCtx, in.LogsPolicy)
	out.Notifications = direct.Slice_ToProto(mapCtx, in.Notifications, JobNotification_ToProto)
	return out
}
func Runnable_Script_Path_ToProto(mapCtx *direct.MapContext, in *string) *pb.Runnable_Script_Path {
	if in == nil {
		return nil
	}
	out := &pb.Runnable_Script_Path{
		Path: direct.ValueOf(in),
	}
	return out
}
func Runnable_Script_Text_ToProto(mapCtx *direct.MapContext, in *string) *pb.Runnable_Script_Text {
	if in == nil {
		return nil
	}
	out := &pb.Runnable_Script_Text{
		Text: direct.ValueOf(in),
	}
	return out
}
func Volume_DeviceName_ToProto(mapCtx *direct.MapContext, in *string) *pb.Volume_DeviceName {
	if in == nil {
		return nil
	}
	out := &pb.Volume_DeviceName{
		DeviceName: direct.ValueOf(in),
	}
	return out
}
func AllocationPolicy_Disk_Image_ToProto(mapCtx *direct.MapContext, in *string) *pb.AllocationPolicy_Disk_Image {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_Disk_Image{
		Image: direct.ValueOf(in),
	}
	return out
}
func AllocationPolicy_Disk_Snapshot_ToProto(mapCtx *direct.MapContext, in *string) *pb.AllocationPolicy_Disk_Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_Disk_Snapshot{
		Snapshot: direct.ValueOf(in),
	}
	return out
}
func AllocationPolicy_AttachedDisk_ExistingDisk_ToProto(mapCtx *direct.MapContext, in *string) *pb.AllocationPolicy_AttachedDisk_ExistingDisk {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_AttachedDisk_ExistingDisk{
		ExistingDisk: direct.ValueOf(in),
	}
	return out
}
func AllocationPolicy_InstancePolicyOrTemplate_InstanceTemplate_ToProto(mapCtx *direct.MapContext, in *string) *pb.AllocationPolicy_InstancePolicyOrTemplate_Policy {
	if in == nil {
		return nil
	}
	out := &pb.AllocationPolicy_InstancePolicyOrTemplate_Policy{
		Policy: &pb.AllocationPolicy_InstancePolicy{
			MachineType: direct.ValueOf(in),
		},
	}
	return out
}
