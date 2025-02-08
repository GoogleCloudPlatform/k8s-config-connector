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
func HcxActivationKey_FromProto(mapCtx *direct.MapContext, in *pb.HcxActivationKey) *krm.HcxActivationKey {
	if in == nil {
		return nil
	}
	out := &krm.HcxActivationKey{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: ActivationKey
	// MISSING: Uid
	return out
}
func HcxActivationKey_ToProto(mapCtx *direct.MapContext, in *krm.HcxActivationKey) *pb.HcxActivationKey {
	if in == nil {
		return nil
	}
	out := &pb.HcxActivationKey{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: ActivationKey
	// MISSING: Uid
	return out
}
func HcxActivationKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.HcxActivationKey) *krm.HcxActivationKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.HcxActivationKeyObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ActivationKey = direct.LazyPtr(in.GetActivationKey())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func HcxActivationKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.HcxActivationKeyObservedState) *pb.HcxActivationKey {
	if in == nil {
		return nil
	}
	out := &pb.HcxActivationKey{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.State = direct.Enum_ToProto[pb.HcxActivationKey_State](mapCtx, in.State)
	out.ActivationKey = direct.ValueOf(in.ActivationKey)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
func VmwareengineHcxActivationKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.HcxActivationKey) *krm.VmwareengineHcxActivationKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineHcxActivationKeyObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: ActivationKey
	// MISSING: Uid
	return out
}
func VmwareengineHcxActivationKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineHcxActivationKeyObservedState) *pb.HcxActivationKey {
	if in == nil {
		return nil
	}
	out := &pb.HcxActivationKey{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: ActivationKey
	// MISSING: Uid
	return out
}
func VmwareengineHcxActivationKeySpec_FromProto(mapCtx *direct.MapContext, in *pb.HcxActivationKey) *krm.VmwareengineHcxActivationKeySpec {
	if in == nil {
		return nil
	}
	out := &krm.VmwareengineHcxActivationKeySpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: ActivationKey
	// MISSING: Uid
	return out
}
func VmwareengineHcxActivationKeySpec_ToProto(mapCtx *direct.MapContext, in *krm.VmwareengineHcxActivationKeySpec) *pb.HcxActivationKey {
	if in == nil {
		return nil
	}
	out := &pb.HcxActivationKey{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: ActivationKey
	// MISSING: Uid
	return out
}
