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
	pb "cloud.google.com/go/monitoring/metricsscope/apiv1/metricsscopepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func MonitoredProject_FromProto(mapCtx *direct.MapContext, in *pb.MonitoredProject) *krm.MonitoredProject {
	if in == nil {
		return nil
	}
	out := &krm.MonitoredProject{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	return out
}
func MonitoredProject_ToProto(mapCtx *direct.MapContext, in *krm.MonitoredProject) *pb.MonitoredProject {
	if in == nil {
		return nil
	}
	out := &pb.MonitoredProject{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	return out
}
func MonitoredProjectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MonitoredProject) *krm.MonitoredProjectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MonitoredProjectObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func MonitoredProjectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MonitoredProjectObservedState) *pb.MonitoredProject {
	if in == nil {
		return nil
	}
	out := &pb.MonitoredProject{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func MonitoringMonitoredProjectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MonitoredProject) *krm.MonitoringMonitoredProjectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringMonitoredProjectObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func MonitoringMonitoredProjectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringMonitoredProjectObservedState) *pb.MonitoredProject {
	if in == nil {
		return nil
	}
	out := &pb.MonitoredProject{}
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func MonitoringMonitoredProjectSpec_FromProto(mapCtx *direct.MapContext, in *pb.MonitoredProject) *krm.MonitoringMonitoredProjectSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringMonitoredProjectSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
func MonitoringMonitoredProjectSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringMonitoredProjectSpec) *pb.MonitoredProject {
	if in == nil {
		return nil
	}
	out := &pb.MonitoredProject{}
	// MISSING: Name
	// MISSING: CreateTime
	return out
}
