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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/telcoautomation/apiv1/telcoautomationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/telcoautomation/v1alpha1"
)
func Deployment_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.Deployment {
	if in == nil {
		return nil
	}
	out := &krm.Deployment{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: RevisionID
	out.SourceBlueprintRevision = direct.LazyPtr(in.GetSourceBlueprintRevision())
	// MISSING: RevisionCreateTime
	// MISSING: State
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Repository
	out.Files = direct.Slice_FromProto(mapCtx, in.Files, File_FromProto)
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourceProvider
	out.WorkloadCluster = direct.LazyPtr(in.GetWorkloadCluster())
	// MISSING: DeploymentLevel
	// MISSING: RollbackSupport
	return out
}
func Deployment_ToProto(mapCtx *direct.MapContext, in *krm.Deployment) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: RevisionID
	out.SourceBlueprintRevision = direct.ValueOf(in.SourceBlueprintRevision)
	// MISSING: RevisionCreateTime
	// MISSING: State
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Repository
	out.Files = direct.Slice_ToProto(mapCtx, in.Files, File_ToProto)
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourceProvider
	out.WorkloadCluster = direct.ValueOf(in.WorkloadCluster)
	// MISSING: DeploymentLevel
	// MISSING: RollbackSupport
	return out
}
func DeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.DeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeploymentObservedState{}
	// MISSING: Name
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	// MISSING: SourceBlueprintRevision
	out.RevisionCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevisionCreateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: DisplayName
	out.Repository = direct.LazyPtr(in.GetRepository())
	// MISSING: Files
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.SourceProvider = direct.LazyPtr(in.GetSourceProvider())
	// MISSING: WorkloadCluster
	out.DeploymentLevel = direct.Enum_FromProto(mapCtx, in.GetDeploymentLevel())
	out.RollbackSupport = direct.LazyPtr(in.GetRollbackSupport())
	return out
}
func DeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeploymentObservedState) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	// MISSING: Name
	out.RevisionId = direct.ValueOf(in.RevisionID)
	// MISSING: SourceBlueprintRevision
	out.RevisionCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.RevisionCreateTime)
	out.State = direct.Enum_ToProto[pb.Deployment_State](mapCtx, in.State)
	// MISSING: DisplayName
	out.Repository = direct.ValueOf(in.Repository)
	// MISSING: Files
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.SourceProvider = direct.ValueOf(in.SourceProvider)
	// MISSING: WorkloadCluster
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
func TelcoautomationDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.TelcoautomationDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TelcoautomationDeploymentObservedState{}
	// MISSING: Name
	// MISSING: RevisionID
	// MISSING: SourceBlueprintRevision
	// MISSING: RevisionCreateTime
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Repository
	// MISSING: Files
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourceProvider
	// MISSING: WorkloadCluster
	// MISSING: DeploymentLevel
	// MISSING: RollbackSupport
	return out
}
func TelcoautomationDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TelcoautomationDeploymentObservedState) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	// MISSING: Name
	// MISSING: RevisionID
	// MISSING: SourceBlueprintRevision
	// MISSING: RevisionCreateTime
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Repository
	// MISSING: Files
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourceProvider
	// MISSING: WorkloadCluster
	// MISSING: DeploymentLevel
	// MISSING: RollbackSupport
	return out
}
func TelcoautomationDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.TelcoautomationDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.TelcoautomationDeploymentSpec{}
	// MISSING: Name
	// MISSING: RevisionID
	// MISSING: SourceBlueprintRevision
	// MISSING: RevisionCreateTime
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Repository
	// MISSING: Files
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourceProvider
	// MISSING: WorkloadCluster
	// MISSING: DeploymentLevel
	// MISSING: RollbackSupport
	return out
}
func TelcoautomationDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.TelcoautomationDeploymentSpec) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	// MISSING: Name
	// MISSING: RevisionID
	// MISSING: SourceBlueprintRevision
	// MISSING: RevisionCreateTime
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Repository
	// MISSING: Files
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: SourceProvider
	// MISSING: WorkloadCluster
	// MISSING: DeploymentLevel
	// MISSING: RollbackSupport
	return out
}
