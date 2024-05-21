package monitoring

import (
	"strings"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
)

func DashboardBlank_FromProto(ctx *MapContext, in *emptypb.Empty) *krm.DashboardBlank {
	if in == nil {
		return nil
	}
	out := &krm.DashboardBlank{}
	return out
}
func DashboardBlank_ToProto(ctx *MapContext, in *krm.DashboardBlank) *emptypb.Empty {
	if in == nil {
		return nil
	}
	out := &emptypb.Empty{}
	return out
}

func Aggregation_AlignmentPeriod_FromProto(ctx *MapContext, in *durationpb.Duration) *string {
	return SecondsString_FromProto(ctx, in)
}

func Aggregation_AlignmentPeriod_ToProto(ctx *MapContext, in *string) *durationpb.Duration {
	return SecondsString_ToProto(ctx, in, "alignmentPeriod")
}

// TODO: The format is not documented, we need to validate
func DataSet_MinAlignmentPeriod_FromProto(ctx *MapContext, in *durationpb.Duration) *string {
	return SecondsString_FromProto(ctx, in)
}

// TODO: The format is not documented, we need to validate
func DataSet_MinAlignmentPeriod_ToProto(ctx *MapContext, in *string) *durationpb.Duration {
	return SecondsString_ToProto(ctx, in, "minAlignmentPeriod")
}

// TODO: The format is not documented, we need to validate
func SparkChartView_MinAlignmentPeriod_FromProto(ctx *MapContext, in *durationpb.Duration) *string {
	return SecondsString_FromProto(ctx, in)
}

// TODO: The format is not documented, we need to validate
func SparkChartView_MinAlignmentPeriod_ToProto(ctx *MapContext, in *string) *durationpb.Duration {
	return SecondsString_ToProto(ctx, in, "minAlignmentPeriod")
}

// TODO: The format is not documented, we need to validate
func XyChart_TimeshiftDuration_FromProto(ctx *MapContext, in *durationpb.Duration) *string {
	return Duration_FromProto(ctx, in)
}

// TODO: The format is not documented, we need to validate
func XyChart_TimeshiftDuration_ToProto(ctx *MapContext, in *string) *durationpb.Duration {
	return Duration_ToProto(ctx, in)
}

func TableDataSet_MinAlignmentPeriod_FromProto(ctx *MapContext, in *durationpb.Duration) *string {
	return Duration_FromProto(ctx, in)
}

func TableDataSet_MinAlignmentPeriod_ToProto(ctx *MapContext, in *string) *durationpb.Duration {
	return Duration_ToProto(ctx, in)
}

func TimeSeriesQuery_TimeSeriesQueryLanguage_ToProto(ctx *MapContext, in *string) *pb.TimeSeriesQuery_TimeSeriesQueryLanguage {
	if in == nil {
		return nil
	}

	return &pb.TimeSeriesQuery_TimeSeriesQueryLanguage{
		TimeSeriesQueryLanguage: ValueOf(in),
	}
}

func DashboardLogsPanel_ResourceNames_FromProto(ctx *MapContext, in []string) []v1alpha1.ResourceRef {
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
			ctx.Errorf("resourceName %q was not recognized", v)
		}
	}
	return out
}

func DashboardLogsPanel_ResourceNames_ToProto(ctx *MapContext, in []v1alpha1.ResourceRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, ref := range in {
		if ref.External == "" {
			ctx.Errorf("reference was not pre-resolved")
		}
		out = append(out, ref.External)
	}
	return out
}

func DashboardTimeSeriesQuery_TimeSeriesQueryLanguage_ToProto(ctx *MapContext, in *string) *pb.TimeSeriesQuery_TimeSeriesQueryLanguage {
	if in == nil {
		return nil
	}
	return &pb.TimeSeriesQuery_TimeSeriesQueryLanguage{
		TimeSeriesQueryLanguage: *in,
	}
}
