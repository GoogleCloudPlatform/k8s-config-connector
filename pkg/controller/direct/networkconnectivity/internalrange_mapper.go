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

package networkconnectivity

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Migration_FromProto(mapCtx *direct.MapContext, in *pb.Migration) *krm.Migration {
	if in == nil {
		return nil
	}
	out := &krm.Migration{}
	out.Source = direct.LazyPtr(in.GetSource())
	out.Target = direct.LazyPtr(in.GetTarget())
	return out
}
func Migration_ToProto(mapCtx *direct.MapContext, in *krm.Migration) *pb.Migration {
	if in == nil {
		return nil
	}
	out := &pb.Migration{}
	out.Source = direct.ValueOf(in.Source)
	out.Target = direct.ValueOf(in.Target)
	return out
}
func NetworkConnectivityInternalRangeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InternalRange) *krm.NetworkConnectivityInternalRangeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityInternalRangeObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Users = in.Users
	return out
}
func NetworkConnectivityInternalRangeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityInternalRangeObservedState) *pb.InternalRange {
	if in == nil {
		return nil
	}
	out := &pb.InternalRange{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Users = in.Users
	return out
}
func NetworkConnectivityInternalRangeSpec_FromProto(mapCtx *direct.MapContext, in *pb.InternalRange) *krm.NetworkConnectivityInternalRangeSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityInternalRangeSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.IPCIDRRange = direct.LazyPtr(in.GetIpCidrRange())
	out.Labels = in.Labels
	out.Migration = Migration_FromProto(mapCtx, in.GetMigration())
	out.Name = direct.LazyPtr(in.GetName())
	if in.GetNetwork() != "" {
		out.NetworkRef = &refsv1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.Overlaps = in.Overlaps
	out.Peering = direct.LazyPtr(in.GetPeering())
	out.PrefixLength = direct.LazyPtr(in.GetPrefixLength())
	out.TargetCIDRRange = in.TargetCidrRange
	out.Usage = direct.LazyPtr(in.GetUsage())
	return out
}
func NetworkConnectivityInternalRangeSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityInternalRangeSpec) *pb.InternalRange {
	if in == nil {
		return nil
	}
	out := &pb.InternalRange{}
	out.Description = direct.ValueOf(in.Description)
	out.IpCidrRange = direct.ValueOf(in.IPCIDRRange)
	out.Labels = in.Labels
	out.Migration = Migration_ToProto(mapCtx, in.Migration)
	out.Name = direct.ValueOf(in.Name)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.Overlaps = in.Overlaps
	out.Peering = direct.ValueOf(in.Peering)
	out.PrefixLength = direct.ValueOf(in.PrefixLength)
	out.TargetCidrRange = in.TargetCIDRRange
	out.Usage = direct.ValueOf(in.Usage)
	return out
}
