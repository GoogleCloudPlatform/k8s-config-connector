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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
)
func ResourcemanagerTagHoldObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TagHold) *krm.ResourcemanagerTagHoldObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerTagHoldObservedState{}
	// MISSING: Name
	// MISSING: Holder
	// MISSING: Origin
	// MISSING: HelpLink
	// MISSING: CreateTime
	return out
}
func ResourcemanagerTagHoldObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerTagHoldObservedState) *pb.TagHold {
	if in == nil {
		return nil
	}
	out := &pb.TagHold{}
	// MISSING: Name
	// MISSING: Holder
	// MISSING: Origin
	// MISSING: HelpLink
	// MISSING: CreateTime
	return out
}
func ResourcemanagerTagHoldSpec_FromProto(mapCtx *direct.MapContext, in *pb.TagHold) *krm.ResourcemanagerTagHoldSpec {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerTagHoldSpec{}
	// MISSING: Name
	// MISSING: Holder
	// MISSING: Origin
	// MISSING: HelpLink
	// MISSING: CreateTime
	return out
}
func ResourcemanagerTagHoldSpec_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerTagHoldSpec) *pb.TagHold {
	if in == nil {
		return nil
	}
	out := &pb.TagHold{}
	// MISSING: Name
	// MISSING: Holder
	// MISSING: Origin
	// MISSING: HelpLink
	// MISSING: CreateTime
	return out
}
func TagHold_FromProto(mapCtx *direct.MapContext, in *pb.TagHold) *krm.TagHold {
	if in == nil {
		return nil
	}
	out := &krm.TagHold{}
	// MISSING: Name
	out.Holder = direct.LazyPtr(in.GetHolder())
	out.Origin = direct.LazyPtr(in.GetOrigin())
	out.HelpLink = direct.LazyPtr(in.GetHelpLink())
	// MISSING: CreateTime
	return out
}
func TagHold_ToProto(mapCtx *direct.MapContext, in *krm.TagHold) *pb.TagHold {
	if in == nil {
		return nil
	}
	out := &pb.TagHold{}
	// MISSING: Name
	out.Holder = direct.ValueOf(in.Holder)
	out.Origin = direct.ValueOf(in.Origin)
	out.HelpLink = direct.ValueOf(in.HelpLink)
	// MISSING: CreateTime
	return out
}
func TagHoldObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TagHold) *krm.TagHoldObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TagHoldObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Holder
	// MISSING: Origin
	// MISSING: HelpLink
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func TagHoldObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TagHoldObservedState) *pb.TagHold {
	if in == nil {
		return nil
	}
	out := &pb.TagHold{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Holder
	// MISSING: Origin
	// MISSING: HelpLink
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
