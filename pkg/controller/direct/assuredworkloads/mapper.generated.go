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

package assuredworkloads

import (
	pb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/assuredworkloads/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func AssuredworkloadsViolationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Violation) *krm.AssuredworkloadsViolationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssuredworkloadsViolationObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: BeginTime
	// MISSING: UpdateTime
	// MISSING: ResolveTime
	// MISSING: Category
	// MISSING: State
	// MISSING: OrgPolicyConstraint
	// MISSING: AuditLogLink
	// MISSING: NonCompliantOrgPolicy
	// MISSING: Remediation
	// MISSING: Acknowledged
	// MISSING: AcknowledgementTime
	// MISSING: ExceptionAuditLogLink
	return out
}
func AssuredworkloadsViolationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssuredworkloadsViolationObservedState) *pb.Violation {
	if in == nil {
		return nil
	}
	out := &pb.Violation{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: BeginTime
	// MISSING: UpdateTime
	// MISSING: ResolveTime
	// MISSING: Category
	// MISSING: State
	// MISSING: OrgPolicyConstraint
	// MISSING: AuditLogLink
	// MISSING: NonCompliantOrgPolicy
	// MISSING: Remediation
	// MISSING: Acknowledged
	// MISSING: AcknowledgementTime
	// MISSING: ExceptionAuditLogLink
	return out
}
func AssuredworkloadsViolationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Violation) *krm.AssuredworkloadsViolationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AssuredworkloadsViolationSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: BeginTime
	// MISSING: UpdateTime
	// MISSING: ResolveTime
	// MISSING: Category
	// MISSING: State
	// MISSING: OrgPolicyConstraint
	// MISSING: AuditLogLink
	// MISSING: NonCompliantOrgPolicy
	// MISSING: Remediation
	// MISSING: Acknowledged
	// MISSING: AcknowledgementTime
	// MISSING: ExceptionAuditLogLink
	return out
}
func AssuredworkloadsViolationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AssuredworkloadsViolationSpec) *pb.Violation {
	if in == nil {
		return nil
	}
	out := &pb.Violation{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: BeginTime
	// MISSING: UpdateTime
	// MISSING: ResolveTime
	// MISSING: Category
	// MISSING: State
	// MISSING: OrgPolicyConstraint
	// MISSING: AuditLogLink
	// MISSING: NonCompliantOrgPolicy
	// MISSING: Remediation
	// MISSING: Acknowledged
	// MISSING: AcknowledgementTime
	// MISSING: ExceptionAuditLogLink
	return out
}
func AssuredworkloadsWorkloadObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.AssuredworkloadsWorkloadObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssuredworkloadsWorkloadObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Resources
	// MISSING: ComplianceRegime
	// MISSING: CreateTime
	// MISSING: BillingAccount
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: ProvisionedResourcesParent
	// MISSING: KMSSettings
	// MISSING: ResourceSettings
	// MISSING: KajEnrollmentState
	// MISSING: EnableSovereignControls
	// MISSING: SaaEnrollmentResponse
	// MISSING: CompliantButDisallowedServices
	// MISSING: Partner
	return out
}
func AssuredworkloadsWorkloadObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssuredworkloadsWorkloadObservedState) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Resources
	// MISSING: ComplianceRegime
	// MISSING: CreateTime
	// MISSING: BillingAccount
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: ProvisionedResourcesParent
	// MISSING: KMSSettings
	// MISSING: ResourceSettings
	// MISSING: KajEnrollmentState
	// MISSING: EnableSovereignControls
	// MISSING: SaaEnrollmentResponse
	// MISSING: CompliantButDisallowedServices
	// MISSING: Partner
	return out
}
func AssuredworkloadsWorkloadSpec_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.AssuredworkloadsWorkloadSpec {
	if in == nil {
		return nil
	}
	out := &krm.AssuredworkloadsWorkloadSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Resources
	// MISSING: ComplianceRegime
	// MISSING: CreateTime
	// MISSING: BillingAccount
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: ProvisionedResourcesParent
	// MISSING: KMSSettings
	// MISSING: ResourceSettings
	// MISSING: KajEnrollmentState
	// MISSING: EnableSovereignControls
	// MISSING: SaaEnrollmentResponse
	// MISSING: CompliantButDisallowedServices
	// MISSING: Partner
	return out
}
func AssuredworkloadsWorkloadSpec_ToProto(mapCtx *direct.MapContext, in *krm.AssuredworkloadsWorkloadSpec) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Resources
	// MISSING: ComplianceRegime
	// MISSING: CreateTime
	// MISSING: BillingAccount
	// MISSING: Etag
	// MISSING: Labels
	// MISSING: ProvisionedResourcesParent
	// MISSING: KMSSettings
	// MISSING: ResourceSettings
	// MISSING: KajEnrollmentState
	// MISSING: EnableSovereignControls
	// MISSING: SaaEnrollmentResponse
	// MISSING: CompliantButDisallowedServices
	// MISSING: Partner
	return out
}
func Violation_FromProto(mapCtx *direct.MapContext, in *pb.Violation) *krm.Violation {
	if in == nil {
		return nil
	}
	out := &krm.Violation{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: BeginTime
	// MISSING: UpdateTime
	// MISSING: ResolveTime
	// MISSING: Category
	// MISSING: State
	// MISSING: OrgPolicyConstraint
	// MISSING: AuditLogLink
	// MISSING: NonCompliantOrgPolicy
	// MISSING: Remediation
	// MISSING: Acknowledged
	out.AcknowledgementTime = direct.StringTimestamp_FromProto(mapCtx, in.GetAcknowledgementTime())
	// MISSING: ExceptionAuditLogLink
	return out
}
func Violation_ToProto(mapCtx *direct.MapContext, in *krm.Violation) *pb.Violation {
	if in == nil {
		return nil
	}
	out := &pb.Violation{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: BeginTime
	// MISSING: UpdateTime
	// MISSING: ResolveTime
	// MISSING: Category
	// MISSING: State
	// MISSING: OrgPolicyConstraint
	// MISSING: AuditLogLink
	// MISSING: NonCompliantOrgPolicy
	// MISSING: Remediation
	// MISSING: Acknowledged
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.AcknowledgementTime); oneof != nil {
		out.AcknowledgementTime = &pb.Violation_AcknowledgementTime{AcknowledgementTime: oneof}
	}
	// MISSING: ExceptionAuditLogLink
	return out
}
func ViolationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Violation) *krm.ViolationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ViolationObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.BeginTime = direct.StringTimestamp_FromProto(mapCtx, in.GetBeginTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ResolveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetResolveTime())
	out.Category = direct.LazyPtr(in.GetCategory())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.OrgPolicyConstraint = direct.LazyPtr(in.GetOrgPolicyConstraint())
	out.AuditLogLink = direct.LazyPtr(in.GetAuditLogLink())
	out.NonCompliantOrgPolicy = direct.LazyPtr(in.GetNonCompliantOrgPolicy())
	out.Remediation = Violation_Remediation_FromProto(mapCtx, in.GetRemediation())
	out.Acknowledged = direct.LazyPtr(in.GetAcknowledged())
	// MISSING: AcknowledgementTime
	out.ExceptionAuditLogLink = direct.LazyPtr(in.GetExceptionAuditLogLink())
	return out
}
func ViolationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ViolationObservedState) *pb.Violation {
	if in == nil {
		return nil
	}
	out := &pb.Violation{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.BeginTime = direct.StringTimestamp_ToProto(mapCtx, in.BeginTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ResolveTime = direct.StringTimestamp_ToProto(mapCtx, in.ResolveTime)
	out.Category = direct.ValueOf(in.Category)
	out.State = direct.Enum_ToProto[pb.Violation_State](mapCtx, in.State)
	out.OrgPolicyConstraint = direct.ValueOf(in.OrgPolicyConstraint)
	out.AuditLogLink = direct.ValueOf(in.AuditLogLink)
	out.NonCompliantOrgPolicy = direct.ValueOf(in.NonCompliantOrgPolicy)
	out.Remediation = Violation_Remediation_ToProto(mapCtx, in.Remediation)
	out.Acknowledged = direct.ValueOf(in.Acknowledged)
	// MISSING: AcknowledgementTime
	out.ExceptionAuditLogLink = direct.ValueOf(in.ExceptionAuditLogLink)
	return out
}
func Violation_Remediation_FromProto(mapCtx *direct.MapContext, in *pb.Violation_Remediation) *krm.Violation_Remediation {
	if in == nil {
		return nil
	}
	out := &krm.Violation_Remediation{}
	out.Instructions = Violation_Remediation_Instructions_FromProto(mapCtx, in.GetInstructions())
	out.CompliantValues = in.CompliantValues
	// MISSING: RemediationType
	return out
}
func Violation_Remediation_ToProto(mapCtx *direct.MapContext, in *krm.Violation_Remediation) *pb.Violation_Remediation {
	if in == nil {
		return nil
	}
	out := &pb.Violation_Remediation{}
	out.Instructions = Violation_Remediation_Instructions_ToProto(mapCtx, in.Instructions)
	out.CompliantValues = in.CompliantValues
	// MISSING: RemediationType
	return out
}
func Violation_RemediationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Violation_Remediation) *krm.Violation_RemediationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Violation_RemediationObservedState{}
	// MISSING: Instructions
	// MISSING: CompliantValues
	out.RemediationType = direct.Enum_FromProto(mapCtx, in.GetRemediationType())
	return out
}
func Violation_RemediationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Violation_RemediationObservedState) *pb.Violation_Remediation {
	if in == nil {
		return nil
	}
	out := &pb.Violation_Remediation{}
	// MISSING: Instructions
	// MISSING: CompliantValues
	out.RemediationType = direct.Enum_ToProto[pb.Violation_Remediation_RemediationType](mapCtx, in.RemediationType)
	return out
}
func Violation_Remediation_Instructions_FromProto(mapCtx *direct.MapContext, in *pb.Violation_Remediation_Instructions) *krm.Violation_Remediation_Instructions {
	if in == nil {
		return nil
	}
	out := &krm.Violation_Remediation_Instructions{}
	out.GcloudInstructions = Violation_Remediation_Instructions_Gcloud_FromProto(mapCtx, in.GetGcloudInstructions())
	out.ConsoleInstructions = Violation_Remediation_Instructions_Console_FromProto(mapCtx, in.GetConsoleInstructions())
	return out
}
func Violation_Remediation_Instructions_ToProto(mapCtx *direct.MapContext, in *krm.Violation_Remediation_Instructions) *pb.Violation_Remediation_Instructions {
	if in == nil {
		return nil
	}
	out := &pb.Violation_Remediation_Instructions{}
	out.GcloudInstructions = Violation_Remediation_Instructions_Gcloud_ToProto(mapCtx, in.GcloudInstructions)
	out.ConsoleInstructions = Violation_Remediation_Instructions_Console_ToProto(mapCtx, in.ConsoleInstructions)
	return out
}
func Violation_Remediation_Instructions_Console_FromProto(mapCtx *direct.MapContext, in *pb.Violation_Remediation_Instructions_Console) *krm.Violation_Remediation_Instructions_Console {
	if in == nil {
		return nil
	}
	out := &krm.Violation_Remediation_Instructions_Console{}
	out.ConsoleUris = in.ConsoleUris
	out.Steps = in.Steps
	out.AdditionalLinks = in.AdditionalLinks
	return out
}
func Violation_Remediation_Instructions_Console_ToProto(mapCtx *direct.MapContext, in *krm.Violation_Remediation_Instructions_Console) *pb.Violation_Remediation_Instructions_Console {
	if in == nil {
		return nil
	}
	out := &pb.Violation_Remediation_Instructions_Console{}
	out.ConsoleUris = in.ConsoleUris
	out.Steps = in.Steps
	out.AdditionalLinks = in.AdditionalLinks
	return out
}
func Violation_Remediation_Instructions_Gcloud_FromProto(mapCtx *direct.MapContext, in *pb.Violation_Remediation_Instructions_Gcloud) *krm.Violation_Remediation_Instructions_Gcloud {
	if in == nil {
		return nil
	}
	out := &krm.Violation_Remediation_Instructions_Gcloud{}
	out.GcloudCommands = in.GcloudCommands
	out.Steps = in.Steps
	out.AdditionalLinks = in.AdditionalLinks
	return out
}
func Violation_Remediation_Instructions_Gcloud_ToProto(mapCtx *direct.MapContext, in *krm.Violation_Remediation_Instructions_Gcloud) *pb.Violation_Remediation_Instructions_Gcloud {
	if in == nil {
		return nil
	}
	out := &pb.Violation_Remediation_Instructions_Gcloud{}
	out.GcloudCommands = in.GcloudCommands
	out.Steps = in.Steps
	out.AdditionalLinks = in.AdditionalLinks
	return out
}
