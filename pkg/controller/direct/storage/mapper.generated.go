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

package storage

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Folder_FromProto(mapCtx *direct.MapContext, in *pb.Folder) *krm.Folder {
	if in == nil {
		return nil
	}
	out := &krm.Folder{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Metageneration
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PendingRenameInfo
	return out
}
func Folder_ToProto(mapCtx *direct.MapContext, in *krm.Folder) *pb.Folder {
	if in == nil {
		return nil
	}
	out := &pb.Folder{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Metageneration
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PendingRenameInfo
	return out
}
func FolderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Folder) *krm.FolderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FolderObservedState{}
	// MISSING: Name
	out.Metageneration = direct.LazyPtr(in.GetMetageneration())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.PendingRenameInfo = PendingRenameInfo_FromProto(mapCtx, in.GetPendingRenameInfo())
	return out
}
func FolderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FolderObservedState) *pb.Folder {
	if in == nil {
		return nil
	}
	out := &pb.Folder{}
	// MISSING: Name
	out.Metageneration = direct.ValueOf(in.Metageneration)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.PendingRenameInfo = PendingRenameInfo_ToProto(mapCtx, in.PendingRenameInfo)
	return out
}
func PendingRenameInfo_FromProto(mapCtx *direct.MapContext, in *pb.PendingRenameInfo) *krm.PendingRenameInfo {
	if in == nil {
		return nil
	}
	out := &krm.PendingRenameInfo{}
	// MISSING: Operation
	return out
}
func PendingRenameInfo_ToProto(mapCtx *direct.MapContext, in *krm.PendingRenameInfo) *pb.PendingRenameInfo {
	if in == nil {
		return nil
	}
	out := &pb.PendingRenameInfo{}
	// MISSING: Operation
	return out
}
func PendingRenameInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PendingRenameInfo) *krm.PendingRenameInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PendingRenameInfoObservedState{}
	out.Operation = direct.LazyPtr(in.GetOperation())
	return out
}
func PendingRenameInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PendingRenameInfoObservedState) *pb.PendingRenameInfo {
	if in == nil {
		return nil
	}
	out := &pb.PendingRenameInfo{}
	out.Operation = direct.ValueOf(in.Operation)
	return out
}
func StorageFolderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Folder) *krm.StorageFolderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StorageFolderObservedState{}
	// MISSING: Name
	// MISSING: Metageneration
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PendingRenameInfo
	return out
}
func StorageFolderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StorageFolderObservedState) *pb.Folder {
	if in == nil {
		return nil
	}
	out := &pb.Folder{}
	// MISSING: Name
	// MISSING: Metageneration
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PendingRenameInfo
	return out
}
func StorageFolderSpec_FromProto(mapCtx *direct.MapContext, in *pb.Folder) *krm.StorageFolderSpec {
	if in == nil {
		return nil
	}
	out := &krm.StorageFolderSpec{}
	// MISSING: Name
	// MISSING: Metageneration
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PendingRenameInfo
	return out
}
func StorageFolderSpec_ToProto(mapCtx *direct.MapContext, in *krm.StorageFolderSpec) *pb.Folder {
	if in == nil {
		return nil
	}
	out := &pb.Folder{}
	// MISSING: Name
	// MISSING: Metageneration
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: PendingRenameInfo
	return out
}
