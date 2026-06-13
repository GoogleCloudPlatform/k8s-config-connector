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
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeFirewallPolicyAssociationSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyAssociation) *krm.ComputeFirewallPolicyAssociationSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFirewallPolicyAssociationSpec{}
	if in.GetAttachmentTarget() != "" {
		out.AttachmentTargetRef = k8sv1alpha1.ResourceRef{External: in.GetAttachmentTarget()}
	}
	return out
}

func ComputeFirewallPolicyAssociationSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFirewallPolicyAssociationSpec) *pb.FirewallPolicyAssociation {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyAssociation{}
	if in.AttachmentTargetRef.External != "" {
		out.AttachmentTarget = &in.AttachmentTargetRef.External
	}
	return out
}
