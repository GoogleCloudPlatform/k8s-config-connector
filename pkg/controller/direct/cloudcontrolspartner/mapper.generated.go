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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudcontrolspartner/apiv1beta/cloudcontrolspartnerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudcontrolspartner/v1alpha1"
)
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
