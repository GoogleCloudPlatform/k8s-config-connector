// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	monitoringpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/monitoring_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
)

// ServiceLevelObjectiveServer implements the gRPC interface for ServiceLevelObjective.
type ServiceLevelObjectiveServer struct{}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum enum from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(e monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum_name[int32(e)]; ok {
		e := monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(n[len("MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum enum from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(e monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum_name[int32(e)]; ok {
		e := monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(n[len("MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum enum from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(e monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum_name[int32(e)]; ok {
		e := monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(n[len("MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum enum from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(e monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum_name[int32(e)]; ok {
		e := monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(n[len("MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveCalendarPeriodEnum converts a ServiceLevelObjectiveCalendarPeriodEnum enum from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveCalendarPeriodEnum(e monitoringpb.MonitoringServiceLevelObjectiveCalendarPeriodEnum) *monitoring.ServiceLevelObjectiveCalendarPeriodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringServiceLevelObjectiveCalendarPeriodEnum_name[int32(e)]; ok {
		e := monitoring.ServiceLevelObjectiveCalendarPeriodEnum(n[len("MonitoringServiceLevelObjectiveCalendarPeriodEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveServiceLevelIndicator converts a ServiceLevelObjectiveServiceLevelIndicator object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicator(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicator) *monitoring.ServiceLevelObjectiveServiceLevelIndicator {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicator{
		BasicSli:     ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSli(p.GetBasicSli()),
		RequestBased: ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBased(p.GetRequestBased()),
		WindowsBased: ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBased(p.GetWindowsBased()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSli converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSli object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSli(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSli) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSli {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSli{
		Availability:          ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability(p.GetAvailability()),
		Latency:               ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency(p.GetLatency()),
		OperationAvailability: ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability(p.GetOperationAvailability()),
		OperationLatency:      ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency(p.GetOperationLatency()),
	}
	for _, r := range p.GetMethod() {
		obj.Method = append(obj.Method, r)
	}
	for _, r := range p.GetLocation() {
		obj.Location = append(obj.Location, r)
	}
	for _, r := range p.GetVersion() {
		obj.Version = append(obj.Version, r)
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability{}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency{
		Threshold:  dcl.StringOrNil(p.GetThreshold()),
		Experience: ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(p.GetExperience()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability{}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency{
		Threshold:  dcl.StringOrNil(p.GetThreshold()),
		Experience: ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(p.GetExperience()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorRequestBased converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBased object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBased(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBased) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorRequestBased {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorRequestBased{
		GoodTotalRatio:  ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio(p.GetGoodTotalRatio()),
		DistributionCut: ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut(p.GetDistributionCut()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio{
		GoodServiceFilter:  dcl.StringOrNil(p.GetGoodServiceFilter()),
		BadServiceFilter:   dcl.StringOrNil(p.GetBadServiceFilter()),
		TotalServiceFilter: dcl.StringOrNil(p.GetTotalServiceFilter()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut{
		DistributionFilter: dcl.StringOrNil(p.GetDistributionFilter()),
		Range:              ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange(p.GetRange()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange{
		Min: dcl.Float64OrNil(p.GetMin()),
		Max: dcl.Float64OrNil(p.GetMax()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBased converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBased object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBased(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBased) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBased {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBased{
		GoodBadMetricFilter:     dcl.StringOrNil(p.GetGoodBadMetricFilter()),
		GoodTotalRatioThreshold: ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold(p.GetGoodTotalRatioThreshold()),
		MetricMeanInRange:       ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange(p.GetMetricMeanInRange()),
		MetricSumInRange:        ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange(p.GetMetricSumInRange()),
		WindowPeriod:            dcl.StringOrNil(p.GetWindowPeriod()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold{
		Performance:         ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance(p.GetPerformance()),
		BasicSliPerformance: ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance(p.GetBasicSliPerformance()),
		Threshold:           dcl.Float64OrNil(p.GetThreshold()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance{
		GoodTotalRatio:  ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio(p.GetGoodTotalRatio()),
		DistributionCut: ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut(p.GetDistributionCut()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio{
		GoodServiceFilter:  dcl.StringOrNil(p.GetGoodServiceFilter()),
		BadServiceFilter:   dcl.StringOrNil(p.GetBadServiceFilter()),
		TotalServiceFilter: dcl.StringOrNil(p.GetTotalServiceFilter()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut{
		DistributionFilter: dcl.StringOrNil(p.GetDistributionFilter()),
		Range:              ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange(p.GetRange()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange{
		Min: dcl.Float64OrNil(p.GetMin()),
		Max: dcl.Float64OrNil(p.GetMax()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance{
		Availability:          ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability(p.GetAvailability()),
		Latency:               ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency(p.GetLatency()),
		OperationAvailability: ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability(p.GetOperationAvailability()),
		OperationLatency:      ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency(p.GetOperationLatency()),
	}
	for _, r := range p.GetMethod() {
		obj.Method = append(obj.Method, r)
	}
	for _, r := range p.GetLocation() {
		obj.Location = append(obj.Location, r)
	}
	for _, r := range p.GetVersion() {
		obj.Version = append(obj.Version, r)
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability{}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency{
		Threshold:  dcl.StringOrNil(p.GetThreshold()),
		Experience: ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(p.GetExperience()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability{}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency{
		Threshold:  dcl.StringOrNil(p.GetThreshold()),
		Experience: ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(p.GetExperience()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange{
		TimeSeries: dcl.StringOrNil(p.GetTimeSeries()),
		Range:      ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange(p.GetRange()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange{
		Min: dcl.Float64OrNil(p.GetMin()),
		Max: dcl.Float64OrNil(p.GetMax()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange{
		TimeSeries: dcl.StringOrNil(p.GetTimeSeries()),
		Range:      ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange(p.GetRange()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange object from its proto representation.
func ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange(p *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange) *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange {
	if p == nil {
		return nil
	}
	obj := &monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange{
		Min: dcl.Float64OrNil(p.GetMin()),
		Max: dcl.Float64OrNil(p.GetMax()),
	}
	return obj
}

// ProtoToServiceLevelObjective converts a ServiceLevelObjective resource from its proto representation.
func ProtoToServiceLevelObjective(p *monitoringpb.MonitoringServiceLevelObjective) *monitoring.ServiceLevelObjective {
	obj := &monitoring.ServiceLevelObjective{
		Name:                   dcl.StringOrNil(p.GetName()),
		DisplayName:            dcl.StringOrNil(p.GetDisplayName()),
		ServiceLevelIndicator:  ProtoToMonitoringServiceLevelObjectiveServiceLevelIndicator(p.GetServiceLevelIndicator()),
		Goal:                   dcl.Float64OrNil(p.GetGoal()),
		RollingPeriod:          dcl.StringOrNil(p.GetRollingPeriod()),
		CalendarPeriod:         ProtoToMonitoringServiceLevelObjectiveCalendarPeriodEnum(p.GetCalendarPeriod()),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		DeleteTime:             dcl.StringOrNil(p.GetDeleteTime()),
		ServiceManagementOwned: dcl.Bool(p.GetServiceManagementOwned()),
		Project:                dcl.StringOrNil(p.GetProject()),
		Service:                dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnumToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum enum to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnumToProto(e *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum) monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum {
	if e == nil {
		return monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(0)
	}
	if v, ok := monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum_value["ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(v)
	}
	return monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(0)
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnumToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum enum to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnumToProto(e *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum) monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum {
	if e == nil {
		return monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(0)
	}
	if v, ok := monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum_value["ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(v)
	}
	return monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(0)
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnumToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum enum to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnumToProto(e *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum) monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum {
	if e == nil {
		return monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(0)
	}
	if v, ok := monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum_value["ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(v)
	}
	return monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(0)
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnumToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum enum to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnumToProto(e *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum) monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum {
	if e == nil {
		return monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(0)
	}
	if v, ok := monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum_value["ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(v)
	}
	return monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(0)
}

// ServiceLevelObjectiveCalendarPeriodEnumToProto converts a ServiceLevelObjectiveCalendarPeriodEnum enum to its proto representation.
func MonitoringServiceLevelObjectiveCalendarPeriodEnumToProto(e *monitoring.ServiceLevelObjectiveCalendarPeriodEnum) monitoringpb.MonitoringServiceLevelObjectiveCalendarPeriodEnum {
	if e == nil {
		return monitoringpb.MonitoringServiceLevelObjectiveCalendarPeriodEnum(0)
	}
	if v, ok := monitoringpb.MonitoringServiceLevelObjectiveCalendarPeriodEnum_value["ServiceLevelObjectiveCalendarPeriodEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringServiceLevelObjectiveCalendarPeriodEnum(v)
	}
	return monitoringpb.MonitoringServiceLevelObjectiveCalendarPeriodEnum(0)
}

// ServiceLevelObjectiveServiceLevelIndicatorToProto converts a ServiceLevelObjectiveServiceLevelIndicator object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicator) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicator {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicator{}
	p.SetBasicSli(MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliToProto(o.BasicSli))
	p.SetRequestBased(MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedToProto(o.RequestBased))
	p.SetWindowsBased(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedToProto(o.WindowsBased))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSli object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSli) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSli {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSli{}
	p.SetAvailability(MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailabilityToProto(o.Availability))
	p.SetLatency(MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyToProto(o.Latency))
	p.SetOperationAvailability(MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailabilityToProto(o.OperationAvailability))
	p.SetOperationLatency(MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyToProto(o.OperationLatency))
	sMethod := make([]string, len(o.Method))
	for i, r := range o.Method {
		sMethod[i] = r
	}
	p.SetMethod(sMethod)
	sLocation := make([]string, len(o.Location))
	for i, r := range o.Location {
		sLocation[i] = r
	}
	p.SetLocation(sLocation)
	sVersion := make([]string, len(o.Version))
	for i, r := range o.Version {
		sVersion[i] = r
	}
	p.SetVersion(sVersion)
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailabilityToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailabilityToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability{}
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency{}
	p.SetThreshold(dcl.ValueOrEmptyString(o.Threshold))
	p.SetExperience(MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnumToProto(o.Experience))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailabilityToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailabilityToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability{}
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency{}
	p.SetThreshold(dcl.ValueOrEmptyString(o.Threshold))
	p.SetExperience(MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnumToProto(o.Experience))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorRequestBasedToProto converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBased object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorRequestBased) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBased {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBased{}
	p.SetGoodTotalRatio(MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatioToProto(o.GoodTotalRatio))
	p.SetDistributionCut(MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutToProto(o.DistributionCut))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatioToProto converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatioToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio{}
	p.SetGoodServiceFilter(dcl.ValueOrEmptyString(o.GoodServiceFilter))
	p.SetBadServiceFilter(dcl.ValueOrEmptyString(o.BadServiceFilter))
	p.SetTotalServiceFilter(dcl.ValueOrEmptyString(o.TotalServiceFilter))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutToProto converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut{}
	p.SetDistributionFilter(dcl.ValueOrEmptyString(o.DistributionFilter))
	p.SetRange(MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRangeToProto(o.Range))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRangeToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange{}
	p.SetMin(dcl.ValueOrEmptyDouble(o.Min))
	p.SetMax(dcl.ValueOrEmptyDouble(o.Max))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBased object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBased) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBased {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBased{}
	p.SetGoodBadMetricFilter(dcl.ValueOrEmptyString(o.GoodBadMetricFilter))
	p.SetGoodTotalRatioThreshold(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdToProto(o.GoodTotalRatioThreshold))
	p.SetMetricMeanInRange(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeToProto(o.MetricMeanInRange))
	p.SetMetricSumInRange(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeToProto(o.MetricSumInRange))
	p.SetWindowPeriod(dcl.ValueOrEmptyString(o.WindowPeriod))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold{}
	p.SetPerformance(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceToProto(o.Performance))
	p.SetBasicSliPerformance(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceToProto(o.BasicSliPerformance))
	p.SetThreshold(dcl.ValueOrEmptyDouble(o.Threshold))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance{}
	p.SetGoodTotalRatio(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatioToProto(o.GoodTotalRatio))
	p.SetDistributionCut(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutToProto(o.DistributionCut))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatioToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatioToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio{}
	p.SetGoodServiceFilter(dcl.ValueOrEmptyString(o.GoodServiceFilter))
	p.SetBadServiceFilter(dcl.ValueOrEmptyString(o.BadServiceFilter))
	p.SetTotalServiceFilter(dcl.ValueOrEmptyString(o.TotalServiceFilter))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut{}
	p.SetDistributionFilter(dcl.ValueOrEmptyString(o.DistributionFilter))
	p.SetRange(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRangeToProto(o.Range))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRangeToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange{}
	p.SetMin(dcl.ValueOrEmptyDouble(o.Min))
	p.SetMax(dcl.ValueOrEmptyDouble(o.Max))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance{}
	p.SetAvailability(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailabilityToProto(o.Availability))
	p.SetLatency(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyToProto(o.Latency))
	p.SetOperationAvailability(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailabilityToProto(o.OperationAvailability))
	p.SetOperationLatency(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyToProto(o.OperationLatency))
	sMethod := make([]string, len(o.Method))
	for i, r := range o.Method {
		sMethod[i] = r
	}
	p.SetMethod(sMethod)
	sLocation := make([]string, len(o.Location))
	for i, r := range o.Location {
		sLocation[i] = r
	}
	p.SetLocation(sLocation)
	sVersion := make([]string, len(o.Version))
	for i, r := range o.Version {
		sVersion[i] = r
	}
	p.SetVersion(sVersion)
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailabilityToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailabilityToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability{}
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency{}
	p.SetThreshold(dcl.ValueOrEmptyString(o.Threshold))
	p.SetExperience(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnumToProto(o.Experience))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailabilityToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailabilityToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability{}
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency{}
	p.SetThreshold(dcl.ValueOrEmptyString(o.Threshold))
	p.SetExperience(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnumToProto(o.Experience))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange{}
	p.SetTimeSeries(dcl.ValueOrEmptyString(o.TimeSeries))
	p.SetRange(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRangeToProto(o.Range))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRangeToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange{}
	p.SetMin(dcl.ValueOrEmptyDouble(o.Min))
	p.SetMax(dcl.ValueOrEmptyDouble(o.Max))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange{}
	p.SetTimeSeries(dcl.ValueOrEmptyString(o.TimeSeries))
	p.SetRange(MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRangeToProto(o.Range))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange object to its proto representation.
func MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRangeToProto(o *monitoring.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange) *monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange{}
	p.SetMin(dcl.ValueOrEmptyDouble(o.Min))
	p.SetMax(dcl.ValueOrEmptyDouble(o.Max))
	return p
}

// ServiceLevelObjectiveToProto converts a ServiceLevelObjective resource to its proto representation.
func ServiceLevelObjectiveToProto(resource *monitoring.ServiceLevelObjective) *monitoringpb.MonitoringServiceLevelObjective {
	p := &monitoringpb.MonitoringServiceLevelObjective{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetServiceLevelIndicator(MonitoringServiceLevelObjectiveServiceLevelIndicatorToProto(resource.ServiceLevelIndicator))
	p.SetGoal(dcl.ValueOrEmptyDouble(resource.Goal))
	p.SetRollingPeriod(dcl.ValueOrEmptyString(resource.RollingPeriod))
	p.SetCalendarPeriod(MonitoringServiceLevelObjectiveCalendarPeriodEnumToProto(resource.CalendarPeriod))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetServiceManagementOwned(dcl.ValueOrEmptyBool(resource.ServiceManagementOwned))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetService(dcl.ValueOrEmptyString(resource.Service))
	mUserLabels := make(map[string]string, len(resource.UserLabels))
	for k, r := range resource.UserLabels {
		mUserLabels[k] = r
	}
	p.SetUserLabels(mUserLabels)

	return p
}

// applyServiceLevelObjective handles the gRPC request by passing it to the underlying ServiceLevelObjective Apply() method.
func (s *ServiceLevelObjectiveServer) applyServiceLevelObjective(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringServiceLevelObjectiveRequest) (*monitoringpb.MonitoringServiceLevelObjective, error) {
	p := ProtoToServiceLevelObjective(request.GetResource())
	res, err := c.ApplyServiceLevelObjective(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceLevelObjectiveToProto(res)
	return r, nil
}

// applyMonitoringServiceLevelObjective handles the gRPC request by passing it to the underlying ServiceLevelObjective Apply() method.
func (s *ServiceLevelObjectiveServer) ApplyMonitoringServiceLevelObjective(ctx context.Context, request *monitoringpb.ApplyMonitoringServiceLevelObjectiveRequest) (*monitoringpb.MonitoringServiceLevelObjective, error) {
	cl, err := createConfigServiceLevelObjective(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyServiceLevelObjective(ctx, cl, request)
}

// DeleteServiceLevelObjective handles the gRPC request by passing it to the underlying ServiceLevelObjective Delete() method.
func (s *ServiceLevelObjectiveServer) DeleteMonitoringServiceLevelObjective(ctx context.Context, request *monitoringpb.DeleteMonitoringServiceLevelObjectiveRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServiceLevelObjective(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServiceLevelObjective(ctx, ProtoToServiceLevelObjective(request.GetResource()))

}

// ListMonitoringServiceLevelObjective handles the gRPC request by passing it to the underlying ServiceLevelObjectiveList() method.
func (s *ServiceLevelObjectiveServer) ListMonitoringServiceLevelObjective(ctx context.Context, request *monitoringpb.ListMonitoringServiceLevelObjectiveRequest) (*monitoringpb.ListMonitoringServiceLevelObjectiveResponse, error) {
	cl, err := createConfigServiceLevelObjective(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServiceLevelObjective(ctx, request.GetProject(), request.GetService())
	if err != nil {
		return nil, err
	}
	var protos []*monitoringpb.MonitoringServiceLevelObjective
	for _, r := range resources.Items {
		rp := ServiceLevelObjectiveToProto(r)
		protos = append(protos, rp)
	}
	p := &monitoringpb.ListMonitoringServiceLevelObjectiveResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigServiceLevelObjective(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
