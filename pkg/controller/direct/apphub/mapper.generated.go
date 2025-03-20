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

// +generated:mapper
// krm.group: apphub.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.apphub.v1

package apphub

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/apphub/apiv1/apphubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apphub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AppHubDiscoveredServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveredService) *krm.AppHubDiscoveredServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AppHubDiscoveredServiceObservedState{}
	// MISSING: Name
	out.ServiceReference = ServiceReferenceObservedState_FromProto(mapCtx, in.GetServiceReference())
	out.ServiceProperties = ServicePropertiesObservedState_FromProto(mapCtx, in.GetServiceProperties())
	return out
}
func AppHubDiscoveredServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AppHubDiscoveredServiceObservedState) *pb.DiscoveredService {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveredService{}
	// MISSING: Name
	out.ServiceReference = ServiceReferenceObservedState_ToProto(mapCtx, in.ServiceReference)
	out.ServiceProperties = ServicePropertiesObservedState_ToProto(mapCtx, in.ServiceProperties)
	return out
}
func AppHubDiscoveredServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveredService) *krm.AppHubDiscoveredServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.AppHubDiscoveredServiceSpec{}
	// MISSING: Name
	return out
}
func AppHubDiscoveredServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.AppHubDiscoveredServiceSpec) *pb.DiscoveredService {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveredService{}
	// MISSING: Name
	return out
}
func AppHubDiscoveredWorkloadObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveredWorkload) *krm.AppHubDiscoveredWorkloadObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AppHubDiscoveredWorkloadObservedState{}
	// MISSING: Name
	out.WorkloadReference = WorkloadReferenceObservedState_FromProto(mapCtx, in.GetWorkloadReference())
	out.WorkloadProperties = WorkloadPropertiesObservedState_FromProto(mapCtx, in.GetWorkloadProperties())
	return out
}
func AppHubDiscoveredWorkloadObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AppHubDiscoveredWorkloadObservedState) *pb.DiscoveredWorkload {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveredWorkload{}
	// MISSING: Name
	out.WorkloadReference = WorkloadReferenceObservedState_ToProto(mapCtx, in.WorkloadReference)
	out.WorkloadProperties = WorkloadPropertiesObservedState_ToProto(mapCtx, in.WorkloadProperties)
	return out
}
func AppHubDiscoveredWorkloadSpec_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveredWorkload) *krm.AppHubDiscoveredWorkloadSpec {
	if in == nil {
		return nil
	}
	out := &krm.AppHubDiscoveredWorkloadSpec{}
	// MISSING: Name
	return out
}
func AppHubDiscoveredWorkloadSpec_ToProto(mapCtx *direct.MapContext, in *krm.AppHubDiscoveredWorkloadSpec) *pb.DiscoveredWorkload {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveredWorkload{}
	// MISSING: Name
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
func Scope_FromProto(mapCtx *direct.MapContext, in *pb.Scope) *krm.Scope {
	if in == nil {
		return nil
	}
	out := &krm.Scope{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Scope_ToProto(mapCtx *direct.MapContext, in *krm.Scope) *pb.Scope {
	if in == nil {
		return nil
	}
	out := &pb.Scope{}
	out.Type = direct.Enum_ToProto[pb.Scope_Type](mapCtx, in.Type)
	return out
}
func ServiceProperties_FromProto(mapCtx *direct.MapContext, in *pb.ServiceProperties) *krm.ServiceProperties {
	if in == nil {
		return nil
	}
	out := &krm.ServiceProperties{}
	// MISSING: GcpProject
	// MISSING: Location
	// MISSING: Zone
	return out
}
func ServiceProperties_ToProto(mapCtx *direct.MapContext, in *krm.ServiceProperties) *pb.ServiceProperties {
	if in == nil {
		return nil
	}
	out := &pb.ServiceProperties{}
	// MISSING: GcpProject
	// MISSING: Location
	// MISSING: Zone
	return out
}
func ServicePropertiesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceProperties) *krm.ServicePropertiesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServicePropertiesObservedState{}
	out.GcpProject = direct.LazyPtr(in.GetGcpProject())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Zone = direct.LazyPtr(in.GetZone())
	return out
}
func ServicePropertiesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ServicePropertiesObservedState) *pb.ServiceProperties {
	if in == nil {
		return nil
	}
	out := &pb.ServiceProperties{}
	out.GcpProject = direct.ValueOf(in.GcpProject)
	out.Location = direct.ValueOf(in.Location)
	out.Zone = direct.ValueOf(in.Zone)
	return out
}
func ServiceReference_FromProto(mapCtx *direct.MapContext, in *pb.ServiceReference) *krm.ServiceReference {
	if in == nil {
		return nil
	}
	out := &krm.ServiceReference{}
	// MISSING: URI
	return out
}
func ServiceReference_ToProto(mapCtx *direct.MapContext, in *krm.ServiceReference) *pb.ServiceReference {
	if in == nil {
		return nil
	}
	out := &pb.ServiceReference{}
	// MISSING: URI
	return out
}
func ServiceReferenceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceReference) *krm.ServiceReferenceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServiceReferenceObservedState{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func ServiceReferenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ServiceReferenceObservedState) *pb.ServiceReference {
	if in == nil {
		return nil
	}
	out := &pb.ServiceReference{}
	out.Uri = direct.ValueOf(in.URI)
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
