// Copyright 2024 Google LLC
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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/apps/cloudidentity/groups/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudIdentityGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.CloudIdentityGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityGroupObservedState{}
	out.AdditionalGroupKeys = direct.Slice_FromProto(mapCtx, in.AdditionalGroupKeys, EntityKey_FromProto)
	// MISSING: CreateTime
	// MISSING: DynamicGroupMetadata
	// MISSING: Name
	// MISSING: PosixGroups
	// MISSING: UpdateTime
	return out
}
func CloudIdentityGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudIdentityGroupObservedState) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	out.AdditionalGroupKeys = direct.Slice_ToProto(mapCtx, in.AdditionalGroupKeys, EntityKey_ToProto)
	// MISSING: CreateTime
	// MISSING: DynamicGroupMetadata
	// MISSING: Name
	// MISSING: PosixGroups
	// MISSING: UpdateTime
	return out
}
func CloudIdentityGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.CloudIdentityGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudIdentityGroupSpec{}
	// MISSING: CreateTime
	out.Description = in.Description
	out.DisplayName = in.DisplayName
	// MISSING: DynamicGroupMetadata
	out.GroupKey = EntityKey_FromProto(mapCtx, in.GetGroupKey())
	out.Labels = in.Labels
	// MISSING: Name
	out.Parent = in.Parent
	// MISSING: PosixGroups
	// MISSING: UpdateTime
	return out
}
func CloudIdentityGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudIdentityGroupSpec) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	// MISSING: CreateTime
	out.Description = in.Description
	out.DisplayName = in.DisplayName
	// MISSING: DynamicGroupMetadata
	out.GroupKey = EntityKey_ToProto(mapCtx, in.GroupKey)
	out.Labels = in.Labels
	// MISSING: Name
	out.Parent = in.Parent
	// MISSING: PosixGroups
	// MISSING: UpdateTime
	return out
}
func DynamicGroupQuery_FromProto(mapCtx *direct.MapContext, in *pb.DynamicGroupQuery) *krm.DynamicGroupQuery {
	if in == nil {
		return nil
	}
	out := &krm.DynamicGroupQuery{}
	out.Query = in.Query
	out.ResourceType = in.ResourceType
	return out
}
func DynamicGroupQuery_ToProto(mapCtx *direct.MapContext, in *krm.DynamicGroupQuery) *pb.DynamicGroupQuery {
	if in == nil {
		return nil
	}
	out := &pb.DynamicGroupQuery{}
	out.Query = in.Query
	out.ResourceType = in.ResourceType
	return out
}
func DynamicGroupStatus_FromProto(mapCtx *direct.MapContext, in *pb.DynamicGroupStatus) *krm.DynamicGroupStatus {
	if in == nil {
		return nil
	}
	out := &krm.DynamicGroupStatus{}
	out.Status = in.Status
	out.StatusTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStatusTime())
	return out
}
func DynamicGroupStatus_ToProto(mapCtx *direct.MapContext, in *krm.DynamicGroupStatus) *pb.DynamicGroupStatus {
	if in == nil {
		return nil
	}
	out := &pb.DynamicGroupStatus{}
	out.Status = in.Status
	out.StatusTime = direct.StringTimestamp_ToProto(mapCtx, in.StatusTime)
	return out
}
func EntityKey_FromProto(mapCtx *direct.MapContext, in *pb.EntityKey) *krm.EntityKey {
	if in == nil {
		return nil
	}
	out := &krm.EntityKey{}
	out.ID = in.Id
	out.Namespace = in.Namespace
	return out
}
func EntityKey_ToProto(mapCtx *direct.MapContext, in *krm.EntityKey) *pb.EntityKey {
	if in == nil {
		return nil
	}
	out := &pb.EntityKey{}
	out.Id = in.ID
	out.Namespace = in.Namespace
	return out
}
func PosixGroup_FromProto(mapCtx *direct.MapContext, in *pb.PosixGroup) *krm.PosixGroup {
	if in == nil {
		return nil
	}
	out := &krm.PosixGroup{}
	out.Gid = in.Gid
	out.Name = in.Name
	out.SystemID = in.SystemId
	return out
}
func PosixGroup_ToProto(mapCtx *direct.MapContext, in *krm.PosixGroup) *pb.PosixGroup {
	if in == nil {
		return nil
	}
	out := &pb.PosixGroup{}
	out.Gid = in.Gid
	out.Name = in.Name
	out.SystemId = in.SystemID
	return out
}
