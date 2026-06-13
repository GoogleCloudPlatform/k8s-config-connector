// Copyright 2026 Google LLC
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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeFirewallSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Firewall) *krm.ComputeFirewallSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFirewallSpec{}

	out.Allow = FirewallAllowSlice_v1beta1_FromProto(mapCtx, in.Allowed)
	out.Deny = FirewallDenySlice_v1beta1_FromProto(mapCtx, in.Denied)
	out.Description = in.Description
	out.DestinationRanges = in.DestinationRanges
	out.Direction = in.Direction
	out.Disabled = in.Disabled
	out.LogConfig = FirewallLogConfig_v1beta1_FromProto(mapCtx, in.LogConfig)
	if in.LogConfig != nil && in.LogConfig.Enable != nil {
		out.EnableLogging = in.LogConfig.Enable
	}
	if in.GetNetwork() != "" {
		out.NetworkRef = &krm.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.Priority != nil {
		out.Priority = direct.LazyPtr(int64(*in.Priority))
	}
	out.SourceRanges = in.SourceRanges
	out.SourceServiceAccounts = ComputeFirewallSpec_SourceServiceAccounts_FromProto(mapCtx, in.SourceServiceAccounts)
	out.SourceTags = in.SourceTags
	out.TargetServiceAccounts = ComputeFirewallSpec_TargetServiceAccounts_FromProto(mapCtx, in.TargetServiceAccounts)
	out.TargetTags = in.TargetTags

	return out
}

func ComputeFirewallSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFirewallSpec) *pb.Firewall {
	if in == nil {
		return nil
	}
	out := &pb.Firewall{}

	out.Allowed = FirewallAllowSlice_v1beta1_ToProto(mapCtx, in.Allow)
	out.Denied = FirewallDenySlice_v1beta1_ToProto(mapCtx, in.Deny)
	out.Description = in.Description
	out.DestinationRanges = in.DestinationRanges
	out.Direction = in.Direction
	out.Disabled = in.Disabled
	out.LogConfig = FirewallLogConfig_v1beta1_ToProto(mapCtx, in.LogConfig)
	if in.EnableLogging != nil {
		if out.LogConfig == nil {
			out.LogConfig = &pb.FirewallLogConfig{}
		}
		out.LogConfig.Enable = in.EnableLogging
	}
	if in.NetworkRef != nil {
		out.Network = direct.LazyPtr(in.NetworkRef.External)
	}
	if in.Priority != nil {
		out.Priority = direct.LazyPtr(int32(*in.Priority))
	}
	out.SourceRanges = in.SourceRanges
	out.SourceServiceAccounts = ComputeFirewallSpec_SourceServiceAccounts_ToProto(mapCtx, in.SourceServiceAccounts)
	out.SourceTags = in.SourceTags
	out.TargetServiceAccounts = ComputeFirewallSpec_TargetServiceAccounts_ToProto(mapCtx, in.TargetServiceAccounts)
	out.TargetTags = in.TargetTags

	return out
}

func FirewallAllowSlice_v1beta1_FromProto(mapCtx *direct.MapContext, in []*pb.Allowed) []krm.FirewallAllow {
	if in == nil {
		return nil
	}
	var out []krm.FirewallAllow
	for _, item := range in {
		if item == nil {
			continue
		}
		out = append(out, krm.FirewallAllow{
			Ports:    item.Ports,
			Protocol: item.GetIPProtocol(),
		})
	}
	return out
}

func FirewallAllowSlice_v1beta1_ToProto(mapCtx *direct.MapContext, in []krm.FirewallAllow) []*pb.Allowed {
	if in == nil {
		return nil
	}
	var out []*pb.Allowed
	for _, item := range in {
		out = append(out, &pb.Allowed{
			Ports:      item.Ports,
			IPProtocol: direct.LazyPtr(item.Protocol),
		})
	}
	return out
}

func FirewallDenySlice_v1beta1_FromProto(mapCtx *direct.MapContext, in []*pb.Denied) []krm.FirewallDeny {
	if in == nil {
		return nil
	}
	var out []krm.FirewallDeny
	for _, item := range in {
		if item == nil {
			continue
		}
		out = append(out, krm.FirewallDeny{
			Ports:    item.Ports,
			Protocol: item.GetIPProtocol(),
		})
	}
	return out
}

func FirewallDenySlice_v1beta1_ToProto(mapCtx *direct.MapContext, in []krm.FirewallDeny) []*pb.Denied {
	if in == nil {
		return nil
	}
	var out []*pb.Denied
	for _, item := range in {
		out = append(out, &pb.Denied{
			Ports:      item.Ports,
			IPProtocol: direct.LazyPtr(item.Protocol),
		})
	}
	return out
}

func FirewallLogConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallLogConfig) *krm.FirewallLogConfig {
	if in == nil {
		return nil
	}
	out := &krm.FirewallLogConfig{}
	out.Metadata = in.GetMetadata()
	return out
}

func FirewallLogConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.FirewallLogConfig) *pb.FirewallLogConfig {
	if in == nil {
		return nil
	}
	out := &pb.FirewallLogConfig{}
	out.Metadata = direct.LazyPtr(in.Metadata)
	return out
}

func ComputeFirewallSpec_SourceServiceAccounts_ToProto(mapCtx *direct.MapContext, in []*refs.IAMServiceAccountRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, i := range in {
		if i == nil {
			continue
		}
		if i.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", i.Name)
		}
		out = append(out, i.External)
	}
	return out
}

func ComputeFirewallSpec_TargetServiceAccounts_ToProto(mapCtx *direct.MapContext, in []*refs.IAMServiceAccountRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, i := range in {
		if i == nil {
			continue
		}
		if i.External == "" {
			mapCtx.Errorf("reference %s was not pre-resolved", i.Name)
		}
		out = append(out, i.External)
	}
	return out
}

func ComputeFirewallSpec_SourceServiceAccounts_FromProto(mapCtx *direct.MapContext, in []string) []*refs.IAMServiceAccountRef {
	if in == nil {
		return nil
	}
	var out []*refs.IAMServiceAccountRef
	for _, i := range in {
		out = append(out, &refs.IAMServiceAccountRef{
			External: i,
		})
	}
	return out
}

func ComputeFirewallSpec_TargetServiceAccounts_FromProto(mapCtx *direct.MapContext, in []string) []*refs.IAMServiceAccountRef {
	if in == nil {
		return nil
	}
	var out []*refs.IAMServiceAccountRef
	for _, i := range in {
		out = append(out, &refs.IAMServiceAccountRef{
			External: i,
		})
	}
	return out
}
