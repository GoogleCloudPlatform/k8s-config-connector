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

// +generated:mapper
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1beta

package compute

import (
	pb "cloud.google.com/go/compute/apiv1beta/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TargetTCPProxy_FromProto(mapCtx *direct.MapContext, in *pb.TargetTcpProxy) *krm.TargetTCPProxy {
	if in == nil {
		return nil
	}
	out := &krm.TargetTCPProxy{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Description = in.Description
	out.ID = in.Id
	out.Kind = in.Kind
	out.Name = in.Name
	out.ProxyBind = in.ProxyBind
	out.ProxyHeader = in.ProxyHeader
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	out.Service = in.Service
	return out
}
func TargetTCPProxy_ToProto(mapCtx *direct.MapContext, in *krm.TargetTCPProxy) *pb.TargetTcpProxy {
	if in == nil {
		return nil
	}
	out := &pb.TargetTcpProxy{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Description = in.Description
	out.Id = in.ID
	out.Kind = in.Kind
	out.Name = in.Name
	out.ProxyBind = in.ProxyBind
	out.ProxyHeader = in.ProxyHeader
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	out.Service = in.Service
	return out
}
