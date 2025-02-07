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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/apphub/apiv1/apphubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apphub/v1alpha1"
)
func ApphubApplicationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.ApphubApplicationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApphubApplicationObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Scope
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubApplicationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApphubApplicationObservedState) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Scope
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubApplicationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Application) *krm.ApphubApplicationSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApphubApplicationSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Scope
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubApplicationSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApphubApplicationSpec) *pb.Application {
	if in == nil {
		return nil
	}
	out := &pb.Application{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Scope
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubDiscoveredServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveredService) *krm.ApphubDiscoveredServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApphubDiscoveredServiceObservedState{}
	// MISSING: Name
	// MISSING: ServiceReference
	// MISSING: ServiceProperties
	return out
}
func ApphubDiscoveredServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApphubDiscoveredServiceObservedState) *pb.DiscoveredService {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveredService{}
	// MISSING: Name
	// MISSING: ServiceReference
	// MISSING: ServiceProperties
	return out
}
func ApphubDiscoveredServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveredService) *krm.ApphubDiscoveredServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApphubDiscoveredServiceSpec{}
	// MISSING: Name
	// MISSING: ServiceReference
	// MISSING: ServiceProperties
	return out
}
func ApphubDiscoveredServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApphubDiscoveredServiceSpec) *pb.DiscoveredService {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveredService{}
	// MISSING: Name
	// MISSING: ServiceReference
	// MISSING: ServiceProperties
	return out
}
func ApphubServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.ApphubServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApphubServiceObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceReference
	// MISSING: ServiceProperties
	// MISSING: Attributes
	// MISSING: DiscoveredService
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApphubServiceObservedState) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceReference
	// MISSING: ServiceProperties
	// MISSING: Attributes
	// MISSING: DiscoveredService
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubServiceProjectAttachmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceProjectAttachment) *krm.ApphubServiceProjectAttachmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApphubServiceProjectAttachmentObservedState{}
	// MISSING: Name
	// MISSING: ServiceProject
	// MISSING: CreateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubServiceProjectAttachmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApphubServiceProjectAttachmentObservedState) *pb.ServiceProjectAttachment {
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
func ApphubServiceProjectAttachmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceProjectAttachment) *krm.ApphubServiceProjectAttachmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApphubServiceProjectAttachmentSpec{}
	// MISSING: Name
	// MISSING: ServiceProject
	// MISSING: CreateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubServiceProjectAttachmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApphubServiceProjectAttachmentSpec) *pb.ServiceProjectAttachment {
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
func ApphubServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.ApphubServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApphubServiceSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceReference
	// MISSING: ServiceProperties
	// MISSING: Attributes
	// MISSING: DiscoveredService
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApphubServiceSpec) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ServiceReference
	// MISSING: ServiceProperties
	// MISSING: Attributes
	// MISSING: DiscoveredService
	// MISSING: CreateTime
	// MISSING: UpdateTime
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
