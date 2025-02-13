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

package cloudidentity

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudidentity/v1beta1"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/apps/cloudidentity/groups/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudIdentityGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.CloudIdentityGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityGroupObservedState{}
	out.AdditionalGroupKeys = direct.Slice_FromProto(mapCtx, in.AdditionalGroupKeys, EntityKey_FromProto)
	return out
}
func CloudIdentityGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudIdentityGroupObservedState) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	out.AdditionalGroupKeys = direct.Slice_ToProto(mapCtx, in.AdditionalGroupKeys, EntityKey_ToProto)
	return out
}
func CloudIdentityGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.CloudIdentityGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityGroupSpec{}
	out.Description = in.Description
	out.DisplayName = in.DisplayName
	out.GroupKey = EntityKey_FromProto(mapCtx, in.GetGroupKey())
	out.Labels = in.Labels
	out.Parent = in.Parent
	return out
}
func CloudIdentityGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudIdentityGroupSpec) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	out.Description = in.Description
	out.DisplayName = in.DisplayName
	out.GroupKey = EntityKey_ToProto(mapCtx, in.GroupKey)
	out.Labels = in.Labels
	out.Parent = in.Parent
	return out
}
func CloudIdentityGroupStatus_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.CloudIdentityGroupStatus {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityGroupStatus{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Name = in.Name
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ObservedState = CloudIdentityGroupObservedState_FromProto(mapCtx, in)
	return out
}
func CloudIdentityGroupStatus_ToProto(mapCtx *direct.MapContext, in *krm.CloudIdentityGroupStatus) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Name = in.Name
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	observedState := CloudIdentityGroupObservedState_ToProto(mapCtx, in.ObservedState)
	if observedState != nil {
		out.AdditionalGroupKeys = observedState.AdditionalGroupKeys
	}
	return out
}
func EntityKey_FromProto(mapCtx *direct.MapContext, in *pb.EntityKey) *krm.EntityKey {
	if in == nil {
		return nil
	}
	out := &krm.EntityKey{}
	out.ID = direct.ValueOf(in.Id)
	out.Namespace = in.Namespace
	return out
}
func EntityKey_ToProto(mapCtx *direct.MapContext, in *krm.EntityKey) *pb.EntityKey {
	if in == nil {
		return nil
	}
	out := &pb.EntityKey{}
	out.Id = direct.LazyPtr(in.ID)
	out.Namespace = in.Namespace
	return out
}
