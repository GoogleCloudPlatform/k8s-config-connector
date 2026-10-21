// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compute

import (
	computepb "cloud.google.com/go/compute/apiv1beta/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_FromProto maps a GCP proto SecurityPolicyAssociation to a KRM spec.
// We hand-code this function because AttachmentID and PolicyID are string fields in KRM but *string in proto,
// and PolicyID maps to SecurityPolicyId in the proto message, while ResourceID maps to Name.
func ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *computepb.SecurityPolicyAssociation) *krm.ComputeOrganizationSecurityPolicyAssociationSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeOrganizationSecurityPolicyAssociationSpec{}
	out.AttachmentID = in.GetAttachmentId()
	out.PolicyID = in.GetSecurityPolicyId()
	if in.GetName() != "" {
		out.ResourceID = direct.LazyPtr(in.GetName())
	}
	return out
}

// ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_ToProto maps a KRM spec to a GCP proto SecurityPolicyAssociation.
// We hand-code this function because AttachmentID and PolicyID are string fields in KRM but *string in proto,
// and PolicyID maps to SecurityPolicyId in the proto message, while ResourceID maps to Name.
func ComputeOrganizationSecurityPolicyAssociationSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeOrganizationSecurityPolicyAssociationSpec) *computepb.SecurityPolicyAssociation {
	if in == nil {
		return nil
	}
	out := &computepb.SecurityPolicyAssociation{}
	if in.AttachmentID != "" {
		out.AttachmentId = direct.LazyPtr(in.AttachmentID)
	}
	if in.PolicyID != "" {
		out.SecurityPolicyId = direct.LazyPtr(in.PolicyID)
	}
	if in.ResourceID != nil {
		out.Name = in.ResourceID
	}
	return out
}

// ComputeOrganizationSecurityPolicyAssociationStatus_v1alpha1_FromProto maps a GCP proto SecurityPolicyAssociation to a KRM status.
func ComputeOrganizationSecurityPolicyAssociationStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *computepb.SecurityPolicyAssociation) *krm.ComputeOrganizationSecurityPolicyAssociationStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeOrganizationSecurityPolicyAssociationStatus{}
	out.DisplayName = in.DisplayName
	return out
}

// ComputeOrganizationSecurityPolicyAssociationStatus_v1alpha1_ToProto maps a KRM status to a GCP proto SecurityPolicyAssociation.
func ComputeOrganizationSecurityPolicyAssociationStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeOrganizationSecurityPolicyAssociationStatus) *computepb.SecurityPolicyAssociation {
	if in == nil {
		return nil
	}
	out := &computepb.SecurityPolicyAssociation{}
	out.DisplayName = in.DisplayName
	return out
}
