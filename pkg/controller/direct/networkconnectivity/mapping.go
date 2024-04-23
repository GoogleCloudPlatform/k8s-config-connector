package networkconnectivity

import (
	api "google.golang.org/api/networkconnectivity/v1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
)

func ServiceConnectionPolicySpec_FromProto(ctx *MapContext, in *api.ServiceConnectionPolicy) *krm.ServiceConnectionPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.ServiceConnectionPolicySpec{}

	out.Description = in.Description

	// // Labels: User-defined labels.
	// Labels map[string]string `json:"labels,omitempty"`

	out.Network = in.Network

	out.PscConfig = PscConfig_FromProto(ctx, in.PscConfig)

	out.ServiceClass = in.ServiceClass

	return out
}
func ServiceConnectionPolicySpec_ToProto(ctx *MapContext, in *krm.ServiceConnectionPolicySpec) *api.ServiceConnectionPolicy {
	if in == nil {
		return nil
	}
	out := &api.ServiceConnectionPolicy{}

	out.Description = in.Description

	// // Labels: User-defined labels.
	// Labels map[string]string `json:"labels,omitempty"`

	out.Network = in.Network

	out.PscConfig = PscConfig_ToProto(ctx, in.PscConfig)

	out.ServiceClass = in.ServiceClass
	return out
}

func ServiceConnectionPolicyState_FromProto(ctx *MapContext, in *api.ServiceConnectionPolicy) *krm.ServiceConnectionPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServiceConnectionPolicyObservedState{}

	out.Infrastructure = in.Infrastructure

	out.PscConnections = Slice_FromProto(ctx, in.PscConnections, PscConnection_FromProto)

	return out
}
func ServiceConnectionPolicySpecState_ToProto(ctx *MapContext, in *krm.ServiceConnectionPolicyObservedState) *api.ServiceConnectionPolicy {
	if in == nil {
		return nil
	}
	out := &api.ServiceConnectionPolicy{}

	out.Infrastructure = in.Infrastructure

	out.PscConnections = Slice_ToProto(ctx, in.PscConnections, PscConnection_ToProto)

	return out
}

func PscConfig_FromProto(ctx *MapContext, in *api.PscConfig) *krm.PscConfig {
	if in == nil {
		return nil
	}
	out := &krm.PscConfig{}
	out.Limit = LazyPtr(in.Limit)
	out.Subnetworks = in.Subnetworks
	return out
}
func PscConfig_ToProto(ctx *MapContext, in *krm.PscConfig) *api.PscConfig {
	if in == nil {
		return nil
	}
	out := &api.PscConfig{}
	out.Limit = ValueOf(in.Limit)
	out.Subnetworks = in.Subnetworks
	return out
}
func PscConnection_FromProto(ctx *MapContext, in *api.PscConnection) *krm.PscConnection {
	if in == nil {
		return nil
	}
	out := &krm.PscConnection{}

	out.ConsumerAddress = LazyPtr(in.ConsumerAddress)
	out.ConsumerForwardingRule = LazyPtr(in.ConsumerForwardingRule)
	out.ConsumerTargetProject = LazyPtr(in.ConsumerTargetProject)
	out.PscConnectionId = LazyPtr(in.PscConnectionId)
	out.State = LazyPtr(in.State)
	return out
}

func PscConnection_ToProto(ctx *MapContext, in *krm.PscConnection) *api.PscConnection {
	if in == nil {
		return nil
	}
	out := &api.PscConnection{}
	out.ConsumerAddress = ValueOf(in.ConsumerAddress)
	out.ConsumerForwardingRule = ValueOf(in.ConsumerForwardingRule)
	out.ConsumerTargetProject = ValueOf(in.ConsumerTargetProject)
	out.PscConnectionId = ValueOf(in.PscConnectionId)
	out.State = ValueOf(in.State)
	return out
}
