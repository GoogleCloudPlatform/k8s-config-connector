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
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ResourcemanagerTagBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TagBinding) *krm.ResourcemanagerTagBindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerTagBindingObservedState{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: TagValue
	// MISSING: TagValueNamespacedName
	return out
}
func ResourcemanagerTagBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerTagBindingObservedState) *pb.TagBinding {
	if in == nil {
		return nil
	}
	out := &pb.TagBinding{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: TagValue
	// MISSING: TagValueNamespacedName
	return out
}
func ResourcemanagerTagBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.TagBinding) *krm.ResourcemanagerTagBindingSpec {
	if in == nil {
		return nil
	}
	out := &krm.ResourcemanagerTagBindingSpec{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: TagValue
	// MISSING: TagValueNamespacedName
	return out
}
func ResourcemanagerTagBindingSpec_ToProto(mapCtx *direct.MapContext, in *krm.ResourcemanagerTagBindingSpec) *pb.TagBinding {
	if in == nil {
		return nil
	}
	out := &pb.TagBinding{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: TagValue
	// MISSING: TagValueNamespacedName
	return out
}
func TagBinding_FromProto(mapCtx *direct.MapContext, in *pb.TagBinding) *krm.TagBinding {
	if in == nil {
		return nil
	}
	out := &krm.TagBinding{}
	// MISSING: Name
	out.Parent = direct.LazyPtr(in.GetParent())
	out.TagValue = direct.LazyPtr(in.GetTagValue())
	out.TagValueNamespacedName = direct.LazyPtr(in.GetTagValueNamespacedName())
	return out
}
func TagBinding_ToProto(mapCtx *direct.MapContext, in *krm.TagBinding) *pb.TagBinding {
	if in == nil {
		return nil
	}
	out := &pb.TagBinding{}
	// MISSING: Name
	out.Parent = direct.ValueOf(in.Parent)
	out.TagValue = direct.ValueOf(in.TagValue)
	out.TagValueNamespacedName = direct.ValueOf(in.TagValueNamespacedName)
	return out
}
func TagBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TagBinding) *krm.TagBindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TagBindingObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Parent
	// MISSING: TagValue
	// MISSING: TagValueNamespacedName
	return out
}
func TagBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TagBindingObservedState) *pb.TagBinding {
	if in == nil {
		return nil
	}
	out := &pb.TagBinding{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Parent
	// MISSING: TagValue
	// MISSING: TagValueNamespacedName
	return out
}
