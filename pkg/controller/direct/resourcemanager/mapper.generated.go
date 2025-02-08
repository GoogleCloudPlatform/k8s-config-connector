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

package resourcemanager

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ResourcemanagerTagKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TagKey) *krm.ResourcemanagerTagKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerTagKeyObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ShortName
	// MISSING: NamespacedName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Purpose
	// MISSING: PurposeData
	return out
}
func ResourcemanagerTagKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerTagKeyObservedState) *pb.TagKey {
	if in == nil {
		return nil
	}
	out := &pb.TagKey{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ShortName
	// MISSING: NamespacedName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Purpose
	// MISSING: PurposeData
	return out
}
func ResourcemanagerTagKeySpec_FromProto(mapCtx *direct.MapContext, in *pb.TagKey) *krm.ResourcemanagerTagKeySpec {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerTagKeySpec{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ShortName
	// MISSING: NamespacedName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Purpose
	// MISSING: PurposeData
	return out
}
func ResourcemanagerTagKeySpec_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerTagKeySpec) *pb.TagKey {
	if in == nil {
		return nil
	}
	out := &pb.TagKey{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ShortName
	// MISSING: NamespacedName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: Purpose
	// MISSING: PurposeData
	return out
}
func TagKey_FromProto(mapCtx *direct.MapContext, in *pb.TagKey) *krm.TagKey {
	if in == nil {
		return nil
	}
	out := &krm.TagKey{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Parent = direct.LazyPtr(in.GetParent())
	out.ShortName = direct.LazyPtr(in.GetShortName())
	// MISSING: NamespacedName
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Purpose = direct.Enum_FromProto(mapCtx, in.GetPurpose())
	out.PurposeData = in.PurposeData
	return out
}
func TagKey_ToProto(mapCtx *direct.MapContext, in *krm.TagKey) *pb.TagKey {
	if in == nil {
		return nil
	}
	out := &pb.TagKey{}
	out.Name = direct.ValueOf(in.Name)
	out.Parent = direct.ValueOf(in.Parent)
	out.ShortName = direct.ValueOf(in.ShortName)
	// MISSING: NamespacedName
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.ValueOf(in.Etag)
	out.Purpose = direct.Enum_ToProto[pb.Purpose](mapCtx, in.Purpose)
	out.PurposeData = in.PurposeData
	return out
}
func TagKeyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TagKey) *krm.TagKeyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TagKeyObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ShortName
	out.NamespacedName = direct.LazyPtr(in.GetNamespacedName())
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	// MISSING: Purpose
	// MISSING: PurposeData
	return out
}
func TagKeyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TagKeyObservedState) *pb.TagKey {
	if in == nil {
		return nil
	}
	out := &pb.TagKey{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ShortName
	out.NamespacedName = direct.ValueOf(in.NamespacedName)
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	// MISSING: Purpose
	// MISSING: PurposeData
	return out
}
