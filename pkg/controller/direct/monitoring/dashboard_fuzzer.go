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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
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

		fuzzer.UnimplementedFields.Insert(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter.statistical_time_series_filter")
		fuzzer.UnimplementedFields.Insert(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter.pick_time_series_filter.interval")
		fuzzer.UnimplementedFields.Insert(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter_ratio.statistical_time_series_filter")
		fuzzer.UnimplementedFields.Insert(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter_ratio.pick_time_series_filter.interval")
	}

	fuzzer.StatusFields.Insert(".etag")

	return fuzzer
}