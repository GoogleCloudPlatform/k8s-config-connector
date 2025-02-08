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

package securitycentermanagement

import (
	pb "cloud.google.com/go/securitycentermanagement/apiv1/securitycentermanagementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycentermanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func SecuritycentermanagementSimulatedFindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SimulatedFinding) *krm.SecuritycentermanagementSimulatedFindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycentermanagementSimulatedFindingObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: Category
	// MISSING: State
	// MISSING: SourceProperties
	// MISSING: EventTime
	// MISSING: Severity
	// MISSING: FindingClass
	return out
}
func SecuritycentermanagementSimulatedFindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycentermanagementSimulatedFindingObservedState) *pb.SimulatedFinding {
	if in == nil {
		return nil
	}
	out := &pb.SimulatedFinding{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: Category
	// MISSING: State
	// MISSING: SourceProperties
	// MISSING: EventTime
	// MISSING: Severity
	// MISSING: FindingClass
	return out
}
func SecuritycentermanagementSimulatedFindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.SimulatedFinding) *krm.SecuritycentermanagementSimulatedFindingSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycentermanagementSimulatedFindingSpec{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: Category
	// MISSING: State
	// MISSING: SourceProperties
	// MISSING: EventTime
	// MISSING: Severity
	// MISSING: FindingClass
	return out
}
func SecuritycentermanagementSimulatedFindingSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycentermanagementSimulatedFindingSpec) *pb.SimulatedFinding {
	if in == nil {
		return nil
	}
	out := &pb.SimulatedFinding{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: Category
	// MISSING: State
	// MISSING: SourceProperties
	// MISSING: EventTime
	// MISSING: Severity
	// MISSING: FindingClass
	return out
}
func SimulatedFinding_FromProto(mapCtx *direct.MapContext, in *pb.SimulatedFinding) *krm.SimulatedFinding {
	if in == nil {
		return nil
	}
	out := &krm.SimulatedFinding{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Parent = direct.LazyPtr(in.GetParent())
	out.ResourceName = direct.LazyPtr(in.GetResourceName())
	out.Category = direct.LazyPtr(in.GetCategory())
	// MISSING: State
	// MISSING: SourceProperties
	out.EventTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEventTime())
	out.Severity = direct.Enum_FromProto(mapCtx, in.GetSeverity())
	out.FindingClass = direct.Enum_FromProto(mapCtx, in.GetFindingClass())
	return out
}
func SimulatedFinding_ToProto(mapCtx *direct.MapContext, in *krm.SimulatedFinding) *pb.SimulatedFinding {
	if in == nil {
		return nil
	}
	out := &pb.SimulatedFinding{}
	out.Name = direct.ValueOf(in.Name)
	out.Parent = direct.ValueOf(in.Parent)
	out.ResourceName = direct.ValueOf(in.ResourceName)
	out.Category = direct.ValueOf(in.Category)
	// MISSING: State
	// MISSING: SourceProperties
	out.EventTime = direct.StringTimestamp_ToProto(mapCtx, in.EventTime)
	out.Severity = direct.Enum_ToProto[pb.SimulatedFinding_Severity](mapCtx, in.Severity)
	out.FindingClass = direct.Enum_ToProto[pb.SimulatedFinding_FindingClass](mapCtx, in.FindingClass)
	return out
}
func SimulatedFindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SimulatedFinding) *krm.SimulatedFindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SimulatedFindingObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: Category
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: SourceProperties
	// MISSING: EventTime
	// MISSING: Severity
	// MISSING: FindingClass
	return out
}
func SimulatedFindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SimulatedFindingObservedState) *pb.SimulatedFinding {
	if in == nil {
		return nil
	}
	out := &pb.SimulatedFinding{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ResourceName
	// MISSING: Category
	out.State = direct.Enum_ToProto[pb.SimulatedFinding_State](mapCtx, in.State)
	// MISSING: SourceProperties
	// MISSING: EventTime
	// MISSING: Severity
	// MISSING: FindingClass
	return out
}
