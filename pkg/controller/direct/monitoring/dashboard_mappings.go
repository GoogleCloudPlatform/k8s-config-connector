// Copyright 2024 Google LLC
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
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	monitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Empty_FromProto(mapCtx *direct.MapContext, in *emptypb.Empty) *krm.Empty {
	if in == nil {
		return nil
	}
	out := &krm.Empty{}
	return out
}
func Empty_ToProto(mapCtx *direct.MapContext, in *krm.Empty) *emptypb.Empty {
	if in == nil {
		return nil
	}
	out := &emptypb.Empty{}
	return out
}

func AlertChart_FromProto(mapCtx *direct.MapContext, in *pb.AlertChart) *krm.AlertChart {
	if in == nil {
		return nil
	}
	out := &krm.AlertChart{}
	if in.Name != "" {
		out.AlertPolicyRef = &monitoringv1beta1.MonitoringAlertPolicyRef{
			External: in.Name,
		}
	}
	return out
}
func AlertChart_ToProto(mapCtx *direct.MapContext, in *krm.AlertChart) *pb.AlertChart {
	if in == nil {
		return nil
	}
	out := &pb.AlertChart{}
	if in.AlertPolicyRef != nil {
		out.Name = in.AlertPolicyRef.External
	}
	return out
}

func Aggregation_AlignmentPeriod_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.SecondsString_FromProto(mapCtx, in)
}

func Aggregation_AlignmentPeriod_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.SecondsString_ToProto(mapCtx, in, "alignmentPeriod")
}

// TODO: The format is not documented, we need to validate
func DataSet_MinAlignmentPeriod_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.SecondsString_FromProto(mapCtx, in)
}

// TODO: The format is not documented, we need to validate
func DataSet_MinAlignmentPeriod_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.SecondsString_ToProto(mapCtx, in, "minAlignmentPeriod")
}

// TODO: The format is not documented, we need to validate
func SparkChartView_MinAlignmentPeriod_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.SecondsString_FromProto(mapCtx, in)
}

// TODO: The format is not documented, we need to validate
func SparkChartView_MinAlignmentPeriod_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.SecondsString_ToProto(mapCtx, in, "minAlignmentPeriod")
}

// TODO: The format is not documented, we need to validate
func XyChart_TimeshiftDuration_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.Duration_FromProto(mapCtx, in)
}

// TODO: The format is not documented, we need to validate
func XyChart_TimeshiftDuration_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.Duration_ToProto(mapCtx, in)
}

func PieChartDataSet_MinAlignmentPeriod_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.Duration_FromProto(mapCtx, in)
}

func PieChartDataSet_MinAlignmentPeriod_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.Duration_ToProto(mapCtx, in)
}

func TableDataSet_MinAlignmentPeriod_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.Duration_FromProto(mapCtx, in)
}

func TableDataSet_MinAlignmentPeriod_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.Duration_ToProto(mapCtx, in)
}

func TimeSeriesQuery_TimeSeriesQueryLanguage_ToProto(mapCtx *direct.MapContext, in *string) *pb.TimeSeriesQuery_TimeSeriesQueryLanguage {
	if in == nil {
		return nil
	}

	return &pb.TimeSeriesQuery_TimeSeriesQueryLanguage{
		TimeSeriesQueryLanguage: direct.ValueOf(in),
	}
}

func LogsPanel_ResourceNames_FromProto(mapCtx *direct.MapContext, in []string) []v1alpha1.ResourceRef {
	if in == nil {
		return nil
	}
	var out []v1alpha1.ResourceRef
	for _, v := range in {
		tokens := strings.Split(v, "/")
		if len(tokens) == 2 && tokens[0] == "projects" {
			out = append(out, v1alpha1.ResourceRef{
				Kind:     "Project",
				External: v,
			})
		} else {
			mapCtx.Errorf("resourceName %q was not recognized", v)
		}
	}
	return out
}

func LogsPanel_ResourceNames_ToProto(mapCtx *direct.MapContext, in []v1alpha1.ResourceRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, ref := range in {
		if ref.External == "" {
			mapCtx.Errorf("reference was not pre-resolved")
		}
		out = append(out, ref.External)
	}
	return out
}

func DashboardTimeSeriesQuery_TimeSeriesQueryLanguage_ToProto(mapCtx *direct.MapContext, in *string) *pb.TimeSeriesQuery_TimeSeriesQueryLanguage {
	if in == nil {
		return nil
	}
	return &pb.TimeSeriesQuery_TimeSeriesQueryLanguage{
		TimeSeriesQueryLanguage: *in,
	}
}

func ErrorReportingPanel_FromProto(mapCtx *direct.MapContext, in *pb.ErrorReportingPanel) *krm.ErrorReportingPanel {
	if in == nil {
		return nil
	}
	out := &krm.ErrorReportingPanel{}
	for _, projectName := range in.ProjectNames {
		out.ProjectRefs = append(out.ProjectRefs, refs.ProjectRef{
			External: projectName,
		})
	}
	out.Services = in.Services
	out.Versions = in.Versions
	return out
}

func ErrorReportingPanel_ToProto(mapCtx *direct.MapContext, in *krm.ErrorReportingPanel) *pb.ErrorReportingPanel {
	if in == nil {
		return nil
	}
	out := &pb.ErrorReportingPanel{}
	for _, projectRef := range in.ProjectRefs {
		out.ProjectNames = append(out.ProjectNames, projectRef.External)
	}
	out.Services = in.Services
	out.Versions = in.Versions
	return out
}

func TimeSeriesQuery_PrometheusQuery_ToProto(mapCtx *direct.MapContext, in *string) *pb.TimeSeriesQuery_PrometheusQuery {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesQuery_PrometheusQuery{}
	out.PrometheusQuery = *in
	return out
}

func TimeSeriesTable_ColumnSettings_FromProto(mapCtx *direct.MapContext, in *pb.TimeSeriesTable_ColumnSettings) *krm.TimeSeriesTable_ColumnSettings {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesTable_ColumnSettings{}
	out.Column = direct.LazyPtr(in.GetColumn())

	// We want to always output the visible field, i.e. `visible: false`
	// We probably can automate this, because the visible field is required.
	out.Visible = direct.PtrTo(in.GetVisible())
	return out
}

func MonitoredResource_FromProto(mapCtx *direct.MapContext, in *monitoredres.MonitoredResource) *krm.MonitoredResource {
	if in == nil {
		return nil
	}
	out := &krm.MonitoredResource{}
	out.Type = direct.LazyPtr(in.GetType())
	out.Labels = in.Labels
	return out
}
func MonitoredResource_ToProto(mapCtx *direct.MapContext, in *krm.MonitoredResource) *monitoredres.MonitoredResource {
	if in == nil {
		return nil
	}
	out := &monitoredres.MonitoredResource{}
	out.Type = direct.ValueOf(in.Type)
	out.Labels = in.Labels
	return out
}

func DashboardFilter_StringValue_ToProto(mapCtx *direct.MapContext, in *string) *pb.DashboardFilter_StringValue {
	if in == nil {
		return nil
	}
	out := &pb.DashboardFilter_StringValue{}
	out.StringValue = *in
	return out
}

func IncidentList_FromProto(mapCtx *direct.MapContext, in *pb.IncidentList) *krm.IncidentList {
	if in == nil {
		return nil
	}
	out := &krm.IncidentList{}
	out.MonitoredResources = direct.Slice_FromProto(mapCtx, in.MonitoredResources, MonitoredResource_FromProto)
	for _, policyName := range in.PolicyNames {
		out.PolicyRefs = append(out.PolicyRefs, monitoringv1beta1.MonitoringAlertPolicyRef{External: policyName})
	}
	return out
}
func IncidentList_ToProto(mapCtx *direct.MapContext, in *krm.IncidentList) *pb.IncidentList {
	if in == nil {
		return nil
	}
	out := &pb.IncidentList{}
	out.MonitoredResources = direct.Slice_ToProto(mapCtx, in.MonitoredResources, MonitoredResource_ToProto)
	for _, policyRef := range in.PolicyRefs {
		out.PolicyNames = append(out.PolicyNames, policyRef.External)
	}
	return out
}

func BlankView_FromProto(mapCtx *direct.MapContext, in *emptypb.Empty) *krm.BlankView {
	if in == nil {
		return nil
	}
	out := &krm.BlankView{}
	return out
}
func BlankView_ToProto(mapCtx *direct.MapContext, in *krm.BlankView) *emptypb.Empty {
	if in == nil {
		return nil
	}
	out := &emptypb.Empty{}
	return out
}
