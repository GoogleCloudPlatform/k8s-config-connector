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

func ComputeSSLPolicySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SslPolicy) *krm.ComputeSSLPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeSSLPolicySpec{}
	out.CustomFeatures = in.CustomFeatures
	out.Description = in.Description
	out.MinTlsVersion = in.MinTlsVersion
	out.Profile = in.Profile
	out.ResourceID = in.Name
	return out
}

func ComputeSSLPolicySpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeSSLPolicySpec) *pb.SslPolicy {
	if in == nil {
		return nil
	}
	out := &pb.SslPolicy{}
	out.CustomFeatures = in.CustomFeatures
	out.Description = in.Description
	out.MinTlsVersion = in.MinTlsVersion
	out.Profile = in.Profile
	out.Name = in.ResourceID
	return out
}

func ComputeSSLPolicyStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SslPolicy) *krm.ComputeSSLPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeSSLPolicyStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.EnabledFeatures = in.EnabledFeatures
	out.Fingerprint = in.Fingerprint
	out.SelfLink = in.SelfLink
	return out
}

func ComputeSSLPolicyStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeSSLPolicyStatus) *pb.SslPolicy {
	if in == nil {
		return nil
	}
	out := &pb.SslPolicy{}
	out.CreationTimestamp = in.CreationTimestamp
	out.EnabledFeatures = in.EnabledFeatures
	out.Fingerprint = in.Fingerprint
	out.SelfLink = in.SelfLink
	return out
}
