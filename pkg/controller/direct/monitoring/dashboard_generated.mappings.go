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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Aggregation_FromProto(mapCtx *direct.MapContext, in *pb.Aggregation) *krm.Aggregation {
	if in == nil {
		return nil
	}
	out := &krm.Aggregation{}
	out.AlignmentPeriod = Aggregation_AlignmentPeriod_FromProto(mapCtx, in.GetAlignmentPeriod())
	out.PerSeriesAligner = direct.Enum_FromProto(mapCtx, in.PerSeriesAligner)
	out.CrossSeriesReducer = direct.Enum_FromProto(mapCtx, in.CrossSeriesReducer)
	out.GroupByFields = in.GroupByFields
	return out
}
func Aggregation_ToProto(mapCtx *direct.MapContext, in *krm.Aggregation) *pb.Aggregation {
	if in == nil {
		return nil
	}
	out := &pb.Aggregation{}
	out.AlignmentPeriod = Aggregation_AlignmentPeriod_ToProto(mapCtx, in.AlignmentPeriod)
	out.PerSeriesAligner = direct.Enum_ToProto[pb.Aggregation_Aligner](mapCtx, in.PerSeriesAligner)
	out.CrossSeriesReducer = direct.Enum_ToProto[pb.Aggregation_Reducer](mapCtx, in.CrossSeriesReducer)
	out.GroupByFields = in.GroupByFields
	return out
}

func ChartOptions_FromProto(mapCtx *direct.MapContext, in *pb.ChartOptions) *krm.ChartOptions {
	if in == nil {
		return nil
	}
	out := &krm.ChartOptions{}
	out.Mode = direct.Enum_FromProto(mapCtx, in.Mode)
	return out
}
func ChartOptions_ToProto(mapCtx *direct.MapContext, in *krm.ChartOptions) *pb.ChartOptions {
	if in == nil {
		return nil
	}
	out := &pb.ChartOptions{}
	out.Mode = direct.Enum_ToProto[pb.ChartOptions_Mode](mapCtx, in.Mode)
	return out
}
func CollapsibleGroup_FromProto(mapCtx *direct.MapContext, in *pb.CollapsibleGroup) *krm.CollapsibleGroup {
	if in == nil {
		return nil
	}
	out := &krm.CollapsibleGroup{}
	out.Collapsed = direct.LazyPtr(in.GetCollapsed())
	return out
}
func CollapsibleGroup_ToProto(mapCtx *direct.MapContext, in *krm.CollapsibleGroup) *pb.CollapsibleGroup {
	if in == nil {
		return nil
	}
	out := &pb.CollapsibleGroup{}
	out.Collapsed = direct.ValueOf(in.Collapsed)
	return out
}
func ColumnLayout_FromProto(mapCtx *direct.MapContext, in *pb.ColumnLayout) *krm.ColumnLayout {
	if in == nil {
		return nil
	}
	out := &krm.ColumnLayout{}
	out.Columns = direct.Slice_FromProto(mapCtx, in.Columns, ColumnLayout_Column_FromProto)
	return out
}
func ColumnLayout_ToProto(mapCtx *direct.MapContext, in *krm.ColumnLayout) *pb.ColumnLayout {
	if in == nil {
		return nil
	}
	out := &pb.ColumnLayout{}
	out.Columns = direct.Slice_ToProto(mapCtx, in.Columns, ColumnLayout_Column_ToProto)
	return out
}
func ColumnLayout_Column_FromProto(mapCtx *direct.MapContext, in *pb.ColumnLayout_Column) *krm.ColumnLayout_Column {
	if in == nil {
		return nil
	}
	out := &krm.ColumnLayout_Column{}
	out.Weight = direct.LazyPtr(in.GetWeight())
	out.Widgets = direct.Slice_FromProto(mapCtx, in.Widgets, Widget_FromProto)
	return out
}
func ColumnLayout_Column_ToProto(mapCtx *direct.MapContext, in *krm.ColumnLayout_Column) *pb.ColumnLayout_Column {
	if in == nil {
		return nil
	}
	out := &pb.ColumnLayout_Column{}
	out.Weight = direct.ValueOf(in.Weight)
	out.Widgets = direct.Slice_ToProto(mapCtx, in.Widgets, Widget_ToProto)
	return out
}
func DashboardFilter_FromProto(mapCtx *direct.MapContext, in *pb.DashboardFilter) *krm.DashboardFilter {
	if in == nil {
		return nil
	}
	out := &krm.DashboardFilter{}
	out.LabelKey = direct.LazyPtr(in.GetLabelKey())
	out.TemplateVariable = direct.LazyPtr(in.GetTemplateVariable())
	out.StringValue = direct.LazyPtr(in.GetStringValue())
	out.FilterType = direct.Enum_FromProto(mapCtx, in.FilterType)
	return out
}
func DashboardFilter_ToProto(mapCtx *direct.MapContext, in *krm.DashboardFilter) *pb.DashboardFilter {
	if in == nil {
		return nil
	}
	out := &pb.DashboardFilter{}
	out.LabelKey = direct.ValueOf(in.LabelKey)
	out.TemplateVariable = direct.ValueOf(in.TemplateVariable)
	if oneof := DashboardFilter_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.DefaultValue = oneof
	}
	out.FilterType = direct.Enum_ToProto[pb.DashboardFilter_FilterType](mapCtx, in.FilterType)
	return out
}
func GridLayout_FromProto(mapCtx *direct.MapContext, in *pb.GridLayout) *krm.GridLayout {
	if in == nil {
		return nil
	}
	out := &krm.GridLayout{}
	out.Columns = direct.LazyPtr(in.GetColumns())
	out.Widgets = direct.Slice_FromProto(mapCtx, in.Widgets, Widget_FromProto)
	return out
}
func GridLayout_ToProto(mapCtx *direct.MapContext, in *krm.GridLayout) *pb.GridLayout {
	if in == nil {
		return nil
	}
	out := &pb.GridLayout{}
	out.Columns = direct.ValueOf(in.Columns)
	out.Widgets = direct.Slice_ToProto(mapCtx, in.Widgets, Widget_ToProto)
	return out
}

func LogsPanel_FromProto(mapCtx *direct.MapContext, in *pb.LogsPanel) *krm.LogsPanel {
	if in == nil {
		return nil
	}
	out := &krm.LogsPanel{}
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.ResourceNames = LogsPanel_ResourceNames_FromProto(mapCtx, in.ResourceNames)
	return out
}

func LogsPanel_ToProto(mapCtx *direct.MapContext, in *krm.LogsPanel) *pb.LogsPanel {
	if in == nil {
		return nil
	}
	out := &pb.LogsPanel{}
	out.Filter = direct.ValueOf(in.Filter)
	out.ResourceNames = LogsPanel_ResourceNames_ToProto(mapCtx, in.ResourceNames)
	return out
}

func MonitoringDashboardSpec_FromProto(mapCtx *direct.MapContext, in *pb.Dashboard) *krm.MonitoringDashboardSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringDashboardSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Etag
	out.GridLayout = GridLayout_FromProto(mapCtx, in.GetGridLayout())
	out.MosaicLayout = MosaicLayout_FromProto(mapCtx, in.GetMosaicLayout())
	out.RowLayout = RowLayout_FromProto(mapCtx, in.GetRowLayout())
	out.ColumnLayout = ColumnLayout_FromProto(mapCtx, in.GetColumnLayout())
	out.DashboardFilters = direct.Slice_FromProto(mapCtx, in.DashboardFilters, DashboardFilter_FromProto)
	// MISSING: Labels
	return out
}
func MonitoringDashboardSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringDashboardSpec) *pb.Dashboard {
	if in == nil {
		return nil
	}
	out := &pb.Dashboard{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
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
	out.DashboardFilters = direct.Slice_ToProto(mapCtx, in.DashboardFilters, DashboardFilter_ToProto)
	// MISSING: Labels
	return out
}
func MonitoringDashboardStatus_FromProto(mapCtx *direct.MapContext, in *pb.Dashboard) *krm.MonitoringDashboardStatus {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringDashboardStatus{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: GridLayout
	// MISSING: MosaicLayout
	// MISSING: RowLayout
	// MISSING: ColumnLayout
	// MISSING: DashboardFilters
	// MISSING: Labels
	return out
}
func MonitoringDashboardStatus_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringDashboardStatus) *pb.Dashboard {
	if in == nil {
		return nil
	}
	out := &pb.Dashboard{}
	// MISSING: Name
	// MISSING: DisplayName
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: GridLayout
	// MISSING: MosaicLayout
	// MISSING: RowLayout
	// MISSING: ColumnLayout
	// MISSING: DashboardFilters
	// MISSING: Labels
	return out
}
func MosaicLayout_FromProto(mapCtx *direct.MapContext, in *pb.MosaicLayout) *krm.MosaicLayout {
	if in == nil {
		return nil
	}
	out := &krm.MosaicLayout{}
	out.Columns = direct.LazyPtr(in.GetColumns())
	out.Tiles = direct.Slice_FromProto(mapCtx, in.Tiles, MosaicLayout_Tile_FromProto)
	return out
}
func MosaicLayout_ToProto(mapCtx *direct.MapContext, in *krm.MosaicLayout) *pb.MosaicLayout {
	if in == nil {
		return nil
	}
	out := &pb.MosaicLayout{}
	out.Columns = direct.ValueOf(in.Columns)
	out.Tiles = direct.Slice_ToProto(mapCtx, in.Tiles, MosaicLayout_Tile_ToProto)
	return out
}
func MosaicLayout_Tile_FromProto(mapCtx *direct.MapContext, in *pb.MosaicLayout_Tile) *krm.MosaicLayout_Tile {
	if in == nil {
		return nil
	}
	out := &krm.MosaicLayout_Tile{}
	out.XPos = direct.LazyPtr(in.GetXPos())
	out.YPos = direct.LazyPtr(in.GetYPos())
	out.Width = direct.LazyPtr(in.GetWidth())
	out.Height = direct.LazyPtr(in.GetHeight())
	out.Widget = Widget_FromProto(mapCtx, in.GetWidget())
	return out
}
func MosaicLayout_Tile_ToProto(mapCtx *direct.MapContext, in *krm.MosaicLayout_Tile) *pb.MosaicLayout_Tile {
	if in == nil {
		return nil
	}
	out := &pb.MosaicLayout_Tile{}
	out.XPos = direct.ValueOf(in.XPos)
	out.YPos = direct.ValueOf(in.YPos)
	out.Width = direct.ValueOf(in.Width)
	out.Height = direct.ValueOf(in.Height)
	out.Widget = Widget_ToProto(mapCtx, in.Widget)
	return out
}
func PickTimeSeriesFilter_FromProto(mapCtx *direct.MapContext, in *pb.PickTimeSeriesFilter) *krm.PickTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &krm.PickTimeSeriesFilter{}
	out.RankingMethod = direct.Enum_FromProto(mapCtx, in.RankingMethod)
	out.NumTimeSeries = direct.LazyPtr(in.GetNumTimeSeries())
	out.Direction = direct.Enum_FromProto(mapCtx, in.Direction)
	// MISSING: Interval
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
	// MISSING: Interval
	return out
}

func PieChart_FromProto(mapCtx *direct.MapContext, in *pb.PieChart) *krm.PieChart {
	if in == nil {
		return nil
	}
	out := &krm.PieChart{}
	out.DataSets = direct.Slice_FromProto(mapCtx, in.DataSets, PieChart_PieChartDataSet_FromProto)
	out.ChartType = direct.Enum_FromProto(mapCtx, in.ChartType)
	out.ShowLabels = direct.LazyPtr(in.GetShowLabels())
	return out
}
func PieChart_ToProto(mapCtx *direct.MapContext, in *krm.PieChart) *pb.PieChart {
	if in == nil {
		return nil
	}
	out := &pb.PieChart{}
	out.DataSets = direct.Slice_ToProto(mapCtx, in.DataSets, PieChart_PieChartDataSet_ToProto)
	out.ChartType = direct.Enum_ToProto[pb.PieChart_PieChartType](mapCtx, in.ChartType)
	out.ShowLabels = direct.ValueOf(in.ShowLabels)
	return out
}
func PieChart_PieChartDataSet_FromProto(mapCtx *direct.MapContext, in *pb.PieChart_PieChartDataSet) *krm.PieChart_PieChartDataSet {
	if in == nil {
		return nil
	}
	out := &krm.PieChart_PieChartDataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(mapCtx, in.GetTimeSeriesQuery())
	out.SliceNameTemplate = direct.LazyPtr(in.GetSliceNameTemplate())
	out.MinAlignmentPeriod = PieChartDataSet_MinAlignmentPeriod_FromProto(mapCtx, in.GetMinAlignmentPeriod())
	return out
}
func PieChart_PieChartDataSet_ToProto(mapCtx *direct.MapContext, in *krm.PieChart_PieChartDataSet) *pb.PieChart_PieChartDataSet {
	if in == nil {
		return nil
	}
	out := &pb.PieChart_PieChartDataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_ToProto(mapCtx, in.TimeSeriesQuery)
	out.SliceNameTemplate = direct.ValueOf(in.SliceNameTemplate)
	out.MinAlignmentPeriod = PieChartDataSet_MinAlignmentPeriod_ToProto(mapCtx, in.MinAlignmentPeriod)
	return out
}

func RowLayout_FromProto(mapCtx *direct.MapContext, in *pb.RowLayout) *krm.RowLayout {
	if in == nil {
		return nil
	}
	out := &krm.RowLayout{}
	out.Rows = direct.Slice_FromProto(mapCtx, in.Rows, RowLayout_Row_FromProto)
	return out
}
func RowLayout_ToProto(mapCtx *direct.MapContext, in *krm.RowLayout) *pb.RowLayout {
	if in == nil {
		return nil
	}
	out := &pb.RowLayout{}
	out.Rows = direct.Slice_ToProto(mapCtx, in.Rows, RowLayout_Row_ToProto)
	return out
}
func RowLayout_Row_FromProto(mapCtx *direct.MapContext, in *pb.RowLayout_Row) *krm.RowLayout_Row {
	if in == nil {
		return nil
	}
	out := &krm.RowLayout_Row{}
	out.Weight = direct.LazyPtr(in.GetWeight())
	out.Widgets = direct.Slice_FromProto(mapCtx, in.Widgets, Widget_FromProto)
	return out
}
func RowLayout_Row_ToProto(mapCtx *direct.MapContext, in *krm.RowLayout_Row) *pb.RowLayout_Row {
	if in == nil {
		return nil
	}
	out := &pb.RowLayout_Row{}
	out.Weight = direct.ValueOf(in.Weight)
	out.Widgets = direct.Slice_ToProto(mapCtx, in.Widgets, Widget_ToProto)
	return out
}
func Scorecard_FromProto(mapCtx *direct.MapContext, in *pb.Scorecard) *krm.Scorecard {
	if in == nil {
		return nil
	}
	out := &krm.Scorecard{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(mapCtx, in.GetTimeSeriesQuery())
	out.GaugeView = Scorecard_GaugeView_FromProto(mapCtx, in.GetGaugeView())
	out.SparkChartView = Scorecard_SparkChartView_FromProto(mapCtx, in.GetSparkChartView())
	out.BlankView = BlankView_FromProto(mapCtx, in.GetBlankView())
	out.Thresholds = direct.Slice_FromProto(mapCtx, in.Thresholds, Threshold_FromProto)
	return out
}
func Scorecard_ToProto(mapCtx *direct.MapContext, in *krm.Scorecard) *pb.Scorecard {
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
	if oneof := BlankView_ToProto(mapCtx, in.BlankView); oneof != nil {
		out.DataView = &pb.Scorecard_BlankView{BlankView: oneof}
	}
	out.Thresholds = direct.Slice_ToProto(mapCtx, in.Thresholds, Threshold_ToProto)
	return out
}
func Scorecard_GaugeView_FromProto(mapCtx *direct.MapContext, in *pb.Scorecard_GaugeView) *krm.Scorecard_GaugeView {
	if in == nil {
		return nil
	}
	out := &krm.Scorecard_GaugeView{}
	out.LowerBound = direct.LazyPtr(in.GetLowerBound())
	out.UpperBound = direct.LazyPtr(in.GetUpperBound())
	return out
}
func Scorecard_GaugeView_ToProto(mapCtx *direct.MapContext, in *krm.Scorecard_GaugeView) *pb.Scorecard_GaugeView {
	if in == nil {
		return nil
	}
	out := &pb.Scorecard_GaugeView{}
	out.LowerBound = direct.ValueOf(in.LowerBound)
	out.UpperBound = direct.ValueOf(in.UpperBound)
	return out
}
func Scorecard_SparkChartView_FromProto(mapCtx *direct.MapContext, in *pb.Scorecard_SparkChartView) *krm.Scorecard_SparkChartView {
	if in == nil {
		return nil
	}
	out := &krm.Scorecard_SparkChartView{}
	out.SparkChartType = direct.Enum_FromProto(mapCtx, in.SparkChartType)
	out.MinAlignmentPeriod = SparkChartView_MinAlignmentPeriod_FromProto(mapCtx, in.GetMinAlignmentPeriod())
	return out
}
func Scorecard_SparkChartView_ToProto(mapCtx *direct.MapContext, in *krm.Scorecard_SparkChartView) *pb.Scorecard_SparkChartView {
	if in == nil {
		return nil
	}
	out := &pb.Scorecard_SparkChartView{}
	out.SparkChartType = direct.Enum_ToProto[pb.SparkChartType](mapCtx, in.SparkChartType)
	out.MinAlignmentPeriod = SparkChartView_MinAlignmentPeriod_ToProto(mapCtx, in.MinAlignmentPeriod)
	return out
}
func SectionHeader_FromProto(mapCtx *direct.MapContext, in *pb.SectionHeader) *krm.SectionHeader {
	if in == nil {
		return nil
	}
	out := &krm.SectionHeader{}
	out.Subtitle = direct.LazyPtr(in.GetSubtitle())
	out.DividerBelow = direct.LazyPtr(in.GetDividerBelow())
	return out
}
func SectionHeader_ToProto(mapCtx *direct.MapContext, in *krm.SectionHeader) *pb.SectionHeader {
	if in == nil {
		return nil
	}
	out := &pb.SectionHeader{}
	out.Subtitle = direct.ValueOf(in.Subtitle)
	out.DividerBelow = direct.ValueOf(in.DividerBelow)
	return out
}
func SingleViewGroup_FromProto(mapCtx *direct.MapContext, in *pb.SingleViewGroup) *krm.SingleViewGroup {
	if in == nil {
		return nil
	}
	out := &krm.SingleViewGroup{}
	return out
}
func SingleViewGroup_ToProto(mapCtx *direct.MapContext, in *krm.SingleViewGroup) *pb.SingleViewGroup {
	if in == nil {
		return nil
	}
	out := &pb.SingleViewGroup{}
	return out
}
func StatisticalTimeSeriesFilter_FromProto(mapCtx *direct.MapContext, in *pb.StatisticalTimeSeriesFilter) *krm.StatisticalTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &krm.StatisticalTimeSeriesFilter{}
	out.RankingMethod = direct.Enum_FromProto(mapCtx, in.RankingMethod)
	out.NumTimeSeries = direct.LazyPtr(in.GetNumTimeSeries())
	return out
}
func StatisticalTimeSeriesFilter_ToProto(mapCtx *direct.MapContext, in *krm.StatisticalTimeSeriesFilter) *pb.StatisticalTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &pb.StatisticalTimeSeriesFilter{}
	out.RankingMethod = direct.Enum_ToProto[pb.StatisticalTimeSeriesFilter_Method](mapCtx, in.RankingMethod)
	out.NumTimeSeries = direct.ValueOf(in.NumTimeSeries)
	return out
}
func TableDisplayOptions_FromProto(mapCtx *direct.MapContext, in *pb.TableDisplayOptions) *krm.TableDisplayOptions {
	if in == nil {
		return nil
	}
	out := &krm.TableDisplayOptions{}
	out.ShownColumns = in.ShownColumns
	return out
}
func TableDisplayOptions_ToProto(mapCtx *direct.MapContext, in *krm.TableDisplayOptions) *pb.TableDisplayOptions {
	if in == nil {
		return nil
	}
	out := &pb.TableDisplayOptions{}
	out.ShownColumns = in.ShownColumns
	return out
}
func Text_FromProto(mapCtx *direct.MapContext, in *pb.Text) *krm.Text {
	if in == nil {
		return nil
	}
	out := &krm.Text{}
	out.Content = direct.LazyPtr(in.GetContent())
	out.Format = direct.Enum_FromProto(mapCtx, in.Format)
	out.Style = Text_TextStyle_FromProto(mapCtx, in.GetStyle())
	return out
}
func Text_ToProto(mapCtx *direct.MapContext, in *krm.Text) *pb.Text {
	if in == nil {
		return nil
	}
	out := &pb.Text{}
	out.Content = direct.ValueOf(in.Content)
	out.Format = direct.Enum_ToProto[pb.Text_Format](mapCtx, in.Format)
	out.Style = Text_TextStyle_ToProto(mapCtx, in.Style)
	return out
}
func Text_TextStyle_FromProto(mapCtx *direct.MapContext, in *pb.Text_TextStyle) *krm.Text_TextStyle {
	if in == nil {
		return nil
	}
	out := &krm.Text_TextStyle{}
	out.BackgroundColor = direct.LazyPtr(in.GetBackgroundColor())
	out.TextColor = direct.LazyPtr(in.GetTextColor())
	out.HorizontalAlignment = direct.Enum_FromProto(mapCtx, in.HorizontalAlignment)
	out.VerticalAlignment = direct.Enum_FromProto(mapCtx, in.VerticalAlignment)
	out.Padding = direct.Enum_FromProto(mapCtx, in.Padding)
	out.FontSize = direct.Enum_FromProto(mapCtx, in.FontSize)
	out.PointerLocation = direct.Enum_FromProto(mapCtx, in.PointerLocation)
	return out
}
func Text_TextStyle_ToProto(mapCtx *direct.MapContext, in *krm.Text_TextStyle) *pb.Text_TextStyle {
	if in == nil {
		return nil
	}
	out := &pb.Text_TextStyle{}
	out.BackgroundColor = direct.ValueOf(in.BackgroundColor)
	out.TextColor = direct.ValueOf(in.TextColor)
	out.HorizontalAlignment = direct.Enum_ToProto[pb.Text_TextStyle_HorizontalAlignment](mapCtx, in.HorizontalAlignment)
	out.VerticalAlignment = direct.Enum_ToProto[pb.Text_TextStyle_VerticalAlignment](mapCtx, in.VerticalAlignment)
	out.Padding = direct.Enum_ToProto[pb.Text_TextStyle_PaddingSize](mapCtx, in.Padding)
	out.FontSize = direct.Enum_ToProto[pb.Text_TextStyle_FontSize](mapCtx, in.FontSize)
	out.PointerLocation = direct.Enum_ToProto[pb.Text_TextStyle_PointerLocation](mapCtx, in.PointerLocation)
	return out
}
func Threshold_FromProto(mapCtx *direct.MapContext, in *pb.Threshold) *krm.Threshold {
	if in == nil {
		return nil
	}
	out := &krm.Threshold{}
	out.Label = direct.LazyPtr(in.GetLabel())
	out.Value = direct.LazyPtr(in.GetValue())
	out.Color = direct.Enum_FromProto(mapCtx, in.Color)
	out.Direction = direct.Enum_FromProto(mapCtx, in.Direction)
	out.TargetAxis = direct.Enum_FromProto(mapCtx, in.TargetAxis)
	return out
}
func Threshold_ToProto(mapCtx *direct.MapContext, in *krm.Threshold) *pb.Threshold {
	if in == nil {
		return nil
	}
	out := &pb.Threshold{}
	out.Label = direct.ValueOf(in.Label)
	out.Value = direct.ValueOf(in.Value)
	out.Color = direct.Enum_ToProto[pb.Threshold_Color](mapCtx, in.Color)
	out.Direction = direct.Enum_ToProto[pb.Threshold_Direction](mapCtx, in.Direction)
	out.TargetAxis = direct.Enum_ToProto[pb.Threshold_TargetAxis](mapCtx, in.TargetAxis)
	return out
}
func TimeSeriesFilter_FromProto(mapCtx *direct.MapContext, in *pb.TimeSeriesFilter) *krm.TimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesFilter{}
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.Aggregation = Aggregation_FromProto(mapCtx, in.GetAggregation())
	out.SecondaryAggregation = Aggregation_FromProto(mapCtx, in.GetSecondaryAggregation())
	out.PickTimeSeriesFilter = PickTimeSeriesFilter_FromProto(mapCtx, in.GetPickTimeSeriesFilter())
	// MISSING: StatisticalTimeSeriesFilter
	return out
}
func TimeSeriesFilter_ToProto(mapCtx *direct.MapContext, in *krm.TimeSeriesFilter) *pb.TimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesFilter{}
	out.Filter = direct.ValueOf(in.Filter)
	out.Aggregation = Aggregation_ToProto(mapCtx, in.Aggregation)
	out.SecondaryAggregation = Aggregation_ToProto(mapCtx, in.SecondaryAggregation)
	if oneof := PickTimeSeriesFilter_ToProto(mapCtx, in.PickTimeSeriesFilter); oneof != nil {
		out.OutputFilter = &pb.TimeSeriesFilter_PickTimeSeriesFilter{PickTimeSeriesFilter: oneof}
	}
	// MISSING: StatisticalTimeSeriesFilter
	return out
}
func TimeSeriesFilterRatio_FromProto(mapCtx *direct.MapContext, in *pb.TimeSeriesFilterRatio) *krm.TimeSeriesFilterRatio {
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
func TimeSeriesFilterRatio_ToProto(mapCtx *direct.MapContext, in *krm.TimeSeriesFilterRatio) *pb.TimeSeriesFilterRatio {
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
func TimeSeriesFilterRatio_RatioPart_FromProto(mapCtx *direct.MapContext, in *pb.TimeSeriesFilterRatio_RatioPart) *krm.TimeSeriesFilterRatio_RatioPart {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesFilterRatio_RatioPart{}
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.Aggregation = Aggregation_FromProto(mapCtx, in.GetAggregation())
	return out
}
func TimeSeriesFilterRatio_RatioPart_ToProto(mapCtx *direct.MapContext, in *krm.TimeSeriesFilterRatio_RatioPart) *pb.TimeSeriesFilterRatio_RatioPart {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesFilterRatio_RatioPart{}
	out.Filter = direct.ValueOf(in.Filter)
	out.Aggregation = Aggregation_ToProto(mapCtx, in.Aggregation)
	return out
}
func TimeSeriesQuery_FromProto(mapCtx *direct.MapContext, in *pb.TimeSeriesQuery) *krm.TimeSeriesQuery {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesQuery{}
	out.TimeSeriesFilter = TimeSeriesFilter_FromProto(mapCtx, in.GetTimeSeriesFilter())
	out.TimeSeriesFilterRatio = TimeSeriesFilterRatio_FromProto(mapCtx, in.GetTimeSeriesFilterRatio())
	out.TimeSeriesQueryLanguage = direct.LazyPtr(in.GetTimeSeriesQueryLanguage())
	out.PrometheusQuery = direct.LazyPtr(in.GetPrometheusQuery())
	out.UnitOverride = direct.LazyPtr(in.GetUnitOverride())
	out.OutputFullDuration = direct.LazyPtr(in.GetOutputFullDuration())
	return out
}
func TimeSeriesQuery_ToProto(mapCtx *direct.MapContext, in *krm.TimeSeriesQuery) *pb.TimeSeriesQuery {
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
	if oneof := TimeSeriesQuery_PrometheusQuery_ToProto(mapCtx, in.PrometheusQuery); oneof != nil {
		out.Source = oneof
	}
	out.UnitOverride = direct.ValueOf(in.UnitOverride)
	out.OutputFullDuration = direct.ValueOf(in.OutputFullDuration)
	return out
}
func TimeSeriesTable_FromProto(mapCtx *direct.MapContext, in *pb.TimeSeriesTable) *krm.TimeSeriesTable {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesTable{}
	out.DataSets = direct.Slice_FromProto(mapCtx, in.DataSets, TimeSeriesTable_TableDataSet_FromProto)
	out.MetricVisualization = direct.Enum_FromProto(mapCtx, in.MetricVisualization)
	out.ColumnSettings = direct.Slice_FromProto(mapCtx, in.ColumnSettings, TimeSeriesTable_ColumnSettings_FromProto)
	return out
}
func TimeSeriesTable_ToProto(mapCtx *direct.MapContext, in *krm.TimeSeriesTable) *pb.TimeSeriesTable {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesTable{}
	out.DataSets = direct.Slice_ToProto(mapCtx, in.DataSets, TimeSeriesTable_TableDataSet_ToProto)
	out.MetricVisualization = direct.Enum_ToProto[pb.TimeSeriesTable_MetricVisualization](mapCtx, in.MetricVisualization)
	out.ColumnSettings = direct.Slice_ToProto(mapCtx, in.ColumnSettings, TimeSeriesTable_ColumnSettings_ToProto)
	return out
}

func TimeSeriesTable_ColumnSettings_ToProto(mapCtx *direct.MapContext, in *krm.TimeSeriesTable_ColumnSettings) *pb.TimeSeriesTable_ColumnSettings {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesTable_ColumnSettings{}
	out.Column = direct.ValueOf(in.Column)
	out.Visible = direct.ValueOf(in.Visible)
	return out
}
func TimeSeriesTable_TableDataSet_FromProto(mapCtx *direct.MapContext, in *pb.TimeSeriesTable_TableDataSet) *krm.TimeSeriesTable_TableDataSet {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesTable_TableDataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(mapCtx, in.GetTimeSeriesQuery())
	out.TableTemplate = direct.LazyPtr(in.GetTableTemplate())
	out.MinAlignmentPeriod = TableDataSet_MinAlignmentPeriod_FromProto(mapCtx, in.GetMinAlignmentPeriod())
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
	out.MinAlignmentPeriod = TableDataSet_MinAlignmentPeriod_ToProto(mapCtx, in.MinAlignmentPeriod)
	out.TableDisplayOptions = TableDisplayOptions_ToProto(mapCtx, in.TableDisplayOptions)
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

func XyChart_FromProto(mapCtx *direct.MapContext, in *pb.XyChart) *krm.XyChart {
	if in == nil {
		return nil
	}
	out := &krm.XyChart{}
	out.DataSets = direct.Slice_FromProto(mapCtx, in.DataSets, XyChart_DataSet_FromProto)
	out.TimeshiftDuration = XyChart_TimeshiftDuration_FromProto(mapCtx, in.GetTimeshiftDuration())
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
	out.TimeshiftDuration = XyChart_TimeshiftDuration_ToProto(mapCtx, in.TimeshiftDuration)
	out.Thresholds = direct.Slice_ToProto(mapCtx, in.Thresholds, Threshold_ToProto)
	out.XAxis = XyChart_Axis_ToProto(mapCtx, in.XAxis)
	out.YAxis = XyChart_Axis_ToProto(mapCtx, in.YAxis)
	out.Y2Axis = XyChart_Axis_ToProto(mapCtx, in.Y2Axis)
	out.ChartOptions = ChartOptions_ToProto(mapCtx, in.ChartOptions)
	return out
}
func XyChart_Axis_FromProto(mapCtx *direct.MapContext, in *pb.XyChart_Axis) *krm.XyChart_Axis {
	if in == nil {
		return nil
	}
	out := &krm.XyChart_Axis{}
	out.Label = direct.LazyPtr(in.GetLabel())
	out.Scale = direct.Enum_FromProto(mapCtx, in.Scale)
	return out
}
func XyChart_Axis_ToProto(mapCtx *direct.MapContext, in *krm.XyChart_Axis) *pb.XyChart_Axis {
	if in == nil {
		return nil
	}
	out := &pb.XyChart_Axis{}
	out.Label = direct.ValueOf(in.Label)
	out.Scale = direct.Enum_ToProto[pb.XyChart_Axis_Scale](mapCtx, in.Scale)
	return out
}
func XyChart_DataSet_FromProto(mapCtx *direct.MapContext, in *pb.XyChart_DataSet) *krm.XyChart_DataSet {
	if in == nil {
		return nil
	}
	out := &krm.XyChart_DataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(mapCtx, in.GetTimeSeriesQuery())
	out.PlotType = direct.Enum_FromProto(mapCtx, in.PlotType)
	out.LegendTemplate = direct.LazyPtr(in.GetLegendTemplate())
	out.MinAlignmentPeriod = DataSet_MinAlignmentPeriod_FromProto(mapCtx, in.GetMinAlignmentPeriod())
	out.TargetAxis = direct.Enum_FromProto(mapCtx, in.TargetAxis)
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
	out.MinAlignmentPeriod = DataSet_MinAlignmentPeriod_ToProto(mapCtx, in.MinAlignmentPeriod)
	out.TargetAxis = direct.Enum_ToProto[pb.XyChart_DataSet_TargetAxis](mapCtx, in.TargetAxis)
	return out
}
