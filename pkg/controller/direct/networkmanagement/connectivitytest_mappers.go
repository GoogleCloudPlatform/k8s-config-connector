package networkmanagement

import (
	pb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkmanagement/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

func EndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.EndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EndpointObservedState{}
	// MISSING: IPAddress
	// MISSING: Port
	// MISSING: Instance
	// MISSING: ForwardingRule
	out.ForwardingRuleTarget = direct.Enum_FromProto(mapCtx, in.GetForwardingRuleTarget())
	out.LoadBalancerID = in.LoadBalancerId
	out.LoadBalancerType = direct.Enum_FromProto(mapCtx, in.GetLoadBalancerType())
	// MISSING: GKEMasterCluster
	// MISSING: Fqdn
	// MISSING: CloudSQLInstance
	// MISSING: RedisInstance
	// MISSING: RedisCluster
	// MISSING: CloudFunction
	// MISSING: AppEngineVersion
	// MISSING: CloudRunRevision
	// MISSING: Network
	// MISSING: NetworkType
	// MISSING: ProjectID
	return out
}
func EndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EndpointObservedState) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	// MISSING: IPAddress
	// MISSING: Port
	// MISSING: Instance
	// MISSING: ForwardingRule
	out.ForwardingRuleTarget = direct.PtrTo(direct.Enum_ToProto[pb.Endpoint_ForwardingRuleTarget](mapCtx, in.ForwardingRuleTarget))
	out.LoadBalancerId = in.LoadBalancerID
	out.LoadBalancerType = direct.PtrTo(direct.Enum_ToProto[pb.LoadBalancerType](mapCtx, in.LoadBalancerType))
	// MISSING: GKEMasterCluster
	// MISSING: Fqdn
	// MISSING: CloudSQLInstance
	// MISSING: RedisInstance
	// MISSING: RedisCluster
	// MISSING: CloudFunction
	// MISSING: AppEngineVersion
	// MISSING: CloudRunRevision
	// MISSING: Network
	// MISSING: NetworkType
	// MISSING: ProjectID
	return out
}
func StatusObservedState_FromProto(mapCtx *direct.MapContext, in *status.Status) *krm.StatusObservedState {
	if in == nil {
		return nil
	}

	out := &krm.StatusObservedState{
		Code:    direct.LazyPtr(in.Code),
		Message: direct.LazyPtr(in.Message),
	}
	if len(in.Details) == 0 {
		return out
	}
	detailsOut := make([]krm.Any, 0)
	for _, d := range in.Details {
		if d == nil {
			continue
		}
		dOut := krm.Any{
			TypeURL: direct.LazyPtr(d.TypeUrl),
			Value:   direct.ByteSliceToStringPtr(mapCtx, d.Value),
		}
		detailsOut = append(detailsOut, dOut)
	}
	if len(detailsOut) > 0 {
		out.Details = detailsOut
	}
	return out
}
func StatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StatusObservedState) *status.Status {
	if in == nil {
		return nil
	}

	out := &status.Status{
		Code:    direct.ValueOf(in.Code),
		Message: direct.ValueOf(in.Message),
	}
	if len(in.Details) == 0 {
		return out
	}
	detailsOut := make([]*anypb.Any, 0)
	for _, d := range in.Details {
		dOut := &anypb.Any{
			TypeUrl: direct.ValueOf(d.TypeURL),
			Value:   direct.StringPtrToByteSlice(mapCtx, d.Value),
		}
		detailsOut = append(detailsOut, dOut)
	}
	if len(detailsOut) > 0 {
		out.Details = detailsOut
	}
	return out
}
func NetworkManagementConnectivityTestSpec_RelatedProjects_FromProto(mapCtx *direct.MapContext, in []string) []refs.ProjectRef {
	if len(in) == 0 {
		return nil
	}
	out := make([]refs.ProjectRef, 0)
	for _, p := range in {
		if p == "" {
			continue
		}
		projectRef := refs.ProjectRef{External: p}
		out = append(out, projectRef)
	}
	return out
}
func NetworkManagementConnectivityTestSpec_RelatedProjects_ToProto(mapCtx *direct.MapContext, in []refs.ProjectRef) []string {
	if len(in) == 0 {
		return nil
	}
	out := make([]string, 0)
	for _, p := range in {
		if p.External == "" {
			continue
		}
		out = append(out, p.External)
	}
	return out
}
