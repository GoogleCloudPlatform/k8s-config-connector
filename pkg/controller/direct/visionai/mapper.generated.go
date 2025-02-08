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

package visionai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Collection_FromProto(mapCtx *direct.MapContext, in *pb.Collection) *krm.Collection {
	if in == nil {
		return nil
	}
	out := &krm.Collection{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func Collection_ToProto(mapCtx *direct.MapContext, in *krm.Collection) *pb.Collection {
	if in == nil {
		return nil
	}
	out := &pb.Collection{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	return out
}
func CollectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Collection) *krm.CollectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CollectionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func CollectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CollectionObservedState) *pb.Collection {
	if in == nil {
		return nil
	}
	out := &pb.Collection{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func VisionaiCollectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Collection) *krm.VisionaiCollectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiCollectionObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func VisionaiCollectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiCollectionObservedState) *pb.Collection {
	if in == nil {
		return nil
	}
	out := &pb.Collection{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func VisionaiCollectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Collection) *krm.VisionaiCollectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiCollectionSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
func VisionaiCollectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiCollectionSpec) *pb.Collection {
	if in == nil {
		return nil
	}
	out := &pb.Collection{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	return out
}
