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
func BasicSli_FromProto(mapCtx *direct.MapContext, in *pb.BasicSli) *krm.BasicSli {
	if in == nil {
		return nil
	}
	out := &krm.BasicSli{}
	out.Method = in.Method
	out.Location = in.Location
	out.Version = in.Version
	out.Availability = BasicSli_AvailabilityCriteria_FromProto(mapCtx, in.GetAvailability())
	out.Latency = BasicSli_LatencyCriteria_FromProto(mapCtx, in.GetLatency())
	return out
}
func BasicSli_ToProto(mapCtx *direct.MapContext, in *krm.BasicSli) *pb.BasicSli {
	if in == nil {
		return nil
	}
	out := &pb.BasicSli{}
	out.Method = in.Method
	out.Location = in.Location
	out.Version = in.Version
	if oneof := BasicSli_AvailabilityCriteria_ToProto(mapCtx, in.Availability); oneof != nil {
		out.SliCriteria = &pb.BasicSli_Availability{Availability: oneof}
	}
	if oneof := BasicSli_LatencyCriteria_ToProto(mapCtx, in.Latency); oneof != nil {
		out.SliCriteria = &pb.BasicSli_Latency{Latency: oneof}
	}
	return out
}
func BasicSli_AvailabilityCriteria_FromProto(mapCtx *direct.MapContext, in *pb.BasicSli_AvailabilityCriteria) *krm.BasicSli_AvailabilityCriteria {
	if in == nil {
		return nil
	}
	out := &krm.BasicSli_AvailabilityCriteria{}
	return out
}
func BasicSli_AvailabilityCriteria_ToProto(mapCtx *direct.MapContext, in *krm.BasicSli_AvailabilityCriteria) *pb.BasicSli_AvailabilityCriteria {
	if in == nil {
		return nil
	}
	out := &pb.BasicSli_AvailabilityCriteria{}
	return out
}
func BasicSli_LatencyCriteria_FromProto(mapCtx *direct.MapContext, in *pb.BasicSli_LatencyCriteria) *krm.BasicSli_LatencyCriteria {
	if in == nil {
		return nil
	}
	out := &krm.BasicSli_LatencyCriteria{}
	out.Threshold = direct.StringDuration_FromProto(mapCtx, in.GetThreshold())
	return out
}
func BasicSli_LatencyCriteria_ToProto(mapCtx *direct.MapContext, in *krm.BasicSli_LatencyCriteria) *pb.BasicSli_LatencyCriteria {
	if in == nil {
		return nil
	}
	out := &pb.BasicSli_LatencyCriteria{}
	out.Threshold = direct.StringDuration_ToProto(mapCtx, in.Threshold)
	return out
}
func DistributionCut_FromProto(mapCtx *direct.MapContext, in *pb.DistributionCut) *krm.DistributionCut {
	if in == nil {
		return nil
	}
	out := &krm.DistributionCut{}
	out.DistributionFilter = direct.LazyPtr(in.GetDistributionFilter())
	out.Range = Range_FromProto(mapCtx, in.GetRange())
	return out
}
func DistributionCut_ToProto(mapCtx *direct.MapContext, in *krm.DistributionCut) *pb.DistributionCut {
	if in == nil {
		return nil
	}
	out := &pb.DistributionCut{}
	out.DistributionFilter = direct.ValueOf(in.DistributionFilter)
	out.Range = Range_ToProto(mapCtx, in.Range)
	return out
}
func MonitoringServiceLevelObjectiveObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceLevelObjective) *krm.MonitoringServiceLevelObjectiveObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringServiceLevelObjectiveObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ServiceLevelIndicator
	// MISSING: Goal
	// MISSING: RollingPeriod
	// MISSING: CalendarPeriod
	// MISSING: UserLabels
	return out
}
func MonitoringServiceLevelObjectiveObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringServiceLevelObjectiveObservedState) *pb.ServiceLevelObjective {
	if in == nil {
		return nil
	}
	out := &pb.ServiceLevelObjective{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ServiceLevelIndicator
	// MISSING: Goal
	// MISSING: RollingPeriod
	// MISSING: CalendarPeriod
	// MISSING: UserLabels
	return out
}
func MonitoringServiceLevelObjectiveSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceLevelObjective) *krm.MonitoringServiceLevelObjectiveSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringServiceLevelObjectiveSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ServiceLevelIndicator
	// MISSING: Goal
	// MISSING: RollingPeriod
	// MISSING: CalendarPeriod
	// MISSING: UserLabels
	return out
}
func MonitoringServiceLevelObjectiveSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringServiceLevelObjectiveSpec) *pb.ServiceLevelObjective {
	if in == nil {
		return nil
	}
	out := &pb.ServiceLevelObjective{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ServiceLevelIndicator
	// MISSING: Goal
	// MISSING: RollingPeriod
	// MISSING: CalendarPeriod
	// MISSING: UserLabels
	return out
}
func Range_FromProto(mapCtx *direct.MapContext, in *pb.Range) *krm.Range {
	if in == nil {
		return nil
	}
	out := &krm.Range{}
	out.Min = direct.LazyPtr(in.GetMin())
	out.Max = direct.LazyPtr(in.GetMax())
	return out
}
func Range_ToProto(mapCtx *direct.MapContext, in *krm.Range) *pb.Range {
	if in == nil {
		return nil
	}
	out := &pb.Range{}
	out.Min = direct.ValueOf(in.Min)
	out.Max = direct.ValueOf(in.Max)
	return out
}
func RequestBasedSli_FromProto(mapCtx *direct.MapContext, in *pb.RequestBasedSli) *krm.RequestBasedSli {
	if in == nil {
		return nil
	}
	out := &krm.RequestBasedSli{}
	out.GoodTotalRatio = TimeSeriesRatio_FromProto(mapCtx, in.GetGoodTotalRatio())
	out.DistributionCut = DistributionCut_FromProto(mapCtx, in.GetDistributionCut())
	return out
}
func RequestBasedSli_ToProto(mapCtx *direct.MapContext, in *krm.RequestBasedSli) *pb.RequestBasedSli {
	if in == nil {
		return nil
	}
	out := &pb.RequestBasedSli{}
	if oneof := TimeSeriesRatio_ToProto(mapCtx, in.GoodTotalRatio); oneof != nil {
		out.Method = &pb.RequestBasedSli_GoodTotalRatio{GoodTotalRatio: oneof}
	}
	if oneof := DistributionCut_ToProto(mapCtx, in.DistributionCut); oneof != nil {
		out.Method = &pb.RequestBasedSli_DistributionCut{DistributionCut: oneof}
	}
	return out
}
func ServiceLevelIndicator_FromProto(mapCtx *direct.MapContext, in *pb.ServiceLevelIndicator) *krm.ServiceLevelIndicator {
	if in == nil {
		return nil
	}
	out := &krm.ServiceLevelIndicator{}
	out.BasicSli = BasicSli_FromProto(mapCtx, in.GetBasicSli())
	out.RequestBased = RequestBasedSli_FromProto(mapCtx, in.GetRequestBased())
	out.WindowsBased = WindowsBasedSli_FromProto(mapCtx, in.GetWindowsBased())
	return out
}
func ServiceLevelIndicator_ToProto(mapCtx *direct.MapContext, in *krm.ServiceLevelIndicator) *pb.ServiceLevelIndicator {
	if in == nil {
		return nil
	}
	out := &pb.ServiceLevelIndicator{}
	if oneof := BasicSli_ToProto(mapCtx, in.BasicSli); oneof != nil {
		out.Type = &pb.ServiceLevelIndicator_BasicSli{BasicSli: oneof}
	}
	if oneof := RequestBasedSli_ToProto(mapCtx, in.RequestBased); oneof != nil {
		out.Type = &pb.ServiceLevelIndicator_RequestBased{RequestBased: oneof}
	}
	if oneof := WindowsBasedSli_ToProto(mapCtx, in.WindowsBased); oneof != nil {
		out.Type = &pb.ServiceLevelIndicator_WindowsBased{WindowsBased: oneof}
	}
	return out
}
func ServiceLevelObjective_FromProto(mapCtx *direct.MapContext, in *pb.ServiceLevelObjective) *krm.ServiceLevelObjective {
	if in == nil {
		return nil
	}
	out := &krm.ServiceLevelObjective{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ServiceLevelIndicator = ServiceLevelIndicator_FromProto(mapCtx, in.GetServiceLevelIndicator())
	out.Goal = direct.LazyPtr(in.GetGoal())
	out.RollingPeriod = direct.StringDuration_FromProto(mapCtx, in.GetRollingPeriod())
	out.CalendarPeriod = direct.Enum_FromProto(mapCtx, in.GetCalendarPeriod())
	out.UserLabels = in.UserLabels
	return out
}
func ServiceLevelObjective_ToProto(mapCtx *direct.MapContext, in *krm.ServiceLevelObjective) *pb.ServiceLevelObjective {
	if in == nil {
		return nil
	}
	out := &pb.ServiceLevelObjective{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ServiceLevelIndicator = ServiceLevelIndicator_ToProto(mapCtx, in.ServiceLevelIndicator)
	out.Goal = direct.ValueOf(in.Goal)
	if oneof := direct.StringDuration_ToProto(mapCtx, in.RollingPeriod); oneof != nil {
		out.Period = &pb.ServiceLevelObjective_RollingPeriod{RollingPeriod: oneof}
	}
	if oneof := ServiceLevelObjective_CalendarPeriod_ToProto(mapCtx, in.CalendarPeriod); oneof != nil {
		out.Period = oneof
	}
	out.UserLabels = in.UserLabels
	return out
}
func TimeSeriesRatio_FromProto(mapCtx *direct.MapContext, in *pb.TimeSeriesRatio) *krm.TimeSeriesRatio {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesRatio{}
	out.GoodServiceFilter = direct.LazyPtr(in.GetGoodServiceFilter())
	out.BadServiceFilter = direct.LazyPtr(in.GetBadServiceFilter())
	out.TotalServiceFilter = direct.LazyPtr(in.GetTotalServiceFilter())
	return out
}
func TimeSeriesRatio_ToProto(mapCtx *direct.MapContext, in *krm.TimeSeriesRatio) *pb.TimeSeriesRatio {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesRatio{}
	out.GoodServiceFilter = direct.ValueOf(in.GoodServiceFilter)
	out.BadServiceFilter = direct.ValueOf(in.BadServiceFilter)
	out.TotalServiceFilter = direct.ValueOf(in.TotalServiceFilter)
	return out
}
func WindowsBasedSli_FromProto(mapCtx *direct.MapContext, in *pb.WindowsBasedSli) *krm.WindowsBasedSli {
	if in == nil {
		return nil
	}
	out := &krm.WindowsBasedSli{}
	out.GoodBadMetricFilter = direct.LazyPtr(in.GetGoodBadMetricFilter())
	out.GoodTotalRatioThreshold = WindowsBasedSli_PerformanceThreshold_FromProto(mapCtx, in.GetGoodTotalRatioThreshold())
	out.MetricMeanInRange = WindowsBasedSli_MetricRange_FromProto(mapCtx, in.GetMetricMeanInRange())
	out.MetricSumInRange = WindowsBasedSli_MetricRange_FromProto(mapCtx, in.GetMetricSumInRange())
	out.WindowPeriod = direct.StringDuration_FromProto(mapCtx, in.GetWindowPeriod())
	return out
}
func WindowsBasedSli_ToProto(mapCtx *direct.MapContext, in *krm.WindowsBasedSli) *pb.WindowsBasedSli {
	if in == nil {
		return nil
	}
	out := &pb.WindowsBasedSli{}
	if oneof := WindowsBasedSli_GoodBadMetricFilter_ToProto(mapCtx, in.GoodBadMetricFilter); oneof != nil {
		out.WindowCriterion = oneof
	}
	if oneof := WindowsBasedSli_PerformanceThreshold_ToProto(mapCtx, in.GoodTotalRatioThreshold); oneof != nil {
		out.WindowCriterion = &pb.WindowsBasedSli_GoodTotalRatioThreshold{GoodTotalRatioThreshold: oneof}
	}
	if oneof := WindowsBasedSli_MetricRange_ToProto(mapCtx, in.MetricMeanInRange); oneof != nil {
		out.WindowCriterion = &pb.WindowsBasedSli_MetricMeanInRange{MetricMeanInRange: oneof}
	}
	if oneof := WindowsBasedSli_MetricRange_ToProto(mapCtx, in.MetricSumInRange); oneof != nil {
		out.WindowCriterion = &pb.WindowsBasedSli_MetricSumInRange{MetricSumInRange: oneof}
	}
	out.WindowPeriod = direct.StringDuration_ToProto(mapCtx, in.WindowPeriod)
	return out
}
func WindowsBasedSli_MetricRange_FromProto(mapCtx *direct.MapContext, in *pb.WindowsBasedSli_MetricRange) *krm.WindowsBasedSli_MetricRange {
	if in == nil {
		return nil
	}
	out := &krm.WindowsBasedSli_MetricRange{}
	out.TimeSeries = direct.LazyPtr(in.GetTimeSeries())
	out.Range = Range_FromProto(mapCtx, in.GetRange())
	return out
}
func WindowsBasedSli_MetricRange_ToProto(mapCtx *direct.MapContext, in *krm.WindowsBasedSli_MetricRange) *pb.WindowsBasedSli_MetricRange {
	if in == nil {
		return nil
	}
	out := &pb.WindowsBasedSli_MetricRange{}
	out.TimeSeries = direct.ValueOf(in.TimeSeries)
	out.Range = Range_ToProto(mapCtx, in.Range)
	return out
}
func WindowsBasedSli_PerformanceThreshold_FromProto(mapCtx *direct.MapContext, in *pb.WindowsBasedSli_PerformanceThreshold) *krm.WindowsBasedSli_PerformanceThreshold {
	if in == nil {
		return nil
	}
	out := &krm.WindowsBasedSli_PerformanceThreshold{}
	out.Performance = RequestBasedSli_FromProto(mapCtx, in.GetPerformance())
	out.BasicSliPerformance = BasicSli_FromProto(mapCtx, in.GetBasicSliPerformance())
	out.Threshold = direct.LazyPtr(in.GetThreshold())
	return out
}
func WindowsBasedSli_PerformanceThreshold_ToProto(mapCtx *direct.MapContext, in *krm.WindowsBasedSli_PerformanceThreshold) *pb.WindowsBasedSli_PerformanceThreshold {
	if in == nil {
		return nil
	}
	out := &pb.WindowsBasedSli_PerformanceThreshold{}
	if oneof := RequestBasedSli_ToProto(mapCtx, in.Performance); oneof != nil {
		out.Type = &pb.WindowsBasedSli_PerformanceThreshold_Performance{Performance: oneof}
	}
	if oneof := BasicSli_ToProto(mapCtx, in.BasicSliPerformance); oneof != nil {
		out.Type = &pb.WindowsBasedSli_PerformanceThreshold_BasicSliPerformance{BasicSliPerformance: oneof}
	}
	out.Threshold = direct.ValueOf(in.Threshold)
	return out
}
