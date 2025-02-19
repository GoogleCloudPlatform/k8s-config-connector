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

func AppHubServiceProjectAttachmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceProjectAttachment) *krm.AppHubServiceProjectAttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AppHubServiceProjectAttachmentObservedState{}
	// MISSING: Name
	// MISSING: ServiceProject
	// MISSING: CreateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func AppHubServiceProjectAttachmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AppHubServiceProjectAttachmentObservedState) *pb.ServiceProjectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.ServiceProjectAttachment{}
	// MISSING: Name
	// MISSING: ServiceProject
	// MISSING: CreateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func AppHubServiceProjectAttachmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceProjectAttachment) *krm.AppHubServiceProjectAttachmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.AppHubServiceProjectAttachmentSpec{}
	// MISSING: Name
	// MISSING: ServiceProject
	// MISSING: CreateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func AppHubServiceProjectAttachmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.AppHubServiceProjectAttachmentSpec) *pb.ServiceProjectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.ServiceProjectAttachment{}
	// MISSING: Name
	// MISSING: ServiceProject
	// MISSING: CreateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func ServiceProjectAttachment_FromProto(mapCtx *direct.MapContext, in *pb.ServiceProjectAttachment) *krm.ServiceProjectAttachment {
	if in == nil {
		return nil
	}
	out := &krm.ServiceProjectAttachment{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ServiceProject = direct.LazyPtr(in.GetServiceProject())
	// MISSING: CreateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func ServiceProjectAttachment_ToProto(mapCtx *direct.MapContext, in *krm.ServiceProjectAttachment) *pb.ServiceProjectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.ServiceProjectAttachment{}
	out.Name = direct.ValueOf(in.Name)
	out.ServiceProject = direct.ValueOf(in.ServiceProject)
	// MISSING: CreateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func ServiceProjectAttachmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceProjectAttachment) *krm.ServiceProjectAttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServiceProjectAttachmentObservedState{}
	// MISSING: Name
	// MISSING: ServiceProject
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func ServiceProjectAttachmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ServiceProjectAttachmentObservedState) *pb.ServiceProjectAttachment {
	if in == nil {
		return nil
	}
	out := &pb.ServiceProjectAttachment{}
	// MISSING: Name
	// MISSING: ServiceProject
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Uid = direct.ValueOf(in.Uid)
	out.State = direct.Enum_ToProto[pb.ServiceProjectAttachment_State](mapCtx, in.State)
	return out
}
