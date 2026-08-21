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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeTargetGRPCProxySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TargetGrpcProxy) *krm.ComputeTargetGRPCProxySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetGRPCProxySpec{}
	out.Description = in.Description
	if in.UrlMap != nil {
		out.UrlMapRef = &krm.ComputeURLMapRef{External: direct.ValueOf(in.UrlMap)}
	}
	out.ValidateForProxyless = in.ValidateForProxyless
	return out
}

func ComputeTargetGRPCProxySpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeTargetGRPCProxySpec) *pb.TargetGrpcProxy {
	if in == nil {
		return nil
	}
	out := &pb.TargetGrpcProxy{}
	out.Description = in.Description
	if in.UrlMapRef != nil {
		out.UrlMap = direct.LazyPtr(in.UrlMapRef.External)
	}
	out.ValidateForProxyless = in.ValidateForProxyless
	return out
}

func ComputeTargetGRPCProxyStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TargetGrpcProxy) *krm.ComputeTargetGRPCProxyStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetGRPCProxyStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Fingerprint = in.Fingerprint
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithId
	return out
}

func ComputeTargetGRPCProxyStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeTargetGRPCProxyStatus) *pb.TargetGrpcProxy {
	if in == nil {
		return nil
	}
	out := &pb.TargetGrpcProxy{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Fingerprint = in.Fingerprint
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithId
	return out
}
