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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1alpha1"
)
func ResourcemanagerTagValueObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TagValue) *krm.ResourcemanagerTagValueObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerTagValueObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ShortName
	// MISSING: NamespacedName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func ResourcemanagerTagValueObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerTagValueObservedState) *pb.TagValue {
	if in == nil {
		return nil
	}
	out := &pb.TagValue{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ShortName
	// MISSING: NamespacedName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func ResourcemanagerTagValueSpec_FromProto(mapCtx *direct.MapContext, in *pb.TagValue) *krm.ResourcemanagerTagValueSpec {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerTagValueSpec{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ShortName
	// MISSING: NamespacedName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func ResourcemanagerTagValueSpec_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerTagValueSpec) *pb.TagValue {
	if in == nil {
		return nil
	}
	out := &pb.TagValue{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ShortName
	// MISSING: NamespacedName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	return out
}
func TagValue_FromProto(mapCtx *direct.MapContext, in *pb.TagValue) *krm.TagValue {
	if in == nil {
		return nil
	}
	out := &krm.TagValue{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Parent = direct.LazyPtr(in.GetParent())
	out.ShortName = direct.LazyPtr(in.GetShortName())
	// MISSING: NamespacedName
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func TagValue_ToProto(mapCtx *direct.MapContext, in *krm.TagValue) *pb.TagValue {
	if in == nil {
		return nil
	}
	out := &pb.TagValue{}
	out.Name = direct.ValueOf(in.Name)
	out.Parent = direct.ValueOf(in.Parent)
	out.ShortName = direct.ValueOf(in.ShortName)
	// MISSING: NamespacedName
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func TagValueObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TagValue) *krm.TagValueObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TagValueObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ShortName
	out.NamespacedName = direct.LazyPtr(in.GetNamespacedName())
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	return out
}
func TagValueObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TagValueObservedState) *pb.TagValue {
	if in == nil {
		return nil
	}
	out := &pb.TagValue{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ShortName
	out.NamespacedName = direct.ValueOf(in.NamespacedName)
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	return out
}
