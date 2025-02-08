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

package edgenetwork

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/edgenetwork/apiv1/edgenetworkpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgenetwork/v1alpha1"
)
func EdgenetworkInterconnectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Interconnect) *krm.EdgenetworkInterconnectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgenetworkInterconnectObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: InterconnectType
	// MISSING: Uuid
	// MISSING: DeviceCloudResourceName
	// MISSING: PhysicalPorts
	return out
}
func EdgenetworkInterconnectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgenetworkInterconnectObservedState) *pb.Interconnect {
	if in == nil {
		return nil
	}
	out := &pb.Interconnect{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: InterconnectType
	// MISSING: Uuid
	// MISSING: DeviceCloudResourceName
	// MISSING: PhysicalPorts
	return out
}
func EdgenetworkInterconnectSpec_FromProto(mapCtx *direct.MapContext, in *pb.Interconnect) *krm.EdgenetworkInterconnectSpec {
	if in == nil {
		return nil
	}
	out := &krm.EdgenetworkInterconnectSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: InterconnectType
	// MISSING: Uuid
	// MISSING: DeviceCloudResourceName
	// MISSING: PhysicalPorts
	return out
}
func EdgenetworkInterconnectSpec_ToProto(mapCtx *direct.MapContext, in *krm.EdgenetworkInterconnectSpec) *pb.Interconnect {
	if in == nil {
		return nil
	}
	out := &pb.Interconnect{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: InterconnectType
	// MISSING: Uuid
	// MISSING: DeviceCloudResourceName
	// MISSING: PhysicalPorts
	return out
}
func Interconnect_FromProto(mapCtx *direct.MapContext, in *pb.Interconnect) *krm.Interconnect {
	if in == nil {
		return nil
	}
	out := &krm.Interconnect{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.InterconnectType = direct.Enum_FromProto(mapCtx, in.GetInterconnectType())
	// MISSING: Uuid
	// MISSING: DeviceCloudResourceName
	// MISSING: PhysicalPorts
	return out
}
func Interconnect_ToProto(mapCtx *direct.MapContext, in *krm.Interconnect) *pb.Interconnect {
	if in == nil {
		return nil
	}
	out := &pb.Interconnect{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.InterconnectType = direct.Enum_ToProto[pb.Interconnect_InterconnectType](mapCtx, in.InterconnectType)
	// MISSING: Uuid
	// MISSING: DeviceCloudResourceName
	// MISSING: PhysicalPorts
	return out
}
func InterconnectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Interconnect) *krm.InterconnectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InterconnectObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: InterconnectType
	out.Uuid = direct.LazyPtr(in.GetUuid())
	out.DeviceCloudResourceName = direct.LazyPtr(in.GetDeviceCloudResourceName())
	out.PhysicalPorts = in.PhysicalPorts
	return out
}
func InterconnectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InterconnectObservedState) *pb.Interconnect {
	if in == nil {
		return nil
	}
	out := &pb.Interconnect{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: InterconnectType
	out.Uuid = direct.ValueOf(in.Uuid)
	out.DeviceCloudResourceName = direct.ValueOf(in.DeviceCloudResourceName)
	out.PhysicalPorts = in.PhysicalPorts
	return out
}
