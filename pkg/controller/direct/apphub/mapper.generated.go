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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/apphub/apiv1/apphubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apphub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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
func DiscoveredService_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveredService) *krm.DiscoveredService {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveredService{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ServiceReference
	// MISSING: ServiceProperties
	return out
}
func DiscoveredService_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveredService) *pb.DiscoveredService {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveredService{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ServiceReference
	// MISSING: ServiceProperties
	return out
}
func DiscoveredServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveredService) *krm.DiscoveredServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveredServiceObservedState{}
	// MISSING: Name
	out.ServiceReference = ServiceReference_FromProto(mapCtx, in.GetServiceReference())
	out.ServiceProperties = ServiceProperties_FromProto(mapCtx, in.GetServiceProperties())
	return out
}
func DiscoveredServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveredServiceObservedState) *pb.DiscoveredService {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveredService{}
	// MISSING: Name
	out.ServiceReference = ServiceReference_ToProto(mapCtx, in.ServiceReference)
	out.ServiceProperties = ServiceProperties_ToProto(mapCtx, in.ServiceProperties)
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
