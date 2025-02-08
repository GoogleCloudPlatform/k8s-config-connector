// Copyright 2025 Google LLC
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

package networkservices

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
)
func ExtensionChain_FromProto(mapCtx *direct.MapContext, in *pb.ExtensionChain) *krm.ExtensionChain {
	if in == nil {
		return nil
	}
	out := &krm.ExtensionChain{}
	out.Name = direct.LazyPtr(in.GetName())
	out.MatchCondition = ExtensionChain_MatchCondition_FromProto(mapCtx, in.GetMatchCondition())
	out.Extensions = direct.Slice_FromProto(mapCtx, in.Extensions, ExtensionChain_Extension_FromProto)
	return out
}
func ExtensionChain_ToProto(mapCtx *direct.MapContext, in *krm.ExtensionChain) *pb.ExtensionChain {
	if in == nil {
		return nil
	}
	out := &pb.ExtensionChain{}
	out.Name = direct.ValueOf(in.Name)
	out.MatchCondition = ExtensionChain_MatchCondition_ToProto(mapCtx, in.MatchCondition)
	out.Extensions = direct.Slice_ToProto(mapCtx, in.Extensions, ExtensionChain_Extension_ToProto)
	return out
}
func ExtensionChain_Extension_FromProto(mapCtx *direct.MapContext, in *pb.ExtensionChain_Extension) *krm.ExtensionChain_Extension {
	if in == nil {
		return nil
	}
	out := &krm.ExtensionChain_Extension{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Authority = direct.LazyPtr(in.GetAuthority())
	out.Service = direct.LazyPtr(in.GetService())
	out.SupportedEvents = direct.EnumSlice_FromProto(mapCtx, in.SupportedEvents)
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.FailOpen = direct.LazyPtr(in.GetFailOpen())
	out.ForwardHeaders = in.ForwardHeaders
	return out
}
func ExtensionChain_Extension_ToProto(mapCtx *direct.MapContext, in *krm.ExtensionChain_Extension) *pb.ExtensionChain_Extension {
	if in == nil {
		return nil
	}
	out := &pb.ExtensionChain_Extension{}
	out.Name = direct.ValueOf(in.Name)
	out.Authority = direct.ValueOf(in.Authority)
	out.Service = direct.ValueOf(in.Service)
	out.SupportedEvents = direct.EnumSlice_ToProto[pb.EventType](mapCtx, in.SupportedEvents)
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.FailOpen = direct.ValueOf(in.FailOpen)
	out.ForwardHeaders = in.ForwardHeaders
	return out
}
func ExtensionChain_MatchCondition_FromProto(mapCtx *direct.MapContext, in *pb.ExtensionChain_MatchCondition) *krm.ExtensionChain_MatchCondition {
	if in == nil {
		return nil
	}
	out := &krm.ExtensionChain_MatchCondition{}
	out.CelExpression = direct.LazyPtr(in.GetCelExpression())
	return out
}
func ExtensionChain_MatchCondition_ToProto(mapCtx *direct.MapContext, in *krm.ExtensionChain_MatchCondition) *pb.ExtensionChain_MatchCondition {
	if in == nil {
		return nil
	}
	out := &pb.ExtensionChain_MatchCondition{}
	out.CelExpression = direct.ValueOf(in.CelExpression)
	return out
}
func LbTrafficExtension_FromProto(mapCtx *direct.MapContext, in *pb.LbTrafficExtension) *krm.LbTrafficExtension {
	if in == nil {
		return nil
	}
	out := &krm.LbTrafficExtension{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.ForwardingRules = in.ForwardingRules
	out.ExtensionChains = direct.Slice_FromProto(mapCtx, in.ExtensionChains, ExtensionChain_FromProto)
	out.LoadBalancingScheme = direct.Enum_FromProto(mapCtx, in.GetLoadBalancingScheme())
	out.Metadata = Metadata_FromProto(mapCtx, in.GetMetadata())
	return out
}
func LbTrafficExtension_ToProto(mapCtx *direct.MapContext, in *krm.LbTrafficExtension) *pb.LbTrafficExtension {
	if in == nil {
		return nil
	}
	out := &pb.LbTrafficExtension{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.ForwardingRules = in.ForwardingRules
	out.ExtensionChains = direct.Slice_ToProto(mapCtx, in.ExtensionChains, ExtensionChain_ToProto)
	out.LoadBalancingScheme = direct.Enum_ToProto[pb.LoadBalancingScheme](mapCtx, in.LoadBalancingScheme)
	out.Metadata = Metadata_ToProto(mapCtx, in.Metadata)
	return out
}
func LbTrafficExtensionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LbTrafficExtension) *krm.LbTrafficExtensionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LbTrafficExtensionObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Description
	// MISSING: Labels
	// MISSING: ForwardingRules
	// MISSING: ExtensionChains
	// MISSING: LoadBalancingScheme
	// MISSING: Metadata
	return out
}
func LbTrafficExtensionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LbTrafficExtensionObservedState) *pb.LbTrafficExtension {
	if in == nil {
		return nil
	}
	out := &pb.LbTrafficExtension{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Description
	// MISSING: Labels
	// MISSING: ForwardingRules
	// MISSING: ExtensionChains
	// MISSING: LoadBalancingScheme
	// MISSING: Metadata
	return out
}
func NetworkservicesLbTrafficExtensionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LbTrafficExtension) *krm.NetworkservicesLbTrafficExtensionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesLbTrafficExtensionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Labels
	// MISSING: ForwardingRules
	// MISSING: ExtensionChains
	// MISSING: LoadBalancingScheme
	// MISSING: Metadata
	return out
}
func NetworkservicesLbTrafficExtensionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesLbTrafficExtensionObservedState) *pb.LbTrafficExtension {
	if in == nil {
		return nil
	}
	out := &pb.LbTrafficExtension{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Labels
	// MISSING: ForwardingRules
	// MISSING: ExtensionChains
	// MISSING: LoadBalancingScheme
	// MISSING: Metadata
	return out
}
func NetworkservicesLbTrafficExtensionSpec_FromProto(mapCtx *direct.MapContext, in *pb.LbTrafficExtension) *krm.NetworkservicesLbTrafficExtensionSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesLbTrafficExtensionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Labels
	// MISSING: ForwardingRules
	// MISSING: ExtensionChains
	// MISSING: LoadBalancingScheme
	// MISSING: Metadata
	return out
}
func NetworkservicesLbTrafficExtensionSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesLbTrafficExtensionSpec) *pb.LbTrafficExtension {
	if in == nil {
		return nil
	}
	out := &pb.LbTrafficExtension{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Description
	// MISSING: Labels
	// MISSING: ForwardingRules
	// MISSING: ExtensionChains
	// MISSING: LoadBalancingScheme
	// MISSING: Metadata
	return out
}
