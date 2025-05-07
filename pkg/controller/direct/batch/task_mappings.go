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
// krm.group: batch.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.batch.v1

package batch

import (
	pb "cloud.google.com/go/batch/apiv1/batchpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/batch/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BatchTaskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.BatchTaskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BatchTaskObservedState{}
	out.Status = TaskStatus_FromProto(mapCtx, in.GetStatus())
	return out
}
func BatchTaskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BatchTaskObservedState) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	out.Status = TaskStatus_ToProto(mapCtx, in.Status)
	return out
}
func BatchTaskSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.BatchTaskSpec {
	if in == nil {
		return nil
	}
	out := &krm.BatchTaskSpec{}
	return out
}
func BatchTaskSpec_ToProto(mapCtx *direct.MapContext, in *krm.BatchTaskSpec) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
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
