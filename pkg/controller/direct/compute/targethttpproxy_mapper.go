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

// ComputeTargetHTTPProxySpec_v1beta1_ToProto converts a KRM Spec to a Proto message.
// This is handcoded because of UrlMapRef external resolution and HttpKeepAliveTimeoutSec type conversion.
func ComputeTargetHTTPProxySpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeTargetHTTPProxySpec) *pb.TargetHttpProxy {
	if in == nil {
		return nil
	}
	out := &pb.TargetHttpProxy{}
	if in.HttpKeepAliveTimeoutSec != nil {
		timeout := int32(*in.HttpKeepAliveTimeoutSec)
		out.HttpKeepAliveTimeoutSec = &timeout
	}
	out.ProxyBind = in.ProxyBind
	out.Description = in.Description

	if in.UrlMapRef != nil {
		out.UrlMap = &in.UrlMapRef.External
	}

	return out
}

// ComputeTargetHTTPProxySpec_v1beta1_FromProto converts a Proto message to a KRM Spec.
// This is handcoded because of UrlMapRef external resolution and HttpKeepAliveTimeoutSec type conversion.
func ComputeTargetHTTPProxySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TargetHttpProxy) *krm.ComputeTargetHTTPProxySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetHTTPProxySpec{}
	if in.HttpKeepAliveTimeoutSec != nil {
		timeout := int(*in.HttpKeepAliveTimeoutSec)
		out.HttpKeepAliveTimeoutSec = &timeout
	}
	out.ProxyBind = in.ProxyBind
	out.Description = in.Description

	if in.UrlMap != nil {
		out.UrlMapRef = &krm.ComputeURLMapRef{External: *in.UrlMap}
	}

	return out
}

// ComputeTargetHTTPProxyStatus_v1beta1_FromProto converts a Proto message to a KRM Status.
// This is handcoded because of ProxyId conversion.
func ComputeTargetHTTPProxyStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TargetHttpProxy) *krm.ComputeTargetHTTPProxyStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetHTTPProxyStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	if in.Id != nil {
		id := int64(*in.Id)
		out.ProxyId = &id
	}
	out.SelfLink = in.SelfLink
	return out
}
