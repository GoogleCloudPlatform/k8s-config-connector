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

package contactcenterinsights

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ContactcenterinsightsViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.View) *krm.ContactcenterinsightsViewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsViewObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Value
	return out
}
func ContactcenterinsightsViewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsViewObservedState) *pb.View {
	if in == nil {
		return nil
	}
	out := &pb.View{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Value
	return out
}
func ContactcenterinsightsViewSpec_FromProto(mapCtx *direct.MapContext, in *pb.View) *krm.ContactcenterinsightsViewSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContactcenterinsightsViewSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Value
	return out
}
func ContactcenterinsightsViewSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContactcenterinsightsViewSpec) *pb.View {
	if in == nil {
		return nil
	}
	out := &pb.View{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Value
	return out
}
func View_FromProto(mapCtx *direct.MapContext, in *pb.View) *krm.View {
	if in == nil {
		return nil
	}
	out := &krm.View{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func View_ToProto(mapCtx *direct.MapContext, in *krm.View) *pb.View {
	if in == nil {
		return nil
	}
	out := &pb.View{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Value = direct.ValueOf(in.Value)
	return out
}
func ViewObservedState_FromProto(mapCtx *direct.MapContext, in *pb.View) *krm.ViewObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ViewObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Value
	return out
}
func ViewObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ViewObservedState) *pb.View {
	if in == nil {
		return nil
	}
	out := &pb.View{}
	// MISSING: Name
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Value
	return out
}
