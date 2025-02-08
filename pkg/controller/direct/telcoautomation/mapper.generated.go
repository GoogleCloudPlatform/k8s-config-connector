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

package telcoautomation

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/telcoautomation/apiv1/telcoautomationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/telcoautomation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func EdgeSlm_FromProto(mapCtx *direct.MapContext, in *pb.EdgeSlm) *krm.EdgeSlm {
	if in == nil {
		return nil
	}
	out := &krm.EdgeSlm{}
	out.Name = direct.LazyPtr(in.GetName())
	out.OrchestrationCluster = direct.LazyPtr(in.GetOrchestrationCluster())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: TnaVersion
	// MISSING: State
	out.WorkloadClusterType = direct.Enum_FromProto(mapCtx, in.GetWorkloadClusterType())
	return out
}
func EdgeSlm_ToProto(mapCtx *direct.MapContext, in *krm.EdgeSlm) *pb.EdgeSlm {
	if in == nil {
		return nil
	}
	out := &pb.EdgeSlm{}
	out.Name = direct.ValueOf(in.Name)
	out.OrchestrationCluster = direct.ValueOf(in.OrchestrationCluster)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	// MISSING: TnaVersion
	// MISSING: State
	out.WorkloadClusterType = direct.Enum_ToProto[pb.EdgeSlm_WorkloadClusterType](mapCtx, in.WorkloadClusterType)
	return out
}
func EdgeSlmObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EdgeSlm) *krm.EdgeSlmObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgeSlmObservedState{}
	// MISSING: Name
	// MISSING: OrchestrationCluster
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	out.TnaVersion = direct.LazyPtr(in.GetTnaVersion())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: WorkloadClusterType
	return out
}
func EdgeSlmObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgeSlmObservedState) *pb.EdgeSlm {
	if in == nil {
		return nil
	}
	out := &pb.EdgeSlm{}
	// MISSING: Name
	// MISSING: OrchestrationCluster
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	out.TnaVersion = direct.ValueOf(in.TnaVersion)
	out.State = direct.Enum_ToProto[pb.EdgeSlm_State](mapCtx, in.State)
	// MISSING: WorkloadClusterType
	return out
}
func TelcoautomationEdgeSlmObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EdgeSlm) *krm.TelcoautomationEdgeSlmObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TelcoautomationEdgeSlmObservedState{}
	// MISSING: Name
	// MISSING: OrchestrationCluster
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: TnaVersion
	// MISSING: State
	// MISSING: WorkloadClusterType
	return out
}
func TelcoautomationEdgeSlmObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TelcoautomationEdgeSlmObservedState) *pb.EdgeSlm {
	if in == nil {
		return nil
	}
	out := &pb.EdgeSlm{}
	// MISSING: Name
	// MISSING: OrchestrationCluster
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: TnaVersion
	// MISSING: State
	// MISSING: WorkloadClusterType
	return out
}
func TelcoautomationEdgeSlmSpec_FromProto(mapCtx *direct.MapContext, in *pb.EdgeSlm) *krm.TelcoautomationEdgeSlmSpec {
	if in == nil {
		return nil
	}
	out := &krm.TelcoautomationEdgeSlmSpec{}
	// MISSING: Name
	// MISSING: OrchestrationCluster
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: TnaVersion
	// MISSING: State
	// MISSING: WorkloadClusterType
	return out
}
func TelcoautomationEdgeSlmSpec_ToProto(mapCtx *direct.MapContext, in *krm.TelcoautomationEdgeSlmSpec) *pb.EdgeSlm {
	if in == nil {
		return nil
	}
	out := &pb.EdgeSlm{}
	// MISSING: Name
	// MISSING: OrchestrationCluster
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: TnaVersion
	// MISSING: State
	// MISSING: WorkloadClusterType
	return out
}
