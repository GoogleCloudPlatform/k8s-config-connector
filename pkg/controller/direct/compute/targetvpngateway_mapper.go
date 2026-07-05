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

// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeTargetVPNGatewayStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TargetVpnGateway) *krm.ComputeTargetVPNGatewayStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetVPNGatewayStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	if in.Id != nil {
		val := int64(*in.Id)
		out.ID = &val
	}
	out.SelfLink = in.SelfLink
	return out
}

func ComputeTargetVPNGatewayStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeTargetVPNGatewayStatus) *pb.TargetVpnGateway {
	if in == nil {
		return nil
	}
	out := &pb.TargetVpnGateway{}
	out.CreationTimestamp = in.CreationTimestamp
	if in.ID != nil {
		val := uint64(*in.ID)
		out.Id = &val
	}
	out.SelfLink = in.SelfLink
	return out
}
