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

package appengine

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
)
func AppengineInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.AppengineInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AppengineInstanceObservedState{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: AppEngineRelease
	// MISSING: Availability
	// MISSING: VmName
	// MISSING: VmZoneName
	// MISSING: VmID
	// MISSING: StartTime
	// MISSING: Requests
	// MISSING: Errors
	// MISSING: Qps
	// MISSING: AverageLatency
	// MISSING: MemoryUsage
	// MISSING: VmStatus
	// MISSING: VmDebugEnabled
	// MISSING: VmIP
	// MISSING: VmLiveness
	return out
}
func AppengineInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AppengineInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: AppEngineRelease
	// MISSING: Availability
	// MISSING: VmName
	// MISSING: VmZoneName
	// MISSING: VmID
	// MISSING: StartTime
	// MISSING: Requests
	// MISSING: Errors
	// MISSING: Qps
	// MISSING: AverageLatency
	// MISSING: MemoryUsage
	// MISSING: VmStatus
	// MISSING: VmDebugEnabled
	// MISSING: VmIP
	// MISSING: VmLiveness
	return out
}
func AppengineInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.AppengineInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.AppengineInstanceSpec{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: AppEngineRelease
	// MISSING: Availability
	// MISSING: VmName
	// MISSING: VmZoneName
	// MISSING: VmID
	// MISSING: StartTime
	// MISSING: Requests
	// MISSING: Errors
	// MISSING: Qps
	// MISSING: AverageLatency
	// MISSING: MemoryUsage
	// MISSING: VmStatus
	// MISSING: VmDebugEnabled
	// MISSING: VmIP
	// MISSING: VmLiveness
	return out
}
func AppengineInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.AppengineInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: AppEngineRelease
	// MISSING: Availability
	// MISSING: VmName
	// MISSING: VmZoneName
	// MISSING: VmID
	// MISSING: StartTime
	// MISSING: Requests
	// MISSING: Errors
	// MISSING: Qps
	// MISSING: AverageLatency
	// MISSING: MemoryUsage
	// MISSING: VmStatus
	// MISSING: VmDebugEnabled
	// MISSING: VmIP
	// MISSING: VmLiveness
	return out
}
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: AppEngineRelease
	// MISSING: Availability
	// MISSING: VmName
	// MISSING: VmZoneName
	// MISSING: VmID
	// MISSING: StartTime
	// MISSING: Requests
	// MISSING: Errors
	// MISSING: Qps
	// MISSING: AverageLatency
	// MISSING: MemoryUsage
	// MISSING: VmStatus
	// MISSING: VmDebugEnabled
	// MISSING: VmIP
	// MISSING: VmLiveness
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: ID
	// MISSING: AppEngineRelease
	// MISSING: Availability
	// MISSING: VmName
	// MISSING: VmZoneName
	// MISSING: VmID
	// MISSING: StartTime
	// MISSING: Requests
	// MISSING: Errors
	// MISSING: Qps
	// MISSING: AverageLatency
	// MISSING: MemoryUsage
	// MISSING: VmStatus
	// MISSING: VmDebugEnabled
	// MISSING: VmIP
	// MISSING: VmLiveness
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ID = direct.LazyPtr(in.GetId())
	out.AppEngineRelease = direct.LazyPtr(in.GetAppEngineRelease())
	out.Availability = direct.Enum_FromProto(mapCtx, in.GetAvailability())
	out.VmName = direct.LazyPtr(in.GetVmName())
	out.VmZoneName = direct.LazyPtr(in.GetVmZoneName())
	out.VmID = direct.LazyPtr(in.GetVmId())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.Requests = direct.LazyPtr(in.GetRequests())
	out.Errors = direct.LazyPtr(in.GetErrors())
	out.Qps = direct.LazyPtr(in.GetQps())
	out.AverageLatency = direct.LazyPtr(in.GetAverageLatency())
	out.MemoryUsage = direct.LazyPtr(in.GetMemoryUsage())
	out.VmStatus = direct.LazyPtr(in.GetVmStatus())
	out.VmDebugEnabled = direct.LazyPtr(in.GetVmDebugEnabled())
	out.VmIP = direct.LazyPtr(in.GetVmIp())
	out.VmLiveness = direct.Enum_FromProto(mapCtx, in.GetVmLiveness())
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	out.Id = direct.ValueOf(in.ID)
	out.AppEngineRelease = direct.ValueOf(in.AppEngineRelease)
	out.Availability = direct.Enum_ToProto[pb.Instance_Availability](mapCtx, in.Availability)
	out.VmName = direct.ValueOf(in.VmName)
	out.VmZoneName = direct.ValueOf(in.VmZoneName)
	out.VmId = direct.ValueOf(in.VmID)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.Requests = direct.ValueOf(in.Requests)
	out.Errors = direct.ValueOf(in.Errors)
	out.Qps = direct.ValueOf(in.Qps)
	out.AverageLatency = direct.ValueOf(in.AverageLatency)
	out.MemoryUsage = direct.ValueOf(in.MemoryUsage)
	out.VmStatus = direct.ValueOf(in.VmStatus)
	out.VmDebugEnabled = direct.ValueOf(in.VmDebugEnabled)
	out.VmIp = direct.ValueOf(in.VmIP)
	out.VmLiveness = direct.Enum_ToProto[pb.Instance_Liveness_LivenessState](mapCtx, in.VmLiveness)
	return out
}
