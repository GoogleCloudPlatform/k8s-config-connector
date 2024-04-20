package redis

import (
	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1alpha1"
)

func ClusterSpec_FromProto(ctx *MapContext, in *pb.Cluster) *krm.RedisClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.RedisClusterSpec{}
	out.ReplicaCount = (in.ReplicaCount)
	out.AuthorizationMode = Enum_FromProto(ctx, &in.AuthorizationMode)
	out.TransitEncryptionMode = Enum_FromProto(ctx, &in.TransitEncryptionMode)
	out.ShardCount = (in.ShardCount)
	out.PscConfigs = Slice_FromProto(ctx, in.PscConfigs, PscConfig_FromProto)
	return out
}
func ClusterSpec_ToProto(ctx *MapContext, in *krm.RedisClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.ReplicaCount = (in.ReplicaCount)
	out.AuthorizationMode = Enum_ToProto[pb.AuthorizationMode](ctx, in.AuthorizationMode)
	out.TransitEncryptionMode = Enum_ToProto[pb.TransitEncryptionMode](ctx, in.TransitEncryptionMode)
	out.ShardCount = (in.ShardCount)
	Slice_ToProto(ctx, in.PscConfigs, &out.PscConfigs, PscConfig_ToProto)
	return out
}

func ClusterState_FromProto(ctx *MapContext, in *pb.Cluster) *krm.RedisClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RedisClusterObservedState{}
	// out.CreateTime = Timestamp_FromProto(ctx, in.CreateTime)
	// out.State = Enum_FromProto(ctx, &in.State)
	// out.Uid = LazyPtr(in.Uid)
	out.SizeGb = (in.SizeGb)
	out.DiscoveryEndpoints = Slice_FromProto(ctx, in.DiscoveryEndpoints, DiscoveryEndpoint_FromProto)
	out.PscConnections = Slice_FromProto(ctx, in.PscConnections, PscConnection_FromProto)

	out.StateInfo = Cluster_StateInfo_FromProto(ctx, in.StateInfo)
	return out
}
func ClusterState_ToProto(ctx *MapContext, in *krm.RedisClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// out.CreateTime = Timestamp_ToProto(ctx, in.CreateTime)
	// out.State = Enum_FromProto(ctx, &in.State)
	// out.Uid = LazyPtr(in.Uid)
	out.SizeGb = (in.SizeGb)
	Slice_ToProto(ctx, in.DiscoveryEndpoints, &out.DiscoveryEndpoints, DiscoveryEndpoint_ToProto)
	Slice_ToProto(ctx, in.PscConnections, &out.PscConnections, PscConnection_ToProto)
	out.StateInfo = Cluster_StateInfo_ToProto(ctx, in.StateInfo)
	return out
}

func Cluster_StateInfo_FromProto(ctx *MapContext, in *pb.Cluster_StateInfo) *krm.Cluster_StateInfo {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_StateInfo{}
	out.UpdateInfo = Cluster_StateInfo_UpdateInfo_FromProto(ctx, in.GetUpdateInfo())
	return out
}
func Cluster_StateInfo_ToProto(ctx *MapContext, in *krm.Cluster_StateInfo) *pb.Cluster_StateInfo {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_StateInfo{}
	if in.UpdateInfo != nil {
		oneof := &pb.Cluster_StateInfo_UpdateInfo_{}
		out.Info = oneof
		oneof.UpdateInfo = Cluster_StateInfo_UpdateInfo_ToProto(ctx, in.UpdateInfo)
	}

	return out
}
func Cluster_StateInfo_UpdateInfo_FromProto(ctx *MapContext, in *pb.Cluster_StateInfo_UpdateInfo) *krm.Cluster_StateInfo_UpdateInfo {
	if in == nil {
		return nil
	}
	out := &krm.Cluster_StateInfo_UpdateInfo{}
	out.TargetShardCount = (in.TargetShardCount)
	out.TargetReplicaCount = (in.TargetReplicaCount)
	return out
}
func Cluster_StateInfo_UpdateInfo_ToProto(ctx *MapContext, in *krm.Cluster_StateInfo_UpdateInfo) *pb.Cluster_StateInfo_UpdateInfo {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_StateInfo_UpdateInfo{}
	out.TargetShardCount = in.TargetShardCount
	out.TargetReplicaCount = in.TargetReplicaCount
	return out
}
func PscConfig_FromProto(ctx *MapContext, in *pb.PscConfig) *krm.PscConfig {
	if in == nil {
		return nil
	}
	out := &krm.PscConfig{}
	out.Network = LazyPtr(in.Network)
	return out
}
func PscConfig_ToProto(ctx *MapContext, in *krm.PscConfig) *pb.PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.PscConfig{}
	out.Network = ValueOf(in.Network)
	return out
}
func DiscoveryEndpoint_FromProto(ctx *MapContext, in *pb.DiscoveryEndpoint) *krm.DiscoveryEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryEndpoint{}
	out.Address = LazyPtr(in.Address)
	out.Port = LazyPtr(in.Port)
	out.PscConfig = PscConfig_FromProto(ctx, in.PscConfig)
	return out
}
func DiscoveryEndpoint_ToProto(ctx *MapContext, in *krm.DiscoveryEndpoint) *pb.DiscoveryEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveryEndpoint{}
	out.Address = ValueOf(in.Address)
	out.Port = ValueOf(in.Port)
	out.PscConfig = PscConfig_ToProto(ctx, in.PscConfig)
	return out
}
func PscConnection_FromProto(ctx *MapContext, in *pb.PscConnection) *krm.PscConnection {
	if in == nil {
		return nil
	}
	out := &krm.PscConnection{}
	out.PscConnectionId = LazyPtr(in.PscConnectionId)
	out.Address = LazyPtr(in.Address)
	out.ForwardingRule = LazyPtr(in.ForwardingRule)
	out.ProjectId = LazyPtr(in.ProjectId)
	out.Network = LazyPtr(in.Network)
	return out
}
func PscConnection_ToProto(ctx *MapContext, in *krm.PscConnection) *pb.PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscConnection{}
	out.PscConnectionId = ValueOf(in.PscConnectionId)
	out.Address = ValueOf(in.Address)
	out.ForwardingRule = ValueOf(in.ForwardingRule)
	out.ProjectId = ValueOf(in.ProjectId)
	out.Network = ValueOf(in.Network)
	return out
}

// func DashboardSpec_ToProto(ctx *MapContext, in *krm.MonitoringDashboardSpec) *pb.Dashboard {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.Dashboard{}

// 	// out.Name = ValueOf(in.ResourceID)
// 	out.DisplayName = in.DisplayName

// 	if in.ColumnLayout != nil {
// 		oneof := &pb.Dashboard_ColumnLayout{}
// 		ColumnLayout_ToProto(ctx, in.ColumnLayout, &oneof.ColumnLayout)
// 		out.Layout = oneof
// 	}

// 	if in.GridLayout != nil {
// 		oneof := &pb.Dashboard_GridLayout{}
// 		GridLayout_ToProto(ctx, in.GridLayout, &oneof.GridLayout)
// 		out.Layout = oneof
// 	}

// 	if in.MosaicLayout != nil {
// 		oneof := &pb.Dashboard_MosaicLayout{}
// 		MosaicLayout_ToProto(ctx, in.MosaicLayout, &oneof.MosaicLayout)
// 		out.Layout = oneof
// 	}

// 	if in.RowLayout != nil {
// 		oneof := &pb.Dashboard_RowLayout{}
// 		RowLayout_ToProto(ctx, in.RowLayout, &oneof.RowLayout)
// 		out.Layout = oneof
// 	}

// 	// TOOD: More types

// 	// TODO: Filters
// 	// MapSlice(out.DashboardFilters, in.Filters, DashboardFilter_ToProto)

// 	// TODO: Labels

// 	return out
// }

// func DashboardSpec_FromProto(ctx *MapContext, in *pb.Dashboard) *krm.MonitoringDashboardSpec {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.MonitoringDashboardSpec{}

// 	// out.ResourceID = ValueOf(in.Name)
// 	out.DisplayName = in.DisplayName

// 	if oneof := in.GetColumnLayout(); oneof != nil {
// 		out.ColumnLayout = ColumnLayout_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetGridLayout(); oneof != nil {
// 		out.GridLayout = GridLayout_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetMosaicLayout(); oneof != nil {
// 		out.MosaicLayout = MosaicLayout_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetRowLayout(); oneof != nil {
// 		out.RowLayout = RowLayout_FromProto(ctx, oneof)
// 	}

// 	// TOOD: More types

// 	// TODO: Filters
// 	// MapSlice(out.DashboardFilters, in.Filters, DashboardFilter_ToProto)

// 	// TODO: Labels

// 	return out
// }

// func ColumnLayout_ToProto(ctx *MapContext, in *krm.DashboardColumnLayout, dest **pb.ColumnLayout) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.ColumnLayout{}
// 	Slice_ToProto(ctx, in.Columns, &out.Columns, ColumnLayout_Column_ToProto)
// 	*dest = out
// }

// func ColumnLayout_FromProto(ctx *MapContext, in *pb.ColumnLayout) *krm.DashboardColumnLayout {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardColumnLayout{}
// 	out.Columns = Slice_FromProto(ctx, in.Columns, ColumnLayout_Column_FromProto)
// 	return out

// }

// func ColumnLayout_Column_ToProto(ctx *MapContext, in *krm.DashboardColumns) *pb.ColumnLayout_Column {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.ColumnLayout_Column{}

// 	out.Weight = int64(ValueOf(in.Weight))
// 	Slice_ToProto(ctx, in.Widgets, &out.Widgets, DashboardWidgets_ToProto)
// 	return out
// }

// func ColumnLayout_Column_FromProto(ctx *MapContext, in *pb.ColumnLayout_Column) *krm.DashboardColumns {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardColumns{}

// 	out.Weight = LazyPtr(in.Weight)
// 	out.Widgets = Slice_FromProto(ctx, in.Widgets, DashboardWidgets_FromProto)
// 	return out
// }

// func GridLayout_ToProto(ctx *MapContext, in *krm.DashboardGridLayout, dest **pb.GridLayout) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.GridLayout{}
// 	*dest = out

// 	out.Columns = int64(ValueOf(in.Columns))
// 	Slice_ToProto(ctx, in.Widgets, &out.Widgets, DashboardWidgets_ToProto)
// }

// func GridLayout_FromProto(ctx *MapContext, in *pb.GridLayout) *krm.DashboardGridLayout {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardGridLayout{}

// 	out.Columns = LazyPtr(in.Columns)
// 	out.Widgets = Slice_FromProto(ctx, in.Widgets, DashboardWidgets_FromProto)
// 	return nil
// }

// func MosaicLayout_ToProto(ctx *MapContext, in *krm.DashboardMosaicLayout, dest **pb.MosaicLayout) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.MosaicLayout{}
// 	*dest = out

// 	out.Columns = int32(ValueOf(in.Columns))
// 	Slice_ToProto(ctx, in.Tiles, &out.Tiles, MosaicLayout_Tile_ToProto)
// }

// func MosaicLayout_FromProto(ctx *MapContext, in *pb.MosaicLayout) *krm.DashboardMosaicLayout {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardMosaicLayout{}

// 	out.Columns = LazyPtr(int64(in.Columns))
// 	out.Tiles = Slice_FromProto(ctx, in.Tiles, MosaicLayout_Tile_FromProto)

// 	return out
// }

// func RowLayout_ToProto(ctx *MapContext, in *krm.DashboardRowLayout, dest **pb.RowLayout) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.RowLayout{}
// 	*dest = out

// 	Slice_ToProto(ctx, in.Rows, &out.Rows, Rows_ToProto)
// }

// func RowLayout_FromProto(ctx *MapContext, in *pb.RowLayout) *krm.DashboardRowLayout {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardRowLayout{}

// 	out.Rows = Slice_FromProto(ctx, in.Rows, Rows_FromProto)
// 	return out
// }

// func Rows_ToProto(ctx *MapContext, in *krm.DashboardRows) *pb.RowLayout_Row {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.RowLayout_Row{}

// 	out.Weight = int64(ValueOf(in.Weight))
// 	Slice_ToProto(ctx, in.Widgets, &out.Widgets, DashboardWidgets_ToProto)
// 	return out
// }

// func Rows_FromProto(ctx *MapContext, in *pb.RowLayout_Row) *krm.DashboardRows {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardRows{}

// 	out.Weight = LazyPtr(in.Weight)
// 	out.Widgets = Slice_FromProto(ctx, in.Widgets, DashboardWidgets_FromProto)
// 	return out
// }

// func MosaicLayout_Tile_FromProto(ctx *MapContext, in *pb.MosaicLayout_Tile) *krm.MosaicLayout_Tile {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.MosaicLayout_Tile{}
// 	out.XPos = LazyPtr(in.XPos)
// 	out.YPos = LazyPtr(in.YPos)
// 	out.Width = LazyPtr(in.Width)
// 	out.Height = LazyPtr(in.Height)
// 	out.Widget = DashboardWidget_FromProto(ctx, in.Widget)
// 	return out
// }
// func MosaicLayout_Tile_ToProto(ctx *MapContext, in *krm.MosaicLayout_Tile) *pb.MosaicLayout_Tile {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.MosaicLayout_Tile{}
// 	out.XPos = ValueOf(in.XPos)
// 	out.YPos = ValueOf(in.YPos)
// 	out.Width = ValueOf(in.Width)
// 	out.Height = ValueOf(in.Height)
// 	out.Widget = DashboardWidget_ToProto(ctx, in.Widget)
// 	return out
// }

// func DashboardWidget_ToProto(ctx *MapContext, in *krm.DashboardWidget) *pb.Widget {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.Widget{}

// 	out.Title = ValueOf(in.Title)
// 	if in.Blank != nil {
// 		out.Content = &pb.Widget_Blank{
// 			Blank: &emptypb.Empty{},
// 		}
// 	}
// 	if in.LogsPanel != nil {
// 		oneof := &pb.Widget_LogsPanel{}
// 		LogsPanel_ToProto(ctx, in.LogsPanel, &oneof.LogsPanel)
// 		out.Content = oneof
// 	}

// 	if in.Scorecard != nil {
// 		oneof := &pb.Widget_Scorecard{}
// 		Scorecard_ToProto(ctx, in.Scorecard, &oneof.Scorecard)
// 		out.Content = oneof
// 	}

// 	if in.Text != nil {
// 		out.Content = &pb.Widget_Text{
// 			Text: DashboardText_ToProto(ctx, in.Text),
// 		}
// 	}

// 	if in.XyChart != nil {
// 		oneof := &pb.Widget_XyChart{}
// 		XyChart_ToProto(ctx, in.XyChart, &oneof.XyChart)
// 		out.Content = oneof
// 	}

// 	if in.AlertChart != nil {
// 		out.Content = &pb.Widget_AlertChart{
// 			AlertChart: AlertChart_ToProto(ctx, in.AlertChart),
// 		}
// 	}

// 	return out
// }

// func DashboardWidget_FromProto(ctx *MapContext, in *pb.Widget) *krm.DashboardWidget {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardWidget{}

// 	out.Title = LazyPtr(in.Title)
// 	if oneof := in.GetBlank(); oneof != nil {
// 		out.Blank = &krm.DashboardBlank{}
// 	}
// 	if oneof := in.GetLogsPanel(); oneof != nil {
// 		out.LogsPanel = LogsPanel_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetScorecard(); oneof != nil {
// 		out.Scorecard = Scorecard_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetText(); oneof != nil {
// 		out.Text = DashboardText_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetXyChart(); oneof != nil {
// 		out.XyChart = XyChart_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetAlertChart(); oneof != nil {
// 		out.AlertChart = AlertChart_FromProto(ctx, oneof)
// 	}

// 	return out
// }

// func DashboardWidgets_ToProto(ctx *MapContext, in *krm.DashboardWidgets) *pb.Widget {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.Widget{}

// 	out.Title = ValueOf(in.Title)
// 	if in.Blank != nil {
// 		out.Content = &pb.Widget_Blank{
// 			Blank: &emptypb.Empty{},
// 		}
// 	}
// 	if in.LogsPanel != nil {
// 		oneof := &pb.Widget_LogsPanel{}
// 		LogsPanel_ToProto(ctx, in.LogsPanel, &oneof.LogsPanel)
// 		out.Content = oneof
// 	}

// 	if in.Scorecard != nil {
// 		oneof := &pb.Widget_Scorecard{}
// 		Scorecard_ToProto(ctx, in.Scorecard, &oneof.Scorecard)
// 		out.Content = oneof
// 	}

// 	if in.Text != nil {
// 		out.Content = &pb.Widget_Text{
// 			Text: DashboardText_ToProto(ctx, in.Text),
// 		}
// 	}

// 	if in.XyChart != nil {
// 		oneof := &pb.Widget_XyChart{}
// 		XyChart_ToProto(ctx, in.XyChart, &oneof.XyChart)
// 		out.Content = oneof
// 	}

// 	if in.AlertChart != nil {
// 		out.Content = &pb.Widget_AlertChart{
// 			AlertChart: AlertChart_ToProto(ctx, in.AlertChart),
// 		}
// 	}

// 	return out
// }

// func DashboardWidgets_FromProto(ctx *MapContext, in *pb.Widget) *krm.DashboardWidgets {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardWidgets{}

// 	out.Title = LazyPtr(in.Title)
// 	if oneof := in.GetBlank(); oneof != nil {
// 		out.Blank = &krm.DashboardBlank{}
// 	}
// 	if oneof := in.GetLogsPanel(); oneof != nil {
// 		out.LogsPanel = LogsPanel_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetScorecard(); oneof != nil {
// 		out.Scorecard = Scorecard_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetText(); oneof != nil {
// 		out.Text = DashboardText_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetXyChart(); oneof != nil {
// 		out.XyChart = XyChart_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetAlertChart(); oneof != nil {
// 		out.AlertChart = AlertChart_FromProto(ctx, oneof)
// 	}

// 	return out
// }

// func DashboardText_ToProto(ctx *MapContext, in *krm.DashboardText) *pb.Text {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.Text{}

// 	out.Content = ValueOf(in.Content)
// 	Enum_ToProto(ctx, in.Format, &out.Format)
// 	return out
// }

// func DashboardText_FromProto(ctx *MapContext, in *pb.Text) *krm.DashboardText {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardText{}

// 	out.Content = LazyPtr(in.Content)
// 	out.Format = Enum_FromProto(ctx, &in.Format)

// 	return out
// }

// func XyChart_ToProto(ctx *MapContext, in *krm.DashboardXyChart, dest **pb.XyChart) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.XyChart{}
// 	*dest = out

// 	ChartOptions_ToProto(ctx, in.ChartOptions, &out.ChartOptions)

// 	Slice_ToProto(ctx, in.DataSets, &out.DataSets, DashboardDataSets_ToProto)

// 	Slice_ToProto(ctx, in.Thresholds, &out.Thresholds, Threshold_ToProto)

// 	Duration_ToProto(ctx, in.TimeshiftDuration, &out.TimeshiftDuration)

// 	DashboardXAxis_ToProto(ctx, in.XAxis, &out.XAxis)

// 	DashboardYAxis_ToProto(ctx, in.YAxis, &out.YAxis)
// }

// func XyChart_FromProto(ctx *MapContext, in *pb.XyChart) *krm.DashboardXyChart {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardXyChart{}

// 	out.ChartOptions = ChartOptions_FromProto(ctx, in.ChartOptions)

// 	out.DataSets = Slice_FromProto(ctx, in.DataSets, DashboardDataSets_FromProto)

// 	out.Thresholds = Slice_FromProto(ctx, in.Thresholds, Threshold_FromProto)

// 	out.TimeshiftDuration = Duration_FromProto(ctx, in.TimeshiftDuration)

// 	out.XAxis = DashboardXAxis_FromProto(ctx, in.XAxis)

// 	out.YAxis = DashboardYAxis_FromProto(ctx, in.YAxis)

// 	return out
// }

// func AlertChart_ToProto(ctx *MapContext, in *krm.DashboardAlertChart) *pb.AlertChart {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.AlertChart{}
// 	out.Name = ValueOf(in.Name)

// 	return out
// }

// func AlertChart_FromProto(ctx *MapContext, in *pb.AlertChart) *krm.DashboardAlertChart {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardAlertChart{}
// 	out.Name = LazyPtr(in.Name)

// 	return out
// }

// func ChartOptions_ToProto(ctx *MapContext, in *krm.DashboardChartOptions, dest **pb.ChartOptions) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.ChartOptions{}
// 	*dest = out

// 	Enum_ToProto(ctx, in.Mode, &out.Mode)
// }

// func ChartOptions_FromProto(ctx *MapContext, in *pb.ChartOptions) *krm.DashboardChartOptions {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardChartOptions{}
// 	out.Mode = Enum_FromProto(ctx, &in.Mode)
// 	return out
// }

// func DashboardDataSets_ToProto(ctx *MapContext, in *krm.DashboardDataSets) *pb.XyChart_DataSet {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.XyChart_DataSet{}

// 	out.TimeSeriesQuery = TimeSeriesQuery_ToProto(ctx, in.TimeSeriesQuery)
// 	Enum_ToProto(ctx, in.PlotType, &out.PlotType)
// 	out.LegendTemplate = ValueOf(in.LegendTemplate)
// 	Duration_ToProto(ctx, in.MinAlignmentPeriod, &out.MinAlignmentPeriod)

// 	// TODO: TargetAxis
// 	// // Optional. The target axis to use for plotting the metric.
// 	// TargetAxis XyChart_DataSet_TargetAxis `protobuf:"varint,5,opt,name=target_axis,json=targetAxis,proto3,enum=google.monitoring.dashboard.v1.XyChart_DataSet_TargetAxis" json:"target_axis,omitempty"`
// 	// Enum_ToProto(ctx, in.TargetAxis, &out.TargetAxis)

// 	return out
// }

// func DashboardDataSets_FromProto(ctx *MapContext, in *pb.XyChart_DataSet) *krm.DashboardDataSets {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardDataSets{}

// 	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(ctx, in.TimeSeriesQuery)
// 	out.PlotType = Enum_FromProto(ctx, &in.PlotType)
// 	out.LegendTemplate = LazyPtr(in.LegendTemplate)
// 	out.MinAlignmentPeriod = Duration_FromProto(ctx, in.MinAlignmentPeriod)

// 	// TODO: TargetAxis
// 	// // Optional. The target axis to use for plotting the metric.
// 	// TargetAxis XyChart_DataSet_TargetAxis `protobuf:"varint,5,opt,name=target_axis,json=targetAxis,proto3,enum=google.monitoring.dashboard.v1.XyChart_DataSet_TargetAxis" json:"target_axis,omitempty"`
// 	// Enum_ToProto(ctx, in.TargetAxis, &out.TargetAxis)

// 	return out
// }

// func DashboardXAxis_ToProto(ctx *MapContext, in *krm.DashboardXAxis, dest **pb.XyChart_Axis) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.XyChart_Axis{}
// 	*dest = out

// 	out.Label = ValueOf(in.Label)
// 	Enum_ToProto(ctx, in.Scale, &out.Scale)
// }

// func DashboardXAxis_FromProto(ctx *MapContext, in *pb.XyChart_Axis) *krm.DashboardXAxis {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardXAxis{}

// 	out.Label = LazyPtr(in.Label)
// 	out.Scale = Enum_FromProto(ctx, &in.Scale)

// 	return out
// }

// func DashboardYAxis_ToProto(ctx *MapContext, in *krm.DashboardYAxis, dest **pb.XyChart_Axis) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.XyChart_Axis{}
// 	*dest = out

// 	out.Label = ValueOf(in.Label)
// 	Enum_ToProto(ctx, in.Scale, &out.Scale)
// }

// func DashboardYAxis_FromProto(ctx *MapContext, in *pb.XyChart_Axis) *krm.DashboardYAxis {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardYAxis{}

// 	out.Label = LazyPtr(in.Label)
// 	out.Scale = Enum_FromProto(ctx, &in.Scale)

// 	return out
// }

// func Scorecard_ToProto(ctx *MapContext, in *krm.DashboardScorecard, dest **pb.Scorecard) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.Scorecard{}
// 	*dest = out

// 	if in.GaugeView != nil {
// 		oneof := &pb.Scorecard_GaugeView_{}
// 		GaugeView_ToProto(ctx, in.GaugeView, &oneof.GaugeView)
// 		out.DataView = oneof
// 	}

// 	if in.SparkChartView != nil {
// 		oneof := &pb.Scorecard_SparkChartView_{}
// 		SparkChartView_ToProto(ctx, in.SparkChartView, &oneof.SparkChartView)
// 		out.DataView = oneof
// 	}

// 	Slice_ToProto(ctx, in.Thresholds, &out.Thresholds, Threshold_ToProto)

// 	out.TimeSeriesQuery = TimeSeriesQuery_ToProto(ctx, in.TimeSeriesQuery)
// }

// func Scorecard_FromProto(ctx *MapContext, in *pb.Scorecard) *krm.DashboardScorecard {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardScorecard{}

// 	if oneof := in.GetGaugeView(); oneof != nil {
// 		out.GaugeView = GaugeView_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetSparkChartView(); oneof != nil {
// 		out.SparkChartView = SparkChartView_FromProto(ctx, oneof)
// 	}

// 	out.Thresholds = Slice_FromProto(ctx, in.Thresholds, Threshold_FromProto)

// 	out.TimeSeriesQuery = TimeSeriesQuery_FromProto(ctx, in.TimeSeriesQuery)

// 	return out
// }

// func Threshold_ToProto(ctx *MapContext, in *krm.DashboardThresholds) *pb.Threshold {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.Threshold{}

// 	out.Label = ValueOf(in.Label)
// 	out.Value = ValueOf(in.Value)
// 	Enum_ToProto(ctx, in.Color, &out.Color)
// 	Enum_ToProto(ctx, in.Direction, &out.Direction)

// 	Enum_ToProto(ctx, in.TargetAxis, &out.TargetAxis)

// 	return out
// }

// func Threshold_FromProto(ctx *MapContext, in *pb.Threshold) *krm.DashboardThresholds {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardThresholds{}

// 	out.Label = LazyPtr(in.Label)
// 	out.Value = LazyPtr(in.Value)
// 	out.Color = Enum_FromProto(ctx, &in.Color)
// 	out.Direction = Enum_FromProto(ctx, &in.Direction)

// 	out.TargetAxis = Enum_FromProto(ctx, &in.TargetAxis)

// 	return out
// }

// func SparkChartView_ToProto(ctx *MapContext, in *krm.DashboardSparkChartView, dest **pb.Scorecard_SparkChartView) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.Scorecard_SparkChartView{}
// 	*dest = out

// 	Enum_ToProto(ctx, &in.SparkChartType, &out.SparkChartType)
// 	SecondsString_ToProto(ctx, in.MinAlignmentPeriod, &out.MinAlignmentPeriod)
// }

// func SparkChartView_FromProto(ctx *MapContext, in *pb.Scorecard_SparkChartView) *krm.DashboardSparkChartView {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardSparkChartView{}

// 	out.SparkChartType = ValueOf(Enum_FromProto(ctx, &in.SparkChartType))
// 	out.MinAlignmentPeriod = SecondsString_FromProto(ctx, in.MinAlignmentPeriod)

// 	return out
// }
// func GaugeView_ToProto(ctx *MapContext, in *krm.DashboardGaugeView, dest **pb.Scorecard_GaugeView) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.Scorecard_GaugeView{}
// 	*dest = out

// 	out.LowerBound = ValueOf(in.LowerBound)
// 	out.UpperBound = ValueOf(in.UpperBound)
// }

// func GaugeView_FromProto(ctx *MapContext, in *pb.Scorecard_GaugeView) *krm.DashboardGaugeView {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardGaugeView{}

// 	out.LowerBound = LazyPtr(in.LowerBound)
// 	out.UpperBound = LazyPtr(in.UpperBound)

// 	return out
// }

// func TimeSeriesQuery_ToProto(ctx *MapContext, in *krm.DashboardTimeSeriesQuery) *pb.TimeSeriesQuery {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.TimeSeriesQuery{}

// 	out.UnitOverride = ValueOf(in.UnitOverride)

// 	if in.TimeSeriesFilter != nil {
// 		oneof := &pb.TimeSeriesQuery_TimeSeriesFilter{}
// 		TimeSeriesFilter_ToProto(ctx, in.TimeSeriesFilter, &oneof.TimeSeriesFilter)
// 		out.Source = oneof
// 	}

// 	if in.TimeSeriesFilterRatio != nil {
// 		oneof := &pb.TimeSeriesQuery_TimeSeriesFilterRatio{}
// 		TimeSeriesFilterRatio_ToProto(ctx, in.TimeSeriesFilterRatio, &oneof.TimeSeriesFilterRatio)
// 		out.Source = oneof
// 	}

// 	if in.TimeSeriesQueryLanguage != nil {
// 		out.Source = &pb.TimeSeriesQuery_TimeSeriesQueryLanguage{
// 			TimeSeriesQueryLanguage: ValueOf(in.TimeSeriesQueryLanguage),
// 		}
// 	}

// 	// TODO: Prometheus query

// 	return out

// }

// func TimeSeriesQuery_FromProto(ctx *MapContext, in *pb.TimeSeriesQuery) *krm.DashboardTimeSeriesQuery {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardTimeSeriesQuery{}

// 	out.UnitOverride = LazyPtr(in.UnitOverride)

// 	if oneof := in.GetTimeSeriesFilter(); oneof != nil {
// 		out.TimeSeriesFilter = TimeSeriesFilter_FromProto(ctx, oneof)
// 	}

// 	if oneof := in.GetTimeSeriesFilterRatio(); oneof != nil {
// 		out.TimeSeriesFilterRatio = TimeSeriesFilterRatio_FromProto(ctx, oneof)
// 	}

// 	if v := in.GetTimeSeriesQueryLanguage(); v != "" {
// 		out.TimeSeriesQueryLanguage = &v
// 	}

// 	// TODO: Prometheus query

// 	return out
// }

// func TimeSeriesFilterRatio_ToProto(ctx *MapContext, in *krm.DashboardTimeSeriesFilterRatio, dest **pb.TimeSeriesFilterRatio) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.TimeSeriesFilterRatio{}
// 	*dest = out

// 	DashboardNumerator_ToProto(ctx, in.Numerator, &out.Numerator)

// 	DashboardDenominator_ToProto(ctx, in.Denominator, &out.Denominator)

// 	out.SecondaryAggregation = &pb.Aggregation{}
// 	SecondaryAggregation_ToProto(ctx, in.SecondaryAggregation, &out.SecondaryAggregation)

// 	if in.PickTimeSeriesFilter != nil {
// 		oneof := &pb.TimeSeriesFilterRatio_PickTimeSeriesFilter{}
// 		PickTimeSeriesFilter_ToProto(ctx, in.PickTimeSeriesFilter, &oneof.PickTimeSeriesFilter)
// 		out.OutputFilter = oneof
// 	}

// 	// TODO: TimeSeriesFilterRatio_StatisticalTimeSeriesFilter
// }

// func TimeSeriesFilterRatio_FromProto(ctx *MapContext, in *pb.TimeSeriesFilterRatio) *krm.DashboardTimeSeriesFilterRatio {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardTimeSeriesFilterRatio{}

// 	out.Numerator = DashboardNumerator_FromProto(ctx, in.Numerator)
// 	out.Denominator = DashboardDenominator_FromProto(ctx, in.Denominator)

// 	out.SecondaryAggregation = SecondaryAggregation_FromProto(ctx, in.SecondaryAggregation)

// 	if oneof := in.GetPickTimeSeriesFilter(); oneof != nil {
// 		out.PickTimeSeriesFilter = PickTimeSeriesFilter_FromProto(ctx, oneof)
// 	}

// 	return out
// }

// func DashboardNumerator_ToProto(ctx *MapContext, in *krm.DashboardNumerator, dest **pb.TimeSeriesFilterRatio_RatioPart) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.TimeSeriesFilterRatio_RatioPart{}
// 	*dest = out

// 	out.Filter = in.Filter
// 	Aggregation_ToProto(ctx, in.Aggregation, &out.Aggregation)
// }

// func DashboardNumerator_FromProto(ctx *MapContext, in *pb.TimeSeriesFilterRatio_RatioPart) *krm.DashboardNumerator {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardNumerator{}

// 	out.Filter = in.Filter
// 	out.Aggregation = Aggregation_FromProto(ctx, in.Aggregation)
// 	return out
// }

// func DashboardDenominator_ToProto(ctx *MapContext, in *krm.DashboardDenominator, dest **pb.TimeSeriesFilterRatio_RatioPart) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.TimeSeriesFilterRatio_RatioPart{}
// 	*dest = out

// 	out.Filter = in.Filter
// 	Aggregation_ToProto(ctx, in.Aggregation, &out.Aggregation)
// }

// func DashboardDenominator_FromProto(ctx *MapContext, in *pb.TimeSeriesFilterRatio_RatioPart) *krm.DashboardDenominator {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardDenominator{}

// 	out.Filter = in.Filter
// 	out.Aggregation = Aggregation_FromProto(ctx, in.Aggregation)
// 	return out
// }

// func PickTimeSeriesFilter_ToProto(ctx *MapContext, in *krm.DashboardPickTimeSeriesFilter, dest **pb.PickTimeSeriesFilter) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.PickTimeSeriesFilter{}
// 	*dest = out

// 	out.NumTimeSeries = int32(ValueOf(in.NumTimeSeries))
// 	Enum_ToProto(ctx, in.RankingMethod, &out.RankingMethod)
// 	Enum_ToProto(ctx, in.Direction, &out.Direction)
// }

// func PickTimeSeriesFilter_FromProto(ctx *MapContext, in *pb.PickTimeSeriesFilter) *krm.DashboardPickTimeSeriesFilter {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardPickTimeSeriesFilter{}

// 	out.NumTimeSeries = LazyPtr(int64(in.NumTimeSeries))
// 	out.RankingMethod = Enum_FromProto(ctx, &in.RankingMethod)
// 	out.Direction = Enum_FromProto(ctx, &in.Direction)

// 	return out
// }

// func TimeSeriesFilter_ToProto(ctx *MapContext, in *krm.DashboardTimeSeriesFilter, dest **pb.TimeSeriesFilter) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.TimeSeriesFilter{}
// 	*dest = out

// 	out.Filter = in.Filter

// 	Aggregation_ToProto(ctx, in.Aggregation, &out.Aggregation)

// 	SecondaryAggregation_ToProto(ctx, in.SecondaryAggregation, &out.SecondaryAggregation)

// 	if in.PickTimeSeriesFilter != nil {
// 		oneof := &pb.TimeSeriesFilter_PickTimeSeriesFilter{}
// 		PickTimeSeriesFilter_ToProto(ctx, in.PickTimeSeriesFilter, &oneof.PickTimeSeriesFilter)
// 		out.OutputFilter = oneof
// 	}

// 	// TODO: TimeSeriesFilter_StatisticalTimeSeriesFilter
// }

// func TimeSeriesFilter_FromProto(ctx *MapContext, in *pb.TimeSeriesFilter) *krm.DashboardTimeSeriesFilter {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardTimeSeriesFilter{}

// 	out.Filter = in.Filter

// 	out.Aggregation = Aggregation_FromProto(ctx, in.Aggregation)

// 	out.SecondaryAggregation = SecondaryAggregation_FromProto(ctx, in.SecondaryAggregation)

// 	if oneof := in.GetPickTimeSeriesFilter(); oneof != nil {
// 		out.PickTimeSeriesFilter = PickTimeSeriesFilter_FromProto(ctx, oneof)
// 	}

// 	// TODO: TimeSeriesFilter_StatisticalTimeSeriesFilter

// 	return out
// }

// func Aggregation_ToProto(ctx *MapContext, in *krm.DashboardAggregation, dest **pb.Aggregation) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.Aggregation{}
// 	*dest = out

// 	SecondsString_ToProto(ctx, in.AlignmentPeriod, &out.AlignmentPeriod)
// 	Enum_ToProto(ctx, in.PerSeriesAligner, &out.PerSeriesAligner)
// 	Enum_ToProto(ctx, in.CrossSeriesReducer, &out.CrossSeriesReducer)
// 	out.GroupByFields = in.GroupByFields
// }

// func Aggregation_FromProto(ctx *MapContext, in *pb.Aggregation) *krm.DashboardAggregation {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardAggregation{}

// 	out.AlignmentPeriod = SecondsString_FromProto(ctx, in.AlignmentPeriod)
// 	out.PerSeriesAligner = Enum_FromProto(ctx, &in.PerSeriesAligner)
// 	out.CrossSeriesReducer = Enum_FromProto(ctx, &in.CrossSeriesReducer)
// 	out.GroupByFields = in.GroupByFields

// 	return out
// }

// func SecondaryAggregation_ToProto(ctx *MapContext, in *krm.DashboardSecondaryAggregation, dest **pb.Aggregation) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.Aggregation{}
// 	*dest = out

// 	SecondsString_ToProto(ctx, in.AlignmentPeriod, &out.AlignmentPeriod)
// 	Enum_ToProto(ctx, in.PerSeriesAligner, &out.PerSeriesAligner)
// 	Enum_ToProto(ctx, in.CrossSeriesReducer, &out.CrossSeriesReducer)
// 	out.GroupByFields = in.GroupByFields
// }

// func SecondaryAggregation_FromProto(ctx *MapContext, in *pb.Aggregation) *krm.DashboardSecondaryAggregation {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardSecondaryAggregation{}

// 	out.AlignmentPeriod = SecondsString_FromProto(ctx, in.AlignmentPeriod)
// 	out.PerSeriesAligner = Enum_FromProto(ctx, &in.PerSeriesAligner)
// 	out.CrossSeriesReducer = Enum_FromProto(ctx, &in.CrossSeriesReducer)
// 	out.GroupByFields = in.GroupByFields

// 	return out
// }

// func SecondsString_ToProto(ctx *MapContext, in *string, dest **durationpb.Duration) {
// 	if in == nil {
// 		return
// 	}
// 	n, err := strconv.ParseInt(*in, 10, 64)
// 	if err != nil {
// 		ctx.Errorf("expected an integer, got %q", *in)
// 	}
// 	out := &durationpb.Duration{
// 		Seconds: n,
// 	}
// 	*dest = out
// }

// func SecondsString_FromProto(ctx *MapContext, in *durationpb.Duration) *string {
// 	if in == nil {
// 		return nil
// 	}
// 	seconds := in.Seconds
// 	s := strconv.FormatInt(seconds, 10)
// 	return &s
// }

// func LogsPanel_ToProto(ctx *MapContext, in *krm.DashboardLogsPanel, dest **pb.LogsPanel) {
// 	if in == nil {
// 		return
// 	}
// 	out := &pb.LogsPanel{}
// 	*dest = out

// 	out.Filter = ValueOf(in.Filter)

// 	for _, resource := range in.ResourceNames {
// 		resolved, err := ctx.ResolveRef(&resource)
// 		if err != nil {
// 			ctx.Errorf("resolving reference: %w", err)
// 			continue
// 		}

// 		out.ResourceNames = append(out.ResourceNames, ValueOf(resolved.External))
// 	}
// }

// func LogsPanel_FromProto(ctx *MapContext, in *pb.LogsPanel) *krm.DashboardLogsPanel {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DashboardLogsPanel{}

// 	out.Filter = LazyPtr(in.Filter)

// 	for _, resource := range in.ResourceNames {
// 		resource := resource
// 		ref := krm.DashboardResourceNames{
// 			External: &resource,
// 		}
// 		out.ResourceNames = append(out.ResourceNames, ref)
// 	}

// 	return out
// }

// func Dashboard_ToStatus(ctx *MapContext, in *pb.Dashboard) *krm.MonitoringDashboardStatus {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.MonitoringDashboardStatus{}
// 	return out
// }
