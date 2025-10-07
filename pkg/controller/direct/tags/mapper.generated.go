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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tags/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TagsTagBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TagBinding) *krm.TagsTagBindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TagsTagBindingObservedState{}
	// MISSING: Name
	// MISSING: TagValueNamespacedName
	return out
}
func TagsTagBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TagsTagBindingObservedState) *pb.TagBinding {
	if in == nil {
		return nil
	}
	out := &pb.TagBinding{}
	// MISSING: Name
	// MISSING: TagValueNamespacedName
	return out
}
func TagsTagBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.TagBinding) *krm.TagsTagBindingSpec {
	if in == nil {
		return nil
	}
	out := &krm.TagsTagBindingSpec{}
	// MISSING: Name
	if in.GetParent() != "" {
		out.ParentRef = &krm.ParentRef{External: in.GetParent()}
	}
	if in.GetTagValue() != "" {
		out.TagValueRef = &refsv1beta1.TagValueRef{External: in.GetTagValue()}
	}
	// MISSING: TagValueNamespacedName
	return out
}
func TagsTagBindingSpec_ToProto(mapCtx *direct.MapContext, in *krm.TagsTagBindingSpec) *pb.TagBinding {
	if in == nil {
		return nil
	}
	out := &pb.TagBinding{}
	// MISSING: Name
	if in.ParentRef != nil {
		out.Parent = in.ParentRef.External
	}
	if in.TagValueRef != nil {
		out.TagValue = in.TagValueRef.External
	}
	// MISSING: TagValueNamespacedName
	return out
}
