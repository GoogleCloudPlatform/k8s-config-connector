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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/beta/monitoring_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/beta"
)

// ServiceLevelObjectiveServer implements the gRPC interface for ServiceLevelObjective.
type ServiceLevelObjectiveServer struct{}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum enum from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(e betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum) *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum_name[int32(e)]; ok {
		e := beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(n[len("MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum enum from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(e betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum) *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum_name[int32(e)]; ok {
		e := beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(n[len("MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum enum from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(e betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum_name[int32(e)]; ok {
		e := beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(n[len("MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum enum from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(e betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum_name[int32(e)]; ok {
		e := beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(n[len("MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveCalendarPeriodEnum converts a ServiceLevelObjectiveCalendarPeriodEnum enum from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveCalendarPeriodEnum(e betapb.MonitoringBetaServiceLevelObjectiveCalendarPeriodEnum) *beta.ServiceLevelObjectiveCalendarPeriodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.MonitoringBetaServiceLevelObjectiveCalendarPeriodEnum_name[int32(e)]; ok {
		e := beta.ServiceLevelObjectiveCalendarPeriodEnum(n[len("MonitoringBetaServiceLevelObjectiveCalendarPeriodEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveServiceLevelIndicator converts a ServiceLevelObjectiveServiceLevelIndicator object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicator(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicator) *beta.ServiceLevelObjectiveServiceLevelIndicator {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicator{
		BasicSli:     ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSli(p.GetBasicSli()),
		RequestBased: ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBased(p.GetRequestBased()),
		WindowsBased: ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBased(p.GetWindowsBased()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSli converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSli object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSli(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSli) *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSli {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSli{
		Availability:          ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability(p.GetAvailability()),
		Latency:               ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency(p.GetLatency()),
		OperationAvailability: ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability(p.GetOperationAvailability()),
		OperationLatency:      ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency(p.GetOperationLatency()),
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
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability) *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability{}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency) *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency{
		Threshold:  dcl.StringOrNil(p.GetThreshold()),
		Experience: ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(p.GetExperience()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability) *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability{}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency) *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency{
		Threshold:  dcl.StringOrNil(p.GetThreshold()),
		Experience: ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(p.GetExperience()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorRequestBased converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBased object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBased(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBased) *beta.ServiceLevelObjectiveServiceLevelIndicatorRequestBased {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorRequestBased{
		GoodTotalRatio:  ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio(p.GetGoodTotalRatio()),
		DistributionCut: ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut(p.GetDistributionCut()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio) *beta.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio{
		GoodServiceFilter:  dcl.StringOrNil(p.GetGoodServiceFilter()),
		BadServiceFilter:   dcl.StringOrNil(p.GetBadServiceFilter()),
		TotalServiceFilter: dcl.StringOrNil(p.GetTotalServiceFilter()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut) *beta.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut{
		DistributionFilter: dcl.StringOrNil(p.GetDistributionFilter()),
		Range:              ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange(p.GetRange()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange) *beta.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange{
		Min: dcl.Float64OrNil(p.GetMin()),
		Max: dcl.Float64OrNil(p.GetMax()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBased converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBased object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBased(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBased) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBased {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBased{
		GoodBadMetricFilter:     dcl.StringOrNil(p.GetGoodBadMetricFilter()),
		GoodTotalRatioThreshold: ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold(p.GetGoodTotalRatioThreshold()),
		MetricMeanInRange:       ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange(p.GetMetricMeanInRange()),
		MetricSumInRange:        ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange(p.GetMetricSumInRange()),
		WindowPeriod:            dcl.StringOrNil(p.GetWindowPeriod()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold{
		Performance:         ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance(p.GetPerformance()),
		BasicSliPerformance: ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance(p.GetBasicSliPerformance()),
		Threshold:           dcl.Float64OrNil(p.GetThreshold()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance{
		GoodTotalRatio:  ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio(p.GetGoodTotalRatio()),
		DistributionCut: ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut(p.GetDistributionCut()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio{
		GoodServiceFilter:  dcl.StringOrNil(p.GetGoodServiceFilter()),
		BadServiceFilter:   dcl.StringOrNil(p.GetBadServiceFilter()),
		TotalServiceFilter: dcl.StringOrNil(p.GetTotalServiceFilter()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut{
		DistributionFilter: dcl.StringOrNil(p.GetDistributionFilter()),
		Range:              ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange(p.GetRange()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange{
		Min: dcl.Float64OrNil(p.GetMin()),
		Max: dcl.Float64OrNil(p.GetMax()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance{
		Availability:          ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability(p.GetAvailability()),
		Latency:               ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency(p.GetLatency()),
		OperationAvailability: ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability(p.GetOperationAvailability()),
		OperationLatency:      ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency(p.GetOperationLatency()),
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
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability{}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency{
		Threshold:  dcl.StringOrNil(p.GetThreshold()),
		Experience: ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(p.GetExperience()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability{}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency{
		Threshold:  dcl.StringOrNil(p.GetThreshold()),
		Experience: ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(p.GetExperience()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange{
		TimeSeries: dcl.StringOrNil(p.GetTimeSeries()),
		Range:      ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange(p.GetRange()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange{
		Min: dcl.Float64OrNil(p.GetMin()),
		Max: dcl.Float64OrNil(p.GetMax()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange{
		TimeSeries: dcl.StringOrNil(p.GetTimeSeries()),
		Range:      ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange(p.GetRange()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange object from its proto representation.
func ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange(p *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange) *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange{
		Min: dcl.Float64OrNil(p.GetMin()),
		Max: dcl.Float64OrNil(p.GetMax()),
	}
	return obj
}

// ProtoToServiceLevelObjective converts a ServiceLevelObjective resource from its proto representation.
func ProtoToServiceLevelObjective(p *betapb.MonitoringBetaServiceLevelObjective) *beta.ServiceLevelObjective {
	obj := &beta.ServiceLevelObjective{
		Name:                   dcl.StringOrNil(p.GetName()),
		DisplayName:            dcl.StringOrNil(p.GetDisplayName()),
		ServiceLevelIndicator:  ProtoToMonitoringBetaServiceLevelObjectiveServiceLevelIndicator(p.GetServiceLevelIndicator()),
		Goal:                   dcl.Float64OrNil(p.GetGoal()),
		RollingPeriod:          dcl.StringOrNil(p.GetRollingPeriod()),
		CalendarPeriod:         ProtoToMonitoringBetaServiceLevelObjectiveCalendarPeriodEnum(p.GetCalendarPeriod()),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		DeleteTime:             dcl.StringOrNil(p.GetDeleteTime()),
		ServiceManagementOwned: dcl.Bool(p.GetServiceManagementOwned()),
		Project:                dcl.StringOrNil(p.GetProject()),
		Service:                dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnumToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum enum to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnumToProto(e *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum) betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum {
	if e == nil {
		return betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(0)
	}
	if v, ok := betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum_value["ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(v)
	}
	return betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(0)
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnumToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum enum to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnumToProto(e *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum) betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum {
	if e == nil {
		return betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(0)
	}
	if v, ok := betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum_value["ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(v)
	}
	return betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(0)
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnumToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum enum to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnumToProto(e *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum) betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum {
	if e == nil {
		return betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(0)
	}
	if v, ok := betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum_value["ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(v)
	}
	return betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(0)
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnumToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum enum to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnumToProto(e *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum) betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum {
	if e == nil {
		return betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(0)
	}
	if v, ok := betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum_value["ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(v)
	}
	return betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(0)
}

// ServiceLevelObjectiveCalendarPeriodEnumToProto converts a ServiceLevelObjectiveCalendarPeriodEnum enum to its proto representation.
func MonitoringBetaServiceLevelObjectiveCalendarPeriodEnumToProto(e *beta.ServiceLevelObjectiveCalendarPeriodEnum) betapb.MonitoringBetaServiceLevelObjectiveCalendarPeriodEnum {
	if e == nil {
		return betapb.MonitoringBetaServiceLevelObjectiveCalendarPeriodEnum(0)
	}
	if v, ok := betapb.MonitoringBetaServiceLevelObjectiveCalendarPeriodEnum_value["ServiceLevelObjectiveCalendarPeriodEnum"+string(*e)]; ok {
		return betapb.MonitoringBetaServiceLevelObjectiveCalendarPeriodEnum(v)
	}
	return betapb.MonitoringBetaServiceLevelObjectiveCalendarPeriodEnum(0)
}

// ServiceLevelObjectiveServiceLevelIndicatorToProto converts a ServiceLevelObjectiveServiceLevelIndicator object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicator) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicator {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicator{}
	p.SetBasicSli(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliToProto(o.BasicSli))
	p.SetRequestBased(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedToProto(o.RequestBased))
	p.SetWindowsBased(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedToProto(o.WindowsBased))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSli object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSli) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSli {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSli{}
	p.SetAvailability(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailabilityToProto(o.Availability))
	p.SetLatency(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyToProto(o.Latency))
	p.SetOperationAvailability(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailabilityToProto(o.OperationAvailability))
	p.SetOperationLatency(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyToProto(o.OperationLatency))
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
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailabilityToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability{}
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency{}
	p.SetThreshold(dcl.ValueOrEmptyString(o.Threshold))
	p.SetExperience(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnumToProto(o.Experience))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailabilityToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailabilityToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability{}
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency{}
	p.SetThreshold(dcl.ValueOrEmptyString(o.Threshold))
	p.SetExperience(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnumToProto(o.Experience))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorRequestBasedToProto converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBased object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorRequestBased) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBased {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBased{}
	p.SetGoodTotalRatio(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatioToProto(o.GoodTotalRatio))
	p.SetDistributionCut(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutToProto(o.DistributionCut))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatioToProto converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatioToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio{}
	p.SetGoodServiceFilter(dcl.ValueOrEmptyString(o.GoodServiceFilter))
	p.SetBadServiceFilter(dcl.ValueOrEmptyString(o.BadServiceFilter))
	p.SetTotalServiceFilter(dcl.ValueOrEmptyString(o.TotalServiceFilter))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutToProto converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut{}
	p.SetDistributionFilter(dcl.ValueOrEmptyString(o.DistributionFilter))
	p.SetRange(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRangeToProto(o.Range))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRangeToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange{}
	p.SetMin(dcl.ValueOrEmptyDouble(o.Min))
	p.SetMax(dcl.ValueOrEmptyDouble(o.Max))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBased object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBased) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBased {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBased{}
	p.SetGoodBadMetricFilter(dcl.ValueOrEmptyString(o.GoodBadMetricFilter))
	p.SetGoodTotalRatioThreshold(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdToProto(o.GoodTotalRatioThreshold))
	p.SetMetricMeanInRange(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeToProto(o.MetricMeanInRange))
	p.SetMetricSumInRange(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeToProto(o.MetricSumInRange))
	p.SetWindowPeriod(dcl.ValueOrEmptyString(o.WindowPeriod))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold{}
	p.SetPerformance(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceToProto(o.Performance))
	p.SetBasicSliPerformance(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceToProto(o.BasicSliPerformance))
	p.SetThreshold(dcl.ValueOrEmptyDouble(o.Threshold))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance{}
	p.SetGoodTotalRatio(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatioToProto(o.GoodTotalRatio))
	p.SetDistributionCut(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutToProto(o.DistributionCut))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatioToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatioToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio{}
	p.SetGoodServiceFilter(dcl.ValueOrEmptyString(o.GoodServiceFilter))
	p.SetBadServiceFilter(dcl.ValueOrEmptyString(o.BadServiceFilter))
	p.SetTotalServiceFilter(dcl.ValueOrEmptyString(o.TotalServiceFilter))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut{}
	p.SetDistributionFilter(dcl.ValueOrEmptyString(o.DistributionFilter))
	p.SetRange(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRangeToProto(o.Range))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRangeToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange{}
	p.SetMin(dcl.ValueOrEmptyDouble(o.Min))
	p.SetMax(dcl.ValueOrEmptyDouble(o.Max))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance{}
	p.SetAvailability(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailabilityToProto(o.Availability))
	p.SetLatency(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyToProto(o.Latency))
	p.SetOperationAvailability(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailabilityToProto(o.OperationAvailability))
	p.SetOperationLatency(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyToProto(o.OperationLatency))
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
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailabilityToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability{}
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency{}
	p.SetThreshold(dcl.ValueOrEmptyString(o.Threshold))
	p.SetExperience(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnumToProto(o.Experience))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailabilityToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailabilityToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability{}
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency{}
	p.SetThreshold(dcl.ValueOrEmptyString(o.Threshold))
	p.SetExperience(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnumToProto(o.Experience))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange{}
	p.SetTimeSeries(dcl.ValueOrEmptyString(o.TimeSeries))
	p.SetRange(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRangeToProto(o.Range))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRangeToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange{}
	p.SetMin(dcl.ValueOrEmptyDouble(o.Min))
	p.SetMax(dcl.ValueOrEmptyDouble(o.Max))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange{}
	p.SetTimeSeries(dcl.ValueOrEmptyString(o.TimeSeries))
	p.SetRange(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRangeToProto(o.Range))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange object to its proto representation.
func MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRangeToProto(o *beta.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange) *betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange {
	if o == nil {
		return nil
	}
	p := &betapb.MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange{}
	p.SetMin(dcl.ValueOrEmptyDouble(o.Min))
	p.SetMax(dcl.ValueOrEmptyDouble(o.Max))
	return p
}

// ServiceLevelObjectiveToProto converts a ServiceLevelObjective resource to its proto representation.
func ServiceLevelObjectiveToProto(resource *beta.ServiceLevelObjective) *betapb.MonitoringBetaServiceLevelObjective {
	p := &betapb.MonitoringBetaServiceLevelObjective{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetServiceLevelIndicator(MonitoringBetaServiceLevelObjectiveServiceLevelIndicatorToProto(resource.ServiceLevelIndicator))
	p.SetGoal(dcl.ValueOrEmptyDouble(resource.Goal))
	p.SetRollingPeriod(dcl.ValueOrEmptyString(resource.RollingPeriod))
	p.SetCalendarPeriod(MonitoringBetaServiceLevelObjectiveCalendarPeriodEnumToProto(resource.CalendarPeriod))
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
func (s *ServiceLevelObjectiveServer) applyServiceLevelObjective(ctx context.Context, c *beta.Client, request *betapb.ApplyMonitoringBetaServiceLevelObjectiveRequest) (*betapb.MonitoringBetaServiceLevelObjective, error) {
	p := ProtoToServiceLevelObjective(request.GetResource())
	res, err := c.ApplyServiceLevelObjective(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceLevelObjectiveToProto(res)
	return r, nil
}

// applyMonitoringBetaServiceLevelObjective handles the gRPC request by passing it to the underlying ServiceLevelObjective Apply() method.
func (s *ServiceLevelObjectiveServer) ApplyMonitoringBetaServiceLevelObjective(ctx context.Context, request *betapb.ApplyMonitoringBetaServiceLevelObjectiveRequest) (*betapb.MonitoringBetaServiceLevelObjective, error) {
	cl, err := createConfigServiceLevelObjective(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyServiceLevelObjective(ctx, cl, request)
}

// DeleteServiceLevelObjective handles the gRPC request by passing it to the underlying ServiceLevelObjective Delete() method.
func (s *ServiceLevelObjectiveServer) DeleteMonitoringBetaServiceLevelObjective(ctx context.Context, request *betapb.DeleteMonitoringBetaServiceLevelObjectiveRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServiceLevelObjective(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServiceLevelObjective(ctx, ProtoToServiceLevelObjective(request.GetResource()))

}

// ListMonitoringBetaServiceLevelObjective handles the gRPC request by passing it to the underlying ServiceLevelObjectiveList() method.
func (s *ServiceLevelObjectiveServer) ListMonitoringBetaServiceLevelObjective(ctx context.Context, request *betapb.ListMonitoringBetaServiceLevelObjectiveRequest) (*betapb.ListMonitoringBetaServiceLevelObjectiveResponse, error) {
	cl, err := createConfigServiceLevelObjective(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServiceLevelObjective(ctx, request.GetProject(), request.GetService())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.MonitoringBetaServiceLevelObjective
	for _, r := range resources.Items {
		rp := ServiceLevelObjectiveToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListMonitoringBetaServiceLevelObjectiveResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigServiceLevelObjective(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
