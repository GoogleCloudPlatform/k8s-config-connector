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
	"math/rand"
	"strings"
	"testing"

	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	dashboardpb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/util/sets"
)

func FuzzMonitoringDashboardSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &dashboardpb.Dashboard{}
		fuzz.FillWithRandom(t, randStream, p1)

		// We don't expect output fields to round-trip
		outputFields := sets.New(".etag")

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New(
			".name",
			".labels",
		)

		// Widgets are under a few paths
		widgetPaths := []string{
			".grid_layout.widgets[]",
			".mosaic_layout.tiles[].widget",
			".column_layout.columns[].widgets[]",
			".row_layout.rows[].widgets[]",
		}
		for _, widgetPath := range widgetPaths {
			unimplementedFields.Insert(widgetPath + ".pie_chart.data_sets[].time_series_query.time_series_filter.statistical_time_series_filter")
			unimplementedFields.Insert(widgetPath + ".pie_chart.data_sets[].time_series_query.time_series_filter.pick_time_series_filter.interval")
			unimplementedFields.Insert(widgetPath + ".pie_chart.data_sets[].time_series_query.time_series_filter_ratio.statistical_time_series_filter")
			unimplementedFields.Insert(widgetPath + ".pie_chart.data_sets[].time_series_query.time_series_filter_ratio.pick_time_series_filter.interval")

			unimplementedFields.Insert(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter.statistical_time_series_filter")
			unimplementedFields.Insert(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter.pick_time_series_filter.interval")
			unimplementedFields.Insert(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter_ratio.statistical_time_series_filter")
			unimplementedFields.Insert(widgetPath + ".time_series_table.data_sets[].time_series_query.time_series_filter_ratio.pick_time_series_filter.interval")
		}

		// Remove any output only or known-unimplemented fields
		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields.Union(outputFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		// Force resource_names to be valid
		r := &fuzz.ReplaceFields{}
		r.Func = func(path string, val protoreflect.Value) (protoreflect.Value, bool) {
			// resource_names should be valid projects
			if strings.HasSuffix(path, ".resource_names[]") {
				return protoreflect.ValueOfString("projects/" + val.String()), true
			}
			// alignment_period only supports seconds
			if strings.HasSuffix(path, ".alignment_period.nanos") {
				return protoreflect.ValueOfInt32(0), true
			}
			// min_alignment_period only supports seconds
			if strings.HasSuffix(path, ".min_alignment_period.nanos") {
				return protoreflect.ValueOfInt32(0), true
			}
			return protoreflect.Value{}, false
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, r)

		ctx := &direct.MapContext{}
		k := MonitoringDashboardSpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := MonitoringDashboardSpec_ToProto(ctx, k)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from krm to proto: %v", ctx.Err())
		}

		if diff := cmp.Diff(p1, p2, protocmp.Transform()); diff != "" {
			t.Logf("p1 = %v", prototext.Format(p1))
			t.Logf("p2 = %v", prototext.Format(p2))
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}
	})
}

func FuzzMonitoringServiceSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.Service{}
		fuzz.FillWithRandom(t, randStream, p1)

		// We don't expect output fields to round-trip
		statusFields := sets.New(".etag")

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New(
			".name",
			".app_engine",
			".basic_service",
			".cloud_endpoints",
			".cloud_run",
			".cluster_istio",
			".istio_canonical_service",
			".gke_service",
			".gke_namespace",
			".gke_workload",
			".mesh_istio",
			".custom",
			".user_labels",
		)

		// Remove any output only or known-unimplemented fields
		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields.Union(statusFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		ctx := &direct.MapContext{}
		k := MonitoringServiceSpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := MonitoringServiceSpec_ToProto(ctx, k)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from krm to proto: %v", ctx.Err())
		}

		if diff := cmp.Diff(p1, p2, protocmp.Transform()); diff != "" {
			t.Logf("p1 = %v", prototext.Format(p1))
			t.Logf("p2 = %v", prototext.Format(p2))
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}
	})
}
