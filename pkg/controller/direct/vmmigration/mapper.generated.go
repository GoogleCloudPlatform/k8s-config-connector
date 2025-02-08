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

package vmmigration

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmmigration/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
)
func TargetProject_FromProto(mapCtx *direct.MapContext, in *pb.TargetProject) *krm.TargetProject {
	if in == nil {
		return nil
	}
	out := &krm.TargetProject{}
	// MISSING: Name
	out.Project = direct.LazyPtr(in.GetProject())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TargetProject_ToProto(mapCtx *direct.MapContext, in *krm.TargetProject) *pb.TargetProject {
	if in == nil {
		return nil
	}
	out := &pb.TargetProject{}
	// MISSING: Name
	out.Project = direct.ValueOf(in.Project)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func TargetProjectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TargetProject) *krm.TargetProjectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TargetProjectObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Project
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func TargetProjectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TargetProjectObservedState) *pb.TargetProject {
	if in == nil {
		return nil
	}
	out := &pb.TargetProject{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Project
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func VmmigrationTargetProjectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TargetProject) *krm.VmmigrationTargetProjectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationTargetProjectObservedState{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func VmmigrationTargetProjectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationTargetProjectObservedState) *pb.TargetProject {
	if in == nil {
		return nil
	}
	out := &pb.TargetProject{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func VmmigrationTargetProjectSpec_FromProto(mapCtx *direct.MapContext, in *pb.TargetProject) *krm.VmmigrationTargetProjectSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationTargetProjectSpec{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func VmmigrationTargetProjectSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationTargetProjectSpec) *pb.TargetProject {
	if in == nil {
		return nil
	}
	out := &pb.TargetProject{}
	// MISSING: Name
	// MISSING: Project
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
