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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
func Attributes_FromProto(mapCtx *direct.MapContext, in *pb.Attributes) *krm.Attributes {
	if in == nil {
		return nil
	}
	out := &krm.Attributes{}
	out.Criticality = Criticality_FromProto(mapCtx, in.GetCriticality())
	out.Environment = Environment_FromProto(mapCtx, in.GetEnvironment())
	out.DeveloperOwners = direct.Slice_FromProto(mapCtx, in.DeveloperOwners, ContactInfo_FromProto)
	out.OperatorOwners = direct.Slice_FromProto(mapCtx, in.OperatorOwners, ContactInfo_FromProto)
	out.BusinessOwners = direct.Slice_FromProto(mapCtx, in.BusinessOwners, ContactInfo_FromProto)
	return out
}
func Attributes_ToProto(mapCtx *direct.MapContext, in *krm.Attributes) *pb.Attributes {
	if in == nil {
		return nil
	}
	out := &pb.Attributes{}
	out.Criticality = Criticality_ToProto(mapCtx, in.Criticality)
	out.Environment = Environment_ToProto(mapCtx, in.Environment)
	out.DeveloperOwners = direct.Slice_ToProto(mapCtx, in.DeveloperOwners, ContactInfo_ToProto)
	out.OperatorOwners = direct.Slice_ToProto(mapCtx, in.OperatorOwners, ContactInfo_ToProto)
	out.BusinessOwners = direct.Slice_ToProto(mapCtx, in.BusinessOwners, ContactInfo_ToProto)
	return out
}
func ContactInfo_FromProto(mapCtx *direct.MapContext, in *pb.ContactInfo) *krm.ContactInfo {
	if in == nil {
		return nil
	}
	out := &krm.ContactInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Email = direct.LazyPtr(in.GetEmail())
	return out
}
func ContactInfo_ToProto(mapCtx *direct.MapContext, in *krm.ContactInfo) *pb.ContactInfo {
	if in == nil {
		return nil
	}
	out := &pb.ContactInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Email = direct.ValueOf(in.Email)
	return out
}
func Criticality_FromProto(mapCtx *direct.MapContext, in *pb.Criticality) *krm.Criticality {
	if in == nil {
		return nil
	}
	out := &krm.Criticality{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Criticality_ToProto(mapCtx *direct.MapContext, in *krm.Criticality) *pb.Criticality {
	if in == nil {
		return nil
	}
	out := &pb.Criticality{}
	out.Type = direct.Enum_ToProto[pb.Criticality_Type](mapCtx, in.Type)
	return out
}
func Environment_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.Environment {
	if in == nil {
		return nil
	}
	out := &krm.Environment{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Environment_ToProto(mapCtx *direct.MapContext, in *krm.Environment) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	out.Type = direct.Enum_ToProto[pb.Environment_Type](mapCtx, in.Type)
	return out
}
func Workload_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.Workload {
	if in == nil {
		return nil
	}
	out := &krm.Workload{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: WorkloadReference
	// MISSING: WorkloadProperties
	out.DiscoveredWorkload = direct.LazyPtr(in.GetDiscoveredWorkload())
	out.Attributes = Attributes_FromProto(mapCtx, in.GetAttributes())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func Workload_ToProto(mapCtx *direct.MapContext, in *krm.Workload) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: WorkloadReference
	// MISSING: WorkloadProperties
	out.DiscoveredWorkload = direct.ValueOf(in.DiscoveredWorkload)
	out.Attributes = Attributes_ToProto(mapCtx, in.Attributes)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Uid
	// MISSING: State
	return out
}
func WorkloadObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.WorkloadObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.WorkloadReference = WorkloadReference_FromProto(mapCtx, in.GetWorkloadReference())
	out.WorkloadProperties = WorkloadProperties_FromProto(mapCtx, in.GetWorkloadProperties())
	// MISSING: DiscoveredWorkload
	// MISSING: Attributes
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Uid = direct.LazyPtr(in.GetUid())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func WorkloadObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadObservedState) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.WorkloadReference = WorkloadReference_ToProto(mapCtx, in.WorkloadReference)
	out.WorkloadProperties = WorkloadProperties_ToProto(mapCtx, in.WorkloadProperties)
	// MISSING: DiscoveredWorkload
	// MISSING: Attributes
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Uid = direct.ValueOf(in.Uid)
	out.State = direct.Enum_ToProto[pb.Workload_State](mapCtx, in.State)
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
