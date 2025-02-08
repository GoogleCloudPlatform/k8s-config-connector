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

package resourcemanager

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Project_FromProto(mapCtx *direct.MapContext, in *pb.Project) *krm.Project {
	if in == nil {
		return nil
	}
	out := &krm.Project{}
	// MISSING: Name
	out.Parent = direct.LazyPtr(in.GetParent())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	// MISSING: State
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	out.Labels = in.Labels
	return out
}
func Project_ToProto(mapCtx *direct.MapContext, in *krm.Project) *pb.Project {
	if in == nil {
		return nil
	}
	out := &pb.Project{}
	// MISSING: Name
	out.Parent = direct.ValueOf(in.Parent)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	// MISSING: State
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	out.Labels = in.Labels
	return out
}
func ProjectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Project) *krm.ProjectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProjectObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Parent
	// MISSING: ProjectID
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: Labels
	return out
}
func ProjectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProjectObservedState) *pb.Project {
	if in == nil {
		return nil
	}
	out := &pb.Project{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Parent
	// MISSING: ProjectID
	out.State = direct.Enum_ToProto[pb.Project_State](mapCtx, in.State)
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: Labels
	return out
}
func ResourcemanagerProjectObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Project) *krm.ResourcemanagerProjectObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerProjectObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ProjectID
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	// MISSING: Labels
	return out
}
func ResourcemanagerProjectObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerProjectObservedState) *pb.Project {
	if in == nil {
		return nil
	}
	out := &pb.Project{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ProjectID
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	// MISSING: Labels
	return out
}
func ResourcemanagerProjectSpec_FromProto(mapCtx *direct.MapContext, in *pb.Project) *krm.ResourcemanagerProjectSpec {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerProjectSpec{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ProjectID
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	// MISSING: Labels
	return out
}
func ResourcemanagerProjectSpec_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerProjectSpec) *pb.Project {
	if in == nil {
		return nil
	}
	out := &pb.Project{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: ProjectID
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	// MISSING: Labels
	return out
}
