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
	"strconv"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krmcomputev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeNetworkEndpointSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEndpoint) *krmcomputev1alpha1.ComputeNetworkEndpointSpec {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.ComputeNetworkEndpointSpec{}
	if in.GetInstance() != "" {
		out.InstanceRef = &krm.InstanceRef{External: in.GetInstance()}
	}
	if in.GetIpAddress() != "" {
		out.IpAddress = in.GetIpAddress()
	}
	if in.Port != nil {
		portStr := strconv.Itoa(int(*in.Port))
		out.ResourceID = &portStr
	}
	return out
}

func ComputeNetworkEndpointSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.ComputeNetworkEndpointSpec) *pb.NetworkEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEndpoint{}
	if in.InstanceRef != nil {
		out.Instance = &in.InstanceRef.External
	}
	if in.IpAddress != "" {
		out.IpAddress = &in.IpAddress
	}
	if in.ResourceID != nil {
		portVal, err := strconv.Atoi(*in.ResourceID)
		if err != nil {
			mapCtx.Errorf("invalid port %q: %v", *in.ResourceID, err)
		} else {
			portInt32 := int32(portVal)
			out.Port = &portInt32
		}
	}
	return out
}

func ComputeNetworkEndpointStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEndpoint) *krmcomputev1alpha1.ComputeNetworkEndpointStatus {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.ComputeNetworkEndpointStatus{}
	return out
}

func ComputeNetworkEndpointStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.ComputeNetworkEndpointStatus) *pb.NetworkEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEndpoint{}
	return out
}
