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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/alpha/monitoring_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/alpha"
)

// ServiceLevelObjectiveServer implements the gRPC interface for ServiceLevelObjective.
type ServiceLevelObjectiveServer struct{}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum enum from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(e alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum) *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum_name[int32(e)]; ok {
		e := alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(n[len("MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum enum from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(e alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum) *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum_name[int32(e)]; ok {
		e := alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(n[len("MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum enum from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(e alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum_name[int32(e)]; ok {
		e := alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(n[len("MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum enum from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(e alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum_name[int32(e)]; ok {
		e := alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(n[len("MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveCalendarPeriodEnum converts a ServiceLevelObjectiveCalendarPeriodEnum enum from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveCalendarPeriodEnum(e alphapb.MonitoringAlphaServiceLevelObjectiveCalendarPeriodEnum) *alpha.ServiceLevelObjectiveCalendarPeriodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaServiceLevelObjectiveCalendarPeriodEnum_name[int32(e)]; ok {
		e := alpha.ServiceLevelObjectiveCalendarPeriodEnum(n[len("MonitoringAlphaServiceLevelObjectiveCalendarPeriodEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLevelObjectiveServiceLevelIndicator converts a ServiceLevelObjectiveServiceLevelIndicator object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicator(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicator) *alpha.ServiceLevelObjectiveServiceLevelIndicator {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicator{
		BasicSli:     ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSli(p.GetBasicSli()),
		RequestBased: ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBased(p.GetRequestBased()),
		WindowsBased: ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBased(p.GetWindowsBased()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSli converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSli object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSli(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSli) *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSli {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSli{
		Availability:          ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability(p.GetAvailability()),
		Latency:               ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency(p.GetLatency()),
		OperationAvailability: ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability(p.GetOperationAvailability()),
		OperationLatency:      ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency(p.GetOperationLatency()),
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
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability) *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability{}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency) *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency{
		Threshold:  dcl.StringOrNil(p.GetThreshold()),
		Experience: ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(p.GetExperience()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability) *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability{}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency) *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency{
		Threshold:  dcl.StringOrNil(p.GetThreshold()),
		Experience: ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(p.GetExperience()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorRequestBased converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBased object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBased(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBased) *alpha.ServiceLevelObjectiveServiceLevelIndicatorRequestBased {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorRequestBased{
		GoodTotalRatio:  ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio(p.GetGoodTotalRatio()),
		DistributionCut: ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut(p.GetDistributionCut()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio) *alpha.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio{
		GoodServiceFilter:  dcl.StringOrNil(p.GetGoodServiceFilter()),
		BadServiceFilter:   dcl.StringOrNil(p.GetBadServiceFilter()),
		TotalServiceFilter: dcl.StringOrNil(p.GetTotalServiceFilter()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut) *alpha.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut{
		DistributionFilter: dcl.StringOrNil(p.GetDistributionFilter()),
		Range:              ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange(p.GetRange()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange) *alpha.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange{
		Min: dcl.Float64OrNil(p.GetMin()),
		Max: dcl.Float64OrNil(p.GetMax()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBased converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBased object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBased(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBased) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBased {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBased{
		GoodBadMetricFilter:     dcl.StringOrNil(p.GetGoodBadMetricFilter()),
		GoodTotalRatioThreshold: ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold(p.GetGoodTotalRatioThreshold()),
		MetricMeanInRange:       ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange(p.GetMetricMeanInRange()),
		MetricSumInRange:        ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange(p.GetMetricSumInRange()),
		WindowPeriod:            dcl.StringOrNil(p.GetWindowPeriod()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold{
		Performance:         ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance(p.GetPerformance()),
		BasicSliPerformance: ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance(p.GetBasicSliPerformance()),
		Threshold:           dcl.Float64OrNil(p.GetThreshold()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance{
		GoodTotalRatio:  ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio(p.GetGoodTotalRatio()),
		DistributionCut: ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut(p.GetDistributionCut()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio{
		GoodServiceFilter:  dcl.StringOrNil(p.GetGoodServiceFilter()),
		BadServiceFilter:   dcl.StringOrNil(p.GetBadServiceFilter()),
		TotalServiceFilter: dcl.StringOrNil(p.GetTotalServiceFilter()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut{
		DistributionFilter: dcl.StringOrNil(p.GetDistributionFilter()),
		Range:              ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange(p.GetRange()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange{
		Min: dcl.Float64OrNil(p.GetMin()),
		Max: dcl.Float64OrNil(p.GetMax()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance{
		Availability:          ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability(p.GetAvailability()),
		Latency:               ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency(p.GetLatency()),
		OperationAvailability: ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability(p.GetOperationAvailability()),
		OperationLatency:      ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency(p.GetOperationLatency()),
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
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability{}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency{
		Threshold:  dcl.StringOrNil(p.GetThreshold()),
		Experience: ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(p.GetExperience()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability{}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency{
		Threshold:  dcl.StringOrNil(p.GetThreshold()),
		Experience: ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(p.GetExperience()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange{
		TimeSeries: dcl.StringOrNil(p.GetTimeSeries()),
		Range:      ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange(p.GetRange()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange{
		Min: dcl.Float64OrNil(p.GetMin()),
		Max: dcl.Float64OrNil(p.GetMax()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange{
		TimeSeries: dcl.StringOrNil(p.GetTimeSeries()),
		Range:      ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange(p.GetRange()),
	}
	return obj
}

// ProtoToServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange object from its proto representation.
func ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange(p *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange) *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange{
		Min: dcl.Float64OrNil(p.GetMin()),
		Max: dcl.Float64OrNil(p.GetMax()),
	}
	return obj
}

// ProtoToServiceLevelObjective converts a ServiceLevelObjective resource from its proto representation.
func ProtoToServiceLevelObjective(p *alphapb.MonitoringAlphaServiceLevelObjective) *alpha.ServiceLevelObjective {
	obj := &alpha.ServiceLevelObjective{
		Name:                   dcl.StringOrNil(p.GetName()),
		DisplayName:            dcl.StringOrNil(p.GetDisplayName()),
		ServiceLevelIndicator:  ProtoToMonitoringAlphaServiceLevelObjectiveServiceLevelIndicator(p.GetServiceLevelIndicator()),
		Goal:                   dcl.Float64OrNil(p.GetGoal()),
		RollingPeriod:          dcl.StringOrNil(p.GetRollingPeriod()),
		CalendarPeriod:         ProtoToMonitoringAlphaServiceLevelObjectiveCalendarPeriodEnum(p.GetCalendarPeriod()),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		DeleteTime:             dcl.StringOrNil(p.GetDeleteTime()),
		ServiceManagementOwned: dcl.Bool(p.GetServiceManagementOwned()),
		Project:                dcl.StringOrNil(p.GetProject()),
		Service:                dcl.StringOrNil(p.GetService()),
	}
	return obj
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnumToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum enum to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnumToProto(e *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum) alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum {
	if e == nil {
		return alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum_value["ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(v)
	}
	return alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(0)
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnumToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum enum to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnumToProto(e *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum) alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum {
	if e == nil {
		return alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum_value["ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(v)
	}
	return alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(0)
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnumToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum enum to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnumToProto(e *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum) alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum {
	if e == nil {
		return alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum_value["ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(v)
	}
	return alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(0)
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnumToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum enum to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnumToProto(e *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum) alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum {
	if e == nil {
		return alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum_value["ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(v)
	}
	return alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(0)
}

// ServiceLevelObjectiveCalendarPeriodEnumToProto converts a ServiceLevelObjectiveCalendarPeriodEnum enum to its proto representation.
func MonitoringAlphaServiceLevelObjectiveCalendarPeriodEnumToProto(e *alpha.ServiceLevelObjectiveCalendarPeriodEnum) alphapb.MonitoringAlphaServiceLevelObjectiveCalendarPeriodEnum {
	if e == nil {
		return alphapb.MonitoringAlphaServiceLevelObjectiveCalendarPeriodEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaServiceLevelObjectiveCalendarPeriodEnum_value["ServiceLevelObjectiveCalendarPeriodEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaServiceLevelObjectiveCalendarPeriodEnum(v)
	}
	return alphapb.MonitoringAlphaServiceLevelObjectiveCalendarPeriodEnum(0)
}

// ServiceLevelObjectiveServiceLevelIndicatorToProto converts a ServiceLevelObjectiveServiceLevelIndicator object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicator) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicator {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicator{}
	p.SetBasicSli(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliToProto(o.BasicSli))
	p.SetRequestBased(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedToProto(o.RequestBased))
	p.SetWindowsBased(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedToProto(o.WindowsBased))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSli object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSli) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSli {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSli{}
	p.SetAvailability(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailabilityToProto(o.Availability))
	p.SetLatency(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyToProto(o.Latency))
	p.SetOperationAvailability(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailabilityToProto(o.OperationAvailability))
	p.SetOperationLatency(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyToProto(o.OperationLatency))
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
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailabilityToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability{}
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency{}
	p.SetThreshold(dcl.ValueOrEmptyString(o.Threshold))
	p.SetExperience(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnumToProto(o.Experience))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailabilityToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailabilityToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability{}
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyToProto converts a ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency{}
	p.SetThreshold(dcl.ValueOrEmptyString(o.Threshold))
	p.SetExperience(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnumToProto(o.Experience))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorRequestBasedToProto converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBased object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorRequestBased) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBased {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBased{}
	p.SetGoodTotalRatio(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatioToProto(o.GoodTotalRatio))
	p.SetDistributionCut(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutToProto(o.DistributionCut))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatioToProto converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatioToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio{}
	p.SetGoodServiceFilter(dcl.ValueOrEmptyString(o.GoodServiceFilter))
	p.SetBadServiceFilter(dcl.ValueOrEmptyString(o.BadServiceFilter))
	p.SetTotalServiceFilter(dcl.ValueOrEmptyString(o.TotalServiceFilter))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutToProto converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut{}
	p.SetDistributionFilter(dcl.ValueOrEmptyString(o.DistributionFilter))
	p.SetRange(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRangeToProto(o.Range))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRangeToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange{}
	p.SetMin(dcl.ValueOrEmptyDouble(o.Min))
	p.SetMax(dcl.ValueOrEmptyDouble(o.Max))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBased object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBased) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBased {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBased{}
	p.SetGoodBadMetricFilter(dcl.ValueOrEmptyString(o.GoodBadMetricFilter))
	p.SetGoodTotalRatioThreshold(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdToProto(o.GoodTotalRatioThreshold))
	p.SetMetricMeanInRange(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeToProto(o.MetricMeanInRange))
	p.SetMetricSumInRange(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeToProto(o.MetricSumInRange))
	p.SetWindowPeriod(dcl.ValueOrEmptyString(o.WindowPeriod))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold{}
	p.SetPerformance(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceToProto(o.Performance))
	p.SetBasicSliPerformance(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceToProto(o.BasicSliPerformance))
	p.SetThreshold(dcl.ValueOrEmptyDouble(o.Threshold))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance{}
	p.SetGoodTotalRatio(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatioToProto(o.GoodTotalRatio))
	p.SetDistributionCut(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutToProto(o.DistributionCut))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatioToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatioToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio{}
	p.SetGoodServiceFilter(dcl.ValueOrEmptyString(o.GoodServiceFilter))
	p.SetBadServiceFilter(dcl.ValueOrEmptyString(o.BadServiceFilter))
	p.SetTotalServiceFilter(dcl.ValueOrEmptyString(o.TotalServiceFilter))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut{}
	p.SetDistributionFilter(dcl.ValueOrEmptyString(o.DistributionFilter))
	p.SetRange(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRangeToProto(o.Range))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRangeToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange{}
	p.SetMin(dcl.ValueOrEmptyDouble(o.Min))
	p.SetMax(dcl.ValueOrEmptyDouble(o.Max))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance{}
	p.SetAvailability(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailabilityToProto(o.Availability))
	p.SetLatency(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyToProto(o.Latency))
	p.SetOperationAvailability(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailabilityToProto(o.OperationAvailability))
	p.SetOperationLatency(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyToProto(o.OperationLatency))
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
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailabilityToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability{}
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency{}
	p.SetThreshold(dcl.ValueOrEmptyString(o.Threshold))
	p.SetExperience(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnumToProto(o.Experience))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailabilityToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailabilityToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability{}
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency{}
	p.SetThreshold(dcl.ValueOrEmptyString(o.Threshold))
	p.SetExperience(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnumToProto(o.Experience))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange{}
	p.SetTimeSeries(dcl.ValueOrEmptyString(o.TimeSeries))
	p.SetRange(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRangeToProto(o.Range))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRangeToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange{}
	p.SetMin(dcl.ValueOrEmptyDouble(o.Min))
	p.SetMax(dcl.ValueOrEmptyDouble(o.Max))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange{}
	p.SetTimeSeries(dcl.ValueOrEmptyString(o.TimeSeries))
	p.SetRange(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRangeToProto(o.Range))
	return p
}

// ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRangeToProto converts a ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange object to its proto representation.
func MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRangeToProto(o *alpha.ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange) *alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange{}
	p.SetMin(dcl.ValueOrEmptyDouble(o.Min))
	p.SetMax(dcl.ValueOrEmptyDouble(o.Max))
	return p
}

// ServiceLevelObjectiveToProto converts a ServiceLevelObjective resource to its proto representation.
func ServiceLevelObjectiveToProto(resource *alpha.ServiceLevelObjective) *alphapb.MonitoringAlphaServiceLevelObjective {
	p := &alphapb.MonitoringAlphaServiceLevelObjective{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetServiceLevelIndicator(MonitoringAlphaServiceLevelObjectiveServiceLevelIndicatorToProto(resource.ServiceLevelIndicator))
	p.SetGoal(dcl.ValueOrEmptyDouble(resource.Goal))
	p.SetRollingPeriod(dcl.ValueOrEmptyString(resource.RollingPeriod))
	p.SetCalendarPeriod(MonitoringAlphaServiceLevelObjectiveCalendarPeriodEnumToProto(resource.CalendarPeriod))
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
func (s *ServiceLevelObjectiveServer) applyServiceLevelObjective(ctx context.Context, c *alpha.Client, request *alphapb.ApplyMonitoringAlphaServiceLevelObjectiveRequest) (*alphapb.MonitoringAlphaServiceLevelObjective, error) {
	p := ProtoToServiceLevelObjective(request.GetResource())
	res, err := c.ApplyServiceLevelObjective(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceLevelObjectiveToProto(res)
	return r, nil
}

// applyMonitoringAlphaServiceLevelObjective handles the gRPC request by passing it to the underlying ServiceLevelObjective Apply() method.
func (s *ServiceLevelObjectiveServer) ApplyMonitoringAlphaServiceLevelObjective(ctx context.Context, request *alphapb.ApplyMonitoringAlphaServiceLevelObjectiveRequest) (*alphapb.MonitoringAlphaServiceLevelObjective, error) {
	cl, err := createConfigServiceLevelObjective(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyServiceLevelObjective(ctx, cl, request)
}

// DeleteServiceLevelObjective handles the gRPC request by passing it to the underlying ServiceLevelObjective Delete() method.
func (s *ServiceLevelObjectiveServer) DeleteMonitoringAlphaServiceLevelObjective(ctx context.Context, request *alphapb.DeleteMonitoringAlphaServiceLevelObjectiveRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServiceLevelObjective(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServiceLevelObjective(ctx, ProtoToServiceLevelObjective(request.GetResource()))

}

// ListMonitoringAlphaServiceLevelObjective handles the gRPC request by passing it to the underlying ServiceLevelObjectiveList() method.
func (s *ServiceLevelObjectiveServer) ListMonitoringAlphaServiceLevelObjective(ctx context.Context, request *alphapb.ListMonitoringAlphaServiceLevelObjectiveRequest) (*alphapb.ListMonitoringAlphaServiceLevelObjectiveResponse, error) {
	cl, err := createConfigServiceLevelObjective(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServiceLevelObjective(ctx, request.GetProject(), request.GetService())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.MonitoringAlphaServiceLevelObjective
	for _, r := range resources.Items {
		rp := ServiceLevelObjectiveToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListMonitoringAlphaServiceLevelObjectiveResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigServiceLevelObjective(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
