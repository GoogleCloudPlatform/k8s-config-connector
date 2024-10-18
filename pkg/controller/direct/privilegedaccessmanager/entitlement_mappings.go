// Copyright 2024 Google LLC
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

package privilegedaccessmanager

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privilegedaccessmanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	pb "cloud.google.com/go/privilegedaccessmanager/apiv1/privilegedaccessmanagerpb"
)

type gcpIAMAccessResource struct {
	resourceType string
	resource     string
}

func GcpIamAccess_FromProto(mapCtx *direct.MapContext, in *pb.PrivilegedAccess_GcpIamAccess) *krm.GcpIamAccess {
	if in == nil {
		return nil
	}
	out := &krm.GcpIamAccess{}
	out.RoleBindings = direct.Slice_FromProto(mapCtx, in.RoleBindings, RoleBinding_FromProto)
	return out
}
func GcpIamAccess_ToProto(mapCtx *direct.MapContext, in *krm.GcpIamAccess, hiddenFields gcpIAMAccessResource) *pb.PrivilegedAccess_GcpIamAccess {
	if in == nil {
		return nil
	}
	out := &pb.PrivilegedAccess_GcpIamAccess{}
	out.ResourceType = hiddenFields.resourceType
	out.Resource = hiddenFields.resource
	out.RoleBindings = direct.Slice_ToProto(mapCtx, in.RoleBindings, RoleBinding_ToProto)
	return out
}
func ManualApprovals_FromProto(mapCtx *direct.MapContext, in *pb.ManualApprovals) *krm.ManualApprovals {
	if in == nil {
		return nil
	}
	out := &krm.ManualApprovals{}
	out.RequireApproverJustification = direct.LazyPtr(in.GetRequireApproverJustification())
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, Step_FromProto)
	return out
}
func ManualApprovals_ToProto(mapCtx *direct.MapContext, in *krm.ManualApprovals) *pb.ManualApprovals {
	if in == nil {
		return nil
	}
	out := &pb.ManualApprovals{}
	out.RequireApproverJustification = direct.ValueOf(in.RequireApproverJustification)
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, Step_ToProto)
	return out
}
func PrivilegedAccess_FromProto(mapCtx *direct.MapContext, in *pb.PrivilegedAccess) *krm.PrivilegedAccess {
	if in == nil {
		return nil
	}
	out := &krm.PrivilegedAccess{}
	out.GcpIAMAccess = GcpIamAccess_FromProto(mapCtx, in.GetGcpIamAccess())
	return out
}
func PrivilegedAccess_ToProto(mapCtx *direct.MapContext, in *krm.PrivilegedAccess, hiddenFields gcpIAMAccessResource) *pb.PrivilegedAccess {
	if in == nil {
		return nil
	}
	out := &pb.PrivilegedAccess{}
	if oneof := GcpIamAccess_ToProto(mapCtx, in.GcpIAMAccess, hiddenFields); oneof != nil {
		out.AccessType = &pb.PrivilegedAccess_GcpIamAccess_{GcpIamAccess: oneof}
	}
	return out
}
func PrivilegedAccessManagerEntitlementSpec_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement) *krm.PrivilegedAccessManagerEntitlementSpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivilegedAccessManagerEntitlementSpec{}
	out.EligibleUsers = direct.Slice_FromProto(mapCtx, in.EligibleUsers, AccessControlEntry_FromProto)
	out.ApprovalWorkflow = ApprovalWorkflow_FromProto(mapCtx, in.GetApprovalWorkflow())
	out.PrivilegedAccess = PrivilegedAccess_FromProto(mapCtx, in.GetPrivilegedAccess())
	out.MaxRequestDuration = direct.StringDuration_FromProto(mapCtx, in.GetMaxRequestDuration())
	out.RequesterJustificationType = RequesterJustificationType_FromProto(mapCtx, in.GetRequesterJustificationConfig())
	out.AdditionalNotificationTargets = AdditionalNotificationTargets_FromProto(mapCtx, in.GetAdditionalNotificationTargets())
	return out
}
func PrivilegedAccessManagerEntitlementSpec_ToProto(mapCtx *direct.MapContext, in *krm.PrivilegedAccessManagerEntitlementSpec, hiddenFields gcpIAMAccessResource) *pb.Entitlement {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement{}
	// MISSING: Name
	out.EligibleUsers = direct.Slice_ToProto(mapCtx, in.EligibleUsers, AccessControlEntry_ToProto)
	out.ApprovalWorkflow = ApprovalWorkflow_ToProto(mapCtx, in.ApprovalWorkflow)
	out.PrivilegedAccess = PrivilegedAccess_ToProto(mapCtx, in.PrivilegedAccess, hiddenFields)
	out.MaxRequestDuration = direct.StringDuration_ToProto(mapCtx, in.MaxRequestDuration)
	out.RequesterJustificationConfig = RequesterJustificationType_ToProto(mapCtx, in.RequesterJustificationType)
	out.AdditionalNotificationTargets = AdditionalNotificationTargets_ToProto(mapCtx, in.AdditionalNotificationTargets)
	return out
}
func PrivilegedAccessManagerEntitlementStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement) *krm.PrivilegedAccessManagerEntitlementObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivilegedAccessManagerEntitlementObservedState{}
	out.State = direct.LazyPtr(in.State.String())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.UpdateTime)
	out.Etag = direct.LazyPtr(in.Etag)
	return out
}
func RequesterJustificationType_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement_RequesterJustificationConfig) *string {
	if in == nil {
		return nil
	}
	outVal := ""
	if in.GetNotMandatory() != nil {
		outVal = krm.RequesterJustificationTypeNotMandatory
	} else if in.GetUnstructured() != nil {
		outVal = krm.RequesterJustificationTypeUnstructured
	} else {
		mapCtx.Errorf("neither 'notMandatory' nor 'unstructured' set under 'requesterJustificationConfig': one of them must be set")
		return nil
	}
	return &outVal
}
func RequesterJustificationType_ToProto(mapCtx *direct.MapContext, in *string) *pb.Entitlement_RequesterJustificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement_RequesterJustificationConfig{}
	switch *in {
	case krm.RequesterJustificationTypeNotMandatory:
		out.JustificationType = &pb.Entitlement_RequesterJustificationConfig_NotMandatory_{
			NotMandatory: &pb.Entitlement_RequesterJustificationConfig_NotMandatory{},
		}
	case krm.RequesterJustificationTypeUnstructured:
		out.JustificationType = &pb.Entitlement_RequesterJustificationConfig_Unstructured_{
			Unstructured: &pb.Entitlement_RequesterJustificationConfig_Unstructured{},
		}
	default:
		mapCtx.Errorf("unknown enum value %q for 'spec.requesterJustificationType' (valid values are %v)",
			*in, krm.ValidRequesterJustificationTypes)
	}
	return out
}
func Step_FromProto(mapCtx *direct.MapContext, in *pb.ManualApprovals_Step) *krm.Step {
	if in == nil {
		return nil
	}
	out := &krm.Step{}
	out.Approvers = direct.Slice_FromProto(mapCtx, in.Approvers, AccessControlEntry_FromProto)
	out.ApprovalsNeeded = direct.LazyPtr(in.GetApprovalsNeeded())
	out.ApproverEmailRecipients = in.ApproverEmailRecipients
	return out
}
func Step_ToProto(mapCtx *direct.MapContext, in *krm.Step) *pb.ManualApprovals_Step {
	if in == nil {
		return nil
	}
	out := &pb.ManualApprovals_Step{}
	out.Approvers = direct.Slice_ToProto(mapCtx, in.Approvers, AccessControlEntry_ToProto)
	out.ApprovalsNeeded = direct.ValueOf(in.ApprovalsNeeded)
	out.ApproverEmailRecipients = in.ApproverEmailRecipients
	return out
}
