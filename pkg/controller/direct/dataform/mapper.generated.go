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

package dataform

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DataformRepositoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.DataformRepositoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataformRepositoryObservedState{}
	// MISSING: Name
	// MISSING: Labels
	return out
}
func DataformRepositoryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataformRepositoryObservedState) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	// MISSING: Name
	// MISSING: Labels
	return out
}
func DataformWorkspaceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workspace) *krm.DataformWorkspaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataformWorkspaceObservedState{}
	// MISSING: Name
	return out
}
func DataformWorkspaceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataformWorkspaceObservedState) *pb.Workspace {
	if in == nil {
		return nil
	}
	out := &pb.Workspace{}
	// MISSING: Name
	return out
}
func DataformWorkspaceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Workspace) *krm.DataformWorkspaceSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataformWorkspaceSpec{}
	// MISSING: Name
	return out
}
func DataformWorkspaceSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataformWorkspaceSpec) *pb.Workspace {
	if in == nil {
		return nil
	}
	out := &pb.Workspace{}
	// MISSING: Name
	return out
}
func Workspace_FromProto(mapCtx *direct.MapContext, in *pb.Workspace) *krm.Workspace {
	if in == nil {
		return nil
	}
	out := &krm.Workspace{}
	// MISSING: Name
	return out
}
func Workspace_ToProto(mapCtx *direct.MapContext, in *krm.Workspace) *pb.Workspace {
	if in == nil {
		return nil
	}
	out := &pb.Workspace{}
	// MISSING: Name
	return out
}
func WorkspaceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workspace) *krm.WorkspaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkspaceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func WorkspaceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkspaceObservedState) *pb.Workspace {
	if in == nil {
		return nil
	}
	out := &pb.Workspace{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
