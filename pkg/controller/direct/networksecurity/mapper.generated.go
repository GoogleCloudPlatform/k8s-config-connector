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

package networksecurity

import (
	pb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AuthorizationPolicy_Rule_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizationPolicy_Rule) *krm.AuthorizationPolicy_Rule {
	if in == nil {
		return nil
	}
	out := &krm.AuthorizationPolicy_Rule{}
	out.Sources = direct.Slice_FromProto(mapCtx, in.Sources, AuthorizationPolicy_Rule_Source_FromProto)
	out.Destinations = direct.Slice_FromProto(mapCtx, in.Destinations, AuthorizationPolicy_Rule_Destination_FromProto)
	return out
}
func AuthorizationPolicy_Rule_ToProto(mapCtx *direct.MapContext, in *krm.AuthorizationPolicy_Rule) *pb.AuthorizationPolicy_Rule {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizationPolicy_Rule{}
	out.Sources = direct.Slice_ToProto(mapCtx, in.Sources, AuthorizationPolicy_Rule_Source_ToProto)
	out.Destinations = direct.Slice_ToProto(mapCtx, in.Destinations, AuthorizationPolicy_Rule_Destination_ToProto)
	return out
}
func AuthorizationPolicy_Rule_Destination_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizationPolicy_Rule_Destination) *krm.AuthorizationPolicy_Rule_Destination {
	if in == nil {
		return nil
	}
	out := &krm.AuthorizationPolicy_Rule_Destination{}
	out.Hosts = in.Hosts
	out.Ports = in.Ports
	out.Methods = in.Methods
	out.HTTPHeaderMatch = AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch_FromProto(mapCtx, in.GetHttpHeaderMatch())
	return out
}
func AuthorizationPolicy_Rule_Destination_ToProto(mapCtx *direct.MapContext, in *krm.AuthorizationPolicy_Rule_Destination) *pb.AuthorizationPolicy_Rule_Destination {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizationPolicy_Rule_Destination{}
	out.Hosts = in.Hosts
	out.Ports = in.Ports
	out.Methods = in.Methods
	out.HttpHeaderMatch = AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch_ToProto(mapCtx, in.HTTPHeaderMatch)
	return out
}
func AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizationPolicy_Rule_Destination_HttpHeaderMatch) *krm.AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch {
	if in == nil {
		return nil
	}
	out := &krm.AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch{}
	out.RegexMatch = direct.LazyPtr(in.GetRegexMatch())
	out.HeaderName = direct.LazyPtr(in.GetHeaderName())
	return out
}
func AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch_ToProto(mapCtx *direct.MapContext, in *krm.AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch) *pb.AuthorizationPolicy_Rule_Destination_HttpHeaderMatch {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizationPolicy_Rule_Destination_HttpHeaderMatch{}
	if oneof := AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch_RegexMatch_ToProto(mapCtx, in.RegexMatch); oneof != nil {
		out.Type = oneof
	}
	out.HeaderName = direct.ValueOf(in.HeaderName)
	return out
}
func AuthorizationPolicy_Rule_Source_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizationPolicy_Rule_Source) *krm.AuthorizationPolicy_Rule_Source {
	if in == nil {
		return nil
	}
	out := &krm.AuthorizationPolicy_Rule_Source{}
	out.Principals = in.Principals
	out.IPBlocks = in.IpBlocks
	return out
}
func AuthorizationPolicy_Rule_Source_ToProto(mapCtx *direct.MapContext, in *krm.AuthorizationPolicy_Rule_Source) *pb.AuthorizationPolicy_Rule_Source {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizationPolicy_Rule_Source{}
	out.Principals = in.Principals
	out.IpBlocks = in.IPBlocks
	return out
}
func NetworkSecurityAuthorizationPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizationPolicy) *krm.NetworkSecurityAuthorizationPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityAuthorizationPolicySpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.Action = direct.Enum_FromProto(mapCtx, in.GetAction())
	out.Rules = direct.Slice_FromProto(mapCtx, in.Rules, AuthorizationPolicy_Rule_FromProto)
	return out
}
func NetworkSecurityAuthorizationPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityAuthorizationPolicySpec) *pb.AuthorizationPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizationPolicy{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.Action = direct.Enum_ToProto[pb.AuthorizationPolicy_Action](mapCtx, in.Action)
	out.Rules = direct.Slice_ToProto(mapCtx, in.Rules, AuthorizationPolicy_Rule_ToProto)
	return out
}
func NetworkSecurityAuthorizationPolicyStatus_FromProto(mapCtx *direct.MapContext, in *pb.AuthorizationPolicy) *krm.NetworkSecurityAuthorizationPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityAuthorizationPolicyStatus{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Action
	// MISSING: Rules
	return out
}
func NetworkSecurityAuthorizationPolicyStatus_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityAuthorizationPolicyStatus) *pb.AuthorizationPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AuthorizationPolicy{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Action
	// MISSING: Rules
	return out
}
