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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeNetworkFirewallPolicyStatus_v1beta1_FromProto is handcoded because status fields are mapped
// directly to status (rather than a nested observedState object) for strict schema compatibility.
// It also handles type conversion between protobuf uint64/int32 and KRM string/int64.
func ComputeNetworkFirewallPolicyStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicy) *krm.ComputeNetworkFirewallPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkFirewallPolicyStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Fingerprint = in.Fingerprint
	if in.Id != nil {
		idStr := strconv.FormatUint(*in.Id, 10)
		out.NetworkFirewallPolicyId = &idStr
	}
	if in.RuleTupleCount != nil {
		count := int64(*in.RuleTupleCount)
		out.RuleTupleCount = &count
	}
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithId
	return out
}

// ComputeNetworkFirewallPolicyStatus_v1beta1_ToProto is handcoded because status fields are mapped
// directly to status (rather than a nested observedState object) for strict schema compatibility.
// It also handles type conversion between KRM string/int64 and protobuf uint64/int32.
func ComputeNetworkFirewallPolicyStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkFirewallPolicyStatus) *pb.FirewallPolicy {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicy{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Fingerprint = in.Fingerprint
	if in.NetworkFirewallPolicyId != nil {
		id, err := strconv.ParseUint(*in.NetworkFirewallPolicyId, 10, 64)
		if err != nil {
			mapCtx.Errorf("error parsing ID %q: %v", *in.NetworkFirewallPolicyId, err)
		} else {
			out.Id = &id
		}
	}
	if in.RuleTupleCount != nil {
		count := int32(*in.RuleTupleCount)
		out.RuleTupleCount = &count
	}
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithId
	return out
}
