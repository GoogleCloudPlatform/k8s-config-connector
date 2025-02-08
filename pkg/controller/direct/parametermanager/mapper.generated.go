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

package parametermanager

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/parametermanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/parametermanager/apiv1/parametermanagerpb"
)
func Parameter_FromProto(mapCtx *direct.MapContext, in *pb.Parameter) *krm.Parameter {
	if in == nil {
		return nil
	}
	out := &krm.Parameter{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	// MISSING: PolicyMember
	return out
}
func Parameter_ToProto(mapCtx *direct.MapContext, in *krm.Parameter) *pb.Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Parameter{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Format = direct.Enum_ToProto[pb.ParameterFormat](mapCtx, in.Format)
	// MISSING: PolicyMember
	return out
}
func ParameterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Parameter) *krm.ParameterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ParameterObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Format
	out.PolicyMember = ResourcePolicyMember_FromProto(mapCtx, in.GetPolicyMember())
	return out
}
func ParameterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ParameterObservedState) *pb.Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Parameter{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Format
	out.PolicyMember = ResourcePolicyMember_ToProto(mapCtx, in.PolicyMember)
	return out
}
func ParametermanagerParameterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Parameter) *krm.ParametermanagerParameterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ParametermanagerParameterObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Format
	// MISSING: PolicyMember
	return out
}
func ParametermanagerParameterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ParametermanagerParameterObservedState) *pb.Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Parameter{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Format
	// MISSING: PolicyMember
	return out
}
func ParametermanagerParameterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Parameter) *krm.ParametermanagerParameterSpec {
	if in == nil {
		return nil
	}
	out := &krm.ParametermanagerParameterSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Format
	// MISSING: PolicyMember
	return out
}
func ParametermanagerParameterSpec_ToProto(mapCtx *direct.MapContext, in *krm.ParametermanagerParameterSpec) *pb.Parameter {
	if in == nil {
		return nil
	}
	out := &pb.Parameter{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Format
	// MISSING: PolicyMember
	return out
}
