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

func ComputeNetworkFirewallPolicyAssociationStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyAssociation) *krm.ComputeNetworkFirewallPolicyAssociationStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkFirewallPolicyAssociationStatus{}
	if in.ShortName != nil {
		out.ShortName = in.ShortName
	}
	return out
}

func ComputeNetworkFirewallPolicyAssociationStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkFirewallPolicyAssociationStatus) *pb.FirewallPolicyAssociation {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyAssociation{}
	if in.ShortName != nil {
		out.ShortName = in.ShortName
	}
	return out
}
