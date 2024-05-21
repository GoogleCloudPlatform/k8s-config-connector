package monitoring

import (
	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
)

func DashboardAggregation_FromProto(ctx *MapContext, in *pb.Aggregation) *krm.DashboardAggregation {
	if in == nil {
		return nil
	}
	out := &krm.DashboardAggregation{}
	out.AlignmentPeriod = Aggregation_AlignmentPeriod_FromProto(ctx, in.GetAlignmentPeriod())
	out.PerSeriesAligner = Enum_FromProto(ctx, in.PerSeriesAligner)
	out.CrossSeriesReducer = Enum_FromProto(ctx, in.CrossSeriesReducer)
	out.GroupByFields = in.GroupByFields
	return out
}
func DashboardAggregation_ToProto(ctx *MapContext, in *krm.DashboardAggregation) *pb.Aggregation {
	if in == nil {
		return nil
	}
	out := &pb.Aggregation{}
	out.AlignmentPeriod = Aggregation_AlignmentPeriod_ToProto(ctx, in.AlignmentPeriod)
	out.PerSeriesAligner = Enum_ToProto[pb.Aggregation_Aligner](ctx, in.PerSeriesAligner)
	out.CrossSeriesReducer = Enum_ToProto[pb.Aggregation_Reducer](ctx, in.CrossSeriesReducer)
	out.GroupByFields = in.GroupByFields
	return out
}
func DashboardChartOptions_FromProto(ctx *MapContext, in *pb.ChartOptions) *krm.DashboardChartOptions {
	if in == nil {
		return nil
	}
	out := &krm.DashboardChartOptions{}
	out.Mode = Enum_FromProto(ctx, in.Mode)
	return out
}
func DashboardChartOptions_ToProto(ctx *MapContext, in *krm.DashboardChartOptions) *pb.ChartOptions {
	if in == nil {
		return nil
	}
	out := &pb.ChartOptions{}
	out.Mode = Enum_ToProto[pb.ChartOptions_Mode](ctx, in.Mode)
	return out
}
func DashboardColumnLayout_FromProto(ctx *MapContext, in *pb.ColumnLayout) *krm.DashboardColumnLayout {
	if in == nil {
		return nil
	}
	out := &krm.DashboardColumnLayout{}
	out.Columns = Slice_FromProto(ctx, in.Columns, DashboardColumns_FromProto)
	return out
}
func DashboardColumnLayout_ToProto(ctx *MapContext, in *krm.DashboardColumnLayout) *pb.ColumnLayout {
	if in == nil {
		return nil
	}
	out := &pb.ColumnLayout{}
	out.Columns = Slice_ToProto(ctx, in.Columns, DashboardColumns_ToProto)
	return out
}
func DashboardColumns_FromProto(ctx *MapContext, in *pb.ColumnLayout_Column) *krm.DashboardColumns {
	if in == nil {
		return nil
	}
	out := &krm.DashboardColumns{}
	out.Weight = LazyPtr(in.GetWeight())
	out.Widgets = Slice_FromProto(ctx, in.Widgets, DashboardWidget_FromProto)
	return out
}
func DashboardColumns_ToProto(ctx *MapContext, in *krm.DashboardColumns) *pb.ColumnLayout_Column {
	if in == nil {
		return nil
	}
	out := &pb.ColumnLayout_Column{}
	out.Weight = ValueOf(in.Weight)
	out.Widgets = Slice_ToProto(ctx, in.Widgets, DashboardWidget_ToProto)
	return out
}
func DashboardDataSets_FromProto(ctx *MapContext, in *pb.XyChart_DataSet) *krm.DashboardDataSets {
	if in == nil {
		return nil
	}
	out := &krm.DashboardDataSets{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(ctx, in.GetTimeSeriesQuery())
	out.PlotType = Enum_FromProto(ctx, in.PlotType)
	out.LegendTemplate = LazyPtr(in.GetLegendTemplate())
	out.MinAlignmentPeriod = DataSet_MinAlignmentPeriod_FromProto(ctx, in.GetMinAlignmentPeriod())
	// MISSING: TargetAxis
	return out
}
func DashboardDataSets_ToProto(ctx *MapContext, in *krm.DashboardDataSets) *pb.XyChart_DataSet {
	if in == nil {
		return nil
	}
	out := &pb.XyChart_DataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_ToProto(ctx, in.TimeSeriesQuery)
	out.PlotType = Enum_ToProto[pb.XyChart_DataSet_PlotType](ctx, in.PlotType)
	out.LegendTemplate = ValueOf(in.LegendTemplate)
	out.MinAlignmentPeriod = DataSet_MinAlignmentPeriod_ToProto(ctx, in.MinAlignmentPeriod)
	// MISSING: TargetAxis
	return out
}
func DashboardDenominator_FromProto(ctx *MapContext, in *pb.TimeSeriesFilterRatio_RatioPart) *krm.DashboardDenominator {
	if in == nil {
		return nil
	}
	out := &krm.DashboardDenominator{}
	out.Filter = LazyPtr(in.GetFilter())
	out.Aggregation = DashboardAggregation_FromProto(ctx, in.GetAggregation())
	return out
}
func DashboardDenominator_ToProto(ctx *MapContext, in *krm.DashboardDenominator) *pb.TimeSeriesFilterRatio_RatioPart {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesFilterRatio_RatioPart{}
	out.Filter = ValueOf(in.Filter)
	out.Aggregation = DashboardAggregation_ToProto(ctx, in.Aggregation)
	return out
}
func DashboardGaugeView_FromProto(ctx *MapContext, in *pb.Scorecard_GaugeView) *krm.DashboardGaugeView {
	if in == nil {
		return nil
	}
	out := &krm.DashboardGaugeView{}
	out.LowerBound = LazyPtr(in.GetLowerBound())
	out.UpperBound = LazyPtr(in.GetUpperBound())
	return out
}
func DashboardGaugeView_ToProto(ctx *MapContext, in *krm.DashboardGaugeView) *pb.Scorecard_GaugeView {
	if in == nil {
		return nil
	}
	out := &pb.Scorecard_GaugeView{}
	out.LowerBound = ValueOf(in.LowerBound)
	out.UpperBound = ValueOf(in.UpperBound)
	return out
}
func DashboardGridLayout_FromProto(ctx *MapContext, in *pb.GridLayout) *krm.DashboardGridLayout {
	if in == nil {
		return nil
	}
	out := &krm.DashboardGridLayout{}
	out.Columns = LazyPtr(in.GetColumns())
	out.Widgets = Slice_FromProto(ctx, in.Widgets, DashboardWidget_FromProto)
	return out
}
func DashboardGridLayout_ToProto(ctx *MapContext, in *krm.DashboardGridLayout) *pb.GridLayout {
	if in == nil {
		return nil
	}
	out := &pb.GridLayout{}
	out.Columns = ValueOf(in.Columns)
	out.Widgets = Slice_ToProto(ctx, in.Widgets, DashboardWidget_ToProto)
	return out
}
func DashboardLogsPanel_FromProto(ctx *MapContext, in *pb.LogsPanel) *krm.DashboardLogsPanel {
	if in == nil {
		return nil
	}
	out := &krm.DashboardLogsPanel{}
	out.Filter = LazyPtr(in.GetFilter())
	out.ResourceNames = DashboardLogsPanel_ResourceNames_FromProto(ctx, in.ResourceNames)
	return out
}
func DashboardLogsPanel_ToProto(ctx *MapContext, in *krm.DashboardLogsPanel) *pb.LogsPanel {
	if in == nil {
		return nil
	}
	out := &pb.LogsPanel{}
	out.Filter = ValueOf(in.Filter)
	out.ResourceNames = DashboardLogsPanel_ResourceNames_ToProto(ctx, in.ResourceNames)
	return out
}
func DashboardMosaicLayout_FromProto(ctx *MapContext, in *pb.MosaicLayout) *krm.DashboardMosaicLayout {
	if in == nil {
		return nil
	}
	out := &krm.DashboardMosaicLayout{}
	out.Columns = LazyPtr(in.GetColumns())
	out.Tiles = Slice_FromProto(ctx, in.Tiles, DashboardTiles_FromProto)
	return out
}
func DashboardMosaicLayout_ToProto(ctx *MapContext, in *krm.DashboardMosaicLayout) *pb.MosaicLayout {
	if in == nil {
		return nil
	}
	out := &pb.MosaicLayout{}
	out.Columns = ValueOf(in.Columns)
	out.Tiles = Slice_ToProto(ctx, in.Tiles, DashboardTiles_ToProto)
	return out
}
func DashboardNumerator_FromProto(ctx *MapContext, in *pb.TimeSeriesFilterRatio_RatioPart) *krm.DashboardNumerator {
	if in == nil {
		return nil
	}
	out := &krm.DashboardNumerator{}
	out.Filter = LazyPtr(in.GetFilter())
	out.Aggregation = DashboardAggregation_FromProto(ctx, in.GetAggregation())
	return out
}
func DashboardNumerator_ToProto(ctx *MapContext, in *krm.DashboardNumerator) *pb.TimeSeriesFilterRatio_RatioPart {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesFilterRatio_RatioPart{}
	out.Filter = ValueOf(in.Filter)
	out.Aggregation = DashboardAggregation_ToProto(ctx, in.Aggregation)
	return out
}
func DashboardPickTimeSeriesFilter_FromProto(ctx *MapContext, in *pb.PickTimeSeriesFilter) *krm.DashboardPickTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &krm.DashboardPickTimeSeriesFilter{}
	out.RankingMethod = Enum_FromProto(ctx, in.RankingMethod)
	out.NumTimeSeries = LazyPtr(in.GetNumTimeSeries())
	out.Direction = Enum_FromProto(ctx, in.Direction)
	// MISSING: Interval
	return out
}
func DashboardPickTimeSeriesFilter_ToProto(ctx *MapContext, in *krm.DashboardPickTimeSeriesFilter) *pb.PickTimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &pb.PickTimeSeriesFilter{}
	out.RankingMethod = Enum_ToProto[pb.PickTimeSeriesFilter_Method](ctx, in.RankingMethod)
	out.NumTimeSeries = ValueOf(in.NumTimeSeries)
	out.Direction = Enum_ToProto[pb.PickTimeSeriesFilter_Direction](ctx, in.Direction)
	// MISSING: Interval
	return out
}
func DashboardRowLayout_FromProto(ctx *MapContext, in *pb.RowLayout) *krm.DashboardRowLayout {
	if in == nil {
		return nil
	}
	out := &krm.DashboardRowLayout{}
	out.Rows = Slice_FromProto(ctx, in.Rows, DashboardRows_FromProto)
	return out
}
func DashboardRowLayout_ToProto(ctx *MapContext, in *krm.DashboardRowLayout) *pb.RowLayout {
	if in == nil {
		return nil
	}
	out := &pb.RowLayout{}
	out.Rows = Slice_ToProto(ctx, in.Rows, DashboardRows_ToProto)
	return out
}
func DashboardRows_FromProto(ctx *MapContext, in *pb.RowLayout_Row) *krm.DashboardRows {
	if in == nil {
		return nil
	}
	out := &krm.DashboardRows{}
	out.Weight = LazyPtr(in.GetWeight())
	out.Widgets = Slice_FromProto(ctx, in.Widgets, DashboardWidget_FromProto)
	return out
}
func DashboardRows_ToProto(ctx *MapContext, in *krm.DashboardRows) *pb.RowLayout_Row {
	if in == nil {
		return nil
	}
	out := &pb.RowLayout_Row{}
	out.Weight = ValueOf(in.Weight)
	out.Widgets = Slice_ToProto(ctx, in.Widgets, DashboardWidget_ToProto)
	return out
}
func DashboardScorecard_FromProto(ctx *MapContext, in *pb.Scorecard) *krm.DashboardScorecard {
	if in == nil {
		return nil
	}
	out := &krm.DashboardScorecard{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(ctx, in.GetTimeSeriesQuery())
	out.GaugeView = DashboardGaugeView_FromProto(ctx, in.GetGaugeView())
	out.SparkChartView = DashboardSparkChartView_FromProto(ctx, in.GetSparkChartView())
	// MISSING: BlankView
	out.Thresholds = Slice_FromProto(ctx, in.Thresholds, DashboardThresholds_FromProto)
	return out
}
func DashboardScorecard_ToProto(ctx *MapContext, in *krm.DashboardScorecard) *pb.Scorecard {
	if in == nil {
		return nil
	}
	out := &pb.Scorecard{}
	out.TimeSeriesQuery = TimeSeriesQuery_ToProto(ctx, in.TimeSeriesQuery)
	if oneof := DashboardGaugeView_ToProto(ctx, in.GaugeView); oneof != nil {
		out.DataView = &pb.Scorecard_GaugeView_{GaugeView: oneof}
	}
	if oneof := DashboardSparkChartView_ToProto(ctx, in.SparkChartView); oneof != nil {
		out.DataView = &pb.Scorecard_SparkChartView_{SparkChartView: oneof}
	}
	// MISSING: BlankView
	out.Thresholds = Slice_ToProto(ctx, in.Thresholds, DashboardThresholds_ToProto)
	return out
}
func DashboardSparkChartView_FromProto(ctx *MapContext, in *pb.Scorecard_SparkChartView) *krm.DashboardSparkChartView {
	if in == nil {
		return nil
	}
	out := &krm.DashboardSparkChartView{}
	out.SparkChartType = Enum_FromProto(ctx, in.SparkChartType)
	out.MinAlignmentPeriod = SparkChartView_MinAlignmentPeriod_FromProto(ctx, in.GetMinAlignmentPeriod())
	return out
}
func DashboardSparkChartView_ToProto(ctx *MapContext, in *krm.DashboardSparkChartView) *pb.Scorecard_SparkChartView {
	if in == nil {
		return nil
	}
	out := &pb.Scorecard_SparkChartView{}
	out.SparkChartType = Enum_ToProto[pb.SparkChartType](ctx, in.SparkChartType)
	out.MinAlignmentPeriod = SparkChartView_MinAlignmentPeriod_ToProto(ctx, in.MinAlignmentPeriod)
	return out
}
func DashboardText_FromProto(ctx *MapContext, in *pb.Text) *krm.DashboardText {
	if in == nil {
		return nil
	}
	out := &krm.DashboardText{}
	out.Content = LazyPtr(in.GetContent())
	out.Format = Enum_FromProto(ctx, in.Format)
	out.Style = Text_TextStyle_FromProto(ctx, in.GetStyle())
	return out
}
func DashboardText_ToProto(ctx *MapContext, in *krm.DashboardText) *pb.Text {
	if in == nil {
		return nil
	}
	out := &pb.Text{}
	out.Content = ValueOf(in.Content)
	out.Format = Enum_ToProto[pb.Text_Format](ctx, in.Format)
	out.Style = Text_TextStyle_ToProto(ctx, in.Style)
	return out
}
func DashboardThresholds_FromProto(ctx *MapContext, in *pb.Threshold) *krm.DashboardThresholds {
	if in == nil {
		return nil
	}
	out := &krm.DashboardThresholds{}
	out.Label = LazyPtr(in.GetLabel())
	out.Value = LazyPtr(in.GetValue())
	out.Color = Enum_FromProto(ctx, in.Color)
	out.Direction = Enum_FromProto(ctx, in.Direction)
	// MISSING: TargetAxis
	return out
}
func DashboardThresholds_ToProto(ctx *MapContext, in *krm.DashboardThresholds) *pb.Threshold {
	if in == nil {
		return nil
	}
	out := &pb.Threshold{}
	out.Label = ValueOf(in.Label)
	out.Value = ValueOf(in.Value)
	out.Color = Enum_ToProto[pb.Threshold_Color](ctx, in.Color)
	out.Direction = Enum_ToProto[pb.Threshold_Direction](ctx, in.Direction)
	// MISSING: TargetAxis
	return out
}
func DashboardTiles_FromProto(ctx *MapContext, in *pb.MosaicLayout_Tile) *krm.DashboardTiles {
	if in == nil {
		return nil
	}
	out := &krm.DashboardTiles{}
	out.XPos = LazyPtr(in.GetXPos())
	out.YPos = LazyPtr(in.GetYPos())
	out.Width = LazyPtr(in.GetWidth())
	out.Height = LazyPtr(in.GetHeight())
	out.Widget = DashboardWidget_FromProto(ctx, in.GetWidget())
	return out
}
func DashboardTiles_ToProto(ctx *MapContext, in *krm.DashboardTiles) *pb.MosaicLayout_Tile {
	if in == nil {
		return nil
	}
	out := &pb.MosaicLayout_Tile{}
	out.XPos = ValueOf(in.XPos)
	out.YPos = ValueOf(in.YPos)
	out.Width = ValueOf(in.Width)
	out.Height = ValueOf(in.Height)
	out.Widget = DashboardWidget_ToProto(ctx, in.Widget)
	return out
}
func DashboardTimeSeriesFilterRatio_FromProto(ctx *MapContext, in *pb.TimeSeriesFilterRatio) *krm.DashboardTimeSeriesFilterRatio {
	if in == nil {
		return nil
	}
	out := &krm.DashboardTimeSeriesFilterRatio{}
	out.Numerator = DashboardNumerator_FromProto(ctx, in.GetNumerator())
	out.Denominator = DashboardDenominator_FromProto(ctx, in.GetDenominator())
	out.SecondaryAggregation = DashboardAggregation_FromProto(ctx, in.GetSecondaryAggregation())
	out.PickTimeSeriesFilter = DashboardPickTimeSeriesFilter_FromProto(ctx, in.GetPickTimeSeriesFilter())
	// MISSING: StatisticalTimeSeriesFilter
	return out
}
func DashboardTimeSeriesFilterRatio_ToProto(ctx *MapContext, in *krm.DashboardTimeSeriesFilterRatio) *pb.TimeSeriesFilterRatio {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesFilterRatio{}
	out.Numerator = DashboardNumerator_ToProto(ctx, in.Numerator)
	out.Denominator = DashboardDenominator_ToProto(ctx, in.Denominator)
	out.SecondaryAggregation = DashboardAggregation_ToProto(ctx, in.SecondaryAggregation)
	if oneof := DashboardPickTimeSeriesFilter_ToProto(ctx, in.PickTimeSeriesFilter); oneof != nil {
		out.OutputFilter = &pb.TimeSeriesFilterRatio_PickTimeSeriesFilter{PickTimeSeriesFilter: oneof}
	}
	// MISSING: StatisticalTimeSeriesFilter
	return out
}
func DashboardWidget_FromProto(ctx *MapContext, in *pb.Widget) *krm.DashboardWidget {
	if in == nil {
		return nil
	}
	out := &krm.DashboardWidget{}
	out.Title = LazyPtr(in.GetTitle())
	out.XyChart = DashboardXyChart_FromProto(ctx, in.GetXyChart())
	out.Scorecard = DashboardScorecard_FromProto(ctx, in.GetScorecard())
	out.Text = DashboardText_FromProto(ctx, in.GetText())
	out.Blank = DashboardBlank_FromProto(ctx, in.GetBlank())
	// MISSING: AlertChart
	out.TimeSeriesTable = TimeSeriesTable_FromProto(ctx, in.GetTimeSeriesTable())
	// MISSING: CollapsibleGroup
	out.LogsPanel = DashboardLogsPanel_FromProto(ctx, in.GetLogsPanel())
	// MISSING: IncidentList
	// MISSING: PieChart
	// MISSING: ErrorReportingPanel
	// MISSING: SectionHeader
	// MISSING: SingleViewGroup
	// MISSING: Id
	return out
}
func DashboardWidget_ToProto(ctx *MapContext, in *krm.DashboardWidget) *pb.Widget {
	if in == nil {
		return nil
	}
	out := &pb.Widget{}
	out.Title = ValueOf(in.Title)
	if oneof := DashboardXyChart_ToProto(ctx, in.XyChart); oneof != nil {
		out.Content = &pb.Widget_XyChart{XyChart: oneof}
	}
	if oneof := DashboardScorecard_ToProto(ctx, in.Scorecard); oneof != nil {
		out.Content = &pb.Widget_Scorecard{Scorecard: oneof}
	}
	if oneof := DashboardText_ToProto(ctx, in.Text); oneof != nil {
		out.Content = &pb.Widget_Text{Text: oneof}
	}
	if oneof := DashboardBlank_ToProto(ctx, in.Blank); oneof != nil {
		out.Content = &pb.Widget_Blank{Blank: oneof}
	}
	// MISSING: AlertChart
	if oneof := TimeSeriesTable_ToProto(ctx, in.TimeSeriesTable); oneof != nil {
		out.Content = &pb.Widget_TimeSeriesTable{TimeSeriesTable: oneof}
	}
	// MISSING: CollapsibleGroup
	if oneof := DashboardLogsPanel_ToProto(ctx, in.LogsPanel); oneof != nil {
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
func DashboardXAxis_FromProto(ctx *MapContext, in *pb.XyChart_Axis) *krm.DashboardXAxis {
	if in == nil {
		return nil
	}
	out := &krm.DashboardXAxis{}
	out.Label = LazyPtr(in.GetLabel())
	out.Scale = Enum_FromProto(ctx, in.Scale)
	return out
}
func DashboardXAxis_ToProto(ctx *MapContext, in *krm.DashboardXAxis) *pb.XyChart_Axis {
	if in == nil {
		return nil
	}
	out := &pb.XyChart_Axis{}
	out.Label = ValueOf(in.Label)
	out.Scale = Enum_ToProto[pb.XyChart_Axis_Scale](ctx, in.Scale)
	return out
}
func DashboardXyChart_FromProto(ctx *MapContext, in *pb.XyChart) *krm.DashboardXyChart {
	if in == nil {
		return nil
	}
	out := &krm.DashboardXyChart{}
	out.DataSets = Slice_FromProto(ctx, in.DataSets, DashboardDataSets_FromProto)
	out.TimeshiftDuration = XyChart_TimeshiftDuration_FromProto(ctx, in.GetTimeshiftDuration())
	out.Thresholds = Slice_FromProto(ctx, in.Thresholds, DashboardThresholds_FromProto)
	out.XAxis = DashboardXAxis_FromProto(ctx, in.GetXAxis())
	out.YAxis = DashboardYAxis_FromProto(ctx, in.GetYAxis())
	// MISSING: Y2Axis
	out.ChartOptions = DashboardChartOptions_FromProto(ctx, in.GetChartOptions())
	return out
}
func DashboardXyChart_ToProto(ctx *MapContext, in *krm.DashboardXyChart) *pb.XyChart {
	if in == nil {
		return nil
	}
	out := &pb.XyChart{}
	out.DataSets = Slice_ToProto(ctx, in.DataSets, DashboardDataSets_ToProto)
	out.TimeshiftDuration = XyChart_TimeshiftDuration_ToProto(ctx, in.TimeshiftDuration)
	out.Thresholds = Slice_ToProto(ctx, in.Thresholds, DashboardThresholds_ToProto)
	out.XAxis = DashboardXAxis_ToProto(ctx, in.XAxis)
	out.YAxis = DashboardYAxis_ToProto(ctx, in.YAxis)
	// MISSING: Y2Axis
	out.ChartOptions = DashboardChartOptions_ToProto(ctx, in.ChartOptions)
	return out
}
func DashboardYAxis_FromProto(ctx *MapContext, in *pb.XyChart_Axis) *krm.DashboardYAxis {
	if in == nil {
		return nil
	}
	out := &krm.DashboardYAxis{}
	out.Label = LazyPtr(in.GetLabel())
	out.Scale = Enum_FromProto(ctx, in.Scale)
	return out
}
func DashboardYAxis_ToProto(ctx *MapContext, in *krm.DashboardYAxis) *pb.XyChart_Axis {
	if in == nil {
		return nil
	}
	out := &pb.XyChart_Axis{}
	out.Label = ValueOf(in.Label)
	out.Scale = Enum_ToProto[pb.XyChart_Axis_Scale](ctx, in.Scale)
	return out
}
func MonitoringDashboardSpec_FromProto(ctx *MapContext, in *pb.Dashboard) *krm.MonitoringDashboardSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringDashboardSpec{}
	// MISSING: Name
	out.DisplayName = LazyPtr(in.GetDisplayName())
	// MISSING: Etag
	out.GridLayout = DashboardGridLayout_FromProto(ctx, in.GetGridLayout())
	out.MosaicLayout = DashboardMosaicLayout_FromProto(ctx, in.GetMosaicLayout())
	out.RowLayout = DashboardRowLayout_FromProto(ctx, in.GetRowLayout())
	out.ColumnLayout = DashboardColumnLayout_FromProto(ctx, in.GetColumnLayout())
	// MISSING: DashboardFilters
	// MISSING: Labels
	return out
}
func MonitoringDashboardSpec_ToProto(ctx *MapContext, in *krm.MonitoringDashboardSpec) *pb.Dashboard {
	if in == nil {
		return nil
	}
	out := &pb.Dashboard{}
	// MISSING: Name
	out.DisplayName = ValueOf(in.DisplayName)
	// MISSING: Etag
	if oneof := DashboardGridLayout_ToProto(ctx, in.GridLayout); oneof != nil {
		out.Layout = &pb.Dashboard_GridLayout{GridLayout: oneof}
	}
	if oneof := DashboardMosaicLayout_ToProto(ctx, in.MosaicLayout); oneof != nil {
		out.Layout = &pb.Dashboard_MosaicLayout{MosaicLayout: oneof}
	}
	if oneof := DashboardRowLayout_ToProto(ctx, in.RowLayout); oneof != nil {
		out.Layout = &pb.Dashboard_RowLayout{RowLayout: oneof}
	}
	if oneof := DashboardColumnLayout_ToProto(ctx, in.ColumnLayout); oneof != nil {
		out.Layout = &pb.Dashboard_ColumnLayout{ColumnLayout: oneof}
	}
	// MISSING: DashboardFilters
	// MISSING: Labels
	return out
}
func MonitoringDashboardStatus_FromProto(ctx *MapContext, in *pb.Dashboard) *krm.MonitoringDashboardStatus {
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
func MonitoringDashboardStatus_ToProto(ctx *MapContext, in *krm.MonitoringDashboardStatus) *pb.Dashboard {
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
func TableDisplayOptions_FromProto(ctx *MapContext, in *pb.TableDisplayOptions) *krm.TableDisplayOptions {
	if in == nil {
		return nil
	}
	out := &krm.TableDisplayOptions{}
	// MISSING: ShownColumns
	return out
}
func TableDisplayOptions_ToProto(ctx *MapContext, in *krm.TableDisplayOptions) *pb.TableDisplayOptions {
	if in == nil {
		return nil
	}
	out := &pb.TableDisplayOptions{}
	// MISSING: ShownColumns
	return out
}
func Text_TextStyle_FromProto(ctx *MapContext, in *pb.Text_TextStyle) *krm.Text_TextStyle {
	if in == nil {
		return nil
	}
	out := &krm.Text_TextStyle{}
	out.BackgroundColor = LazyPtr(in.GetBackgroundColor())
	out.TextColor = LazyPtr(in.GetTextColor())
	out.HorizontalAlignment = Enum_FromProto(ctx, in.HorizontalAlignment)
	out.VerticalAlignment = Enum_FromProto(ctx, in.VerticalAlignment)
	out.Padding = Enum_FromProto(ctx, in.Padding)
	out.FontSize = Enum_FromProto(ctx, in.FontSize)
	out.PointerLocation = Enum_FromProto(ctx, in.PointerLocation)
	return out
}
func Text_TextStyle_ToProto(ctx *MapContext, in *krm.Text_TextStyle) *pb.Text_TextStyle {
	if in == nil {
		return nil
	}
	out := &pb.Text_TextStyle{}
	out.BackgroundColor = ValueOf(in.BackgroundColor)
	out.TextColor = ValueOf(in.TextColor)
	out.HorizontalAlignment = Enum_ToProto[pb.Text_TextStyle_HorizontalAlignment](ctx, in.HorizontalAlignment)
	out.VerticalAlignment = Enum_ToProto[pb.Text_TextStyle_VerticalAlignment](ctx, in.VerticalAlignment)
	out.Padding = Enum_ToProto[pb.Text_TextStyle_PaddingSize](ctx, in.Padding)
	out.FontSize = Enum_ToProto[pb.Text_TextStyle_FontSize](ctx, in.FontSize)
	out.PointerLocation = Enum_ToProto[pb.Text_TextStyle_PointerLocation](ctx, in.PointerLocation)
	return out
}
func TimeSeriesFilter_FromProto(ctx *MapContext, in *pb.TimeSeriesFilter) *krm.TimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesFilter{}
	out.Filter = LazyPtr(in.GetFilter())
	out.Aggregation = DashboardAggregation_FromProto(ctx, in.GetAggregation())
	out.SecondaryAggregation = DashboardAggregation_FromProto(ctx, in.GetSecondaryAggregation())
	out.PickTimeSeriesFilter = DashboardPickTimeSeriesFilter_FromProto(ctx, in.GetPickTimeSeriesFilter())
	// MISSING: StatisticalTimeSeriesFilter
	return out
}
func TimeSeriesFilter_ToProto(ctx *MapContext, in *krm.TimeSeriesFilter) *pb.TimeSeriesFilter {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesFilter{}
	out.Filter = ValueOf(in.Filter)
	out.Aggregation = DashboardAggregation_ToProto(ctx, in.Aggregation)
	out.SecondaryAggregation = DashboardAggregation_ToProto(ctx, in.SecondaryAggregation)
	if oneof := DashboardPickTimeSeriesFilter_ToProto(ctx, in.PickTimeSeriesFilter); oneof != nil {
		out.OutputFilter = &pb.TimeSeriesFilter_PickTimeSeriesFilter{PickTimeSeriesFilter: oneof}
	}
	// MISSING: StatisticalTimeSeriesFilter
	return out
}
func TimeSeriesQuery_FromProto(ctx *MapContext, in *pb.TimeSeriesQuery) *krm.TimeSeriesQuery {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesQuery{}
	out.TimeSeriesFilter = TimeSeriesFilter_FromProto(ctx, in.GetTimeSeriesFilter())
	out.TimeSeriesFilterRatio = DashboardTimeSeriesFilterRatio_FromProto(ctx, in.GetTimeSeriesFilterRatio())
	out.TimeSeriesQueryLanguage = LazyPtr(in.GetTimeSeriesQueryLanguage())
	// MISSING: PrometheusQuery
	out.UnitOverride = LazyPtr(in.GetUnitOverride())
	// MISSING: OutputFullDuration
	return out
}
func TimeSeriesQuery_ToProto(ctx *MapContext, in *krm.TimeSeriesQuery) *pb.TimeSeriesQuery {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesQuery{}
	if oneof := TimeSeriesFilter_ToProto(ctx, in.TimeSeriesFilter); oneof != nil {
		out.Source = &pb.TimeSeriesQuery_TimeSeriesFilter{TimeSeriesFilter: oneof}
	}
	if oneof := DashboardTimeSeriesFilterRatio_ToProto(ctx, in.TimeSeriesFilterRatio); oneof != nil {
		out.Source = &pb.TimeSeriesQuery_TimeSeriesFilterRatio{TimeSeriesFilterRatio: oneof}
	}
	if oneof := TimeSeriesQuery_TimeSeriesQueryLanguage_ToProto(ctx, in.TimeSeriesQueryLanguage); oneof != nil {
		out.Source = oneof
	}
	// MISSING: PrometheusQuery
	out.UnitOverride = ValueOf(in.UnitOverride)
	// MISSING: OutputFullDuration
	return out
}
func TimeSeriesTable_FromProto(ctx *MapContext, in *pb.TimeSeriesTable) *krm.TimeSeriesTable {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesTable{}
	out.DataSets = Slice_FromProto(ctx, in.DataSets, TimeSeriesTable_TableDataSet_FromProto)
	out.MetricVisualization = Enum_FromProto(ctx, in.MetricVisualization)
	out.ColumnSettings = Slice_FromProto(ctx, in.ColumnSettings, TimeSeriesTable_ColumnSettings_FromProto)
	return out
}
func TimeSeriesTable_ToProto(ctx *MapContext, in *krm.TimeSeriesTable) *pb.TimeSeriesTable {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesTable{}
	out.DataSets = Slice_ToProto(ctx, in.DataSets, TimeSeriesTable_TableDataSet_ToProto)
	out.MetricVisualization = Enum_ToProto[pb.TimeSeriesTable_MetricVisualization](ctx, in.MetricVisualization)
	out.ColumnSettings = Slice_ToProto(ctx, in.ColumnSettings, TimeSeriesTable_ColumnSettings_ToProto)
	return out
}
func TimeSeriesTable_ColumnSettings_FromProto(ctx *MapContext, in *pb.TimeSeriesTable_ColumnSettings) *krm.TimeSeriesTable_ColumnSettings {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesTable_ColumnSettings{}
	out.Column = LazyPtr(in.GetColumn())
	out.Visible = LazyPtr(in.GetVisible())
	return out
}
func TimeSeriesTable_ColumnSettings_ToProto(ctx *MapContext, in *krm.TimeSeriesTable_ColumnSettings) *pb.TimeSeriesTable_ColumnSettings {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesTable_ColumnSettings{}
	out.Column = ValueOf(in.Column)
	out.Visible = ValueOf(in.Visible)
	return out
}
func TimeSeriesTable_TableDataSet_FromProto(ctx *MapContext, in *pb.TimeSeriesTable_TableDataSet) *krm.TimeSeriesTable_TableDataSet {
	if in == nil {
		return nil
	}
	out := &krm.TimeSeriesTable_TableDataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(ctx, in.GetTimeSeriesQuery())
	out.TableTemplate = LazyPtr(in.GetTableTemplate())
	out.MinAlignmentPeriod = TableDataSet_MinAlignmentPeriod_FromProto(ctx, in.GetMinAlignmentPeriod())
	out.TableDisplayOptions = TableDisplayOptions_FromProto(ctx, in.GetTableDisplayOptions())
	return out
}
func TimeSeriesTable_TableDataSet_ToProto(ctx *MapContext, in *krm.TimeSeriesTable_TableDataSet) *pb.TimeSeriesTable_TableDataSet {
	if in == nil {
		return nil
	}
	out := &pb.TimeSeriesTable_TableDataSet{}
	out.TimeSeriesQuery = TimeSeriesQuery_ToProto(ctx, in.TimeSeriesQuery)
	out.TableTemplate = ValueOf(in.TableTemplate)
	out.MinAlignmentPeriod = TableDataSet_MinAlignmentPeriod_ToProto(ctx, in.MinAlignmentPeriod)
	out.TableDisplayOptions = TableDisplayOptions_ToProto(ctx, in.TableDisplayOptions)
	return out
}
