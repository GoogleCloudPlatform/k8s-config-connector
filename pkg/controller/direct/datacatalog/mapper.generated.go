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

package datacatalog

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datacatalog/lineage/apiv1/lineagepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DatacatalogRunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Run) *krm.DatacatalogRunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatacatalogRunObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Attributes
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	return out
}
func DatacatalogRunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatacatalogRunObservedState) *pb.Run {
	if in == nil {
		return nil
	}
	out := &pb.Run{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Attributes
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	return out
}
func DatacatalogRunSpec_FromProto(mapCtx *direct.MapContext, in *pb.Run) *krm.DatacatalogRunSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatacatalogRunSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Attributes
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	return out
}
func DatacatalogRunSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatacatalogRunSpec) *pb.Run {
	if in == nil {
		return nil
	}
	out := &pb.Run{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Attributes
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: State
	return out
}
func Run_FromProto(mapCtx *direct.MapContext, in *pb.Run) *krm.Run {
	if in == nil {
		return nil
	}
	out := &krm.Run{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Attributes
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func Run_ToProto(mapCtx *direct.MapContext, in *krm.Run) *pb.Run {
	if in == nil {
		return nil
	}
	out := &pb.Run{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Attributes
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.State = direct.Enum_ToProto[pb.Run_State](mapCtx, in.State)
	return out
}
