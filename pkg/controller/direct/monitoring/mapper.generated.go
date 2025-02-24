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

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Group_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.Group {
	if in == nil {
		return nil
	}
	out := &krm.Group{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ParentName = direct.LazyPtr(in.GetParentName())
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.IsCluster = direct.LazyPtr(in.GetIsCluster())
	return out
}
func Group_ToProto(mapCtx *direct.MapContext, in *krm.Group) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ParentName = direct.ValueOf(in.ParentName)
	out.Filter = direct.ValueOf(in.Filter)
	out.IsCluster = direct.ValueOf(in.IsCluster)
	return out
}
func MonitoringGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.MonitoringGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringGroupObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ParentName
	// MISSING: Filter
	// MISSING: IsCluster
	return out
}
func MonitoringGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringGroupObservedState) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ParentName
	// MISSING: Filter
	// MISSING: IsCluster
	return out
}
func MonitoringGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.Group) *krm.MonitoringGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringGroupSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ParentName
	// MISSING: Filter
	// MISSING: IsCluster
	return out
}
func MonitoringGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringGroupSpec) *pb.Group {
	if in == nil {
		return nil
	}
	out := &pb.Group{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ParentName
	// MISSING: Filter
	// MISSING: IsCluster
	return out
}
