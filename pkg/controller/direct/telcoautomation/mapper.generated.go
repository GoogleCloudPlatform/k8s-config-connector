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

package telcoautomation

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/telcoautomation/apiv1/telcoautomationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/telcoautomation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Blueprint_FromProto(mapCtx *direct.MapContext, in *pb.Blueprint) *krm.Blueprint {
	if in == nil {
		return nil
	}
	out := &krm.Blueprint{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: RevisionID
	out.SourceBlueprint = direct.LazyPtr(in.GetSourceBlueprint())
	// MISSING: RevisionCreateTime
	// MISSING: ApprovalState
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Repository
	out.Files = direct.Slice_FromProto(mapCtx, in.Files, File_FromProto)
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourceProvider
	// MISSING: DeploymentLevel
	// MISSING: RollbackSupport
	return out
}
func Blueprint_ToProto(mapCtx *direct.MapContext, in *krm.Blueprint) *pb.Blueprint {
	if in == nil {
		return nil
	}
	out := &pb.Blueprint{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: RevisionID
	out.SourceBlueprint = direct.ValueOf(in.SourceBlueprint)
	// MISSING: RevisionCreateTime
	// MISSING: ApprovalState
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Repository
	out.Files = direct.Slice_ToProto(mapCtx, in.Files, File_ToProto)
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourceProvider
	// MISSING: DeploymentLevel
	// MISSING: RollbackSupport
	return out
}
func BlueprintObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Blueprint) *krm.BlueprintObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BlueprintObservedState{}
	// MISSING: Name
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	// MISSING: SourceBlueprint
	out.RevisionCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevisionCreateTime())
	out.ApprovalState = direct.Enum_FromProto(mapCtx, in.GetApprovalState())
	// MISSING: DisplayName
	out.Repository = direct.LazyPtr(in.GetRepository())
	// MISSING: Files
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.SourceProvider = direct.LazyPtr(in.GetSourceProvider())
	out.DeploymentLevel = direct.Enum_FromProto(mapCtx, in.GetDeploymentLevel())
	out.RollbackSupport = direct.LazyPtr(in.GetRollbackSupport())
	return out
}
func BlueprintObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BlueprintObservedState) *pb.Blueprint {
	if in == nil {
		return nil
	}
	out := &pb.Blueprint{}
	// MISSING: Name
	out.RevisionId = direct.ValueOf(in.RevisionID)
	// MISSING: SourceBlueprint
	out.RevisionCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.RevisionCreateTime)
	out.ApprovalState = direct.Enum_ToProto[pb.Blueprint_ApprovalState](mapCtx, in.ApprovalState)
	// MISSING: DisplayName
	out.Repository = direct.ValueOf(in.Repository)
	// MISSING: Files
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.SourceProvider = direct.ValueOf(in.SourceProvider)
	out.DeploymentLevel = direct.Enum_ToProto[pb.DeploymentLevel](mapCtx, in.DeploymentLevel)
	out.RollbackSupport = direct.ValueOf(in.RollbackSupport)
	return out
}
func File_FromProto(mapCtx *direct.MapContext, in *pb.File) *krm.File {
	if in == nil {
		return nil
	}
	out := &krm.File{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.Content = direct.LazyPtr(in.GetContent())
	out.Deleted = direct.LazyPtr(in.GetDeleted())
	out.Editable = direct.LazyPtr(in.GetEditable())
	return out
}
func File_ToProto(mapCtx *direct.MapContext, in *krm.File) *pb.File {
	if in == nil {
		return nil
	}
	out := &pb.File{}
	out.Path = direct.ValueOf(in.Path)
	out.Content = direct.ValueOf(in.Content)
	out.Deleted = direct.ValueOf(in.Deleted)
	out.Editable = direct.ValueOf(in.Editable)
	return out
}
func TelcoautomationBlueprintObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Blueprint) *krm.TelcoautomationBlueprintObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TelcoautomationBlueprintObservedState{}
	// MISSING: Name
	// MISSING: RevisionID
	// MISSING: SourceBlueprint
	// MISSING: RevisionCreateTime
	// MISSING: ApprovalState
	// MISSING: DisplayName
	// MISSING: Repository
	// MISSING: Files
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourceProvider
	// MISSING: DeploymentLevel
	// MISSING: RollbackSupport
	return out
}
func TelcoautomationBlueprintObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TelcoautomationBlueprintObservedState) *pb.Blueprint {
	if in == nil {
		return nil
	}
	out := &pb.Blueprint{}
	// MISSING: Name
	// MISSING: RevisionID
	// MISSING: SourceBlueprint
	// MISSING: RevisionCreateTime
	// MISSING: ApprovalState
	// MISSING: DisplayName
	// MISSING: Repository
	// MISSING: Files
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourceProvider
	// MISSING: DeploymentLevel
	// MISSING: RollbackSupport
	return out
}
func TelcoautomationBlueprintSpec_FromProto(mapCtx *direct.MapContext, in *pb.Blueprint) *krm.TelcoautomationBlueprintSpec {
	if in == nil {
		return nil
	}
	out := &krm.TelcoautomationBlueprintSpec{}
	// MISSING: Name
	// MISSING: RevisionID
	// MISSING: SourceBlueprint
	// MISSING: RevisionCreateTime
	// MISSING: ApprovalState
	// MISSING: DisplayName
	// MISSING: Repository
	// MISSING: Files
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourceProvider
	// MISSING: DeploymentLevel
	// MISSING: RollbackSupport
	return out
}
func TelcoautomationBlueprintSpec_ToProto(mapCtx *direct.MapContext, in *krm.TelcoautomationBlueprintSpec) *pb.Blueprint {
	if in == nil {
		return nil
	}
	out := &pb.Blueprint{}
	// MISSING: Name
	// MISSING: RevisionID
	// MISSING: SourceBlueprint
	// MISSING: RevisionCreateTime
	// MISSING: ApprovalState
	// MISSING: DisplayName
	// MISSING: Repository
	// MISSING: Files
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourceProvider
	// MISSING: DeploymentLevel
	// MISSING: RollbackSupport
	return out
}
