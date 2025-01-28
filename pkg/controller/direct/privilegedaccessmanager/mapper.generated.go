// Copyright 2025 Google LLC
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
	pb "cloud.google.com/go/privilegedaccessmanager/apiv1/privilegedaccessmanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privilegedaccessmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AccessControlEntry_FromProto(mapCtx *direct.MapContext, in *pb.AccessControlEntry) *krm.AccessControlEntry {
	if in == nil {
		return nil
	}
	out := &krm.AccessControlEntry{}
	out.Principals = in.Principals
	return out
}
func AccessControlEntry_ToProto(mapCtx *direct.MapContext, in *krm.AccessControlEntry) *pb.AccessControlEntry {
	if in == nil {
		return nil
	}
	out := &pb.AccessControlEntry{}
	out.Principals = in.Principals
	return out
}
func AdditionalNotificationTargets_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement_AdditionalNotificationTargets) *krm.AdditionalNotificationTargets {
	if in == nil {
		return nil
	}
	out := &krm.AdditionalNotificationTargets{}
	out.AdminEmailRecipients = in.AdminEmailRecipients
	out.RequesterEmailRecipients = in.RequesterEmailRecipients
	return out
}
func AdditionalNotificationTargets_ToProto(mapCtx *direct.MapContext, in *krm.AdditionalNotificationTargets) *pb.Entitlement_AdditionalNotificationTargets {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement_AdditionalNotificationTargets{}
	out.AdminEmailRecipients = in.AdminEmailRecipients
	out.RequesterEmailRecipients = in.RequesterEmailRecipients
	return out
}
func ApprovalWorkflow_FromProto(mapCtx *direct.MapContext, in *pb.ApprovalWorkflow) *krm.ApprovalWorkflow {
	if in == nil {
		return nil
	}
	out := &krm.ApprovalWorkflow{}
	out.ManualApprovals = ManualApprovals_FromProto(mapCtx, in.GetManualApprovals())
	return out
}
func ApprovalWorkflow_ToProto(mapCtx *direct.MapContext, in *krm.ApprovalWorkflow) *pb.ApprovalWorkflow {
	if in == nil {
		return nil
	}
	out := &pb.ApprovalWorkflow{}
	if oneof := ManualApprovals_ToProto(mapCtx, in.ManualApprovals); oneof != nil {
		out.ApprovalWorkflow = &pb.ApprovalWorkflow_ManualApprovals{ManualApprovals: oneof}
	}
	return out
}
func PrivilegedAccessManagerEntitlementObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement) *krm.PrivilegedAccessManagerEntitlementObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivilegedAccessManagerEntitlementObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func PrivilegedAccessManagerEntitlementObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivilegedAccessManagerEntitlementObservedState) *pb.Entitlement {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Entitlement_State](mapCtx, in.State)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func RoleBinding_FromProto(mapCtx *direct.MapContext, in *pb.PrivilegedAccess_GcpIamAccess_RoleBinding) *krm.RoleBinding {
	if in == nil {
		return nil
	}
	out := &krm.RoleBinding{}
	out.Role = direct.LazyPtr(in.GetRole())
	out.ConditionExpression = direct.LazyPtr(in.GetConditionExpression())
	return out
}
func RoleBinding_ToProto(mapCtx *direct.MapContext, in *krm.RoleBinding) *pb.PrivilegedAccess_GcpIamAccess_RoleBinding {
	if in == nil {
		return nil
	}
	out := &pb.PrivilegedAccess_GcpIamAccess_RoleBinding{}
	out.Role = direct.ValueOf(in.Role)
	out.ConditionExpression = direct.ValueOf(in.ConditionExpression)
	return out
}
