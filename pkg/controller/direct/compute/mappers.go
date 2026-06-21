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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeOrganizationSecurityPolicySpec_v1alpha1_FromProto converts a v1.SecurityPolicy proto to a v1alpha1.ComputeOrganizationSecurityPolicySpec.
// We hand-code this function because KRM displayName and parent map to proto ShortName and Parent.
func ComputeOrganizationSecurityPolicySpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicy) *krmcomputev1alpha1.ComputeOrganizationSecurityPolicySpec {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.ComputeOrganizationSecurityPolicySpec{}
	out.Description = in.Description
	out.DisplayName = in.GetShortName()
	out.Parent = in.GetParent()
	out.Type = in.Type
	return out
}

// ComputeOrganizationSecurityPolicySpec_v1alpha1_ToProto converts a v1alpha1.ComputeOrganizationSecurityPolicySpec to a v1.SecurityPolicy proto.
// We hand-code this function because KRM displayName and parent map to proto ShortName and Parent.
func ComputeOrganizationSecurityPolicySpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.ComputeOrganizationSecurityPolicySpec) *pb.SecurityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicy{}
	out.Description = in.Description
	if in.DisplayName != "" {
		out.ShortName = &in.DisplayName
	}
	if in.Parent != "" {
		out.Parent = &in.Parent
	}
	out.Type = in.Type
	return out
}

// ComputeOrganizationSecurityPolicyStatus_v1alpha1_FromProto converts a v1.SecurityPolicy proto to a v1alpha1.ComputeOrganizationSecurityPolicyStatus.
// We hand-code this function because policyId maps to Id (which requires string to uint64 conversion).
func ComputeOrganizationSecurityPolicyStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicy) *krmcomputev1alpha1.ComputeOrganizationSecurityPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.ComputeOrganizationSecurityPolicyStatus{}
	out.Fingerprint = in.Fingerprint
	if in.Id != nil {
		idStr := strconv.FormatUint(*in.Id, 10)
		out.PolicyId = &idStr
	}
	return out
}

// ComputeOrganizationSecurityPolicyStatus_v1alpha1_ToProto converts a v1alpha1.ComputeOrganizationSecurityPolicyStatus to a v1.SecurityPolicy proto.
// We hand-code this function because policyId maps to Id (which requires string to uint64 conversion).
func ComputeOrganizationSecurityPolicyStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.ComputeOrganizationSecurityPolicyStatus) *pb.SecurityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicy{}
	out.Fingerprint = in.Fingerprint
	if in.PolicyId != nil {
		idVal, err := strconv.ParseUint(*in.PolicyId, 10, 64)
		if err != nil {
			mapCtx.Errorf("parsing policyId %q: %v", *in.PolicyId, err)
		} else {
			out.Id = &idVal
		}
	}
	return out
}
