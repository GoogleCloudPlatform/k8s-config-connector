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

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
)

func Empty_FromProto(mapCtx *MapContext, in *emptypb.Empty) *krm.Empty {
	if in == nil {
		return nil
	}
	out := &krm.Empty{}
	return out
}
func Empty_ToProto(mapCtx *MapContext, in *krm.Empty) *emptypb.Empty {
	if in == nil {
		return nil
	}
	out := &emptypb.Empty{}
	return out
}

func AlertChart_FromProto(mapCtx *MapContext, in *pb.AlertChart) *krm.AlertChart {
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
func AlertChart_ToProto(mapCtx *MapContext, in *krm.AlertChart) *pb.AlertChart {
	if in == nil {
		return nil
	}
	out := &pb.AlertChart{}
	if in.AlertPolicyRef != nil {
		out.Name = in.AlertPolicyRef.External
	}
	return out
}

func Aggregation_AlignmentPeriod_FromProto(mapCtx *MapContext, in *durationpb.Duration) *string {
	return SecondsString_FromProto(mapCtx, in)
}

func Aggregation_AlignmentPeriod_ToProto(mapCtx *MapContext, in *string) *durationpb.Duration {
	return SecondsString_ToProto(mapCtx, in, "alignmentPeriod")
}

// TODO: The format is not documented, we need to validate
func DataSet_MinAlignmentPeriod_FromProto(mapCtx *MapContext, in *durationpb.Duration) *string {
	return SecondsString_FromProto(mapCtx, in)
}

// TODO: The format is not documented, we need to validate
func DataSet_MinAlignmentPeriod_ToProto(mapCtx *MapContext, in *string) *durationpb.Duration {
	return SecondsString_ToProto(mapCtx, in, "minAlignmentPeriod")
}

// TODO: The format is not documented, we need to validate
func SparkChartView_MinAlignmentPeriod_FromProto(mapCtx *MapContext, in *durationpb.Duration) *string {
	return SecondsString_FromProto(mapCtx, in)
}

// TODO: The format is not documented, we need to validate
func SparkChartView_MinAlignmentPeriod_ToProto(mapCtx *MapContext, in *string) *durationpb.Duration {
	return SecondsString_ToProto(mapCtx, in, "minAlignmentPeriod")
}

// TODO: The format is not documented, we need to validate
func XyChart_TimeshiftDuration_FromProto(mapCtx *MapContext, in *durationpb.Duration) *string {
	return Duration_FromProto(mapCtx, in)
}

// TODO: The format is not documented, we need to validate
func XyChart_TimeshiftDuration_ToProto(mapCtx *MapContext, in *string) *durationpb.Duration {
	return Duration_ToProto(mapCtx, in)
}

func PieChartDataSet_MinAlignmentPeriod_FromProto(mapCtx *MapContext, in *durationpb.Duration) *string {
	return Duration_FromProto(mapCtx, in)
}

func PieChartDataSet_MinAlignmentPeriod_ToProto(mapCtx *MapContext, in *string) *durationpb.Duration {
	return Duration_ToProto(mapCtx, in)
}

func TableDataSet_MinAlignmentPeriod_FromProto(mapCtx *MapContext, in *durationpb.Duration) *string {
	return Duration_FromProto(mapCtx, in)
}

func TableDataSet_MinAlignmentPeriod_ToProto(mapCtx *MapContext, in *string) *durationpb.Duration {
	return Duration_ToProto(mapCtx, in)
}

func TimeSeriesQuery_TimeSeriesQueryLanguage_ToProto(mapCtx *MapContext, in *string) *pb.TimeSeriesQuery_TimeSeriesQueryLanguage {
	if in == nil {
		return nil
	}

	return &pb.TimeSeriesQuery_TimeSeriesQueryLanguage{
		TimeSeriesQueryLanguage: ValueOf(in),
	}
}

func LogsPanel_ResourceNames_FromProto(mapCtx *MapContext, in []string) []v1alpha1.ResourceRef {
	if in == nil {
		return nil
	}
	var out []v1alpha1.ResourceRef
	for _, v := range in {
		tokens := strings.Split(v, "/")
		if len(tokens) == 2 && tokens[0] == "projects" {
			out = append(out, v1alpha1.ResourceRef{
				Kind:     "Project",
				External: v,
			})
		} else {
			mapCtx.Errorf("resourceName %q was not recognized", v)
		}
	}
	return out
}

func LogsPanel_ResourceNames_ToProto(mapCtx *MapContext, in []v1alpha1.ResourceRef) []string {
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

func DashboardTimeSeriesQuery_TimeSeriesQueryLanguage_ToProto(mapCtx *MapContext, in *string) *pb.TimeSeriesQuery_TimeSeriesQueryLanguage {
	if in == nil {
		return nil
	}
	return &pb.TimeSeriesQuery_TimeSeriesQueryLanguage{
		TimeSeriesQueryLanguage: *in,
	}
}
