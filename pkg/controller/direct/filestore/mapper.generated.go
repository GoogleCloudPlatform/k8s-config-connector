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

package filestore

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/filestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/filestore/apiv1/filestorepb"
)
func FilestoreSnapshotObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.FilestoreSnapshotObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FilestoreSnapshotObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: FilesystemUsedBytes
	return out
}
func FilestoreSnapshotObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FilestoreSnapshotObservedState) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: FilesystemUsedBytes
	return out
}
func FilestoreSnapshotSpec_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.FilestoreSnapshotSpec {
	if in == nil {
		return nil
	}
	out := &krm.FilestoreSnapshotSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: FilesystemUsedBytes
	return out
}
func FilestoreSnapshotSpec_ToProto(mapCtx *direct.MapContext, in *krm.FilestoreSnapshotSpec) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: Labels
	// MISSING: FilesystemUsedBytes
	return out
}
func Snapshot_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.Snapshot {
	if in == nil {
		return nil
	}
	out := &krm.Snapshot{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	// MISSING: CreateTime
	out.Labels = in.Labels
	// MISSING: FilesystemUsedBytes
	return out
}
func Snapshot_ToProto(mapCtx *direct.MapContext, in *krm.Snapshot) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	// MISSING: CreateTime
	out.Labels = in.Labels
	// MISSING: FilesystemUsedBytes
	return out
}
func SnapshotObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.SnapshotObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SnapshotObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Labels
	out.FilesystemUsedBytes = direct.LazyPtr(in.GetFilesystemUsedBytes())
	return out
}
func SnapshotObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SnapshotObservedState) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.Snapshot_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Labels
	out.FilesystemUsedBytes = direct.ValueOf(in.FilesystemUsedBytes)
	return out
}
