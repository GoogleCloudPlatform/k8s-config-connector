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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeGlobalNetworkEndpointGroupSpec_v1alpha1_FromProto maps GCP proto definition to KRM.
// This is handcoded because of:
// 1. DefaultPort type mismatch (proto int32 vs KRM int64).
// 2. NetworkEndpointType type mismatch (proto *string vs KRM string).
// 3. ProjectRef cannot be automatically mapped from proto since proto doesn't have projectRef field.
func ComputeGlobalNetworkEndpointGroupSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEndpointGroup) *krm.ComputeGlobalNetworkEndpointGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeGlobalNetworkEndpointGroupSpec{}
	if in.DefaultPort != nil {
		val := int64(*in.DefaultPort)
		out.DefaultPort = &val
	}
	out.Description = in.Description
	if in.NetworkEndpointType != nil {
		out.NetworkEndpointType = *in.NetworkEndpointType
	}
	return out
}

// ComputeGlobalNetworkEndpointGroupSpec_v1alpha1_ToProto maps KRM definition to GCP proto.
// This is handcoded because of:
// 1. DefaultPort type mismatch (proto int32 vs KRM int64).
// 2. NetworkEndpointType type mismatch (proto *string vs KRM string).
// 3. ProjectRef cannot be automatically mapped from proto since proto doesn't have projectRef field.
func ComputeGlobalNetworkEndpointGroupSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeGlobalNetworkEndpointGroupSpec) *pb.NetworkEndpointGroup {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEndpointGroup{}
	if in.DefaultPort != nil {
		val := int32(*in.DefaultPort)
		out.DefaultPort = &val
	}
	out.Description = in.Description
	if in.NetworkEndpointType != "" {
		out.NetworkEndpointType = &in.NetworkEndpointType
	}
	return out
}
