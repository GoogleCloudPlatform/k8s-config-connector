// Copyright 2024 Google LLC
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

package regiontargettcpproxy

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeRegionTargetTCPProxySpec_FromProto(mapCtx *direct.MapContext, in *pb.TargetTcpProxy) *krm.ComputeRegionTargetTCPProxySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRegionTargetTCPProxySpec{}
	out.Description = in.Description
	out.ProxyBind = in.ProxyBind
	out.ProxyHeader = in.ProxyHeader
	out.BackendServiceRef = ComputeRegionTargetTCPProxySpec_BackendServiceRef_FromProto(mapCtx, direct.ValueOf(in.Service))
	return out
}
func ComputeRegionTargetTCPProxySpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRegionTargetTCPProxySpec) *pb.TargetTcpProxy {
	if in == nil {
		return nil
	}
	out := &pb.TargetTcpProxy{}
	out.Description = in.Description
	out.ProxyBind = in.ProxyBind
	out.ProxyHeader = in.ProxyHeader
	out.Service = ComputeRegionTargetTCPProxySpec_BackendServiceRef_ToProto(mapCtx, in.BackendServiceRef)
	return out
}
func ComputeRegionTargetTCPProxyStatus_FromProto(mapCtx *direct.MapContext, in *pb.TargetTcpProxy) *krm.ComputeRegionTargetTCPProxyStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRegionTargetTCPProxyStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Kind = in.Kind
	out.ProxyId = direct.LazyPtr(int64(in.GetId()))
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	return out
}
func ComputeRegionTargetTCPProxyStatus_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRegionTargetTCPProxyStatus) *pb.TargetTcpProxy {
	if in == nil {
		return nil
	}
	out := &pb.TargetTcpProxy{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Id = direct.LazyPtr(uint64(*in.ProxyId))
	out.Kind = in.Kind
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	return out
}
