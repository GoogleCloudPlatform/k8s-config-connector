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

package migrationcenter

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/migrationcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func MigrationcenterSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Source) *krm.MigrationcenterSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterSourceObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Type
	// MISSING: Priority
	// MISSING: Managed
	// MISSING: PendingFrameCount
	// MISSING: ErrorFrameCount
	// MISSING: State
	return out
}
func MigrationcenterSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterSourceObservedState) *pb.Source {
	if in == nil {
		return nil
	}
	out := &pb.Source{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Type
	// MISSING: Priority
	// MISSING: Managed
	// MISSING: PendingFrameCount
	// MISSING: ErrorFrameCount
	// MISSING: State
	return out
}
func MigrationcenterSourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Source) *krm.MigrationcenterSourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterSourceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Type
	// MISSING: Priority
	// MISSING: Managed
	// MISSING: PendingFrameCount
	// MISSING: ErrorFrameCount
	// MISSING: State
	return out
}
func MigrationcenterSourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterSourceSpec) *pb.Source {
	if in == nil {
		return nil
	}
	out := &pb.Source{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Type
	// MISSING: Priority
	// MISSING: Managed
	// MISSING: PendingFrameCount
	// MISSING: ErrorFrameCount
	// MISSING: State
	return out
}
func Source_FromProto(mapCtx *direct.MapContext, in *pb.Source) *krm.Source {
	if in == nil {
		return nil
	}
	out := &krm.Source{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.Managed = direct.LazyPtr(in.GetManaged())
	// MISSING: PendingFrameCount
	// MISSING: ErrorFrameCount
	// MISSING: State
	return out
}
func Source_ToProto(mapCtx *direct.MapContext, in *krm.Source) *pb.Source {
	if in == nil {
		return nil
	}
	out := &pb.Source{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Type = direct.Enum_ToProto[pb.Source_SourceType](mapCtx, in.Type)
	out.Priority = direct.ValueOf(in.Priority)
	out.Managed = direct.ValueOf(in.Managed)
	// MISSING: PendingFrameCount
	// MISSING: ErrorFrameCount
	// MISSING: State
	return out
}
func SourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Source) *krm.SourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SourceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Type
	// MISSING: Priority
	// MISSING: Managed
	out.PendingFrameCount = direct.LazyPtr(in.GetPendingFrameCount())
	out.ErrorFrameCount = direct.LazyPtr(in.GetErrorFrameCount())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func SourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SourceObservedState) *pb.Source {
	if in == nil {
		return nil
	}
	out := &pb.Source{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Type
	// MISSING: Priority
	// MISSING: Managed
	out.PendingFrameCount = direct.ValueOf(in.PendingFrameCount)
	out.ErrorFrameCount = direct.ValueOf(in.ErrorFrameCount)
	out.State = direct.Enum_ToProto[pb.Source_State](mapCtx, in.State)
	return out
}
