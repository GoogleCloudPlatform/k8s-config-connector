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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privilegedaccessmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privilegedaccessmanager/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/privilegedaccessmanager/apiv1/privilegedaccessmanagerpb"
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
func Grant_FromProto(mapCtx *direct.MapContext, in *pb.Grant) *krm.Grant {
	if in == nil {
		return nil
	}
	out := &krm.Grant{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Requester
	out.RequestedDuration = direct.StringDuration_FromProto(mapCtx, in.GetRequestedDuration())
	out.Justification = Justification_FromProto(mapCtx, in.GetJustification())
	// MISSING: State
	// MISSING: Timeline
	// MISSING: PrivilegedAccess
	// MISSING: AuditTrail
	out.AdditionalEmailRecipients = in.AdditionalEmailRecipients
	// MISSING: ExternallyModified
	return out
}
func Grant_ToProto(mapCtx *direct.MapContext, in *krm.Grant) *pb.Grant {
	if in == nil {
		return nil
	}
	out := &pb.Grant{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Requester
	out.RequestedDuration = direct.StringDuration_ToProto(mapCtx, in.RequestedDuration)
	out.Justification = Justification_ToProto(mapCtx, in.Justification)
	// MISSING: State
	// MISSING: Timeline
	// MISSING: PrivilegedAccess
	// MISSING: AuditTrail
	out.AdditionalEmailRecipients = in.AdditionalEmailRecipients
	// MISSING: ExternallyModified
	return out
}
func GrantObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Grant) *krm.GrantObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GrantObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Requester = direct.LazyPtr(in.GetRequester())
	// MISSING: RequestedDuration
	// MISSING: Justification
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Timeline = Grant_Timeline_FromProto(mapCtx, in.GetTimeline())
	out.PrivilegedAccess = PrivilegedAccess_FromProto(mapCtx, in.GetPrivilegedAccess())
	out.AuditTrail = Grant_AuditTrail_FromProto(mapCtx, in.GetAuditTrail())
	// MISSING: AdditionalEmailRecipients
	out.ExternallyModified = direct.LazyPtr(in.GetExternallyModified())
	return out
}
func GrantObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GrantObservedState) *pb.Grant {
	if in == nil {
		return nil
	}
	out := &pb.Grant{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Requester = direct.ValueOf(in.Requester)
	// MISSING: RequestedDuration
	// MISSING: Justification
	out.State = direct.Enum_ToProto[pb.Grant_State](mapCtx, in.State)
	out.Timeline = Grant_Timeline_ToProto(mapCtx, in.Timeline)
	out.PrivilegedAccess = PrivilegedAccess_ToProto(mapCtx, in.PrivilegedAccess)
	out.AuditTrail = Grant_AuditTrail_ToProto(mapCtx, in.AuditTrail)
	// MISSING: AdditionalEmailRecipients
	out.ExternallyModified = direct.ValueOf(in.ExternallyModified)
	return out
}
func Grant_AuditTrail_FromProto(mapCtx *direct.MapContext, in *pb.Grant_AuditTrail) *krm.Grant_AuditTrail {
	if in == nil {
		return nil
	}
	out := &krm.Grant_AuditTrail{}
	// MISSING: AccessGrantTime
	// MISSING: AccessRemoveTime
	return out
}
func Grant_AuditTrail_ToProto(mapCtx *direct.MapContext, in *krm.Grant_AuditTrail) *pb.Grant_AuditTrail {
	if in == nil {
		return nil
	}
	out := &pb.Grant_AuditTrail{}
	// MISSING: AccessGrantTime
	// MISSING: AccessRemoveTime
	return out
}
func Grant_AuditTrailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Grant_AuditTrail) *krm.Grant_AuditTrailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Grant_AuditTrailObservedState{}
	out.AccessGrantTime = direct.StringTimestamp_FromProto(mapCtx, in.GetAccessGrantTime())
	out.AccessRemoveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetAccessRemoveTime())
	return out
}
func Grant_AuditTrailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Grant_AuditTrailObservedState) *pb.Grant_AuditTrail {
	if in == nil {
		return nil
	}
	out := &pb.Grant_AuditTrail{}
	out.AccessGrantTime = direct.StringTimestamp_ToProto(mapCtx, in.AccessGrantTime)
	out.AccessRemoveTime = direct.StringTimestamp_ToProto(mapCtx, in.AccessRemoveTime)
	return out
}
func Grant_Timeline_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline) *krm.Grant_Timeline {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline{}
	// MISSING: Events
	return out
}
func Grant_Timeline_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline) *pb.Grant_Timeline {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline{}
	// MISSING: Events
	return out
}
func Grant_TimelineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline) *krm.Grant_TimelineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Grant_TimelineObservedState{}
	out.Events = direct.Slice_FromProto(mapCtx, in.Events, Grant_Timeline_Event_FromProto)
	return out
}
func Grant_TimelineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Grant_TimelineObservedState) *pb.Grant_Timeline {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline{}
	out.Events = direct.Slice_ToProto(mapCtx, in.Events, Grant_Timeline_Event_ToProto)
	return out
}
func Grant_Timeline_Event_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event) *krm.Grant_Timeline_Event {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event{}
	out.Requested = Grant_Timeline_Event_Requested_FromProto(mapCtx, in.GetRequested())
	out.Approved = Grant_Timeline_Event_Approved_FromProto(mapCtx, in.GetApproved())
	out.Denied = Grant_Timeline_Event_Denied_FromProto(mapCtx, in.GetDenied())
	out.Revoked = Grant_Timeline_Event_Revoked_FromProto(mapCtx, in.GetRevoked())
	out.Scheduled = Grant_Timeline_Event_Scheduled_FromProto(mapCtx, in.GetScheduled())
	out.Activated = Grant_Timeline_Event_Activated_FromProto(mapCtx, in.GetActivated())
	out.ActivationFailed = Grant_Timeline_Event_ActivationFailed_FromProto(mapCtx, in.GetActivationFailed())
	out.Expired = Grant_Timeline_Event_Expired_FromProto(mapCtx, in.GetExpired())
	out.Ended = Grant_Timeline_Event_Ended_FromProto(mapCtx, in.GetEnded())
	out.ExternallyModified = Grant_Timeline_Event_ExternallyModified_FromProto(mapCtx, in.GetExternallyModified())
	// MISSING: EventTime
	return out
}
func Grant_Timeline_Event_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event) *pb.Grant_Timeline_Event {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event{}
	if oneof := Grant_Timeline_Event_Requested_ToProto(mapCtx, in.Requested); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Requested_{Requested: oneof}
	}
	if oneof := Grant_Timeline_Event_Approved_ToProto(mapCtx, in.Approved); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Approved_{Approved: oneof}
	}
	if oneof := Grant_Timeline_Event_Denied_ToProto(mapCtx, in.Denied); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Denied_{Denied: oneof}
	}
	if oneof := Grant_Timeline_Event_Revoked_ToProto(mapCtx, in.Revoked); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Revoked_{Revoked: oneof}
	}
	if oneof := Grant_Timeline_Event_Scheduled_ToProto(mapCtx, in.Scheduled); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Scheduled_{Scheduled: oneof}
	}
	if oneof := Grant_Timeline_Event_Activated_ToProto(mapCtx, in.Activated); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Activated_{Activated: oneof}
	}
	if oneof := Grant_Timeline_Event_ActivationFailed_ToProto(mapCtx, in.ActivationFailed); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_ActivationFailed_{ActivationFailed: oneof}
	}
	if oneof := Grant_Timeline_Event_Expired_ToProto(mapCtx, in.Expired); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Expired_{Expired: oneof}
	}
	if oneof := Grant_Timeline_Event_Ended_ToProto(mapCtx, in.Ended); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Ended_{Ended: oneof}
	}
	if oneof := Grant_Timeline_Event_ExternallyModified_ToProto(mapCtx, in.ExternallyModified); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_ExternallyModified_{ExternallyModified: oneof}
	}
	// MISSING: EventTime
	return out
}
func Grant_Timeline_EventObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event) *krm.Grant_Timeline_EventObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_EventObservedState{}
	out.Requested = Grant_Timeline_Event_RequestedObservedState_FromProto(mapCtx, in.GetRequested())
	out.Approved = Grant_Timeline_Event_ApprovedObservedState_FromProto(mapCtx, in.GetApproved())
	out.Denied = Grant_Timeline_Event_DeniedObservedState_FromProto(mapCtx, in.GetDenied())
	out.Revoked = Grant_Timeline_Event_RevokedObservedState_FromProto(mapCtx, in.GetRevoked())
	out.Scheduled = Grant_Timeline_Event_ScheduledObservedState_FromProto(mapCtx, in.GetScheduled())
	// MISSING: Activated
	out.ActivationFailed = Grant_Timeline_Event_ActivationFailedObservedState_FromProto(mapCtx, in.GetActivationFailed())
	// MISSING: Expired
	// MISSING: Ended
	// MISSING: ExternallyModified
	out.EventTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEventTime())
	return out
}
func Grant_Timeline_EventObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_EventObservedState) *pb.Grant_Timeline_Event {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event{}
	if oneof := Grant_Timeline_Event_RequestedObservedState_ToProto(mapCtx, in.Requested); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Requested_{Requested: oneof}
	}
	if oneof := Grant_Timeline_Event_ApprovedObservedState_ToProto(mapCtx, in.Approved); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Approved_{Approved: oneof}
	}
	if oneof := Grant_Timeline_Event_DeniedObservedState_ToProto(mapCtx, in.Denied); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Denied_{Denied: oneof}
	}
	if oneof := Grant_Timeline_Event_RevokedObservedState_ToProto(mapCtx, in.Revoked); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Revoked_{Revoked: oneof}
	}
	if oneof := Grant_Timeline_Event_ScheduledObservedState_ToProto(mapCtx, in.Scheduled); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_Scheduled_{Scheduled: oneof}
	}
	// MISSING: Activated
	if oneof := Grant_Timeline_Event_ActivationFailedObservedState_ToProto(mapCtx, in.ActivationFailed); oneof != nil {
		out.Event = &pb.Grant_Timeline_Event_ActivationFailed_{ActivationFailed: oneof}
	}
	// MISSING: Expired
	// MISSING: Ended
	// MISSING: ExternallyModified
	out.EventTime = direct.StringTimestamp_ToProto(mapCtx, in.EventTime)
	return out
}
func Grant_Timeline_Event_Activated_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Activated) *krm.Grant_Timeline_Event_Activated {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_Activated{}
	return out
}
func Grant_Timeline_Event_Activated_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_Activated) *pb.Grant_Timeline_Event_Activated {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Activated{}
	return out
}
func Grant_Timeline_Event_ActivationFailed_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_ActivationFailed) *krm.Grant_Timeline_Event_ActivationFailed {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_ActivationFailed{}
	// MISSING: Error
	return out
}
func Grant_Timeline_Event_ActivationFailed_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_ActivationFailed) *pb.Grant_Timeline_Event_ActivationFailed {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_ActivationFailed{}
	// MISSING: Error
	return out
}
func Grant_Timeline_Event_ActivationFailedObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_ActivationFailed) *krm.Grant_Timeline_Event_ActivationFailedObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_ActivationFailedObservedState{}
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}
func Grant_Timeline_Event_ActivationFailedObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_ActivationFailedObservedState) *pb.Grant_Timeline_Event_ActivationFailed {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_ActivationFailed{}
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}
func Grant_Timeline_Event_Approved_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Approved) *krm.Grant_Timeline_Event_Approved {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_Approved{}
	// MISSING: Reason
	// MISSING: Actor
	return out
}
func Grant_Timeline_Event_Approved_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_Approved) *pb.Grant_Timeline_Event_Approved {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Approved{}
	// MISSING: Reason
	// MISSING: Actor
	return out
}
func Grant_Timeline_Event_ApprovedObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Approved) *krm.Grant_Timeline_Event_ApprovedObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_ApprovedObservedState{}
	out.Reason = direct.LazyPtr(in.GetReason())
	out.Actor = direct.LazyPtr(in.GetActor())
	return out
}
func Grant_Timeline_Event_ApprovedObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_ApprovedObservedState) *pb.Grant_Timeline_Event_Approved {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Approved{}
	out.Reason = direct.ValueOf(in.Reason)
	out.Actor = direct.ValueOf(in.Actor)
	return out
}
func Grant_Timeline_Event_Denied_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Denied) *krm.Grant_Timeline_Event_Denied {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_Denied{}
	// MISSING: Reason
	// MISSING: Actor
	return out
}
func Grant_Timeline_Event_Denied_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_Denied) *pb.Grant_Timeline_Event_Denied {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Denied{}
	// MISSING: Reason
	// MISSING: Actor
	return out
}
func Grant_Timeline_Event_DeniedObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Denied) *krm.Grant_Timeline_Event_DeniedObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_DeniedObservedState{}
	out.Reason = direct.LazyPtr(in.GetReason())
	out.Actor = direct.LazyPtr(in.GetActor())
	return out
}
func Grant_Timeline_Event_DeniedObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_DeniedObservedState) *pb.Grant_Timeline_Event_Denied {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Denied{}
	out.Reason = direct.ValueOf(in.Reason)
	out.Actor = direct.ValueOf(in.Actor)
	return out
}
func Grant_Timeline_Event_Ended_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Ended) *krm.Grant_Timeline_Event_Ended {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_Ended{}
	return out
}
func Grant_Timeline_Event_Ended_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_Ended) *pb.Grant_Timeline_Event_Ended {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Ended{}
	return out
}
func Grant_Timeline_Event_Expired_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Expired) *krm.Grant_Timeline_Event_Expired {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_Expired{}
	return out
}
func Grant_Timeline_Event_Expired_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_Expired) *pb.Grant_Timeline_Event_Expired {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Expired{}
	return out
}
func Grant_Timeline_Event_ExternallyModified_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_ExternallyModified) *krm.Grant_Timeline_Event_ExternallyModified {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_ExternallyModified{}
	return out
}
func Grant_Timeline_Event_ExternallyModified_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_ExternallyModified) *pb.Grant_Timeline_Event_ExternallyModified {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_ExternallyModified{}
	return out
}
func Grant_Timeline_Event_Requested_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Requested) *krm.Grant_Timeline_Event_Requested {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_Requested{}
	// MISSING: ExpireTime
	return out
}
func Grant_Timeline_Event_Requested_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_Requested) *pb.Grant_Timeline_Event_Requested {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Requested{}
	// MISSING: ExpireTime
	return out
}
func Grant_Timeline_Event_RequestedObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Requested) *krm.Grant_Timeline_Event_RequestedObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_RequestedObservedState{}
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}
func Grant_Timeline_Event_RequestedObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_RequestedObservedState) *pb.Grant_Timeline_Event_Requested {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Requested{}
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}
func Grant_Timeline_Event_Revoked_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Revoked) *krm.Grant_Timeline_Event_Revoked {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_Revoked{}
	// MISSING: Reason
	// MISSING: Actor
	return out
}
func Grant_Timeline_Event_Revoked_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_Revoked) *pb.Grant_Timeline_Event_Revoked {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Revoked{}
	// MISSING: Reason
	// MISSING: Actor
	return out
}
func Grant_Timeline_Event_RevokedObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Revoked) *krm.Grant_Timeline_Event_RevokedObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_RevokedObservedState{}
	out.Reason = direct.LazyPtr(in.GetReason())
	out.Actor = direct.LazyPtr(in.GetActor())
	return out
}
func Grant_Timeline_Event_RevokedObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_RevokedObservedState) *pb.Grant_Timeline_Event_Revoked {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Revoked{}
	out.Reason = direct.ValueOf(in.Reason)
	out.Actor = direct.ValueOf(in.Actor)
	return out
}
func Grant_Timeline_Event_Scheduled_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Scheduled) *krm.Grant_Timeline_Event_Scheduled {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_Scheduled{}
	// MISSING: ScheduledActivationTime
	return out
}
func Grant_Timeline_Event_Scheduled_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_Scheduled) *pb.Grant_Timeline_Event_Scheduled {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Scheduled{}
	// MISSING: ScheduledActivationTime
	return out
}
func Grant_Timeline_Event_ScheduledObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Grant_Timeline_Event_Scheduled) *krm.Grant_Timeline_Event_ScheduledObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Grant_Timeline_Event_ScheduledObservedState{}
	out.ScheduledActivationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduledActivationTime())
	return out
}
func Grant_Timeline_Event_ScheduledObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Grant_Timeline_Event_ScheduledObservedState) *pb.Grant_Timeline_Event_Scheduled {
	if in == nil {
		return nil
	}
	out := &pb.Grant_Timeline_Event_Scheduled{}
	out.ScheduledActivationTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduledActivationTime)
	return out
}
func Justification_FromProto(mapCtx *direct.MapContext, in *pb.Justification) *krm.Justification {
	if in == nil {
		return nil
	}
	out := &krm.Justification{}
	out.UnstructuredJustification = direct.LazyPtr(in.GetUnstructuredJustification())
	return out
}
func Justification_ToProto(mapCtx *direct.MapContext, in *krm.Justification) *pb.Justification {
	if in == nil {
		return nil
	}
	out := &pb.Justification{}
	if oneof := Justification_UnstructuredJustification_ToProto(mapCtx, in.UnstructuredJustification); oneof != nil {
		out.Justification = oneof
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
func PrivilegedAccess_GcpIamAccess_FromProto(mapCtx *direct.MapContext, in *pb.PrivilegedAccess_GcpIamAccess) *krm.PrivilegedAccess_GcpIamAccess {
	if in == nil {
		return nil
	}
	out := &krm.PrivilegedAccess_GcpIamAccess{}
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.Resource = direct.LazyPtr(in.GetResource())
	out.RoleBindings = direct.Slice_FromProto(mapCtx, in.RoleBindings, PrivilegedAccess_GcpIamAccess_RoleBinding_FromProto)
	return out
}
func PrivilegedAccess_GcpIamAccess_ToProto(mapCtx *direct.MapContext, in *krm.PrivilegedAccess_GcpIamAccess) *pb.PrivilegedAccess_GcpIamAccess {
	if in == nil {
		return nil
	}
	out := &pb.PrivilegedAccess_GcpIamAccess{}
	out.ResourceType = direct.ValueOf(in.ResourceType)
	out.Resource = direct.ValueOf(in.Resource)
	out.RoleBindings = direct.Slice_ToProto(mapCtx, in.RoleBindings, PrivilegedAccess_GcpIamAccess_RoleBinding_ToProto)
	return out
}
func PrivilegedAccess_GcpIamAccess_RoleBinding_FromProto(mapCtx *direct.MapContext, in *pb.PrivilegedAccess_GcpIamAccess_RoleBinding) *krm.PrivilegedAccess_GcpIamAccess_RoleBinding {
	if in == nil {
		return nil
	}
	out := &krm.PrivilegedAccess_GcpIamAccess_RoleBinding{}
	out.Role = direct.LazyPtr(in.GetRole())
	out.ConditionExpression = direct.LazyPtr(in.GetConditionExpression())
	return out
}
func PrivilegedAccess_GcpIamAccess_RoleBinding_ToProto(mapCtx *direct.MapContext, in *krm.PrivilegedAccess_GcpIamAccess_RoleBinding) *pb.PrivilegedAccess_GcpIamAccess_RoleBinding {
	if in == nil {
		return nil
	}
	out := &pb.PrivilegedAccess_GcpIamAccess_RoleBinding{}
	out.Role = direct.ValueOf(in.Role)
	out.ConditionExpression = direct.ValueOf(in.ConditionExpression)
	return out
}
func PrivilegedaccessmanagerGrantObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Grant) *krm.PrivilegedaccessmanagerGrantObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivilegedaccessmanagerGrantObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Requester
	// MISSING: RequestedDuration
	// MISSING: Justification
	// MISSING: State
	// MISSING: Timeline
	// MISSING: PrivilegedAccess
	// MISSING: AuditTrail
	// MISSING: AdditionalEmailRecipients
	// MISSING: ExternallyModified
	return out
}
func PrivilegedaccessmanagerGrantObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivilegedaccessmanagerGrantObservedState) *pb.Grant {
	if in == nil {
		return nil
	}
	out := &pb.Grant{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Requester
	// MISSING: RequestedDuration
	// MISSING: Justification
	// MISSING: State
	// MISSING: Timeline
	// MISSING: PrivilegedAccess
	// MISSING: AuditTrail
	// MISSING: AdditionalEmailRecipients
	// MISSING: ExternallyModified
	return out
}
func PrivilegedaccessmanagerGrantSpec_FromProto(mapCtx *direct.MapContext, in *pb.Grant) *krm.PrivilegedaccessmanagerGrantSpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivilegedaccessmanagerGrantSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Requester
	// MISSING: RequestedDuration
	// MISSING: Justification
	// MISSING: State
	// MISSING: Timeline
	// MISSING: PrivilegedAccess
	// MISSING: AuditTrail
	// MISSING: AdditionalEmailRecipients
	// MISSING: ExternallyModified
	return out
}
func PrivilegedaccessmanagerGrantSpec_ToProto(mapCtx *direct.MapContext, in *krm.PrivilegedaccessmanagerGrantSpec) *pb.Grant {
	if in == nil {
		return nil
	}
	out := &pb.Grant{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Requester
	// MISSING: RequestedDuration
	// MISSING: Justification
	// MISSING: State
	// MISSING: Timeline
	// MISSING: PrivilegedAccess
	// MISSING: AuditTrail
	// MISSING: AdditionalEmailRecipients
	// MISSING: ExternallyModified
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
