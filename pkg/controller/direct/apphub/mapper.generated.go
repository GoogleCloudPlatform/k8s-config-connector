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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apphub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/apphub/apiv1/apphubpb"
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
func ApphubDiscoveredWorkloadObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveredWorkload) *krm.ApphubDiscoveredWorkloadObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApphubDiscoveredWorkloadObservedState{}
	// MISSING: Name
	// MISSING: WorkloadReference
	// MISSING: WorkloadProperties
	return out
}
func ApphubDiscoveredWorkloadObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApphubDiscoveredWorkloadObservedState) *pb.DiscoveredWorkload {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveredWorkload{}
	// MISSING: Name
	// MISSING: WorkloadReference
	// MISSING: WorkloadProperties
	return out
}
func ApphubDiscoveredWorkloadSpec_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveredWorkload) *krm.ApphubDiscoveredWorkloadSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApphubDiscoveredWorkloadSpec{}
	// MISSING: Name
	// MISSING: WorkloadReference
	// MISSING: WorkloadProperties
	return out
}
func ApphubDiscoveredWorkloadSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApphubDiscoveredWorkloadSpec) *pb.DiscoveredWorkload {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveredWorkload{}
	// MISSING: Name
	// MISSING: WorkloadReference
	// MISSING: WorkloadProperties
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
func ApphubWorkloadObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.ApphubWorkloadObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ApphubWorkloadObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: WorkloadReference
	// MISSING: WorkloadProperties
	// MISSING: DiscoveredWorkload
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubWorkloadObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ApphubWorkloadObservedState) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: WorkloadReference
	// MISSING: WorkloadProperties
	// MISSING: DiscoveredWorkload
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubWorkloadSpec_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.ApphubWorkloadSpec {
	if in == nil {
		return nil
	}
	out := &krm.ApphubWorkloadSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: WorkloadReference
	// MISSING: WorkloadProperties
	// MISSING: DiscoveredWorkload
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func ApphubWorkloadSpec_ToProto(mapCtx *direct.MapContext, in *krm.ApphubWorkloadSpec) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: WorkloadReference
	// MISSING: WorkloadProperties
	// MISSING: DiscoveredWorkload
	// MISSING: Attributes
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func DiscoveredWorkload_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveredWorkload) *krm.DiscoveredWorkload {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveredWorkload{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: WorkloadReference
	// MISSING: WorkloadProperties
	return out
}
func DiscoveredWorkload_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveredWorkload) *pb.DiscoveredWorkload {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveredWorkload{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: WorkloadReference
	// MISSING: WorkloadProperties
	return out
}
func DiscoveredWorkloadObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveredWorkload) *krm.DiscoveredWorkloadObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveredWorkloadObservedState{}
	// MISSING: Name
	out.WorkloadReference = WorkloadReference_FromProto(mapCtx, in.GetWorkloadReference())
	out.WorkloadProperties = WorkloadProperties_FromProto(mapCtx, in.GetWorkloadProperties())
	return out
}
func DiscoveredWorkloadObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveredWorkloadObservedState) *pb.DiscoveredWorkload {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveredWorkload{}
	// MISSING: Name
	out.WorkloadReference = WorkloadReference_ToProto(mapCtx, in.WorkloadReference)
	out.WorkloadProperties = WorkloadProperties_ToProto(mapCtx, in.WorkloadProperties)
	return out
}
func WorkloadProperties_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadProperties) *krm.WorkloadProperties {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadProperties{}
	// MISSING: GcpProject
	// MISSING: Location
	// MISSING: Zone
	return out
}
func WorkloadProperties_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadProperties) *pb.WorkloadProperties {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadProperties{}
	// MISSING: GcpProject
	// MISSING: Location
	// MISSING: Zone
	return out
}
func WorkloadPropertiesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadProperties) *krm.WorkloadPropertiesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadPropertiesObservedState{}
	out.GcpProject = direct.LazyPtr(in.GetGcpProject())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Zone = direct.LazyPtr(in.GetZone())
	return out
}
func WorkloadPropertiesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadPropertiesObservedState) *pb.WorkloadProperties {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadProperties{}
	out.GcpProject = direct.ValueOf(in.GcpProject)
	out.Location = direct.ValueOf(in.Location)
	out.Zone = direct.ValueOf(in.Zone)
	return out
}
func WorkloadReference_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadReference) *krm.WorkloadReference {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadReference{}
	// MISSING: URI
	return out
}
func WorkloadReference_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadReference) *pb.WorkloadReference {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadReference{}
	// MISSING: URI
	return out
}
func WorkloadReferenceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadReference) *krm.WorkloadReferenceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadReferenceObservedState{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func WorkloadReferenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadReferenceObservedState) *pb.WorkloadReference {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadReference{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
