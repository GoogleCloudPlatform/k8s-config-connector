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

// +tool:fuzz-gen
// proto.message: google.monitoring.dashboard.v1.Dashboard
// krm.group: monitoring.cnrm.cloud.google.com
// krm.kind: MonitoringDashboard

package monitoring

import (
	"strings"

	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dashboardFuzzer())
}

func dashboardFuzzer() fuzztesting.KRMFuzzer {
	fuzzer := fuzztesting.NewKRMTypedFuzzer[*pb.Dashboard, krm.MonitoringDashboardSpec, krm.MonitoringDashboardStatus](
		&pb.Dashboard{},
		MonitoringDashboardSpec_FromProto,
		MonitoringDashboardSpec_ToProto,
		MonitoringDashboardStatus_FromProto,
		MonitoringDashboardStatus_ToProto,
	)

	fuzzer.FilterSpec = func(in *pb.Dashboard) {
		fuzztesting.VisitValues(in, func(path string, fd protoreflect.FieldDescriptor, v protoreflect.Value) protoreflect.Value {
			if fd.Kind() == protoreflect.MessageKind && !fd.IsList() && !fd.IsMap() {
				switch val := v.Message().Interface().(type) {
				case *durationpb.Duration:
					if strings.HasSuffix(path, ".min_alignment_period") || strings.HasSuffix(path, ".alignment_period") {
						val.Nanos = 0
					}

				case *pb.TimeSeriesFilterRatio:
					// StatisticalTimeSeriesFilter is deprecated upstream, so don't try to map it.
					if val.GetStatisticalTimeSeriesFilter() != nil {
						val.OutputFilter = nil
					}

				case *pb.TimeSeriesFilter:
					// StatisticalTimeSeriesFilter is deprecated upstream, so don't try to map it.
					if val.GetStatisticalTimeSeriesFilter() != nil {
						val.OutputFilter = nil
					}
				}
			}
			return v
		})
	}

	// Field comparison between KRM Spec (MonitoringDashboardSpec) and GCP Proto (google.monitoring.dashboard.v1.Dashboard):
	// - ProjectRef (KRM Spec) maps to GCP project URL segment, not part of standard Dashboard fields.
	// - ResourceID (KRM Spec) maps to the name segment of the Dashboard GCP name.
	// - DisplayName (KRM Spec) maps to `.display_name` (GCP Proto).
	// - GridLayout (KRM Spec) maps to `.grid_layout` (GCP Proto).
	// - MosaicLayout (KRM Spec) maps to `.mosaic_layout` (GCP Proto).
	// - RowLayout (KRM Spec) maps to `.row_layout` (GCP Proto).
	// - ColumnLayout (KRM Spec) maps to `.column_layout` (GCP Proto).
	// - DashboardFilters (KRM Spec) maps to `.dashboard_filters` (GCP Proto).

	fuzzer.SpecField(".display_name")
	fuzzer.SpecField(".grid_layout")
	fuzzer.SpecField(".mosaic_layout")
	fuzzer.SpecField(".row_layout")
	fuzzer.SpecField(".column_layout")
	fuzzer.SpecField(".dashboard_filters")

	// Identity and Unimplemented fields
	fuzzer.Unimplemented_Identity(".name")
	fuzzer.Unimplemented_NotYetTriaged(".labels")

	widgetPaths := []string{
		".grid_layout.widgets[]",
		".mosaic_layout.tiles[].widget",
		".column_layout.columns[].widgets[]",
		".row_layout.rows[].widgets[]",
	}
	for _, widgetPath := range widgetPaths {
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".pie_chart.data_sets[].time_series_query.time_series_filter.statistical_time_series_filter")
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".pie_chart.data_sets[].time_series_query.time_series_filter.pick_time_series_filter.interval")
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".pie_chart.data_sets[].time_series_query.time_series_filter_ratio.statistical_time_series_filter")
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".pie_chart.data_sets[].time_series_query.time_series_filter_ratio.pick_time_series_filter.interval")

		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".scorecard.time_series_query.time_series_filter.statistical_time_series_filter")
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".scorecard.time_series_query.time_series_filter.pick_time_series_filter.interval")
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".scorecard.time_series_query.time_series_filter_ratio.statistical_time_series_filter")
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".scorecard.time_series_query.time_series_filter_ratio.pick_time_series_filter.interval")

		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter.statistical_time_series_filter")
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter.pick_time_series_filter.interval")
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter_ratio.statistical_time_series_filter")
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter_ratio.pick_time_series_filter.interval")

		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".xy_chart.data_sets[].time_series_query.time_series_filter.statistical_time_series_filter")
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".xy_chart.data_sets[].time_series_query.time_series_filter.pick_time_series_filter.interval")
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".xy_chart.data_sets[].time_series_query.time_series_filter_ratio.statistical_time_series_filter")
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".xy_chart.data_sets[].time_series_query.time_series_filter_ratio.pick_time_series_filter.interval")

		// Resource names need to be in a special format.
		fuzzer.Unimplemented_NotYetTriaged(widgetPath + ".logs_panel.resource_names") // the proto 'resource_names' needs to be a special format.
	}

	// Status fields
	fuzzer.StatusField(".etag")

	return fuzzer
}
