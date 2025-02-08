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

package gkehub

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/gkehub/apiv1beta/gkehubpb"
)
func CommonFeatureSpec_FromProto(mapCtx *direct.MapContext, in *pb.CommonFeatureSpec) *krm.CommonFeatureSpec {
	if in == nil {
		return nil
	}
	out := &krm.CommonFeatureSpec{}
	out.Multiclusteringress = FeatureSpec_FromProto(mapCtx, in.GetMulticlusteringress())
	return out
}
func CommonFeatureSpec_ToProto(mapCtx *direct.MapContext, in *krm.CommonFeatureSpec) *pb.CommonFeatureSpec {
	if in == nil {
		return nil
	}
	out := &pb.CommonFeatureSpec{}
	if oneof := FeatureSpec_ToProto(mapCtx, in.Multiclusteringress); oneof != nil {
		out.FeatureSpec = &pb.CommonFeatureSpec_Multiclusteringress{Multiclusteringress: oneof}
	}
	return out
}
func CommonFeatureState_FromProto(mapCtx *direct.MapContext, in *pb.CommonFeatureState) *krm.CommonFeatureState {
	if in == nil {
		return nil
	}
	out := &krm.CommonFeatureState{}
	// MISSING: State
	return out
}
func CommonFeatureState_ToProto(mapCtx *direct.MapContext, in *krm.CommonFeatureState) *pb.CommonFeatureState {
	if in == nil {
		return nil
	}
	out := &pb.CommonFeatureState{}
	// MISSING: State
	return out
}
func CommonFeatureStateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CommonFeatureState) *krm.CommonFeatureStateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CommonFeatureStateObservedState{}
	out.State = FeatureState_FromProto(mapCtx, in.GetState())
	return out
}
func CommonFeatureStateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CommonFeatureStateObservedState) *pb.CommonFeatureState {
	if in == nil {
		return nil
	}
	out := &pb.CommonFeatureState{}
	out.State = FeatureState_ToProto(mapCtx, in.State)
	return out
}
func Feature_FromProto(mapCtx *direct.MapContext, in *pb.Feature) *krm.Feature {
	if in == nil {
		return nil
	}
	out := &krm.Feature{}
	// MISSING: Name
	out.Labels = in.Labels
	// MISSING: ResourceState
	out.Spec = CommonFeatureSpec_FromProto(mapCtx, in.GetSpec())
	// MISSING: MembershipSpecs
	// MISSING: State
	// MISSING: MembershipStates
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	return out
}
func Feature_ToProto(mapCtx *direct.MapContext, in *krm.Feature) *pb.Feature {
	if in == nil {
		return nil
	}
	out := &pb.Feature{}
	// MISSING: Name
	out.Labels = in.Labels
	// MISSING: ResourceState
	out.Spec = CommonFeatureSpec_ToProto(mapCtx, in.Spec)
	// MISSING: MembershipSpecs
	// MISSING: State
	// MISSING: MembershipStates
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	return out
}
func FeatureObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Feature) *krm.FeatureObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FeatureObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Labels
	out.ResourceState = FeatureResourceState_FromProto(mapCtx, in.GetResourceState())
	// MISSING: Spec
	// MISSING: MembershipSpecs
	out.State = CommonFeatureState_FromProto(mapCtx, in.GetState())
	// MISSING: MembershipStates
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	return out
}
func FeatureObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FeatureObservedState) *pb.Feature {
	if in == nil {
		return nil
	}
	out := &pb.Feature{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Labels
	out.ResourceState = FeatureResourceState_ToProto(mapCtx, in.ResourceState)
	// MISSING: Spec
	// MISSING: MembershipSpecs
	out.State = CommonFeatureState_ToProto(mapCtx, in.State)
	// MISSING: MembershipStates
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	return out
}
func FeatureResourceState_FromProto(mapCtx *direct.MapContext, in *pb.FeatureResourceState) *krm.FeatureResourceState {
	if in == nil {
		return nil
	}
	out := &krm.FeatureResourceState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func FeatureResourceState_ToProto(mapCtx *direct.MapContext, in *krm.FeatureResourceState) *pb.FeatureResourceState {
	if in == nil {
		return nil
	}
	out := &pb.FeatureResourceState{}
	out.State = direct.Enum_ToProto[pb.FeatureResourceState_State](mapCtx, in.State)
	return out
}
func FeatureState_FromProto(mapCtx *direct.MapContext, in *pb.FeatureState) *krm.FeatureState {
	if in == nil {
		return nil
	}
	out := &krm.FeatureState{}
	out.Code = direct.Enum_FromProto(mapCtx, in.GetCode())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func FeatureState_ToProto(mapCtx *direct.MapContext, in *krm.FeatureState) *pb.FeatureState {
	if in == nil {
		return nil
	}
	out := &pb.FeatureState{}
	out.Code = direct.Enum_ToProto[pb.FeatureState_Code](mapCtx, in.Code)
	out.Description = direct.ValueOf(in.Description)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func MembershipFeatureSpec_FromProto(mapCtx *direct.MapContext, in *pb.MembershipFeatureSpec) *krm.MembershipFeatureSpec {
	if in == nil {
		return nil
	}
	out := &krm.MembershipFeatureSpec{}
	out.Configmanagement = MembershipSpec_FromProto(mapCtx, in.GetConfigmanagement())
	out.Mesh = MembershipSpec_FromProto(mapCtx, in.GetMesh())
	out.Policycontroller = MembershipSpec_FromProto(mapCtx, in.GetPolicycontroller())
	return out
}
func MembershipFeatureSpec_ToProto(mapCtx *direct.MapContext, in *krm.MembershipFeatureSpec) *pb.MembershipFeatureSpec {
	if in == nil {
		return nil
	}
	out := &pb.MembershipFeatureSpec{}
	if oneof := MembershipSpec_ToProto(mapCtx, in.Configmanagement); oneof != nil {
		out.FeatureSpec = &pb.MembershipFeatureSpec_Configmanagement{Configmanagement: oneof}
	}
	if oneof := MembershipSpec_ToProto(mapCtx, in.Mesh); oneof != nil {
		out.FeatureSpec = &pb.MembershipFeatureSpec_Mesh{Mesh: oneof}
	}
	if oneof := MembershipSpec_ToProto(mapCtx, in.Policycontroller); oneof != nil {
		out.FeatureSpec = &pb.MembershipFeatureSpec_Policycontroller{Policycontroller: oneof}
	}
	return out
}
func MembershipFeatureState_FromProto(mapCtx *direct.MapContext, in *pb.MembershipFeatureState) *krm.MembershipFeatureState {
	if in == nil {
		return nil
	}
	out := &krm.MembershipFeatureState{}
	out.Servicemesh = MembershipState_FromProto(mapCtx, in.GetServicemesh())
	out.Metering = MembershipState_FromProto(mapCtx, in.GetMetering())
	out.Configmanagement = MembershipState_FromProto(mapCtx, in.GetConfigmanagement())
	out.Policycontroller = MembershipState_FromProto(mapCtx, in.GetPolicycontroller())
	out.State = FeatureState_FromProto(mapCtx, in.GetState())
	return out
}
func MembershipFeatureState_ToProto(mapCtx *direct.MapContext, in *krm.MembershipFeatureState) *pb.MembershipFeatureState {
	if in == nil {
		return nil
	}
	out := &pb.MembershipFeatureState{}
	if oneof := MembershipState_ToProto(mapCtx, in.Servicemesh); oneof != nil {
		out.FeatureState = &pb.MembershipFeatureState_Servicemesh{Servicemesh: oneof}
	}
	if oneof := MembershipState_ToProto(mapCtx, in.Metering); oneof != nil {
		out.FeatureState = &pb.MembershipFeatureState_Metering{Metering: oneof}
	}
	if oneof := MembershipState_ToProto(mapCtx, in.Configmanagement); oneof != nil {
		out.FeatureState = &pb.MembershipFeatureState_Configmanagement{Configmanagement: oneof}
	}
	if oneof := MembershipState_ToProto(mapCtx, in.Policycontroller); oneof != nil {
		out.FeatureState = &pb.MembershipFeatureState_Policycontroller{Policycontroller: oneof}
	}
	out.State = FeatureState_ToProto(mapCtx, in.State)
	return out
}
