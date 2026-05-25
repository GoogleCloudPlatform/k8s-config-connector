package networkmanagement

import (
        pb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
        krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkmanagement/v1alpha1"
        "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ProbingDetails_SingleEdgeResponseObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProbingDetails_SingleEdgeResponse) *krm.ProbingDetails_SingleEdgeResponseObservedState {
        if in == nil {
                return nil
        }
        mapCtx.NotImplemented()
        return nil
}

func ProbingDetails_SingleEdgeResponseObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProbingDetails_SingleEdgeResponseObservedState) *pb.ProbingDetails_SingleEdgeResponse {
        if in == nil {
                return nil
        }
        mapCtx.NotImplemented()
        return nil
}
