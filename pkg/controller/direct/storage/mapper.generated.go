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

package storage

import (
	pb "cloud.google.com/go/storage/control/apiv2/controlpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func StorageManagedFolderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagedFolder) *krm.StorageManagedFolderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StorageManagedFolderObservedState{}
	// MISSING: Name
	out.Metageneration = direct.LazyPtr(in.GetMetageneration())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func StorageManagedFolderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StorageManagedFolderObservedState) *pb.ManagedFolder {
	if in == nil {
		return nil
	}
	out := &pb.ManagedFolder{}
	// MISSING: Name
	out.Metageneration = direct.ValueOf(in.Metageneration)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func StorageManagedFolderSpec_FromProto(mapCtx *direct.MapContext, in *pb.ManagedFolder) *krm.StorageManagedFolderSpec {
	if in == nil {
		return nil
	}
	out := &krm.StorageManagedFolderSpec{}
	// MISSING: Name
	return out
}
func StorageManagedFolderSpec_ToProto(mapCtx *direct.MapContext, in *krm.StorageManagedFolderSpec) *pb.ManagedFolder {
	if in == nil {
		return nil
	}
	out := &pb.ManagedFolder{}
	// MISSING: Name
	return out
}
