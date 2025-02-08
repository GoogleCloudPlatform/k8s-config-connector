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
	pb "cloud.google.com/go/resourcemanager/apiv2/resourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Folder_FromProto(mapCtx *direct.MapContext, in *pb.Folder) *krm.Folder {
	if in == nil {
		return nil
	}
	out := &krm.Folder{}
	// MISSING: Name
	out.Parent = direct.LazyPtr(in.GetParent())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: LifecycleState
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Folder_ToProto(mapCtx *direct.MapContext, in *krm.Folder) *pb.Folder {
	if in == nil {
		return nil
	}
	out := &pb.Folder{}
	// MISSING: Name
	out.Parent = direct.ValueOf(in.Parent)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: LifecycleState
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func FolderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Folder) *krm.FolderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FolderObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Parent
	// MISSING: DisplayName
	out.LifecycleState = direct.Enum_FromProto(mapCtx, in.GetLifecycleState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func FolderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FolderObservedState) *pb.Folder {
	if in == nil {
		return nil
	}
	out := &pb.Folder{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Parent
	// MISSING: DisplayName
	out.LifecycleState = direct.Enum_ToProto[pb.Folder_LifecycleState](mapCtx, in.LifecycleState)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func ResourcemanagerFolderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Folder) *krm.ResourcemanagerFolderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerFolderObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: DisplayName
	// MISSING: LifecycleState
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ResourcemanagerFolderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerFolderObservedState) *pb.Folder {
	if in == nil {
		return nil
	}
	out := &pb.Folder{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: DisplayName
	// MISSING: LifecycleState
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ResourcemanagerFolderSpec_FromProto(mapCtx *direct.MapContext, in *pb.Folder) *krm.ResourcemanagerFolderSpec {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerFolderSpec{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: DisplayName
	// MISSING: LifecycleState
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ResourcemanagerFolderSpec_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerFolderSpec) *pb.Folder {
	if in == nil {
		return nil
	}
	out := &pb.Folder{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: DisplayName
	// MISSING: LifecycleState
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
