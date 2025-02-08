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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ExternalAddress_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAddress) *krm.ExternalAddress {
	if in == nil {
		return nil
	}
	out := &krm.ExternalAddress{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	// MISSING: ExternalIP
	// MISSING: State
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func ExternalAddress_ToProto(mapCtx *direct.MapContext, in *krm.ExternalAddress) *pb.ExternalAddress {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAddress{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.InternalIp = direct.ValueOf(in.InternalIP)
	// MISSING: ExternalIP
	// MISSING: State
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	return out
}
func ExternalAddressObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAddress) *krm.ExternalAddressObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ExternalAddressObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: InternalIP
	out.ExternalIP = direct.LazyPtr(in.GetExternalIp())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Description
	return out
}
func ExternalAddressObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ExternalAddressObservedState) *pb.ExternalAddress {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAddress{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: InternalIP
	out.ExternalIp = direct.ValueOf(in.ExternalIP)
	out.State = direct.Enum_ToProto[pb.ExternalAddress_State](mapCtx, in.State)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Description
	return out
}
func VmwareengineExternalAddressObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAddress) *krm.VmwareengineExternalAddressObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineExternalAddressObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: InternalIP
	// MISSING: ExternalIP
	// MISSING: State
	// MISSING: Uid
	// MISSING: Description
	return out
}
func VmwareengineExternalAddressObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineExternalAddressObservedState) *pb.ExternalAddress {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAddress{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: InternalIP
	// MISSING: ExternalIP
	// MISSING: State
	// MISSING: Uid
	// MISSING: Description
	return out
}
func VmwareengineExternalAddressSpec_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAddress) *krm.VmwareengineExternalAddressSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineExternalAddressSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: InternalIP
	// MISSING: ExternalIP
	// MISSING: State
	// MISSING: Uid
	// MISSING: Description
	return out
}
func VmwareengineExternalAddressSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineExternalAddressSpec) *pb.ExternalAddress {
	if in == nil {
		return nil
	}
	out := &pb.ExternalAddress{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: InternalIP
	// MISSING: ExternalIP
	// MISSING: State
	// MISSING: Uid
	// MISSING: Description
	return out
}
