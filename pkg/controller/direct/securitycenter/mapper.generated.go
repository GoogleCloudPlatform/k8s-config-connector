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

package securitycenter

import (
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func SecuritycenterSimulationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Simulation) *krm.SecuritycenterSimulationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterSimulationObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: ResourceValueConfigsMetadata
	// MISSING: CloudProvider
	return out
}
func SecuritycenterSimulationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterSimulationObservedState) *pb.Simulation {
	if in == nil {
		return nil
	}
	out := &pb.Simulation{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: ResourceValueConfigsMetadata
	// MISSING: CloudProvider
	return out
}
func SecuritycenterSimulationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Simulation) *krm.SecuritycenterSimulationSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterSimulationSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: ResourceValueConfigsMetadata
	// MISSING: CloudProvider
	return out
}
func SecuritycenterSimulationSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterSimulationSpec) *pb.Simulation {
	if in == nil {
		return nil
	}
	out := &pb.Simulation{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: ResourceValueConfigsMetadata
	// MISSING: CloudProvider
	return out
}
func Simulation_FromProto(mapCtx *direct.MapContext, in *pb.Simulation) *krm.Simulation {
	if in == nil {
		return nil
	}
	out := &krm.Simulation{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	out.ResourceValueConfigsMetadata = direct.Slice_FromProto(mapCtx, in.ResourceValueConfigsMetadata, ResourceValueConfigMetadata_FromProto)
	out.CloudProvider = direct.Enum_FromProto(mapCtx, in.GetCloudProvider())
	return out
}
func Simulation_ToProto(mapCtx *direct.MapContext, in *krm.Simulation) *pb.Simulation {
	if in == nil {
		return nil
	}
	out := &pb.Simulation{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	out.ResourceValueConfigsMetadata = direct.Slice_ToProto(mapCtx, in.ResourceValueConfigsMetadata, ResourceValueConfigMetadata_ToProto)
	out.CloudProvider = direct.Enum_ToProto[pb.CloudProvider](mapCtx, in.CloudProvider)
	return out
}
func SimulationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Simulation) *krm.SimulationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SimulationObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: ResourceValueConfigsMetadata
	// MISSING: CloudProvider
	return out
}
func SimulationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SimulationObservedState) *pb.Simulation {
	if in == nil {
		return nil
	}
	out := &pb.Simulation{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: ResourceValueConfigsMetadata
	// MISSING: CloudProvider
	return out
}
