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
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privilegedaccessmanager/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	pb "cloud.google.com/go/privilegedaccessmanager/apiv1/privilegedaccessmanagerpb"
	"k8s.io/apimachinery/pkg/runtime"
)

func GcpIamAccess_FromProto(mapCtx *direct.MapContext, in *pb.PrivilegedAccess_GcpIamAccess) *krm.GcpIamAccess {
	if in == nil {
		return nil
	}
	out := &krm.GcpIamAccess{}
	resourceType := in.GetResourceType()
	out.ResourceType = direct.LazyPtr(resourceType)
	resource := in.GetResource()
	externalVal := strings.TrimPrefix(resource, "//cloudresourcemanager.googleapis.com/")
	switch resourceType {
	case "cloudresourcemanager.googleapis.com/Project":
		out.ProjectRef = &refs.ProjectRef{
			External: externalVal,
		}
	case "cloudresourcemanager.googleapis.com/Folder":
		out.FolderRef = &refs.FolderRef{
			External: externalVal,
		}
	case "cloudresourcemanager.googleapis.com/Organization":
		out.OrganizationRef = &refs.OrganizationRef{
			External: externalVal,
		}
	}
	out.RoleBindings = direct.SliceOfPointers_FromProto(mapCtx, in.RoleBindings, RoleBinding_FromProto)
	return out
}
func GcpIamAccess_ToProto(mapCtx *direct.MapContext, in *krm.GcpIamAccess) *pb.PrivilegedAccess_GcpIamAccess {
	if in == nil {
		return nil
	}
	out := &pb.PrivilegedAccess_GcpIamAccess{}
	resourceType := in.ResourceType
	out.ResourceType = direct.ValueOf(resourceType)
	switch *resourceType {
	case "cloudresourcemanager.googleapis.com/Project":
		out.Resource = fmt.Sprintf("//cloudresourcemanager.googleapis.com/%s", in.ProjectRef.External)
	case "cloudresourcemanager.googleapis.com/Folder":
		out.Resource = fmt.Sprintf("//cloudresourcemanager.googleapis.com/%s", in.FolderRef.External)
	case "cloudresourcemanager.googleapis.com/Organization":
		out.Resource = fmt.Sprintf("//cloudresourcemanager.googleapis.com/%s", in.OrganizationRef.External)
	}
	out.RoleBindings = direct.SliceOfPointers_ToProto(mapCtx, in.RoleBindings, RoleBinding_ToProto)
	return out
}
func ManualApprovals_FromProto(mapCtx *direct.MapContext, in *pb.ManualApprovals) *krm.ManualApprovals {
	if in == nil {
		return nil
	}
	out := &krm.ManualApprovals{}
	out.RequireApproverJustification = direct.LazyPtr(in.GetRequireApproverJustification())
	out.Steps = direct.SliceOfPointers_FromProto(mapCtx, in.Steps, Step_FromProto)
	return out
}
func ManualApprovals_ToProto(mapCtx *direct.MapContext, in *krm.ManualApprovals) *pb.ManualApprovals {
	if in == nil {
		return nil
	}
	out := &pb.ManualApprovals{}
	out.RequireApproverJustification = direct.ValueOf(in.RequireApproverJustification)
	out.Steps = direct.SliceOfPointers_ToProto(mapCtx, in.Steps, Step_ToProto)
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
func PrivilegedAccess_ToProto(mapCtx *direct.MapContext, in *krm.PrivilegedAccess) *pb.PrivilegedAccess {
	if in == nil {
		return nil
	}
	out := &pb.PrivilegedAccess{}
	if oneof := GcpIamAccess_ToProto(mapCtx, in.GcpIAMAccess); oneof != nil {
		out.AccessType = &pb.PrivilegedAccess_GcpIamAccess_{GcpIamAccess: oneof}
	}
	return out
}
func PrivilegedAccessManagerEntitlementSpec_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement) *krm.PrivilegedAccessManagerEntitlementSpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivilegedAccessManagerEntitlementSpec{}
	out.EligibleUsers = direct.SliceOfPointers_FromProto(mapCtx, in.EligibleUsers, AccessControlEntry_FromProto)
	out.ApprovalWorkflow = ApprovalWorkflow_FromProto(mapCtx, in.GetApprovalWorkflow())
	out.PrivilegedAccess = PrivilegedAccess_FromProto(mapCtx, in.GetPrivilegedAccess())
	out.MaxRequestDuration = direct.StringDuration_FromProto(mapCtx, in.GetMaxRequestDuration())
	out.RequesterJustificationConfig = RequesterJustificationConfig_FromProto(mapCtx, in.GetRequesterJustificationConfig())
	out.AdditionalNotificationTargets = AdditionalNotificationTargets_FromProto(mapCtx, in.GetAdditionalNotificationTargets())
	return out
}
func PrivilegedAccessManagerEntitlementSpec_ToProto(mapCtx *direct.MapContext, in *krm.PrivilegedAccessManagerEntitlementSpec) *pb.Entitlement {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement{}
	// MISSING: Name
	out.EligibleUsers = direct.SliceOfPointers_ToProto(mapCtx, in.EligibleUsers, AccessControlEntry_ToProto)
	out.ApprovalWorkflow = ApprovalWorkflow_ToProto(mapCtx, in.ApprovalWorkflow)
	out.PrivilegedAccess = PrivilegedAccess_ToProto(mapCtx, in.PrivilegedAccess)
	out.MaxRequestDuration = direct.StringDuration_ToProto(mapCtx, in.MaxRequestDuration)
	out.RequesterJustificationConfig = RequesterJustificationConfig_ToProto(mapCtx, in.RequesterJustificationConfig)
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
func RequesterJustificationConfig_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement_RequesterJustificationConfig) *krm.RequesterJustificationConfig {
	if in == nil {
		return nil
	}
	out := &krm.RequesterJustificationConfig{}
	if in.GetNotMandatory() != nil {
		out.NotMandatory = &runtime.RawExtension{Raw: []byte("{}")}
	} else if in.GetUnstructured() != nil {
		out.Unstructured = &runtime.RawExtension{Raw: []byte("{}")}
	}
	return out
}
func RequesterJustificationConfig_ToProto(mapCtx *direct.MapContext, in *krm.RequesterJustificationConfig) *pb.Entitlement_RequesterJustificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement_RequesterJustificationConfig{}
	if in.NotMandatory != nil {
		out.JustificationType = &pb.Entitlement_RequesterJustificationConfig_NotMandatory_{
			NotMandatory: &pb.Entitlement_RequesterJustificationConfig_NotMandatory{},
		}
	} else if in.Unstructured != nil {
		out.JustificationType = &pb.Entitlement_RequesterJustificationConfig_Unstructured_{
			Unstructured: &pb.Entitlement_RequesterJustificationConfig_Unstructured{},
		}
	}
	return out
}
func Step_FromProto(mapCtx *direct.MapContext, in *pb.ManualApprovals_Step) *krm.Step {
	if in == nil {
		return nil
	}
	out := &krm.Step{}
	out.Approvers = direct.SliceOfPointers_FromProto(mapCtx, in.Approvers, AccessControlEntry_FromProto)
	out.ApprovalsNeeded = direct.LazyPtr(in.GetApprovalsNeeded())
	out.ApproverEmailRecipients = in.ApproverEmailRecipients
	return out
}
func Step_ToProto(mapCtx *direct.MapContext, in *krm.Step) *pb.ManualApprovals_Step {
	if in == nil {
		return nil
	}
	out := &pb.ManualApprovals_Step{}
	out.Approvers = direct.SliceOfPointers_ToProto(mapCtx, in.Approvers, AccessControlEntry_ToProto)
	out.ApprovalsNeeded = direct.ValueOf(in.ApprovalsNeeded)
	out.ApproverEmailRecipients = in.ApproverEmailRecipients
	return out
}
