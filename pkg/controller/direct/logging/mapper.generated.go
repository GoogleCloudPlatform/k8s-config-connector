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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func LogView_FromProto(mapCtx *direct.MapContext, in *pb.LogView) *krm.LogView {
	if in == nil {
		return nil
	}
	out := &krm.LogView{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Filter = direct.LazyPtr(in.GetFilter())
	return out
}
func LogView_ToProto(mapCtx *direct.MapContext, in *krm.LogView) *pb.LogView {
	if in == nil {
		return nil
	}
	out := &pb.LogView{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Filter = direct.ValueOf(in.Filter)
	return out
}
func LogViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogView) *krm.LogViewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LogViewObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Filter
	return out
}
func LogViewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LogViewObservedState) *pb.LogView {
	if in == nil {
		return nil
	}
	out := &pb.LogView{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Filter
	return out
}
func LoggingLogViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogView) *krm.LoggingLogViewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogViewObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Filter
	return out
}
func LoggingLogViewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogViewObservedState) *pb.LogView {
	if in == nil {
		return nil
	}
	out := &pb.LogView{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Filter
	return out
}
func LoggingLogViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogView) *krm.LoggingLogViewSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogViewSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Filter
	return out
}
func LoggingLogViewSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogViewSpec) *pb.LogView {
	if in == nil {
		return nil
	}
	out := &pb.LogView{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Filter
	return out
}
