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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_FromProto converts a proto FirewallPolicyAssociation to a KRM ComputeOrganizationSecurityPolicyAssociationSpec.
// We handcode this because in KRM, AttachmentTarget and FirewallPolicyID are non-pointers (required strings), but in proto they are pointer strings.
func ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyAssociation) *krm.ComputeOrganizationSecurityPolicyAssociationSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeOrganizationSecurityPolicyAssociationSpec{}
	out.AttachmentTarget = in.GetAttachmentTarget()
	out.FirewallPolicyID = in.GetFirewallPolicyId()
	return out
}

// ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_ToProto converts a KRM ComputeOrganizationSecurityPolicyAssociationSpec to a proto FirewallPolicyAssociation.
// We handcode this because in KRM, AttachmentTarget and FirewallPolicyID are non-pointers (required strings), but in proto they are pointer strings.
func ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeOrganizationSecurityPolicyAssociationSpec) *pb.FirewallPolicyAssociation {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyAssociation{}
	if in.AttachmentTarget != "" {
		out.AttachmentTarget = &in.AttachmentTarget
	}
	if in.FirewallPolicyID != "" {
		out.FirewallPolicyId = &in.FirewallPolicyID
	}
	return out
}
