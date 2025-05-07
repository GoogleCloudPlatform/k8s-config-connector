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
