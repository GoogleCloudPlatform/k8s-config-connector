// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.orgpolicy.v2.Policy
// api.group: orgpolicy.cnrm.cloud.google.com

package orgpolicy

import (
	pb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/orgpolicy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func OrgPolicyPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.OrgPolicyPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrgPolicyPolicyObservedState{}
	out.Spec = PolicySpecObservedState_FromProto(mapCtx, in.GetSpec())
	out.DryRunSpec = PolicySpecObservedState_FromProto(mapCtx, in.GetDryRunSpec())
	return out
}
func OrgPolicyPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrgPolicyPolicyObservedState) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	out.Spec = PolicySpecObservedState_ToProto(mapCtx, in.Spec)
	out.DryRunSpec = PolicySpecObservedState_ToProto(mapCtx, in.DryRunSpec)
	return out
}
