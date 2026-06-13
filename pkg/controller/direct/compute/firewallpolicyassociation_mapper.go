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
	"strings"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeFirewallPolicyAssociationSpec_v1beta1_FromProto maps a GCP proto FirewallPolicyAssociation to a KRM spec.
// This is handwritten because the reference fields (attachmentTargetRef and firewallPolicyRef) have a different type
// structure than the simple string fields in the proto message.
func ComputeFirewallPolicyAssociationSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyAssociation) *krm.ComputeFirewallPolicyAssociationSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFirewallPolicyAssociationSpec{}
	if in.GetAttachmentTarget() != "" {
		out.AttachmentTargetRef.External = in.GetAttachmentTarget()
		if strings.HasPrefix(in.GetAttachmentTarget(), "folders/") {
			out.AttachmentTargetRef.Kind = "Folder"
		}
	}
	if in.GetFirewallPolicyId() != "" {
		out.FirewallPolicyRef.External = in.GetFirewallPolicyId()
	}
	if in.GetName() != "" {
		out.ResourceID = direct.LazyPtr(in.GetName())
	}
	return out
}

// ComputeFirewallPolicyAssociationSpec_v1beta1_ToProto maps a KRM spec to a GCP proto FirewallPolicyAssociation.
// This is handwritten because the reference fields (attachmentTargetRef and firewallPolicyRef) have a different type
// structure than the simple string fields in the proto message.
func ComputeFirewallPolicyAssociationSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFirewallPolicyAssociationSpec) *pb.FirewallPolicyAssociation {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyAssociation{}
	if in.AttachmentTargetRef.External != "" {
		out.AttachmentTarget = direct.LazyPtr(in.AttachmentTargetRef.External)
	}
	if in.FirewallPolicyRef.External != "" {
		out.FirewallPolicyId = direct.LazyPtr(in.FirewallPolicyRef.External)
	}
	if in.ResourceID != nil {
		out.Name = in.ResourceID
	}
	return out
}

// ComputeFirewallPolicyAssociationStatus_v1beta1_FromProto maps a GCP proto FirewallPolicyAssociation to a KRM status.
func ComputeFirewallPolicyAssociationStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyAssociation) *krm.ComputeFirewallPolicyAssociationStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeFirewallPolicyAssociationStatus{}
	if in.GetShortName() != "" {
		out.ShortName = direct.LazyPtr(in.GetShortName())
	}
	return out
}

// ComputeFirewallPolicyAssociationStatus_v1beta1_ToProto maps a KRM status to a GCP proto FirewallPolicyAssociation.
func ComputeFirewallPolicyAssociationStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeFirewallPolicyAssociationStatus) *pb.FirewallPolicyAssociation {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyAssociation{}
	if in.ShortName != nil {
		out.ShortName = in.ShortName
	}
	return out
}
