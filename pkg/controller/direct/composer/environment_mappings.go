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

package composer

import (
	pb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/composer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func MasterAuthorizedNetworksConfig_FromProto(mapCtx *direct.MapContext, in *pb.MasterAuthorizedNetworksConfig) *krm.MasterAuthorizedNetworksConfig {
	if in == nil {
		return nil
	}
	out := &krm.MasterAuthorizedNetworksConfig{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.CIDRBlocks = direct.Slice_FromProto(mapCtx, in.CidrBlocks, MasterAuthorizedNetworksConfig_CIDRBlock_FromProto)
	return out
}
func MasterAuthorizedNetworksConfig_ToProto(mapCtx *direct.MapContext, in *krm.MasterAuthorizedNetworksConfig) *pb.MasterAuthorizedNetworksConfig {
	if in == nil {
		return nil
	}
	out := &pb.MasterAuthorizedNetworksConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.CidrBlocks = direct.Slice_ToProto(mapCtx, in.CIDRBlocks, MasterAuthorizedNetworksConfig_CIDRBlock_ToProto)
	return out
}

func WebServerNetworkAccessControl_FromProto(mapCtx *direct.MapContext, in *pb.WebServerNetworkAccessControl) *krm.WebServerNetworkAccessControl {
	if in == nil {
		return nil
	}
	out := &krm.WebServerNetworkAccessControl{}
	out.AllowedIPRanges = direct.Slice_FromProto(mapCtx, in.AllowedIpRanges, WebServerNetworkAccessControl_AllowedIPRange_FromProto)
	return out
}
func WebServerNetworkAccessControl_ToProto(mapCtx *direct.MapContext, in *krm.WebServerNetworkAccessControl) *pb.WebServerNetworkAccessControl {
	if in == nil {
		return nil
	}
	out := &pb.WebServerNetworkAccessControl{}
	out.AllowedIpRanges = direct.Slice_ToProto(mapCtx, in.AllowedIPRanges, WebServerNetworkAccessControl_AllowedIPRange_ToProto)
	return out
}

func IPAllocationPolicy_FromProto(mapCtx *direct.MapContext, in *pb.IPAllocationPolicy) *krm.IPAllocationPolicy {
	if in == nil {
		return nil
	}
	out := &krm.IPAllocationPolicy{}
	out.UseIPAliases = direct.LazyPtr(in.GetUseIpAliases())
	out.ClusterSecondaryRangeName = direct.LazyPtr(in.GetClusterSecondaryRangeName())
	out.ClusterIPV4CIDRBlock = direct.LazyPtr(in.GetClusterIpv4CidrBlock())
	out.ServicesSecondaryRangeName = direct.LazyPtr(in.GetServicesSecondaryRangeName())
	out.ServicesIPV4CIDRBlock = direct.LazyPtr(in.GetServicesIpv4CidrBlock())
	return out
}
func IPAllocationPolicy_ToProto(mapCtx *direct.MapContext, in *krm.IPAllocationPolicy) *pb.IPAllocationPolicy {
	if in == nil {
		return nil
	}
	out := &pb.IPAllocationPolicy{}
	out.UseIpAliases = direct.ValueOf(in.UseIPAliases)
	if in.ClusterSecondaryRangeName != nil {
		out.ClusterIpAllocation = &pb.IPAllocationPolicy_ClusterSecondaryRangeName{
			ClusterSecondaryRangeName: direct.ValueOf(in.ClusterSecondaryRangeName),
		}
	}
	if in.ClusterIPV4CIDRBlock != nil {
		out.ClusterIpAllocation = &pb.IPAllocationPolicy_ClusterIpv4CidrBlock{
			ClusterIpv4CidrBlock: direct.ValueOf(in.ClusterIPV4CIDRBlock),
		}
	}
	if in.ServicesSecondaryRangeName != nil {
		out.ServicesIpAllocation = &pb.IPAllocationPolicy_ServicesSecondaryRangeName{
			ServicesSecondaryRangeName: direct.ValueOf(in.ServicesSecondaryRangeName),
		}
	}
	if in.ServicesIPV4CIDRBlock != nil {
		out.ServicesIpAllocation = &pb.IPAllocationPolicy_ServicesIpv4CidrBlock{
			ServicesIpv4CidrBlock: direct.ValueOf(in.ServicesIPV4CIDRBlock),
		}
	}
	return out
}
