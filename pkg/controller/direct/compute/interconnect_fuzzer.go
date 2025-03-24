import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
)

func ComputeInterconnectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Interconnect) *krm.ComputeInterconnectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeInterconnectObservedState{}
	out.AvailableFeatures = in.GetAvailableFeatures()
	out.CircuitInfos = direct.Slice_FromProto(mapCtx, in.CircuitInfos, InterconnectCircuitInfo_FromProto)
	out.CreationTimestamp = direct.LazyPtr(in.GetCreationTimestamp())
	out.ExpectedOutages = direct.Slice_FromProto(mapCtx, in.ExpectedOutages, InterconnectOutageNotification_FromProto)
	out.GoogleIPAddress = direct.LazyPtr(in.GetGoogleIpAddress())
	out.GoogleReferenceID = direct.LazyPtr(in.GetGoogleReferenceId())
	out.ID = direct.Uint64Value_FromProto(mapCtx, in.GetId())
	out.InterconnectAttachments = in.GetInterconnectAttachments()
	out.Kind = direct.LazyPtr(in.GetKind())
	out.OperationalStatus = direct.Enum_FromProto(mapCtx, in.GetOperationalStatus())
	out.PeerIPAddress = direct.LazyPtr(in.GetPeerIpAddress())
	out.ProvisionedLinkCount = direct.LazyPtr(in.GetProvisionedLinkCount())
	out.SatisfiesPzs = in.GetSatisfiesPzs()
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}

