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
func EndpointMatcher_FromProto(mapCtx *direct.MapContext, in *pb.EndpointMatcher) *krm.EndpointMatcher {
	if in == nil {
		return nil
	}
	out := &krm.EndpointMatcher{}
	out.MetadataLabelMatcher = EndpointMatcher_MetadataLabelMatcher_FromProto(mapCtx, in.GetMetadataLabelMatcher())
	return out
}
func EndpointMatcher_ToProto(mapCtx *direct.MapContext, in *krm.EndpointMatcher) *pb.EndpointMatcher {
	if in == nil {
		return nil
	}
	out := &pb.EndpointMatcher{}
	if oneof := EndpointMatcher_MetadataLabelMatcher_ToProto(mapCtx, in.MetadataLabelMatcher); oneof != nil {
		out.MatcherType = &pb.EndpointMatcher_MetadataLabelMatcher_{MetadataLabelMatcher: oneof}
	}
	return out
}
func EndpointMatcher_MetadataLabelMatcher_FromProto(mapCtx *direct.MapContext, in *pb.EndpointMatcher_MetadataLabelMatcher) *krm.EndpointMatcher_MetadataLabelMatcher {
	if in == nil {
		return nil
	}
	out := &krm.EndpointMatcher_MetadataLabelMatcher{}
	out.MetadataLabelMatchCriteria = direct.Enum_FromProto(mapCtx, in.GetMetadataLabelMatchCriteria())
	out.MetadataLabels = direct.Slice_FromProto(mapCtx, in.MetadataLabels, EndpointMatcher_MetadataLabelMatcher_MetadataLabels_FromProto)
	return out
}
func EndpointMatcher_MetadataLabelMatcher_ToProto(mapCtx *direct.MapContext, in *krm.EndpointMatcher_MetadataLabelMatcher) *pb.EndpointMatcher_MetadataLabelMatcher {
	if in == nil {
		return nil
	}
	out := &pb.EndpointMatcher_MetadataLabelMatcher{}
	out.MetadataLabelMatchCriteria = direct.Enum_ToProto[pb.EndpointMatcher_MetadataLabelMatcher_MetadataLabelMatchCriteria](mapCtx, in.MetadataLabelMatchCriteria)
	out.MetadataLabels = direct.Slice_ToProto(mapCtx, in.MetadataLabels, EndpointMatcher_MetadataLabelMatcher_MetadataLabels_ToProto)
	return out
}
func EndpointMatcher_MetadataLabelMatcher_MetadataLabels_FromProto(mapCtx *direct.MapContext, in *pb.EndpointMatcher_MetadataLabelMatcher_MetadataLabels) *krm.EndpointMatcher_MetadataLabelMatcher_MetadataLabels {
	if in == nil {
		return nil
	}
	out := &krm.EndpointMatcher_MetadataLabelMatcher_MetadataLabels{}
	out.LabelName = direct.LazyPtr(in.GetLabelName())
	out.LabelValue = direct.LazyPtr(in.GetLabelValue())
	return out
}
func EndpointMatcher_MetadataLabelMatcher_MetadataLabels_ToProto(mapCtx *direct.MapContext, in *krm.EndpointMatcher_MetadataLabelMatcher_MetadataLabels) *pb.EndpointMatcher_MetadataLabelMatcher_MetadataLabels {
	if in == nil {
		return nil
	}
	out := &pb.EndpointMatcher_MetadataLabelMatcher_MetadataLabels{}
	out.LabelName = direct.ValueOf(in.LabelName)
	out.LabelValue = direct.ValueOf(in.LabelValue)
	return out
}
func EndpointPolicy_FromProto(mapCtx *direct.MapContext, in *pb.EndpointPolicy) *krm.EndpointPolicy {
	if in == nil {
		return nil
	}
	out := &krm.EndpointPolicy{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.AuthorizationPolicy = direct.LazyPtr(in.GetAuthorizationPolicy())
	out.EndpointMatcher = EndpointMatcher_FromProto(mapCtx, in.GetEndpointMatcher())
	out.TrafficPortSelector = TrafficPortSelector_FromProto(mapCtx, in.GetTrafficPortSelector())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ServerTlsPolicy = direct.LazyPtr(in.GetServerTlsPolicy())
	out.ClientTlsPolicy = direct.LazyPtr(in.GetClientTlsPolicy())
	return out
}
func EndpointPolicy_ToProto(mapCtx *direct.MapContext, in *krm.EndpointPolicy) *pb.EndpointPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EndpointPolicy{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Type = direct.Enum_ToProto[pb.EndpointPolicy_EndpointPolicyType](mapCtx, in.Type)
	out.AuthorizationPolicy = direct.ValueOf(in.AuthorizationPolicy)
	out.EndpointMatcher = EndpointMatcher_ToProto(mapCtx, in.EndpointMatcher)
	out.TrafficPortSelector = TrafficPortSelector_ToProto(mapCtx, in.TrafficPortSelector)
	out.Description = direct.ValueOf(in.Description)
	out.ServerTlsPolicy = direct.ValueOf(in.ServerTlsPolicy)
	out.ClientTlsPolicy = direct.ValueOf(in.ClientTlsPolicy)
	return out
}
func EndpointPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EndpointPolicy) *krm.EndpointPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EndpointPolicyObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Type
	// MISSING: AuthorizationPolicy
	// MISSING: EndpointMatcher
	// MISSING: TrafficPortSelector
	// MISSING: Description
	// MISSING: ServerTlsPolicy
	// MISSING: ClientTlsPolicy
	return out
}
func EndpointPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EndpointPolicyObservedState) *pb.EndpointPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EndpointPolicy{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Type
	// MISSING: AuthorizationPolicy
	// MISSING: EndpointMatcher
	// MISSING: TrafficPortSelector
	// MISSING: Description
	// MISSING: ServerTlsPolicy
	// MISSING: ClientTlsPolicy
	return out
}
func NetworkservicesEndpointPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EndpointPolicy) *krm.NetworkservicesEndpointPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesEndpointPolicyObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	// MISSING: AuthorizationPolicy
	// MISSING: EndpointMatcher
	// MISSING: TrafficPortSelector
	// MISSING: Description
	// MISSING: ServerTlsPolicy
	// MISSING: ClientTlsPolicy
	return out
}
func NetworkservicesEndpointPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesEndpointPolicyObservedState) *pb.EndpointPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EndpointPolicy{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	// MISSING: AuthorizationPolicy
	// MISSING: EndpointMatcher
	// MISSING: TrafficPortSelector
	// MISSING: Description
	// MISSING: ServerTlsPolicy
	// MISSING: ClientTlsPolicy
	return out
}
func NetworkservicesEndpointPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.EndpointPolicy) *krm.NetworkservicesEndpointPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesEndpointPolicySpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	// MISSING: AuthorizationPolicy
	// MISSING: EndpointMatcher
	// MISSING: TrafficPortSelector
	// MISSING: Description
	// MISSING: ServerTlsPolicy
	// MISSING: ClientTlsPolicy
	return out
}
func NetworkservicesEndpointPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesEndpointPolicySpec) *pb.EndpointPolicy {
	if in == nil {
		return nil
	}
	out := &pb.EndpointPolicy{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Type
	// MISSING: AuthorizationPolicy
	// MISSING: EndpointMatcher
	// MISSING: TrafficPortSelector
	// MISSING: Description
	// MISSING: ServerTlsPolicy
	// MISSING: ClientTlsPolicy
	return out
}
func TrafficPortSelector_FromProto(mapCtx *direct.MapContext, in *pb.TrafficPortSelector) *krm.TrafficPortSelector {
	if in == nil {
		return nil
	}
	out := &krm.TrafficPortSelector{}
	out.Ports = in.Ports
	return out
}
func TrafficPortSelector_ToProto(mapCtx *direct.MapContext, in *krm.TrafficPortSelector) *pb.TrafficPortSelector {
	if in == nil {
		return nil
	}
	out := &pb.TrafficPortSelector{}
	out.Ports = in.Ports
	return out
}
