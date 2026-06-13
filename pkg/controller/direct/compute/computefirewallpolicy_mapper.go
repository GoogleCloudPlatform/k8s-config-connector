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

// ComputeFirewallPolicySpec_v1beta1_FromProto is handcoded because ShortName is a pointer *string in the proto but a required non-pointer string in the KRM schema.
func ComputeFirewallPolicySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicy) *krm.ComputeFirewallPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFirewallPolicySpec{}
	out.Description = in.Description
	out.ShortName = direct.ValueOf(in.ShortName)
	return out
}

// ComputeFirewallPolicySpec_v1beta1_ToProto is handcoded because ShortName is a pointer *string in the proto but a required non-pointer string in the KRM schema.
func ComputeFirewallPolicySpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFirewallPolicySpec) *pb.FirewallPolicy {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicy{}
	out.Description = in.Description
	out.ShortName = direct.LazyPtr(in.ShortName)
	return out
}

func ComputeFirewallPolicyStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicy) *krm.ComputeFirewallPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFirewallPolicyStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Fingerprint = in.Fingerprint
	if in.Id != nil {
		idStr := strconv.FormatUint(*in.Id, 10)
		out.ID = &idStr
	}
	if in.RuleTupleCount != nil {
		rtc := int64(*in.RuleTupleCount)
		out.RuleTupleCount = &rtc
	}
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithId
	return out
}

func ComputeFirewallPolicyStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFirewallPolicyStatus) *pb.FirewallPolicy {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicy{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Fingerprint = in.Fingerprint
	if in.ID != nil {
		idVal, err := strconv.ParseUint(*in.ID, 10, 64)
		if err != nil {
			mapCtx.Errorf("error converting ID string %s to uint64: %v", *in.ID, err)
		} else {
			out.Id = &idVal
		}
	}
	if in.RuleTupleCount != nil {
		rtc := int32(*in.RuleTupleCount)
		out.RuleTupleCount = &rtc
	}
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithId
	return out
}
