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

package apphub

import (
	pb "cloud.google.com/go/apphub/apiv1/apphubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apphub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AppHubApplicationStatus_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.AppHubApplicationStatus {
	if in == nil {
		return nil
	}
	out := &krm.AppHubApplicationStatus{}
	out.ObservedState = AppHubApplicationObservedState_FromProto(mapCtx, in)
	return out
}

func AppHubApplicationStatus_ToProto(mapCtx *direct.MapContext, in *krm.AppHubApplicationStatus) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	out = AppHubApplicationObservedState_ToProto(mapCtx, in.ObservedState)
	return out
}

func AppHubApplicationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.AppHubApplicationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AppHubApplicationObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func AppHubApplicationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AppHubApplicationObservedState) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Uid = direct.ValueOf(in.Uid)
	out.State = direct.Enum_ToProto[pb.Application_State](mapCtx, in.State)
	return out
}
func AppHubApplicationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.AppHubApplicationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AppHubApplicationSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Attributes = Attributes_FromProto(mapCtx, in.GetAttributes())
	out.Scope = Scope_FromProto(mapCtx, in.GetScope())
	return out
}
func AppHubApplicationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AppHubApplicationSpec) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Attributes = Attributes_ToProto(mapCtx, in.Attributes)
	out.Scope = Scope_ToProto(mapCtx, in.Scope)
	return out
}
