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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/batch/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/batch/apiv1/batchpb"
)
func BatchJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.BatchJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BatchJobObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Priority
	// MISSING: TaskGroups
	// MISSING: AllocationPolicy
	// MISSING: Labels
	// MISSING: Status
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LogsPolicy
	// MISSING: Notifications
	return out
}
func BatchJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BatchJobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Priority
	// MISSING: TaskGroups
	// MISSING: AllocationPolicy
	// MISSING: Labels
	// MISSING: Status
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LogsPolicy
	// MISSING: Notifications
	return out
}
func BatchJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.BatchJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.BatchJobSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Priority
	// MISSING: TaskGroups
	// MISSING: AllocationPolicy
	// MISSING: Labels
	// MISSING: Status
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LogsPolicy
	// MISSING: Notifications
	return out
}
func BatchJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.BatchJobSpec) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Priority
	// MISSING: TaskGroups
	// MISSING: AllocationPolicy
	// MISSING: Labels
	// MISSING: Status
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: LogsPolicy
	// MISSING: Notifications
	return out
}
func BatchTaskGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TaskGroup) *krm.BatchTaskGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BatchTaskGroupObservedState{}
	// MISSING: Name
	// MISSING: TaskSpec
	// MISSING: TaskCount
	// MISSING: Parallelism
	// MISSING: SchedulingPolicy
	// MISSING: TaskEnvironments
	// MISSING: TaskCountPerNode
	// MISSING: RequireHostsFile
	// MISSING: PermissiveSSH
	// MISSING: RunAsNonRoot
	return out
}
func BatchTaskGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BatchTaskGroupObservedState) *pb.TaskGroup {
	if in == nil {
		return nil
	}
	out := &pb.TaskGroup{}
	// MISSING: Name
	// MISSING: TaskSpec
	// MISSING: TaskCount
	// MISSING: Parallelism
	// MISSING: SchedulingPolicy
	// MISSING: TaskEnvironments
	// MISSING: TaskCountPerNode
	// MISSING: RequireHostsFile
	// MISSING: PermissiveSSH
	// MISSING: RunAsNonRoot
	return out
}
func BatchTaskGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.TaskGroup) *krm.BatchTaskGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.BatchTaskGroupSpec{}
	// MISSING: Name
	// MISSING: TaskSpec
	// MISSING: TaskCount
	// MISSING: Parallelism
	// MISSING: SchedulingPolicy
	// MISSING: TaskEnvironments
	// MISSING: TaskCountPerNode
	// MISSING: RequireHostsFile
	// MISSING: PermissiveSSH
	// MISSING: RunAsNonRoot
	return out
}
func BatchTaskGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.BatchTaskGroupSpec) *pb.TaskGroup {
	if in == nil {
		return nil
	}
	out := &pb.TaskGroup{}
	// MISSING: Name
	// MISSING: TaskSpec
	// MISSING: TaskCount
	// MISSING: Parallelism
	// MISSING: SchedulingPolicy
	// MISSING: TaskEnvironments
	// MISSING: TaskCountPerNode
	// MISSING: RequireHostsFile
	// MISSING: PermissiveSSH
	// MISSING: RunAsNonRoot
	return out
}
func BatchTaskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.BatchTaskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BatchTaskObservedState{}
	// MISSING: Name
	// MISSING: Status
	return out
}
func BatchTaskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BatchTaskObservedState) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	// MISSING: Name
	// MISSING: Status
	return out
}
func BatchTaskSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.BatchTaskSpec {
	if in == nil {
		return nil
	}
	out := &krm.BatchTaskSpec{}
	// MISSING: Name
	// MISSING: Status
	return out
}
func BatchTaskSpec_ToProto(mapCtx *direct.MapContext, in *krm.BatchTaskSpec) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	// MISSING: Name
	// MISSING: Status
	return out
}
func StatusEvent_FromProto(mapCtx *direct.MapContext, in *pb.StatusEvent) *krm.StatusEvent {
	if in == nil {
		return nil
	}
	out := &krm.StatusEvent{}
	out.Type = direct.LazyPtr(in.GetType())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.EventTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEventTime())
	out.TaskExecution = TaskExecution_FromProto(mapCtx, in.GetTaskExecution())
	out.TaskState = direct.Enum_FromProto(mapCtx, in.GetTaskState())
	return out
}
func StatusEvent_ToProto(mapCtx *direct.MapContext, in *krm.StatusEvent) *pb.StatusEvent {
	if in == nil {
		return nil
	}
	out := &pb.StatusEvent{}
	out.Type = direct.ValueOf(in.Type)
	out.Description = direct.ValueOf(in.Description)
	out.EventTime = direct.StringTimestamp_ToProto(mapCtx, in.EventTime)
	out.TaskExecution = TaskExecution_ToProto(mapCtx, in.TaskExecution)
	out.TaskState = direct.Enum_ToProto[pb.TaskStatus_State](mapCtx, in.TaskState)
	return out
}
func Task_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.Task {
	if in == nil {
		return nil
	}
	out := &krm.Task{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Status = TaskStatus_FromProto(mapCtx, in.GetStatus())
	return out
}
func Task_ToProto(mapCtx *direct.MapContext, in *krm.Task) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	out.Name = direct.ValueOf(in.Name)
	out.Status = TaskStatus_ToProto(mapCtx, in.Status)
	return out
}
func TaskExecution_FromProto(mapCtx *direct.MapContext, in *pb.TaskExecution) *krm.TaskExecution {
	if in == nil {
		return nil
	}
	out := &krm.TaskExecution{}
	out.ExitCode = direct.LazyPtr(in.GetExitCode())
	return out
}
func TaskExecution_ToProto(mapCtx *direct.MapContext, in *krm.TaskExecution) *pb.TaskExecution {
	if in == nil {
		return nil
	}
	out := &pb.TaskExecution{}
	out.ExitCode = direct.ValueOf(in.ExitCode)
	return out
}
func TaskStatus_FromProto(mapCtx *direct.MapContext, in *pb.TaskStatus) *krm.TaskStatus {
	if in == nil {
		return nil
	}
	out := &krm.TaskStatus{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StatusEvents = direct.Slice_FromProto(mapCtx, in.StatusEvents, StatusEvent_FromProto)
	return out
}
func TaskStatus_ToProto(mapCtx *direct.MapContext, in *krm.TaskStatus) *pb.TaskStatus {
	if in == nil {
		return nil
	}
	out := &pb.TaskStatus{}
	out.State = direct.Enum_ToProto[pb.TaskStatus_State](mapCtx, in.State)
	out.StatusEvents = direct.Slice_ToProto(mapCtx, in.StatusEvents, StatusEvent_ToProto)
	return out
}
