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

package datacatalog

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb"
)
func PolicyTag_FromProto(mapCtx *direct.MapContext, in *pb.PolicyTag) *krm.PolicyTag {
	if in == nil {
		return nil
	}
	out := &krm.PolicyTag{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ParentPolicyTag = direct.LazyPtr(in.GetParentPolicyTag())
	// MISSING: ChildPolicyTags
	return out
}
func PolicyTag_ToProto(mapCtx *direct.MapContext, in *krm.PolicyTag) *pb.PolicyTag {
	if in == nil {
		return nil
	}
	out := &pb.PolicyTag{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.ParentPolicyTag = direct.ValueOf(in.ParentPolicyTag)
	// MISSING: ChildPolicyTags
	return out
}
func PolicyTagObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PolicyTag) *krm.PolicyTagObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PolicyTagObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ParentPolicyTag
	out.ChildPolicyTags = in.ChildPolicyTags
	return out
}
func PolicyTagObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PolicyTagObservedState) *pb.PolicyTag {
	if in == nil {
		return nil
	}
	out := &pb.PolicyTag{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ParentPolicyTag
	out.ChildPolicyTags = in.ChildPolicyTags
	return out
}
