// Copyright 2026 Google LLC
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

package migrationcenter

import (
	"math"

	pb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/migrationcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ReportSummary_HistogramChartData_Bucket_FromProto(mapCtx *direct.MapContext, in *pb.ReportSummary_HistogramChartData_Bucket) *krm.ReportSummary_HistogramChartData_Bucket {
	if in == nil {
		return nil
	}
	out := &krm.ReportSummary_HistogramChartData_Bucket{}
	out.LowerBound = direct.LazyPtr(in.GetLowerBound())
	if in.GetUpperBound() != math.MaxInt64 {
		out.UpperBound = direct.LazyPtr(in.GetUpperBound())
	}
	out.Count = direct.LazyPtr(in.GetCount())
	return out
}

func ReportSummary_HistogramChartData_Bucket_ToProto(mapCtx *direct.MapContext, in *krm.ReportSummary_HistogramChartData_Bucket) *pb.ReportSummary_HistogramChartData_Bucket {
	if in == nil {
		return nil
	}
	out := &pb.ReportSummary_HistogramChartData_Bucket{}
	out.LowerBound = direct.ValueOf(in.LowerBound)
	out.UpperBound = direct.ValueOf(in.UpperBound)
	out.Count = direct.ValueOf(in.Count)
	return out
}
