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

package deploy

import (
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/deploy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CustomTargetSkaffoldActions_FromProto(mapCtx *direct.MapContext, in *pb.CustomTargetSkaffoldActions) *krm.CustomTargetSkaffoldActions {
	if in == nil {
		return nil
	}
	out := &krm.CustomTargetSkaffoldActions{}
	out.RenderAction = direct.LazyPtr(in.GetRenderAction())
	out.DeployAction = direct.LazyPtr(in.GetDeployAction())
	out.IncludeSkaffoldModules = direct.Slice_FromProto(mapCtx, in.IncludeSkaffoldModules, SkaffoldModules_FromProto)
	return out
}
func CustomTargetSkaffoldActions_ToProto(mapCtx *direct.MapContext, in *krm.CustomTargetSkaffoldActions) *pb.CustomTargetSkaffoldActions {
	if in == nil {
		return nil
	}
	out := &pb.CustomTargetSkaffoldActions{}
	out.RenderAction = direct.ValueOf(in.RenderAction)
	out.DeployAction = direct.ValueOf(in.DeployAction)
	out.IncludeSkaffoldModules = direct.Slice_ToProto(mapCtx, in.IncludeSkaffoldModules, SkaffoldModules_ToProto)
	return out
}
func CustomTargetType_FromProto(mapCtx *direct.MapContext, in *pb.CustomTargetType) *krm.CustomTargetType {
	if in == nil {
		return nil
	}
	out := &krm.CustomTargetType{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CustomTargetTypeID
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.CustomActions = CustomTargetSkaffoldActions_FromProto(mapCtx, in.GetCustomActions())
	return out
}
func CustomTargetType_ToProto(mapCtx *direct.MapContext, in *krm.CustomTargetType) *pb.CustomTargetType {
	if in == nil {
		return nil
	}
	out := &pb.CustomTargetType{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CustomTargetTypeID
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	out.Annotations = in.Annotations
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Etag = direct.ValueOf(in.Etag)
	if oneof := CustomTargetSkaffoldActions_ToProto(mapCtx, in.CustomActions); oneof != nil {
		out.Definition = &pb.CustomTargetType_CustomActions{CustomActions: oneof}
	}
	return out
}
func CustomTargetTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomTargetType) *krm.CustomTargetTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomTargetTypeObservedState{}
	// MISSING: Name
	out.CustomTargetTypeID = direct.LazyPtr(in.GetCustomTargetTypeId())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	// MISSING: CustomActions
	return out
}
func CustomTargetTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomTargetTypeObservedState) *pb.CustomTargetType {
	if in == nil {
		return nil
	}
	out := &pb.CustomTargetType{}
	// MISSING: Name
	out.CustomTargetTypeId = direct.ValueOf(in.CustomTargetTypeID)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	// MISSING: CustomActions
	return out
}
func DeployCustomTargetTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomTargetType) *krm.DeployCustomTargetTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployCustomTargetTypeObservedState{}
	// MISSING: Name
	// MISSING: CustomTargetTypeID
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: CustomActions
	return out
}
func DeployCustomTargetTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployCustomTargetTypeObservedState) *pb.CustomTargetType {
	if in == nil {
		return nil
	}
	out := &pb.CustomTargetType{}
	// MISSING: Name
	// MISSING: CustomTargetTypeID
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: CustomActions
	return out
}
func DeployCustomTargetTypeSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomTargetType) *krm.DeployCustomTargetTypeSpec {
	if in == nil {
		return nil
	}
	out := &krm.DeployCustomTargetTypeSpec{}
	// MISSING: Name
	// MISSING: CustomTargetTypeID
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: CustomActions
	return out
}
func DeployCustomTargetTypeSpec_ToProto(mapCtx *direct.MapContext, in *krm.DeployCustomTargetTypeSpec) *pb.CustomTargetType {
	if in == nil {
		return nil
	}
	out := &pb.CustomTargetType{}
	// MISSING: Name
	// MISSING: CustomTargetTypeID
	// MISSING: Uid
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Etag
	// MISSING: CustomActions
	return out
}
func SkaffoldModules_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldModules) *krm.SkaffoldModules {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldModules{}
	out.Configs = in.Configs
	out.Git = SkaffoldModules_SkaffoldGitSource_FromProto(mapCtx, in.GetGit())
	out.GoogleCloudStorage = SkaffoldModules_SkaffoldGCSSource_FromProto(mapCtx, in.GetGoogleCloudStorage())
	out.GoogleCloudBuildRepo = SkaffoldModules_SkaffoldGcbRepoSource_FromProto(mapCtx, in.GetGoogleCloudBuildRepo())
	return out
}
func SkaffoldModules_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldModules) *pb.SkaffoldModules {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldModules{}
	out.Configs = in.Configs
	if oneof := SkaffoldModules_SkaffoldGitSource_ToProto(mapCtx, in.Git); oneof != nil {
		out.Source = &pb.SkaffoldModules_Git{Git: oneof}
	}
	if oneof := SkaffoldModules_SkaffoldGCSSource_ToProto(mapCtx, in.GoogleCloudStorage); oneof != nil {
		out.Source = &pb.SkaffoldModules_GoogleCloudStorage{GoogleCloudStorage: oneof}
	}
	if oneof := SkaffoldModules_SkaffoldGcbRepoSource_ToProto(mapCtx, in.GoogleCloudBuildRepo); oneof != nil {
		out.Source = &pb.SkaffoldModules_GoogleCloudBuildRepo{GoogleCloudBuildRepo: oneof}
	}
	return out
}
func SkaffoldModules_SkaffoldGCSSource_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldModules_SkaffoldGCSSource) *krm.SkaffoldModules_SkaffoldGCSSource {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldModules_SkaffoldGCSSource{}
	out.Source = direct.LazyPtr(in.GetSource())
	out.Path = direct.LazyPtr(in.GetPath())
	return out
}
func SkaffoldModules_SkaffoldGCSSource_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldModules_SkaffoldGCSSource) *pb.SkaffoldModules_SkaffoldGCSSource {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldModules_SkaffoldGCSSource{}
	out.Source = direct.ValueOf(in.Source)
	out.Path = direct.ValueOf(in.Path)
	return out
}
func SkaffoldModules_SkaffoldGcbRepoSource_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldModules_SkaffoldGCBRepoSource) *krm.SkaffoldModules_SkaffoldGcbRepoSource {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldModules_SkaffoldGcbRepoSource{}
	out.Repository = direct.LazyPtr(in.GetRepository())
	out.Path = direct.LazyPtr(in.GetPath())
	out.Ref = direct.LazyPtr(in.GetRef())
	return out
}
func SkaffoldModules_SkaffoldGcbRepoSource_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldModules_SkaffoldGcbRepoSource) *pb.SkaffoldModules_SkaffoldGCBRepoSource {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldModules_SkaffoldGCBRepoSource{}
	out.Repository = direct.ValueOf(in.Repository)
	out.Path = direct.ValueOf(in.Path)
	out.Ref = direct.ValueOf(in.Ref)
	return out
}
func SkaffoldModules_SkaffoldGitSource_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldModules_SkaffoldGitSource) *krm.SkaffoldModules_SkaffoldGitSource {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldModules_SkaffoldGitSource{}
	out.Repo = direct.LazyPtr(in.GetRepo())
	out.Path = direct.LazyPtr(in.GetPath())
	out.Ref = direct.LazyPtr(in.GetRef())
	return out
}
func SkaffoldModules_SkaffoldGitSource_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldModules_SkaffoldGitSource) *pb.SkaffoldModules_SkaffoldGitSource {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldModules_SkaffoldGitSource{}
	out.Repo = direct.ValueOf(in.Repo)
	out.Path = direct.ValueOf(in.Path)
	out.Ref = direct.ValueOf(in.Ref)
	return out
}
