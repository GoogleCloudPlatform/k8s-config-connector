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

// +generated:mapper
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.compute.v1

package compute

import (
	"strconv"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeGlobalNetworkEndpointSpec_v1alpha1_FromProto maps a pb.NetworkEndpoint to a krm.ComputeGlobalNetworkEndpointSpec
func ComputeGlobalNetworkEndpointSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEndpoint) *krm.ComputeGlobalNetworkEndpointSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeGlobalNetworkEndpointSpec{}
	out.FQDN = in.Fqdn
	out.IPAddress = in.IpAddress
	if in.Port != nil {
		portStr := strconv.FormatInt(int64(*in.Port), 10)
		out.ResourceID = &portStr
	}
	return out
}

// ComputeGlobalNetworkEndpointSpec_v1alpha1_ToProto maps a krm.ComputeGlobalNetworkEndpointSpec to a pb.NetworkEndpoint
func ComputeGlobalNetworkEndpointSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeGlobalNetworkEndpointSpec) *pb.NetworkEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEndpoint{}
	out.Fqdn = in.FQDN
	out.IpAddress = in.IPAddress
	if in.ResourceID != nil {
		port, err := strconv.ParseInt(*in.ResourceID, 10, 32)
		if err == nil {
			p := int32(port)
			out.Port = &p
		} else {
			mapCtx.Errorf("invalid port value %q: %v", *in.ResourceID, err)
		}
	}
	return out
}
