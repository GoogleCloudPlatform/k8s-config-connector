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

package netapp

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
)
func NetappSnapshotObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.NetappSnapshotObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetappSnapshotObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: Description
	// MISSING: UsedBytes
	// MISSING: CreateTime
	// MISSING: Labels
	return out
}
func NetappSnapshotObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetappSnapshotObservedState) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	// MISSING: Name
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: Description
	// MISSING: UsedBytes
	// MISSING: CreateTime
	// MISSING: Labels
	return out
}
func NetappSnapshotSpec_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.NetappSnapshotSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetappSnapshotSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: Description
	// MISSING: UsedBytes
	// MISSING: CreateTime
	// MISSING: Labels
	return out
}
func NetappSnapshotSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetappSnapshotSpec) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	// MISSING: Name
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: Description
	// MISSING: UsedBytes
	// MISSING: CreateTime
	// MISSING: Labels
	return out
}
func Snapshot_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.Snapshot {
	if in == nil {
		return nil
	}
	out := &krm.Snapshot{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: State
	// MISSING: StateDetails
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: UsedBytes
	// MISSING: CreateTime
	out.Labels = in.Labels
	return out
}
func Snapshot_ToProto(mapCtx *direct.MapContext, in *krm.Snapshot) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: State
	// MISSING: StateDetails
	out.Description = direct.ValueOf(in.Description)
	// MISSING: UsedBytes
	// MISSING: CreateTime
	out.Labels = in.Labels
	return out
}
func SnapshotObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Snapshot) *krm.SnapshotObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SnapshotObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateDetails = direct.LazyPtr(in.GetStateDetails())
	// MISSING: Description
	out.UsedBytes = direct.LazyPtr(in.GetUsedBytes())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Labels
	return out
}
func SnapshotObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SnapshotObservedState) *pb.Snapshot {
	if in == nil {
		return nil
	}
	out := &pb.Snapshot{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.Snapshot_State](mapCtx, in.State)
	out.StateDetails = direct.ValueOf(in.StateDetails)
	// MISSING: Description
	out.UsedBytes = direct.ValueOf(in.UsedBytes)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Labels
	return out
}
