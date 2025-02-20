package logging

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
        krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
        metric "google.golang.org/genproto/googleapis/api/metric"
        pb "cloud.google.com/go/logging/apiv2/loggingpb"
)

func MetricDescriptor_FromProto(mapCtx *direct.MapContext, in interface{}) *krm.MetricDescriptor {
	return nil
}

func Distribution_BucketOptions_FromProto(mapCtx *direct.MapContext, in interface{}) *krm.Distribution_BucketOptions {
	return nil
}

func MetricDescriptor_ToProto(mapCtx *direct.MapContext, in interface{}) *metric.MetricDescriptor {
	return nil
}

func Distribution_BucketOptions_ToProto(mapCtx *direct.MapContext, in interface{}) *pb.Distribution_BucketOptions {
	return nil
}
