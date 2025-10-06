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

// +generated:mapper
// krm.group: tags.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.resourcemanager.v3

package tags

import (
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tags/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TagBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TagBinding) *krm.TagBindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TagBindingObservedState{}
	// MISSING: Name
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
	// MISSING: Name
	// MISSING: Parent
	// MISSING: TagValue
	// MISSING: TagValueNamespacedName
	return out
}
func TagBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.TagBinding) *krm.TagBindingSpec {
	if in == nil {
		return nil
	}
	out := &krm.TagBindingSpec{}
	// MISSING: Name
	// MISSING: Parent
	// MISSING: TagValue
	// MISSING: TagValueNamespacedName
	return out
}
func TagBindingSpec_ToProto(mapCtx *direct.MapContext, in *krm.TagBindingSpec) *pb.TagBinding {
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
