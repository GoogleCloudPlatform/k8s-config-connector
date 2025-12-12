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

	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(dashboardFuzzer())
}

func dashboardFuzzer() fuzztesting.KRMFuzzer {
	fuzzer := fuzztesting.NewKRMTypedSpecFuzzer(
		&pb.Dashboard{},
		MonitoringDashboardSpec_FromProto,
		MonitoringDashboardSpec_ToProto,
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

	fuzzer.UnimplementedFields.Insert(".name")
	fuzzer.UnimplementedFields.Insert(".labels")

	widgetPaths := []string{
		".grid_layout.widgets[]",
		".mosaic_layout.tiles[].widget",
		".column_layout.columns[].widgets[]",
		".row_layout.rows[].widgets[]",
	}
	for _, widgetPath := range widgetPaths {
		fuzzer.UnimplementedFields.Insert(widgetPath + ".pie_chart.data_sets[].time_series_query.time_series_filter.statistical_time_series_filter")
		fuzzer.UnimplementedFields.Insert(widgetPath + ".pie_chart.data_sets[].time_series_query.time_series_filter.pick_time_series_filter.interval")
		fuzzer.UnimplementedFields.Insert(widgetPath + ".pie_chart.data_sets[].time_series_query.time_series_filter_ratio.statistical_time_series_filter")
		fuzzer.UnimplementedFields.Insert(widgetPath + ".pie_chart.data_sets[].time_series_query.time_series_filter_ratio.pick_time_series_filter.interval")

		fuzzer.UnimplementedFields.Insert(widgetPath + ".scorecard.time_series_query.time_series_filter.statistical_time_series_filter")
		fuzzer.UnimplementedFields.Insert(widgetPath + ".scorecard.time_series_query.time_series_filter.pick_time_series_filter.interval")
		fuzzer.UnimplementedFields.Insert(widgetPath + ".scorecard.time_series_query.time_series_filter_ratio.statistical_time_series_filter")
		fuzzer.UnimplementedFields.Insert(widgetPath + ".scorecard.time_series_query.time_series_filter_ratio.pick_time_series_filter.interval")

		fuzzer.UnimplementedFields.Insert(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter.statistical_time_series_filter")
		fuzzer.UnimplementedFields.Insert(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter.pick_time_series_filter.interval")
		fuzzer.UnimplementedFields.Insert(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter_ratio.statistical_time_series_filter")
		fuzzer.UnimplementedFields.Insert(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter_ratio.pick_time_series_filter.interval")

		// Resource names need to be in a special format.
		fuzzer.UnimplementedFields.Insert(widgetPath + ".logs_panel.resource_names") // the proto 'resource_names' needs to be a special format.
	}

	fuzzer.StatusFields.Insert(".etag")

	return fuzzer
}
