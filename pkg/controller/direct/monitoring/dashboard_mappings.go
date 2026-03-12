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
	"fmt"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"
	"google.golang.org/genproto/googleapis/type/interval"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func MonitoringDashboardSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringDashboardSpec) *pb.Dashboard {
	if in == nil {
		return nil
	}
	out := &pb.Dashboard{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if in.GridLayout != nil {
		out.Layout = &pb.Dashboard_GridLayout{GridLayout: GridLayout_ToProto(mapCtx, in.GridLayout)}
	}
	if in.MosaicLayout != nil {
		out.Layout = &pb.Dashboard_MosaicLayout{MosaicLayout: MosaicLayout_ToProto(mapCtx, in.MosaicLayout)}
	}
	if in.RowLayout != nil {
		out.Layout = &pb.Dashboard_RowLayout{RowLayout: RowLayout_ToProto(mapCtx, in.RowLayout)}
	}
	if in.ColumnLayout != nil {
		out.Layout = &pb.Dashboard_ColumnLayout{ColumnLayout: ColumnLayout_ToProto(mapCtx, in.ColumnLayout)}
	}
	out.DashboardFilters = direct.Slice_ToProto(mapCtx, in.DashboardFilters, DashboardFilter_ToProto)
	out.Labels = in.Labels
	return out
}

func MonitoringDashboardSpec_FromProto(mapCtx *direct.MapContext, in *pb.Dashboard) *krm.MonitoringDashboardSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringDashboardSpec{}
	out.DisplayName = direct.LazyPtr(in.DisplayName)
	if in.GetGridLayout() != nil {
		out.GridLayout = GridLayout_FromProto(mapCtx, in.GetGridLayout())
	}
	if in.GetMosaicLayout() != nil {
		out.MosaicLayout = MosaicLayout_FromProto(mapCtx, in.GetMosaicLayout())
	}
	if in.GetRowLayout() != nil {
		out.RowLayout = RowLayout_FromProto(mapCtx, in.GetRowLayout())
	}
	if in.GetColumnLayout() != nil {
		out.ColumnLayout = ColumnLayout_FromProto(mapCtx, in.GetColumnLayout())
	}
	out.DashboardFilters = direct.Slice_FromProto(mapCtx, in.DashboardFilters, DashboardFilter_FromProto)
	out.Labels = in.Labels
	return out
}

func PickTimeSeriesFilter_ToProto(mapCtx *direct.MapContext, in *krm.PickTimeSeriesFilter) *pb.PickTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &pb.PickTimeSeriesFilter{}
	out.RankingMethod = direct.Enum_ToProto[pb.PickTimeSeriesFilter_Method](mapCtx, in.RankingMethod)
	out.NumTimeSeries = direct.ValueOf(in.NumTimeSeries)
	out.Direction = direct.Enum_ToProto[pb.PickTimeSeriesFilter_Direction](mapCtx, in.Direction)
	out.Interval = Interval_ToProto(mapCtx, in.Interval)
	return out
}

func PickTimeSeriesFilter_FromProto(mapCtx *direct.MapContext, in *pb.PickTimeSeriesFilter) *krm.PickTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &krm.PickTimeSeriesFilter{}
	out.RankingMethod = direct.Enum_FromProto(mapCtx, in.RankingMethod)
	out.NumTimeSeries = direct.LazyPtr(in.NumTimeSeries)
	out.Direction = direct.Enum_FromProto(mapCtx, in.Direction)
	out.Interval = Interval_FromProto(mapCtx, in.Interval)
	return out
}

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
		out.AlertPolicyRef = &krm.MonitoringAlertPolicyRef{
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

func Aggregation_FromProto(mapCtx *direct.MapContext, in *pb.Aggregation) *krm.Aggregation {
	if in == nil {
		return nil
	}
	out := &krm.Aggregation{}
	out.AlignmentPeriod = direct.SecondsString_FromProto(mapCtx, in.GetAlignmentPeriod())
	out.PerSeriesAligner = direct.Enum_FromProto(mapCtx, in.GetPerSeriesAligner())
	out.CrossSeriesReducer = direct.Enum_FromProto(mapCtx, in.GetCrossSeriesReducer())
	out.GroupByFields = in.GroupByFields
	return out
}

func Aggregation_ToProto(mapCtx *direct.MapContext, in *krm.Aggregation) *pb.Aggregation {
	if in == nil {
		return nil
	}
	out := &pb.Aggregation{}
	out.AlignmentPeriod = direct.SecondsString_ToProto(mapCtx, in.AlignmentPeriod, "alignmentPeriod")
	out.PerSeriesAligner = direct.Enum_ToProto[pb.Aggregation_Aligner](mapCtx, in.PerSeriesAligner)
	out.CrossSeriesReducer = direct.Enum_ToProto[pb.Aggregation_Reducer](mapCtx, in.CrossSeriesReducer)
	out.GroupByFields = in.GroupByFields
	return out
}

func Aggregation_AlignmentPeriod_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.SecondsString_FromProto(mapCtx, in)
}

func Aggregation_AlignmentPeriod_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.SecondsString_ToProto(mapCtx, in, "alignmentPeriod")
}

// TODO: The format is not documented, we need to validate
func XyChart_DataSet_FromProto(mapCtx *direct.MapContext, in *pb.XyChart_DataSet) *krm.XyChart_DataSet {
	if in == nil {
		return nil
	}
	out := &krm.XyChart_DataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(mapCtx, in.GetTimeSeriesQuery())
	out.PlotType = direct.Enum_FromProto(mapCtx, in.GetPlotType())
	out.LegendTemplate = direct.LazyPtr(in.GetLegendTemplate())
	out.MinAlignmentPeriod = direct.SecondsString_FromProto(mapCtx, in.GetMinAlignmentPeriod())
	out.TargetAxis = direct.Enum_FromProto(mapCtx, in.GetTargetAxis())
	return out
}

func XyChart_DataSet_ToProto(mapCtx *direct.MapContext, in *krm.XyChart_DataSet) *pb.XyChart_DataSet {
	if in == nil {
		return nil
	}
	out := &pb.XyChart_DataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_ToProto(mapCtx, in.TimeSeriesQuery)
	out.PlotType = direct.Enum_ToProto[pb.XyChart_DataSet_PlotType](mapCtx, in.PlotType)
	out.LegendTemplate = direct.ValueOf(in.LegendTemplate)
	out.MinAlignmentPeriod = direct.SecondsString_ToProto(mapCtx, in.MinAlignmentPeriod, "minAlignmentPeriod")
	out.TargetAxis = direct.Enum_ToProto[pb.XyChart_DataSet_TargetAxis](mapCtx, in.TargetAxis)
	return out
}

func DataSet_MinAlignmentPeriod_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.SecondsString_FromProto(mapCtx, in)
}

// TODO: The format is not documented, we need to validate
func DataSet_MinAlignmentPeriod_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.SecondsString_ToProto(mapCtx, in, "minAlignmentPeriod")
}

// TODO: The format is not documented, we need to validate
func Scorecard_SparkChartView_FromProto(mapCtx *direct.MapContext, in *pb.Scorecard_SparkChartView) *krm.Scorecard_SparkChartView {
	if in == nil {
		return nil
	}
	out := &krm.Scorecard_SparkChartView{}
	out.SparkChartType = direct.Enum_FromProto(mapCtx, in.GetSparkChartType())
	out.MinAlignmentPeriod = direct.SecondsString_FromProto(mapCtx, in.GetMinAlignmentPeriod())
	return out
}

func Scorecard_SparkChartView_ToProto(mapCtx *direct.MapContext, in *krm.Scorecard_SparkChartView) *pb.Scorecard_SparkChartView {
	if in == nil {
		return nil
	}
	out := &pb.Scorecard_SparkChartView{}
	out.SparkChartType = direct.Enum_ToProto[pb.SparkChartType](mapCtx, in.SparkChartType)
	out.MinAlignmentPeriod = direct.SecondsString_ToProto(mapCtx, in.MinAlignmentPeriod, "minAlignmentPeriod")
	return out
}

func SparkChartView_MinAlignmentPeriod_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.SecondsString_FromProto(mapCtx, in)
}

// TODO: The format is not documented, we need to validate
func SparkChartView_MinAlignmentPeriod_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.SecondsString_ToProto(mapCtx, in, "minAlignmentPeriod")
}

// TODO: The format is not documented, we need to validate
func XyChart_FromProto(mapCtx *direct.MapContext, in *pb.XyChart) *krm.XyChart {
	if in == nil {
		return nil
	}
	out := &krm.XyChart{}
	out.DataSets = direct.Slice_FromProto(mapCtx, in.DataSets, XyChart_DataSet_FromProto)
	out.TimeshiftDuration = direct.Duration_FromProto(mapCtx, in.GetTimeshiftDuration())
	out.Thresholds = direct.Slice_FromProto(mapCtx, in.Thresholds, Threshold_FromProto)
	out.XAxis = XyChart_Axis_FromProto(mapCtx, in.GetXAxis())
	out.YAxis = XyChart_Axis_FromProto(mapCtx, in.GetYAxis())
	out.Y2Axis = XyChart_Axis_FromProto(mapCtx, in.GetY2Axis())
	out.ChartOptions = ChartOptions_FromProto(mapCtx, in.GetChartOptions())
	return out
}

func XyChart_ToProto(mapCtx *direct.MapContext, in *krm.XyChart) *pb.XyChart {
	if in == nil {
		return nil
	}
	out := &pb.XyChart{}
	out.DataSets = direct.Slice_ToProto(mapCtx, in.DataSets, XyChart_DataSet_ToProto)
	out.TimeshiftDuration = direct.Duration_ToProto(mapCtx, in.TimeshiftDuration)
	out.Thresholds = direct.Slice_ToProto(mapCtx, in.Thresholds, Threshold_ToProto)
	out.XAxis = XyChart_Axis_ToProto(mapCtx, in.XAxis)
	out.YAxis = XyChart_Axis_ToProto(mapCtx, in.YAxis)
	out.Y2Axis = XyChart_Axis_ToProto(mapCtx, in.Y2Axis)
	out.ChartOptions = ChartOptions_ToProto(mapCtx, in.ChartOptions)
	return out
}

func XyChart_TimeshiftDuration_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.Duration_FromProto(mapCtx, in)
}

// TODO: The format is not documented, we need to validate
func XyChart_TimeshiftDuration_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.Duration_ToProto(mapCtx, in)
}

func PieChart_PieChartDataSet_FromProto(mapCtx *direct.MapContext, in *pb.PieChart_PieChartDataSet) *krm.PieChart_PieChartDataSet {
	if in == nil {
		return nil
	}
	out := &krm.PieChart_PieChartDataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(mapCtx, in.GetTimeSeriesQuery())
	out.SliceNameTemplate = direct.LazyPtr(in.GetSliceNameTemplate())
	out.MinAlignmentPeriod = direct.Duration_FromProto(mapCtx, in.GetMinAlignmentPeriod())
	return out
}

func PieChart_PieChartDataSet_ToProto(mapCtx *direct.MapContext, in *krm.PieChart_PieChartDataSet) *pb.PieChart_PieChartDataSet {
	if in == nil {
		return nil
	}
	out := &pb.PieChart_PieChartDataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_ToProto(mapCtx, in.TimeSeriesQuery)
	out.SliceNameTemplate = direct.ValueOf(in.SliceNameTemplate)
	out.MinAlignmentPeriod = direct.Duration_ToProto(mapCtx, in.MinAlignmentPeriod)
	return out
}

func PieChartDataSet_MinAlignmentPeriod_FromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	return direct.Duration_FromProto(mapCtx, in)
}

func PieChartDataSet_MinAlignmentPeriod_ToProto(mapCtx *direct.MapContext, in *string) *durationpb.Duration {
	return direct.Duration_ToProto(mapCtx, in)
}

func TimeSeriesTable_TableDataSet_FromProto(mapCtx *direct.MapContext, in *pb.TimeSeriesTable_TableDataSet) *krm.TimeSeriesTable_TableDataSet {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesTable_TableDataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(mapCtx, in.GetTimeSeriesQuery())
	out.TableTemplate = direct.LazyPtr(in.GetTableTemplate())
	out.MinAlignmentPeriod = direct.Duration_FromProto(mapCtx, in.GetMinAlignmentPeriod())
	out.TableDisplayOptions = TableDisplayOptions_FromProto(mapCtx, in.GetTableDisplayOptions())
	return out
}

func TimeSeriesTable_TableDataSet_ToProto(mapCtx *direct.MapContext, in *krm.TimeSeriesTable_TableDataSet) *pb.TimeSeriesTable_TableDataSet {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesTable_TableDataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_ToProto(mapCtx, in.TimeSeriesQuery)
	out.TableTemplate = direct.ValueOf(in.TableTemplate)
	out.MinAlignmentPeriod = direct.Duration_ToProto(mapCtx, in.MinAlignmentPeriod)
	out.TableDisplayOptions = TableDisplayOptions_ToProto(mapCtx, in.TableDisplayOptions)
	return out
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

func LogsPanel_ResourceNames_FromProto(mapCtx *direct.MapContext, in []string) []krm.LogsPanelResourceRef {
	if in == nil {
		return nil
	}
	var out []krm.LogsPanelResourceRef
	for _, v := range in {
		tokens := strings.Split(v, "/")
		if len(tokens) == 2 && tokens[0] == "projects" {
			out = append(out, krm.LogsPanelResourceRef{
				Kind:     "Project",
				External: v,
			})
		} else {
			mapCtx.Errorf("resourceName %q was not recognized", v)
		}
	}
	return out
}

func LogsPanel_ResourceNames_ToProto(mapCtx *direct.MapContext, in []krm.LogsPanelResourceRef) []string {
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
	out.StringValue = direct.ValueOf(in)
	return out
}

func Interval_FromProto(mapCtx *direct.MapContext, in *interval.Interval) *krm.Interval {
	if in == nil {
		return nil
	}
	out := &krm.Interval{}
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}

func Interval_ToProto(mapCtx *direct.MapContext, in *krm.Interval) *interval.Interval {
	if in == nil {
		return nil
	}
	out := &interval.Interval{}
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}

func IncidentList_FromProto(mapCtx *direct.MapContext, in *pb.IncidentList) *krm.IncidentList {
	if in == nil {
		return nil
	}
	out := &krm.IncidentList{}
	out.MonitoredResources = direct.Slice_FromProto(mapCtx, in.MonitoredResources, MonitoredResource_FromProto)
	for _, policyName := range in.PolicyNames {
		out.PolicyRefs = append(out.PolicyRefs, krm.MonitoringAlertPolicyRef{External: policyName})
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

func Widget_FromProto(mapCtx *direct.MapContext, in *pb.Widget) *krm.Widget {
	if in == nil {
		return nil
	}
	out := &krm.Widget{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.XyChart = XyChart_FromProto(mapCtx, in.GetXyChart())
	out.Scorecard = Scorecard_FromProto(mapCtx, in.GetScorecard())
	out.Text = Text_FromProto(mapCtx, in.GetText())
	out.Blank = Empty_FromProto(mapCtx, in.GetBlank())
	out.AlertChart = AlertChart_FromProto(mapCtx, in.GetAlertChart())
	out.TimeSeriesTable = TimeSeriesTable_FromProto(mapCtx, in.GetTimeSeriesTable())
	out.CollapsibleGroup = CollapsibleGroup_FromProto(mapCtx, in.GetCollapsibleGroup())
	out.LogsPanel = LogsPanel_FromProto(mapCtx, in.GetLogsPanel())
	out.IncidentList = IncidentList_FromProto(mapCtx, in.GetIncidentList())
	out.PieChart = PieChart_FromProto(mapCtx, in.GetPieChart())
	out.ErrorReportingPanel = ErrorReportingPanel_FromProto(mapCtx, in.GetErrorReportingPanel())
	out.SectionHeader = SectionHeader_FromProto(mapCtx, in.GetSectionHeader())
	out.SingleViewGroup = SingleViewGroup_FromProto(mapCtx, in.GetSingleViewGroup())
	out.Id = direct.LazyPtr(in.GetId())
	return out
}

func Widget_ToProto(mapCtx *direct.MapContext, in *krm.Widget) *pb.Widget {
	if in == nil {
		return nil
	}
	if in.Id != nil {
		fmt.Printf("DEBUG: Widget_ToProto with Id: %s\n", *in.Id)
	}
	out := &pb.Widget{}
	out.Title = direct.ValueOf(in.Title)
	if oneof := XyChart_ToProto(mapCtx, in.XyChart); oneof != nil {
		out.Content = &pb.Widget_XyChart{XyChart: oneof}
	}
	if oneof := Scorecard_ToProto(mapCtx, in.Scorecard); oneof != nil {
		out.Content = &pb.Widget_Scorecard{Scorecard: oneof}
	}
	if oneof := Text_ToProto(mapCtx, in.Text); oneof != nil {
		out.Content = &pb.Widget_Text{Text: oneof}
	}
	if oneof := Empty_ToProto(mapCtx, in.Blank); oneof != nil {
		out.Content = &pb.Widget_Blank{Blank: oneof}
	}
	if oneof := AlertChart_ToProto(mapCtx, in.AlertChart); oneof != nil {
		out.Content = &pb.Widget_AlertChart{AlertChart: oneof}
	}
	if oneof := TimeSeriesTable_ToProto(mapCtx, in.TimeSeriesTable); oneof != nil {
		out.Content = &pb.Widget_TimeSeriesTable{TimeSeriesTable: oneof}
	}
	if oneof := CollapsibleGroup_ToProto(mapCtx, in.CollapsibleGroup); oneof != nil {
		out.Content = &pb.Widget_CollapsibleGroup{CollapsibleGroup: oneof}
	}
	if oneof := LogsPanel_ToProto(mapCtx, in.LogsPanel); oneof != nil {
		out.Content = &pb.Widget_LogsPanel{LogsPanel: oneof}
	}
	if oneof := IncidentList_ToProto(mapCtx, in.IncidentList); oneof != nil {
		out.Content = &pb.Widget_IncidentList{IncidentList: oneof}
	}
	if oneof := PieChart_ToProto(mapCtx, in.PieChart); oneof != nil {
		out.Content = &pb.Widget_PieChart{PieChart: oneof}
	}
	if oneof := ErrorReportingPanel_ToProto(mapCtx, in.ErrorReportingPanel); oneof != nil {
		out.Content = &pb.Widget_ErrorReportingPanel{ErrorReportingPanel: oneof}
	}
	if oneof := SectionHeader_ToProto(mapCtx, in.SectionHeader); oneof != nil {
		out.Content = &pb.Widget_SectionHeader{SectionHeader: oneof}
	}
	if oneof := SingleViewGroup_ToProto(mapCtx, in.SingleViewGroup); oneof != nil {
		out.Content = &pb.Widget_SingleViewGroup{SingleViewGroup: oneof}
	}
	out.Id = direct.ValueOf(in.Id)
	return out
}
