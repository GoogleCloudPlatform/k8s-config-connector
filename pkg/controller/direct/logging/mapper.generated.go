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

package logging

import (
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func LogExclusion_FromProto(mapCtx *direct.MapContext, in *pb.LogExclusion) *krm.LogExclusion {
	if in == nil {
		return nil
	}
	out := &krm.LogExclusion{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LogExclusion_ToProto(mapCtx *direct.MapContext, in *krm.LogExclusion) *pb.LogExclusion {
	if in == nil {
		return nil
	}
	out := &pb.LogExclusion{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Filter = direct.ValueOf(in.Filter)
	out.Disabled = direct.ValueOf(in.Disabled)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LogExclusionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogExclusion) *krm.LogExclusionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LogExclusionObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Disabled
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func LogExclusionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LogExclusionObservedState) *pb.LogExclusion {
	if in == nil {
		return nil
	}
	out := &pb.LogExclusion{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Disabled
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func LoggingLogExclusionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogExclusion) *krm.LoggingLogExclusionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogExclusionObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Disabled
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LoggingLogExclusionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogExclusionObservedState) *pb.LogExclusion {
	if in == nil {
		return nil
	}
	out := &pb.LogExclusion{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Disabled
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LoggingLogExclusionSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogExclusion) *krm.LoggingLogExclusionSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogExclusionSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Disabled
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func LoggingLogExclusionSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogExclusionSpec) *pb.LogExclusion {
	if in == nil {
		return nil
	}
	out := &pb.LogExclusion{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Disabled
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
