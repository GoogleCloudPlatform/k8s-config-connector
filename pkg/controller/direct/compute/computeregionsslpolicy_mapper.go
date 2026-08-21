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
// krm.version: v1alpha1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeRegionSSLPolicySpec_v1alpha1_FromProto maps SslPolicy proto to ComputeRegionSSLPolicySpec.
// This function is handcoded because the proto field 'region' is an optional *string (pointer)
// while the KRM field 'Region' is a required/immutable string (non-pointer).
func ComputeRegionSSLPolicySpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SslPolicy) *krm.ComputeRegionSSLPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRegionSSLPolicySpec{}
	out.CustomFeatures = in.CustomFeatures
	out.Description = in.Description
	out.MinTLSVersion = in.MinTlsVersion
	out.Profile = in.Profile
	if in.Region != nil {
		out.Region = *in.Region
	}
	return out
}

// ComputeRegionSSLPolicySpec_v1alpha1_ToProto maps ComputeRegionSSLPolicySpec to SslPolicy proto.
// This function is handcoded because the proto field 'region' is an optional *string (pointer)
// while the KRM field 'Region' is a required/immutable string (non-pointer).
func ComputeRegionSSLPolicySpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRegionSSLPolicySpec) *pb.SslPolicy {
	if in == nil {
		return nil
	}
	out := &pb.SslPolicy{}
	out.CustomFeatures = in.CustomFeatures
	out.Description = in.Description
	out.MinTlsVersion = in.MinTLSVersion
	out.Profile = in.Profile
	if in.Region != "" {
		out.Region = &in.Region
	}
	return out
}

// ComputeRegionSSLPolicyStatus_v1alpha1_FromProto maps SslPolicy proto to ComputeRegionSSLPolicyStatus.
// This function is handcoded to map the output-only status fields directly from the SslPolicy proto.
func ComputeRegionSSLPolicyStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SslPolicy) *krm.ComputeRegionSSLPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeRegionSSLPolicyStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.EnabledFeatures = in.EnabledFeatures
	out.Fingerprint = in.Fingerprint
	out.SelfLink = in.SelfLink
	return out
}

// ComputeRegionSSLPolicyStatus_v1alpha1_ToProto maps ComputeRegionSSLPolicyStatus to SslPolicy proto.
// This function is handcoded to map the output-only status fields directly to the SslPolicy proto.
func ComputeRegionSSLPolicyStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeRegionSSLPolicyStatus) *pb.SslPolicy {
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
