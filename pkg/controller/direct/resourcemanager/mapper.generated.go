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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1alpha1"
)
func Folder_FromProto(mapCtx *direct.MapContext, in *pb.Folder) *krm.Folder {
	if in == nil {
		return nil
	}
	out := &krm.Folder{}
	// MISSING: Name
	out.Parent = direct.LazyPtr(in.GetParent())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	return out
}
func Folder_ToProto(mapCtx *direct.MapContext, in *krm.Folder) *pb.Folder {
	if in == nil {
		return nil
	}
	out := &pb.Folder{}
	// MISSING: Name
	out.Parent = direct.ValueOf(in.Parent)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DeleteTime
	// MISSING: Etag
	return out
}
func FolderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Folder) *krm.FolderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FolderObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Parent
	// MISSING: DisplayName
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func FolderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FolderObservedState) *pb.Folder {
	if in == nil {
		return nil
	}
	out := &pb.Folder{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Parent
	// MISSING: DisplayName
	out.State = direct.Enum_ToProto[pb.Folder_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
