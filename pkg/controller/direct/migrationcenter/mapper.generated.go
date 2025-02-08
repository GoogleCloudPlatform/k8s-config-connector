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
func Group_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.Group {
	if in == nil {
		return nil
	}
	out := &krm.Group{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func Group_ToProto(mapCtx *direct.MapContext, in *krm.Group) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func GroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.GroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GroupObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func GroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GroupObservedState) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func MigrationcenterGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.MigrationcenterGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterGroupObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func MigrationcenterGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterGroupObservedState) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func MigrationcenterGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.MigrationcenterGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterGroupSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func MigrationcenterGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterGroupSpec) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
