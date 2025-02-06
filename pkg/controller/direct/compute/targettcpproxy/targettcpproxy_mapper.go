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

package targettcpproxy

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"strings"
)

func ComputeTargetTCPProxySpec_FromProto(mapCtx *direct.MapContext, in *pb.TargetTcpProxy) *krm.ComputeTargetTCPProxySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetTCPProxySpec{}
	out.Description = in.Description
	out.ProxyBind = in.ProxyBind
	out.ProxyHeader = in.ProxyHeader
	if in.Region != nil {
		// Convert `https://www.googleapis.com/compute/v1/projects/projectId/regions/europe-west4` to `europe-west4`
		tokens := strings.Split(direct.ValueOf(in.Region), "/")
		out.Location = direct.LazyPtr(tokens[len(tokens)-1])
	}
	if in.Service != nil {
		out.BackendServiceRef = &refs.ComputeBackendServiceRef{External: direct.ValueOf(in.Service)}
	}
	return out
}
func ComputeTargetTCPProxySpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeTargetTCPProxySpec) *pb.TargetTcpProxy {
	if in == nil {
		return nil
	}
	out := &pb.TargetTcpProxy{}
	out.Description = in.Description
	out.ProxyBind = in.ProxyBind
	out.ProxyHeader = in.ProxyHeader
	out.Region = in.Location
	if in.BackendServiceRef != nil {
		out.Service = direct.LazyPtr(in.BackendServiceRef.External)
	}
	return out
}
func ComputeTargetTCPProxyStatus_FromProto(mapCtx *direct.MapContext, in *pb.TargetTcpProxy) *krm.ComputeTargetTCPProxyStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetTCPProxyStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	// Kind
	out.ProxyId = direct.LazyPtr(int64(in.GetId()))
	out.SelfLink = in.SelfLink
	return out
}
func ComputeTargetTCPProxyStatus_ToProto(mapCtx *direct.MapContext, in *krm.ComputeTargetTCPProxyStatus) *pb.TargetTcpProxy {
	if in == nil {
		return nil
	}
	out := &pb.TargetTcpProxy{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Id = direct.LazyPtr(uint64(*in.ProxyId))
	// Kind
	out.SelfLink = in.SelfLink
	return out
}
