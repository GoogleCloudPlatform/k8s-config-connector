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

package cloudcontrolspartner

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudcontrolspartner/apiv1/cloudcontrolspartnerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudcontrolspartner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CloudcontrolspartnerCustomerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Customer) *krm.CloudcontrolspartnerCustomerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerCustomerObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomerOnboardingState
	// MISSING: IsOnboarded
	return out
}
func CloudcontrolspartnerCustomerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerCustomerObservedState) *pb.Customer {
	if in == nil {
		return nil
	}
	out := &pb.Customer{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomerOnboardingState
	// MISSING: IsOnboarded
	return out
}
func CloudcontrolspartnerCustomerSpec_FromProto(mapCtx *direct.MapContext, in *pb.Customer) *krm.CloudcontrolspartnerCustomerSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerCustomerSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomerOnboardingState
	// MISSING: IsOnboarded
	return out
}
func CloudcontrolspartnerCustomerSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerCustomerSpec) *pb.Customer {
	if in == nil {
		return nil
	}
	out := &pb.Customer{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CustomerOnboardingState
	// MISSING: IsOnboarded
	return out
}
func CloudcontrolspartnerEkmConnectionsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnections) *krm.CloudcontrolspartnerEkmConnectionsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerEkmConnectionsObservedState{}
	// MISSING: Name
	// MISSING: EkmConnections
	return out
}
func CloudcontrolspartnerEkmConnectionsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerEkmConnectionsObservedState) *pb.EkmConnections {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnections{}
	// MISSING: Name
	// MISSING: EkmConnections
	return out
}
func CloudcontrolspartnerEkmConnectionsSpec_FromProto(mapCtx *direct.MapContext, in *pb.EkmConnections) *krm.CloudcontrolspartnerEkmConnectionsSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerEkmConnectionsSpec{}
	// MISSING: Name
	// MISSING: EkmConnections
	return out
}
func CloudcontrolspartnerEkmConnectionsSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerEkmConnectionsSpec) *pb.EkmConnections {
	if in == nil {
		return nil
	}
	out := &pb.EkmConnections{}
	// MISSING: Name
	// MISSING: EkmConnections
	return out
}
func CloudcontrolspartnerPartnerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Partner) *krm.CloudcontrolspartnerPartnerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerPartnerObservedState{}
	// MISSING: Name
	// MISSING: Skus
	// MISSING: EkmSolutions
	// MISSING: OperatedCloudRegions
	// MISSING: PartnerProjectID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func CloudcontrolspartnerPartnerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerPartnerObservedState) *pb.Partner {
	if in == nil {
		return nil
	}
	out := &pb.Partner{}
	// MISSING: Name
	// MISSING: Skus
	// MISSING: EkmSolutions
	// MISSING: OperatedCloudRegions
	// MISSING: PartnerProjectID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func CloudcontrolspartnerPartnerPermissionsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PartnerPermissions) *krm.CloudcontrolspartnerPartnerPermissionsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerPartnerPermissionsObservedState{}
	// MISSING: Name
	// MISSING: PartnerPermissions
	return out
}
func CloudcontrolspartnerPartnerPermissionsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerPartnerPermissionsObservedState) *pb.PartnerPermissions {
	if in == nil {
		return nil
	}
	out := &pb.PartnerPermissions{}
	// MISSING: Name
	// MISSING: PartnerPermissions
	return out
}
func CloudcontrolspartnerPartnerPermissionsSpec_FromProto(mapCtx *direct.MapContext, in *pb.PartnerPermissions) *krm.CloudcontrolspartnerPartnerPermissionsSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerPartnerPermissionsSpec{}
	// MISSING: Name
	// MISSING: PartnerPermissions
	return out
}
func CloudcontrolspartnerPartnerPermissionsSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerPartnerPermissionsSpec) *pb.PartnerPermissions {
	if in == nil {
		return nil
	}
	out := &pb.PartnerPermissions{}
	// MISSING: Name
	// MISSING: PartnerPermissions
	return out
}
func CloudcontrolspartnerPartnerSpec_FromProto(mapCtx *direct.MapContext, in *pb.Partner) *krm.CloudcontrolspartnerPartnerSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerPartnerSpec{}
	// MISSING: Name
	// MISSING: Skus
	// MISSING: EkmSolutions
	// MISSING: OperatedCloudRegions
	// MISSING: PartnerProjectID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func CloudcontrolspartnerPartnerSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerPartnerSpec) *pb.Partner {
	if in == nil {
		return nil
	}
	out := &pb.Partner{}
	// MISSING: Name
	// MISSING: Skus
	// MISSING: EkmSolutions
	// MISSING: OperatedCloudRegions
	// MISSING: PartnerProjectID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func CloudcontrolspartnerViolationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Violation) *krm.CloudcontrolspartnerViolationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerViolationObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: BeginTime
	// MISSING: UpdateTime
	// MISSING: ResolveTime
	// MISSING: Category
	// MISSING: State
	// MISSING: NonCompliantOrgPolicy
	// MISSING: FolderID
	// MISSING: Remediation
	return out
}
func CloudcontrolspartnerViolationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerViolationObservedState) *pb.Violation {
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
	// MISSING: NonCompliantOrgPolicy
	// MISSING: FolderID
	// MISSING: Remediation
	return out
}
func CloudcontrolspartnerViolationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Violation) *krm.CloudcontrolspartnerViolationSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerViolationSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: BeginTime
	// MISSING: UpdateTime
	// MISSING: ResolveTime
	// MISSING: Category
	// MISSING: State
	// MISSING: NonCompliantOrgPolicy
	// MISSING: FolderID
	// MISSING: Remediation
	return out
}
func CloudcontrolspartnerViolationSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerViolationSpec) *pb.Violation {
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
	// MISSING: NonCompliantOrgPolicy
	// MISSING: FolderID
	// MISSING: Remediation
	return out
}
func CloudcontrolspartnerWorkloadObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.CloudcontrolspartnerWorkloadObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerWorkloadObservedState{}
	// MISSING: Name
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	// MISSING: WorkloadOnboardingState
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func CloudcontrolspartnerWorkloadObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerWorkloadObservedState) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	// MISSING: WorkloadOnboardingState
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func CloudcontrolspartnerWorkloadSpec_FromProto(mapCtx *direct.MapContext, in *pb.Workload) *krm.CloudcontrolspartnerWorkloadSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudcontrolspartnerWorkloadSpec{}
	// MISSING: Name
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	// MISSING: WorkloadOnboardingState
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func CloudcontrolspartnerWorkloadSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudcontrolspartnerWorkloadSpec) *pb.Workload {
	if in == nil {
		return nil
	}
	out := &pb.Workload{}
	// MISSING: Name
	// MISSING: FolderID
	// MISSING: CreateTime
	// MISSING: Folder
	// MISSING: WorkloadOnboardingState
	// MISSING: IsOnboarded
	// MISSING: KeyManagementProjectID
	// MISSING: Location
	// MISSING: Partner
	return out
}
func Violation_FromProto(mapCtx *direct.MapContext, in *pb.Violation) *krm.Violation {
	if in == nil {
		return nil
	}
	out := &krm.Violation{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	// MISSING: BeginTime
	// MISSING: UpdateTime
	// MISSING: ResolveTime
	// MISSING: Category
	// MISSING: State
	// MISSING: NonCompliantOrgPolicy
	out.FolderID = direct.LazyPtr(in.GetFolderId())
	// MISSING: Remediation
	return out
}
func Violation_ToProto(mapCtx *direct.MapContext, in *krm.Violation) *pb.Violation {
	if in == nil {
		return nil
	}
	out := &pb.Violation{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	// MISSING: BeginTime
	// MISSING: UpdateTime
	// MISSING: ResolveTime
	// MISSING: Category
	// MISSING: State
	// MISSING: NonCompliantOrgPolicy
	out.FolderId = direct.ValueOf(in.FolderID)
	// MISSING: Remediation
	return out
}
func ViolationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Violation) *krm.ViolationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ViolationObservedState{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.BeginTime = direct.StringTimestamp_FromProto(mapCtx, in.GetBeginTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ResolveTime = direct.StringTimestamp_FromProto(mapCtx, in.GetResolveTime())
	out.Category = direct.LazyPtr(in.GetCategory())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.NonCompliantOrgPolicy = direct.LazyPtr(in.GetNonCompliantOrgPolicy())
	// MISSING: FolderID
	out.Remediation = Violation_Remediation_FromProto(mapCtx, in.GetRemediation())
	return out
}
func ViolationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ViolationObservedState) *pb.Violation {
	if in == nil {
		return nil
	}
	out := &pb.Violation{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.BeginTime = direct.StringTimestamp_ToProto(mapCtx, in.BeginTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ResolveTime = direct.StringTimestamp_ToProto(mapCtx, in.ResolveTime)
	out.Category = direct.ValueOf(in.Category)
	out.State = direct.Enum_ToProto[pb.Violation_State](mapCtx, in.State)
	out.NonCompliantOrgPolicy = direct.ValueOf(in.NonCompliantOrgPolicy)
	// MISSING: FolderID
	out.Remediation = Violation_Remediation_ToProto(mapCtx, in.Remediation)
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
