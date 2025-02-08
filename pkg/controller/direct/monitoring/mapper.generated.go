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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func MonitoringSnoozeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Snooze) *krm.MonitoringSnoozeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringSnoozeObservedState{}
	// MISSING: Name
	// MISSING: Criteria
	// MISSING: Interval
	// MISSING: DisplayName
	return out
}
func MonitoringSnoozeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringSnoozeObservedState) *pb.Snooze {
	if in == nil {
		return nil
	}
	out := &pb.Snooze{}
	// MISSING: Name
	// MISSING: Criteria
	// MISSING: Interval
	// MISSING: DisplayName
	return out
}
func MonitoringSnoozeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Snooze) *krm.MonitoringSnoozeSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringSnoozeSpec{}
	// MISSING: Name
	// MISSING: Criteria
	// MISSING: Interval
	// MISSING: DisplayName
	return out
}
func MonitoringSnoozeSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringSnoozeSpec) *pb.Snooze {
	if in == nil {
		return nil
	}
	out := &pb.Snooze{}
	// MISSING: Name
	// MISSING: Criteria
	// MISSING: Interval
	// MISSING: DisplayName
	return out
}
func Snooze_FromProto(mapCtx *direct.MapContext, in *pb.Snooze) *krm.Snooze {
	if in == nil {
		return nil
	}
	out := &krm.Snooze{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Criteria = Snooze_Criteria_FromProto(mapCtx, in.GetCriteria())
	out.Interval = TimeInterval_FromProto(mapCtx, in.GetInterval())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func Snooze_ToProto(mapCtx *direct.MapContext, in *krm.Snooze) *pb.Snooze {
	if in == nil {
		return nil
	}
	out := &pb.Snooze{}
	out.Name = direct.ValueOf(in.Name)
	out.Criteria = Snooze_Criteria_ToProto(mapCtx, in.Criteria)
	out.Interval = TimeInterval_ToProto(mapCtx, in.Interval)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func Snooze_Criteria_FromProto(mapCtx *direct.MapContext, in *pb.Snooze_Criteria) *krm.Snooze_Criteria {
	if in == nil {
		return nil
	}
	out := &krm.Snooze_Criteria{}
	out.Policies = in.Policies
	return out
}
func Snooze_Criteria_ToProto(mapCtx *direct.MapContext, in *krm.Snooze_Criteria) *pb.Snooze_Criteria {
	if in == nil {
		return nil
	}
	out := &pb.Snooze_Criteria{}
	out.Policies = in.Policies
	return out
}
func TimeInterval_FromProto(mapCtx *direct.MapContext, in *pb.TimeInterval) *krm.TimeInterval {
	if in == nil {
		return nil
	}
	out := &krm.TimeInterval{}
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	return out
}
func TimeInterval_ToProto(mapCtx *direct.MapContext, in *krm.TimeInterval) *pb.TimeInterval {
	if in == nil {
		return nil
	}
	out := &pb.TimeInterval{}
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	return out
}
