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

package vmwareengine

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
)
func ManagementDnsZoneBinding_FromProto(mapCtx *direct.MapContext, in *pb.ManagementDnsZoneBinding) *krm.ManagementDnsZoneBinding {
	if in == nil {
		return nil
	}
	out := &krm.ManagementDnsZoneBinding{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	out.Description = direct.LazyPtr(in.GetDescription())
	out.VpcNetwork = direct.LazyPtr(in.GetVpcNetwork())
	out.VmwareEngineNetwork = direct.LazyPtr(in.GetVmwareEngineNetwork())
	// MISSING: Uid
	return out
}
func ManagementDnsZoneBinding_ToProto(mapCtx *direct.MapContext, in *krm.ManagementDnsZoneBinding) *pb.ManagementDnsZoneBinding {
	if in == nil {
		return nil
	}
	out := &pb.ManagementDnsZoneBinding{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	out.Description = direct.ValueOf(in.Description)
	if oneof := ManagementDnsZoneBinding_VpcNetwork_ToProto(mapCtx, in.VpcNetwork); oneof != nil {
		out.BindNetwork = oneof
	}
	if oneof := ManagementDnsZoneBinding_VmwareEngineNetwork_ToProto(mapCtx, in.VmwareEngineNetwork); oneof != nil {
		out.BindNetwork = oneof
	}
	// MISSING: Uid
	return out
}
func ManagementDnsZoneBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementDnsZoneBinding) *krm.ManagementDnsZoneBindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagementDnsZoneBindingObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Description
	// MISSING: VpcNetwork
	// MISSING: VmwareEngineNetwork
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func ManagementDnsZoneBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagementDnsZoneBindingObservedState) *pb.ManagementDnsZoneBinding {
	if in == nil {
		return nil
	}
	out := &pb.ManagementDnsZoneBinding{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.ManagementDnsZoneBinding_State](mapCtx, in.State)
	// MISSING: Description
	// MISSING: VpcNetwork
	// MISSING: VmwareEngineNetwork
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
func VmwareengineManagementDnsZoneBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementDnsZoneBinding) *krm.VmwareengineManagementDnsZoneBindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineManagementDnsZoneBindingObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Description
	// MISSING: VpcNetwork
	// MISSING: VmwareEngineNetwork
	// MISSING: Uid
	return out
}
func VmwareengineManagementDnsZoneBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineManagementDnsZoneBindingObservedState) *pb.ManagementDnsZoneBinding {
	if in == nil {
		return nil
	}
	out := &pb.ManagementDnsZoneBinding{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Description
	// MISSING: VpcNetwork
	// MISSING: VmwareEngineNetwork
	// MISSING: Uid
	return out
}
func VmwareengineManagementDnsZoneBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.ManagementDnsZoneBinding) *krm.VmwareengineManagementDnsZoneBindingSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineManagementDnsZoneBindingSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Description
	// MISSING: VpcNetwork
	// MISSING: VmwareEngineNetwork
	// MISSING: Uid
	return out
}
func VmwareengineManagementDnsZoneBindingSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineManagementDnsZoneBindingSpec) *pb.ManagementDnsZoneBinding {
	if in == nil {
		return nil
	}
	out := &pb.ManagementDnsZoneBinding{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: Description
	// MISSING: VpcNetwork
	// MISSING: VmwareEngineNetwork
	// MISSING: Uid
	return out
}
