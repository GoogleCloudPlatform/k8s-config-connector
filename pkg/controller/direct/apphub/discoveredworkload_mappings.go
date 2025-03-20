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
