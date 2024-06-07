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
	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
)

func Aggregation_FromProto(mapCtx *MapContext, in *pb.Aggregation) *krm.Aggregation {
	if in == nil {
		return nil
	}
	out := &krm.Aggregation{}
	out.AlignmentPeriod = Aggregation_AlignmentPeriod_FromProto(mapCtx, in.GetAlignmentPeriod())
	out.PerSeriesAligner = Enum_FromProto(mapCtx, in.PerSeriesAligner)
	out.CrossSeriesReducer = Enum_FromProto(mapCtx, in.CrossSeriesReducer)
	out.GroupByFields = in.GroupByFields
	return out
}
func Aggregation_ToProto(mapCtx *MapContext, in *krm.Aggregation) *pb.Aggregation {
	if in == nil {
		return nil
	}
	out := &pb.Aggregation{}
	out.AlignmentPeriod = Aggregation_AlignmentPeriod_ToProto(mapCtx, in.AlignmentPeriod)
	out.PerSeriesAligner = Enum_ToProto[pb.Aggregation_Aligner](mapCtx, in.PerSeriesAligner)
	out.CrossSeriesReducer = Enum_ToProto[pb.Aggregation_Reducer](mapCtx, in.CrossSeriesReducer)
	out.GroupByFields = in.GroupByFields
	return out
}
func AlertChart_FromProto(mapCtx *MapContext, in *pb.AlertChart) *krm.AlertChart {
	if in == nil {
		return nil
	}
	out := &krm.AlertChart{}
	out.Name = LazyPtr(in.GetName())
	return out
}
func AlertChart_ToProto(mapCtx *MapContext, in *krm.AlertChart) *pb.AlertChart {
	if in == nil {
		return nil
	}
	out := &pb.AlertChart{}
	out.Name = ValueOf(in.Name)
	return out
}
func ChartOptions_FromProto(mapCtx *MapContext, in *pb.ChartOptions) *krm.ChartOptions {
	if in == nil {
		return nil
	}
	out := &krm.ChartOptions{}
	out.Mode = Enum_FromProto(mapCtx, in.Mode)
	return out
}
func ChartOptions_ToProto(mapCtx *MapContext, in *krm.ChartOptions) *pb.ChartOptions {
	if in == nil {
		return nil
	}
	out := &pb.ChartOptions{}
	out.Mode = Enum_ToProto[pb.ChartOptions_Mode](mapCtx, in.Mode)
	return out
}
func CollapsibleGroup_FromProto(mapCtx *MapContext, in *pb.CollapsibleGroup) *krm.CollapsibleGroup {
	if in == nil {
		return nil
	}
	out := &krm.CollapsibleGroup{}
	out.Collapsed = LazyPtr(in.GetCollapsed())
	return out
}
func CollapsibleGroup_ToProto(mapCtx *MapContext, in *krm.CollapsibleGroup) *pb.CollapsibleGroup {
	if in == nil {
		return nil
	}
	out := &pb.CollapsibleGroup{}
	out.Collapsed = ValueOf(in.Collapsed)
	return out
}
func ColumnLayout_FromProto(mapCtx *MapContext, in *pb.ColumnLayout) *krm.ColumnLayout {
	if in == nil {
		return nil
	}
	out := &krm.ColumnLayout{}
	out.Columns = Slice_FromProto(mapCtx, in.Columns, ColumnLayout_Column_FromProto)
	return out
}
func ColumnLayout_ToProto(mapCtx *MapContext, in *krm.ColumnLayout) *pb.ColumnLayout {
	if in == nil {
		return nil
	}
	out := &pb.ColumnLayout{}
	out.Columns = Slice_ToProto(mapCtx, in.Columns, ColumnLayout_Column_ToProto)
	return out
}
func ColumnLayout_Column_FromProto(mapCtx *MapContext, in *pb.ColumnLayout_Column) *krm.ColumnLayout_Column {
	if in == nil {
		return nil
	}
	out := &krm.ColumnLayout_Column{}
	out.Weight = LazyPtr(in.GetWeight())
	out.Widgets = Slice_FromProto(mapCtx, in.Widgets, Widget_FromProto)
	return out
}
func ColumnLayout_Column_ToProto(mapCtx *MapContext, in *krm.ColumnLayout_Column) *pb.ColumnLayout_Column {
	if in == nil {
		return nil
	}
	out := &pb.ColumnLayout_Column{}
	out.Weight = ValueOf(in.Weight)
	out.Widgets = Slice_ToProto(mapCtx, in.Widgets, Widget_ToProto)
	return out
}

// func DashboardFilter_FromProto(mapCtx *MapContext, in *pb.DashboardFilter) *krm.DashboardFilter {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardFilter{}
// 	out.LabelKey = LazyPtr(in.GetLabelKey())
// 	out.TemplateVariable = LazyPtr(in.GetTemplateVariable())
// 	out.StringValue = LazyPtr(in.GetStringValue())
// 	out.FilterType = Enum_FromProto(mapCtx, in.FilterType)
// 	return out
// }
// func DashboardFilter_ToProto(mapCtx *MapContext, in *krm.DashboardFilter) *pb.DashboardFilter {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.DashboardFilter{}
// 	out.LabelKey = ValueOf(in.LabelKey)
// 	out.TemplateVariable = ValueOf(in.TemplateVariable)
// 	if oneof := DashboardFilter_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
// 		out.DefaultValue = oneof
// 	}
// 	out.FilterType = Enum_ToProto[pb.DashboardFilter_FilterType](mapCtx, in.FilterType)
// 	return out
// }

//	func Dashboard_LabelsEntry_FromProto(mapCtx *MapContext, in *pb.Dashboard_LabelsEntry) *krm.Dashboard_LabelsEntry {
//		if in == nil {
//			return nil
//		}
//		out := &krm.Dashboard_LabelsEntry{}
//		out.Key = LazyPtr(in.GetKey())
//		out.Value = LazyPtr(in.GetValue())
//		return out
//	}
//
//	func Dashboard_LabelsEntry_ToProto(mapCtx *MapContext, in *krm.Dashboard_LabelsEntry) *pb.Dashboard_LabelsEntry {
//		if in == nil {
//			return nil
//		}
//		out := &pb.Dashboard_LabelsEntry{}
//		out.Key = ValueOf(in.Key)
//		out.Value = ValueOf(in.Value)
//		return out
//	}
func ErrorReportingPanel_FromProto(mapCtx *MapContext, in *pb.ErrorReportingPanel) *krm.ErrorReportingPanel {
	if in == nil {
		return nil
	}
	out := &krm.ErrorReportingPanel{}
	out.ProjectNames = in.ProjectNames
	out.Services = in.Services
	out.Versions = in.Versions
	return out
}
func ErrorReportingPanel_ToProto(mapCtx *MapContext, in *krm.ErrorReportingPanel) *pb.ErrorReportingPanel {
	if in == nil {
		return nil
	}
	out := &pb.ErrorReportingPanel{}
	out.ProjectNames = in.ProjectNames
	out.Services = in.Services
	out.Versions = in.Versions
	return out
}
func GridLayout_FromProto(mapCtx *MapContext, in *pb.GridLayout) *krm.GridLayout {
	if in == nil {
		return nil
	}
	out := &krm.GridLayout{}
	out.Columns = LazyPtr(in.GetColumns())
	out.Widgets = Slice_FromProto(mapCtx, in.Widgets, Widget_FromProto)
	return out
}
func GridLayout_ToProto(mapCtx *MapContext, in *krm.GridLayout) *pb.GridLayout {
	if in == nil {
		return nil
	}
	out := &pb.GridLayout{}
	out.Columns = ValueOf(in.Columns)
	out.Widgets = Slice_ToProto(mapCtx, in.Widgets, Widget_ToProto)
	return out
}

//	func IncidentList_FromProto(mapCtx *MapContext, in *pb.IncidentList) *krm.IncidentList {
//		if in == nil {
//			return nil
//		}
//		out := &krm.IncidentList{}
//		out.MonitoredResources = Slice_FromProto(mapCtx, in.MonitoredResources, string_FromProto)
//		out.PolicyNames = in.PolicyNames
//		return out
//	}
//
//	func IncidentList_ToProto(mapCtx *MapContext, in *krm.IncidentList) *pb.IncidentList {
//		if in == nil {
//			return nil
//		}
//		out := &pb.IncidentList{}
//		out.MonitoredResources = Slice_ToProto(mapCtx, in.MonitoredResources, string_ToProto)
//		out.PolicyNames = in.PolicyNames
//		return out
//	}
func LogsPanel_FromProto(mapCtx *MapContext, in *pb.LogsPanel) *krm.LogsPanel {
	if in == nil {
		return nil
	}
	out := &krm.LogsPanel{}
	out.Filter = LazyPtr(in.GetFilter())
	out.ResourceNames = LogsPanel_ResourceNames_FromProto(mapCtx, in.ResourceNames)
	return out
}

func LogsPanel_ToProto(mapCtx *MapContext, in *krm.LogsPanel) *pb.LogsPanel {
	if in == nil {
		return nil
	}
	out := &pb.LogsPanel{}
	out.Filter = ValueOf(in.Filter)
	out.ResourceNames = LogsPanel_ResourceNames_ToProto(mapCtx, in.ResourceNames)
	return out
}
func MonitoringDashboardSpec_FromProto(mapCtx *MapContext, in *pb.Dashboard) *krm.MonitoringDashboardSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringDashboardSpec{}
	// MISSING: Name
	out.DisplayName = LazyPtr(in.GetDisplayName())
	// MISSING: Etag
	out.GridLayout = GridLayout_FromProto(mapCtx, in.GetGridLayout())
	out.MosaicLayout = MosaicLayout_FromProto(mapCtx, in.GetMosaicLayout())
	out.RowLayout = RowLayout_FromProto(mapCtx, in.GetRowLayout())
	out.ColumnLayout = ColumnLayout_FromProto(mapCtx, in.GetColumnLayout())
	// MISSING: DashboardFilters
	// MISSING: Labels
	return out
}
func MonitoringDashboardSpec_ToProto(mapCtx *MapContext, in *krm.MonitoringDashboardSpec) *pb.Dashboard {
	if in == nil {
		return nil
	}
	out := &pb.Dashboard{}
	// MISSING: Name
	out.DisplayName = ValueOf(in.DisplayName)
	// MISSING: Etag
	if oneof := GridLayout_ToProto(mapCtx, in.GridLayout); oneof != nil {
		out.Layout = &pb.Dashboard_GridLayout{GridLayout: oneof}
	}
	if oneof := MosaicLayout_ToProto(mapCtx, in.MosaicLayout); oneof != nil {
		out.Layout = &pb.Dashboard_MosaicLayout{MosaicLayout: oneof}
	}
	if oneof := RowLayout_ToProto(mapCtx, in.RowLayout); oneof != nil {
		out.Layout = &pb.Dashboard_RowLayout{RowLayout: oneof}
	}
	if oneof := ColumnLayout_ToProto(mapCtx, in.ColumnLayout); oneof != nil {
		out.Layout = &pb.Dashboard_ColumnLayout{ColumnLayout: oneof}
	}
	// MISSING: DashboardFilters
	// MISSING: Labels
	return out
}
func MonitoringDashboardStatus_FromProto(mapCtx *MapContext, in *pb.Dashboard) *krm.MonitoringDashboardStatus {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringDashboardStatus{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Etag = LazyPtr(in.GetEtag())
	// MISSING: GridLayout
	// MISSING: MosaicLayout
	// MISSING: RowLayout
	// MISSING: ColumnLayout
	// MISSING: DashboardFilters
	// MISSING: Labels
	return out
}
func MonitoringDashboardStatus_ToProto(mapCtx *MapContext, in *krm.MonitoringDashboardStatus) *pb.Dashboard {
	if in == nil {
		return nil
	}
	out := &pb.Dashboard{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Etag = ValueOf(in.Etag)
	// MISSING: GridLayout
	// MISSING: MosaicLayout
	// MISSING: RowLayout
	// MISSING: ColumnLayout
	// MISSING: DashboardFilters
	// MISSING: Labels
	return out
}
func MosaicLayout_FromProto(mapCtx *MapContext, in *pb.MosaicLayout) *krm.MosaicLayout {
	if in == nil {
		return nil
	}
	out := &krm.MosaicLayout{}
	out.Columns = LazyPtr(in.GetColumns())
	out.Tiles = Slice_FromProto(mapCtx, in.Tiles, MosaicLayout_Tile_FromProto)
	return out
}
func MosaicLayout_ToProto(mapCtx *MapContext, in *krm.MosaicLayout) *pb.MosaicLayout {
	if in == nil {
		return nil
	}
	out := &pb.MosaicLayout{}
	out.Columns = ValueOf(in.Columns)
	out.Tiles = Slice_ToProto(mapCtx, in.Tiles, MosaicLayout_Tile_ToProto)
	return out
}
func MosaicLayout_Tile_FromProto(mapCtx *MapContext, in *pb.MosaicLayout_Tile) *krm.MosaicLayout_Tile {
	if in == nil {
		return nil
	}
	out := &krm.MosaicLayout_Tile{}
	out.XPos = LazyPtr(in.GetXPos())
	out.YPos = LazyPtr(in.GetYPos())
	out.Width = LazyPtr(in.GetWidth())
	out.Height = LazyPtr(in.GetHeight())
	out.Widget = Widget_FromProto(mapCtx, in.GetWidget())
	return out
}
func MosaicLayout_Tile_ToProto(mapCtx *MapContext, in *krm.MosaicLayout_Tile) *pb.MosaicLayout_Tile {
	if in == nil {
		return nil
	}
	out := &pb.MosaicLayout_Tile{}
	out.XPos = ValueOf(in.XPos)
	out.YPos = ValueOf(in.YPos)
	out.Width = ValueOf(in.Width)
	out.Height = ValueOf(in.Height)
	out.Widget = Widget_ToProto(mapCtx, in.Widget)
	return out
}
func PickTimeSeriesFilter_FromProto(mapCtx *MapContext, in *pb.PickTimeSeriesFilter) *krm.PickTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &krm.PickTimeSeriesFilter{}
	out.RankingMethod = Enum_FromProto(mapCtx, in.RankingMethod)
	out.NumTimeSeries = LazyPtr(in.GetNumTimeSeries())
	out.Direction = Enum_FromProto(mapCtx, in.Direction)
	// MISSING: Interval
	return out
}
func PickTimeSeriesFilter_ToProto(mapCtx *MapContext, in *krm.PickTimeSeriesFilter) *pb.PickTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &pb.PickTimeSeriesFilter{}
	out.RankingMethod = Enum_ToProto[pb.PickTimeSeriesFilter_Method](mapCtx, in.RankingMethod)
	out.NumTimeSeries = ValueOf(in.NumTimeSeries)
	out.Direction = Enum_ToProto[pb.PickTimeSeriesFilter_Direction](mapCtx, in.Direction)
	// MISSING: Interval
	return out
}

// func PieChart_FromProto(mapCtx *MapContext, in *pb.PieChart) *krm.PieChart {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.PieChart{}
// 	out.DataSets = Slice_FromProto(mapCtx, in.DataSets, PieChart_PieChartDataSet_FromProto)
// 	out.ChartType = Enum_FromProto(mapCtx, in.ChartType)
// 	out.ShowLabels = LazyPtr(in.GetShowLabels())
// 	return out
// }
// func PieChart_ToProto(mapCtx *MapContext, in *krm.PieChart) *pb.PieChart {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.PieChart{}
// 	out.DataSets = Slice_ToProto(mapCtx, in.DataSets, PieChart_PieChartDataSet_ToProto)
// 	out.ChartType = Enum_ToProto[pb.PieChart_PieChartType](mapCtx, in.ChartType)
// 	out.ShowLabels = ValueOf(in.ShowLabels)
// 	return out
// }

//	func PieChart_PieChartDataSet_FromProto(mapCtx *MapContext, in *pb.PieChart_PieChartDataSet) *krm.PieChart_PieChartDataSet {
//		if in == nil {
//			return nil
//		}
//		out := &krm.PieChart_PieChartDataSet{}
//		out.TimeSeriesQuery = TimeSeriesQuery_FromProto(mapCtx, in.GetTimeSeriesQuery())
//		out.SliceNameTemplate = LazyPtr(in.GetSliceNameTemplate())
//		out.MinAlignmentPeriod = PieChartDataSet_MinAlignmentPeriod_FromProto(mapCtx, in.GetMinAlignmentPeriod())
//		return out
//	}
//
//	func PieChart_PieChartDataSet_ToProto(mapCtx *MapContext, in *krm.PieChart_PieChartDataSet) *pb.PieChart_PieChartDataSet {
//		if in == nil {
//			return nil
//		}
//		out := &pb.PieChart_PieChartDataSet{}
//		out.TimeSeriesQuery = TimeSeriesQuery_ToProto(mapCtx, in.TimeSeriesQuery)
//		out.SliceNameTemplate = ValueOf(in.SliceNameTemplate)
//		out.MinAlignmentPeriod = PieChartDataSet_MinAlignmentPeriod_ToProto(mapCtx, in.MinAlignmentPeriod)
//		return out
//	}
func RowLayout_FromProto(mapCtx *MapContext, in *pb.RowLayout) *krm.RowLayout {
	if in == nil {
		return nil
	}
	out := &krm.RowLayout{}
	out.Rows = Slice_FromProto(mapCtx, in.Rows, RowLayout_Row_FromProto)
	return out
}
func RowLayout_ToProto(mapCtx *MapContext, in *krm.RowLayout) *pb.RowLayout {
	if in == nil {
		return nil
	}
	out := &pb.RowLayout{}
	out.Rows = Slice_ToProto(mapCtx, in.Rows, RowLayout_Row_ToProto)
	return out
}
func RowLayout_Row_FromProto(mapCtx *MapContext, in *pb.RowLayout_Row) *krm.RowLayout_Row {
	if in == nil {
		return nil
	}
	out := &krm.RowLayout_Row{}
	out.Weight = LazyPtr(in.GetWeight())
	out.Widgets = Slice_FromProto(mapCtx, in.Widgets, Widget_FromProto)
	return out
}
func RowLayout_Row_ToProto(mapCtx *MapContext, in *krm.RowLayout_Row) *pb.RowLayout_Row {
	if in == nil {
		return nil
	}
	out := &pb.RowLayout_Row{}
	out.Weight = ValueOf(in.Weight)
	out.Widgets = Slice_ToProto(mapCtx, in.Widgets, Widget_ToProto)
	return out
}
func Scorecard_FromProto(mapCtx *MapContext, in *pb.Scorecard) *krm.Scorecard {
	if in == nil {
		return nil
	}
	out := &krm.Scorecard{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(mapCtx, in.GetTimeSeriesQuery())
	out.GaugeView = Scorecard_GaugeView_FromProto(mapCtx, in.GetGaugeView())
	out.SparkChartView = Scorecard_SparkChartView_FromProto(mapCtx, in.GetSparkChartView())
	// MISSING: BlankView
	out.Thresholds = Slice_FromProto(mapCtx, in.Thresholds, Threshold_FromProto)
	return out
}
func Scorecard_ToProto(mapCtx *MapContext, in *krm.Scorecard) *pb.Scorecard {
	if in == nil {
		return nil
	}
	out := &pb.Scorecard{}
	out.TimeSeriesQuery = TimeSeriesQuery_ToProto(mapCtx, in.TimeSeriesQuery)
	if oneof := Scorecard_GaugeView_ToProto(mapCtx, in.GaugeView); oneof != nil {
		out.DataView = &pb.Scorecard_GaugeView_{GaugeView: oneof}
	}
	if oneof := Scorecard_SparkChartView_ToProto(mapCtx, in.SparkChartView); oneof != nil {
		out.DataView = &pb.Scorecard_SparkChartView_{SparkChartView: oneof}
	}
	// MISSING: BlankView
	out.Thresholds = Slice_ToProto(mapCtx, in.Thresholds, Threshold_ToProto)
	return out
}
func Scorecard_GaugeView_FromProto(mapCtx *MapContext, in *pb.Scorecard_GaugeView) *krm.Scorecard_GaugeView {
	if in == nil {
		return nil
	}
	out := &krm.Scorecard_GaugeView{}
	out.LowerBound = LazyPtr(in.GetLowerBound())
	out.UpperBound = LazyPtr(in.GetUpperBound())
	return out
}
func Scorecard_GaugeView_ToProto(mapCtx *MapContext, in *krm.Scorecard_GaugeView) *pb.Scorecard_GaugeView {
	if in == nil {
		return nil
	}
	out := &pb.Scorecard_GaugeView{}
	out.LowerBound = ValueOf(in.LowerBound)
	out.UpperBound = ValueOf(in.UpperBound)
	return out
}
func Scorecard_SparkChartView_FromProto(mapCtx *MapContext, in *pb.Scorecard_SparkChartView) *krm.Scorecard_SparkChartView {
	if in == nil {
		return nil
	}
	out := &krm.Scorecard_SparkChartView{}
	out.SparkChartType = Enum_FromProto(mapCtx, in.SparkChartType)
	out.MinAlignmentPeriod = SparkChartView_MinAlignmentPeriod_FromProto(mapCtx, in.GetMinAlignmentPeriod())
	return out
}
func Scorecard_SparkChartView_ToProto(mapCtx *MapContext, in *krm.Scorecard_SparkChartView) *pb.Scorecard_SparkChartView {
	if in == nil {
		return nil
	}
	out := &pb.Scorecard_SparkChartView{}
	out.SparkChartType = Enum_ToProto[pb.SparkChartType](mapCtx, in.SparkChartType)
	out.MinAlignmentPeriod = SparkChartView_MinAlignmentPeriod_ToProto(mapCtx, in.MinAlignmentPeriod)
	return out
}
func SectionHeader_FromProto(mapCtx *MapContext, in *pb.SectionHeader) *krm.SectionHeader {
	if in == nil {
		return nil
	}
	out := &krm.SectionHeader{}
	out.Subtitle = LazyPtr(in.GetSubtitle())
	out.DividerBelow = LazyPtr(in.GetDividerBelow())
	return out
}
func SectionHeader_ToProto(mapCtx *MapContext, in *krm.SectionHeader) *pb.SectionHeader {
	if in == nil {
		return nil
	}
	out := &pb.SectionHeader{}
	out.Subtitle = ValueOf(in.Subtitle)
	out.DividerBelow = ValueOf(in.DividerBelow)
	return out
}
func SingleViewGroup_FromProto(mapCtx *MapContext, in *pb.SingleViewGroup) *krm.SingleViewGroup {
	if in == nil {
		return nil
	}
	out := &krm.SingleViewGroup{}
	return out
}
func SingleViewGroup_ToProto(mapCtx *MapContext, in *krm.SingleViewGroup) *pb.SingleViewGroup {
	if in == nil {
		return nil
	}
	out := &pb.SingleViewGroup{}
	return out
}
func StatisticalTimeSeriesFilter_FromProto(mapCtx *MapContext, in *pb.StatisticalTimeSeriesFilter) *krm.StatisticalTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &krm.StatisticalTimeSeriesFilter{}
	out.RankingMethod = Enum_FromProto(mapCtx, in.RankingMethod)
	out.NumTimeSeries = LazyPtr(in.GetNumTimeSeries())
	return out
}
func StatisticalTimeSeriesFilter_ToProto(mapCtx *MapContext, in *krm.StatisticalTimeSeriesFilter) *pb.StatisticalTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &pb.StatisticalTimeSeriesFilter{}
	out.RankingMethod = Enum_ToProto[pb.StatisticalTimeSeriesFilter_Method](mapCtx, in.RankingMethod)
	out.NumTimeSeries = ValueOf(in.NumTimeSeries)
	return out
}
func TableDisplayOptions_FromProto(mapCtx *MapContext, in *pb.TableDisplayOptions) *krm.TableDisplayOptions {
	if in == nil {
		return nil
	}
	out := &krm.TableDisplayOptions{}
	out.ShownColumns = in.ShownColumns
	return out
}
func TableDisplayOptions_ToProto(mapCtx *MapContext, in *krm.TableDisplayOptions) *pb.TableDisplayOptions {
	if in == nil {
		return nil
	}
	out := &pb.TableDisplayOptions{}
	out.ShownColumns = in.ShownColumns
	return out
}
func Text_FromProto(mapCtx *MapContext, in *pb.Text) *krm.Text {
	if in == nil {
		return nil
	}
	out := &krm.Text{}
	out.Content = LazyPtr(in.GetContent())
	out.Format = Enum_FromProto(mapCtx, in.Format)
	// MISSING: Style
	return out
}
func Text_ToProto(mapCtx *MapContext, in *krm.Text) *pb.Text {
	if in == nil {
		return nil
	}
	out := &pb.Text{}
	out.Content = ValueOf(in.Content)
	out.Format = Enum_ToProto[pb.Text_Format](mapCtx, in.Format)
	// MISSING: Style
	return out
}
func Text_TextStyle_FromProto(mapCtx *MapContext, in *pb.Text_TextStyle) *krm.Text_TextStyle {
	if in == nil {
		return nil
	}
	out := &krm.Text_TextStyle{}
	out.BackgroundColor = LazyPtr(in.GetBackgroundColor())
	out.TextColor = LazyPtr(in.GetTextColor())
	out.HorizontalAlignment = Enum_FromProto(mapCtx, in.HorizontalAlignment)
	out.VerticalAlignment = Enum_FromProto(mapCtx, in.VerticalAlignment)
	out.Padding = Enum_FromProto(mapCtx, in.Padding)
	out.FontSize = Enum_FromProto(mapCtx, in.FontSize)
	out.PointerLocation = Enum_FromProto(mapCtx, in.PointerLocation)
	return out
}
func Text_TextStyle_ToProto(mapCtx *MapContext, in *krm.Text_TextStyle) *pb.Text_TextStyle {
	if in == nil {
		return nil
	}
	out := &pb.Text_TextStyle{}
	out.BackgroundColor = ValueOf(in.BackgroundColor)
	out.TextColor = ValueOf(in.TextColor)
	out.HorizontalAlignment = Enum_ToProto[pb.Text_TextStyle_HorizontalAlignment](mapCtx, in.HorizontalAlignment)
	out.VerticalAlignment = Enum_ToProto[pb.Text_TextStyle_VerticalAlignment](mapCtx, in.VerticalAlignment)
	out.Padding = Enum_ToProto[pb.Text_TextStyle_PaddingSize](mapCtx, in.Padding)
	out.FontSize = Enum_ToProto[pb.Text_TextStyle_FontSize](mapCtx, in.FontSize)
	out.PointerLocation = Enum_ToProto[pb.Text_TextStyle_PointerLocation](mapCtx, in.PointerLocation)
	return out
}
func Threshold_FromProto(mapCtx *MapContext, in *pb.Threshold) *krm.Threshold {
	if in == nil {
		return nil
	}
	out := &krm.Threshold{}
	out.Label = LazyPtr(in.GetLabel())
	out.Value = LazyPtr(in.GetValue())
	out.Color = Enum_FromProto(mapCtx, in.Color)
	out.Direction = Enum_FromProto(mapCtx, in.Direction)
	// MISSING: TargetAxis
	return out
}
func Threshold_ToProto(mapCtx *MapContext, in *krm.Threshold) *pb.Threshold {
	if in == nil {
		return nil
	}
	out := &pb.Threshold{}
	out.Label = ValueOf(in.Label)
	out.Value = ValueOf(in.Value)
	out.Color = Enum_ToProto[pb.Threshold_Color](mapCtx, in.Color)
	out.Direction = Enum_ToProto[pb.Threshold_Direction](mapCtx, in.Direction)
	// MISSING: TargetAxis
	return out
}
func TimeSeriesFilter_FromProto(mapCtx *MapContext, in *pb.TimeSeriesFilter) *krm.TimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesFilter{}
	out.Filter = LazyPtr(in.GetFilter())
	out.Aggregation = Aggregation_FromProto(mapCtx, in.GetAggregation())
	out.SecondaryAggregation = Aggregation_FromProto(mapCtx, in.GetSecondaryAggregation())
	out.PickTimeSeriesFilter = PickTimeSeriesFilter_FromProto(mapCtx, in.GetPickTimeSeriesFilter())
	// MISSING: StatisticalTimeSeriesFilter
	return out
}
func TimeSeriesFilter_ToProto(mapCtx *MapContext, in *krm.TimeSeriesFilter) *pb.TimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesFilter{}
	out.Filter = ValueOf(in.Filter)
	out.Aggregation = Aggregation_ToProto(mapCtx, in.Aggregation)
	out.SecondaryAggregation = Aggregation_ToProto(mapCtx, in.SecondaryAggregation)
	if oneof := PickTimeSeriesFilter_ToProto(mapCtx, in.PickTimeSeriesFilter); oneof != nil {
		out.OutputFilter = &pb.TimeSeriesFilter_PickTimeSeriesFilter{PickTimeSeriesFilter: oneof}
	}
	// MISSING: StatisticalTimeSeriesFilter
	return out
}
func TimeSeriesFilterRatio_FromProto(mapCtx *MapContext, in *pb.TimeSeriesFilterRatio) *krm.TimeSeriesFilterRatio {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesFilterRatio{}
	out.Numerator = TimeSeriesFilterRatio_RatioPart_FromProto(mapCtx, in.GetNumerator())
	out.Denominator = TimeSeriesFilterRatio_RatioPart_FromProto(mapCtx, in.GetDenominator())
	out.SecondaryAggregation = Aggregation_FromProto(mapCtx, in.GetSecondaryAggregation())
	out.PickTimeSeriesFilter = PickTimeSeriesFilter_FromProto(mapCtx, in.GetPickTimeSeriesFilter())
	// MISSING: StatisticalTimeSeriesFilter
	return out
}
func TimeSeriesFilterRatio_ToProto(mapCtx *MapContext, in *krm.TimeSeriesFilterRatio) *pb.TimeSeriesFilterRatio {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesFilterRatio{}
	out.Numerator = TimeSeriesFilterRatio_RatioPart_ToProto(mapCtx, in.Numerator)
	out.Denominator = TimeSeriesFilterRatio_RatioPart_ToProto(mapCtx, in.Denominator)
	out.SecondaryAggregation = Aggregation_ToProto(mapCtx, in.SecondaryAggregation)
	if oneof := PickTimeSeriesFilter_ToProto(mapCtx, in.PickTimeSeriesFilter); oneof != nil {
		out.OutputFilter = &pb.TimeSeriesFilterRatio_PickTimeSeriesFilter{PickTimeSeriesFilter: oneof}
	}
	// MISSING: StatisticalTimeSeriesFilter
	return out
}
func TimeSeriesFilterRatio_RatioPart_FromProto(mapCtx *MapContext, in *pb.TimeSeriesFilterRatio_RatioPart) *krm.TimeSeriesFilterRatio_RatioPart {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesFilterRatio_RatioPart{}
	out.Filter = LazyPtr(in.GetFilter())
	out.Aggregation = Aggregation_FromProto(mapCtx, in.GetAggregation())
	return out
}
func TimeSeriesFilterRatio_RatioPart_ToProto(mapCtx *MapContext, in *krm.TimeSeriesFilterRatio_RatioPart) *pb.TimeSeriesFilterRatio_RatioPart {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesFilterRatio_RatioPart{}
	out.Filter = ValueOf(in.Filter)
	out.Aggregation = Aggregation_ToProto(mapCtx, in.Aggregation)
	return out
}
func TimeSeriesQuery_FromProto(mapCtx *MapContext, in *pb.TimeSeriesQuery) *krm.TimeSeriesQuery {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesQuery{}
	out.TimeSeriesFilter = TimeSeriesFilter_FromProto(mapCtx, in.GetTimeSeriesFilter())
	out.TimeSeriesFilterRatio = TimeSeriesFilterRatio_FromProto(mapCtx, in.GetTimeSeriesFilterRatio())
	out.TimeSeriesQueryLanguage = LazyPtr(in.GetTimeSeriesQueryLanguage())
	// MISSING: PrometheusQuery
	out.UnitOverride = LazyPtr(in.GetUnitOverride())
	// MISSING: OutputFullDuration
	return out
}
func TimeSeriesQuery_ToProto(mapCtx *MapContext, in *krm.TimeSeriesQuery) *pb.TimeSeriesQuery {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesQuery{}
	if oneof := TimeSeriesFilter_ToProto(mapCtx, in.TimeSeriesFilter); oneof != nil {
		out.Source = &pb.TimeSeriesQuery_TimeSeriesFilter{TimeSeriesFilter: oneof}
	}
	if oneof := TimeSeriesFilterRatio_ToProto(mapCtx, in.TimeSeriesFilterRatio); oneof != nil {
		out.Source = &pb.TimeSeriesQuery_TimeSeriesFilterRatio{TimeSeriesFilterRatio: oneof}
	}
	if oneof := TimeSeriesQuery_TimeSeriesQueryLanguage_ToProto(mapCtx, in.TimeSeriesQueryLanguage); oneof != nil {
		out.Source = oneof
	}
	// MISSING: PrometheusQuery
	out.UnitOverride = ValueOf(in.UnitOverride)
	// MISSING: OutputFullDuration
	return out
}
func TimeSeriesTable_FromProto(mapCtx *MapContext, in *pb.TimeSeriesTable) *krm.TimeSeriesTable {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesTable{}
	out.DataSets = Slice_FromProto(mapCtx, in.DataSets, TimeSeriesTable_TableDataSet_FromProto)
	// MISSING: MetricVisualization
	out.ColumnSettings = Slice_FromProto(mapCtx, in.ColumnSettings, TimeSeriesTable_ColumnSettings_FromProto)
	return out
}
func TimeSeriesTable_ToProto(mapCtx *MapContext, in *krm.TimeSeriesTable) *pb.TimeSeriesTable {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesTable{}
	out.DataSets = Slice_ToProto(mapCtx, in.DataSets, TimeSeriesTable_TableDataSet_ToProto)
	// MISSING: MetricVisualization
	out.ColumnSettings = Slice_ToProto(mapCtx, in.ColumnSettings, TimeSeriesTable_ColumnSettings_ToProto)
	return out
}
func TimeSeriesTable_ColumnSettings_FromProto(mapCtx *MapContext, in *pb.TimeSeriesTable_ColumnSettings) *krm.TimeSeriesTable_ColumnSettings {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesTable_ColumnSettings{}
	out.Column = LazyPtr(in.GetColumn())
	out.Visible = LazyPtr(in.GetVisible())
	return out
}
func TimeSeriesTable_ColumnSettings_ToProto(mapCtx *MapContext, in *krm.TimeSeriesTable_ColumnSettings) *pb.TimeSeriesTable_ColumnSettings {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesTable_ColumnSettings{}
	out.Column = ValueOf(in.Column)
	out.Visible = ValueOf(in.Visible)
	return out
}
func TimeSeriesTable_TableDataSet_FromProto(mapCtx *MapContext, in *pb.TimeSeriesTable_TableDataSet) *krm.TimeSeriesTable_TableDataSet {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesTable_TableDataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(mapCtx, in.GetTimeSeriesQuery())
	out.TableTemplate = LazyPtr(in.GetTableTemplate())
	out.MinAlignmentPeriod = TableDataSet_MinAlignmentPeriod_FromProto(mapCtx, in.GetMinAlignmentPeriod())
	out.TableDisplayOptions = TableDisplayOptions_FromProto(mapCtx, in.GetTableDisplayOptions())
	return out
}
func TimeSeriesTable_TableDataSet_ToProto(mapCtx *MapContext, in *krm.TimeSeriesTable_TableDataSet) *pb.TimeSeriesTable_TableDataSet {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesTable_TableDataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_ToProto(mapCtx, in.TimeSeriesQuery)
	out.TableTemplate = ValueOf(in.TableTemplate)
	out.MinAlignmentPeriod = TableDataSet_MinAlignmentPeriod_ToProto(mapCtx, in.MinAlignmentPeriod)
	out.TableDisplayOptions = TableDisplayOptions_ToProto(mapCtx, in.TableDisplayOptions)
	return out
}
func Widget_FromProto(mapCtx *MapContext, in *pb.Widget) *krm.Widget {
	if in == nil {
		return nil
	}
	out := &krm.Widget{}
	out.Title = LazyPtr(in.GetTitle())
	out.XyChart = XyChart_FromProto(mapCtx, in.GetXyChart())
	out.Scorecard = Scorecard_FromProto(mapCtx, in.GetScorecard())
	out.Text = Text_FromProto(mapCtx, in.GetText())
	out.Blank = Empty_FromProto(mapCtx, in.GetBlank())
	// MISSING: AlertChart
	// MISSING: TimeSeriesTable
	// MISSING: CollapsibleGroup
	out.LogsPanel = LogsPanel_FromProto(mapCtx, in.GetLogsPanel())
	// MISSING: IncidentList
	// MISSING: PieChart
	// MISSING: ErrorReportingPanel
	// MISSING: SectionHeader
	// MISSING: SingleViewGroup
	// MISSING: Id
	return out
}
func Widget_ToProto(mapCtx *MapContext, in *krm.Widget) *pb.Widget {
	if in == nil {
		return nil
	}
	out := &pb.Widget{}
	out.Title = ValueOf(in.Title)
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
	// MISSING: AlertChart
	// MISSING: TimeSeriesTable
	// MISSING: CollapsibleGroup
	if oneof := LogsPanel_ToProto(mapCtx, in.LogsPanel); oneof != nil {
		out.Content = &pb.Widget_LogsPanel{LogsPanel: oneof}
	}
	// MISSING: IncidentList
	// MISSING: PieChart
	// MISSING: ErrorReportingPanel
	// MISSING: SectionHeader
	// MISSING: SingleViewGroup
	// MISSING: Id
	return out
}
func XyChart_FromProto(mapCtx *MapContext, in *pb.XyChart) *krm.XyChart {
	if in == nil {
		return nil
	}
	out := &krm.XyChart{}
	out.DataSets = Slice_FromProto(mapCtx, in.DataSets, XyChart_DataSet_FromProto)
	out.TimeshiftDuration = XyChart_TimeshiftDuration_FromProto(mapCtx, in.GetTimeshiftDuration())
	out.Thresholds = Slice_FromProto(mapCtx, in.Thresholds, Threshold_FromProto)
	out.XAxis = XyChart_Axis_FromProto(mapCtx, in.GetXAxis())
	out.YAxis = XyChart_Axis_FromProto(mapCtx, in.GetYAxis())
	// MISSING: Y2Axis
	out.ChartOptions = ChartOptions_FromProto(mapCtx, in.GetChartOptions())
	return out
}
func XyChart_ToProto(mapCtx *MapContext, in *krm.XyChart) *pb.XyChart {
	if in == nil {
		return nil
	}
	out := &pb.XyChart{}
	out.DataSets = Slice_ToProto(mapCtx, in.DataSets, XyChart_DataSet_ToProto)
	out.TimeshiftDuration = XyChart_TimeshiftDuration_ToProto(mapCtx, in.TimeshiftDuration)
	out.Thresholds = Slice_ToProto(mapCtx, in.Thresholds, Threshold_ToProto)
	out.XAxis = XyChart_Axis_ToProto(mapCtx, in.XAxis)
	out.YAxis = XyChart_Axis_ToProto(mapCtx, in.YAxis)
	// MISSING: Y2Axis
	out.ChartOptions = ChartOptions_ToProto(mapCtx, in.ChartOptions)
	return out
}
func XyChart_Axis_FromProto(mapCtx *MapContext, in *pb.XyChart_Axis) *krm.XyChart_Axis {
	if in == nil {
		return nil
	}
	out := &krm.XyChart_Axis{}
	out.Label = LazyPtr(in.GetLabel())
	out.Scale = Enum_FromProto(mapCtx, in.Scale)
	return out
}
func XyChart_Axis_ToProto(mapCtx *MapContext, in *krm.XyChart_Axis) *pb.XyChart_Axis {
	if in == nil {
		return nil
	}
	out := &pb.XyChart_Axis{}
	out.Label = ValueOf(in.Label)
	out.Scale = Enum_ToProto[pb.XyChart_Axis_Scale](mapCtx, in.Scale)
	return out
}
func XyChart_DataSet_FromProto(mapCtx *MapContext, in *pb.XyChart_DataSet) *krm.XyChart_DataSet {
	if in == nil {
		return nil
	}
	out := &krm.XyChart_DataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(mapCtx, in.GetTimeSeriesQuery())
	out.PlotType = Enum_FromProto(mapCtx, in.PlotType)
	out.LegendTemplate = LazyPtr(in.GetLegendTemplate())
	out.MinAlignmentPeriod = DataSet_MinAlignmentPeriod_FromProto(mapCtx, in.GetMinAlignmentPeriod())
	// MISSING: TargetAxis
	return out
}
func XyChart_DataSet_ToProto(mapCtx *MapContext, in *krm.XyChart_DataSet) *pb.XyChart_DataSet {
	if in == nil {
		return nil
	}
	out := &pb.XyChart_DataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_ToProto(mapCtx, in.TimeSeriesQuery)
	out.PlotType = Enum_ToProto[pb.XyChart_DataSet_PlotType](mapCtx, in.PlotType)
	out.LegendTemplate = ValueOf(in.LegendTemplate)
	out.MinAlignmentPeriod = DataSet_MinAlignmentPeriod_ToProto(mapCtx, in.MinAlignmentPeriod)
	// MISSING: TargetAxis
	return out
}
