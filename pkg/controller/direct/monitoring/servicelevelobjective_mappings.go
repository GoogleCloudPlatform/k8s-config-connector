// Copyright 2026 Google LLC
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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// MonitoringServiceLevelObjectiveSpec is hand-coded because the Goal field of type
// float64 is a non-pointer (value) type in KRM which cannot be passed to direct.ValueOf.
func MonitoringServiceLevelObjectiveSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceLevelObjective) *krm.MonitoringServiceLevelObjectiveSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringServiceLevelObjectiveSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ServiceLevelIndicator = ServiceLevelIndicator_FromProto(mapCtx, in.GetServiceLevelIndicator())
	out.Goal = in.GetGoal()
	out.RollingPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRollingPeriod())
	out.CalendarPeriod = direct.Enum_FromProto(mapCtx, in.GetCalendarPeriod())
	return out
}

func MonitoringServiceLevelObjectiveSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringServiceLevelObjectiveSpec) *pb.ServiceLevelObjective {
	if in == nil {
		return nil
	}
	out := &pb.ServiceLevelObjective{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ServiceLevelIndicator = ServiceLevelIndicator_ToProto(mapCtx, in.ServiceLevelIndicator)
	out.Goal = in.Goal
	if oneof := direct.StringDuration_ToProto(mapCtx, in.RollingPeriod); oneof != nil {
		out.Period = &pb.ServiceLevelObjective_RollingPeriod{RollingPeriod: oneof}
	}
	if oneof := MonitoringServiceLevelObjectiveSpec_CalendarPeriod_ToProto(mapCtx, in.CalendarPeriod); oneof != nil {
		out.Period = oneof
	}
	return out
}

func MonitoringServiceLevelObjectiveStatus_FromProto(mapCtx *direct.MapContext, in *pb.ServiceLevelObjective) *krm.MonitoringServiceLevelObjectiveStatus {
	return nil
}

func MonitoringServiceLevelObjectiveStatus_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringServiceLevelObjectiveStatus) *pb.ServiceLevelObjective {
	return nil
}
